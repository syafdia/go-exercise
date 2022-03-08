package e2e

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/dghubble/sling"
)

func Test_PostAPIV1AuthorisationCheck_400(t *testing.T) {
	client := &http.Client{}
	req, _ := sling.New().Post("http://0.0.0.0:8080/api/v1/authorisation/check").Request()

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	if !reflect.DeepEqual(resp.StatusCode, http.StatusBadRequest) {
		t.Errorf("http status is not 400, got %d", resp.StatusCode)
	}
}

func Test_PostAPIV1AuthorisationCheck_200(t *testing.T) {
	client := &http.Client{}
	req, _ := sling.New().
		Post("http://0.0.0.0:8080/api/v1/authorisation/check").
		BodyJSON(map[string]interface{}{
			"actions": []string{"compliant"},
			"resource": map[string]interface{}{
				"kind": "VA",
			},
			"principal": map[string]interface{}{
				"kind": "BUSINESS_ID",
				"id":   "abcd-0001-56789-asdf",
			},
		}).
		Request()

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	if !reflect.DeepEqual(resp.StatusCode, http.StatusOK) {
		t.Errorf("http status is not 200, got %d", resp.StatusCode)
	}

	rawBody, _ := ioutil.ReadAll(resp.Body)
	expectedBody := `{"results":[{"kind":"VA","actions":{"compliant":"EFFECT_ALLOWED"}}]}`

	if !reflect.DeepEqual(string(rawBody), expectedBody) {
		t.Errorf("response body is not valid,, got %s", expectedBody)
	}
}
