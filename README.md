# RebootX On-Prem

![CI](https://github.com/c100k/rebootx-on-prem/actions/workflows/quality.yml/badge.svg)

## What is RebootX On-Prem?

**RebootX On-Prem** is an open source specification for defining a custom server in order to manage on-premise _runnables_ and _dashboards_ in the [RebootX](https://c100k.eu/p/rebootx) app.

A _Runnable_ is anything that _runs_, can be _stopped_ and _rebooted_. For instance, Virtual Machines (VMs), Dedicated servers, Containers, PaaS Applications, Databases... are all valid concretions of a _runnable_.

A _Dashboard_ is a collection of numeric metrics. For example, number of nodes, number of orders, latency... are all eligible metrics for a dashboard.

If you have already used the [RebootX](https://c100k.eu/p/rebootx) app, you are already familiar with how it works when you connect public cloud providers ([Amazon Web Services (AWS)](https://aws.amazon.com), [Microsoft Azure](https://azure.microsoft.com), [Clever Cloud](https://www.clever-cloud.com), [Google Cloud (GCP)](https://cloud.google.com), [OVH](https://www.ovhcloud.com), [Scaleway](https://www.scaleway.com)) or dashboard managers ([Grafana](https://grafana.com)).

Why should you use this specification ?

- You have servers in your local network that you want to manage via an app
- You hack around small devices like the Raspberry Pi and you want to manage it via an app
- You have dedicated servers in a datacenter that do not have a central administration console and you want to manage them via an app
- You have metrics that you want to observe but don't have the time or the need for a big tool like Grafana

Of course, these are only examples and the only limit is your imagination.

## Getting Started

<p align="center">
  <img width="100%" src="./docs/endpoints.png">
</p>

The spec follows the [OpenAPI Specification](https://swagger.io/specification) allowing a high level of compatibility with existing tools.

You can play with it by loading it locally in [SwaggerUI](https://swagger.io/tools/swagger-ui) with [Docker Compose](https://docs.docker.com/compose):

```sh
# Generate swagger.json (optional since it's already present in the repository)
docker run --rm -v $(pwd):/app oven/bun run /app/spec/generate-swagger.ts
pnpm lint

# Generate Go code with OpenAPI Generator
(cd impl/http-server-go && go mod vendor)
docker run --rm \
-v $(pwd):/app \
openapitools/openapi-generator-cli:v7.4.0 generate \
-i /app/spec/_generated/swagger.json \
-g go \
-o /app/impl/http-server-go/vendor/openapi

# Run Swagger UI calling the http-server-go (see below for more details)
docker compose up
```

You can then access http://localhost:9002 via your browser and test the endpoints. See `docker-compose.yml` to have the `apiKey`.

You can also directly test the server with cURL:

```sh
curl -v -H "Authorization: <apiKey>" http://localhost:9001/cd5331ba/runnables
curl -v -X POST -H "Authorization: <apiKey>" http://localhost:9001/cd5331ba/runnables/123/reboot
curl -v -X POST -H "Authorization: <apiKey>" http://localhost:9001/cd5331ba/runnables/123/stop
curl -v -H "Authorization: <apiKey>" http://localhost:9001/cd5331ba/dashboards
```

## Creating your own server

You can develop your own server to fit your personal needs and keep control on everything that runs on your infra. It is also possible to add endpoints to any existing REST API.

As long as you respect the specification, you can use the language of your choice. To speed things up, you can generate some code using [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator).

For instance, you can generate [Rust](https://www.rust-lang.org) code with the following command :

```sh
docker run --rm \
-v $(pwd):/app \
openapitools/openapi-generator-cli:v7.4.0 generate \
-i /app/spec/_generated/swagger.json \
-g rust \
-o /app/impl/http-server-rust/openapi
```

Although it can be handful, we do not recommend relying on all the generated code for a **production server**, because it contains too much boilerplate, making it harder to maintain. It's fine to use the generated `structs`, `interfaces`, `enums`, `traits`, though. That being said, it's up to you.

Once ready, your server should get your _runnables_ and _dashboards_ from a local file, another API, network calls or a database (remote or local)... Once again, it's up to you.

It should be deployed to an independent location, accessible by the phone you want to use : Internet if you want access from anywhere / LAN if local access is enough for your use case.

Of course, except in some cases, you must serve it with `HTTPS` to secure the connection between the app and your server.

## Existing servers

### HTTP Server Go

This is the implementation used in the `docker-compose` command above, made in Go. You can browse the code at [impl/http-server-go](./impl/http-server-go).

If you don't want or cannot build your own server, you can use this one.

You can download the latest release from the [releases](https://github.com/c100k/rebootx-on-prem/releases) page or build it yourself from source.

The server can run in different modes depending on your use case :

- **Runnables**
  - `self` : it returns the host as a _runnable_. Be careful if you run this on a machine as a privileged user. It relies on [syscall](https://pkg.go.dev/syscall) and [exec](https://pkg.go.dev/os/exec) so it can actually `reboot` or `stop` the machine for real
  - `fileJson` (default) : it reads the _runnables_ from a JSON file that must respect the schema in order to be unmarshalled into an array of `Runnable` (see [servers.example.json](./data/servers.example.json))
- **Dashboards**
  - `fileJson` (default) : it reads the dashboards from a JSON file that must respect the schema in order to be unmarshalled into an array of `Dashboard` (see [dashboards.example.json](./data/dashboards.example.json))

To override the default behavior, see `docker-compose.yml` or `config.go` and update the appropriate environment variables accordingly. 

## Contributing

We welcome all kind of contributions. Feel free to reach out to us or to publish a PR.

If you want to develop an implementation in a specific language, it must be clean, follow good practice and be self contained in its own directory in [./impl](./impl).

In any case, please read the [CONTRIBUTING.md](./CONTRIBUTING.md) guide first.
