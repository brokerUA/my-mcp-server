package tools

import (
	"context"
	"embed"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

//go:embed ui/story-app/index.html
var uiFS embed.FS

const (
	storyAppResourceURI = "ui://story-app/index.html"
	storyAppMIME        = "text/html;profile=mcp-app"
	storyAppToolName    = "show_story_app"
)

type StoryAppParams struct{}
type StoryAppResult struct{}

func init() {
	registerTool(StoryAppTool())
	registerResource(StoryAppResource())
}

func StoryAppResource() MCPResource {
	// Read file from embedded FS
	htmlData, err := uiFS.ReadFile("ui/story-app/index.html")
	if err != nil {
		panic("failed to read embedded html: " + err.Error())
	}

	return MCPResource{
		Name:        "Story App",
		URI:         storyAppResourceURI,
		Description: "HTML page for story generation",
		MIMEType:    storyAppMIME,
		Meta: mcp.Meta{
			"ui": true,
		},
		Handler: func(ctx context.Context, cc *mcp.ServerSession, params *mcp.ReadResourceParams) (*mcp.ReadResourceResult, error) {
			return &mcp.ReadResourceResult{
				Contents: []*mcp.ResourceContents{
					{
						URI:      storyAppResourceURI,
						MIMEType: storyAppMIME,
						Text:     string(htmlData),
						Meta: mcp.Meta{
							"ui": map[string]any{
								"permissions": map[string]any{},
							},
						},
					},
				},
			}, nil
		},
	}
}

func StoryAppTool() MCPTool[StoryAppParams, StoryAppResult] {
	return MCPTool[StoryAppParams, StoryAppResult]{
		Name:        storyAppToolName,
		Description: "Show story generation app",
		Meta: mcp.Meta{
			"ui": map[string]any{
				"resourceUri": storyAppResourceURI,
			},
		},
		Handler: func(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[StoryAppParams]) (*mcp.CallToolResultFor[StoryAppResult], error) {
			return &mcp.CallToolResultFor[StoryAppResult]{
				Content: []mcp.Content{
					&mcp.TextContent{
						Text: "The story generation form is open. Wait for the final prompt from the user.",
					},
				},
			}, nil
		},
	}
}
