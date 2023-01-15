// Code generated by nrpc; DO NOT EDIT.

package example

import (
	"strings"
	"time"
)

type Server interface {
	Start() error
	Stop() error
}

// An option that can be used on both the server
// and client.
type Opt interface {
	ServerOpt
	ClientOpt
}

func newOpt(so func(*ServerOptions) error, co func(*ClientOptions) error) Opt {
	return &opt{so, co}
}

// A server configuration value
type ServerOpt interface {
	setServer(*ServerOptions) error
}

// A client configuration value
type ClientOpt interface {
	setClient(*ClientOptions) error
}

// Client configuration values
type ClientOptions struct {
	// Timeout sets the amount of time that a client will
	// wait for a response from the server
	Timeout   time.Duration
	// Namespace will be added to the beginning of all NATS
	// subjects used by the client, effectively allowing
	// multiple servers to be run and accessed manually.
	Namespace string
}

func defaultClientOptions() *ClientOptions {
	return &ClientOptions{
		Timeout: 10 * time.Second,
	}
}

type ServerOptions struct {
	// Namespace will be added to the beginning of all NATS
	// subjects the server listens on, effectively allowing
	// multiple servers to be run.
	Namespace    string
	// A handler for errors that occur during service calls.
	// Errors will be given
	ErrorHandler func(error)
	// The maximum number of pending messages that each endpoint
	// in the server supports. This size is per-endpoint.
	BufferSize   int
}

func defaultServerOptions() *ServerOptions {
	return &ServerOptions{
		BufferSize: 64,
	}
}

type opt struct {
	so func(*ServerOptions) error
	co func(*ClientOptions) error
}

func (o *opt) setServer(so *ServerOptions) error {
	return o.so(so)
}

func (o *opt) setClient(co *ClientOptions) error {
	return o.co(co)
}

type serverOptFunc func(o *ServerOptions) error

func (sof serverOptFunc) setServer(o *ServerOptions) error {
	return sof(o)
}

type clientOptFunc func(o *ClientOptions) error

func (cof clientOptFunc) setClient(o *ClientOptions) error {
	return cof(o)
}

// Set the namespace used for all NATS subjects
func Namespace(ns string) ClientOpt {
	return newOpt(
		func(o *ServerOptions) error {
			if ns != "" {
				o.Namespace = strings.Trim(ns, ".") + "."
			}
			return nil
		},
		func(o *ClientOptions) error {
			if ns != "" {
				o.Namespace = strings.Trim(ns, ".") + "."
			}
			return nil
		},
	)
}

// Set the maximum number of buffered messages
// for each server endpoint
func ServerBufferSize(bs int) ServerOpt {
	return serverOptFunc(func(o *ServerOptions) error {
		if bs >= 0 {
			o.BufferSize = bs
		}
		return nil
	})
}

// Set the error handler for the server
func ServerErrorHandler(eh func(error)) ServerOpt {
	return serverOptFunc(func(o *ServerOptions) error {
		o.ErrorHandler = eh
		return nil
	})
}

// Set the maximum amount of time the client will wait
// for a response from a server
func ClientTimeout(to time.Duration) ClientOpt {
	return clientOptFunc(func(o *ClientOptions) error {
		if to > 0 {
			o.Timeout = to
		}
		return nil
	})
}
