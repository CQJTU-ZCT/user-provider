/**
 * @Desc
 * @author zjhfyq
 * @data 2018/4/13 17:48.
 */
package models

type UsersResp struct {
	UsersInfo         Users         `json:"usersInfo"`
	AccountStatusInfo AccountStatus `json:"accountStatusInfo"`
	RoleInfo          Role          `json:"roleInfo"`
}
