package server

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
)

const (
	API_GET = iota << 1
	API_POST
	API_BOTH = 0xF
)

type defaultApp struct {
	template *template.Template
	data     map[string]string
	filedir  http.Dir
}

func (d *defaultApp) loadTemplate(tfile, js, style string) {
	f, err := os.Open(tfile)
	if err != nil {
		log.Errorln("Tpl err", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Errorln("Tpl read err", err)
		os.Exit(1)
	}
	tpl, err := template.New("app").Parse(string(b))
	if err != nil {
		log.Errorln("Tpl parse err", err)
		os.Exit(1)
	}
	d.data = map[string]string{}
	d.data["Js"] = js
	d.data["Style"] = style

	d.template = tpl
}

func (s *Server) MapRoutes() {
	r := s.router
	cwd, _ := os.Getwd()
	static := path.Join(cwd, s.config.static)
	// create the default app (the route used to serve the client app)
	app := defaultApp{filedir: http.Dir(static)}
	app.loadTemplate(s.config.template, s.config.js, s.config.style)
	// if it's not an api call then we use the app, after first checking
	// if there's a file matching the route
	r.NotFound = app
}

func (r defaultApp) checkStaticFile(w http.ResponseWriter, req *http.Request) bool {
	f, _ := r.filedir.Open(req.RequestURI)
	if f != nil {
		defer f.Close()
		i, err := f.Stat()
		if i.IsDir() {
			return true
		}
		if err != nil {
			log.Errorln("fstat", err)
			http.Error(w, err.Error(), 500)
			return false
		}
		buf := make([]byte, i.Size())
		_, err = f.Read(buf)
		if err != nil {
			log.Errorln("fread", err)
			http.Error(w, err.Error(), 500)
			return false
		}

		_, err = w.Write(buf)
		if err != nil {
			log.Errorln("fwrite", err)
			http.Error(w, err.Error(), 500)
			return false
		}
		return false
	}

	return true
}

func (r defaultApp) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Infof("%s\t%s\t%s", req.RemoteAddr, req.Method, req.RequestURI)
	// check for static file match
	serveApp := r.checkStaticFile(w, req)
	// no file match, let client take care of routing
	if serveApp {
		if err := r.template.Execute(w, r.data); err != nil {
			log.Errorln("tpl exec", err)
			http.Error(w, err.Error(), 500)
			return
		}
	}
}

// Adds an api endpoint
func (s *Server) Endpoint(pattern string, opts int, h httprouter.Handle) {
	log.Debugln("adding endpoint", pattern)
	fpat := path.Join(base(s.config.api), base(pattern))
	if opts&API_GET == API_GET {
		s.router.GET(fpat, h)
	}
	if opts&API_POST == API_POST {
		s.router.POST(fpat, h)
	}
}

func base(s string) string {
	return path.Join("/", s)
}
