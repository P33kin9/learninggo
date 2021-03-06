package main

import "fmt"

// notifier is an interface that defined notificaiton
// type behavior.
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// main is the entry point for the application.
func main() {
	// Create a value of type User and send a notification.
	u := user{"Bill", "bill@email.com"}

	sendNotification(&u)

	// ./listing36.go:32: cannot use u (type user) as type
	// 		notifier in argument to sendNotification:
	// user does not implement notifer
	//			(notify method has pointer receiver)
}

// sendNotification accepts value that implement the notifier
// interface and sends notification.
func sendNotification(n notifier) {
	n.notify()
}
