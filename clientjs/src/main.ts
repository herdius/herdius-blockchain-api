import * as secp from './lib/secp256k1'
import * as fs from 'fs'
const sha256 = require('sha256');


const util = require('util');
const fetch = require('node-fetch');

interface ITxRequest {
    /** TxRequest tx */
    tx?: (ITx | null);
}

interface ITx {

    /** Tx senderAddress */
    sender_address?: (string | null);

    /** Tx senderPubkey */
    sender_pubkey?: (string | null);

    /** Tx recieverAddress */
    reciever_address?: (string | null);

    /** Tx asset */
    asset?: (IAsset | null);

    /** Tx message */
    message?: (string | null);

    /** Tx sign */
    sign?: (string | null);

    /** Tx type */
    type?: (string | null);

    /** Tx status */
    status?: (string | null);
}

interface IAsset {

    /** Asset category */
    category?: (string | null);

    /** Asset symbol */
    symbol?: (string | null);

    /** Asset network */
    network?: (string | null);

    /** Asset value */
    value?: (number | null);

    /** Asset fee */
    fee?: (number | null);

    /** Asset nonce */
    nonce?: (number | null);
}

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


let sendTx = async () => {
    let endpoint = "http://" + "127.0.0.1" + ":80/tx"
    let senderData = await LoakKeys(dataPath + "1_peer_id.json")

    let privD = JSON.parse(senderData.toString()).priv_key
    let sender: Key = { type: privD.type, value: privD.value }
    let sendkerKey = getKeys(sender)


    let recieverData = await LoakKeys(dataPath + "2_peer_id.json")
    let rec = JSON.parse(recieverData.toString()).priv_key

    let reciever: Key = { type: rec.type, value: rec.value }
    let recieverKey = getKeys(reciever)

    let asset: IAsset = {
        category: "crypto",
        symbol: "HER",
        network: "Herdius",
        value: 100,
        fee: 1,
        nonce:10
        }

    let tx:any = {
        sender_address: sendkerKey.getAddress(),
        sender_pubkey: sendkerKey.getPublicKey().toString('base64'),
        reciever_address: recieverKey.getAddress(),
        asset: asset,
        message: "Send Her Token",
    }

    const msg = Buffer.from(sha256(JSON.stringify(tx), { asBytes: true }));

    const signedTx = sendkerKey.sign(msg);
    tx.sign = signedTx.signature.toString('base64')


    let req: ITxRequest = { tx: tx }

    let request = await fetch(endpoint, {
        method: "POST",
        body: JSON.stringify(req),
        headers: {
            "Content-Type": "application/json"
        }
    })

    let result = await request.json()
    console.log(result)

}


sendTx();

