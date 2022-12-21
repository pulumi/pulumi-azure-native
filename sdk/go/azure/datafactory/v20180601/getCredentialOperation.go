


package v20180601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCredentialOperation(ctx *pulumi.Context, args *LookupCredentialOperationArgs, opts ...pulumi.InvokeOption) (*LookupCredentialOperationResult, error) {
	var rv LookupCredentialOperationResult
	err := ctx.Invoke("azure-native:datafactory/v20180601:getCredentialOperation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCredentialOperationArgs struct {
	CredentialName    string `pulumi:"credentialName"`
	FactoryName       string `pulumi:"factoryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCredentialOperationResult struct {
	Etag       string                            `pulumi:"etag"`
	Id         string                            `pulumi:"id"`
	Name       string                            `pulumi:"name"`
	Properties ManagedIdentityCredentialResponse `pulumi:"properties"`
	Type       string                            `pulumi:"type"`
}

func LookupCredentialOperationOutput(ctx *pulumi.Context, args LookupCredentialOperationOutputArgs, opts ...pulumi.InvokeOption) LookupCredentialOperationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCredentialOperationResult, error) {
			args := v.(LookupCredentialOperationArgs)
			r, err := LookupCredentialOperation(ctx, &args, opts...)
			var s LookupCredentialOperationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCredentialOperationResultOutput)
}

type LookupCredentialOperationOutputArgs struct {
	CredentialName    pulumi.StringInput `pulumi:"credentialName"`
	FactoryName       pulumi.StringInput `pulumi:"factoryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCredentialOperationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCredentialOperationArgs)(nil)).Elem()
}


type LookupCredentialOperationResultOutput struct{ *pulumi.OutputState }

func (LookupCredentialOperationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCredentialOperationResult)(nil)).Elem()
}

func (o LookupCredentialOperationResultOutput) ToLookupCredentialOperationResultOutput() LookupCredentialOperationResultOutput {
	return o
}

func (o LookupCredentialOperationResultOutput) ToLookupCredentialOperationResultOutputWithContext(ctx context.Context) LookupCredentialOperationResultOutput {
	return o
}

func (o LookupCredentialOperationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialOperationResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupCredentialOperationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialOperationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCredentialOperationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialOperationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCredentialOperationResultOutput) Properties() ManagedIdentityCredentialResponseOutput {
	return o.ApplyT(func(v LookupCredentialOperationResult) ManagedIdentityCredentialResponse { return v.Properties }).(ManagedIdentityCredentialResponseOutput)
}

func (o LookupCredentialOperationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialOperationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCredentialOperationResultOutput{})
}
