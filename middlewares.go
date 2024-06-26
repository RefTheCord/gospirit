
package server

import (
	"net"
	"net/http"
	"net/url"
	"github.com/RefTheCord/gospirit/configuration"
	"log"
	"strings"
)

// Generally is not a good idea to serve all requests on the blog IP even with empty or unknown host header.
// The good practice is to serve requests with correct 'Host' header and return 400 otherwise.
// See rfc2616 for details.

func CheckHost(next http.Handler) http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		parsed, err := url.Parse(configuration.Config.Url)
		if err != nil {
			log.Fatal("Error: Couldn't parse the Config.Url:", err)
		}
		host, _, _ := net.SplitHostPort(parsed.Host)

		if !strings.EqualFold(r.Host, "") {
			if (strings.EqualFold(r.Host, host) || strings.EqualFold(r.Host, parsed.Host)) {
				next.ServeHTTP(w, r)
				return
			}
		}

		http.Error(w, http.StatusText(400), 400)
		return
	})
}
