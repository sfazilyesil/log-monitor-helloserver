package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestHelloServer(t *testing.T) {
	fmt.Println("Test etmeden olmaz!")

	// uygulamayı başlat
	go func() {
		NewHelloHttpServer("8080").Run()
	}()

	// sunucunun başlaması için biraz bekle (300ms)
	time.Sleep(time.Millisecond * 300)

	// GET /hello isteği yap
	resp, err := http.Get("http://localhost:8080/hello")
	fmt.Printf("resp: %+v , err: %+v \n", resp, err)

	// hata olmasın
	if err != nil {
		t.Error("Not expected an error but got one. ", err)
		t.FailNow()
	}

	// cevap nil olmasın
	if resp == nil {
		t.Error("Response is nil.")
		t.FailNow()
	}

	// http status code 200 OK olsun
	if resp.StatusCode != http.StatusOK {
		t.Error("Expected status 200 OK, but got ", http.StatusOK)
		t.FailNow()
	}

	var body []byte
	expectedBody := `{"msg":"Merhaba ziyaretçi...", "count":1}`
	defer resp.Body.Close()

	// body yi hatasız okuyabilelim
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		t.Error("Not expected an error but got one. ", err)
		t.FailNow()
	}

	// body beklenen JSON mesaj olsun
	if string(body) != expectedBody {
		t.Error("Expected ", string(body), " to be ", expectedBody)
		t.FailNow()
	}

	fmt.Println("Test bitti!")
}
