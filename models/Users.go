/**
 * @Desc
 * @author zjhfyq
 * @data 2018/4/13 11:59.
 */
package models

type Users struct {
	IdCard          string `json:"idCard"`
	AccountStatusId int    `json:"accountStatusId"`
	RoleId          int    `json:"roleId"`
	RealName        string `json:"realName"`
	Phone           string `json:"phone"`
	Password        string `json:"-"`
	Mail            string `json:"mail"`
}
