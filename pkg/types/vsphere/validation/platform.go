package validation

import (
	"k8s.io/apimachinery/pkg/util/validation/field"
	"github.com/openshift/installer/pkg/types/vsphere"
)

func ValidatePlatform(p *vsphere.Platform, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	if len(p.VCenter) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("vCenter"), "must specify the name of the vCenter"))
	}
	if len(p.Username) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("username"), "must specify the username"))
	}
	if len(p.Password) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("password"), "must specify the password"))
	}
	if len(p.Datacenter) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("datacenter"), "must specify the datacenter"))
	}
	if len(p.DefaultDatastore) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("defaultDatastore"), "must specify the default datastore"))
	}
	return allErrs
}
