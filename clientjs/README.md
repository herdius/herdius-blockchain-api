## Intent

JS Herdius Client 

## Installation

```
cd clientjs
yarn install

```

#### Send Transaction
```
yarn start

```
### secp256k1 Library

To use secp256k1 in existing code

1) import * as secp from './lib/secp256k1'
2) Create secp object
```
    let secp256 =  new secp.Secp256k1Generator(randomValue)
```
3) Get Private key and Public key
```
let privateKey = secp256.getPrivateKey()
let publicKey = secp256.getPublicKey()

```

4) Sign Message

```
let messageObj = {
    ....
}

Sign needs sha256 message 

const msg = Buffer.from(sha256(JSON.stringify(messageStruct), { asBytes: true }));

let signed = secp256.sign(msg)

```

