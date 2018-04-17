/**
 * @Desc
 * @author zjhfyq
 * @data 2018/4/13 14:04.
 */
package models

import (
	"encoding/json"
	"log"
)

type Message struct {
	Code int                    `json:"code"`
	Info string                 `json:"info"`
	Map  map[string]interface{} `json:"map"`
}

func (this Message) Put(key string, value interface{}) Message {
	if this.Map == nil {
		this.Map = make(map[string]interface{})
	}
	this.Map[key] = value
	return this
}

func (this Message) ToJson() string {
	var result string
	js, err := json.Marshal(this)
	if err != nil {
		log.Println(err)
		messgae := Message{}
		messgae.Code = 500
		messgae.Info = "服务器忙..." + err.Error()
		js, _ = json.Marshal(messgae)
	}
	result = string(js)
	return result
}
