


package v20170101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupProfile(ctx *pulumi.Context, args *LookupProfileArgs, opts ...pulumi.InvokeOption) (*LookupProfileResult, error) {
	var rv LookupProfileResult
	err := ctx.Invoke("azure-native:customerinsights/v20170101:getProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProfileArgs struct {
	HubName           string  `pulumi:"hubName"`
	LocaleCode        *string `pulumi:"localeCode"`
	ProfileName       string  `pulumi:"profileName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupProfileResult struct {
	ApiEntitySetName    *string                      `pulumi:"apiEntitySetName"`
	Attributes          map[string][]string          `pulumi:"attributes"`
	Description         map[string]string            `pulumi:"description"`
	DisplayName         map[string]string            `pulumi:"displayName"`
	EntityType          *string                      `pulumi:"entityType"`
	Fields              []PropertyDefinitionResponse `pulumi:"fields"`
	Id                  string                       `pulumi:"id"`
	InstancesCount      *int                         `pulumi:"instancesCount"`
	LargeImage          *string                      `pulumi:"largeImage"`
	LastChangedUtc      string                       `pulumi:"lastChangedUtc"`
	LocalizedAttributes map[string]map[string]string `pulumi:"localizedAttributes"`
	MediumImage         *string                      `pulumi:"mediumImage"`
	Name                string                       `pulumi:"name"`
	ProvisioningState   string                       `pulumi:"provisioningState"`
	SchemaItemTypeLink  *string                      `pulumi:"schemaItemTypeLink"`
	SmallImage          *string                      `pulumi:"smallImage"`
	StrongIds           []StrongIdResponse           `pulumi:"strongIds"`
	TenantId            string                       `pulumi:"tenantId"`
	TimestampFieldName  *string                      `pulumi:"timestampFieldName"`
	Type                string                       `pulumi:"type"`
	TypeName            *string                      `pulumi:"typeName"`
}

func LookupProfileOutput(ctx *pulumi.Context, args LookupProfileOutputArgs, opts ...pulumi.InvokeOption) LookupProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProfileResult, error) {
			args := v.(LookupProfileArgs)
			r, err := LookupProfile(ctx, &args, opts...)
			var s LookupProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProfileResultOutput)
}

type LookupProfileOutputArgs struct {
	HubName           pulumi.StringInput    `pulumi:"hubName"`
	LocaleCode        pulumi.StringPtrInput `pulumi:"localeCode"`
	ProfileName       pulumi.StringInput    `pulumi:"profileName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileArgs)(nil)).Elem()
}


type LookupProfileResultOutput struct{ *pulumi.OutputState }

func (LookupProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileResult)(nil)).Elem()
}

func (o LookupProfileResultOutput) ToLookupProfileResultOutput() LookupProfileResultOutput {
	return o
}

func (o LookupProfileResultOutput) ToLookupProfileResultOutputWithContext(ctx context.Context) LookupProfileResultOutput {
	return o
}

func (o LookupProfileResultOutput) ApiEntitySetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.ApiEntitySetName }).(pulumi.StringPtrOutput)
}

func (o LookupProfileResultOutput) Attributes() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v LookupProfileResult) map[string][]string { return v.Attributes }).(pulumi.StringArrayMapOutput)
}

func (o LookupProfileResultOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProfileResult) map[string]string { return v.Description }).(pulumi.StringMapOutput)
}

func (o LookupProfileResultOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProfileResult) map[string]string { return v.DisplayName }).(pulumi.StringMapOutput)
}

func (o LookupProfileResultOutput) EntityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.EntityType }).(pulumi.StringPtrOutput)
}

func (o LookupProfileResultOutput) Fields() PropertyDefinitionResponseArrayOutput {
	return o.ApplyT(func(v LookupProfileResult) []PropertyDefinitionResponse { return v.Fields }).(PropertyDefinitionResponseArrayOutput)
}

func (o LookupProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProfileResultOutput) InstancesCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *int { return v.InstancesCount }).(pulumi.IntPtrOutput)
}

func (o LookupProfileResultOutput) LargeImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.LargeImage }).(pulumi.StringPtrOutput)
}

func (o LookupProfileResultOutput) LastChangedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.LastChangedUtc }).(pulumi.StringOutput)
}

func (o LookupProfileResultOutput) LocalizedAttributes() pulumi.StringMapMapOutput {
	return o.ApplyT(func(v LookupProfileResult) map[string]map[string]string { return v.LocalizedAttributes }).(pulumi.StringMapMapOutput)
}

func (o LookupProfileResultOutput) MediumImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.MediumImage }).(pulumi.StringPtrOutput)
}

func (o LookupProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProfileResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupProfileResultOutput) SchemaItemTypeLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.SchemaItemTypeLink }).(pulumi.StringPtrOutput)
}

func (o LookupProfileResultOutput) SmallImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.SmallImage }).(pulumi.StringPtrOutput)
}

func (o LookupProfileResultOutput) StrongIds() StrongIdResponseArrayOutput {
	return o.ApplyT(func(v LookupProfileResult) []StrongIdResponse { return v.StrongIds }).(StrongIdResponseArrayOutput)
}

func (o LookupProfileResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupProfileResultOutput) TimestampFieldName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.TimestampFieldName }).(pulumi.StringPtrOutput)
}

func (o LookupProfileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupProfileResultOutput) TypeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.TypeName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProfileResultOutput{})
}
