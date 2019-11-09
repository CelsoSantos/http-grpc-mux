package kncloudevents

import (
	"net"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/transport/http"
)

// NewDefaultClient (optionally) receives a listener and a target
func NewDefaultClient(l net.Listener, target ...string) (cloudevents.Client, error) {
	tOpts := []http.Option{cloudevents.WithBinaryEncoding()}
	if len(target) > 0 && target[0] != "" {
		tOpts = append(tOpts, cloudevents.WithTarget(target[0]))
	}

	if l != nil {
		tOpts = append(tOpts, http.WithListener(l))
	}

	// Make an http transport for the CloudEvents client.
	t, err := cloudevents.NewHTTPTransport(tOpts...)
	if err != nil {
		return nil, err
	}

	// Use the transport to make a new CloudEvents client.
	c, err := cloudevents.NewClient(t,
		cloudevents.WithUUIDs(),
		cloudevents.WithTimeNow(),
	)

	if err != nil {
		return nil, err
	}
	return c, nil
}
