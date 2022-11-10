package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {
	c, err := ftp.Dial("192.168.59.131:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("anonymous", "anonymous")
	if err != nil {
		log.Fatal(err)
	}

	dirs, err := c.List("")
	if err == nil {
		if len(dirs) > 0 {
			for i := 0; i < len(dirs); i++ {
				fmt.Print("\n" + dirs[i].Name)
			}
		}
	}

	// r, err := c.Retr("1.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer r.Close()

	// buf, err := ioutil.ReadAll(r)
	// err = ioutil.WriteFile("1.txt", buf, 0644)
	// if err != nil {
	// 	panic(err)
	// }
	// Do something with the FTP conn

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}
