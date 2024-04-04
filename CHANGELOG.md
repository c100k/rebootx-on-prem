# CHANGELOG

## http-server-go/unreleased

* refactor(http-server-go)!: modularize runnables => 💥 BREAKING CHANGE : You need to adjust the environment variable names if you run the server with custom ones (see `impl/http-server-go/config.go`)

## http-server-go/v0.2.1 (2024-03-29)

* chore: make docker-compose env vars overridable by @c100k in https://github.com/c100k/rebootx-on-prem/pull/12
* fix(http-server-go): adjust json file service implementation by @c100k in https://github.com/c100k/rebootx-on-prem/pull/13

## spec/v0.2.1 (2024-03-29)

* chore(deps-dev): bump @typescript-eslint/eslint-plugin from 7.3.1 to 7.4.0 by @dependabot in https://github.com/c100k/rebootx-on-prem/pull/8
* chore(deps-dev): bump @typescript-eslint/parser from 7.3.1 to 7.4.0 by @dependabot in https://github.com/c100k/rebootx-on-prem/pull/9
* chore(deps-dev): bump fast-check from 3.17.0 to 3.17.1 by @dependabot in https://github.com/c100k/rebootx-on-prem/pull/10
* chore(deps-dev): bump eslint-plugin-sonarjs from 0.24.0 to 0.25.0 by @dependabot in https://github.com/c100k/rebootx-on-prem/pull/11

## http-server-go/v0.2.0 (2024-03-25)

* docs: update docker run volume args by @c100k in https://github.com/c100k/rebootx-on-prem/pull/5
* feat!: respect rest convention by @c100k in https://github.com/c100k/rebootx-on-prem/pull/6

## spec/v0.2.0 (2024-03-25)

* docs: update docker run volume args by @c100k in https://github.com/c100k/rebootx-on-prem/pull/5
* feat!: respect rest convention by @c100k in https://github.com/c100k/rebootx-on-prem/pull/6

## http-server-go/v0.1.0 (2024-03-25)

This is the first release of the http-server-go, implementing the [spec](https://github.com/c100k/rebootx-on-prem/releases/tag/spec%2Fv0.1.0).

Give it a try in a couple of seconds :

```sh
# Make the downloaded binary executable
chmod u+x ~/Downloads/rebootx-on-prem-http-server-go-linux-amd64

# Run it in a debian docker container
docker run --rm -e RBTX_API_KEY="supersecret" -e RBTX_PATH_PREFIX="covfefe" -p "8080:8080" -v ~/Downloads/rebootx-on-prem-http-server-go-linux-amd64:/http-server-go debian:latest /http-server-go

# Make a request to the server
curl -v -H "Authorization: supersecret" http://localhost:8080/covfefe/runnables
```

## spec/v0.1.0 (2024-03-25)

This is the first release of the spec containing the following endpoints : `list`, `reboot`, `stop`.

To browse the spec, the fastest way is to clone this repository and execute `docker-compose up`. It will spin up [SwaggerUI](https://swagger.io/tools/swagger-ui) connected to a basic noop server.

You can also import the `swagger.json` file in any compatible tool of your choice (e.g. [Insomnia](https://github.com/Kong/insomnia/releases), [Postman](https://github.com/postmanlabs)).
