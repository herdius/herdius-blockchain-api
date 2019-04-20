import { randomBytes } from "crypto";
import secp256k1 from 'secp256k1';

let genPrivateKey = (random: Buffer): any => {
  let privKey
  do {
    privKey = randomBytes(32)
  } while (!secp256k1.privateKeyVerify(privKey))
  return privKey
}

let genPublicKey = (privateKey: any): any => {
  return secp256k1.publicKeyCreate(privateKey)
}

let sign = (msg: any, privKey: any): any => {
  return secp256k1.sign(msg, privKey)
}

let verify = (msg: any, signObj: any, pubKey: any): any => {
  return secp256k1.verify(msg, signObj.signature, pubKey)
}


export { genPrivateKey, genPublicKey, sign, verify }
