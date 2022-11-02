package webclient

import (
	"net/http"
	"time"
)

// Conn is a http client
type Conn struct {
	W *http.Client
}

// New is a constructor for Conn
func New(opts ...Option) *Conn {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 5 * time.Second,
	}
	for _, opt := range opts {
		opt(tr)
	}
	client := &http.Client{Transport: tr}
	return &Conn{W: client}
}
