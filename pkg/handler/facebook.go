package handler

import (
	"github.com/go-sre/core/exchange"
	"github.com/go-sre/core/runtime"
	"github.com/go-sre/example-host/pkg/facebook"
	"net/http"
)

func FacebookHandler(w http.ResponseWriter, r *http.Request) {
	buf, status := facebook.Get[runtime.LogError](runtime.ContextWithRequest(r))
	exchange.WriteResponse(w, buf, status)
}
