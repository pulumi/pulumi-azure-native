


package v20221111preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDevBoxDefinition(ctx *pulumi.Context, args *LookupDevBoxDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupDevBoxDefinitionResult, error) {
	var rv LookupDevBoxDefinitionResult
	err := ctx.Invoke("azure-native:devcenter/v20221111preview:getDevBoxDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDevBoxDefinitionArgs struct {
	DevBoxDefinitionName string `pulumi:"devBoxDefinitionName"`
	DevCenterName        string `pulumi:"devCenterName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupDevBoxDefinitionResult struct {
	ActiveImageReference        ImageReferenceResponse              `pulumi:"activeImageReference"`
	HibernateSupport            *string                             `pulumi:"hibernateSupport"`
	Id                          string                              `pulumi:"id"`
	ImageReference              ImageReferenceResponse              `pulumi:"imageReference"`
	ImageValidationErrorDetails ImageValidationErrorDetailsResponse `pulumi:"imageValidationErrorDetails"`
	ImageValidationStatus       string                              `pulumi:"imageValidationStatus"`
	Location                    string                              `pulumi:"location"`
	Name                        string                              `pulumi:"name"`
	OsStorageType               string                              `pulumi:"osStorageType"`
	ProvisioningState           string                              `pulumi:"provisioningState"`
	Sku                         SkuResponse                         `pulumi:"sku"`
	SystemData                  SystemDataResponse                  `pulumi:"systemData"`
	Tags                        map[string]string                   `pulumi:"tags"`
	Type                        string                              `pulumi:"type"`
}

func LookupDevBoxDefinitionOutput(ctx *pulumi.Context, args LookupDevBoxDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupDevBoxDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDevBoxDefinitionResult, error) {
			args := v.(LookupDevBoxDefinitionArgs)
			r, err := LookupDevBoxDefinition(ctx, &args, opts...)
			var s LookupDevBoxDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDevBoxDefinitionResultOutput)
}

type LookupDevBoxDefinitionOutputArgs struct {
	DevBoxDefinitionName pulumi.StringInput `pulumi:"devBoxDefinitionName"`
	DevCenterName        pulumi.StringInput `pulumi:"devCenterName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDevBoxDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDevBoxDefinitionArgs)(nil)).Elem()
}


type LookupDevBoxDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupDevBoxDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDevBoxDefinitionResult)(nil)).Elem()
}

func (o LookupDevBoxDefinitionResultOutput) ToLookupDevBoxDefinitionResultOutput() LookupDevBoxDefinitionResultOutput {
	return o
}

func (o LookupDevBoxDefinitionResultOutput) ToLookupDevBoxDefinitionResultOutputWithContext(ctx context.Context) LookupDevBoxDefinitionResultOutput {
	return o
}

func (o LookupDevBoxDefinitionResultOutput) ActiveImageReference() ImageReferenceResponseOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) ImageReferenceResponse { return v.ActiveImageReference }).(ImageReferenceResponseOutput)
}

func (o LookupDevBoxDefinitionResultOutput) HibernateSupport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) *string { return v.HibernateSupport }).(pulumi.StringPtrOutput)
}

func (o LookupDevBoxDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDevBoxDefinitionResultOutput) ImageReference() ImageReferenceResponseOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) ImageReferenceResponse { return v.ImageReference }).(ImageReferenceResponseOutput)
}

func (o LookupDevBoxDefinitionResultOutput) ImageValidationErrorDetails() ImageValidationErrorDetailsResponseOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) ImageValidationErrorDetailsResponse {
		return v.ImageValidationErrorDetails
	}).(ImageValidationErrorDetailsResponseOutput)
}

func (o LookupDevBoxDefinitionResultOutput) ImageValidationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) string { return v.ImageValidationStatus }).(pulumi.StringOutput)
}

func (o LookupDevBoxDefinitionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDevBoxDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDevBoxDefinitionResultOutput) OsStorageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) string { return v.OsStorageType }).(pulumi.StringOutput)
}

func (o LookupDevBoxDefinitionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDevBoxDefinitionResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupDevBoxDefinitionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDevBoxDefinitionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDevBoxDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDevBoxDefinitionResultOutput{})
}
