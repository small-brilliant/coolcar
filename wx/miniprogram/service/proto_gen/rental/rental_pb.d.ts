import * as $protobuf from "protobufjs";
/** Namespace rental. */
export namespace rental {

    /** Namespace v1. */
    namespace v1 {

        /** Properties of a Location. */
        interface ILocation {

            /** Location latitude */
            latitude?: (number|null);

            /** Location longitude */
            longitude?: (number|null);
        }

        /** Represents a Location. */
        class Location implements ILocation {

            /**
             * Constructs a new Location.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.ILocation);

            /** Location latitude. */
            public latitude: number;

            /** Location longitude. */
            public longitude: number;

            /**
             * Creates a new Location instance using the specified properties.
             * @param [properties] Properties to set
             * @returns Location instance
             */
            public static create(properties?: rental.v1.ILocation): rental.v1.Location;

            /**
             * Encodes the specified Location message. Does not implicitly {@link rental.v1.Location.verify|verify} messages.
             * @param message Location message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.ILocation, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified Location message, length delimited. Does not implicitly {@link rental.v1.Location.verify|verify} messages.
             * @param message Location message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.ILocation, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a Location message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns Location
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.Location;

            /**
             * Decodes a Location message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns Location
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.Location;

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
            public static fromObject(object: { [k: string]: any }): rental.v1.Location;

            /**
             * Creates a plain object from a Location message. Also converts values to other types if specified.
             * @param message Location
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.Location, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this Location to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a LocationStatus. */
        interface ILocationStatus {

            /** LocationStatus location */
            location?: (rental.v1.ILocation|null);

            /** LocationStatus feeCent */
            feeCent?: (number|null);

            /** LocationStatus kmDriven */
            kmDriven?: (number|null);

            /** LocationStatus poiName */
            poiName?: (string|null);

            /** LocationStatus timestampSec */
            timestampSec?: (number|null);
        }

        /** Represents a LocationStatus. */
        class LocationStatus implements ILocationStatus {

            /**
             * Constructs a new LocationStatus.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.ILocationStatus);

            /** LocationStatus location. */
            public location?: (rental.v1.ILocation|null);

            /** LocationStatus feeCent. */
            public feeCent: number;

            /** LocationStatus kmDriven. */
            public kmDriven: number;

            /** LocationStatus poiName. */
            public poiName: string;

            /** LocationStatus timestampSec. */
            public timestampSec: number;

            /**
             * Creates a new LocationStatus instance using the specified properties.
             * @param [properties] Properties to set
             * @returns LocationStatus instance
             */
            public static create(properties?: rental.v1.ILocationStatus): rental.v1.LocationStatus;

            /**
             * Encodes the specified LocationStatus message. Does not implicitly {@link rental.v1.LocationStatus.verify|verify} messages.
             * @param message LocationStatus message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.ILocationStatus, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified LocationStatus message, length delimited. Does not implicitly {@link rental.v1.LocationStatus.verify|verify} messages.
             * @param message LocationStatus message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.ILocationStatus, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a LocationStatus message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns LocationStatus
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.LocationStatus;

            /**
             * Decodes a LocationStatus message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns LocationStatus
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.LocationStatus;

            /**
             * Verifies a LocationStatus message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates a LocationStatus message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns LocationStatus
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.LocationStatus;

            /**
             * Creates a plain object from a LocationStatus message. Also converts values to other types if specified.
             * @param message LocationStatus
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.LocationStatus, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this LocationStatus to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** TripStatus enum. */
        enum TripStatus {
            TS_NOT_SPECIFIED = 0,
            IN_PROGRESS = 1,
            FINISHED = 2
        }

        /** Properties of a TripEntity. */
        interface ITripEntity {

            /** TripEntity id */
            id?: (string|null);

            /** TripEntity trip */
            trip?: (rental.v1.ITrip|null);
        }

        /** Represents a TripEntity. */
        class TripEntity implements ITripEntity {

            /**
             * Constructs a new TripEntity.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.ITripEntity);

            /** TripEntity id. */
            public id: string;

            /** TripEntity trip. */
            public trip?: (rental.v1.ITrip|null);

            /**
             * Creates a new TripEntity instance using the specified properties.
             * @param [properties] Properties to set
             * @returns TripEntity instance
             */
            public static create(properties?: rental.v1.ITripEntity): rental.v1.TripEntity;

            /**
             * Encodes the specified TripEntity message. Does not implicitly {@link rental.v1.TripEntity.verify|verify} messages.
             * @param message TripEntity message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.ITripEntity, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified TripEntity message, length delimited. Does not implicitly {@link rental.v1.TripEntity.verify|verify} messages.
             * @param message TripEntity message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.ITripEntity, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a TripEntity message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns TripEntity
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.TripEntity;

            /**
             * Decodes a TripEntity message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns TripEntity
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.TripEntity;

            /**
             * Verifies a TripEntity message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates a TripEntity message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns TripEntity
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.TripEntity;

            /**
             * Creates a plain object from a TripEntity message. Also converts values to other types if specified.
             * @param message TripEntity
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.TripEntity, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this TripEntity to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a Trip. */
        interface ITrip {

            /** Trip accountId */
            accountId?: (string|null);

            /** Trip carId */
            carId?: (string|null);

            /** Trip start */
            start?: (rental.v1.ILocationStatus|null);

            /** Trip current */
            current?: (rental.v1.ILocationStatus|null);

