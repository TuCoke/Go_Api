package defs

type UserCredential struct {
	 UserName string `json:"user_name"`
	 Pwd string `json:"pwd"`
}

// video Model
type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}


type Comment struct{
	Id string
	VideoId string
	AuthorName string
	Content string
}