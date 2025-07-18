package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"log/slog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	// Set up slog to log to a file
func main() {
	log.Println("Starting MCP server...")
func main() {
	slog.Info("Starting MCP server...")
	s := server.NewMCPServer(
		"Datadog Demo Agent 🚀",
	slog.Info("MCP server created.")
		server.WithToolCapabilities(false),
	log.Println("MCP server created.")

	slog.Info("MCP server created.")
	getMetricsTool := mcp.NewTool("getmetrics",
		mcp.WithDescription("Retrieve metrics from Datadog based on a query"),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("Datadog metrics query string"),
	slog.Info("Added tool: getmetrics")
	)
	s.AddTool(getMetricsTool, getMetricsHandler)
	slog.Info("Starting stdio server...")

	slog.Info("Added tool: getmetrics")
	log.Println("Starting stdio server...")
	if err := server.ServeStdio(s); err != nil {
	slog.Info("Starting stdio server...")
	if err := server.ServeStdio(s); err != nil {
		slog.Error("Server error", "error", err)
	}
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewMetricsApi(apiClient)
	args, ok := request.Params.Arguments.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid arguments type")
	}
	query, ok := args["query"].(string)
	slog.Info("getMetricsHandler called", "query", query)
		return nil, fmt.Errorf("query argument is missing or not a string")
	}
	log.Printf("getMetricsHandler called for %s\n", query)
		slog.Error("Error when calling MetricsApi.QueryMetrics", "error", err)
	slog.Info("getMetricsHandler called", "query", query)
	if !ok {
		return nil, fmt.Errorf("query argument is missing or not a string")
	}
	resp, r, err := api.QueryMetrics(ddctx, time.Now().AddDate(0, 0, -1).Unix(), time.Now().Unix(), query)
	if err != nil {
		slog.Error("Error when calling MetricsApi.QueryMetrics", "error", err)
		slog.Error("Full HTTP response", "response", r)
	}
		return nil, fmt.Errorf("invalid arguments type")
