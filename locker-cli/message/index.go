package message

const Host1 string = "No Host Found. Try: lk set host <host-name>"
const Host2 string = `
Try: lk set host <host-name>
ex.: lk set host http://localhost:8080
`

const LogIn1 string = "You must log in first. Try: lk login <username> <password>"
const LogIn2 string = "Try: lk login <username> <password>"

const CreateRepo1 string = "Try: lk create repo <repo-name>"

const CreateEnv1 string = "Try: lk create env <repo-name>/<env-name>"

const CreateSecret1 string = "Try: lk create secret <repo-name>/<env-name>/<secret-name> <secret-value>"

const GetSecret1 string = "Try: lk get <repo>/<env>/<secret-name>"
