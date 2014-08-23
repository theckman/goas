// Tideland Go Application Support - Scene
//
// Copyright (C) 2014 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package scene

//--------------------
// IMPORTS
//--------------------

import (
	"time"

	"github.com/tideland/goas/v2/loop"
	"github.com/tideland/goas/v3/errors"
)

//--------------------
// SCENE
//--------------------

// CleanupFunc is a function for the cleanup of props after
// a scene ended.
type CleanupFunc func(key string, prop interface{}) error

// box contains a prop and a possible cleanup function.
type box struct {
	key     string
	prop    interface{}
	cleanup CleanupFunc
}

// subscription contains a topic and a signal channel.
type subscription struct {
	topic      string
	signalChan chan struct{}
}

const (
	storeProp = iota
	fetchProp
	disposeProp
	signal
	subscribe
)

// envelope contains information transfered between client and scene.
type envelope struct {
	kind         int
	box          *box
	subscription *subscription
	err          error
	respChan     chan *envelope
}

// Scene is the access point to one scene. It has to be created once
// for a continuous flow of operations and then passed between all
// functions and goroutine which are actors of the scene.
type Scene interface {
	// Stop tells the scene to end and waits until it is done.
	Stop() error

	// Abort tells the scene to end due to the passed error.
	// Here only the first error will be stored for later evaluation.
	Abort(err error)

	// Wait blocks the caller until the scene ended and returns a
	// possible error or nil.
	Wait() error

	// Store stores a prop with a given key. The key must not exist.
	Store(key string, prop interface{}) error

	// StoreClean stores a prop with a given key and a cleanup
	// function called when a scene ends. The key must not exist.
	StoreClean(key string, prop interface{}, cleanup CleanupFunc) error

	// Fetch retrieves a prop.
	Fetch(key string) (interface{}, error)

	// Dispose retrieves a prop and deletes it from the store.
	Dispose(key string) (interface{}, error)

	// Signal allows to signal a topic to interested listeners.
	Signal(topic string) error

	// WaitSignal waits until the passed topic has been signalled.
	WaitSignal(topic string) error

	// WaitSignalLimited waits until the passed topic has been signalled
	// or the timeout happened.
	WaitSignalLimited(topic string, timeout time.Duration) error
}

// scene implements Scene.
type scene struct {
	props         map[string]*box
	subscriptions map[string][]chan struct{}
	inactivity    time.Duration
	absolute      time.Duration
	commandChan   chan *envelope
	backend       loop.Loop
}

// Start creates and runs a new scene.
func Start() Scene {
	return StartLimited(0, 0)
}

// StartLimited creates and runs a new scene with an inactivity
// and an absolute timeout. They may be zero.
func StartLimited(inactivity, absolute time.Duration) Scene {
	s := &scene{
		props:         make(map[string]*box),
		subscriptions: make(map[string][]chan struct{}),
		inactivity:    inactivity,
		absolute:      absolute,
		commandChan:   make(chan *envelope, 1),
	}
	s.backend = loop.Go(s.backendLoop)
	return s
}

// Stop is specified on the Scene interface.
func (s *scene) Stop() error {
	return s.backend.Stop()
}

// Abort is specified on the Scene interface.
func (s *scene) Abort(err error) {
	s.backend.Kill(err)
}

// Wait is specified on the Scene interface.
func (s *scene) Wait() error {
	return s.backend.Wait()
}

// Store is specified on the Scene interface.
func (s *scene) Store(key string, prop interface{}) error {
	return s.StoreClean(key, prop, nil)
}

// StoreClean is specified on the Scene interface.
func (s *scene) StoreClean(key string, prop interface{}, cleanup CleanupFunc) error {
	command := &envelope{
		kind: storeProp,
		box: &box{
			key:     key,
			prop:    prop,
			cleanup: cleanup,
		},
		respChan: make(chan *envelope, 1),
	}
	_, err := s.command(command)
	return err
}

// Fetch is specified on the Scene interface.
func (s *scene) Fetch(key string) (interface{}, error) {
	command := &envelope{
		kind: fetchProp,
		box: &box{
			key: key,
		},
		respChan: make(chan *envelope, 1),
	}
	resp, err := s.command(command)
	if err != nil {
		return nil, err
	}
	return resp.box.prop, nil
}

// Dispose is specified on the Scene interface.
func (s *scene) Dispose(key string) (interface{}, error) {
	command := &envelope{
		kind: disposeProp,
		box: &box{
			key: key,
		},
		respChan: make(chan *envelope, 1),
	}
	resp, err := s.command(command)
	if err != nil {
		return nil, err
	}
	return resp.box.prop, nil
}

// Signal is specified on the Scene interface.
func (s *scene) Signal(topic string) error {
	command := &envelope{
		kind: signal,
		subscription: &subscription{
			topic: topic,
		},
		respChan: make(chan *envelope, 1),
	}
	_, err := s.command(command)
	return err
}

