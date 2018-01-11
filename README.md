# ejabberd-go
Script for external authentication to [Ejabberd](https://www.ejabberd.im/) server
# Usage
1. Import the library and implement one of interfaces:
* `Authorizator` - to provide custom authorization
* `UserChecker` - to provide custom user validation
* `PasswordChanger` - to provide custom password change
* `UserRegister` - to provide custom user registration
* `UserRemover` or `UserRemover3` - to provide custom user remove
2. Pass your object to `NewExternal`
```golang
dummy := Dummy{}
external := ejabberd.NewExternal(dummy)

external.Start()
```
# Example
`dummy.go` - Accepts all authorizations with any user and password
1. Compile examples/dummy.go
2. Edit your ejabberd configuration file to:
```yaml
auth_method: [external]
extauth_program: "/etc/ejabberd/compiledGoCode"
extauth_instances: 3
auth_use_cache: false
```
3. Run your server and authorize under any user with any password

# Other
For more information see [Ejabberd developers guide](https://www.ejabberd.im/files/doc/dev.html)
