# Tideland Go Application Support

## Description

The *Tideland Go Application Support* (GOAS) is a collection of smaller 
helpful packages for multiple purposes. They are those little helpers we
always need. See their descriptions below.

## Installation

```
go get github.com/tideland/goas/v3/errors
go get github.com/tideland/goas/v2/identifier
go get github.com/tideland/goas/v2/logger
go get github.com/tideland/goas/v2/loop
go get github.com/tideland/goas/v2/monitoring
go get github.com/tideland/goas/v2/timex
go get github.com/tideland/goas/v1/version
```

## Usage

### Errors

Typical errors in Go are often created using `errors.New()` or `fmt.Errorf()`. Those
errors only contain a string as information. When trying to differentiate between
errors or to carry helpful payload own types are needed.

The errors package allows to easily create formatted errors with `errors.New()` or 
`errors.Annotate()` like with the `fmt.Errorf()` function, but also with an error code. 
This easily can be tested with `errors.IsError(err, code)`. 

[![GoDoc](https://godoc.org/github.com/tideland/goas/v3/errors?status.svg)](https://godoc.org/github.com/tideland/goas/v3/errors)

### Identifier

The identifier packge provides different ways to produce identifiers like UUIDs. The
UUID generation follows version 4 (based on random numbers), other identifier types are
based on passed data or types. Here the individual parts are harmonized and concatenated
by the passed seperators. It is the users responsibility to check if the identifier is
unique in its context.

[![GoDoc](https://godoc.org/github.com/tideland/goas/v2/identifier?status.svg)](https://godoc.org/github.com/tideland/goas/v2/identifier)

### Logger

This package helps to log with different levels and on different backends like on an
io.Writer or the Go logging. Own backends can be defined based on a simple interface.
Setting the level controls what will be logged. Debug and critical logging also print
filename, function name and line number.

[![GoDoc](https://godoc.org/github.com/tideland/goas/v2/logger?status.svg)](https://godoc.org/github.com/tideland/goas/v2/logger)

### Loop

A typical Go idiom for concurrent applications is running a loop in the background doing
a select on one or more channels. Stopping those loops or getting aware of internal errors
requires extra efforts. The loop package helps to manage this kind of goroutines.

The loop function, or method, looks like

```
func (f *Foo) backendLoop(l loop.Loop) error {
        select {
        case <-l.ShallStop():
                return nil
        case bar := <-f.barChan:
                ...
                if err != nil {
                        return err
                }
        case baz := <-f.bazChan:
                ...
                if err != nil {
                        return err
                }
        }
}
```

The loop is started with `l := loop.Go(f.backendLoop)`. Now the loop can be stopped with
`l.Stop()` which will return `nil` in case of no error or the error that has been returned
by the loop during message processing. The loop also can be killed with `l.Kill(err)` or
the status retrieved with `status, err := l.Error()`. The status is one of

- `loop.Running`,
- `loop.Stopping`, or 
- `loop.Stopped`.

Another variant is `loop.GoRecoverable(f.backendLoop, f.recoverFunc)`. Here a loop error
or the value of a recovering after a panic are passed to the recover function. It then
can decide if the loop shall be restarted or really terminated.

[![GoDoc](https://godoc.org/github.com/tideland/goas/v2/loop?status.svg)](https://godoc.org/github.com/tideland/goas/v2/loop)

### Monitoring

The monitoring package supports three kinds of system monitoring. They are helpful to
understand what's happening inside a system during runtime. So execution times can be
measured and analyzed, stay-set indicators integrated and dynamic control value retrieval
provided.

```
etm := monitoring.BeginMeasuring("foo")
defer etm.EndMeasuring()

monitoring.SetVariable("bar", 4711)
monitoring.IncrVariable("bar")
monitoring.DecrVariable("bar")

monitoring.Register("baz", func() (string, error) { ... })
```

[![GoDoc](https://godoc.org/github.com/tideland/goas/v2/monitoring?status.svg)](https://godoc.org/github.com/tideland/goas/v2/monitoring)

### Timex

The timex package supports the work with dates and times. Additionally it provides a
simple crontab.

[![GoDoc](https://godoc.org/github.com/tideland/goas/v2/timex?status.svg)](https://godoc.org/github.com/tideland/goas/v2/timex)

### Version

The version package allows other packages to provide information about their version
and to compare it to other versions. It follows the idea of [semantic versioning](http://semver.org). 
Here given a version number MAJOR.MINOR.PATCH, increment the:

- MAJOR version when you make incompatible API changes,
- MINOR version when you add functionality in a backwards-compatible manner, and
- PATCH version when you make backwards-compatible bug fixes.

Additional labels for pre-release and build metadata are available as extensions to the 
MAJOR.MINOR.PATCH format.

[![GoDoc](https://godoc.org/github.com/tideland/goas/v1/version?status.svg)](https://godoc.org/github.com/tideland/goas/v1/version)

And now have fun. ;)

## Contributors

- Frank Mueller - <mue@tideland.biz>
- Benedikt Lang - <github@benediktlang.de>

## License

*Tideland Go Application Support* is distributed under the terms of the BSD 3-Clause license.
