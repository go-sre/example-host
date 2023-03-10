package handler

import (
	"github.com/gotemplates/core/exchange"
	"github.com/gotemplates/core/runtime"
	"github.com/gotemplates/example-host/pkg/facebook"
	"net/http"
)

func FacebookHandler(w http.ResponseWriter, r *http.Request) {
	buf, status := facebook.Get[runtime.LogError](runtime.ContextWithRequest(r))
	exchange.WriteResponse(w, buf, status)
}
