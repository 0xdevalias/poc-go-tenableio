# poc-go-tenable.io

PoC/partial implementation of the [Tenable.io API](https://cloud.tenable.com/api) as a Golang client library.

I figured [Tenable.io](https://cloud.tenable.com/api) had a [Python SDK](https://github.com/tenable/Tenable.io-SDK-for-Python), and a [Java SDK](https://github.com/tenable/Tenable.io-SDK-for-Java), so it should have a partially implemented Go SDK as well.. because reasons.

## Usage

```go
import tio "github.com/0xdevalias/poc-go-tenableio/api

c := api.DefaultClient(
		"TODO-ACCESS-KEY",
		"TODO-SECRET-KEY",
	)
	
assets, _ := c.Workbenches.Assets()
```