// WaitSignal is specified on the Scene interface.
func (s *scene) WaitSignal(topic string) error {
	return s.WaitSignalLimited(topic, 0)
}

// WaitSignalLimited is specified on the Scene interface.
func (s *scene) WaitSignalLimited(topic string, timeout time.Duration) error {
	// Add signal channel.
	command := &envelope{
		kind: subscribe,
		subscription: &subscription{
			topic:      topic,
			signalChan: make(chan struct{}, 1),
		},
		respChan: make(chan *envelope, 1),
	}
	_, err := s.command(command)
	if err != nil {
		return err
	}
	// Wait for signal.
	var timeoutChan <-chan time.Time
	if timeout > 0 {
		timeoutChan = time.After(timeout)
	}
	select {
	case <-s.backend.IsStopping():
		err = s.Wait()
		if err == nil {
			err = errors.New(ErrSceneEnded, errorMessages)
		}
		return err
	case <-command.subscription.signalChan:
		return nil
	case <-timeoutChan:
		return errors.New(ErrWaitedTooLong, errorMessages, topic)
	}
}

// command sends a command envelope to the backend and
// waits for the response.
func (s *scene) command(command *envelope) (*envelope, error) {
	select {
	case s.commandChan <- command:
	case <-s.backend.IsStopping():
		err := s.Wait()
		if err == nil {
			err = errors.New(ErrSceneEnded, errorMessages)
		}
		return nil, err
	}
	select {
	case <-s.backend.IsStopping():
		err := s.Wait()
		if err == nil {
			err = errors.New(ErrSceneEnded, errorMessages)
		}
		return nil, err
	case resp := <-command.respChan:
		if resp.err != nil {
			return nil, resp.err
		}
		return resp, nil
	}
}

// backendLoop runs the backend loop of the scene.
func (s *scene) backendLoop(l loop.Loop) (err error) {
	// Defer cleanup.
	defer func() {
		cerr := s.cleanupAllProps()
		if err == nil {
			err = cerr
		}
	}()
	// Init timers.
	var watchdog <-chan time.Time
	var clapperboard <-chan time.Time
	if s.absolute > 0 {
		clapperboard = time.After(s.absolute)
	}
	// Run loop.
	for {
		if s.inactivity > 0 {
			watchdog = time.After(s.inactivity)
		}
		select {
		case <-l.ShallStop():
			return nil
		case timeout := <-watchdog:
			return errors.New(ErrTimeout, errorMessages, "inactivity", timeout)
		case timeout := <-clapperboard:
			return errors.New(ErrTimeout, errorMessages, "absolute", timeout)
		case command := <-s.commandChan:
			s.processCommand(command)
		}
	}
}

// processCommand processes the sent commands.
func (s *scene) processCommand(command *envelope) {
	switch command.kind {
	case storeProp:
		// Add a new prop.
		_, ok := s.props[command.box.key]
		if ok {
			command.err = errors.New(ErrPropAlreadyExist, errorMessages, command.box.key)
		} else {
			s.props[command.box.key] = command.box
		}
	case fetchProp:
		// Retrieve a prop.
		box, ok := s.props[command.box.key]
		if !ok {
			command.err = errors.New(ErrPropNotFound, errorMessages, command.box.key)
		} else {
			command.box = box
		}
	case disposeProp:
		// Remove a prop.
		box, ok := s.props[command.box.key]
		if !ok {
			command.err = errors.New(ErrPropNotFound, errorMessages, command.box.key)
		} else {
			delete(s.props, command.box.key)
			command.box = box
			if box.cleanup != nil {
				cerr := box.cleanup(box.key, box.prop)
				if cerr != nil {
					command.err = errors.Annotate(cerr, ErrCleanupFailed, errorMessages, box.key)
				}
			}
		}
	case signal:
		// Signal a topic.
		subscribers, ok := s.subscriptions[command.subscription.topic]
		if !ok {
			command.err = errors.New(ErrNoSubscriber, errorMessages, command.subscription.topic)
		} else {
			delete(s.subscriptions, command.subscription.topic)
			for _, subscriber := range subscribers {
				subscriber <- struct{}{}
			}
		}
	case subscribe:
		// Add a subscriber channel.
		subscribers, ok := s.subscriptions[command.subscription.topic]
		if !ok {
			s.subscriptions[command.subscription.topic] = []chan struct{}{command.subscription.signalChan}
		} else {
			s.subscriptions[command.subscription.topic] = append(subscribers, command.subscription.signalChan)
		}
	default:
		panic("illegal command")
	}
	// Return the changed command as response.
	command.respChan <- command
}

// cleanupAllProps cleans all props.
func (s *scene) cleanupAllProps() error {
	for _, box := range s.props {
		if box.cleanup != nil {
			err := box.cleanup(box.key, box.prop)
			if err != nil {
				return errors.Annotate(err, ErrCleanupFailed, errorMessages, box.key)
			}
		}
	}
	return nil
}

// EOF
