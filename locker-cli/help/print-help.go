package help

import (
	"lkcli/logger"
)

const HelpText string = `
lk create repo <repo-name>
lk update repo <repo-name> <new-repo-name>
lk delete repo <repo-name>

lk create env <repo-name>/<env-name>
lk update env <repo-name>/<env-name> <new-env-name>
lk delete env <repo-name>/<env-name>

lk create secret <repo-name>/<env-name>/<secret-name> <secret-value>
lk update secret <repo-name>/<env-name>/<secret-name> <secret-value>
lk delete secret <repo-name>/<env-name>/<secret-name>

lk get <repo-name>/<env-name>/<secret-name>
lk get <repo-name>/<env-name>
lk get <repo-name>

--- EXAMPLES ---
lk create repo book-api

lk create env book-api/dev
lk create env book-api/stg
lk create env book-api/prd

lk create secret book-api/dev/POSTGRESS_USERNAME postgres

lk create secret book-api/dev/POSTGRESS_PASSWORD $(lk rand 40)
`

func PrintHelpThenExit() {
	logger.Exit(HelpText)
}
