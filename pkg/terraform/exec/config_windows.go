package exec

import (
	"path/filepath"
	"syscall"
	"unsafe"
)

var (
	shell			= syscall.MustLoadDLL("Shell32.dll")
	getFolderPath	= shell.MustFindProc("SHGetFolderPathW")
)

const CSIDL_APPDATA = 26

func configDir() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dir, err := homeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "terraform.d"), nil
}
func homeDir() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b := make([]uint16, syscall.MAX_PATH)
	r, _, err := getFolderPath.Call(0, CSIDL_APPDATA, 0, 0, uintptr(unsafe.Pointer(&b[0])))
	if uint32(r) != 0 {
		return "", err
	}
	return syscall.UTF16ToString(b), nil
}
