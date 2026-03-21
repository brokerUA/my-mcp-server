package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"my-mcp-server/internal/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var (
	httpAddr = flag.String("http", "", "if set, use streamable HTTP to serve MCP (on this address), instead of stdin/stdout")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	impl := &mcp.Implementation {
	  Name: "my-mcp-server",
	  Version: "0.1.0",
  }
	server := mcp.NewServer(impl, nil)

	// Register tools
	tools.AddToolsToServer(server)

	// Register resources
	tools.AddResourcesToServer(server)

	// Start server with appropriate transport
	if *httpAddr != "" {
		handler := mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server {
			return server
		}, nil)
		log.Printf("MCP server listening at %s", *httpAddr)
		return http.ListenAndServe(*httpAddr, handler)
	} else {
		t := mcp.NewLoggingTransport(mcp.NewStdioTransport(), os.Stderr)
		return server.Run(context.Background(), t)
	}
}
