package itihost

import (
	"net/http"
)

// HostMatcher Store the schema to match with the request Path
type HostMatcher struct {
	hosts []string
}

// New is the constructor to ItiHost
func New(template ...string) *HostMatcher {
	return &HostMatcher{
		hosts: template,
	}
}

// Match returns if the request can be handled by this Route.
func (h *HostMatcher) Match(req *http.Request) bool {
	for _, v := range h.hosts {
		if v == req.Host {
			return true
		}
	}
	return false
}
