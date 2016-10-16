package serfer

import (
	"bytes"
	"sync"
	"testing"

	"github.com/hashicorp/serf/serf"
)

// TODO: Replace this shitty test
func TestSerferShortMessage(t *testing.T) {
	// Setup router
	cfg := serf.DefaultConfig()
	cfg.NodeName = "master"
	cfg.MemberlistConfig.BindPort = 49510
	cfg.MemberlistConfig.AdvertisePort = 49510

	events := make(chan serf.Event)
	cfg.EventCh = events

	sc, err := serf.Create(cfg)
	if err != nil {
		t.Fatalf("failed to create serf client: %s", err)
	}

	msgName := "test-msg"
	msgBody := []byte("test-text")

	sr := NewSerfEventRouter(sc, events)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	sr.RegisterRoute(MatcherUserEvent, func(event serf.Event) {
		userEvent, ok := event.(serf.UserEvent)
		if !ok {
			t.Fatal("unexpected event type")
		}
		if userEvent.Name != msgName {
			t.Fatalf("wrong message name: %s", userEvent.Name)
		}
		if !bytes.Equal(userEvent.Payload, msgBody) {
			t.Fatalf("wrong message body: %s", string(userEvent.Payload))
		}
		wg.Done()
	})
	go sr.Start()

	if err := sc.UserEvent(msgName, msgBody, true); err != nil {
		t.Fatalf("failed to send message: %s", err)
	}

	wg.Wait()
}
