package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func nameParamHandle(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	fmt.Fprintf(w, "Hello, %s", v["PARAM"])
}

func badHandle(w http.ResponseWriter, r *http.Request) {
	status := http.StatusText(http.StatusInternalServerError)
	http.Error(w, status, http.StatusInternalServerError)
}

func dataHandle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Can not read req body", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "I got message: \n%s", body)
}

func headersHandle(w http.ResponseWriter, r *http.Request) {
	res := 0
	for _, v := range []string{"a", "b"} {
		head := r.Header.Get(v)
		number, err := strconv.Atoi(head)
		if err != nil {
			status := fmt.Sprintf("header %s is not a number %d", v, number)
			http.Error(w, status, http.StatusBadRequest)
			return
		}
		res += number
	}
	w.Header().Add("a+b", strconv.Itoa(res))
}
