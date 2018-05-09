package routes
import (
	// "net/http"
	//"net/url"
	// "path/filepath"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
	//"net/url"
	"path/filepath"
	// "github.com/codegangsta/negroni"
	// "github.com/gorilla/mux"
	"go-select/dbservices"
	"encoding/json"
	"go-select/entity"
	"fmt"
	"io/ioutil"  

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
	mx.HandleFunc("/update/single/id",UpdateSingleByIdHandler(formatter)).Methods("POST")
	mx.HandleFunc("/update/multi/id",UpdateMultiByIdHandler(formatter)).Methods("POST")
	mx.HandleFunc("/belongs", GetBelongsHandler(formatter)).Methods("POST")
	mx.HandleFunc("/create/single", InsertSingleHandler(formatter)).Methods("POST")
	mx.HandleFunc("/create/multi", InsertMultiHandler(formatter)).Methods("POST")
	mx.HandleFunc("/single/get/by/id",GetSinglesByIdHandler(formatter)).Methods("POST")
	// mx.HandleFunc("/multiple/create",
	// mx.HandleFunc("/multiple/get/all",)
	mx.HandleFunc("/exam/create", AddExamHandler(formatter)).Methods("POST")
	mx.HandleFunc("/exam/get/all", GetAllExamsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/exam/get/{id:[_a-zA-Z0-9]+}",GetExamByIdHandler(formatter)).Methods("GET")
	mx.HandleFunc("/user/login", LoginHandler(formatter)).Methods("POST")
}

func GetBelongsHandler(formatter *render.Render)http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request) {
		req.ParseForm();
		p,_ := ioutil.ReadAll(req.Body)
		type Id struct {
			Id string `json:"id"`
		}
		var data Id;
		if err := json.Unmarshal(p, &data); err != nil {
			formatter.Text(w,500,err.Error());
		} else {
			err,r := dbservices.GetBelongs(data.Id);
			if err != nil {
				formatter.Text(w,201,err.Error());
			} else {
				formatter.JSON(w,200,r)
			}
		}
	}
}

func UpdateSingleByIdHandler(formatter *render.Render) http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		ck, er := req.Cookie("token")
		if er != nil {
			formatter.Text(w,403,er.Error());
			return
		}
		if ok,role := TokenVerify(ck.Value); !ok || role != "teacher" {
			formatter.Text(w,403,er.Error());
			return
		}
		p,_ := ioutil.ReadAll(req.Body)
		var data entity.SingleReq
		if err := json.Unmarshal(p, &data); err != nil {
			formatter.Text(w,500,err.Error());
		} else {
			fmt.Println(data)
			err := dbservices.UpdateSingleById(data);
			if err != nil {
				formatter.Text(w,201,err.Error());
			} else {
				formatter.JSON(w,200,"OK")
			}
		}
	}
}

func UpdateMultiByIdHandler(formatter *render.Render) http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		ck, er := req.Cookie("token")
		if er != nil {
			formatter.Text(w,403,er.Error());
			return
		}
		if ok,role := TokenVerify(ck.Value); !ok || role != "teacher" {
			formatter.Text(w,403,er.Error());
			return
		}
		p,_ := ioutil.ReadAll(req.Body)
		var data entity.MultiReq
		if err := json.Unmarshal(p, &data); err != nil {
			formatter.Text(w,500,err.Error());
		} else {
			fmt.Println(data)
			err := dbservices.UpdateMultiById(data);
			if err != nil {
				formatter.Text(w,201,err.Error());
			} else {
				formatter.JSON(w,200,"OK")
			}
		}
	}
}






func AddExamHandler(formatter *render.Render) http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		ck, er := req.Cookie("token")
		if er != nil {
			formatter.Text(w,403,er.Error());
			return
		}
		if ok,role := TokenVerify(ck.Value); !ok || role != "teacher" {
			formatter.Text(w,403,er.Error());
			return
		}
		p,_ := ioutil.ReadAll(req.Body)
		var data entity.Exam;
		if err := json.Unmarshal(p, &data); err != nil {
			formatter.Text(w,201,err.Error());
		} else {
			fmt.Println(data)
			err := dbservices.AddExam(data);
			if err != nil {
				formatter.Text(w,201,err.Error());
			} else {
				formatter.JSON(w,200,"OK")
			}
		}
	}	
}

func GetAllExamsHandler(formatter *render.Render) http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request) {
		err, data := dbservices.FindAllExams();
		if err != nil {
			formatter.Text(w,201,err.Error())
		} else {
			formatter.JSON(w,200,data);
		}
	}
}


func GetExamByIdHandler(formatter *render.Render) http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		path := filepath.FromSlash(req.RequestURI)
		t,id := filepath.Split(path)
		fmt.Println(t,id)
		err,res := dbservices.GetExamById(id);
		if err != nil {
			formatter.Text(w,201,err.Error());
		} else {
			formatter.JSON(w,200,res)
		}
	
	}
}

