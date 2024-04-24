package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	_ "encoding/json"

	"github.com/cristalhq/jwt/v5"
)

type Claims struct {
	UserRole string `json:"role"`
	Email    string `json:"email"`
}

func encryptPasswordSHA256(password string) string {
	h := sha256.Sum256([]byte(password))
	return hex.EncodeToString([]byte(h[:]))
}

func decriptedJWT(token string) Claims {
	var myClaims Claims
	if token[0:1] == "B" {
		token = token[7:]
	}
	jwtToken, _ := jwt.ParseNoVerify([]byte(token))
	json.Unmarshal(jwtToken.Claims(), &myClaims)
	return myClaims
}
func getJWT(email string, role string) string {
	key := []byte(`1234`)
	signer, _ := jwt.NewSignerHS(jwt.HS256, key)
	builder := jwt.NewBuilder(signer)

	claims := &Claims{
		UserRole: role,
		Email:    email,
	}

	token, _ := builder.Build(claims)
	return token.String()
}

//func main() {
//	//// TODO secret
//	//key := []byte(`1234`)
//	//signer, _ := jwt.NewSignerHS(jwt.HS256, key)
//	//builder := jwt.NewBuilder(signer)
//	//
//	//claims := &Claims{
//	//	UserRole: "user",
//	//	Email:    "foo@bar.baz",
//	//}
//	//
//	//token, _ := builder.Build(claims)
//	//
//	//var retrivedClaims Claims
//	//json.Unmarshal(token.Claims(), &retrivedClaims)
//
//	//mypass := encryptPasswordSHA256("1234")
//	//sql := "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"
//	//fmt.Println("\n", mypass)
//	//expr := (sql == mypass)
//	//fmt.Println("\n", expr)
//}
