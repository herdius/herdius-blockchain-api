import * as $protobuf from "protobufjs";
/** Namespace protobuf. */
export namespace protobuf {

    /** Properties of a Timestamp. */
    interface ITimestamp {

        /** Timestamp seconds */
        seconds?: (number|Long|null);

        /** Timestamp nanos */
        nanos?: (number|Long|null);
    }

    /** Represents a Timestamp. */
    class Timestamp implements ITimestamp {

        /**
         * Constructs a new Timestamp.
         * @param [properties] Properties to set
         */
        constructor(properties?: protobuf.ITimestamp);

        /** Timestamp seconds. */
        public seconds: (number|Long);

        /** Timestamp nanos. */
        public nanos: (number|Long);

        /**
         * Creates a new Timestamp instance using the specified properties.
         * @param [properties] Properties to set
         * @returns Timestamp instance
         */
        public static create(properties?: protobuf.ITimestamp): protobuf.Timestamp;

        /**
         * Encodes the specified Timestamp message. Does not implicitly {@link protobuf.Timestamp.verify|verify} messages.
         * @param message Timestamp message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: protobuf.ITimestamp, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Timestamp message, length delimited. Does not implicitly {@link protobuf.Timestamp.verify|verify} messages.
         * @param message Timestamp message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: protobuf.ITimestamp, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Timestamp message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Timestamp
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): protobuf.Timestamp;

        /**
         * Decodes a Timestamp message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns Timestamp
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): protobuf.Timestamp;

        /**
         * Verifies a Timestamp message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a Timestamp message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Timestamp
         */
        public static fromObject(object: { [k: string]: any }): protobuf.Timestamp;

        /**
         * Creates a plain object from a Timestamp message. Also converts values to other types if specified.
         * @param message Timestamp
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: protobuf.Timestamp, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Timestamp to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of a BlockHeightRequest. */
    interface IBlockHeightRequest {

        /** BlockHeightRequest blockHeight */
        blockHeight?: (number|Long|null);
    }

    /** Represents a BlockHeightRequest. */
    class BlockHeightRequest implements IBlockHeightRequest {

        /**
         * Constructs a new BlockHeightRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: protobuf.IBlockHeightRequest);

        /** BlockHeightRequest blockHeight. */
        public blockHeight: (number|Long);

        /**
         * Creates a new BlockHeightRequest instance using the specified properties.
         * @param [properties] Properties to set
         * @returns BlockHeightRequest instance
         */
        public static create(properties?: protobuf.IBlockHeightRequest): protobuf.BlockHeightRequest;

        /**
         * Encodes the specified BlockHeightRequest message. Does not implicitly {@link protobuf.BlockHeightRequest.verify|verify} messages.
         * @param message BlockHeightRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: protobuf.IBlockHeightRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified BlockHeightRequest message, length delimited. Does not implicitly {@link protobuf.BlockHeightRequest.verify|verify} messages.
         * @param message BlockHeightRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: protobuf.IBlockHeightRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a BlockHeightRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns BlockHeightRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): protobuf.BlockHeightRequest;

        /**
         * Decodes a BlockHeightRequest message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns BlockHeightRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): protobuf.BlockHeightRequest;

        /**
         * Verifies a BlockHeightRequest message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a BlockHeightRequest message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns BlockHeightRequest
         */
        public static fromObject(object: { [k: string]: any }): protobuf.BlockHeightRequest;

        /**
         * Creates a plain object from a BlockHeightRequest message. Also converts values to other types if specified.
         * @param message BlockHeightRequest
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: protobuf.BlockHeightRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this BlockHeightRequest to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of a BlockResponse. */
    interface IBlockResponse {

        /** BlockResponse blockHeight */
        blockHeight?: (number|Long|null);

        /** BlockResponse time */
        time?: (protobuf.ITimestamp|null);

        /** BlockResponse transactions */
        transactions?: (number|Long|null);

        /** BlockResponse supervisorAddress */
        supervisorAddress?: (string|null);

        /** BlockResponse stateRoot */
        stateRoot?: (Uint8Array|null);
    }

