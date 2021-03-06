package dao

import (
	"database/sql"
	"douyin/pkg/dbx"
	"errors"
	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
)

type Profile struct {
	UserID        int64                  `gorm:"primary_key;autoIncrement:false" json:"id"`
	Name          string                 `gorm:"size:63" json:"name"`
	FollowCount   *int64                 `json:"follow_count"`
	FollowerCount *int64                 `json:"follower_count"`
	Followers     []*Profile             `gorm:"many2many:profile_followers;" json:"-"`
	Videos        []Video                `gorm:"foreignKey:AuthorID;references:UserID" json:"-"`
	Favorites     []*Video               `gorm:"many2many:video_favorites;" json:"-"`
	Comments      []Comment              `gorm:"foreignKey:AuthorID;references:UserID" json:"-"`
	Version       optimisticlock.Version `json:"-"`

	Avatar          string `gorm:"size:255" json:"avatar"`
	BackgroundImage string `gorm:"size:255" json:"background_image"`
	Signature       string `gorm:"size:255" json:"signature"`

	// post-loads
	IsFollow bool `gorm:"-" json:"is_follow,omitempty"`
}

type ProfileFollower struct {
	Model
	ProfileUserID  int64   `gorm:"primaryKey;autoIncrement:false"`
	FollowerUserID int64   `gorm:"primaryKey;autoIncrement:false"`
	User           Profile `gorm:"foreignKey:ProfileUserID;references:UserID" json:"-"`
	Follower       Profile `gorm:"foreignKey:FollowerUserID;references:UserID" json:"-"`
}

func GetProfileByUserId(userId int64) *Profile {
	var profile Profile
	db.First(&profile, userId)
	if profile.UserID > 0 {
		return &profile
	}
	return nil
}

func HasFollower(userId, followerId int64) (bool, error) {
	return dbx.Exists(db, &ProfileFollower{},
		"profile_user_id = ? and follower_user_id = ?",
		userId, followerId)
}

func RemoveFollower(userId, followerId int64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		return removeFollower(tx, userId, followerId)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func removeFollower(tx *gorm.DB, userId, followerId int64) (err error) {
	var follower ProfileFollower
	err = tx.Unscoped().First(&follower,
		"profile_user_id = ? and follower_user_id = ?",
		userId, followerId).Error
	// if no records or record is soft-deleted, no need to delete
	if errors.Is(err, gorm.ErrRecordNotFound) || follower.DeletedAt.Valid {
		return nil
	}
	if err != nil {
		return err
	}

	// decrease FollowerCount with optimistic lock
	err = addFollowerCount(tx, userId, -1)
	if err != nil {
		return err
	}
	// decrease FollowCount with optimistic lock
	err = addFollowCount(tx, followerId, -1)
	if err != nil {
		return err
	}

	// soft delete the record
	return tx.Delete(&follower).Error
}

func AddFollower(userId, followerId int64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		return addFollower(tx, userId, followerId)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func addFollower(tx *gorm.DB, userId, followerId int64) (err error) {
	var follower ProfileFollower
	err = tx.Unscoped().First(&follower,
		"profile_user_id = ? and follower_user_id = ?",
		userId, followerId).Error
	// if there is an undeleted record, no need to add
	if !errors.Is(err, gorm.ErrRecordNotFound) && !follower.DeletedAt.Valid {
		if err != nil {
			return err
		}
		return nil
	}

	// increase FollowerCount with optimistic lock
	err = addFollowerCount(tx, userId, 1)
	if err != nil {
		return err
	}
	// increase FollowCount with optimistic lock
	err = addFollowCount(tx, followerId, 1)
	if err != nil {
		return err
	}

	// if there is a record and soft-deleted, set deleted false
	if follower.DeletedAt.Valid {
		follower.DeletedAt.Valid = false
		return tx.Unscoped().Updates(&follower).Error
	}
	// assign ids and update/insert
	follower.ProfileUserID = userId
	follower.FollowerUserID = followerId
	return tx.Save(&follower).Error
}

// addFollowerCount with optimistic lock
func addFollowerCount(tx *gorm.DB, userId int64, amount int64) error {
	return dbx.SpinOptimisticLock(tx, userId, func(user *Profile) {
		if user.FollowerCount == nil {
			var one int64 = 1
			user.FollowerCount = &one
			return
		}
		*user.FollowerCount += amount
	})
}

// addFollowCount with optimistic lock
func addFollowCount(tx *gorm.DB, userId int64, amount int64) error {
	return dbx.SpinOptimisticLock(tx, userId, func(user *Profile) {
		if user.FollowCount == nil {
			var one int64 = 1
			user.FollowCount = &one
			return
		}
		*user.FollowCount += amount
	})
}

func GetFollows(userId int64) ([]*Profile, error) {
	var profiles []*Profile
	err := db.Joins(
		"inner join profile_followers pf"+
			" on profiles.user_id = pf.profile_user_id"+
			" and pf.follower_user_id = ?",
		userId,
	).Where("pf.deleted_at is null").Find(&profiles).Error
	return profiles, err
}

func GetFollowers(userId int64) ([]*Profile, error) {
	var profiles []*Profile
	err := db.Joins(
		"inner join profile_followers pf"+
			" on profiles.user_id = pf.follower_user_id"+
			" and pf.profile_user_id = ?",
		userId,
	).Where("pf.deleted_at is null").Find(&profiles).Error
	return profiles, err
}
