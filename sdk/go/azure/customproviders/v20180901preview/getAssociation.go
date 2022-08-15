


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAssociation(ctx *pulumi.Context, args *LookupAssociationArgs, opts ...pulumi.InvokeOption) (*LookupAssociationResult, error) {
	var rv LookupAssociationResult
	err := ctx.Invoke("azure-native:customproviders/v20180901preview:getAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssociationArgs struct {
	AssociationName string `pulumi:"associationName"`
	Scope           string `pulumi:"scope"`
}


type LookupAssociationResult struct {
	Id                string  `pulumi:"id"`
	Name              string  `pulumi:"name"`
	ProvisioningState string  `pulumi:"provisioningState"`
	TargetResourceId  *string `pulumi:"targetResourceId"`
	Type              string  `pulumi:"type"`
}

func LookupAssociationOutput(ctx *pulumi.Context, args LookupAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAssociationResult, error) {
			args := v.(LookupAssociationArgs)
			r, err := LookupAssociation(ctx, &args, opts...)
			var s LookupAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAssociationResultOutput)
}

type LookupAssociationOutputArgs struct {
	AssociationName pulumi.StringInput `pulumi:"associationName"`
	Scope           pulumi.StringInput `pulumi:"scope"`
}

func (LookupAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssociationArgs)(nil)).Elem()
}


type LookupAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssociationResult)(nil)).Elem()
}

func (o LookupAssociationResultOutput) ToLookupAssociationResultOutput() LookupAssociationResultOutput {
	return o
}

func (o LookupAssociationResultOutput) ToLookupAssociationResultOutputWithContext(ctx context.Context) LookupAssociationResultOutput {
	return o
}

func (o LookupAssociationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssociationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAssociationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssociationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAssociationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssociationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAssociationResultOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssociationResult) *string { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupAssociationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssociationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAssociationResultOutput{})
}
