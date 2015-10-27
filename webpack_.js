var webpack = require('webpack');
var path = require('path');
var toml = require('toml');
var fs = require('fs');
var config = toml.parse(String(fs.readFileSync("config.toml")));

var plugins = [];
var entries = [
	path.join(__dirname, "client", config.client.entry)
];
var publicPath = path.join("/", config.common.static, "/");

// if the hot setting is set, inject the clientside code
if(config.common.hot) {
	if(!Object.prototype.hasOwnProperty.call(config.common, 'hmr')) {
		throw 'Cannot run hot mode without a hmr server. (hint: set the hmr config option)';
	}
	if(!Object.prototype.hasOwnProperty.call(config.common, 'env') || config.common.env !== 'development') {
		throw 'Cannot run hot mode in production';
	}
	entries.push('webpack-hot-middleware/client?path=http://'+path.join(config.common.hmr, "__webpack_hmr"));
	plugins.push(new webpack.optimize.OccurenceOrderPlugin());
	plugins.push(new webpack.HotModuleReplacementPlugin());
	plugins.push(new webpack.NoErrorsPlugin());

	// update the public path to serve from the hmr server	
	publicPath = "http://" + path.join(config.common.hmr, "/", config.common.static, "/");
}

var webpackConfig = {
	entry: entries,
	output: {
		path: path.join(__dirname, config.common.static), 
		publicPath: publicPath,
		filename: config.common.js
	},
	plugins: plugins
};

var compiler = webpack(webpackConfig);

if(config.common.hot) {
	// our hot reload server
	var app = require('express')();

	// this middleware gives memory-fs capabilities to webpack
	app.use(require("webpack-dev-middleware")(compiler, {
		noInfo: true, publicPath: webpackConfig.output.publicPath
	}));
			
	// this middleware serves hot bundles via EventSource
	app.use(require("webpack-hot-middleware")(compiler));

	// parse the host/port from config
	var host, port, parts;
	parts = config.common.hmr.split(":");
	host = parts[0];
	port = parseInt(parts[1]);
	delete parts;

	// start up the server
	app.listen(port, host, function(err) {
		if(typeof err !== 'undefined') {
			console.log("HMR server bootstrap err: " + err);
			return;
		}
		console.log('Started hmr server at ' + config.common.hmr);
	});
} else {
	compiler.watch({}, function(err, stats) {
		console.log(stats.toString({
			colors: true,
		}))
	});
}
