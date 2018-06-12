package ejabberd

import (
	"bufio"
	"encoding/binary"
	"os"
	"strings"
)

type Authorizator interface {
	Auth(user, server, password string) bool
}

type UserChecker interface {
	IsUser(user, server string) bool
}

type PasswordChanger interface {
	SetPassword(user, server, password string) bool
}

type UserRegister interface {
	TryRegister(user, server, password string) bool
}

type UserRemover interface {
	RemoveUser(user, server string) bool
}

type UserRemover3 interface {
	RemoveUser3(user, server, password string) bool
}

type External struct {
	ExternalMethod interface{}
}

func NewExternal(method interface{}) External {
	return External{
		ExternalMethod: method,
	}
}

func (e *External) Start() error {
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	var (
		success bool
		length  uint16
		result  uint16
	)
	for {
		if err := binary.Read(input, binary.BigEndian, &length); err != nil {
			return err
		}

		buffer := make([]byte, length)
		_, err := input.Read(buffer)
		if err != nil {
			return err
		}

		data := strings.Split(string(buffer), ":")
		switch data[0] {
		case "auth":
			if auth, ok := e.ExternalMethod.(Authorizator); ok {
				user, server, password := data[1], data[2], data[3]
				success = auth.Auth(user, server, password)
			}
		case "isuser":
			if userCheck, ok := e.ExternalMethod.(UserChecker); ok {
				user, server := data[1], data[2]
				success = userCheck.IsUser(user, server)
			}
		case "setpass":
			if passwordChanger, ok := e.ExternalMethod.(PasswordChanger); ok {
				user, server, password := data[1], data[2], data[3]
				success = passwordChanger.SetPassword(user, server, password)
			}
		case "tryregister":
			if register, ok := e.ExternalMethod.(UserRegister); ok {
				user, server, password := data[1], data[2], data[3]
				success = register.TryRegister(user, server, password)
			}
		case "removeuser":
			if removeUser, ok := e.ExternalMethod.(UserRemover); ok {
				user, server := data[1], data[2]
				success = removeUser.RemoveUser(user, server)
			}
		case "removeuser3":
			if removeUser, ok := e.ExternalMethod.(UserRemover3); ok {
				user, server, password := data[1], data[2], data[3]
				success = removeUser.RemoveUser3(user, server, password)
			}
		default:
		}

		result = 0
		if success {
			result = 1
		}

		length = 2
		binary.Write(output, binary.BigEndian, &length)
		binary.Write(output, binary.BigEndian, &result)
		output.Flush()
	}
}
