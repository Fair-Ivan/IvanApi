package Model

type LoginInfo struct {
	Id   string `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
}

type LoginResult struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

func LoginCheck(loginReq *LoginInfo) (bool, *LoginInfo) {
	flag := false
	if loginReq.Id == "1" && loginReq.Name == "1" {
		flag = true
	}
	return flag, loginReq
}
