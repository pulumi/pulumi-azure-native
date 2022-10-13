


package app

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContainerApp(ctx *pulumi.Context, args *LookupContainerAppArgs, opts ...pulumi.InvokeOption) (*LookupContainerAppResult, error) {
	var rv LookupContainerAppResult
	err := ctx.Invoke("azure-native:app:getContainerApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupContainerAppArgs struct {
	ContainerAppName  string `pulumi:"containerAppName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupContainerAppResult struct {
	Configuration              *ConfigurationResponse          `pulumi:"configuration"`
	CustomDomainVerificationId string                          `pulumi:"customDomainVerificationId"`
	Id                         string                          `pulumi:"id"`
	Identity                   *ManagedServiceIdentityResponse `pulumi:"identity"`
	LatestRevisionFqdn         string                          `pulumi:"latestRevisionFqdn"`
	LatestRevisionName         string                          `pulumi:"latestRevisionName"`
	Location                   string                          `pulumi:"location"`
	ManagedEnvironmentId       *string                         `pulumi:"managedEnvironmentId"`
	Name                       string                          `pulumi:"name"`
	OutboundIpAddresses        []string                        `pulumi:"outboundIpAddresses"`
	ProvisioningState          string                          `pulumi:"provisioningState"`
	SystemData                 SystemDataResponse              `pulumi:"systemData"`
	Tags                       map[string]string               `pulumi:"tags"`
	Template                   *TemplateResponse               `pulumi:"template"`
	Type                       string                          `pulumi:"type"`
}


func (val *LookupContainerAppResult) Defaults() *LookupContainerAppResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Configuration = tmp.Configuration.Defaults()

	return &tmp
}

func LookupContainerAppOutput(ctx *pulumi.Context, args LookupContainerAppOutputArgs, opts ...pulumi.InvokeOption) LookupContainerAppResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerAppResult, error) {
			args := v.(LookupContainerAppArgs)
			r, err := LookupContainerApp(ctx, &args, opts...)
			var s LookupContainerAppResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContainerAppResultOutput)
}

type LookupContainerAppOutputArgs struct {
	ContainerAppName  pulumi.StringInput `pulumi:"containerAppName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupContainerAppOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerAppArgs)(nil)).Elem()
}


type LookupContainerAppResultOutput struct{ *pulumi.OutputState }

func (LookupContainerAppResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerAppResult)(nil)).Elem()
}

func (o LookupContainerAppResultOutput) ToLookupContainerAppResultOutput() LookupContainerAppResultOutput {
	return o
}

func (o LookupContainerAppResultOutput) ToLookupContainerAppResultOutputWithContext(ctx context.Context) LookupContainerAppResultOutput {
	return o
}

func (o LookupContainerAppResultOutput) Configuration() ConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppResult) *ConfigurationResponse { return v.Configuration }).(ConfigurationResponsePtrOutput)
}

func (o LookupContainerAppResultOutput) CustomDomainVerificationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppResult) string { return v.CustomDomainVerificationId }).(pulumi.StringOutput)
}

func (o LookupContainerAppResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupContainerAppResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupContainerAppResultOutput) LatestRevisionFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppResult) string { return v.LatestRevisionFqdn }).(pulumi.StringOutput)
}

func (o LookupContainerAppResultOutput) LatestRevisionName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppResult) string { return v.LatestRevisionName }).(pulumi.StringOutput)
}

func (o LookupContainerAppResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupContainerAppResultOutput) ManagedEnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerAppResult) *string { return v.ManagedEnvironmentId }).(pulumi.StringPtrOutput)
}

func (o LookupContainerAppResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupContainerAppResultOutput) OutboundIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupContainerAppResult) []string { return v.OutboundIpAddresses }).(pulumi.StringArrayOutput)
}

func (o LookupContainerAppResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupContainerAppResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupContainerAppResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupContainerAppResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContainerAppResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupContainerAppResultOutput) Template() TemplateResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppResult) *TemplateResponse { return v.Template }).(TemplateResponsePtrOutput)
}

func (o LookupContainerAppResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerAppResultOutput{})
}
