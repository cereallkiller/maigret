package chrome

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"

	log "github.com/sirupsen/logrus"
)

const listeningURL string = "127.0.0.1"

type forwardingProxy struct {
	targetURL *url.URL
	server    *httputil.ReverseProxy
	listener  net.Listener
	port      int
}

func (proxy *forwardingProxy) start() error {
	log.WithFields(log.Fields{"target-url": proxy.targetURL}).Debug("Initializing functions for forwarding proxy")

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	proxy.targetURL.Path = "/"
	proxy.server = httputil.NewSingleHostReverseProxy(proxy.targetURL)
	proxy.server.Transport = transport

	var err error
	proxy.listener, err = net.Listen("tcp", listeningURL+":0")
	if err != nil {
		return err
	}

	proxy.port = proxy.listener.Addr().(*net.TCPAddr).Port

}
