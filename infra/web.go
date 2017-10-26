package infra

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func WriteJson(w http.ResponseWriter, data interface{}) {
	if bytes, err := json.Marshal(data); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"message\":" + err.Error() + "\"}"))
	} else {
		w.Write(bytes)
	}
}

func ReadJson(r *http.Request, data interface{}) error {
	if bytes, err := ioutil.ReadAll(r.Body); err != nil {
		return err
	} else if err := json.Unmarshal(bytes, data); err != nil {
		return err
	}
	return nil
}

func WriteError(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	WriteJson(w, map[string]interface{} {
		"message": err.Error(),
	})
}