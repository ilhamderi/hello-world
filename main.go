package main

import(
	fmt
	net/http
)

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"<h1>Hello World</h1>")
}

func check(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"<h1>Hello World</h1>")
}

func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/health-check", check)
	fmt.Println("server starting...")
	http.ListenerAndServe(":3000", nil)
}
