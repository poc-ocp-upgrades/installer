package exec

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func globalPluginDirs(datadir string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ret []string
	cdir, err := configDir()
	if err != nil {
		return ret, fmt.Errorf("error finding global config directory: %s", err)
	}
	for _, d := range []string{cdir, datadir} {
		machineDir := fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH)
		ret = append(ret, filepath.Join(d, "plugins"))
		ret = append(ret, filepath.Join(d, "plugins", machineDir))
	}
	return ret, nil
}
