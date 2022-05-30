/*
 * 测试内容：注册模块接口：POST请求，http://localhost:8080/douyin/user/register/?username=""&password=""
 * 测试计划：
 * 1. 用户名为空，密码为空 ["", ""]
 * 2. 用户名为空，密码不为空 ["", "123456789"]
 * 3. 用户名不为空，密码为空 ["douyintest", ""]
 * 4. 用户名小于4位， 密码8-32位 ["dou", "123456789"]
 * 5. 用户名4-32位， 密码小于8位 ["douyintest2", "123456"]
 * 6. 用户名等于32位，密码等于32位 ["1234567891111111111111111111111", "1234567891111111111111111111111"]
 * 7. 用户名大于32位，密码大于32位 ["1234567891111111111111111111111111", "1234567891111111111111111111111111"]
 * 8. 用户名等于4位，密码等于8位 ["1234", "12345678"]
 * 9. 用户名正确，密码格式正确 ["douyin", "123456789"], 查看数据库中是否存储信息
 * 10. 用户名已存在，密码8-32位 ["douyin", "123456"]
 */
package main

import (
	"douyin/app/dao"
	"douyin/app/data"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

//注册接口测试(功能测试)
func TestRegister(t *testing.T) {
	router := Setup("../test/")
	dao.TruncateAllTables()
	//模拟一个post提交请求
	res := data.Read_json(".\\data\\register.json")
	// 获取[]byte类型的json测试数据
	testData := data.UsrPwdData{}
	// 解析byte数组获取解析后的测试数据
	testData = data.ParseRegister(res)
	for i := 0; i < len(testData.Data); i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "http://localhost:8080/douyin/user/register/?username="+
			string(testData.Data[i].Username)+"&password="+string(testData.Data[i].Password), nil)
		router.ServeHTTP(w, req)
		// 服务端响应 401 未授权的
		assert.Equal(t, 200, w.Code)
		//fmt.Println("-----")
	}
}

//注册接口压力测试
func BenchmarkRegister(b *testing.B) {
	b.StopTimer()  // 调用该函数停止压力测试的时间计数
	b.N = 10       // 发送的请求数量
	b.StartTimer() // 重新开始时间计时
	router := Setup("../test/")
	dao.TruncateAllTables()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST",
			"http://localhost:8080/douyin/user/register/?username='douyintest'"+
				"&password='123456789'", nil)
		router.ServeHTTP(w, req)
	}
}
