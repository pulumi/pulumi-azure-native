


package v20160331

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDatabaseAccountCassandraKeyspace(ctx *pulumi.Context, args *LookupDatabaseAccountCassandraKeyspaceArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountCassandraKeyspaceResult, error) {
	var rv LookupDatabaseAccountCassandraKeyspaceResult
	err := ctx.Invoke("azure-native:documentdb/v20160331:getDatabaseAccountCassandraKeyspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseAccountCassandraKeyspaceArgs struct {
	AccountName       string `pulumi:"accountName"`
	KeyspaceName      string `pulumi:"keyspaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatabaseAccountCassandraKeyspaceResult struct {
	Id       string            `pulumi:"id"`
	Location *string           `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Tags     map[string]string `pulumi:"tags"`
	Type     string            `pulumi:"type"`
}

func LookupDatabaseAccountCassandraKeyspaceOutput(ctx *pulumi.Context, args LookupDatabaseAccountCassandraKeyspaceOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseAccountCassandraKeyspaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseAccountCassandraKeyspaceResult, error) {
			args := v.(LookupDatabaseAccountCassandraKeyspaceArgs)
			r, err := LookupDatabaseAccountCassandraKeyspace(ctx, &args, opts...)
			var s LookupDatabaseAccountCassandraKeyspaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseAccountCassandraKeyspaceResultOutput)
}

type LookupDatabaseAccountCassandraKeyspaceOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	KeyspaceName      pulumi.StringInput `pulumi:"keyspaceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatabaseAccountCassandraKeyspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountCassandraKeyspaceArgs)(nil)).Elem()
}


type LookupDatabaseAccountCassandraKeyspaceResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseAccountCassandraKeyspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountCassandraKeyspaceResult)(nil)).Elem()
}

func (o LookupDatabaseAccountCassandraKeyspaceResultOutput) ToLookupDatabaseAccountCassandraKeyspaceResultOutput() LookupDatabaseAccountCassandraKeyspaceResultOutput {
	return o
}

func (o LookupDatabaseAccountCassandraKeyspaceResultOutput) ToLookupDatabaseAccountCassandraKeyspaceResultOutputWithContext(ctx context.Context) LookupDatabaseAccountCassandraKeyspaceResultOutput {
	return o
}

func (o LookupDatabaseAccountCassandraKeyspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountCassandraKeyspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountCassandraKeyspaceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountCassandraKeyspaceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountCassandraKeyspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountCassandraKeyspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountCassandraKeyspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseAccountCassandraKeyspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDatabaseAccountCassandraKeyspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountCassandraKeyspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseAccountCassandraKeyspaceResultOutput{})
}
