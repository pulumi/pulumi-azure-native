


package v20220515

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIotConnectorFhirDestination(ctx *pulumi.Context, args *LookupIotConnectorFhirDestinationArgs, opts ...pulumi.InvokeOption) (*LookupIotConnectorFhirDestinationResult, error) {
	var rv LookupIotConnectorFhirDestinationResult
	err := ctx.Invoke("azure-native:healthcareapis/v20220515:getIotConnectorFhirDestination", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotConnectorFhirDestinationArgs struct {
	FhirDestinationName string `pulumi:"fhirDestinationName"`
	IotConnectorName    string `pulumi:"iotConnectorName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	WorkspaceName       string `pulumi:"workspaceName"`
}


type LookupIotConnectorFhirDestinationResult struct {
	Etag                           *string                      `pulumi:"etag"`
	FhirMapping                    IotMappingPropertiesResponse `pulumi:"fhirMapping"`
	FhirServiceResourceId          string                       `pulumi:"fhirServiceResourceId"`
	Id                             string                       `pulumi:"id"`
	Location                       *string                      `pulumi:"location"`
	Name                           string                       `pulumi:"name"`
	ResourceIdentityResolutionType string                       `pulumi:"resourceIdentityResolutionType"`
	SystemData                     SystemDataResponse           `pulumi:"systemData"`
	Type                           string                       `pulumi:"type"`
}

func LookupIotConnectorFhirDestinationOutput(ctx *pulumi.Context, args LookupIotConnectorFhirDestinationOutputArgs, opts ...pulumi.InvokeOption) LookupIotConnectorFhirDestinationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIotConnectorFhirDestinationResult, error) {
			args := v.(LookupIotConnectorFhirDestinationArgs)
			r, err := LookupIotConnectorFhirDestination(ctx, &args, opts...)
			var s LookupIotConnectorFhirDestinationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIotConnectorFhirDestinationResultOutput)
}

type LookupIotConnectorFhirDestinationOutputArgs struct {
	FhirDestinationName pulumi.StringInput `pulumi:"fhirDestinationName"`
	IotConnectorName    pulumi.StringInput `pulumi:"iotConnectorName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName       pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupIotConnectorFhirDestinationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotConnectorFhirDestinationArgs)(nil)).Elem()
}


type LookupIotConnectorFhirDestinationResultOutput struct{ *pulumi.OutputState }

func (LookupIotConnectorFhirDestinationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotConnectorFhirDestinationResult)(nil)).Elem()
}

func (o LookupIotConnectorFhirDestinationResultOutput) ToLookupIotConnectorFhirDestinationResultOutput() LookupIotConnectorFhirDestinationResultOutput {
	return o
}

func (o LookupIotConnectorFhirDestinationResultOutput) ToLookupIotConnectorFhirDestinationResultOutputWithContext(ctx context.Context) LookupIotConnectorFhirDestinationResultOutput {
	return o
}

func (o LookupIotConnectorFhirDestinationResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotConnectorFhirDestinationResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupIotConnectorFhirDestinationResultOutput) FhirMapping() IotMappingPropertiesResponseOutput {
	return o.ApplyT(func(v LookupIotConnectorFhirDestinationResult) IotMappingPropertiesResponse { return v.FhirMapping }).(IotMappingPropertiesResponseOutput)
}

func (o LookupIotConnectorFhirDestinationResultOutput) FhirServiceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotConnectorFhirDestinationResult) string { return v.FhirServiceResourceId }).(pulumi.StringOutput)
}

func (o LookupIotConnectorFhirDestinationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotConnectorFhirDestinationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIotConnectorFhirDestinationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotConnectorFhirDestinationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupIotConnectorFhirDestinationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotConnectorFhirDestinationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIotConnectorFhirDestinationResultOutput) ResourceIdentityResolutionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotConnectorFhirDestinationResult) string { return v.ResourceIdentityResolutionType }).(pulumi.StringOutput)
}

func (o LookupIotConnectorFhirDestinationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIotConnectorFhirDestinationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupIotConnectorFhirDestinationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotConnectorFhirDestinationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIotConnectorFhirDestinationResultOutput{})
}
