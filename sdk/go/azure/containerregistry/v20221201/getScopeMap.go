


package v20221201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScopeMap(ctx *pulumi.Context, args *LookupScopeMapArgs, opts ...pulumi.InvokeOption) (*LookupScopeMapResult, error) {
	var rv LookupScopeMapResult
	err := ctx.Invoke("azure-native:containerregistry/v20221201:getScopeMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScopeMapArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ScopeMapName      string `pulumi:"scopeMapName"`
}


type LookupScopeMapResult struct {
	Actions           []string           `pulumi:"actions"`
	CreationDate      string             `pulumi:"creationDate"`
	Description       *string            `pulumi:"description"`
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}

func LookupScopeMapOutput(ctx *pulumi.Context, args LookupScopeMapOutputArgs, opts ...pulumi.InvokeOption) LookupScopeMapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScopeMapResult, error) {
			args := v.(LookupScopeMapArgs)
			r, err := LookupScopeMap(ctx, &args, opts...)
			var s LookupScopeMapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScopeMapResultOutput)
}

type LookupScopeMapOutputArgs struct {
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ScopeMapName      pulumi.StringInput `pulumi:"scopeMapName"`
}

func (LookupScopeMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScopeMapArgs)(nil)).Elem()
}


type LookupScopeMapResultOutput struct{ *pulumi.OutputState }

func (LookupScopeMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScopeMapResult)(nil)).Elem()
}

func (o LookupScopeMapResultOutput) ToLookupScopeMapResultOutput() LookupScopeMapResultOutput {
	return o
}

func (o LookupScopeMapResultOutput) ToLookupScopeMapResultOutputWithContext(ctx context.Context) LookupScopeMapResultOutput {
	return o
}

func (o LookupScopeMapResultOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupScopeMapResult) []string { return v.Actions }).(pulumi.StringArrayOutput)
}

func (o LookupScopeMapResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeMapResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupScopeMapResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeMapResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupScopeMapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeMapResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScopeMapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeMapResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScopeMapResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeMapResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupScopeMapResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupScopeMapResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupScopeMapResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeMapResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScopeMapResultOutput{})
}
