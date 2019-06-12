package controller

import "net/http"
import "encoding/json"

type CardController struct{}

func NewCardController() *CardController {
	return &CardController{}
}

func (controller *CardController) GetCVV(res http.ResponseWriter, req *http.Request) {

}

func (controller *CardController) Ping(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode("success")
}
