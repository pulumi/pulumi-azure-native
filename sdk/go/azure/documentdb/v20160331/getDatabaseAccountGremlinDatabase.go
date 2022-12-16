


package v20160331

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDatabaseAccountGremlinDatabase(ctx *pulumi.Context, args *LookupDatabaseAccountGremlinDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountGremlinDatabaseResult, error) {
	var rv LookupDatabaseAccountGremlinDatabaseResult
	err := ctx.Invoke("azure-native:documentdb/v20160331:getDatabaseAccountGremlinDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseAccountGremlinDatabaseArgs struct {
	AccountName       string `pulumi:"accountName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatabaseAccountGremlinDatabaseResult struct {
	Etag     *string           `pulumi:"etag"`
	Id       string            `pulumi:"id"`
	Location *string           `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Rid      *string           `pulumi:"rid"`
	Tags     map[string]string `pulumi:"tags"`
	Ts       interface{}       `pulumi:"ts"`
	Type     string            `pulumi:"type"`
}

func LookupDatabaseAccountGremlinDatabaseOutput(ctx *pulumi.Context, args LookupDatabaseAccountGremlinDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseAccountGremlinDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseAccountGremlinDatabaseResult, error) {
			args := v.(LookupDatabaseAccountGremlinDatabaseArgs)
			r, err := LookupDatabaseAccountGremlinDatabase(ctx, &args, opts...)
			var s LookupDatabaseAccountGremlinDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseAccountGremlinDatabaseResultOutput)
}

type LookupDatabaseAccountGremlinDatabaseOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatabaseAccountGremlinDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountGremlinDatabaseArgs)(nil)).Elem()
}


type LookupDatabaseAccountGremlinDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseAccountGremlinDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountGremlinDatabaseResult)(nil)).Elem()
}

func (o LookupDatabaseAccountGremlinDatabaseResultOutput) ToLookupDatabaseAccountGremlinDatabaseResultOutput() LookupDatabaseAccountGremlinDatabaseResultOutput {
	return o
}

func (o LookupDatabaseAccountGremlinDatabaseResultOutput) ToLookupDatabaseAccountGremlinDatabaseResultOutputWithContext(ctx context.Context) LookupDatabaseAccountGremlinDatabaseResultOutput {
	return o
}

func (o LookupDatabaseAccountGremlinDatabaseResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinDatabaseResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountGremlinDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountGremlinDatabaseResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinDatabaseResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountGremlinDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountGremlinDatabaseResultOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinDatabaseResult) *string { return v.Rid }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountGremlinDatabaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinDatabaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDatabaseAccountGremlinDatabaseResultOutput) Ts() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinDatabaseResult) interface{} { return v.Ts }).(pulumi.AnyOutput)
}

func (o LookupDatabaseAccountGremlinDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountGremlinDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseAccountGremlinDatabaseResultOutput{})
}
