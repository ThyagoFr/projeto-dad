  
var express = require("express");
var morgan = require("morgan");
var compression = require('compression');
var helmet = require('helmet');

var app = express();
app.use(helmet());
app.use(compression()); 

// Serve the static files from the build folder
app.use(express.static( __dirname + "/build"));
// Listen to port 3000
app.listen(3000);