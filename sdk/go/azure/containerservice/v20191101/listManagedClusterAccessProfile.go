


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListManagedClusterAccessProfile(ctx *pulumi.Context, args *ListManagedClusterAccessProfileArgs, opts ...pulumi.InvokeOption) (*ListManagedClusterAccessProfileResult, error) {
	var rv ListManagedClusterAccessProfileResult
	err := ctx.Invoke("azure-native:containerservice/v20191101:listManagedClusterAccessProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagedClusterAccessProfileArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
	RoleName          string `pulumi:"roleName"`
}


type ListManagedClusterAccessProfileResult struct {
	Id         string            `pulumi:"id"`
	KubeConfig *string           `pulumi:"kubeConfig"`
	Location   string            `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

func ListManagedClusterAccessProfileOutput(ctx *pulumi.Context, args ListManagedClusterAccessProfileOutputArgs, opts ...pulumi.InvokeOption) ListManagedClusterAccessProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListManagedClusterAccessProfileResult, error) {
			args := v.(ListManagedClusterAccessProfileArgs)
			r, err := ListManagedClusterAccessProfile(ctx, &args, opts...)
			var s ListManagedClusterAccessProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListManagedClusterAccessProfileResultOutput)
}

type ListManagedClusterAccessProfileOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
	RoleName          pulumi.StringInput `pulumi:"roleName"`
}

func (ListManagedClusterAccessProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterAccessProfileArgs)(nil)).Elem()
}


type ListManagedClusterAccessProfileResultOutput struct{ *pulumi.OutputState }

func (ListManagedClusterAccessProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterAccessProfileResult)(nil)).Elem()
}

func (o ListManagedClusterAccessProfileResultOutput) ToListManagedClusterAccessProfileResultOutput() ListManagedClusterAccessProfileResultOutput {
	return o
}

func (o ListManagedClusterAccessProfileResultOutput) ToListManagedClusterAccessProfileResultOutputWithContext(ctx context.Context) ListManagedClusterAccessProfileResultOutput {
	return o
}

func (o ListManagedClusterAccessProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListManagedClusterAccessProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListManagedClusterAccessProfileResultOutput) KubeConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListManagedClusterAccessProfileResult) *string { return v.KubeConfig }).(pulumi.StringPtrOutput)
}

func (o ListManagedClusterAccessProfileResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ListManagedClusterAccessProfileResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o ListManagedClusterAccessProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListManagedClusterAccessProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListManagedClusterAccessProfileResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListManagedClusterAccessProfileResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ListManagedClusterAccessProfileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListManagedClusterAccessProfileResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListManagedClusterAccessProfileResultOutput{})
}
