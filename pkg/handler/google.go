package handler

import (
	"github.com/gotemplates/core/exchange"
	"github.com/gotemplates/core/runtime"
	"github.com/gotemplates/example-host/pkg/google"
	"net/http"
)

func GoogleHandler(w http.ResponseWriter, r *http.Request) {
	buf, status := google.Search[runtime.LogError](runtime.ContextWithRequest(r), r.URL)
	exchange.WriteResponse(w, buf, status)
}
