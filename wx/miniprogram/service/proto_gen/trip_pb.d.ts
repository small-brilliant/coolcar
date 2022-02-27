import * as $protobuf from "protobufjs";
/** Namespace coolar. */
export namespace coolar {

    /** Properties of a Location. */
    interface ILocation {

        /** Location latitude */
        latitude?: (number|null);

        /** Location longtitude */
        longtitude?: (number|null);
    }

    /** Represents a Location. */
    class Location implements ILocation {

        /**
         * Constructs a new Location.
         * @param [properties] Properties to set
         */
        constructor(properties?: coolar.ILocation);

        /** Location latitude. */
        public latitude: number;

        /** Location longtitude. */
        public longtitude: number;

        /**
         * Creates a new Location instance using the specified properties.
         * @param [properties] Properties to set
         * @returns Location instance
         */
        public static create(properties?: coolar.ILocation): coolar.Location;

        /**
         * Encodes the specified Location message. Does not implicitly {@link coolar.Location.verify|verify} messages.
         * @param message Location message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: coolar.ILocation, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Location message, length delimited. Does not implicitly {@link coolar.Location.verify|verify} messages.
         * @param message Location message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: coolar.ILocation, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Location message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Location
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): coolar.Location;

        /**
         * Decodes a Location message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns Location
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): coolar.Location;

        /**
         * Verifies a Location message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a Location message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Location
         */
        public static fromObject(object: { [k: string]: any }): coolar.Location;

        /**
         * Creates a plain object from a Location message. Also converts values to other types if specified.
         * @param message Location
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: coolar.Location, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Location to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of a Trip. */
    interface ITrip {

        /** Trip start */
        start?: (string|null);

        /** Trip end */
        end?: (string|null);

        /** Trip durationSec */
        durationSec?: (number|null);

        /** Trip feeCent */
        feeCent?: (number|null);

        /** Trip startPos */
        startPos?: (coolar.ILocation|null);

        /** Trip pathLocation */
        pathLocation?: (coolar.ILocation[]|null);

        /** Trip status */
        status?: (coolar.TripStatus|null);
    }

    /** Represents a Trip. */
    class Trip implements ITrip {

        /**
         * Constructs a new Trip.
         * @param [properties] Properties to set
         */
        constructor(properties?: coolar.ITrip);

        /** Trip start. */
        public start: string;

        /** Trip end. */
        public end: string;

        /** Trip durationSec. */
        public durationSec: number;

        /** Trip feeCent. */
        public feeCent: number;

        /** Trip startPos. */
        public startPos?: (coolar.ILocation|null);

        /** Trip pathLocation. */
        public pathLocation: coolar.ILocation[];

        /** Trip status. */
        public status: coolar.TripStatus;

        /**
         * Creates a new Trip instance using the specified properties.
         * @param [properties] Properties to set
         * @returns Trip instance
         */
        public static create(properties?: coolar.ITrip): coolar.Trip;

        /**
         * Encodes the specified Trip message. Does not implicitly {@link coolar.Trip.verify|verify} messages.
         * @param message Trip message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: coolar.ITrip, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Trip message, length delimited. Does not implicitly {@link coolar.Trip.verify|verify} messages.
         * @param message Trip message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: coolar.ITrip, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Trip message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Trip
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): coolar.Trip;

        /**
         * Decodes a Trip message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns Trip
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): coolar.Trip;

        /**
         * Verifies a Trip message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a Trip message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Trip
         */
        public static fromObject(object: { [k: string]: any }): coolar.Trip;

        /**
         * Creates a plain object from a Trip message. Also converts values to other types if specified.
         * @param message Trip
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: coolar.Trip, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Trip to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** TripStatus enum. */
    enum TripStatus {
        TS_NOT_SPECIFIDE = 0,
        NOT_STARTED = 1,
        IN_PROGRESS = 2,
        FINISHED = 3,
        PAID = 4
    }

    /** Properties of a GetTripRequest. */
    interface IGetTripRequest {

        /** GetTripRequest id */
        id?: (string|null);
    }

    /** Represents a GetTripRequest. */
    class GetTripRequest implements IGetTripRequest {

        /**
         * Constructs a new GetTripRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: coolar.IGetTripRequest);

        /** GetTripRequest id. */
        public id: string;

