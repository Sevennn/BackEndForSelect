package entity
import (
	"github.com/dgrijalva/jwt-go"
)
type UserCredentials struct {
    Username string `json:"username"`
}


type User struct {
	Username string `json:"username"`
	Role string `json:"role"`
}

type CustomerClaim struct {
	Role string `json:"role"`
	jwt.StandardClaims
}