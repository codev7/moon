var toml = require('toml');
var fs = require('fs');
var path = require('path');

var config = function() {
	var cfg = toml.parse(String(fs.readFileSync("config.toml")));
	return Object.assign({}, cfg.common, cfg.client);
}();

module.exports = {
	entry: path.join(__dirname, config.dir, config.entry),
	output: {
		path: path.join(__dirname, config.static), 
		filename: config.js
	}
};
