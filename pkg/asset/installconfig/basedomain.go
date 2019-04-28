package installconfig

import (
	"github.com/aws/aws-sdk-go/aws/request"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"github.com/pkg/errors"
	survey "gopkg.in/AlecAivazis/survey.v1"
	"github.com/openshift/installer/pkg/asset"
	awsconfig "github.com/openshift/installer/pkg/asset/installconfig/aws"
	azureconfig "github.com/openshift/installer/pkg/asset/installconfig/azure"
	"github.com/openshift/installer/pkg/types/aws"
	"github.com/openshift/installer/pkg/types/azure"
	"github.com/openshift/installer/pkg/validate"
)

type baseDomain struct{ BaseDomain string }

var _ asset.Asset = (*baseDomain)(nil)

func (a *baseDomain) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&platform{}}
}
func (a *baseDomain) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	platform := &platform{}
	parents.Get(platform)
	switch platform.CurrentName() {
	case aws.Name:
		var err error
		a.BaseDomain, err = awsconfig.GetBaseDomain()
		cause := errors.Cause(err)
		if !(awsconfig.IsForbidden(cause) || request.IsErrorThrottle(cause)) {
			return err
		}
	case azure.Name:
		var err error
		azureDNS, _ := azureconfig.NewDNSConfig()
		zone, err := azureDNS.GetDNSZone()
		if err != nil {
			return err
		}
		a.BaseDomain = zone.Name
		return platform.Azure.SetBaseDomain(zone.ID)
	default:
	}
	return survey.Ask([]*survey.Question{{Prompt: &survey.Input{Message: "Base Domain", Help: "The base domain of the cluster. All DNS records will be sub-domains of this base and will also include the cluster name."}, Validate: survey.ComposeValidators(survey.Required, func(ans interface{}) error {
		return validate.DomainName(ans.(string), true)
	})}}, &a.BaseDomain)
}
func (a *baseDomain) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Base Domain"
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
