package handlers

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler func(c *gin.Context, message map[string]interface{})

type Router map[string]map[string]Handler
type Subrouter map[string]Handler

var MainRouter Router
var SubrouterMes Subrouter
var SubrouterOther Subrouter

func init() {

	MainRouter = make(Router)
	SubrouterMes = make(Subrouter)
	SubrouterOther = make(Subrouter)

	SubrouterMes["Normal"] = NormalMessageHandler
	SubrouterMes["Image"] = ImgHandler
	SubrouterMes["Speak"] = SpeakHandler
	SubrouterMes["ShinoSpeak"] = ShinoSpeakHandler

	SubrouterOther["Poke"] = PokeHandler

	MainRouter["Message"] = SubrouterMes
	MainRouter["Other"] = SubrouterOther

}

func Do(mainType string, subType string, c *gin.Context, message map[string]interface{}) {

	MainRouter[mainType][subType](c, message)
}

func SplitMessageType(message map[string]interface{}) (mainType string, subType string, validPost bool) {

	var _mainType string
	var _subType string
	var _validPost bool = true

	if message["post_type"].(string) == "message" {
		_mainType = "Message"
		_subType = SplitRawMessage(message["raw_message"].(string))
	} else if message["post_type"].(string) == "meta_event" {
		_validPost = false
	} else {
		fmt.Printf("herer question")
		_mainType = "Other"
		_subType = SplitOther(message)
	}

	return _mainType, _subType, _validPost

}

func SplitRawMessage(rawMessage string) (subType string) {

	var _subType string
	indexForSpeak := strings.Index(rawMessage, "please read:")
	indexForImage := strings.Index(rawMessage, "setu!")
	indexForShinoSpeak := strings.Index(rawMessage,"kaka!")

	if indexForSpeak != -1 {
		_subType = "Speak"
	} else if indexForImage != -1 {
		_subType = "Image"
	} else if indexForShinoSpeak != -1{
		_subType = "ShinoSpeak"
	}else{
		_subType = "Normal"
	}

	return _subType

}

func SplitOther(message map[string]interface{}) (subType string) {

	var _subType string

	_subType = "Poke"

	return _subType
}