    /** Represents a BlockResponse. */
    class BlockResponse implements IBlockResponse {

        /**
         * Constructs a new BlockResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: protobuf.IBlockResponse);

        /** BlockResponse blockHeight. */
        public blockHeight: (number|Long);

        /** BlockResponse time. */
        public time?: (protobuf.ITimestamp|null);

        /** BlockResponse transactions. */
        public transactions: (number|Long);

        /** BlockResponse supervisorAddress. */
        public supervisorAddress: string;

        /** BlockResponse stateRoot. */
        public stateRoot: Uint8Array;

        /**
         * Creates a new BlockResponse instance using the specified properties.
         * @param [properties] Properties to set
         * @returns BlockResponse instance
         */
        public static create(properties?: protobuf.IBlockResponse): protobuf.BlockResponse;

        /**
         * Encodes the specified BlockResponse message. Does not implicitly {@link protobuf.BlockResponse.verify|verify} messages.
         * @param message BlockResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: protobuf.IBlockResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified BlockResponse message, length delimited. Does not implicitly {@link protobuf.BlockResponse.verify|verify} messages.
         * @param message BlockResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: protobuf.IBlockResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a BlockResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns BlockResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): protobuf.BlockResponse;

        /**
         * Decodes a BlockResponse message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns BlockResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): protobuf.BlockResponse;

        /**
         * Verifies a BlockResponse message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a BlockResponse message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns BlockResponse
         */
        public static fromObject(object: { [k: string]: any }): protobuf.BlockResponse;

        /**
         * Creates a plain object from a BlockResponse message. Also converts values to other types if specified.
         * @param message BlockResponse
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: protobuf.BlockResponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this BlockResponse to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of an AccountRequest. */
    interface IAccountRequest {

        /** AccountRequest address */
        address?: (string|null);
    }

    /** Represents an AccountRequest. */
    class AccountRequest implements IAccountRequest {

        /**
         * Constructs a new AccountRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: protobuf.IAccountRequest);

        /** AccountRequest address. */
        public address: string;

        /**
         * Creates a new AccountRequest instance using the specified properties.
         * @param [properties] Properties to set
         * @returns AccountRequest instance
         */
        public static create(properties?: protobuf.IAccountRequest): protobuf.AccountRequest;

        /**
         * Encodes the specified AccountRequest message. Does not implicitly {@link protobuf.AccountRequest.verify|verify} messages.
         * @param message AccountRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: protobuf.IAccountRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified AccountRequest message, length delimited. Does not implicitly {@link protobuf.AccountRequest.verify|verify} messages.
         * @param message AccountRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: protobuf.IAccountRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an AccountRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns AccountRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): protobuf.AccountRequest;

        /**
         * Decodes an AccountRequest message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns AccountRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): protobuf.AccountRequest;

        /**
         * Verifies an AccountRequest message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates an AccountRequest message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns AccountRequest
         */
        public static fromObject(object: { [k: string]: any }): protobuf.AccountRequest;

        /**
         * Creates a plain object from an AccountRequest message. Also converts values to other types if specified.
         * @param message AccountRequest
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: protobuf.AccountRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this AccountRequest to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of an AccountResponse. */
    interface IAccountResponse {

        /** AccountResponse address */
        address?: (string|null);

        /** AccountResponse nonce */
        nonce?: (number|Long|null);

        /** AccountResponse storageRoot */
        storageRoot?: (string|null);

        /** AccountResponse publicKey */
        publicKey?: (string|null);

        /** AccountResponse balance */
        balance?: (number|Long|null);

        /** AccountResponse balances */
        balances?: ({ [k: string]: (number|Long) }|null);
    }

    /** Represents an AccountResponse. */
    class AccountResponse implements IAccountResponse {

        /**
         * Constructs a new AccountResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: protobuf.IAccountResponse);

        /** AccountResponse address. */
        public address: string;

        /** AccountResponse nonce. */
        public nonce: (number|Long);

        /** AccountResponse storageRoot. */
        public storageRoot: string;

