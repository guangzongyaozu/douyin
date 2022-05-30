/*
 * 文件功能：读取json文件工具类
 */
package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//解析注册的json数据
func ParseRegister(wait []byte) UsrPwdData {
	if wait == nil {
		return UsrPwdData{}
	}
	res := UsrPwdData{}
	//fmt.Println("wait:", wait)
	json.Unmarshal(wait, &res)
	//fmt.Println("parse: ", res)
	return res
}

//解析注册返回的json数据
func ParseResponseRegister(wait []byte) ResponseUsrPwd {
	if wait == nil {
		return ResponseUsrPwd{}
	}
	res := ResponseUsrPwd{}
	//fmt.Println("wait:", wait)
	json.Unmarshal(wait, &res)
	//fmt.Println("parse: ", res)
	return res
}

//【读取json文件】
func Read_json(filepath string) []byte {
	filePtr, err := os.Open(filepath)
	if err != nil {
		fmt.Println("文件打开失败 [Err:%s]", err.Error())
		return nil
	}
	defer filePtr.Close()
	// 创建json解码器
	byteValue, _ := ioutil.ReadAll(filePtr)
	//fmt.Println(string(byteValue))
	return byteValue
}
