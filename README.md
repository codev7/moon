# moon

moon is a minimal backend web application framework written in Go. It achieves its purpose of rapid prototyping for the SPA pattern with opinionated routing and webpack.

The webpack config is intentionally basic - the purpose of the project is not to mandate frontend decisions.

React is a good choice for a frontend framework to integrate with moon.

# Config

_toml_ is used to configure moon applications.

```toml
[common]	
env = "dev" # dev|prod
js = "bundle.js" # Auto-prefixed with root
style = "bundle.css" # Auto-prefixed with root
static = "static" # Images, stubs, etc
api = "/api/" # prefix for api endpoints 

[server]
template = "template.html.tpl"
address = "localhost:8888"

[client]
entry = "entry.js"
dir = "client"
```

# Routing

The router is setup to serve 3 types of routes:

1. API endpoints - prefixed with the `api` config entry 
2. Static files stored in the directory set in the `static` config entry
3. Default to client routing. If the prior two cases are not satisfied, the server will send the frontend application code. The application takes the form of an html5 template with a link and script tag pointing to the css and js bundles that were output by webpack. 

# Usage

Things you will need:

1. Go
2. Node.js
3. npm

First thing you want to do is clone this repo. Enter the cloned directory and type:

`npm install`

Open up a separate tty and type:

`webpack --watch --colors --progress`

Make sure you have the Go dependencies to run the server. Assuming you have set your GOPATH, just manually install the 3 dependencies:

```bash
go get github.com/Sirupsen/logrus
go get github.com/BurntSushi/toml
go get github.com/julienschmidt/httprouter
```

A simple Makefile is provided to build and run the server:

`make run-server`

Navigate to `http://localhost:8888` and verify that "Hello World" is output in your browser's console.

Check `http://localhost:8888/api/version` to see the endpoint support.

# Adding Endpoints

A simple version endpoint is provided in the `/server/api.go` file. You can manage all of your endpoints from this file. Endpoints must implement `httprouter.Handle`. In other words they need the same signature as the `VersionEndpoint` function. 

To add your own endpoint, create a new function with the same signature as `VersionEndpoint`. Then add the endpoint as shown below. 

Endpoints can have parameterized syntax provided by `httprouter`. Just remember that the endpoint will be prefixed with whatever you have set in your config. For example:

```go
s.Endpoint("/user/:id/:action", API_POST, UserEndpoint)
```

Could be reached at `http://localhost/api/user/1/edit`

Note that `API_GET`, `API_POST`, and `API_BOTH` are bit masked in order to determine allowed methods for the router.

# Todo

- live reloading
