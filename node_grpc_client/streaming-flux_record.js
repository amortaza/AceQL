var PROTO_PATH = __dirname + '/../bsn/streaming_grpc_record.proto';

var grpc = require('@grpc/grpc-js');
var deasync = require('deasync');

var protoLoader = require('@grpc/proto-loader');

var options = { keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true };

var packageDefinition = protoLoader.loadSync( PROTO_PATH, options );

var record_proto = grpc.loadPackageDefinition(packageDefinition).grpc_record;

var client = new record_proto.RecordService("localhost:9000", grpc.credentials.createInsecure());

module.exports = function(table) {
    var that = this;
    this.table = table;

    var stream = client.BabaSays();

    this.callInProgress = false;
    this.lastResponse = null;
    this.streamEnded = false;

    this.type = "cool"

    stream.on("data", (response) => {
        console.log("on-data: " + JSON.stringify(response));
        that.lastResponse = response;
        that.callInProgress = false;
    });

    stream.on("error", (e) => {
        that.lastResponse = { fault: e };
        that.callInProgress = false;
        console.log("on-error: " + e) }
    );

    stream.on("end", () => {
        that.callInProgress = false;
        that.streamEnded = true;
        that.lastResponse = null;
        console.log("on-end: ended");  }
    );

    this.callInProgress = true;
    stream.write({operation:"NewRecord()", param: table})
    deasync.loopWhile(function(){return that.callInProgress;});

    this.next = function() {
        if (!this._checkStream("next()")) return;

        this.callInProgress = true;
        stream.write({operation:"Next()", param: ""})
        deasync.loopWhile(function(){return that.callInProgress;});

        return this.lastResponse;
    }

    this.query = function() {
        if (!this._checkStream("query()")) return;

        this.callInProgress = true;
        stream.write({operation:"Query()", param: ""})
        deasync.loopWhile(function(){return that.callInProgress;});

        return this.lastResponse;
    }

    this._checkStream = function(name) {
        if (this.streamEnded) {
            console.log("stream has ended...cannot call " + name );
            return false;
        }

        return true;
    }

    return this;
};
