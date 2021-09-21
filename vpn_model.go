package vpn

type VpnRes struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}
type VpnResult struct {
	Code    int                    `json:"code"`
	Success bool                   `json:"success"`
	Result  map[string]interface{} `json:"result"`
}

type AddAccountParams struct {
	Name        string `json:"name"`
	Note        string `json:"note"`
	ParentGroup string `json:"parent_group"`
	Passwd      string `json:"passwd"`
	RoleName    string `json:"role_name"`
}

type DelAccountParams struct {
	Names string `json:"names"`
}

type GetAccountDetailParams struct {
	Username string `json:"username"`
}
type ResetPwdParams struct {
	NewName     string `json:"new_name"`
	Note        string `json:"note"`
	OldName     string `json:"old_name"`
	ParentGroup string `json:"parent_group"`
	Passwd      string `json:"passwd"`
	RoleName    string `json:"role_name"`
}
type GetUserList struct {
	Limit string `json:"limit"`
}
type DisConnectUser struct {
	Users string `json:"users"`
}
