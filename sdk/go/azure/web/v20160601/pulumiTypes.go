


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiConnectionDefinitionProperties struct {
	Api                      *ApiReference                `pulumi:"api"`
	ChangedTime              *string                      `pulumi:"changedTime"`
	CreatedTime              *string                      `pulumi:"createdTime"`
	CustomParameterValues    map[string]string            `pulumi:"customParameterValues"`
	DisplayName              *string                      `pulumi:"displayName"`
	NonSecretParameterValues map[string]string            `pulumi:"nonSecretParameterValues"`
	ParameterValues          map[string]string            `pulumi:"parameterValues"`
	Statuses                 []ConnectionStatusDefinition `pulumi:"statuses"`
	TestLinks                []ApiConnectionTestLink      `pulumi:"testLinks"`
}





type ApiConnectionDefinitionPropertiesInput interface {
	pulumi.Input

	ToApiConnectionDefinitionPropertiesOutput() ApiConnectionDefinitionPropertiesOutput
	ToApiConnectionDefinitionPropertiesOutputWithContext(context.Context) ApiConnectionDefinitionPropertiesOutput
}

type ApiConnectionDefinitionPropertiesArgs struct {
	Api                      ApiReferencePtrInput                 `pulumi:"api"`
	ChangedTime              pulumi.StringPtrInput                `pulumi:"changedTime"`
	CreatedTime              pulumi.StringPtrInput                `pulumi:"createdTime"`
	CustomParameterValues    pulumi.StringMapInput                `pulumi:"customParameterValues"`
	DisplayName              pulumi.StringPtrInput                `pulumi:"displayName"`
	NonSecretParameterValues pulumi.StringMapInput                `pulumi:"nonSecretParameterValues"`
	ParameterValues          pulumi.StringMapInput                `pulumi:"parameterValues"`
	Statuses                 ConnectionStatusDefinitionArrayInput `pulumi:"statuses"`
	TestLinks                ApiConnectionTestLinkArrayInput      `pulumi:"testLinks"`
}

func (ApiConnectionDefinitionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiConnectionDefinitionProperties)(nil)).Elem()
}

func (i ApiConnectionDefinitionPropertiesArgs) ToApiConnectionDefinitionPropertiesOutput() ApiConnectionDefinitionPropertiesOutput {
	return i.ToApiConnectionDefinitionPropertiesOutputWithContext(context.Background())
}

func (i ApiConnectionDefinitionPropertiesArgs) ToApiConnectionDefinitionPropertiesOutputWithContext(ctx context.Context) ApiConnectionDefinitionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConnectionDefinitionPropertiesOutput)
}

func (i ApiConnectionDefinitionPropertiesArgs) ToApiConnectionDefinitionPropertiesPtrOutput() ApiConnectionDefinitionPropertiesPtrOutput {
	return i.ToApiConnectionDefinitionPropertiesPtrOutputWithContext(context.Background())
}

func (i ApiConnectionDefinitionPropertiesArgs) ToApiConnectionDefinitionPropertiesPtrOutputWithContext(ctx context.Context) ApiConnectionDefinitionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConnectionDefinitionPropertiesOutput).ToApiConnectionDefinitionPropertiesPtrOutputWithContext(ctx)
}









type ApiConnectionDefinitionPropertiesPtrInput interface {
	pulumi.Input

	ToApiConnectionDefinitionPropertiesPtrOutput() ApiConnectionDefinitionPropertiesPtrOutput
	ToApiConnectionDefinitionPropertiesPtrOutputWithContext(context.Context) ApiConnectionDefinitionPropertiesPtrOutput
}

type apiConnectionDefinitionPropertiesPtrType ApiConnectionDefinitionPropertiesArgs

func ApiConnectionDefinitionPropertiesPtr(v *ApiConnectionDefinitionPropertiesArgs) ApiConnectionDefinitionPropertiesPtrInput {
	return (*apiConnectionDefinitionPropertiesPtrType)(v)
}

func (*apiConnectionDefinitionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiConnectionDefinitionProperties)(nil)).Elem()
}

func (i *apiConnectionDefinitionPropertiesPtrType) ToApiConnectionDefinitionPropertiesPtrOutput() ApiConnectionDefinitionPropertiesPtrOutput {
	return i.ToApiConnectionDefinitionPropertiesPtrOutputWithContext(context.Background())
}

func (i *apiConnectionDefinitionPropertiesPtrType) ToApiConnectionDefinitionPropertiesPtrOutputWithContext(ctx context.Context) ApiConnectionDefinitionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConnectionDefinitionPropertiesPtrOutput)
}

type ApiConnectionDefinitionPropertiesOutput struct{ *pulumi.OutputState }

func (ApiConnectionDefinitionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiConnectionDefinitionProperties)(nil)).Elem()
}

func (o ApiConnectionDefinitionPropertiesOutput) ToApiConnectionDefinitionPropertiesOutput() ApiConnectionDefinitionPropertiesOutput {
	return o
}

func (o ApiConnectionDefinitionPropertiesOutput) ToApiConnectionDefinitionPropertiesOutputWithContext(ctx context.Context) ApiConnectionDefinitionPropertiesOutput {
	return o
}

func (o ApiConnectionDefinitionPropertiesOutput) ToApiConnectionDefinitionPropertiesPtrOutput() ApiConnectionDefinitionPropertiesPtrOutput {
	return o.ToApiConnectionDefinitionPropertiesPtrOutputWithContext(context.Background())
}

func (o ApiConnectionDefinitionPropertiesOutput) ToApiConnectionDefinitionPropertiesPtrOutputWithContext(ctx context.Context) ApiConnectionDefinitionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiConnectionDefinitionProperties) *ApiConnectionDefinitionProperties {
		return &v
	}).(ApiConnectionDefinitionPropertiesPtrOutput)
}

func (o ApiConnectionDefinitionPropertiesOutput) Api() ApiReferencePtrOutput {
	return o.ApplyT(func(v ApiConnectionDefinitionProperties) *ApiReference { return v.Api }).(ApiReferencePtrOutput)
}

func (o ApiConnectionDefinitionPropertiesOutput) ChangedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiConnectionDefinitionProperties) *string { return v.ChangedTime }).(pulumi.StringPtrOutput)
}

func (o ApiConnectionDefinitionPropertiesOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiConnectionDefinitionProperties) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o ApiConnectionDefinitionPropertiesOutput) CustomParameterValues() pulumi.StringMapOutput {
	return o.ApplyT(func(v ApiConnectionDefinitionProperties) map[string]string { return v.CustomParameterValues }).(pulumi.StringMapOutput)
}

func (o ApiConnectionDefinitionPropertiesOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiConnectionDefinitionProperties) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ApiConnectionDefinitionPropertiesOutput) NonSecretParameterValues() pulumi.StringMapOutput {
	return o.ApplyT(func(v ApiConnectionDefinitionProperties) map[string]string { return v.NonSecretParameterValues }).(pulumi.StringMapOutput)
}

func (o ApiConnectionDefinitionPropertiesOutput) ParameterValues() pulumi.StringMapOutput {
	return o.ApplyT(func(v ApiConnectionDefinitionProperties) map[string]string { return v.ParameterValues }).(pulumi.StringMapOutput)
}

func (o ApiConnectionDefinitionPropertiesOutput) Statuses() ConnectionStatusDefinitionArrayOutput {
	return o.ApplyT(func(v ApiConnectionDefinitionProperties) []ConnectionStatusDefinition { return v.Statuses }).(ConnectionStatusDefinitionArrayOutput)
}

func (o ApiConnectionDefinitionPropertiesOutput) TestLinks() ApiConnectionTestLinkArrayOutput {
	return o.ApplyT(func(v ApiConnectionDefinitionProperties) []ApiConnectionTestLink { return v.TestLinks }).(ApiConnectionTestLinkArrayOutput)
}

type ApiConnectionDefinitionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ApiConnectionDefinitionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiConnectionDefinitionProperties)(nil)).Elem()
}

func (o ApiConnectionDefinitionPropertiesPtrOutput) ToApiConnectionDefinitionPropertiesPtrOutput() ApiConnectionDefinitionPropertiesPtrOutput {
	return o
}

func (o ApiConnectionDefinitionPropertiesPtrOutput) ToApiConnectionDefinitionPropertiesPtrOutputWithContext(ctx context.Context) ApiConnectionDefinitionPropertiesPtrOutput {
	return o
}

func (o ApiConnectionDefinitionPropertiesPtrOutput) Elem() ApiConnectionDefinitionPropertiesOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionProperties) ApiConnectionDefinitionProperties {
		if v != nil {
			return *v
		}
		var ret ApiConnectionDefinitionProperties
		return ret
	}).(ApiConnectionDefinitionPropertiesOutput)
}

func (o ApiConnectionDefinitionPropertiesPtrOutput) Api() ApiReferencePtrOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionProperties) *ApiReference {
		if v == nil {
			return nil
		}
		return v.Api
	}).(ApiReferencePtrOutput)
}

func (o ApiConnectionDefinitionPropertiesPtrOutput) ChangedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionProperties) *string {
		if v == nil {
			return nil
		}
		return v.ChangedTime
	}).(pulumi.StringPtrOutput)
}

func (o ApiConnectionDefinitionPropertiesPtrOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionProperties) *string {
		if v == nil {
			return nil
		}
		return v.CreatedTime
	}).(pulumi.StringPtrOutput)
}

func (o ApiConnectionDefinitionPropertiesPtrOutput) CustomParameterValues() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.CustomParameterValues
	}).(pulumi.StringMapOutput)
}

func (o ApiConnectionDefinitionPropertiesPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionProperties) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ApiConnectionDefinitionPropertiesPtrOutput) NonSecretParameterValues() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.NonSecretParameterValues
	}).(pulumi.StringMapOutput)
}

func (o ApiConnectionDefinitionPropertiesPtrOutput) ParameterValues() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.ParameterValues
	}).(pulumi.StringMapOutput)
}

func (o ApiConnectionDefinitionPropertiesPtrOutput) Statuses() ConnectionStatusDefinitionArrayOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionProperties) []ConnectionStatusDefinition {
		if v == nil {
			return nil
		}
		return v.Statuses
	}).(ConnectionStatusDefinitionArrayOutput)
}

func (o ApiConnectionDefinitionPropertiesPtrOutput) TestLinks() ApiConnectionTestLinkArrayOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionProperties) []ApiConnectionTestLink {
		if v == nil {
			return nil
		}
		return v.TestLinks
	}).(ApiConnectionTestLinkArrayOutput)
}

type ApiConnectionDefinitionResponseProperties struct {
	Api                      *ApiReferenceResponse                `pulumi:"api"`
	ChangedTime              *string                              `pulumi:"changedTime"`
	CreatedTime              *string                              `pulumi:"createdTime"`
	CustomParameterValues    map[string]string                    `pulumi:"customParameterValues"`
	DisplayName              *string                              `pulumi:"displayName"`
	NonSecretParameterValues map[string]string                    `pulumi:"nonSecretParameterValues"`
	ParameterValues          map[string]string                    `pulumi:"parameterValues"`
	Statuses                 []ConnectionStatusDefinitionResponse `pulumi:"statuses"`
	TestLinks                []ApiConnectionTestLinkResponse      `pulumi:"testLinks"`
}





type ApiConnectionDefinitionResponsePropertiesInput interface {
	pulumi.Input

	ToApiConnectionDefinitionResponsePropertiesOutput() ApiConnectionDefinitionResponsePropertiesOutput
	ToApiConnectionDefinitionResponsePropertiesOutputWithContext(context.Context) ApiConnectionDefinitionResponsePropertiesOutput
}

type ApiConnectionDefinitionResponsePropertiesArgs struct {
	Api                      ApiReferenceResponsePtrInput                 `pulumi:"api"`
	ChangedTime              pulumi.StringPtrInput                        `pulumi:"changedTime"`
	CreatedTime              pulumi.StringPtrInput                        `pulumi:"createdTime"`
	CustomParameterValues    pulumi.StringMapInput                        `pulumi:"customParameterValues"`
	DisplayName              pulumi.StringPtrInput                        `pulumi:"displayName"`
	NonSecretParameterValues pulumi.StringMapInput                        `pulumi:"nonSecretParameterValues"`
	ParameterValues          pulumi.StringMapInput                        `pulumi:"parameterValues"`
	Statuses                 ConnectionStatusDefinitionResponseArrayInput `pulumi:"statuses"`
	TestLinks                ApiConnectionTestLinkResponseArrayInput      `pulumi:"testLinks"`
}

func (ApiConnectionDefinitionResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiConnectionDefinitionResponseProperties)(nil)).Elem()
}

func (i ApiConnectionDefinitionResponsePropertiesArgs) ToApiConnectionDefinitionResponsePropertiesOutput() ApiConnectionDefinitionResponsePropertiesOutput {
	return i.ToApiConnectionDefinitionResponsePropertiesOutputWithContext(context.Background())
}

func (i ApiConnectionDefinitionResponsePropertiesArgs) ToApiConnectionDefinitionResponsePropertiesOutputWithContext(ctx context.Context) ApiConnectionDefinitionResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConnectionDefinitionResponsePropertiesOutput)
}

func (i ApiConnectionDefinitionResponsePropertiesArgs) ToApiConnectionDefinitionResponsePropertiesPtrOutput() ApiConnectionDefinitionResponsePropertiesPtrOutput {
	return i.ToApiConnectionDefinitionResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i ApiConnectionDefinitionResponsePropertiesArgs) ToApiConnectionDefinitionResponsePropertiesPtrOutputWithContext(ctx context.Context) ApiConnectionDefinitionResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConnectionDefinitionResponsePropertiesOutput).ToApiConnectionDefinitionResponsePropertiesPtrOutputWithContext(ctx)
}









type ApiConnectionDefinitionResponsePropertiesPtrInput interface {
	pulumi.Input

	ToApiConnectionDefinitionResponsePropertiesPtrOutput() ApiConnectionDefinitionResponsePropertiesPtrOutput
	ToApiConnectionDefinitionResponsePropertiesPtrOutputWithContext(context.Context) ApiConnectionDefinitionResponsePropertiesPtrOutput
}

type apiConnectionDefinitionResponsePropertiesPtrType ApiConnectionDefinitionResponsePropertiesArgs

func ApiConnectionDefinitionResponsePropertiesPtr(v *ApiConnectionDefinitionResponsePropertiesArgs) ApiConnectionDefinitionResponsePropertiesPtrInput {
	return (*apiConnectionDefinitionResponsePropertiesPtrType)(v)
}

func (*apiConnectionDefinitionResponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiConnectionDefinitionResponseProperties)(nil)).Elem()
}

func (i *apiConnectionDefinitionResponsePropertiesPtrType) ToApiConnectionDefinitionResponsePropertiesPtrOutput() ApiConnectionDefinitionResponsePropertiesPtrOutput {
	return i.ToApiConnectionDefinitionResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *apiConnectionDefinitionResponsePropertiesPtrType) ToApiConnectionDefinitionResponsePropertiesPtrOutputWithContext(ctx context.Context) ApiConnectionDefinitionResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConnectionDefinitionResponsePropertiesPtrOutput)
}

type ApiConnectionDefinitionResponsePropertiesOutput struct{ *pulumi.OutputState }

func (ApiConnectionDefinitionResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiConnectionDefinitionResponseProperties)(nil)).Elem()
}

func (o ApiConnectionDefinitionResponsePropertiesOutput) ToApiConnectionDefinitionResponsePropertiesOutput() ApiConnectionDefinitionResponsePropertiesOutput {
	return o
}

func (o ApiConnectionDefinitionResponsePropertiesOutput) ToApiConnectionDefinitionResponsePropertiesOutputWithContext(ctx context.Context) ApiConnectionDefinitionResponsePropertiesOutput {
	return o
}

func (o ApiConnectionDefinitionResponsePropertiesOutput) ToApiConnectionDefinitionResponsePropertiesPtrOutput() ApiConnectionDefinitionResponsePropertiesPtrOutput {
	return o.ToApiConnectionDefinitionResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o ApiConnectionDefinitionResponsePropertiesOutput) ToApiConnectionDefinitionResponsePropertiesPtrOutputWithContext(ctx context.Context) ApiConnectionDefinitionResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiConnectionDefinitionResponseProperties) *ApiConnectionDefinitionResponseProperties {
		return &v
	}).(ApiConnectionDefinitionResponsePropertiesPtrOutput)
}

func (o ApiConnectionDefinitionResponsePropertiesOutput) Api() ApiReferenceResponsePtrOutput {
	return o.ApplyT(func(v ApiConnectionDefinitionResponseProperties) *ApiReferenceResponse { return v.Api }).(ApiReferenceResponsePtrOutput)
}

func (o ApiConnectionDefinitionResponsePropertiesOutput) ChangedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiConnectionDefinitionResponseProperties) *string { return v.ChangedTime }).(pulumi.StringPtrOutput)
}

func (o ApiConnectionDefinitionResponsePropertiesOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiConnectionDefinitionResponseProperties) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o ApiConnectionDefinitionResponsePropertiesOutput) CustomParameterValues() pulumi.StringMapOutput {
	return o.ApplyT(func(v ApiConnectionDefinitionResponseProperties) map[string]string { return v.CustomParameterValues }).(pulumi.StringMapOutput)
}

func (o ApiConnectionDefinitionResponsePropertiesOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiConnectionDefinitionResponseProperties) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ApiConnectionDefinitionResponsePropertiesOutput) NonSecretParameterValues() pulumi.StringMapOutput {
	return o.ApplyT(func(v ApiConnectionDefinitionResponseProperties) map[string]string { return v.NonSecretParameterValues }).(pulumi.StringMapOutput)
}

func (o ApiConnectionDefinitionResponsePropertiesOutput) ParameterValues() pulumi.StringMapOutput {
	return o.ApplyT(func(v ApiConnectionDefinitionResponseProperties) map[string]string { return v.ParameterValues }).(pulumi.StringMapOutput)
}

func (o ApiConnectionDefinitionResponsePropertiesOutput) Statuses() ConnectionStatusDefinitionResponseArrayOutput {
	return o.ApplyT(func(v ApiConnectionDefinitionResponseProperties) []ConnectionStatusDefinitionResponse {
		return v.Statuses
	}).(ConnectionStatusDefinitionResponseArrayOutput)
}

func (o ApiConnectionDefinitionResponsePropertiesOutput) TestLinks() ApiConnectionTestLinkResponseArrayOutput {
	return o.ApplyT(func(v ApiConnectionDefinitionResponseProperties) []ApiConnectionTestLinkResponse { return v.TestLinks }).(ApiConnectionTestLinkResponseArrayOutput)
}

type ApiConnectionDefinitionResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ApiConnectionDefinitionResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiConnectionDefinitionResponseProperties)(nil)).Elem()
}

func (o ApiConnectionDefinitionResponsePropertiesPtrOutput) ToApiConnectionDefinitionResponsePropertiesPtrOutput() ApiConnectionDefinitionResponsePropertiesPtrOutput {
	return o
}

func (o ApiConnectionDefinitionResponsePropertiesPtrOutput) ToApiConnectionDefinitionResponsePropertiesPtrOutputWithContext(ctx context.Context) ApiConnectionDefinitionResponsePropertiesPtrOutput {
	return o
}

func (o ApiConnectionDefinitionResponsePropertiesPtrOutput) Elem() ApiConnectionDefinitionResponsePropertiesOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionResponseProperties) ApiConnectionDefinitionResponseProperties {
		if v != nil {
			return *v
		}
		var ret ApiConnectionDefinitionResponseProperties
		return ret
	}).(ApiConnectionDefinitionResponsePropertiesOutput)
}

func (o ApiConnectionDefinitionResponsePropertiesPtrOutput) Api() ApiReferenceResponsePtrOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionResponseProperties) *ApiReferenceResponse {
		if v == nil {
			return nil
		}
		return v.Api
	}).(ApiReferenceResponsePtrOutput)
}

func (o ApiConnectionDefinitionResponsePropertiesPtrOutput) ChangedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.ChangedTime
	}).(pulumi.StringPtrOutput)
}

func (o ApiConnectionDefinitionResponsePropertiesPtrOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.CreatedTime
	}).(pulumi.StringPtrOutput)
}

func (o ApiConnectionDefinitionResponsePropertiesPtrOutput) CustomParameterValues() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionResponseProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.CustomParameterValues
	}).(pulumi.StringMapOutput)
}

func (o ApiConnectionDefinitionResponsePropertiesPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ApiConnectionDefinitionResponsePropertiesPtrOutput) NonSecretParameterValues() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionResponseProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.NonSecretParameterValues
	}).(pulumi.StringMapOutput)
}

func (o ApiConnectionDefinitionResponsePropertiesPtrOutput) ParameterValues() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionResponseProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.ParameterValues
	}).(pulumi.StringMapOutput)
}

func (o ApiConnectionDefinitionResponsePropertiesPtrOutput) Statuses() ConnectionStatusDefinitionResponseArrayOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionResponseProperties) []ConnectionStatusDefinitionResponse {
		if v == nil {
			return nil
		}
		return v.Statuses
	}).(ConnectionStatusDefinitionResponseArrayOutput)
}

func (o ApiConnectionDefinitionResponsePropertiesPtrOutput) TestLinks() ApiConnectionTestLinkResponseArrayOutput {
	return o.ApplyT(func(v *ApiConnectionDefinitionResponseProperties) []ApiConnectionTestLinkResponse {
		if v == nil {
			return nil
		}
		return v.TestLinks
	}).(ApiConnectionTestLinkResponseArrayOutput)
}

type ApiConnectionTestLink struct {
	Method     *string `pulumi:"method"`
	RequestUri *string `pulumi:"requestUri"`
}





type ApiConnectionTestLinkInput interface {
	pulumi.Input

	ToApiConnectionTestLinkOutput() ApiConnectionTestLinkOutput
	ToApiConnectionTestLinkOutputWithContext(context.Context) ApiConnectionTestLinkOutput
}

type ApiConnectionTestLinkArgs struct {
	Method     pulumi.StringPtrInput `pulumi:"method"`
	RequestUri pulumi.StringPtrInput `pulumi:"requestUri"`
}

func (ApiConnectionTestLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiConnectionTestLink)(nil)).Elem()
}

func (i ApiConnectionTestLinkArgs) ToApiConnectionTestLinkOutput() ApiConnectionTestLinkOutput {
	return i.ToApiConnectionTestLinkOutputWithContext(context.Background())
}

func (i ApiConnectionTestLinkArgs) ToApiConnectionTestLinkOutputWithContext(ctx context.Context) ApiConnectionTestLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConnectionTestLinkOutput)
}





type ApiConnectionTestLinkArrayInput interface {
	pulumi.Input

	ToApiConnectionTestLinkArrayOutput() ApiConnectionTestLinkArrayOutput
	ToApiConnectionTestLinkArrayOutputWithContext(context.Context) ApiConnectionTestLinkArrayOutput
}

type ApiConnectionTestLinkArray []ApiConnectionTestLinkInput

func (ApiConnectionTestLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiConnectionTestLink)(nil)).Elem()
}

func (i ApiConnectionTestLinkArray) ToApiConnectionTestLinkArrayOutput() ApiConnectionTestLinkArrayOutput {
	return i.ToApiConnectionTestLinkArrayOutputWithContext(context.Background())
}

func (i ApiConnectionTestLinkArray) ToApiConnectionTestLinkArrayOutputWithContext(ctx context.Context) ApiConnectionTestLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConnectionTestLinkArrayOutput)
}

type ApiConnectionTestLinkOutput struct{ *pulumi.OutputState }

func (ApiConnectionTestLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiConnectionTestLink)(nil)).Elem()
}

func (o ApiConnectionTestLinkOutput) ToApiConnectionTestLinkOutput() ApiConnectionTestLinkOutput {
	return o
}

func (o ApiConnectionTestLinkOutput) ToApiConnectionTestLinkOutputWithContext(ctx context.Context) ApiConnectionTestLinkOutput {
	return o
}

func (o ApiConnectionTestLinkOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiConnectionTestLink) *string { return v.Method }).(pulumi.StringPtrOutput)
}

func (o ApiConnectionTestLinkOutput) RequestUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiConnectionTestLink) *string { return v.RequestUri }).(pulumi.StringPtrOutput)
}

type ApiConnectionTestLinkArrayOutput struct{ *pulumi.OutputState }

