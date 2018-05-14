# rebase-bot

Forked to make some changes for deployment and TLS support.

## About

the rebase-bot takes care of ensuring proper protocol is followed when working with
pull-requests.

specifically it…

… automatically rebases master once LGTM'd to ensure PR is up-to-date  
… waits for tests and merge automatically when green

## deployment

See depop/platform repo

## development

The rebase-bot has lots of unit tests to ensure it's working as intended. Also
there are some integration tests which operate on a locally cloned repository with pre-defined scenarios.

To execute all tests run `go test ./...`.

For prototyping it's sometimes useful to run the bot locally. This is best done via [ngrok](https://ngrok.io):

1. start ngrok: `ngrok http 8080`
2. start the bot with the endpoint provided by ngrok: `go build . && ./github-rebase-bot -public-dns <ngrok-endpoint> -repos <repo> -merge-label LGTM -addr :8080`

To update the integration test scenarios unarchive `scenarios/rebase-conflict.zip`, change the repository and archive it again using zip: `zip -r ../rebase-conflict.zip .`

## installation

`go get -u github.com/depop/github-rebase-bot`
