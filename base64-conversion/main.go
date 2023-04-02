package main 
import (
	"fmt"
	"encoding/base64"
)


func main () {
	resp:= []byte("\x00"+"email@gmail.com"+"\x00" + "password")
	fmt.Println(string(resp), resp)
	sEnc:= base64.StdEncoding.EncodeToString([]byte(resp))
	fmt.Println(sEnc)
}