


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListLocalUserKeys(ctx *pulumi.Context, args *ListLocalUserKeysArgs, opts ...pulumi.InvokeOption) (*ListLocalUserKeysResult, error) {
	var rv ListLocalUserKeysResult
	err := ctx.Invoke("azure-native:storage/v20220501:listLocalUserKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListLocalUserKeysArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Username          string `pulumi:"username"`
}


type ListLocalUserKeysResult struct {
	SharedKey         string                 `pulumi:"sharedKey"`
	SshAuthorizedKeys []SshPublicKeyResponse `pulumi:"sshAuthorizedKeys"`
}

func ListLocalUserKeysOutput(ctx *pulumi.Context, args ListLocalUserKeysOutputArgs, opts ...pulumi.InvokeOption) ListLocalUserKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListLocalUserKeysResult, error) {
			args := v.(ListLocalUserKeysArgs)
			r, err := ListLocalUserKeys(ctx, &args, opts...)
			var s ListLocalUserKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListLocalUserKeysResultOutput)
}

type ListLocalUserKeysOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Username          pulumi.StringInput `pulumi:"username"`
}

func (ListLocalUserKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListLocalUserKeysArgs)(nil)).Elem()
}


type ListLocalUserKeysResultOutput struct{ *pulumi.OutputState }

func (ListLocalUserKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListLocalUserKeysResult)(nil)).Elem()
}

func (o ListLocalUserKeysResultOutput) ToListLocalUserKeysResultOutput() ListLocalUserKeysResultOutput {
	return o
}

func (o ListLocalUserKeysResultOutput) ToListLocalUserKeysResultOutputWithContext(ctx context.Context) ListLocalUserKeysResultOutput {
	return o
}

func (o ListLocalUserKeysResultOutput) SharedKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListLocalUserKeysResult) string { return v.SharedKey }).(pulumi.StringOutput)
}

func (o ListLocalUserKeysResultOutput) SshAuthorizedKeys() SshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v ListLocalUserKeysResult) []SshPublicKeyResponse { return v.SshAuthorizedKeys }).(SshPublicKeyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListLocalUserKeysResultOutput{})
}
