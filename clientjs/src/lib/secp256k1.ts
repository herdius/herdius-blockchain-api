import * as crypto from "crypto";
import secp256k1 from 'secp256k1';
const sha3_256 = require('js-sha3').sha3_256;
const bs58 = require('bs58');
const sha256 = require('sha256');
const toBuffer = require('typedarray-to-buffer');
const atob = require('atob');

import * as ethUtils from 'ethereumjs-util';




interface IGenerator {
  getPrivateKey(): Buffer;
  getPublicKey(privateKey: any): any;
  sign(msg: Buffer, privateKey: any): any;
  verify(msg: any, signObj: any, pubKey: any): any;
}

class Secp256k1Generator implements IGenerator {

  private _key: string
  private _privKey!: Buffer;

  constructor(key: string) {
    console.log(key)
    this._key = key
    let pkBuffer = base64ToArrayBuffer(key);
    this._privKey = pkBuffer
  }
  getPrivateKey = (): Buffer => {
    let pkBuffer = base64ToArrayBuffer(this._key);
    this._privKey = pkBuffer
    return this._privKey
  }

  setPrivKey = (key: string): Buffer => {
    let pkBuffer = base64ToArrayBuffer(key);
    this._privKey = pkBuffer
    return this._privKey
  }

  getPublicKey = (): Buffer => {
    return secp256k1.publicKeyCreate(this._privKey)
  }

  getAddress = (): any => {
    console.log(this._privKey)
    let pubKey = secp256k1.publicKeyCreate(this._privKey)
    let hash = sha3_256(pubKey);
    const hash20 = new Uint8Array(hexToArrayBuffer(hash).slice(12));
    let append40 = new Uint8Array(1);
    append40[0] = 40;
    append40 = concatTypedArrays(append40, hash20);

    const doubleHash = sha256.x2(append40, { asBytes: true });
    const checksum = doubleHash.slice(0, 4);
    const address = concatTypedArrays(append40, checksum);
    let senderaddress = bs58.encode(toBuffer(address));
    return senderaddress
  }

  sign = (msg: any): any => {
    // let mm = crypto.createHash("sha256").update(msg).digest();
    // console.log("-------hash-------",mm)

    return secp256k1.sign(msg, this.getPrivateKey())
    // let gsigned = secp256k1.signatureNormalize(signed.signature);

    // Convert to DER
    // return secp256k1.signatureExport(gsigned);
    //  const mm = ethUtils.hashPersonalMessage(ethUtils.toBuffer(msg));
    //     const signed = ethUtils.ecsign(
    //       Buffer.from(mm),
    //       this._privKey
    //     );
    //     const combined = Buffer.concat([
    //       Buffer.from([signed.v]),
    //       Buffer.from(signed.r),
    //       Buffer.from(signed.s)
    //     ]);
    //     const combinedHex = combined.toString('base64');
    //     return combinedHex
  }

  verify = (msg: Buffer, signObj: any): any => {
    return secp256k1.verify(msg, signObj, this.getPublicKey())
  }
}

function hexToArrayBuffer(hex: any) {
  const typedArray = new Uint8Array(hex.match(/[\da-f]{2}/gi).map(function (h: any) {
    return parseInt(h, 16);
  }));
  return typedArray.buffer;
}

function concatTypedArrays(a: any, b: any) {
  let c = new (a.constructor)(a.length + b.length);
  c.set(a, 0);
  c.set(b, a.length);
  return c;
}

function base64ToArrayBuffer(base64: any) {
  const binary_string = atob(base64);
  const len = binary_string.length;
  const bytes = new Uint8Array(len);
  for (let i = 0; i < len; i++) {
    bytes[i] = binary_string.charCodeAt(i);
  }
  return toBuffer(bytes.buffer);
}




export { Secp256k1Generator }