            /** Trip end */
            end?: (rental.v1.ILocationStatus|null);

            /** Trip status */
            status?: (rental.v1.TripStatus|null);

            /** Trip identityId */
            identityId?: (string|null);
        }

        /** Represents a Trip. */
        class Trip implements ITrip {

            /**
             * Constructs a new Trip.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.ITrip);

            /** Trip accountId. */
            public accountId: string;

            /** Trip carId. */
            public carId: string;

            /** Trip start. */
            public start?: (rental.v1.ILocationStatus|null);

            /** Trip current. */
            public current?: (rental.v1.ILocationStatus|null);

            /** Trip end. */
            public end?: (rental.v1.ILocationStatus|null);

            /** Trip status. */
            public status: rental.v1.TripStatus;

            /** Trip identityId. */
            public identityId: string;

            /**
             * Creates a new Trip instance using the specified properties.
             * @param [properties] Properties to set
             * @returns Trip instance
             */
            public static create(properties?: rental.v1.ITrip): rental.v1.Trip;

            /**
             * Encodes the specified Trip message. Does not implicitly {@link rental.v1.Trip.verify|verify} messages.
             * @param message Trip message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.ITrip, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified Trip message, length delimited. Does not implicitly {@link rental.v1.Trip.verify|verify} messages.
             * @param message Trip message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.ITrip, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a Trip message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns Trip
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.Trip;

            /**
             * Decodes a Trip message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns Trip
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.Trip;

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
            public static fromObject(object: { [k: string]: any }): rental.v1.Trip;

            /**
             * Creates a plain object from a Trip message. Also converts values to other types if specified.
             * @param message Trip
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.Trip, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this Trip to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a CreateTripRequest. */
        interface ICreateTripRequest {

            /** CreateTripRequest start */
            start?: (rental.v1.ILocation|null);

            /** CreateTripRequest carId */
            carId?: (string|null);

            /** CreateTripRequest avatarUrl */
            avatarUrl?: (string|null);
        }

        /** Represents a CreateTripRequest. */
        class CreateTripRequest implements ICreateTripRequest {

            /**
             * Constructs a new CreateTripRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.ICreateTripRequest);

            /** CreateTripRequest start. */
            public start?: (rental.v1.ILocation|null);

            /** CreateTripRequest carId. */
            public carId: string;

            /** CreateTripRequest avatarUrl. */
            public avatarUrl: string;

            /**
             * Creates a new CreateTripRequest instance using the specified properties.
             * @param [properties] Properties to set
             * @returns CreateTripRequest instance
             */
            public static create(properties?: rental.v1.ICreateTripRequest): rental.v1.CreateTripRequest;

            /**
             * Encodes the specified CreateTripRequest message. Does not implicitly {@link rental.v1.CreateTripRequest.verify|verify} messages.
             * @param message CreateTripRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.ICreateTripRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified CreateTripRequest message, length delimited. Does not implicitly {@link rental.v1.CreateTripRequest.verify|verify} messages.
             * @param message CreateTripRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.ICreateTripRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a CreateTripRequest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns CreateTripRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.CreateTripRequest;

            /**
             * Decodes a CreateTripRequest message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns CreateTripRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.CreateTripRequest;

            /**
             * Verifies a CreateTripRequest message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates a CreateTripRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns CreateTripRequest
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.CreateTripRequest;

            /**
             * Creates a plain object from a CreateTripRequest message. Also converts values to other types if specified.
             * @param message CreateTripRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.CreateTripRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this CreateTripRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
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
            constructor(properties?: rental.v1.IGetTripRequest);

            /** GetTripRequest id. */
            public id: string;

            /**
             * Creates a new GetTripRequest instance using the specified properties.
             * @param [properties] Properties to set
             * @returns GetTripRequest instance
             */
            public static create(properties?: rental.v1.IGetTripRequest): rental.v1.GetTripRequest;

            /**
             * Encodes the specified GetTripRequest message. Does not implicitly {@link rental.v1.GetTripRequest.verify|verify} messages.
             * @param message GetTripRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.IGetTripRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified GetTripRequest message, length delimited. Does not implicitly {@link rental.v1.GetTripRequest.verify|verify} messages.
             * @param message GetTripRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.IGetTripRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a GetTripRequest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns GetTripRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.GetTripRequest;

            /**
             * Decodes a GetTripRequest message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns GetTripRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.GetTripRequest;

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
            public static fromObject(object: { [k: string]: any }): rental.v1.GetTripRequest;

            /**
             * Creates a plain object from a GetTripRequest message. Also converts values to other types if specified.
             * @param message GetTripRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.GetTripRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this GetTripRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a GetTripsRequest. */
        interface IGetTripsRequest {

            /** GetTripsRequest status */
            status?: (rental.v1.TripStatus|null);
        }

        /** Represents a GetTripsRequest. */
        class GetTripsRequest implements IGetTripsRequest {

            /**
             * Constructs a new GetTripsRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IGetTripsRequest);

            /** GetTripsRequest status. */
            public status: rental.v1.TripStatus;

            /**
             * Creates a new GetTripsRequest instance using the specified properties.
             * @param [properties] Properties to set
             * @returns GetTripsRequest instance
             */
            public static create(properties?: rental.v1.IGetTripsRequest): rental.v1.GetTripsRequest;

