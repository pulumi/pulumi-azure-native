


package v20200202preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTransparentDataEncryption(ctx *pulumi.Context, args *LookupTransparentDataEncryptionArgs, opts ...pulumi.InvokeOption) (*LookupTransparentDataEncryptionResult, error) {
	var rv LookupTransparentDataEncryptionResult
	err := ctx.Invoke("azure-native:sql/v20200202preview:getTransparentDataEncryption", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTransparentDataEncryptionArgs struct {
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
	TdeName           string `pulumi:"tdeName"`
}


type LookupTransparentDataEncryptionResult struct {
	Id    string `pulumi:"id"`
	Name  string `pulumi:"name"`
	State string `pulumi:"state"`
	Type  string `pulumi:"type"`
}

func LookupTransparentDataEncryptionOutput(ctx *pulumi.Context, args LookupTransparentDataEncryptionOutputArgs, opts ...pulumi.InvokeOption) LookupTransparentDataEncryptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTransparentDataEncryptionResult, error) {
			args := v.(LookupTransparentDataEncryptionArgs)
			r, err := LookupTransparentDataEncryption(ctx, &args, opts...)
			var s LookupTransparentDataEncryptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTransparentDataEncryptionResultOutput)
}

type LookupTransparentDataEncryptionOutputArgs struct {
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
	TdeName           pulumi.StringInput `pulumi:"tdeName"`
}

func (LookupTransparentDataEncryptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransparentDataEncryptionArgs)(nil)).Elem()
}


type LookupTransparentDataEncryptionResultOutput struct{ *pulumi.OutputState }

func (LookupTransparentDataEncryptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransparentDataEncryptionResult)(nil)).Elem()
}

func (o LookupTransparentDataEncryptionResultOutput) ToLookupTransparentDataEncryptionResultOutput() LookupTransparentDataEncryptionResultOutput {
	return o
}

func (o LookupTransparentDataEncryptionResultOutput) ToLookupTransparentDataEncryptionResultOutputWithContext(ctx context.Context) LookupTransparentDataEncryptionResultOutput {
	return o
}

func (o LookupTransparentDataEncryptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransparentDataEncryptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTransparentDataEncryptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransparentDataEncryptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTransparentDataEncryptionResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransparentDataEncryptionResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupTransparentDataEncryptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransparentDataEncryptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTransparentDataEncryptionResultOutput{})
}
