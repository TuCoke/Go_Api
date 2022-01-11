package main

import(
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main(){
	router :=httprouter.New()
	router.GET(
		"/CreateUser",
		CreateUser,
	)
	router.GET("/user/:user_name",Login)
	http.ListenAndServe(":8080",router)
}
