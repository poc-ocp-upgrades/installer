package azure

import (
	"context"
	"fmt"
	"time"
	azdns "github.com/Azure/azure-sdk-for-go/profiles/latest/dns/mgmt/dns"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

type DNSConfig struct{ Session *Session }
type ZonesGetter interface {
	GetAllPublicZones() (map[string]string, error)
}
type ZonesClient struct{ azureClient azdns.ZonesClient }
type Zone struct {
	ID	string
	Name	string
}

func (z Zone) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s", z.Name)
}
func transformZone(f func(s string) *Zone) survey.Transformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(ans interface{}) interface{} {
		if "" == ans.(string) {
			return nil
		}
		s, ok := ans.(string)
		if !ok {
			return nil
		}
		return f(s)
	}
}
func (config DNSConfig) GetDNSZoneID(rgName string, zoneName string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/dnszones/%s", config.Session.Credentials.SubscriptionID, rgName, zoneName)
}
func (config DNSConfig) GetDNSZone() (*Zone, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	zonesClient := newZonesClient(config.Session)
	allZones, _ := zonesClient.GetAllPublicZones()
	if len(allZones) == 0 {
		return nil, errors.New("no public dns zone found in your subscription")
	}
	zoneNames := []string{}
	for zoneName := range allZones {
		zoneNames = append(zoneNames, zoneName)
	}
	var zoneName string
	err := survey.Ask([]*survey.Question{{Prompt: &survey.Select{Message: "Base Domain", Help: "The base domain of the cluster. All DNS records will be sub-domains of this base and will also include the cluster name.\n\nIf you don't see you intended base-domain listed, create a new Azure DNS Zone and rerun the installer.", Options: zoneNames}}}, &zoneName)
	if err != nil {
		return nil, err
	}
	return &Zone{ID: allZones[zoneName], Name: zoneName}, nil
}
func NewDNSConfig() (*DNSConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	session, err := GetSession()
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve session information")
	}
	return &DNSConfig{Session: session}, nil
}
func newZonesClient(session *Session) ZonesGetter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	azureClient := azdns.NewZonesClient(session.Credentials.SubscriptionID)
	azureClient.Authorizer = session.Authorizer
	return &ZonesClient{azureClient: azureClient}
}
func (client *ZonesClient) GetAllPublicZones() (map[string]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	allZones := map[string]string{}
	for zonesPage, err := client.azureClient.List(ctx, to.Int32Ptr(100)); zonesPage.NotDone(); err = zonesPage.NextWithContext(ctx) {
		if err != nil {
			return nil, err
		}
		for _, zone := range zonesPage.Values() {
			allZones[to.String(zone.Name)] = to.String(zone.ID)
		}
	}
	return allZones, nil
}
