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
func GetAllMultiHandler(formatter *render.Render) http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request) {
		err,res := dbservices.FindAllMulti();
		if err != nil {
			formatter.Text(w,200,string(err.Error()))
		} else {
			formatter.JSON(w,200,res)
		}
	}
}

func InsertMultiHandler(formatter *render.Render) http.HandlerFunc {
	return func (w http.ResponseWriter,req *http.Request) {
		req.ParseForm()
		p,_ := ioutil.ReadAll(req.Body)
		var data []entity.Multiple;
		if err := json.Unmarshal(p, &data); err != nil {
			formatter.Text(w,201,err.Error());
		} else {
			fmt.Println(data)
			err := dbservices.InsertMultis(data)
			if err != nil {
				fmt.Println(err)
			}
			formatter.Text(w,200,"OK")
		}
	}
}