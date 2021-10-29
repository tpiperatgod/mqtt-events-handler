package mqtt_events_handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	log "k8s.io/klog/v2"
)

func MQTTHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "Get request body failed")
		return
	}

	content := string(reqBody)
	msg := fmt.Sprintf("Receive msg from device: %s\n", content)
	log.Info(msg)
	w.WriteHeader(200)
	fmt.Fprint(w, msg)
}
