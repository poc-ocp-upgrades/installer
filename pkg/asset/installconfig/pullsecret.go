package installconfig

import (
	survey "gopkg.in/AlecAivazis/survey.v1"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/validate"
)

type pullSecret struct{ PullSecret string }

var _ asset.Asset = (*pullSecret)(nil)

func (a *pullSecret) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (a *pullSecret) Generate(asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return survey.Ask([]*survey.Question{{Prompt: &survey.Password{Message: "Pull Secret", Help: "The container registry pull secret for this cluster, as a single line of JSON (e.g. {\"auths\": {...}}).\n\nYou can get this secret from https://cloud.openshift.com/clusters/install#pull-secret"}, Validate: survey.ComposeValidators(survey.Required, func(ans interface{}) error {
		return validate.ImagePullSecret(ans.(string))
	})}}, &a.PullSecret)
}
func (a *pullSecret) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Pull Secret"
}
