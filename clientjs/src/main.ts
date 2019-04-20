import * as secp from './lib/secp256k1'
import {randomBytes} from 'crypto'

//Generate public Key


const msg = randomBytes(32)
const msg1 = randomBytes(32)


let privateKey = secp.genPrivateKey(msg)
let PublicKey = secp.genPublicKey(privateKey)
console.log(privateKey)
console.log(PublicKey)
let signedObj = secp.sign(msg,privateKey)
console.log(signedObj)
let verify = secp.verify(msg,signedObj,PublicKey)
console.log(verify)
verify = secp.verify(msg1,signedObj,PublicKey)
console.log(verify)




