const express = require('express')
const app = express()
const port = 8080
const bucket = require("./bucket")

let db = new bucket()

app.get('/get', (req, res) => {
    let resp = ""
    for (let i = 0; i < 50; i++) {
        resp += db.Get(123)
    }
    res.send(resp)
})

app.listen(port, () => {
    console.log(`Example app listening at http://localhost:${port}`)
})