        /**
         * Creates a new GetTripRequest instance using the specified properties.
         * @param [properties] Properties to set
         * @returns GetTripRequest instance
         */
        public static create(properties?: coolar.IGetTripRequest): coolar.GetTripRequest;

        /**
         * Encodes the specified GetTripRequest message. Does not implicitly {@link coolar.GetTripRequest.verify|verify} messages.
         * @param message GetTripRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: coolar.IGetTripRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified GetTripRequest message, length delimited. Does not implicitly {@link coolar.GetTripRequest.verify|verify} messages.
         * @param message GetTripRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: coolar.IGetTripRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a GetTripRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns GetTripRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): coolar.GetTripRequest;

        /**
         * Decodes a GetTripRequest message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns GetTripRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): coolar.GetTripRequest;

        /**
         * Verifies a GetTripRequest message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a GetTripRequest message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns GetTripRequest
         */
        public static fromObject(object: { [k: string]: any }): coolar.GetTripRequest;

        /**
         * Creates a plain object from a GetTripRequest message. Also converts values to other types if specified.
         * @param message GetTripRequest
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: coolar.GetTripRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this GetTripRequest to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of a GetTripResponse. */
    interface IGetTripResponse {

        /** GetTripResponse id */
        id?: (string|null);

        /** GetTripResponse trip */
        trip?: (coolar.ITrip|null);
    }

    /** Represents a GetTripResponse. */
    class GetTripResponse implements IGetTripResponse {

        /**
         * Constructs a new GetTripResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: coolar.IGetTripResponse);

        /** GetTripResponse id. */
        public id: string;

        /** GetTripResponse trip. */
        public trip?: (coolar.ITrip|null);

        /**
         * Creates a new GetTripResponse instance using the specified properties.
         * @param [properties] Properties to set
         * @returns GetTripResponse instance
         */
        public static create(properties?: coolar.IGetTripResponse): coolar.GetTripResponse;

        /**
         * Encodes the specified GetTripResponse message. Does not implicitly {@link coolar.GetTripResponse.verify|verify} messages.
         * @param message GetTripResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: coolar.IGetTripResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified GetTripResponse message, length delimited. Does not implicitly {@link coolar.GetTripResponse.verify|verify} messages.
         * @param message GetTripResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: coolar.IGetTripResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a GetTripResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns GetTripResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): coolar.GetTripResponse;

        /**
         * Decodes a GetTripResponse message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns GetTripResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): coolar.GetTripResponse;

        /**
         * Verifies a GetTripResponse message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a GetTripResponse message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns GetTripResponse
         */
        public static fromObject(object: { [k: string]: any }): coolar.GetTripResponse;

        /**
         * Creates a plain object from a GetTripResponse message. Also converts values to other types if specified.
         * @param message GetTripResponse
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: coolar.GetTripResponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this GetTripResponse to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Represents a TripService */
    class TripService extends $protobuf.rpc.Service {

        /**
         * Constructs a new TripService service.
         * @param rpcImpl RPC implementation
         * @param [requestDelimited=false] Whether requests are length-delimited
         * @param [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Creates new TripService service using the specified rpc implementation.
         * @param rpcImpl RPC implementation
         * @param [requestDelimited=false] Whether requests are length-delimited
         * @param [responseDelimited=false] Whether responses are length-delimited
         * @returns RPC service. Useful where requests and/or responses are streamed.
         */
        public static create(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean): TripService;

        /**
         * Calls GetTrip.
         * @param request GetTripRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and GetTripResponse
         */
        public getTrip(request: coolar.IGetTripRequest, callback: coolar.TripService.GetTripCallback): void;

        /**
         * Calls GetTrip.
         * @param request GetTripRequest message or plain object
         * @returns Promise
         */
        public getTrip(request: coolar.IGetTripRequest): Promise<coolar.GetTripResponse>;
    }

    namespace TripService {

        /**
         * Callback as used by {@link coolar.TripService#getTrip}.
         * @param error Error, if any
         * @param [response] GetTripResponse
         */
        type GetTripCallback = (error: (Error|null), response?: coolar.GetTripResponse) => void;
    }
}
