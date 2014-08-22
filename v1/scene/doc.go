// Tideland Go Application Support - Scene
//
// Copyright (C) 2014 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

// Scene provides a shared access to common used data in a larger context.
//
// By definition a scene is a sequence of continuous action in a play,
// movie, opera, or book. Applications do know these kind of scenes too.
// Here aspects of the action have to passed between the actors, very
// often they are interwoven and depending.
//
// This also happens in software, especially in concurrent software,
// where different parts have to exchange data in a simple and secure
// way. Here the scene package helps. Beside a simple atomic way to
// store and retrieve information together with optional cleanup
// functions it handles inactivity and absolute timeouts.
//
// A scene without timeouts is started with
//
//    scn := scene.Start()
//
// Now props can be stored, fetched, and disposed.
//
//    err := scn.Store("foo", myFoo)
//    foo, err := scn.Fetch("foo")
//    foo, err := scn.Dispose("foo")
//
// It's also possible to cleanup if a prop is disposed or the whole
// scene is stopped or aborted.
//
//    myCleanup := func(key string, prop interface{}) error {
//        // Cleanup, e.g. return the prop into a pool
//        // or close handles.
//        ...
//        return nil
//    }
//    err := scn.StoreClean("foo", myFoo, myCleanup)
//
// The cleanup is called individually per prop when disposing it, when the
// scene ends due to a timeout, or when it is stopped with
//
//    err := scn.Stop()
//
// or
//
//    scn.Abort(myError)
//
// A scene knows two different timeouts. The first is the time of inactivity,
// the second is the absolute maximum time of a scene.
//
//    inactivityTimeout := 5 * time.Minutes
//    absoluteTimeout := 60 * time.Minutes
//    scn := scene.StartLimited(inactivityTimeout, absoluteTimeout)
//
// Now the scene is stopped after 5 minutes without any access or at the
// latest 60 minutes after the start. Both value may be zero if not needed.
// So scene.StartLimited(0, 0) is the same as scene.Start().
package scene

// EOF
