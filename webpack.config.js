var path = require('path');
var toml = require('toml');
var fs = require('fs');
var config = toml.parse(String(fs.readFileSync("config.toml")));

module.exports = {
	entry: path.join(__dirname, "client", config.client.entry),
	output: {
		path: path.join(__dirname, config.common.static), 
		filename: config.common.js
	}
};
