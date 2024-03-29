package main
import (
	"net/http"
	"html/template"
	"flag"
	"log"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/gorilla/mux"
	"code.google.com/p/gcfg"
)


type Config struct {
	Admin struct {
		Id	string
		Passwd	string
	}
	Session struct{
		Secretkey string
	}
}

type Scheme struct {
	Id	int
	HTML	string
}

var(
	assets = flag.String("a","./public_html/.","url startpoint")
	config  = Config{}
)
	
var store sessions.Store

const(
 	SESSION_NAME = "PHPSESSIONID" // joke
	CONFIG_NAME = "config.gcfg"
)


func test(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("template/base.html"))
	t.Execute(w,Scheme{1,"foofoofoo"}) 
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
	
	err := gcfg.ReadFileInto(&config, CONFIG_NAME)
	if err != nil {
		log.Fatal("ReadFileInto:", err)
	}

	store = sessions.NewCookieStore([]byte(config.Session.Secretkey))
	
	path	= http.FileServer(http.Dir(*assets))

	r 	:= mux.NewRouter()
	r.Handle("/",path)
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/check", CheckHandler)
	r.HandleFunc("/test", test)
	http.Handle("/", r)

	err = http.ListenAndServe(":11111",nil)
	if err != nil{
		log.Fatal("ListenAndSeve: ", err);
	}
}
