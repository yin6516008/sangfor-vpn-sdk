package vpn

import (
	"encoding/json"
	"fmt"
)

// 创建用户
func (v *VpnClient) AddAccount(p *AddAccountParams) (*VpnRes, error) {
	a := NewCommonParams("AddUserCloud", "User")
	m := v.GetSignMap(p, a)
	params, err := HandleParams(p, m)
	if err != nil {
		return nil, err
	}
	query := ParamsToSortQuery(params)
	res, err := v.VpnPost("/cgi-bin/php-cgi/html/delegatemodule/WebApi.php?controler=User&action=AddUserCloud", query)
	if err != nil {
		return nil, err
	}
	addVpnRes := &VpnRes{}
	if err := json.Unmarshal(res, addVpnRes); err != nil {
		return nil, err
	}

	if !addVpnRes.Success {
		return nil, fmt.Errorf("add account failed, err: %s", addVpnRes.Message)
	}

	return addVpnRes, err
}

// 删除用户
func (v *VpnClient) DelAccount(p *DelAccountParams) (*VpnRes, error) {
	a := NewCommonParams("DelUserByNameCloud", "User")
	m := v.GetSignMap(p, a)
	params, err := HandleParams(p, m)
	if err != nil {
		return nil, err
	}
	query := ParamsToSortQuery(params)
	res, err := v.VpnPost("/cgi-bin/php-cgi/html/delegatemodule/WebApi.php?controler=User&action=DelUserByNameCloud", query)
	if err != nil {
		return nil, err
	}
	delVpnRes := &VpnRes{}
	if err := json.Unmarshal(res, delVpnRes); err != nil {
		return nil, err
	}

	if !delVpnRes.Success {
		return nil, fmt.Errorf("del account failed, err: %s", delVpnRes.Message)
	}
	return delVpnRes, err
}

func (v *VpnClient) GetAccountDetail(p *GetAccountDetailParams) (*VpnRes, error) {
	a := NewCommonParams("ExGetUserInfo", "User")
	m := v.GetSignMap(p, a)
	params, err := HandleParams(p, m)
	if err != nil {
		return nil, err
	}
	query := ParamsToSortQuery(params)
	res, err := v.VpnPost("/cgi-bin/php-cgi/html/delegatemodule/WebApi.php?controler=User&action=ExGetUserInfo", query)

	if err != nil {
		return nil, err
	}
	getAccountRes := &VpnRes{}
	if err := json.Unmarshal(res, getAccountRes); err != nil {
		return nil, err
	}

	if !getAccountRes.Success {
		return nil, fmt.Errorf("get account detail failed, err: %s", getAccountRes.Message)
	}
	return getAccountRes, err
}

//
func (v *VpnClient) ResetPwd(p *ResetPwdParams) (*VpnRes, error) {
	a := NewCommonParams("UpdateUserCloud", "User")
	m := v.GetSignMap(p, a)
	params, err := HandleParams(p, m)
	if err != nil {
		return nil, err
	}
	query := ParamsToSortQuery(params)
	res, err := v.VpnPost("/cgi-bin/php-cgi/html/delegatemodule/WebApi.php?controler=User&action=UpdateUserCloud", query)
	if err != nil {
		return nil, err
	}
	resetPwd := &VpnRes{}
	if err := json.Unmarshal(res, resetPwd); err != nil {
		return nil, err
	}

	if !resetPwd.Success {
		return nil, fmt.Errorf("reset pwd failed, err: %s", resetPwd.Message)
	}
	return resetPwd, err
}

// 获取用户列表
func (v *VpnClient) GetUserList(p *GetUserList) (*VpnRes, error) {
	a := NewCommonParams("GetSearchData", "User")
	m := v.GetSignMap(p, a)
	params, err := HandleParams(p, m)
	if err != nil {
		return nil, err
	}
	query := ParamsToSortQuery(params)
	res, err := v.VpnPost("/cgi-bin/php-cgi/html/delegatemodule/WebApi.php?controler=User&action=GetSearchData", query)
	if err != nil {
		return nil, err
	}
	getUserList := &VpnRes{}
	if err := json.Unmarshal(res, getUserList); err != nil {
		return nil, err
	}

	if !getUserList.Success {
		return nil, fmt.Errorf("get user list failed, err: %s", getUserList.Message)
	}
	return getUserList, err
}

// 断开用户连接
func (v *VpnClient) DisConnectUser(p *DisConnectUser) (*VpnRes, error) {
	a := NewCommonParams("KillOnlineUserCloud", "State")
	m := v.GetSignMap(p, a)
	params, err := HandleParams(p, m)
	if err != nil {
		return nil, err
	}
	query := ParamsToSortQuery(params)
	res, err := v.VpnPost("/cgi-bin/php-cgi/html/delegatemodule/WebApi.php?controler=State&action=KillOnlineUserCloud", query)
	if err != nil {
		return nil, err
	}
	addVpnRes := &VpnRes{}
	if err := json.Unmarshal(res, addVpnRes); err != nil {
		return nil, err
	}

	if !addVpnRes.Success {
		return nil, fmt.Errorf("disconnect user failed, err: %s", addVpnRes.Message)
	}
	return addVpnRes, err
}