        /** AccountResponse publicKey. */
        public publicKey: string;

        /** AccountResponse balance. */
        public balance: (number|Long);

        /** AccountResponse balances. */
        public balances: { [k: string]: (number|Long) };

        /**
         * Creates a new AccountResponse instance using the specified properties.
         * @param [properties] Properties to set
         * @returns AccountResponse instance
         */
        public static create(properties?: protobuf.IAccountResponse): protobuf.AccountResponse;

        /**
         * Encodes the specified AccountResponse message. Does not implicitly {@link protobuf.AccountResponse.verify|verify} messages.
         * @param message AccountResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: protobuf.IAccountResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified AccountResponse message, length delimited. Does not implicitly {@link protobuf.AccountResponse.verify|verify} messages.
         * @param message AccountResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: protobuf.IAccountResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an AccountResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns AccountResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): protobuf.AccountResponse;

        /**
         * Decodes an AccountResponse message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns AccountResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): protobuf.AccountResponse;

        /**
         * Verifies an AccountResponse message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates an AccountResponse message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns AccountResponse
         */
        public static fromObject(object: { [k: string]: any }): protobuf.AccountResponse;

        /**
         * Creates a plain object from an AccountResponse message. Also converts values to other types if specified.
         * @param message AccountResponse
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: protobuf.AccountResponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this AccountResponse to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of a Tx. */
    interface ITx {

        /** Tx senderAddress */
        senderAddress?: (string|null);

        /** Tx senderPubkey */
        senderPubkey?: (string|null);

        /** Tx recieverAddress */
        recieverAddress?: (string|null);

        /** Tx asset */
        asset?: (protobuf.IAsset|null);

        /** Tx message */
        message?: (string|null);

        /** Tx sign */
        sign?: (string|null);

        /** Tx type */
        type?: (string|null);

        /** Tx status */
        status?: (string|null);
    }

    /** Represents a Tx. */
    class Tx implements ITx {

        /**
         * Constructs a new Tx.
         * @param [properties] Properties to set
         */
        constructor(properties?: protobuf.ITx);

        /** Tx senderAddress. */
        public senderAddress: string;

        /** Tx senderPubkey. */
        public senderPubkey: string;

        /** Tx recieverAddress. */
        public recieverAddress: string;

        /** Tx asset. */
        public asset?: (protobuf.IAsset|null);

        /** Tx message. */
        public message: string;

        /** Tx sign. */
        public sign: string;

        /** Tx type. */
        public type: string;

        /** Tx status. */
        public status: string;

        /**
         * Creates a new Tx instance using the specified properties.
         * @param [properties] Properties to set
         * @returns Tx instance
         */
        public static create(properties?: protobuf.ITx): protobuf.Tx;

        /**
         * Encodes the specified Tx message. Does not implicitly {@link protobuf.Tx.verify|verify} messages.
         * @param message Tx message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: protobuf.ITx, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Tx message, length delimited. Does not implicitly {@link protobuf.Tx.verify|verify} messages.
         * @param message Tx message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: protobuf.ITx, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Tx message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Tx
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): protobuf.Tx;

        /**
         * Decodes a Tx message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns Tx
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): protobuf.Tx;

        /**
         * Verifies a Tx message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a Tx message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Tx
         */
        public static fromObject(object: { [k: string]: any }): protobuf.Tx;

        /**
         * Creates a plain object from a Tx message. Also converts values to other types if specified.
         * @param message Tx
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: protobuf.Tx, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Tx to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of an Asset. */
    interface IAsset {

        /** Asset category */
        category?: (string|null);

        /** Asset symbol */
        symbol?: (string|null);

        /** Asset network */
        network?: (string|null);

        /** Asset value */
        value?: (number|Long|null);

        /** Asset fee */
        fee?: (number|Long|null);

        /** Asset nonce */
        nonce?: (number|Long|null);
    }

    /** Represents an Asset. */
    class Asset implements IAsset {

