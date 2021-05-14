const bucket = require("./bucket")

let db = new bucket();
console.log(db.Get(123));