func (ApiConnectionTestLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiConnectionTestLink)(nil)).Elem()
}

func (o ApiConnectionTestLinkArrayOutput) ToApiConnectionTestLinkArrayOutput() ApiConnectionTestLinkArrayOutput {
	return o
}

func (o ApiConnectionTestLinkArrayOutput) ToApiConnectionTestLinkArrayOutputWithContext(ctx context.Context) ApiConnectionTestLinkArrayOutput {
	return o
}

func (o ApiConnectionTestLinkArrayOutput) Index(i pulumi.IntInput) ApiConnectionTestLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiConnectionTestLink {
		return vs[0].([]ApiConnectionTestLink)[vs[1].(int)]
	}).(ApiConnectionTestLinkOutput)
}

type ApiConnectionTestLinkResponse struct {
	Method     *string `pulumi:"method"`
	RequestUri *string `pulumi:"requestUri"`
}





type ApiConnectionTestLinkResponseInput interface {
	pulumi.Input

	ToApiConnectionTestLinkResponseOutput() ApiConnectionTestLinkResponseOutput
	ToApiConnectionTestLinkResponseOutputWithContext(context.Context) ApiConnectionTestLinkResponseOutput
}

type ApiConnectionTestLinkResponseArgs struct {
	Method     pulumi.StringPtrInput `pulumi:"method"`
	RequestUri pulumi.StringPtrInput `pulumi:"requestUri"`
}

func (ApiConnectionTestLinkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiConnectionTestLinkResponse)(nil)).Elem()
}

func (i ApiConnectionTestLinkResponseArgs) ToApiConnectionTestLinkResponseOutput() ApiConnectionTestLinkResponseOutput {
	return i.ToApiConnectionTestLinkResponseOutputWithContext(context.Background())
}

func (i ApiConnectionTestLinkResponseArgs) ToApiConnectionTestLinkResponseOutputWithContext(ctx context.Context) ApiConnectionTestLinkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConnectionTestLinkResponseOutput)
}





type ApiConnectionTestLinkResponseArrayInput interface {
	pulumi.Input

	ToApiConnectionTestLinkResponseArrayOutput() ApiConnectionTestLinkResponseArrayOutput
	ToApiConnectionTestLinkResponseArrayOutputWithContext(context.Context) ApiConnectionTestLinkResponseArrayOutput
}

type ApiConnectionTestLinkResponseArray []ApiConnectionTestLinkResponseInput

func (ApiConnectionTestLinkResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiConnectionTestLinkResponse)(nil)).Elem()
}

func (i ApiConnectionTestLinkResponseArray) ToApiConnectionTestLinkResponseArrayOutput() ApiConnectionTestLinkResponseArrayOutput {
	return i.ToApiConnectionTestLinkResponseArrayOutputWithContext(context.Background())
}

func (i ApiConnectionTestLinkResponseArray) ToApiConnectionTestLinkResponseArrayOutputWithContext(ctx context.Context) ApiConnectionTestLinkResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConnectionTestLinkResponseArrayOutput)
}

type ApiConnectionTestLinkResponseOutput struct{ *pulumi.OutputState }

func (ApiConnectionTestLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiConnectionTestLinkResponse)(nil)).Elem()
}

func (o ApiConnectionTestLinkResponseOutput) ToApiConnectionTestLinkResponseOutput() ApiConnectionTestLinkResponseOutput {
	return o
}

func (o ApiConnectionTestLinkResponseOutput) ToApiConnectionTestLinkResponseOutputWithContext(ctx context.Context) ApiConnectionTestLinkResponseOutput {
	return o
}

func (o ApiConnectionTestLinkResponseOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiConnectionTestLinkResponse) *string { return v.Method }).(pulumi.StringPtrOutput)
}

func (o ApiConnectionTestLinkResponseOutput) RequestUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiConnectionTestLinkResponse) *string { return v.RequestUri }).(pulumi.StringPtrOutput)
}

type ApiConnectionTestLinkResponseArrayOutput struct{ *pulumi.OutputState }

func (ApiConnectionTestLinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiConnectionTestLinkResponse)(nil)).Elem()
}

func (o ApiConnectionTestLinkResponseArrayOutput) ToApiConnectionTestLinkResponseArrayOutput() ApiConnectionTestLinkResponseArrayOutput {
	return o
}

func (o ApiConnectionTestLinkResponseArrayOutput) ToApiConnectionTestLinkResponseArrayOutputWithContext(ctx context.Context) ApiConnectionTestLinkResponseArrayOutput {
	return o
}

func (o ApiConnectionTestLinkResponseArrayOutput) Index(i pulumi.IntInput) ApiConnectionTestLinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiConnectionTestLinkResponse {
		return vs[0].([]ApiConnectionTestLinkResponse)[vs[1].(int)]
	}).(ApiConnectionTestLinkResponseOutput)
}

type ApiOAuthSettings struct {
	ClientId         *string                              `pulumi:"clientId"`
	ClientSecret     *string                              `pulumi:"clientSecret"`
	CustomParameters map[string]ApiOAuthSettingsParameter `pulumi:"customParameters"`
	IdentityProvider *string                              `pulumi:"identityProvider"`
	Properties       interface{}                          `pulumi:"properties"`
	RedirectUrl      *string                              `pulumi:"redirectUrl"`
	Scopes           []string                             `pulumi:"scopes"`
}





type ApiOAuthSettingsInput interface {
	pulumi.Input

	ToApiOAuthSettingsOutput() ApiOAuthSettingsOutput
	ToApiOAuthSettingsOutputWithContext(context.Context) ApiOAuthSettingsOutput
}

type ApiOAuthSettingsArgs struct {
	ClientId         pulumi.StringPtrInput             `pulumi:"clientId"`
	ClientSecret     pulumi.StringPtrInput             `pulumi:"clientSecret"`
	CustomParameters ApiOAuthSettingsParameterMapInput `pulumi:"customParameters"`
	IdentityProvider pulumi.StringPtrInput             `pulumi:"identityProvider"`
	Properties       pulumi.Input                      `pulumi:"properties"`
	RedirectUrl      pulumi.StringPtrInput             `pulumi:"redirectUrl"`
	Scopes           pulumi.StringArrayInput           `pulumi:"scopes"`
}

func (ApiOAuthSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOAuthSettings)(nil)).Elem()
}

func (i ApiOAuthSettingsArgs) ToApiOAuthSettingsOutput() ApiOAuthSettingsOutput {
	return i.ToApiOAuthSettingsOutputWithContext(context.Background())
}

func (i ApiOAuthSettingsArgs) ToApiOAuthSettingsOutputWithContext(ctx context.Context) ApiOAuthSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsOutput)
}

func (i ApiOAuthSettingsArgs) ToApiOAuthSettingsPtrOutput() ApiOAuthSettingsPtrOutput {
	return i.ToApiOAuthSettingsPtrOutputWithContext(context.Background())
}

func (i ApiOAuthSettingsArgs) ToApiOAuthSettingsPtrOutputWithContext(ctx context.Context) ApiOAuthSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsOutput).ToApiOAuthSettingsPtrOutputWithContext(ctx)
}









type ApiOAuthSettingsPtrInput interface {
	pulumi.Input

	ToApiOAuthSettingsPtrOutput() ApiOAuthSettingsPtrOutput
	ToApiOAuthSettingsPtrOutputWithContext(context.Context) ApiOAuthSettingsPtrOutput
}

type apiOAuthSettingsPtrType ApiOAuthSettingsArgs

func ApiOAuthSettingsPtr(v *ApiOAuthSettingsArgs) ApiOAuthSettingsPtrInput {
	return (*apiOAuthSettingsPtrType)(v)
}

func (*apiOAuthSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOAuthSettings)(nil)).Elem()
}

func (i *apiOAuthSettingsPtrType) ToApiOAuthSettingsPtrOutput() ApiOAuthSettingsPtrOutput {
	return i.ToApiOAuthSettingsPtrOutputWithContext(context.Background())
}

func (i *apiOAuthSettingsPtrType) ToApiOAuthSettingsPtrOutputWithContext(ctx context.Context) ApiOAuthSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsPtrOutput)
}

type ApiOAuthSettingsOutput struct{ *pulumi.OutputState }

func (ApiOAuthSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOAuthSettings)(nil)).Elem()
}

func (o ApiOAuthSettingsOutput) ToApiOAuthSettingsOutput() ApiOAuthSettingsOutput {
	return o
}

func (o ApiOAuthSettingsOutput) ToApiOAuthSettingsOutputWithContext(ctx context.Context) ApiOAuthSettingsOutput {
	return o
}

func (o ApiOAuthSettingsOutput) ToApiOAuthSettingsPtrOutput() ApiOAuthSettingsPtrOutput {
	return o.ToApiOAuthSettingsPtrOutputWithContext(context.Background())
}

func (o ApiOAuthSettingsOutput) ToApiOAuthSettingsPtrOutputWithContext(ctx context.Context) ApiOAuthSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiOAuthSettings) *ApiOAuthSettings {
		return &v
	}).(ApiOAuthSettingsPtrOutput)
}

func (o ApiOAuthSettingsOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettings) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettings) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsOutput) CustomParameters() ApiOAuthSettingsParameterMapOutput {
	return o.ApplyT(func(v ApiOAuthSettings) map[string]ApiOAuthSettingsParameter { return v.CustomParameters }).(ApiOAuthSettingsParameterMapOutput)
}

func (o ApiOAuthSettingsOutput) IdentityProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettings) *string { return v.IdentityProvider }).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v ApiOAuthSettings) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o ApiOAuthSettingsOutput) RedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettings) *string { return v.RedirectUrl }).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApiOAuthSettings) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

type ApiOAuthSettingsPtrOutput struct{ *pulumi.OutputState }

func (ApiOAuthSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOAuthSettings)(nil)).Elem()
}

func (o ApiOAuthSettingsPtrOutput) ToApiOAuthSettingsPtrOutput() ApiOAuthSettingsPtrOutput {
	return o
}

func (o ApiOAuthSettingsPtrOutput) ToApiOAuthSettingsPtrOutputWithContext(ctx context.Context) ApiOAuthSettingsPtrOutput {
	return o
}

func (o ApiOAuthSettingsPtrOutput) Elem() ApiOAuthSettingsOutput {
	return o.ApplyT(func(v *ApiOAuthSettings) ApiOAuthSettings {
		if v != nil {
			return *v
		}
		var ret ApiOAuthSettings
		return ret
	}).(ApiOAuthSettingsOutput)
}

func (o ApiOAuthSettingsPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiOAuthSettings) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsPtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiOAuthSettings) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsPtrOutput) CustomParameters() ApiOAuthSettingsParameterMapOutput {
	return o.ApplyT(func(v *ApiOAuthSettings) map[string]ApiOAuthSettingsParameter {
		if v == nil {
			return nil
		}
		return v.CustomParameters
	}).(ApiOAuthSettingsParameterMapOutput)
}

func (o ApiOAuthSettingsPtrOutput) IdentityProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiOAuthSettings) *string {
		if v == nil {
			return nil
		}
		return v.IdentityProvider
	}).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsPtrOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *ApiOAuthSettings) interface{} {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.AnyOutput)
}

func (o ApiOAuthSettingsPtrOutput) RedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiOAuthSettings) *string {
		if v == nil {
			return nil
		}
		return v.RedirectUrl
	}).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsPtrOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiOAuthSettings) []string {
		if v == nil {
			return nil
		}
		return v.Scopes
	}).(pulumi.StringArrayOutput)
}

type ApiOAuthSettingsParameter struct {
	Options      interface{} `pulumi:"options"`
	UiDefinition interface{} `pulumi:"uiDefinition"`
	Value        *string     `pulumi:"value"`
}





type ApiOAuthSettingsParameterInput interface {
	pulumi.Input

	ToApiOAuthSettingsParameterOutput() ApiOAuthSettingsParameterOutput
	ToApiOAuthSettingsParameterOutputWithContext(context.Context) ApiOAuthSettingsParameterOutput
}

type ApiOAuthSettingsParameterArgs struct {
	Options      pulumi.Input          `pulumi:"options"`
	UiDefinition pulumi.Input          `pulumi:"uiDefinition"`
	Value        pulumi.StringPtrInput `pulumi:"value"`
}

func (ApiOAuthSettingsParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOAuthSettingsParameter)(nil)).Elem()
}

func (i ApiOAuthSettingsParameterArgs) ToApiOAuthSettingsParameterOutput() ApiOAuthSettingsParameterOutput {
	return i.ToApiOAuthSettingsParameterOutputWithContext(context.Background())
}

func (i ApiOAuthSettingsParameterArgs) ToApiOAuthSettingsParameterOutputWithContext(ctx context.Context) ApiOAuthSettingsParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsParameterOutput)
}





type ApiOAuthSettingsParameterMapInput interface {
	pulumi.Input

	ToApiOAuthSettingsParameterMapOutput() ApiOAuthSettingsParameterMapOutput
	ToApiOAuthSettingsParameterMapOutputWithContext(context.Context) ApiOAuthSettingsParameterMapOutput
}

type ApiOAuthSettingsParameterMap map[string]ApiOAuthSettingsParameterInput

func (ApiOAuthSettingsParameterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiOAuthSettingsParameter)(nil)).Elem()
}

func (i ApiOAuthSettingsParameterMap) ToApiOAuthSettingsParameterMapOutput() ApiOAuthSettingsParameterMapOutput {
	return i.ToApiOAuthSettingsParameterMapOutputWithContext(context.Background())
}

func (i ApiOAuthSettingsParameterMap) ToApiOAuthSettingsParameterMapOutputWithContext(ctx context.Context) ApiOAuthSettingsParameterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsParameterMapOutput)
}

type ApiOAuthSettingsParameterOutput struct{ *pulumi.OutputState }

func (ApiOAuthSettingsParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOAuthSettingsParameter)(nil)).Elem()
}

func (o ApiOAuthSettingsParameterOutput) ToApiOAuthSettingsParameterOutput() ApiOAuthSettingsParameterOutput {
	return o
}

func (o ApiOAuthSettingsParameterOutput) ToApiOAuthSettingsParameterOutputWithContext(ctx context.Context) ApiOAuthSettingsParameterOutput {
	return o
}

func (o ApiOAuthSettingsParameterOutput) Options() pulumi.AnyOutput {
	return o.ApplyT(func(v ApiOAuthSettingsParameter) interface{} { return v.Options }).(pulumi.AnyOutput)
}

func (o ApiOAuthSettingsParameterOutput) UiDefinition() pulumi.AnyOutput {
	return o.ApplyT(func(v ApiOAuthSettingsParameter) interface{} { return v.UiDefinition }).(pulumi.AnyOutput)
}

func (o ApiOAuthSettingsParameterOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettingsParameter) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ApiOAuthSettingsParameterMapOutput struct{ *pulumi.OutputState }

func (ApiOAuthSettingsParameterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiOAuthSettingsParameter)(nil)).Elem()
}

func (o ApiOAuthSettingsParameterMapOutput) ToApiOAuthSettingsParameterMapOutput() ApiOAuthSettingsParameterMapOutput {
	return o
}

func (o ApiOAuthSettingsParameterMapOutput) ToApiOAuthSettingsParameterMapOutputWithContext(ctx context.Context) ApiOAuthSettingsParameterMapOutput {
	return o
}

func (o ApiOAuthSettingsParameterMapOutput) MapIndex(k pulumi.StringInput) ApiOAuthSettingsParameterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApiOAuthSettingsParameter {
		return vs[0].(map[string]ApiOAuthSettingsParameter)[vs[1].(string)]
	}).(ApiOAuthSettingsParameterOutput)
}

type ApiOAuthSettingsParameterResponse struct {
	Options      interface{} `pulumi:"options"`
	UiDefinition interface{} `pulumi:"uiDefinition"`
	Value        *string     `pulumi:"value"`
}





type ApiOAuthSettingsParameterResponseInput interface {
	pulumi.Input

	ToApiOAuthSettingsParameterResponseOutput() ApiOAuthSettingsParameterResponseOutput
	ToApiOAuthSettingsParameterResponseOutputWithContext(context.Context) ApiOAuthSettingsParameterResponseOutput
}

type ApiOAuthSettingsParameterResponseArgs struct {
	Options      pulumi.Input          `pulumi:"options"`
	UiDefinition pulumi.Input          `pulumi:"uiDefinition"`
	Value        pulumi.StringPtrInput `pulumi:"value"`
}

func (ApiOAuthSettingsParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOAuthSettingsParameterResponse)(nil)).Elem()
}

func (i ApiOAuthSettingsParameterResponseArgs) ToApiOAuthSettingsParameterResponseOutput() ApiOAuthSettingsParameterResponseOutput {
	return i.ToApiOAuthSettingsParameterResponseOutputWithContext(context.Background())
}

func (i ApiOAuthSettingsParameterResponseArgs) ToApiOAuthSettingsParameterResponseOutputWithContext(ctx context.Context) ApiOAuthSettingsParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsParameterResponseOutput)
}





type ApiOAuthSettingsParameterResponseMapInput interface {
	pulumi.Input

	ToApiOAuthSettingsParameterResponseMapOutput() ApiOAuthSettingsParameterResponseMapOutput
	ToApiOAuthSettingsParameterResponseMapOutputWithContext(context.Context) ApiOAuthSettingsParameterResponseMapOutput
}

type ApiOAuthSettingsParameterResponseMap map[string]ApiOAuthSettingsParameterResponseInput

func (ApiOAuthSettingsParameterResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiOAuthSettingsParameterResponse)(nil)).Elem()
}

func (i ApiOAuthSettingsParameterResponseMap) ToApiOAuthSettingsParameterResponseMapOutput() ApiOAuthSettingsParameterResponseMapOutput {
	return i.ToApiOAuthSettingsParameterResponseMapOutputWithContext(context.Background())
}

func (i ApiOAuthSettingsParameterResponseMap) ToApiOAuthSettingsParameterResponseMapOutputWithContext(ctx context.Context) ApiOAuthSettingsParameterResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsParameterResponseMapOutput)
}

type ApiOAuthSettingsParameterResponseOutput struct{ *pulumi.OutputState }

func (ApiOAuthSettingsParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOAuthSettingsParameterResponse)(nil)).Elem()
}

func (o ApiOAuthSettingsParameterResponseOutput) ToApiOAuthSettingsParameterResponseOutput() ApiOAuthSettingsParameterResponseOutput {
	return o
}

func (o ApiOAuthSettingsParameterResponseOutput) ToApiOAuthSettingsParameterResponseOutputWithContext(ctx context.Context) ApiOAuthSettingsParameterResponseOutput {
	return o
}

func (o ApiOAuthSettingsParameterResponseOutput) Options() pulumi.AnyOutput {
	return o.ApplyT(func(v ApiOAuthSettingsParameterResponse) interface{} { return v.Options }).(pulumi.AnyOutput)
}

func (o ApiOAuthSettingsParameterResponseOutput) UiDefinition() pulumi.AnyOutput {
	return o.ApplyT(func(v ApiOAuthSettingsParameterResponse) interface{} { return v.UiDefinition }).(pulumi.AnyOutput)
}

func (o ApiOAuthSettingsParameterResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettingsParameterResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ApiOAuthSettingsParameterResponseMapOutput struct{ *pulumi.OutputState }

func (ApiOAuthSettingsParameterResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiOAuthSettingsParameterResponse)(nil)).Elem()
}

func (o ApiOAuthSettingsParameterResponseMapOutput) ToApiOAuthSettingsParameterResponseMapOutput() ApiOAuthSettingsParameterResponseMapOutput {
	return o
}

func (o ApiOAuthSettingsParameterResponseMapOutput) ToApiOAuthSettingsParameterResponseMapOutputWithContext(ctx context.Context) ApiOAuthSettingsParameterResponseMapOutput {
	return o
}

func (o ApiOAuthSettingsParameterResponseMapOutput) MapIndex(k pulumi.StringInput) ApiOAuthSettingsParameterResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApiOAuthSettingsParameterResponse {
		return vs[0].(map[string]ApiOAuthSettingsParameterResponse)[vs[1].(string)]
	}).(ApiOAuthSettingsParameterResponseOutput)
}

type ApiOAuthSettingsResponse struct {
	ClientId         *string                                      `pulumi:"clientId"`
	ClientSecret     *string                                      `pulumi:"clientSecret"`
	CustomParameters map[string]ApiOAuthSettingsParameterResponse `pulumi:"customParameters"`
	IdentityProvider *string                                      `pulumi:"identityProvider"`
	Properties       interface{}                                  `pulumi:"properties"`
	RedirectUrl      *string                                      `pulumi:"redirectUrl"`
	Scopes           []string                                     `pulumi:"scopes"`
}





type ApiOAuthSettingsResponseInput interface {
	pulumi.Input

	ToApiOAuthSettingsResponseOutput() ApiOAuthSettingsResponseOutput
	ToApiOAuthSettingsResponseOutputWithContext(context.Context) ApiOAuthSettingsResponseOutput
}

type ApiOAuthSettingsResponseArgs struct {
	ClientId         pulumi.StringPtrInput                     `pulumi:"clientId"`
	ClientSecret     pulumi.StringPtrInput                     `pulumi:"clientSecret"`
	CustomParameters ApiOAuthSettingsParameterResponseMapInput `pulumi:"customParameters"`
	IdentityProvider pulumi.StringPtrInput                     `pulumi:"identityProvider"`
	Properties       pulumi.Input                              `pulumi:"properties"`
	RedirectUrl      pulumi.StringPtrInput                     `pulumi:"redirectUrl"`
	Scopes           pulumi.StringArrayInput                   `pulumi:"scopes"`
}

func (ApiOAuthSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOAuthSettingsResponse)(nil)).Elem()
}

func (i ApiOAuthSettingsResponseArgs) ToApiOAuthSettingsResponseOutput() ApiOAuthSettingsResponseOutput {
	return i.ToApiOAuthSettingsResponseOutputWithContext(context.Background())
}

func (i ApiOAuthSettingsResponseArgs) ToApiOAuthSettingsResponseOutputWithContext(ctx context.Context) ApiOAuthSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsResponseOutput)
}

func (i ApiOAuthSettingsResponseArgs) ToApiOAuthSettingsResponsePtrOutput() ApiOAuthSettingsResponsePtrOutput {
	return i.ToApiOAuthSettingsResponsePtrOutputWithContext(context.Background())
}

func (i ApiOAuthSettingsResponseArgs) ToApiOAuthSettingsResponsePtrOutputWithContext(ctx context.Context) ApiOAuthSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsResponseOutput).ToApiOAuthSettingsResponsePtrOutputWithContext(ctx)
}









type ApiOAuthSettingsResponsePtrInput interface {
	pulumi.Input

	ToApiOAuthSettingsResponsePtrOutput() ApiOAuthSettingsResponsePtrOutput
	ToApiOAuthSettingsResponsePtrOutputWithContext(context.Context) ApiOAuthSettingsResponsePtrOutput
}

type apiOAuthSettingsResponsePtrType ApiOAuthSettingsResponseArgs

func ApiOAuthSettingsResponsePtr(v *ApiOAuthSettingsResponseArgs) ApiOAuthSettingsResponsePtrInput {
	return (*apiOAuthSettingsResponsePtrType)(v)
}

func (*apiOAuthSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOAuthSettingsResponse)(nil)).Elem()
}

func (i *apiOAuthSettingsResponsePtrType) ToApiOAuthSettingsResponsePtrOutput() ApiOAuthSettingsResponsePtrOutput {
	return i.ToApiOAuthSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *apiOAuthSettingsResponsePtrType) ToApiOAuthSettingsResponsePtrOutputWithContext(ctx context.Context) ApiOAuthSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsResponsePtrOutput)
}

type ApiOAuthSettingsResponseOutput struct{ *pulumi.OutputState }

func (ApiOAuthSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOAuthSettingsResponse)(nil)).Elem()
}

func (o ApiOAuthSettingsResponseOutput) ToApiOAuthSettingsResponseOutput() ApiOAuthSettingsResponseOutput {
	return o
}

func (o ApiOAuthSettingsResponseOutput) ToApiOAuthSettingsResponseOutputWithContext(ctx context.Context) ApiOAuthSettingsResponseOutput {
	return o
}

func (o ApiOAuthSettingsResponseOutput) ToApiOAuthSettingsResponsePtrOutput() ApiOAuthSettingsResponsePtrOutput {
	return o.ToApiOAuthSettingsResponsePtrOutputWithContext(context.Background())
}