            /**
             * Encodes the specified GetTripsRequest message. Does not implicitly {@link rental.v1.GetTripsRequest.verify|verify} messages.
             * @param message GetTripsRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.IGetTripsRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified GetTripsRequest message, length delimited. Does not implicitly {@link rental.v1.GetTripsRequest.verify|verify} messages.
             * @param message GetTripsRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.IGetTripsRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a GetTripsRequest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns GetTripsRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.GetTripsRequest;

            /**
             * Decodes a GetTripsRequest message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns GetTripsRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.GetTripsRequest;

            /**
             * Verifies a GetTripsRequest message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates a GetTripsRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns GetTripsRequest
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.GetTripsRequest;

            /**
             * Creates a plain object from a GetTripsRequest message. Also converts values to other types if specified.
             * @param message GetTripsRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.GetTripsRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this GetTripsRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a GetTripsResponse. */
        interface IGetTripsResponse {

            /** GetTripsResponse trips */
            trips?: (rental.v1.ITripEntity[]|null);
        }

        /** Represents a GetTripsResponse. */
        class GetTripsResponse implements IGetTripsResponse {

            /**
             * Constructs a new GetTripsResponse.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IGetTripsResponse);

            /** GetTripsResponse trips. */
            public trips: rental.v1.ITripEntity[];

            /**
             * Creates a new GetTripsResponse instance using the specified properties.
             * @param [properties] Properties to set
             * @returns GetTripsResponse instance
             */
            public static create(properties?: rental.v1.IGetTripsResponse): rental.v1.GetTripsResponse;

            /**
             * Encodes the specified GetTripsResponse message. Does not implicitly {@link rental.v1.GetTripsResponse.verify|verify} messages.
             * @param message GetTripsResponse message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.IGetTripsResponse, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified GetTripsResponse message, length delimited. Does not implicitly {@link rental.v1.GetTripsResponse.verify|verify} messages.
             * @param message GetTripsResponse message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.IGetTripsResponse, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a GetTripsResponse message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns GetTripsResponse
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.GetTripsResponse;

            /**
             * Decodes a GetTripsResponse message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns GetTripsResponse
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.GetTripsResponse;

            /**
             * Verifies a GetTripsResponse message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates a GetTripsResponse message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns GetTripsResponse
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.GetTripsResponse;

            /**
             * Creates a plain object from a GetTripsResponse message. Also converts values to other types if specified.
             * @param message GetTripsResponse
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.GetTripsResponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this GetTripsResponse to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of an UpdateTripRequest. */
        interface IUpdateTripRequest {

            /** UpdateTripRequest id */
            id?: (string|null);

            /** UpdateTripRequest current */
            current?: (rental.v1.ILocation|null);

            /** UpdateTripRequest endTrip */
            endTrip?: (boolean|null);
        }

        /** Represents an UpdateTripRequest. */
        class UpdateTripRequest implements IUpdateTripRequest {

            /**
             * Constructs a new UpdateTripRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IUpdateTripRequest);

            /** UpdateTripRequest id. */
            public id: string;

            /** UpdateTripRequest current. */
            public current?: (rental.v1.ILocation|null);

            /** UpdateTripRequest endTrip. */
            public endTrip: boolean;

            /**
             * Creates a new UpdateTripRequest instance using the specified properties.
             * @param [properties] Properties to set
             * @returns UpdateTripRequest instance
             */
            public static create(properties?: rental.v1.IUpdateTripRequest): rental.v1.UpdateTripRequest;

            /**
             * Encodes the specified UpdateTripRequest message. Does not implicitly {@link rental.v1.UpdateTripRequest.verify|verify} messages.
             * @param message UpdateTripRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.IUpdateTripRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified UpdateTripRequest message, length delimited. Does not implicitly {@link rental.v1.UpdateTripRequest.verify|verify} messages.
             * @param message UpdateTripRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.IUpdateTripRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes an UpdateTripRequest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns UpdateTripRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.UpdateTripRequest;

            /**
             * Decodes an UpdateTripRequest message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns UpdateTripRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.UpdateTripRequest;

            /**
             * Verifies an UpdateTripRequest message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates an UpdateTripRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns UpdateTripRequest
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.UpdateTripRequest;

            /**
             * Creates a plain object from an UpdateTripRequest message. Also converts values to other types if specified.
             * @param message UpdateTripRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.UpdateTripRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this UpdateTripRequest to JSON.
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
             * Calls CreateTrip.
             * @param request CreateTripRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and TripEntity
             */
            public createTrip(request: rental.v1.ICreateTripRequest, callback: rental.v1.TripService.CreateTripCallback): void;

            /**
             * Calls CreateTrip.
             * @param request CreateTripRequest message or plain object
             * @returns Promise
             */
            public createTrip(request: rental.v1.ICreateTripRequest): Promise<rental.v1.TripEntity>;

            /**
             * Calls GetTrip.
             * @param request GetTripRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and Trip
             */
            public getTrip(request: rental.v1.IGetTripRequest, callback: rental.v1.TripService.GetTripCallback): void;

            /**
             * Calls GetTrip.
             * @param request GetTripRequest message or plain object
             * @returns Promise
             */
            public getTrip(request: rental.v1.IGetTripRequest): Promise<rental.v1.Trip>;

            /**
             * Calls GetTrips.
             * @param request GetTripsRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and GetTripsResponse
             */
            public getTrips(request: rental.v1.IGetTripsRequest, callback: rental.v1.TripService.GetTripsCallback): void;

