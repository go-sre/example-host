package host

import (
	"fmt"
	"github.com/gotemplates/core/runtime"
	"net/http"
)

func Example_Startup() {
	_, status := Startup[runtime.DebugError, runtime.StdOutput](http.NewServeMux())
	fmt.Printf("test: Startup() -> %v\n", status)

	//Output:
	//startup successful for resource [github.com/idiomatic-go/resiliency/service/pkg/facebook] : 0s
	//startup successful for resource [github.com/idiomatic-go/resiliency/service/pkg/google] : 0s
	//startup successful for resource [github.com/idiomatic-go/resiliency/service/pkg/twitter] : 0s
	//test: Startup() -> OK

}
