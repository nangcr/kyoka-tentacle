package kyokatentacle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type API struct {
	Endpoint string
	Referer  string
	Client   HttpClient
}

func NewAPI() (*API, error) {
	return NewAPIWithClient(ENDPOINT, REFERER, &http.Client{})
}

func NewAPIWIthEndpoint(endpoint, referer string) (*API, error) {
	return NewAPIWithClient(endpoint, referer, &http.Client{})
}

func NewAPIWithClient(endpoint, referer string, client HttpClient) (*API, error) {
	api := &API{
		endpoint,
		referer,
		client,
	}
	return api, nil
}

func (api *API) sendRequest(method, query, payload string) ([]byte, error) {
	request, err := http.NewRequest(method, api.Endpoint+query, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}
	request.Header.Set(`Content-Type`, `application/json`)
	request.Header.Set(`Referer`, `https://kengxxiao.github.io/Kyouka/`)
	response, err := api.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return body, nil
}

// GetLine 获取排名档线
func (api *API) GetLine() ([]Clan, time.Time, error) {
	query := `line`
	payload := `{"history":0}`
	body, err := api.sendRequest(`POST`, query, payload)
	if err != nil {
		return nil, time.Time{}, err
	}

	result := &ApiResponse{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, time.Time{}, err
	}

	return result.Data, result.Ts, nil
}

// GetByRank 按排名查询
func (api *API) GetByRank(rank int) (Clan, time.Time, error) {
	query := fmt.Sprintf(`rank/%d`, rank)
	payload := `{"history":0}`
	body, err := api.sendRequest(`POST`, query, payload)
	if err != nil {
		return Clan{}, time.Time{}, err
	}

	result := &ApiResponse{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return Clan{}, time.Time{}, err
	}

	return result.Data[0], result.Ts, nil
}

// GetByName 按行会名查询
func (api *API) GetByName(name string, page int) ([]Clan, time.Time, int64, error) {
	query := fmt.Sprintf(`name/%d`, page)
	payload := fmt.Sprintf(`{"history":0,"clanName":"%s"}`, name)
	body, err := api.sendRequest(`POST`, query, payload)
	if err != nil {
		return nil, time.Time{}, 0, err
	}

	result := &ApiResponse{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, time.Time{}, 0, err
	}

	return result.Data, result.Ts, result.Full, nil
}

// GetByLeader 按会长名查询
func (api *API) GetByLeader(leader string, page int) ([]Clan, time.Time, int64, error) {
	query := fmt.Sprintf(`leader/%d`, page)
	payload := fmt.Sprintf(`{"history":0,"leaderName":"%s"}`, leader)
	body, err := api.sendRequest(`POST`, query, payload)
	if err != nil {
		return nil, time.Time{}, 0, err
	}

	result := &ApiResponse{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, time.Time{}, 0, err
	}

	return result.Data, result.Ts, result.Full, nil
}
