package webclient

import (
	"net/http"
	"time"
)

// Option is a type of functions-setters
type Option func(*http.Transport)

// MaxConn sets up max idle connections for client
func MaxConn(mc int) Option {
	return func(tr *http.Transport) {
		if mc > 0 {
			tr.MaxIdleConns = mc
		}
	}
}

// IdleTimeout sets up max idle timeout for client
func IdleTimeout(timeout time.Duration) Option {
	return func(tr *http.Transport) {
		if timeout.Seconds() != 0 {
			tr.IdleConnTimeout = timeout
		}
	}
}
