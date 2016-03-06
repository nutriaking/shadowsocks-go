package web

import (
	"github.com/orvice/shadowsocks-go/mu/user"
	ss "github.com/shadowsocks/shadowsocks-go/shadowsocks"
	"strconv"
)

type User struct {
	Id             int    `json:"id"`
	Port           int    `json:"port"`
	Passwd         string `json:"passwd"`
	Method         string `json:"method"`
	Enable         int    `json:"enable"`
	TransferEnable int64  `json:"transfer_enable"`
	U              int64  `json:"u"`
	D              int64  `json:"d"`
}

func (u User) GetPort() int {
	return u.Port
}

func (u User) GetPasswd() string {
	return u.Passwd
}

func (u User) GetMethod() string {
	return u.Method
}

func (u User) IsEnable() bool {
	if u.Enable == 0 {
		return false
	}
	if u.TransferEnable < (u.U + u.D) {
		return false
	}
	return true
}

func (u User) GetCipher() (*ss.Cipher, error) {
	return ss.NewCipher(u.Method, u.Passwd)
}

func (u User) UpdateTraffic(storageSize int) error {
	dStr := strconv.Itoa(storageSize)
	uStr := string('0')
	return client.UpdateTraffic(u.Id, uStr, dStr)
}

func (u User) GetUserInfo() user.UserInfo {
	user := user.UserInfo{
		Passwd: u.Passwd,
		Port:   u.Port,
		Method: u.Method,
	}
	return user
}