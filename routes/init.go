package routes
import (
	// "net/http"
	//"net/url"
	// "path/filepath"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"

)
func NewServer() *negroni.Negroni {
	
		formatter := render.New()
	
		n := negroni.Classic()
		mx := mux.NewRouter()
	
		initApiRoutes(mx, formatter)
	
		n.UseHandler(mx)
		return n
}

func initApiRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/multi/get/all", GetAllMultiHandler(formatter)).Methods("GET")
	mx.HandleFunc("/single/get/all", GetAllSingleHandler(formatter)).Methods("GET")
	mx.HandleFunc("/update/single/by/id",UpdateSinglesByIdHandler(formatter)).Methods("POST")
	mx.HandleFunc("/create/single", InsertSingleHandler(formatter)).Methods("POST")
	mx.HandleFunc("/create/multi", InsertMultiHandler(formatter)).Methods("POST")
	mx.HandleFunc("/single/get/by/id",GetSinglesByIdHandler(formatter)).Methods("POST")
	// mx.HandleFunc("/multiple/create",
	// mx.HandleFunc("/multiple/get/all",)
}