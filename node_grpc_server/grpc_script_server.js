var PROTO_PATH = __dirname + '/../bsn/grpc_script.proto';

var grpc = require('@grpc/grpc-js');

var protoLoader = require('@grpc/proto-loader');

var options = { keepCase: true,
                longs: String,
                enums: String,
                defaults: true,
                oneofs: true };

var packageDefinition = protoLoader.loadSync( PROTO_PATH, options );

var grpc_script = grpc.loadPackageDefinition(packageDefinition).grpc_script;

// Implements the RPC method.
async function onBusinessRule(call, callback) {
    try {
        var scriptPath = call.request.scriptPath;
        var action = call.request.action;
        var table = call.request.table;
        var record_id = call.request.record_id;
        var originals = call.request.originals;
        var current = call.request.current;

        // console.log("onBussinessRule: current is " + JSON.stringify(current));

        delete require.cache[require.resolve(scriptPath)]
        
        var script = require(scriptPath)
        var businessRuleResponse = await script( action, table, record_id, originals, current );

        console.log("onBusinessRule: response " + JSON.stringify(businessRuleResponse));

        callback(null, businessRuleResponse);
    } catch ( e ) {
        console.log('****************** error: ' + e )
        callback(null, {error_to_display: e});
    }
}

function onScriptCall(call, callback) {
    try {
        var scriptPath = call.request.scriptPath;
        var params = call.request.params;

        delete require.cache[require.resolve(scriptPath)]

        var script = require(scriptPath)
        var scriptResponse = script( params )

        console.log("onScriptCall: response " + JSON.stringify(scriptResponse));

        callback(null, {answer: scriptResponse});
    } catch ( e ) {
        console.log('****************** error: ' + e )
        callback(null, {result: 'there was an error'});
    }
}

function onImportSet(call, callback) {
    try {
        var scriptPath = "../js/importset_adapters/" + call.request.adapter;
        var page = call.request.page;
        var pagesize = call.request.pagesize;

        delete require.cache[require.resolve(scriptPath)]

        var script = require(scriptPath)
        var importsetResponse = script( page, pagesize )

        callback(null, importsetResponse);
    } catch ( e ) {
        console.log('****************** error: ' + e )
        callback(null, {result: 'there was an error'});
    }
}

function main() {
  var server = new grpc.Server();

  server.addService(
      grpc_script.ScriptService.service, {
          onBusinessRule: onBusinessRule,
          onScriptCall: onScriptCall,
          onImportSet: onImportSet
      });

  server.bindAsync('0.0.0.0:50051', grpc.ServerCredentials.createInsecure(), () => {
      console.log('starting SCRIPT server...');
      console.log('...business rules');
      console.log('...importsets');
      console.log('...general scripts');

      server.start();
  });
}

main();
