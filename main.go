package hosts

import (
	"net/http"
	"net/url"

	"github.com/nektro/go-util/util"
)

var (
	allowed = []string{}
)

func Allow(s string) {
	allowed = append(allowed, s)
}

func Allowed(r *http.Request) bool {
	u, _ := url.Parse("http://" + r.Host)
	return util.Contains(allowed, u.Hostname())
}