        /**
         * Constructs a new Asset.
         * @param [properties] Properties to set
         */
        constructor(properties?: protobuf.IAsset);

        /** Asset category. */
        public category: string;

        /** Asset symbol. */
        public symbol: string;

        /** Asset network. */
        public network: string;

        /** Asset value. */
        public value: (number|Long);

        /** Asset fee. */
        public fee: (number|Long);

        /** Asset nonce. */
        public nonce: (number|Long);

        /**
         * Creates a new Asset instance using the specified properties.
         * @param [properties] Properties to set
         * @returns Asset instance
         */
        public static create(properties?: protobuf.IAsset): protobuf.Asset;

        /**
         * Encodes the specified Asset message. Does not implicitly {@link protobuf.Asset.verify|verify} messages.
         * @param message Asset message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: protobuf.IAsset, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Asset message, length delimited. Does not implicitly {@link protobuf.Asset.verify|verify} messages.
         * @param message Asset message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: protobuf.IAsset, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an Asset message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Asset
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): protobuf.Asset;

        /**
         * Decodes an Asset message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns Asset
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): protobuf.Asset;

        /**
         * Verifies an Asset message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates an Asset message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Asset
         */
        public static fromObject(object: { [k: string]: any }): protobuf.Asset;

        /**
         * Creates a plain object from an Asset message. Also converts values to other types if specified.
         * @param message Asset
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: protobuf.Asset, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Asset to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of a TxRequest. */
    interface ITxRequest {

        /** TxRequest tx */
        tx?: (protobuf.ITx|null);
    }

    /** Represents a TxRequest. */
    class TxRequest implements ITxRequest {

        /**
         * Constructs a new TxRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: protobuf.ITxRequest);

        /** TxRequest tx. */
        public tx?: (protobuf.ITx|null);

        /**
         * Creates a new TxRequest instance using the specified properties.
         * @param [properties] Properties to set
         * @returns TxRequest instance
         */
        public static create(properties?: protobuf.ITxRequest): protobuf.TxRequest;

        /**
         * Encodes the specified TxRequest message. Does not implicitly {@link protobuf.TxRequest.verify|verify} messages.
         * @param message TxRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: protobuf.ITxRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified TxRequest message, length delimited. Does not implicitly {@link protobuf.TxRequest.verify|verify} messages.
         * @param message TxRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: protobuf.ITxRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a TxRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns TxRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): protobuf.TxRequest;

        /**
         * Decodes a TxRequest message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns TxRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): protobuf.TxRequest;

        /**
         * Verifies a TxRequest message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a TxRequest message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns TxRequest
         */
        public static fromObject(object: { [k: string]: any }): protobuf.TxRequest;

        /**
         * Creates a plain object from a TxRequest message. Also converts values to other types if specified.
         * @param message TxRequest
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: protobuf.TxRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this TxRequest to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of a TxResponse. */
    interface ITxResponse {

        /** TxResponse txId */
        txId?: (string|null);

        /** TxResponse pending */
        pending?: (number|Long|null);

        /** TxResponse queued */
        queued?: (number|Long|null);

        /** TxResponse status */
        status?: (string|null);
    }

    /** Represents a TxResponse. */
    class TxResponse implements ITxResponse {

        /**
         * Constructs a new TxResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: protobuf.ITxResponse);

        /** TxResponse txId. */
        public txId: string;

        /** TxResponse pending. */
        public pending: (number|Long);

        /** TxResponse queued. */
        public queued: (number|Long);

        /** TxResponse status. */
        public status: string;

        /**
         * Creates a new TxResponse instance using the specified properties.
         * @param [properties] Properties to set
         * @returns TxResponse instance
         */
        public static create(properties?: protobuf.ITxResponse): protobuf.TxResponse;

        /**
         * Encodes the specified TxResponse message. Does not implicitly {@link protobuf.TxResponse.verify|verify} messages.
         * @param message TxResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: protobuf.ITxResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified TxResponse message, length delimited. Does not implicitly {@link protobuf.TxResponse.verify|verify} messages.
         * @param message TxResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: protobuf.ITxResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a TxResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns TxResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): protobuf.TxResponse;

        /**
         * Decodes a TxResponse message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns TxResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): protobuf.TxResponse;

        /**
         * Verifies a TxResponse message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a TxResponse message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns TxResponse
         */
        public static fromObject(object: { [k: string]: any }): protobuf.TxResponse;

