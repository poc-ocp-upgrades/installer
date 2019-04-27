package tls

import (
	"fmt"
	"testing"
	"github.com/openshift/installer/pkg/asset"
	"github.com/stretchr/testify/assert"
)

type mockCertKey struct {
	cert	string
	key	string
}

func (mck *mockCertKey) Cert() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []byte(mck.cert)
}
func (mck *mockCertKey) Key() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []byte(mck.key)
}
func Test_CertBundleGenerate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		input		[]string
		expBundle	string
		expErr		string
	}{{expErr: "atleast one certificate required for a bundle"}, {input: []string{`-----BEGIN CERTIFICATE-----
MIICYTCCAcqgAwIBAgIJAI2kA+uXAbhOMA0GCSqGSIb3DQEBCwUAMEgxCzAJBgNV
BAYTAlVTMQswCQYDVQQIDAJDQTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzEUMBIG
A1UECgwLUmVkIEhhdCBJbmMwHhcNMTkwMjEyMTkzMjUzWhcNMTkwMjEzMTkzMjUz
WjBIMQswCQYDVQQGEwJVUzELMAkGA1UECAwCQ0ExFjAUBgNVBAcMDVNhbiBGcmFu
Y2lzY28xFDASBgNVBAoMC1JlZCBIYXQgSW5jMIGfMA0GCSqGSIb3DQEBAQUAA4GN
ADCBiQKBgQC+HOC0mKig/oINAKPo88LqxDJ4l7lozdLtp5oGeqWrLUXSfkvXAkQY
2QYdvPAjpRfH7Ii7G0Asx+HTKdvula7B5fXDjc6NYKuEpTJZRV1ugntI97bozF/E
C2BBmxxEnJN3+Xe8RYXMjz5Q4aqPw9vZhlWN+0hrREl1Ea/zHuWFIQIDAQABo1Mw
UTAdBgNVHQ4EFgQUvTS1XjlvOdsufSyWxukyQu3LriEwHwYDVR0jBBgwFoAUvTS1
XjlvOdsufSyWxukyQu3LriEwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsF
AAOBgQB9gFcOXnzJrM65QqxeCB9Z5l5JMjp45UFC9Bj2cgwDHP80Zvi4omlaacC6
aavmnLd67zm9PbYDWRaOIWAMeB916Iwaw/v6I0jwhAk/VxX5Fl6cGlZu9jZ3zbFE
2sDqkwzIuSjCG2A23s6d4M1S3IXCCydoCSLMu+WhLkbboK6jEg==
-----END CERTIFICATE-----`}, expBundle: `-----BEGIN CERTIFICATE-----
MIICYTCCAcqgAwIBAgIJAI2kA+uXAbhOMA0GCSqGSIb3DQEBCwUAMEgxCzAJBgNV
BAYTAlVTMQswCQYDVQQIDAJDQTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzEUMBIG
A1UECgwLUmVkIEhhdCBJbmMwHhcNMTkwMjEyMTkzMjUzWhcNMTkwMjEzMTkzMjUz
WjBIMQswCQYDVQQGEwJVUzELMAkGA1UECAwCQ0ExFjAUBgNVBAcMDVNhbiBGcmFu
Y2lzY28xFDASBgNVBAoMC1JlZCBIYXQgSW5jMIGfMA0GCSqGSIb3DQEBAQUAA4GN
ADCBiQKBgQC+HOC0mKig/oINAKPo88LqxDJ4l7lozdLtp5oGeqWrLUXSfkvXAkQY
2QYdvPAjpRfH7Ii7G0Asx+HTKdvula7B5fXDjc6NYKuEpTJZRV1ugntI97bozF/E
C2BBmxxEnJN3+Xe8RYXMjz5Q4aqPw9vZhlWN+0hrREl1Ea/zHuWFIQIDAQABo1Mw
UTAdBgNVHQ4EFgQUvTS1XjlvOdsufSyWxukyQu3LriEwHwYDVR0jBBgwFoAUvTS1
XjlvOdsufSyWxukyQu3LriEwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsF
AAOBgQB9gFcOXnzJrM65QqxeCB9Z5l5JMjp45UFC9Bj2cgwDHP80Zvi4omlaacC6
aavmnLd67zm9PbYDWRaOIWAMeB916Iwaw/v6I0jwhAk/VxX5Fl6cGlZu9jZ3zbFE
2sDqkwzIuSjCG2A23s6d4M1S3IXCCydoCSLMu+WhLkbboK6jEg==
-----END CERTIFICATE-----
`}, {input: []string{`-----BEGIN CERTIFICATE-----
MIICYTCCAcqgAwIBAgIJAI2kA+uXAbhOMA0GCSqGSIb3DQEBCwUAMEgxCzAJBgNV
BAYTAlVTMQswCQYDVQQIDAJDQTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzEUMBIG
A1UECgwLUmVkIEhhdCBJbmMwHhcNMTkwMjEyMTkzMjUzWhcNMTkwMjEzMTkzMjUz
WjBIMQswCQYDVQQGEwJVUzELMAkGA1UECAwCQ0ExFjAUBgNVBAcMDVNhbiBGcmFu
Y2lzY28xFDASBgNVBAoMC1JlZCBIYXQgSW5jMIGfMA0GCSqGSIb3DQEBAQUAA4GN
ADCBiQKBgQC+HOC0mKig/oINAKPo88LqxDJ4l7lozdLtp5oGeqWrLUXSfkvXAkQY
2QYdvPAjpRfH7Ii7G0Asx+HTKdvula7B5fXDjc6NYKuEpTJZRV1ugntI97bozF/E
C2BBmxxEnJN3+Xe8RYXMjz5Q4aqPw9vZhlWN+0hrREl1Ea/zHuWFIQIDAQABo1Mw
UTAdBgNVHQ4EFgQUvTS1XjlvOdsufSyWxukyQu3LriEwHwYDVR0jBBgwFoAUvTS1
XjlvOdsufSyWxukyQu3LriEwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsF
AAOBgQB9gFcOXnzJrM65QqxeCB9Z5l5JMjp45UFC9Bj2cgwDHP80Zvi4omlaacC6
aavmnLd67zm9PbYDWRaOIWAMeB916Iwaw/v6I0jwhAk/VxX5Fl6cGlZu9jZ3zbFE
2sDqkwzIuSjCG2A23s6d4M1S3IXCCydoCSLMu+WhLkbboK6jEg==
-----END CERTIFICATE-----`, `-----BEGIN CERTIFICATE-----
MIICVTCCAb6gAwIBAgIJAKDi1rywgIWHMA0GCSqGSIb3DQEBCwUAMEIxCzAJBgNV
BAYTAlVTMRUwEwYDVQQHDAxEZWZhdWx0IENpdHkxHDAaBgNVBAoME0RlZmF1bHQg
Q29tcGFueSBMdGQwHhcNMTkwMjEyMTkzNTA5WhcNMTkwMjEzMTkzNTA5WjBCMQsw
CQYDVQQGEwJVUzEVMBMGA1UEBwwMRGVmYXVsdCBDaXR5MRwwGgYDVQQKDBNEZWZh
dWx0IENvbXBhbnkgTHRkMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC0IwdO
//CIUspBaLVDVPlcvbE4koN5v4dx/G0ClcfDDDcUFhw3yu/MiAt/EwDDzucHoM+l
mVrNZV3A4+QjswHxwl7jKAeZ/bpTsDh7j/r5zpoY8hD5za9+BZ5Pr0k5q3ONS6Iy
Uu8VHCPjIzMETcEJ+LCB0iD4xMAvb7daNmimVQIDAQABo1MwUTAdBgNVHQ4EFgQU
pKgTmnFxt741aiUOMy2bbcEP2AgwHwYDVR0jBBgwFoAUpKgTmnFxt741aiUOMy2b
bcEP2AgwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOBgQBW/U1DAbEX
0zk4FNxl4d/82ax44MVaZ5Owrhgr6kWBXDR2kRyYlq1yfLVHLjMqIkCe5VcBBfwy
s5Q0Xv1T6UKcWvIHwNNxo/dYDtnmjdllrEeVKKW0kmotCCsGLU/ZBa++Rl/GYpwv
CjH8bTNT3u6KYZKRVH0A2/EpJHC+TSSe3A==
-----END CERTIFICATE-----`}, expBundle: `-----BEGIN CERTIFICATE-----
MIICYTCCAcqgAwIBAgIJAI2kA+uXAbhOMA0GCSqGSIb3DQEBCwUAMEgxCzAJBgNV
BAYTAlVTMQswCQYDVQQIDAJDQTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzEUMBIG
A1UECgwLUmVkIEhhdCBJbmMwHhcNMTkwMjEyMTkzMjUzWhcNMTkwMjEzMTkzMjUz
WjBIMQswCQYDVQQGEwJVUzELMAkGA1UECAwCQ0ExFjAUBgNVBAcMDVNhbiBGcmFu
Y2lzY28xFDASBgNVBAoMC1JlZCBIYXQgSW5jMIGfMA0GCSqGSIb3DQEBAQUAA4GN
ADCBiQKBgQC+HOC0mKig/oINAKPo88LqxDJ4l7lozdLtp5oGeqWrLUXSfkvXAkQY
2QYdvPAjpRfH7Ii7G0Asx+HTKdvula7B5fXDjc6NYKuEpTJZRV1ugntI97bozF/E
C2BBmxxEnJN3+Xe8RYXMjz5Q4aqPw9vZhlWN+0hrREl1Ea/zHuWFIQIDAQABo1Mw
UTAdBgNVHQ4EFgQUvTS1XjlvOdsufSyWxukyQu3LriEwHwYDVR0jBBgwFoAUvTS1
XjlvOdsufSyWxukyQu3LriEwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsF
AAOBgQB9gFcOXnzJrM65QqxeCB9Z5l5JMjp45UFC9Bj2cgwDHP80Zvi4omlaacC6
aavmnLd67zm9PbYDWRaOIWAMeB916Iwaw/v6I0jwhAk/VxX5Fl6cGlZu9jZ3zbFE
2sDqkwzIuSjCG2A23s6d4M1S3IXCCydoCSLMu+WhLkbboK6jEg==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIICVTCCAb6gAwIBAgIJAKDi1rywgIWHMA0GCSqGSIb3DQEBCwUAMEIxCzAJBgNV
BAYTAlVTMRUwEwYDVQQHDAxEZWZhdWx0IENpdHkxHDAaBgNVBAoME0RlZmF1bHQg
Q29tcGFueSBMdGQwHhcNMTkwMjEyMTkzNTA5WhcNMTkwMjEzMTkzNTA5WjBCMQsw
CQYDVQQGEwJVUzEVMBMGA1UEBwwMRGVmYXVsdCBDaXR5MRwwGgYDVQQKDBNEZWZh
dWx0IENvbXBhbnkgTHRkMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC0IwdO
//CIUspBaLVDVPlcvbE4koN5v4dx/G0ClcfDDDcUFhw3yu/MiAt/EwDDzucHoM+l
mVrNZV3A4+QjswHxwl7jKAeZ/bpTsDh7j/r5zpoY8hD5za9+BZ5Pr0k5q3ONS6Iy
Uu8VHCPjIzMETcEJ+LCB0iD4xMAvb7daNmimVQIDAQABo1MwUTAdBgNVHQ4EFgQU
pKgTmnFxt741aiUOMy2bbcEP2AgwHwYDVR0jBBgwFoAUpKgTmnFxt741aiUOMy2b
bcEP2AgwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOBgQBW/U1DAbEX
0zk4FNxl4d/82ax44MVaZ5Owrhgr6kWBXDR2kRyYlq1yfLVHLjMqIkCe5VcBBfwy
s5Q0Xv1T6UKcWvIHwNNxo/dYDtnmjdllrEeVKKW0kmotCCsGLU/ZBa++Rl/GYpwv
CjH8bTNT3u6KYZKRVH0A2/EpJHC+TSSe3A==
-----END CERTIFICATE-----
`}}
	for idx, test := range tests {
		t.Run(fmt.Sprintf("#%d", idx), func(t *testing.T) {
			var certkeys []CertInterface
			for _, c := range test.input {
				certkeys = append(certkeys, &mockCertKey{cert: c})
			}
			bundle := CertBundle{}
			err := bundle.Generate("test-bundle", certkeys...)
			if test.expErr == "" {
				assert.NoError(t, err)
				assert.Equal(t, string(bundle.BundleRaw), test.expBundle)
				assert.Equal(t, bundle.FileList, []*asset.File{{Filename: "tls/test-bundle.crt", Data: []byte(test.expBundle)}})
			} else {
				assert.EqualError(t, err, test.expErr)
			}
		})
	}
}
