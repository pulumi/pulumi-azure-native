


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLocalUser(ctx *pulumi.Context, args *LookupLocalUserArgs, opts ...pulumi.InvokeOption) (*LookupLocalUserResult, error) {
	var rv LookupLocalUserResult
	err := ctx.Invoke("azure-native:storage/v20220501:getLocalUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLocalUserArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Username          string `pulumi:"username"`
}


type LookupLocalUserResult struct {
	HasSharedKey      *bool                     `pulumi:"hasSharedKey"`
	HasSshKey         *bool                     `pulumi:"hasSshKey"`
	HasSshPassword    *bool                     `pulumi:"hasSshPassword"`
	HomeDirectory     *string                   `pulumi:"homeDirectory"`
	Id                string                    `pulumi:"id"`
	Name              string                    `pulumi:"name"`
	PermissionScopes  []PermissionScopeResponse `pulumi:"permissionScopes"`
	Sid               string                    `pulumi:"sid"`
	SshAuthorizedKeys []SshPublicKeyResponse    `pulumi:"sshAuthorizedKeys"`
	SystemData        SystemDataResponse        `pulumi:"systemData"`
	Type              string                    `pulumi:"type"`
}

func LookupLocalUserOutput(ctx *pulumi.Context, args LookupLocalUserOutputArgs, opts ...pulumi.InvokeOption) LookupLocalUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocalUserResult, error) {
			args := v.(LookupLocalUserArgs)
			r, err := LookupLocalUser(ctx, &args, opts...)
			var s LookupLocalUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLocalUserResultOutput)
}

type LookupLocalUserOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Username          pulumi.StringInput `pulumi:"username"`
}

func (LookupLocalUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalUserArgs)(nil)).Elem()
}


type LookupLocalUserResultOutput struct{ *pulumi.OutputState }

func (LookupLocalUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalUserResult)(nil)).Elem()
}

func (o LookupLocalUserResultOutput) ToLookupLocalUserResultOutput() LookupLocalUserResultOutput {
	return o
}

func (o LookupLocalUserResultOutput) ToLookupLocalUserResultOutputWithContext(ctx context.Context) LookupLocalUserResultOutput {
	return o
}

func (o LookupLocalUserResultOutput) HasSharedKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalUserResult) *bool { return v.HasSharedKey }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalUserResultOutput) HasSshKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalUserResult) *bool { return v.HasSshKey }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalUserResultOutput) HasSshPassword() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalUserResult) *bool { return v.HasSshPassword }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalUserResultOutput) HomeDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalUserResult) *string { return v.HomeDirectory }).(pulumi.StringPtrOutput)
}

func (o LookupLocalUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalUserResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLocalUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalUserResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLocalUserResultOutput) PermissionScopes() PermissionScopeResponseArrayOutput {
	return o.ApplyT(func(v LookupLocalUserResult) []PermissionScopeResponse { return v.PermissionScopes }).(PermissionScopeResponseArrayOutput)
}

func (o LookupLocalUserResultOutput) Sid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalUserResult) string { return v.Sid }).(pulumi.StringOutput)
}

func (o LookupLocalUserResultOutput) SshAuthorizedKeys() SshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v LookupLocalUserResult) []SshPublicKeyResponse { return v.SshAuthorizedKeys }).(SshPublicKeyResponseArrayOutput)
}

func (o LookupLocalUserResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLocalUserResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupLocalUserResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalUserResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocalUserResultOutput{})
}
