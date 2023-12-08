package cors

import (
	"net/http"

	"github.com/rs/cors"
)

// DevelopmentHandler creates a lax CORS handler to be shared across all services
//
// TODO create a non-dev CORS handler
var DevelopmentHandler = cors.New(cors.Options{
	AllowedMethods: []string{
		http.MethodGet,
		http.MethodPost,
	},
	AllowedOrigins: []string{"*"}, // Set to all origins for ease of development
	AllowedHeaders: []string{
		"Accept-Encoding",
		"Content-Encoding",
		"Content-Type",
		"Connect-Protocol-Version",
		"Connect-Timeout-Ms",
		"Connect-Accept-Encoding",  // Unused in web browsers, but added for future-proofing
		"Connect-Content-Encoding", // Unused in web browsers, but added for future-proofing
		"Grpc-Timeout",             // Used for gRPC-web
		"X-Grpc-Web",               // Used for gRPC-web
		"X-User-Agent",             // Used for gRPC-web
	},
	ExposedHeaders: []string{
		"Content-Encoding",         // Unused in web browsers, but added for future-proofing
		"Connect-Content-Encoding", // Unused in web browsers, but added for future-proofing
		"Grpc-Status",              // Required for gRPC-web
		"Grpc-Message",             // Required for gRPC-web
	},
})
