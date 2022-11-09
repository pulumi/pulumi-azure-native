


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTable(ctx *pulumi.Context, args *LookupTableArgs, opts ...pulumi.InvokeOption) (*LookupTableResult, error) {
	var rv LookupTableResult
	err := ctx.Invoke("azure-native:storage/v20220501:getTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTableArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TableName         string `pulumi:"tableName"`
}


type LookupTableResult struct {
	Id                string                          `pulumi:"id"`
	Name              string                          `pulumi:"name"`
	SignedIdentifiers []TableSignedIdentifierResponse `pulumi:"signedIdentifiers"`
	TableName         string                          `pulumi:"tableName"`
	Type              string                          `pulumi:"type"`
}

func LookupTableOutput(ctx *pulumi.Context, args LookupTableOutputArgs, opts ...pulumi.InvokeOption) LookupTableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTableResult, error) {
			args := v.(LookupTableArgs)
			r, err := LookupTable(ctx, &args, opts...)
			var s LookupTableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTableResultOutput)
}

type LookupTableOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TableName         pulumi.StringInput `pulumi:"tableName"`
}

func (LookupTableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableArgs)(nil)).Elem()
}


type LookupTableResultOutput struct{ *pulumi.OutputState }

func (LookupTableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableResult)(nil)).Elem()
}

func (o LookupTableResultOutput) ToLookupTableResultOutput() LookupTableResultOutput {
	return o
}

func (o LookupTableResultOutput) ToLookupTableResultOutputWithContext(ctx context.Context) LookupTableResultOutput {
	return o
}

func (o LookupTableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTableResultOutput) SignedIdentifiers() TableSignedIdentifierResponseArrayOutput {
	return o.ApplyT(func(v LookupTableResult) []TableSignedIdentifierResponse { return v.SignedIdentifiers }).(TableSignedIdentifierResponseArrayOutput)
}

func (o LookupTableResultOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableResult) string { return v.TableName }).(pulumi.StringOutput)
}

func (o LookupTableResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTableResultOutput{})
}
