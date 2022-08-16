


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMap(ctx *pulumi.Context, args *LookupMapArgs, opts ...pulumi.InvokeOption) (*LookupMapResult, error) {
	var rv LookupMapResult
	err := ctx.Invoke("azure-native:logic/v20160601:getMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMapArgs struct {
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	MapName                string `pulumi:"mapName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupMapResult struct {
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

func LookupMapOutput(ctx *pulumi.Context, args LookupMapOutputArgs, opts ...pulumi.InvokeOption) LookupMapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMapResult, error) {
			args := v.(LookupMapArgs)
			r, err := LookupMap(ctx, &args, opts...)
			var s LookupMapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMapResultOutput)
}

type LookupMapOutputArgs struct {
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	MapName                pulumi.StringInput `pulumi:"mapName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMapArgs)(nil)).Elem()
}


type LookupMapResultOutput struct{ *pulumi.OutputState }

func (LookupMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMapResult)(nil)).Elem()
}

func (o LookupMapResultOutput) ToLookupMapResultOutput() LookupMapResultOutput {
	return o
}

func (o LookupMapResultOutput) ToLookupMapResultOutputWithContext(ctx context.Context) LookupMapResultOutput {
	return o
}

func (o LookupMapResultOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMapResult) string { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o LookupMapResultOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMapResult) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o LookupMapResultOutput) ContentLink() ContentLinkResponseOutput {
	return o.ApplyT(func(v LookupMapResult) ContentLinkResponse { return v.ContentLink }).(ContentLinkResponseOutput)
}

func (o LookupMapResultOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMapResult) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o LookupMapResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMapResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o LookupMapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMapResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMapResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMapResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupMapResultOutput) MapType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMapResult) string { return v.MapType }).(pulumi.StringOutput)
}

func (o LookupMapResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupMapResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupMapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMapResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMapResultOutput) ParametersSchema() IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput {
	return o.ApplyT(func(v LookupMapResult) *IntegrationAccountMapPropertiesResponseParametersSchema {
		return v.ParametersSchema
	}).(IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput)
}

func (o LookupMapResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMapResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMapResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMapResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMapResultOutput{})
}
