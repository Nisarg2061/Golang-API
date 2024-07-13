package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomeSuccess(t *testing.T){
  testServer := httptest.NewServer(http.HandlerFunc(home))
  defer testServer.Close()

  testclient := testServer.Client()
  res, err := testclient.Get(testServer.URL)
  if err != nil {
    t.Error(err)
  }
  
  body, err := io.ReadAll(res.Body)
  if err != nil {
    t.Error(err)
  }

  assert.Equal(t, "Hello World!\n", string(body))
  assert.Equal(t, 200, res.StatusCode)
}

func TestHomeFail(t *testing.T){
  testServer := httptest.NewServer(http.HandlerFunc(home))
  defer testServer.Close()

  testclient := testServer.Client()

  body := strings.NewReader("something...")
  res, err := testclient.Post(testServer.URL, "application/json", body)
  if err != nil {
    t.Error(err)
  }
  
  assert.Equal(t, 405, res.StatusCode)
}

func TestNewEnd(t *testing.T){
  testServer := httptest.NewServer(http.HandlerFunc(new_end))
  defer testServer.Close()

  testclient := testServer.Client()
  res, err := testclient.Get(testServer.URL)
  if err != nil {
    t.Error(err)
  }
  
  body, err := io.ReadAll(res.Body)
  if err != nil {
    t.Error(err)
  }

  assert.Equal(t, "New End Point!\n", string(body))
  assert.Equal(t, 200, res.StatusCode)
}

func TestRoute(t *testing.T){
  testServer := httptest.NewServer(http.HandlerFunc(route))
  defer testServer.Close()

  testclient := testServer.Client()
  res, err := testclient.Get(testServer.URL)
  if err != nil {
    t.Error(err)
  }
  
  body, err := io.ReadAll(res.Body)
  if err != nil {
    t.Error(err)
  }

  assert.Equal(t, "Hidden route found!\n", string(body))
  assert.Equal(t, 200, res.StatusCode)
}
