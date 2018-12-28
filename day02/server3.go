package main

func main() {

}

func handler(w,http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"%s %s %s\n",r.Method,r.URL,r.Proto)
}
