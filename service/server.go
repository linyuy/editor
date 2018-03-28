package service

import (
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			//fmt.Println(root)
		}
	}

	mx.NotFoundHandler = NotImplementedHandler()
	mx.MethodNotAllowedHandler = NotImplementedHandler()
	// mx.HandleFunc("/login", loginHandler(formatter))
	mx.HandleFunc("/upload", saveImageHandler(formatter)).Methods("POST")
	mx.HandleFunc("/", homeHandler(formatter)).Methods("GET")
	mx.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/assets/"))))
	// mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))

}
