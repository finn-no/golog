# Golog
Golog is a wrapper around logrus which conforms to Finn's logging requirements.

## Getting started
Golog can be dropped in place of Logrus and used in a very similar manner:
```
package main

import log "github.com/finn-no/golog"

func main() {
    log.Info("hello, golog!")
}
```

## Using FIAAS_ENVIRONMENT
Golog looks for the environment variable `FIAAS_ENVIRONMENT`. If it is non
empty, Golog will conform to all of Finn's logging expectations. If it is empty,
a standard Logrus logger will be used.

## Setting level
The logging level can be set much in the same way as in Logrus:
```
log.SetLevel("debug")
```
`SetLevel` returns an error which can be checked like so:
```
if err := log.SetLevel("debugg"); err != nil {
    log.Fatalf("setting logging level: %v", err)
}
```

## Using the logger
When you call `golog.Info("some message")`, the message's calling file will be
from the `golog` package itself, rather than the actual calling file. This
behaviour can be mitigated by using the logger directly:
```
logger := log.Logger()
logger.Info("hello, golog!")
```
This is the same `logrus.Logger` used by golog internally.
