package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Pickausernaame/HighloadHW2/application"
	"github.com/Pickausernaame/HighloadHW2/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func getConfig(t *testing.T, path string) *models.Config {
	configBytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("unable to open configuration file '" + os.Args[1] + "': " + err.Error())
	}
	config := &models.Config{}
	err = json.NewDecoder(bytes.NewReader(configBytes)).Decode(&config)
	if err != nil {
		t.Errorf("unable to parse configuration file ./config.json : " +
			"it should be json with all fields : " + err.Error())
	}
	return config
}

func executeRequest(req *http.Request, testApp *application.App) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	testApp.Router.ServeHTTP(recorder, req)
	return recorder
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestSuccessGenerateRequest(t *testing.T) {
	testApp, err := application.New(getConfig(t, "../config.json"))
	if err != nil {
		t.Errorf(err.Error())
	}
	succesRequest, _ := http.NewRequest("POST", "/api/generate", strings.NewReader(""))
	succesRequest.Header.Set("Content-Type", "application/json")
	response_201 := executeRequest(succesRequest, testApp)
	checkResponseCode(t, http.StatusCreated, response_201.Code)

}

func TestSuccessGeneration(t *testing.T) {
	testApp, err := application.New(getConfig(t, "../config.json"))
	if err != nil {
		t.Errorf(err.Error())
	}
	succesGenRequest, _ := http.NewRequest("POST", "/api/generate", strings.NewReader(""))
	succesGenRequest.Header.Set("Content-Type", "application/json")
	response_201 := executeRequest(succesGenRequest, testApp)
	checkResponseCode(t, http.StatusCreated, response_201.Code)
	var gen models.Generation
	err = json.NewDecoder(bytes.NewReader(response_201.Body.Bytes())).Decode(&gen)
	id := gen.Id
	succesRetRequest, _ := http.NewRequest("GET", fmt.Sprintf("/api/retrieve/%d", id), strings.NewReader(""))
	succesRetRequest.Header.Set("Content-Type", "application/json")
	response_200 := executeRequest(succesRetRequest, testApp)
	checkResponseCode(t, http.StatusOK, response_200.Code)
	var ret models.Generation
	err = json.NewDecoder(bytes.NewReader(response_200.Body.Bytes())).Decode(&ret)
	if err != nil {
		t.Errorf(err.Error())
	}

	if gen.Id != ret.Id || gen.Data != ret.Data {
		t.Errorf(fmt.Sprintf("Bad retrieve\n ID %d(generated) != %d(retrieved) \n GENERATION  %d(generated) != %d(retrieved)", gen.Id, ret.Id, gen.Data, ret.Id))
	}
}

func TestWrongId(t *testing.T) {
	testApp, err := application.New(getConfig(t, "../config.json"))
	if err != nil {
		t.Errorf(err.Error())
	}

	succesRetRequest, _ := http.NewRequest("GET", "/api/retrieve/woooooooop", strings.NewReader(""))
	succesRetRequest.Header.Set("Content-Type", "application/json")
	response_409 := executeRequest(succesRetRequest, testApp)
	checkResponseCode(t, http.StatusConflict, response_409.Code)
}

func TestMissingId(t *testing.T) {
	testApp, err := application.New(getConfig(t, "../config.json"))
	if err != nil {
		t.Errorf(err.Error())
	}

	succesRetRequest, _ := http.NewRequest("GET", "/api/retrieve/123", strings.NewReader(""))
	succesRetRequest.Header.Set("Content-Type", "application/json")
	response_404 := executeRequest(succesRetRequest, testApp)
	checkResponseCode(t, http.StatusNotFound, response_404.Code)
}
