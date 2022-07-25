var PROTO_PATH = __dirname + '/../bsn/grpc_record.proto';

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
    that.table = table;

    client.BabaSays({operation: "op1", param:"param1"}, function(err, response) {
        console.log('Got answer:', JSON.stringify(response));
    });

    return this;
};