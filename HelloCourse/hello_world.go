package HelloCourse

import (
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func (h HelloCourse) HelloWorld() (string, error) {
	return "Hello, World", nil
}

func makeHelloWorldEndpoint(svc HelloCoursecontract) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.HelloWorld()
		if err != nil {
			return helloWorldResponse{v, err.Error()}, nil
		}
		return helloWorldResponse{v, ""}, nil
	}
}

type helloWorldRequest struct{}

type helloWorldResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

func decodeHelloWorldRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return helloWorldRequest{}, nil
}
