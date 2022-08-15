


package v20160331

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDatabaseAccountMongoDBDatabase(ctx *pulumi.Context, args *LookupDatabaseAccountMongoDBDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountMongoDBDatabaseResult, error) {
	var rv LookupDatabaseAccountMongoDBDatabaseResult
	err := ctx.Invoke("azure-native:documentdb/v20160331:getDatabaseAccountMongoDBDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseAccountMongoDBDatabaseArgs struct {
	AccountName       string `pulumi:"accountName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatabaseAccountMongoDBDatabaseResult struct {
	Id       string            `pulumi:"id"`
	Location *string           `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Tags     map[string]string `pulumi:"tags"`
	Type     string            `pulumi:"type"`
}

func LookupDatabaseAccountMongoDBDatabaseOutput(ctx *pulumi.Context, args LookupDatabaseAccountMongoDBDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseAccountMongoDBDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseAccountMongoDBDatabaseResult, error) {
			args := v.(LookupDatabaseAccountMongoDBDatabaseArgs)
			r, err := LookupDatabaseAccountMongoDBDatabase(ctx, &args, opts...)
			var s LookupDatabaseAccountMongoDBDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseAccountMongoDBDatabaseResultOutput)
}

type LookupDatabaseAccountMongoDBDatabaseOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatabaseAccountMongoDBDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountMongoDBDatabaseArgs)(nil)).Elem()
}


type LookupDatabaseAccountMongoDBDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseAccountMongoDBDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountMongoDBDatabaseResult)(nil)).Elem()
}

func (o LookupDatabaseAccountMongoDBDatabaseResultOutput) ToLookupDatabaseAccountMongoDBDatabaseResultOutput() LookupDatabaseAccountMongoDBDatabaseResultOutput {
	return o
}

func (o LookupDatabaseAccountMongoDBDatabaseResultOutput) ToLookupDatabaseAccountMongoDBDatabaseResultOutputWithContext(ctx context.Context) LookupDatabaseAccountMongoDBDatabaseResultOutput {
	return o
}

func (o LookupDatabaseAccountMongoDBDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountMongoDBDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountMongoDBDatabaseResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountMongoDBDatabaseResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountMongoDBDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountMongoDBDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountMongoDBDatabaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseAccountMongoDBDatabaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDatabaseAccountMongoDBDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountMongoDBDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseAccountMongoDBDatabaseResultOutput{})
}
