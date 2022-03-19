package response

import systemmodel "github.com/GanweiYee/feuer/server/model/system"

type LoginResponse struct {
	UserInfo *systemmodel.UserInfo `json:"UserInfo"`
	Token    string                `json:"token"`
}
