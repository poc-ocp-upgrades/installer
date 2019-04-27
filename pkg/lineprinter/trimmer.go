package lineprinter

import (
	"strings"
)

type Trimmer struct{ WrappedPrint Print }

func (t *Trimmer) Print(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) > 0 {
		i := len(args) - 1
		arg, ok := args[i].(string)
		if ok {
			args[i] = strings.TrimRight(arg, "\n")
		}
	}
	t.WrappedPrint(args...)
}
