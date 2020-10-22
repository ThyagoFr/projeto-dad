  
var express = require("express");
var morgan = require("morgan");
var compression = require('compression');
var helmet = require('helmet');

var app = express();
app.use(helmet());
app.use(compression()); 

// Serve the static files from the build folder
app.use(express.static( __dirname + "/build"));

app.use(express.static(path.join(__dirname, 'build')));

if(process.env.NODE_ENV === 'production') {
  app.get('/*', function (req, res) {
   	res.sendFile(path.join(__dirname, '/build', 'index.html'));
  });
}
// Listen to port 3000
app.listen(3000);