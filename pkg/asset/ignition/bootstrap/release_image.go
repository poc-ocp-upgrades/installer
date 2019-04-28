package bootstrap

import (
	"fmt"
	"strings"
)

var (
	defaultReleaseImageOriginal	= "registry.svc.ci.openshift.org/origin/release:v4.1"
	defaultReleaseImagePadded	= "\x00_RELEASE_IMAGE_LOCATION_\x00XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\x00"
	defaultReleaseImagePrefix	= "\x00_RELEASE_IMAGE_LOCATION_\x00"
	defaultReleaseImageLength	= len(defaultReleaseImagePadded)
)

func DefaultReleaseImage() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if strings.HasPrefix(defaultReleaseImagePadded, defaultReleaseImagePrefix) {
		return defaultReleaseImageOriginal, nil
	}
	nullTerminator := strings.IndexByte(defaultReleaseImagePadded, '\x00')
	if nullTerminator == -1 {
		return "", fmt.Errorf("release image location was replaced but without a null terminator before %d bytes", defaultReleaseImageLength)
	}
	if nullTerminator > len(defaultReleaseImagePadded) {
		return "", fmt.Errorf("release image location contains no null-terminator and constant is corrupted")
	}
	pullspec := defaultReleaseImagePadded[:nullTerminator]
	if len(pullspec) == 0 {
		return "", fmt.Errorf("release image location is empty, this binary was incorrectly generated")
	}
	return pullspec, nil
}
