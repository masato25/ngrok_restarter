package main

import (
	"github.com/masato25/ngrok_restarter/services"
	"time"

	"github.com/masato25/ngrok_restarter/settings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	settings.Execute()
	m := viper.GetString("mode")
	log.Infof("use mode: %s", m)

	switch m {
	case "client":
		// create 2 chan for control
		c := make(chan int, 1)
		o := make(chan int, 1)
		o <- 0
		// force delete procress when termianl it!
		defer func() {
			c <- 2
		}()
		for {
			select {
			case vo := <-o:
				// trigger log reader
				if vo == 1 {
					log.Info("@will post")
					services.ReadAndUpdateUrl()
				} else {
					// run script
					go services.OpenNgrokService(c, o)
				}
			case <-time.After(2 * time.Hour):
				// kill process
				//c <- 2
				log.Fatalln("process existing")
			}
		}
	case "server":
		services.WebService()
	default:
		log.Fatalf("invaild mode: %s", m)
	}
}
