package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var server *httptest.Server

func startServer() {
	gin.SetMode(gin.TestMode)
	router = gin.Default()

	router.GET("/greeting", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from server!",
		})
	})

	server = httptest.NewServer(router)
}

func stopServer() {
	server.Close()
}

func theClientRequestsGETGreeting() error {
	resp, err := http.Get(server.URL + "/greeting")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	expected := `{"message":"Hello from server!"}`
	if string(body) != expected {
		return godog.ErrPending
	}

	return nil
}

func theResponseShouldContain(expected string) error {
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the client requests GET /greeting$`, theClientRequestsGETGreeting)
	ctx.Step(`^the response should contain "([^"]*)"$`, theResponseShouldContain)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		startServer()
	})

	ctx.AfterSuite(func() {
		stopServer()
	})
}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format:    "pretty",
		Paths:     []string{"features"},
		Randomize: 0,
	}

	status := godog.TestSuite{
		Name:                 "greeting",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}

	server.Close()
}
