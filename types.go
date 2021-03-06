package kyokatentacle

import (
	"encoding/json"
	"time"
)

// APIResponse 接口返回json的模型
type APIResponse struct {
	Code int64     `json:"code"`
	Msg  string    `json:"msg"`
	Data []Clan    `json:"data"`
	Ts   time.Time `json:"ts"`
	Full int64     `json:"full"`
}

// Clan 接口返回数据的模型
type Clan struct {
	Rank           int64  `json:"rank"`
	Damage         int64  `json:"damage"`
	ClanName       string `json:"clan_name"`
	MemberNum      int64  `json:"member_num"`
	LeaderName     string `json:"leader_name"`
	LeaderViewerID int64  `json:"leader_viewer_id"`
}

func (api *APIResponse) MarshalJSON() ([]byte, error) {
	type Alias APIResponse
	return json.Marshal(&struct {
		Ts int64 `json:"ts"`
		*Alias
	}{
		Ts:    api.Ts.Unix(),
		Alias: (*Alias)(api),
	})
}

func (api *APIResponse) UnmarshalJSON(data []byte) error {
	type Alias APIResponse
	aux := &struct {
		Ts int64 `json:"ts"`
		*Alias
	}{
		Alias: (*Alias)(api),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	api.Ts = time.Unix(aux.Ts, 0)
	return nil
}
