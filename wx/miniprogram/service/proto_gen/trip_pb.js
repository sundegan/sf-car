import * as $protobuf from "protobufjs";

// Common aliases
const $util = $protobuf.util;

// Exported root namespace
const $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

export const sfcar = $root.sfcar = (() => {

    /**
     * Namespace sfcar.
     * @exports sfcar
     * @namespace
     */
    const sfcar = {};

    /**
     * TripStatus enum.
     * @name sfcar.TripStatus
     * @enum {number}
     * @property {number} TS_NOT_SPECIFIED=0 TS_NOT_SPECIFIED value
     * @property {number} IN_PROGRESS=1 IN_PROGRESS value
     * @property {number} FINISHED=2 FINISHED value
     */
    sfcar.TripStatus = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "TS_NOT_SPECIFIED"] = 0;
        values[valuesById[1] = "IN_PROGRESS"] = 1;
        values[valuesById[2] = "FINISHED"] = 2;
        return values;
    })();

    sfcar.Location = (function() {

        /**
         * Properties of a Location.
         * @memberof sfcar
         * @interface ILocation
         * @property {number|null} [latitude] Location latitude
         * @property {number|null} [longitude] Location longitude
         */

        /**
         * Constructs a new Location.
         * @memberof sfcar
         * @classdesc Represents a Location.
         * @implements ILocation
         * @constructor
         * @param {sfcar.ILocation=} [properties] Properties to set
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
         * @memberof sfcar.Location
         * @instance
         */
        Location.prototype.latitude = 0;

        /**
         * Location longitude.
         * @member {number} longitude
         * @memberof sfcar.Location
         * @instance
         */
        Location.prototype.longitude = 0;

        /**
         * Creates a Location message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof sfcar.Location
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {sfcar.Location} Location
         */
        Location.fromObject = function fromObject(object) {
            if (object instanceof $root.sfcar.Location)
                return object;
            let message = new $root.sfcar.Location();
            if (object.latitude != null)
                message.latitude = Number(object.latitude);
            if (object.longitude != null)
                message.longitude = Number(object.longitude);
            return message;
        };

        /**
         * Creates a plain object from a Location message. Also converts values to other types if specified.
         * @function toObject
         * @memberof sfcar.Location
         * @static
         * @param {sfcar.Location} message Location
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Location.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.latitude = 0;
                object.longitude = 0;
            }
            if (message.latitude != null && message.hasOwnProperty("latitude"))
                object.latitude = options.json && !isFinite(message.latitude) ? String(message.latitude) : message.latitude;
            if (message.longitude != null && message.hasOwnProperty("longitude"))
                object.longitude = options.json && !isFinite(message.longitude) ? String(message.longitude) : message.longitude;
            return object;
        };

        /**
         * Converts this Location to JSON.
         * @function toJSON
         * @memberof sfcar.Location
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Location.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Location
         * @function getTypeUrl
         * @memberof sfcar.Location
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Location.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/sfcar.Location";
        };

        return Location;
    })();

    sfcar.Trip = (function() {

        /**
         * Properties of a Trip.
         * @memberof sfcar
         * @interface ITrip
         * @property {string|null} [start] Trip start
         * @property {string|null} [end] Trip end
         * @property {number|null} [durationSec] Trip durationSec
         * @property {number|null} [feeCent] Trip feeCent
         * @property {sfcar.ILocation|null} [startPos] Trip startPos
         * @property {sfcar.ILocation|null} [endPos] Trip endPos
         * @property {sfcar.TripStatus|null} [status] Trip status
         */

        /**
         * Constructs a new Trip.
         * @memberof sfcar
         * @classdesc Represents a Trip.
         * @implements ITrip
         * @constructor
         * @param {sfcar.ITrip=} [properties] Properties to set
         */
        function Trip(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Trip start.
         * @member {string} start
         * @memberof sfcar.Trip
         * @instance
         */
        Trip.prototype.start = "";

        /**
         * Trip end.
         * @member {string} end
         * @memberof sfcar.Trip
         * @instance
         */
        Trip.prototype.end = "";

        /**
         * Trip durationSec.
         * @member {number} durationSec
         * @memberof sfcar.Trip
         * @instance
         */
        Trip.prototype.durationSec = 0;

        /**
         * Trip feeCent.
         * @member {number} feeCent
         * @memberof sfcar.Trip
         * @instance
         */
        Trip.prototype.feeCent = 0;

        /**
         * Trip startPos.
         * @member {sfcar.ILocation|null|undefined} startPos
         * @memberof sfcar.Trip
         * @instance
         */
        Trip.prototype.startPos = null;

        /**
         * Trip endPos.
         * @member {sfcar.ILocation|null|undefined} endPos
         * @memberof sfcar.Trip
         * @instance
         */
        Trip.prototype.endPos = null;

        /**
         * Trip status.
         * @member {sfcar.TripStatus} status
         * @memberof sfcar.Trip
         * @instance
         */
        Trip.prototype.status = 0;

        /**
         * Creates a Trip message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof sfcar.Trip
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {sfcar.Trip} Trip
         */
        Trip.fromObject = function fromObject(object) {
            if (object instanceof $root.sfcar.Trip)
                return object;
            let message = new $root.sfcar.Trip();
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
                    throw TypeError(".sfcar.Trip.startPos: object expected");
                message.startPos = $root.sfcar.Location.fromObject(object.startPos);
            }
            if (object.endPos != null) {
                if (typeof object.endPos !== "object")
                    throw TypeError(".sfcar.Trip.endPos: object expected");
                message.endPos = $root.sfcar.Location.fromObject(object.endPos);
            }
            switch (object.status) {
            default:
                if (typeof object.status === "number") {
                    message.status = object.status;
                    break;
                }
                break;
            case "TS_NOT_SPECIFIED":
            case 0:
                message.status = 0;
                break;
            case "IN_PROGRESS":
            case 1:
                message.status = 1;
                break;
            case "FINISHED":
            case 2:
                message.status = 2;
                break;
            }
            return message;
        };

        /**
         * Creates a plain object from a Trip message. Also converts values to other types if specified.
         * @function toObject
         * @memberof sfcar.Trip
         * @static
         * @param {sfcar.Trip} message Trip
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Trip.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.start = "";
                object.end = "";
                object.durationSec = 0;
                object.feeCent = 0;
                object.startPos = null;
                object.endPos = null;
                object.status = options.enums === String ? "TS_NOT_SPECIFIED" : 0;
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
                object.startPos = $root.sfcar.Location.toObject(message.startPos, options);
            if (message.endPos != null && message.hasOwnProperty("endPos"))
                object.endPos = $root.sfcar.Location.toObject(message.endPos, options);
            if (message.status != null && message.hasOwnProperty("status"))
                object.status = options.enums === String ? $root.sfcar.TripStatus[message.status] === undefined ? message.status : $root.sfcar.TripStatus[message.status] : message.status;
            return object;
        };

        /**
         * Converts this Trip to JSON.
         * @function toJSON
         * @memberof sfcar.Trip
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Trip.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Trip
         * @function getTypeUrl
         * @memberof sfcar.Trip
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Trip.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/sfcar.Trip";
        };

        return Trip;
    })();

    sfcar.GetTripRequest = (function() {

        /**
         * Properties of a GetTripRequest.
         * @memberof sfcar
         * @interface IGetTripRequest
         * @property {string|null} [tripId] GetTripRequest tripId
         */

        /**
         * Constructs a new GetTripRequest.
         * @memberof sfcar
         * @classdesc Represents a GetTripRequest.
         * @implements IGetTripRequest
         * @constructor
         * @param {sfcar.IGetTripRequest=} [properties] Properties to set
         */
        function GetTripRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * GetTripRequest tripId.
         * @member {string} tripId
         * @memberof sfcar.GetTripRequest
         * @instance
         */
        GetTripRequest.prototype.tripId = "";

        /**
         * Creates a GetTripRequest message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof sfcar.GetTripRequest
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {sfcar.GetTripRequest} GetTripRequest
         */
        GetTripRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.sfcar.GetTripRequest)
                return object;
            let message = new $root.sfcar.GetTripRequest();
            if (object.tripId != null)
                message.tripId = String(object.tripId);
            return message;
        };

        /**
         * Creates a plain object from a GetTripRequest message. Also converts values to other types if specified.
         * @function toObject
         * @memberof sfcar.GetTripRequest
         * @static
         * @param {sfcar.GetTripRequest} message GetTripRequest
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        GetTripRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.tripId = "";
            if (message.tripId != null && message.hasOwnProperty("tripId"))
                object.tripId = message.tripId;
            return object;
        };

        /**
         * Converts this GetTripRequest to JSON.
         * @function toJSON
         * @memberof sfcar.GetTripRequest
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        GetTripRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for GetTripRequest
         * @function getTypeUrl
         * @memberof sfcar.GetTripRequest
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        GetTripRequest.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/sfcar.GetTripRequest";
        };

        return GetTripRequest;
    })();

    sfcar.GetTripResponse = (function() {

        /**
         * Properties of a GetTripResponse.
         * @memberof sfcar
         * @interface IGetTripResponse
         * @property {string|null} [id] GetTripResponse id
         * @property {sfcar.ITrip|null} [trip] GetTripResponse trip
         */

        /**
         * Constructs a new GetTripResponse.
         * @memberof sfcar
         * @classdesc Represents a GetTripResponse.
         * @implements IGetTripResponse
         * @constructor
         * @param {sfcar.IGetTripResponse=} [properties] Properties to set
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
         * @memberof sfcar.GetTripResponse
         * @instance
         */
        GetTripResponse.prototype.id = "";

        /**
         * GetTripResponse trip.
         * @member {sfcar.ITrip|null|undefined} trip
         * @memberof sfcar.GetTripResponse
         * @instance
         */
        GetTripResponse.prototype.trip = null;

        /**
         * Creates a GetTripResponse message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof sfcar.GetTripResponse
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {sfcar.GetTripResponse} GetTripResponse
         */
        GetTripResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.sfcar.GetTripResponse)
                return object;
            let message = new $root.sfcar.GetTripResponse();
            if (object.id != null)
                message.id = String(object.id);
            if (object.trip != null) {
                if (typeof object.trip !== "object")
                    throw TypeError(".sfcar.GetTripResponse.trip: object expected");
                message.trip = $root.sfcar.Trip.fromObject(object.trip);
            }
            return message;
        };

        /**
         * Creates a plain object from a GetTripResponse message. Also converts values to other types if specified.
         * @function toObject
         * @memberof sfcar.GetTripResponse
         * @static
         * @param {sfcar.GetTripResponse} message GetTripResponse
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
                object.trip = $root.sfcar.Trip.toObject(message.trip, options);
            return object;
        };

        /**
         * Converts this GetTripResponse to JSON.
         * @function toJSON
         * @memberof sfcar.GetTripResponse
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        GetTripResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for GetTripResponse
         * @function getTypeUrl
         * @memberof sfcar.GetTripResponse
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        GetTripResponse.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/sfcar.GetTripResponse";
        };

        return GetTripResponse;
    })();

    sfcar.TripService = (function() {

        /**
         * Constructs a new TripService service.
         * @memberof sfcar
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
         * Callback as used by {@link sfcar.TripService#getTrip}.
         * @memberof sfcar.TripService
         * @typedef GetTripCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {sfcar.GetTripResponse} [response] GetTripResponse
         */

        /**
         * Calls GetTrip.
         * @function getTrip
         * @memberof sfcar.TripService
         * @instance
         * @param {sfcar.IGetTripRequest} request GetTripRequest message or plain object
         * @param {sfcar.TripService.GetTripCallback} callback Node-style callback called with the error, if any, and GetTripResponse
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(TripService.prototype.getTrip = function getTrip(request, callback) {
            return this.rpcCall(getTrip, $root.sfcar.GetTripRequest, $root.sfcar.GetTripResponse, request, callback);
        }, "name", { value: "GetTrip" });

        /**
         * Calls GetTrip.
         * @function getTrip
         * @memberof sfcar.TripService
         * @instance
         * @param {sfcar.IGetTripRequest} request GetTripRequest message or plain object
         * @returns {Promise<sfcar.GetTripResponse>} Promise
         * @variation 2
         */

        return TripService;
    })();

    return sfcar;
})();