var Flux = require('./promising-flux_record');

async function main() {
    var grUser = Flux("x_user");

    console.log(1);
    var r = await grUser.newRecord();
    console.log(">> " + JSON.stringify(r));

    console.log(2);

    result = await grUser.query();
    console.log(">> " + JSON.stringify(r));

    console.log(3);

    // grUser.newRecord().then( (response) => {
    //     grUser.query().then( (response) => {
    //         console.log(JSON.stringify(response));
    //     })
    // });

    // var resp = ;
    // console.log(JSON.stringify(resp));
    //
    // var resp = grUser.query();


}

main();
