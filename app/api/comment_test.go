package api

import (
	"douyin/app/config"
	"douyin/app/dao"
	"douyin/app/service"
	"douyin/pkg/security"
	"douyin/pkg/validate"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

//初始化配置文件
func init() {
	config.Load("../../")
	security.Setup(config.Val.Jwt)
	dao.Setup()
	validate.Setup()
}

//为测试使用创建 *gin.Engine实例
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(security.Middleware)
	fmt.Println("SetupRouter")
	router.POST("/douyin/comment/action", CommentAction)
	router.GET("/douyin/comment/list", CommentList)
	// initRouter(router)
	fmt.Println("SetupRouter")
	return router
}

func TestCommentAction(t *testing.T) {

	r := SetupRouter()

	//登录获取token 此处用户和密码为已注册的账号密码
	_, token := service.Login("wangsi", "6666666666")
	fmt.Println(*token)

	//发起post请求
	req, _ := http.NewRequest("POST", "/douyin/comment/action?user_id=1&video_id=1&action_type=1&comment_text=测试&comment_id=0", nil)
	req.Header.Add("Authorization", "Bearer "+*token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println(w.Result())
	assert.Equal(t, http.StatusOK, w.Code)

}
func TestCommentList(t *testing.T) {
	r := SetupRouter()

	//登录获取token 此处用户和密码为已注册的账号密码
	_, token := service.Login("wangsi", "6666666666")
	fmt.Println(*token)
	//向注册的路有发起请求
	req, _ := http.NewRequest("GET", "/douyin/comment/list?video_id=1", nil)
	req.Header.Add("Authorization", "Bearer "+*token)
	w := httptest.NewRecorder()
	//模拟http服务处理请求
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
