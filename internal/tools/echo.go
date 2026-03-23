package tools

import (
	"context"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const echoToolName = "echo"

type EchoParams struct {
	Message string `json:"message" jsonschema:"The message to echo back"`
}

type EchoResult struct{}

func init() {
	registerTool(EchoTool())
}

func EchoTool() MCPTool[EchoParams, EchoResult] {
	return MCPTool[EchoParams, EchoResult]{
		Name:        echoToolName,
		Description: "A simple tool that returns the input message as is (echo functionality).",
		Handler: func(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[EchoParams]) (*mcp.CallToolResultFor[EchoResult], error) {
			return &mcp.CallToolResultFor[EchoResult]{
				Content: []mcp.Content{
					&mcp.TextContent{
						Text: params.Arguments.Message,
					},
				},
			}, nil
		},
	}
}
