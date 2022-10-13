


package v20200401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupCassandraResourceCassandraKeyspace(ctx *pulumi.Context, args *LookupCassandraResourceCassandraKeyspaceArgs, opts ...pulumi.InvokeOption) (*LookupCassandraResourceCassandraKeyspaceResult, error) {
	var rv LookupCassandraResourceCassandraKeyspaceResult
	err := ctx.Invoke("azure-native:documentdb/v20200401:getCassandraResourceCassandraKeyspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCassandraResourceCassandraKeyspaceArgs struct {
	AccountName       string `pulumi:"accountName"`
	KeyspaceName      string `pulumi:"keyspaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCassandraResourceCassandraKeyspaceResult struct {
	Id       string                                          `pulumi:"id"`
	Location *string                                         `pulumi:"location"`
	Name     string                                          `pulumi:"name"`
	Options  *CassandraKeyspaceGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *CassandraKeyspaceGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                               `pulumi:"tags"`
	Type     string                                          `pulumi:"type"`
}

func LookupCassandraResourceCassandraKeyspaceOutput(ctx *pulumi.Context, args LookupCassandraResourceCassandraKeyspaceOutputArgs, opts ...pulumi.InvokeOption) LookupCassandraResourceCassandraKeyspaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCassandraResourceCassandraKeyspaceResult, error) {
			args := v.(LookupCassandraResourceCassandraKeyspaceArgs)
			r, err := LookupCassandraResourceCassandraKeyspace(ctx, &args, opts...)
			var s LookupCassandraResourceCassandraKeyspaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCassandraResourceCassandraKeyspaceResultOutput)
}

type LookupCassandraResourceCassandraKeyspaceOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	KeyspaceName      pulumi.StringInput `pulumi:"keyspaceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCassandraResourceCassandraKeyspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCassandraResourceCassandraKeyspaceArgs)(nil)).Elem()
}


type LookupCassandraResourceCassandraKeyspaceResultOutput struct{ *pulumi.OutputState }

func (LookupCassandraResourceCassandraKeyspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCassandraResourceCassandraKeyspaceResult)(nil)).Elem()
}

func (o LookupCassandraResourceCassandraKeyspaceResultOutput) ToLookupCassandraResourceCassandraKeyspaceResultOutput() LookupCassandraResourceCassandraKeyspaceResultOutput {
	return o
}

func (o LookupCassandraResourceCassandraKeyspaceResultOutput) ToLookupCassandraResourceCassandraKeyspaceResultOutputWithContext(ctx context.Context) LookupCassandraResourceCassandraKeyspaceResultOutput {
	return o
}

func (o LookupCassandraResourceCassandraKeyspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraKeyspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCassandraResourceCassandraKeyspaceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraKeyspaceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupCassandraResourceCassandraKeyspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraKeyspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCassandraResourceCassandraKeyspaceResultOutput) Options() CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraKeyspaceResult) *CassandraKeyspaceGetPropertiesResponseOptions {
		return v.Options
	}).(CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput)
}

func (o LookupCassandraResourceCassandraKeyspaceResultOutput) Resource() CassandraKeyspaceGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraKeyspaceResult) *CassandraKeyspaceGetPropertiesResponseResource {
		return v.Resource
	}).(CassandraKeyspaceGetPropertiesResponseResourcePtrOutput)
}

func (o LookupCassandraResourceCassandraKeyspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraKeyspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCassandraResourceCassandraKeyspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraKeyspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCassandraResourceCassandraKeyspaceResultOutput{})
}
