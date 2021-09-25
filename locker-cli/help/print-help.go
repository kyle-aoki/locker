package help

import (
	"lkcli/logger"
)

const HelpText string = `
lk create repo <repo>
lk update repo <repo> <new-repo>
lk delete repo <repo>

lk create env <repo>/<env>
lk update env <repo>/<env> <new-env>
lk delete env <repo>/<env>

lk create secret <repo>/<env>/<secret> <secret-value>
lk update secret <repo>/<env>/<secret> <secret-value>
lk delete secret <repo>/<env>/<secret>

lk get <repo>/<env>/<secret>
lk get <repo>/<env>
lk get <repo>

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
