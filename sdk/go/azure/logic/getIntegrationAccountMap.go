


package logic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationAccountMap(ctx *pulumi.Context, args *LookupIntegrationAccountMapArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountMapResult, error) {
	var rv LookupIntegrationAccountMapResult
	err := ctx.Invoke("azure-native:logic:getIntegrationAccountMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountMapArgs struct {
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	MapName                string `pulumi:"mapName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupIntegrationAccountMapResult struct {
	ChangedTime      string                                                   `pulumi:"changedTime"`
	Content          *string                                                  `pulumi:"content"`
	ContentLink      ContentLinkResponse                                      `pulumi:"contentLink"`
	ContentType      *string                                                  `pulumi:"contentType"`
	CreatedTime      string                                                   `pulumi:"createdTime"`
	Id               string                                                   `pulumi:"id"`
	Location         *string                                                  `pulumi:"location"`
	MapType          string                                                   `pulumi:"mapType"`
	Metadata         interface{}                                              `pulumi:"metadata"`
	Name             string                                                   `pulumi:"name"`
	ParametersSchema *IntegrationAccountMapPropertiesResponseParametersSchema `pulumi:"parametersSchema"`
	Tags             map[string]string                                        `pulumi:"tags"`
	Type             string                                                   `pulumi:"type"`
}

func LookupIntegrationAccountMapOutput(ctx *pulumi.Context, args LookupIntegrationAccountMapOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationAccountMapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationAccountMapResult, error) {
			args := v.(LookupIntegrationAccountMapArgs)
			r, err := LookupIntegrationAccountMap(ctx, &args, opts...)
			var s LookupIntegrationAccountMapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationAccountMapResultOutput)
}

type LookupIntegrationAccountMapOutputArgs struct {
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	MapName                pulumi.StringInput `pulumi:"mapName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIntegrationAccountMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountMapArgs)(nil)).Elem()
}


type LookupIntegrationAccountMapResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationAccountMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountMapResult)(nil)).Elem()
}

func (o LookupIntegrationAccountMapResultOutput) ToLookupIntegrationAccountMapResultOutput() LookupIntegrationAccountMapResultOutput {
	return o
}

func (o LookupIntegrationAccountMapResultOutput) ToLookupIntegrationAccountMapResultOutputWithContext(ctx context.Context) LookupIntegrationAccountMapResultOutput {
	return o
}

func (o LookupIntegrationAccountMapResultOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) string { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountMapResultOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationAccountMapResultOutput) ContentLink() ContentLinkResponseOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) ContentLinkResponse { return v.ContentLink }).(ContentLinkResponseOutput)
}

func (o LookupIntegrationAccountMapResultOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationAccountMapResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountMapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountMapResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationAccountMapResultOutput) MapType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) string { return v.MapType }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountMapResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupIntegrationAccountMapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountMapResultOutput) ParametersSchema() IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) *IntegrationAccountMapPropertiesResponseParametersSchema {
		return v.ParametersSchema
	}).(IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput)
}

func (o LookupIntegrationAccountMapResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIntegrationAccountMapResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationAccountMapResultOutput{})
}
