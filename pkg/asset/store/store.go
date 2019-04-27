package store

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/openshift/installer/pkg/asset"
)

const (
	stateFileName = ".openshift_install_state.json"
)

type assetSource int

const (
	unfetched	assetSource	= iota
	generatedSource
	onDiskSource
	stateFileSource
)

type assetState struct {
	asset		asset.Asset
	source		assetSource
	anyParentsDirty	bool
	presentOnDisk	bool
}
type storeImpl struct {
	directory	string
	assets		map[reflect.Type]*assetState
	stateFileAssets	map[string]json.RawMessage
	fileFetcher	asset.FileFetcher
}

func NewStore(dir string) (asset.Store, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newStore(dir)
}
func newStore(dir string) (*storeImpl, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	store := &storeImpl{directory: dir, fileFetcher: &fileFetcher{directory: dir}, assets: map[reflect.Type]*assetState{}}
	if err := store.loadStateFile(); err != nil {
		return nil, err
	}
	return store, nil
}
func (s *storeImpl) Fetch(a asset.Asset, preserved ...asset.WritableAsset) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := s.fetch(a, ""); err != nil {
		return err
	}
	if err := s.saveStateFile(); err != nil {
		return errors.Wrap(err, "failed to save state")
	}
	if wa, ok := a.(asset.WritableAsset); ok {
		return errors.Wrap(s.purge(append(preserved, wa)), "failed to purge asset")
	}
	return nil
}
func (s *storeImpl) Destroy(a asset.Asset) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if sa, ok := s.assets[reflect.TypeOf(a)]; ok {
		reflect.ValueOf(a).Elem().Set(reflect.ValueOf(sa.asset).Elem())
	} else if s.isAssetInState(a) {
		if err := s.loadAssetFromState(a); err != nil {
			return err
		}
	} else {
		return nil
	}
	if wa, ok := a.(asset.WritableAsset); ok {
		if err := asset.DeleteAssetFromDisk(wa, s.directory); err != nil {
			return err
		}
	}
	delete(s.assets, reflect.TypeOf(a))
	delete(s.stateFileAssets, reflect.TypeOf(a).String())
	return s.saveStateFile()
}
func (s *storeImpl) DestroyState() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.stateFileAssets = nil
	path := filepath.Join(s.directory, stateFileName)
	err := os.Remove(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return nil
}
func (s *storeImpl) loadStateFile() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	path := filepath.Join(s.directory, stateFileName)
	assets := map[string]json.RawMessage{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	err = json.Unmarshal(data, &assets)
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal state file %q", path)
	}
	s.stateFileAssets = assets
	return nil
}
func (s *storeImpl) loadAssetFromState(a asset.Asset) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, ok := s.stateFileAssets[reflect.TypeOf(a).String()]
	if !ok {
		return errors.Errorf("asset %q is not found in the state file", a.Name())
	}
	return json.Unmarshal(bytes, a)
}
func (s *storeImpl) isAssetInState(a asset.Asset) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := s.stateFileAssets[reflect.TypeOf(a).String()]
	return ok
}
func (s *storeImpl) saveStateFile() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s.stateFileAssets == nil {
		s.stateFileAssets = map[string]json.RawMessage{}
	}
	for k, v := range s.assets {
		if v.source == unfetched {
			continue
		}
		data, err := json.MarshalIndent(v.asset, "", "    ")
		if err != nil {
			return err
		}
		s.stateFileAssets[k.String()] = json.RawMessage(data)
	}
	data, err := json.MarshalIndent(s.stateFileAssets, "", "    ")
	if err != nil {
		return err
	}
	path := filepath.Join(s.directory, stateFileName)
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	if err := ioutil.WriteFile(path, data, 0644); err != nil {
		return err
	}
	return nil
}
func (s *storeImpl) fetch(a asset.Asset, indent string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	logrus.Debugf("%sFetching %q...", indent, a.Name())
	assetState, ok := s.assets[reflect.TypeOf(a)]
	if !ok {
		if _, err := s.load(a, ""); err != nil {
			return err
		}
		assetState = s.assets[reflect.TypeOf(a)]
	}
	if assetState.source != unfetched {
		logrus.Debugf("%sReusing previously-fetched %q", indent, a.Name())
		reflect.ValueOf(a).Elem().Set(reflect.ValueOf(assetState.asset).Elem())
		return nil
	}
	dependencies := a.Dependencies()
	parents := make(asset.Parents, len(dependencies))
	for _, d := range dependencies {
		if err := s.fetch(d, increaseIndent(indent)); err != nil {
			return errors.Wrapf(err, "failed to fetch dependency of %q", a.Name())
		}
		parents.Add(d)
	}
	logrus.Debugf("%sGenerating %q...", indent, a.Name())
	if err := a.Generate(parents); err != nil {
		return errors.Wrapf(err, "failed to generate asset %q", a.Name())
	}
	assetState.asset = a
	assetState.source = generatedSource
	return nil
}
func (s *storeImpl) load(a asset.Asset, indent string) (*assetState, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	logrus.Debugf("%sLoading %q...", indent, a.Name())
	if state, ok := s.assets[reflect.TypeOf(a)]; ok {
		return state, nil
	}
	anyParentsDirty := false
	for _, d := range a.Dependencies() {
		state, err := s.load(d, increaseIndent(indent))
		if err != nil {
			return nil, err
		}
		if state.anyParentsDirty || state.source == onDiskSource {
			anyParentsDirty = true
		}
	}
	var (
		onDiskAsset	asset.WritableAsset
		foundOnDisk	bool
	)
	if _, isWritable := a.(asset.WritableAsset); isWritable {
		onDiskAsset = reflect.New(reflect.TypeOf(a).Elem()).Interface().(asset.WritableAsset)
		var err error
		foundOnDisk, err = onDiskAsset.Load(s.fileFetcher)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to load asset %q", a.Name())
		}
	}
	var (
		stateFileAsset		asset.Asset
		foundInStateFile	bool
		onDiskMatchesStateFile	bool
	)
	if !anyParentsDirty {
		foundInStateFile = s.isAssetInState(a)
		if foundInStateFile {
			stateFileAsset = reflect.New(reflect.TypeOf(a).Elem()).Interface().(asset.Asset)
			if err := s.loadAssetFromState(stateFileAsset); err != nil {
				return nil, errors.Wrapf(err, "failed to load asset %q from state file", a.Name())
			}
		}
		if foundOnDisk && foundInStateFile {
			logrus.Debugf("%sLoading %q from both state file and target directory", indent, a.Name())
			onDiskMatchesStateFile = reflect.DeepEqual(onDiskAsset, stateFileAsset)
			if onDiskMatchesStateFile {
				logrus.Debugf("%sOn-disk %q matches asset in state file", indent, a.Name())
			}
		}
	}
	var (
		assetToStore	asset.Asset
		source		assetSource
	)
	switch {
	case anyParentsDirty:
		if foundOnDisk {
			logrus.Warningf("%sDiscarding the %q that was provided in the target directory because its dependencies are dirty and it needs to be regenerated", indent, a.Name())
		}
		source = unfetched
	case foundOnDisk && !onDiskMatchesStateFile:
		logrus.Debugf("%sUsing %q loaded from target directory", indent, a.Name())
		assetToStore = onDiskAsset
		source = onDiskSource
	case foundInStateFile:
		logrus.Debugf("%sUsing %q loaded from state file", indent, a.Name())
		assetToStore = stateFileAsset
		source = stateFileSource
	default:
		source = unfetched
	}
	state := &assetState{asset: assetToStore, source: source, anyParentsDirty: anyParentsDirty, presentOnDisk: foundOnDisk}
	s.assets[reflect.TypeOf(a)] = state
	return state, nil
}
func (s *storeImpl) purge(excluded []asset.WritableAsset) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	excl := make(map[reflect.Type]bool, len(excluded))
	for _, a := range excluded {
		excl[reflect.TypeOf(a)] = true
	}
	for _, assetState := range s.assets {
		if !assetState.presentOnDisk || excl[reflect.TypeOf(assetState.asset)] {
			continue
		}
		logrus.Infof("Consuming %q from target directory", assetState.asset.Name())
		if err := asset.DeleteAssetFromDisk(assetState.asset.(asset.WritableAsset), s.directory); err != nil {
			return err
		}
		assetState.presentOnDisk = false
	}
	return nil
}
func increaseIndent(indent string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return indent + "  "
}