func (o ApiOAuthSettingsResponseOutput) ToApiOAuthSettingsResponsePtrOutputWithContext(ctx context.Context) ApiOAuthSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiOAuthSettingsResponse) *ApiOAuthSettingsResponse {
		return &v
	}).(ApiOAuthSettingsResponsePtrOutput)
}

func (o ApiOAuthSettingsResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettingsResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsResponseOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettingsResponse) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsResponseOutput) CustomParameters() ApiOAuthSettingsParameterResponseMapOutput {
	return o.ApplyT(func(v ApiOAuthSettingsResponse) map[string]ApiOAuthSettingsParameterResponse {
		return v.CustomParameters
	}).(ApiOAuthSettingsParameterResponseMapOutput)
}

func (o ApiOAuthSettingsResponseOutput) IdentityProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettingsResponse) *string { return v.IdentityProvider }).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsResponseOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v ApiOAuthSettingsResponse) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o ApiOAuthSettingsResponseOutput) RedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettingsResponse) *string { return v.RedirectUrl }).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsResponseOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApiOAuthSettingsResponse) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

type ApiOAuthSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiOAuthSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOAuthSettingsResponse)(nil)).Elem()
}

func (o ApiOAuthSettingsResponsePtrOutput) ToApiOAuthSettingsResponsePtrOutput() ApiOAuthSettingsResponsePtrOutput {
	return o
}

func (o ApiOAuthSettingsResponsePtrOutput) ToApiOAuthSettingsResponsePtrOutputWithContext(ctx context.Context) ApiOAuthSettingsResponsePtrOutput {
	return o
}

func (o ApiOAuthSettingsResponsePtrOutput) Elem() ApiOAuthSettingsResponseOutput {
	return o.ApplyT(func(v *ApiOAuthSettingsResponse) ApiOAuthSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ApiOAuthSettingsResponse
		return ret
	}).(ApiOAuthSettingsResponseOutput)
}

func (o ApiOAuthSettingsResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiOAuthSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsResponsePtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiOAuthSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsResponsePtrOutput) CustomParameters() ApiOAuthSettingsParameterResponseMapOutput {
	return o.ApplyT(func(v *ApiOAuthSettingsResponse) map[string]ApiOAuthSettingsParameterResponse {
		if v == nil {
			return nil
		}
		return v.CustomParameters
	}).(ApiOAuthSettingsParameterResponseMapOutput)
}

func (o ApiOAuthSettingsResponsePtrOutput) IdentityProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiOAuthSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.IdentityProvider
	}).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsResponsePtrOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *ApiOAuthSettingsResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.AnyOutput)
}

func (o ApiOAuthSettingsResponsePtrOutput) RedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiOAuthSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.RedirectUrl
	}).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsResponsePtrOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiOAuthSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.Scopes
	}).(pulumi.StringArrayOutput)
}

type ApiReference struct {
	BrandColor  *string     `pulumi:"brandColor"`
	Description *string     `pulumi:"description"`
	DisplayName *string     `pulumi:"displayName"`
	IconUri     *string     `pulumi:"iconUri"`
	Id          *string     `pulumi:"id"`
	Name        *string     `pulumi:"name"`
	Swagger     interface{} `pulumi:"swagger"`
	Type        *string     `pulumi:"type"`
}





type ApiReferenceInput interface {
	pulumi.Input

	ToApiReferenceOutput() ApiReferenceOutput
	ToApiReferenceOutputWithContext(context.Context) ApiReferenceOutput
}

type ApiReferenceArgs struct {
	BrandColor  pulumi.StringPtrInput `pulumi:"brandColor"`
	Description pulumi.StringPtrInput `pulumi:"description"`
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	IconUri     pulumi.StringPtrInput `pulumi:"iconUri"`
	Id          pulumi.StringPtrInput `pulumi:"id"`
	Name        pulumi.StringPtrInput `pulumi:"name"`
	Swagger     pulumi.Input          `pulumi:"swagger"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (ApiReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiReference)(nil)).Elem()
}

func (i ApiReferenceArgs) ToApiReferenceOutput() ApiReferenceOutput {
	return i.ToApiReferenceOutputWithContext(context.Background())
}

func (i ApiReferenceArgs) ToApiReferenceOutputWithContext(ctx context.Context) ApiReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiReferenceOutput)
}

func (i ApiReferenceArgs) ToApiReferencePtrOutput() ApiReferencePtrOutput {
	return i.ToApiReferencePtrOutputWithContext(context.Background())
}

func (i ApiReferenceArgs) ToApiReferencePtrOutputWithContext(ctx context.Context) ApiReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiReferenceOutput).ToApiReferencePtrOutputWithContext(ctx)
}









type ApiReferencePtrInput interface {
	pulumi.Input

	ToApiReferencePtrOutput() ApiReferencePtrOutput
	ToApiReferencePtrOutputWithContext(context.Context) ApiReferencePtrOutput
}

type apiReferencePtrType ApiReferenceArgs

func ApiReferencePtr(v *ApiReferenceArgs) ApiReferencePtrInput {
	return (*apiReferencePtrType)(v)
}

func (*apiReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiReference)(nil)).Elem()
}

func (i *apiReferencePtrType) ToApiReferencePtrOutput() ApiReferencePtrOutput {
	return i.ToApiReferencePtrOutputWithContext(context.Background())
}

func (i *apiReferencePtrType) ToApiReferencePtrOutputWithContext(ctx context.Context) ApiReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiReferencePtrOutput)
}

type ApiReferenceOutput struct{ *pulumi.OutputState }

func (ApiReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiReference)(nil)).Elem()
}

func (o ApiReferenceOutput) ToApiReferenceOutput() ApiReferenceOutput {
	return o
}

func (o ApiReferenceOutput) ToApiReferenceOutputWithContext(ctx context.Context) ApiReferenceOutput {
	return o
}

func (o ApiReferenceOutput) ToApiReferencePtrOutput() ApiReferencePtrOutput {
	return o.ToApiReferencePtrOutputWithContext(context.Background())
}

func (o ApiReferenceOutput) ToApiReferencePtrOutputWithContext(ctx context.Context) ApiReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiReference) *ApiReference {
		return &v
	}).(ApiReferencePtrOutput)
}

func (o ApiReferenceOutput) BrandColor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiReference) *string { return v.BrandColor }).(pulumi.StringPtrOutput)
}

func (o ApiReferenceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiReference) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ApiReferenceOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiReference) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ApiReferenceOutput) IconUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiReference) *string { return v.IconUri }).(pulumi.StringPtrOutput)
}

func (o ApiReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApiReferenceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiReference) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApiReferenceOutput) Swagger() pulumi.AnyOutput {
	return o.ApplyT(func(v ApiReference) interface{} { return v.Swagger }).(pulumi.AnyOutput)
}

func (o ApiReferenceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiReference) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ApiReferencePtrOutput struct{ *pulumi.OutputState }

func (ApiReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiReference)(nil)).Elem()
}

func (o ApiReferencePtrOutput) ToApiReferencePtrOutput() ApiReferencePtrOutput {
	return o
}

func (o ApiReferencePtrOutput) ToApiReferencePtrOutputWithContext(ctx context.Context) ApiReferencePtrOutput {
	return o
}

func (o ApiReferencePtrOutput) Elem() ApiReferenceOutput {
	return o.ApplyT(func(v *ApiReference) ApiReference {
		if v != nil {
			return *v
		}
		var ret ApiReference
		return ret
	}).(ApiReferenceOutput)
}

func (o ApiReferencePtrOutput) BrandColor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiReference) *string {
		if v == nil {
			return nil
		}
		return v.BrandColor
	}).(pulumi.StringPtrOutput)
}

func (o ApiReferencePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiReference) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ApiReferencePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiReference) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ApiReferencePtrOutput) IconUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiReference) *string {
		if v == nil {
			return nil
		}
		return v.IconUri
	}).(pulumi.StringPtrOutput)
}

func (o ApiReferencePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiReference) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ApiReferencePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiReference) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApiReferencePtrOutput) Swagger() pulumi.AnyOutput {
	return o.ApplyT(func(v *ApiReference) interface{} {
		if v == nil {
			return nil
		}
		return v.Swagger
	}).(pulumi.AnyOutput)
}

func (o ApiReferencePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiReference) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ApiReferenceResponse struct {
	BrandColor  *string     `pulumi:"brandColor"`
	Description *string     `pulumi:"description"`
	DisplayName *string     `pulumi:"displayName"`
	IconUri     *string     `pulumi:"iconUri"`
	Id          *string     `pulumi:"id"`
	Name        *string     `pulumi:"name"`
	Swagger     interface{} `pulumi:"swagger"`
	Type        *string     `pulumi:"type"`
}





type ApiReferenceResponseInput interface {
	pulumi.Input

	ToApiReferenceResponseOutput() ApiReferenceResponseOutput
	ToApiReferenceResponseOutputWithContext(context.Context) ApiReferenceResponseOutput
}

type ApiReferenceResponseArgs struct {
	BrandColor  pulumi.StringPtrInput `pulumi:"brandColor"`
	Description pulumi.StringPtrInput `pulumi:"description"`
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	IconUri     pulumi.StringPtrInput `pulumi:"iconUri"`
	Id          pulumi.StringPtrInput `pulumi:"id"`
	Name        pulumi.StringPtrInput `pulumi:"name"`
	Swagger     pulumi.Input          `pulumi:"swagger"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (ApiReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiReferenceResponse)(nil)).Elem()
}

func (i ApiReferenceResponseArgs) ToApiReferenceResponseOutput() ApiReferenceResponseOutput {
	return i.ToApiReferenceResponseOutputWithContext(context.Background())
}

func (i ApiReferenceResponseArgs) ToApiReferenceResponseOutputWithContext(ctx context.Context) ApiReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiReferenceResponseOutput)
}

func (i ApiReferenceResponseArgs) ToApiReferenceResponsePtrOutput() ApiReferenceResponsePtrOutput {
	return i.ToApiReferenceResponsePtrOutputWithContext(context.Background())
}

func (i ApiReferenceResponseArgs) ToApiReferenceResponsePtrOutputWithContext(ctx context.Context) ApiReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiReferenceResponseOutput).ToApiReferenceResponsePtrOutputWithContext(ctx)
}









type ApiReferenceResponsePtrInput interface {
	pulumi.Input

	ToApiReferenceResponsePtrOutput() ApiReferenceResponsePtrOutput
	ToApiReferenceResponsePtrOutputWithContext(context.Context) ApiReferenceResponsePtrOutput
}

type apiReferenceResponsePtrType ApiReferenceResponseArgs

func ApiReferenceResponsePtr(v *ApiReferenceResponseArgs) ApiReferenceResponsePtrInput {
	return (*apiReferenceResponsePtrType)(v)
}

func (*apiReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiReferenceResponse)(nil)).Elem()
}

func (i *apiReferenceResponsePtrType) ToApiReferenceResponsePtrOutput() ApiReferenceResponsePtrOutput {
	return i.ToApiReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *apiReferenceResponsePtrType) ToApiReferenceResponsePtrOutputWithContext(ctx context.Context) ApiReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiReferenceResponsePtrOutput)
}

type ApiReferenceResponseOutput struct{ *pulumi.OutputState }

func (ApiReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiReferenceResponse)(nil)).Elem()
}

func (o ApiReferenceResponseOutput) ToApiReferenceResponseOutput() ApiReferenceResponseOutput {
	return o
}

func (o ApiReferenceResponseOutput) ToApiReferenceResponseOutputWithContext(ctx context.Context) ApiReferenceResponseOutput {
	return o
}

func (o ApiReferenceResponseOutput) ToApiReferenceResponsePtrOutput() ApiReferenceResponsePtrOutput {
	return o.ToApiReferenceResponsePtrOutputWithContext(context.Background())
}

func (o ApiReferenceResponseOutput) ToApiReferenceResponsePtrOutputWithContext(ctx context.Context) ApiReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiReferenceResponse) *ApiReferenceResponse {
		return &v
	}).(ApiReferenceResponsePtrOutput)
}

func (o ApiReferenceResponseOutput) BrandColor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiReferenceResponse) *string { return v.BrandColor }).(pulumi.StringPtrOutput)
}

func (o ApiReferenceResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiReferenceResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ApiReferenceResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiReferenceResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ApiReferenceResponseOutput) IconUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiReferenceResponse) *string { return v.IconUri }).(pulumi.StringPtrOutput)
}

func (o ApiReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApiReferenceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiReferenceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApiReferenceResponseOutput) Swagger() pulumi.AnyOutput {
	return o.ApplyT(func(v ApiReferenceResponse) interface{} { return v.Swagger }).(pulumi.AnyOutput)
}

func (o ApiReferenceResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiReferenceResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ApiReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiReferenceResponse)(nil)).Elem()
}

func (o ApiReferenceResponsePtrOutput) ToApiReferenceResponsePtrOutput() ApiReferenceResponsePtrOutput {
	return o
}

func (o ApiReferenceResponsePtrOutput) ToApiReferenceResponsePtrOutputWithContext(ctx context.Context) ApiReferenceResponsePtrOutput {
	return o
}

func (o ApiReferenceResponsePtrOutput) Elem() ApiReferenceResponseOutput {
	return o.ApplyT(func(v *ApiReferenceResponse) ApiReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ApiReferenceResponse
		return ret
	}).(ApiReferenceResponseOutput)
}

func (o ApiReferenceResponsePtrOutput) BrandColor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.BrandColor
	}).(pulumi.StringPtrOutput)
}

func (o ApiReferenceResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ApiReferenceResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ApiReferenceResponsePtrOutput) IconUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.IconUri
	}).(pulumi.StringPtrOutput)
}

func (o ApiReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ApiReferenceResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApiReferenceResponsePtrOutput) Swagger() pulumi.AnyOutput {
	return o.ApplyT(func(v *ApiReferenceResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Swagger
	}).(pulumi.AnyOutput)
}

func (o ApiReferenceResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ApiResourceBackendService struct {
	ServiceUrl *string `pulumi:"serviceUrl"`
}





type ApiResourceBackendServiceInput interface {
	pulumi.Input

	ToApiResourceBackendServiceOutput() ApiResourceBackendServiceOutput
	ToApiResourceBackendServiceOutputWithContext(context.Context) ApiResourceBackendServiceOutput
}

type ApiResourceBackendServiceArgs struct {
	ServiceUrl pulumi.StringPtrInput `pulumi:"serviceUrl"`
}

func (ApiResourceBackendServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiResourceBackendService)(nil)).Elem()
}

func (i ApiResourceBackendServiceArgs) ToApiResourceBackendServiceOutput() ApiResourceBackendServiceOutput {
	return i.ToApiResourceBackendServiceOutputWithContext(context.Background())
}

func (i ApiResourceBackendServiceArgs) ToApiResourceBackendServiceOutputWithContext(ctx context.Context) ApiResourceBackendServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiResourceBackendServiceOutput)
}

func (i ApiResourceBackendServiceArgs) ToApiResourceBackendServicePtrOutput() ApiResourceBackendServicePtrOutput {
	return i.ToApiResourceBackendServicePtrOutputWithContext(context.Background())
}

func (i ApiResourceBackendServiceArgs) ToApiResourceBackendServicePtrOutputWithContext(ctx context.Context) ApiResourceBackendServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiResourceBackendServiceOutput).ToApiResourceBackendServicePtrOutputWithContext(ctx)
}









type ApiResourceBackendServicePtrInput interface {
	pulumi.Input

	ToApiResourceBackendServicePtrOutput() ApiResourceBackendServicePtrOutput
	ToApiResourceBackendServicePtrOutputWithContext(context.Context) ApiResourceBackendServicePtrOutput
}

type apiResourceBackendServicePtrType ApiResourceBackendServiceArgs

func ApiResourceBackendServicePtr(v *ApiResourceBackendServiceArgs) ApiResourceBackendServicePtrInput {
	return (*apiResourceBackendServicePtrType)(v)
}

func (*apiResourceBackendServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiResourceBackendService)(nil)).Elem()
}

func (i *apiResourceBackendServicePtrType) ToApiResourceBackendServicePtrOutput() ApiResourceBackendServicePtrOutput {
	return i.ToApiResourceBackendServicePtrOutputWithContext(context.Background())
}

func (i *apiResourceBackendServicePtrType) ToApiResourceBackendServicePtrOutputWithContext(ctx context.Context) ApiResourceBackendServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiResourceBackendServicePtrOutput)
}

type ApiResourceBackendServiceOutput struct{ *pulumi.OutputState }

func (ApiResourceBackendServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiResourceBackendService)(nil)).Elem()
}

func (o ApiResourceBackendServiceOutput) ToApiResourceBackendServiceOutput() ApiResourceBackendServiceOutput {
	return o
}

func (o ApiResourceBackendServiceOutput) ToApiResourceBackendServiceOutputWithContext(ctx context.Context) ApiResourceBackendServiceOutput {
	return o
}

func (o ApiResourceBackendServiceOutput) ToApiResourceBackendServicePtrOutput() ApiResourceBackendServicePtrOutput {
	return o.ToApiResourceBackendServicePtrOutputWithContext(context.Background())
}

func (o ApiResourceBackendServiceOutput) ToApiResourceBackendServicePtrOutputWithContext(ctx context.Context) ApiResourceBackendServicePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiResourceBackendService) *ApiResourceBackendService {
		return &v
	}).(ApiResourceBackendServicePtrOutput)
}

func (o ApiResourceBackendServiceOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceBackendService) *string { return v.ServiceUrl }).(pulumi.StringPtrOutput)
}

type ApiResourceBackendServicePtrOutput struct{ *pulumi.OutputState }

func (ApiResourceBackendServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiResourceBackendService)(nil)).Elem()
}

func (o ApiResourceBackendServicePtrOutput) ToApiResourceBackendServicePtrOutput() ApiResourceBackendServicePtrOutput {
	return o
}

func (o ApiResourceBackendServicePtrOutput) ToApiResourceBackendServicePtrOutputWithContext(ctx context.Context) ApiResourceBackendServicePtrOutput {
	return o
}

func (o ApiResourceBackendServicePtrOutput) Elem() ApiResourceBackendServiceOutput {
	return o.ApplyT(func(v *ApiResourceBackendService) ApiResourceBackendService {
		if v != nil {
			return *v
		}
		var ret ApiResourceBackendService
		return ret
	}).(ApiResourceBackendServiceOutput)
}

func (o ApiResourceBackendServicePtrOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiResourceBackendService) *string {
		if v == nil {
			return nil
		}
		return v.ServiceUrl
	}).(pulumi.StringPtrOutput)
}

type ApiResourceBackendServiceResponse struct {
	ServiceUrl *string `pulumi:"serviceUrl"`
}





type ApiResourceBackendServiceResponseInput interface {
	pulumi.Input

	ToApiResourceBackendServiceResponseOutput() ApiResourceBackendServiceResponseOutput
	ToApiResourceBackendServiceResponseOutputWithContext(context.Context) ApiResourceBackendServiceResponseOutput
}

type ApiResourceBackendServiceResponseArgs struct {
	ServiceUrl pulumi.StringPtrInput `pulumi:"serviceUrl"`
}

func (ApiResourceBackendServiceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiResourceBackendServiceResponse)(nil)).Elem()
}

func (i ApiResourceBackendServiceResponseArgs) ToApiResourceBackendServiceResponseOutput() ApiResourceBackendServiceResponseOutput {
	return i.ToApiResourceBackendServiceResponseOutputWithContext(context.Background())
}

func (i ApiResourceBackendServiceResponseArgs) ToApiResourceBackendServiceResponseOutputWithContext(ctx context.Context) ApiResourceBackendServiceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiResourceBackendServiceResponseOutput)
}

func (i ApiResourceBackendServiceResponseArgs) ToApiResourceBackendServiceResponsePtrOutput() ApiResourceBackendServiceResponsePtrOutput {
	return i.ToApiResourceBackendServiceResponsePtrOutputWithContext(context.Background())
}

func (i ApiResourceBackendServiceResponseArgs) ToApiResourceBackendServiceResponsePtrOutputWithContext(ctx context.Context) ApiResourceBackendServiceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiResourceBackendServiceResponseOutput).ToApiResourceBackendServiceResponsePtrOutputWithContext(ctx)
}









type ApiResourceBackendServiceResponsePtrInput interface {
	pulumi.Input

	ToApiResourceBackendServiceResponsePtrOutput() ApiResourceBackendServiceResponsePtrOutput
	ToApiResourceBackendServiceResponsePtrOutputWithContext(context.Context) ApiResourceBackendServiceResponsePtrOutput
}

type apiResourceBackendServiceResponsePtrType ApiResourceBackendServiceResponseArgs

func ApiResourceBackendServiceResponsePtr(v *ApiResourceBackendServiceResponseArgs) ApiResourceBackendServiceResponsePtrInput {
	return (*apiResourceBackendServiceResponsePtrType)(v)
}

func (*apiResourceBackendServiceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiResourceBackendServiceResponse)(nil)).Elem()
}

func (i *apiResourceBackendServiceResponsePtrType) ToApiResourceBackendServiceResponsePtrOutput() ApiResourceBackendServiceResponsePtrOutput {
	return i.ToApiResourceBackendServiceResponsePtrOutputWithContext(context.Background())
}

func (i *apiResourceBackendServiceResponsePtrType) ToApiResourceBackendServiceResponsePtrOutputWithContext(ctx context.Context) ApiResourceBackendServiceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiResourceBackendServiceResponsePtrOutput)
}

type ApiResourceBackendServiceResponseOutput struct{ *pulumi.OutputState }

func (ApiResourceBackendServiceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiResourceBackendServiceResponse)(nil)).Elem()
}

func (o ApiResourceBackendServiceResponseOutput) ToApiResourceBackendServiceResponseOutput() ApiResourceBackendServiceResponseOutput {
	return o
}

func (o ApiResourceBackendServiceResponseOutput) ToApiResourceBackendServiceResponseOutputWithContext(ctx context.Context) ApiResourceBackendServiceResponseOutput {
	return o
}

func (o ApiResourceBackendServiceResponseOutput) ToApiResourceBackendServiceResponsePtrOutput() ApiResourceBackendServiceResponsePtrOutput {
	return o.ToApiResourceBackendServiceResponsePtrOutputWithContext(context.Background())
}

func (o ApiResourceBackendServiceResponseOutput) ToApiResourceBackendServiceResponsePtrOutputWithContext(ctx context.Context) ApiResourceBackendServiceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiResourceBackendServiceResponse) *ApiResourceBackendServiceResponse {
		return &v
	}).(ApiResourceBackendServiceResponsePtrOutput)
}

func (o ApiResourceBackendServiceResponseOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceBackendServiceResponse) *string { return v.ServiceUrl }).(pulumi.StringPtrOutput)
}

type ApiResourceBackendServiceResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiResourceBackendServiceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiResourceBackendServiceResponse)(nil)).Elem()
}

func (o ApiResourceBackendServiceResponsePtrOutput) ToApiResourceBackendServiceResponsePtrOutput() ApiResourceBackendServiceResponsePtrOutput {
	return o
}

func (o ApiResourceBackendServiceResponsePtrOutput) ToApiResourceBackendServiceResponsePtrOutputWithContext(ctx context.Context) ApiResourceBackendServiceResponsePtrOutput {
	return o
}

func (o ApiResourceBackendServiceResponsePtrOutput) Elem() ApiResourceBackendServiceResponseOutput {
	return o.ApplyT(func(v *ApiResourceBackendServiceResponse) ApiResourceBackendServiceResponse {
		if v != nil {
			return *v
		}
		var ret ApiResourceBackendServiceResponse
		return ret
	}).(ApiResourceBackendServiceResponseOutput)
}

func (o ApiResourceBackendServiceResponsePtrOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiResourceBackendServiceResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceUrl
	}).(pulumi.StringPtrOutput)
}

type ApiResourceDefinitions struct {
	ModifiedSwaggerUrl *string `pulumi:"modifiedSwaggerUrl"`
	OriginalSwaggerUrl *string `pulumi:"originalSwaggerUrl"`
}





type ApiResourceDefinitionsInput interface {
	pulumi.Input

	ToApiResourceDefinitionsOutput() ApiResourceDefinitionsOutput
	ToApiResourceDefinitionsOutputWithContext(context.Context) ApiResourceDefinitionsOutput
}

