import * as $protobuf from "protobufjs";
// Common aliases
const $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
const $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

export const coolar = $root.coolar = (() => {

    /**
     * Namespace coolar.
     * @exports coolar
     * @namespace
     */
    const coolar = {};

    coolar.Location = (function() {

        /**
         * Properties of a Location.
         * @memberof coolar
         * @interface ILocation
         * @property {number|null} [latitude] Location latitude
         * @property {number|null} [longtitude] Location longtitude
         */

        /**
         * Constructs a new Location.
         * @memberof coolar
         * @classdesc Represents a Location.
         * @implements ILocation
         * @constructor
         * @param {coolar.ILocation=} [properties] Properties to set
         */
        function Location(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Location latitude.
         * @member {number} latitude
         * @memberof coolar.Location
         * @instance
         */
        Location.prototype.latitude = 0;

        /**
         * Location longtitude.
         * @member {number} longtitude
         * @memberof coolar.Location
         * @instance
         */
        Location.prototype.longtitude = 0;

        /**
         * Creates a new Location instance using the specified properties.
         * @function create
         * @memberof coolar.Location
         * @static
         * @param {coolar.ILocation=} [properties] Properties to set
         * @returns {coolar.Location} Location instance
         */
        Location.create = function create(properties) {
            return new Location(properties);
        };

        /**
         * Encodes the specified Location message. Does not implicitly {@link coolar.Location.verify|verify} messages.
         * @function encode
         * @memberof coolar.Location
         * @static
         * @param {coolar.ILocation} message Location message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Location.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.latitude != null && Object.hasOwnProperty.call(message, "latitude"))
                writer.uint32(/* id 1, wireType 1 =*/9).double(message.latitude);
            if (message.longtitude != null && Object.hasOwnProperty.call(message, "longtitude"))
                writer.uint32(/* id 2, wireType 1 =*/17).double(message.longtitude);
            return writer;
        };

        /**
         * Encodes the specified Location message, length delimited. Does not implicitly {@link coolar.Location.verify|verify} messages.
         * @function encodeDelimited
         * @memberof coolar.Location
         * @static
         * @param {coolar.ILocation} message Location message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Location.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Location message from the specified reader or buffer.
         * @function decode
         * @memberof coolar.Location
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {coolar.Location} Location
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Location.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.coolar.Location();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.latitude = reader.double();
                    break;
                case 2:
                    message.longtitude = reader.double();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Location message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof coolar.Location
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {coolar.Location} Location
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Location.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Location message.
         * @function verify
         * @memberof coolar.Location
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Location.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.latitude != null && message.hasOwnProperty("latitude"))
                if (typeof message.latitude !== "number")
                    return "latitude: number expected";
            if (message.longtitude != null && message.hasOwnProperty("longtitude"))
                if (typeof message.longtitude !== "number")
                    return "longtitude: number expected";
            return null;
        };

        /**
         * Creates a Location message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof coolar.Location
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {coolar.Location} Location
         */
        Location.fromObject = function fromObject(object) {
            if (object instanceof $root.coolar.Location)
                return object;
            let message = new $root.coolar.Location();
            if (object.latitude != null)
                message.latitude = Number(object.latitude);
            if (object.longtitude != null)
                message.longtitude = Number(object.longtitude);
            return message;
        };

        /**
         * Creates a plain object from a Location message. Also converts values to other types if specified.
         * @function toObject
         * @memberof coolar.Location
         * @static
         * @param {coolar.Location} message Location
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Location.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.latitude = 0;
                object.longtitude = 0;
            }
            if (message.latitude != null && message.hasOwnProperty("latitude"))
                object.latitude = options.json && !isFinite(message.latitude) ? String(message.latitude) : message.latitude;
            if (message.longtitude != null && message.hasOwnProperty("longtitude"))
                object.longtitude = options.json && !isFinite(message.longtitude) ? String(message.longtitude) : message.longtitude;
            return object;
        };

        /**
         * Converts this Location to JSON.
         * @function toJSON
         * @memberof coolar.Location
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Location.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Location;
    })();

    coolar.Trip = (function() {

        /**
         * Properties of a Trip.
         * @memberof coolar
         * @interface ITrip
         * @property {string|null} [start] Trip start
         * @property {string|null} [end] Trip end
         * @property {number|null} [durationSec] Trip durationSec
         * @property {number|null} [feeCent] Trip feeCent
         * @property {coolar.ILocation|null} [startPos] Trip startPos
         * @property {Array.<coolar.ILocation>|null} [pathLocation] Trip pathLocation
         * @property {coolar.TripStatus|null} [status] Trip status
         */

        /**
         * Constructs a new Trip.
         * @memberof coolar
         * @classdesc Represents a Trip.
         * @implements ITrip
         * @constructor
         * @param {coolar.ITrip=} [properties] Properties to set
         */
        function Trip(properties) {
            this.pathLocation = [];
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Trip start.
         * @member {string} start
         * @memberof coolar.Trip
         * @instance
         */
        Trip.prototype.start = "";

        /**
         * Trip end.
         * @member {string} end
         * @memberof coolar.Trip
         * @instance
         */
        Trip.prototype.end = "";

        /**
         * Trip durationSec.
         * @member {number} durationSec
         * @memberof coolar.Trip
         * @instance
         */
        Trip.prototype.durationSec = 0;

        /**
         * Trip feeCent.
         * @member {number} feeCent
         * @memberof coolar.Trip
         * @instance
         */
        Trip.prototype.feeCent = 0;

        /**
         * Trip startPos.
         * @member {coolar.ILocation|null|undefined} startPos
         * @memberof coolar.Trip
         * @instance
         */
        Trip.prototype.startPos = null;

        /**
         * Trip pathLocation.
         * @member {Array.<coolar.ILocation>} pathLocation
         * @memberof coolar.Trip
         * @instance
         */
        Trip.prototype.pathLocation = $util.emptyArray;

        /**
         * Trip status.
         * @member {coolar.TripStatus} status
         * @memberof coolar.Trip
         * @instance
         */
        Trip.prototype.status = 0;

        /**
         * Creates a new Trip instance using the specified properties.
         * @function create
         * @memberof coolar.Trip
         * @static
         * @param {coolar.ITrip=} [properties] Properties to set
         * @returns {coolar.Trip} Trip instance
         */
        Trip.create = function create(properties) {
            return new Trip(properties);
        };

        /**
         * Encodes the specified Trip message. Does not implicitly {@link coolar.Trip.verify|verify} messages.
         * @function encode
         * @memberof coolar.Trip
         * @static
         * @param {coolar.ITrip} message Trip message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Trip.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.start != null && Object.hasOwnProperty.call(message, "start"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.start);
            if (message.end != null && Object.hasOwnProperty.call(message, "end"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.end);
            if (message.durationSec != null && Object.hasOwnProperty.call(message, "durationSec"))
                writer.uint32(/* id 3, wireType 0 =*/24).int32(message.durationSec);
            if (message.feeCent != null && Object.hasOwnProperty.call(message, "feeCent"))
                writer.uint32(/* id 4, wireType 0 =*/32).int32(message.feeCent);
            if (message.startPos != null && Object.hasOwnProperty.call(message, "startPos"))
                $root.coolar.Location.encode(message.startPos, writer.uint32(/* id 5, wireType 2 =*/42).fork()).ldelim();
            if (message.pathLocation != null && message.pathLocation.length)
                for (let i = 0; i < message.pathLocation.length; ++i)
                    $root.coolar.Location.encode(message.pathLocation[i], writer.uint32(/* id 6, wireType 2 =*/50).fork()).ldelim();
            if (message.status != null && Object.hasOwnProperty.call(message, "status"))
                writer.uint32(/* id 7, wireType 0 =*/56).int32(message.status);
            return writer;
        };

        /**
         * Encodes the specified Trip message, length delimited. Does not implicitly {@link coolar.Trip.verify|verify} messages.
         * @function encodeDelimited
         * @memberof coolar.Trip
         * @static
         * @param {coolar.ITrip} message Trip message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Trip.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Trip message from the specified reader or buffer.
         * @function decode
         * @memberof coolar.Trip
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {coolar.Trip} Trip
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Trip.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.coolar.Trip();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.start = reader.string();
                    break;
                case 2:
                    message.end = reader.string();
                    break;
                case 3:
                    message.durationSec = reader.int32();
                    break;
                case 4:
                    message.feeCent = reader.int32();
                    break;
                case 5:
                    message.startPos = $root.coolar.Location.decode(reader, reader.uint32());
                    break;
                case 6:
                    if (!(message.pathLocation && message.pathLocation.length))
                        message.pathLocation = [];
                    message.pathLocation.push($root.coolar.Location.decode(reader, reader.uint32()));
                    break;
                case 7:
                    message.status = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Trip message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof coolar.Trip
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {coolar.Trip} Trip
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Trip.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Trip message.
         * @function verify
         * @memberof coolar.Trip
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Trip.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.start != null && message.hasOwnProperty("start"))
                if (!$util.isString(message.start))
                    return "start: string expected";
            if (message.end != null && message.hasOwnProperty("end"))
                if (!$util.isString(message.end))
                    return "end: string expected";
            if (message.durationSec != null && message.hasOwnProperty("durationSec"))
                if (!$util.isInteger(message.durationSec))
                    return "durationSec: integer expected";
            if (message.feeCent != null && message.hasOwnProperty("feeCent"))
                if (!$util.isInteger(message.feeCent))
                    return "feeCent: integer expected";
            if (message.startPos != null && message.hasOwnProperty("startPos")) {
                let error = $root.coolar.Location.verify(message.startPos);
                if (error)
                    return "startPos." + error;
            }
            if (message.pathLocation != null && message.hasOwnProperty("pathLocation")) {
                if (!Array.isArray(message.pathLocation))
                    return "pathLocation: array expected";
                for (let i = 0; i < message.pathLocation.length; ++i) {
                    let error = $root.coolar.Location.verify(message.pathLocation[i]);
                    if (error)
                        return "pathLocation." + error;
                }
            }
            if (message.status != null && message.hasOwnProperty("status"))
                switch (message.status) {
                default:
                    return "status: enum value expected";
                case 0:
                case 1:
                case 2:
                case 3:
                case 4:
                    break;
                }
            return null;
        };

        /**
         * Creates a Trip message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof coolar.Trip
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {coolar.Trip} Trip
         */
        Trip.fromObject = function fromObject(object) {
            if (object instanceof $root.coolar.Trip)
                return object;
            let message = new $root.coolar.Trip();
            if (object.start != null)
                message.start = String(object.start);
            if (object.end != null)
                message.end = String(object.end);
            if (object.durationSec != null)
                message.durationSec = object.durationSec | 0;
            if (object.feeCent != null)
                message.feeCent = object.feeCent | 0;
            if (object.startPos != null) {
                if (typeof object.startPos !== "object")
                    throw TypeError(".coolar.Trip.startPos: object expected");
                message.startPos = $root.coolar.Location.fromObject(object.startPos);
            }
            if (object.pathLocation) {
                if (!Array.isArray(object.pathLocation))
                    throw TypeError(".coolar.Trip.pathLocation: array expected");
                message.pathLocation = [];
                for (let i = 0; i < object.pathLocation.length; ++i) {
                    if (typeof object.pathLocation[i] !== "object")
                        throw TypeError(".coolar.Trip.pathLocation: object expected");
                    message.pathLocation[i] = $root.coolar.Location.fromObject(object.pathLocation[i]);
                }
            }
            switch (object.status) {
            case "TS_NOT_SPECIFIDE":
            case 0:
                message.status = 0;
                break;
            case "NOT_STARTED":
            case 1:
                message.status = 1;
                break;
            case "IN_PROGRESS":
            case 2:
                message.status = 2;
                break;
            case "FINISHED":
            case 3:
                message.status = 3;
                break;
            case "PAID":
            case 4:
                message.status = 4;
                break;
            }
            return message;
        };

        /**
         * Creates a plain object from a Trip message. Also converts values to other types if specified.
         * @function toObject
         * @memberof coolar.Trip
         * @static
         * @param {coolar.Trip} message Trip
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Trip.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.arrays || options.defaults)
                object.pathLocation = [];
            if (options.defaults) {
                object.start = "";
                object.end = "";
                object.durationSec = 0;
                object.feeCent = 0;
                object.startPos = null;
                object.status = options.enums === String ? "TS_NOT_SPECIFIDE" : 0;
            }
            if (message.start != null && message.hasOwnProperty("start"))
                object.start = message.start;
            if (message.end != null && message.hasOwnProperty("end"))
                object.end = message.end;
            if (message.durationSec != null && message.hasOwnProperty("durationSec"))
                object.durationSec = message.durationSec;
            if (message.feeCent != null && message.hasOwnProperty("feeCent"))
                object.feeCent = message.feeCent;
            if (message.startPos != null && message.hasOwnProperty("startPos"))
                object.startPos = $root.coolar.Location.toObject(message.startPos, options);
            if (message.pathLocation && message.pathLocation.length) {
                object.pathLocation = [];
                for (let j = 0; j < message.pathLocation.length; ++j)
                    object.pathLocation[j] = $root.coolar.Location.toObject(message.pathLocation[j], options);
            }
            if (message.status != null && message.hasOwnProperty("status"))
                object.status = options.enums === String ? $root.coolar.TripStatus[message.status] : message.status;
            return object;
        };

        /**
         * Converts this Trip to JSON.
         * @function toJSON
         * @memberof coolar.Trip
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Trip.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Trip;
    })();

    /**
     * TripStatus enum.
     * @name coolar.TripStatus
     * @enum {number}
     * @property {number} TS_NOT_SPECIFIDE=0 TS_NOT_SPECIFIDE value
     * @property {number} NOT_STARTED=1 NOT_STARTED value
     * @property {number} IN_PROGRESS=2 IN_PROGRESS value
     * @property {number} FINISHED=3 FINISHED value
     * @property {number} PAID=4 PAID value
     */
    coolar.TripStatus = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "TS_NOT_SPECIFIDE"] = 0;
        values[valuesById[1] = "NOT_STARTED"] = 1;
        values[valuesById[2] = "IN_PROGRESS"] = 2;
        values[valuesById[3] = "FINISHED"] = 3;
        values[valuesById[4] = "PAID"] = 4;
        return values;
    })();

    coolar.GetTripRequest = (function() {

        /**
         * Properties of a GetTripRequest.
         * @memberof coolar
         * @interface IGetTripRequest
         * @property {string|null} [id] GetTripRequest id
         */

        /**
         * Constructs a new GetTripRequest.
         * @memberof coolar
         * @classdesc Represents a GetTripRequest.
         * @implements IGetTripRequest
         * @constructor
         * @param {coolar.IGetTripRequest=} [properties] Properties to set
         */
        function GetTripRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * GetTripRequest id.
         * @member {string} id
         * @memberof coolar.GetTripRequest
         * @instance
         */
        GetTripRequest.prototype.id = "";

        /**
         * Creates a new GetTripRequest instance using the specified properties.
         * @function create
         * @memberof coolar.GetTripRequest
         * @static
         * @param {coolar.IGetTripRequest=} [properties] Properties to set
         * @returns {coolar.GetTripRequest} GetTripRequest instance
         */
        GetTripRequest.create = function create(properties) {
            return new GetTripRequest(properties);
        };

        /**
         * Encodes the specified GetTripRequest message. Does not implicitly {@link coolar.GetTripRequest.verify|verify} messages.
         * @function encode
         * @memberof coolar.GetTripRequest
         * @static
         * @param {coolar.IGetTripRequest} message GetTripRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        GetTripRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.id != null && Object.hasOwnProperty.call(message, "id"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.id);
            return writer;
        };

        /**
         * Encodes the specified GetTripRequest message, length delimited. Does not implicitly {@link coolar.GetTripRequest.verify|verify} messages.
         * @function encodeDelimited
         * @memberof coolar.GetTripRequest
         * @static
         * @param {coolar.IGetTripRequest} message GetTripRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        GetTripRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a GetTripRequest message from the specified reader or buffer.
         * @function decode
         * @memberof coolar.GetTripRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {coolar.GetTripRequest} GetTripRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        GetTripRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.coolar.GetTripRequest();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a GetTripRequest message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof coolar.GetTripRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {coolar.GetTripRequest} GetTripRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        GetTripRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a GetTripRequest message.
         * @function verify
         * @memberof coolar.GetTripRequest
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        GetTripRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.id != null && message.hasOwnProperty("id"))
                if (!$util.isString(message.id))
                    return "id: string expected";
            return null;
        };

        /**
         * Creates a GetTripRequest message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof coolar.GetTripRequest
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {coolar.GetTripRequest} GetTripRequest
         */
        GetTripRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.coolar.GetTripRequest)
                return object;
            let message = new $root.coolar.GetTripRequest();
            if (object.id != null)
                message.id = String(object.id);
            return message;
        };

        /**
         * Creates a plain object from a GetTripRequest message. Also converts values to other types if specified.
         * @function toObject
         * @memberof coolar.GetTripRequest
         * @static
         * @param {coolar.GetTripRequest} message GetTripRequest
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        GetTripRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.id = "";
            if (message.id != null && message.hasOwnProperty("id"))
                object.id = message.id;
            return object;
        };

        /**
         * Converts this GetTripRequest to JSON.
         * @function toJSON
         * @memberof coolar.GetTripRequest
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        GetTripRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return GetTripRequest;
    })();

    coolar.GetTripResponse = (function() {

        /**
         * Properties of a GetTripResponse.
         * @memberof coolar
         * @interface IGetTripResponse
         * @property {string|null} [id] GetTripResponse id
         * @property {coolar.ITrip|null} [trip] GetTripResponse trip
         */

        /**
         * Constructs a new GetTripResponse.
         * @memberof coolar
         * @classdesc Represents a GetTripResponse.
         * @implements IGetTripResponse
         * @constructor
         * @param {coolar.IGetTripResponse=} [properties] Properties to set
         */
        function GetTripResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * GetTripResponse id.
         * @member {string} id
         * @memberof coolar.GetTripResponse
         * @instance
         */
        GetTripResponse.prototype.id = "";

        /**
         * GetTripResponse trip.
         * @member {coolar.ITrip|null|undefined} trip
         * @memberof coolar.GetTripResponse
         * @instance
         */
        GetTripResponse.prototype.trip = null;

        /**
         * Creates a new GetTripResponse instance using the specified properties.
         * @function create
         * @memberof coolar.GetTripResponse
         * @static
         * @param {coolar.IGetTripResponse=} [properties] Properties to set
         * @returns {coolar.GetTripResponse} GetTripResponse instance
         */
        GetTripResponse.create = function create(properties) {
            return new GetTripResponse(properties);
        };

        /**
         * Encodes the specified GetTripResponse message. Does not implicitly {@link coolar.GetTripResponse.verify|verify} messages.
         * @function encode
         * @memberof coolar.GetTripResponse
         * @static
         * @param {coolar.IGetTripResponse} message GetTripResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        GetTripResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.id != null && Object.hasOwnProperty.call(message, "id"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.id);
            if (message.trip != null && Object.hasOwnProperty.call(message, "trip"))
                $root.coolar.Trip.encode(message.trip, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified GetTripResponse message, length delimited. Does not implicitly {@link coolar.GetTripResponse.verify|verify} messages.
         * @function encodeDelimited
         * @memberof coolar.GetTripResponse
         * @static
         * @param {coolar.IGetTripResponse} message GetTripResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        GetTripResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a GetTripResponse message from the specified reader or buffer.
         * @function decode
         * @memberof coolar.GetTripResponse
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {coolar.GetTripResponse} GetTripResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        GetTripResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.coolar.GetTripResponse();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                case 2:
                    message.trip = $root.coolar.Trip.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a GetTripResponse message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof coolar.GetTripResponse
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {coolar.GetTripResponse} GetTripResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        GetTripResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a GetTripResponse message.
         * @function verify
         * @memberof coolar.GetTripResponse
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        GetTripResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.id != null && message.hasOwnProperty("id"))
                if (!$util.isString(message.id))
                    return "id: string expected";
            if (message.trip != null && message.hasOwnProperty("trip")) {
                let error = $root.coolar.Trip.verify(message.trip);
                if (error)
                    return "trip." + error;
            }
            return null;
        };

        /**
         * Creates a GetTripResponse message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof coolar.GetTripResponse
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {coolar.GetTripResponse} GetTripResponse
         */
        GetTripResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.coolar.GetTripResponse)
                return object;
            let message = new $root.coolar.GetTripResponse();
            if (object.id != null)
                message.id = String(object.id);
            if (object.trip != null) {
                if (typeof object.trip !== "object")
                    throw TypeError(".coolar.GetTripResponse.trip: object expected");
                message.trip = $root.coolar.Trip.fromObject(object.trip);
            }
            return message;
        };

        /**
         * Creates a plain object from a GetTripResponse message. Also converts values to other types if specified.
         * @function toObject
         * @memberof coolar.GetTripResponse
         * @static
         * @param {coolar.GetTripResponse} message GetTripResponse
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        GetTripResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.id = "";
                object.trip = null;
            }
            if (message.id != null && message.hasOwnProperty("id"))
                object.id = message.id;
            if (message.trip != null && message.hasOwnProperty("trip"))
                object.trip = $root.coolar.Trip.toObject(message.trip, options);
            return object;
        };

        /**
         * Converts this GetTripResponse to JSON.
         * @function toJSON
         * @memberof coolar.GetTripResponse
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        GetTripResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return GetTripResponse;
    })();

    coolar.TripService = (function() {

        /**
         * Constructs a new TripService service.
         * @memberof coolar
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
         * @memberof coolar.TripService
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
         * Callback as used by {@link coolar.TripService#getTrip}.
         * @memberof coolar.TripService
         * @typedef GetTripCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {coolar.GetTripResponse} [response] GetTripResponse
         */

        /**
         * Calls GetTrip.
         * @function getTrip
         * @memberof coolar.TripService
         * @instance
         * @param {coolar.IGetTripRequest} request GetTripRequest message or plain object
         * @param {coolar.TripService.GetTripCallback} callback Node-style callback called with the error, if any, and GetTripResponse
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(TripService.prototype.getTrip = function getTrip(request, callback) {
            return this.rpcCall(getTrip, $root.coolar.GetTripRequest, $root.coolar.GetTripResponse, request, callback);
        }, "name", { value: "GetTrip" });

        /**
         * Calls GetTrip.
         * @function getTrip
         * @memberof coolar.TripService
         * @instance
         * @param {coolar.IGetTripRequest} request GetTripRequest message or plain object
         * @returns {Promise<coolar.GetTripResponse>} Promise
         * @variation 2
         */

        return TripService;
    })();

    return coolar;
})();