


package operationsmanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementAssociation(ctx *pulumi.Context, args *LookupManagementAssociationArgs, opts ...pulumi.InvokeOption) (*LookupManagementAssociationResult, error) {
	var rv LookupManagementAssociationResult
	err := ctx.Invoke("azure-native:operationsmanagement:getManagementAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementAssociationArgs struct {
	ManagementAssociationName string `pulumi:"managementAssociationName"`
	ProviderName              string `pulumi:"providerName"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	ResourceName              string `pulumi:"resourceName"`
	ResourceType              string `pulumi:"resourceType"`
}


type LookupManagementAssociationResult struct {
	Id         string                                  `pulumi:"id"`
	Location   *string                                 `pulumi:"location"`
	Name       string                                  `pulumi:"name"`
	Properties ManagementAssociationPropertiesResponse `pulumi:"properties"`
	Type       string                                  `pulumi:"type"`
}

func LookupManagementAssociationOutput(ctx *pulumi.Context, args LookupManagementAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupManagementAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagementAssociationResult, error) {
			args := v.(LookupManagementAssociationArgs)
			r, err := LookupManagementAssociation(ctx, &args, opts...)
			var s LookupManagementAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagementAssociationResultOutput)
}

type LookupManagementAssociationOutputArgs struct {
	ManagementAssociationName pulumi.StringInput `pulumi:"managementAssociationName"`
	ProviderName              pulumi.StringInput `pulumi:"providerName"`
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName              pulumi.StringInput `pulumi:"resourceName"`
	ResourceType              pulumi.StringInput `pulumi:"resourceType"`
}

func (LookupManagementAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementAssociationArgs)(nil)).Elem()
}


type LookupManagementAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupManagementAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementAssociationResult)(nil)).Elem()
}

func (o LookupManagementAssociationResultOutput) ToLookupManagementAssociationResultOutput() LookupManagementAssociationResultOutput {
	return o
}

func (o LookupManagementAssociationResultOutput) ToLookupManagementAssociationResultOutputWithContext(ctx context.Context) LookupManagementAssociationResultOutput {
	return o
}

func (o LookupManagementAssociationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementAssociationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagementAssociationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementAssociationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupManagementAssociationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementAssociationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagementAssociationResultOutput) Properties() ManagementAssociationPropertiesResponseOutput {
	return o.ApplyT(func(v LookupManagementAssociationResult) ManagementAssociationPropertiesResponse { return v.Properties }).(ManagementAssociationPropertiesResponseOutput)
}

func (o LookupManagementAssociationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementAssociationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementAssociationResultOutput{})
}
