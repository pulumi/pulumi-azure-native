


package v20191202preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupRemoteRenderingAccount(ctx *pulumi.Context, args *LookupRemoteRenderingAccountArgs, opts ...pulumi.InvokeOption) (*LookupRemoteRenderingAccountResult, error) {
	var rv LookupRemoteRenderingAccountResult
	err := ctx.Invoke("azure-native:mixedreality/v20191202preview:getRemoteRenderingAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRemoteRenderingAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRemoteRenderingAccountResult struct {
	AccountDomain string            `pulumi:"accountDomain"`
	AccountId     string            `pulumi:"accountId"`
	Id            string            `pulumi:"id"`
	Identity      *IdentityResponse `pulumi:"identity"`
	Location      string            `pulumi:"location"`
	Name          string            `pulumi:"name"`
	Tags          map[string]string `pulumi:"tags"`
	Type          string            `pulumi:"type"`
}

func LookupRemoteRenderingAccountOutput(ctx *pulumi.Context, args LookupRemoteRenderingAccountOutputArgs, opts ...pulumi.InvokeOption) LookupRemoteRenderingAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemoteRenderingAccountResult, error) {
			args := v.(LookupRemoteRenderingAccountArgs)
			r, err := LookupRemoteRenderingAccount(ctx, &args, opts...)
			var s LookupRemoteRenderingAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemoteRenderingAccountResultOutput)
}

type LookupRemoteRenderingAccountOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRemoteRenderingAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteRenderingAccountArgs)(nil)).Elem()
}


type LookupRemoteRenderingAccountResultOutput struct{ *pulumi.OutputState }

func (LookupRemoteRenderingAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteRenderingAccountResult)(nil)).Elem()
}

func (o LookupRemoteRenderingAccountResultOutput) ToLookupRemoteRenderingAccountResultOutput() LookupRemoteRenderingAccountResultOutput {
	return o
}

func (o LookupRemoteRenderingAccountResultOutput) ToLookupRemoteRenderingAccountResultOutputWithContext(ctx context.Context) LookupRemoteRenderingAccountResultOutput {
	return o
}

func (o LookupRemoteRenderingAccountResultOutput) AccountDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteRenderingAccountResult) string { return v.AccountDomain }).(pulumi.StringOutput)
}

func (o LookupRemoteRenderingAccountResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteRenderingAccountResult) string { return v.AccountId }).(pulumi.StringOutput)
}

func (o LookupRemoteRenderingAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteRenderingAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemoteRenderingAccountResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupRemoteRenderingAccountResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupRemoteRenderingAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteRenderingAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupRemoteRenderingAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteRenderingAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRemoteRenderingAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRemoteRenderingAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupRemoteRenderingAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteRenderingAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemoteRenderingAccountResultOutput{})
}
