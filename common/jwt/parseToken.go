package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	SIGN_NAME_SCERET = "abcdefghijklmnopqrstuvwxyz" // 与Java 保持一致签名密钥
)

//创建 tokenString
func createJwt() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["foo"] = "bar"
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	// claims["nbf"] = time.Now().Unix() + 60 // 60 秒后生效
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(SIGN_NAME_SCERET))
	return tokenString, err
}

//验证
//在调用Parse时，会进行加密验证，同时如果提供了exp，会进行过期验证；
//如果提供了iat，会进行发行时间验证;如果提供了nbf，会进行生效时间验证．
//解析tokenString

func ParseJwt(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(SIGN_NAME_SCERET), nil
	})

	var claims jwt.MapClaims
	var ok bool

	if claims, ok = token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

	return claims
}

func main() {
	fmt.Println("Hello World!")

	tokenString, err := createJwt()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(tokenString)

	claims := ParseJwt(tokenString)
	fmt.Println(claims)

}
