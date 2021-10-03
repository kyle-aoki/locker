package help

import (
	"lkcli/logger"
)

const helpText string = `-------------------- COMMANDS --------------------

lk set host <lk-server-host>
lk login <username <password>

lk create repo    <repo>
lk create env     <repo>/<env>
lk create secret  <repo>/<env>/<secret> <secret-value>
 
lk update repo    <repo> <new-repo>
lk update env     <repo>/<env> <new-env>
lk update secret  <repo>/<env>/<secret> <new-secret-value>
 
lk delete repo    <repo>
lk delete env     <repo>/<env>
lk delete secret  <repo>/<env>/<secret>

lk get <repo>/<env>/<secret>
lk get <repo>/<env>
lk get <repo>

lk missing <repo>/<env> <target-env>
`

func PrintHelpThenExit() {
	logger.Exit(helpText)
}
