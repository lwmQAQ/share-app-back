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
	Sex        int       `gorm:"type:int;comment:'性别 1为男性,2为女性'" json:"sex"`                                                                              // 性别 1为男性，2为女性
	IPInfo     string    `gorm:"type:text;comment:'ip信息'" json:"ip_info"`                                                                                 // ip信息 (JSON 类型)
	Status     int       `gorm:"type:int;default:0;comment:'使用状态 0.正常 1拉黑'" json:"status"`                                                                // 使用状态 0.正常 1拉黑
	CreateTime time.Time `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);comment:'创建时间'" json:"create_time"`                                // 创建时间
	UpdateTime time.Time `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);comment:'修改时间'" json:"update_time"` // 修改时间
	Email      string    `gorm:"type:varchar(255);unique;not null;comment:'用户邮箱'" json:"email"`                                                           // 用户邮箱
	Password   string    `gorm:"type:varchar(255);not null;comment:'用户密码'" json:"password"`                                                               // 用户密码                                                                                        // 软删除时间
}

type UserBuilder struct {
	user User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		user: User{},
	}
}

func (b *UserBuilder) SetID(id uint64) *UserBuilder {
	b.user.ID = id
	return b
}

func (b *UserBuilder) SetName(name string) *UserBuilder {
	b.user.Name = name
	return b
}

func (b *UserBuilder) SetAvatar(avatar string) *UserBuilder {
	b.user.Avatar = avatar
	return b
}

func (b *UserBuilder) SetSex(sex int) *UserBuilder {
	b.user.Sex = sex
	return b
}

func (b *UserBuilder) SetIPInfo(ipInfo []string) *UserBuilder {
	// 将 []string 转换为逗号分隔的字符串
	b.user.IPInfo = strings.Join(ipInfo, ",")
	return b
}

func (b *UserBuilder) SetStatus(status int) *UserBuilder {
	b.user.Status = status
	return b
}

func (b *UserBuilder) SetEmail(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) SetPassword(password string) *UserBuilder {
	b.user.Password = password
	return b
}

func (b *UserBuilder) Build() User {
	return b.user
}
