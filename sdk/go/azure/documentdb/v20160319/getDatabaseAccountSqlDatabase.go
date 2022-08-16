


package v20160319

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDatabaseAccountSqlDatabase(ctx *pulumi.Context, args *LookupDatabaseAccountSqlDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountSqlDatabaseResult, error) {
	var rv LookupDatabaseAccountSqlDatabaseResult
	err := ctx.Invoke("azure-native:documentdb/v20160319:getDatabaseAccountSqlDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseAccountSqlDatabaseArgs struct {
	AccountName       string `pulumi:"accountName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatabaseAccountSqlDatabaseResult struct {
	Colls    *string           `pulumi:"colls"`
	Etag     *string           `pulumi:"etag"`
	Id       string            `pulumi:"id"`
	Location *string           `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Rid      *string           `pulumi:"rid"`
	Tags     map[string]string `pulumi:"tags"`
	Ts       interface{}       `pulumi:"ts"`
	Type     string            `pulumi:"type"`
	Users    *string           `pulumi:"users"`
}

func LookupDatabaseAccountSqlDatabaseOutput(ctx *pulumi.Context, args LookupDatabaseAccountSqlDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseAccountSqlDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseAccountSqlDatabaseResult, error) {
			args := v.(LookupDatabaseAccountSqlDatabaseArgs)
			r, err := LookupDatabaseAccountSqlDatabase(ctx, &args, opts...)
			var s LookupDatabaseAccountSqlDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseAccountSqlDatabaseResultOutput)
}

type LookupDatabaseAccountSqlDatabaseOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatabaseAccountSqlDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountSqlDatabaseArgs)(nil)).Elem()
}


type LookupDatabaseAccountSqlDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseAccountSqlDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountSqlDatabaseResult)(nil)).Elem()
}

func (o LookupDatabaseAccountSqlDatabaseResultOutput) ToLookupDatabaseAccountSqlDatabaseResultOutput() LookupDatabaseAccountSqlDatabaseResultOutput {
	return o
}

func (o LookupDatabaseAccountSqlDatabaseResultOutput) ToLookupDatabaseAccountSqlDatabaseResultOutputWithContext(ctx context.Context) LookupDatabaseAccountSqlDatabaseResultOutput {
	return o
}

func (o LookupDatabaseAccountSqlDatabaseResultOutput) Colls() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlDatabaseResult) *string { return v.Colls }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountSqlDatabaseResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlDatabaseResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountSqlDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountSqlDatabaseResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlDatabaseResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountSqlDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountSqlDatabaseResultOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlDatabaseResult) *string { return v.Rid }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountSqlDatabaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlDatabaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDatabaseAccountSqlDatabaseResultOutput) Ts() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlDatabaseResult) interface{} { return v.Ts }).(pulumi.AnyOutput)
}

func (o LookupDatabaseAccountSqlDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountSqlDatabaseResultOutput) Users() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountSqlDatabaseResult) *string { return v.Users }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseAccountSqlDatabaseResultOutput{})
}
