package handler

import (
	"github.com/gotemplates/core/exchange"
	"github.com/gotemplates/core/runtime"
	"github.com/gotemplates/example-host/pkg/twitter"
	"net/http"
)

func TwitterHandler(w http.ResponseWriter, r *http.Request) {
	buf, status := twitter.Get[runtime.LogError](runtime.ContextWithRequest(r))
	exchange.WriteResponse(w, buf, status)
}
