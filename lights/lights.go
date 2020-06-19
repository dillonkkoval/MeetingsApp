package lights

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	JACK   = "1"
	STAIRS = "2"
	DILLON = "3"
	NICK   = "4"
	RED    = "0"
	GREEN  = "20000"
	YELLOW = "10000"
)

func ChangeLight(lightNum string, color string) {
	client := &http.Client{}
	req, err := http.NewRequest(
		http.MethodPut,
		"http://192.168.1.26/api/daRMbHlh9cO8Xz1LRISXt6la75F7dGRG94LkjeRS/lights/"+lightNum+"/state",
		strings.NewReader(`{"hue":`+color+`}`))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}
