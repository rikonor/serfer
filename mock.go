package serfer

// MockSerfEventRouter is a mockable SerfEventRouter
type MockSerfEventRouter struct {
	RegisterRouteFn        func(MatcherFunc, RouteFunc)
	RegisterRouteFnInvoked bool

	StartFn        func()
	StartFnInvoked bool
}

// RegisterRoute invokes the underyling RegisterRoute method
func (s *MockSerfEventRouter) RegisterRoute(mfn MatcherFunc, rfn RouteFunc) {
	s.RegisterRouteFnInvoked = true
	s.RegisterRouteFn(mfn, rfn)
}

// Start invokes the underyling Start method
func (s *MockSerfEventRouter) Start() {
	s.StartFnInvoked = true
	s.StartFn()
}
