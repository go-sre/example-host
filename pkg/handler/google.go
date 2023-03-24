package handler

import (
	"github.com/go-sre/core/exchange"
	"github.com/go-sre/core/runtime"
	"github.com/go-sre/example-host/pkg/google"
	"net/http"
)

func GoogleHandler(w http.ResponseWriter, r *http.Request) {
	buf, status := google.Search[runtime.LogError](runtime.ContextWithRequest(r), r.URL)
	exchange.WriteResponse(w, buf, status)
}
