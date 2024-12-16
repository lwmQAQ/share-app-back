package models

import (
	"strings"
	"time"
)

// User represents a row from the 'user' table.
type User struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement;comment:'用户id'" json:"id"`                                                                       // 用户id
	Name       string    `gorm:"type:varchar(20);comment:'用户昵称'" json:"name"`                                                                             // 用户昵称
	Avatar     string    `gorm:"type:varchar(255);comment:'用户头像'" json:"avatar"`                                                                          // 用户头像
	Bio        string    `gorm:"type:text;comment:'个人简介'" json:"bio"`                                                                                     // 个人简介
	Sex        int       `gorm:"type:int;comment:'性别 1为男性,2为女性'" json:"sex"`                                                                              // 性别 1为男性，2为女性
	IPInfo     string    `gorm:"type:text;comment:'ip信息'" json:"ip_info"`                                                                                 // ip信息 (JSON 类型)
	Status     int       `gorm:"type:int;default:0;comment:'使用状态 0.正常 1拉黑'" json:"status"`                                                                // 使用状态 0.正常 1拉黑
	Level      int       `gorm:"type:int;default:0;comment:'用户等级'" json:"level"`                                                                          // 用户等级
	Experience int       `gorm:"type:int;default:0;comment:'用户经验值'" json:"experience"`                                                                    // 用户经验值
	CreateTime time.Time `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);comment:'创建时间'" json:"create_time"`                                // 创建时间
	UpdateTime time.Time `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);comment:'修改时间'" json:"update_time"` // 修改时间
	Email      string    `gorm:"type:varchar(255);unique;not null;comment:'用户邮箱'" json:"email"`                                                           // 用户邮箱
	Password   string    `gorm:"type:varchar(255);not null;comment:'用户密码'" json:"password"`                                                               // 用户密码
}

// UserBuilder is a builder for creating User instances.
type UserBuilder struct {
	user User
}

// NewUserBuilder creates a new UserBuilder.
func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		user: User{},
	}
}

// SetID sets the ID for the user.
func (b *UserBuilder) SetID(id uint64) *UserBuilder {
	b.user.ID = id
	return b
}

// SetName sets the name for the user.
func (b *UserBuilder) SetName(name string) *UserBuilder {
	b.user.Name = name
	return b
}

// SetAvatar sets the avatar for the user.
func (b *UserBuilder) SetAvatar(avatar string) *UserBuilder {
	b.user.Avatar = avatar
	return b
}

// SetBio sets the bio for the user.
func (b *UserBuilder) SetBio(bio string) *UserBuilder {
	b.user.Bio = bio
	return b
}

// SetSex sets the sex for the user.
func (b *UserBuilder) SetSex(sex int) *UserBuilder {
	b.user.Sex = sex
	return b
}

// SetIPInfo sets the IP information for the user.
func (b *UserBuilder) SetIPInfo(ipInfo []string) *UserBuilder {
	// 将 []string 转换为逗号分隔的字符串
	b.user.IPInfo = strings.Join(ipInfo, ",")
	return b
}

// SetStatus sets the status for the user.
func (b *UserBuilder) SetStatus(status int) *UserBuilder {
	b.user.Status = status
	return b
}

// SetLevel sets the level for the user.
func (b *UserBuilder) SetLevel(level int) *UserBuilder {
	b.user.Level = level
	return b
}

// SetExperience sets the experience for the user.
func (b *UserBuilder) SetExperience(experience int) *UserBuilder {
	b.user.Experience = experience
	return b
}

// SetEmail sets the email for the user.
func (b *UserBuilder) SetEmail(email string) *UserBuilder {
	b.user.Email = email
	return b
}

// SetPassword sets the password for the user.
func (b *UserBuilder) SetPassword(password string) *UserBuilder {
	b.user.Password = password
	return b
}

// Build returns the constructed User instance.
func (b *UserBuilder) Build() User {
	return b.user
}
