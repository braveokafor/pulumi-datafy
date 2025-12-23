# Datafy Resource Provider

The Datafy Resource Provider lets you manage [Datafy](https://www.datafy.com) resources.

## Building

```bash
make build              # provider + SDKs
make provider           # provider only
```

Default version: `1.0.0-alpha.0+dev` (override: `PROVIDER_VERSION=x.y.z`)

## Installing


### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/braveokafor/pulumi-datafy/sdk/go/datafy
```

Then import in your code:

```go
import "github.com/braveokafor/pulumi-datafy/sdk/go/datafy"
```

### Manual Plugin Installation

For Automation API:

```bash
cd provider
mkdir -p ~/.pulumi/plugins/resource-datafy-v1.0.0-alpha.0+dev
go build -o ~/.pulumi/plugins/resource-datafy-v1.0.0-alpha.0+dev/pulumi-resource-datafy \
  -ldflags "-X github.com/braveokafor/pulumi-datafy/provider/pkg/version.Version=1.0.0-alpha.0+dev -s -w" \
  ./cmd/pulumi-resource-datafy
```

## Configuration

`datafy:token` - API token (env: `DATAFY_TOKEN`)

See [pkg.go.dev](https://pkg.go.dev/github.com/braveokafor/pulumi-datafy/sdk)

## Reference
