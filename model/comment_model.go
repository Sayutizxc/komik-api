package model

type Comment struct {
	Name       string `json:"name"`
	ProfilePic string `json:"profile_pic"`
	Comment    string `json:"comment"`
	Date       string `json:"date"`
}
