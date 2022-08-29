


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlPoolTransparentDataEncryption(ctx *pulumi.Context, args *LookupSqlPoolTransparentDataEncryptionArgs, opts ...pulumi.InvokeOption) (*LookupSqlPoolTransparentDataEncryptionResult, error) {
	var rv LookupSqlPoolTransparentDataEncryptionResult
	err := ctx.Invoke("azure-native:synapse/v20210301:getSqlPoolTransparentDataEncryption", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlPoolTransparentDataEncryptionArgs struct {
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	SqlPoolName                   string `pulumi:"sqlPoolName"`
	TransparentDataEncryptionName string `pulumi:"transparentDataEncryptionName"`
	WorkspaceName                 string `pulumi:"workspaceName"`
}


type LookupSqlPoolTransparentDataEncryptionResult struct {
	Id       string  `pulumi:"id"`
	Location string  `pulumi:"location"`
	Name     string  `pulumi:"name"`
	Status   *string `pulumi:"status"`
	Type     string  `pulumi:"type"`
}

func LookupSqlPoolTransparentDataEncryptionOutput(ctx *pulumi.Context, args LookupSqlPoolTransparentDataEncryptionOutputArgs, opts ...pulumi.InvokeOption) LookupSqlPoolTransparentDataEncryptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlPoolTransparentDataEncryptionResult, error) {
			args := v.(LookupSqlPoolTransparentDataEncryptionArgs)
			r, err := LookupSqlPoolTransparentDataEncryption(ctx, &args, opts...)
			var s LookupSqlPoolTransparentDataEncryptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlPoolTransparentDataEncryptionResultOutput)
}

type LookupSqlPoolTransparentDataEncryptionOutputArgs struct {
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	SqlPoolName                   pulumi.StringInput `pulumi:"sqlPoolName"`
	TransparentDataEncryptionName pulumi.StringInput `pulumi:"transparentDataEncryptionName"`
	WorkspaceName                 pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupSqlPoolTransparentDataEncryptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolTransparentDataEncryptionArgs)(nil)).Elem()
}


type LookupSqlPoolTransparentDataEncryptionResultOutput struct{ *pulumi.OutputState }

func (LookupSqlPoolTransparentDataEncryptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolTransparentDataEncryptionResult)(nil)).Elem()
}

func (o LookupSqlPoolTransparentDataEncryptionResultOutput) ToLookupSqlPoolTransparentDataEncryptionResultOutput() LookupSqlPoolTransparentDataEncryptionResultOutput {
	return o
}

func (o LookupSqlPoolTransparentDataEncryptionResultOutput) ToLookupSqlPoolTransparentDataEncryptionResultOutputWithContext(ctx context.Context) LookupSqlPoolTransparentDataEncryptionResultOutput {
	return o
}

func (o LookupSqlPoolTransparentDataEncryptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolTransparentDataEncryptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlPoolTransparentDataEncryptionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolTransparentDataEncryptionResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSqlPoolTransparentDataEncryptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolTransparentDataEncryptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlPoolTransparentDataEncryptionResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolTransparentDataEncryptionResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolTransparentDataEncryptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolTransparentDataEncryptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlPoolTransparentDataEncryptionResultOutput{})
}
