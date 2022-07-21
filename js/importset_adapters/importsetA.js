module.exports = function (page, pagesize) {
    console.log("Importset A called! ");
    console.log("page " + page );
    console.log("page size " + pagesize );

    return {
        table: "u_import_a",

        fields: [
            {
                fieldname: "u_firstname",
                fieldtype: "String"
            },
            {
                fieldname: "u_lastname",
                fieldtype: "String"
            },
            {
                fieldname: "u_dob",
                fieldtype: "String"
            }
        ],

        rows: [
            {values: ["David","Nia","01-01-1990"] },
            {values: ["Tom","Morrison","02-02-1992"] },
        ]
    };
}