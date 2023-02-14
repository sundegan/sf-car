import * as $protobuf from "protobufjs";

// Common aliases
const $util = $protobuf.util;

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

        v1.CreateBlobRequest = (function() {

            /**
             * Properties of a CreateBlobRequest.
             * @memberof rental.v1
             * @interface ICreateBlobRequest
             * @property {string|null} [accountId] CreateBlobRequest accountId
             * @property {number|null} [uploadUrlTimeoutSec] CreateBlobRequest uploadUrlTimeoutSec
             */

            /**
             * Constructs a new CreateBlobRequest.
             * @memberof rental.v1
             * @classdesc Represents a CreateBlobRequest.
             * @implements ICreateBlobRequest
             * @constructor
             * @param {rental.v1.ICreateBlobRequest=} [properties] Properties to set
             */
            function CreateBlobRequest(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * CreateBlobRequest accountId.
             * @member {string} accountId
             * @memberof rental.v1.CreateBlobRequest
             * @instance
             */
            CreateBlobRequest.prototype.accountId = "";

            /**
             * CreateBlobRequest uploadUrlTimeoutSec.
             * @member {number} uploadUrlTimeoutSec
             * @memberof rental.v1.CreateBlobRequest
             * @instance
             */
            CreateBlobRequest.prototype.uploadUrlTimeoutSec = 0;

            /**
             * Creates a CreateBlobRequest message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.CreateBlobRequest
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.CreateBlobRequest} CreateBlobRequest
             */
            CreateBlobRequest.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.CreateBlobRequest)
                    return object;
                let message = new $root.rental.v1.CreateBlobRequest();
                if (object.accountId != null)
                    message.accountId = String(object.accountId);
                if (object.uploadUrlTimeoutSec != null)
                    message.uploadUrlTimeoutSec = object.uploadUrlTimeoutSec | 0;
                return message;
            };

            /**
             * Creates a plain object from a CreateBlobRequest message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.CreateBlobRequest
             * @static
             * @param {rental.v1.CreateBlobRequest} message CreateBlobRequest
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            CreateBlobRequest.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults) {
                    object.accountId = "";
                    object.uploadUrlTimeoutSec = 0;
                }
                if (message.accountId != null && message.hasOwnProperty("accountId"))
                    object.accountId = message.accountId;
                if (message.uploadUrlTimeoutSec != null && message.hasOwnProperty("uploadUrlTimeoutSec"))
                    object.uploadUrlTimeoutSec = message.uploadUrlTimeoutSec;
                return object;
            };

            /**
             * Converts this CreateBlobRequest to JSON.
             * @function toJSON
             * @memberof rental.v1.CreateBlobRequest
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            CreateBlobRequest.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            /**
             * Gets the default type url for CreateBlobRequest
             * @function getTypeUrl
             * @memberof rental.v1.CreateBlobRequest
             * @static
             * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns {string} The default type url
             */
            CreateBlobRequest.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
                if (typeUrlPrefix === undefined) {
                    typeUrlPrefix = "type.googleapis.com";
                }
                return typeUrlPrefix + "/rental.v1.CreateBlobRequest";
            };

            return CreateBlobRequest;
        })();

        v1.CreateBlobResponse = (function() {

            /**
             * Properties of a CreateBlobResponse.
             * @memberof rental.v1
             * @interface ICreateBlobResponse
             * @property {string|null} [id] CreateBlobResponse id
             * @property {string|null} [uploadUrl] CreateBlobResponse uploadUrl
             */

            /**
             * Constructs a new CreateBlobResponse.
             * @memberof rental.v1
             * @classdesc Represents a CreateBlobResponse.
             * @implements ICreateBlobResponse
             * @constructor
             * @param {rental.v1.ICreateBlobResponse=} [properties] Properties to set
             */
            function CreateBlobResponse(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * CreateBlobResponse id.
             * @member {string} id
             * @memberof rental.v1.CreateBlobResponse
             * @instance
             */
            CreateBlobResponse.prototype.id = "";

            /**
             * CreateBlobResponse uploadUrl.
             * @member {string} uploadUrl
             * @memberof rental.v1.CreateBlobResponse
             * @instance
             */
            CreateBlobResponse.prototype.uploadUrl = "";

            /**
             * Creates a CreateBlobResponse message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.CreateBlobResponse
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.CreateBlobResponse} CreateBlobResponse
             */
            CreateBlobResponse.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.CreateBlobResponse)
                    return object;
                let message = new $root.rental.v1.CreateBlobResponse();
                if (object.id != null)
                    message.id = String(object.id);
                if (object.uploadUrl != null)
                    message.uploadUrl = String(object.uploadUrl);
                return message;
            };

            /**
             * Creates a plain object from a CreateBlobResponse message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.CreateBlobResponse
             * @static
             * @param {rental.v1.CreateBlobResponse} message CreateBlobResponse
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            CreateBlobResponse.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults) {
                    object.id = "";
                    object.uploadUrl = "";
                }
                if (message.id != null && message.hasOwnProperty("id"))
                    object.id = message.id;
                if (message.uploadUrl != null && message.hasOwnProperty("uploadUrl"))
                    object.uploadUrl = message.uploadUrl;
                return object;
            };

            /**
             * Converts this CreateBlobResponse to JSON.
             * @function toJSON
             * @memberof rental.v1.CreateBlobResponse
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            CreateBlobResponse.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            /**
             * Gets the default type url for CreateBlobResponse
             * @function getTypeUrl
             * @memberof rental.v1.CreateBlobResponse
             * @static
             * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns {string} The default type url
             */
            CreateBlobResponse.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
                if (typeUrlPrefix === undefined) {
                    typeUrlPrefix = "type.googleapis.com";
                }
                return typeUrlPrefix + "/rental.v1.CreateBlobResponse";
            };

            return CreateBlobResponse;
        })();

        v1.GetBlobRequest = (function() {

            /**
             * Properties of a GetBlobRequest.
             * @memberof rental.v1
             * @interface IGetBlobRequest
             * @property {string|null} [id] GetBlobRequest id
             */

            /**
             * Constructs a new GetBlobRequest.
             * @memberof rental.v1
             * @classdesc Represents a GetBlobRequest.
             * @implements IGetBlobRequest
             * @constructor
             * @param {rental.v1.IGetBlobRequest=} [properties] Properties to set
             */
            function GetBlobRequest(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * GetBlobRequest id.
             * @member {string} id
             * @memberof rental.v1.GetBlobRequest
             * @instance
             */
            GetBlobRequest.prototype.id = "";

            /**
             * Creates a GetBlobRequest message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.GetBlobRequest
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.GetBlobRequest} GetBlobRequest
             */
            GetBlobRequest.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.GetBlobRequest)
                    return object;
                let message = new $root.rental.v1.GetBlobRequest();
                if (object.id != null)
                    message.id = String(object.id);
                return message;
            };

            /**
             * Creates a plain object from a GetBlobRequest message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.GetBlobRequest
             * @static
             * @param {rental.v1.GetBlobRequest} message GetBlobRequest
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            GetBlobRequest.toObject = function toObject(message, options) {
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
             * Converts this GetBlobRequest to JSON.
             * @function toJSON
             * @memberof rental.v1.GetBlobRequest
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            GetBlobRequest.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            /**
             * Gets the default type url for GetBlobRequest
             * @function getTypeUrl
             * @memberof rental.v1.GetBlobRequest
             * @static
             * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns {string} The default type url
             */
            GetBlobRequest.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
                if (typeUrlPrefix === undefined) {
                    typeUrlPrefix = "type.googleapis.com";
                }
                return typeUrlPrefix + "/rental.v1.GetBlobRequest";
            };

            return GetBlobRequest;
        })();

        v1.GetBlobResponse = (function() {

            /**
             * Properties of a GetBlobResponse.
             * @memberof rental.v1
             * @interface IGetBlobResponse
             * @property {Uint8Array|null} [data] GetBlobResponse data
             */

            /**
             * Constructs a new GetBlobResponse.
             * @memberof rental.v1
             * @classdesc Represents a GetBlobResponse.
             * @implements IGetBlobResponse
             * @constructor
             * @param {rental.v1.IGetBlobResponse=} [properties] Properties to set
             */
            function GetBlobResponse(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * GetBlobResponse data.
             * @member {Uint8Array} data
             * @memberof rental.v1.GetBlobResponse
             * @instance
             */
            GetBlobResponse.prototype.data = $util.newBuffer([]);

            /**
             * Creates a GetBlobResponse message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.GetBlobResponse
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.GetBlobResponse} GetBlobResponse
             */
            GetBlobResponse.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.GetBlobResponse)
                    return object;
                let message = new $root.rental.v1.GetBlobResponse();
                if (object.data != null)
                    if (typeof object.data === "string")
                        $util.base64.decode(object.data, message.data = $util.newBuffer($util.base64.length(object.data)), 0);
                    else if (object.data.length >= 0)
                        message.data = object.data;
                return message;
            };

            /**
             * Creates a plain object from a GetBlobResponse message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.GetBlobResponse
             * @static
             * @param {rental.v1.GetBlobResponse} message GetBlobResponse
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            GetBlobResponse.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults)
                    if (options.bytes === String)
                        object.data = "";
                    else {
                        object.data = [];
                        if (options.bytes !== Array)
                            object.data = $util.newBuffer(object.data);
                    }
                if (message.data != null && message.hasOwnProperty("data"))
                    object.data = options.bytes === String ? $util.base64.encode(message.data, 0, message.data.length) : options.bytes === Array ? Array.prototype.slice.call(message.data) : message.data;
                return object;
            };

            /**
             * Converts this GetBlobResponse to JSON.
             * @function toJSON
             * @memberof rental.v1.GetBlobResponse
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            GetBlobResponse.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            /**
             * Gets the default type url for GetBlobResponse
             * @function getTypeUrl
             * @memberof rental.v1.GetBlobResponse
             * @static
             * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns {string} The default type url
             */
            GetBlobResponse.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
                if (typeUrlPrefix === undefined) {
                    typeUrlPrefix = "type.googleapis.com";
                }
                return typeUrlPrefix + "/rental.v1.GetBlobResponse";
            };

            return GetBlobResponse;
        })();

        v1.GetBlobURLRequest = (function() {

            /**
             * Properties of a GetBlobURLRequest.
             * @memberof rental.v1
             * @interface IGetBlobURLRequest
             * @property {string|null} [id] GetBlobURLRequest id
             * @property {number|null} [timeoutSec] GetBlobURLRequest timeoutSec
             */

            /**
             * Constructs a new GetBlobURLRequest.
             * @memberof rental.v1
             * @classdesc Represents a GetBlobURLRequest.
             * @implements IGetBlobURLRequest
             * @constructor
             * @param {rental.v1.IGetBlobURLRequest=} [properties] Properties to set
             */
            function GetBlobURLRequest(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * GetBlobURLRequest id.
             * @member {string} id
             * @memberof rental.v1.GetBlobURLRequest
             * @instance
             */
            GetBlobURLRequest.prototype.id = "";

            /**
             * GetBlobURLRequest timeoutSec.
             * @member {number} timeoutSec
             * @memberof rental.v1.GetBlobURLRequest
             * @instance
             */
            GetBlobURLRequest.prototype.timeoutSec = 0;

            /**
             * Creates a GetBlobURLRequest message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.GetBlobURLRequest
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.GetBlobURLRequest} GetBlobURLRequest
             */
            GetBlobURLRequest.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.GetBlobURLRequest)
                    return object;
                let message = new $root.rental.v1.GetBlobURLRequest();
                if (object.id != null)
                    message.id = String(object.id);
                if (object.timeoutSec != null)
                    message.timeoutSec = object.timeoutSec | 0;
                return message;
            };

            /**
             * Creates a plain object from a GetBlobURLRequest message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.GetBlobURLRequest
             * @static
             * @param {rental.v1.GetBlobURLRequest} message GetBlobURLRequest
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            GetBlobURLRequest.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults) {
                    object.id = "";
                    object.timeoutSec = 0;
                }
                if (message.id != null && message.hasOwnProperty("id"))
                    object.id = message.id;
                if (message.timeoutSec != null && message.hasOwnProperty("timeoutSec"))
                    object.timeoutSec = message.timeoutSec;
                return object;
            };

            /**
             * Converts this GetBlobURLRequest to JSON.
             * @function toJSON
             * @memberof rental.v1.GetBlobURLRequest
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            GetBlobURLRequest.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            /**
             * Gets the default type url for GetBlobURLRequest
             * @function getTypeUrl
             * @memberof rental.v1.GetBlobURLRequest
             * @static
             * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns {string} The default type url
             */
            GetBlobURLRequest.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
                if (typeUrlPrefix === undefined) {
                    typeUrlPrefix = "type.googleapis.com";
                }
                return typeUrlPrefix + "/rental.v1.GetBlobURLRequest";
            };

            return GetBlobURLRequest;
        })();

        v1.GetBlobURLResponse = (function() {

            /**
             * Properties of a GetBlobURLResponse.
             * @memberof rental.v1
             * @interface IGetBlobURLResponse
             * @property {string|null} [url] GetBlobURLResponse url
             */

            /**
             * Constructs a new GetBlobURLResponse.
             * @memberof rental.v1
             * @classdesc Represents a GetBlobURLResponse.
             * @implements IGetBlobURLResponse
             * @constructor
             * @param {rental.v1.IGetBlobURLResponse=} [properties] Properties to set
             */
            function GetBlobURLResponse(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * GetBlobURLResponse url.
             * @member {string} url
             * @memberof rental.v1.GetBlobURLResponse
             * @instance
             */
            GetBlobURLResponse.prototype.url = "";

            /**
             * Creates a GetBlobURLResponse message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof rental.v1.GetBlobURLResponse
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {rental.v1.GetBlobURLResponse} GetBlobURLResponse
             */
            GetBlobURLResponse.fromObject = function fromObject(object) {
                if (object instanceof $root.rental.v1.GetBlobURLResponse)
                    return object;
                let message = new $root.rental.v1.GetBlobURLResponse();
                if (object.url != null)
                    message.url = String(object.url);
                return message;
            };

            /**
             * Creates a plain object from a GetBlobURLResponse message. Also converts values to other types if specified.
             * @function toObject
             * @memberof rental.v1.GetBlobURLResponse
             * @static
             * @param {rental.v1.GetBlobURLResponse} message GetBlobURLResponse
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            GetBlobURLResponse.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults)
                    object.url = "";
                if (message.url != null && message.hasOwnProperty("url"))
                    object.url = message.url;
                return object;
            };

            /**
             * Converts this GetBlobURLResponse to JSON.
             * @function toJSON
             * @memberof rental.v1.GetBlobURLResponse
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            GetBlobURLResponse.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            /**
             * Gets the default type url for GetBlobURLResponse
             * @function getTypeUrl
             * @memberof rental.v1.GetBlobURLResponse
             * @static
             * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns {string} The default type url
             */
            GetBlobURLResponse.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
                if (typeUrlPrefix === undefined) {
                    typeUrlPrefix = "type.googleapis.com";
                }
                return typeUrlPrefix + "/rental.v1.GetBlobURLResponse";
            };

            return GetBlobURLResponse;
        })();

        v1.BlobService = (function() {

            /**
             * Constructs a new BlobService service.
             * @memberof rental.v1
             * @classdesc Represents a BlobService
             * @extends $protobuf.rpc.Service
             * @constructor
             * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
             * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
             * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
             */
            function BlobService(rpcImpl, requestDelimited, responseDelimited) {
                $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
            }

            (BlobService.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = BlobService;

            /**
             * Callback as used by {@link rental.v1.BlobService#createBlob}.
             * @memberof rental.v1.BlobService
             * @typedef CreateBlobCallback
             * @type {function}
             * @param {Error|null} error Error, if any
             * @param {rental.v1.CreateBlobResponse} [response] CreateBlobResponse
             */

            /**
             * Calls CreateBlob.
             * @function createBlob
             * @memberof rental.v1.BlobService
             * @instance
             * @param {rental.v1.ICreateBlobRequest} request CreateBlobRequest message or plain object
             * @param {rental.v1.BlobService.CreateBlobCallback} callback Node-style callback called with the error, if any, and CreateBlobResponse
             * @returns {undefined}
             * @variation 1
             */
            Object.defineProperty(BlobService.prototype.createBlob = function createBlob(request, callback) {
                return this.rpcCall(createBlob, $root.rental.v1.CreateBlobRequest, $root.rental.v1.CreateBlobResponse, request, callback);
            }, "name", { value: "CreateBlob" });

            /**
             * Calls CreateBlob.
             * @function createBlob
             * @memberof rental.v1.BlobService
             * @instance
             * @param {rental.v1.ICreateBlobRequest} request CreateBlobRequest message or plain object
             * @returns {Promise<rental.v1.CreateBlobResponse>} Promise
             * @variation 2
             */

            /**
             * Callback as used by {@link rental.v1.BlobService#getBlob}.
             * @memberof rental.v1.BlobService
             * @typedef GetBlobCallback
             * @type {function}
             * @param {Error|null} error Error, if any
             * @param {rental.v1.GetBlobResponse} [response] GetBlobResponse
             */

            /**
             * Calls GetBlob.
             * @function getBlob
             * @memberof rental.v1.BlobService
             * @instance
             * @param {rental.v1.IGetBlobRequest} request GetBlobRequest message or plain object
             * @param {rental.v1.BlobService.GetBlobCallback} callback Node-style callback called with the error, if any, and GetBlobResponse
             * @returns {undefined}
             * @variation 1
             */
            Object.defineProperty(BlobService.prototype.getBlob = function getBlob(request, callback) {
                return this.rpcCall(getBlob, $root.rental.v1.GetBlobRequest, $root.rental.v1.GetBlobResponse, request, callback);
            }, "name", { value: "GetBlob" });

            /**
             * Calls GetBlob.
             * @function getBlob
             * @memberof rental.v1.BlobService
             * @instance
             * @param {rental.v1.IGetBlobRequest} request GetBlobRequest message or plain object
             * @returns {Promise<rental.v1.GetBlobResponse>} Promise
             * @variation 2
             */

            /**
             * Callback as used by {@link rental.v1.BlobService#getBlobURL}.
             * @memberof rental.v1.BlobService
             * @typedef GetBlobURLCallback
             * @type {function}
             * @param {Error|null} error Error, if any
             * @param {rental.v1.GetBlobURLResponse} [response] GetBlobURLResponse
             */

            /**
             * Calls GetBlobURL.
             * @function getBlobURL
             * @memberof rental.v1.BlobService
             * @instance
             * @param {rental.v1.IGetBlobURLRequest} request GetBlobURLRequest message or plain object
             * @param {rental.v1.BlobService.GetBlobURLCallback} callback Node-style callback called with the error, if any, and GetBlobURLResponse
             * @returns {undefined}
             * @variation 1
             */
            Object.defineProperty(BlobService.prototype.getBlobURL = function getBlobURL(request, callback) {
                return this.rpcCall(getBlobURL, $root.rental.v1.GetBlobURLRequest, $root.rental.v1.GetBlobURLResponse, request, callback);
            }, "name", { value: "GetBlobURL" });

            /**
             * Calls GetBlobURL.
             * @function getBlobURL
             * @memberof rental.v1.BlobService
             * @instance
             * @param {rental.v1.IGetBlobURLRequest} request GetBlobURLRequest message or plain object
             * @returns {Promise<rental.v1.GetBlobURLResponse>} Promise
             * @variation 2
             */

            return BlobService;
        })();

        return v1;
    })();

    return rental;
})();