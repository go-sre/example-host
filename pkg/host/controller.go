package host

import (
	"errors"
	"fmt"
	"github.com/gotemplates/example-host/pkg/resource"
	"github.com/gotemplates/host/controller"
	"github.com/gotemplates/host/runtime"
	"net/http"
)

const (
	ingressRoutesNameFmt = "fs/ingress_routes_{env}.json"
	egressRoutesNameFmt  = "fs/egress_routes_{env}.json"
)

func initControllers() []error {
	errs := initIngressControllers()
	if len(errs) == 0 {
		errs = initEgressControllers()
	}
	return errs
}

func initIngressControllers() []error {
	lookup := make(map[string]string, 16)

	errs := controller.InitIngressControllers(
		func() ([]byte, error) {
			return resource.ReadFile(runtime.EnvExpansion(ingressRoutesNameFmt))
		},
		func(routes []controller.Route) error {
			return updateIngressLookup(lookup, routes)
		})
	if len(errs) > 0 {
		return errs
	}

	controller.IngressTable().SetHttpMatcher(func(req *http.Request) (string, bool) {
		if req == nil || req.URL == nil {
			return "", false
		}
		s := lookup[req.URL.Path]
		if s != "" {
			return s, true
		}
		return "", false
	})
	return nil
}

func updateIngressLookup(m map[string]string, routes []controller.Route) error {
	for _, r := range routes {
		if r.Pattern == "" {
			continue
		}
		if _, ok := m[r.Pattern]; ok {
			return errors.New(fmt.Sprintf("invalid route, pattern is duplicated: %v", r.Name))
		}
		m[r.Pattern] = r.Name
	}
	return nil
}

func initEgressControllers() []error {
	lookup := make(map[string]string, 16)

	errs := controller.InitEgressControllers(
		func() ([]byte, error) {
			return resource.ReadFile(runtime.EnvExpansion(egressRoutesNameFmt))
		},
		func(routes []controller.Route) error {
			return updateEgressLookup(lookup, routes)
		})
	if len(errs) > 0 {
		return errs
	}

	controller.EgressTable().SetHttpMatcher(func(req *http.Request) (string, bool) {
		if req == nil || req.URL == nil {
			return "", false
		}
		return lookup[req.URL.Host], true
	})
	return nil
}

func updateEgressLookup(m map[string]string, routes []controller.Route) error {
	for _, r := range routes {
		if r.Pattern == "" {
			continue
		}
		if _, ok := m[r.Pattern]; ok {
			return errors.New(fmt.Sprintf("invalid route, pattern is duplicated: %v", r.Name))
		}
		m[r.Pattern] = r.Name
	}
	return nil
}
