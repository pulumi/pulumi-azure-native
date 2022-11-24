


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnector(ctx *pulumi.Context, args *LookupConnectorArgs, opts ...pulumi.InvokeOption) (*LookupConnectorResult, error) {
	var rv LookupConnectorResult
	err := ctx.Invoke("azure-native:servicelinker/v20221101preview:getConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorArgs struct {
	ConnectorName     string `pulumi:"connectorName"`
	Location          string `pulumi:"location"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupConnectorResult struct {
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

func LookupConnectorOutput(ctx *pulumi.Context, args LookupConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectorResult, error) {
			args := v.(LookupConnectorArgs)
			r, err := LookupConnector(ctx, &args, opts...)
			var s LookupConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectorResultOutput)
}

type LookupConnectorOutputArgs struct {
	ConnectorName     pulumi.StringInput `pulumi:"connectorName"`
	Location          pulumi.StringInput `pulumi:"location"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorArgs)(nil)).Elem()
}


type LookupConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorResult)(nil)).Elem()
}

func (o LookupConnectorResultOutput) ToLookupConnectorResultOutput() LookupConnectorResultOutput {
	return o
}

func (o LookupConnectorResultOutput) ToLookupConnectorResultOutputWithContext(ctx context.Context) LookupConnectorResultOutput {
	return o
}

func (o LookupConnectorResultOutput) AuthInfo() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupConnectorResult) interface{} { return v.AuthInfo }).(pulumi.AnyOutput)
}

func (o LookupConnectorResultOutput) ClientType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.ClientType }).(pulumi.StringPtrOutput)
}

func (o LookupConnectorResultOutput) ConfigurationInfo() ConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *ConfigurationInfoResponse { return v.ConfigurationInfo }).(ConfigurationInfoResponsePtrOutput)
}

func (o LookupConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectorResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupConnectorResultOutput) PublicNetworkSolution() PublicNetworkSolutionResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *PublicNetworkSolutionResponse { return v.PublicNetworkSolution }).(PublicNetworkSolutionResponsePtrOutput)
}

func (o LookupConnectorResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o LookupConnectorResultOutput) SecretStore() SecretStoreResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *SecretStoreResponse { return v.SecretStore }).(SecretStoreResponsePtrOutput)
}

func (o LookupConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupConnectorResultOutput) TargetService() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupConnectorResult) interface{} { return v.TargetService }).(pulumi.AnyOutput)
}

func (o LookupConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupConnectorResultOutput) VNetSolution() VNetSolutionResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *VNetSolutionResponse { return v.VNetSolution }).(VNetSolutionResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectorResultOutput{})
}
