import * as $protobuf from "protobufjs";
// Common aliases
const $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
const $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

export const rental = $root.rental = (() => {

    /**
     * Namespace rental.
     * @exports rental
     * @namespace
     */
    const rental = {};

    rental.v1 = (function() {

        /**
         * Namespace v1.
         * @memberof rental
         * @namespace
         */
        const v1 = {};

        v1.CreateTripRequest = (function() {

            /**
             * Properties of a CreateTripRequest.
             * @memberof rental.v1
             * @interface ICreateTripRequest
             * @property {string|null} [start] CreateTripRequest start
             */

            /**
             * Constructs a new CreateTripRequest.
             * @memberof rental.v1
             * @classdesc Represents a CreateTripRequest.
             * @implements ICreateTripRequest
             * @constructor
             * @param {rental.v1.ICreateTripRequest=} [properties] Properties to set
             */
            function CreateTripRequest(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * CreateTripRequest start.
             * @member {string} start
             * @memberof rental.v1.CreateTripRequest
             * @instance
             */
            CreateTripRequest.prototype.start = "";

            /**
             * Creates a new CreateTripRequest instance using the specified properties.
             * @function create
             * @memberof rental.v1.CreateTripRequest
             * @static
             * @param {rental.v1.ICreateTripRequest=} [properties] Properties to set
             * @returns {rental.v1.CreateTripRequest} CreateTripRequest instance
             */
            CreateTripRequest.create = function create(properties) {
                return new CreateTripRequest(properties);
            };

            /**
             * Encodes the specified CreateTripRequest message. Does not implicitly {@link rental.v1.CreateTripRequest.verify|verify} messages.
             * @function encode
             * @memberof rental.v1.CreateTripRequest
             * @static
             * @param {rental.v1.ICreateTripRequest} message CreateTripRequest message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            CreateTripRequest.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.start != null && Object.hasOwnProperty.call(message, "start"))
                    writer.uint32(/* id 1, wireType 2 =*/10).string(message.start);
                return writer;
            };

            /**
             * Encodes the specified CreateTripRequest message, length delimited. Does not implicitly {@link rental.v1.CreateTripRequest.verify|verify} messages.
             * @function encodeDelimited
             * @memberof rental.v1.CreateTripRequest
             * @static
             * @param {rental.v1.ICreateTripRequest} message CreateTripRequest message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            CreateTripRequest.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a CreateTripRequest message from the specified reader or buffer.
             * @function decode
             * @memberof rental.v1.CreateTripRequest
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {rental.v1.CreateTripRequest} CreateTripRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            CreateTripRequest.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.rental.v1.CreateTripRequest();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.start = reader.string();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a CreateTripRequest message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof rental.v1.CreateTripRequest
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {rental.v1.CreateTripRequest} CreateTripRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            CreateTripRequest.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a CreateTripRequest message.
             * @function verify
             * @memberof rental.v1.CreateTripRequest
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            CreateTripRequest.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.start != null && message.hasOwnProperty("start"))
                    if (!$util.isString(message.start))
                        return "start: string expected";
                return null;
            };

            /**
             * Creates a CreateTripRequest message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.CreateTripRequest
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.CreateTripRequest} CreateTripRequest
             */
            CreateTripRequest.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.CreateTripRequest)
                    return object;
                let message = new $root.rental.v1.CreateTripRequest();
                if (object.start != null)
                    message.start = String(object.start);
                return message;
            };

            /**
             * Creates a plain object from a CreateTripRequest message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.CreateTripRequest
             * @static
             * @param {rental.v1.CreateTripRequest} message CreateTripRequest
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            CreateTripRequest.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults)
                    object.start = "";
                if (message.start != null && message.hasOwnProperty("start"))
                    object.start = message.start;
                return object;
            };

            /**
             * Converts this CreateTripRequest to JSON.
             * @function toJSON
             * @memberof rental.v1.CreateTripRequest
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            CreateTripRequest.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return CreateTripRequest;
        })();

        v1.CreateTripResponse = (function() {

            /**
             * Properties of a CreateTripResponse.
             * @memberof rental.v1
             * @interface ICreateTripResponse
             */

            /**
             * Constructs a new CreateTripResponse.
             * @memberof rental.v1
             * @classdesc Represents a CreateTripResponse.
             * @implements ICreateTripResponse
             * @constructor
             * @param {rental.v1.ICreateTripResponse=} [properties] Properties to set
             */
            function CreateTripResponse(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * Creates a new CreateTripResponse instance using the specified properties.
             * @function create
             * @memberof rental.v1.CreateTripResponse
             * @static
             * @param {rental.v1.ICreateTripResponse=} [properties] Properties to set
             * @returns {rental.v1.CreateTripResponse} CreateTripResponse instance
             */
            CreateTripResponse.create = function create(properties) {
                return new CreateTripResponse(properties);
            };

            /**
             * Encodes the specified CreateTripResponse message. Does not implicitly {@link rental.v1.CreateTripResponse.verify|verify} messages.
             * @function encode
             * @memberof rental.v1.CreateTripResponse
             * @static
             * @param {rental.v1.ICreateTripResponse} message CreateTripResponse message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            CreateTripResponse.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                return writer;
            };

            /**
             * Encodes the specified CreateTripResponse message, length delimited. Does not implicitly {@link rental.v1.CreateTripResponse.verify|verify} messages.
             * @function encodeDelimited
             * @memberof rental.v1.CreateTripResponse
             * @static
             * @param {rental.v1.ICreateTripResponse} message CreateTripResponse message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            CreateTripResponse.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a CreateTripResponse message from the specified reader or buffer.
             * @function decode
             * @memberof rental.v1.CreateTripResponse
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {rental.v1.CreateTripResponse} CreateTripResponse
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            CreateTripResponse.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.rental.v1.CreateTripResponse();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a CreateTripResponse message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof rental.v1.CreateTripResponse
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {rental.v1.CreateTripResponse} CreateTripResponse
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            CreateTripResponse.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a CreateTripResponse message.
             * @function verify
             * @memberof rental.v1.CreateTripResponse
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            CreateTripResponse.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                return null;
            };

            /**
             * Creates a CreateTripResponse message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.CreateTripResponse
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.CreateTripResponse} CreateTripResponse
             */
            CreateTripResponse.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.CreateTripResponse)
                    return object;
                return new $root.rental.v1.CreateTripResponse();
            };

            /**
             * Creates a plain object from a CreateTripResponse message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.CreateTripResponse
             * @static
             * @param {rental.v1.CreateTripResponse} message CreateTripResponse
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            CreateTripResponse.toObject = function toObject() {
                return {};
            };

            /**
             * Converts this CreateTripResponse to JSON.
             * @function toJSON
             * @memberof rental.v1.CreateTripResponse
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            CreateTripResponse.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return CreateTripResponse;
        })();

        v1.TripService = (function() {

            /**
             * Constructs a new TripService service.
             * @memberof rental.v1
             * @classdesc Represents a TripService
             * @extends $protobuf.rpc.Service
             * @constructor
             * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
             * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
             * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
             */
            function TripService(rpcImpl, requestDelimited, responseDelimited) {
                $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
            }

            (TripService.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = TripService;

            /**
             * Creates new TripService service using the specified rpc implementation.
             * @function create
             * @memberof rental.v1.TripService
             * @static
             * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
             * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
             * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
             * @returns {TripService} RPC service. Useful where requests and/or responses are streamed.
             */
            TripService.create = function create(rpcImpl, requestDelimited, responseDelimited) {
                return new this(rpcImpl, requestDelimited, responseDelimited);
            };

            /**
             * Callback as used by {@link rental.v1.TripService#createTrip}.
             * @memberof rental.v1.TripService
             * @typedef CreateTripCallback
             * @type {function}
             * @param {Error|null} error Error, if any
             * @param {rental.v1.CreateTripResponse} [response] CreateTripResponse
             */

            /**
             * Calls CreateTrip.
             * @function createTrip
             * @memberof rental.v1.TripService
             * @instance
             * @param {rental.v1.ICreateTripRequest} request CreateTripRequest message or plain object
             * @param {rental.v1.TripService.CreateTripCallback} callback Node-style callback called with the error, if any, and CreateTripResponse
             * @returns {undefined}
             * @variation 1
             */
            Object.defineProperty(TripService.prototype.createTrip = function createTrip(request, callback) {
                return this.rpcCall(createTrip, $root.rental.v1.CreateTripRequest, $root.rental.v1.CreateTripResponse, request, callback);
            }, "name", { value: "CreateTrip" });

            /**
             * Calls CreateTrip.
             * @function createTrip
             * @memberof rental.v1.TripService
             * @instance
             * @param {rental.v1.ICreateTripRequest} request CreateTripRequest message or plain object
             * @returns {Promise<rental.v1.CreateTripResponse>} Promise
             * @variation 2
             */

            return TripService;
        })();

        return v1;
    })();

    return rental;
})();