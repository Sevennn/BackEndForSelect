package routes
import (
	//"github.com/Sevennn/agenda-go-server/service/entity"
	"net/http"
	//"net/url"
	// "path/filepath"
	// "github.com/codegangsta/negroni"
	// "github.com/gorilla/mux"
	"github.com/unrolled/render"
	"go-select/dbservices"
	"encoding/json"
	"go-select/entity"
	"fmt"
	"io/ioutil"  
	
)


func GetAllSingleHandler(formatter *render.Render) http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request) {
		err,res := dbservices.FindAllSingle();
		if err != nil {
			formatter.Text(w,200,string(err.Error()))
		} else {
			formatter.JSON(w,200,res)
		}
	}
}

func InsertSingleHandler(formatter *render.Render) http.HandlerFunc {
	return func (w http.ResponseWriter,req *http.Request) {
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
		var data []entity.Single;
		if err := json.Unmarshal(p, &data); err != nil {
			formatter.Text(w,201,err.Error());
		} else {
			fmt.Println(data)
			err := dbservices.InsertSingles(data)
			if err != nil {
				fmt.Println(err)
			}
			formatter.Text(w,200,"OK")
		}
	}
}
type ids struct {
	Ids []string `json:"ids"`
}

func GetSinglesByIdHandler(formatter *render.Render) http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		p,_ := ioutil.ReadAll(req.Body)

		var data ids
		if err := json.Unmarshal(p, &data); err != nil {
			formatter.Text(w,201,err.Error());
		} else {
			fmt.Println(data)
			err,res := dbservices.FindSinglesByIds(data.Ids);
			if err != nil {
				formatter.Text(w,201,err.Error());
			} else {
				formatter.JSON(w,200,res)
			}
		}
	}
}

