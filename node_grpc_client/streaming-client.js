var Flux = require('./promising-flux_record');

async function main() {
    var grUser = Flux("x_user");

    var r = await grUser.newRecord();

    result = await grUser.query();
}

main();
