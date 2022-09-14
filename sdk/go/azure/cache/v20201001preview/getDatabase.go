


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabase(ctx *pulumi.Context, args *LookupDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseResult, error) {
	var rv LookupDatabaseResult
	err := ctx.Invoke("azure-native:cache/v20201001preview:getDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatabaseResult struct {
	ClientProtocol    *string          `pulumi:"clientProtocol"`
	ClusteringPolicy  *string          `pulumi:"clusteringPolicy"`
	EvictionPolicy    *string          `pulumi:"evictionPolicy"`
	Id                string           `pulumi:"id"`
	Modules           []ModuleResponse `pulumi:"modules"`
	Name              string           `pulumi:"name"`
	Port              *int             `pulumi:"port"`
	ProvisioningState string           `pulumi:"provisioningState"`
	ResourceState     string           `pulumi:"resourceState"`
	Type              string           `pulumi:"type"`
}

func LookupDatabaseOutput(ctx *pulumi.Context, args LookupDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseResult, error) {
			args := v.(LookupDatabaseArgs)
			r, err := LookupDatabase(ctx, &args, opts...)
			var s LookupDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseResultOutput)
}

type LookupDatabaseOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseArgs)(nil)).Elem()
}


type LookupDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseResult)(nil)).Elem()
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutput() LookupDatabaseResultOutput {
	return o
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutputWithContext(ctx context.Context) LookupDatabaseResultOutput {
	return o
}

func (o LookupDatabaseResultOutput) ClientProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.ClientProtocol }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseResultOutput) ClusteringPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.ClusteringPolicy }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseResultOutput) EvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.EvictionPolicy }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) Modules() ModuleResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseResult) []ModuleResponse { return v.Modules }).(ModuleResponseArrayOutput)
}

func (o LookupDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o LookupDatabaseResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseResultOutput{})
}
