## Intent

The Herdius Blockchain API provides an API interface in which end users can interract with both the core [Herdius Blockchain](https://herdius.com), as well as it's various test chains. At the moment, a fully functional chain is still on its way to deployment, thus only test chains are available for API interraction.

The API node(s) act as a [DMZ](https://searchsecurity.techtarget.com/definition/DMZ) for the secured Herdius Cluster (consisting of the Supervisor and Validator nodes) which, in the case of deployed test chains, operate in a fully private subnet. Thus, DDoS and similar malicious attacks will disrupt service at the API level, but the chains themselves will still function in a completely undisrupted way.

## Usage

#### `POST` new transaction

**Method 1**: A helper package found in `client.go` is created to simplify this process. To `POST` a new transaction to a Herdius Blockchain API node, simply run `go run client.go [API endpoint]` in which the `[API endpoint]` is the IP or DNS address of the API node of the chain to send the transaction to. The default address is `localhost`, meaning that you have an API node listening on your host server/computer.

Example:

```
// Send transaction to primary Herdius test chain (with API node IP address at `3.209.249.184`)
go run client.go 3.209.249.184

// Send transaction to host's API node
go run client.go
go run client.go localhost // equivalent to ^
```

**Method 2**: Manually `POST` to API node

TODO

#### `GET` account details

#### `GET` transaction details

#### `GET` block details

## Versions and Testing

A persistent Herdius test chain lives behind a series of API nodes for which the public can openly test transactions and account interractions. These API nodes are integrated with a CICD pipeline which deploys any new commits to the `master` git branch as they are merged.

The test chain is in a state of transation as we continue to strengthen and reshape the chain structure and logic. However, the API nodes are stable for interraction. Because the chain is in a state of flux, the API nodes may periodically return unexpected results to the calling client.
