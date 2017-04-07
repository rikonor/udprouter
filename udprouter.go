package udprouter

import (
	"fmt"
	"net"
	"strings"
)

const (
	maxBufferSize = 1024
)

// UDPResponseFunc can be used to respond to a request
type UDPResponseFunc func(response []byte)

// UDPHandleFunc defines a UDP route handler
type UDPHandleFunc func(body string, respond UDPResponseFunc)

// UDPRouter can be used to register route handlers and listen to incoming requests
type UDPRouter interface {
	Handle(route string, handleFn UDPHandleFunc)
	Listen(addr string) error
}

type udpRouter struct {
	routeHandlers map[string]UDPHandleFunc
}

// NewUDPRouter will create a new UDPRouter
func NewUDPRouter() UDPRouter {
	routeHandlers := make(map[string]UDPHandleFunc)

	return &udpRouter{
		routeHandlers: routeHandlers,
	}
}

// Handle will register a handler for a given route
func (r *udpRouter) Handle(route string, handleFn UDPHandleFunc) {
	r.routeHandlers[route] = handleFn
}

// Listen will start listening for incoming requests and handle them as they come
func (r *udpRouter) Listen(addr string) error {
	laddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Get requests from connection and route them to handlers
	for {
		buf := make([]byte, maxBufferSize)
		n, raddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			return err
		}

		request := string(buf[:n])

		// break request into route and body
		parts := strings.Split(request, "|")
		if len(parts) != 2 {
			fmt.Printf("error: invalid request: %s\n", request)
			continue
		}

		route := parts[0]
		body := parts[1]

		// route request to handler
		handler, ok := r.routeHandlers[route]
		if !ok {
			fmt.Printf("error: no route found: %s\n", request)
			continue
		}

		respond := func(response []byte) {
			conn.WriteToUDP(response, raddr)
		}

		handler(body, respond)
	}
}
