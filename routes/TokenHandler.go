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
	"go-select/config"
	"github.com/dgrijalva/jwt-go"
)


func LoginHandler(formatter *render.Render) http.HandlerFunc {
	return func (w http.ResponseWriter,req *http.Request) {
		var user entity.UserCredentials
		
			err := json.NewDecoder(req.Body).Decode(&user)
			fmt.Println(user.Username)
			if err != nil {
				// w.WriteHeader(http.StatusForbidden)
				formatter.Text(w, 401, "Error in Request")
				return
			}
		
			err, res := dbservices.FindUserByName(user.Username)
			fmt.Println(err)
			if err != nil {
				// w.WriteHeader(http.StatusInternalServerError)
				formatter.Text(w, 404, "User not found")
				return
			}
			fmt.Println(res.Role)
			if res == nil {
				formatter.Text(w, 404, "User not found")
				return
			}
			claims := entity.CustomerClaim{
				res.Role,
				jwt.StandardClaims{
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			ss, err := token.SignedString(config.MySigningKey)
		
			if err != nil {
				// w.WriteHeader(http.StatusInternalServerError)
				formatter.Text(w, 500, "Error extracting the key")
				return
			}
		
			type Token struct {
				Token string `json:"token"`
				Role string `json:"role"`
			}
			response := Token{ss,res.Role}
			// r, err := json.Marshal(response)
			// fmt.Println(r)
			if err != nil {
				// w.WriteHeader(http.StatusInternalServerError)
				formatter.Text(w, 500, "Error in server side")
				return
			}
			formatter.JSON(w,200,response)
	}
	
}

func TokenVerify(tokenString string) (bool,string) {
	token, err := jwt.ParseWithClaims(tokenString,&entity.CustomerClaim{},func(token *jwt.Token) (interface{}, error) {
		return []byte(config.MySigningKey), nil
	})
	if err != nil {
		return false, ""
	}
	if claims, ok := token.Claims.(*entity.CustomerClaim);ok && token.Valid {
		return true,claims.Role
	} else {
		return false, ""
	}
}