## Intent

The Herdius Blockchain API provides an API interface in which end users can interract with both the core [Herdius Blockchain](https://herdius.com), as well as it's various test chains. At the moment, a fully functional chain is still on its way to deployment, thus only test chains are available for API interraction.

The API node(s) act as a [DMZ](https://searchsecurity.techtarget.com/definition/DMZ) for the secured Herdius Cluster (consisting of the Supervisor and Validator nodes) which, in the case of deployed test chains, operate in a fully private subnet. Thus, DDoS and similar malicious attacks will disrupt service at the API level, but the chains themselves will still function in a completely undisrupted way.

## Usage

#### `POST` new transaction

#### `GET` account details

#### `GET` transaction details

#### `GET` block details
