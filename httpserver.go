package main
import (
	"net/http"
	"flag"
	"log"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/gorilla/mux"
)
var(
	assets = flag.String("a","./public_html/.","url startpoint");
	store = sessions.NewCookieStore([]byte("flag_is_here")) 
)
const(
 	SESSION_NAME = "PHPSESSIONID" // joke
)
func test(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello World!!")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, SESSION_NAME)
	session.Values["user"] = "UserName"
	session.Save(r, w)
	fmt.Fprintf(w,"Create Session?")
}
func CheckHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, SESSION_NAME)
	fmt.Fprintf(w,"Hi! %s",session.Values["user"])
}
func main() {
	var(
		path	http.Handler
		
	)
	flag.Parse()
	path	= http.FileServer(http.Dir(*assets))
	r 	:= mux.NewRouter()
	r.Handle("/",path)
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/check", CheckHandler)

	http.Handle("/", r)
	err := http.ListenAndServe(":11111",nil)
	if err != nil{
		log.Fatal("ListenAndSeve: ", err);
	}
}
