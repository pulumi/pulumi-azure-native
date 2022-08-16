


package v20190601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupEncryptionScope(ctx *pulumi.Context, args *LookupEncryptionScopeArgs, opts ...pulumi.InvokeOption) (*LookupEncryptionScopeResult, error) {
	var rv LookupEncryptionScopeResult
	err := ctx.Invoke("azure-native:storage/v20190601:getEncryptionScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEncryptionScopeArgs struct {
	AccountName         string `pulumi:"accountName"`
	EncryptionScopeName string `pulumi:"encryptionScopeName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupEncryptionScopeResult struct {
	CreationTime       string                                     `pulumi:"creationTime"`
	Id                 string                                     `pulumi:"id"`
	KeyVaultProperties *EncryptionScopeKeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	LastModifiedTime   string                                     `pulumi:"lastModifiedTime"`
	Name               string                                     `pulumi:"name"`
	Source             *string                                    `pulumi:"source"`
	State              *string                                    `pulumi:"state"`
	Type               string                                     `pulumi:"type"`
}

func LookupEncryptionScopeOutput(ctx *pulumi.Context, args LookupEncryptionScopeOutputArgs, opts ...pulumi.InvokeOption) LookupEncryptionScopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEncryptionScopeResult, error) {
			args := v.(LookupEncryptionScopeArgs)
			r, err := LookupEncryptionScope(ctx, &args, opts...)
			var s LookupEncryptionScopeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEncryptionScopeResultOutput)
}

type LookupEncryptionScopeOutputArgs struct {
	AccountName         pulumi.StringInput `pulumi:"accountName"`
	EncryptionScopeName pulumi.StringInput `pulumi:"encryptionScopeName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEncryptionScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEncryptionScopeArgs)(nil)).Elem()
}


type LookupEncryptionScopeResultOutput struct{ *pulumi.OutputState }

func (LookupEncryptionScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEncryptionScopeResult)(nil)).Elem()
}

func (o LookupEncryptionScopeResultOutput) ToLookupEncryptionScopeResultOutput() LookupEncryptionScopeResultOutput {
	return o
}

func (o LookupEncryptionScopeResultOutput) ToLookupEncryptionScopeResultOutputWithContext(ctx context.Context) LookupEncryptionScopeResultOutput {
	return o
}

func (o LookupEncryptionScopeResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEncryptionScopeResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupEncryptionScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEncryptionScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEncryptionScopeResultOutput) KeyVaultProperties() EncryptionScopeKeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupEncryptionScopeResult) *EncryptionScopeKeyVaultPropertiesResponse {
		return v.KeyVaultProperties
	}).(EncryptionScopeKeyVaultPropertiesResponsePtrOutput)
}

func (o LookupEncryptionScopeResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEncryptionScopeResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupEncryptionScopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEncryptionScopeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEncryptionScopeResultOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEncryptionScopeResult) *string { return v.Source }).(pulumi.StringPtrOutput)
}

func (o LookupEncryptionScopeResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEncryptionScopeResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o LookupEncryptionScopeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEncryptionScopeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEncryptionScopeResultOutput{})
}
