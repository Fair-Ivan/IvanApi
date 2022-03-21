package model

type OssForm struct {
	CheckSum string `form:"checksum" json:"checksum"`
	Content  string `form:"content" json:"content"`
}
