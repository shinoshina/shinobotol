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
	SubrouterMes["Remove"] = RemoveHandler

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

	post_type := message["post_type"].(string)
	var msg string
	if post_type == "message" {
		msg = message["raw_message"].(string)
		fmt.Print(msg)
	}

	if post_type == "message" {
		_mainType = "Message"
		_subType = SplitRawMessage(msg)
	} else if post_type == "meta_event" {
		_validPost = false
	} else {
		fmt.Printf("herer question")
		_mainType = "Other"
		_subType = SplitOther(message)
	}

	return _mainType, _subType, _validPost

}

func SplitRawMessage(rawMessage string) (subType string) {

	indexForSpeak := strings.Index(rawMessage, "read:")
	indexForImage := strings.Index(rawMessage, "setu!")
	indexForShinoSpeak := strings.Index(rawMessage, "kaka!")

	okForRemove := (strings.Contains(rawMessage,"[CQ:reply,") && 
	strings.Contains(rawMessage,"bukeyi!") && 
	strings.Contains(rawMessage,"[CQ:at,qq=2037310389]"))

	if okForRemove {
		subType = "Remove"
		return 
	}

	if indexForSpeak != -1 {
		subType = "Speak"
	} else if indexForImage != -1 {
		subType = "Image"
	} else if indexForShinoSpeak != -1 {
		subType = "ShinoSpeak"
	} else {
		subType = "Normal"
	}

	return

}

func SplitOther(message map[string]interface{}) (subType string) {

	var _subType string

	_subType = "Poke"

	return _subType
}
