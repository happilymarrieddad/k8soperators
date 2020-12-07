const express = require('express');
var cors = require('cors')

const app = express();
const port = 8000;

app.use(cors({
    origin: '*',
    optionsSuccessStatus: 200 // some legacy browsers (IE11, various SmartTVs) choke on 204
}))

app.get('/', (req, res) => {
    res.json([
        {id:1,name:'Home 1'},
        {id:2,name:'Home 2'}
    ])
});

app.get('/about', (req, res) => {
    res.json([
        {id:1,name:'About 1'},
        {id:2,name:'About 2'}
    ])
});

app.listen(port, () => console.log(`Server listening on port ${port}`));
