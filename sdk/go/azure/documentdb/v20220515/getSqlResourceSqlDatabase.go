


package v20220515

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlResourceSqlDatabase(ctx *pulumi.Context, args *LookupSqlResourceSqlDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupSqlResourceSqlDatabaseResult, error) {
	var rv LookupSqlResourceSqlDatabaseResult
	err := ctx.Invoke("azure-native:documentdb/v20220515:getSqlResourceSqlDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlResourceSqlDatabaseArgs struct {
	AccountName       string `pulumi:"accountName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupSqlResourceSqlDatabaseResult struct {
	Id       string                                    `pulumi:"id"`
	Location *string                                   `pulumi:"location"`
	Name     string                                    `pulumi:"name"`
	Options  *SqlDatabaseGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *SqlDatabaseGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                         `pulumi:"tags"`
	Type     string                                    `pulumi:"type"`
}

func LookupSqlResourceSqlDatabaseOutput(ctx *pulumi.Context, args LookupSqlResourceSqlDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupSqlResourceSqlDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlResourceSqlDatabaseResult, error) {
			args := v.(LookupSqlResourceSqlDatabaseArgs)
			r, err := LookupSqlResourceSqlDatabase(ctx, &args, opts...)
			var s LookupSqlResourceSqlDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlResourceSqlDatabaseResultOutput)
}

type LookupSqlResourceSqlDatabaseOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupSqlResourceSqlDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlDatabaseArgs)(nil)).Elem()
}


type LookupSqlResourceSqlDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupSqlResourceSqlDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlDatabaseResult)(nil)).Elem()
}

func (o LookupSqlResourceSqlDatabaseResultOutput) ToLookupSqlResourceSqlDatabaseResultOutput() LookupSqlResourceSqlDatabaseResultOutput {
	return o
}

func (o LookupSqlResourceSqlDatabaseResultOutput) ToLookupSqlResourceSqlDatabaseResultOutputWithContext(ctx context.Context) LookupSqlResourceSqlDatabaseResultOutput {
	return o
}

func (o LookupSqlResourceSqlDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlResourceSqlDatabaseResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlDatabaseResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupSqlResourceSqlDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlResourceSqlDatabaseResultOutput) Options() SqlDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlDatabaseResult) *SqlDatabaseGetPropertiesResponseOptions { return v.Options }).(SqlDatabaseGetPropertiesResponseOptionsPtrOutput)
}

func (o LookupSqlResourceSqlDatabaseResultOutput) Resource() SqlDatabaseGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlDatabaseResult) *SqlDatabaseGetPropertiesResponseResource {
		return v.Resource
	}).(SqlDatabaseGetPropertiesResponseResourcePtrOutput)
}

func (o LookupSqlResourceSqlDatabaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlDatabaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSqlResourceSqlDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlResourceSqlDatabaseResultOutput{})
}
