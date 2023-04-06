package facebook

import (
	"context"
	"github.com/go-http-utils/headers"
	"github.com/go-sre/core/exchange"
	"github.com/go-sre/core/runtime"
	"net/http"
)

const (
	uri = "https://www.facebook.com"
)

var homeLoc = pkgPath + "/home"

func Get[E runtime.ErrorHandler](ctx context.Context) ([]byte, *runtime.Status) {
	var e E

	if ctx == nil {
		ctx = context.Background()
	}
	req, err := http.NewRequestWithContext(ctx, "GET", uri, nil)
	if err != nil {
		return nil, e.Handle(ctx, homeLoc, err)
	}
	req.Header.Add("x-request-id", runtime.ContextRequestId(ctx))
	resp, buf, status := exchange.Do[E, exchange.Default, []byte](req)
	if !status.OK() {
	} else {
		status.SetMetadataFromResponse(resp, headers.ContentType)
	}
	return buf, status
}
