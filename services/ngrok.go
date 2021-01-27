package services

import (
	"bytes"
	"fmt"

	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"time"
)

func OpenNgrokService(c chan int, o chan int) {
	if _, err := os.Stat("log.log"); err == nil {
		log.Println("wil delete file log.log")
		e := os.Remove("log.log")
		if e != nil {
			log.Fatal(e)
		}
	}
	var outb, errb bytes.Buffer
	cmd := exec.Command("./ngrok", "http", "8088", "--log", "log.log")
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	o <- 1
	fmt.Println("out:", outb.String(), "err:", errb.String())
	log.Printf("started with pid %d\n", cmd.Process.Pid)
	for {
		select {
		case cc := <-c:
			log.Printf("will close pid %d\n", cmd.Process.Pid)
			if err := cmd.Process.Kill(); err != nil {
				log.Fatalln(err.Error())
			}
			if cc == 1 {
				if err := cmd.Process.Release(); err != nil {
					log.Println(err.Error())
				}
				time.Sleep(2 * time.Second)
				o <- 0
			}
		default:
			time.Sleep(10 * time.Second)
		}
	}

}
