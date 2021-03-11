// Sample program to show how you  can personally mock concrete types when
// you need to for your own packages or tests.
package main

import "goination/ch10/listing06/pubsub"

// publisher is an interface to allow this package to mock the
// pubsub package support.
type publisher interface {
	Publish(key string, v interface{}) error
	Subcribe(key string) error
}

// mock is a concrete type to help support the mocking of the
// pubsub package.
type mock struct{}

// Publish implement the publisher interface for the mock
func (m *mock) Publish(key string, v interface{}) error {

	// ADD YOUR MOCK FOR THE PUBLISH CALL.
	return nil
}

// Subscibe implements the publisher interface for the mock.
func (m *mock) Subcribe(key string) error {

	// ADD YOUR MOCK FOR THE SUBCIRBE CALL.
	return nil
}

func main() {

	// Create a slice of publisher interface values. Assign
	// the address of a pubsub.PubSub value and the address of
	// a mock value.
	pubs := []publisher{
		pubsub.New("localhost"),
		&mock{},
	}

	// Range over the interface value to see how the publisher
	// interface provides the level of decouping the user needs.
	// The pubsub package did not need to provide the interface type.
	for _, p := range pubs {
		p.Publish("key", "value")
		p.Subcribe("key")
	}
}
