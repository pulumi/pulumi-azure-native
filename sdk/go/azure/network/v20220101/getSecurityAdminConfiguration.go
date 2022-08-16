


package v20220101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecurityAdminConfiguration(ctx *pulumi.Context, args *LookupSecurityAdminConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupSecurityAdminConfigurationResult, error) {
	var rv LookupSecurityAdminConfigurationResult
	err := ctx.Invoke("azure-native:network/v20220101:getSecurityAdminConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityAdminConfigurationArgs struct {
	ConfigurationName  string `pulumi:"configurationName"`
	NetworkManagerName string `pulumi:"networkManagerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupSecurityAdminConfigurationResult struct {
	ApplyOnNetworkIntentPolicyBasedServices []string           `pulumi:"applyOnNetworkIntentPolicyBasedServices"`
	Description                             *string            `pulumi:"description"`
	Etag                                    string             `pulumi:"etag"`
	Id                                      string             `pulumi:"id"`
	Name                                    string             `pulumi:"name"`
	ProvisioningState                       string             `pulumi:"provisioningState"`
	SystemData                              SystemDataResponse `pulumi:"systemData"`
	Type                                    string             `pulumi:"type"`
}

func LookupSecurityAdminConfigurationOutput(ctx *pulumi.Context, args LookupSecurityAdminConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityAdminConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityAdminConfigurationResult, error) {
			args := v.(LookupSecurityAdminConfigurationArgs)
			r, err := LookupSecurityAdminConfiguration(ctx, &args, opts...)
			var s LookupSecurityAdminConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityAdminConfigurationResultOutput)
}

type LookupSecurityAdminConfigurationOutputArgs struct {
	ConfigurationName  pulumi.StringInput `pulumi:"configurationName"`
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupSecurityAdminConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityAdminConfigurationArgs)(nil)).Elem()
}


type LookupSecurityAdminConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityAdminConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityAdminConfigurationResult)(nil)).Elem()
}

func (o LookupSecurityAdminConfigurationResultOutput) ToLookupSecurityAdminConfigurationResultOutput() LookupSecurityAdminConfigurationResultOutput {
	return o
}

func (o LookupSecurityAdminConfigurationResultOutput) ToLookupSecurityAdminConfigurationResultOutputWithContext(ctx context.Context) LookupSecurityAdminConfigurationResultOutput {
	return o
}

func (o LookupSecurityAdminConfigurationResultOutput) ApplyOnNetworkIntentPolicyBasedServices() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSecurityAdminConfigurationResult) []string {
		return v.ApplyOnNetworkIntentPolicyBasedServices
	}).(pulumi.StringArrayOutput)
}

func (o LookupSecurityAdminConfigurationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityAdminConfigurationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityAdminConfigurationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityAdminConfigurationResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupSecurityAdminConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityAdminConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSecurityAdminConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityAdminConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSecurityAdminConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityAdminConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSecurityAdminConfigurationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSecurityAdminConfigurationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSecurityAdminConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityAdminConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityAdminConfigurationResultOutput{})
}
