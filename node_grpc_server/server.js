var PROTO_PATH = __dirname + '/../bsn/grpc_hook.proto';

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
function onScriptCall(call, callback) {
    try {
        var scriptPath = call.request.scriptPath;
        var params = call.request.params;
        var ctx = call.request.ctx;

        delete require.cache[require.resolve(scriptPath)]
        
        var script = require(scriptPath)
        var scriptResult = script( ctx, params )

        callback(null, scriptResult);
    } catch ( e ) {
        console.log('****************** error: ' + e )
        callback(null, {result: 'there was an error'});
    }
}

function main() {
  var server = new grpc.Server();

  server.addService(hook_proto.HookService.service, {onScriptCall: onScriptCall} );

  server.bindAsync('0.0.0.0:50051', grpc.ServerCredentials.createInsecure(), () => {
      console.log('Starting Server...')
      server.start();
  });
}

main();
