package host

import (
	"fmt"
	"github.com/gotemplates/host/controller"
	"net/http"
)

func Example_initIngressControllers() {
	errs := initIngressControllers()
	fmt.Printf("test: initIngressControllers() -> [err:%v]\n", errs)

	ctrl := controller.IngressTable.Host()
	c, _ := ctrl.RateLimiter()
	fmt.Printf("test: Ingress.Host() -> [controller:%v] [rateLimiter:%v]\n", ctrl.Name(), c != nil)

	req, _ := http.NewRequest("", "https://www.google.com", nil)
	ctrl = controller.IngressTable.LookupHttp(req)
	c1, _ := ctrl.Timeout()
	fmt.Printf("test: Ingress.Lookup(https://www.google.com) -> [controller:%v] [timeout:%v]\n", ctrl.Name(), c1 != nil)

	req, _ = http.NewRequest("", "https://www.google.com/google?q=test", nil)
	ctrl = controller.IngressTable.LookupHttp(req)
	c1, _ = ctrl.Timeout()
	fmt.Printf("test: Ingress.Lookup(https://www.google.com/google/search) -> [controller:%v] [timeout:%v]\n", ctrl.Name(), c1 != nil)

	//Output:
	//test: initIngressControllers() -> [err:[]]
	//test: Ingress.Host() -> [controller:host] [rateLimiter:true]
	//test: Ingress.Lookup(https://www.google.com) -> [controller:!] [timeout:false]
	//test: Ingress.Lookup(https://www.google.com/google/search) -> [controller:google-search-ingress] [timeout:true]

}

func Example_initEgressControllers() {
	errs := initEgressControllers()
	fmt.Printf("test: initEgressControllers() -> [err:%v]\n", errs)

	name := "google-search"
	ctrl := controller.EgressTable.LookupByName(name)
	fmt.Printf("test: LookupByName(%v) -> [controller:%v]\n", name, ctrl.Name())

	name = "google-home"
	ctrl = controller.EgressTable.LookupByName(name)
	fmt.Printf("test: LookupByName(%v) -> [controller:%v]\n", name, ctrl.Name())

	req, _ := http.NewRequest("", "https://www.google.com/search?q=test", nil)
	ctrl = controller.EgressTable.LookupHttp(req)
	fmt.Printf("test: Lookup(https://www.google.com/search?q=test) -> [controller:%v]\n", ctrl.Name())

	req, _ = http.NewRequest("", "https://www.twitter.com", nil)
	ctrl = controller.EgressTable.LookupHttp(req)
	fmt.Printf("test: Lookup(https://www.twitter.com) -> [controller:%v]\n", ctrl.Name())

	req, _ = http.NewRequest("", "https://instagram.com", nil)
	ctrl = controller.EgressTable.LookupHttp(req)
	fmt.Printf("test: Lookup(https://instagram.com) -> [controller:%v]\n", ctrl.Name())

	//Output:
	//test: initEgressControllers() -> [err:[]]
	//test: LookupByName(google-search) -> [controller:google-search]
	//test: LookupByName(google-home) -> [controller:google-home]
	//test: Lookup(https://www.google.com/search?q=test) -> [controller:google-search]
	//test: Lookup(https://www.twitter.com) -> [controller:twitter-home]
	//test: Lookup(https://instagram.com) -> [controller:*]

}
