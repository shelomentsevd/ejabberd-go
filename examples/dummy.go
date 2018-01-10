package main

import "github.com/shelomentsevd/ejabberd-go"

type Dummy struct {
}

func (dummy Dummy) Auth(user, server, password string) bool {
	return true
}

func (dummy Dummy) IsUser(user, server string) bool {
	return true
}

func (dummy Dummy) SetPassword(user, server, password string) bool {
	return false
}

func (dummy Dummy) TryRegister(user, server, password string) bool {
	return false
}

func (dummy Dummy) RemoveUser(user, server string) bool {
	return false
}

func (dummy Dummy) RemoveUser3(user, server, password string) bool {
	return false
}

func main() {
	external := jabberd.NewExternal(dummy)

	external.Start()
}
