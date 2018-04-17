/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/13 14:12.
 */
package common

import "regexp"

const RegCard  ="\\d{17}[\\d|x]|\\d{15}"

const RegPhone  = "0?(13|14|15|18)[0-9]{9}"

const RegMail  = "\\w[-\\w.+]*@([A-Za-z0-9][-A-Za-z0-9]+\\.)+[A-Za-z]{2,14}"

func IsIdCard(card string) bool {
	var reg = regexp.MustCompile(RegCard)
	return reg.MatchString(card)
}

func IsPhone(phone string) bool  {
	var reg = regexp.MustCompile(RegPhone)
	return reg.MatchString(phone)
}


func IsMail(mail string) bool  {
	var reg = regexp.MustCompile(RegMail)
	return reg.MatchString(mail)
}