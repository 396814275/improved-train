package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

const TokenExpireDuration = time.Hour * 24

var CustomSecret = []byte("小秘密")

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	// 可根据需要自行添加字段
	UserID             int64  `json:"user_id"`
	Username           string `json:"username"`
	jwt.StandardClaims        // 内嵌标准的声明
}

// GenToken 生成JWT
func GenToken(userid int64, username string) (aToken string, err error) {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		userid,   //自定义字段
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: (time.Now().Add(TokenExpireDuration)).Unix(),
			Issuer:    "渣渣辉", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	// 使用指定的secret签名并获得完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(CustomSecret)
	//refresh token 不需要存放任何数据
	//rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
	//	ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 720)),
	//	Issuer:    "渣渣辉",
	//}).SignedString(CustomSecret)
	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	var mc = new(CustomClaims)
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (key interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

//func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
//	// 解析token refrshtoken无效直接返回
//	if _, err = jwt.Parse(rToken, func(token *jwt.Token) (key interface{}, err error) {
//		// 直接使用标准的Claim则可以直接使用Parse方法
//		return CustomSecret, nil
//	}); err != nil {
//		return
//	}
//
//	//从旧access token中解析出claims数据
//	var mc CustomClaims
//	//使用自定义claims需要使用ParseWithClaims方法
//	_, err = jwt.ParseWithClaims(aToken, &mc, func(token *jwt.Token) (key interface{}, err error) {
//		return CustomSecret, nil
//	})
//	//v, _ := err.(*jwt.ValidationError)
//	////当access token是过时错误 并且refresh token没有过期是就创建一个新的access token
//	//if v.Errors == jwt.ValidationErrorExpired {
//	//	return GenToken(mc.UserID, mc.Username)
//	//}
//
//	return
//}
