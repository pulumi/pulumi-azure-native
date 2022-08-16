


package v20200501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	var rv LookupUserResult
	err := ctx.Invoke("azure-native:databoxedge/v20200501preview:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupUserResult struct {
	EncryptedPassword *AsymmetricEncryptedSecretResponse `pulumi:"encryptedPassword"`
	Id                string                             `pulumi:"id"`
	Name              string                             `pulumi:"name"`
	ShareAccessRights []ShareAccessRightResponse         `pulumi:"shareAccessRights"`
	Type              string                             `pulumi:"type"`
	UserType          string                             `pulumi:"userType"`
}

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserResult, error) {
			args := v.(LookupUserArgs)
			r, err := LookupUser(ctx, &args, opts...)
			var s LookupUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserResultOutput)
}

type LookupUserOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}


type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) EncryptedPassword() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v LookupUserResult) *AsymmetricEncryptedSecretResponse { return v.EncryptedPassword }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) ShareAccessRights() ShareAccessRightResponseArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []ShareAccessRightResponse { return v.ShareAccessRights }).(ShareAccessRightResponseArrayOutput)
}

func (o LookupUserResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) UserType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.UserType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}
