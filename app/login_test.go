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
	"douyin/app/dao"
	"douyin/app/data"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 登录测试
func TestLogin(t *testing.T) {
	router := Setup("../test/")
	//模拟一个post提交请求
	res := data.Read_json(".\\data\\login.json")
	// 获取[]byte类型的json测试数据
	testData := data.UsrPwdData{}
	except := []int{400, 400, 200, 200}
	// 解析byte数组获取解析后的测试数据
	testData = data.ParseRegister(res)
	for i := 0; i < len(testData.Data); i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "http://localhost:8080/douyin/user/login/?username="+
			string(testData.Data[i].Username)+"&password="+string(testData.Data[i].Password), nil)
		router.ServeHTTP(w, req)
		// 服务端响应 401 未授权的
		assert.Equal(t, except[i], w.Code)
		fmt.Println(w.Body)

		//fmt.Println("-----")
	}
}

//注册接口压力测试
func BenchmarkLogin(b *testing.B) {
	b.StopTimer()  // 调用该函数停止压力测试的时间计数
	b.StartTimer() // 重新开始时间计时
	router := Setup("../test/")
	dao.TruncateAllTables()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST",
			"http://localhost:8080/douyin/user/login/?username='douyintest'"+
				"&password='123456789'", nil)
		router.ServeHTTP(w, req)
	}
}
