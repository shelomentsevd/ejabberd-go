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
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	for {
		binary.Read(input, binary.BigEndian, &length)

		buffer := make([]byte, length)

		r, _ := input.Read(buffer)
		if r == 0 {
			continue
		}

		data := strings.Split(string(buffer), ":")
		switch data[0] {
		case "auth":
		case "isuser":
		case "setpass":
		case "tryregister":
		case "removeuser":
		case "removeuser3":
		default:
			success = false
		}

		result = 0
		if success {
			result = 1
		}

		length = 2
		binary.Write(input, binary.BigEndian, &result)
		input.Flush()
	}
}
