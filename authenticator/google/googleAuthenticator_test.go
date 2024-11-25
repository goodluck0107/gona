/*
******************************************************

		File Name: main.go
		Author: An
		Mail:lijian@cmcm.com
		Created Time: 14/11/25 - 10:24:49
		Modify Time: 14/11/25 - 10:24:49
	 ******************************************************
*/
package googleAuthenticator

import (
	"fmt"
	"image/png"
	"os"
	"testing"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func createSecret(ga *GAuth) string {
	secret, err := ga.CreateSecret(16)
	if err != nil {
		return ""
	}
	return secret
}

func getCode(ga *GAuth, secret string) string {
	code, err := ga.GetCode(secret)
	if err != nil {
		return "*"
	}
	return code
}

func verifyCode(ga *GAuth, secret, code string) bool {
	// 1:30sec
	ret, err := ga.VerifyCode(secret, code, 1)
	if err != nil {
		return false
	}
	return ret
}

func TestCreateSecret(t *testing.T) {
	// ga := NewGAuth()
	// for i := 0; i < 10; i++ {
	// 	log.Println(createSecret(ga))
	// }
}

func TestGenCodeAndVerifyCode(t *testing.T) {
	// secret := "LC42VPXL3VUMBCAN"
	// //secret := "IU7B5Q3VBL55Q645"
	// ga := NewGAuth()
	// code := getCode(ga, secret)
	// fmt.Println("code", code)
	// if verifyCode(ga, secret, code) {
	// 	fmt.Println("verifyCode success")
	// } else {
	// 	fmt.Println("verifyCode fail")
	// }
}
func TestVerifyCode(t *testing.T) {
	secret := "CRSUHP3VXJ45KL7U"
	//secret := "IU7B5Q3VBL55Q645"
	ga := NewGAuth()
	code := "962563"
	fmt.Println("code", code)
	if verifyCode(ga, secret, code) {
		fmt.Println("verifyCode success")
	} else {
		fmt.Println("verifyCode fail")
	}
}

func TestCreateQRCode(t *testing.T) {
	qrCode, _ := qr.Encode("otpauth://totp/bydaccount?secret=CRSUHP3VXJ45KL7U&issuer=ljtech", qr.M, qr.Auto)
	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)
	// create the output file
	file, _ := os.Create("qrcode.png")
	defer file.Close()
	// encode the barcode as png
	png.Encode(file, qrCode)
}
