package console

import (
	"io/ioutil"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func LoadDispatchJSON(
	path string,
	logger log.Logger,
) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		level.Error(logger).Log("File reading error", err)
		return nil, err
	}
	return data, nil
}
