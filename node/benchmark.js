const bucket = require("./bucket")

function benchmark(method) {
    var start = +(new Date);
    method(function (callback) {
        var end = +(new Date);
        var difference = end - start;
        callback(difference);
    });
}

let db = new bucket();
benchmark(function (next) {
    let COUNT = 1e6
    for (let i = 0; i < COUNT; i++) {
        db.Get(123)
    }
    next(function (difference) {
        console.log((difference / COUNT * 1e6) + ' ns');
    });
})