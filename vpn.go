package vpn

import (
	"encoding/json"
)

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
	return addVpnRes, err
}

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
	return delVpnRes, err
}
func (v *VpnClient) GetAccountDetail(p *GetAccountDetailParams) (*VpnResult, error) {
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
	getAccountRes := &VpnResult{}
	if err := json.Unmarshal(res, getAccountRes); err != nil {
		return nil, err
	}
	return getAccountRes, err
}

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
	return resetPwd, err
}

func (v *VpnClient) GetUserList(p *GetUserList) (*VpnResult, error) {
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
	getUserList := &VpnResult{}
	if err := json.Unmarshal(res, getUserList); err != nil {
		return nil, err
	}
	return getUserList, err
}

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
	return addVpnRes, err
}
