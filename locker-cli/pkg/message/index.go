package message

const Host1 string = "No Host Found. Try: lk set host <host-name>"
const Host2 string = `
Try: lk set host <host-name>
ex.: lk set host http://localhost:8080
`

const LogIn1 string = "You must log in first. Try: lk login <username> <password>"
const LogIn2 string = "Try: lk login <username> <password>"

const CreateRepo1 string = "Try: lk create repo <repo>"

const CreateEnv1 string = "Try: lk create env <repo>/<env>"

const CreateSecret1 string = "Try: lk create secret <repo>/<env>/<secret> <secret-value>"

const GetSecret1 string = "Try: lk get <repo>/<env>/<secret>"

const ListRepos1 string = "Try: lk list repos <limit?> <offset?>"
const ListSecrets1 string = "Try: lk list secrets <repo>/<env>"

const GetHelp1 string = "Try: lk get <repo>\nor\nlk get <repo>/<env>\nor\nlk get <repo>/<env>/<secret>"
