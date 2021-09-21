package vpn

import (
	"fmt"
	"testing"
)

func TestAddVpnAccount(t *testing.T) {
	client := NewVpnClient("", "")
	params := &AddAccountParams{
		Name:        "sdkTest3",
		Note:        "sdk测试3",
		ParentGroup: "/软件研发中心/软件部",
		Passwd:      GetRandom(12),
		RoleName:    "软件部",
	}
	res, err := client.AddAccount(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(params.Passwd, res)
}

func TestDelVpnAccount(t *testing.T) {
	client := NewVpnClient("", "")
	params := &DelAccountParams{Names: "sdkTest3"}
	res, err := client.DelAccount(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
func TestGetAccountDetail(t *testing.T) {
	client := NewVpnClient("", "")
	params := &GetAccountDetailParams{Username: "test0917"}
	res, err := client.GetAccountDetail(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
func TestResetPwd(t *testing.T) {
	client := NewVpnClient("", "")
	params := &ResetPwdParams{
		NewName:     "sdkTest3",
		Note:        "sdk测试3",
		OldName:     "sdkTest3",
		ParentGroup: "/软件研发中心/软件部",
		Passwd:      "123456qaz",
		RoleName:    "软件部",
	}
	res, err := client.ResetPwd(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func TestGetUserList(t *testing.T) {
	client := NewVpnClient("", "")
	params := &GetUserList{Limit: "100"}
	res, err := client.GetUserList(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func TestDisconnectUser(t *testing.T) {
	client := NewVpnClient("", "")
	params := &DisConnectUser{Users: "sdkTest3"}
	res, err := client.DisConnectUser(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
