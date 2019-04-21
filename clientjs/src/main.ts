import * as secp from './lib/secp256k1'
import { randomBytes } from 'crypto'
import * as fs from 'fs'
import { protobuf } from './models/proto'

const util = require('util');
const fetch = require('node-fetch');







//Generate public Key

let dataPath = "../testdata/secp205k1Accts/"

interface Key {
    type: string
    value: string
}


let LoakKeys = async (path: string): Promise<any> => {
    const readFile = util.promisify(fs.readFile);
    return await readFile(path)
}


let getKeys = (privateKey64: Key) => {
    return new secp.Secp256k1Generator(privateKey64.value)
}




let start = async () => {
    let endpoint = "http://" + "localhost" + ":80/tx"

    let senderData = await LoakKeys(dataPath + "1_peer_id.json")

    let privD = JSON.parse(senderData.toString()).priv_key
    let sender: Key = { type: privD.type, value: privD.value }
    let sendkerKey = getKeys(sender)


    let recieverData = await LoakKeys(dataPath + "2_peer_id.json")
    let rec = JSON.parse(recieverData.toString()).priv_key

    let reciever: Key = { type: rec.type, value: rec.value }
    let recieverKey = getKeys(reciever)

    let asset: protobuf.IAsset = { category: "crypto", symbol: "HER", network: "Herdius", value: 100, fee: 0, nonce: 0 }

    let tx: protobuf.ITx = { senderAddress: sendkerKey.getAddress(), recieverAddress: recieverKey.getAddress(), asset: asset, message: "Send Her Token", senderPubkey: sendkerKey.getPublicKey().toString('base64') }
 
    let signedTx = sendkerKey.sign(Buffer.from(JSON.stringify(tx)))
    tx.sign = signedTx.signature.toString('base64')

    let req: protobuf.ITxRequest = { tx: tx }

 
    let request = await fetch(endpoint, {
        method: "POST",
        body: JSON.stringify(req),
        headers: {
            "Content-Type": "application/json"
        },
    })

    let d = await request.json()
    console.log(d)

}


start();

