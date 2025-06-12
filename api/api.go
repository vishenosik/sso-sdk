package api

import (
	"net/http"

	"github.com/vishenosik/gocherry/pkg/errors"
)

type errMap interface {
	Get(err error) int
}

func newErrorsMap(_map_ map[error]int) errMap {
	return errors.NewErrorsMap(http.StatusInternalServerError, _map_)
}
