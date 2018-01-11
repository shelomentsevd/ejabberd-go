package main

import (
	"log"
	"os"

	"github.com/shelomentsevd/ejabberd-go"
)

type Dummy struct {
}

func (dummy Dummy) Auth(user, server, password string) bool {
	log.Printf("Authorization user: %s, server: %s, passwd: %s\n", user, server, password)
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
	log.SetOutput(os.Stderr)

	dummy := Dummy{}
	external := ejabberd.NewExternal(dummy)

	external.Start()
}
