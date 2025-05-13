package main

import (
	"fmt"
	"github.com/sbigtree/go-itsdangerous"
	"github.com/sbigtree/go-itsdangerous/zlib"
)

type User struct {
	User int `json:"user"`
}

func main() {
	// 保持默认压缩级别即可（和 Node 一致）
	deflateString, err := zlib.DeflateString("123")

	if err != nil {
		return
	}
	fmt.Println(deflateString) // 输出 eJwzNDIGAAEtAJc=

	// 创建签名对象，设置密钥、盐、分隔符等参数
	secret := "0123456789abcdef" // 密钥
	salt := "itsdangerous"       // 盐
	//sep := "."                   // 分隔符
	//derivation := "django-concat" // 密钥派生方式
	//digest := sha1.New()          // 使用 SHA-1 算法，确保哈希函数正确初始化

	// 使用 HMAC 算法签名数据
	signer := itsdangerous.NewSignature(secret, salt, "", "", nil, nil)

	//str := ".eJyrViotTi1SsjKsBQAThANn"
	user := User{User: 2}
	//str_b, _ := json.Marshal(user)
	//str := string(str_b)

	//encode, err := signer.Zip("123")
	//fmt.Println("Zip", itsdangerous.Base64Encode(encode))

	//sign, err := signer.Sign(str)
	sign, err := signer.Dumps(user)
	if err != nil {
		return
	}
	fmt.Println("sign", sign)
	//d := ".eJyrViotTi1SsjKqBQAThgNo.Di3+nEgD9CkuHHuFEZEE/qauyv8="
	d := ".eJyrViotTi1SsjKqBQAThgNo.Di3-nEgD9CkuHHuFEZEE_qauyv8"
	unsign, err := signer.Loads(d)
	if err != nil {
		fmt.Println("unsign", err)
		return
	}

	fmt.Println("unsign", unsign)

}
