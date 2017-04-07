package udprouter

// MockUDPRouter is a mockable UDPRouter
type MockUDPRouter struct {
	HandleFn func(route string, handleFn UDPHandleFunc)
	ListenFn func(addr string) error
}

// Handle calls the underlying Handle method
func (r *MockUDPRouter) Handle(route string, handleFn UDPHandleFunc) {
	r.HandleFn(route, handleFn)
}

// Listen calls the underlying Listen method
func (r *MockUDPRouter) Listen(addr string) error {
	return r.ListenFn(addr)
}
