var Flux = require('../client_flux_record.js');

async function main(action, table, record_id, originals, current) {
    var xid = record_id;

    var gr = Flux("x_history");
    await gr.open();
    await gr.set("x_table",table);
    await gr.set("x_record_id",xid);

    console.log("current = " + JSON.stringify(current));

    for(var field in current) {
        var value = current[field];

        console.log(field + ":" + value);

        await gr.set("x_field", field);
        await gr.set("x_new", value);
    }

    await gr.insert();
    await gr.close();

    console.log("victory is mine");

    // string cancel_action = 1;
    // string fault = 2;
    // string message_to_display = 3;
    // string error_to_display = 4;
    //
    return {message_to_display: "success"};
}


module.exports = async function (action, table, record_id, originals, current) {
    return await main(action, table, record_id, originals, current);
};