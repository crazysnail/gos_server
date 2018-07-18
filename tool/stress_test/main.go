package main

import (
    "fmt"
    "io/ioutil"
    //    "log"
    "net/http"
    //    "strings"
	"bytes"

    //"encoding/json"
)
var count = 0

func PostRequest(number int){
	//str2 := fmt.Sprintf("%d", number)
	req := `{"UserName":"junneyang_"}`
	client := &http.Client{}
    req_new := bytes.NewBuffer([]byte(req))
    request, _ := http.NewRequest("POST", "http://0.0.0.0:18810/", req_new)
    request.Header.Set("Content-type", "application/json")
	response, _ := client.Do(request)
	
    if response != nil {
		if response.StatusCode == 200 {
			body, _ := ioutil.ReadAll(response.Body)
			count++
			fmt.Println(string(body),count)
		}
	}else{
		fmt.Println("no response!",number)
	}
}

func main() {
	done := make(chan bool)

	a := func(){
		for i :=0; i<300; i++ {
			PostRequest(i)
		}
	}

	for k:=0; k<200; k++{
		go a()
	}
	

	<-done
}

