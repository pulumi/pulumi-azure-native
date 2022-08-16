


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecurityUserConfiguration(ctx *pulumi.Context, args *LookupSecurityUserConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupSecurityUserConfigurationResult, error) {
	var rv LookupSecurityUserConfigurationResult
	err := ctx.Invoke("azure-native:network:getSecurityUserConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityUserConfigurationArgs struct {
	ConfigurationName  string `pulumi:"configurationName"`
	NetworkManagerName string `pulumi:"networkManagerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupSecurityUserConfigurationResult struct {
	DeleteExistingNSGs *string            `pulumi:"deleteExistingNSGs"`
	Description        *string            `pulumi:"description"`
	DisplayName        *string            `pulumi:"displayName"`
	Etag               string             `pulumi:"etag"`
	Id                 string             `pulumi:"id"`
	Name               string             `pulumi:"name"`
	ProvisioningState  string             `pulumi:"provisioningState"`
	SecurityType       *string            `pulumi:"securityType"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	Type               string             `pulumi:"type"`
}

func LookupSecurityUserConfigurationOutput(ctx *pulumi.Context, args LookupSecurityUserConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityUserConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityUserConfigurationResult, error) {
			args := v.(LookupSecurityUserConfigurationArgs)
			r, err := LookupSecurityUserConfiguration(ctx, &args, opts...)
			var s LookupSecurityUserConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityUserConfigurationResultOutput)
}

type LookupSecurityUserConfigurationOutputArgs struct {
	ConfigurationName  pulumi.StringInput `pulumi:"configurationName"`
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupSecurityUserConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityUserConfigurationArgs)(nil)).Elem()
}


type LookupSecurityUserConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityUserConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityUserConfigurationResult)(nil)).Elem()
}

func (o LookupSecurityUserConfigurationResultOutput) ToLookupSecurityUserConfigurationResultOutput() LookupSecurityUserConfigurationResultOutput {
	return o
}

func (o LookupSecurityUserConfigurationResultOutput) ToLookupSecurityUserConfigurationResultOutputWithContext(ctx context.Context) LookupSecurityUserConfigurationResultOutput {
	return o
}

func (o LookupSecurityUserConfigurationResultOutput) DeleteExistingNSGs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityUserConfigurationResult) *string { return v.DeleteExistingNSGs }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityUserConfigurationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityUserConfigurationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityUserConfigurationResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityUserConfigurationResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityUserConfigurationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityUserConfigurationResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupSecurityUserConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityUserConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSecurityUserConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityUserConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSecurityUserConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityUserConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSecurityUserConfigurationResultOutput) SecurityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityUserConfigurationResult) *string { return v.SecurityType }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityUserConfigurationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSecurityUserConfigurationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSecurityUserConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityUserConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityUserConfigurationResultOutput{})
}
