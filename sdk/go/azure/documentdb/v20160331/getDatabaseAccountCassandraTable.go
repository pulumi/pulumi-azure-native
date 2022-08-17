


package v20160331

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDatabaseAccountCassandraTable(ctx *pulumi.Context, args *LookupDatabaseAccountCassandraTableArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountCassandraTableResult, error) {
	var rv LookupDatabaseAccountCassandraTableResult
	err := ctx.Invoke("azure-native:documentdb/v20160331:getDatabaseAccountCassandraTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseAccountCassandraTableArgs struct {
	AccountName       string `pulumi:"accountName"`
	KeyspaceName      string `pulumi:"keyspaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TableName         string `pulumi:"tableName"`
}


type LookupDatabaseAccountCassandraTableResult struct {
	DefaultTtl *int                     `pulumi:"defaultTtl"`
	Id         string                   `pulumi:"id"`
	Location   *string                  `pulumi:"location"`
	Name       string                   `pulumi:"name"`
	Schema     *CassandraSchemaResponse `pulumi:"schema"`
	Tags       map[string]string        `pulumi:"tags"`
	Type       string                   `pulumi:"type"`
}

func LookupDatabaseAccountCassandraTableOutput(ctx *pulumi.Context, args LookupDatabaseAccountCassandraTableOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseAccountCassandraTableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseAccountCassandraTableResult, error) {
			args := v.(LookupDatabaseAccountCassandraTableArgs)
			r, err := LookupDatabaseAccountCassandraTable(ctx, &args, opts...)
			var s LookupDatabaseAccountCassandraTableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseAccountCassandraTableResultOutput)
}

type LookupDatabaseAccountCassandraTableOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	KeyspaceName      pulumi.StringInput `pulumi:"keyspaceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TableName         pulumi.StringInput `pulumi:"tableName"`
}

func (LookupDatabaseAccountCassandraTableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountCassandraTableArgs)(nil)).Elem()
}


type LookupDatabaseAccountCassandraTableResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseAccountCassandraTableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountCassandraTableResult)(nil)).Elem()
}

func (o LookupDatabaseAccountCassandraTableResultOutput) ToLookupDatabaseAccountCassandraTableResultOutput() LookupDatabaseAccountCassandraTableResultOutput {
	return o
}

func (o LookupDatabaseAccountCassandraTableResultOutput) ToLookupDatabaseAccountCassandraTableResultOutputWithContext(ctx context.Context) LookupDatabaseAccountCassandraTableResultOutput {
	return o
}

func (o LookupDatabaseAccountCassandraTableResultOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountCassandraTableResult) *int { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o LookupDatabaseAccountCassandraTableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountCassandraTableResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountCassandraTableResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountCassandraTableResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountCassandraTableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountCassandraTableResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountCassandraTableResultOutput) Schema() CassandraSchemaResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountCassandraTableResult) *CassandraSchemaResponse { return v.Schema }).(CassandraSchemaResponsePtrOutput)
}

func (o LookupDatabaseAccountCassandraTableResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseAccountCassandraTableResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDatabaseAccountCassandraTableResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountCassandraTableResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseAccountCassandraTableResultOutput{})
}
