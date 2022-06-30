var PROTO_PATH = __dirname + '/../bsn/hook.proto';

var grpc = require('@grpc/grpc-js');

var protoLoader = require('@grpc/proto-loader');

var options = { keepCase: true,
                longs: String,
                enums: String,
                defaults: true,
                oneofs: true };

var packageDefinition = protoLoader.loadSync( PROTO_PATH, options );

var hook_proto = grpc.loadPackageDefinition(packageDefinition).hook;

// Implements the RPC method.
function onRecordUpdate(call, callback) {
    try {
        var scriptName = call.request.name;
        var params = call.request.params;
        var path = "../businessrules/" + scriptName + ".js";

        delete require.cache[require.resolve(path)]
        
        var script = require(path)
        var result = script(params)

        callback(null, {result: result});
    } catch ( e ) {
        console.log('****************** error: ' + e )
        callback(null, {result: 'there was an error'});
    }
}

function main() {
  var server = new grpc.Server();

  server.addService(hook_proto.HookService.service, {onRecordUpdate: onRecordUpdate} );

  server.bindAsync('0.0.0.0:50051', grpc.ServerCredentials.createInsecure(), () => {
      console.log('Starting Server...')
      server.start();
  });
}

main();
