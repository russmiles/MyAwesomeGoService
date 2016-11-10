package HelloCourse

import (
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

type HelloCoursecontract interface {
	HelloWorld() (string, error)
}

type HelloCourse struct{}

// AddServices sets up and starts the service.
func AddServices() {
	ctx := context.Background()
	svc := HelloCourse{}

	HelloCourseHandler := httptransport.NewServer(
		ctx,
		makeHelloWorldEndpoint(svc),
		decodeHelloWorldRequest,
		encodeResponse,
	)

	http.Handle("/", HelloCourseHandler)
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
