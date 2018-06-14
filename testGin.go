package main

import (
	"bytes"
	bs "encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func makeqrcode(base641 string) image.Image {
	log.Println("Original data:", base641)
	code, err := qr.Encode(base641, qr.L, qr.Unicode)
	// code, err := code39.Encode(base64)
	if err != nil {
		log.Fatal(err)
	}
	if base641 != code.Content() {
		log.Fatal("data differs")
	}
	code, err = barcode.Scale(code, 200, 200)
	if err != nil {
		log.Fatal(err)
	}
	return code
}

func main() {
	http.HandleFunc("/", qrcode) //设置访问的路由
	err1 := http.ListenAndServe(":8080", nil)
	if err1 != nil {
		log.Fatal("ListenAndServe:", err1)
	}
}
func qrcode(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	param, _ := req.Form["str"]
	if param == nil {
		fmt.Printf("1a")
		http.Redirect(w, req, "/?str=", http.StatusFound)
	} else {
		base641 := param[0]
		img := makeqrcode(base641)
		emptyBuff := bytes.NewBuffer([]byte{})         //开辟一个新的空buff
		png.Encode(emptyBuff, img)                     //img写入到buff
		dist := make([]byte, 50000)                    //开辟存储空间
		bs.StdEncoding.Encode(dist, emptyBuff.Bytes()) //buff转成base64
		n := 0
		for i := 0; i < len(dist); i++ {
			if dist[i] == 0 {
				n = i
				break
			}
		}
		str := string(dist[:n])
		fmt.Println(str)
		fmt.Println(len(str))
		ls1 := "<!DOCTYPE html><head><meta charset='utf-8' /></head><body><img src='data:image/png;base64,"
		ls2 := "'/></body></html>"
		ls := ls1 + str + ls2
		fmt.Fprintf(w, ls)
	}
}