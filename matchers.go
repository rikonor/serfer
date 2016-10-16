package serfer

import "github.com/hashicorp/serf/serf"

func MatcherAny(event serf.Event) bool {
	return true
}

func MatcherUserEvent(event serf.Event) bool {
	_, ok := event.(serf.UserEvent)
	return ok
}

func MatcherQuery(event serf.Event) bool {
	_, ok := event.(*serf.Query)
	return ok
}
