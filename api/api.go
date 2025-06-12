package api

import (
	"net/http"

	"github.com/vishenosik/gocherry/pkg/errors"
	"google.golang.org/grpc/codes"
)

type httpErrMap interface {
	Get(err error) int
}

type grpcErrMap interface {
	Get(err error) codes.Code
}

func httpErrorsMap(_map_ map[error]int) httpErrMap {
	return errors.NewErrorsMap(http.StatusInternalServerError, _map_)
}

func grpcErrorsMap(_map_ map[error]codes.Code) grpcErrMap {
	return errors.NewErrorsMap(codes.Internal, _map_)
}
