package ejabberd

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
	TryRegister(user, server, password) bool
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

func (e *External) Start() {
	for {

	}
}
