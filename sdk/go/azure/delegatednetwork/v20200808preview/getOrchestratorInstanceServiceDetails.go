


package v20200808preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupOrchestratorInstanceServiceDetails(ctx *pulumi.Context, args *LookupOrchestratorInstanceServiceDetailsArgs, opts ...pulumi.InvokeOption) (*LookupOrchestratorInstanceServiceDetailsResult, error) {
	var rv LookupOrchestratorInstanceServiceDetailsResult
	err := ctx.Invoke("azure-native:delegatednetwork/v20200808preview:getOrchestratorInstanceServiceDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrchestratorInstanceServiceDetailsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupOrchestratorInstanceServiceDetailsResult struct {
	ApiServerEndpoint    *string                       `pulumi:"apiServerEndpoint"`
	ClusterRootCA        *string                       `pulumi:"clusterRootCA"`
	ControllerDetails    ControllerDetailsResponse     `pulumi:"controllerDetails"`
	Id                   string                        `pulumi:"id"`
	Identity             *OrchestratorIdentityResponse `pulumi:"identity"`
	Kind                 string                        `pulumi:"kind"`
	Location             *string                       `pulumi:"location"`
	Name                 string                        `pulumi:"name"`
	OrchestratorAppId    *string                       `pulumi:"orchestratorAppId"`
	OrchestratorTenantId *string                       `pulumi:"orchestratorTenantId"`
	ProvisioningState    string                        `pulumi:"provisioningState"`
	ResourceGuid         string                        `pulumi:"resourceGuid"`
	Tags                 map[string]string             `pulumi:"tags"`
	Type                 string                        `pulumi:"type"`
}

func LookupOrchestratorInstanceServiceDetailsOutput(ctx *pulumi.Context, args LookupOrchestratorInstanceServiceDetailsOutputArgs, opts ...pulumi.InvokeOption) LookupOrchestratorInstanceServiceDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrchestratorInstanceServiceDetailsResult, error) {
			args := v.(LookupOrchestratorInstanceServiceDetailsArgs)
			r, err := LookupOrchestratorInstanceServiceDetails(ctx, &args, opts...)
			var s LookupOrchestratorInstanceServiceDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrchestratorInstanceServiceDetailsResultOutput)
}

type LookupOrchestratorInstanceServiceDetailsOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupOrchestratorInstanceServiceDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrchestratorInstanceServiceDetailsArgs)(nil)).Elem()
}


type LookupOrchestratorInstanceServiceDetailsResultOutput struct{ *pulumi.OutputState }

func (LookupOrchestratorInstanceServiceDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrchestratorInstanceServiceDetailsResult)(nil)).Elem()
}

func (o LookupOrchestratorInstanceServiceDetailsResultOutput) ToLookupOrchestratorInstanceServiceDetailsResultOutput() LookupOrchestratorInstanceServiceDetailsResultOutput {
	return o
}

func (o LookupOrchestratorInstanceServiceDetailsResultOutput) ToLookupOrchestratorInstanceServiceDetailsResultOutputWithContext(ctx context.Context) LookupOrchestratorInstanceServiceDetailsResultOutput {
	return o
}

func (o LookupOrchestratorInstanceServiceDetailsResultOutput) ApiServerEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrchestratorInstanceServiceDetailsResult) *string { return v.ApiServerEndpoint }).(pulumi.StringPtrOutput)
}

func (o LookupOrchestratorInstanceServiceDetailsResultOutput) ClusterRootCA() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrchestratorInstanceServiceDetailsResult) *string { return v.ClusterRootCA }).(pulumi.StringPtrOutput)
}

func (o LookupOrchestratorInstanceServiceDetailsResultOutput) ControllerDetails() ControllerDetailsResponseOutput {
	return o.ApplyT(func(v LookupOrchestratorInstanceServiceDetailsResult) ControllerDetailsResponse {
		return v.ControllerDetails
	}).(ControllerDetailsResponseOutput)
}

func (o LookupOrchestratorInstanceServiceDetailsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrchestratorInstanceServiceDetailsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOrchestratorInstanceServiceDetailsResultOutput) Identity() OrchestratorIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupOrchestratorInstanceServiceDetailsResult) *OrchestratorIdentityResponse {
		return v.Identity
	}).(OrchestratorIdentityResponsePtrOutput)
}

func (o LookupOrchestratorInstanceServiceDetailsResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrchestratorInstanceServiceDetailsResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupOrchestratorInstanceServiceDetailsResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrchestratorInstanceServiceDetailsResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupOrchestratorInstanceServiceDetailsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrchestratorInstanceServiceDetailsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOrchestratorInstanceServiceDetailsResultOutput) OrchestratorAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrchestratorInstanceServiceDetailsResult) *string { return v.OrchestratorAppId }).(pulumi.StringPtrOutput)
}

func (o LookupOrchestratorInstanceServiceDetailsResultOutput) OrchestratorTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrchestratorInstanceServiceDetailsResult) *string { return v.OrchestratorTenantId }).(pulumi.StringPtrOutput)
}

func (o LookupOrchestratorInstanceServiceDetailsResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrchestratorInstanceServiceDetailsResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupOrchestratorInstanceServiceDetailsResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrchestratorInstanceServiceDetailsResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupOrchestratorInstanceServiceDetailsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOrchestratorInstanceServiceDetailsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupOrchestratorInstanceServiceDetailsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrchestratorInstanceServiceDetailsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrchestratorInstanceServiceDetailsResultOutput{})
}