            /**
             * Calls GetTrips.
             * @param request GetTripsRequest message or plain object
             * @returns Promise
             */
            public getTrips(request: rental.v1.IGetTripsRequest): Promise<rental.v1.GetTripsResponse>;

            /**
             * Calls UpdateTrip.
             * @param request UpdateTripRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and Trip
             */
            public updateTrip(request: rental.v1.IUpdateTripRequest, callback: rental.v1.TripService.UpdateTripCallback): void;

            /**
             * Calls UpdateTrip.
             * @param request UpdateTripRequest message or plain object
             * @returns Promise
             */
            public updateTrip(request: rental.v1.IUpdateTripRequest): Promise<rental.v1.Trip>;
        }

        namespace TripService {

            /**
             * Callback as used by {@link rental.v1.TripService#createTrip}.
             * @param error Error, if any
             * @param [response] TripEntity
             */
            type CreateTripCallback = (error: (Error|null), response?: rental.v1.TripEntity) => void;

            /**
             * Callback as used by {@link rental.v1.TripService#getTrip}.
             * @param error Error, if any
             * @param [response] Trip
             */
            type GetTripCallback = (error: (Error|null), response?: rental.v1.Trip) => void;

            /**
             * Callback as used by {@link rental.v1.TripService#getTrips}.
             * @param error Error, if any
             * @param [response] GetTripsResponse
             */
            type GetTripsCallback = (error: (Error|null), response?: rental.v1.GetTripsResponse) => void;

            /**
             * Callback as used by {@link rental.v1.TripService#updateTrip}.
             * @param error Error, if any
             * @param [response] Trip
             */
            type UpdateTripCallback = (error: (Error|null), response?: rental.v1.Trip) => void;
        }

        /** Gender enum. */
        enum Gender {
            G_NOT_SPECIFIED = 0,
            FEMALE = 1,
            MALE = 2
        }

        /** IdentityStatus enum. */
        enum IdentityStatus {
            UNSUBMITTED = 0,
            PENDING = 1,
            VERIFIED = 2
        }

        /** Properties of a Profile. */
        interface IProfile {

            /** Profile identity */
            identity?: (rental.v1.IIdentity|null);

            /** Profile identityStatus */
            identityStatus?: (rental.v1.IdentityStatus|null);
        }

        /** Represents a Profile. */
        class Profile implements IProfile {

            /**
             * Constructs a new Profile.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IProfile);

            /** Profile identity. */
            public identity?: (rental.v1.IIdentity|null);

            /** Profile identityStatus. */
            public identityStatus: rental.v1.IdentityStatus;

            /**
             * Creates a new Profile instance using the specified properties.
             * @param [properties] Properties to set
             * @returns Profile instance
             */
            public static create(properties?: rental.v1.IProfile): rental.v1.Profile;

            /**
             * Encodes the specified Profile message. Does not implicitly {@link rental.v1.Profile.verify|verify} messages.
             * @param message Profile message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.IProfile, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified Profile message, length delimited. Does not implicitly {@link rental.v1.Profile.verify|verify} messages.
             * @param message Profile message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.IProfile, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a Profile message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns Profile
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.Profile;

            /**
             * Decodes a Profile message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns Profile
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.Profile;

            /**
             * Verifies a Profile message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates a Profile message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns Profile
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.Profile;

            /**
             * Creates a plain object from a Profile message. Also converts values to other types if specified.
             * @param message Profile
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.Profile, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this Profile to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of an Identity. */
        interface IIdentity {

            /** Identity licNumber */
            licNumber?: (string|null);

            /** Identity name */
            name?: (string|null);

            /** Identity gender */
            gender?: (rental.v1.Gender|null);

            /** Identity birthDateMillis */
            birthDateMillis?: (number|null);
        }

        /** Represents an Identity. */
        class Identity implements IIdentity {

            /**
             * Constructs a new Identity.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IIdentity);

            /** Identity licNumber. */
            public licNumber: string;

            /** Identity name. */
            public name: string;

            /** Identity gender. */
            public gender: rental.v1.Gender;

            /** Identity birthDateMillis. */
            public birthDateMillis: number;

            /**
             * Creates a new Identity instance using the specified properties.
             * @param [properties] Properties to set
             * @returns Identity instance
             */
            public static create(properties?: rental.v1.IIdentity): rental.v1.Identity;

            /**
             * Encodes the specified Identity message. Does not implicitly {@link rental.v1.Identity.verify|verify} messages.
             * @param message Identity message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.IIdentity, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified Identity message, length delimited. Does not implicitly {@link rental.v1.Identity.verify|verify} messages.
             * @param message Identity message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.IIdentity, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes an Identity message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns Identity
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.Identity;

            /**
             * Decodes an Identity message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns Identity
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.Identity;

            /**
             * Verifies an Identity message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates an Identity message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns Identity
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.Identity;

            /**
             * Creates a plain object from an Identity message. Also converts values to other types if specified.
             * @param message Identity
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.Identity, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this Identity to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a GetProfileRequest. */
        interface IGetProfileRequest {
        }

        /** Represents a GetProfileRequest. */
        class GetProfileRequest implements IGetProfileRequest {

            /**
             * Constructs a new GetProfileRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IGetProfileRequest);

            /**
             * Creates a new GetProfileRequest instance using the specified properties.
             * @param [properties] Properties to set
             * @returns GetProfileRequest instance
             */
            public static create(properties?: rental.v1.IGetProfileRequest): rental.v1.GetProfileRequest;