        /**
         * Creates a plain object from a TxResponse message. Also converts values to other types if specified.
         * @param message TxResponse
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: protobuf.TxResponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this TxResponse to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of an AccountRegisterRequest. */
    interface IAccountRegisterRequest {

        /** AccountRegisterRequest senderPubkey */
        senderPubkey?: (string|null);
    }

    /** Represents an AccountRegisterRequest. */
    class AccountRegisterRequest implements IAccountRegisterRequest {

        /**
         * Constructs a new AccountRegisterRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: protobuf.IAccountRegisterRequest);

        /** AccountRegisterRequest senderPubkey. */
        public senderPubkey: string;

        /**
         * Creates a new AccountRegisterRequest instance using the specified properties.
         * @param [properties] Properties to set
         * @returns AccountRegisterRequest instance
         */
        public static create(properties?: protobuf.IAccountRegisterRequest): protobuf.AccountRegisterRequest;

        /**
         * Encodes the specified AccountRegisterRequest message. Does not implicitly {@link protobuf.AccountRegisterRequest.verify|verify} messages.
         * @param message AccountRegisterRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: protobuf.IAccountRegisterRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified AccountRegisterRequest message, length delimited. Does not implicitly {@link protobuf.AccountRegisterRequest.verify|verify} messages.
         * @param message AccountRegisterRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: protobuf.IAccountRegisterRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an AccountRegisterRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns AccountRegisterRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): protobuf.AccountRegisterRequest;

        /**
         * Decodes an AccountRegisterRequest message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns AccountRegisterRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): protobuf.AccountRegisterRequest;

        /**
         * Verifies an AccountRegisterRequest message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates an AccountRegisterRequest message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns AccountRegisterRequest
         */
        public static fromObject(object: { [k: string]: any }): protobuf.AccountRegisterRequest;

        /**
         * Creates a plain object from an AccountRegisterRequest message. Also converts values to other types if specified.
         * @param message AccountRegisterRequest
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: protobuf.AccountRegisterRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this AccountRegisterRequest to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of an AccountRegisterResponse. */
    interface IAccountRegisterResponse {

        /** AccountRegisterResponse status */
        status?: (boolean|null);
    }

    /** Represents an AccountRegisterResponse. */
    class AccountRegisterResponse implements IAccountRegisterResponse {

        /**
         * Constructs a new AccountRegisterResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: protobuf.IAccountRegisterResponse);

        /** AccountRegisterResponse status. */
        public status: boolean;

        /**
         * Creates a new AccountRegisterResponse instance using the specified properties.
         * @param [properties] Properties to set
         * @returns AccountRegisterResponse instance
         */
        public static create(properties?: protobuf.IAccountRegisterResponse): protobuf.AccountRegisterResponse;

        /**
         * Encodes the specified AccountRegisterResponse message. Does not implicitly {@link protobuf.AccountRegisterResponse.verify|verify} messages.
         * @param message AccountRegisterResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: protobuf.IAccountRegisterResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified AccountRegisterResponse message, length delimited. Does not implicitly {@link protobuf.AccountRegisterResponse.verify|verify} messages.
         * @param message AccountRegisterResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: protobuf.IAccountRegisterResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an AccountRegisterResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns AccountRegisterResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): protobuf.AccountRegisterResponse;

        /**
         * Decodes an AccountRegisterResponse message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns AccountRegisterResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): protobuf.AccountRegisterResponse;

        /**
         * Verifies an AccountRegisterResponse message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates an AccountRegisterResponse message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns AccountRegisterResponse
         */
        public static fromObject(object: { [k: string]: any }): protobuf.AccountRegisterResponse;