type ApiResourceDefinitionsArgs struct {
	ModifiedSwaggerUrl pulumi.StringPtrInput `pulumi:"modifiedSwaggerUrl"`
	OriginalSwaggerUrl pulumi.StringPtrInput `pulumi:"originalSwaggerUrl"`
}

func (ApiResourceDefinitionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiResourceDefinitions)(nil)).Elem()
}

func (i ApiResourceDefinitionsArgs) ToApiResourceDefinitionsOutput() ApiResourceDefinitionsOutput {
	return i.ToApiResourceDefinitionsOutputWithContext(context.Background())
}

func (i ApiResourceDefinitionsArgs) ToApiResourceDefinitionsOutputWithContext(ctx context.Context) ApiResourceDefinitionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiResourceDefinitionsOutput)
}

func (i ApiResourceDefinitionsArgs) ToApiResourceDefinitionsPtrOutput() ApiResourceDefinitionsPtrOutput {
	return i.ToApiResourceDefinitionsPtrOutputWithContext(context.Background())
}

func (i ApiResourceDefinitionsArgs) ToApiResourceDefinitionsPtrOutputWithContext(ctx context.Context) ApiResourceDefinitionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiResourceDefinitionsOutput).ToApiResourceDefinitionsPtrOutputWithContext(ctx)
}









type ApiResourceDefinitionsPtrInput interface {
	pulumi.Input

	ToApiResourceDefinitionsPtrOutput() ApiResourceDefinitionsPtrOutput
	ToApiResourceDefinitionsPtrOutputWithContext(context.Context) ApiResourceDefinitionsPtrOutput
}

type apiResourceDefinitionsPtrType ApiResourceDefinitionsArgs

func ApiResourceDefinitionsPtr(v *ApiResourceDefinitionsArgs) ApiResourceDefinitionsPtrInput {
	return (*apiResourceDefinitionsPtrType)(v)
}

func (*apiResourceDefinitionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiResourceDefinitions)(nil)).Elem()
}

func (i *apiResourceDefinitionsPtrType) ToApiResourceDefinitionsPtrOutput() ApiResourceDefinitionsPtrOutput {
	return i.ToApiResourceDefinitionsPtrOutputWithContext(context.Background())
}

func (i *apiResourceDefinitionsPtrType) ToApiResourceDefinitionsPtrOutputWithContext(ctx context.Context) ApiResourceDefinitionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiResourceDefinitionsPtrOutput)
}

type ApiResourceDefinitionsOutput struct{ *pulumi.OutputState }

func (ApiResourceDefinitionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiResourceDefinitions)(nil)).Elem()
}

func (o ApiResourceDefinitionsOutput) ToApiResourceDefinitionsOutput() ApiResourceDefinitionsOutput {
	return o
}

func (o ApiResourceDefinitionsOutput) ToApiResourceDefinitionsOutputWithContext(ctx context.Context) ApiResourceDefinitionsOutput {
	return o
}

func (o ApiResourceDefinitionsOutput) ToApiResourceDefinitionsPtrOutput() ApiResourceDefinitionsPtrOutput {
	return o.ToApiResourceDefinitionsPtrOutputWithContext(context.Background())
}

func (o ApiResourceDefinitionsOutput) ToApiResourceDefinitionsPtrOutputWithContext(ctx context.Context) ApiResourceDefinitionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiResourceDefinitions) *ApiResourceDefinitions {
		return &v
	}).(ApiResourceDefinitionsPtrOutput)
}

func (o ApiResourceDefinitionsOutput) ModifiedSwaggerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceDefinitions) *string { return v.ModifiedSwaggerUrl }).(pulumi.StringPtrOutput)
}

func (o ApiResourceDefinitionsOutput) OriginalSwaggerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceDefinitions) *string { return v.OriginalSwaggerUrl }).(pulumi.StringPtrOutput)
}

type ApiResourceDefinitionsPtrOutput struct{ *pulumi.OutputState }

func (ApiResourceDefinitionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiResourceDefinitions)(nil)).Elem()
}

func (o ApiResourceDefinitionsPtrOutput) ToApiResourceDefinitionsPtrOutput() ApiResourceDefinitionsPtrOutput {
	return o
}

func (o ApiResourceDefinitionsPtrOutput) ToApiResourceDefinitionsPtrOutputWithContext(ctx context.Context) ApiResourceDefinitionsPtrOutput {
	return o
}

func (o ApiResourceDefinitionsPtrOutput) Elem() ApiResourceDefinitionsOutput {
	return o.ApplyT(func(v *ApiResourceDefinitions) ApiResourceDefinitions {
		if v != nil {
			return *v
		}
		var ret ApiResourceDefinitions
		return ret
	}).(ApiResourceDefinitionsOutput)
}

func (o ApiResourceDefinitionsPtrOutput) ModifiedSwaggerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiResourceDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.ModifiedSwaggerUrl
	}).(pulumi.StringPtrOutput)
}

func (o ApiResourceDefinitionsPtrOutput) OriginalSwaggerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiResourceDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.OriginalSwaggerUrl
	}).(pulumi.StringPtrOutput)
}

type ApiResourceDefinitionsResponse struct {
	ModifiedSwaggerUrl *string `pulumi:"modifiedSwaggerUrl"`
	OriginalSwaggerUrl *string `pulumi:"originalSwaggerUrl"`
}





type ApiResourceDefinitionsResponseInput interface {
	pulumi.Input

	ToApiResourceDefinitionsResponseOutput() ApiResourceDefinitionsResponseOutput
	ToApiResourceDefinitionsResponseOutputWithContext(context.Context) ApiResourceDefinitionsResponseOutput
}

type ApiResourceDefinitionsResponseArgs struct {
	ModifiedSwaggerUrl pulumi.StringPtrInput `pulumi:"modifiedSwaggerUrl"`
	OriginalSwaggerUrl pulumi.StringPtrInput `pulumi:"originalSwaggerUrl"`
}

func (ApiResourceDefinitionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiResourceDefinitionsResponse)(nil)).Elem()
}

func (i ApiResourceDefinitionsResponseArgs) ToApiResourceDefinitionsResponseOutput() ApiResourceDefinitionsResponseOutput {
	return i.ToApiResourceDefinitionsResponseOutputWithContext(context.Background())
}

func (i ApiResourceDefinitionsResponseArgs) ToApiResourceDefinitionsResponseOutputWithContext(ctx context.Context) ApiResourceDefinitionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiResourceDefinitionsResponseOutput)
}

func (i ApiResourceDefinitionsResponseArgs) ToApiResourceDefinitionsResponsePtrOutput() ApiResourceDefinitionsResponsePtrOutput {
	return i.ToApiResourceDefinitionsResponsePtrOutputWithContext(context.Background())
}

func (i ApiResourceDefinitionsResponseArgs) ToApiResourceDefinitionsResponsePtrOutputWithContext(ctx context.Context) ApiResourceDefinitionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiResourceDefinitionsResponseOutput).ToApiResourceDefinitionsResponsePtrOutputWithContext(ctx)
}









type ApiResourceDefinitionsResponsePtrInput interface {
	pulumi.Input

	ToApiResourceDefinitionsResponsePtrOutput() ApiResourceDefinitionsResponsePtrOutput
	ToApiResourceDefinitionsResponsePtrOutputWithContext(context.Context) ApiResourceDefinitionsResponsePtrOutput
}

type apiResourceDefinitionsResponsePtrType ApiResourceDefinitionsResponseArgs

func ApiResourceDefinitionsResponsePtr(v *ApiResourceDefinitionsResponseArgs) ApiResourceDefinitionsResponsePtrInput {
	return (*apiResourceDefinitionsResponsePtrType)(v)
}

func (*apiResourceDefinitionsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiResourceDefinitionsResponse)(nil)).Elem()
}

func (i *apiResourceDefinitionsResponsePtrType) ToApiResourceDefinitionsResponsePtrOutput() ApiResourceDefinitionsResponsePtrOutput {
	return i.ToApiResourceDefinitionsResponsePtrOutputWithContext(context.Background())
}

func (i *apiResourceDefinitionsResponsePtrType) ToApiResourceDefinitionsResponsePtrOutputWithContext(ctx context.Context) ApiResourceDefinitionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiResourceDefinitionsResponsePtrOutput)
}

type ApiResourceDefinitionsResponseOutput struct{ *pulumi.OutputState }

func (ApiResourceDefinitionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiResourceDefinitionsResponse)(nil)).Elem()
}

func (o ApiResourceDefinitionsResponseOutput) ToApiResourceDefinitionsResponseOutput() ApiResourceDefinitionsResponseOutput {
	return o
}

func (o ApiResourceDefinitionsResponseOutput) ToApiResourceDefinitionsResponseOutputWithContext(ctx context.Context) ApiResourceDefinitionsResponseOutput {
	return o
}

func (o ApiResourceDefinitionsResponseOutput) ToApiResourceDefinitionsResponsePtrOutput() ApiResourceDefinitionsResponsePtrOutput {
	return o.ToApiResourceDefinitionsResponsePtrOutputWithContext(context.Background())
}

func (o ApiResourceDefinitionsResponseOutput) ToApiResourceDefinitionsResponsePtrOutputWithContext(ctx context.Context) ApiResourceDefinitionsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiResourceDefinitionsResponse) *ApiResourceDefinitionsResponse {
		return &v
	}).(ApiResourceDefinitionsResponsePtrOutput)
}

func (o ApiResourceDefinitionsResponseOutput) ModifiedSwaggerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceDefinitionsResponse) *string { return v.ModifiedSwaggerUrl }).(pulumi.StringPtrOutput)
}

func (o ApiResourceDefinitionsResponseOutput) OriginalSwaggerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceDefinitionsResponse) *string { return v.OriginalSwaggerUrl }).(pulumi.StringPtrOutput)
}

type ApiResourceDefinitionsResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiResourceDefinitionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiResourceDefinitionsResponse)(nil)).Elem()
}

func (o ApiResourceDefinitionsResponsePtrOutput) ToApiResourceDefinitionsResponsePtrOutput() ApiResourceDefinitionsResponsePtrOutput {
	return o
}

func (o ApiResourceDefinitionsResponsePtrOutput) ToApiResourceDefinitionsResponsePtrOutputWithContext(ctx context.Context) ApiResourceDefinitionsResponsePtrOutput {
	return o
}

func (o ApiResourceDefinitionsResponsePtrOutput) Elem() ApiResourceDefinitionsResponseOutput {
	return o.ApplyT(func(v *ApiResourceDefinitionsResponse) ApiResourceDefinitionsResponse {
		if v != nil {
			return *v
		}
		var ret ApiResourceDefinitionsResponse
		return ret
	}).(ApiResourceDefinitionsResponseOutput)
}

func (o ApiResourceDefinitionsResponsePtrOutput) ModifiedSwaggerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiResourceDefinitionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ModifiedSwaggerUrl
	}).(pulumi.StringPtrOutput)
}

func (o ApiResourceDefinitionsResponsePtrOutput) OriginalSwaggerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiResourceDefinitionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.OriginalSwaggerUrl
	}).(pulumi.StringPtrOutput)
}

type ConnectionError struct {
	Code     *string           `pulumi:"code"`
	Etag     *string           `pulumi:"etag"`
	Location *string           `pulumi:"location"`
	Message  *string           `pulumi:"message"`
	Tags     map[string]string `pulumi:"tags"`
}





type ConnectionErrorInput interface {
	pulumi.Input

	ToConnectionErrorOutput() ConnectionErrorOutput
	ToConnectionErrorOutputWithContext(context.Context) ConnectionErrorOutput
}

type ConnectionErrorArgs struct {
	Code     pulumi.StringPtrInput `pulumi:"code"`
	Etag     pulumi.StringPtrInput `pulumi:"etag"`
	Location pulumi.StringPtrInput `pulumi:"location"`
	Message  pulumi.StringPtrInput `pulumi:"message"`
	Tags     pulumi.StringMapInput `pulumi:"tags"`
}

func (ConnectionErrorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionError)(nil)).Elem()
}

func (i ConnectionErrorArgs) ToConnectionErrorOutput() ConnectionErrorOutput {
	return i.ToConnectionErrorOutputWithContext(context.Background())
}

func (i ConnectionErrorArgs) ToConnectionErrorOutputWithContext(ctx context.Context) ConnectionErrorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionErrorOutput)
}

func (i ConnectionErrorArgs) ToConnectionErrorPtrOutput() ConnectionErrorPtrOutput {
	return i.ToConnectionErrorPtrOutputWithContext(context.Background())
}

func (i ConnectionErrorArgs) ToConnectionErrorPtrOutputWithContext(ctx context.Context) ConnectionErrorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionErrorOutput).ToConnectionErrorPtrOutputWithContext(ctx)
}









type ConnectionErrorPtrInput interface {
	pulumi.Input

	ToConnectionErrorPtrOutput() ConnectionErrorPtrOutput
	ToConnectionErrorPtrOutputWithContext(context.Context) ConnectionErrorPtrOutput
}

type connectionErrorPtrType ConnectionErrorArgs

func ConnectionErrorPtr(v *ConnectionErrorArgs) ConnectionErrorPtrInput {
	return (*connectionErrorPtrType)(v)
}

func (*connectionErrorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionError)(nil)).Elem()
}

func (i *connectionErrorPtrType) ToConnectionErrorPtrOutput() ConnectionErrorPtrOutput {
	return i.ToConnectionErrorPtrOutputWithContext(context.Background())
}

func (i *connectionErrorPtrType) ToConnectionErrorPtrOutputWithContext(ctx context.Context) ConnectionErrorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionErrorPtrOutput)
}

type ConnectionErrorOutput struct{ *pulumi.OutputState }

func (ConnectionErrorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionError)(nil)).Elem()
}

func (o ConnectionErrorOutput) ToConnectionErrorOutput() ConnectionErrorOutput {
	return o
}

func (o ConnectionErrorOutput) ToConnectionErrorOutputWithContext(ctx context.Context) ConnectionErrorOutput {
	return o
}

func (o ConnectionErrorOutput) ToConnectionErrorPtrOutput() ConnectionErrorPtrOutput {
	return o.ToConnectionErrorPtrOutputWithContext(context.Background())
}

func (o ConnectionErrorOutput) ToConnectionErrorPtrOutputWithContext(ctx context.Context) ConnectionErrorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionError) *ConnectionError {
		return &v
	}).(ConnectionErrorPtrOutput)
}

func (o ConnectionErrorOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionError) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionError) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionError) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionError) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ConnectionError) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ConnectionErrorPtrOutput struct{ *pulumi.OutputState }

func (ConnectionErrorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionError)(nil)).Elem()
}

func (o ConnectionErrorPtrOutput) ToConnectionErrorPtrOutput() ConnectionErrorPtrOutput {
	return o
}

func (o ConnectionErrorPtrOutput) ToConnectionErrorPtrOutputWithContext(ctx context.Context) ConnectionErrorPtrOutput {
	return o
}

func (o ConnectionErrorPtrOutput) Elem() ConnectionErrorOutput {
	return o.ApplyT(func(v *ConnectionError) ConnectionError {
		if v != nil {
			return *v
		}
		var ret ConnectionError
		return ret
	}).(ConnectionErrorOutput)
}

func (o ConnectionErrorPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionError) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorPtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionError) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorPtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionError) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionError) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorPtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConnectionError) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

type ConnectionErrorResponse struct {
	Code     *string           `pulumi:"code"`
	Etag     *string           `pulumi:"etag"`
	Id       string            `pulumi:"id"`
	Location *string           `pulumi:"location"`
	Message  *string           `pulumi:"message"`
	Name     string            `pulumi:"name"`
	Tags     map[string]string `pulumi:"tags"`
	Type     string            `pulumi:"type"`
}





type ConnectionErrorResponseInput interface {
	pulumi.Input

	ToConnectionErrorResponseOutput() ConnectionErrorResponseOutput
	ToConnectionErrorResponseOutputWithContext(context.Context) ConnectionErrorResponseOutput
}

