


package v20200113preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCredential(ctx *pulumi.Context, args *LookupCredentialArgs, opts ...pulumi.InvokeOption) (*LookupCredentialResult, error) {
	var rv LookupCredentialResult
	err := ctx.Invoke("azure-native:automation/v20200113preview:getCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCredentialArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	CredentialName        string `pulumi:"credentialName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupCredentialResult struct {
	CreationTime     string  `pulumi:"creationTime"`
	Description      *string `pulumi:"description"`
	Id               string  `pulumi:"id"`
	LastModifiedTime string  `pulumi:"lastModifiedTime"`
	Name             string  `pulumi:"name"`
	Type             string  `pulumi:"type"`
	UserName         string  `pulumi:"userName"`
}

func LookupCredentialOutput(ctx *pulumi.Context, args LookupCredentialOutputArgs, opts ...pulumi.InvokeOption) LookupCredentialResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCredentialResult, error) {
			args := v.(LookupCredentialArgs)
			r, err := LookupCredential(ctx, &args, opts...)
			var s LookupCredentialResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCredentialResultOutput)
}

type LookupCredentialOutputArgs struct {
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	CredentialName        pulumi.StringInput `pulumi:"credentialName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCredentialArgs)(nil)).Elem()
}


type LookupCredentialResultOutput struct{ *pulumi.OutputState }

func (LookupCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCredentialResult)(nil)).Elem()
}

func (o LookupCredentialResultOutput) ToLookupCredentialResultOutput() LookupCredentialResultOutput {
	return o
}

func (o LookupCredentialResultOutput) ToLookupCredentialResultOutputWithContext(ctx context.Context) LookupCredentialResultOutput {
	return o
}

func (o LookupCredentialResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupCredentialResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCredentialResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupCredentialResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCredentialResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupCredentialResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCredentialResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupCredentialResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCredentialResultOutput{})
}
