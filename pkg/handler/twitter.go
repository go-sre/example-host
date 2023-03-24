package handler

import (
	"github.com/go-sre/core/exchange"
	"github.com/go-sre/core/runtime"
	"github.com/go-sre/example-host/pkg/twitter"
	"net/http"
)

func TwitterHandler(w http.ResponseWriter, r *http.Request) {
	buf, status := twitter.Get[runtime.LogError](runtime.ContextWithRequest(r))
	exchange.WriteResponse(w, buf, status)
}
