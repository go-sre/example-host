package host

import (
	"github.com/gotemplates/example-host/pkg/handler"
	"net/http"
	"net/http/pprof"
)

const (
	googlePattern         = "/google"
	twitterPattern        = "/twitter"
	healthLivenessPattern = "/health/liveness"
	facebookPattern       = "/facebook"

	indexPattern   = "/debug/pprof/"
	cmdLinePattern = "/debug/pprof/cmdline"
	profilePattern = "/debug/pprof/profile" // ?seconds=30
	symbolPattern  = "/debug/pprof/symbol"
	tracePattern   = "/debug/pprof/trace"

	IndexRouteName   = "pprof-index"
	CmdLineRouteName = "pprof-cmdline"
	ProfileRouteName = "pprof-profile"
	SymbolRouteName  = "pprof-symbol"
	TraceRouteName   = "pprof-trace"
)

func initMux(r *http.ServeMux) {
	addRoutes(r)
	r.Handle(googlePattern, http.HandlerFunc(handler.GoogleHandler))
	r.Handle(twitterPattern, http.HandlerFunc(handler.TwitterHandler))
	r.Handle(facebookPattern, http.HandlerFunc(handler.FacebookHandler))
	r.Handle(healthLivenessPattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("up"))
	}))
}

func addRoutes(r *http.ServeMux) {
	r.Handle(indexPattern, http.HandlerFunc(pprof.Index))
	r.Handle(cmdLinePattern, http.HandlerFunc(pprof.Cmdline))
	r.Handle(profilePattern, http.HandlerFunc(pprof.Profile))
	r.Handle(symbolPattern, http.HandlerFunc(pprof.Symbol))
	r.Handle(tracePattern, http.HandlerFunc(pprof.Trace))

}
