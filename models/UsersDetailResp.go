/**
 * @Desc
 * @author zjhfyq
 * @data 2018/4/14 15:03.
 */
package models

type UsersDetailResp struct {
	UsersInfo         Users         `json:"usersInfo"`
	AccountStatusInfo AccountStatus `json:"accountStatusInfo"`
	RoleInfo          Role          `json:"roleInfo"`
	UsersDetailInfo   UsersDetail   `json:"usersDetailInfo"`
	SexInfo           Sex           `json:"sexInfo"`
	NationInfo        Nation        `json:"nationInfo"`
	ProfileInfo       Profile       `json:"profileInfo"`
	PhotoInfo         Photo         `json:"photoInfo"`
}
