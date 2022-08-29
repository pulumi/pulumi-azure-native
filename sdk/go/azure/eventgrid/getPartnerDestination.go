


package eventgrid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPartnerDestination(ctx *pulumi.Context, args *LookupPartnerDestinationArgs, opts ...pulumi.InvokeOption) (*LookupPartnerDestinationResult, error) {
	var rv LookupPartnerDestinationResult
	err := ctx.Invoke("azure-native:eventgrid:getPartnerDestination", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPartnerDestinationArgs struct {
	PartnerDestinationName string `pulumi:"partnerDestinationName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupPartnerDestinationResult struct {
	ActivationState                 *string            `pulumi:"activationState"`
	EndpointBaseUrl                 *string            `pulumi:"endpointBaseUrl"`
	EndpointServiceContext          *string            `pulumi:"endpointServiceContext"`
	ExpirationTimeIfNotActivatedUtc *string            `pulumi:"expirationTimeIfNotActivatedUtc"`
	Id                              string             `pulumi:"id"`
	Location                        string             `pulumi:"location"`
	MessageForActivation            *string            `pulumi:"messageForActivation"`
	Name                            string             `pulumi:"name"`
	PartnerRegistrationImmutableId  *string            `pulumi:"partnerRegistrationImmutableId"`
	ProvisioningState               *string            `pulumi:"provisioningState"`
	SystemData                      SystemDataResponse `pulumi:"systemData"`
	Tags                            map[string]string  `pulumi:"tags"`
	Type                            string             `pulumi:"type"`
}

func LookupPartnerDestinationOutput(ctx *pulumi.Context, args LookupPartnerDestinationOutputArgs, opts ...pulumi.InvokeOption) LookupPartnerDestinationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPartnerDestinationResult, error) {
			args := v.(LookupPartnerDestinationArgs)
			r, err := LookupPartnerDestination(ctx, &args, opts...)
			var s LookupPartnerDestinationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPartnerDestinationResultOutput)
}

type LookupPartnerDestinationOutputArgs struct {
	PartnerDestinationName pulumi.StringInput `pulumi:"partnerDestinationName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPartnerDestinationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerDestinationArgs)(nil)).Elem()
}


type LookupPartnerDestinationResultOutput struct{ *pulumi.OutputState }

func (LookupPartnerDestinationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerDestinationResult)(nil)).Elem()
}

func (o LookupPartnerDestinationResultOutput) ToLookupPartnerDestinationResultOutput() LookupPartnerDestinationResultOutput {
	return o
}

func (o LookupPartnerDestinationResultOutput) ToLookupPartnerDestinationResultOutputWithContext(ctx context.Context) LookupPartnerDestinationResultOutput {
	return o
}

func (o LookupPartnerDestinationResultOutput) ActivationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) *string { return v.ActivationState }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerDestinationResultOutput) EndpointBaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) *string { return v.EndpointBaseUrl }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerDestinationResultOutput) EndpointServiceContext() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) *string { return v.EndpointServiceContext }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerDestinationResultOutput) ExpirationTimeIfNotActivatedUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) *string { return v.ExpirationTimeIfNotActivatedUtc }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerDestinationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPartnerDestinationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPartnerDestinationResultOutput) MessageForActivation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) *string { return v.MessageForActivation }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerDestinationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPartnerDestinationResultOutput) PartnerRegistrationImmutableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) *string { return v.PartnerRegistrationImmutableId }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerDestinationResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerDestinationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPartnerDestinationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPartnerDestinationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPartnerDestinationResultOutput{})
}
