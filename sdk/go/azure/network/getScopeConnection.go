


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScopeConnection(ctx *pulumi.Context, args *LookupScopeConnectionArgs, opts ...pulumi.InvokeOption) (*LookupScopeConnectionResult, error) {
	var rv LookupScopeConnectionResult
	err := ctx.Invoke("azure-native:network:getScopeConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScopeConnectionArgs struct {
	NetworkManagerName  string `pulumi:"networkManagerName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	ScopeConnectionName string `pulumi:"scopeConnectionName"`
}


type LookupScopeConnectionResult struct {
	Description *string            `pulumi:"description"`
	Etag        string             `pulumi:"etag"`
	Id          string             `pulumi:"id"`
	Name        string             `pulumi:"name"`
	ResourceId  *string            `pulumi:"resourceId"`
	SystemData  SystemDataResponse `pulumi:"systemData"`
	TenantId    *string            `pulumi:"tenantId"`
	Type        string             `pulumi:"type"`
}

func LookupScopeConnectionOutput(ctx *pulumi.Context, args LookupScopeConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupScopeConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScopeConnectionResult, error) {
			args := v.(LookupScopeConnectionArgs)
			r, err := LookupScopeConnection(ctx, &args, opts...)
			var s LookupScopeConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScopeConnectionResultOutput)
}

type LookupScopeConnectionOutputArgs struct {
	NetworkManagerName  pulumi.StringInput `pulumi:"networkManagerName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	ScopeConnectionName pulumi.StringInput `pulumi:"scopeConnectionName"`
}

func (LookupScopeConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScopeConnectionArgs)(nil)).Elem()
}


type LookupScopeConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupScopeConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScopeConnectionResult)(nil)).Elem()
}

func (o LookupScopeConnectionResultOutput) ToLookupScopeConnectionResultOutput() LookupScopeConnectionResultOutput {
	return o
}

func (o LookupScopeConnectionResultOutput) ToLookupScopeConnectionResultOutputWithContext(ctx context.Context) LookupScopeConnectionResultOutput {
	return o
}

func (o LookupScopeConnectionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeConnectionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupScopeConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupScopeConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScopeConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScopeConnectionResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeConnectionResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupScopeConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupScopeConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupScopeConnectionResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeConnectionResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupScopeConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScopeConnectionResultOutput{})
}
