var PROTO_PATH = __dirname + '/../bsn/streaming_grpc_record.proto';

var grpc = require('@grpc/grpc-js');

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

    var stream = client.GetFluxRecordStream();

    this.resolvers = [];

    stream.on("data", (response) => {
        this.resolvers.pop()(response);
        console.log("on-data: " + JSON.stringify(response));
    });

    stream.on("error", (e) => {
        console.log("on-error: " + e) }
    );

    stream.on("end", () => {
        console.log("on-end: ended");  }
    );

    this.newRecord = function() {
        return new Promise( (resolve, reject ) => {
            that.resolvers.unshift(resolve);
            stream.write({operation: "NewRecord()", param: that.table})
        } );
    }

    this.next = function() {
        return new Promise( (resolve, reject ) => {
            that.resolvers.unshift(resolve);
            stream.write({operation:"Next()", param: ""})
        } )
    }

    this.query = function() {
        return new Promise( (resolve, reject ) => {
            that.resolvers.unshift(resolve);
            stream.write({operation:"Query()", param: ""})
        } )
    }

    return this;
};
