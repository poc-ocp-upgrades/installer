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
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ret []string
	for _, d := range []string{datadir} {
		machineDir := fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH)
		ret = append(ret, filepath.Join(d, "plugins"))
		ret = append(ret, filepath.Join(d, "plugins", machineDir))
	}
	return ret, nil
}
