


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLinker(ctx *pulumi.Context, args *LookupLinkerArgs, opts ...pulumi.InvokeOption) (*LookupLinkerResult, error) {
	var rv LookupLinkerResult
	err := ctx.Invoke("azure-native:servicelinker/v20221101preview:getLinker", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkerArgs struct {
	LinkerName  string `pulumi:"linkerName"`
	ResourceUri string `pulumi:"resourceUri"`
}


type LookupLinkerResult struct {
	AuthInfo              interface{}                    `pulumi:"authInfo"`
	ClientType            *string                        `pulumi:"clientType"`
	ConfigurationInfo     *ConfigurationInfoResponse     `pulumi:"configurationInfo"`
	Id                    string                         `pulumi:"id"`
	Name                  string                         `pulumi:"name"`
	ProvisioningState     string                         `pulumi:"provisioningState"`
	PublicNetworkSolution *PublicNetworkSolutionResponse `pulumi:"publicNetworkSolution"`
	Scope                 *string                        `pulumi:"scope"`
	SecretStore           *SecretStoreResponse           `pulumi:"secretStore"`
	SystemData            SystemDataResponse             `pulumi:"systemData"`
	TargetService         interface{}                    `pulumi:"targetService"`
	Type                  string                         `pulumi:"type"`
	VNetSolution          *VNetSolutionResponse          `pulumi:"vNetSolution"`
}

func LookupLinkerOutput(ctx *pulumi.Context, args LookupLinkerOutputArgs, opts ...pulumi.InvokeOption) LookupLinkerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLinkerResult, error) {
			args := v.(LookupLinkerArgs)
			r, err := LookupLinker(ctx, &args, opts...)
			var s LookupLinkerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLinkerResultOutput)
}

type LookupLinkerOutputArgs struct {
	LinkerName  pulumi.StringInput `pulumi:"linkerName"`
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (LookupLinkerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkerArgs)(nil)).Elem()
}


type LookupLinkerResultOutput struct{ *pulumi.OutputState }

func (LookupLinkerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkerResult)(nil)).Elem()
}

func (o LookupLinkerResultOutput) ToLookupLinkerResultOutput() LookupLinkerResultOutput {
	return o
}

func (o LookupLinkerResultOutput) ToLookupLinkerResultOutputWithContext(ctx context.Context) LookupLinkerResultOutput {
	return o
}

func (o LookupLinkerResultOutput) AuthInfo() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupLinkerResult) interface{} { return v.AuthInfo }).(pulumi.AnyOutput)
}

func (o LookupLinkerResultOutput) ClientType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLinkerResult) *string { return v.ClientType }).(pulumi.StringPtrOutput)
}

func (o LookupLinkerResultOutput) ConfigurationInfo() ConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupLinkerResult) *ConfigurationInfoResponse { return v.ConfigurationInfo }).(ConfigurationInfoResponsePtrOutput)
}

func (o LookupLinkerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLinkerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLinkerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupLinkerResultOutput) PublicNetworkSolution() PublicNetworkSolutionResponsePtrOutput {
	return o.ApplyT(func(v LookupLinkerResult) *PublicNetworkSolutionResponse { return v.PublicNetworkSolution }).(PublicNetworkSolutionResponsePtrOutput)
}

func (o LookupLinkerResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLinkerResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o LookupLinkerResultOutput) SecretStore() SecretStoreResponsePtrOutput {
	return o.ApplyT(func(v LookupLinkerResult) *SecretStoreResponse { return v.SecretStore }).(SecretStoreResponsePtrOutput)
}

func (o LookupLinkerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLinkerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupLinkerResultOutput) TargetService() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupLinkerResult) interface{} { return v.TargetService }).(pulumi.AnyOutput)
}

func (o LookupLinkerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkerResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupLinkerResultOutput) VNetSolution() VNetSolutionResponsePtrOutput {
	return o.ApplyT(func(v LookupLinkerResult) *VNetSolutionResponse { return v.VNetSolution }).(VNetSolutionResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLinkerResultOutput{})
}
