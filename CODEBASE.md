# Codebase Documentation for AI Agent

## Project Summary

This project implements a **Model Context Protocol (MCP)** server in Go. It provides a platform for defining and serving AI tools to be used by MCP-compatible clients (like Claude, IDEs, etc.).

## Architecture

- **MCP Server**: Built using the `github.com/modelcontextprotocol/go-sdk/mcp`.
- **Transports**: Supports both **Stdio** (default) and **Streamable HTTP** (`-http` flag).
- **Tool & Resource Registration**: Uses a dynamic registration pattern where tools and resources are defined in `internal/tools/` and automatically added to the server during startup.
- **MCP Apps (Interactive UI)**: Support for interactive user interfaces via custom resources with `mcp-app` MIME types and tool metadata linking.

## Package Structure

- `cmd/server/`: Contains the main entry point (`main.go`) that initializes and runs the MCP server.
- `internal/tools/`:
    - `all_tools.go`: Defines the `MCPTool` struct and tool registration logic.
    - `all_resources.go`: Defines the `MCPResource` struct and resource registration logic.
    - `story_app.go`: Example of an MCP App implementation (tool + resource + UI).
    - `ui/`: Contains static assets (HTML/JS/CSS) for MCP Apps, embedded into the binary using `go:embed`.
- `kmcp.yaml`: Configuration for the `kmcp` tool used for building and deploying the server.

## Key Components

- **MCPTool Struct**: A generic struct in `internal/tools/all_tools.go` used to define tool names, descriptions, and handlers.
- **MCPResource Struct**: A struct in `internal/tools/all_resources.go` for defining server resources.
- **Dynamic Registration**: Each tool or resource in `internal/tools/` calls `registerTool` or `registerResource` in its `init()` function.

## Development & Deployment

- **kmcp**: The project uses `kmcp` for building Docker images and deploying to Kubernetes as an `MCPServer` custom resource.
- **Docker**: A `Dockerfile` is provided for containerizing the MCP server.
- **Mise**: Uses `mise.toml` for tool and environment management.

## MCP Apps Implementation

- **Resource Metadata**: UI resources are marked with `mcp.Meta{"ui": true}` and use the `text/html;profile=mcp-app` MIME type.
- **Tool-Resource Linking**: Tools can trigger UI display by including the resource URI in their metadata: `mcp.Meta{"ui": map[string]any{"resourceUri": "..."}}`.
- **UI Embedding**: Uses `embed.FS` to include UI assets directly in the Go binary, served as MCP resources.

## Code Patterns

- **Generic Tool Handlers**: Leveraging Go generics in `MCPTool` to handle various input/output types.
- **Decoupled Tools**: Tools are defined in separate files within `internal/tools/`, making it easy to add or remove functionality without modifying the main server logic.
- **Environment Driven**: Configurable via flags (like `-http`) for different execution environments.
