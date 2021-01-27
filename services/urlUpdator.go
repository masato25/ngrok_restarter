package services

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"time"
)

func ReadAndUpdateUrl() {
	APIURL := fmt.Sprintf("http://%v:%v/ngorkme", viper.GetString("server.host"), viper.GetString("server.port"))
	flag := false
	time.AfterFunc(30*time.Second, func() {
		if flag {
			return
		}
		log.Println("readAndUpdateUrl timeout")
		flag = true
	})
	for !flag {
		if _, err := os.Stat("log.log"); err == nil {
			f, err := os.OpenFile("log.log", os.O_RDONLY, 0644)
			if err != nil {
				log.Fatal(err.Error())
			}
			b, err := ioutil.ReadAll(f)
			if err != nil {
				log.Fatal(err.Error())
			} else {
				out := string(b)
				r := regexp.MustCompile("command_line addr.+url=(.+?)\\W{0,3}\n")
				x := r.FindStringSubmatch(out)
				if len(x) > 1 {
					myurl := x[1]
					response, err := http.PostForm(APIURL, url.Values{
						"myurl": {myurl},
					})
					log.Println(err, response, myurl)
					flag = true
				}
			}
			f.Close()
		} else {
			time.Sleep(2 * time.Second)
			log.Println("log.log not found wait 2 second")
		}
	}
}
