/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/14 15:02.
 */
package service

import (
	"user-provider/models"
	"user-provider/db"
	"log"
	"strconv"
)



func GetUsersDetailByIdCard(idCard string) models.UsersDetailResp {
	resp := models.UsersDetailResp{}
	sql:= " SELECT " +
		" users.id_card , users.account_status_id ,users.role_id , users.realname , users.phone , users.`password` , users.mail , " +
		" users_detail.nation_id , users_detail.photo_id , users_detail.profile_id , users_detail.sex_id,users_detail.address , users_detail.birth_y_m_d , " +
		" account_status.description as account_status_desc , " +
		" role.description as role_desc ,nation.nation_desc , " +
		" profile.profile_path , " +
		" photo.photo_path , " +
		" sex.sex_desc " +
		" FROM " +
		" users LEFT JOIN users_detail ON users.id_card = users_detail.id_card " +
		"       LEFT JOIN account_status ON users.account_status_id = account_status.account_status_id " +
		"       LEFT JOIN role ON users.role_id = role.role_id  " +
		"       LEFT JOIN nation ON users_detail.nation_id = nation.nation_id  " +
		"       LEFT JOIN `profile` ON users_detail.profile_id = profile.profile_id " +
		"       LEFT JOIN photo ON users_detail.photo_id = photo.photo_id  " +
		"       LEFT JOIN sex ON users_detail.sex_id = sex.sex_id " +
		"       WHERE users.id_card = ?"
	result , err := db.Select(sql , idCard)
	if err != nil {
		log.Println(err)
	}else {
		if len(result) == 1 {
			//账户信息
			resp.AccountStatusInfo.AccountStatusId,_ = strconv.Atoi(result[0]["account_status_id"])
			resp.AccountStatusInfo.Description,_ = result[0]["account_status_desc"]

			//角色信息
			resp.RoleInfo.RoleId ,_ = strconv.Atoi(result[0]["role_id"])
			resp.RoleInfo.Description = result[0]["role_desc"]

			//用户信息
			resp.UsersInfo.IdCard = result[0]["id_card"]
			resp.UsersInfo.RoleId = resp.RoleInfo.RoleId
			resp.UsersInfo.AccountStatusId = resp.AccountStatusInfo.AccountStatusId
			resp.UsersInfo.Phone = result[0]["phone"]
			resp.UsersInfo.Mail = result[0]["mail"]
			resp.UsersInfo.RealName = result[0]["realname"]
			resp.UsersInfo.Password = result[0]["password"]

			//详细信息
			resp.UsersDetailInfo.IdCard = result[0]["id_card"]
			resp.UsersDetailInfo.Address = result[0]["address"]
			resp.UsersDetailInfo.Birth = result[0]["birth_y_m_d"]
			resp.UsersDetailInfo.NationId ,_= strconv.Atoi(result[0]["nation_id"])
			resp.UsersDetailInfo.PhotoId = result[0]["photo_id"]
			resp.UsersDetailInfo.ProfileId = result[0]["profile_id"]
			resp.UsersDetailInfo.SexId ,_ = strconv.Atoi(result[0]["sex_id"])
			//性别信息
			resp.SexInfo.SexId = resp.UsersDetailInfo.SexId
			resp.SexInfo.SexDesc = result[0]["sex_desc"]

			//民族信息
			resp.NationInfo.NationId = resp.UsersDetailInfo.NationId
			resp.NationInfo.NationDesc = result[0]["nation_desc"]

			//照片信息
			resp.PhotoInfo.PhotoId = resp.UsersDetailInfo.PhotoId
			resp.PhotoInfo.PhotoPath = result[0]["photo_path"]

			//头像信息
			resp.ProfileInfo.ProfileId= resp.UsersDetailInfo.ProfileId
			resp.ProfileInfo.ProfilePath = result[0]["profile_path"]

		}
	}
	return resp
}
