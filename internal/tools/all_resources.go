package tools

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func AddResourcesToServer(server *mcp.Server) {
	for _, addResourceFunc := range resourcesToAdd {
		addResourceFunc(server)
	}
}

var resourcesToAdd []func(server *mcp.Server)

func registerResource(resource MCPResource) {
	resourcesToAdd = append(resourcesToAdd, func(server *mcp.Server) {
		server.AddResource(&mcp.Resource{
			Name:        resource.Name,
			Description: resource.Description,
			URI:         resource.URI,
			MIMEType:    resource.MIMEType,
			Meta:        resource.Meta,
			Annotations: resource.Annotations,
		}, resource.Handler)
	})
}

type MCPResource struct {
	Name        string
	Description string
	URI         string
	MIMEType    string
	Meta        mcp.Meta
	Annotations *mcp.Annotations
	Handler     mcp.ResourceHandler
}
