


package v20220131preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListUserAssignedIdentityAssociatedResources(ctx *pulumi.Context, args *ListUserAssignedIdentityAssociatedResourcesArgs, opts ...pulumi.InvokeOption) (*ListUserAssignedIdentityAssociatedResourcesResult, error) {
	var rv ListUserAssignedIdentityAssociatedResourcesResult
	err := ctx.Invoke("azure-native:managedidentity/v20220131preview:listUserAssignedIdentityAssociatedResources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListUserAssignedIdentityAssociatedResourcesArgs struct {
	Filter            *string `pulumi:"filter"`
	Orderby           *string `pulumi:"orderby"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ResourceName      string  `pulumi:"resourceName"`
	Skip              *int    `pulumi:"skip"`
	Skiptoken         *string `pulumi:"skiptoken"`
	Top               *int    `pulumi:"top"`
}


type ListUserAssignedIdentityAssociatedResourcesResult struct {
	NextLink   string                  `pulumi:"nextLink"`
	TotalCount float64                 `pulumi:"totalCount"`
	Value      []AzureResourceResponse `pulumi:"value"`
}

func ListUserAssignedIdentityAssociatedResourcesOutput(ctx *pulumi.Context, args ListUserAssignedIdentityAssociatedResourcesOutputArgs, opts ...pulumi.InvokeOption) ListUserAssignedIdentityAssociatedResourcesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListUserAssignedIdentityAssociatedResourcesResult, error) {
			args := v.(ListUserAssignedIdentityAssociatedResourcesArgs)
			r, err := ListUserAssignedIdentityAssociatedResources(ctx, &args, opts...)
			var s ListUserAssignedIdentityAssociatedResourcesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListUserAssignedIdentityAssociatedResourcesResultOutput)
}

type ListUserAssignedIdentityAssociatedResourcesOutputArgs struct {
	Filter            pulumi.StringPtrInput `pulumi:"filter"`
	Orderby           pulumi.StringPtrInput `pulumi:"orderby"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput    `pulumi:"resourceName"`
	Skip              pulumi.IntPtrInput    `pulumi:"skip"`
	Skiptoken         pulumi.StringPtrInput `pulumi:"skiptoken"`
	Top               pulumi.IntPtrInput    `pulumi:"top"`
}

func (ListUserAssignedIdentityAssociatedResourcesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListUserAssignedIdentityAssociatedResourcesArgs)(nil)).Elem()
}


type ListUserAssignedIdentityAssociatedResourcesResultOutput struct{ *pulumi.OutputState }

func (ListUserAssignedIdentityAssociatedResourcesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListUserAssignedIdentityAssociatedResourcesResult)(nil)).Elem()
}

func (o ListUserAssignedIdentityAssociatedResourcesResultOutput) ToListUserAssignedIdentityAssociatedResourcesResultOutput() ListUserAssignedIdentityAssociatedResourcesResultOutput {
	return o
}

func (o ListUserAssignedIdentityAssociatedResourcesResultOutput) ToListUserAssignedIdentityAssociatedResourcesResultOutputWithContext(ctx context.Context) ListUserAssignedIdentityAssociatedResourcesResultOutput {
	return o
}

func (o ListUserAssignedIdentityAssociatedResourcesResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListUserAssignedIdentityAssociatedResourcesResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListUserAssignedIdentityAssociatedResourcesResultOutput) TotalCount() pulumi.Float64Output {
	return o.ApplyT(func(v ListUserAssignedIdentityAssociatedResourcesResult) float64 { return v.TotalCount }).(pulumi.Float64Output)
}

func (o ListUserAssignedIdentityAssociatedResourcesResultOutput) Value() AzureResourceResponseArrayOutput {
	return o.ApplyT(func(v ListUserAssignedIdentityAssociatedResourcesResult) []AzureResourceResponse { return v.Value }).(AzureResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListUserAssignedIdentityAssociatedResourcesResultOutput{})
}
