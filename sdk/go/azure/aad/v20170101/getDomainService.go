


package v20170101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDomainService(ctx *pulumi.Context, args *LookupDomainServiceArgs, opts ...pulumi.InvokeOption) (*LookupDomainServiceResult, error) {
	var rv LookupDomainServiceResult
	err := ctx.Invoke("azure-native:aad/v20170101:getDomainService", args, &rv, opts...)
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
	Name                      string                          `pulumi:"name"`
	NotificationSettings      *NotificationSettingsResponse   `pulumi:"notificationSettings"`
	ProvisioningState         string                          `pulumi:"provisioningState"`
	ServiceStatus             string                          `pulumi:"serviceStatus"`
	SubnetId                  *string                         `pulumi:"subnetId"`
	Tags                      map[string]string               `pulumi:"tags"`
	TenantId                  string                          `pulumi:"tenantId"`
	Type                      string                          `pulumi:"type"`
	VnetSiteId                string                          `pulumi:"vnetSiteId"`
}

func LookupDomainServiceOutput(ctx *pulumi.Context, args LookupDomainServiceOutputArgs, opts ...pulumi.InvokeOption) LookupDomainServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainServiceResult, error) {
			args := v.(LookupDomainServiceArgs)
			r, err := LookupDomainService(ctx, &args, opts...)
			var s LookupDomainServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainServiceResultOutput)
}

type LookupDomainServiceOutputArgs struct {
	DomainServiceName pulumi.StringInput `pulumi:"domainServiceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDomainServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainServiceArgs)(nil)).Elem()
}


type LookupDomainServiceResultOutput struct{ *pulumi.OutputState }

func (LookupDomainServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainServiceResult)(nil)).Elem()
}

func (o LookupDomainServiceResultOutput) ToLookupDomainServiceResultOutput() LookupDomainServiceResultOutput {
	return o
}

func (o LookupDomainServiceResultOutput) ToLookupDomainServiceResultOutputWithContext(ctx context.Context) LookupDomainServiceResultOutput {
	return o
}

func (o LookupDomainServiceResultOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.DeploymentId }).(pulumi.StringOutput)
}

func (o LookupDomainServiceResultOutput) DomainControllerIpAddress() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) []string { return v.DomainControllerIpAddress }).(pulumi.StringArrayOutput)
}

func (o LookupDomainServiceResultOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *string { return v.DomainName }).(pulumi.StringPtrOutput)
}

func (o LookupDomainServiceResultOutput) DomainSecuritySettings() DomainSecuritySettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *DomainSecuritySettingsResponse { return v.DomainSecuritySettings }).(DomainSecuritySettingsResponsePtrOutput)
}

func (o LookupDomainServiceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupDomainServiceResultOutput) FilteredSync() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *string { return v.FilteredSync }).(pulumi.StringPtrOutput)
}

func (o LookupDomainServiceResultOutput) HealthAlerts() HealthAlertResponseArrayOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) []HealthAlertResponse { return v.HealthAlerts }).(HealthAlertResponseArrayOutput)
}

func (o LookupDomainServiceResultOutput) HealthLastEvaluated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.HealthLastEvaluated }).(pulumi.StringOutput)
}

func (o LookupDomainServiceResultOutput) HealthMonitors() HealthMonitorResponseArrayOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) []HealthMonitorResponse { return v.HealthMonitors }).(HealthMonitorResponseArrayOutput)
}

func (o LookupDomainServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDomainServiceResultOutput) LdapsSettings() LdapsSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *LdapsSettingsResponse { return v.LdapsSettings }).(LdapsSettingsResponsePtrOutput)
}

func (o LookupDomainServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDomainServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDomainServiceResultOutput) NotificationSettings() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *NotificationSettingsResponse { return v.NotificationSettings }).(NotificationSettingsResponsePtrOutput)
}

func (o LookupDomainServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDomainServiceResultOutput) ServiceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.ServiceStatus }).(pulumi.StringOutput)
}

func (o LookupDomainServiceResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o LookupDomainServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDomainServiceResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupDomainServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDomainServiceResultOutput) VnetSiteId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainServiceResult) string { return v.VnetSiteId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainServiceResultOutput{})
}
