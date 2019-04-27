package validation

import (
	"sort"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"github.com/openshift/installer/pkg/types/aws"
)

var (
	Regions			= map[string]string{"ap-northeast-1": "Tokyo", "ap-northeast-2": "Seoul", "ap-south-1": "Mumbai", "ap-southeast-1": "Singapore", "ap-southeast-2": "Sydney", "ca-central-1": "Central", "eu-central-1": "Frankfurt", "eu-west-1": "Ireland", "eu-west-2": "London", "eu-west-3": "Paris", "sa-east-1": "São Paulo", "us-east-1": "N. Virginia", "us-east-2": "Ohio", "us-west-1": "N. California", "us-west-2": "Oregon"}
	validRegionValues	= func() []string {
		validValues := make([]string, len(Regions))
		i := 0
		for r := range Regions {
			validValues[i] = r
			i++
		}
		sort.Strings(validValues)
		return validValues
	}()
)

func ValidatePlatform(p *aws.Platform, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	if _, ok := Regions[p.Region]; !ok {
		allErrs = append(allErrs, field.NotSupported(fldPath.Child("region"), p.Region, validRegionValues))
	}
	if p.DefaultMachinePlatform != nil {
		allErrs = append(allErrs, ValidateMachinePool(p, p.DefaultMachinePlatform, fldPath.Child("defaultMachinePlatform"))...)
	}
	return allErrs
}
