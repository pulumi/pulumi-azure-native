


package v20170601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDomainService(ctx *pulumi.Context, args *LookupDomainServiceArgs, opts ...pulumi.InvokeOption) (*LookupDomainServiceResult, error) {
	var rv LookupDomainServiceResult
	err := ctx.Invoke("azure-native:aad/v20170601:getDomainService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDomainServiceArgs struct {
	DomainServiceName string `pulumi:"domainServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDomainServiceResult struct {
	DeploymentId              string                          `pulumi:"deploymentId"`
	DomainConfigurationType   *string                         `pulumi:"domainConfigurationType"`
	DomainControllerIpAddress []string                        `pulumi:"domainControllerIpAddress"`
	DomainName                *string                         `pulumi:"domainName"`
	DomainSecuritySettings    *DomainSecuritySettingsResponse `pulumi:"domainSecuritySettings"`
	Etag                      *string                         `pulumi:"etag"`
	FilteredSync              *string                         `pulumi:"filteredSync"`
	HealthAlerts              []HealthAlertResponse           `pulumi:"healthAlerts"`
	HealthLastEvaluated       string                          `pulumi:"healthLastEvaluated"`
	HealthMonitors            []HealthMonitorResponse         `pulumi:"healthMonitors"`
	Id                        string                          `pulumi:"id"`
	LdapsSettings             *LdapsSettingsResponse          `pulumi:"ldapsSettings"`
	Location                  *string                         `pulumi:"location"`
	MigrationProperties       MigrationPropertiesResponse     `pulumi:"migrationProperties"`
	Name                      string                          `pulumi:"name"`
	NotificationSettings      *NotificationSettingsResponse   `pulumi:"notificationSettings"`
	ProvisioningState         string                          `pulumi:"provisioningState"`
	ResourceForestSettings    *ResourceForestSettingsResponse `pulumi:"resourceForestSettings"`
	ServiceStatus             string                          `pulumi:"serviceStatus"`
	Sku                       *string                         `pulumi:"sku"`
	SubnetId                  *string                         `pulumi:"subnetId"`
	Tags                      map[string]string               `pulumi:"tags"`
	TenantId                  string                          `pulumi:"tenantId"`
	Type                      string                          `pulumi:"type"`
	Version                   int                             `pulumi:"version"`
	VnetSiteId                string                          `pulumi:"vnetSiteId"`
}
