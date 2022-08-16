


package v20191101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataCollectionRuleAssociation(ctx *pulumi.Context, args *LookupDataCollectionRuleAssociationArgs, opts ...pulumi.InvokeOption) (*LookupDataCollectionRuleAssociationResult, error) {
	var rv LookupDataCollectionRuleAssociationResult
	err := ctx.Invoke("azure-native:insights/v20191101preview:getDataCollectionRuleAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataCollectionRuleAssociationArgs struct {
	AssociationName string `pulumi:"associationName"`
	ResourceUri     string `pulumi:"resourceUri"`
}


type LookupDataCollectionRuleAssociationResult struct {
	DataCollectionRuleId *string `pulumi:"dataCollectionRuleId"`
	Description          *string `pulumi:"description"`
	Etag                 string  `pulumi:"etag"`
	Id                   string  `pulumi:"id"`
	Name                 string  `pulumi:"name"`
	ProvisioningState    string  `pulumi:"provisioningState"`
	Type                 string  `pulumi:"type"`
}

func LookupDataCollectionRuleAssociationOutput(ctx *pulumi.Context, args LookupDataCollectionRuleAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupDataCollectionRuleAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataCollectionRuleAssociationResult, error) {
			args := v.(LookupDataCollectionRuleAssociationArgs)
			r, err := LookupDataCollectionRuleAssociation(ctx, &args, opts...)
			var s LookupDataCollectionRuleAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataCollectionRuleAssociationResultOutput)
}

type LookupDataCollectionRuleAssociationOutputArgs struct {
	AssociationName pulumi.StringInput `pulumi:"associationName"`
	ResourceUri     pulumi.StringInput `pulumi:"resourceUri"`
}

func (LookupDataCollectionRuleAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataCollectionRuleAssociationArgs)(nil)).Elem()
}


type LookupDataCollectionRuleAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupDataCollectionRuleAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataCollectionRuleAssociationResult)(nil)).Elem()
}

func (o LookupDataCollectionRuleAssociationResultOutput) ToLookupDataCollectionRuleAssociationResultOutput() LookupDataCollectionRuleAssociationResultOutput {
	return o
}

func (o LookupDataCollectionRuleAssociationResultOutput) ToLookupDataCollectionRuleAssociationResultOutputWithContext(ctx context.Context) LookupDataCollectionRuleAssociationResultOutput {
	return o
}

func (o LookupDataCollectionRuleAssociationResultOutput) DataCollectionRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleAssociationResult) *string { return v.DataCollectionRuleId }).(pulumi.StringPtrOutput)
}

func (o LookupDataCollectionRuleAssociationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleAssociationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupDataCollectionRuleAssociationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleAssociationResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupDataCollectionRuleAssociationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleAssociationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataCollectionRuleAssociationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleAssociationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataCollectionRuleAssociationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleAssociationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDataCollectionRuleAssociationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleAssociationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataCollectionRuleAssociationResultOutput{})
}
