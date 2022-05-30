package data

/*
	定义测试需要的一些结构体
*/

// 【定义传参结构】
type UsrPwdData struct {
	Data []UsrPwd `json:"data"`
}

//json数据
type UsrPwd struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//响应数据
type ResponseUsrPwd struct {
	Status_code string
	Status_msg  string
	User_id     string
	token       string
}

// 【返回注册测试数据的结果数组】
func RegisterResult() []string {
	res := []string{"'status_code':1001,'status_msg':'user already exist'"}
	return res
}
