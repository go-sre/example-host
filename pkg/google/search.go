package google

import (
	"context"
	"github.com/go-http-utils/headers"
	"github.com/gotemplates/core/exchange"
	"github.com/gotemplates/core/runtime"
	"net/http"
	"net/url"
)

const (
	homeUri   = "https://www.google.com"
	searchUri = "https://www.google.com/search?q="
)

var searchLocation = pkgPath + "/search"

func Search[E runtime.ErrorHandler](ctx context.Context, uri *url.URL) ([]byte, *runtime.Status) {
	var e E

	if ctx == nil {
		ctx = context.Background()
	}
	req, err := http.NewRequestWithContext(ctx, "GET", createUri(uri), nil)
	if err != nil {
		return nil, e.HandleWithContext(ctx, searchLocation, err)
	}
	req.Header.Add("x-request-id", runtime.ContextRequestId(ctx))
	resp, buf, status := exchange.DoT[E, []byte, exchange.Default](req)
	if !status.OK() {

	} else {
		status.SetMetadataFromResponse(resp, headers.ContentType)
	}
	return buf, status
}

func createUri(uri *url.URL) string {
	if uri == nil {
		return homeUri
	}
	v := uri.Query()
	if v == nil {
		return homeUri
	}
	if str := v["q"]; len(str) > 0 {
		return searchUri + str[0]
	}
	return homeUri
}
