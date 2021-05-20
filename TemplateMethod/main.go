package main

type AbstConnectHelper interface {
	doSecurity(info string) string
	authentication(id string, password string) string
	authorization(userName string) int
	connection(info string) string
	requestConnection(info string) string
}

func main() {
	// if use interface, we use &struct{}. not struct{}
	implementation := &ConcreteImplementation{}
	connectHelper := &ConcreteConnectHelper{
		chelper: implementation,
	}
	connectHelper.requestConnection("info")
}
