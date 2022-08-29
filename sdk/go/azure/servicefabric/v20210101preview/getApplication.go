


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("azure-native:servicefabric/v20210101preview:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	ApplicationName   string `pulumi:"applicationName"`
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupApplicationResult struct {
	Id                string                                    `pulumi:"id"`
	Identity          *ManagedIdentityResponse                  `pulumi:"identity"`
	Location          *string                                   `pulumi:"location"`
	ManagedIdentities []ApplicationUserAssignedIdentityResponse `pulumi:"managedIdentities"`
	Name              string                                    `pulumi:"name"`
	Parameters        map[string]string                         `pulumi:"parameters"`
	ProvisioningState string                                    `pulumi:"provisioningState"`
	SystemData        SystemDataResponse                        `pulumi:"systemData"`
	Tags              map[string]string                         `pulumi:"tags"`
	Type              string                                    `pulumi:"type"`
	UpgradePolicy     *ApplicationUpgradePolicyResponse         `pulumi:"upgradePolicy"`
	Version           *string                                   `pulumi:"version"`
}

func LookupApplicationOutput(ctx *pulumi.Context, args LookupApplicationOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationResult, error) {
			args := v.(LookupApplicationArgs)
			r, err := LookupApplication(ctx, &args, opts...)
			var s LookupApplicationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationResultOutput)
}

type LookupApplicationOutputArgs struct {
	ApplicationName   pulumi.StringInput `pulumi:"applicationName"`
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationArgs)(nil)).Elem()
}


type LookupApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationResult)(nil)).Elem()
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutput() LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutputWithContext(ctx context.Context) LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Identity() ManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *ManagedIdentityResponse { return v.Identity }).(ManagedIdentityResponsePtrOutput)
}

func (o LookupApplicationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) ManagedIdentities() ApplicationUserAssignedIdentityResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []ApplicationUserAssignedIdentityResponse { return v.ManagedIdentities }).(ApplicationUserAssignedIdentityResponseArrayOutput)
}

func (o LookupApplicationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationResult) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o LookupApplicationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApplicationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupApplicationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupApplicationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) UpgradePolicy() ApplicationUpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *ApplicationUpgradePolicyResponse { return v.UpgradePolicy }).(ApplicationUpgradePolicyResponsePtrOutput)
}

func (o LookupApplicationResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationResultOutput{})
}
