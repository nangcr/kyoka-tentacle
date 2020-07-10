package kyokatentacle

import (
	"testing"
)

func TestAPI_GetLine(t *testing.T) {
	api, err := NewAPI()
	if err != nil {
		t.Error(err)
	}

	_, _, err = api.GetLine()
	if err != nil {
		t.Error(err)
	}
}

func TestAPI_GetByLeader(t *testing.T) {
	api, err := NewAPI()
	if err != nil {
		t.Error(err)
	}

	_, _, _, err = api.GetByLeader("南", 0)
	if err != nil {
		t.Error(err)
	}
}

func TestAPI_GetByRank(t *testing.T) {
	api, err := NewAPI()
	if err != nil {
		t.Error(err)
	}

	_, _, err = api.GetByRank(1)
	if err != nil {
		t.Error(err)
	}
}

func TestAPI_GetByName(t *testing.T) {
	api, err := NewAPI()
	if err != nil {
		t.Error(err)
	}

	_, _, _, err = api.GetByName("幻术", 0)
	if err != nil {
		t.Error(err)
	}
}
