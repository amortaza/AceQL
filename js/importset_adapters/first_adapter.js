module.exports = function (page, pagesize) {
    console.log("Importset First Adapter called! ");
    console.log("page " + page );
    console.log("page size " + pagesize );

    return {

        fields: [ "first", "last", "phone" ],

        rows: [
            {values: ["David","Nia","1-800-ace"] },
            {values: ["Tom","Morrison","911-emergency"] },
        ]
    };
}