            /**
             * Encodes the specified GetProfileRequest message. Does not implicitly {@link rental.v1.GetProfileRequest.verify|verify} messages.
             * @param message GetProfileRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.IGetProfileRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified GetProfileRequest message, length delimited. Does not implicitly {@link rental.v1.GetProfileRequest.verify|verify} messages.
             * @param message GetProfileRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.IGetProfileRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a GetProfileRequest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns GetProfileRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.GetProfileRequest;

            /**
             * Decodes a GetProfileRequest message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns GetProfileRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.GetProfileRequest;

            /**
             * Verifies a GetProfileRequest message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates a GetProfileRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns GetProfileRequest
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.GetProfileRequest;

            /**
             * Creates a plain object from a GetProfileRequest message. Also converts values to other types if specified.
             * @param message GetProfileRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.GetProfileRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this GetProfileRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a ClearProfileRequest. */
        interface IClearProfileRequest {
        }

        /** Represents a ClearProfileRequest. */
        class ClearProfileRequest implements IClearProfileRequest {

            /**
             * Constructs a new ClearProfileRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IClearProfileRequest);

            /**
             * Creates a new ClearProfileRequest instance using the specified properties.
             * @param [properties] Properties to set
             * @returns ClearProfileRequest instance
             */
            public static create(properties?: rental.v1.IClearProfileRequest): rental.v1.ClearProfileRequest;

            /**
             * Encodes the specified ClearProfileRequest message. Does not implicitly {@link rental.v1.ClearProfileRequest.verify|verify} messages.
             * @param message ClearProfileRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.IClearProfileRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified ClearProfileRequest message, length delimited. Does not implicitly {@link rental.v1.ClearProfileRequest.verify|verify} messages.
             * @param message ClearProfileRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.IClearProfileRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a ClearProfileRequest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns ClearProfileRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.ClearProfileRequest;

            /**
             * Decodes a ClearProfileRequest message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns ClearProfileRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.ClearProfileRequest;

            /**
             * Verifies a ClearProfileRequest message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates a ClearProfileRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns ClearProfileRequest
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.ClearProfileRequest;

            /**
             * Creates a plain object from a ClearProfileRequest message. Also converts values to other types if specified.
             * @param message ClearProfileRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.ClearProfileRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this ClearProfileRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a GetProfilePhotoRequest. */
        interface IGetProfilePhotoRequest {
        }

        /** Represents a GetProfilePhotoRequest. */
        class GetProfilePhotoRequest implements IGetProfilePhotoRequest {

            /**
             * Constructs a new GetProfilePhotoRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IGetProfilePhotoRequest);

            /**
             * Creates a new GetProfilePhotoRequest instance using the specified properties.
             * @param [properties] Properties to set
             * @returns GetProfilePhotoRequest instance
             */
            public static create(properties?: rental.v1.IGetProfilePhotoRequest): rental.v1.GetProfilePhotoRequest;

            /**
             * Encodes the specified GetProfilePhotoRequest message. Does not implicitly {@link rental.v1.GetProfilePhotoRequest.verify|verify} messages.
             * @param message GetProfilePhotoRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.IGetProfilePhotoRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified GetProfilePhotoRequest message, length delimited. Does not implicitly {@link rental.v1.GetProfilePhotoRequest.verify|verify} messages.
             * @param message GetProfilePhotoRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.IGetProfilePhotoRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a GetProfilePhotoRequest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns GetProfilePhotoRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.GetProfilePhotoRequest;

            /**
             * Decodes a GetProfilePhotoRequest message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns GetProfilePhotoRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.GetProfilePhotoRequest;

            /**
             * Verifies a GetProfilePhotoRequest message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates a GetProfilePhotoRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns GetProfilePhotoRequest
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.GetProfilePhotoRequest;

            /**
             * Creates a plain object from a GetProfilePhotoRequest message. Also converts values to other types if specified.
             * @param message GetProfilePhotoRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.GetProfilePhotoRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this GetProfilePhotoRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a GetProfilePhotoReponse. */
        interface IGetProfilePhotoReponse {

            /** GetProfilePhotoReponse url */
            url?: (string|null);
        }

        /** Represents a GetProfilePhotoReponse. */
        class GetProfilePhotoReponse implements IGetProfilePhotoReponse {

            /**
             * Constructs a new GetProfilePhotoReponse.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IGetProfilePhotoReponse);

            /** GetProfilePhotoReponse url. */
            public url: string;

            /**
             * Creates a new GetProfilePhotoReponse instance using the specified properties.
             * @param [properties] Properties to set
             * @returns GetProfilePhotoReponse instance
             */
            public static create(properties?: rental.v1.IGetProfilePhotoReponse): rental.v1.GetProfilePhotoReponse;

            /**
             * Encodes the specified GetProfilePhotoReponse message. Does not implicitly {@link rental.v1.GetProfilePhotoReponse.verify|verify} messages.
             * @param message GetProfilePhotoReponse message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.IGetProfilePhotoReponse, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified GetProfilePhotoReponse message, length delimited. Does not implicitly {@link rental.v1.GetProfilePhotoReponse.verify|verify} messages.
             * @param message GetProfilePhotoReponse message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.IGetProfilePhotoReponse, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a GetProfilePhotoReponse message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns GetProfilePhotoReponse
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.GetProfilePhotoReponse;

            /**
             * Decodes a GetProfilePhotoReponse message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns GetProfilePhotoReponse
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.GetProfilePhotoReponse;

            /**
             * Verifies a GetProfilePhotoReponse message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates a GetProfilePhotoReponse message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns GetProfilePhotoReponse
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.GetProfilePhotoReponse;

            /**
             * Creates a plain object from a GetProfilePhotoReponse message. Also converts values to other types if specified.
             * @param message GetProfilePhotoReponse
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.GetProfilePhotoReponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this GetProfilePhotoReponse to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a CreateProfilePhotoRequest. */
        interface ICreateProfilePhotoRequest {
        }

