


package v20171001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegistry(ctx *pulumi.Context, args *LookupRegistryArgs, opts ...pulumi.InvokeOption) (*LookupRegistryResult, error) {
	var rv LookupRegistryResult
	err := ctx.Invoke("azure-native:containerregistry/v20171001:getRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRegistryResult struct {
	AdminUserEnabled  *bool                             `pulumi:"adminUserEnabled"`
	CreationDate      string                            `pulumi:"creationDate"`
	Id                string                            `pulumi:"id"`
	Location          string                            `pulumi:"location"`
	LoginServer       string                            `pulumi:"loginServer"`
	Name              string                            `pulumi:"name"`
	NetworkRuleSet    *NetworkRuleSetResponse           `pulumi:"networkRuleSet"`
	ProvisioningState string                            `pulumi:"provisioningState"`
	Sku               SkuResponse                       `pulumi:"sku"`
	Status            StatusResponse                    `pulumi:"status"`
	StorageAccount    *StorageAccountPropertiesResponse `pulumi:"storageAccount"`
	Tags              map[string]string                 `pulumi:"tags"`
	Type              string                            `pulumi:"type"`
}


func (val *LookupRegistryResult) Defaults() *LookupRegistryResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AdminUserEnabled) {
		adminUserEnabled_ := false
		tmp.AdminUserEnabled = &adminUserEnabled_
	}
	tmp.NetworkRuleSet = tmp.NetworkRuleSet.Defaults()

	return &tmp
}

func LookupRegistryOutput(ctx *pulumi.Context, args LookupRegistryOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryResult, error) {
			args := v.(LookupRegistryArgs)
			r, err := LookupRegistry(ctx, &args, opts...)
			var s LookupRegistryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryResultOutput)
}

type LookupRegistryOutputArgs struct {
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRegistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryArgs)(nil)).Elem()
}


type LookupRegistryResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryResult)(nil)).Elem()
}

func (o LookupRegistryResultOutput) ToLookupRegistryResultOutput() LookupRegistryResultOutput {
	return o
}

func (o LookupRegistryResultOutput) ToLookupRegistryResultOutputWithContext(ctx context.Context) LookupRegistryResultOutput {
	return o
}

func (o LookupRegistryResultOutput) AdminUserEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *bool { return v.AdminUserEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupRegistryResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) LoginServer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.LoginServer }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) NetworkRuleSet() NetworkRuleSetResponsePtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *NetworkRuleSetResponse { return v.NetworkRuleSet }).(NetworkRuleSetResponsePtrOutput)
}

func (o LookupRegistryResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupRegistryResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupRegistryResultOutput) Status() StatusResponseOutput {
	return o.ApplyT(func(v LookupRegistryResult) StatusResponse { return v.Status }).(StatusResponseOutput)
}

func (o LookupRegistryResultOutput) StorageAccount() StorageAccountPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *StorageAccountPropertiesResponse { return v.StorageAccount }).(StorageAccountPropertiesResponsePtrOutput)
}

func (o LookupRegistryResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRegistryResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupRegistryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryResultOutput{})
}
