# Codebase Documentation for AI Agent

## Project Summary

This project implements a **Model Context Protocol (MCP)** server in Go. It provides a platform for defining and serving AI tools to be used by MCP-compatible clients (like Claude, IDEs, etc.).

## Architecture

- **MCP Server**: Built using the `github.com/modelcontextprotocol/go-sdk/mcp`.
- **Transports**: Supports both **Stdio** (default) and **Streamable HTTP** (`-http` flag).
- **Tool Registration**: Uses a dynamic registration pattern where tools are defined in `internal/tools/` and automatically added to the server.

## Package Structure

- `cmd/server/`: Contains the main entry point (`main.go`) that initializes and runs the MCP server.
- `internal/tools/`:
    - `all_tools.go`: Defines the `MCPTool` struct and registration logic (`registerTool`, `AddToolsToServer`).
    - Individual tool files (e.g., `echo.go`): Define specific tools and register themselves via `init()`.
- `kmcp.yaml`: Configuration for the `kmcp` tool used for building and deploying the server.
- `.github/workflows/`: CI/CD pipelines for review and evaluation.

## Key Components

- **MCPTool Struct**: A generic struct in `internal/tools/all_tools.go` used to define tool names, descriptions, and handlers.
- **Dynamic Tool Registration**: Each tool in `internal/tools/` calls `registerTool` in its `init()` function, populating a global slice used during server startup.

## Development & Deployment

- **kmcp**: The project uses `kmcp` for building Docker images and deploying to Kubernetes as an `MCPServer` custom resource.
- **Docker**: A `Dockerfile` is provided for containerizing the MCP server.
- **Mise**: Uses `mise.toml` for tool and environment management.

## Code Patterns

- **Generic Tool Handlers**: Leveraging Go generics in `MCPTool` to handle various input/output types.
- **Decoupled Tools**: Tools are defined in separate files within `internal/tools/`, making it easy to add or remove functionality without modifying the main server logic.
- **Environment Driven**: Configurable via flags (like `-http`) for different execution environments.
