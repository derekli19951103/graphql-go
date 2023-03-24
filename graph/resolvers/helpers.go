package resolver

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CKey string

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), CKey("Gin"), c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(CKey("Gin"))
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func EncodeToken(tokenUser string) (string,error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	saltedData := append(salt, []byte(tokenUser)...)
	encodedString := base64.StdEncoding.EncodeToString(saltedData)
	return encodedString, nil
}

type TokenUser struct {
	ID        int
	ExpiresAt string
}

func DecodeToken(token string) (TokenUser, error) {
	 // Decode the encoded string using base64
	 decodedBytes, err := base64.StdEncoding.DecodeString(token)
	 if err != nil {
		 panic(err)
	 }
 
	 // Separate the salt value from the decoded bytes
	 originalBytes := decodedBytes[16:]
	 substrings := strings.Split(string(originalBytes), ",")

    // Print the resulting substrings
    for _, substring := range substrings {
        fmt.Println(substring)
    }
	if len(substrings) != 2 {
		return TokenUser{}, errors.New("invalid token")
	}

	id, err := strconv.Atoi(substrings[0])
    if err != nil {
       return TokenUser{}, errors.New("invalid token user id")
    }

	return TokenUser{
		ID:        id,
		ExpiresAt: substrings[1],
	},nil

}

func GetTokenUser(ctx context.Context) (TokenUser, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return TokenUser{}, err
	}
	auth := gc.Request.Header["Authorization"]
	if len(auth) == 0 {
		return TokenUser{}, errors.New("no auth header")
	}

	return DecodeToken(auth[0])
}