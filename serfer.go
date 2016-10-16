package serfer

import "github.com/hashicorp/serf/serf"

// MatcherFunc is a function that is run against an event to produce a yes/no
type MatcherFunc func(serf.Event) bool

// RouteFunc is a function that should be run when its route is hit
type RouteFunc func(serf.Event)

type routerAction struct {
	mfn MatcherFunc
	rfn RouteFunc
}

// SerfEventRouter is an event router built on top of Serf
type SerfEventRouter interface {
	RegisterRoute(MatcherFunc, RouteFunc)
	Start()
}

type serfEventRouter struct {
	sc      *serf.Serf
	events  chan serf.Event
	actions []*routerAction
}

// NewSerfEventRouter creates a new SerfEventRouter
func NewSerfEventRouter(sc *serf.Serf, events chan serf.Event) SerfEventRouter {
	return &serfEventRouter{
		sc:     sc,
		events: events,
	}
}

func (s *serfEventRouter) RegisterRoute(mfn MatcherFunc, rfn RouteFunc) {
	s.actions = append(s.actions, &routerAction{mfn: mfn, rfn: rfn})
}

func (s *serfEventRouter) Start() {
	for event := range s.events {
		s.routeEvent(event)
	}
}

func (s *serfEventRouter) routeEvent(event serf.Event) {
	for _, action := range s.actions {
		if action.mfn(event) {
			action.rfn(event)
		}
	}
}
