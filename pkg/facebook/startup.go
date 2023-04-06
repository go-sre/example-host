package facebook

import (
	"github.com/go-sre/core/runtime"
	"github.com/go-sre/host/messaging"
	"reflect"
	"sync/atomic"
	"time"
)

type pkg struct{}

var (
	PkgUrl  = pkgPath
	pkgPath = reflect.TypeOf(any(pkg{})).PkgPath()
	c       = make(chan messaging.Message, 1)
	started int64
)

func IsStarted() bool { return atomic.LoadInt64(&started) != 0 }

func init() {
	messaging.RegisterResource(PkgUrl, c)
	go receive()
}

var messageHandler messaging.MessageHandler = func(msg messaging.Message) {
	start := time.Now()
	switch msg.Event {
	case messaging.StartupEvent:
		atomic.StoreInt64(&started, 1)
		messaging.ReplyTo(msg, runtime.NewStatusOK().SetDuration(time.Since(start)))
	case messaging.ShutdownEvent:
	}
}

func receive() {
	for {
		select {
		case msg, open := <-c:
			if !open {
				return
			}
			go messageHandler(msg)
		default:
		}
	}
}
