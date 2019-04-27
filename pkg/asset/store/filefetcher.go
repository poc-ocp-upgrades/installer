package store

import (
	"io/ioutil"
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
)

type fileFetcher struct{ directory string }

func (f *fileFetcher) FetchByName(name string) (*asset.File, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	data, err := ioutil.ReadFile(filepath.Join(f.directory, name))
	if err != nil {
		return nil, err
	}
	return &asset.File{Filename: name, Data: data}, nil
}
func (f *fileFetcher) FetchByPattern(pattern string) (files []*asset.File, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	matches, err := filepath.Glob(filepath.Join(f.directory, pattern))
	if err != nil {
		return nil, err
	}
	files = make([]*asset.File, 0, len(matches))
	for _, path := range matches {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}
		filename, err := filepath.Rel(f.directory, path)
		if err != nil {
			return nil, err
		}
		files = append(files, &asset.File{Filename: filename, Data: data})
	}
	return files, nil
}
