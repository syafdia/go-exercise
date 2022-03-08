package e2e

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/dghubble/sling"
)

func Test_GetPing_200(t *testing.T) {
	client := &http.Client{}
	req, _ := sling.New().Get("http://0.0.0.0:8080/ping").Request()

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	if !reflect.DeepEqual(resp.StatusCode, http.StatusOK) {
		t.Errorf("http status is not 200, got %d", resp.StatusCode)
	}
}
