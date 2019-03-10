package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var getParam string = r.URL.Path
	stringNum := strings.Split(getParam, "/")
	//fmt.Fprintf(w, "%q", stringNum)
	num, err := strconv.Atoi(stringNum[1])

	var responseValue string = ""

	if err != nil {
		fmt.Fprintf(w, "not integer")
	}

	for i := 1; i < 10; i++ {
		var n = num * i
		responseValue += strconv.Itoa(num)
		responseValue += " * "
		responseValue += strconv.Itoa(i)
		responseValue += " = "
		responseValue += strconv.Itoa(n)
		responseValue += " "
	}

	fmt.Fprintf(w, "%q", responseValue)
}
