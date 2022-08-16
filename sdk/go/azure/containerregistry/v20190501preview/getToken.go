


package v20190501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupToken(ctx *pulumi.Context, args *LookupTokenArgs, opts ...pulumi.InvokeOption) (*LookupTokenResult, error) {
	var rv LookupTokenResult
	err := ctx.Invoke("azure-native:containerregistry/v20190501preview:getToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTokenArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TokenName         string `pulumi:"tokenName"`
}


type LookupTokenResult struct {
	CreationDate      string                              `pulumi:"creationDate"`
	Credentials       *TokenCredentialsPropertiesResponse `pulumi:"credentials"`
	Id                string                              `pulumi:"id"`
	Name              string                              `pulumi:"name"`
	ProvisioningState string                              `pulumi:"provisioningState"`
	ScopeMapId        *string                             `pulumi:"scopeMapId"`
	Status            *string                             `pulumi:"status"`
	SystemData        SystemDataResponse                  `pulumi:"systemData"`
	Type              string                              `pulumi:"type"`
}

func LookupTokenOutput(ctx *pulumi.Context, args LookupTokenOutputArgs, opts ...pulumi.InvokeOption) LookupTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTokenResult, error) {
			args := v.(LookupTokenArgs)
			r, err := LookupToken(ctx, &args, opts...)
			var s LookupTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTokenResultOutput)
}

type LookupTokenOutputArgs struct {
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TokenName         pulumi.StringInput `pulumi:"tokenName"`
}

func (LookupTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTokenArgs)(nil)).Elem()
}


type LookupTokenResultOutput struct{ *pulumi.OutputState }

func (LookupTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTokenResult)(nil)).Elem()
}

func (o LookupTokenResultOutput) ToLookupTokenResultOutput() LookupTokenResultOutput {
	return o
}

func (o LookupTokenResultOutput) ToLookupTokenResultOutputWithContext(ctx context.Context) LookupTokenResultOutput {
	return o
}

func (o LookupTokenResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTokenResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupTokenResultOutput) Credentials() TokenCredentialsPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupTokenResult) *TokenCredentialsPropertiesResponse { return v.Credentials }).(TokenCredentialsPropertiesResponsePtrOutput)
}

func (o LookupTokenResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTokenResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTokenResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTokenResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTokenResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTokenResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupTokenResultOutput) ScopeMapId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTokenResult) *string { return v.ScopeMapId }).(pulumi.StringPtrOutput)
}

func (o LookupTokenResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTokenResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupTokenResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTokenResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupTokenResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTokenResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTokenResultOutput{})
}
