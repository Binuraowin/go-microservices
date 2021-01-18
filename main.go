package main
import (
	"fmt"
	"io/ioutil"
)
import "net/http"
import "log"
func main(){
	http.HandleFunc("/", func(rw http.ResponseWriter,r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw,"oops",http.StatusBadRequest)
			//rw.WriteHeader(http.StatusBadRequest)
			//rw.Write([]byte("Ooops"))
			return
		}
		fmt.Fprintf(rw,"Heloo %s",d)
		log.Printf("data is %s",d)
	})
	fmt.Println("welcome")
	http.ListenAndServe(":9090",nil)
}