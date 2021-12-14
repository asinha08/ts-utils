package password

import (
	"testing"
)

var testPassword = "test"
var hashPassword = "$2a$10$g8cR1RkliRua38rmEoyVoebGEpumhnR7CkukpTxc9VuXuv5nRvDEO"

func TestDoPasswordMatch(t *testing.T) {
	got := DoPasswordMatch(hashPassword, testPassword)
	if !got {
		t.Errorf("got %t", got)
	}
}

func TestEncryptPassword(t *testing.T) {
	got, err := EncryptPassword(testPassword, 10)
	if got != "" && err != nil {
		t.Errorf("got %s, wanted %s", got, hashPassword)
	}
}
