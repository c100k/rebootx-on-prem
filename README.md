# RebootX On-Prem

![CI](https://github.com/c100k/rebootx-on-prem/actions/workflows/quality.yml/badge.svg)

## What is RebootX On-Prem?

**RebootX On-Prem** is an open source specification for defining a custom server in order to manage on-premise _runnables_ in the [RebootX](https://c100k.eu/p/rebootx) app.

A _Runnable_ is anything that _runs_, can be _stopped_ and _rebooted_. For instance, Virtual Machines (VMs), Dedicated servers, Containers, PaaS Applications, Databases... are all valid concretions of a _runnable_.

If you have already used the [RebootX](https://c100k.eu/p/rebootx) app, you are already familiar with how it works when you connect public cloud providers like Amazon Web Services (AWS), Microsoft Azure, Clever Cloud, Google Cloud (GCP), OVH or Scaleway.

Why should you use this specification ?

It provides the solution if you are in one of these cases :

- You have servers in your local network that you want to manage via an app
- You hack around small devices like the Raspberry Pi and you want to manage it via an app
- You have dedicated servers in a datacenter that do not have a central administration console and you want to manage them via an app

Of course, these are only examples and the only limit is your imagination.

## Getting Started

The specification is as simple as the following endpoints : `list`, `reboot`, `stop`. Of course, it will evolve overtime.

It follows the [OpenAPI Specification](https://swagger.io/specification) allowing a high level of compatibility with existing tools.

You can play with it by loading it locally in [SwaggerUI](https://swagger.io/tools/swagger-ui) with [Docker Compose](https://docs.docker.com/compose):

```sh
# Generate swagger.json (optional since it's already present in the repository)
docker run --rm \
-v ./spec:/spec \
oven/bun run \
/spec/generate-swagger.ts

# Generate Go code with OpenAPI Generator
docker run --rm \
-v ./:/app \
openapitools/openapi-generator-cli generate \
-i /app/spec/_generated/swagger.json \
-g go \
-o /app/impl/http-server-go/vendor/openapi

# Run Swagger UI calling the Go server
docker-compose up
```

You can then access http://localhost:9002 via your browser and test the endpoints. See `docker-compose.yml` to have the `apiKey`.

You can also directly test the server with cURL:

```sh
curl -v -H "Authorization: <apiKey>" http://localhost:9001/cd5331ba/runnables
curl -v -X POST -H "Authorization: <apiKey>" http://localhost:9001/cd5331ba/runnables/reboot/self
curl -v -X POST -H "Authorization: <apiKey>" http://localhost:9001/cd5331ba/runnables/stop/self
```

## Creating your own server

You should develop your own server to fit your personal needs and keep your infra private.

As long as you respect the specification, you can develop in the language of your choice. To speed things up, you can generate some code using [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator).

For instance, you can generate [Rust](https://www.rust-lang.org) code with the following commands :

```sh
docker run --rm \
-v ./:/app \
openapitools/openapi-generator-cli generate \
-i /app/spec/_generated/swagger.json \
-g rust \
-o /app/impl/http-server-rust/openapi
```

Although it can be handful, we do not recommend relying on all the generated code for a **production server**, because it contains too much boilerplate, making it harder to maintain. It's fine to use the generated `structs`, `interfaces`, `enums`, `traits`, though. That being said, it's up to you.

Once ready, your server should get your _runnables_ from a local file, another API, network calls or a database (remote or local). Once again, it's up to you.

It should be deployed to an independent location, accessible by the phone you want to use : Internet if you want access from anywhere / LAN if local access is enough for your use case.

Of course, you must serve it with `HTTPS` to secure the connection between the app and your server.

## Examples

### HTTP Server Go

This is the implementation, used in the `docker-compose` command above, made in Go. You can browse the code at [impl/http-server-go](./impl/http-server-go).

There are multiple service implementations available :

- `noop` (default) : as its name indicates, it does nothing except returning empty payloads
- `self` : it returns the container as a runnable, hence the usage of self. Since the container is not running as root, you should get an error when calling `reboot` or `stop`. But be careful if you run this on a machine as a privileged user. It relies on [syscall](https://pkg.go.dev/syscall) and [exec](https://pkg.go.dev/os/exec) so it can actually reboot or stop the machine.

You can change the value by setting it int the ad-hoc environment variable (see `docker-compose.yml` or `config.go`).

## Contributing

We welcome all kind of contributions to show examples in other languages. Feel free to reach out to us or to publish a PR. It must be clean and self contained in its own directory in [./impl](./impl).

Please read the [CONTRIBUTING.md](./CONTRIBUTING.md) guide first.
