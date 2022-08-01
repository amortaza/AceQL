var PROTO_PATH = __dirname + '/../bsn/grpc_flux_record.proto';

var grpc = require('@grpc/grpc-js');

var protoLoader = require('@grpc/proto-loader');

var options = { keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true };

var packageDefinition = protoLoader.loadSync( PROTO_PATH, options );

var record_proto = grpc.loadPackageDefinition(packageDefinition).grpc_flux_record;

var client = new record_proto.FluxRecordService("localhost:9000", grpc.credentials.createInsecure());

module.exports = function(table) {
    var that = this;
    this.table = table;

    var stream = client.GetServiceStream();

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

    this.close = function() {
        return new Promise( (resolve, reject ) => {
            that.resolvers.unshift(resolve);
            stream.write({operation: "Close()"})
        } );
    }

    this.insert = function() {
        return new Promise( (resolve, reject ) => {
            that.resolvers.unshift(resolve);
            stream.write({operation: "Insert()"})
        } );
    }

    this.update = function() {
        return new Promise( (resolve, reject ) => {
            that.resolvers.unshift(resolve);
            stream.write({operation: "Update()"})
        } );
    }

    this.set = function(field, value) {
        return new Promise( (resolve, reject ) => {
            that.resolvers.unshift(resolve);
            stream.write({operation: "Set()", param1: field, param2: value})
        } );
    }

    this.get = function(field) {
        return new Promise( (resolve, reject ) => {
            that.resolvers.unshift(resolve);
            stream.write({operation: "Get()", param1: field})
        } );
    }

    this.open = function() {
        return new Promise( (resolve, reject ) => {
            that.resolvers.unshift(resolve);
            stream.write({operation: "Open()", param1: that.table})
        } );
    }

    this.next = function() {
        return new Promise( (resolve, reject ) => {
            that.resolvers.unshift(resolve);
            stream.write({operation:"Next()"})
        } )
    }

    this.query = function() {
        return new Promise( (resolve, reject ) => {
            that.resolvers.unshift(resolve);
            stream.write({operation:"Query()"})
        } )
    }

    return this;
};
