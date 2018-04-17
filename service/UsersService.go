/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/13 14:44.
 */
package service

import (
	"user-provider/models"
	"user-provider/db"
	"log"
	"strconv"
)

func AddUsers(users models.Users) bool {
	isAdd := false
	sql := "INSERT INTO users(id_card , account_status_id , role_id , realname , phone , password , mail) "+
		" VALUES (?,?,?,?,?,?,?)"
	result , err :=db.Insert(sql,users.IdCard , users.AccountStatusId ,users.RoleId , users.RealName ,
			users.Phone , users.Password ,users.Password)
	if err != nil {
		log.Println(err)
	}else {
		row , _ := result.RowsAffected()
		if row >0 {
			isAdd = true
		}
	}
	return isAdd
}


func GetUsersByIdCard(id string) models.UsersResp  {
	users := models.UsersResp{}
	sql:= "SELECT id_card , users.account_status_id as uas_id,users.role_id as ur_id, realname , phone , password ,mail ," +
		" account_status.account_status_id as asas_id,account_status.description  asd," +
		" role.role_id as rr_id, role.description as rd from users , account_status , role " +
		" where id_card = ?  and users.account_status_id = account_status.account_status_id " +
		" and users.role_id = role.role_id "
	result , err := db.Select(sql,id)
	if err != nil {
		log.Println(err)
	}else {
		if len(result) == 1 {
			//账户信息
			users.AccountStatusInfo.AccountStatusId,_ = strconv.Atoi(result[0]["asas_id"])
			users.AccountStatusInfo.Description,_ = result[0]["asd"]

			//角色信息
			users.RoleInfo.RoleId ,_ = strconv.Atoi(result[0]["rr_id"])
			users.RoleInfo.Description = result[0]["rd"]

			//用户信息
			users.UsersInfo.IdCard = result[0]["id_card"]
			users.UsersInfo.RoleId ,_ = strconv.Atoi(result[0]["ur_id"])
			users.UsersInfo.AccountStatusId , _ = strconv.Atoi(result[0]["uas_id"])
			users.UsersInfo.Phone = result[0]["phone"]
			users.UsersInfo.Mail = result[0]["mail"]
			users.UsersInfo.RealName = result[0]["realname"]
			users.UsersInfo.Password = result[0]["password"]
		}
	}
	return users
}

func IsExits(idCard string) bool {
	result := false
	sql := "SELECT id_card from users where id_card=?"
	query , err := db.Select(sql,idCard)
	if err != nil {
		log.Println(err)
	}else {
		if len(query) >= 1 {
			result = true
		}
	}
	return result
}

func QueryPassword(idCard string) string {
	password := ""
	sql:= "SELECT password from  users where id_card =?"
	result ,err := db.Select(sql,idCard)
	if err!= nil {
		log.Println(err)
	}else {
		if len(result) == 1 {
			password = result[0]["password"]
		}
	}
	return  password
}

func UpdatePassword(newPassword string, idCard string) bool  {
	result := false
	sql:= "UPDATE users SET password = ? where id_card = ?"
	query ,err :=db.Update(sql , newPassword, idCard)
	if err != nil {
		log.Println(err)
	}else {
		affected , err := query.RowsAffected()
		if err != nil {
			log.Println(err)
		}else{
			if affected == 1{
				result = true
			}
		}
	}
	return  result
}