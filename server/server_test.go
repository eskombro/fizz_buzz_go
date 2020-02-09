package server

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"
	fb "fizz_buzz_go/fizzbuzz"
)

func TestStartServer(t *testing.T) {
	go StartServer()
	time.Sleep(100 * time.Millisecond)
	resp, err := http.Get("http://localhost:8000/health")
	if err != nil {
		t.Errorf("Error initializing server [StartServer()]")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error initializing server [Health check]")
	}
	if resp.StatusCode != 200 || !reflect.DeepEqual(body, []byte("\"Running\"\n")) {
		t.Errorf("Error initializing server [Health check]")
	}
	testHome(t)
	testFizzBuzz(t)
}

func testHome(t *testing.T) {
	resp, err := http.Get("http://localhost:8000/")
	if err != nil {
		t.Errorf("Error on / route")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Errorf("Error initializing server [Health check]")
	}
}

func testFizzBuzz(t *testing.T) {
	monkey.Patch(fb.AddRequest, func(params fb.FizzBuzzParams) {})
	var jsonStr = []byte(`{"int1":1,"int2":2,"limit":4,"str1":"F","str2":"B"}`)
	req, _ := http.NewRequest(
		"POST",
		"http://localhost:8000/fizzbuzz",
		bytes.NewBuffer(jsonStr),
	)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error on /fizzbuzz route")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error on /fizzbuzz route [response]")
	}
	expected := []byte("\"F, FB, F, FB\"\n")
	if resp.StatusCode != 200 || !reflect.DeepEqual(body, expected) {
		t.Errorf(
			"Error on /fizzbuzz route\nExpected: %s\nGot:      %s", expected, body,
		)
	}
}
