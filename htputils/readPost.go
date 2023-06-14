package httputils

import (
	"io/ioutil"
	"net/http"
)

// ReadForm checkForm values, false:no values,true:has values
// after that,you can use  r.FormValue("xxx") get post value
func ReadForm(r *http.Request) (bool, error) {
	if err := r.ParseForm(); err != nil {
		return false, err
	}
	return len(r.Form) != 0, nil
}

func ReadBody(r *http.Request) ([]byte, error) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
