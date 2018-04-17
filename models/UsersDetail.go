/**
 * @Desc
 * @author zjhfyq
 * @data 2018/4/13 12:00.
 */
package models

type UsersDetail struct {
	IdCard    string `json:"idCard"`
	NationId  int    `json:"nationId"`
	PhotoId   string `json:"photoId"`
	ProfileId string `json:"profileId"`
	SexId     int    `json:"sexId"`
	Address   string `json:"address"`
	Birth     string `json:"birth,string"`
}
