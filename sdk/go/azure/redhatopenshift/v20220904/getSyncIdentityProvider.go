


package v20220904

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSyncIdentityProvider(ctx *pulumi.Context, args *LookupSyncIdentityProviderArgs, opts ...pulumi.InvokeOption) (*LookupSyncIdentityProviderResult, error) {
	var rv LookupSyncIdentityProviderResult
	err := ctx.Invoke("azure-native:redhatopenshift/v20220904:getSyncIdentityProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSyncIdentityProviderArgs struct {
	ChildResourceName string `pulumi:"childResourceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupSyncIdentityProviderResult struct {
	Id         string             `pulumi:"id"`
	Name       string             `pulumi:"name"`
	Resources  *string            `pulumi:"resources"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupSyncIdentityProviderOutput(ctx *pulumi.Context, args LookupSyncIdentityProviderOutputArgs, opts ...pulumi.InvokeOption) LookupSyncIdentityProviderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSyncIdentityProviderResult, error) {
			args := v.(LookupSyncIdentityProviderArgs)
			r, err := LookupSyncIdentityProvider(ctx, &args, opts...)
			var s LookupSyncIdentityProviderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSyncIdentityProviderResultOutput)
}

type LookupSyncIdentityProviderOutputArgs struct {
	ChildResourceName pulumi.StringInput `pulumi:"childResourceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupSyncIdentityProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSyncIdentityProviderArgs)(nil)).Elem()
}


type LookupSyncIdentityProviderResultOutput struct{ *pulumi.OutputState }

func (LookupSyncIdentityProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSyncIdentityProviderResult)(nil)).Elem()
}

func (o LookupSyncIdentityProviderResultOutput) ToLookupSyncIdentityProviderResultOutput() LookupSyncIdentityProviderResultOutput {
	return o
}

func (o LookupSyncIdentityProviderResultOutput) ToLookupSyncIdentityProviderResultOutputWithContext(ctx context.Context) LookupSyncIdentityProviderResultOutput {
	return o
}

func (o LookupSyncIdentityProviderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncIdentityProviderResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSyncIdentityProviderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncIdentityProviderResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSyncIdentityProviderResultOutput) Resources() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncIdentityProviderResult) *string { return v.Resources }).(pulumi.StringPtrOutput)
}

func (o LookupSyncIdentityProviderResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSyncIdentityProviderResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSyncIdentityProviderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncIdentityProviderResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSyncIdentityProviderResultOutput{})
}
