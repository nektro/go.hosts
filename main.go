package hosts

import (
	"net/http"
	"net/url"

	"github.com/nektro/go-util/util"
)

var (
	allowed = []string{}
	filter  = true
)

func Allow(s string) {
	allowed = append(allowed, s)
}

func AllowAll() {
	filter = false
}

func Allowed(r *http.Request) bool {
	if !filter {
		return true
	}
	u, _ := url.Parse("http://" + r.Host)
	return util.Contains(allowed, u.Hostname())
}