        /**
         * Creates a plain object from an AccountRegisterResponse message. Also converts values to other types if specified.
         * @param message AccountRegisterResponse
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: protobuf.AccountRegisterResponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this AccountRegisterResponse to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of a TxDetailRequest. */
    interface ITxDetailRequest {

        /** TxDetailRequest txId */
        txId?: (string|null);
    }

    /** Represents a TxDetailRequest. */
    class TxDetailRequest implements ITxDetailRequest {

        /**
         * Constructs a new TxDetailRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: protobuf.ITxDetailRequest);

        /** TxDetailRequest txId. */
        public txId: string;

        /**
         * Creates a new TxDetailRequest instance using the specified properties.
         * @param [properties] Properties to set
         * @returns TxDetailRequest instance
         */
        public static create(properties?: protobuf.ITxDetailRequest): protobuf.TxDetailRequest;

        /**
         * Encodes the specified TxDetailRequest message. Does not implicitly {@link protobuf.TxDetailRequest.verify|verify} messages.
         * @param message TxDetailRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: protobuf.ITxDetailRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified TxDetailRequest message, length delimited. Does not implicitly {@link protobuf.TxDetailRequest.verify|verify} messages.
         * @param message TxDetailRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: protobuf.ITxDetailRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a TxDetailRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns TxDetailRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): protobuf.TxDetailRequest;

        /**
         * Decodes a TxDetailRequest message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns TxDetailRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): protobuf.TxDetailRequest;

        /**
         * Verifies a TxDetailRequest message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a TxDetailRequest message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns TxDetailRequest
         */
        public static fromObject(object: { [k: string]: any }): protobuf.TxDetailRequest;

        /**
         * Creates a plain object from a TxDetailRequest message. Also converts values to other types if specified.
         * @param message TxDetailRequest
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: protobuf.TxDetailRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this TxDetailRequest to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of a TxDetailResponse. */
    interface ITxDetailResponse {

        /** TxDetailResponse txId */
        txId?: (string|null);

        /** TxDetailResponse tx */
        tx?: (protobuf.ITx|null);

        /** TxDetailResponse creationDt */
        creationDt?: (protobuf.ITimestamp|null);

        /** TxDetailResponse blockId */
        blockId?: (number|Long|null);
    }

    /** Represents a TxDetailResponse. */
    class TxDetailResponse implements ITxDetailResponse {

        /**
         * Constructs a new TxDetailResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: protobuf.ITxDetailResponse);

        /** TxDetailResponse txId. */
        public txId: string;

        /** TxDetailResponse tx. */
        public tx?: (protobuf.ITx|null);

        /** TxDetailResponse creationDt. */
        public creationDt?: (protobuf.ITimestamp|null);

        /** TxDetailResponse blockId. */
        public blockId: (number|Long);

        /**
         * Creates a new TxDetailResponse instance using the specified properties.
         * @param [properties] Properties to set
         * @returns TxDetailResponse instance
         */
        public static create(properties?: protobuf.ITxDetailResponse): protobuf.TxDetailResponse;

        /**
         * Encodes the specified TxDetailResponse message. Does not implicitly {@link protobuf.TxDetailResponse.verify|verify} messages.
         * @param message TxDetailResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: protobuf.ITxDetailResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified TxDetailResponse message, length delimited. Does not implicitly {@link protobuf.TxDetailResponse.verify|verify} messages.
         * @param message TxDetailResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: protobuf.ITxDetailResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a TxDetailResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns TxDetailResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): protobuf.TxDetailResponse;

        /**
         * Decodes a TxDetailResponse message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns TxDetailResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): protobuf.TxDetailResponse;

        /**
         * Verifies a TxDetailResponse message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a TxDetailResponse message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns TxDetailResponse
         */
        public static fromObject(object: { [k: string]: any }): protobuf.TxDetailResponse;

        /**
         * Creates a plain object from a TxDetailResponse message. Also converts values to other types if specified.
         * @param message TxDetailResponse
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: protobuf.TxDetailResponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this TxDetailResponse to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }
}
