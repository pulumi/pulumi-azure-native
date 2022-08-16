


package v20220215preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTableResourceTable(ctx *pulumi.Context, args *LookupTableResourceTableArgs, opts ...pulumi.InvokeOption) (*LookupTableResourceTableResult, error) {
	var rv LookupTableResourceTableResult
	err := ctx.Invoke("azure-native:documentdb/v20220215preview:getTableResourceTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTableResourceTableArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TableName         string `pulumi:"tableName"`
}


type LookupTableResourceTableResult struct {
	Id       string                              `pulumi:"id"`
	Identity *ManagedServiceIdentityResponse     `pulumi:"identity"`
	Location *string                             `pulumi:"location"`
	Name     string                              `pulumi:"name"`
	Options  *TableGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *TableGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                   `pulumi:"tags"`
	Type     string                              `pulumi:"type"`
}

func LookupTableResourceTableOutput(ctx *pulumi.Context, args LookupTableResourceTableOutputArgs, opts ...pulumi.InvokeOption) LookupTableResourceTableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTableResourceTableResult, error) {
			args := v.(LookupTableResourceTableArgs)
			r, err := LookupTableResourceTable(ctx, &args, opts...)
			var s LookupTableResourceTableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTableResourceTableResultOutput)
}

type LookupTableResourceTableOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TableName         pulumi.StringInput `pulumi:"tableName"`
}

func (LookupTableResourceTableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableResourceTableArgs)(nil)).Elem()
}


type LookupTableResourceTableResultOutput struct{ *pulumi.OutputState }

func (LookupTableResourceTableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableResourceTableResult)(nil)).Elem()
}

func (o LookupTableResourceTableResultOutput) ToLookupTableResourceTableResultOutput() LookupTableResourceTableResultOutput {
	return o
}

func (o LookupTableResourceTableResultOutput) ToLookupTableResourceTableResultOutputWithContext(ctx context.Context) LookupTableResourceTableResultOutput {
	return o
}

func (o LookupTableResourceTableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTableResourceTableResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupTableResourceTableResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupTableResourceTableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTableResourceTableResultOutput) Options() TableGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) *TableGetPropertiesResponseOptions { return v.Options }).(TableGetPropertiesResponseOptionsPtrOutput)
}

func (o LookupTableResourceTableResultOutput) Resource() TableGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) *TableGetPropertiesResponseResource { return v.Resource }).(TableGetPropertiesResponseResourcePtrOutput)
}

func (o LookupTableResourceTableResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupTableResourceTableResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTableResourceTableResultOutput{})
}