type ConnectionErrorResponseArgs struct {
	Code     pulumi.StringPtrInput `pulumi:"code"`
	Etag     pulumi.StringPtrInput `pulumi:"etag"`
	Id       pulumi.StringInput    `pulumi:"id"`
	Location pulumi.StringPtrInput `pulumi:"location"`
	Message  pulumi.StringPtrInput `pulumi:"message"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Tags     pulumi.StringMapInput `pulumi:"tags"`
	Type     pulumi.StringInput    `pulumi:"type"`
}

func (ConnectionErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionErrorResponse)(nil)).Elem()
}

func (i ConnectionErrorResponseArgs) ToConnectionErrorResponseOutput() ConnectionErrorResponseOutput {
	return i.ToConnectionErrorResponseOutputWithContext(context.Background())
}

func (i ConnectionErrorResponseArgs) ToConnectionErrorResponseOutputWithContext(ctx context.Context) ConnectionErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionErrorResponseOutput)
}

func (i ConnectionErrorResponseArgs) ToConnectionErrorResponsePtrOutput() ConnectionErrorResponsePtrOutput {
	return i.ToConnectionErrorResponsePtrOutputWithContext(context.Background())
}

func (i ConnectionErrorResponseArgs) ToConnectionErrorResponsePtrOutputWithContext(ctx context.Context) ConnectionErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionErrorResponseOutput).ToConnectionErrorResponsePtrOutputWithContext(ctx)
}









type ConnectionErrorResponsePtrInput interface {
	pulumi.Input

	ToConnectionErrorResponsePtrOutput() ConnectionErrorResponsePtrOutput
	ToConnectionErrorResponsePtrOutputWithContext(context.Context) ConnectionErrorResponsePtrOutput
}

type connectionErrorResponsePtrType ConnectionErrorResponseArgs

func ConnectionErrorResponsePtr(v *ConnectionErrorResponseArgs) ConnectionErrorResponsePtrInput {
	return (*connectionErrorResponsePtrType)(v)
}

func (*connectionErrorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionErrorResponse)(nil)).Elem()
}

func (i *connectionErrorResponsePtrType) ToConnectionErrorResponsePtrOutput() ConnectionErrorResponsePtrOutput {
	return i.ToConnectionErrorResponsePtrOutputWithContext(context.Background())
}

func (i *connectionErrorResponsePtrType) ToConnectionErrorResponsePtrOutputWithContext(ctx context.Context) ConnectionErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionErrorResponsePtrOutput)
}

type ConnectionErrorResponseOutput struct{ *pulumi.OutputState }

func (ConnectionErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionErrorResponse)(nil)).Elem()
}

func (o ConnectionErrorResponseOutput) ToConnectionErrorResponseOutput() ConnectionErrorResponseOutput {
	return o
}

func (o ConnectionErrorResponseOutput) ToConnectionErrorResponseOutputWithContext(ctx context.Context) ConnectionErrorResponseOutput {
	return o
}

func (o ConnectionErrorResponseOutput) ToConnectionErrorResponsePtrOutput() ConnectionErrorResponsePtrOutput {
	return o.ToConnectionErrorResponsePtrOutputWithContext(context.Background())
}

func (o ConnectionErrorResponseOutput) ToConnectionErrorResponsePtrOutputWithContext(ctx context.Context) ConnectionErrorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionErrorResponse) *ConnectionErrorResponse {
		return &v
	}).(ConnectionErrorResponsePtrOutput)
}

func (o ConnectionErrorResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionErrorResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionErrorResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionErrorResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ConnectionErrorResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionErrorResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionErrorResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionErrorResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ConnectionErrorResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ConnectionErrorResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ConnectionErrorResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionErrorResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ConnectionErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectionErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionErrorResponse)(nil)).Elem()
}

func (o ConnectionErrorResponsePtrOutput) ToConnectionErrorResponsePtrOutput() ConnectionErrorResponsePtrOutput {
	return o
}

func (o ConnectionErrorResponsePtrOutput) ToConnectionErrorResponsePtrOutputWithContext(ctx context.Context) ConnectionErrorResponsePtrOutput {
	return o
}

func (o ConnectionErrorResponsePtrOutput) Elem() ConnectionErrorResponseOutput {
	return o.ApplyT(func(v *ConnectionErrorResponse) ConnectionErrorResponse {
		if v != nil {
			return *v
		}
		var ret ConnectionErrorResponse
		return ret
	}).(ConnectionErrorResponseOutput)
}

func (o ConnectionErrorResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionErrorResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionErrorResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConnectionErrorResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o ConnectionErrorResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionErrorResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ConnectionGatewayDefinitionProperties struct {
	BackendUri                    *string                     `pulumi:"backendUri"`
	ConnectionGatewayInstallation *ConnectionGatewayReference `pulumi:"connectionGatewayInstallation"`
	ContactInformation            []string                    `pulumi:"contactInformation"`
	Description                   *string                     `pulumi:"description"`
	DisplayName                   *string                     `pulumi:"displayName"`
	MachineName                   *string                     `pulumi:"machineName"`
	Status                        interface{}                 `pulumi:"status"`
}





type ConnectionGatewayDefinitionPropertiesInput interface {
	pulumi.Input

	ToConnectionGatewayDefinitionPropertiesOutput() ConnectionGatewayDefinitionPropertiesOutput
	ToConnectionGatewayDefinitionPropertiesOutputWithContext(context.Context) ConnectionGatewayDefinitionPropertiesOutput
}

type ConnectionGatewayDefinitionPropertiesArgs struct {
	BackendUri                    pulumi.StringPtrInput              `pulumi:"backendUri"`
	ConnectionGatewayInstallation ConnectionGatewayReferencePtrInput `pulumi:"connectionGatewayInstallation"`
	ContactInformation            pulumi.StringArrayInput            `pulumi:"contactInformation"`
	Description                   pulumi.StringPtrInput              `pulumi:"description"`
	DisplayName                   pulumi.StringPtrInput              `pulumi:"displayName"`
	MachineName                   pulumi.StringPtrInput              `pulumi:"machineName"`
	Status                        pulumi.Input                       `pulumi:"status"`
}

func (ConnectionGatewayDefinitionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionGatewayDefinitionProperties)(nil)).Elem()
}

func (i ConnectionGatewayDefinitionPropertiesArgs) ToConnectionGatewayDefinitionPropertiesOutput() ConnectionGatewayDefinitionPropertiesOutput {
	return i.ToConnectionGatewayDefinitionPropertiesOutputWithContext(context.Background())
}

func (i ConnectionGatewayDefinitionPropertiesArgs) ToConnectionGatewayDefinitionPropertiesOutputWithContext(ctx context.Context) ConnectionGatewayDefinitionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionGatewayDefinitionPropertiesOutput)
}

func (i ConnectionGatewayDefinitionPropertiesArgs) ToConnectionGatewayDefinitionPropertiesPtrOutput() ConnectionGatewayDefinitionPropertiesPtrOutput {
	return i.ToConnectionGatewayDefinitionPropertiesPtrOutputWithContext(context.Background())
}

func (i ConnectionGatewayDefinitionPropertiesArgs) ToConnectionGatewayDefinitionPropertiesPtrOutputWithContext(ctx context.Context) ConnectionGatewayDefinitionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionGatewayDefinitionPropertiesOutput).ToConnectionGatewayDefinitionPropertiesPtrOutputWithContext(ctx)
}









type ConnectionGatewayDefinitionPropertiesPtrInput interface {
	pulumi.Input

	ToConnectionGatewayDefinitionPropertiesPtrOutput() ConnectionGatewayDefinitionPropertiesPtrOutput
	ToConnectionGatewayDefinitionPropertiesPtrOutputWithContext(context.Context) ConnectionGatewayDefinitionPropertiesPtrOutput
}

type connectionGatewayDefinitionPropertiesPtrType ConnectionGatewayDefinitionPropertiesArgs

func ConnectionGatewayDefinitionPropertiesPtr(v *ConnectionGatewayDefinitionPropertiesArgs) ConnectionGatewayDefinitionPropertiesPtrInput {
	return (*connectionGatewayDefinitionPropertiesPtrType)(v)
}

func (*connectionGatewayDefinitionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionGatewayDefinitionProperties)(nil)).Elem()
}

func (i *connectionGatewayDefinitionPropertiesPtrType) ToConnectionGatewayDefinitionPropertiesPtrOutput() ConnectionGatewayDefinitionPropertiesPtrOutput {
	return i.ToConnectionGatewayDefinitionPropertiesPtrOutputWithContext(context.Background())
}

func (i *connectionGatewayDefinitionPropertiesPtrType) ToConnectionGatewayDefinitionPropertiesPtrOutputWithContext(ctx context.Context) ConnectionGatewayDefinitionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionGatewayDefinitionPropertiesPtrOutput)
}

type ConnectionGatewayDefinitionPropertiesOutput struct{ *pulumi.OutputState }

func (ConnectionGatewayDefinitionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionGatewayDefinitionProperties)(nil)).Elem()
}

func (o ConnectionGatewayDefinitionPropertiesOutput) ToConnectionGatewayDefinitionPropertiesOutput() ConnectionGatewayDefinitionPropertiesOutput {
	return o
}

func (o ConnectionGatewayDefinitionPropertiesOutput) ToConnectionGatewayDefinitionPropertiesOutputWithContext(ctx context.Context) ConnectionGatewayDefinitionPropertiesOutput {
	return o
}

func (o ConnectionGatewayDefinitionPropertiesOutput) ToConnectionGatewayDefinitionPropertiesPtrOutput() ConnectionGatewayDefinitionPropertiesPtrOutput {
	return o.ToConnectionGatewayDefinitionPropertiesPtrOutputWithContext(context.Background())
}

func (o ConnectionGatewayDefinitionPropertiesOutput) ToConnectionGatewayDefinitionPropertiesPtrOutputWithContext(ctx context.Context) ConnectionGatewayDefinitionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionGatewayDefinitionProperties) *ConnectionGatewayDefinitionProperties {
		return &v
	}).(ConnectionGatewayDefinitionPropertiesPtrOutput)
}

func (o ConnectionGatewayDefinitionPropertiesOutput) BackendUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionGatewayDefinitionProperties) *string { return v.BackendUri }).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayDefinitionPropertiesOutput) ConnectionGatewayInstallation() ConnectionGatewayReferencePtrOutput {
	return o.ApplyT(func(v ConnectionGatewayDefinitionProperties) *ConnectionGatewayReference {
		return v.ConnectionGatewayInstallation
	}).(ConnectionGatewayReferencePtrOutput)
}

func (o ConnectionGatewayDefinitionPropertiesOutput) ContactInformation() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ConnectionGatewayDefinitionProperties) []string { return v.ContactInformation }).(pulumi.StringArrayOutput)
}

func (o ConnectionGatewayDefinitionPropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionGatewayDefinitionProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayDefinitionPropertiesOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionGatewayDefinitionProperties) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayDefinitionPropertiesOutput) MachineName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionGatewayDefinitionProperties) *string { return v.MachineName }).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayDefinitionPropertiesOutput) Status() pulumi.AnyOutput {
	return o.ApplyT(func(v ConnectionGatewayDefinitionProperties) interface{} { return v.Status }).(pulumi.AnyOutput)
}

type ConnectionGatewayDefinitionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ConnectionGatewayDefinitionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionGatewayDefinitionProperties)(nil)).Elem()
}

func (o ConnectionGatewayDefinitionPropertiesPtrOutput) ToConnectionGatewayDefinitionPropertiesPtrOutput() ConnectionGatewayDefinitionPropertiesPtrOutput {
	return o
}

func (o ConnectionGatewayDefinitionPropertiesPtrOutput) ToConnectionGatewayDefinitionPropertiesPtrOutputWithContext(ctx context.Context) ConnectionGatewayDefinitionPropertiesPtrOutput {
	return o
}

func (o ConnectionGatewayDefinitionPropertiesPtrOutput) Elem() ConnectionGatewayDefinitionPropertiesOutput {
	return o.ApplyT(func(v *ConnectionGatewayDefinitionProperties) ConnectionGatewayDefinitionProperties {
		if v != nil {
			return *v
		}
		var ret ConnectionGatewayDefinitionProperties
		return ret
	}).(ConnectionGatewayDefinitionPropertiesOutput)
}

func (o ConnectionGatewayDefinitionPropertiesPtrOutput) BackendUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionGatewayDefinitionProperties) *string {
		if v == nil {
			return nil
		}
		return v.BackendUri
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayDefinitionPropertiesPtrOutput) ConnectionGatewayInstallation() ConnectionGatewayReferencePtrOutput {
	return o.ApplyT(func(v *ConnectionGatewayDefinitionProperties) *ConnectionGatewayReference {
		if v == nil {
			return nil
		}
		return v.ConnectionGatewayInstallation
	}).(ConnectionGatewayReferencePtrOutput)
}

func (o ConnectionGatewayDefinitionPropertiesPtrOutput) ContactInformation() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ConnectionGatewayDefinitionProperties) []string {
		if v == nil {
			return nil
		}
		return v.ContactInformation
	}).(pulumi.StringArrayOutput)
}

func (o ConnectionGatewayDefinitionPropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionGatewayDefinitionProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayDefinitionPropertiesPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionGatewayDefinitionProperties) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayDefinitionPropertiesPtrOutput) MachineName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionGatewayDefinitionProperties) *string {
		if v == nil {
			return nil
		}
		return v.MachineName
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayDefinitionPropertiesPtrOutput) Status() pulumi.AnyOutput {
	return o.ApplyT(func(v *ConnectionGatewayDefinitionProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.AnyOutput)
}

type ConnectionGatewayDefinitionResponseProperties struct {
	BackendUri                    *string                             `pulumi:"backendUri"`
	ConnectionGatewayInstallation *ConnectionGatewayReferenceResponse `pulumi:"connectionGatewayInstallation"`
	ContactInformation            []string                            `pulumi:"contactInformation"`
	Description                   *string                             `pulumi:"description"`
	DisplayName                   *string                             `pulumi:"displayName"`
	MachineName                   *string                             `pulumi:"machineName"`
	Status                        interface{}                         `pulumi:"status"`
}





type ConnectionGatewayDefinitionResponsePropertiesInput interface {
	pulumi.Input

	ToConnectionGatewayDefinitionResponsePropertiesOutput() ConnectionGatewayDefinitionResponsePropertiesOutput
	ToConnectionGatewayDefinitionResponsePropertiesOutputWithContext(context.Context) ConnectionGatewayDefinitionResponsePropertiesOutput
}

type ConnectionGatewayDefinitionResponsePropertiesArgs struct {
	BackendUri                    pulumi.StringPtrInput                      `pulumi:"backendUri"`
	ConnectionGatewayInstallation ConnectionGatewayReferenceResponsePtrInput `pulumi:"connectionGatewayInstallation"`
	ContactInformation            pulumi.StringArrayInput                    `pulumi:"contactInformation"`
	Description                   pulumi.StringPtrInput                      `pulumi:"description"`
	DisplayName                   pulumi.StringPtrInput                      `pulumi:"displayName"`
	MachineName                   pulumi.StringPtrInput                      `pulumi:"machineName"`
	Status                        pulumi.Input                               `pulumi:"status"`
}

func (ConnectionGatewayDefinitionResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionGatewayDefinitionResponseProperties)(nil)).Elem()
}

func (i ConnectionGatewayDefinitionResponsePropertiesArgs) ToConnectionGatewayDefinitionResponsePropertiesOutput() ConnectionGatewayDefinitionResponsePropertiesOutput {
	return i.ToConnectionGatewayDefinitionResponsePropertiesOutputWithContext(context.Background())
}

func (i ConnectionGatewayDefinitionResponsePropertiesArgs) ToConnectionGatewayDefinitionResponsePropertiesOutputWithContext(ctx context.Context) ConnectionGatewayDefinitionResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionGatewayDefinitionResponsePropertiesOutput)
}

func (i ConnectionGatewayDefinitionResponsePropertiesArgs) ToConnectionGatewayDefinitionResponsePropertiesPtrOutput() ConnectionGatewayDefinitionResponsePropertiesPtrOutput {
	return i.ToConnectionGatewayDefinitionResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i ConnectionGatewayDefinitionResponsePropertiesArgs) ToConnectionGatewayDefinitionResponsePropertiesPtrOutputWithContext(ctx context.Context) ConnectionGatewayDefinitionResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionGatewayDefinitionResponsePropertiesOutput).ToConnectionGatewayDefinitionResponsePropertiesPtrOutputWithContext(ctx)
}









type ConnectionGatewayDefinitionResponsePropertiesPtrInput interface {
	pulumi.Input

	ToConnectionGatewayDefinitionResponsePropertiesPtrOutput() ConnectionGatewayDefinitionResponsePropertiesPtrOutput
	ToConnectionGatewayDefinitionResponsePropertiesPtrOutputWithContext(context.Context) ConnectionGatewayDefinitionResponsePropertiesPtrOutput
}

type connectionGatewayDefinitionResponsePropertiesPtrType ConnectionGatewayDefinitionResponsePropertiesArgs

func ConnectionGatewayDefinitionResponsePropertiesPtr(v *ConnectionGatewayDefinitionResponsePropertiesArgs) ConnectionGatewayDefinitionResponsePropertiesPtrInput {
	return (*connectionGatewayDefinitionResponsePropertiesPtrType)(v)
}

func (*connectionGatewayDefinitionResponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionGatewayDefinitionResponseProperties)(nil)).Elem()
}

func (i *connectionGatewayDefinitionResponsePropertiesPtrType) ToConnectionGatewayDefinitionResponsePropertiesPtrOutput() ConnectionGatewayDefinitionResponsePropertiesPtrOutput {
	return i.ToConnectionGatewayDefinitionResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *connectionGatewayDefinitionResponsePropertiesPtrType) ToConnectionGatewayDefinitionResponsePropertiesPtrOutputWithContext(ctx context.Context) ConnectionGatewayDefinitionResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionGatewayDefinitionResponsePropertiesPtrOutput)
}

type ConnectionGatewayDefinitionResponsePropertiesOutput struct{ *pulumi.OutputState }

func (ConnectionGatewayDefinitionResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionGatewayDefinitionResponseProperties)(nil)).Elem()
}

func (o ConnectionGatewayDefinitionResponsePropertiesOutput) ToConnectionGatewayDefinitionResponsePropertiesOutput() ConnectionGatewayDefinitionResponsePropertiesOutput {
	return o
}

func (o ConnectionGatewayDefinitionResponsePropertiesOutput) ToConnectionGatewayDefinitionResponsePropertiesOutputWithContext(ctx context.Context) ConnectionGatewayDefinitionResponsePropertiesOutput {
	return o
}

func (o ConnectionGatewayDefinitionResponsePropertiesOutput) ToConnectionGatewayDefinitionResponsePropertiesPtrOutput() ConnectionGatewayDefinitionResponsePropertiesPtrOutput {
	return o.ToConnectionGatewayDefinitionResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o ConnectionGatewayDefinitionResponsePropertiesOutput) ToConnectionGatewayDefinitionResponsePropertiesPtrOutputWithContext(ctx context.Context) ConnectionGatewayDefinitionResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionGatewayDefinitionResponseProperties) *ConnectionGatewayDefinitionResponseProperties {
		return &v
	}).(ConnectionGatewayDefinitionResponsePropertiesPtrOutput)
}

func (o ConnectionGatewayDefinitionResponsePropertiesOutput) BackendUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionGatewayDefinitionResponseProperties) *string { return v.BackendUri }).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayDefinitionResponsePropertiesOutput) ConnectionGatewayInstallation() ConnectionGatewayReferenceResponsePtrOutput {
	return o.ApplyT(func(v ConnectionGatewayDefinitionResponseProperties) *ConnectionGatewayReferenceResponse {
		return v.ConnectionGatewayInstallation
	}).(ConnectionGatewayReferenceResponsePtrOutput)
}

func (o ConnectionGatewayDefinitionResponsePropertiesOutput) ContactInformation() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ConnectionGatewayDefinitionResponseProperties) []string { return v.ContactInformation }).(pulumi.StringArrayOutput)
}

func (o ConnectionGatewayDefinitionResponsePropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionGatewayDefinitionResponseProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayDefinitionResponsePropertiesOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionGatewayDefinitionResponseProperties) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayDefinitionResponsePropertiesOutput) MachineName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionGatewayDefinitionResponseProperties) *string { return v.MachineName }).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayDefinitionResponsePropertiesOutput) Status() pulumi.AnyOutput {
	return o.ApplyT(func(v ConnectionGatewayDefinitionResponseProperties) interface{} { return v.Status }).(pulumi.AnyOutput)
}

type ConnectionGatewayDefinitionResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ConnectionGatewayDefinitionResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionGatewayDefinitionResponseProperties)(nil)).Elem()
}

func (o ConnectionGatewayDefinitionResponsePropertiesPtrOutput) ToConnectionGatewayDefinitionResponsePropertiesPtrOutput() ConnectionGatewayDefinitionResponsePropertiesPtrOutput {
	return o
}

func (o ConnectionGatewayDefinitionResponsePropertiesPtrOutput) ToConnectionGatewayDefinitionResponsePropertiesPtrOutputWithContext(ctx context.Context) ConnectionGatewayDefinitionResponsePropertiesPtrOutput {
	return o
}

func (o ConnectionGatewayDefinitionResponsePropertiesPtrOutput) Elem() ConnectionGatewayDefinitionResponsePropertiesOutput {
	return o.ApplyT(func(v *ConnectionGatewayDefinitionResponseProperties) ConnectionGatewayDefinitionResponseProperties {
		if v != nil {
			return *v
		}
		var ret ConnectionGatewayDefinitionResponseProperties
		return ret
	}).(ConnectionGatewayDefinitionResponsePropertiesOutput)
}

func (o ConnectionGatewayDefinitionResponsePropertiesPtrOutput) BackendUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionGatewayDefinitionResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.BackendUri
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayDefinitionResponsePropertiesPtrOutput) ConnectionGatewayInstallation() ConnectionGatewayReferenceResponsePtrOutput {
	return o.ApplyT(func(v *ConnectionGatewayDefinitionResponseProperties) *ConnectionGatewayReferenceResponse {
		if v == nil {
			return nil
		}
		return v.ConnectionGatewayInstallation
	}).(ConnectionGatewayReferenceResponsePtrOutput)
}

func (o ConnectionGatewayDefinitionResponsePropertiesPtrOutput) ContactInformation() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ConnectionGatewayDefinitionResponseProperties) []string {
		if v == nil {
			return nil
		}
		return v.ContactInformation
	}).(pulumi.StringArrayOutput)
}

func (o ConnectionGatewayDefinitionResponsePropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionGatewayDefinitionResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayDefinitionResponsePropertiesPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionGatewayDefinitionResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayDefinitionResponsePropertiesPtrOutput) MachineName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionGatewayDefinitionResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.MachineName
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayDefinitionResponsePropertiesPtrOutput) Status() pulumi.AnyOutput {
	return o.ApplyT(func(v *ConnectionGatewayDefinitionResponseProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.AnyOutput)
}

type ConnectionGatewayReference struct {
	Id       *string `pulumi:"id"`
	Location *string `pulumi:"location"`
	Name     *string `pulumi:"name"`
	Type     *string `pulumi:"type"`
}





type ConnectionGatewayReferenceInput interface {
	pulumi.Input

	ToConnectionGatewayReferenceOutput() ConnectionGatewayReferenceOutput
	ToConnectionGatewayReferenceOutputWithContext(context.Context) ConnectionGatewayReferenceOutput
}

type ConnectionGatewayReferenceArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Location pulumi.StringPtrInput `pulumi:"location"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Type     pulumi.StringPtrInput `pulumi:"type"`
}

func (ConnectionGatewayReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionGatewayReference)(nil)).Elem()
}

func (i ConnectionGatewayReferenceArgs) ToConnectionGatewayReferenceOutput() ConnectionGatewayReferenceOutput {
	return i.ToConnectionGatewayReferenceOutputWithContext(context.Background())
}

func (i ConnectionGatewayReferenceArgs) ToConnectionGatewayReferenceOutputWithContext(ctx context.Context) ConnectionGatewayReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionGatewayReferenceOutput)
}

func (i ConnectionGatewayReferenceArgs) ToConnectionGatewayReferencePtrOutput() ConnectionGatewayReferencePtrOutput {
	return i.ToConnectionGatewayReferencePtrOutputWithContext(context.Background())
}

func (i ConnectionGatewayReferenceArgs) ToConnectionGatewayReferencePtrOutputWithContext(ctx context.Context) ConnectionGatewayReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionGatewayReferenceOutput).ToConnectionGatewayReferencePtrOutputWithContext(ctx)
}









type ConnectionGatewayReferencePtrInput interface {
	pulumi.Input

	ToConnectionGatewayReferencePtrOutput() ConnectionGatewayReferencePtrOutput
	ToConnectionGatewayReferencePtrOutputWithContext(context.Context) ConnectionGatewayReferencePtrOutput
}

type connectionGatewayReferencePtrType ConnectionGatewayReferenceArgs

func ConnectionGatewayReferencePtr(v *ConnectionGatewayReferenceArgs) ConnectionGatewayReferencePtrInput {
	return (*connectionGatewayReferencePtrType)(v)
}

func (*connectionGatewayReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionGatewayReference)(nil)).Elem()
}

func (i *connectionGatewayReferencePtrType) ToConnectionGatewayReferencePtrOutput() ConnectionGatewayReferencePtrOutput {
	return i.ToConnectionGatewayReferencePtrOutputWithContext(context.Background())
}

func (i *connectionGatewayReferencePtrType) ToConnectionGatewayReferencePtrOutputWithContext(ctx context.Context) ConnectionGatewayReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionGatewayReferencePtrOutput)
}

type ConnectionGatewayReferenceOutput struct{ *pulumi.OutputState }

func (ConnectionGatewayReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionGatewayReference)(nil)).Elem()
}

func (o ConnectionGatewayReferenceOutput) ToConnectionGatewayReferenceOutput() ConnectionGatewayReferenceOutput {
	return o
}

func (o ConnectionGatewayReferenceOutput) ToConnectionGatewayReferenceOutputWithContext(ctx context.Context) ConnectionGatewayReferenceOutput {
	return o
}

func (o ConnectionGatewayReferenceOutput) ToConnectionGatewayReferencePtrOutput() ConnectionGatewayReferencePtrOutput {
	return o.ToConnectionGatewayReferencePtrOutputWithContext(context.Background())
}

func (o ConnectionGatewayReferenceOutput) ToConnectionGatewayReferencePtrOutputWithContext(ctx context.Context) ConnectionGatewayReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionGatewayReference) *ConnectionGatewayReference {
		return &v
	}).(ConnectionGatewayReferencePtrOutput)
}

func (o ConnectionGatewayReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionGatewayReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayReferenceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionGatewayReference) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayReferenceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionGatewayReference) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayReferenceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionGatewayReference) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ConnectionGatewayReferencePtrOutput struct{ *pulumi.OutputState }

func (ConnectionGatewayReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionGatewayReference)(nil)).Elem()
}

func (o ConnectionGatewayReferencePtrOutput) ToConnectionGatewayReferencePtrOutput() ConnectionGatewayReferencePtrOutput {
	return o
}

func (o ConnectionGatewayReferencePtrOutput) ToConnectionGatewayReferencePtrOutputWithContext(ctx context.Context) ConnectionGatewayReferencePtrOutput {
	return o
}

func (o ConnectionGatewayReferencePtrOutput) Elem() ConnectionGatewayReferenceOutput {
	return o.ApplyT(func(v *ConnectionGatewayReference) ConnectionGatewayReference {
		if v != nil {
			return *v
		}
		var ret ConnectionGatewayReference
		return ret
	}).(ConnectionGatewayReferenceOutput)
}

func (o ConnectionGatewayReferencePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionGatewayReference) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayReferencePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionGatewayReference) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayReferencePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionGatewayReference) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayReferencePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionGatewayReference) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ConnectionGatewayReferenceResponse struct {
	Id       *string `pulumi:"id"`
	Location *string `pulumi:"location"`
	Name     *string `pulumi:"name"`
	Type     *string `pulumi:"type"`
}





type ConnectionGatewayReferenceResponseInput interface {
	pulumi.Input

	ToConnectionGatewayReferenceResponseOutput() ConnectionGatewayReferenceResponseOutput
	ToConnectionGatewayReferenceResponseOutputWithContext(context.Context) ConnectionGatewayReferenceResponseOutput
}

type ConnectionGatewayReferenceResponseArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Location pulumi.StringPtrInput `pulumi:"location"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Type     pulumi.StringPtrInput `pulumi:"type"`
}

func (ConnectionGatewayReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionGatewayReferenceResponse)(nil)).Elem()
}

func (i ConnectionGatewayReferenceResponseArgs) ToConnectionGatewayReferenceResponseOutput() ConnectionGatewayReferenceResponseOutput {
	return i.ToConnectionGatewayReferenceResponseOutputWithContext(context.Background())
}

func (i ConnectionGatewayReferenceResponseArgs) ToConnectionGatewayReferenceResponseOutputWithContext(ctx context.Context) ConnectionGatewayReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionGatewayReferenceResponseOutput)
}

func (i ConnectionGatewayReferenceResponseArgs) ToConnectionGatewayReferenceResponsePtrOutput() ConnectionGatewayReferenceResponsePtrOutput {
	return i.ToConnectionGatewayReferenceResponsePtrOutputWithContext(context.Background())
}

func (i ConnectionGatewayReferenceResponseArgs) ToConnectionGatewayReferenceResponsePtrOutputWithContext(ctx context.Context) ConnectionGatewayReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionGatewayReferenceResponseOutput).ToConnectionGatewayReferenceResponsePtrOutputWithContext(ctx)
}









type ConnectionGatewayReferenceResponsePtrInput interface {
	pulumi.Input

	ToConnectionGatewayReferenceResponsePtrOutput() ConnectionGatewayReferenceResponsePtrOutput
	ToConnectionGatewayReferenceResponsePtrOutputWithContext(context.Context) ConnectionGatewayReferenceResponsePtrOutput
}

type connectionGatewayReferenceResponsePtrType ConnectionGatewayReferenceResponseArgs

func ConnectionGatewayReferenceResponsePtr(v *ConnectionGatewayReferenceResponseArgs) ConnectionGatewayReferenceResponsePtrInput {
	return (*connectionGatewayReferenceResponsePtrType)(v)
}

func (*connectionGatewayReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionGatewayReferenceResponse)(nil)).Elem()
}

func (i *connectionGatewayReferenceResponsePtrType) ToConnectionGatewayReferenceResponsePtrOutput() ConnectionGatewayReferenceResponsePtrOutput {
	return i.ToConnectionGatewayReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *connectionGatewayReferenceResponsePtrType) ToConnectionGatewayReferenceResponsePtrOutputWithContext(ctx context.Context) ConnectionGatewayReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionGatewayReferenceResponsePtrOutput)
}

type ConnectionGatewayReferenceResponseOutput struct{ *pulumi.OutputState }

func (ConnectionGatewayReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionGatewayReferenceResponse)(nil)).Elem()
}

func (o ConnectionGatewayReferenceResponseOutput) ToConnectionGatewayReferenceResponseOutput() ConnectionGatewayReferenceResponseOutput {
	return o
}

func (o ConnectionGatewayReferenceResponseOutput) ToConnectionGatewayReferenceResponseOutputWithContext(ctx context.Context) ConnectionGatewayReferenceResponseOutput {
	return o
}

func (o ConnectionGatewayReferenceResponseOutput) ToConnectionGatewayReferenceResponsePtrOutput() ConnectionGatewayReferenceResponsePtrOutput {
	return o.ToConnectionGatewayReferenceResponsePtrOutputWithContext(context.Background())
}

func (o ConnectionGatewayReferenceResponseOutput) ToConnectionGatewayReferenceResponsePtrOutputWithContext(ctx context.Context) ConnectionGatewayReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionGatewayReferenceResponse) *ConnectionGatewayReferenceResponse {
		return &v
	}).(ConnectionGatewayReferenceResponsePtrOutput)
}

func (o ConnectionGatewayReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionGatewayReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayReferenceResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionGatewayReferenceResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayReferenceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionGatewayReferenceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayReferenceResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionGatewayReferenceResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ConnectionGatewayReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectionGatewayReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionGatewayReferenceResponse)(nil)).Elem()
}

func (o ConnectionGatewayReferenceResponsePtrOutput) ToConnectionGatewayReferenceResponsePtrOutput() ConnectionGatewayReferenceResponsePtrOutput {
	return o
}

func (o ConnectionGatewayReferenceResponsePtrOutput) ToConnectionGatewayReferenceResponsePtrOutputWithContext(ctx context.Context) ConnectionGatewayReferenceResponsePtrOutput {
	return o
}

func (o ConnectionGatewayReferenceResponsePtrOutput) Elem() ConnectionGatewayReferenceResponseOutput {
	return o.ApplyT(func(v *ConnectionGatewayReferenceResponse) ConnectionGatewayReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ConnectionGatewayReferenceResponse
		return ret
	}).(ConnectionGatewayReferenceResponseOutput)
}

func (o ConnectionGatewayReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionGatewayReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayReferenceResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionGatewayReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayReferenceResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionGatewayReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionGatewayReferenceResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionGatewayReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ConnectionParameter struct {
	OAuthSettings *ApiOAuthSettings        `pulumi:"oAuthSettings"`
	Type          *ConnectionParameterType `pulumi:"type"`
}





type ConnectionParameterInput interface {
	pulumi.Input

	ToConnectionParameterOutput() ConnectionParameterOutput
	ToConnectionParameterOutputWithContext(context.Context) ConnectionParameterOutput
}

type ConnectionParameterArgs struct {
	OAuthSettings ApiOAuthSettingsPtrInput        `pulumi:"oAuthSettings"`
	Type          ConnectionParameterTypePtrInput `pulumi:"type"`
}

func (ConnectionParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionParameter)(nil)).Elem()
}

func (i ConnectionParameterArgs) ToConnectionParameterOutput() ConnectionParameterOutput {
	return i.ToConnectionParameterOutputWithContext(context.Background())
}

func (i ConnectionParameterArgs) ToConnectionParameterOutputWithContext(ctx context.Context) ConnectionParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionParameterOutput)
}





type ConnectionParameterMapInput interface {
	pulumi.Input

	ToConnectionParameterMapOutput() ConnectionParameterMapOutput
	ToConnectionParameterMapOutputWithContext(context.Context) ConnectionParameterMapOutput
}

type ConnectionParameterMap map[string]ConnectionParameterInput

func (ConnectionParameterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConnectionParameter)(nil)).Elem()
}

func (i ConnectionParameterMap) ToConnectionParameterMapOutput() ConnectionParameterMapOutput {
	return i.ToConnectionParameterMapOutputWithContext(context.Background())
}

func (i ConnectionParameterMap) ToConnectionParameterMapOutputWithContext(ctx context.Context) ConnectionParameterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionParameterMapOutput)
}

type ConnectionParameterOutput struct{ *pulumi.OutputState }

func (ConnectionParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionParameter)(nil)).Elem()
}

func (o ConnectionParameterOutput) ToConnectionParameterOutput() ConnectionParameterOutput {
	return o
}

func (o ConnectionParameterOutput) ToConnectionParameterOutputWithContext(ctx context.Context) ConnectionParameterOutput {
	return o
}

func (o ConnectionParameterOutput) OAuthSettings() ApiOAuthSettingsPtrOutput {
	return o.ApplyT(func(v ConnectionParameter) *ApiOAuthSettings { return v.OAuthSettings }).(ApiOAuthSettingsPtrOutput)
}

func (o ConnectionParameterOutput) Type() ConnectionParameterTypePtrOutput {
	return o.ApplyT(func(v ConnectionParameter) *ConnectionParameterType { return v.Type }).(ConnectionParameterTypePtrOutput)
}

type ConnectionParameterMapOutput struct{ *pulumi.OutputState }

func (ConnectionParameterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConnectionParameter)(nil)).Elem()
}

func (o ConnectionParameterMapOutput) ToConnectionParameterMapOutput() ConnectionParameterMapOutput {
	return o
}

func (o ConnectionParameterMapOutput) ToConnectionParameterMapOutputWithContext(ctx context.Context) ConnectionParameterMapOutput {
	return o
}

func (o ConnectionParameterMapOutput) MapIndex(k pulumi.StringInput) ConnectionParameterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ConnectionParameter {
		return vs[0].(map[string]ConnectionParameter)[vs[1].(string)]
	}).(ConnectionParameterOutput)
}

type ConnectionParameterResponse struct {
	OAuthSettings *ApiOAuthSettingsResponse `pulumi:"oAuthSettings"`
	Type          *string                   `pulumi:"type"`
}





type ConnectionParameterResponseInput interface {
	pulumi.Input

	ToConnectionParameterResponseOutput() ConnectionParameterResponseOutput
	ToConnectionParameterResponseOutputWithContext(context.Context) ConnectionParameterResponseOutput
}

type ConnectionParameterResponseArgs struct {
	OAuthSettings ApiOAuthSettingsResponsePtrInput `pulumi:"oAuthSettings"`
	Type          pulumi.StringPtrInput            `pulumi:"type"`
}

func (ConnectionParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionParameterResponse)(nil)).Elem()
}

func (i ConnectionParameterResponseArgs) ToConnectionParameterResponseOutput() ConnectionParameterResponseOutput {
	return i.ToConnectionParameterResponseOutputWithContext(context.Background())
}

func (i ConnectionParameterResponseArgs) ToConnectionParameterResponseOutputWithContext(ctx context.Context) ConnectionParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionParameterResponseOutput)
}





type ConnectionParameterResponseMapInput interface {
	pulumi.Input

	ToConnectionParameterResponseMapOutput() ConnectionParameterResponseMapOutput
	ToConnectionParameterResponseMapOutputWithContext(context.Context) ConnectionParameterResponseMapOutput
}

type ConnectionParameterResponseMap map[string]ConnectionParameterResponseInput

func (ConnectionParameterResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConnectionParameterResponse)(nil)).Elem()
}

func (i ConnectionParameterResponseMap) ToConnectionParameterResponseMapOutput() ConnectionParameterResponseMapOutput {
	return i.ToConnectionParameterResponseMapOutputWithContext(context.Background())
}

func (i ConnectionParameterResponseMap) ToConnectionParameterResponseMapOutputWithContext(ctx context.Context) ConnectionParameterResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionParameterResponseMapOutput)
}

type ConnectionParameterResponseOutput struct{ *pulumi.OutputState }

func (ConnectionParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionParameterResponse)(nil)).Elem()
}

func (o ConnectionParameterResponseOutput) ToConnectionParameterResponseOutput() ConnectionParameterResponseOutput {
	return o
}

func (o ConnectionParameterResponseOutput) ToConnectionParameterResponseOutputWithContext(ctx context.Context) ConnectionParameterResponseOutput {
	return o
}

func (o ConnectionParameterResponseOutput) OAuthSettings() ApiOAuthSettingsResponsePtrOutput {
	return o.ApplyT(func(v ConnectionParameterResponse) *ApiOAuthSettingsResponse { return v.OAuthSettings }).(ApiOAuthSettingsResponsePtrOutput)
}

func (o ConnectionParameterResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionParameterResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ConnectionParameterResponseMapOutput struct{ *pulumi.OutputState }

func (ConnectionParameterResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConnectionParameterResponse)(nil)).Elem()
}

func (o ConnectionParameterResponseMapOutput) ToConnectionParameterResponseMapOutput() ConnectionParameterResponseMapOutput {
	return o
}

func (o ConnectionParameterResponseMapOutput) ToConnectionParameterResponseMapOutputWithContext(ctx context.Context) ConnectionParameterResponseMapOutput {
	return o
}

func (o ConnectionParameterResponseMapOutput) MapIndex(k pulumi.StringInput) ConnectionParameterResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ConnectionParameterResponse {
		return vs[0].(map[string]ConnectionParameterResponse)[vs[1].(string)]
	}).(ConnectionParameterResponseOutput)
}

type ConnectionStatusDefinition struct {
	Error  *ConnectionError `pulumi:"error"`
	Status *string          `pulumi:"status"`
	Target *string          `pulumi:"target"`
}





type ConnectionStatusDefinitionInput interface {
	pulumi.Input

	ToConnectionStatusDefinitionOutput() ConnectionStatusDefinitionOutput
	ToConnectionStatusDefinitionOutputWithContext(context.Context) ConnectionStatusDefinitionOutput
}

type ConnectionStatusDefinitionArgs struct {
	Error  ConnectionErrorPtrInput `pulumi:"error"`
	Status pulumi.StringPtrInput   `pulumi:"status"`
	Target pulumi.StringPtrInput   `pulumi:"target"`
}

func (ConnectionStatusDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStatusDefinition)(nil)).Elem()
}

func (i ConnectionStatusDefinitionArgs) ToConnectionStatusDefinitionOutput() ConnectionStatusDefinitionOutput {
	return i.ToConnectionStatusDefinitionOutputWithContext(context.Background())
}

func (i ConnectionStatusDefinitionArgs) ToConnectionStatusDefinitionOutputWithContext(ctx context.Context) ConnectionStatusDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStatusDefinitionOutput)
}





type ConnectionStatusDefinitionArrayInput interface {
	pulumi.Input

	ToConnectionStatusDefinitionArrayOutput() ConnectionStatusDefinitionArrayOutput
	ToConnectionStatusDefinitionArrayOutputWithContext(context.Context) ConnectionStatusDefinitionArrayOutput
}

type ConnectionStatusDefinitionArray []ConnectionStatusDefinitionInput

func (ConnectionStatusDefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectionStatusDefinition)(nil)).Elem()
}

func (i ConnectionStatusDefinitionArray) ToConnectionStatusDefinitionArrayOutput() ConnectionStatusDefinitionArrayOutput {
	return i.ToConnectionStatusDefinitionArrayOutputWithContext(context.Background())
}

func (i ConnectionStatusDefinitionArray) ToConnectionStatusDefinitionArrayOutputWithContext(ctx context.Context) ConnectionStatusDefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStatusDefinitionArrayOutput)
}

type ConnectionStatusDefinitionOutput struct{ *pulumi.OutputState }

func (ConnectionStatusDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStatusDefinition)(nil)).Elem()
}

func (o ConnectionStatusDefinitionOutput) ToConnectionStatusDefinitionOutput() ConnectionStatusDefinitionOutput {
	return o
}

func (o ConnectionStatusDefinitionOutput) ToConnectionStatusDefinitionOutputWithContext(ctx context.Context) ConnectionStatusDefinitionOutput {
	return o
}

func (o ConnectionStatusDefinitionOutput) Error() ConnectionErrorPtrOutput {
	return o.ApplyT(func(v ConnectionStatusDefinition) *ConnectionError { return v.Error }).(ConnectionErrorPtrOutput)
}

func (o ConnectionStatusDefinitionOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStatusDefinition) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ConnectionStatusDefinitionOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStatusDefinition) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type ConnectionStatusDefinitionArrayOutput struct{ *pulumi.OutputState }

func (ConnectionStatusDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectionStatusDefinition)(nil)).Elem()
}

func (o ConnectionStatusDefinitionArrayOutput) ToConnectionStatusDefinitionArrayOutput() ConnectionStatusDefinitionArrayOutput {
	return o
}

func (o ConnectionStatusDefinitionArrayOutput) ToConnectionStatusDefinitionArrayOutputWithContext(ctx context.Context) ConnectionStatusDefinitionArrayOutput {
	return o
}

func (o ConnectionStatusDefinitionArrayOutput) Index(i pulumi.IntInput) ConnectionStatusDefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnectionStatusDefinition {
		return vs[0].([]ConnectionStatusDefinition)[vs[1].(int)]
	}).(ConnectionStatusDefinitionOutput)
}

type ConnectionStatusDefinitionResponse struct {
	Error  *ConnectionErrorResponse `pulumi:"error"`
	Status *string                  `pulumi:"status"`
	Target *string                  `pulumi:"target"`
}





type ConnectionStatusDefinitionResponseInput interface {
	pulumi.Input

	ToConnectionStatusDefinitionResponseOutput() ConnectionStatusDefinitionResponseOutput
	ToConnectionStatusDefinitionResponseOutputWithContext(context.Context) ConnectionStatusDefinitionResponseOutput
}

type ConnectionStatusDefinitionResponseArgs struct {
	Error  ConnectionErrorResponsePtrInput `pulumi:"error"`
	Status pulumi.StringPtrInput           `pulumi:"status"`
	Target pulumi.StringPtrInput           `pulumi:"target"`
}

func (ConnectionStatusDefinitionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStatusDefinitionResponse)(nil)).Elem()
}

func (i ConnectionStatusDefinitionResponseArgs) ToConnectionStatusDefinitionResponseOutput() ConnectionStatusDefinitionResponseOutput {
	return i.ToConnectionStatusDefinitionResponseOutputWithContext(context.Background())
}

func (i ConnectionStatusDefinitionResponseArgs) ToConnectionStatusDefinitionResponseOutputWithContext(ctx context.Context) ConnectionStatusDefinitionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStatusDefinitionResponseOutput)
}





type ConnectionStatusDefinitionResponseArrayInput interface {
	pulumi.Input

	ToConnectionStatusDefinitionResponseArrayOutput() ConnectionStatusDefinitionResponseArrayOutput
	ToConnectionStatusDefinitionResponseArrayOutputWithContext(context.Context) ConnectionStatusDefinitionResponseArrayOutput
}

type ConnectionStatusDefinitionResponseArray []ConnectionStatusDefinitionResponseInput

func (ConnectionStatusDefinitionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectionStatusDefinitionResponse)(nil)).Elem()
}

func (i ConnectionStatusDefinitionResponseArray) ToConnectionStatusDefinitionResponseArrayOutput() ConnectionStatusDefinitionResponseArrayOutput {
	return i.ToConnectionStatusDefinitionResponseArrayOutputWithContext(context.Background())
}

func (i ConnectionStatusDefinitionResponseArray) ToConnectionStatusDefinitionResponseArrayOutputWithContext(ctx context.Context) ConnectionStatusDefinitionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStatusDefinitionResponseArrayOutput)
}

type ConnectionStatusDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ConnectionStatusDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStatusDefinitionResponse)(nil)).Elem()
}

func (o ConnectionStatusDefinitionResponseOutput) ToConnectionStatusDefinitionResponseOutput() ConnectionStatusDefinitionResponseOutput {
	return o
}

func (o ConnectionStatusDefinitionResponseOutput) ToConnectionStatusDefinitionResponseOutputWithContext(ctx context.Context) ConnectionStatusDefinitionResponseOutput {
	return o
}

func (o ConnectionStatusDefinitionResponseOutput) Error() ConnectionErrorResponsePtrOutput {
	return o.ApplyT(func(v ConnectionStatusDefinitionResponse) *ConnectionErrorResponse { return v.Error }).(ConnectionErrorResponsePtrOutput)
}

func (o ConnectionStatusDefinitionResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStatusDefinitionResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ConnectionStatusDefinitionResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStatusDefinitionResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type ConnectionStatusDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (ConnectionStatusDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectionStatusDefinitionResponse)(nil)).Elem()
}

func (o ConnectionStatusDefinitionResponseArrayOutput) ToConnectionStatusDefinitionResponseArrayOutput() ConnectionStatusDefinitionResponseArrayOutput {
	return o
}

func (o ConnectionStatusDefinitionResponseArrayOutput) ToConnectionStatusDefinitionResponseArrayOutputWithContext(ctx context.Context) ConnectionStatusDefinitionResponseArrayOutput {
	return o
}

func (o ConnectionStatusDefinitionResponseArrayOutput) Index(i pulumi.IntInput) ConnectionStatusDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnectionStatusDefinitionResponse {
		return vs[0].([]ConnectionStatusDefinitionResponse)[vs[1].(int)]
	}).(ConnectionStatusDefinitionResponseOutput)
}

type ConsentLinkDefinitionResponse struct {
	DisplayName        *string `pulumi:"displayName"`
	FirstPartyLoginUri *string `pulumi:"firstPartyLoginUri"`
	Link               *string `pulumi:"link"`
	Status             *string `pulumi:"status"`
}





type ConsentLinkDefinitionResponseInput interface {
	pulumi.Input

	ToConsentLinkDefinitionResponseOutput() ConsentLinkDefinitionResponseOutput
	ToConsentLinkDefinitionResponseOutputWithContext(context.Context) ConsentLinkDefinitionResponseOutput
}

type ConsentLinkDefinitionResponseArgs struct {
	DisplayName        pulumi.StringPtrInput `pulumi:"displayName"`
	FirstPartyLoginUri pulumi.StringPtrInput `pulumi:"firstPartyLoginUri"`
	Link               pulumi.StringPtrInput `pulumi:"link"`
	Status             pulumi.StringPtrInput `pulumi:"status"`
}

func (ConsentLinkDefinitionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsentLinkDefinitionResponse)(nil)).Elem()
}

func (i ConsentLinkDefinitionResponseArgs) ToConsentLinkDefinitionResponseOutput() ConsentLinkDefinitionResponseOutput {
	return i.ToConsentLinkDefinitionResponseOutputWithContext(context.Background())
}

func (i ConsentLinkDefinitionResponseArgs) ToConsentLinkDefinitionResponseOutputWithContext(ctx context.Context) ConsentLinkDefinitionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsentLinkDefinitionResponseOutput)
}





type ConsentLinkDefinitionResponseArrayInput interface {
	pulumi.Input

	ToConsentLinkDefinitionResponseArrayOutput() ConsentLinkDefinitionResponseArrayOutput
	ToConsentLinkDefinitionResponseArrayOutputWithContext(context.Context) ConsentLinkDefinitionResponseArrayOutput
}

type ConsentLinkDefinitionResponseArray []ConsentLinkDefinitionResponseInput

func (ConsentLinkDefinitionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConsentLinkDefinitionResponse)(nil)).Elem()
}

func (i ConsentLinkDefinitionResponseArray) ToConsentLinkDefinitionResponseArrayOutput() ConsentLinkDefinitionResponseArrayOutput {
	return i.ToConsentLinkDefinitionResponseArrayOutputWithContext(context.Background())
}

func (i ConsentLinkDefinitionResponseArray) ToConsentLinkDefinitionResponseArrayOutputWithContext(ctx context.Context) ConsentLinkDefinitionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsentLinkDefinitionResponseArrayOutput)
}

type ConsentLinkDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ConsentLinkDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsentLinkDefinitionResponse)(nil)).Elem()
}

func (o ConsentLinkDefinitionResponseOutput) ToConsentLinkDefinitionResponseOutput() ConsentLinkDefinitionResponseOutput {
	return o
}

func (o ConsentLinkDefinitionResponseOutput) ToConsentLinkDefinitionResponseOutputWithContext(ctx context.Context) ConsentLinkDefinitionResponseOutput {
	return o
}

func (o ConsentLinkDefinitionResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsentLinkDefinitionResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ConsentLinkDefinitionResponseOutput) FirstPartyLoginUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsentLinkDefinitionResponse) *string { return v.FirstPartyLoginUri }).(pulumi.StringPtrOutput)
}

func (o ConsentLinkDefinitionResponseOutput) Link() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsentLinkDefinitionResponse) *string { return v.Link }).(pulumi.StringPtrOutput)
}

func (o ConsentLinkDefinitionResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsentLinkDefinitionResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ConsentLinkDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (ConsentLinkDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConsentLinkDefinitionResponse)(nil)).Elem()
}

func (o ConsentLinkDefinitionResponseArrayOutput) ToConsentLinkDefinitionResponseArrayOutput() ConsentLinkDefinitionResponseArrayOutput {
	return o
}

func (o ConsentLinkDefinitionResponseArrayOutput) ToConsentLinkDefinitionResponseArrayOutputWithContext(ctx context.Context) ConsentLinkDefinitionResponseArrayOutput {
	return o
}

func (o ConsentLinkDefinitionResponseArrayOutput) Index(i pulumi.IntInput) ConsentLinkDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConsentLinkDefinitionResponse {
		return vs[0].([]ConsentLinkDefinitionResponse)[vs[1].(int)]
	}).(ConsentLinkDefinitionResponseOutput)
}

type ConsentLinkParameterDefinition struct {
	ObjectId      *string `pulumi:"objectId"`
	ParameterName *string `pulumi:"parameterName"`
	RedirectUrl   *string `pulumi:"redirectUrl"`
	TenantId      *string `pulumi:"tenantId"`
}





type ConsentLinkParameterDefinitionInput interface {
	pulumi.Input

	ToConsentLinkParameterDefinitionOutput() ConsentLinkParameterDefinitionOutput
	ToConsentLinkParameterDefinitionOutputWithContext(context.Context) ConsentLinkParameterDefinitionOutput
}

type ConsentLinkParameterDefinitionArgs struct {
	ObjectId      pulumi.StringPtrInput `pulumi:"objectId"`
	ParameterName pulumi.StringPtrInput `pulumi:"parameterName"`
	RedirectUrl   pulumi.StringPtrInput `pulumi:"redirectUrl"`
	TenantId      pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (ConsentLinkParameterDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsentLinkParameterDefinition)(nil)).Elem()
}

func (i ConsentLinkParameterDefinitionArgs) ToConsentLinkParameterDefinitionOutput() ConsentLinkParameterDefinitionOutput {
	return i.ToConsentLinkParameterDefinitionOutputWithContext(context.Background())
}

func (i ConsentLinkParameterDefinitionArgs) ToConsentLinkParameterDefinitionOutputWithContext(ctx context.Context) ConsentLinkParameterDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsentLinkParameterDefinitionOutput)
}





type ConsentLinkParameterDefinitionArrayInput interface {
	pulumi.Input

	ToConsentLinkParameterDefinitionArrayOutput() ConsentLinkParameterDefinitionArrayOutput
	ToConsentLinkParameterDefinitionArrayOutputWithContext(context.Context) ConsentLinkParameterDefinitionArrayOutput
}

type ConsentLinkParameterDefinitionArray []ConsentLinkParameterDefinitionInput

func (ConsentLinkParameterDefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConsentLinkParameterDefinition)(nil)).Elem()
}

func (i ConsentLinkParameterDefinitionArray) ToConsentLinkParameterDefinitionArrayOutput() ConsentLinkParameterDefinitionArrayOutput {
	return i.ToConsentLinkParameterDefinitionArrayOutputWithContext(context.Background())
}

func (i ConsentLinkParameterDefinitionArray) ToConsentLinkParameterDefinitionArrayOutputWithContext(ctx context.Context) ConsentLinkParameterDefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsentLinkParameterDefinitionArrayOutput)
}

type ConsentLinkParameterDefinitionOutput struct{ *pulumi.OutputState }

func (ConsentLinkParameterDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsentLinkParameterDefinition)(nil)).Elem()
}

func (o ConsentLinkParameterDefinitionOutput) ToConsentLinkParameterDefinitionOutput() ConsentLinkParameterDefinitionOutput {
	return o
}

func (o ConsentLinkParameterDefinitionOutput) ToConsentLinkParameterDefinitionOutputWithContext(ctx context.Context) ConsentLinkParameterDefinitionOutput {
	return o
}

func (o ConsentLinkParameterDefinitionOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsentLinkParameterDefinition) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o ConsentLinkParameterDefinitionOutput) ParameterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsentLinkParameterDefinition) *string { return v.ParameterName }).(pulumi.StringPtrOutput)
}

func (o ConsentLinkParameterDefinitionOutput) RedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsentLinkParameterDefinition) *string { return v.RedirectUrl }).(pulumi.StringPtrOutput)
}

func (o ConsentLinkParameterDefinitionOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsentLinkParameterDefinition) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type ConsentLinkParameterDefinitionArrayOutput struct{ *pulumi.OutputState }

func (ConsentLinkParameterDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConsentLinkParameterDefinition)(nil)).Elem()
}

func (o ConsentLinkParameterDefinitionArrayOutput) ToConsentLinkParameterDefinitionArrayOutput() ConsentLinkParameterDefinitionArrayOutput {
	return o
}

func (o ConsentLinkParameterDefinitionArrayOutput) ToConsentLinkParameterDefinitionArrayOutputWithContext(ctx context.Context) ConsentLinkParameterDefinitionArrayOutput {
	return o
}

func (o ConsentLinkParameterDefinitionArrayOutput) Index(i pulumi.IntInput) ConsentLinkParameterDefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConsentLinkParameterDefinition {
		return vs[0].([]ConsentLinkParameterDefinition)[vs[1].(int)]
	}).(ConsentLinkParameterDefinitionOutput)
}

type CustomApiPropertiesDefinition struct {
	ApiDefinitions       *ApiResourceDefinitions        `pulumi:"apiDefinitions"`
	ApiType              *string                        `pulumi:"apiType"`
	BackendService       *ApiResourceBackendService     `pulumi:"backendService"`
	BrandColor           *string                        `pulumi:"brandColor"`
	Capabilities         []string                       `pulumi:"capabilities"`
	ConnectionParameters map[string]ConnectionParameter `pulumi:"connectionParameters"`
	Description          *string                        `pulumi:"description"`
	DisplayName          *string                        `pulumi:"displayName"`
	IconUri              *string                        `pulumi:"iconUri"`
	RuntimeUrls          []string                       `pulumi:"runtimeUrls"`
	Swagger              interface{}                    `pulumi:"swagger"`
	WsdlDefinition       *WsdlDefinition                `pulumi:"wsdlDefinition"`
}





type CustomApiPropertiesDefinitionInput interface {
	pulumi.Input

	ToCustomApiPropertiesDefinitionOutput() CustomApiPropertiesDefinitionOutput
	ToCustomApiPropertiesDefinitionOutputWithContext(context.Context) CustomApiPropertiesDefinitionOutput
}

type CustomApiPropertiesDefinitionArgs struct {
	ApiDefinitions       ApiResourceDefinitionsPtrInput    `pulumi:"apiDefinitions"`
	ApiType              pulumi.StringPtrInput             `pulumi:"apiType"`
	BackendService       ApiResourceBackendServicePtrInput `pulumi:"backendService"`
	BrandColor           pulumi.StringPtrInput             `pulumi:"brandColor"`
	Capabilities         pulumi.StringArrayInput           `pulumi:"capabilities"`
	ConnectionParameters ConnectionParameterMapInput       `pulumi:"connectionParameters"`
	Description          pulumi.StringPtrInput             `pulumi:"description"`
	DisplayName          pulumi.StringPtrInput             `pulumi:"displayName"`
	IconUri              pulumi.StringPtrInput             `pulumi:"iconUri"`
	RuntimeUrls          pulumi.StringArrayInput           `pulumi:"runtimeUrls"`
	Swagger              pulumi.Input                      `pulumi:"swagger"`
	WsdlDefinition       WsdlDefinitionPtrInput            `pulumi:"wsdlDefinition"`
}

func (CustomApiPropertiesDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomApiPropertiesDefinition)(nil)).Elem()
}

func (i CustomApiPropertiesDefinitionArgs) ToCustomApiPropertiesDefinitionOutput() CustomApiPropertiesDefinitionOutput {
	return i.ToCustomApiPropertiesDefinitionOutputWithContext(context.Background())
}

func (i CustomApiPropertiesDefinitionArgs) ToCustomApiPropertiesDefinitionOutputWithContext(ctx context.Context) CustomApiPropertiesDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomApiPropertiesDefinitionOutput)
}

func (i CustomApiPropertiesDefinitionArgs) ToCustomApiPropertiesDefinitionPtrOutput() CustomApiPropertiesDefinitionPtrOutput {
	return i.ToCustomApiPropertiesDefinitionPtrOutputWithContext(context.Background())
}

func (i CustomApiPropertiesDefinitionArgs) ToCustomApiPropertiesDefinitionPtrOutputWithContext(ctx context.Context) CustomApiPropertiesDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomApiPropertiesDefinitionOutput).ToCustomApiPropertiesDefinitionPtrOutputWithContext(ctx)
}









type CustomApiPropertiesDefinitionPtrInput interface {
	pulumi.Input

	ToCustomApiPropertiesDefinitionPtrOutput() CustomApiPropertiesDefinitionPtrOutput
	ToCustomApiPropertiesDefinitionPtrOutputWithContext(context.Context) CustomApiPropertiesDefinitionPtrOutput
}

type customApiPropertiesDefinitionPtrType CustomApiPropertiesDefinitionArgs

func CustomApiPropertiesDefinitionPtr(v *CustomApiPropertiesDefinitionArgs) CustomApiPropertiesDefinitionPtrInput {
	return (*customApiPropertiesDefinitionPtrType)(v)
}

func (*customApiPropertiesDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomApiPropertiesDefinition)(nil)).Elem()
}

func (i *customApiPropertiesDefinitionPtrType) ToCustomApiPropertiesDefinitionPtrOutput() CustomApiPropertiesDefinitionPtrOutput {
	return i.ToCustomApiPropertiesDefinitionPtrOutputWithContext(context.Background())
}

func (i *customApiPropertiesDefinitionPtrType) ToCustomApiPropertiesDefinitionPtrOutputWithContext(ctx context.Context) CustomApiPropertiesDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomApiPropertiesDefinitionPtrOutput)
}

type CustomApiPropertiesDefinitionOutput struct{ *pulumi.OutputState }

func (CustomApiPropertiesDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomApiPropertiesDefinition)(nil)).Elem()
}

func (o CustomApiPropertiesDefinitionOutput) ToCustomApiPropertiesDefinitionOutput() CustomApiPropertiesDefinitionOutput {
	return o
}

func (o CustomApiPropertiesDefinitionOutput) ToCustomApiPropertiesDefinitionOutputWithContext(ctx context.Context) CustomApiPropertiesDefinitionOutput {
	return o
}

func (o CustomApiPropertiesDefinitionOutput) ToCustomApiPropertiesDefinitionPtrOutput() CustomApiPropertiesDefinitionPtrOutput {
	return o.ToCustomApiPropertiesDefinitionPtrOutputWithContext(context.Background())
}

func (o CustomApiPropertiesDefinitionOutput) ToCustomApiPropertiesDefinitionPtrOutputWithContext(ctx context.Context) CustomApiPropertiesDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomApiPropertiesDefinition) *CustomApiPropertiesDefinition {
		return &v
	}).(CustomApiPropertiesDefinitionPtrOutput)
}

func (o CustomApiPropertiesDefinitionOutput) ApiDefinitions() ApiResourceDefinitionsPtrOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinition) *ApiResourceDefinitions { return v.ApiDefinitions }).(ApiResourceDefinitionsPtrOutput)
}

func (o CustomApiPropertiesDefinitionOutput) ApiType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinition) *string { return v.ApiType }).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionOutput) BackendService() ApiResourceBackendServicePtrOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinition) *ApiResourceBackendService { return v.BackendService }).(ApiResourceBackendServicePtrOutput)
}

func (o CustomApiPropertiesDefinitionOutput) BrandColor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinition) *string { return v.BrandColor }).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionOutput) Capabilities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinition) []string { return v.Capabilities }).(pulumi.StringArrayOutput)
}

func (o CustomApiPropertiesDefinitionOutput) ConnectionParameters() ConnectionParameterMapOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinition) map[string]ConnectionParameter { return v.ConnectionParameters }).(ConnectionParameterMapOutput)
}

func (o CustomApiPropertiesDefinitionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinition) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionOutput) IconUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinition) *string { return v.IconUri }).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionOutput) RuntimeUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinition) []string { return v.RuntimeUrls }).(pulumi.StringArrayOutput)
}

func (o CustomApiPropertiesDefinitionOutput) Swagger() pulumi.AnyOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinition) interface{} { return v.Swagger }).(pulumi.AnyOutput)
}

func (o CustomApiPropertiesDefinitionOutput) WsdlDefinition() WsdlDefinitionPtrOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinition) *WsdlDefinition { return v.WsdlDefinition }).(WsdlDefinitionPtrOutput)
}

type CustomApiPropertiesDefinitionPtrOutput struct{ *pulumi.OutputState }

func (CustomApiPropertiesDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomApiPropertiesDefinition)(nil)).Elem()
}

func (o CustomApiPropertiesDefinitionPtrOutput) ToCustomApiPropertiesDefinitionPtrOutput() CustomApiPropertiesDefinitionPtrOutput {
	return o
}

func (o CustomApiPropertiesDefinitionPtrOutput) ToCustomApiPropertiesDefinitionPtrOutputWithContext(ctx context.Context) CustomApiPropertiesDefinitionPtrOutput {
	return o
}

func (o CustomApiPropertiesDefinitionPtrOutput) Elem() CustomApiPropertiesDefinitionOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinition) CustomApiPropertiesDefinition {
		if v != nil {
			return *v
		}
		var ret CustomApiPropertiesDefinition
		return ret
	}).(CustomApiPropertiesDefinitionOutput)
}

func (o CustomApiPropertiesDefinitionPtrOutput) ApiDefinitions() ApiResourceDefinitionsPtrOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinition) *ApiResourceDefinitions {
		if v == nil {
			return nil
		}
		return v.ApiDefinitions
	}).(ApiResourceDefinitionsPtrOutput)
}

func (o CustomApiPropertiesDefinitionPtrOutput) ApiType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinition) *string {
		if v == nil {
			return nil
		}
		return v.ApiType
	}).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionPtrOutput) BackendService() ApiResourceBackendServicePtrOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinition) *ApiResourceBackendService {
		if v == nil {
			return nil
		}
		return v.BackendService
	}).(ApiResourceBackendServicePtrOutput)
}

func (o CustomApiPropertiesDefinitionPtrOutput) BrandColor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinition) *string {
		if v == nil {
			return nil
		}
		return v.BrandColor
	}).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionPtrOutput) Capabilities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinition) []string {
		if v == nil {
			return nil
		}
		return v.Capabilities
	}).(pulumi.StringArrayOutput)
}

func (o CustomApiPropertiesDefinitionPtrOutput) ConnectionParameters() ConnectionParameterMapOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinition) map[string]ConnectionParameter {
		if v == nil {
			return nil
		}
		return v.ConnectionParameters
	}).(ConnectionParameterMapOutput)
}

func (o CustomApiPropertiesDefinitionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinition) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionPtrOutput) IconUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinition) *string {
		if v == nil {
			return nil
		}
		return v.IconUri
	}).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionPtrOutput) RuntimeUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinition) []string {
		if v == nil {
			return nil
		}
		return v.RuntimeUrls
	}).(pulumi.StringArrayOutput)
}

func (o CustomApiPropertiesDefinitionPtrOutput) Swagger() pulumi.AnyOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinition) interface{} {
		if v == nil {
			return nil
		}
		return v.Swagger
	}).(pulumi.AnyOutput)
}

func (o CustomApiPropertiesDefinitionPtrOutput) WsdlDefinition() WsdlDefinitionPtrOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinition) *WsdlDefinition {
		if v == nil {
			return nil
		}
		return v.WsdlDefinition
	}).(WsdlDefinitionPtrOutput)
}

type CustomApiPropertiesDefinitionResponse struct {
	ApiDefinitions       *ApiResourceDefinitionsResponse        `pulumi:"apiDefinitions"`
	ApiType              *string                                `pulumi:"apiType"`
	BackendService       *ApiResourceBackendServiceResponse     `pulumi:"backendService"`
	BrandColor           *string                                `pulumi:"brandColor"`
	Capabilities         []string                               `pulumi:"capabilities"`
	ConnectionParameters map[string]ConnectionParameterResponse `pulumi:"connectionParameters"`
	Description          *string                                `pulumi:"description"`
	DisplayName          *string                                `pulumi:"displayName"`
	IconUri              *string                                `pulumi:"iconUri"`
	RuntimeUrls          []string                               `pulumi:"runtimeUrls"`
	Swagger              interface{}                            `pulumi:"swagger"`
	WsdlDefinition       *WsdlDefinitionResponse                `pulumi:"wsdlDefinition"`
}





type CustomApiPropertiesDefinitionResponseInput interface {
	pulumi.Input

	ToCustomApiPropertiesDefinitionResponseOutput() CustomApiPropertiesDefinitionResponseOutput
	ToCustomApiPropertiesDefinitionResponseOutputWithContext(context.Context) CustomApiPropertiesDefinitionResponseOutput
}

type CustomApiPropertiesDefinitionResponseArgs struct {
	ApiDefinitions       ApiResourceDefinitionsResponsePtrInput    `pulumi:"apiDefinitions"`
	ApiType              pulumi.StringPtrInput                     `pulumi:"apiType"`
	BackendService       ApiResourceBackendServiceResponsePtrInput `pulumi:"backendService"`
	BrandColor           pulumi.StringPtrInput                     `pulumi:"brandColor"`
	Capabilities         pulumi.StringArrayInput                   `pulumi:"capabilities"`
	ConnectionParameters ConnectionParameterResponseMapInput       `pulumi:"connectionParameters"`
	Description          pulumi.StringPtrInput                     `pulumi:"description"`
	DisplayName          pulumi.StringPtrInput                     `pulumi:"displayName"`
	IconUri              pulumi.StringPtrInput                     `pulumi:"iconUri"`
	RuntimeUrls          pulumi.StringArrayInput                   `pulumi:"runtimeUrls"`
	Swagger              pulumi.Input                              `pulumi:"swagger"`
	WsdlDefinition       WsdlDefinitionResponsePtrInput            `pulumi:"wsdlDefinition"`
}

func (CustomApiPropertiesDefinitionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomApiPropertiesDefinitionResponse)(nil)).Elem()
}

func (i CustomApiPropertiesDefinitionResponseArgs) ToCustomApiPropertiesDefinitionResponseOutput() CustomApiPropertiesDefinitionResponseOutput {
	return i.ToCustomApiPropertiesDefinitionResponseOutputWithContext(context.Background())
}

func (i CustomApiPropertiesDefinitionResponseArgs) ToCustomApiPropertiesDefinitionResponseOutputWithContext(ctx context.Context) CustomApiPropertiesDefinitionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomApiPropertiesDefinitionResponseOutput)
}

func (i CustomApiPropertiesDefinitionResponseArgs) ToCustomApiPropertiesDefinitionResponsePtrOutput() CustomApiPropertiesDefinitionResponsePtrOutput {
	return i.ToCustomApiPropertiesDefinitionResponsePtrOutputWithContext(context.Background())
}

func (i CustomApiPropertiesDefinitionResponseArgs) ToCustomApiPropertiesDefinitionResponsePtrOutputWithContext(ctx context.Context) CustomApiPropertiesDefinitionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomApiPropertiesDefinitionResponseOutput).ToCustomApiPropertiesDefinitionResponsePtrOutputWithContext(ctx)
}









type CustomApiPropertiesDefinitionResponsePtrInput interface {
	pulumi.Input

	ToCustomApiPropertiesDefinitionResponsePtrOutput() CustomApiPropertiesDefinitionResponsePtrOutput
	ToCustomApiPropertiesDefinitionResponsePtrOutputWithContext(context.Context) CustomApiPropertiesDefinitionResponsePtrOutput
}

type customApiPropertiesDefinitionResponsePtrType CustomApiPropertiesDefinitionResponseArgs

func CustomApiPropertiesDefinitionResponsePtr(v *CustomApiPropertiesDefinitionResponseArgs) CustomApiPropertiesDefinitionResponsePtrInput {
	return (*customApiPropertiesDefinitionResponsePtrType)(v)
}

func (*customApiPropertiesDefinitionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomApiPropertiesDefinitionResponse)(nil)).Elem()
}

func (i *customApiPropertiesDefinitionResponsePtrType) ToCustomApiPropertiesDefinitionResponsePtrOutput() CustomApiPropertiesDefinitionResponsePtrOutput {
	return i.ToCustomApiPropertiesDefinitionResponsePtrOutputWithContext(context.Background())
}

func (i *customApiPropertiesDefinitionResponsePtrType) ToCustomApiPropertiesDefinitionResponsePtrOutputWithContext(ctx context.Context) CustomApiPropertiesDefinitionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomApiPropertiesDefinitionResponsePtrOutput)
}

type CustomApiPropertiesDefinitionResponseOutput struct{ *pulumi.OutputState }

func (CustomApiPropertiesDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomApiPropertiesDefinitionResponse)(nil)).Elem()
}

func (o CustomApiPropertiesDefinitionResponseOutput) ToCustomApiPropertiesDefinitionResponseOutput() CustomApiPropertiesDefinitionResponseOutput {
	return o
}

func (o CustomApiPropertiesDefinitionResponseOutput) ToCustomApiPropertiesDefinitionResponseOutputWithContext(ctx context.Context) CustomApiPropertiesDefinitionResponseOutput {
	return o
}

func (o CustomApiPropertiesDefinitionResponseOutput) ToCustomApiPropertiesDefinitionResponsePtrOutput() CustomApiPropertiesDefinitionResponsePtrOutput {
	return o.ToCustomApiPropertiesDefinitionResponsePtrOutputWithContext(context.Background())
}

func (o CustomApiPropertiesDefinitionResponseOutput) ToCustomApiPropertiesDefinitionResponsePtrOutputWithContext(ctx context.Context) CustomApiPropertiesDefinitionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomApiPropertiesDefinitionResponse) *CustomApiPropertiesDefinitionResponse {
		return &v
	}).(CustomApiPropertiesDefinitionResponsePtrOutput)
}

func (o CustomApiPropertiesDefinitionResponseOutput) ApiDefinitions() ApiResourceDefinitionsResponsePtrOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinitionResponse) *ApiResourceDefinitionsResponse { return v.ApiDefinitions }).(ApiResourceDefinitionsResponsePtrOutput)
}

func (o CustomApiPropertiesDefinitionResponseOutput) ApiType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinitionResponse) *string { return v.ApiType }).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionResponseOutput) BackendService() ApiResourceBackendServiceResponsePtrOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinitionResponse) *ApiResourceBackendServiceResponse {
		return v.BackendService
	}).(ApiResourceBackendServiceResponsePtrOutput)
}

func (o CustomApiPropertiesDefinitionResponseOutput) BrandColor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinitionResponse) *string { return v.BrandColor }).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionResponseOutput) Capabilities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinitionResponse) []string { return v.Capabilities }).(pulumi.StringArrayOutput)
}

func (o CustomApiPropertiesDefinitionResponseOutput) ConnectionParameters() ConnectionParameterResponseMapOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinitionResponse) map[string]ConnectionParameterResponse {
		return v.ConnectionParameters
	}).(ConnectionParameterResponseMapOutput)
}

func (o CustomApiPropertiesDefinitionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinitionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinitionResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionResponseOutput) IconUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinitionResponse) *string { return v.IconUri }).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionResponseOutput) RuntimeUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinitionResponse) []string { return v.RuntimeUrls }).(pulumi.StringArrayOutput)
}

func (o CustomApiPropertiesDefinitionResponseOutput) Swagger() pulumi.AnyOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinitionResponse) interface{} { return v.Swagger }).(pulumi.AnyOutput)
}

func (o CustomApiPropertiesDefinitionResponseOutput) WsdlDefinition() WsdlDefinitionResponsePtrOutput {
	return o.ApplyT(func(v CustomApiPropertiesDefinitionResponse) *WsdlDefinitionResponse { return v.WsdlDefinition }).(WsdlDefinitionResponsePtrOutput)
}

type CustomApiPropertiesDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomApiPropertiesDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomApiPropertiesDefinitionResponse)(nil)).Elem()
}

func (o CustomApiPropertiesDefinitionResponsePtrOutput) ToCustomApiPropertiesDefinitionResponsePtrOutput() CustomApiPropertiesDefinitionResponsePtrOutput {
	return o
}

func (o CustomApiPropertiesDefinitionResponsePtrOutput) ToCustomApiPropertiesDefinitionResponsePtrOutputWithContext(ctx context.Context) CustomApiPropertiesDefinitionResponsePtrOutput {
	return o
}

func (o CustomApiPropertiesDefinitionResponsePtrOutput) Elem() CustomApiPropertiesDefinitionResponseOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinitionResponse) CustomApiPropertiesDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret CustomApiPropertiesDefinitionResponse
		return ret
	}).(CustomApiPropertiesDefinitionResponseOutput)
}

func (o CustomApiPropertiesDefinitionResponsePtrOutput) ApiDefinitions() ApiResourceDefinitionsResponsePtrOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinitionResponse) *ApiResourceDefinitionsResponse {
		if v == nil {
			return nil
		}
		return v.ApiDefinitions
	}).(ApiResourceDefinitionsResponsePtrOutput)
}

func (o CustomApiPropertiesDefinitionResponsePtrOutput) ApiType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApiType
	}).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionResponsePtrOutput) BackendService() ApiResourceBackendServiceResponsePtrOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinitionResponse) *ApiResourceBackendServiceResponse {
		if v == nil {
			return nil
		}
		return v.BackendService
	}).(ApiResourceBackendServiceResponsePtrOutput)
}

func (o CustomApiPropertiesDefinitionResponsePtrOutput) BrandColor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.BrandColor
	}).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionResponsePtrOutput) Capabilities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinitionResponse) []string {
		if v == nil {
			return nil
		}
		return v.Capabilities
	}).(pulumi.StringArrayOutput)
}

func (o CustomApiPropertiesDefinitionResponsePtrOutput) ConnectionParameters() ConnectionParameterResponseMapOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinitionResponse) map[string]ConnectionParameterResponse {
		if v == nil {
			return nil
		}
		return v.ConnectionParameters
	}).(ConnectionParameterResponseMapOutput)
}

func (o CustomApiPropertiesDefinitionResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionResponsePtrOutput) IconUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.IconUri
	}).(pulumi.StringPtrOutput)
}

func (o CustomApiPropertiesDefinitionResponsePtrOutput) RuntimeUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinitionResponse) []string {
		if v == nil {
			return nil
		}
		return v.RuntimeUrls
	}).(pulumi.StringArrayOutput)
}

func (o CustomApiPropertiesDefinitionResponsePtrOutput) Swagger() pulumi.AnyOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinitionResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Swagger
	}).(pulumi.AnyOutput)
}

func (o CustomApiPropertiesDefinitionResponsePtrOutput) WsdlDefinition() WsdlDefinitionResponsePtrOutput {
	return o.ApplyT(func(v *CustomApiPropertiesDefinitionResponse) *WsdlDefinitionResponse {
		if v == nil {
			return nil
		}
		return v.WsdlDefinition
	}).(WsdlDefinitionResponsePtrOutput)
}

type WsdlDefinition struct {
	Content      *string      `pulumi:"content"`
	ImportMethod *string      `pulumi:"importMethod"`
	Service      *WsdlService `pulumi:"service"`
	Url          *string      `pulumi:"url"`
}





type WsdlDefinitionInput interface {
	pulumi.Input

	ToWsdlDefinitionOutput() WsdlDefinitionOutput
	ToWsdlDefinitionOutputWithContext(context.Context) WsdlDefinitionOutput
}

type WsdlDefinitionArgs struct {
	Content      pulumi.StringPtrInput `pulumi:"content"`
	ImportMethod pulumi.StringPtrInput `pulumi:"importMethod"`
	Service      WsdlServicePtrInput   `pulumi:"service"`
	Url          pulumi.StringPtrInput `pulumi:"url"`
}

func (WsdlDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WsdlDefinition)(nil)).Elem()
}

func (i WsdlDefinitionArgs) ToWsdlDefinitionOutput() WsdlDefinitionOutput {
	return i.ToWsdlDefinitionOutputWithContext(context.Background())
}

func (i WsdlDefinitionArgs) ToWsdlDefinitionOutputWithContext(ctx context.Context) WsdlDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WsdlDefinitionOutput)
}

func (i WsdlDefinitionArgs) ToWsdlDefinitionPtrOutput() WsdlDefinitionPtrOutput {
	return i.ToWsdlDefinitionPtrOutputWithContext(context.Background())
}

func (i WsdlDefinitionArgs) ToWsdlDefinitionPtrOutputWithContext(ctx context.Context) WsdlDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WsdlDefinitionOutput).ToWsdlDefinitionPtrOutputWithContext(ctx)
}









type WsdlDefinitionPtrInput interface {
	pulumi.Input

	ToWsdlDefinitionPtrOutput() WsdlDefinitionPtrOutput
	ToWsdlDefinitionPtrOutputWithContext(context.Context) WsdlDefinitionPtrOutput
}

type wsdlDefinitionPtrType WsdlDefinitionArgs

func WsdlDefinitionPtr(v *WsdlDefinitionArgs) WsdlDefinitionPtrInput {
	return (*wsdlDefinitionPtrType)(v)
}

func (*wsdlDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WsdlDefinition)(nil)).Elem()
}

func (i *wsdlDefinitionPtrType) ToWsdlDefinitionPtrOutput() WsdlDefinitionPtrOutput {
	return i.ToWsdlDefinitionPtrOutputWithContext(context.Background())
}

func (i *wsdlDefinitionPtrType) ToWsdlDefinitionPtrOutputWithContext(ctx context.Context) WsdlDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WsdlDefinitionPtrOutput)
}

type WsdlDefinitionOutput struct{ *pulumi.OutputState }

func (WsdlDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WsdlDefinition)(nil)).Elem()
}

func (o WsdlDefinitionOutput) ToWsdlDefinitionOutput() WsdlDefinitionOutput {
	return o
}

func (o WsdlDefinitionOutput) ToWsdlDefinitionOutputWithContext(ctx context.Context) WsdlDefinitionOutput {
	return o
}

func (o WsdlDefinitionOutput) ToWsdlDefinitionPtrOutput() WsdlDefinitionPtrOutput {
	return o.ToWsdlDefinitionPtrOutputWithContext(context.Background())
}

func (o WsdlDefinitionOutput) ToWsdlDefinitionPtrOutputWithContext(ctx context.Context) WsdlDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WsdlDefinition) *WsdlDefinition {
		return &v
	}).(WsdlDefinitionPtrOutput)
}

func (o WsdlDefinitionOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsdlDefinition) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o WsdlDefinitionOutput) ImportMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsdlDefinition) *string { return v.ImportMethod }).(pulumi.StringPtrOutput)
}

func (o WsdlDefinitionOutput) Service() WsdlServicePtrOutput {
	return o.ApplyT(func(v WsdlDefinition) *WsdlService { return v.Service }).(WsdlServicePtrOutput)
}

func (o WsdlDefinitionOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsdlDefinition) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type WsdlDefinitionPtrOutput struct{ *pulumi.OutputState }

func (WsdlDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WsdlDefinition)(nil)).Elem()
}

func (o WsdlDefinitionPtrOutput) ToWsdlDefinitionPtrOutput() WsdlDefinitionPtrOutput {
	return o
}

func (o WsdlDefinitionPtrOutput) ToWsdlDefinitionPtrOutputWithContext(ctx context.Context) WsdlDefinitionPtrOutput {
	return o
}

func (o WsdlDefinitionPtrOutput) Elem() WsdlDefinitionOutput {
	return o.ApplyT(func(v *WsdlDefinition) WsdlDefinition {
		if v != nil {
			return *v
		}
		var ret WsdlDefinition
		return ret
	}).(WsdlDefinitionOutput)
}

func (o WsdlDefinitionPtrOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsdlDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Content
	}).(pulumi.StringPtrOutput)
}

func (o WsdlDefinitionPtrOutput) ImportMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsdlDefinition) *string {
		if v == nil {
			return nil
		}
		return v.ImportMethod
	}).(pulumi.StringPtrOutput)
}

func (o WsdlDefinitionPtrOutput) Service() WsdlServicePtrOutput {
	return o.ApplyT(func(v *WsdlDefinition) *WsdlService {
		if v == nil {
			return nil
		}
		return v.Service
	}).(WsdlServicePtrOutput)
}

func (o WsdlDefinitionPtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsdlDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type WsdlDefinitionResponse struct {
	Content      *string              `pulumi:"content"`
	ImportMethod *string              `pulumi:"importMethod"`
	Service      *WsdlServiceResponse `pulumi:"service"`
	Url          *string              `pulumi:"url"`
}





type WsdlDefinitionResponseInput interface {
	pulumi.Input

	ToWsdlDefinitionResponseOutput() WsdlDefinitionResponseOutput
	ToWsdlDefinitionResponseOutputWithContext(context.Context) WsdlDefinitionResponseOutput
}

type WsdlDefinitionResponseArgs struct {
	Content      pulumi.StringPtrInput       `pulumi:"content"`
	ImportMethod pulumi.StringPtrInput       `pulumi:"importMethod"`
	Service      WsdlServiceResponsePtrInput `pulumi:"service"`
	Url          pulumi.StringPtrInput       `pulumi:"url"`
}

func (WsdlDefinitionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WsdlDefinitionResponse)(nil)).Elem()
}

func (i WsdlDefinitionResponseArgs) ToWsdlDefinitionResponseOutput() WsdlDefinitionResponseOutput {
	return i.ToWsdlDefinitionResponseOutputWithContext(context.Background())
}

func (i WsdlDefinitionResponseArgs) ToWsdlDefinitionResponseOutputWithContext(ctx context.Context) WsdlDefinitionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WsdlDefinitionResponseOutput)
}

func (i WsdlDefinitionResponseArgs) ToWsdlDefinitionResponsePtrOutput() WsdlDefinitionResponsePtrOutput {
	return i.ToWsdlDefinitionResponsePtrOutputWithContext(context.Background())
}

func (i WsdlDefinitionResponseArgs) ToWsdlDefinitionResponsePtrOutputWithContext(ctx context.Context) WsdlDefinitionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WsdlDefinitionResponseOutput).ToWsdlDefinitionResponsePtrOutputWithContext(ctx)
}









type WsdlDefinitionResponsePtrInput interface {
	pulumi.Input

	ToWsdlDefinitionResponsePtrOutput() WsdlDefinitionResponsePtrOutput
	ToWsdlDefinitionResponsePtrOutputWithContext(context.Context) WsdlDefinitionResponsePtrOutput
}

type wsdlDefinitionResponsePtrType WsdlDefinitionResponseArgs

func WsdlDefinitionResponsePtr(v *WsdlDefinitionResponseArgs) WsdlDefinitionResponsePtrInput {
	return (*wsdlDefinitionResponsePtrType)(v)
}

func (*wsdlDefinitionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WsdlDefinitionResponse)(nil)).Elem()
}

func (i *wsdlDefinitionResponsePtrType) ToWsdlDefinitionResponsePtrOutput() WsdlDefinitionResponsePtrOutput {
	return i.ToWsdlDefinitionResponsePtrOutputWithContext(context.Background())
}

func (i *wsdlDefinitionResponsePtrType) ToWsdlDefinitionResponsePtrOutputWithContext(ctx context.Context) WsdlDefinitionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WsdlDefinitionResponsePtrOutput)
}

type WsdlDefinitionResponseOutput struct{ *pulumi.OutputState }

func (WsdlDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WsdlDefinitionResponse)(nil)).Elem()
}

func (o WsdlDefinitionResponseOutput) ToWsdlDefinitionResponseOutput() WsdlDefinitionResponseOutput {
	return o
}

func (o WsdlDefinitionResponseOutput) ToWsdlDefinitionResponseOutputWithContext(ctx context.Context) WsdlDefinitionResponseOutput {
	return o
}

func (o WsdlDefinitionResponseOutput) ToWsdlDefinitionResponsePtrOutput() WsdlDefinitionResponsePtrOutput {
	return o.ToWsdlDefinitionResponsePtrOutputWithContext(context.Background())
}

func (o WsdlDefinitionResponseOutput) ToWsdlDefinitionResponsePtrOutputWithContext(ctx context.Context) WsdlDefinitionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WsdlDefinitionResponse) *WsdlDefinitionResponse {
		return &v
	}).(WsdlDefinitionResponsePtrOutput)
}

func (o WsdlDefinitionResponseOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsdlDefinitionResponse) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o WsdlDefinitionResponseOutput) ImportMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsdlDefinitionResponse) *string { return v.ImportMethod }).(pulumi.StringPtrOutput)
}

func (o WsdlDefinitionResponseOutput) Service() WsdlServiceResponsePtrOutput {
	return o.ApplyT(func(v WsdlDefinitionResponse) *WsdlServiceResponse { return v.Service }).(WsdlServiceResponsePtrOutput)
}

func (o WsdlDefinitionResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsdlDefinitionResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type WsdlDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (WsdlDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WsdlDefinitionResponse)(nil)).Elem()
}

func (o WsdlDefinitionResponsePtrOutput) ToWsdlDefinitionResponsePtrOutput() WsdlDefinitionResponsePtrOutput {
	return o
}

func (o WsdlDefinitionResponsePtrOutput) ToWsdlDefinitionResponsePtrOutputWithContext(ctx context.Context) WsdlDefinitionResponsePtrOutput {
	return o
}

func (o WsdlDefinitionResponsePtrOutput) Elem() WsdlDefinitionResponseOutput {
	return o.ApplyT(func(v *WsdlDefinitionResponse) WsdlDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret WsdlDefinitionResponse
		return ret
	}).(WsdlDefinitionResponseOutput)
}

func (o WsdlDefinitionResponsePtrOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsdlDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Content
	}).(pulumi.StringPtrOutput)
}

func (o WsdlDefinitionResponsePtrOutput) ImportMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsdlDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ImportMethod
	}).(pulumi.StringPtrOutput)
}

func (o WsdlDefinitionResponsePtrOutput) Service() WsdlServiceResponsePtrOutput {
	return o.ApplyT(func(v *WsdlDefinitionResponse) *WsdlServiceResponse {
		if v == nil {
			return nil
		}
		return v.Service
	}).(WsdlServiceResponsePtrOutput)
}

func (o WsdlDefinitionResponsePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsdlDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type WsdlService struct {
	EndpointQualifiedNames []string `pulumi:"endpointQualifiedNames"`
	QualifiedName          string   `pulumi:"qualifiedName"`
}





type WsdlServiceInput interface {
	pulumi.Input

	ToWsdlServiceOutput() WsdlServiceOutput
	ToWsdlServiceOutputWithContext(context.Context) WsdlServiceOutput
}

type WsdlServiceArgs struct {
	EndpointQualifiedNames pulumi.StringArrayInput `pulumi:"endpointQualifiedNames"`
	QualifiedName          pulumi.StringInput      `pulumi:"qualifiedName"`
}

func (WsdlServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WsdlService)(nil)).Elem()
}

func (i WsdlServiceArgs) ToWsdlServiceOutput() WsdlServiceOutput {
	return i.ToWsdlServiceOutputWithContext(context.Background())
}

func (i WsdlServiceArgs) ToWsdlServiceOutputWithContext(ctx context.Context) WsdlServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WsdlServiceOutput)
}

func (i WsdlServiceArgs) ToWsdlServicePtrOutput() WsdlServicePtrOutput {
	return i.ToWsdlServicePtrOutputWithContext(context.Background())
}

func (i WsdlServiceArgs) ToWsdlServicePtrOutputWithContext(ctx context.Context) WsdlServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WsdlServiceOutput).ToWsdlServicePtrOutputWithContext(ctx)
}









type WsdlServicePtrInput interface {
	pulumi.Input

	ToWsdlServicePtrOutput() WsdlServicePtrOutput
	ToWsdlServicePtrOutputWithContext(context.Context) WsdlServicePtrOutput
}

type wsdlServicePtrType WsdlServiceArgs

func WsdlServicePtr(v *WsdlServiceArgs) WsdlServicePtrInput {
	return (*wsdlServicePtrType)(v)
}

func (*wsdlServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WsdlService)(nil)).Elem()
}

func (i *wsdlServicePtrType) ToWsdlServicePtrOutput() WsdlServicePtrOutput {
	return i.ToWsdlServicePtrOutputWithContext(context.Background())
}

func (i *wsdlServicePtrType) ToWsdlServicePtrOutputWithContext(ctx context.Context) WsdlServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WsdlServicePtrOutput)
}

type WsdlServiceOutput struct{ *pulumi.OutputState }

func (WsdlServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WsdlService)(nil)).Elem()
}

func (o WsdlServiceOutput) ToWsdlServiceOutput() WsdlServiceOutput {
	return o
}

func (o WsdlServiceOutput) ToWsdlServiceOutputWithContext(ctx context.Context) WsdlServiceOutput {
	return o
}

func (o WsdlServiceOutput) ToWsdlServicePtrOutput() WsdlServicePtrOutput {
	return o.ToWsdlServicePtrOutputWithContext(context.Background())
}

func (o WsdlServiceOutput) ToWsdlServicePtrOutputWithContext(ctx context.Context) WsdlServicePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WsdlService) *WsdlService {
		return &v
	}).(WsdlServicePtrOutput)
}

func (o WsdlServiceOutput) EndpointQualifiedNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WsdlService) []string { return v.EndpointQualifiedNames }).(pulumi.StringArrayOutput)
}

func (o WsdlServiceOutput) QualifiedName() pulumi.StringOutput {
	return o.ApplyT(func(v WsdlService) string { return v.QualifiedName }).(pulumi.StringOutput)
}

type WsdlServicePtrOutput struct{ *pulumi.OutputState }

func (WsdlServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WsdlService)(nil)).Elem()
}

func (o WsdlServicePtrOutput) ToWsdlServicePtrOutput() WsdlServicePtrOutput {
	return o
}

func (o WsdlServicePtrOutput) ToWsdlServicePtrOutputWithContext(ctx context.Context) WsdlServicePtrOutput {
	return o
}

func (o WsdlServicePtrOutput) Elem() WsdlServiceOutput {
	return o.ApplyT(func(v *WsdlService) WsdlService {
		if v != nil {
			return *v
		}
		var ret WsdlService
		return ret
	}).(WsdlServiceOutput)
}

func (o WsdlServicePtrOutput) EndpointQualifiedNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WsdlService) []string {
		if v == nil {
			return nil
		}
		return v.EndpointQualifiedNames
	}).(pulumi.StringArrayOutput)
}

func (o WsdlServicePtrOutput) QualifiedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsdlService) *string {
		if v == nil {
			return nil
		}
		return &v.QualifiedName
	}).(pulumi.StringPtrOutput)
}

type WsdlServiceResponse struct {
	EndpointQualifiedNames []string `pulumi:"endpointQualifiedNames"`
	QualifiedName          string   `pulumi:"qualifiedName"`
}





type WsdlServiceResponseInput interface {
	pulumi.Input

	ToWsdlServiceResponseOutput() WsdlServiceResponseOutput
	ToWsdlServiceResponseOutputWithContext(context.Context) WsdlServiceResponseOutput
}

type WsdlServiceResponseArgs struct {
	EndpointQualifiedNames pulumi.StringArrayInput `pulumi:"endpointQualifiedNames"`
	QualifiedName          pulumi.StringInput      `pulumi:"qualifiedName"`
}

func (WsdlServiceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WsdlServiceResponse)(nil)).Elem()
}

func (i WsdlServiceResponseArgs) ToWsdlServiceResponseOutput() WsdlServiceResponseOutput {
	return i.ToWsdlServiceResponseOutputWithContext(context.Background())
}

func (i WsdlServiceResponseArgs) ToWsdlServiceResponseOutputWithContext(ctx context.Context) WsdlServiceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WsdlServiceResponseOutput)
}

func (i WsdlServiceResponseArgs) ToWsdlServiceResponsePtrOutput() WsdlServiceResponsePtrOutput {
	return i.ToWsdlServiceResponsePtrOutputWithContext(context.Background())
}

func (i WsdlServiceResponseArgs) ToWsdlServiceResponsePtrOutputWithContext(ctx context.Context) WsdlServiceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WsdlServiceResponseOutput).ToWsdlServiceResponsePtrOutputWithContext(ctx)
}









type WsdlServiceResponsePtrInput interface {
	pulumi.Input

	ToWsdlServiceResponsePtrOutput() WsdlServiceResponsePtrOutput
	ToWsdlServiceResponsePtrOutputWithContext(context.Context) WsdlServiceResponsePtrOutput
}

type wsdlServiceResponsePtrType WsdlServiceResponseArgs

func WsdlServiceResponsePtr(v *WsdlServiceResponseArgs) WsdlServiceResponsePtrInput {
	return (*wsdlServiceResponsePtrType)(v)
}

func (*wsdlServiceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WsdlServiceResponse)(nil)).Elem()
}

func (i *wsdlServiceResponsePtrType) ToWsdlServiceResponsePtrOutput() WsdlServiceResponsePtrOutput {
	return i.ToWsdlServiceResponsePtrOutputWithContext(context.Background())
}

func (i *wsdlServiceResponsePtrType) ToWsdlServiceResponsePtrOutputWithContext(ctx context.Context) WsdlServiceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WsdlServiceResponsePtrOutput)
}





type WsdlServiceResponseArrayInput interface {
	pulumi.Input

	ToWsdlServiceResponseArrayOutput() WsdlServiceResponseArrayOutput
	ToWsdlServiceResponseArrayOutputWithContext(context.Context) WsdlServiceResponseArrayOutput
}

type WsdlServiceResponseArray []WsdlServiceResponseInput

func (WsdlServiceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WsdlServiceResponse)(nil)).Elem()
}

func (i WsdlServiceResponseArray) ToWsdlServiceResponseArrayOutput() WsdlServiceResponseArrayOutput {
	return i.ToWsdlServiceResponseArrayOutputWithContext(context.Background())
}

func (i WsdlServiceResponseArray) ToWsdlServiceResponseArrayOutputWithContext(ctx context.Context) WsdlServiceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WsdlServiceResponseArrayOutput)
}

type WsdlServiceResponseOutput struct{ *pulumi.OutputState }

func (WsdlServiceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WsdlServiceResponse)(nil)).Elem()
}

func (o WsdlServiceResponseOutput) ToWsdlServiceResponseOutput() WsdlServiceResponseOutput {
	return o
}

func (o WsdlServiceResponseOutput) ToWsdlServiceResponseOutputWithContext(ctx context.Context) WsdlServiceResponseOutput {
	return o
}

func (o WsdlServiceResponseOutput) ToWsdlServiceResponsePtrOutput() WsdlServiceResponsePtrOutput {
	return o.ToWsdlServiceResponsePtrOutputWithContext(context.Background())
}

func (o WsdlServiceResponseOutput) ToWsdlServiceResponsePtrOutputWithContext(ctx context.Context) WsdlServiceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WsdlServiceResponse) *WsdlServiceResponse {
		return &v
	}).(WsdlServiceResponsePtrOutput)
}

func (o WsdlServiceResponseOutput) EndpointQualifiedNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WsdlServiceResponse) []string { return v.EndpointQualifiedNames }).(pulumi.StringArrayOutput)
}

func (o WsdlServiceResponseOutput) QualifiedName() pulumi.StringOutput {
	return o.ApplyT(func(v WsdlServiceResponse) string { return v.QualifiedName }).(pulumi.StringOutput)
}

type WsdlServiceResponsePtrOutput struct{ *pulumi.OutputState }

func (WsdlServiceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WsdlServiceResponse)(nil)).Elem()
}

func (o WsdlServiceResponsePtrOutput) ToWsdlServiceResponsePtrOutput() WsdlServiceResponsePtrOutput {
	return o
}

func (o WsdlServiceResponsePtrOutput) ToWsdlServiceResponsePtrOutputWithContext(ctx context.Context) WsdlServiceResponsePtrOutput {
	return o
}

func (o WsdlServiceResponsePtrOutput) Elem() WsdlServiceResponseOutput {
	return o.ApplyT(func(v *WsdlServiceResponse) WsdlServiceResponse {
		if v != nil {
			return *v
		}
		var ret WsdlServiceResponse
		return ret
	}).(WsdlServiceResponseOutput)
}

func (o WsdlServiceResponsePtrOutput) EndpointQualifiedNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WsdlServiceResponse) []string {
		if v == nil {
			return nil
		}
		return v.EndpointQualifiedNames
	}).(pulumi.StringArrayOutput)
}

func (o WsdlServiceResponsePtrOutput) QualifiedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsdlServiceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.QualifiedName
	}).(pulumi.StringPtrOutput)
}

type WsdlServiceResponseArrayOutput struct{ *pulumi.OutputState }

func (WsdlServiceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WsdlServiceResponse)(nil)).Elem()
}

func (o WsdlServiceResponseArrayOutput) ToWsdlServiceResponseArrayOutput() WsdlServiceResponseArrayOutput {
	return o
}

func (o WsdlServiceResponseArrayOutput) ToWsdlServiceResponseArrayOutputWithContext(ctx context.Context) WsdlServiceResponseArrayOutput {
	return o
}

func (o WsdlServiceResponseArrayOutput) Index(i pulumi.IntInput) WsdlServiceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WsdlServiceResponse {
		return vs[0].([]WsdlServiceResponse)[vs[1].(int)]
	}).(WsdlServiceResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiConnectionDefinitionPropertiesOutput{})
	pulumi.RegisterOutputType(ApiConnectionDefinitionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ApiConnectionDefinitionResponsePropertiesOutput{})
	pulumi.RegisterOutputType(ApiConnectionDefinitionResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ApiConnectionTestLinkOutput{})
	pulumi.RegisterOutputType(ApiConnectionTestLinkArrayOutput{})
	pulumi.RegisterOutputType(ApiConnectionTestLinkResponseOutput{})
	pulumi.RegisterOutputType(ApiConnectionTestLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(ApiOAuthSettingsOutput{})
	pulumi.RegisterOutputType(ApiOAuthSettingsPtrOutput{})
	pulumi.RegisterOutputType(ApiOAuthSettingsParameterOutput{})
	pulumi.RegisterOutputType(ApiOAuthSettingsParameterMapOutput{})
	pulumi.RegisterOutputType(ApiOAuthSettingsParameterResponseOutput{})
	pulumi.RegisterOutputType(ApiOAuthSettingsParameterResponseMapOutput{})
	pulumi.RegisterOutputType(ApiOAuthSettingsResponseOutput{})
	pulumi.RegisterOutputType(ApiOAuthSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(ApiReferenceOutput{})
	pulumi.RegisterOutputType(ApiReferencePtrOutput{})
	pulumi.RegisterOutputType(ApiReferenceResponseOutput{})
	pulumi.RegisterOutputType(ApiReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(ApiResourceBackendServiceOutput{})
	pulumi.RegisterOutputType(ApiResourceBackendServicePtrOutput{})
	pulumi.RegisterOutputType(ApiResourceBackendServiceResponseOutput{})
	pulumi.RegisterOutputType(ApiResourceBackendServiceResponsePtrOutput{})
	pulumi.RegisterOutputType(ApiResourceDefinitionsOutput{})
	pulumi.RegisterOutputType(ApiResourceDefinitionsPtrOutput{})
	pulumi.RegisterOutputType(ApiResourceDefinitionsResponseOutput{})
	pulumi.RegisterOutputType(ApiResourceDefinitionsResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectionErrorOutput{})
	pulumi.RegisterOutputType(ConnectionErrorPtrOutput{})
	pulumi.RegisterOutputType(ConnectionErrorResponseOutput{})
	pulumi.RegisterOutputType(ConnectionErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectionGatewayDefinitionPropertiesOutput{})
	pulumi.RegisterOutputType(ConnectionGatewayDefinitionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConnectionGatewayDefinitionResponsePropertiesOutput{})
	pulumi.RegisterOutputType(ConnectionGatewayDefinitionResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConnectionGatewayReferenceOutput{})
	pulumi.RegisterOutputType(ConnectionGatewayReferencePtrOutput{})
	pulumi.RegisterOutputType(ConnectionGatewayReferenceResponseOutput{})
	pulumi.RegisterOutputType(ConnectionGatewayReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectionParameterOutput{})
	pulumi.RegisterOutputType(ConnectionParameterMapOutput{})
	pulumi.RegisterOutputType(ConnectionParameterResponseOutput{})
	pulumi.RegisterOutputType(ConnectionParameterResponseMapOutput{})
	pulumi.RegisterOutputType(ConnectionStatusDefinitionOutput{})
	pulumi.RegisterOutputType(ConnectionStatusDefinitionArrayOutput{})
	pulumi.RegisterOutputType(ConnectionStatusDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ConnectionStatusDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(ConsentLinkDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ConsentLinkDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(ConsentLinkParameterDefinitionOutput{})
	pulumi.RegisterOutputType(ConsentLinkParameterDefinitionArrayOutput{})
	pulumi.RegisterOutputType(CustomApiPropertiesDefinitionOutput{})
	pulumi.RegisterOutputType(CustomApiPropertiesDefinitionPtrOutput{})
	pulumi.RegisterOutputType(CustomApiPropertiesDefinitionResponseOutput{})
	pulumi.RegisterOutputType(CustomApiPropertiesDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(WsdlDefinitionOutput{})
	pulumi.RegisterOutputType(WsdlDefinitionPtrOutput{})
	pulumi.RegisterOutputType(WsdlDefinitionResponseOutput{})
	pulumi.RegisterOutputType(WsdlDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(WsdlServiceOutput{})
	pulumi.RegisterOutputType(WsdlServicePtrOutput{})
	pulumi.RegisterOutputType(WsdlServiceResponseOutput{})
	pulumi.RegisterOutputType(WsdlServiceResponsePtrOutput{})
	pulumi.RegisterOutputType(WsdlServiceResponseArrayOutput{})
}
