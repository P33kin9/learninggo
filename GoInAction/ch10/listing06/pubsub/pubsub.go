// Sample pubsub simulates a package that provides
// pubulication/subcription type services.
package pubsub

// PubSub provides access to a queue system.
type PubSub struct {
	host string

	// PRETEND THERE ARE MORE FIELDS.
}

// New creates a pubsub value for use.
func New(host string) *PubSub {
	ps := PubSub{
		host: host,
	}

	// PRETEND THERE IS A SPECIFIC IPLEMENTATION.

	return &ps
}

// Publish sends the data for the specified key.
func (ps *PubSub) Publish(key string, v interface{}) error {

	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}

// Subscribe sets up an request to receive messages for the
// specified key.
func (ps *PubSub) Subcribe(key string) error {

	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}
