package asset

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"github.com/pkg/errors"
)

type State struct{ Contents []Content }
type Content struct {
	Name	string
	Data	[]byte
}

func (s *State) PersistToFile(directory string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s == nil {
		return nil
	}
	for _, c := range s.Contents {
		if c.Name == "" {
			continue
		}
		path := filepath.Join(directory, c.Name)
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return errors.Wrap(err, "failed to create dir")
		}
		if err := ioutil.WriteFile(path, c.Data, 0644); err != nil {
			return errors.Wrap(err, "failed to write file")
		}
	}
	return nil
}
