package main
	
	import (
		"fmt"
		"net/http"
		"time"
	)
	
	func greet(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! %s", time.Now())
	}
	
	func main() {
		http.HandleFunc("/", greet)
		fmt.Println( http.ListenAndServe(":3000", nil))
		fmt.Println("serving on PORT 3000")
	}
