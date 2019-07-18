/**
 * @fileoverview gRPC-Web generated client stub for ubiquiv
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.ubiquiv = require('./keyvalue_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.ubiquiv.KeyValueClient =
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
proto.ubiquiv.KeyValuePromiseClient =
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
 *   !proto.ubiquiv.Operation,
 *   !proto.ubiquiv.Result>}
 */
const methodInfo_KeyValue_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.ubiquiv.Result,
  /** @param {!proto.ubiquiv.Operation} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.ubiquiv.Result.deserializeBinary
);


/**
 * @param {!proto.ubiquiv.Operation} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.ubiquiv.Result)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.ubiquiv.Result>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ubiquiv.KeyValueClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ubiquiv.KeyValue/Get',
      request,
      metadata || {},
      methodInfo_KeyValue_Get,
      callback);
};


/**
 * @param {!proto.ubiquiv.Operation} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.ubiquiv.Result>}
 *     A native promise that resolves to the response
 */
proto.ubiquiv.KeyValuePromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ubiquiv.KeyValue/Get',
      request,
      metadata || {},
      methodInfo_KeyValue_Get);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ubiquiv.Operation,
 *   !proto.ubiquiv.Result>}
 */
const methodInfo_KeyValue_Put = new grpc.web.AbstractClientBase.MethodInfo(
  proto.ubiquiv.Result,
  /** @param {!proto.ubiquiv.Operation} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.ubiquiv.Result.deserializeBinary
);


/**
 * @param {!proto.ubiquiv.Operation} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.ubiquiv.Result)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.ubiquiv.Result>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ubiquiv.KeyValueClient.prototype.put =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ubiquiv.KeyValue/Put',
      request,
      metadata || {},
      methodInfo_KeyValue_Put,
      callback);
};


/**
 * @param {!proto.ubiquiv.Operation} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.ubiquiv.Result>}
 *     A native promise that resolves to the response
 */
proto.ubiquiv.KeyValuePromiseClient.prototype.put =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ubiquiv.KeyValue/Put',
      request,
      metadata || {},
      methodInfo_KeyValue_Put);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ubiquiv.Operation,
 *   !proto.ubiquiv.Result>}
 */
const methodInfo_KeyValue_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  proto.ubiquiv.Result,
  /** @param {!proto.ubiquiv.Operation} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.ubiquiv.Result.deserializeBinary
);


/**
 * @param {!proto.ubiquiv.Operation} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.ubiquiv.Result)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.ubiquiv.Result>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ubiquiv.KeyValueClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ubiquiv.KeyValue/Delete',
      request,
      metadata || {},
      methodInfo_KeyValue_Delete,
      callback);
};


/**
 * @param {!proto.ubiquiv.Operation} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.ubiquiv.Result>}
 *     A native promise that resolves to the response
 */
proto.ubiquiv.KeyValuePromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ubiquiv.KeyValue/Delete',
      request,
      metadata || {},
      methodInfo_KeyValue_Delete);
};


module.exports = proto.ubiquiv;