        /** Represents a CreateProfilePhotoRequest. */
        class CreateProfilePhotoRequest implements ICreateProfilePhotoRequest {

            /**
             * Constructs a new CreateProfilePhotoRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.ICreateProfilePhotoRequest);

            /**
             * Creates a new CreateProfilePhotoRequest instance using the specified properties.
             * @param [properties] Properties to set
             * @returns CreateProfilePhotoRequest instance
             */
            public static create(properties?: rental.v1.ICreateProfilePhotoRequest): rental.v1.CreateProfilePhotoRequest;

            /**
             * Encodes the specified CreateProfilePhotoRequest message. Does not implicitly {@link rental.v1.CreateProfilePhotoRequest.verify|verify} messages.
             * @param message CreateProfilePhotoRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.ICreateProfilePhotoRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified CreateProfilePhotoRequest message, length delimited. Does not implicitly {@link rental.v1.CreateProfilePhotoRequest.verify|verify} messages.
             * @param message CreateProfilePhotoRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.ICreateProfilePhotoRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a CreateProfilePhotoRequest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns CreateProfilePhotoRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.CreateProfilePhotoRequest;

            /**
             * Decodes a CreateProfilePhotoRequest message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns CreateProfilePhotoRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.CreateProfilePhotoRequest;

            /**
             * Verifies a CreateProfilePhotoRequest message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates a CreateProfilePhotoRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns CreateProfilePhotoRequest
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.CreateProfilePhotoRequest;

            /**
             * Creates a plain object from a CreateProfilePhotoRequest message. Also converts values to other types if specified.
             * @param message CreateProfilePhotoRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.CreateProfilePhotoRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this CreateProfilePhotoRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a CreateProfilePhotoReponse. */
        interface ICreateProfilePhotoReponse {

            /** CreateProfilePhotoReponse uploadUrl */
            uploadUrl?: (string|null);
        }

        /** Represents a CreateProfilePhotoReponse. */
        class CreateProfilePhotoReponse implements ICreateProfilePhotoReponse {

            /**
             * Constructs a new CreateProfilePhotoReponse.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.ICreateProfilePhotoReponse);

            /** CreateProfilePhotoReponse uploadUrl. */
            public uploadUrl: string;

            /**
             * Creates a new CreateProfilePhotoReponse instance using the specified properties.
             * @param [properties] Properties to set
             * @returns CreateProfilePhotoReponse instance
             */
            public static create(properties?: rental.v1.ICreateProfilePhotoReponse): rental.v1.CreateProfilePhotoReponse;

            /**
             * Encodes the specified CreateProfilePhotoReponse message. Does not implicitly {@link rental.v1.CreateProfilePhotoReponse.verify|verify} messages.
             * @param message CreateProfilePhotoReponse message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.ICreateProfilePhotoReponse, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified CreateProfilePhotoReponse message, length delimited. Does not implicitly {@link rental.v1.CreateProfilePhotoReponse.verify|verify} messages.
             * @param message CreateProfilePhotoReponse message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.ICreateProfilePhotoReponse, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a CreateProfilePhotoReponse message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns CreateProfilePhotoReponse
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.CreateProfilePhotoReponse;

            /**
             * Decodes a CreateProfilePhotoReponse message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns CreateProfilePhotoReponse
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.CreateProfilePhotoReponse;

            /**
             * Verifies a CreateProfilePhotoReponse message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates a CreateProfilePhotoReponse message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns CreateProfilePhotoReponse
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.CreateProfilePhotoReponse;

            /**
             * Creates a plain object from a CreateProfilePhotoReponse message. Also converts values to other types if specified.
             * @param message CreateProfilePhotoReponse
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.CreateProfilePhotoReponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this CreateProfilePhotoReponse to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a CompleteProfilePhotoRequest. */
        interface ICompleteProfilePhotoRequest {
        }

        /** Represents a CompleteProfilePhotoRequest. */
        class CompleteProfilePhotoRequest implements ICompleteProfilePhotoRequest {

            /**
             * Constructs a new CompleteProfilePhotoRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.ICompleteProfilePhotoRequest);

            /**
             * Creates a new CompleteProfilePhotoRequest instance using the specified properties.
             * @param [properties] Properties to set
             * @returns CompleteProfilePhotoRequest instance
             */
            public static create(properties?: rental.v1.ICompleteProfilePhotoRequest): rental.v1.CompleteProfilePhotoRequest;

