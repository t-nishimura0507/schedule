package schedule

import (
	"testing"
)

func TestGetSuccess(t *testing.T) {

	// param
	var param = "2020-01-21"

	// exec
	_, err := Get(param)

	// assertion
	if err != nil {
		t.Error("Response is Error...")
	}
}

func TestGetValidationError(t *testing.T) {

	// param
	var param = "間違ったパラメータ"

	// exec
	_, err := Get(param)

	// assertion
	if err == nil {
		t.Error("Response is Success...")
	}
}
