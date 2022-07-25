var newFluxRecord = require('./flux_record');

function main() {
    var grUser = newFluxRecord("x_user");
    // grUser.query();

    // has = grUser.next();

    // setTimeout( () => {
    //     console.log("qurey");
    //     grUser.Query();
    // }, 0);

  //
  //   setTimeout( callme, 1000);
  //
  //   function callme() {
  //       console.log("callme()");
  //
  //
  //       !done && setTimeout( callme, 1000);
  //   }
  //
  //   console.log("bye");
}

main();

// var PROTO_PATH = __dirname + '/../bsn/grpc_record.proto';
//
// var grpc = require('@grpc/grpc-js');
//
// var protoLoader = require('@grpc/proto-loader');
//
// var options = { keepCase: true,
//                 longs: String,
//                 enums: String,
//                 defaults: true,
//                 oneofs: true };
//
// var packageDefinition = protoLoader.loadSync( PROTO_PATH, options );
//
// var record_proto = grpc.loadPackageDefinition(packageDefinition).grpc_record;
//
// var client = new record_proto.RecordService("localhost:9000", grpc.credentials.createInsecure());

//   const stream = client.BabaSays();
//
// var done = false;
//
//   stream.on("data", (data) => {
//       if (data.answer == "no more") {
//           done = true;
//           //stream.close();
//       }
//       console.log(data);
//   });
//   stream.on("error", (e) => { console.log("error! " + e); } );
//   stream.on("end", () => { console.log("ended");  } );
//
//   setTimeout( callme, 1000);
//
//   function callme() {
//       console.log("callme()");
//       stream.write({operation:"Next", param:"param3 from js"})
//
//       !done && setTimeout( callme, 1000);
//   }
//
//   console.log("bye");