            /**
             * Encodes the specified CompleteProfilePhotoRequest message. Does not implicitly {@link rental.v1.CompleteProfilePhotoRequest.verify|verify} messages.
             * @param message CompleteProfilePhotoRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.ICompleteProfilePhotoRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified CompleteProfilePhotoRequest message, length delimited. Does not implicitly {@link rental.v1.CompleteProfilePhotoRequest.verify|verify} messages.
             * @param message CompleteProfilePhotoRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.ICompleteProfilePhotoRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a CompleteProfilePhotoRequest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns CompleteProfilePhotoRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.CompleteProfilePhotoRequest;

            /**
             * Decodes a CompleteProfilePhotoRequest message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns CompleteProfilePhotoRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.CompleteProfilePhotoRequest;

            /**
             * Verifies a CompleteProfilePhotoRequest message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates a CompleteProfilePhotoRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns CompleteProfilePhotoRequest
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.CompleteProfilePhotoRequest;

            /**
             * Creates a plain object from a CompleteProfilePhotoRequest message. Also converts values to other types if specified.
             * @param message CompleteProfilePhotoRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.CompleteProfilePhotoRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this CompleteProfilePhotoRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a ClearProfilePhotoRequest. */
        interface IClearProfilePhotoRequest {
        }

        /** Represents a ClearProfilePhotoRequest. */
        class ClearProfilePhotoRequest implements IClearProfilePhotoRequest {

            /**
             * Constructs a new ClearProfilePhotoRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IClearProfilePhotoRequest);

            /**
             * Creates a new ClearProfilePhotoRequest instance using the specified properties.
             * @param [properties] Properties to set
             * @returns ClearProfilePhotoRequest instance
             */
            public static create(properties?: rental.v1.IClearProfilePhotoRequest): rental.v1.ClearProfilePhotoRequest;

            /**
             * Encodes the specified ClearProfilePhotoRequest message. Does not implicitly {@link rental.v1.ClearProfilePhotoRequest.verify|verify} messages.
             * @param message ClearProfilePhotoRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.IClearProfilePhotoRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified ClearProfilePhotoRequest message, length delimited. Does not implicitly {@link rental.v1.ClearProfilePhotoRequest.verify|verify} messages.
             * @param message ClearProfilePhotoRequest message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.IClearProfilePhotoRequest, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a ClearProfilePhotoRequest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns ClearProfilePhotoRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.ClearProfilePhotoRequest;

            /**
             * Decodes a ClearProfilePhotoRequest message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns ClearProfilePhotoRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.ClearProfilePhotoRequest;

            /**
             * Verifies a ClearProfilePhotoRequest message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates a ClearProfilePhotoRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns ClearProfilePhotoRequest
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.ClearProfilePhotoRequest;

            /**
             * Creates a plain object from a ClearProfilePhotoRequest message. Also converts values to other types if specified.
             * @param message ClearProfilePhotoRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.ClearProfilePhotoRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this ClearProfilePhotoRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a ClearProfilePhotoReponse. */
        interface IClearProfilePhotoReponse {
        }

        /** Represents a ClearProfilePhotoReponse. */
        class ClearProfilePhotoReponse implements IClearProfilePhotoReponse {

            /**
             * Constructs a new ClearProfilePhotoReponse.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IClearProfilePhotoReponse);

            /**
             * Creates a new ClearProfilePhotoReponse instance using the specified properties.
             * @param [properties] Properties to set
             * @returns ClearProfilePhotoReponse instance
             */
            public static create(properties?: rental.v1.IClearProfilePhotoReponse): rental.v1.ClearProfilePhotoReponse;

            /**
             * Encodes the specified ClearProfilePhotoReponse message. Does not implicitly {@link rental.v1.ClearProfilePhotoReponse.verify|verify} messages.
             * @param message ClearProfilePhotoReponse message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: rental.v1.IClearProfilePhotoReponse, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified ClearProfilePhotoReponse message, length delimited. Does not implicitly {@link rental.v1.ClearProfilePhotoReponse.verify|verify} messages.
             * @param message ClearProfilePhotoReponse message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: rental.v1.IClearProfilePhotoReponse, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a ClearProfilePhotoReponse message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns ClearProfilePhotoReponse
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): rental.v1.ClearProfilePhotoReponse;

            /**
             * Decodes a ClearProfilePhotoReponse message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns ClearProfilePhotoReponse
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): rental.v1.ClearProfilePhotoReponse;

            /**
             * Verifies a ClearProfilePhotoReponse message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates a ClearProfilePhotoReponse message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns ClearProfilePhotoReponse
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.ClearProfilePhotoReponse;

            /**
             * Creates a plain object from a ClearProfilePhotoReponse message. Also converts values to other types if specified.
             * @param message ClearProfilePhotoReponse
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.ClearProfilePhotoReponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this ClearProfilePhotoReponse to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Represents a ProfileService */
        class ProfileService extends $protobuf.rpc.Service {

            /**
             * Constructs a new ProfileService service.
             * @param rpcImpl RPC implementation
             * @param [requestDelimited=false] Whether requests are length-delimited
             * @param [responseDelimited=false] Whether responses are length-delimited
             */
            constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

            /**
             * Creates new ProfileService service using the specified rpc implementation.
             * @param rpcImpl RPC implementation
             * @param [requestDelimited=false] Whether requests are length-delimited
             * @param [responseDelimited=false] Whether responses are length-delimited
             * @returns RPC service. Useful where requests and/or responses are streamed.
             */
            public static create(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean): ProfileService;

            /**
             * Calls GetProfile.
             * @param request GetProfileRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and Profile
             */
            public getProfile(request: rental.v1.IGetProfileRequest, callback: rental.v1.ProfileService.GetProfileCallback): void;

