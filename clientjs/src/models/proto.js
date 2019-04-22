/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.protobuf = (function() {

    /**
     * Namespace protobuf.
     * @exports protobuf
     * @namespace
     */
    var protobuf = {};

    protobuf.Timestamp = (function() {

        /**
         * Properties of a Timestamp.
         * @memberof protobuf
         * @interface ITimestamp
         * @property {number|Long|null} [seconds] Timestamp seconds
         * @property {number|Long|null} [nanos] Timestamp nanos
         */

        /**
         * Constructs a new Timestamp.
         * @memberof protobuf
         * @classdesc Represents a Timestamp.
         * @implements ITimestamp
         * @constructor
         * @param {protobuf.ITimestamp=} [properties] Properties to set
         */
        function Timestamp(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Timestamp seconds.
         * @member {number|Long} seconds
         * @memberof protobuf.Timestamp
         * @instance
         */
        Timestamp.prototype.seconds = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Timestamp nanos.
         * @member {number|Long} nanos
         * @memberof protobuf.Timestamp
         * @instance
         */
        Timestamp.prototype.nanos = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Creates a new Timestamp instance using the specified properties.
         * @function create
         * @memberof protobuf.Timestamp
         * @static
         * @param {protobuf.ITimestamp=} [properties] Properties to set
         * @returns {protobuf.Timestamp} Timestamp instance
         */
        Timestamp.create = function create(properties) {
            return new Timestamp(properties);
        };

        /**
         * Encodes the specified Timestamp message. Does not implicitly {@link protobuf.Timestamp.verify|verify} messages.
         * @function encode
         * @memberof protobuf.Timestamp
         * @static
         * @param {protobuf.ITimestamp} message Timestamp message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Timestamp.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.seconds != null && message.hasOwnProperty("seconds"))
                writer.uint32(/* id 1, wireType 0 =*/8).int64(message.seconds);
            if (message.nanos != null && message.hasOwnProperty("nanos"))
                writer.uint32(/* id 2, wireType 0 =*/16).int64(message.nanos);
            return writer;
        };

        /**
         * Encodes the specified Timestamp message, length delimited. Does not implicitly {@link protobuf.Timestamp.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protobuf.Timestamp
         * @static
         * @param {protobuf.ITimestamp} message Timestamp message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Timestamp.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Timestamp message from the specified reader or buffer.
         * @function decode
         * @memberof protobuf.Timestamp
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protobuf.Timestamp} Timestamp
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Timestamp.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.protobuf.Timestamp();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.seconds = reader.int64();
                    break;
                case 2:
                    message.nanos = reader.int64();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Timestamp message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protobuf.Timestamp
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protobuf.Timestamp} Timestamp
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Timestamp.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Timestamp message.
         * @function verify
         * @memberof protobuf.Timestamp
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Timestamp.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.seconds != null && message.hasOwnProperty("seconds"))
                if (!$util.isInteger(message.seconds) && !(message.seconds && $util.isInteger(message.seconds.low) && $util.isInteger(message.seconds.high)))
                    return "seconds: integer|Long expected";
            if (message.nanos != null && message.hasOwnProperty("nanos"))
                if (!$util.isInteger(message.nanos) && !(message.nanos && $util.isInteger(message.nanos.low) && $util.isInteger(message.nanos.high)))
                    return "nanos: integer|Long expected";
            return null;
        };

        /**
         * Creates a Timestamp message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protobuf.Timestamp
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protobuf.Timestamp} Timestamp
         */
        Timestamp.fromObject = function fromObject(object) {
            if (object instanceof $root.protobuf.Timestamp)
                return object;
            var message = new $root.protobuf.Timestamp();
            if (object.seconds != null)
                if ($util.Long)
                    (message.seconds = $util.Long.fromValue(object.seconds)).unsigned = false;
                else if (typeof object.seconds === "string")
                    message.seconds = parseInt(object.seconds, 10);
                else if (typeof object.seconds === "number")
                    message.seconds = object.seconds;
                else if (typeof object.seconds === "object")
                    message.seconds = new $util.LongBits(object.seconds.low >>> 0, object.seconds.high >>> 0).toNumber();
            if (object.nanos != null)
                if ($util.Long)
                    (message.nanos = $util.Long.fromValue(object.nanos)).unsigned = false;
                else if (typeof object.nanos === "string")
                    message.nanos = parseInt(object.nanos, 10);
                else if (typeof object.nanos === "number")
                    message.nanos = object.nanos;
                else if (typeof object.nanos === "object")
                    message.nanos = new $util.LongBits(object.nanos.low >>> 0, object.nanos.high >>> 0).toNumber();
            return message;
        };

        /**
         * Creates a plain object from a Timestamp message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protobuf.Timestamp
         * @static
         * @param {protobuf.Timestamp} message Timestamp
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Timestamp.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.seconds = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.seconds = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.nanos = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.nanos = options.longs === String ? "0" : 0;
            }
            if (message.seconds != null && message.hasOwnProperty("seconds"))
                if (typeof message.seconds === "number")
                    object.seconds = options.longs === String ? String(message.seconds) : message.seconds;
                else
                    object.seconds = options.longs === String ? $util.Long.prototype.toString.call(message.seconds) : options.longs === Number ? new $util.LongBits(message.seconds.low >>> 0, message.seconds.high >>> 0).toNumber() : message.seconds;
            if (message.nanos != null && message.hasOwnProperty("nanos"))
                if (typeof message.nanos === "number")
                    object.nanos = options.longs === String ? String(message.nanos) : message.nanos;
                else
                    object.nanos = options.longs === String ? $util.Long.prototype.toString.call(message.nanos) : options.longs === Number ? new $util.LongBits(message.nanos.low >>> 0, message.nanos.high >>> 0).toNumber() : message.nanos;
            return object;
        };

        /**
         * Converts this Timestamp to JSON.
         * @function toJSON
         * @memberof protobuf.Timestamp
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Timestamp.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Timestamp;
    })();

    protobuf.BlockHeightRequest = (function() {

        /**
         * Properties of a BlockHeightRequest.
         * @memberof protobuf
         * @interface IBlockHeightRequest
         * @property {number|Long|null} [blockHeight] BlockHeightRequest blockHeight
         */

        /**
         * Constructs a new BlockHeightRequest.
         * @memberof protobuf
         * @classdesc Represents a BlockHeightRequest.
         * @implements IBlockHeightRequest
         * @constructor
         * @param {protobuf.IBlockHeightRequest=} [properties] Properties to set
         */
        function BlockHeightRequest(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * BlockHeightRequest blockHeight.
         * @member {number|Long} blockHeight
         * @memberof protobuf.BlockHeightRequest
         * @instance
         */
        BlockHeightRequest.prototype.blockHeight = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * Creates a new BlockHeightRequest instance using the specified properties.
         * @function create
         * @memberof protobuf.BlockHeightRequest
         * @static
         * @param {protobuf.IBlockHeightRequest=} [properties] Properties to set
         * @returns {protobuf.BlockHeightRequest} BlockHeightRequest instance
         */
        BlockHeightRequest.create = function create(properties) {
            return new BlockHeightRequest(properties);
        };

        /**
         * Encodes the specified BlockHeightRequest message. Does not implicitly {@link protobuf.BlockHeightRequest.verify|verify} messages.
         * @function encode
         * @memberof protobuf.BlockHeightRequest
         * @static
         * @param {protobuf.IBlockHeightRequest} message BlockHeightRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        BlockHeightRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.blockHeight != null && message.hasOwnProperty("blockHeight"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint64(message.blockHeight);
            return writer;
        };

        /**
         * Encodes the specified BlockHeightRequest message, length delimited. Does not implicitly {@link protobuf.BlockHeightRequest.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protobuf.BlockHeightRequest
         * @static
         * @param {protobuf.IBlockHeightRequest} message BlockHeightRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        BlockHeightRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a BlockHeightRequest message from the specified reader or buffer.
         * @function decode
         * @memberof protobuf.BlockHeightRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protobuf.BlockHeightRequest} BlockHeightRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        BlockHeightRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.protobuf.BlockHeightRequest();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.blockHeight = reader.uint64();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a BlockHeightRequest message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protobuf.BlockHeightRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protobuf.BlockHeightRequest} BlockHeightRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        BlockHeightRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a BlockHeightRequest message.
         * @function verify
         * @memberof protobuf.BlockHeightRequest
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        BlockHeightRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.blockHeight != null && message.hasOwnProperty("blockHeight"))
                if (!$util.isInteger(message.blockHeight) && !(message.blockHeight && $util.isInteger(message.blockHeight.low) && $util.isInteger(message.blockHeight.high)))
                    return "blockHeight: integer|Long expected";
            return null;
        };

        /**
         * Creates a BlockHeightRequest message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protobuf.BlockHeightRequest
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protobuf.BlockHeightRequest} BlockHeightRequest
         */
        BlockHeightRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.protobuf.BlockHeightRequest)
                return object;
            var message = new $root.protobuf.BlockHeightRequest();
            if (object.blockHeight != null)
                if ($util.Long)
                    (message.blockHeight = $util.Long.fromValue(object.blockHeight)).unsigned = true;
                else if (typeof object.blockHeight === "string")
                    message.blockHeight = parseInt(object.blockHeight, 10);
                else if (typeof object.blockHeight === "number")
                    message.blockHeight = object.blockHeight;
                else if (typeof object.blockHeight === "object")
                    message.blockHeight = new $util.LongBits(object.blockHeight.low >>> 0, object.blockHeight.high >>> 0).toNumber(true);
            return message;
        };

        /**
         * Creates a plain object from a BlockHeightRequest message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protobuf.BlockHeightRequest
         * @static
         * @param {protobuf.BlockHeightRequest} message BlockHeightRequest
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        BlockHeightRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.blockHeight = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.blockHeight = options.longs === String ? "0" : 0;
            if (message.blockHeight != null && message.hasOwnProperty("blockHeight"))
                if (typeof message.blockHeight === "number")
                    object.blockHeight = options.longs === String ? String(message.blockHeight) : message.blockHeight;
                else
                    object.blockHeight = options.longs === String ? $util.Long.prototype.toString.call(message.blockHeight) : options.longs === Number ? new $util.LongBits(message.blockHeight.low >>> 0, message.blockHeight.high >>> 0).toNumber(true) : message.blockHeight;
            return object;
        };

        /**
         * Converts this BlockHeightRequest to JSON.
         * @function toJSON
         * @memberof protobuf.BlockHeightRequest
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        BlockHeightRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return BlockHeightRequest;
    })();

    protobuf.BlockResponse = (function() {

        /**
         * Properties of a BlockResponse.
         * @memberof protobuf
         * @interface IBlockResponse
         * @property {number|Long|null} [blockHeight] BlockResponse blockHeight
         * @property {protobuf.ITimestamp|null} [time] BlockResponse time
         * @property {number|Long|null} [transactions] BlockResponse transactions
         * @property {string|null} [supervisorAddress] BlockResponse supervisorAddress
         * @property {Uint8Array|null} [stateRoot] BlockResponse stateRoot
         */

        /**
         * Constructs a new BlockResponse.
         * @memberof protobuf
         * @classdesc Represents a BlockResponse.
         * @implements IBlockResponse
         * @constructor
         * @param {protobuf.IBlockResponse=} [properties] Properties to set
         */
        function BlockResponse(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * BlockResponse blockHeight.
         * @member {number|Long} blockHeight
         * @memberof protobuf.BlockResponse
         * @instance
         */
        BlockResponse.prototype.blockHeight = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * BlockResponse time.
         * @member {protobuf.ITimestamp|null|undefined} time
         * @memberof protobuf.BlockResponse
         * @instance
         */
        BlockResponse.prototype.time = null;

        /**
         * BlockResponse transactions.
         * @member {number|Long} transactions
         * @memberof protobuf.BlockResponse
         * @instance
         */
        BlockResponse.prototype.transactions = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * BlockResponse supervisorAddress.
         * @member {string} supervisorAddress
         * @memberof protobuf.BlockResponse
         * @instance
         */
        BlockResponse.prototype.supervisorAddress = "";

        /**
         * BlockResponse stateRoot.
         * @member {Uint8Array} stateRoot
         * @memberof protobuf.BlockResponse
         * @instance
         */
        BlockResponse.prototype.stateRoot = $util.newBuffer([]);

        /**
         * Creates a new BlockResponse instance using the specified properties.
         * @function create
         * @memberof protobuf.BlockResponse
         * @static
         * @param {protobuf.IBlockResponse=} [properties] Properties to set
         * @returns {protobuf.BlockResponse} BlockResponse instance
         */
        BlockResponse.create = function create(properties) {
            return new BlockResponse(properties);
        };

        /**
         * Encodes the specified BlockResponse message. Does not implicitly {@link protobuf.BlockResponse.verify|verify} messages.
         * @function encode
         * @memberof protobuf.BlockResponse
         * @static
         * @param {protobuf.IBlockResponse} message BlockResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        BlockResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.blockHeight != null && message.hasOwnProperty("blockHeight"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint64(message.blockHeight);
            if (message.time != null && message.hasOwnProperty("time"))
                $root.protobuf.Timestamp.encode(message.time, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            if (message.transactions != null && message.hasOwnProperty("transactions"))
                writer.uint32(/* id 3, wireType 0 =*/24).uint64(message.transactions);
            if (message.supervisorAddress != null && message.hasOwnProperty("supervisorAddress"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.supervisorAddress);
            if (message.stateRoot != null && message.hasOwnProperty("stateRoot"))
                writer.uint32(/* id 5, wireType 2 =*/42).bytes(message.stateRoot);
            return writer;
        };

        /**
         * Encodes the specified BlockResponse message, length delimited. Does not implicitly {@link protobuf.BlockResponse.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protobuf.BlockResponse
         * @static
         * @param {protobuf.IBlockResponse} message BlockResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        BlockResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a BlockResponse message from the specified reader or buffer.
         * @function decode
         * @memberof protobuf.BlockResponse
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protobuf.BlockResponse} BlockResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        BlockResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.protobuf.BlockResponse();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.blockHeight = reader.uint64();
                    break;
                case 2:
                    message.time = $root.protobuf.Timestamp.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.transactions = reader.uint64();
                    break;
                case 4:
                    message.supervisorAddress = reader.string();
                    break;
                case 5:
                    message.stateRoot = reader.bytes();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a BlockResponse message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protobuf.BlockResponse
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protobuf.BlockResponse} BlockResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        BlockResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a BlockResponse message.
         * @function verify
         * @memberof protobuf.BlockResponse
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        BlockResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.blockHeight != null && message.hasOwnProperty("blockHeight"))
                if (!$util.isInteger(message.blockHeight) && !(message.blockHeight && $util.isInteger(message.blockHeight.low) && $util.isInteger(message.blockHeight.high)))
                    return "blockHeight: integer|Long expected";
            if (message.time != null && message.hasOwnProperty("time")) {
                var error = $root.protobuf.Timestamp.verify(message.time);
                if (error)
                    return "time." + error;
            }
            if (message.transactions != null && message.hasOwnProperty("transactions"))
                if (!$util.isInteger(message.transactions) && !(message.transactions && $util.isInteger(message.transactions.low) && $util.isInteger(message.transactions.high)))
                    return "transactions: integer|Long expected";
            if (message.supervisorAddress != null && message.hasOwnProperty("supervisorAddress"))
                if (!$util.isString(message.supervisorAddress))
                    return "supervisorAddress: string expected";
            if (message.stateRoot != null && message.hasOwnProperty("stateRoot"))
                if (!(message.stateRoot && typeof message.stateRoot.length === "number" || $util.isString(message.stateRoot)))
                    return "stateRoot: buffer expected";
            return null;
        };

        /**
         * Creates a BlockResponse message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protobuf.BlockResponse
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protobuf.BlockResponse} BlockResponse
         */
        BlockResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.protobuf.BlockResponse)
                return object;
            var message = new $root.protobuf.BlockResponse();
            if (object.blockHeight != null)
                if ($util.Long)
                    (message.blockHeight = $util.Long.fromValue(object.blockHeight)).unsigned = true;
                else if (typeof object.blockHeight === "string")
                    message.blockHeight = parseInt(object.blockHeight, 10);
                else if (typeof object.blockHeight === "number")
                    message.blockHeight = object.blockHeight;
                else if (typeof object.blockHeight === "object")
                    message.blockHeight = new $util.LongBits(object.blockHeight.low >>> 0, object.blockHeight.high >>> 0).toNumber(true);
            if (object.time != null) {
                if (typeof object.time !== "object")
                    throw TypeError(".protobuf.BlockResponse.time: object expected");
                message.time = $root.protobuf.Timestamp.fromObject(object.time);
            }
            if (object.transactions != null)
                if ($util.Long)
                    (message.transactions = $util.Long.fromValue(object.transactions)).unsigned = true;
                else if (typeof object.transactions === "string")
                    message.transactions = parseInt(object.transactions, 10);
                else if (typeof object.transactions === "number")
                    message.transactions = object.transactions;
                else if (typeof object.transactions === "object")
                    message.transactions = new $util.LongBits(object.transactions.low >>> 0, object.transactions.high >>> 0).toNumber(true);
            if (object.supervisorAddress != null)
                message.supervisorAddress = String(object.supervisorAddress);
            if (object.stateRoot != null)
                if (typeof object.stateRoot === "string")
                    $util.base64.decode(object.stateRoot, message.stateRoot = $util.newBuffer($util.base64.length(object.stateRoot)), 0);
                else if (object.stateRoot.length)
                    message.stateRoot = object.stateRoot;
            return message;
        };

        /**
         * Creates a plain object from a BlockResponse message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protobuf.BlockResponse
         * @static
         * @param {protobuf.BlockResponse} message BlockResponse
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        BlockResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.blockHeight = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.blockHeight = options.longs === String ? "0" : 0;
                object.time = null;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.transactions = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.transactions = options.longs === String ? "0" : 0;
                object.supervisorAddress = "";
                if (options.bytes === String)
                    object.stateRoot = "";
                else {
                    object.stateRoot = [];
                    if (options.bytes !== Array)
                        object.stateRoot = $util.newBuffer(object.stateRoot);
                }
            }
            if (message.blockHeight != null && message.hasOwnProperty("blockHeight"))
                if (typeof message.blockHeight === "number")
                    object.blockHeight = options.longs === String ? String(message.blockHeight) : message.blockHeight;
                else
                    object.blockHeight = options.longs === String ? $util.Long.prototype.toString.call(message.blockHeight) : options.longs === Number ? new $util.LongBits(message.blockHeight.low >>> 0, message.blockHeight.high >>> 0).toNumber(true) : message.blockHeight;
            if (message.time != null && message.hasOwnProperty("time"))
                object.time = $root.protobuf.Timestamp.toObject(message.time, options);
            if (message.transactions != null && message.hasOwnProperty("transactions"))
                if (typeof message.transactions === "number")
                    object.transactions = options.longs === String ? String(message.transactions) : message.transactions;
                else
                    object.transactions = options.longs === String ? $util.Long.prototype.toString.call(message.transactions) : options.longs === Number ? new $util.LongBits(message.transactions.low >>> 0, message.transactions.high >>> 0).toNumber(true) : message.transactions;
            if (message.supervisorAddress != null && message.hasOwnProperty("supervisorAddress"))
                object.supervisorAddress = message.supervisorAddress;
            if (message.stateRoot != null && message.hasOwnProperty("stateRoot"))
                object.stateRoot = options.bytes === String ? $util.base64.encode(message.stateRoot, 0, message.stateRoot.length) : options.bytes === Array ? Array.prototype.slice.call(message.stateRoot) : message.stateRoot;
            return object;
        };

        /**
         * Converts this BlockResponse to JSON.
         * @function toJSON
         * @memberof protobuf.BlockResponse
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        BlockResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return BlockResponse;
    })();

    protobuf.AccountRequest = (function() {

        /**
         * Properties of an AccountRequest.
         * @memberof protobuf
         * @interface IAccountRequest
         * @property {string|null} [address] AccountRequest address
         */

        /**
         * Constructs a new AccountRequest.
         * @memberof protobuf
         * @classdesc Represents an AccountRequest.
         * @implements IAccountRequest
         * @constructor
         * @param {protobuf.IAccountRequest=} [properties] Properties to set
         */
        function AccountRequest(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * AccountRequest address.
         * @member {string} address
         * @memberof protobuf.AccountRequest
         * @instance
         */
        AccountRequest.prototype.address = "";

        /**
         * Creates a new AccountRequest instance using the specified properties.
         * @function create
         * @memberof protobuf.AccountRequest
         * @static
         * @param {protobuf.IAccountRequest=} [properties] Properties to set
         * @returns {protobuf.AccountRequest} AccountRequest instance
         */
        AccountRequest.create = function create(properties) {
            return new AccountRequest(properties);
        };

        /**
         * Encodes the specified AccountRequest message. Does not implicitly {@link protobuf.AccountRequest.verify|verify} messages.
         * @function encode
         * @memberof protobuf.AccountRequest
         * @static
         * @param {protobuf.IAccountRequest} message AccountRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        AccountRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.address != null && message.hasOwnProperty("address"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.address);
            return writer;
        };

        /**
         * Encodes the specified AccountRequest message, length delimited. Does not implicitly {@link protobuf.AccountRequest.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protobuf.AccountRequest
         * @static
         * @param {protobuf.IAccountRequest} message AccountRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        AccountRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an AccountRequest message from the specified reader or buffer.
         * @function decode
         * @memberof protobuf.AccountRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protobuf.AccountRequest} AccountRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        AccountRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.protobuf.AccountRequest();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an AccountRequest message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protobuf.AccountRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protobuf.AccountRequest} AccountRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        AccountRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an AccountRequest message.
         * @function verify
         * @memberof protobuf.AccountRequest
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        AccountRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.address != null && message.hasOwnProperty("address"))
                if (!$util.isString(message.address))
                    return "address: string expected";
            return null;
        };

        /**
         * Creates an AccountRequest message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protobuf.AccountRequest
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protobuf.AccountRequest} AccountRequest
         */
        AccountRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.protobuf.AccountRequest)
                return object;
            var message = new $root.protobuf.AccountRequest();
            if (object.address != null)
                message.address = String(object.address);
            return message;
        };

        /**
         * Creates a plain object from an AccountRequest message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protobuf.AccountRequest
         * @static
         * @param {protobuf.AccountRequest} message AccountRequest
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        AccountRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.address = "";
            if (message.address != null && message.hasOwnProperty("address"))
                object.address = message.address;
            return object;
        };

        /**
         * Converts this AccountRequest to JSON.
         * @function toJSON
         * @memberof protobuf.AccountRequest
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        AccountRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return AccountRequest;
    })();

    protobuf.AccountResponse = (function() {

        /**
         * Properties of an AccountResponse.
         * @memberof protobuf
         * @interface IAccountResponse
         * @property {string|null} [address] AccountResponse address
         * @property {number|Long|null} [nonce] AccountResponse nonce
         * @property {string|null} [storageRoot] AccountResponse storageRoot
         * @property {string|null} [publicKey] AccountResponse publicKey
         * @property {number|Long|null} [balance] AccountResponse balance
         * @property {Object.<string,number|Long>|null} [balances] AccountResponse balances
         */

        /**
         * Constructs a new AccountResponse.
         * @memberof protobuf
         * @classdesc Represents an AccountResponse.
         * @implements IAccountResponse
         * @constructor
         * @param {protobuf.IAccountResponse=} [properties] Properties to set
         */
        function AccountResponse(properties) {
            this.balances = {};
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * AccountResponse address.
         * @member {string} address
         * @memberof protobuf.AccountResponse
         * @instance
         */
        AccountResponse.prototype.address = "";

        /**
         * AccountResponse nonce.
         * @member {number|Long} nonce
         * @memberof protobuf.AccountResponse
         * @instance
         */
        AccountResponse.prototype.nonce = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * AccountResponse storageRoot.
         * @member {string} storageRoot
         * @memberof protobuf.AccountResponse
         * @instance
         */
        AccountResponse.prototype.storageRoot = "";

        /**
         * AccountResponse publicKey.
         * @member {string} publicKey
         * @memberof protobuf.AccountResponse
         * @instance
         */
        AccountResponse.prototype.publicKey = "";

        /**
         * AccountResponse balance.
         * @member {number|Long} balance
         * @memberof protobuf.AccountResponse
         * @instance
         */
        AccountResponse.prototype.balance = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * AccountResponse balances.
         * @member {Object.<string,number|Long>} balances
         * @memberof protobuf.AccountResponse
         * @instance
         */
        AccountResponse.prototype.balances = $util.emptyObject;

        /**
         * Creates a new AccountResponse instance using the specified properties.
         * @function create
         * @memberof protobuf.AccountResponse
         * @static
         * @param {protobuf.IAccountResponse=} [properties] Properties to set
         * @returns {protobuf.AccountResponse} AccountResponse instance
         */
        AccountResponse.create = function create(properties) {
            return new AccountResponse(properties);
        };

        /**
         * Encodes the specified AccountResponse message. Does not implicitly {@link protobuf.AccountResponse.verify|verify} messages.
         * @function encode
         * @memberof protobuf.AccountResponse
         * @static
         * @param {protobuf.IAccountResponse} message AccountResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        AccountResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.address != null && message.hasOwnProperty("address"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.address);
            if (message.nonce != null && message.hasOwnProperty("nonce"))
                writer.uint32(/* id 2, wireType 0 =*/16).uint64(message.nonce);
            if (message.storageRoot != null && message.hasOwnProperty("storageRoot"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.storageRoot);
            if (message.publicKey != null && message.hasOwnProperty("publicKey"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.publicKey);
            if (message.balance != null && message.hasOwnProperty("balance"))
                writer.uint32(/* id 5, wireType 0 =*/40).uint64(message.balance);
            if (message.balances != null && message.hasOwnProperty("balances"))
                for (var keys = Object.keys(message.balances), i = 0; i < keys.length; ++i)
                    writer.uint32(/* id 6, wireType 2 =*/50).fork().uint32(/* id 1, wireType 2 =*/10).string(keys[i]).uint32(/* id 2, wireType 0 =*/16).uint64(message.balances[keys[i]]).ldelim();
            return writer;
        };

        /**
         * Encodes the specified AccountResponse message, length delimited. Does not implicitly {@link protobuf.AccountResponse.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protobuf.AccountResponse
         * @static
         * @param {protobuf.IAccountResponse} message AccountResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        AccountResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an AccountResponse message from the specified reader or buffer.
         * @function decode
         * @memberof protobuf.AccountResponse
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protobuf.AccountResponse} AccountResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        AccountResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.protobuf.AccountResponse(), key;
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                case 2:
                    message.nonce = reader.uint64();
                    break;
                case 3:
                    message.storageRoot = reader.string();
                    break;
                case 4:
                    message.publicKey = reader.string();
                    break;
                case 5:
                    message.balance = reader.uint64();
                    break;
                case 6:
                    reader.skip().pos++;
                    if (message.balances === $util.emptyObject)
                        message.balances = {};
                    key = reader.string();
                    reader.pos++;
                    message.balances[key] = reader.uint64();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an AccountResponse message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protobuf.AccountResponse
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protobuf.AccountResponse} AccountResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        AccountResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an AccountResponse message.
         * @function verify
         * @memberof protobuf.AccountResponse
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        AccountResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.address != null && message.hasOwnProperty("address"))
                if (!$util.isString(message.address))
                    return "address: string expected";
            if (message.nonce != null && message.hasOwnProperty("nonce"))
                if (!$util.isInteger(message.nonce) && !(message.nonce && $util.isInteger(message.nonce.low) && $util.isInteger(message.nonce.high)))
                    return "nonce: integer|Long expected";
            if (message.storageRoot != null && message.hasOwnProperty("storageRoot"))
                if (!$util.isString(message.storageRoot))
                    return "storageRoot: string expected";
            if (message.publicKey != null && message.hasOwnProperty("publicKey"))
                if (!$util.isString(message.publicKey))
                    return "publicKey: string expected";
            if (message.balance != null && message.hasOwnProperty("balance"))
                if (!$util.isInteger(message.balance) && !(message.balance && $util.isInteger(message.balance.low) && $util.isInteger(message.balance.high)))
                    return "balance: integer|Long expected";
            if (message.balances != null && message.hasOwnProperty("balances")) {
                if (!$util.isObject(message.balances))
                    return "balances: object expected";
                var key = Object.keys(message.balances);
                for (var i = 0; i < key.length; ++i)
                    if (!$util.isInteger(message.balances[key[i]]) && !(message.balances[key[i]] && $util.isInteger(message.balances[key[i]].low) && $util.isInteger(message.balances[key[i]].high)))
                        return "balances: integer|Long{k:string} expected";
            }
            return null;
        };

        /**
         * Creates an AccountResponse message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protobuf.AccountResponse
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protobuf.AccountResponse} AccountResponse
         */
        AccountResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.protobuf.AccountResponse)
                return object;
            var message = new $root.protobuf.AccountResponse();
            if (object.address != null)
                message.address = String(object.address);
            if (object.nonce != null)
                if ($util.Long)
                    (message.nonce = $util.Long.fromValue(object.nonce)).unsigned = true;
                else if (typeof object.nonce === "string")
                    message.nonce = parseInt(object.nonce, 10);
                else if (typeof object.nonce === "number")
                    message.nonce = object.nonce;
                else if (typeof object.nonce === "object")
                    message.nonce = new $util.LongBits(object.nonce.low >>> 0, object.nonce.high >>> 0).toNumber(true);
            if (object.storageRoot != null)
                message.storageRoot = String(object.storageRoot);
            if (object.publicKey != null)
                message.publicKey = String(object.publicKey);
            if (object.balance != null)
                if ($util.Long)
                    (message.balance = $util.Long.fromValue(object.balance)).unsigned = true;
                else if (typeof object.balance === "string")
                    message.balance = parseInt(object.balance, 10);
                else if (typeof object.balance === "number")
                    message.balance = object.balance;
                else if (typeof object.balance === "object")
                    message.balance = new $util.LongBits(object.balance.low >>> 0, object.balance.high >>> 0).toNumber(true);
            if (object.balances) {
                if (typeof object.balances !== "object")
                    throw TypeError(".protobuf.AccountResponse.balances: object expected");
                message.balances = {};
                for (var keys = Object.keys(object.balances), i = 0; i < keys.length; ++i)
                    if ($util.Long)
                        (message.balances[keys[i]] = $util.Long.fromValue(object.balances[keys[i]])).unsigned = true;
                    else if (typeof object.balances[keys[i]] === "string")
                        message.balances[keys[i]] = parseInt(object.balances[keys[i]], 10);
                    else if (typeof object.balances[keys[i]] === "number")
                        message.balances[keys[i]] = object.balances[keys[i]];
                    else if (typeof object.balances[keys[i]] === "object")
                        message.balances[keys[i]] = new $util.LongBits(object.balances[keys[i]].low >>> 0, object.balances[keys[i]].high >>> 0).toNumber(true);
            }
            return message;
        };

        /**
         * Creates a plain object from an AccountResponse message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protobuf.AccountResponse
         * @static
         * @param {protobuf.AccountResponse} message AccountResponse
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        AccountResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.objects || options.defaults)
                object.balances = {};
            if (options.defaults) {
                object.address = "";
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.nonce = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.nonce = options.longs === String ? "0" : 0;
                object.storageRoot = "";
                object.publicKey = "";
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.balance = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.balance = options.longs === String ? "0" : 0;
            }
            if (message.address != null && message.hasOwnProperty("address"))
                object.address = message.address;
            if (message.nonce != null && message.hasOwnProperty("nonce"))
                if (typeof message.nonce === "number")
                    object.nonce = options.longs === String ? String(message.nonce) : message.nonce;
                else
                    object.nonce = options.longs === String ? $util.Long.prototype.toString.call(message.nonce) : options.longs === Number ? new $util.LongBits(message.nonce.low >>> 0, message.nonce.high >>> 0).toNumber(true) : message.nonce;
            if (message.storageRoot != null && message.hasOwnProperty("storageRoot"))
                object.storageRoot = message.storageRoot;
            if (message.publicKey != null && message.hasOwnProperty("publicKey"))
                object.publicKey = message.publicKey;
            if (message.balance != null && message.hasOwnProperty("balance"))
                if (typeof message.balance === "number")
                    object.balance = options.longs === String ? String(message.balance) : message.balance;
                else
                    object.balance = options.longs === String ? $util.Long.prototype.toString.call(message.balance) : options.longs === Number ? new $util.LongBits(message.balance.low >>> 0, message.balance.high >>> 0).toNumber(true) : message.balance;
            var keys2;
            if (message.balances && (keys2 = Object.keys(message.balances)).length) {
                object.balances = {};
                for (var j = 0; j < keys2.length; ++j)
                    if (typeof message.balances[keys2[j]] === "number")
                        object.balances[keys2[j]] = options.longs === String ? String(message.balances[keys2[j]]) : message.balances[keys2[j]];
                    else
                        object.balances[keys2[j]] = options.longs === String ? $util.Long.prototype.toString.call(message.balances[keys2[j]]) : options.longs === Number ? new $util.LongBits(message.balances[keys2[j]].low >>> 0, message.balances[keys2[j]].high >>> 0).toNumber(true) : message.balances[keys2[j]];
            }
            return object;
        };

        /**
         * Converts this AccountResponse to JSON.
         * @function toJSON
         * @memberof protobuf.AccountResponse
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        AccountResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return AccountResponse;
    })();

    protobuf.Tx = (function() {

        /**
         * Properties of a Tx.
         * @memberof protobuf
         * @interface ITx
         * @property {string|null} [senderAddress] Tx senderAddress
         * @property {string|null} [senderPubkey] Tx senderPubkey
         * @property {string|null} [recieverAddress] Tx recieverAddress
         * @property {protobuf.IAsset|null} [asset] Tx asset
         * @property {string|null} [message] Tx message
         * @property {string|null} [sign] Tx sign
         * @property {string|null} [type] Tx type
         * @property {string|null} [status] Tx status
         */

        /**
         * Constructs a new Tx.
         * @memberof protobuf
         * @classdesc Represents a Tx.
         * @implements ITx
         * @constructor
         * @param {protobuf.ITx=} [properties] Properties to set
         */
        function Tx(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Tx senderAddress.
         * @member {string} senderAddress
         * @memberof protobuf.Tx
         * @instance
         */
        Tx.prototype.senderAddress = "";

        /**
         * Tx senderPubkey.
         * @member {string} senderPubkey
         * @memberof protobuf.Tx
         * @instance
         */
        Tx.prototype.senderPubkey = "";

        /**
         * Tx recieverAddress.
         * @member {string} recieverAddress
         * @memberof protobuf.Tx
         * @instance
         */
        Tx.prototype.recieverAddress = "";

        /**
         * Tx asset.
         * @member {protobuf.IAsset|null|undefined} asset
         * @memberof protobuf.Tx
         * @instance
         */
        Tx.prototype.asset = null;

        /**
         * Tx message.
         * @member {string} message
         * @memberof protobuf.Tx
         * @instance
         */
        Tx.prototype.message = "";

        /**
         * Tx sign.
         * @member {string} sign
         * @memberof protobuf.Tx
         * @instance
         */
        Tx.prototype.sign = "";

        /**
         * Tx type.
         * @member {string} type
         * @memberof protobuf.Tx
         * @instance
         */
        Tx.prototype.type = "";

        /**
         * Tx status.
         * @member {string} status
         * @memberof protobuf.Tx
         * @instance
         */
        Tx.prototype.status = "";

        /**
         * Creates a new Tx instance using the specified properties.
         * @function create
         * @memberof protobuf.Tx
         * @static
         * @param {protobuf.ITx=} [properties] Properties to set
         * @returns {protobuf.Tx} Tx instance
         */
        Tx.create = function create(properties) {
            return new Tx(properties);
        };

        /**
         * Encodes the specified Tx message. Does not implicitly {@link protobuf.Tx.verify|verify} messages.
         * @function encode
         * @memberof protobuf.Tx
         * @static
         * @param {protobuf.ITx} message Tx message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Tx.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.senderAddress != null && message.hasOwnProperty("senderAddress"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.senderAddress);
            if (message.senderPubkey != null && message.hasOwnProperty("senderPubkey"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.senderPubkey);
            if (message.recieverAddress != null && message.hasOwnProperty("recieverAddress"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.recieverAddress);
            if (message.asset != null && message.hasOwnProperty("asset"))
                $root.protobuf.Asset.encode(message.asset, writer.uint32(/* id 4, wireType 2 =*/34).fork()).ldelim();
            if (message.message != null && message.hasOwnProperty("message"))
                writer.uint32(/* id 5, wireType 2 =*/42).string(message.message);
            if (message.sign != null && message.hasOwnProperty("sign"))
                writer.uint32(/* id 6, wireType 2 =*/50).string(message.sign);
            if (message.type != null && message.hasOwnProperty("type"))
                writer.uint32(/* id 7, wireType 2 =*/58).string(message.type);
            if (message.status != null && message.hasOwnProperty("status"))
                writer.uint32(/* id 8, wireType 2 =*/66).string(message.status);
            return writer;
        };

        /**
         * Encodes the specified Tx message, length delimited. Does not implicitly {@link protobuf.Tx.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protobuf.Tx
         * @static
         * @param {protobuf.ITx} message Tx message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Tx.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Tx message from the specified reader or buffer.
         * @function decode
         * @memberof protobuf.Tx
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protobuf.Tx} Tx
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Tx.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.protobuf.Tx();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.senderAddress = reader.string();
                    break;
                case 2:
                    message.senderPubkey = reader.string();
                    break;
                case 3:
                    message.recieverAddress = reader.string();
                    break;
                case 4:
                    message.asset = $root.protobuf.Asset.decode(reader, reader.uint32());
                    break;
                case 5:
                    message.message = reader.string();
                    break;
                case 6:
                    message.sign = reader.string();
                    break;
                case 7:
                    message.type = reader.string();
                    break;
                case 8:
                    message.status = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Tx message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protobuf.Tx
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protobuf.Tx} Tx
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Tx.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Tx message.
         * @function verify
         * @memberof protobuf.Tx
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Tx.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.senderAddress != null && message.hasOwnProperty("senderAddress"))
                if (!$util.isString(message.senderAddress))
                    return "senderAddress: string expected";
            if (message.senderPubkey != null && message.hasOwnProperty("senderPubkey"))
                if (!$util.isString(message.senderPubkey))
                    return "senderPubkey: string expected";
            if (message.recieverAddress != null && message.hasOwnProperty("recieverAddress"))
                if (!$util.isString(message.recieverAddress))
                    return "recieverAddress: string expected";
            if (message.asset != null && message.hasOwnProperty("asset")) {
                var error = $root.protobuf.Asset.verify(message.asset);
                if (error)
                    return "asset." + error;
            }
            if (message.message != null && message.hasOwnProperty("message"))
                if (!$util.isString(message.message))
                    return "message: string expected";
            if (message.sign != null && message.hasOwnProperty("sign"))
                if (!$util.isString(message.sign))
                    return "sign: string expected";
            if (message.type != null && message.hasOwnProperty("type"))
                if (!$util.isString(message.type))
                    return "type: string expected";
            if (message.status != null && message.hasOwnProperty("status"))
                if (!$util.isString(message.status))
                    return "status: string expected";
            return null;
        };

        /**
         * Creates a Tx message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protobuf.Tx
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protobuf.Tx} Tx
         */
        Tx.fromObject = function fromObject(object) {
            if (object instanceof $root.protobuf.Tx)
                return object;
            var message = new $root.protobuf.Tx();
            if (object.senderAddress != null)
                message.senderAddress = String(object.senderAddress);
            if (object.senderPubkey != null)
                message.senderPubkey = String(object.senderPubkey);
            if (object.recieverAddress != null)
                message.recieverAddress = String(object.recieverAddress);
            if (object.asset != null) {
                if (typeof object.asset !== "object")
                    throw TypeError(".protobuf.Tx.asset: object expected");
                message.asset = $root.protobuf.Asset.fromObject(object.asset);
            }
            if (object.message != null)
                message.message = String(object.message);
            if (object.sign != null)
                message.sign = String(object.sign);
            if (object.type != null)
                message.type = String(object.type);
            if (object.status != null)
                message.status = String(object.status);
            return message;
        };

        /**
         * Creates a plain object from a Tx message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protobuf.Tx
         * @static
         * @param {protobuf.Tx} message Tx
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Tx.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.senderAddress = "";
                object.senderPubkey = "";
                object.recieverAddress = "";
                object.asset = null;
                object.message = "";
                object.sign = "";
                object.type = "";
                object.status = "";
            }
            if (message.senderAddress != null && message.hasOwnProperty("senderAddress"))
                object.senderAddress = message.senderAddress;
            if (message.senderPubkey != null && message.hasOwnProperty("senderPubkey"))
                object.senderPubkey = message.senderPubkey;
            if (message.recieverAddress != null && message.hasOwnProperty("recieverAddress"))
                object.recieverAddress = message.recieverAddress;
            if (message.asset != null && message.hasOwnProperty("asset"))
                object.asset = $root.protobuf.Asset.toObject(message.asset, options);
            if (message.message != null && message.hasOwnProperty("message"))
                object.message = message.message;
            if (message.sign != null && message.hasOwnProperty("sign"))
                object.sign = message.sign;
            if (message.type != null && message.hasOwnProperty("type"))
                object.type = message.type;
            if (message.status != null && message.hasOwnProperty("status"))
                object.status = message.status;
            return object;
        };

        /**
         * Converts this Tx to JSON.
         * @function toJSON
         * @memberof protobuf.Tx
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Tx.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Tx;
    })();

    protobuf.Asset = (function() {

        /**
         * Properties of an Asset.
         * @memberof protobuf
         * @interface IAsset
         * @property {string|null} [category] Asset category
         * @property {string|null} [symbol] Asset symbol
         * @property {string|null} [network] Asset network
         * @property {number|Long|null} [value] Asset value
         * @property {number|Long|null} [fee] Asset fee
         * @property {number|Long|null} [nonce] Asset nonce
         */

        /**
         * Constructs a new Asset.
         * @memberof protobuf
         * @classdesc Represents an Asset.
         * @implements IAsset
         * @constructor
         * @param {protobuf.IAsset=} [properties] Properties to set
         */
        function Asset(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Asset category.
         * @member {string} category
         * @memberof protobuf.Asset
         * @instance
         */
        Asset.prototype.category = "";

        /**
         * Asset symbol.
         * @member {string} symbol
         * @memberof protobuf.Asset
         * @instance
         */
        Asset.prototype.symbol = "";

        /**
         * Asset network.
         * @member {string} network
         * @memberof protobuf.Asset
         * @instance
         */
        Asset.prototype.network = "";

        /**
         * Asset value.
         * @member {number|Long} value
         * @memberof protobuf.Asset
         * @instance
         */
        Asset.prototype.value = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * Asset fee.
         * @member {number|Long} fee
         * @memberof protobuf.Asset
         * @instance
         */
        Asset.prototype.fee = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * Asset nonce.
         * @member {number|Long} nonce
         * @memberof protobuf.Asset
         * @instance
         */
        Asset.prototype.nonce = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * Creates a new Asset instance using the specified properties.
         * @function create
         * @memberof protobuf.Asset
         * @static
         * @param {protobuf.IAsset=} [properties] Properties to set
         * @returns {protobuf.Asset} Asset instance
         */
        Asset.create = function create(properties) {
            return new Asset(properties);
        };

        /**
         * Encodes the specified Asset message. Does not implicitly {@link protobuf.Asset.verify|verify} messages.
         * @function encode
         * @memberof protobuf.Asset
         * @static
         * @param {protobuf.IAsset} message Asset message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Asset.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.category != null && message.hasOwnProperty("category"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.category);
            if (message.symbol != null && message.hasOwnProperty("symbol"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.symbol);
            if (message.network != null && message.hasOwnProperty("network"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.network);
            if (message.value != null && message.hasOwnProperty("value"))
                writer.uint32(/* id 4, wireType 0 =*/32).uint64(message.value);
            if (message.fee != null && message.hasOwnProperty("fee"))
                writer.uint32(/* id 5, wireType 0 =*/40).uint64(message.fee);
            if (message.nonce != null && message.hasOwnProperty("nonce"))
                writer.uint32(/* id 6, wireType 0 =*/48).uint64(message.nonce);
            return writer;
        };

        /**
         * Encodes the specified Asset message, length delimited. Does not implicitly {@link protobuf.Asset.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protobuf.Asset
         * @static
         * @param {protobuf.IAsset} message Asset message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Asset.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an Asset message from the specified reader or buffer.
         * @function decode
         * @memberof protobuf.Asset
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protobuf.Asset} Asset
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Asset.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.protobuf.Asset();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.category = reader.string();
                    break;
                case 2:
                    message.symbol = reader.string();
                    break;
                case 3:
                    message.network = reader.string();
                    break;
                case 4:
                    message.value = reader.uint64();
                    break;
                case 5:
                    message.fee = reader.uint64();
                    break;
                case 6:
                    message.nonce = reader.uint64();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an Asset message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protobuf.Asset
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protobuf.Asset} Asset
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Asset.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an Asset message.
         * @function verify
         * @memberof protobuf.Asset
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Asset.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.category != null && message.hasOwnProperty("category"))
                if (!$util.isString(message.category))
                    return "category: string expected";
            if (message.symbol != null && message.hasOwnProperty("symbol"))
                if (!$util.isString(message.symbol))
                    return "symbol: string expected";
            if (message.network != null && message.hasOwnProperty("network"))
                if (!$util.isString(message.network))
                    return "network: string expected";
            if (message.value != null && message.hasOwnProperty("value"))
                if (!$util.isInteger(message.value) && !(message.value && $util.isInteger(message.value.low) && $util.isInteger(message.value.high)))
                    return "value: integer|Long expected";
            if (message.fee != null && message.hasOwnProperty("fee"))
                if (!$util.isInteger(message.fee) && !(message.fee && $util.isInteger(message.fee.low) && $util.isInteger(message.fee.high)))
                    return "fee: integer|Long expected";
            if (message.nonce != null && message.hasOwnProperty("nonce"))
                if (!$util.isInteger(message.nonce) && !(message.nonce && $util.isInteger(message.nonce.low) && $util.isInteger(message.nonce.high)))
                    return "nonce: integer|Long expected";
            return null;
        };

        /**
         * Creates an Asset message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protobuf.Asset
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protobuf.Asset} Asset
         */
        Asset.fromObject = function fromObject(object) {
            if (object instanceof $root.protobuf.Asset)
                return object;
            var message = new $root.protobuf.Asset();
            if (object.category != null)
                message.category = String(object.category);
            if (object.symbol != null)
                message.symbol = String(object.symbol);
            if (object.network != null)
                message.network = String(object.network);
            if (object.value != null)
                if ($util.Long)
                    (message.value = $util.Long.fromValue(object.value)).unsigned = true;
                else if (typeof object.value === "string")
                    message.value = parseInt(object.value, 10);
                else if (typeof object.value === "number")
                    message.value = object.value;
                else if (typeof object.value === "object")
                    message.value = new $util.LongBits(object.value.low >>> 0, object.value.high >>> 0).toNumber(true);
            if (object.fee != null)
                if ($util.Long)
                    (message.fee = $util.Long.fromValue(object.fee)).unsigned = true;
                else if (typeof object.fee === "string")
                    message.fee = parseInt(object.fee, 10);
                else if (typeof object.fee === "number")
                    message.fee = object.fee;
                else if (typeof object.fee === "object")
                    message.fee = new $util.LongBits(object.fee.low >>> 0, object.fee.high >>> 0).toNumber(true);
            if (object.nonce != null)
                if ($util.Long)
                    (message.nonce = $util.Long.fromValue(object.nonce)).unsigned = true;
                else if (typeof object.nonce === "string")
                    message.nonce = parseInt(object.nonce, 10);
                else if (typeof object.nonce === "number")
                    message.nonce = object.nonce;
                else if (typeof object.nonce === "object")
                    message.nonce = new $util.LongBits(object.nonce.low >>> 0, object.nonce.high >>> 0).toNumber(true);
            return message;
        };

        /**
         * Creates a plain object from an Asset message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protobuf.Asset
         * @static
         * @param {protobuf.Asset} message Asset
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Asset.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.category = "";
                object.symbol = "";
                object.network = "";
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.value = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.value = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.fee = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.fee = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.nonce = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.nonce = options.longs === String ? "0" : 0;
            }
            if (message.category != null && message.hasOwnProperty("category"))
                object.category = message.category;
            if (message.symbol != null && message.hasOwnProperty("symbol"))
                object.symbol = message.symbol;
            if (message.network != null && message.hasOwnProperty("network"))
                object.network = message.network;
            if (message.value != null && message.hasOwnProperty("value"))
                if (typeof message.value === "number")
                    object.value = options.longs === String ? String(message.value) : message.value;
                else
                    object.value = options.longs === String ? $util.Long.prototype.toString.call(message.value) : options.longs === Number ? new $util.LongBits(message.value.low >>> 0, message.value.high >>> 0).toNumber(true) : message.value;
            if (message.fee != null && message.hasOwnProperty("fee"))
                if (typeof message.fee === "number")
                    object.fee = options.longs === String ? String(message.fee) : message.fee;
                else
                    object.fee = options.longs === String ? $util.Long.prototype.toString.call(message.fee) : options.longs === Number ? new $util.LongBits(message.fee.low >>> 0, message.fee.high >>> 0).toNumber(true) : message.fee;
            if (message.nonce != null && message.hasOwnProperty("nonce"))
                if (typeof message.nonce === "number")
                    object.nonce = options.longs === String ? String(message.nonce) : message.nonce;
                else
                    object.nonce = options.longs === String ? $util.Long.prototype.toString.call(message.nonce) : options.longs === Number ? new $util.LongBits(message.nonce.low >>> 0, message.nonce.high >>> 0).toNumber(true) : message.nonce;
            return object;
        };

        /**
         * Converts this Asset to JSON.
         * @function toJSON
         * @memberof protobuf.Asset
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Asset.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Asset;
    })();

    protobuf.TxRequest = (function() {

        /**
         * Properties of a TxRequest.
         * @memberof protobuf
         * @interface ITxRequest
         * @property {protobuf.ITx|null} [tx] TxRequest tx
         */

        /**
         * Constructs a new TxRequest.
         * @memberof protobuf
         * @classdesc Represents a TxRequest.
         * @implements ITxRequest
         * @constructor
         * @param {protobuf.ITxRequest=} [properties] Properties to set
         */
        function TxRequest(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * TxRequest tx.
         * @member {protobuf.ITx|null|undefined} tx
         * @memberof protobuf.TxRequest
         * @instance
         */
        TxRequest.prototype.tx = null;

        /**
         * Creates a new TxRequest instance using the specified properties.
         * @function create
         * @memberof protobuf.TxRequest
         * @static
         * @param {protobuf.ITxRequest=} [properties] Properties to set
         * @returns {protobuf.TxRequest} TxRequest instance
         */
        TxRequest.create = function create(properties) {
            return new TxRequest(properties);
        };

        /**
         * Encodes the specified TxRequest message. Does not implicitly {@link protobuf.TxRequest.verify|verify} messages.
         * @function encode
         * @memberof protobuf.TxRequest
         * @static
         * @param {protobuf.ITxRequest} message TxRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        TxRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.tx != null && message.hasOwnProperty("tx"))
                $root.protobuf.Tx.encode(message.tx, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified TxRequest message, length delimited. Does not implicitly {@link protobuf.TxRequest.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protobuf.TxRequest
         * @static
         * @param {protobuf.ITxRequest} message TxRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        TxRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a TxRequest message from the specified reader or buffer.
         * @function decode
         * @memberof protobuf.TxRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protobuf.TxRequest} TxRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        TxRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.protobuf.TxRequest();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.tx = $root.protobuf.Tx.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a TxRequest message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protobuf.TxRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protobuf.TxRequest} TxRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        TxRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a TxRequest message.
         * @function verify
         * @memberof protobuf.TxRequest
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        TxRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.tx != null && message.hasOwnProperty("tx")) {
                var error = $root.protobuf.Tx.verify(message.tx);
                if (error)
                    return "tx." + error;
            }
            return null;
        };

        /**
         * Creates a TxRequest message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protobuf.TxRequest
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protobuf.TxRequest} TxRequest
         */
        TxRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.protobuf.TxRequest)
                return object;
            var message = new $root.protobuf.TxRequest();
            if (object.tx != null) {
                if (typeof object.tx !== "object")
                    throw TypeError(".protobuf.TxRequest.tx: object expected");
                message.tx = $root.protobuf.Tx.fromObject(object.tx);
            }
            return message;
        };

        /**
         * Creates a plain object from a TxRequest message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protobuf.TxRequest
         * @static
         * @param {protobuf.TxRequest} message TxRequest
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        TxRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.tx = null;
            if (message.tx != null && message.hasOwnProperty("tx"))
                object.tx = $root.protobuf.Tx.toObject(message.tx, options);
            return object;
        };

        /**
         * Converts this TxRequest to JSON.
         * @function toJSON
         * @memberof protobuf.TxRequest
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        TxRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return TxRequest;
    })();

    protobuf.TxResponse = (function() {

        /**
         * Properties of a TxResponse.
         * @memberof protobuf
         * @interface ITxResponse
         * @property {string|null} [txId] TxResponse txId
         * @property {number|Long|null} [pending] TxResponse pending
         * @property {number|Long|null} [queued] TxResponse queued
         * @property {string|null} [status] TxResponse status
         */

        /**
         * Constructs a new TxResponse.
         * @memberof protobuf
         * @classdesc Represents a TxResponse.
         * @implements ITxResponse
         * @constructor
         * @param {protobuf.ITxResponse=} [properties] Properties to set
         */
        function TxResponse(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * TxResponse txId.
         * @member {string} txId
         * @memberof protobuf.TxResponse
         * @instance
         */
        TxResponse.prototype.txId = "";

        /**
         * TxResponse pending.
         * @member {number|Long} pending
         * @memberof protobuf.TxResponse
         * @instance
         */
        TxResponse.prototype.pending = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * TxResponse queued.
         * @member {number|Long} queued
         * @memberof protobuf.TxResponse
         * @instance
         */
        TxResponse.prototype.queued = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * TxResponse status.
         * @member {string} status
         * @memberof protobuf.TxResponse
         * @instance
         */
        TxResponse.prototype.status = "";

        /**
         * Creates a new TxResponse instance using the specified properties.
         * @function create
         * @memberof protobuf.TxResponse
         * @static
         * @param {protobuf.ITxResponse=} [properties] Properties to set
         * @returns {protobuf.TxResponse} TxResponse instance
         */
        TxResponse.create = function create(properties) {
            return new TxResponse(properties);
        };

        /**
         * Encodes the specified TxResponse message. Does not implicitly {@link protobuf.TxResponse.verify|verify} messages.
         * @function encode
         * @memberof protobuf.TxResponse
         * @static
         * @param {protobuf.ITxResponse} message TxResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        TxResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.txId != null && message.hasOwnProperty("txId"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.txId);
            if (message.pending != null && message.hasOwnProperty("pending"))
                writer.uint32(/* id 2, wireType 0 =*/16).int64(message.pending);
            if (message.queued != null && message.hasOwnProperty("queued"))
                writer.uint32(/* id 3, wireType 0 =*/24).int64(message.queued);
            if (message.status != null && message.hasOwnProperty("status"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.status);
            return writer;
        };

        /**
         * Encodes the specified TxResponse message, length delimited. Does not implicitly {@link protobuf.TxResponse.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protobuf.TxResponse
         * @static
         * @param {protobuf.ITxResponse} message TxResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        TxResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a TxResponse message from the specified reader or buffer.
         * @function decode
         * @memberof protobuf.TxResponse
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protobuf.TxResponse} TxResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        TxResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.protobuf.TxResponse();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.txId = reader.string();
                    break;
                case 2:
                    message.pending = reader.int64();
                    break;
                case 3:
                    message.queued = reader.int64();
                    break;
                case 4:
                    message.status = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a TxResponse message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protobuf.TxResponse
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protobuf.TxResponse} TxResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        TxResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a TxResponse message.
         * @function verify
         * @memberof protobuf.TxResponse
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        TxResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.txId != null && message.hasOwnProperty("txId"))
                if (!$util.isString(message.txId))
                    return "txId: string expected";
            if (message.pending != null && message.hasOwnProperty("pending"))
                if (!$util.isInteger(message.pending) && !(message.pending && $util.isInteger(message.pending.low) && $util.isInteger(message.pending.high)))
                    return "pending: integer|Long expected";
            if (message.queued != null && message.hasOwnProperty("queued"))
                if (!$util.isInteger(message.queued) && !(message.queued && $util.isInteger(message.queued.low) && $util.isInteger(message.queued.high)))
                    return "queued: integer|Long expected";
            if (message.status != null && message.hasOwnProperty("status"))
                if (!$util.isString(message.status))
                    return "status: string expected";
            return null;
        };

        /**
         * Creates a TxResponse message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protobuf.TxResponse
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protobuf.TxResponse} TxResponse
         */
        TxResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.protobuf.TxResponse)
                return object;
            var message = new $root.protobuf.TxResponse();
            if (object.txId != null)
                message.txId = String(object.txId);
            if (object.pending != null)
                if ($util.Long)
                    (message.pending = $util.Long.fromValue(object.pending)).unsigned = false;
                else if (typeof object.pending === "string")
                    message.pending = parseInt(object.pending, 10);
                else if (typeof object.pending === "number")
                    message.pending = object.pending;
                else if (typeof object.pending === "object")
                    message.pending = new $util.LongBits(object.pending.low >>> 0, object.pending.high >>> 0).toNumber();
            if (object.queued != null)
                if ($util.Long)
                    (message.queued = $util.Long.fromValue(object.queued)).unsigned = false;
                else if (typeof object.queued === "string")
                    message.queued = parseInt(object.queued, 10);
                else if (typeof object.queued === "number")
                    message.queued = object.queued;
                else if (typeof object.queued === "object")
                    message.queued = new $util.LongBits(object.queued.low >>> 0, object.queued.high >>> 0).toNumber();
            if (object.status != null)
                message.status = String(object.status);
            return message;
        };

        /**
         * Creates a plain object from a TxResponse message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protobuf.TxResponse
         * @static
         * @param {protobuf.TxResponse} message TxResponse
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        TxResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.txId = "";
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.pending = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.pending = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.queued = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.queued = options.longs === String ? "0" : 0;
                object.status = "";
            }
            if (message.txId != null && message.hasOwnProperty("txId"))
                object.txId = message.txId;
            if (message.pending != null && message.hasOwnProperty("pending"))
                if (typeof message.pending === "number")
                    object.pending = options.longs === String ? String(message.pending) : message.pending;
                else
                    object.pending = options.longs === String ? $util.Long.prototype.toString.call(message.pending) : options.longs === Number ? new $util.LongBits(message.pending.low >>> 0, message.pending.high >>> 0).toNumber() : message.pending;
            if (message.queued != null && message.hasOwnProperty("queued"))
                if (typeof message.queued === "number")
                    object.queued = options.longs === String ? String(message.queued) : message.queued;
                else
                    object.queued = options.longs === String ? $util.Long.prototype.toString.call(message.queued) : options.longs === Number ? new $util.LongBits(message.queued.low >>> 0, message.queued.high >>> 0).toNumber() : message.queued;
            if (message.status != null && message.hasOwnProperty("status"))
                object.status = message.status;
            return object;
        };

        /**
         * Converts this TxResponse to JSON.
         * @function toJSON
         * @memberof protobuf.TxResponse
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        TxResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return TxResponse;
    })();

    protobuf.AccountRegisterRequest = (function() {

        /**
         * Properties of an AccountRegisterRequest.
         * @memberof protobuf
         * @interface IAccountRegisterRequest
         * @property {string|null} [senderPubkey] AccountRegisterRequest senderPubkey
         */

        /**
         * Constructs a new AccountRegisterRequest.
         * @memberof protobuf
         * @classdesc Represents an AccountRegisterRequest.
         * @implements IAccountRegisterRequest
         * @constructor
         * @param {protobuf.IAccountRegisterRequest=} [properties] Properties to set
         */
        function AccountRegisterRequest(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * AccountRegisterRequest senderPubkey.
         * @member {string} senderPubkey
         * @memberof protobuf.AccountRegisterRequest
         * @instance
         */
        AccountRegisterRequest.prototype.senderPubkey = "";

        /**
         * Creates a new AccountRegisterRequest instance using the specified properties.
         * @function create
         * @memberof protobuf.AccountRegisterRequest
         * @static
         * @param {protobuf.IAccountRegisterRequest=} [properties] Properties to set
         * @returns {protobuf.AccountRegisterRequest} AccountRegisterRequest instance
         */
        AccountRegisterRequest.create = function create(properties) {
            return new AccountRegisterRequest(properties);
        };

        /**
         * Encodes the specified AccountRegisterRequest message. Does not implicitly {@link protobuf.AccountRegisterRequest.verify|verify} messages.
         * @function encode
         * @memberof protobuf.AccountRegisterRequest
         * @static
         * @param {protobuf.IAccountRegisterRequest} message AccountRegisterRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        AccountRegisterRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.senderPubkey != null && message.hasOwnProperty("senderPubkey"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.senderPubkey);
            return writer;
        };

        /**
         * Encodes the specified AccountRegisterRequest message, length delimited. Does not implicitly {@link protobuf.AccountRegisterRequest.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protobuf.AccountRegisterRequest
         * @static
         * @param {protobuf.IAccountRegisterRequest} message AccountRegisterRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        AccountRegisterRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an AccountRegisterRequest message from the specified reader or buffer.
         * @function decode
         * @memberof protobuf.AccountRegisterRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protobuf.AccountRegisterRequest} AccountRegisterRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        AccountRegisterRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.protobuf.AccountRegisterRequest();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.senderPubkey = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an AccountRegisterRequest message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protobuf.AccountRegisterRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protobuf.AccountRegisterRequest} AccountRegisterRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        AccountRegisterRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an AccountRegisterRequest message.
         * @function verify
         * @memberof protobuf.AccountRegisterRequest
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        AccountRegisterRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.senderPubkey != null && message.hasOwnProperty("senderPubkey"))
                if (!$util.isString(message.senderPubkey))
                    return "senderPubkey: string expected";
            return null;
        };

        /**
         * Creates an AccountRegisterRequest message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protobuf.AccountRegisterRequest
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protobuf.AccountRegisterRequest} AccountRegisterRequest
         */
        AccountRegisterRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.protobuf.AccountRegisterRequest)
                return object;
            var message = new $root.protobuf.AccountRegisterRequest();
            if (object.senderPubkey != null)
                message.senderPubkey = String(object.senderPubkey);
            return message;
        };

        /**
         * Creates a plain object from an AccountRegisterRequest message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protobuf.AccountRegisterRequest
         * @static
         * @param {protobuf.AccountRegisterRequest} message AccountRegisterRequest
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        AccountRegisterRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.senderPubkey = "";
            if (message.senderPubkey != null && message.hasOwnProperty("senderPubkey"))
                object.senderPubkey = message.senderPubkey;
            return object;
        };

        /**
         * Converts this AccountRegisterRequest to JSON.
         * @function toJSON
         * @memberof protobuf.AccountRegisterRequest
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        AccountRegisterRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return AccountRegisterRequest;
    })();

    protobuf.AccountRegisterResponse = (function() {

        /**
         * Properties of an AccountRegisterResponse.
         * @memberof protobuf
         * @interface IAccountRegisterResponse
         * @property {boolean|null} [status] AccountRegisterResponse status
         */

        /**
         * Constructs a new AccountRegisterResponse.
         * @memberof protobuf
         * @classdesc Represents an AccountRegisterResponse.
         * @implements IAccountRegisterResponse
         * @constructor
         * @param {protobuf.IAccountRegisterResponse=} [properties] Properties to set
         */
        function AccountRegisterResponse(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * AccountRegisterResponse status.
         * @member {boolean} status
         * @memberof protobuf.AccountRegisterResponse
         * @instance
         */
        AccountRegisterResponse.prototype.status = false;

        /**
         * Creates a new AccountRegisterResponse instance using the specified properties.
         * @function create
         * @memberof protobuf.AccountRegisterResponse
         * @static
         * @param {protobuf.IAccountRegisterResponse=} [properties] Properties to set
         * @returns {protobuf.AccountRegisterResponse} AccountRegisterResponse instance
         */
        AccountRegisterResponse.create = function create(properties) {
            return new AccountRegisterResponse(properties);
        };

        /**
         * Encodes the specified AccountRegisterResponse message. Does not implicitly {@link protobuf.AccountRegisterResponse.verify|verify} messages.
         * @function encode
         * @memberof protobuf.AccountRegisterResponse
         * @static
         * @param {protobuf.IAccountRegisterResponse} message AccountRegisterResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        AccountRegisterResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.status != null && message.hasOwnProperty("status"))
                writer.uint32(/* id 1, wireType 0 =*/8).bool(message.status);
            return writer;
        };

        /**
         * Encodes the specified AccountRegisterResponse message, length delimited. Does not implicitly {@link protobuf.AccountRegisterResponse.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protobuf.AccountRegisterResponse
         * @static
         * @param {protobuf.IAccountRegisterResponse} message AccountRegisterResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        AccountRegisterResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an AccountRegisterResponse message from the specified reader or buffer.
         * @function decode
         * @memberof protobuf.AccountRegisterResponse
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protobuf.AccountRegisterResponse} AccountRegisterResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        AccountRegisterResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.protobuf.AccountRegisterResponse();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.status = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an AccountRegisterResponse message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protobuf.AccountRegisterResponse
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protobuf.AccountRegisterResponse} AccountRegisterResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        AccountRegisterResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an AccountRegisterResponse message.
         * @function verify
         * @memberof protobuf.AccountRegisterResponse
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        AccountRegisterResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.status != null && message.hasOwnProperty("status"))
                if (typeof message.status !== "boolean")
                    return "status: boolean expected";
            return null;
        };

        /**
         * Creates an AccountRegisterResponse message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protobuf.AccountRegisterResponse
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protobuf.AccountRegisterResponse} AccountRegisterResponse
         */
        AccountRegisterResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.protobuf.AccountRegisterResponse)
                return object;
            var message = new $root.protobuf.AccountRegisterResponse();
            if (object.status != null)
                message.status = Boolean(object.status);
            return message;
        };

        /**
         * Creates a plain object from an AccountRegisterResponse message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protobuf.AccountRegisterResponse
         * @static
         * @param {protobuf.AccountRegisterResponse} message AccountRegisterResponse
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        AccountRegisterResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.status = false;
            if (message.status != null && message.hasOwnProperty("status"))
                object.status = message.status;
            return object;
        };

        /**
         * Converts this AccountRegisterResponse to JSON.
         * @function toJSON
         * @memberof protobuf.AccountRegisterResponse
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        AccountRegisterResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return AccountRegisterResponse;
    })();

    protobuf.TxDetailRequest = (function() {

        /**
         * Properties of a TxDetailRequest.
         * @memberof protobuf
         * @interface ITxDetailRequest
         * @property {string|null} [txId] TxDetailRequest txId
         */

        /**
         * Constructs a new TxDetailRequest.
         * @memberof protobuf
         * @classdesc Represents a TxDetailRequest.
         * @implements ITxDetailRequest
         * @constructor
         * @param {protobuf.ITxDetailRequest=} [properties] Properties to set
         */
        function TxDetailRequest(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * TxDetailRequest txId.
         * @member {string} txId
         * @memberof protobuf.TxDetailRequest
         * @instance
         */
        TxDetailRequest.prototype.txId = "";

        /**
         * Creates a new TxDetailRequest instance using the specified properties.
         * @function create
         * @memberof protobuf.TxDetailRequest
         * @static
         * @param {protobuf.ITxDetailRequest=} [properties] Properties to set
         * @returns {protobuf.TxDetailRequest} TxDetailRequest instance
         */
        TxDetailRequest.create = function create(properties) {
            return new TxDetailRequest(properties);
        };

        /**
         * Encodes the specified TxDetailRequest message. Does not implicitly {@link protobuf.TxDetailRequest.verify|verify} messages.
         * @function encode
         * @memberof protobuf.TxDetailRequest
         * @static
         * @param {protobuf.ITxDetailRequest} message TxDetailRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        TxDetailRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.txId != null && message.hasOwnProperty("txId"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.txId);
            return writer;
        };

        /**
         * Encodes the specified TxDetailRequest message, length delimited. Does not implicitly {@link protobuf.TxDetailRequest.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protobuf.TxDetailRequest
         * @static
         * @param {protobuf.ITxDetailRequest} message TxDetailRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        TxDetailRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a TxDetailRequest message from the specified reader or buffer.
         * @function decode
         * @memberof protobuf.TxDetailRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protobuf.TxDetailRequest} TxDetailRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        TxDetailRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.protobuf.TxDetailRequest();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.txId = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a TxDetailRequest message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protobuf.TxDetailRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protobuf.TxDetailRequest} TxDetailRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        TxDetailRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a TxDetailRequest message.
         * @function verify
         * @memberof protobuf.TxDetailRequest
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        TxDetailRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.txId != null && message.hasOwnProperty("txId"))
                if (!$util.isString(message.txId))
                    return "txId: string expected";
            return null;
        };

        /**
         * Creates a TxDetailRequest message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protobuf.TxDetailRequest
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protobuf.TxDetailRequest} TxDetailRequest
         */
        TxDetailRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.protobuf.TxDetailRequest)
                return object;
            var message = new $root.protobuf.TxDetailRequest();
            if (object.txId != null)
                message.txId = String(object.txId);
            return message;
        };

        /**
         * Creates a plain object from a TxDetailRequest message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protobuf.TxDetailRequest
         * @static
         * @param {protobuf.TxDetailRequest} message TxDetailRequest
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        TxDetailRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.txId = "";
            if (message.txId != null && message.hasOwnProperty("txId"))
                object.txId = message.txId;
            return object;
        };

        /**
         * Converts this TxDetailRequest to JSON.
         * @function toJSON
         * @memberof protobuf.TxDetailRequest
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        TxDetailRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return TxDetailRequest;
    })();

    protobuf.TxDetailResponse = (function() {

        /**
         * Properties of a TxDetailResponse.
         * @memberof protobuf
         * @interface ITxDetailResponse
         * @property {string|null} [txId] TxDetailResponse txId
         * @property {protobuf.ITx|null} [tx] TxDetailResponse tx
         * @property {protobuf.ITimestamp|null} [creationDt] TxDetailResponse creationDt
         * @property {number|Long|null} [blockId] TxDetailResponse blockId
         */

        /**
         * Constructs a new TxDetailResponse.
         * @memberof protobuf
         * @classdesc Represents a TxDetailResponse.
         * @implements ITxDetailResponse
         * @constructor
         * @param {protobuf.ITxDetailResponse=} [properties] Properties to set
         */
        function TxDetailResponse(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * TxDetailResponse txId.
         * @member {string} txId
         * @memberof protobuf.TxDetailResponse
         * @instance
         */
        TxDetailResponse.prototype.txId = "";

        /**
         * TxDetailResponse tx.
         * @member {protobuf.ITx|null|undefined} tx
         * @memberof protobuf.TxDetailResponse
         * @instance
         */
        TxDetailResponse.prototype.tx = null;

        /**
         * TxDetailResponse creationDt.
         * @member {protobuf.ITimestamp|null|undefined} creationDt
         * @memberof protobuf.TxDetailResponse
         * @instance
         */
        TxDetailResponse.prototype.creationDt = null;

        /**
         * TxDetailResponse blockId.
         * @member {number|Long} blockId
         * @memberof protobuf.TxDetailResponse
         * @instance
         */
        TxDetailResponse.prototype.blockId = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * Creates a new TxDetailResponse instance using the specified properties.
         * @function create
         * @memberof protobuf.TxDetailResponse
         * @static
         * @param {protobuf.ITxDetailResponse=} [properties] Properties to set
         * @returns {protobuf.TxDetailResponse} TxDetailResponse instance
         */
        TxDetailResponse.create = function create(properties) {
            return new TxDetailResponse(properties);
        };

        /**
         * Encodes the specified TxDetailResponse message. Does not implicitly {@link protobuf.TxDetailResponse.verify|verify} messages.
         * @function encode
         * @memberof protobuf.TxDetailResponse
         * @static
         * @param {protobuf.ITxDetailResponse} message TxDetailResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        TxDetailResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.txId != null && message.hasOwnProperty("txId"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.txId);
            if (message.tx != null && message.hasOwnProperty("tx"))
                $root.protobuf.Tx.encode(message.tx, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            if (message.creationDt != null && message.hasOwnProperty("creationDt"))
                $root.protobuf.Timestamp.encode(message.creationDt, writer.uint32(/* id 3, wireType 2 =*/26).fork()).ldelim();
            if (message.blockId != null && message.hasOwnProperty("blockId"))
                writer.uint32(/* id 4, wireType 0 =*/32).uint64(message.blockId);
            return writer;
        };

        /**
         * Encodes the specified TxDetailResponse message, length delimited. Does not implicitly {@link protobuf.TxDetailResponse.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protobuf.TxDetailResponse
         * @static
         * @param {protobuf.ITxDetailResponse} message TxDetailResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        TxDetailResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a TxDetailResponse message from the specified reader or buffer.
         * @function decode
         * @memberof protobuf.TxDetailResponse
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protobuf.TxDetailResponse} TxDetailResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        TxDetailResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.protobuf.TxDetailResponse();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.txId = reader.string();
                    break;
                case 2:
                    message.tx = $root.protobuf.Tx.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.creationDt = $root.protobuf.Timestamp.decode(reader, reader.uint32());
                    break;
                case 4:
                    message.blockId = reader.uint64();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a TxDetailResponse message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protobuf.TxDetailResponse
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protobuf.TxDetailResponse} TxDetailResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        TxDetailResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a TxDetailResponse message.
         * @function verify
         * @memberof protobuf.TxDetailResponse
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        TxDetailResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.txId != null && message.hasOwnProperty("txId"))
                if (!$util.isString(message.txId))
                    return "txId: string expected";
            if (message.tx != null && message.hasOwnProperty("tx")) {
                var error = $root.protobuf.Tx.verify(message.tx);
                if (error)
                    return "tx." + error;
            }
            if (message.creationDt != null && message.hasOwnProperty("creationDt")) {
                var error = $root.protobuf.Timestamp.verify(message.creationDt);
                if (error)
                    return "creationDt." + error;
            }
            if (message.blockId != null && message.hasOwnProperty("blockId"))
                if (!$util.isInteger(message.blockId) && !(message.blockId && $util.isInteger(message.blockId.low) && $util.isInteger(message.blockId.high)))
                    return "blockId: integer|Long expected";
            return null;
        };

        /**
         * Creates a TxDetailResponse message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protobuf.TxDetailResponse
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protobuf.TxDetailResponse} TxDetailResponse
         */
        TxDetailResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.protobuf.TxDetailResponse)
                return object;
            var message = new $root.protobuf.TxDetailResponse();
            if (object.txId != null)
                message.txId = String(object.txId);
            if (object.tx != null) {
                if (typeof object.tx !== "object")
                    throw TypeError(".protobuf.TxDetailResponse.tx: object expected");
                message.tx = $root.protobuf.Tx.fromObject(object.tx);
            }
            if (object.creationDt != null) {
                if (typeof object.creationDt !== "object")
                    throw TypeError(".protobuf.TxDetailResponse.creationDt: object expected");
                message.creationDt = $root.protobuf.Timestamp.fromObject(object.creationDt);
            }
            if (object.blockId != null)
                if ($util.Long)
                    (message.blockId = $util.Long.fromValue(object.blockId)).unsigned = true;
                else if (typeof object.blockId === "string")
                    message.blockId = parseInt(object.blockId, 10);
                else if (typeof object.blockId === "number")
                    message.blockId = object.blockId;
                else if (typeof object.blockId === "object")
                    message.blockId = new $util.LongBits(object.blockId.low >>> 0, object.blockId.high >>> 0).toNumber(true);
            return message;
        };

        /**
         * Creates a plain object from a TxDetailResponse message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protobuf.TxDetailResponse
         * @static
         * @param {protobuf.TxDetailResponse} message TxDetailResponse
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        TxDetailResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.txId = "";
                object.tx = null;
                object.creationDt = null;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.blockId = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.blockId = options.longs === String ? "0" : 0;
            }
            if (message.txId != null && message.hasOwnProperty("txId"))
                object.txId = message.txId;
            if (message.tx != null && message.hasOwnProperty("tx"))
                object.tx = $root.protobuf.Tx.toObject(message.tx, options);
            if (message.creationDt != null && message.hasOwnProperty("creationDt"))
                object.creationDt = $root.protobuf.Timestamp.toObject(message.creationDt, options);
            if (message.blockId != null && message.hasOwnProperty("blockId"))
                if (typeof message.blockId === "number")
                    object.blockId = options.longs === String ? String(message.blockId) : message.blockId;
                else
                    object.blockId = options.longs === String ? $util.Long.prototype.toString.call(message.blockId) : options.longs === Number ? new $util.LongBits(message.blockId.low >>> 0, message.blockId.high >>> 0).toNumber(true) : message.blockId;
            return object;
        };

        /**
         * Converts this TxDetailResponse to JSON.
         * @function toJSON
         * @memberof protobuf.TxDetailResponse
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        TxDetailResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return TxDetailResponse;
    })();

    return protobuf;
})();

module.exports = $root;
