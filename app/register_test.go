/*
 * 创建时间：2022/5/21
 * 测试内容：注册模块接口：POST请求，http://localhost:8080/douyin/user/register/?username=""&password=""
 * 测试人员：YuYaoMao
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
	"douyin/app/config"
	"douyin/app/dao"
	"douyin/test/tools"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

// 清空所有的表
func prepareDao() func() {
	config.Setup("E:\\awesomeProject\\douyin\\test\\config.yaml")
	dao.Setup()
	dao.TruncateAllTables()
	return dao.TruncateAllTables
}

// 注册测试
func TestRegister(t *testing.T) {
	//prepareDao()
	//模拟一个post提交请求
	res := tools.Read_json("E:\\awesomeProject\\douyin\\test\\Mock\\register.json")
	// 获取[]byte类型的json测试数据
	testData := tools.UsrPwdData{}
	// 解析byte数组获取解析后的测试数据
	testData = tools.ParseRegister(res)
	for i := 0; i < len(testData.Data); i++ {
		postInfo := "http://localhost:8080/douyin/user/register/?username=" +
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
