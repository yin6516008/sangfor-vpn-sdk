package vpn

import (
	"encoding/json"
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
		Passwd:      "ffdfdf",
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
	params := &DisConnectUser{Users: ""}
	res, err := client.DisConnectUser(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func TestIsUserExist(t *testing.T) {
	client := NewVpnClient("", "")
	params := &IsUserExist{Username: ""}
	res, err := client.IsUserExist(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, err := json.Marshal(res)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}

func TestIsGroupExist(t *testing.T) {
	client := NewVpnClient("", "")
	params := &IsGroupExist{GroupName: ""}
	res, err := client.IsGroupExist(params)
	if err != nil {
		fmt.Println("111")
		fmt.Println(err)
		return
	}
	marshal, err := json.Marshal(res)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}

func TestAddGroup(t *testing.T) {
	client := NewVpnClient("", "")
	params := &AddGroup{Name: "", ParentGroup: "", Note: ""}
	res, err := client.AddGroup(params)
	if err != nil {
		fmt.Println("111")
		fmt.Println(err)
		return
	}
	marshal, err := json.Marshal(res)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}
