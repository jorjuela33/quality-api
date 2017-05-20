package authentication

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jorjuela33/quality-api/database"
	"github.com/jorjuela33/quality-api/mssql"
)

const Database = "database"

func New() gin.HandlerFunc {
	return func(context *gin.Context) {
		token, error := jwt.Parse("tokenString", func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return Shared().PublicKey, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["foo"], claims["nbf"])
			database := mssql.New(&database.Options{
				ServerName:   "181.49.12.194",
				DatabaseName: "BD_TEMP",
			})
			_ = database.NewSession()
			context.Set(Database, database)
			context.Next()
		} else {
			fmt.Println(error)
		}
	}
}
