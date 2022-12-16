


package v20210827

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDataConnection(ctx *pulumi.Context, args *LookupDataConnectionArgs, opts ...pulumi.InvokeOption) (*LookupDataConnectionResult, error) {
	var rv LookupDataConnectionResult
	err := ctx.Invoke("azure-native:kusto/v20210827:getDataConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataConnectionArgs struct {
	ClusterName        string `pulumi:"clusterName"`
	DataConnectionName string `pulumi:"dataConnectionName"`
	DatabaseName       string `pulumi:"databaseName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupDataConnectionResult struct {
	Id       string  `pulumi:"id"`
	Kind     string  `pulumi:"kind"`
	Location *string `pulumi:"location"`
	Name     string  `pulumi:"name"`
	Type     string  `pulumi:"type"`
}

func LookupDataConnectionOutput(ctx *pulumi.Context, args LookupDataConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupDataConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataConnectionResult, error) {
			args := v.(LookupDataConnectionArgs)
			r, err := LookupDataConnection(ctx, &args, opts...)
			var s LookupDataConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataConnectionResultOutput)
}

type LookupDataConnectionOutputArgs struct {
	ClusterName        pulumi.StringInput `pulumi:"clusterName"`
	DataConnectionName pulumi.StringInput `pulumi:"dataConnectionName"`
	DatabaseName       pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDataConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataConnectionArgs)(nil)).Elem()
}


type LookupDataConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupDataConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataConnectionResult)(nil)).Elem()
}

func (o LookupDataConnectionResultOutput) ToLookupDataConnectionResultOutput() LookupDataConnectionResultOutput {
	return o
}

func (o LookupDataConnectionResultOutput) ToLookupDataConnectionResultOutputWithContext(ctx context.Context) LookupDataConnectionResultOutput {
	return o
}

func (o LookupDataConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataConnectionResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataConnectionResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupDataConnectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataConnectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDataConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataConnectionResultOutput{})
}
