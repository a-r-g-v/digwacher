package main
import (
	"net/http"
	"flag"
	"log"
	"fmt"
)
var(
	assets = flag.String("a","./public_html/.","url startpoint");

)
func test(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello World!!")
}

func main() {
	var(
		mux 	= http.NewServeMux()
		path	http.Handler
	)
	flag.Parse()
	path	= http.FileServer(http.Dir(*assets))
	mux.Handle("/", path)
	mux.Handle("/login", test)
	err := http.ListenAndServe(":11111",mux)
	if err != nil{
		log.Fatal("ListenAndSeve: ", err);
	}
}
