/**
 * @fileoverview gRPC-Web generated client stub for ubivolt
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.ubivolt = require('./keyvalue_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.ubivolt.KeyValueClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.ubivolt.KeyValuePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ubivolt.Operation,
 *   !proto.ubivolt.Result>}
 */
const methodInfo_KeyValue_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.ubivolt.Result,
  /** @param {!proto.ubivolt.Operation} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.ubivolt.Result.deserializeBinary
);


/**
 * @param {!proto.ubivolt.Operation} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.ubivolt.Result)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.ubivolt.Result>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ubivolt.KeyValueClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ubivolt.KeyValue/Get',
      request,
      metadata || {},
      methodInfo_KeyValue_Get,
      callback);
};


/**
 * @param {!proto.ubivolt.Operation} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.ubivolt.Result>}
 *     A native promise that resolves to the response
 */
proto.ubivolt.KeyValuePromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ubivolt.KeyValue/Get',
      request,
      metadata || {},
      methodInfo_KeyValue_Get);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ubivolt.Operation,
 *   !proto.ubivolt.Result>}
 */
const methodInfo_KeyValue_Put = new grpc.web.AbstractClientBase.MethodInfo(
  proto.ubivolt.Result,
  /** @param {!proto.ubivolt.Operation} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.ubivolt.Result.deserializeBinary
);


/**
 * @param {!proto.ubivolt.Operation} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.ubivolt.Result)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.ubivolt.Result>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ubivolt.KeyValueClient.prototype.put =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ubivolt.KeyValue/Put',
      request,
      metadata || {},
      methodInfo_KeyValue_Put,
      callback);
};


/**
 * @param {!proto.ubivolt.Operation} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.ubivolt.Result>}
 *     A native promise that resolves to the response
 */
proto.ubivolt.KeyValuePromiseClient.prototype.put =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ubivolt.KeyValue/Put',
      request,
      metadata || {},
      methodInfo_KeyValue_Put);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ubivolt.Operation,
 *   !proto.ubivolt.Result>}
 */
const methodInfo_KeyValue_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  proto.ubivolt.Result,
  /** @param {!proto.ubivolt.Operation} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.ubivolt.Result.deserializeBinary
);


/**
 * @param {!proto.ubivolt.Operation} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.ubivolt.Result)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.ubivolt.Result>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ubivolt.KeyValueClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ubivolt.KeyValue/Delete',
      request,
      metadata || {},
      methodInfo_KeyValue_Delete,
      callback);
};


/**
 * @param {!proto.ubivolt.Operation} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.ubivolt.Result>}
 *     A native promise that resolves to the response
 */
proto.ubivolt.KeyValuePromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ubivolt.KeyValue/Delete',
      request,
      metadata || {},
      methodInfo_KeyValue_Delete);
};


module.exports = proto.ubivolt;

