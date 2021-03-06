/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var basic_pb = require('./basic_pb.js');
goog.exportSymbol('proto.pb.ErrorCode', null, global);
goog.exportSymbol('proto.pb.GCError', null, global);
goog.exportSymbol('proto.pb.gcerror', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.GCError = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.GCError, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.pb.GCError.displayName = 'proto.pb.GCError';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.GCError.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.GCError.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.GCError} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GCError.toObject = function(includeInstance, msg) {
  var f, obj = {
    errorcode: jspb.Message.getField(msg, 1)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.GCError}
 */
proto.pb.GCError.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.GCError;
  return proto.pb.GCError.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.GCError} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.GCError}
 */
proto.pb.GCError.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.pb.ErrorCode} */ (reader.readEnum());
      msg.setErrorcode(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.GCError.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.GCError.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.GCError} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GCError.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {!proto.pb.ErrorCode} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeEnum(
      1,
      f
    );
  }
};


/**
 * required ErrorCode errorCode = 1;
 * @return {!proto.pb.ErrorCode}
 */
proto.pb.GCError.prototype.getErrorcode = function() {
  return /** @type {!proto.pb.ErrorCode} */ (jspb.Message.getFieldWithDefault(this, 1, 1001));
};


/** @param {!proto.pb.ErrorCode} value */
proto.pb.GCError.prototype.setErrorcode = function(value) {
  jspb.Message.setField(this, 1, value);
};


proto.pb.GCError.prototype.clearErrorcode = function() {
  jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.pb.GCError.prototype.hasErrorcode = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * @enum {number}
 */
proto.pb.ErrorCode = {
  ROOMNOEXIST: 1001,
  AUTHTIMEOUT: 1002,
  UNKNOWN: 0
};


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `gcerror`.
 * @type {!jspb.ExtensionFieldInfo.<!proto.pb.GCError>}
 */
proto.pb.gcerror = new jspb.ExtensionFieldInfo(
    201,
    {gcerror: 0},
    proto.pb.GCError,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         proto.pb.GCError.toObject),
    0);

basic_pb.Message.extensionsBinary[201] = new jspb.ExtensionFieldBinaryInfo(
    proto.pb.gcerror,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    proto.pb.GCError.serializeBinaryToWriter,
    proto.pb.GCError.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
basic_pb.Message.extensions[201] = proto.pb.gcerror;

goog.object.extend(exports, proto.pb);
