package services

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func WebService() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello ^______________^")
		return
	})
	http.HandleFunc("/ngorkme", helloServer)
	serverHost := fmt.Sprintf(":%v", viper.GetString("server.port"))
	log.Infof("server host %v \n", serverHost)
	http.ListenAndServe(serverHost, nil)
}

var myurl string = "null"
var datetime time.Time = time.Now()

func helloServer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		myurl = r.FormValue("myurl")
		datetime = time.Now()
		fmt.Fprintf(w, "hello, here is your updated url: %s\n", myurl)
		fmt.Fprintf(w, "lasted update at: %v\n", datetime)
		return
	}

	fmt.Fprintf(w, "hello, here is your url: %s\n", myurl)
	fmt.Fprintf(w, "lasted update at: %v\n", datetime)
}