            /**
             * Calls GetProfile.
             * @param request GetProfileRequest message or plain object
             * @returns Promise
             */
            public getProfile(request: rental.v1.IGetProfileRequest): Promise<rental.v1.Profile>;

            /**
             * Calls SubmitProfile.
             * @param request Identity message or plain object
             * @param callback Node-style callback called with the error, if any, and Profile
             */
            public submitProfile(request: rental.v1.IIdentity, callback: rental.v1.ProfileService.SubmitProfileCallback): void;

            /**
             * Calls SubmitProfile.
             * @param request Identity message or plain object
             * @returns Promise
             */
            public submitProfile(request: rental.v1.IIdentity): Promise<rental.v1.Profile>;

            /**
             * Calls ClearProfile.
             * @param request ClearProfileRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and Profile
             */
            public clearProfile(request: rental.v1.IClearProfileRequest, callback: rental.v1.ProfileService.ClearProfileCallback): void;

            /**
             * Calls ClearProfile.
             * @param request ClearProfileRequest message or plain object
             * @returns Promise
             */
            public clearProfile(request: rental.v1.IClearProfileRequest): Promise<rental.v1.Profile>;

            /**
             * Calls GetProfilePhoto.
             * @param request GetProfilePhotoRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and GetProfilePhotoReponse
             */
            public getProfilePhoto(request: rental.v1.IGetProfilePhotoRequest, callback: rental.v1.ProfileService.GetProfilePhotoCallback): void;

            /**
             * Calls GetProfilePhoto.
             * @param request GetProfilePhotoRequest message or plain object
             * @returns Promise
             */
            public getProfilePhoto(request: rental.v1.IGetProfilePhotoRequest): Promise<rental.v1.GetProfilePhotoReponse>;

            /**
             * Calls CreateProfilePhoto.
             * @param request CreateProfilePhotoRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and CreateProfilePhotoReponse
             */
            public createProfilePhoto(request: rental.v1.ICreateProfilePhotoRequest, callback: rental.v1.ProfileService.CreateProfilePhotoCallback): void;

            /**
             * Calls CreateProfilePhoto.
             * @param request CreateProfilePhotoRequest message or plain object
             * @returns Promise
             */
            public createProfilePhoto(request: rental.v1.ICreateProfilePhotoRequest): Promise<rental.v1.CreateProfilePhotoReponse>;

            /**
             * Calls CompleteProfilePhoto.
             * @param request CompleteProfilePhotoRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and Identity
             */
            public completeProfilePhoto(request: rental.v1.ICompleteProfilePhotoRequest, callback: rental.v1.ProfileService.CompleteProfilePhotoCallback): void;

            /**
             * Calls CompleteProfilePhoto.
             * @param request CompleteProfilePhotoRequest message or plain object
             * @returns Promise
             */
            public completeProfilePhoto(request: rental.v1.ICompleteProfilePhotoRequest): Promise<rental.v1.Identity>;

            /**
             * Calls ClearProfilePhoto.
             * @param request ClearProfilePhotoRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and ClearProfilePhotoReponse
             */
            public clearProfilePhoto(request: rental.v1.IClearProfilePhotoRequest, callback: rental.v1.ProfileService.ClearProfilePhotoCallback): void;

            /**
             * Calls ClearProfilePhoto.
             * @param request ClearProfilePhotoRequest message or plain object
             * @returns Promise
             */
            public clearProfilePhoto(request: rental.v1.IClearProfilePhotoRequest): Promise<rental.v1.ClearProfilePhotoReponse>;
        }

        namespace ProfileService {

            /**
             * Callback as used by {@link rental.v1.ProfileService#getProfile}.
             * @param error Error, if any
             * @param [response] Profile
             */
            type GetProfileCallback = (error: (Error|null), response?: rental.v1.Profile) => void;

            /**
             * Callback as used by {@link rental.v1.ProfileService#submitProfile}.
             * @param error Error, if any
             * @param [response] Profile
             */
            type SubmitProfileCallback = (error: (Error|null), response?: rental.v1.Profile) => void;

            /**
             * Callback as used by {@link rental.v1.ProfileService#clearProfile}.
             * @param error Error, if any
             * @param [response] Profile
             */
            type ClearProfileCallback = (error: (Error|null), response?: rental.v1.Profile) => void;

            /**
             * Callback as used by {@link rental.v1.ProfileService#getProfilePhoto}.
             * @param error Error, if any
             * @param [response] GetProfilePhotoReponse
             */
            type GetProfilePhotoCallback = (error: (Error|null), response?: rental.v1.GetProfilePhotoReponse) => void;

            /**
             * Callback as used by {@link rental.v1.ProfileService#createProfilePhoto}.
             * @param error Error, if any
             * @param [response] CreateProfilePhotoReponse
             */
            type CreateProfilePhotoCallback = (error: (Error|null), response?: rental.v1.CreateProfilePhotoReponse) => void;

            /**
             * Callback as used by {@link rental.v1.ProfileService#completeProfilePhoto}.
             * @param error Error, if any
             * @param [response] Identity
             */
            type CompleteProfilePhotoCallback = (error: (Error|null), response?: rental.v1.Identity) => void;

            /**
             * Callback as used by {@link rental.v1.ProfileService#clearProfilePhoto}.
             * @param error Error, if any
             * @param [response] ClearProfilePhotoReponse
             */
            type ClearProfilePhotoCallback = (error: (Error|null), response?: rental.v1.ClearProfilePhotoReponse) => void;
        }
    }
}
