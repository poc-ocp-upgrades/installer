package rhcos

import (
	"context"
	"net/url"
	"github.com/pkg/errors"
)

func QEMU(ctx context.Context) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	meta, err := fetchRHCOSBuild(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed to fetch RHCOS metadata")
	}
	base, err := url.Parse(meta.BaseURI)
	if err != nil {
		return "", err
	}
	relQEMU, err := url.Parse(meta.Images.QEMU.Path)
	if err != nil {
		return "", err
	}
	return base.ResolveReference(relQEMU).String(), nil
}
