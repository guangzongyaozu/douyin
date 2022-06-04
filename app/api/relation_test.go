package api

import (
	"douyin/app/config"
	"douyin/app/dao"
	"douyin/app/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func prepareDao() func() {
	config.Load("../../test/")
	dao.Setup()
	dao.TruncateAllTables()
	return dao.TruncateAllTables
}

func preRegister() func() {

	expectedUsr1 := "relationtest1"
	expectedPwd1 := "relationtest_pwd1"
	expectedUsr2 := "relationtest2"
	expectedPwd2 := "relationtest_pwd2"
	// 注册
	service.Register(expectedUsr1, expectedPwd1)
	service.Register(expectedUsr2, expectedPwd2)

	return nil
}

func TestRelationAction(t *testing.T) {
	restore := prepareDao()
	defer restore()
	preRegister()
	//关注
	addErr := service.Follow(1, 2, false)
	addErr2 := service.Follow(1, 1, false)
	removeErr := service.Follow(1, 2, true)
	removeErr2 := service.Follow(1, 1, true)
	assert.Nil(t, removeErr)
	assert.Nil(t, addErr)
	assert.NotNil(t, removeErr2)
	assert.NotNil(t, addErr2)

}

func TestFollowList(t *testing.T) {
	restore := prepareDao()
	defer restore()
	preRegister()
	s := make([]*dao.Profile, 0)

	follows, err := dao.GetFollows(1)
	assert.Equal(t, s, follows)
	assert.Nil(t, err)

	service.Follow(2, 1, false)
	follows2, err2 := dao.GetFollows(1)
	assert.NotNil(t, follows2)
	assert.Nil(t, err2)

	service.Follow(2, 1, true)
	follows3, err3 := dao.GetFollows(1)
	assert.Equal(t, s, follows3)
	assert.Nil(t, err3)

}
func TestFollowerList(t *testing.T) {
	restore := prepareDao()
	defer restore()
	preRegister()
	s := make([]*dao.Profile, 0)

	follows, err := dao.GetFollowers(1)
	assert.Equal(t, s, follows)
	assert.Nil(t, err)

	service.Follow(1, 2, false)
	follows2, err2 := dao.GetFollowers(1)
	assert.NotNil(t, follows2)
	assert.Nil(t, err2)

	service.Follow(1, 2, true)
	follows3, err3 := dao.GetFollowers(1)
	assert.Equal(t, s, follows3)
	assert.Nil(t, err3)
}
