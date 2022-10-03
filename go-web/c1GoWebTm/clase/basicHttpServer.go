package clase

import (
	"fmt"
	"net/http"
)

func mihandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hola\n%+v\n", req.URL.Query())
}

func HttpServerBasic() {
	http.HandleFunc("/Pattern", mihandler)
	http.ListenAndServe(":8080", nil)
}
