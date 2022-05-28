/*
 * 创建时间：2022/5/21
 * 测试内容：登录模块接口
 * 测试人员：YuYaoMao
 * 测试计划：
 * 1. 用户名为空，密码为空 ["", ""]
 * 2. 用户名正确，密码不正确
 * 3. 用户名错误，密码正确
 * 4. 用户名正确，密码正确
 */

package main

import (
	"douyin/test/tools"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

// 登录测试
func TestLogin(t *testing.T) {
	//prepareDao()
	//模拟一个post提交请求
	res := tools.Read_json("E:\\awesomeProject\\douyin\\test\\Mock\\login.json")
	// 获取[]byte类型的json测试数据
	testData := tools.UsrPwdData{}
	// 解析byte数组获取解析后的测试数据
	testData = tools.ParseRegister(res)
	for i := 0; i < len(testData.Data); i++ {
		postInfo := "http://localhost:8080/douyin/user/login/?username=" +
			testData.Data[i].Username + "&password=" + testData.Data[i].Password
		resp, err := http.Post(postInfo, "application/x-www-form-urlencoded", strings.NewReader("id=1"))
		if err != nil {
			fmt.Println(err)
			continue
		}
		//关闭连接
		defer resp.Body.Close()
		//读取报文中所有内容
		body, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)
		//输出内容
		fmt.Println(string(body))
		time.Sleep(2)
	}

}
