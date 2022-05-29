package homeworkForerrGroup

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

var CommandStart="start"

func Server() {
    args:=os.Args
    if len(args)<2 || args[1]!=CommandStart{
        log.Fatal("Server starts is fail.")
    }
	g := new(errgroup.Group)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello,GopherCon SG")
	})
	g.Go(func() error {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			return err
		}
		return nil
	})
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully is running")
	} else {
		fmt.Println("Server is stop")
	}
}
