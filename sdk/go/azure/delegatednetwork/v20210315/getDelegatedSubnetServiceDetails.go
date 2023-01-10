


package v20210315

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDelegatedSubnetServiceDetails(ctx *pulumi.Context, args *LookupDelegatedSubnetServiceDetailsArgs, opts ...pulumi.InvokeOption) (*LookupDelegatedSubnetServiceDetailsResult, error) {
	var rv LookupDelegatedSubnetServiceDetailsResult
	err := ctx.Invoke("azure-native:delegatednetwork/v20210315:getDelegatedSubnetServiceDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDelegatedSubnetServiceDetailsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupDelegatedSubnetServiceDetailsResult struct {
	ControllerDetails *ControllerDetailsResponse `pulumi:"controllerDetails"`
	Id                string                     `pulumi:"id"`
	Location          *string                    `pulumi:"location"`
	Name              string                     `pulumi:"name"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	ResourceGuid      string                     `pulumi:"resourceGuid"`
	SubnetDetails     *SubnetDetailsResponse     `pulumi:"subnetDetails"`
	Tags              map[string]string          `pulumi:"tags"`
	Type              string                     `pulumi:"type"`
}

func LookupDelegatedSubnetServiceDetailsOutput(ctx *pulumi.Context, args LookupDelegatedSubnetServiceDetailsOutputArgs, opts ...pulumi.InvokeOption) LookupDelegatedSubnetServiceDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDelegatedSubnetServiceDetailsResult, error) {
			args := v.(LookupDelegatedSubnetServiceDetailsArgs)
			r, err := LookupDelegatedSubnetServiceDetails(ctx, &args, opts...)
			var s LookupDelegatedSubnetServiceDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDelegatedSubnetServiceDetailsResultOutput)
}

type LookupDelegatedSubnetServiceDetailsOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupDelegatedSubnetServiceDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDelegatedSubnetServiceDetailsArgs)(nil)).Elem()
}


type LookupDelegatedSubnetServiceDetailsResultOutput struct{ *pulumi.OutputState }

func (LookupDelegatedSubnetServiceDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDelegatedSubnetServiceDetailsResult)(nil)).Elem()
}

func (o LookupDelegatedSubnetServiceDetailsResultOutput) ToLookupDelegatedSubnetServiceDetailsResultOutput() LookupDelegatedSubnetServiceDetailsResultOutput {
	return o
}

func (o LookupDelegatedSubnetServiceDetailsResultOutput) ToLookupDelegatedSubnetServiceDetailsResultOutputWithContext(ctx context.Context) LookupDelegatedSubnetServiceDetailsResultOutput {
	return o
}

func (o LookupDelegatedSubnetServiceDetailsResultOutput) ControllerDetails() ControllerDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupDelegatedSubnetServiceDetailsResult) *ControllerDetailsResponse {
		return v.ControllerDetails
	}).(ControllerDetailsResponsePtrOutput)
}

func (o LookupDelegatedSubnetServiceDetailsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDelegatedSubnetServiceDetailsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDelegatedSubnetServiceDetailsResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDelegatedSubnetServiceDetailsResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDelegatedSubnetServiceDetailsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDelegatedSubnetServiceDetailsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDelegatedSubnetServiceDetailsResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDelegatedSubnetServiceDetailsResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDelegatedSubnetServiceDetailsResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDelegatedSubnetServiceDetailsResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupDelegatedSubnetServiceDetailsResultOutput) SubnetDetails() SubnetDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupDelegatedSubnetServiceDetailsResult) *SubnetDetailsResponse { return v.SubnetDetails }).(SubnetDetailsResponsePtrOutput)
}

func (o LookupDelegatedSubnetServiceDetailsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDelegatedSubnetServiceDetailsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDelegatedSubnetServiceDetailsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDelegatedSubnetServiceDetailsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDelegatedSubnetServiceDetailsResultOutput{})
}
