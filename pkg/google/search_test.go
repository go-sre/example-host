package google

import (
	"context"
	"fmt"
	"github.com/gotemplates/core/exchange/httptest"
	"github.com/gotemplates/core/runtime"
	"github.com/gotemplates/host/controller"
	"github.com/gotemplates/host/middleware"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"
)

//const (
//	searchTestUri = "proxy://www.google.com/search?q="
//)

func init() {
	name := "google-search"
	middleware.ControllerWrapTransport(nil)
	controller.EgressTable.SetDefaultController(controller.NewRoute(controller.DefaultControllerName, "egress", "", false, controller.NewTimeoutConfig(time.Millisecond*2000, http.StatusGatewayTimeout)))
	controller.EgressTable.AddController(controller.NewRoute(name, "egress", "", false))
	controller.EgressTable.SetHttpMatcher(func(req *http.Request) (string, bool) {
		return name, true
	},
	)
}

var searchTestCtx = runtime.ContextWithProxy(context.Background(), searchTestProxy)

func searchTestProxy(req *http.Request) (*http.Response, error) {
	if req == nil || req.URL == nil {
		return nil, nil
	}
	switch req.URL.String() {
	case searchUri:
		resp := httptest.NewResponse(http.StatusOK, []byte("Override test content"), "content-type", "text/plain")
		return resp, nil
	default:
	}
	return nil, nil
}

func Example_createUri() {
	uri := createUri(nil)
	fmt.Printf("test: createUri(nil) -> %v\n", uri)

	req, _ := http.NewRequest("", "https://www.google.com/search?q=test", nil)
	uri = createUri(req.URL)
	fmt.Printf("test: createUri(nil) -> %v\n", uri)

	req, _ = http.NewRequest("", "proxy://www.google.com/search?q=test", nil)
	uri = createUri(req.URL)
	fmt.Printf("test: createUri(nil) -> %v\n", uri)

	//Output:
	//test: createUri(nil) -> https://www.google.com
	//test: createUri(nil) -> https://www.google.com/search?q=test
	//test: createUri(nil) -> https://www.google.com/search?q=test

}

func ExampleSearch_Success() {
	buff, status := Search[runtime.DebugError](nil, nil)
	fmt.Printf("test: Search() -> [%v] [content:%v]\n", status, buff != nil)

	//Output:
	//test: Search() -> [OK] [content:true]

}

func ExampleSearch_TestContent() {
	u, _ := url.Parse(searchUri)
	buff, status := Search[runtime.DebugError](searchTestCtx, u)
	fmt.Printf("test: Search() -> [%v] [content:%v]\n", status, string(buff))

	//Output:
	//test: Search() -> [OK] [content:Override test content]

}

func ExampleSearch_HttpError() {
	u, _ := url.Parse(searchUri)
	buff, status := Search[runtime.DebugError](runtime.ContextWithProxy(context.Background(), httptest.HttpErrorProxy), u)
	fmt.Printf("test: Search() -> [%v] [content:%v]\n", status, string(buff))

	//Output:
	//[[] github.com/gotemplates/core/exchange/do [http: connection has been hijacked]]
	//test: Search() -> [Internal] [content:]

}

func ExampleSearch_BodyIOError() {
	u, _ := url.Parse(searchUri)
	buff, status := Search[runtime.DebugError](runtime.ContextWithProxy(context.Background(), httptest.BodyIOErrorProxy), u)
	fmt.Printf("test: Search() -> [%v] [content:%v]\n", status, string(buff))

	//Output:
	//[[] github.com/gotemplates/core/exchange/deserialize [unexpected EOF]]
	//test: Search() -> [I/O Failure] [content:]

}

func TestSearch(t *testing.T) {
	type args struct {
		ctx context.Context
		uri *url.URL
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 *runtime.Status
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Search[runtime.DebugError](tt.args.ctx, tt.args.uri)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Search() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

/*
func Test_createUri(t *testing.T) {
	type args struct {
		uri *url.URL
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createUri(tt.args.uri); got != tt.want {
				t.Errorf("createUri() = %v, want %v", got, tt.want)
			}
		})
	}
}


*/
