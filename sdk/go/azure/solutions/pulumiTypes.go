


package solutions

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationArtifactResponse struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
	Uri  string `pulumi:"uri"`
}





type ApplicationArtifactResponseInput interface {
	pulumi.Input

	ToApplicationArtifactResponseOutput() ApplicationArtifactResponseOutput
	ToApplicationArtifactResponseOutputWithContext(context.Context) ApplicationArtifactResponseOutput
}

type ApplicationArtifactResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Type pulumi.StringInput `pulumi:"type"`
	Uri  pulumi.StringInput `pulumi:"uri"`
}

func (ApplicationArtifactResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationArtifactResponse)(nil)).Elem()
}

func (i ApplicationArtifactResponseArgs) ToApplicationArtifactResponseOutput() ApplicationArtifactResponseOutput {
	return i.ToApplicationArtifactResponseOutputWithContext(context.Background())
}

func (i ApplicationArtifactResponseArgs) ToApplicationArtifactResponseOutputWithContext(ctx context.Context) ApplicationArtifactResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationArtifactResponseOutput)
}





type ApplicationArtifactResponseArrayInput interface {
	pulumi.Input

	ToApplicationArtifactResponseArrayOutput() ApplicationArtifactResponseArrayOutput
	ToApplicationArtifactResponseArrayOutputWithContext(context.Context) ApplicationArtifactResponseArrayOutput
}

type ApplicationArtifactResponseArray []ApplicationArtifactResponseInput

func (ApplicationArtifactResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationArtifactResponse)(nil)).Elem()
}

func (i ApplicationArtifactResponseArray) ToApplicationArtifactResponseArrayOutput() ApplicationArtifactResponseArrayOutput {
	return i.ToApplicationArtifactResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationArtifactResponseArray) ToApplicationArtifactResponseArrayOutputWithContext(ctx context.Context) ApplicationArtifactResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationArtifactResponseArrayOutput)
}

type ApplicationArtifactResponseOutput struct{ *pulumi.OutputState }

func (ApplicationArtifactResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationArtifactResponse)(nil)).Elem()
}

func (o ApplicationArtifactResponseOutput) ToApplicationArtifactResponseOutput() ApplicationArtifactResponseOutput {
	return o
}

func (o ApplicationArtifactResponseOutput) ToApplicationArtifactResponseOutputWithContext(ctx context.Context) ApplicationArtifactResponseOutput {
	return o
}

func (o ApplicationArtifactResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationArtifactResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationArtifactResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationArtifactResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ApplicationArtifactResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationArtifactResponse) string { return v.Uri }).(pulumi.StringOutput)
}

type ApplicationArtifactResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationArtifactResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationArtifactResponse)(nil)).Elem()
}

func (o ApplicationArtifactResponseArrayOutput) ToApplicationArtifactResponseArrayOutput() ApplicationArtifactResponseArrayOutput {
	return o
}

func (o ApplicationArtifactResponseArrayOutput) ToApplicationArtifactResponseArrayOutputWithContext(ctx context.Context) ApplicationArtifactResponseArrayOutput {
	return o
}

func (o ApplicationArtifactResponseArrayOutput) Index(i pulumi.IntInput) ApplicationArtifactResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationArtifactResponse {
		return vs[0].([]ApplicationArtifactResponse)[vs[1].(int)]
	}).(ApplicationArtifactResponseOutput)
}

type ApplicationAuthorization struct {
	PrincipalId      string `pulumi:"principalId"`
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
}





type ApplicationAuthorizationInput interface {
	pulumi.Input

	ToApplicationAuthorizationOutput() ApplicationAuthorizationOutput
	ToApplicationAuthorizationOutputWithContext(context.Context) ApplicationAuthorizationOutput
}

type ApplicationAuthorizationArgs struct {
	PrincipalId      pulumi.StringInput `pulumi:"principalId"`
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
}

func (ApplicationAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationAuthorization)(nil)).Elem()
}

func (i ApplicationAuthorizationArgs) ToApplicationAuthorizationOutput() ApplicationAuthorizationOutput {
	return i.ToApplicationAuthorizationOutputWithContext(context.Background())
}

func (i ApplicationAuthorizationArgs) ToApplicationAuthorizationOutputWithContext(ctx context.Context) ApplicationAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationAuthorizationOutput)
}





type ApplicationAuthorizationArrayInput interface {
	pulumi.Input

	ToApplicationAuthorizationArrayOutput() ApplicationAuthorizationArrayOutput
	ToApplicationAuthorizationArrayOutputWithContext(context.Context) ApplicationAuthorizationArrayOutput
}

type ApplicationAuthorizationArray []ApplicationAuthorizationInput

func (ApplicationAuthorizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationAuthorization)(nil)).Elem()
}

func (i ApplicationAuthorizationArray) ToApplicationAuthorizationArrayOutput() ApplicationAuthorizationArrayOutput {
	return i.ToApplicationAuthorizationArrayOutputWithContext(context.Background())
}

func (i ApplicationAuthorizationArray) ToApplicationAuthorizationArrayOutputWithContext(ctx context.Context) ApplicationAuthorizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationAuthorizationArrayOutput)
}

type ApplicationAuthorizationOutput struct{ *pulumi.OutputState }

func (ApplicationAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationAuthorization)(nil)).Elem()
}

func (o ApplicationAuthorizationOutput) ToApplicationAuthorizationOutput() ApplicationAuthorizationOutput {
	return o
}

func (o ApplicationAuthorizationOutput) ToApplicationAuthorizationOutputWithContext(ctx context.Context) ApplicationAuthorizationOutput {
	return o
}

func (o ApplicationAuthorizationOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationAuthorization) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ApplicationAuthorizationOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationAuthorization) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type ApplicationAuthorizationArrayOutput struct{ *pulumi.OutputState }

func (ApplicationAuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationAuthorization)(nil)).Elem()
}

func (o ApplicationAuthorizationArrayOutput) ToApplicationAuthorizationArrayOutput() ApplicationAuthorizationArrayOutput {
	return o
}

func (o ApplicationAuthorizationArrayOutput) ToApplicationAuthorizationArrayOutputWithContext(ctx context.Context) ApplicationAuthorizationArrayOutput {
	return o
}

func (o ApplicationAuthorizationArrayOutput) Index(i pulumi.IntInput) ApplicationAuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationAuthorization {
		return vs[0].([]ApplicationAuthorization)[vs[1].(int)]
	}).(ApplicationAuthorizationOutput)
}

type ApplicationAuthorizationResponse struct {
	PrincipalId      string `pulumi:"principalId"`
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
}





type ApplicationAuthorizationResponseInput interface {
	pulumi.Input

	ToApplicationAuthorizationResponseOutput() ApplicationAuthorizationResponseOutput
	ToApplicationAuthorizationResponseOutputWithContext(context.Context) ApplicationAuthorizationResponseOutput
}

type ApplicationAuthorizationResponseArgs struct {
	PrincipalId      pulumi.StringInput `pulumi:"principalId"`
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
}

func (ApplicationAuthorizationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationAuthorizationResponse)(nil)).Elem()
}

func (i ApplicationAuthorizationResponseArgs) ToApplicationAuthorizationResponseOutput() ApplicationAuthorizationResponseOutput {
	return i.ToApplicationAuthorizationResponseOutputWithContext(context.Background())
}

func (i ApplicationAuthorizationResponseArgs) ToApplicationAuthorizationResponseOutputWithContext(ctx context.Context) ApplicationAuthorizationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationAuthorizationResponseOutput)
}





type ApplicationAuthorizationResponseArrayInput interface {
	pulumi.Input

	ToApplicationAuthorizationResponseArrayOutput() ApplicationAuthorizationResponseArrayOutput
	ToApplicationAuthorizationResponseArrayOutputWithContext(context.Context) ApplicationAuthorizationResponseArrayOutput
}

type ApplicationAuthorizationResponseArray []ApplicationAuthorizationResponseInput

func (ApplicationAuthorizationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationAuthorizationResponse)(nil)).Elem()
}

func (i ApplicationAuthorizationResponseArray) ToApplicationAuthorizationResponseArrayOutput() ApplicationAuthorizationResponseArrayOutput {
	return i.ToApplicationAuthorizationResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationAuthorizationResponseArray) ToApplicationAuthorizationResponseArrayOutputWithContext(ctx context.Context) ApplicationAuthorizationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationAuthorizationResponseArrayOutput)
}

type ApplicationAuthorizationResponseOutput struct{ *pulumi.OutputState }

func (ApplicationAuthorizationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationAuthorizationResponse)(nil)).Elem()
}

func (o ApplicationAuthorizationResponseOutput) ToApplicationAuthorizationResponseOutput() ApplicationAuthorizationResponseOutput {
	return o
}

func (o ApplicationAuthorizationResponseOutput) ToApplicationAuthorizationResponseOutputWithContext(ctx context.Context) ApplicationAuthorizationResponseOutput {
	return o
}

func (o ApplicationAuthorizationResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationAuthorizationResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ApplicationAuthorizationResponseOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationAuthorizationResponse) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type ApplicationAuthorizationResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationAuthorizationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationAuthorizationResponse)(nil)).Elem()
}

func (o ApplicationAuthorizationResponseArrayOutput) ToApplicationAuthorizationResponseArrayOutput() ApplicationAuthorizationResponseArrayOutput {
	return o
}

func (o ApplicationAuthorizationResponseArrayOutput) ToApplicationAuthorizationResponseArrayOutputWithContext(ctx context.Context) ApplicationAuthorizationResponseArrayOutput {
	return o
}

func (o ApplicationAuthorizationResponseArrayOutput) Index(i pulumi.IntInput) ApplicationAuthorizationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationAuthorizationResponse {
		return vs[0].([]ApplicationAuthorizationResponse)[vs[1].(int)]
	}).(ApplicationAuthorizationResponseOutput)
}

type ApplicationBillingDetailsDefinitionResponse struct {
	ResourceUsageId *string `pulumi:"resourceUsageId"`
}





type ApplicationBillingDetailsDefinitionResponseInput interface {
	pulumi.Input

	ToApplicationBillingDetailsDefinitionResponseOutput() ApplicationBillingDetailsDefinitionResponseOutput
	ToApplicationBillingDetailsDefinitionResponseOutputWithContext(context.Context) ApplicationBillingDetailsDefinitionResponseOutput
}

type ApplicationBillingDetailsDefinitionResponseArgs struct {
	ResourceUsageId pulumi.StringPtrInput `pulumi:"resourceUsageId"`
}

func (ApplicationBillingDetailsDefinitionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationBillingDetailsDefinitionResponse)(nil)).Elem()
}

func (i ApplicationBillingDetailsDefinitionResponseArgs) ToApplicationBillingDetailsDefinitionResponseOutput() ApplicationBillingDetailsDefinitionResponseOutput {
	return i.ToApplicationBillingDetailsDefinitionResponseOutputWithContext(context.Background())
}

func (i ApplicationBillingDetailsDefinitionResponseArgs) ToApplicationBillingDetailsDefinitionResponseOutputWithContext(ctx context.Context) ApplicationBillingDetailsDefinitionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationBillingDetailsDefinitionResponseOutput)
}

func (i ApplicationBillingDetailsDefinitionResponseArgs) ToApplicationBillingDetailsDefinitionResponsePtrOutput() ApplicationBillingDetailsDefinitionResponsePtrOutput {
	return i.ToApplicationBillingDetailsDefinitionResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationBillingDetailsDefinitionResponseArgs) ToApplicationBillingDetailsDefinitionResponsePtrOutputWithContext(ctx context.Context) ApplicationBillingDetailsDefinitionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationBillingDetailsDefinitionResponseOutput).ToApplicationBillingDetailsDefinitionResponsePtrOutputWithContext(ctx)
}









type ApplicationBillingDetailsDefinitionResponsePtrInput interface {
	pulumi.Input

	ToApplicationBillingDetailsDefinitionResponsePtrOutput() ApplicationBillingDetailsDefinitionResponsePtrOutput
	ToApplicationBillingDetailsDefinitionResponsePtrOutputWithContext(context.Context) ApplicationBillingDetailsDefinitionResponsePtrOutput
}

type applicationBillingDetailsDefinitionResponsePtrType ApplicationBillingDetailsDefinitionResponseArgs

func ApplicationBillingDetailsDefinitionResponsePtr(v *ApplicationBillingDetailsDefinitionResponseArgs) ApplicationBillingDetailsDefinitionResponsePtrInput {
	return (*applicationBillingDetailsDefinitionResponsePtrType)(v)
}

func (*applicationBillingDetailsDefinitionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationBillingDetailsDefinitionResponse)(nil)).Elem()
}

func (i *applicationBillingDetailsDefinitionResponsePtrType) ToApplicationBillingDetailsDefinitionResponsePtrOutput() ApplicationBillingDetailsDefinitionResponsePtrOutput {
	return i.ToApplicationBillingDetailsDefinitionResponsePtrOutputWithContext(context.Background())
}

func (i *applicationBillingDetailsDefinitionResponsePtrType) ToApplicationBillingDetailsDefinitionResponsePtrOutputWithContext(ctx context.Context) ApplicationBillingDetailsDefinitionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationBillingDetailsDefinitionResponsePtrOutput)
}

type ApplicationBillingDetailsDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ApplicationBillingDetailsDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationBillingDetailsDefinitionResponse)(nil)).Elem()
}

func (o ApplicationBillingDetailsDefinitionResponseOutput) ToApplicationBillingDetailsDefinitionResponseOutput() ApplicationBillingDetailsDefinitionResponseOutput {
	return o
}

func (o ApplicationBillingDetailsDefinitionResponseOutput) ToApplicationBillingDetailsDefinitionResponseOutputWithContext(ctx context.Context) ApplicationBillingDetailsDefinitionResponseOutput {
	return o
}

func (o ApplicationBillingDetailsDefinitionResponseOutput) ToApplicationBillingDetailsDefinitionResponsePtrOutput() ApplicationBillingDetailsDefinitionResponsePtrOutput {
	return o.ToApplicationBillingDetailsDefinitionResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationBillingDetailsDefinitionResponseOutput) ToApplicationBillingDetailsDefinitionResponsePtrOutputWithContext(ctx context.Context) ApplicationBillingDetailsDefinitionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationBillingDetailsDefinitionResponse) *ApplicationBillingDetailsDefinitionResponse {
		return &v
	}).(ApplicationBillingDetailsDefinitionResponsePtrOutput)
}

func (o ApplicationBillingDetailsDefinitionResponseOutput) ResourceUsageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationBillingDetailsDefinitionResponse) *string { return v.ResourceUsageId }).(pulumi.StringPtrOutput)
}

type ApplicationBillingDetailsDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationBillingDetailsDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationBillingDetailsDefinitionResponse)(nil)).Elem()
}

func (o ApplicationBillingDetailsDefinitionResponsePtrOutput) ToApplicationBillingDetailsDefinitionResponsePtrOutput() ApplicationBillingDetailsDefinitionResponsePtrOutput {
	return o
}

func (o ApplicationBillingDetailsDefinitionResponsePtrOutput) ToApplicationBillingDetailsDefinitionResponsePtrOutputWithContext(ctx context.Context) ApplicationBillingDetailsDefinitionResponsePtrOutput {
	return o
}

func (o ApplicationBillingDetailsDefinitionResponsePtrOutput) Elem() ApplicationBillingDetailsDefinitionResponseOutput {
	return o.ApplyT(func(v *ApplicationBillingDetailsDefinitionResponse) ApplicationBillingDetailsDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationBillingDetailsDefinitionResponse
		return ret
	}).(ApplicationBillingDetailsDefinitionResponseOutput)
}

func (o ApplicationBillingDetailsDefinitionResponsePtrOutput) ResourceUsageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationBillingDetailsDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceUsageId
	}).(pulumi.StringPtrOutput)
}

type ApplicationClientDetailsResponse struct {
	ApplicationId *string `pulumi:"applicationId"`
	Oid           *string `pulumi:"oid"`
	Puid          *string `pulumi:"puid"`
}





type ApplicationClientDetailsResponseInput interface {
	pulumi.Input

	ToApplicationClientDetailsResponseOutput() ApplicationClientDetailsResponseOutput
	ToApplicationClientDetailsResponseOutputWithContext(context.Context) ApplicationClientDetailsResponseOutput
}

type ApplicationClientDetailsResponseArgs struct {
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
	Oid           pulumi.StringPtrInput `pulumi:"oid"`
	Puid          pulumi.StringPtrInput `pulumi:"puid"`
}

func (ApplicationClientDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationClientDetailsResponse)(nil)).Elem()
}

func (i ApplicationClientDetailsResponseArgs) ToApplicationClientDetailsResponseOutput() ApplicationClientDetailsResponseOutput {
	return i.ToApplicationClientDetailsResponseOutputWithContext(context.Background())
}

func (i ApplicationClientDetailsResponseArgs) ToApplicationClientDetailsResponseOutputWithContext(ctx context.Context) ApplicationClientDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationClientDetailsResponseOutput)
}

func (i ApplicationClientDetailsResponseArgs) ToApplicationClientDetailsResponsePtrOutput() ApplicationClientDetailsResponsePtrOutput {
	return i.ToApplicationClientDetailsResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationClientDetailsResponseArgs) ToApplicationClientDetailsResponsePtrOutputWithContext(ctx context.Context) ApplicationClientDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationClientDetailsResponseOutput).ToApplicationClientDetailsResponsePtrOutputWithContext(ctx)
}









type ApplicationClientDetailsResponsePtrInput interface {
	pulumi.Input

	ToApplicationClientDetailsResponsePtrOutput() ApplicationClientDetailsResponsePtrOutput
	ToApplicationClientDetailsResponsePtrOutputWithContext(context.Context) ApplicationClientDetailsResponsePtrOutput
}

type applicationClientDetailsResponsePtrType ApplicationClientDetailsResponseArgs

func ApplicationClientDetailsResponsePtr(v *ApplicationClientDetailsResponseArgs) ApplicationClientDetailsResponsePtrInput {
	return (*applicationClientDetailsResponsePtrType)(v)
}

func (*applicationClientDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationClientDetailsResponse)(nil)).Elem()
}

func (i *applicationClientDetailsResponsePtrType) ToApplicationClientDetailsResponsePtrOutput() ApplicationClientDetailsResponsePtrOutput {
	return i.ToApplicationClientDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *applicationClientDetailsResponsePtrType) ToApplicationClientDetailsResponsePtrOutputWithContext(ctx context.Context) ApplicationClientDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationClientDetailsResponsePtrOutput)
}

type ApplicationClientDetailsResponseOutput struct{ *pulumi.OutputState }

func (ApplicationClientDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationClientDetailsResponse)(nil)).Elem()
}

func (o ApplicationClientDetailsResponseOutput) ToApplicationClientDetailsResponseOutput() ApplicationClientDetailsResponseOutput {
	return o
}

func (o ApplicationClientDetailsResponseOutput) ToApplicationClientDetailsResponseOutputWithContext(ctx context.Context) ApplicationClientDetailsResponseOutput {
	return o
}

func (o ApplicationClientDetailsResponseOutput) ToApplicationClientDetailsResponsePtrOutput() ApplicationClientDetailsResponsePtrOutput {
	return o.ToApplicationClientDetailsResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationClientDetailsResponseOutput) ToApplicationClientDetailsResponsePtrOutputWithContext(ctx context.Context) ApplicationClientDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationClientDetailsResponse) *ApplicationClientDetailsResponse {
		return &v
	}).(ApplicationClientDetailsResponsePtrOutput)
}

func (o ApplicationClientDetailsResponseOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationClientDetailsResponse) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

func (o ApplicationClientDetailsResponseOutput) Oid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationClientDetailsResponse) *string { return v.Oid }).(pulumi.StringPtrOutput)
}

func (o ApplicationClientDetailsResponseOutput) Puid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationClientDetailsResponse) *string { return v.Puid }).(pulumi.StringPtrOutput)
}

type ApplicationClientDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationClientDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationClientDetailsResponse)(nil)).Elem()
}

func (o ApplicationClientDetailsResponsePtrOutput) ToApplicationClientDetailsResponsePtrOutput() ApplicationClientDetailsResponsePtrOutput {
	return o
}

func (o ApplicationClientDetailsResponsePtrOutput) ToApplicationClientDetailsResponsePtrOutputWithContext(ctx context.Context) ApplicationClientDetailsResponsePtrOutput {
	return o
}

func (o ApplicationClientDetailsResponsePtrOutput) Elem() ApplicationClientDetailsResponseOutput {
	return o.ApplyT(func(v *ApplicationClientDetailsResponse) ApplicationClientDetailsResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationClientDetailsResponse
		return ret
	}).(ApplicationClientDetailsResponseOutput)
}

func (o ApplicationClientDetailsResponsePtrOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationClientDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationClientDetailsResponsePtrOutput) Oid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationClientDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Oid
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationClientDetailsResponsePtrOutput) Puid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationClientDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Puid
	}).(pulumi.StringPtrOutput)
}

type ApplicationDefinitionArtifact struct {
	Name string                  `pulumi:"name"`
	Type ApplicationArtifactType `pulumi:"type"`
	Uri  string                  `pulumi:"uri"`
}





type ApplicationDefinitionArtifactInput interface {
	pulumi.Input

	ToApplicationDefinitionArtifactOutput() ApplicationDefinitionArtifactOutput
	ToApplicationDefinitionArtifactOutputWithContext(context.Context) ApplicationDefinitionArtifactOutput
}

type ApplicationDefinitionArtifactArgs struct {
	Name pulumi.StringInput           `pulumi:"name"`
	Type ApplicationArtifactTypeInput `pulumi:"type"`
	Uri  pulumi.StringInput           `pulumi:"uri"`
}

func (ApplicationDefinitionArtifactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationDefinitionArtifact)(nil)).Elem()
}

func (i ApplicationDefinitionArtifactArgs) ToApplicationDefinitionArtifactOutput() ApplicationDefinitionArtifactOutput {
	return i.ToApplicationDefinitionArtifactOutputWithContext(context.Background())
}

func (i ApplicationDefinitionArtifactArgs) ToApplicationDefinitionArtifactOutputWithContext(ctx context.Context) ApplicationDefinitionArtifactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationDefinitionArtifactOutput)
}





type ApplicationDefinitionArtifactArrayInput interface {
	pulumi.Input

	ToApplicationDefinitionArtifactArrayOutput() ApplicationDefinitionArtifactArrayOutput
	ToApplicationDefinitionArtifactArrayOutputWithContext(context.Context) ApplicationDefinitionArtifactArrayOutput
}

type ApplicationDefinitionArtifactArray []ApplicationDefinitionArtifactInput

func (ApplicationDefinitionArtifactArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationDefinitionArtifact)(nil)).Elem()
}

func (i ApplicationDefinitionArtifactArray) ToApplicationDefinitionArtifactArrayOutput() ApplicationDefinitionArtifactArrayOutput {
	return i.ToApplicationDefinitionArtifactArrayOutputWithContext(context.Background())
}

func (i ApplicationDefinitionArtifactArray) ToApplicationDefinitionArtifactArrayOutputWithContext(ctx context.Context) ApplicationDefinitionArtifactArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationDefinitionArtifactArrayOutput)
}

type ApplicationDefinitionArtifactOutput struct{ *pulumi.OutputState }

func (ApplicationDefinitionArtifactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationDefinitionArtifact)(nil)).Elem()
}

func (o ApplicationDefinitionArtifactOutput) ToApplicationDefinitionArtifactOutput() ApplicationDefinitionArtifactOutput {
	return o
}

func (o ApplicationDefinitionArtifactOutput) ToApplicationDefinitionArtifactOutputWithContext(ctx context.Context) ApplicationDefinitionArtifactOutput {
	return o
}

func (o ApplicationDefinitionArtifactOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationDefinitionArtifact) string { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationDefinitionArtifactOutput) Type() ApplicationArtifactTypeOutput {
	return o.ApplyT(func(v ApplicationDefinitionArtifact) ApplicationArtifactType { return v.Type }).(ApplicationArtifactTypeOutput)
}

func (o ApplicationDefinitionArtifactOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationDefinitionArtifact) string { return v.Uri }).(pulumi.StringOutput)
}

type ApplicationDefinitionArtifactArrayOutput struct{ *pulumi.OutputState }

func (ApplicationDefinitionArtifactArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationDefinitionArtifact)(nil)).Elem()
}

func (o ApplicationDefinitionArtifactArrayOutput) ToApplicationDefinitionArtifactArrayOutput() ApplicationDefinitionArtifactArrayOutput {
	return o
}

func (o ApplicationDefinitionArtifactArrayOutput) ToApplicationDefinitionArtifactArrayOutputWithContext(ctx context.Context) ApplicationDefinitionArtifactArrayOutput {
	return o
}

func (o ApplicationDefinitionArtifactArrayOutput) Index(i pulumi.IntInput) ApplicationDefinitionArtifactOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationDefinitionArtifact {
		return vs[0].([]ApplicationDefinitionArtifact)[vs[1].(int)]
	}).(ApplicationDefinitionArtifactOutput)
}

type ApplicationDefinitionArtifactResponse struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
	Uri  string `pulumi:"uri"`
}





type ApplicationDefinitionArtifactResponseInput interface {
	pulumi.Input

	ToApplicationDefinitionArtifactResponseOutput() ApplicationDefinitionArtifactResponseOutput
	ToApplicationDefinitionArtifactResponseOutputWithContext(context.Context) ApplicationDefinitionArtifactResponseOutput
}

type ApplicationDefinitionArtifactResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Type pulumi.StringInput `pulumi:"type"`
	Uri  pulumi.StringInput `pulumi:"uri"`
}

func (ApplicationDefinitionArtifactResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationDefinitionArtifactResponse)(nil)).Elem()
}

func (i ApplicationDefinitionArtifactResponseArgs) ToApplicationDefinitionArtifactResponseOutput() ApplicationDefinitionArtifactResponseOutput {
	return i.ToApplicationDefinitionArtifactResponseOutputWithContext(context.Background())
}

func (i ApplicationDefinitionArtifactResponseArgs) ToApplicationDefinitionArtifactResponseOutputWithContext(ctx context.Context) ApplicationDefinitionArtifactResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationDefinitionArtifactResponseOutput)
}





type ApplicationDefinitionArtifactResponseArrayInput interface {
	pulumi.Input

	ToApplicationDefinitionArtifactResponseArrayOutput() ApplicationDefinitionArtifactResponseArrayOutput
	ToApplicationDefinitionArtifactResponseArrayOutputWithContext(context.Context) ApplicationDefinitionArtifactResponseArrayOutput
}

type ApplicationDefinitionArtifactResponseArray []ApplicationDefinitionArtifactResponseInput

func (ApplicationDefinitionArtifactResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationDefinitionArtifactResponse)(nil)).Elem()
}

func (i ApplicationDefinitionArtifactResponseArray) ToApplicationDefinitionArtifactResponseArrayOutput() ApplicationDefinitionArtifactResponseArrayOutput {
	return i.ToApplicationDefinitionArtifactResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationDefinitionArtifactResponseArray) ToApplicationDefinitionArtifactResponseArrayOutputWithContext(ctx context.Context) ApplicationDefinitionArtifactResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationDefinitionArtifactResponseArrayOutput)
}

type ApplicationDefinitionArtifactResponseOutput struct{ *pulumi.OutputState }

func (ApplicationDefinitionArtifactResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationDefinitionArtifactResponse)(nil)).Elem()
}

func (o ApplicationDefinitionArtifactResponseOutput) ToApplicationDefinitionArtifactResponseOutput() ApplicationDefinitionArtifactResponseOutput {
	return o
}

func (o ApplicationDefinitionArtifactResponseOutput) ToApplicationDefinitionArtifactResponseOutputWithContext(ctx context.Context) ApplicationDefinitionArtifactResponseOutput {
	return o
}

func (o ApplicationDefinitionArtifactResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationDefinitionArtifactResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationDefinitionArtifactResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationDefinitionArtifactResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ApplicationDefinitionArtifactResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationDefinitionArtifactResponse) string { return v.Uri }).(pulumi.StringOutput)
}

type ApplicationDefinitionArtifactResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationDefinitionArtifactResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationDefinitionArtifactResponse)(nil)).Elem()
}

func (o ApplicationDefinitionArtifactResponseArrayOutput) ToApplicationDefinitionArtifactResponseArrayOutput() ApplicationDefinitionArtifactResponseArrayOutput {
	return o
}

func (o ApplicationDefinitionArtifactResponseArrayOutput) ToApplicationDefinitionArtifactResponseArrayOutputWithContext(ctx context.Context) ApplicationDefinitionArtifactResponseArrayOutput {
	return o
}

func (o ApplicationDefinitionArtifactResponseArrayOutput) Index(i pulumi.IntInput) ApplicationDefinitionArtifactResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationDefinitionArtifactResponse {
		return vs[0].([]ApplicationDefinitionArtifactResponse)[vs[1].(int)]
	}).(ApplicationDefinitionArtifactResponseOutput)
}

type ApplicationDeploymentPolicy struct {
	DeploymentMode string `pulumi:"deploymentMode"`
}





type ApplicationDeploymentPolicyInput interface {
	pulumi.Input

	ToApplicationDeploymentPolicyOutput() ApplicationDeploymentPolicyOutput
	ToApplicationDeploymentPolicyOutputWithContext(context.Context) ApplicationDeploymentPolicyOutput
}

type ApplicationDeploymentPolicyArgs struct {
	DeploymentMode pulumi.StringInput `pulumi:"deploymentMode"`
}

func (ApplicationDeploymentPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationDeploymentPolicy)(nil)).Elem()
}

func (i ApplicationDeploymentPolicyArgs) ToApplicationDeploymentPolicyOutput() ApplicationDeploymentPolicyOutput {
	return i.ToApplicationDeploymentPolicyOutputWithContext(context.Background())
}

func (i ApplicationDeploymentPolicyArgs) ToApplicationDeploymentPolicyOutputWithContext(ctx context.Context) ApplicationDeploymentPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationDeploymentPolicyOutput)
}

func (i ApplicationDeploymentPolicyArgs) ToApplicationDeploymentPolicyPtrOutput() ApplicationDeploymentPolicyPtrOutput {
	return i.ToApplicationDeploymentPolicyPtrOutputWithContext(context.Background())
}

func (i ApplicationDeploymentPolicyArgs) ToApplicationDeploymentPolicyPtrOutputWithContext(ctx context.Context) ApplicationDeploymentPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationDeploymentPolicyOutput).ToApplicationDeploymentPolicyPtrOutputWithContext(ctx)
}









type ApplicationDeploymentPolicyPtrInput interface {
	pulumi.Input

	ToApplicationDeploymentPolicyPtrOutput() ApplicationDeploymentPolicyPtrOutput
	ToApplicationDeploymentPolicyPtrOutputWithContext(context.Context) ApplicationDeploymentPolicyPtrOutput
}

type applicationDeploymentPolicyPtrType ApplicationDeploymentPolicyArgs

func ApplicationDeploymentPolicyPtr(v *ApplicationDeploymentPolicyArgs) ApplicationDeploymentPolicyPtrInput {
	return (*applicationDeploymentPolicyPtrType)(v)
}

func (*applicationDeploymentPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationDeploymentPolicy)(nil)).Elem()
}

func (i *applicationDeploymentPolicyPtrType) ToApplicationDeploymentPolicyPtrOutput() ApplicationDeploymentPolicyPtrOutput {
	return i.ToApplicationDeploymentPolicyPtrOutputWithContext(context.Background())
}

func (i *applicationDeploymentPolicyPtrType) ToApplicationDeploymentPolicyPtrOutputWithContext(ctx context.Context) ApplicationDeploymentPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationDeploymentPolicyPtrOutput)
}

type ApplicationDeploymentPolicyOutput struct{ *pulumi.OutputState }

func (ApplicationDeploymentPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationDeploymentPolicy)(nil)).Elem()
}

func (o ApplicationDeploymentPolicyOutput) ToApplicationDeploymentPolicyOutput() ApplicationDeploymentPolicyOutput {
	return o
}

func (o ApplicationDeploymentPolicyOutput) ToApplicationDeploymentPolicyOutputWithContext(ctx context.Context) ApplicationDeploymentPolicyOutput {
	return o
}

func (o ApplicationDeploymentPolicyOutput) ToApplicationDeploymentPolicyPtrOutput() ApplicationDeploymentPolicyPtrOutput {
	return o.ToApplicationDeploymentPolicyPtrOutputWithContext(context.Background())
}

func (o ApplicationDeploymentPolicyOutput) ToApplicationDeploymentPolicyPtrOutputWithContext(ctx context.Context) ApplicationDeploymentPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationDeploymentPolicy) *ApplicationDeploymentPolicy {
		return &v
	}).(ApplicationDeploymentPolicyPtrOutput)
}

func (o ApplicationDeploymentPolicyOutput) DeploymentMode() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationDeploymentPolicy) string { return v.DeploymentMode }).(pulumi.StringOutput)
}

type ApplicationDeploymentPolicyPtrOutput struct{ *pulumi.OutputState }

func (ApplicationDeploymentPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationDeploymentPolicy)(nil)).Elem()
}

func (o ApplicationDeploymentPolicyPtrOutput) ToApplicationDeploymentPolicyPtrOutput() ApplicationDeploymentPolicyPtrOutput {
	return o
}

func (o ApplicationDeploymentPolicyPtrOutput) ToApplicationDeploymentPolicyPtrOutputWithContext(ctx context.Context) ApplicationDeploymentPolicyPtrOutput {
	return o
}

func (o ApplicationDeploymentPolicyPtrOutput) Elem() ApplicationDeploymentPolicyOutput {
	return o.ApplyT(func(v *ApplicationDeploymentPolicy) ApplicationDeploymentPolicy {
		if v != nil {
			return *v
		}
		var ret ApplicationDeploymentPolicy
		return ret
	}).(ApplicationDeploymentPolicyOutput)
}

func (o ApplicationDeploymentPolicyPtrOutput) DeploymentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationDeploymentPolicy) *string {
		if v == nil {
			return nil
		}
		return &v.DeploymentMode
	}).(pulumi.StringPtrOutput)
}

type ApplicationDeploymentPolicyResponse struct {
	DeploymentMode string `pulumi:"deploymentMode"`
}





type ApplicationDeploymentPolicyResponseInput interface {
	pulumi.Input

	ToApplicationDeploymentPolicyResponseOutput() ApplicationDeploymentPolicyResponseOutput
	ToApplicationDeploymentPolicyResponseOutputWithContext(context.Context) ApplicationDeploymentPolicyResponseOutput
}

type ApplicationDeploymentPolicyResponseArgs struct {
	DeploymentMode pulumi.StringInput `pulumi:"deploymentMode"`
}

func (ApplicationDeploymentPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationDeploymentPolicyResponse)(nil)).Elem()
}

func (i ApplicationDeploymentPolicyResponseArgs) ToApplicationDeploymentPolicyResponseOutput() ApplicationDeploymentPolicyResponseOutput {
	return i.ToApplicationDeploymentPolicyResponseOutputWithContext(context.Background())
}

func (i ApplicationDeploymentPolicyResponseArgs) ToApplicationDeploymentPolicyResponseOutputWithContext(ctx context.Context) ApplicationDeploymentPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationDeploymentPolicyResponseOutput)
}

func (i ApplicationDeploymentPolicyResponseArgs) ToApplicationDeploymentPolicyResponsePtrOutput() ApplicationDeploymentPolicyResponsePtrOutput {
	return i.ToApplicationDeploymentPolicyResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationDeploymentPolicyResponseArgs) ToApplicationDeploymentPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationDeploymentPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationDeploymentPolicyResponseOutput).ToApplicationDeploymentPolicyResponsePtrOutputWithContext(ctx)
}









type ApplicationDeploymentPolicyResponsePtrInput interface {
	pulumi.Input

	ToApplicationDeploymentPolicyResponsePtrOutput() ApplicationDeploymentPolicyResponsePtrOutput
	ToApplicationDeploymentPolicyResponsePtrOutputWithContext(context.Context) ApplicationDeploymentPolicyResponsePtrOutput
}

type applicationDeploymentPolicyResponsePtrType ApplicationDeploymentPolicyResponseArgs

func ApplicationDeploymentPolicyResponsePtr(v *ApplicationDeploymentPolicyResponseArgs) ApplicationDeploymentPolicyResponsePtrInput {
	return (*applicationDeploymentPolicyResponsePtrType)(v)
}

func (*applicationDeploymentPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationDeploymentPolicyResponse)(nil)).Elem()
}

func (i *applicationDeploymentPolicyResponsePtrType) ToApplicationDeploymentPolicyResponsePtrOutput() ApplicationDeploymentPolicyResponsePtrOutput {
	return i.ToApplicationDeploymentPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *applicationDeploymentPolicyResponsePtrType) ToApplicationDeploymentPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationDeploymentPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationDeploymentPolicyResponsePtrOutput)
}

type ApplicationDeploymentPolicyResponseOutput struct{ *pulumi.OutputState }

func (ApplicationDeploymentPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationDeploymentPolicyResponse)(nil)).Elem()
}

func (o ApplicationDeploymentPolicyResponseOutput) ToApplicationDeploymentPolicyResponseOutput() ApplicationDeploymentPolicyResponseOutput {
	return o
}

func (o ApplicationDeploymentPolicyResponseOutput) ToApplicationDeploymentPolicyResponseOutputWithContext(ctx context.Context) ApplicationDeploymentPolicyResponseOutput {
	return o
}

func (o ApplicationDeploymentPolicyResponseOutput) ToApplicationDeploymentPolicyResponsePtrOutput() ApplicationDeploymentPolicyResponsePtrOutput {
	return o.ToApplicationDeploymentPolicyResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationDeploymentPolicyResponseOutput) ToApplicationDeploymentPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationDeploymentPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationDeploymentPolicyResponse) *ApplicationDeploymentPolicyResponse {
		return &v
	}).(ApplicationDeploymentPolicyResponsePtrOutput)
}

func (o ApplicationDeploymentPolicyResponseOutput) DeploymentMode() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationDeploymentPolicyResponse) string { return v.DeploymentMode }).(pulumi.StringOutput)
}

type ApplicationDeploymentPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationDeploymentPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationDeploymentPolicyResponse)(nil)).Elem()
}

func (o ApplicationDeploymentPolicyResponsePtrOutput) ToApplicationDeploymentPolicyResponsePtrOutput() ApplicationDeploymentPolicyResponsePtrOutput {
	return o
}

func (o ApplicationDeploymentPolicyResponsePtrOutput) ToApplicationDeploymentPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationDeploymentPolicyResponsePtrOutput {
	return o
}

func (o ApplicationDeploymentPolicyResponsePtrOutput) Elem() ApplicationDeploymentPolicyResponseOutput {
	return o.ApplyT(func(v *ApplicationDeploymentPolicyResponse) ApplicationDeploymentPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationDeploymentPolicyResponse
		return ret
	}).(ApplicationDeploymentPolicyResponseOutput)
}

func (o ApplicationDeploymentPolicyResponsePtrOutput) DeploymentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationDeploymentPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DeploymentMode
	}).(pulumi.StringPtrOutput)
}

type ApplicationJitAccessPolicy struct {
	JitAccessEnabled         bool                    `pulumi:"jitAccessEnabled"`
	JitApprovalMode          *string                 `pulumi:"jitApprovalMode"`
	JitApprovers             []JitApproverDefinition `pulumi:"jitApprovers"`
	MaximumJitAccessDuration *string                 `pulumi:"maximumJitAccessDuration"`
}





type ApplicationJitAccessPolicyInput interface {
	pulumi.Input

	ToApplicationJitAccessPolicyOutput() ApplicationJitAccessPolicyOutput
	ToApplicationJitAccessPolicyOutputWithContext(context.Context) ApplicationJitAccessPolicyOutput
}

type ApplicationJitAccessPolicyArgs struct {
	JitAccessEnabled         pulumi.BoolInput                `pulumi:"jitAccessEnabled"`
	JitApprovalMode          pulumi.StringPtrInput           `pulumi:"jitApprovalMode"`
	JitApprovers             JitApproverDefinitionArrayInput `pulumi:"jitApprovers"`
	MaximumJitAccessDuration pulumi.StringPtrInput           `pulumi:"maximumJitAccessDuration"`
}

func (ApplicationJitAccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationJitAccessPolicy)(nil)).Elem()
}

func (i ApplicationJitAccessPolicyArgs) ToApplicationJitAccessPolicyOutput() ApplicationJitAccessPolicyOutput {
	return i.ToApplicationJitAccessPolicyOutputWithContext(context.Background())
}

func (i ApplicationJitAccessPolicyArgs) ToApplicationJitAccessPolicyOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationJitAccessPolicyOutput)
}

func (i ApplicationJitAccessPolicyArgs) ToApplicationJitAccessPolicyPtrOutput() ApplicationJitAccessPolicyPtrOutput {
	return i.ToApplicationJitAccessPolicyPtrOutputWithContext(context.Background())
}

func (i ApplicationJitAccessPolicyArgs) ToApplicationJitAccessPolicyPtrOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationJitAccessPolicyOutput).ToApplicationJitAccessPolicyPtrOutputWithContext(ctx)
}









type ApplicationJitAccessPolicyPtrInput interface {
	pulumi.Input

	ToApplicationJitAccessPolicyPtrOutput() ApplicationJitAccessPolicyPtrOutput
	ToApplicationJitAccessPolicyPtrOutputWithContext(context.Context) ApplicationJitAccessPolicyPtrOutput
}

type applicationJitAccessPolicyPtrType ApplicationJitAccessPolicyArgs

func ApplicationJitAccessPolicyPtr(v *ApplicationJitAccessPolicyArgs) ApplicationJitAccessPolicyPtrInput {
	return (*applicationJitAccessPolicyPtrType)(v)
}

func (*applicationJitAccessPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationJitAccessPolicy)(nil)).Elem()
}

func (i *applicationJitAccessPolicyPtrType) ToApplicationJitAccessPolicyPtrOutput() ApplicationJitAccessPolicyPtrOutput {
	return i.ToApplicationJitAccessPolicyPtrOutputWithContext(context.Background())
}

func (i *applicationJitAccessPolicyPtrType) ToApplicationJitAccessPolicyPtrOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationJitAccessPolicyPtrOutput)
}

type ApplicationJitAccessPolicyOutput struct{ *pulumi.OutputState }

func (ApplicationJitAccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationJitAccessPolicy)(nil)).Elem()
}

func (o ApplicationJitAccessPolicyOutput) ToApplicationJitAccessPolicyOutput() ApplicationJitAccessPolicyOutput {
	return o
}

func (o ApplicationJitAccessPolicyOutput) ToApplicationJitAccessPolicyOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyOutput {
	return o
}

func (o ApplicationJitAccessPolicyOutput) ToApplicationJitAccessPolicyPtrOutput() ApplicationJitAccessPolicyPtrOutput {
	return o.ToApplicationJitAccessPolicyPtrOutputWithContext(context.Background())
}

func (o ApplicationJitAccessPolicyOutput) ToApplicationJitAccessPolicyPtrOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationJitAccessPolicy) *ApplicationJitAccessPolicy {
		return &v
	}).(ApplicationJitAccessPolicyPtrOutput)
}

func (o ApplicationJitAccessPolicyOutput) JitAccessEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ApplicationJitAccessPolicy) bool { return v.JitAccessEnabled }).(pulumi.BoolOutput)
}

func (o ApplicationJitAccessPolicyOutput) JitApprovalMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationJitAccessPolicy) *string { return v.JitApprovalMode }).(pulumi.StringPtrOutput)
}

func (o ApplicationJitAccessPolicyOutput) JitApprovers() JitApproverDefinitionArrayOutput {
	return o.ApplyT(func(v ApplicationJitAccessPolicy) []JitApproverDefinition { return v.JitApprovers }).(JitApproverDefinitionArrayOutput)
}

func (o ApplicationJitAccessPolicyOutput) MaximumJitAccessDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationJitAccessPolicy) *string { return v.MaximumJitAccessDuration }).(pulumi.StringPtrOutput)
}

type ApplicationJitAccessPolicyPtrOutput struct{ *pulumi.OutputState }

func (ApplicationJitAccessPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationJitAccessPolicy)(nil)).Elem()
}

func (o ApplicationJitAccessPolicyPtrOutput) ToApplicationJitAccessPolicyPtrOutput() ApplicationJitAccessPolicyPtrOutput {
	return o
}

func (o ApplicationJitAccessPolicyPtrOutput) ToApplicationJitAccessPolicyPtrOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyPtrOutput {
	return o
}

func (o ApplicationJitAccessPolicyPtrOutput) Elem() ApplicationJitAccessPolicyOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicy) ApplicationJitAccessPolicy {
		if v != nil {
			return *v
		}
		var ret ApplicationJitAccessPolicy
		return ret
	}).(ApplicationJitAccessPolicyOutput)
}

func (o ApplicationJitAccessPolicyPtrOutput) JitAccessEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicy) *bool {
		if v == nil {
			return nil
		}
		return &v.JitAccessEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationJitAccessPolicyPtrOutput) JitApprovalMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicy) *string {
		if v == nil {
			return nil
		}
		return v.JitApprovalMode
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationJitAccessPolicyPtrOutput) JitApprovers() JitApproverDefinitionArrayOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicy) []JitApproverDefinition {
		if v == nil {
			return nil
		}
		return v.JitApprovers
	}).(JitApproverDefinitionArrayOutput)
}

func (o ApplicationJitAccessPolicyPtrOutput) MaximumJitAccessDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicy) *string {
		if v == nil {
			return nil
		}
		return v.MaximumJitAccessDuration
	}).(pulumi.StringPtrOutput)
}

type ApplicationJitAccessPolicyResponse struct {
	JitAccessEnabled         bool                            `pulumi:"jitAccessEnabled"`
	JitApprovalMode          *string                         `pulumi:"jitApprovalMode"`
	JitApprovers             []JitApproverDefinitionResponse `pulumi:"jitApprovers"`
	MaximumJitAccessDuration *string                         `pulumi:"maximumJitAccessDuration"`
}





type ApplicationJitAccessPolicyResponseInput interface {
	pulumi.Input

	ToApplicationJitAccessPolicyResponseOutput() ApplicationJitAccessPolicyResponseOutput
	ToApplicationJitAccessPolicyResponseOutputWithContext(context.Context) ApplicationJitAccessPolicyResponseOutput
}

type ApplicationJitAccessPolicyResponseArgs struct {
	JitAccessEnabled         pulumi.BoolInput                        `pulumi:"jitAccessEnabled"`
	JitApprovalMode          pulumi.StringPtrInput                   `pulumi:"jitApprovalMode"`
	JitApprovers             JitApproverDefinitionResponseArrayInput `pulumi:"jitApprovers"`
	MaximumJitAccessDuration pulumi.StringPtrInput                   `pulumi:"maximumJitAccessDuration"`
}

func (ApplicationJitAccessPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationJitAccessPolicyResponse)(nil)).Elem()
}

func (i ApplicationJitAccessPolicyResponseArgs) ToApplicationJitAccessPolicyResponseOutput() ApplicationJitAccessPolicyResponseOutput {
	return i.ToApplicationJitAccessPolicyResponseOutputWithContext(context.Background())
}

func (i ApplicationJitAccessPolicyResponseArgs) ToApplicationJitAccessPolicyResponseOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationJitAccessPolicyResponseOutput)
}

func (i ApplicationJitAccessPolicyResponseArgs) ToApplicationJitAccessPolicyResponsePtrOutput() ApplicationJitAccessPolicyResponsePtrOutput {
	return i.ToApplicationJitAccessPolicyResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationJitAccessPolicyResponseArgs) ToApplicationJitAccessPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationJitAccessPolicyResponseOutput).ToApplicationJitAccessPolicyResponsePtrOutputWithContext(ctx)
}









type ApplicationJitAccessPolicyResponsePtrInput interface {
	pulumi.Input

	ToApplicationJitAccessPolicyResponsePtrOutput() ApplicationJitAccessPolicyResponsePtrOutput
	ToApplicationJitAccessPolicyResponsePtrOutputWithContext(context.Context) ApplicationJitAccessPolicyResponsePtrOutput
}

type applicationJitAccessPolicyResponsePtrType ApplicationJitAccessPolicyResponseArgs

func ApplicationJitAccessPolicyResponsePtr(v *ApplicationJitAccessPolicyResponseArgs) ApplicationJitAccessPolicyResponsePtrInput {
	return (*applicationJitAccessPolicyResponsePtrType)(v)
}

func (*applicationJitAccessPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationJitAccessPolicyResponse)(nil)).Elem()
}

func (i *applicationJitAccessPolicyResponsePtrType) ToApplicationJitAccessPolicyResponsePtrOutput() ApplicationJitAccessPolicyResponsePtrOutput {
	return i.ToApplicationJitAccessPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *applicationJitAccessPolicyResponsePtrType) ToApplicationJitAccessPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationJitAccessPolicyResponsePtrOutput)
}

type ApplicationJitAccessPolicyResponseOutput struct{ *pulumi.OutputState }

func (ApplicationJitAccessPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationJitAccessPolicyResponse)(nil)).Elem()
}

func (o ApplicationJitAccessPolicyResponseOutput) ToApplicationJitAccessPolicyResponseOutput() ApplicationJitAccessPolicyResponseOutput {
	return o
}

func (o ApplicationJitAccessPolicyResponseOutput) ToApplicationJitAccessPolicyResponseOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyResponseOutput {
	return o
}

func (o ApplicationJitAccessPolicyResponseOutput) ToApplicationJitAccessPolicyResponsePtrOutput() ApplicationJitAccessPolicyResponsePtrOutput {
	return o.ToApplicationJitAccessPolicyResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationJitAccessPolicyResponseOutput) ToApplicationJitAccessPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationJitAccessPolicyResponse) *ApplicationJitAccessPolicyResponse {
		return &v
	}).(ApplicationJitAccessPolicyResponsePtrOutput)
}

func (o ApplicationJitAccessPolicyResponseOutput) JitAccessEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ApplicationJitAccessPolicyResponse) bool { return v.JitAccessEnabled }).(pulumi.BoolOutput)
}

func (o ApplicationJitAccessPolicyResponseOutput) JitApprovalMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationJitAccessPolicyResponse) *string { return v.JitApprovalMode }).(pulumi.StringPtrOutput)
}

func (o ApplicationJitAccessPolicyResponseOutput) JitApprovers() JitApproverDefinitionResponseArrayOutput {
	return o.ApplyT(func(v ApplicationJitAccessPolicyResponse) []JitApproverDefinitionResponse { return v.JitApprovers }).(JitApproverDefinitionResponseArrayOutput)
}

func (o ApplicationJitAccessPolicyResponseOutput) MaximumJitAccessDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationJitAccessPolicyResponse) *string { return v.MaximumJitAccessDuration }).(pulumi.StringPtrOutput)
}

type ApplicationJitAccessPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationJitAccessPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationJitAccessPolicyResponse)(nil)).Elem()
}

func (o ApplicationJitAccessPolicyResponsePtrOutput) ToApplicationJitAccessPolicyResponsePtrOutput() ApplicationJitAccessPolicyResponsePtrOutput {
	return o
}

func (o ApplicationJitAccessPolicyResponsePtrOutput) ToApplicationJitAccessPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyResponsePtrOutput {
	return o
}

func (o ApplicationJitAccessPolicyResponsePtrOutput) Elem() ApplicationJitAccessPolicyResponseOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicyResponse) ApplicationJitAccessPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationJitAccessPolicyResponse
		return ret
	}).(ApplicationJitAccessPolicyResponseOutput)
}

func (o ApplicationJitAccessPolicyResponsePtrOutput) JitAccessEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.JitAccessEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationJitAccessPolicyResponsePtrOutput) JitApprovalMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.JitApprovalMode
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationJitAccessPolicyResponsePtrOutput) JitApprovers() JitApproverDefinitionResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicyResponse) []JitApproverDefinitionResponse {
		if v == nil {
			return nil
		}
		return v.JitApprovers
	}).(JitApproverDefinitionResponseArrayOutput)
}

func (o ApplicationJitAccessPolicyResponsePtrOutput) MaximumJitAccessDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.MaximumJitAccessDuration
	}).(pulumi.StringPtrOutput)
}

type ApplicationManagementPolicy struct {
	Mode *string `pulumi:"mode"`
}





type ApplicationManagementPolicyInput interface {
	pulumi.Input

	ToApplicationManagementPolicyOutput() ApplicationManagementPolicyOutput
	ToApplicationManagementPolicyOutputWithContext(context.Context) ApplicationManagementPolicyOutput
}

type ApplicationManagementPolicyArgs struct {
	Mode pulumi.StringPtrInput `pulumi:"mode"`
}

func (ApplicationManagementPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationManagementPolicy)(nil)).Elem()
}

func (i ApplicationManagementPolicyArgs) ToApplicationManagementPolicyOutput() ApplicationManagementPolicyOutput {
	return i.ToApplicationManagementPolicyOutputWithContext(context.Background())
}

func (i ApplicationManagementPolicyArgs) ToApplicationManagementPolicyOutputWithContext(ctx context.Context) ApplicationManagementPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationManagementPolicyOutput)
}

func (i ApplicationManagementPolicyArgs) ToApplicationManagementPolicyPtrOutput() ApplicationManagementPolicyPtrOutput {
	return i.ToApplicationManagementPolicyPtrOutputWithContext(context.Background())
}

func (i ApplicationManagementPolicyArgs) ToApplicationManagementPolicyPtrOutputWithContext(ctx context.Context) ApplicationManagementPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationManagementPolicyOutput).ToApplicationManagementPolicyPtrOutputWithContext(ctx)
}









type ApplicationManagementPolicyPtrInput interface {
	pulumi.Input

	ToApplicationManagementPolicyPtrOutput() ApplicationManagementPolicyPtrOutput
	ToApplicationManagementPolicyPtrOutputWithContext(context.Context) ApplicationManagementPolicyPtrOutput
}

type applicationManagementPolicyPtrType ApplicationManagementPolicyArgs

func ApplicationManagementPolicyPtr(v *ApplicationManagementPolicyArgs) ApplicationManagementPolicyPtrInput {
	return (*applicationManagementPolicyPtrType)(v)
}

func (*applicationManagementPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationManagementPolicy)(nil)).Elem()
}

func (i *applicationManagementPolicyPtrType) ToApplicationManagementPolicyPtrOutput() ApplicationManagementPolicyPtrOutput {
	return i.ToApplicationManagementPolicyPtrOutputWithContext(context.Background())
}

func (i *applicationManagementPolicyPtrType) ToApplicationManagementPolicyPtrOutputWithContext(ctx context.Context) ApplicationManagementPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationManagementPolicyPtrOutput)
}

type ApplicationManagementPolicyOutput struct{ *pulumi.OutputState }

func (ApplicationManagementPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationManagementPolicy)(nil)).Elem()
}

func (o ApplicationManagementPolicyOutput) ToApplicationManagementPolicyOutput() ApplicationManagementPolicyOutput {
	return o
}

func (o ApplicationManagementPolicyOutput) ToApplicationManagementPolicyOutputWithContext(ctx context.Context) ApplicationManagementPolicyOutput {
	return o
}

func (o ApplicationManagementPolicyOutput) ToApplicationManagementPolicyPtrOutput() ApplicationManagementPolicyPtrOutput {
	return o.ToApplicationManagementPolicyPtrOutputWithContext(context.Background())
}

func (o ApplicationManagementPolicyOutput) ToApplicationManagementPolicyPtrOutputWithContext(ctx context.Context) ApplicationManagementPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationManagementPolicy) *ApplicationManagementPolicy {
		return &v
	}).(ApplicationManagementPolicyPtrOutput)
}

func (o ApplicationManagementPolicyOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationManagementPolicy) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

type ApplicationManagementPolicyPtrOutput struct{ *pulumi.OutputState }

func (ApplicationManagementPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationManagementPolicy)(nil)).Elem()
}

func (o ApplicationManagementPolicyPtrOutput) ToApplicationManagementPolicyPtrOutput() ApplicationManagementPolicyPtrOutput {
	return o
}

func (o ApplicationManagementPolicyPtrOutput) ToApplicationManagementPolicyPtrOutputWithContext(ctx context.Context) ApplicationManagementPolicyPtrOutput {
	return o
}

func (o ApplicationManagementPolicyPtrOutput) Elem() ApplicationManagementPolicyOutput {
	return o.ApplyT(func(v *ApplicationManagementPolicy) ApplicationManagementPolicy {
		if v != nil {
			return *v
		}
		var ret ApplicationManagementPolicy
		return ret
	}).(ApplicationManagementPolicyOutput)
}

func (o ApplicationManagementPolicyPtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationManagementPolicy) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

type ApplicationManagementPolicyResponse struct {
	Mode *string `pulumi:"mode"`
}





type ApplicationManagementPolicyResponseInput interface {
	pulumi.Input

	ToApplicationManagementPolicyResponseOutput() ApplicationManagementPolicyResponseOutput
	ToApplicationManagementPolicyResponseOutputWithContext(context.Context) ApplicationManagementPolicyResponseOutput
}

type ApplicationManagementPolicyResponseArgs struct {
	Mode pulumi.StringPtrInput `pulumi:"mode"`
}

func (ApplicationManagementPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationManagementPolicyResponse)(nil)).Elem()
}

func (i ApplicationManagementPolicyResponseArgs) ToApplicationManagementPolicyResponseOutput() ApplicationManagementPolicyResponseOutput {
	return i.ToApplicationManagementPolicyResponseOutputWithContext(context.Background())
}

func (i ApplicationManagementPolicyResponseArgs) ToApplicationManagementPolicyResponseOutputWithContext(ctx context.Context) ApplicationManagementPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationManagementPolicyResponseOutput)
}

func (i ApplicationManagementPolicyResponseArgs) ToApplicationManagementPolicyResponsePtrOutput() ApplicationManagementPolicyResponsePtrOutput {
	return i.ToApplicationManagementPolicyResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationManagementPolicyResponseArgs) ToApplicationManagementPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationManagementPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationManagementPolicyResponseOutput).ToApplicationManagementPolicyResponsePtrOutputWithContext(ctx)
}









type ApplicationManagementPolicyResponsePtrInput interface {
	pulumi.Input

	ToApplicationManagementPolicyResponsePtrOutput() ApplicationManagementPolicyResponsePtrOutput
	ToApplicationManagementPolicyResponsePtrOutputWithContext(context.Context) ApplicationManagementPolicyResponsePtrOutput
}

type applicationManagementPolicyResponsePtrType ApplicationManagementPolicyResponseArgs

func ApplicationManagementPolicyResponsePtr(v *ApplicationManagementPolicyResponseArgs) ApplicationManagementPolicyResponsePtrInput {
	return (*applicationManagementPolicyResponsePtrType)(v)
}

func (*applicationManagementPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationManagementPolicyResponse)(nil)).Elem()
}

func (i *applicationManagementPolicyResponsePtrType) ToApplicationManagementPolicyResponsePtrOutput() ApplicationManagementPolicyResponsePtrOutput {
	return i.ToApplicationManagementPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *applicationManagementPolicyResponsePtrType) ToApplicationManagementPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationManagementPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationManagementPolicyResponsePtrOutput)
}

type ApplicationManagementPolicyResponseOutput struct{ *pulumi.OutputState }

func (ApplicationManagementPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationManagementPolicyResponse)(nil)).Elem()
}

func (o ApplicationManagementPolicyResponseOutput) ToApplicationManagementPolicyResponseOutput() ApplicationManagementPolicyResponseOutput {
	return o
}

func (o ApplicationManagementPolicyResponseOutput) ToApplicationManagementPolicyResponseOutputWithContext(ctx context.Context) ApplicationManagementPolicyResponseOutput {
	return o
}

func (o ApplicationManagementPolicyResponseOutput) ToApplicationManagementPolicyResponsePtrOutput() ApplicationManagementPolicyResponsePtrOutput {
	return o.ToApplicationManagementPolicyResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationManagementPolicyResponseOutput) ToApplicationManagementPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationManagementPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationManagementPolicyResponse) *ApplicationManagementPolicyResponse {
		return &v
	}).(ApplicationManagementPolicyResponsePtrOutput)
}

func (o ApplicationManagementPolicyResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationManagementPolicyResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

type ApplicationManagementPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationManagementPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationManagementPolicyResponse)(nil)).Elem()
}

func (o ApplicationManagementPolicyResponsePtrOutput) ToApplicationManagementPolicyResponsePtrOutput() ApplicationManagementPolicyResponsePtrOutput {
	return o
}

func (o ApplicationManagementPolicyResponsePtrOutput) ToApplicationManagementPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationManagementPolicyResponsePtrOutput {
	return o
}

func (o ApplicationManagementPolicyResponsePtrOutput) Elem() ApplicationManagementPolicyResponseOutput {
	return o.ApplyT(func(v *ApplicationManagementPolicyResponse) ApplicationManagementPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationManagementPolicyResponse
		return ret
	}).(ApplicationManagementPolicyResponseOutput)
}

func (o ApplicationManagementPolicyResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationManagementPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

type ApplicationNotificationEndpoint struct {
	Uri string `pulumi:"uri"`
}





type ApplicationNotificationEndpointInput interface {
	pulumi.Input

	ToApplicationNotificationEndpointOutput() ApplicationNotificationEndpointOutput
	ToApplicationNotificationEndpointOutputWithContext(context.Context) ApplicationNotificationEndpointOutput
}

type ApplicationNotificationEndpointArgs struct {
	Uri pulumi.StringInput `pulumi:"uri"`
}

func (ApplicationNotificationEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationNotificationEndpoint)(nil)).Elem()
}

func (i ApplicationNotificationEndpointArgs) ToApplicationNotificationEndpointOutput() ApplicationNotificationEndpointOutput {
	return i.ToApplicationNotificationEndpointOutputWithContext(context.Background())
}

func (i ApplicationNotificationEndpointArgs) ToApplicationNotificationEndpointOutputWithContext(ctx context.Context) ApplicationNotificationEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationNotificationEndpointOutput)
}





type ApplicationNotificationEndpointArrayInput interface {
	pulumi.Input

	ToApplicationNotificationEndpointArrayOutput() ApplicationNotificationEndpointArrayOutput
	ToApplicationNotificationEndpointArrayOutputWithContext(context.Context) ApplicationNotificationEndpointArrayOutput
}

type ApplicationNotificationEndpointArray []ApplicationNotificationEndpointInput

func (ApplicationNotificationEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationNotificationEndpoint)(nil)).Elem()
}

func (i ApplicationNotificationEndpointArray) ToApplicationNotificationEndpointArrayOutput() ApplicationNotificationEndpointArrayOutput {
	return i.ToApplicationNotificationEndpointArrayOutputWithContext(context.Background())
}

func (i ApplicationNotificationEndpointArray) ToApplicationNotificationEndpointArrayOutputWithContext(ctx context.Context) ApplicationNotificationEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationNotificationEndpointArrayOutput)
}

type ApplicationNotificationEndpointOutput struct{ *pulumi.OutputState }

func (ApplicationNotificationEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationNotificationEndpoint)(nil)).Elem()
}

func (o ApplicationNotificationEndpointOutput) ToApplicationNotificationEndpointOutput() ApplicationNotificationEndpointOutput {
	return o
}

func (o ApplicationNotificationEndpointOutput) ToApplicationNotificationEndpointOutputWithContext(ctx context.Context) ApplicationNotificationEndpointOutput {
	return o
}

func (o ApplicationNotificationEndpointOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationNotificationEndpoint) string { return v.Uri }).(pulumi.StringOutput)
}

type ApplicationNotificationEndpointArrayOutput struct{ *pulumi.OutputState }

func (ApplicationNotificationEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationNotificationEndpoint)(nil)).Elem()
}

func (o ApplicationNotificationEndpointArrayOutput) ToApplicationNotificationEndpointArrayOutput() ApplicationNotificationEndpointArrayOutput {
	return o
}

func (o ApplicationNotificationEndpointArrayOutput) ToApplicationNotificationEndpointArrayOutputWithContext(ctx context.Context) ApplicationNotificationEndpointArrayOutput {
	return o
}

func (o ApplicationNotificationEndpointArrayOutput) Index(i pulumi.IntInput) ApplicationNotificationEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationNotificationEndpoint {
		return vs[0].([]ApplicationNotificationEndpoint)[vs[1].(int)]
	}).(ApplicationNotificationEndpointOutput)
}

type ApplicationNotificationEndpointResponse struct {
	Uri string `pulumi:"uri"`
}





type ApplicationNotificationEndpointResponseInput interface {
	pulumi.Input

	ToApplicationNotificationEndpointResponseOutput() ApplicationNotificationEndpointResponseOutput
	ToApplicationNotificationEndpointResponseOutputWithContext(context.Context) ApplicationNotificationEndpointResponseOutput
}

type ApplicationNotificationEndpointResponseArgs struct {
	Uri pulumi.StringInput `pulumi:"uri"`
}

func (ApplicationNotificationEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationNotificationEndpointResponse)(nil)).Elem()
}

func (i ApplicationNotificationEndpointResponseArgs) ToApplicationNotificationEndpointResponseOutput() ApplicationNotificationEndpointResponseOutput {
	return i.ToApplicationNotificationEndpointResponseOutputWithContext(context.Background())
}

func (i ApplicationNotificationEndpointResponseArgs) ToApplicationNotificationEndpointResponseOutputWithContext(ctx context.Context) ApplicationNotificationEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationNotificationEndpointResponseOutput)
}





type ApplicationNotificationEndpointResponseArrayInput interface {
	pulumi.Input

	ToApplicationNotificationEndpointResponseArrayOutput() ApplicationNotificationEndpointResponseArrayOutput
	ToApplicationNotificationEndpointResponseArrayOutputWithContext(context.Context) ApplicationNotificationEndpointResponseArrayOutput
}

type ApplicationNotificationEndpointResponseArray []ApplicationNotificationEndpointResponseInput

func (ApplicationNotificationEndpointResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationNotificationEndpointResponse)(nil)).Elem()
}

func (i ApplicationNotificationEndpointResponseArray) ToApplicationNotificationEndpointResponseArrayOutput() ApplicationNotificationEndpointResponseArrayOutput {
	return i.ToApplicationNotificationEndpointResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationNotificationEndpointResponseArray) ToApplicationNotificationEndpointResponseArrayOutputWithContext(ctx context.Context) ApplicationNotificationEndpointResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationNotificationEndpointResponseArrayOutput)
}

type ApplicationNotificationEndpointResponseOutput struct{ *pulumi.OutputState }

func (ApplicationNotificationEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationNotificationEndpointResponse)(nil)).Elem()
}

func (o ApplicationNotificationEndpointResponseOutput) ToApplicationNotificationEndpointResponseOutput() ApplicationNotificationEndpointResponseOutput {
	return o
}

func (o ApplicationNotificationEndpointResponseOutput) ToApplicationNotificationEndpointResponseOutputWithContext(ctx context.Context) ApplicationNotificationEndpointResponseOutput {
	return o
}

func (o ApplicationNotificationEndpointResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationNotificationEndpointResponse) string { return v.Uri }).(pulumi.StringOutput)
}

type ApplicationNotificationEndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationNotificationEndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationNotificationEndpointResponse)(nil)).Elem()
}

func (o ApplicationNotificationEndpointResponseArrayOutput) ToApplicationNotificationEndpointResponseArrayOutput() ApplicationNotificationEndpointResponseArrayOutput {
	return o
}

func (o ApplicationNotificationEndpointResponseArrayOutput) ToApplicationNotificationEndpointResponseArrayOutputWithContext(ctx context.Context) ApplicationNotificationEndpointResponseArrayOutput {
	return o
}

func (o ApplicationNotificationEndpointResponseArrayOutput) Index(i pulumi.IntInput) ApplicationNotificationEndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationNotificationEndpointResponse {
		return vs[0].([]ApplicationNotificationEndpointResponse)[vs[1].(int)]
	}).(ApplicationNotificationEndpointResponseOutput)
}

type ApplicationNotificationPolicy struct {
	NotificationEndpoints []ApplicationNotificationEndpoint `pulumi:"notificationEndpoints"`
}





type ApplicationNotificationPolicyInput interface {
	pulumi.Input

	ToApplicationNotificationPolicyOutput() ApplicationNotificationPolicyOutput
	ToApplicationNotificationPolicyOutputWithContext(context.Context) ApplicationNotificationPolicyOutput
}

type ApplicationNotificationPolicyArgs struct {
	NotificationEndpoints ApplicationNotificationEndpointArrayInput `pulumi:"notificationEndpoints"`
}

func (ApplicationNotificationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationNotificationPolicy)(nil)).Elem()
}

func (i ApplicationNotificationPolicyArgs) ToApplicationNotificationPolicyOutput() ApplicationNotificationPolicyOutput {
	return i.ToApplicationNotificationPolicyOutputWithContext(context.Background())
}

func (i ApplicationNotificationPolicyArgs) ToApplicationNotificationPolicyOutputWithContext(ctx context.Context) ApplicationNotificationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationNotificationPolicyOutput)
}

func (i ApplicationNotificationPolicyArgs) ToApplicationNotificationPolicyPtrOutput() ApplicationNotificationPolicyPtrOutput {
	return i.ToApplicationNotificationPolicyPtrOutputWithContext(context.Background())
}

func (i ApplicationNotificationPolicyArgs) ToApplicationNotificationPolicyPtrOutputWithContext(ctx context.Context) ApplicationNotificationPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationNotificationPolicyOutput).ToApplicationNotificationPolicyPtrOutputWithContext(ctx)
}









type ApplicationNotificationPolicyPtrInput interface {
	pulumi.Input

	ToApplicationNotificationPolicyPtrOutput() ApplicationNotificationPolicyPtrOutput
	ToApplicationNotificationPolicyPtrOutputWithContext(context.Context) ApplicationNotificationPolicyPtrOutput
}

type applicationNotificationPolicyPtrType ApplicationNotificationPolicyArgs

func ApplicationNotificationPolicyPtr(v *ApplicationNotificationPolicyArgs) ApplicationNotificationPolicyPtrInput {
	return (*applicationNotificationPolicyPtrType)(v)
}

func (*applicationNotificationPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationNotificationPolicy)(nil)).Elem()
}

func (i *applicationNotificationPolicyPtrType) ToApplicationNotificationPolicyPtrOutput() ApplicationNotificationPolicyPtrOutput {
	return i.ToApplicationNotificationPolicyPtrOutputWithContext(context.Background())
}

func (i *applicationNotificationPolicyPtrType) ToApplicationNotificationPolicyPtrOutputWithContext(ctx context.Context) ApplicationNotificationPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationNotificationPolicyPtrOutput)
}

type ApplicationNotificationPolicyOutput struct{ *pulumi.OutputState }

func (ApplicationNotificationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationNotificationPolicy)(nil)).Elem()
}

func (o ApplicationNotificationPolicyOutput) ToApplicationNotificationPolicyOutput() ApplicationNotificationPolicyOutput {
	return o
}

func (o ApplicationNotificationPolicyOutput) ToApplicationNotificationPolicyOutputWithContext(ctx context.Context) ApplicationNotificationPolicyOutput {
	return o
}

func (o ApplicationNotificationPolicyOutput) ToApplicationNotificationPolicyPtrOutput() ApplicationNotificationPolicyPtrOutput {
	return o.ToApplicationNotificationPolicyPtrOutputWithContext(context.Background())
}

func (o ApplicationNotificationPolicyOutput) ToApplicationNotificationPolicyPtrOutputWithContext(ctx context.Context) ApplicationNotificationPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationNotificationPolicy) *ApplicationNotificationPolicy {
		return &v
	}).(ApplicationNotificationPolicyPtrOutput)
}

func (o ApplicationNotificationPolicyOutput) NotificationEndpoints() ApplicationNotificationEndpointArrayOutput {
	return o.ApplyT(func(v ApplicationNotificationPolicy) []ApplicationNotificationEndpoint {
		return v.NotificationEndpoints
	}).(ApplicationNotificationEndpointArrayOutput)
}

type ApplicationNotificationPolicyPtrOutput struct{ *pulumi.OutputState }

func (ApplicationNotificationPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationNotificationPolicy)(nil)).Elem()
}

func (o ApplicationNotificationPolicyPtrOutput) ToApplicationNotificationPolicyPtrOutput() ApplicationNotificationPolicyPtrOutput {
	return o
}

func (o ApplicationNotificationPolicyPtrOutput) ToApplicationNotificationPolicyPtrOutputWithContext(ctx context.Context) ApplicationNotificationPolicyPtrOutput {
	return o
}

func (o ApplicationNotificationPolicyPtrOutput) Elem() ApplicationNotificationPolicyOutput {
	return o.ApplyT(func(v *ApplicationNotificationPolicy) ApplicationNotificationPolicy {
		if v != nil {
			return *v
		}
		var ret ApplicationNotificationPolicy
		return ret
	}).(ApplicationNotificationPolicyOutput)
}

func (o ApplicationNotificationPolicyPtrOutput) NotificationEndpoints() ApplicationNotificationEndpointArrayOutput {
	return o.ApplyT(func(v *ApplicationNotificationPolicy) []ApplicationNotificationEndpoint {
		if v == nil {
			return nil
		}
		return v.NotificationEndpoints
	}).(ApplicationNotificationEndpointArrayOutput)
}

type ApplicationNotificationPolicyResponse struct {
	NotificationEndpoints []ApplicationNotificationEndpointResponse `pulumi:"notificationEndpoints"`
}





type ApplicationNotificationPolicyResponseInput interface {
	pulumi.Input

	ToApplicationNotificationPolicyResponseOutput() ApplicationNotificationPolicyResponseOutput
	ToApplicationNotificationPolicyResponseOutputWithContext(context.Context) ApplicationNotificationPolicyResponseOutput
}

type ApplicationNotificationPolicyResponseArgs struct {
	NotificationEndpoints ApplicationNotificationEndpointResponseArrayInput `pulumi:"notificationEndpoints"`
}

func (ApplicationNotificationPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationNotificationPolicyResponse)(nil)).Elem()
}

func (i ApplicationNotificationPolicyResponseArgs) ToApplicationNotificationPolicyResponseOutput() ApplicationNotificationPolicyResponseOutput {
	return i.ToApplicationNotificationPolicyResponseOutputWithContext(context.Background())
}

func (i ApplicationNotificationPolicyResponseArgs) ToApplicationNotificationPolicyResponseOutputWithContext(ctx context.Context) ApplicationNotificationPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationNotificationPolicyResponseOutput)
}

func (i ApplicationNotificationPolicyResponseArgs) ToApplicationNotificationPolicyResponsePtrOutput() ApplicationNotificationPolicyResponsePtrOutput {
	return i.ToApplicationNotificationPolicyResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationNotificationPolicyResponseArgs) ToApplicationNotificationPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationNotificationPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationNotificationPolicyResponseOutput).ToApplicationNotificationPolicyResponsePtrOutputWithContext(ctx)
}









type ApplicationNotificationPolicyResponsePtrInput interface {
	pulumi.Input

	ToApplicationNotificationPolicyResponsePtrOutput() ApplicationNotificationPolicyResponsePtrOutput
	ToApplicationNotificationPolicyResponsePtrOutputWithContext(context.Context) ApplicationNotificationPolicyResponsePtrOutput
}

type applicationNotificationPolicyResponsePtrType ApplicationNotificationPolicyResponseArgs

func ApplicationNotificationPolicyResponsePtr(v *ApplicationNotificationPolicyResponseArgs) ApplicationNotificationPolicyResponsePtrInput {
	return (*applicationNotificationPolicyResponsePtrType)(v)
}

func (*applicationNotificationPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationNotificationPolicyResponse)(nil)).Elem()
}

func (i *applicationNotificationPolicyResponsePtrType) ToApplicationNotificationPolicyResponsePtrOutput() ApplicationNotificationPolicyResponsePtrOutput {
	return i.ToApplicationNotificationPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *applicationNotificationPolicyResponsePtrType) ToApplicationNotificationPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationNotificationPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationNotificationPolicyResponsePtrOutput)
}

type ApplicationNotificationPolicyResponseOutput struct{ *pulumi.OutputState }

func (ApplicationNotificationPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationNotificationPolicyResponse)(nil)).Elem()
}

func (o ApplicationNotificationPolicyResponseOutput) ToApplicationNotificationPolicyResponseOutput() ApplicationNotificationPolicyResponseOutput {
	return o
}

func (o ApplicationNotificationPolicyResponseOutput) ToApplicationNotificationPolicyResponseOutputWithContext(ctx context.Context) ApplicationNotificationPolicyResponseOutput {
	return o
}

func (o ApplicationNotificationPolicyResponseOutput) ToApplicationNotificationPolicyResponsePtrOutput() ApplicationNotificationPolicyResponsePtrOutput {
	return o.ToApplicationNotificationPolicyResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationNotificationPolicyResponseOutput) ToApplicationNotificationPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationNotificationPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationNotificationPolicyResponse) *ApplicationNotificationPolicyResponse {
		return &v
	}).(ApplicationNotificationPolicyResponsePtrOutput)
}

func (o ApplicationNotificationPolicyResponseOutput) NotificationEndpoints() ApplicationNotificationEndpointResponseArrayOutput {
	return o.ApplyT(func(v ApplicationNotificationPolicyResponse) []ApplicationNotificationEndpointResponse {
		return v.NotificationEndpoints
	}).(ApplicationNotificationEndpointResponseArrayOutput)
}

type ApplicationNotificationPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationNotificationPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationNotificationPolicyResponse)(nil)).Elem()
}

func (o ApplicationNotificationPolicyResponsePtrOutput) ToApplicationNotificationPolicyResponsePtrOutput() ApplicationNotificationPolicyResponsePtrOutput {
	return o
}

func (o ApplicationNotificationPolicyResponsePtrOutput) ToApplicationNotificationPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationNotificationPolicyResponsePtrOutput {
	return o
}

func (o ApplicationNotificationPolicyResponsePtrOutput) Elem() ApplicationNotificationPolicyResponseOutput {
	return o.ApplyT(func(v *ApplicationNotificationPolicyResponse) ApplicationNotificationPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationNotificationPolicyResponse
		return ret
	}).(ApplicationNotificationPolicyResponseOutput)
}

func (o ApplicationNotificationPolicyResponsePtrOutput) NotificationEndpoints() ApplicationNotificationEndpointResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationNotificationPolicyResponse) []ApplicationNotificationEndpointResponse {
		if v == nil {
			return nil
		}
		return v.NotificationEndpoints
	}).(ApplicationNotificationEndpointResponseArrayOutput)
}

type ApplicationPackageContactResponse struct {
	ContactName *string `pulumi:"contactName"`
	Email       string  `pulumi:"email"`
	Phone       string  `pulumi:"phone"`
}





type ApplicationPackageContactResponseInput interface {
	pulumi.Input

	ToApplicationPackageContactResponseOutput() ApplicationPackageContactResponseOutput
	ToApplicationPackageContactResponseOutputWithContext(context.Context) ApplicationPackageContactResponseOutput
}

type ApplicationPackageContactResponseArgs struct {
	ContactName pulumi.StringPtrInput `pulumi:"contactName"`
	Email       pulumi.StringInput    `pulumi:"email"`
	Phone       pulumi.StringInput    `pulumi:"phone"`
}

func (ApplicationPackageContactResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackageContactResponse)(nil)).Elem()
}

func (i ApplicationPackageContactResponseArgs) ToApplicationPackageContactResponseOutput() ApplicationPackageContactResponseOutput {
	return i.ToApplicationPackageContactResponseOutputWithContext(context.Background())
}

func (i ApplicationPackageContactResponseArgs) ToApplicationPackageContactResponseOutputWithContext(ctx context.Context) ApplicationPackageContactResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPackageContactResponseOutput)
}

func (i ApplicationPackageContactResponseArgs) ToApplicationPackageContactResponsePtrOutput() ApplicationPackageContactResponsePtrOutput {
	return i.ToApplicationPackageContactResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationPackageContactResponseArgs) ToApplicationPackageContactResponsePtrOutputWithContext(ctx context.Context) ApplicationPackageContactResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPackageContactResponseOutput).ToApplicationPackageContactResponsePtrOutputWithContext(ctx)
}









type ApplicationPackageContactResponsePtrInput interface {
	pulumi.Input

	ToApplicationPackageContactResponsePtrOutput() ApplicationPackageContactResponsePtrOutput
	ToApplicationPackageContactResponsePtrOutputWithContext(context.Context) ApplicationPackageContactResponsePtrOutput
}

type applicationPackageContactResponsePtrType ApplicationPackageContactResponseArgs

func ApplicationPackageContactResponsePtr(v *ApplicationPackageContactResponseArgs) ApplicationPackageContactResponsePtrInput {
	return (*applicationPackageContactResponsePtrType)(v)
}

func (*applicationPackageContactResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPackageContactResponse)(nil)).Elem()
}

func (i *applicationPackageContactResponsePtrType) ToApplicationPackageContactResponsePtrOutput() ApplicationPackageContactResponsePtrOutput {
	return i.ToApplicationPackageContactResponsePtrOutputWithContext(context.Background())
}

func (i *applicationPackageContactResponsePtrType) ToApplicationPackageContactResponsePtrOutputWithContext(ctx context.Context) ApplicationPackageContactResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPackageContactResponsePtrOutput)
}

type ApplicationPackageContactResponseOutput struct{ *pulumi.OutputState }

func (ApplicationPackageContactResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackageContactResponse)(nil)).Elem()
}

func (o ApplicationPackageContactResponseOutput) ToApplicationPackageContactResponseOutput() ApplicationPackageContactResponseOutput {
	return o
}

func (o ApplicationPackageContactResponseOutput) ToApplicationPackageContactResponseOutputWithContext(ctx context.Context) ApplicationPackageContactResponseOutput {
	return o
}

func (o ApplicationPackageContactResponseOutput) ToApplicationPackageContactResponsePtrOutput() ApplicationPackageContactResponsePtrOutput {
	return o.ToApplicationPackageContactResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationPackageContactResponseOutput) ToApplicationPackageContactResponsePtrOutputWithContext(ctx context.Context) ApplicationPackageContactResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationPackageContactResponse) *ApplicationPackageContactResponse {
		return &v
	}).(ApplicationPackageContactResponsePtrOutput)
}

func (o ApplicationPackageContactResponseOutput) ContactName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageContactResponse) *string { return v.ContactName }).(pulumi.StringPtrOutput)
}

func (o ApplicationPackageContactResponseOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageContactResponse) string { return v.Email }).(pulumi.StringOutput)
}

func (o ApplicationPackageContactResponseOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageContactResponse) string { return v.Phone }).(pulumi.StringOutput)
}

type ApplicationPackageContactResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationPackageContactResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPackageContactResponse)(nil)).Elem()
}

func (o ApplicationPackageContactResponsePtrOutput) ToApplicationPackageContactResponsePtrOutput() ApplicationPackageContactResponsePtrOutput {
	return o
}

func (o ApplicationPackageContactResponsePtrOutput) ToApplicationPackageContactResponsePtrOutputWithContext(ctx context.Context) ApplicationPackageContactResponsePtrOutput {
	return o
}

func (o ApplicationPackageContactResponsePtrOutput) Elem() ApplicationPackageContactResponseOutput {
	return o.ApplyT(func(v *ApplicationPackageContactResponse) ApplicationPackageContactResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationPackageContactResponse
		return ret
	}).(ApplicationPackageContactResponseOutput)
}

func (o ApplicationPackageContactResponsePtrOutput) ContactName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationPackageContactResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContactName
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationPackageContactResponsePtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationPackageContactResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Email
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationPackageContactResponsePtrOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationPackageContactResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Phone
	}).(pulumi.StringPtrOutput)
}

type ApplicationPackageLockingPolicyDefinition struct {
	AllowedActions []string `pulumi:"allowedActions"`
}





type ApplicationPackageLockingPolicyDefinitionInput interface {
	pulumi.Input

	ToApplicationPackageLockingPolicyDefinitionOutput() ApplicationPackageLockingPolicyDefinitionOutput
	ToApplicationPackageLockingPolicyDefinitionOutputWithContext(context.Context) ApplicationPackageLockingPolicyDefinitionOutput
}

type ApplicationPackageLockingPolicyDefinitionArgs struct {
	AllowedActions pulumi.StringArrayInput `pulumi:"allowedActions"`
}

func (ApplicationPackageLockingPolicyDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackageLockingPolicyDefinition)(nil)).Elem()
}

func (i ApplicationPackageLockingPolicyDefinitionArgs) ToApplicationPackageLockingPolicyDefinitionOutput() ApplicationPackageLockingPolicyDefinitionOutput {
	return i.ToApplicationPackageLockingPolicyDefinitionOutputWithContext(context.Background())
}

func (i ApplicationPackageLockingPolicyDefinitionArgs) ToApplicationPackageLockingPolicyDefinitionOutputWithContext(ctx context.Context) ApplicationPackageLockingPolicyDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPackageLockingPolicyDefinitionOutput)
}

func (i ApplicationPackageLockingPolicyDefinitionArgs) ToApplicationPackageLockingPolicyDefinitionPtrOutput() ApplicationPackageLockingPolicyDefinitionPtrOutput {
	return i.ToApplicationPackageLockingPolicyDefinitionPtrOutputWithContext(context.Background())
}

func (i ApplicationPackageLockingPolicyDefinitionArgs) ToApplicationPackageLockingPolicyDefinitionPtrOutputWithContext(ctx context.Context) ApplicationPackageLockingPolicyDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPackageLockingPolicyDefinitionOutput).ToApplicationPackageLockingPolicyDefinitionPtrOutputWithContext(ctx)
}









type ApplicationPackageLockingPolicyDefinitionPtrInput interface {
	pulumi.Input

	ToApplicationPackageLockingPolicyDefinitionPtrOutput() ApplicationPackageLockingPolicyDefinitionPtrOutput
	ToApplicationPackageLockingPolicyDefinitionPtrOutputWithContext(context.Context) ApplicationPackageLockingPolicyDefinitionPtrOutput
}

type applicationPackageLockingPolicyDefinitionPtrType ApplicationPackageLockingPolicyDefinitionArgs

func ApplicationPackageLockingPolicyDefinitionPtr(v *ApplicationPackageLockingPolicyDefinitionArgs) ApplicationPackageLockingPolicyDefinitionPtrInput {
	return (*applicationPackageLockingPolicyDefinitionPtrType)(v)
}

func (*applicationPackageLockingPolicyDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPackageLockingPolicyDefinition)(nil)).Elem()
}

func (i *applicationPackageLockingPolicyDefinitionPtrType) ToApplicationPackageLockingPolicyDefinitionPtrOutput() ApplicationPackageLockingPolicyDefinitionPtrOutput {
	return i.ToApplicationPackageLockingPolicyDefinitionPtrOutputWithContext(context.Background())
}

func (i *applicationPackageLockingPolicyDefinitionPtrType) ToApplicationPackageLockingPolicyDefinitionPtrOutputWithContext(ctx context.Context) ApplicationPackageLockingPolicyDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPackageLockingPolicyDefinitionPtrOutput)
}

type ApplicationPackageLockingPolicyDefinitionOutput struct{ *pulumi.OutputState }

func (ApplicationPackageLockingPolicyDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackageLockingPolicyDefinition)(nil)).Elem()
}

func (o ApplicationPackageLockingPolicyDefinitionOutput) ToApplicationPackageLockingPolicyDefinitionOutput() ApplicationPackageLockingPolicyDefinitionOutput {
	return o
}

func (o ApplicationPackageLockingPolicyDefinitionOutput) ToApplicationPackageLockingPolicyDefinitionOutputWithContext(ctx context.Context) ApplicationPackageLockingPolicyDefinitionOutput {
	return o
}

func (o ApplicationPackageLockingPolicyDefinitionOutput) ToApplicationPackageLockingPolicyDefinitionPtrOutput() ApplicationPackageLockingPolicyDefinitionPtrOutput {
	return o.ToApplicationPackageLockingPolicyDefinitionPtrOutputWithContext(context.Background())
}

func (o ApplicationPackageLockingPolicyDefinitionOutput) ToApplicationPackageLockingPolicyDefinitionPtrOutputWithContext(ctx context.Context) ApplicationPackageLockingPolicyDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationPackageLockingPolicyDefinition) *ApplicationPackageLockingPolicyDefinition {
		return &v
	}).(ApplicationPackageLockingPolicyDefinitionPtrOutput)
}

func (o ApplicationPackageLockingPolicyDefinitionOutput) AllowedActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApplicationPackageLockingPolicyDefinition) []string { return v.AllowedActions }).(pulumi.StringArrayOutput)
}

type ApplicationPackageLockingPolicyDefinitionPtrOutput struct{ *pulumi.OutputState }

func (ApplicationPackageLockingPolicyDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPackageLockingPolicyDefinition)(nil)).Elem()
}

func (o ApplicationPackageLockingPolicyDefinitionPtrOutput) ToApplicationPackageLockingPolicyDefinitionPtrOutput() ApplicationPackageLockingPolicyDefinitionPtrOutput {
	return o
}

func (o ApplicationPackageLockingPolicyDefinitionPtrOutput) ToApplicationPackageLockingPolicyDefinitionPtrOutputWithContext(ctx context.Context) ApplicationPackageLockingPolicyDefinitionPtrOutput {
	return o
}

func (o ApplicationPackageLockingPolicyDefinitionPtrOutput) Elem() ApplicationPackageLockingPolicyDefinitionOutput {
	return o.ApplyT(func(v *ApplicationPackageLockingPolicyDefinition) ApplicationPackageLockingPolicyDefinition {
		if v != nil {
			return *v
		}
		var ret ApplicationPackageLockingPolicyDefinition
		return ret
	}).(ApplicationPackageLockingPolicyDefinitionOutput)
}

func (o ApplicationPackageLockingPolicyDefinitionPtrOutput) AllowedActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationPackageLockingPolicyDefinition) []string {
		if v == nil {
			return nil
		}
		return v.AllowedActions
	}).(pulumi.StringArrayOutput)
}

type ApplicationPackageLockingPolicyDefinitionResponse struct {
	AllowedActions []string `pulumi:"allowedActions"`
}





type ApplicationPackageLockingPolicyDefinitionResponseInput interface {
	pulumi.Input

	ToApplicationPackageLockingPolicyDefinitionResponseOutput() ApplicationPackageLockingPolicyDefinitionResponseOutput
	ToApplicationPackageLockingPolicyDefinitionResponseOutputWithContext(context.Context) ApplicationPackageLockingPolicyDefinitionResponseOutput
}

type ApplicationPackageLockingPolicyDefinitionResponseArgs struct {
	AllowedActions pulumi.StringArrayInput `pulumi:"allowedActions"`
}

func (ApplicationPackageLockingPolicyDefinitionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackageLockingPolicyDefinitionResponse)(nil)).Elem()
}

func (i ApplicationPackageLockingPolicyDefinitionResponseArgs) ToApplicationPackageLockingPolicyDefinitionResponseOutput() ApplicationPackageLockingPolicyDefinitionResponseOutput {
	return i.ToApplicationPackageLockingPolicyDefinitionResponseOutputWithContext(context.Background())
}

func (i ApplicationPackageLockingPolicyDefinitionResponseArgs) ToApplicationPackageLockingPolicyDefinitionResponseOutputWithContext(ctx context.Context) ApplicationPackageLockingPolicyDefinitionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPackageLockingPolicyDefinitionResponseOutput)
}

func (i ApplicationPackageLockingPolicyDefinitionResponseArgs) ToApplicationPackageLockingPolicyDefinitionResponsePtrOutput() ApplicationPackageLockingPolicyDefinitionResponsePtrOutput {
	return i.ToApplicationPackageLockingPolicyDefinitionResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationPackageLockingPolicyDefinitionResponseArgs) ToApplicationPackageLockingPolicyDefinitionResponsePtrOutputWithContext(ctx context.Context) ApplicationPackageLockingPolicyDefinitionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPackageLockingPolicyDefinitionResponseOutput).ToApplicationPackageLockingPolicyDefinitionResponsePtrOutputWithContext(ctx)
}









type ApplicationPackageLockingPolicyDefinitionResponsePtrInput interface {
	pulumi.Input

	ToApplicationPackageLockingPolicyDefinitionResponsePtrOutput() ApplicationPackageLockingPolicyDefinitionResponsePtrOutput
	ToApplicationPackageLockingPolicyDefinitionResponsePtrOutputWithContext(context.Context) ApplicationPackageLockingPolicyDefinitionResponsePtrOutput
}

type applicationPackageLockingPolicyDefinitionResponsePtrType ApplicationPackageLockingPolicyDefinitionResponseArgs

func ApplicationPackageLockingPolicyDefinitionResponsePtr(v *ApplicationPackageLockingPolicyDefinitionResponseArgs) ApplicationPackageLockingPolicyDefinitionResponsePtrInput {
	return (*applicationPackageLockingPolicyDefinitionResponsePtrType)(v)
}

func (*applicationPackageLockingPolicyDefinitionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPackageLockingPolicyDefinitionResponse)(nil)).Elem()
}

func (i *applicationPackageLockingPolicyDefinitionResponsePtrType) ToApplicationPackageLockingPolicyDefinitionResponsePtrOutput() ApplicationPackageLockingPolicyDefinitionResponsePtrOutput {
	return i.ToApplicationPackageLockingPolicyDefinitionResponsePtrOutputWithContext(context.Background())
}

func (i *applicationPackageLockingPolicyDefinitionResponsePtrType) ToApplicationPackageLockingPolicyDefinitionResponsePtrOutputWithContext(ctx context.Context) ApplicationPackageLockingPolicyDefinitionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPackageLockingPolicyDefinitionResponsePtrOutput)
}

type ApplicationPackageLockingPolicyDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ApplicationPackageLockingPolicyDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackageLockingPolicyDefinitionResponse)(nil)).Elem()
}

func (o ApplicationPackageLockingPolicyDefinitionResponseOutput) ToApplicationPackageLockingPolicyDefinitionResponseOutput() ApplicationPackageLockingPolicyDefinitionResponseOutput {
	return o
}

func (o ApplicationPackageLockingPolicyDefinitionResponseOutput) ToApplicationPackageLockingPolicyDefinitionResponseOutputWithContext(ctx context.Context) ApplicationPackageLockingPolicyDefinitionResponseOutput {
	return o
}

func (o ApplicationPackageLockingPolicyDefinitionResponseOutput) ToApplicationPackageLockingPolicyDefinitionResponsePtrOutput() ApplicationPackageLockingPolicyDefinitionResponsePtrOutput {
	return o.ToApplicationPackageLockingPolicyDefinitionResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationPackageLockingPolicyDefinitionResponseOutput) ToApplicationPackageLockingPolicyDefinitionResponsePtrOutputWithContext(ctx context.Context) ApplicationPackageLockingPolicyDefinitionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationPackageLockingPolicyDefinitionResponse) *ApplicationPackageLockingPolicyDefinitionResponse {
		return &v
	}).(ApplicationPackageLockingPolicyDefinitionResponsePtrOutput)
}

func (o ApplicationPackageLockingPolicyDefinitionResponseOutput) AllowedActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApplicationPackageLockingPolicyDefinitionResponse) []string { return v.AllowedActions }).(pulumi.StringArrayOutput)
}

type ApplicationPackageLockingPolicyDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationPackageLockingPolicyDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPackageLockingPolicyDefinitionResponse)(nil)).Elem()
}

func (o ApplicationPackageLockingPolicyDefinitionResponsePtrOutput) ToApplicationPackageLockingPolicyDefinitionResponsePtrOutput() ApplicationPackageLockingPolicyDefinitionResponsePtrOutput {
	return o
}

func (o ApplicationPackageLockingPolicyDefinitionResponsePtrOutput) ToApplicationPackageLockingPolicyDefinitionResponsePtrOutputWithContext(ctx context.Context) ApplicationPackageLockingPolicyDefinitionResponsePtrOutput {
	return o
}

func (o ApplicationPackageLockingPolicyDefinitionResponsePtrOutput) Elem() ApplicationPackageLockingPolicyDefinitionResponseOutput {
	return o.ApplyT(func(v *ApplicationPackageLockingPolicyDefinitionResponse) ApplicationPackageLockingPolicyDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationPackageLockingPolicyDefinitionResponse
		return ret
	}).(ApplicationPackageLockingPolicyDefinitionResponseOutput)
}

func (o ApplicationPackageLockingPolicyDefinitionResponsePtrOutput) AllowedActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationPackageLockingPolicyDefinitionResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedActions
	}).(pulumi.StringArrayOutput)
}

type ApplicationPackageSupportUrlsResponse struct {
	GovernmentCloud *string `pulumi:"governmentCloud"`
	PublicAzure     *string `pulumi:"publicAzure"`
}





type ApplicationPackageSupportUrlsResponseInput interface {
	pulumi.Input

	ToApplicationPackageSupportUrlsResponseOutput() ApplicationPackageSupportUrlsResponseOutput
	ToApplicationPackageSupportUrlsResponseOutputWithContext(context.Context) ApplicationPackageSupportUrlsResponseOutput
}

type ApplicationPackageSupportUrlsResponseArgs struct {
	GovernmentCloud pulumi.StringPtrInput `pulumi:"governmentCloud"`
	PublicAzure     pulumi.StringPtrInput `pulumi:"publicAzure"`
}

func (ApplicationPackageSupportUrlsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackageSupportUrlsResponse)(nil)).Elem()
}

func (i ApplicationPackageSupportUrlsResponseArgs) ToApplicationPackageSupportUrlsResponseOutput() ApplicationPackageSupportUrlsResponseOutput {
	return i.ToApplicationPackageSupportUrlsResponseOutputWithContext(context.Background())
}

func (i ApplicationPackageSupportUrlsResponseArgs) ToApplicationPackageSupportUrlsResponseOutputWithContext(ctx context.Context) ApplicationPackageSupportUrlsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPackageSupportUrlsResponseOutput)
}

func (i ApplicationPackageSupportUrlsResponseArgs) ToApplicationPackageSupportUrlsResponsePtrOutput() ApplicationPackageSupportUrlsResponsePtrOutput {
	return i.ToApplicationPackageSupportUrlsResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationPackageSupportUrlsResponseArgs) ToApplicationPackageSupportUrlsResponsePtrOutputWithContext(ctx context.Context) ApplicationPackageSupportUrlsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPackageSupportUrlsResponseOutput).ToApplicationPackageSupportUrlsResponsePtrOutputWithContext(ctx)
}









type ApplicationPackageSupportUrlsResponsePtrInput interface {
	pulumi.Input

	ToApplicationPackageSupportUrlsResponsePtrOutput() ApplicationPackageSupportUrlsResponsePtrOutput
	ToApplicationPackageSupportUrlsResponsePtrOutputWithContext(context.Context) ApplicationPackageSupportUrlsResponsePtrOutput
}

type applicationPackageSupportUrlsResponsePtrType ApplicationPackageSupportUrlsResponseArgs

func ApplicationPackageSupportUrlsResponsePtr(v *ApplicationPackageSupportUrlsResponseArgs) ApplicationPackageSupportUrlsResponsePtrInput {
	return (*applicationPackageSupportUrlsResponsePtrType)(v)
}

func (*applicationPackageSupportUrlsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPackageSupportUrlsResponse)(nil)).Elem()
}

func (i *applicationPackageSupportUrlsResponsePtrType) ToApplicationPackageSupportUrlsResponsePtrOutput() ApplicationPackageSupportUrlsResponsePtrOutput {
	return i.ToApplicationPackageSupportUrlsResponsePtrOutputWithContext(context.Background())
}

func (i *applicationPackageSupportUrlsResponsePtrType) ToApplicationPackageSupportUrlsResponsePtrOutputWithContext(ctx context.Context) ApplicationPackageSupportUrlsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPackageSupportUrlsResponsePtrOutput)
}

type ApplicationPackageSupportUrlsResponseOutput struct{ *pulumi.OutputState }

func (ApplicationPackageSupportUrlsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackageSupportUrlsResponse)(nil)).Elem()
}

func (o ApplicationPackageSupportUrlsResponseOutput) ToApplicationPackageSupportUrlsResponseOutput() ApplicationPackageSupportUrlsResponseOutput {
	return o
}

func (o ApplicationPackageSupportUrlsResponseOutput) ToApplicationPackageSupportUrlsResponseOutputWithContext(ctx context.Context) ApplicationPackageSupportUrlsResponseOutput {
	return o
}

func (o ApplicationPackageSupportUrlsResponseOutput) ToApplicationPackageSupportUrlsResponsePtrOutput() ApplicationPackageSupportUrlsResponsePtrOutput {
	return o.ToApplicationPackageSupportUrlsResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationPackageSupportUrlsResponseOutput) ToApplicationPackageSupportUrlsResponsePtrOutputWithContext(ctx context.Context) ApplicationPackageSupportUrlsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationPackageSupportUrlsResponse) *ApplicationPackageSupportUrlsResponse {
		return &v
	}).(ApplicationPackageSupportUrlsResponsePtrOutput)
}

func (o ApplicationPackageSupportUrlsResponseOutput) GovernmentCloud() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageSupportUrlsResponse) *string { return v.GovernmentCloud }).(pulumi.StringPtrOutput)
}

func (o ApplicationPackageSupportUrlsResponseOutput) PublicAzure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageSupportUrlsResponse) *string { return v.PublicAzure }).(pulumi.StringPtrOutput)
}

type ApplicationPackageSupportUrlsResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationPackageSupportUrlsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPackageSupportUrlsResponse)(nil)).Elem()
}

func (o ApplicationPackageSupportUrlsResponsePtrOutput) ToApplicationPackageSupportUrlsResponsePtrOutput() ApplicationPackageSupportUrlsResponsePtrOutput {
	return o
}

func (o ApplicationPackageSupportUrlsResponsePtrOutput) ToApplicationPackageSupportUrlsResponsePtrOutputWithContext(ctx context.Context) ApplicationPackageSupportUrlsResponsePtrOutput {
	return o
}

func (o ApplicationPackageSupportUrlsResponsePtrOutput) Elem() ApplicationPackageSupportUrlsResponseOutput {
	return o.ApplyT(func(v *ApplicationPackageSupportUrlsResponse) ApplicationPackageSupportUrlsResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationPackageSupportUrlsResponse
		return ret
	}).(ApplicationPackageSupportUrlsResponseOutput)
}

func (o ApplicationPackageSupportUrlsResponsePtrOutput) GovernmentCloud() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationPackageSupportUrlsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GovernmentCloud
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationPackageSupportUrlsResponsePtrOutput) PublicAzure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationPackageSupportUrlsResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicAzure
	}).(pulumi.StringPtrOutput)
}

type ApplicationPolicy struct {
	Name               *string `pulumi:"name"`
	Parameters         *string `pulumi:"parameters"`
	PolicyDefinitionId *string `pulumi:"policyDefinitionId"`
}





type ApplicationPolicyInput interface {
	pulumi.Input

	ToApplicationPolicyOutput() ApplicationPolicyOutput
	ToApplicationPolicyOutputWithContext(context.Context) ApplicationPolicyOutput
}

type ApplicationPolicyArgs struct {
	Name               pulumi.StringPtrInput `pulumi:"name"`
	Parameters         pulumi.StringPtrInput `pulumi:"parameters"`
	PolicyDefinitionId pulumi.StringPtrInput `pulumi:"policyDefinitionId"`
}

func (ApplicationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPolicy)(nil)).Elem()
}

func (i ApplicationPolicyArgs) ToApplicationPolicyOutput() ApplicationPolicyOutput {
	return i.ToApplicationPolicyOutputWithContext(context.Background())
}

func (i ApplicationPolicyArgs) ToApplicationPolicyOutputWithContext(ctx context.Context) ApplicationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPolicyOutput)
}





type ApplicationPolicyArrayInput interface {
	pulumi.Input

	ToApplicationPolicyArrayOutput() ApplicationPolicyArrayOutput
	ToApplicationPolicyArrayOutputWithContext(context.Context) ApplicationPolicyArrayOutput
}

type ApplicationPolicyArray []ApplicationPolicyInput

func (ApplicationPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationPolicy)(nil)).Elem()
}

func (i ApplicationPolicyArray) ToApplicationPolicyArrayOutput() ApplicationPolicyArrayOutput {
	return i.ToApplicationPolicyArrayOutputWithContext(context.Background())
}

func (i ApplicationPolicyArray) ToApplicationPolicyArrayOutputWithContext(ctx context.Context) ApplicationPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPolicyArrayOutput)
}

type ApplicationPolicyOutput struct{ *pulumi.OutputState }

func (ApplicationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPolicy)(nil)).Elem()
}

func (o ApplicationPolicyOutput) ToApplicationPolicyOutput() ApplicationPolicyOutput {
	return o
}

func (o ApplicationPolicyOutput) ToApplicationPolicyOutputWithContext(ctx context.Context) ApplicationPolicyOutput {
	return o
}

func (o ApplicationPolicyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPolicy) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationPolicyOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPolicy) *string { return v.Parameters }).(pulumi.StringPtrOutput)
}

func (o ApplicationPolicyOutput) PolicyDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPolicy) *string { return v.PolicyDefinitionId }).(pulumi.StringPtrOutput)
}

type ApplicationPolicyArrayOutput struct{ *pulumi.OutputState }

func (ApplicationPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationPolicy)(nil)).Elem()
}

func (o ApplicationPolicyArrayOutput) ToApplicationPolicyArrayOutput() ApplicationPolicyArrayOutput {
	return o
}

func (o ApplicationPolicyArrayOutput) ToApplicationPolicyArrayOutputWithContext(ctx context.Context) ApplicationPolicyArrayOutput {
	return o
}

func (o ApplicationPolicyArrayOutput) Index(i pulumi.IntInput) ApplicationPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationPolicy {
		return vs[0].([]ApplicationPolicy)[vs[1].(int)]
	}).(ApplicationPolicyOutput)
}

type ApplicationPolicyResponse struct {
	Name               *string `pulumi:"name"`
	Parameters         *string `pulumi:"parameters"`
	PolicyDefinitionId *string `pulumi:"policyDefinitionId"`
}





type ApplicationPolicyResponseInput interface {
	pulumi.Input

	ToApplicationPolicyResponseOutput() ApplicationPolicyResponseOutput
	ToApplicationPolicyResponseOutputWithContext(context.Context) ApplicationPolicyResponseOutput
}

type ApplicationPolicyResponseArgs struct {
	Name               pulumi.StringPtrInput `pulumi:"name"`
	Parameters         pulumi.StringPtrInput `pulumi:"parameters"`
	PolicyDefinitionId pulumi.StringPtrInput `pulumi:"policyDefinitionId"`
}

func (ApplicationPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPolicyResponse)(nil)).Elem()
}

func (i ApplicationPolicyResponseArgs) ToApplicationPolicyResponseOutput() ApplicationPolicyResponseOutput {
	return i.ToApplicationPolicyResponseOutputWithContext(context.Background())
}

func (i ApplicationPolicyResponseArgs) ToApplicationPolicyResponseOutputWithContext(ctx context.Context) ApplicationPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPolicyResponseOutput)
}





type ApplicationPolicyResponseArrayInput interface {
	pulumi.Input

	ToApplicationPolicyResponseArrayOutput() ApplicationPolicyResponseArrayOutput
	ToApplicationPolicyResponseArrayOutputWithContext(context.Context) ApplicationPolicyResponseArrayOutput
}

type ApplicationPolicyResponseArray []ApplicationPolicyResponseInput

func (ApplicationPolicyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationPolicyResponse)(nil)).Elem()
}

func (i ApplicationPolicyResponseArray) ToApplicationPolicyResponseArrayOutput() ApplicationPolicyResponseArrayOutput {
	return i.ToApplicationPolicyResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationPolicyResponseArray) ToApplicationPolicyResponseArrayOutputWithContext(ctx context.Context) ApplicationPolicyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPolicyResponseArrayOutput)
}

type ApplicationPolicyResponseOutput struct{ *pulumi.OutputState }

func (ApplicationPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPolicyResponse)(nil)).Elem()
}

func (o ApplicationPolicyResponseOutput) ToApplicationPolicyResponseOutput() ApplicationPolicyResponseOutput {
	return o
}

func (o ApplicationPolicyResponseOutput) ToApplicationPolicyResponseOutputWithContext(ctx context.Context) ApplicationPolicyResponseOutput {
	return o
}

func (o ApplicationPolicyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPolicyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationPolicyResponseOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPolicyResponse) *string { return v.Parameters }).(pulumi.StringPtrOutput)
}

func (o ApplicationPolicyResponseOutput) PolicyDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPolicyResponse) *string { return v.PolicyDefinitionId }).(pulumi.StringPtrOutput)
}

type ApplicationPolicyResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationPolicyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationPolicyResponse)(nil)).Elem()
}

func (o ApplicationPolicyResponseArrayOutput) ToApplicationPolicyResponseArrayOutput() ApplicationPolicyResponseArrayOutput {
	return o
}

func (o ApplicationPolicyResponseArrayOutput) ToApplicationPolicyResponseArrayOutputWithContext(ctx context.Context) ApplicationPolicyResponseArrayOutput {
	return o
}

func (o ApplicationPolicyResponseArrayOutput) Index(i pulumi.IntInput) ApplicationPolicyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationPolicyResponse {
		return vs[0].([]ApplicationPolicyResponse)[vs[1].(int)]
	}).(ApplicationPolicyResponseOutput)
}

type Identity struct {
	Type                   *ResourceIdentityType  `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type                   ResourceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput              `pulumi:"userAssignedIdentities"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v Identity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

func (o IdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v Identity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

func (o IdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *Identity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type IdentityResponse struct {
	PrincipalId            string                                          `pulumi:"principalId"`
	TenantId               string                                          `pulumi:"tenantId"`
	Type                   *string                                         `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedResourceIdentityResponse `pulumi:"userAssignedIdentities"`
}





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId            pulumi.StringInput                           `pulumi:"principalId"`
	TenantId               pulumi.StringInput                           `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput                        `pulumi:"type"`
	UserAssignedIdentities UserAssignedResourceIdentityResponseMapInput `pulumi:"userAssignedIdentities"`
}

func (IdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (i IdentityResponseArgs) ToIdentityResponseOutput() IdentityResponseOutput {
	return i.ToIdentityResponseOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput)
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput).ToIdentityResponsePtrOutputWithContext(ctx)
}









type IdentityResponsePtrInput interface {
	pulumi.Input

	ToIdentityResponsePtrOutput() IdentityResponsePtrOutput
	ToIdentityResponsePtrOutputWithContext(context.Context) IdentityResponsePtrOutput
}

type identityResponsePtrType IdentityResponseArgs

func IdentityResponsePtr(v *IdentityResponseArgs) IdentityResponsePtrInput {
	return (*identityResponsePtrType)(v)
}

func (*identityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponsePtrOutput)
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityResponse) *IdentityResponse {
		return &v
	}).(IdentityResponsePtrOutput)
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o IdentityResponseOutput) UserAssignedIdentities() UserAssignedResourceIdentityResponseMapOutput {
	return o.ApplyT(func(v IdentityResponse) map[string]UserAssignedResourceIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedResourceIdentityResponseMapOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedResourceIdentityResponseMapOutput {
	return o.ApplyT(func(v *IdentityResponse) map[string]UserAssignedResourceIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedResourceIdentityResponseMapOutput)
}

type JitApproverDefinition struct {
	DisplayName *string `pulumi:"displayName"`
	Id          string  `pulumi:"id"`
	Type        *string `pulumi:"type"`
}





type JitApproverDefinitionInput interface {
	pulumi.Input

	ToJitApproverDefinitionOutput() JitApproverDefinitionOutput
	ToJitApproverDefinitionOutputWithContext(context.Context) JitApproverDefinitionOutput
}

type JitApproverDefinitionArgs struct {
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	Id          pulumi.StringInput    `pulumi:"id"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (JitApproverDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitApproverDefinition)(nil)).Elem()
}

func (i JitApproverDefinitionArgs) ToJitApproverDefinitionOutput() JitApproverDefinitionOutput {
	return i.ToJitApproverDefinitionOutputWithContext(context.Background())
}

func (i JitApproverDefinitionArgs) ToJitApproverDefinitionOutputWithContext(ctx context.Context) JitApproverDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitApproverDefinitionOutput)
}





type JitApproverDefinitionArrayInput interface {
	pulumi.Input

	ToJitApproverDefinitionArrayOutput() JitApproverDefinitionArrayOutput
	ToJitApproverDefinitionArrayOutputWithContext(context.Context) JitApproverDefinitionArrayOutput
}

type JitApproverDefinitionArray []JitApproverDefinitionInput

func (JitApproverDefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitApproverDefinition)(nil)).Elem()
}

func (i JitApproverDefinitionArray) ToJitApproverDefinitionArrayOutput() JitApproverDefinitionArrayOutput {
	return i.ToJitApproverDefinitionArrayOutputWithContext(context.Background())
}

func (i JitApproverDefinitionArray) ToJitApproverDefinitionArrayOutputWithContext(ctx context.Context) JitApproverDefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitApproverDefinitionArrayOutput)
}

type JitApproverDefinitionOutput struct{ *pulumi.OutputState }

func (JitApproverDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitApproverDefinition)(nil)).Elem()
}

func (o JitApproverDefinitionOutput) ToJitApproverDefinitionOutput() JitApproverDefinitionOutput {
	return o
}

func (o JitApproverDefinitionOutput) ToJitApproverDefinitionOutputWithContext(ctx context.Context) JitApproverDefinitionOutput {
	return o
}

func (o JitApproverDefinitionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitApproverDefinition) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o JitApproverDefinitionOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v JitApproverDefinition) string { return v.Id }).(pulumi.StringOutput)
}

func (o JitApproverDefinitionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitApproverDefinition) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type JitApproverDefinitionArrayOutput struct{ *pulumi.OutputState }

func (JitApproverDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitApproverDefinition)(nil)).Elem()
}

func (o JitApproverDefinitionArrayOutput) ToJitApproverDefinitionArrayOutput() JitApproverDefinitionArrayOutput {
	return o
}

func (o JitApproverDefinitionArrayOutput) ToJitApproverDefinitionArrayOutputWithContext(ctx context.Context) JitApproverDefinitionArrayOutput {
	return o
}

func (o JitApproverDefinitionArrayOutput) Index(i pulumi.IntInput) JitApproverDefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitApproverDefinition {
		return vs[0].([]JitApproverDefinition)[vs[1].(int)]
	}).(JitApproverDefinitionOutput)
}

type JitApproverDefinitionResponse struct {
	DisplayName *string `pulumi:"displayName"`
	Id          string  `pulumi:"id"`
	Type        *string `pulumi:"type"`
}





type JitApproverDefinitionResponseInput interface {
	pulumi.Input

	ToJitApproverDefinitionResponseOutput() JitApproverDefinitionResponseOutput
	ToJitApproverDefinitionResponseOutputWithContext(context.Context) JitApproverDefinitionResponseOutput
}

type JitApproverDefinitionResponseArgs struct {
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	Id          pulumi.StringInput    `pulumi:"id"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (JitApproverDefinitionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitApproverDefinitionResponse)(nil)).Elem()
}

func (i JitApproverDefinitionResponseArgs) ToJitApproverDefinitionResponseOutput() JitApproverDefinitionResponseOutput {
	return i.ToJitApproverDefinitionResponseOutputWithContext(context.Background())
}

func (i JitApproverDefinitionResponseArgs) ToJitApproverDefinitionResponseOutputWithContext(ctx context.Context) JitApproverDefinitionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitApproverDefinitionResponseOutput)
}





type JitApproverDefinitionResponseArrayInput interface {
	pulumi.Input

	ToJitApproverDefinitionResponseArrayOutput() JitApproverDefinitionResponseArrayOutput
	ToJitApproverDefinitionResponseArrayOutputWithContext(context.Context) JitApproverDefinitionResponseArrayOutput
}

type JitApproverDefinitionResponseArray []JitApproverDefinitionResponseInput

func (JitApproverDefinitionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitApproverDefinitionResponse)(nil)).Elem()
}

func (i JitApproverDefinitionResponseArray) ToJitApproverDefinitionResponseArrayOutput() JitApproverDefinitionResponseArrayOutput {
	return i.ToJitApproverDefinitionResponseArrayOutputWithContext(context.Background())
}

func (i JitApproverDefinitionResponseArray) ToJitApproverDefinitionResponseArrayOutputWithContext(ctx context.Context) JitApproverDefinitionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitApproverDefinitionResponseArrayOutput)
}

type JitApproverDefinitionResponseOutput struct{ *pulumi.OutputState }

func (JitApproverDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitApproverDefinitionResponse)(nil)).Elem()
}

func (o JitApproverDefinitionResponseOutput) ToJitApproverDefinitionResponseOutput() JitApproverDefinitionResponseOutput {
	return o
}

func (o JitApproverDefinitionResponseOutput) ToJitApproverDefinitionResponseOutputWithContext(ctx context.Context) JitApproverDefinitionResponseOutput {
	return o
}

func (o JitApproverDefinitionResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitApproverDefinitionResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o JitApproverDefinitionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v JitApproverDefinitionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o JitApproverDefinitionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitApproverDefinitionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type JitApproverDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (JitApproverDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitApproverDefinitionResponse)(nil)).Elem()
}

func (o JitApproverDefinitionResponseArrayOutput) ToJitApproverDefinitionResponseArrayOutput() JitApproverDefinitionResponseArrayOutput {
	return o
}

func (o JitApproverDefinitionResponseArrayOutput) ToJitApproverDefinitionResponseArrayOutputWithContext(ctx context.Context) JitApproverDefinitionResponseArrayOutput {
	return o
}

func (o JitApproverDefinitionResponseArrayOutput) Index(i pulumi.IntInput) JitApproverDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitApproverDefinitionResponse {
		return vs[0].([]JitApproverDefinitionResponse)[vs[1].(int)]
	}).(JitApproverDefinitionResponseOutput)
}

type JitAuthorizationPolicies struct {
	PrincipalId      string `pulumi:"principalId"`
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
}





type JitAuthorizationPoliciesInput interface {
	pulumi.Input

	ToJitAuthorizationPoliciesOutput() JitAuthorizationPoliciesOutput
	ToJitAuthorizationPoliciesOutputWithContext(context.Context) JitAuthorizationPoliciesOutput
}

type JitAuthorizationPoliciesArgs struct {
	PrincipalId      pulumi.StringInput `pulumi:"principalId"`
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
}

func (JitAuthorizationPoliciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitAuthorizationPolicies)(nil)).Elem()
}

func (i JitAuthorizationPoliciesArgs) ToJitAuthorizationPoliciesOutput() JitAuthorizationPoliciesOutput {
	return i.ToJitAuthorizationPoliciesOutputWithContext(context.Background())
}

func (i JitAuthorizationPoliciesArgs) ToJitAuthorizationPoliciesOutputWithContext(ctx context.Context) JitAuthorizationPoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitAuthorizationPoliciesOutput)
}





type JitAuthorizationPoliciesArrayInput interface {
	pulumi.Input

	ToJitAuthorizationPoliciesArrayOutput() JitAuthorizationPoliciesArrayOutput
	ToJitAuthorizationPoliciesArrayOutputWithContext(context.Context) JitAuthorizationPoliciesArrayOutput
}

type JitAuthorizationPoliciesArray []JitAuthorizationPoliciesInput

func (JitAuthorizationPoliciesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitAuthorizationPolicies)(nil)).Elem()
}

func (i JitAuthorizationPoliciesArray) ToJitAuthorizationPoliciesArrayOutput() JitAuthorizationPoliciesArrayOutput {
	return i.ToJitAuthorizationPoliciesArrayOutputWithContext(context.Background())
}

func (i JitAuthorizationPoliciesArray) ToJitAuthorizationPoliciesArrayOutputWithContext(ctx context.Context) JitAuthorizationPoliciesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitAuthorizationPoliciesArrayOutput)
}

type JitAuthorizationPoliciesOutput struct{ *pulumi.OutputState }

func (JitAuthorizationPoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitAuthorizationPolicies)(nil)).Elem()
}

func (o JitAuthorizationPoliciesOutput) ToJitAuthorizationPoliciesOutput() JitAuthorizationPoliciesOutput {
	return o
}

func (o JitAuthorizationPoliciesOutput) ToJitAuthorizationPoliciesOutputWithContext(ctx context.Context) JitAuthorizationPoliciesOutput {
	return o
}

func (o JitAuthorizationPoliciesOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v JitAuthorizationPolicies) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o JitAuthorizationPoliciesOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v JitAuthorizationPolicies) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type JitAuthorizationPoliciesArrayOutput struct{ *pulumi.OutputState }

func (JitAuthorizationPoliciesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitAuthorizationPolicies)(nil)).Elem()
}

func (o JitAuthorizationPoliciesArrayOutput) ToJitAuthorizationPoliciesArrayOutput() JitAuthorizationPoliciesArrayOutput {
	return o
}

func (o JitAuthorizationPoliciesArrayOutput) ToJitAuthorizationPoliciesArrayOutputWithContext(ctx context.Context) JitAuthorizationPoliciesArrayOutput {
	return o
}

func (o JitAuthorizationPoliciesArrayOutput) Index(i pulumi.IntInput) JitAuthorizationPoliciesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitAuthorizationPolicies {
		return vs[0].([]JitAuthorizationPolicies)[vs[1].(int)]
	}).(JitAuthorizationPoliciesOutput)
}

type JitAuthorizationPoliciesResponse struct {
	PrincipalId      string `pulumi:"principalId"`
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
}





type JitAuthorizationPoliciesResponseInput interface {
	pulumi.Input

	ToJitAuthorizationPoliciesResponseOutput() JitAuthorizationPoliciesResponseOutput
	ToJitAuthorizationPoliciesResponseOutputWithContext(context.Context) JitAuthorizationPoliciesResponseOutput
}

type JitAuthorizationPoliciesResponseArgs struct {
	PrincipalId      pulumi.StringInput `pulumi:"principalId"`
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
}

func (JitAuthorizationPoliciesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitAuthorizationPoliciesResponse)(nil)).Elem()
}

func (i JitAuthorizationPoliciesResponseArgs) ToJitAuthorizationPoliciesResponseOutput() JitAuthorizationPoliciesResponseOutput {
	return i.ToJitAuthorizationPoliciesResponseOutputWithContext(context.Background())
}

func (i JitAuthorizationPoliciesResponseArgs) ToJitAuthorizationPoliciesResponseOutputWithContext(ctx context.Context) JitAuthorizationPoliciesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitAuthorizationPoliciesResponseOutput)
}





type JitAuthorizationPoliciesResponseArrayInput interface {
	pulumi.Input

	ToJitAuthorizationPoliciesResponseArrayOutput() JitAuthorizationPoliciesResponseArrayOutput
	ToJitAuthorizationPoliciesResponseArrayOutputWithContext(context.Context) JitAuthorizationPoliciesResponseArrayOutput
}

type JitAuthorizationPoliciesResponseArray []JitAuthorizationPoliciesResponseInput

func (JitAuthorizationPoliciesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitAuthorizationPoliciesResponse)(nil)).Elem()
}

func (i JitAuthorizationPoliciesResponseArray) ToJitAuthorizationPoliciesResponseArrayOutput() JitAuthorizationPoliciesResponseArrayOutput {
	return i.ToJitAuthorizationPoliciesResponseArrayOutputWithContext(context.Background())
}

func (i JitAuthorizationPoliciesResponseArray) ToJitAuthorizationPoliciesResponseArrayOutputWithContext(ctx context.Context) JitAuthorizationPoliciesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitAuthorizationPoliciesResponseArrayOutput)
}

type JitAuthorizationPoliciesResponseOutput struct{ *pulumi.OutputState }

func (JitAuthorizationPoliciesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitAuthorizationPoliciesResponse)(nil)).Elem()
}

func (o JitAuthorizationPoliciesResponseOutput) ToJitAuthorizationPoliciesResponseOutput() JitAuthorizationPoliciesResponseOutput {
	return o
}

func (o JitAuthorizationPoliciesResponseOutput) ToJitAuthorizationPoliciesResponseOutputWithContext(ctx context.Context) JitAuthorizationPoliciesResponseOutput {
	return o
}

func (o JitAuthorizationPoliciesResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v JitAuthorizationPoliciesResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o JitAuthorizationPoliciesResponseOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v JitAuthorizationPoliciesResponse) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type JitAuthorizationPoliciesResponseArrayOutput struct{ *pulumi.OutputState }

func (JitAuthorizationPoliciesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitAuthorizationPoliciesResponse)(nil)).Elem()
}

func (o JitAuthorizationPoliciesResponseArrayOutput) ToJitAuthorizationPoliciesResponseArrayOutput() JitAuthorizationPoliciesResponseArrayOutput {
	return o
}

func (o JitAuthorizationPoliciesResponseArrayOutput) ToJitAuthorizationPoliciesResponseArrayOutputWithContext(ctx context.Context) JitAuthorizationPoliciesResponseArrayOutput {
	return o
}

func (o JitAuthorizationPoliciesResponseArrayOutput) Index(i pulumi.IntInput) JitAuthorizationPoliciesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitAuthorizationPoliciesResponse {
		return vs[0].([]JitAuthorizationPoliciesResponse)[vs[1].(int)]
	}).(JitAuthorizationPoliciesResponseOutput)
}

type JitSchedulingPolicy struct {
	Duration  string `pulumi:"duration"`
	StartTime string `pulumi:"startTime"`
}





type JitSchedulingPolicyInput interface {
	pulumi.Input

	ToJitSchedulingPolicyOutput() JitSchedulingPolicyOutput
	ToJitSchedulingPolicyOutputWithContext(context.Context) JitSchedulingPolicyOutput
}

type JitSchedulingPolicyArgs struct {
	Duration  pulumi.StringInput `pulumi:"duration"`
	StartTime pulumi.StringInput `pulumi:"startTime"`
}

func (JitSchedulingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitSchedulingPolicy)(nil)).Elem()
}

func (i JitSchedulingPolicyArgs) ToJitSchedulingPolicyOutput() JitSchedulingPolicyOutput {
	return i.ToJitSchedulingPolicyOutputWithContext(context.Background())
}

func (i JitSchedulingPolicyArgs) ToJitSchedulingPolicyOutputWithContext(ctx context.Context) JitSchedulingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitSchedulingPolicyOutput)
}

func (i JitSchedulingPolicyArgs) ToJitSchedulingPolicyPtrOutput() JitSchedulingPolicyPtrOutput {
	return i.ToJitSchedulingPolicyPtrOutputWithContext(context.Background())
}

func (i JitSchedulingPolicyArgs) ToJitSchedulingPolicyPtrOutputWithContext(ctx context.Context) JitSchedulingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitSchedulingPolicyOutput).ToJitSchedulingPolicyPtrOutputWithContext(ctx)
}









type JitSchedulingPolicyPtrInput interface {
	pulumi.Input

	ToJitSchedulingPolicyPtrOutput() JitSchedulingPolicyPtrOutput
	ToJitSchedulingPolicyPtrOutputWithContext(context.Context) JitSchedulingPolicyPtrOutput
}

type jitSchedulingPolicyPtrType JitSchedulingPolicyArgs

func JitSchedulingPolicyPtr(v *JitSchedulingPolicyArgs) JitSchedulingPolicyPtrInput {
	return (*jitSchedulingPolicyPtrType)(v)
}

func (*jitSchedulingPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JitSchedulingPolicy)(nil)).Elem()
}

func (i *jitSchedulingPolicyPtrType) ToJitSchedulingPolicyPtrOutput() JitSchedulingPolicyPtrOutput {
	return i.ToJitSchedulingPolicyPtrOutputWithContext(context.Background())
}

func (i *jitSchedulingPolicyPtrType) ToJitSchedulingPolicyPtrOutputWithContext(ctx context.Context) JitSchedulingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitSchedulingPolicyPtrOutput)
}

type JitSchedulingPolicyOutput struct{ *pulumi.OutputState }

func (JitSchedulingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitSchedulingPolicy)(nil)).Elem()
}

func (o JitSchedulingPolicyOutput) ToJitSchedulingPolicyOutput() JitSchedulingPolicyOutput {
	return o
}

func (o JitSchedulingPolicyOutput) ToJitSchedulingPolicyOutputWithContext(ctx context.Context) JitSchedulingPolicyOutput {
	return o
}

func (o JitSchedulingPolicyOutput) ToJitSchedulingPolicyPtrOutput() JitSchedulingPolicyPtrOutput {
	return o.ToJitSchedulingPolicyPtrOutputWithContext(context.Background())
}

func (o JitSchedulingPolicyOutput) ToJitSchedulingPolicyPtrOutputWithContext(ctx context.Context) JitSchedulingPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JitSchedulingPolicy) *JitSchedulingPolicy {
		return &v
	}).(JitSchedulingPolicyPtrOutput)
}

func (o JitSchedulingPolicyOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v JitSchedulingPolicy) string { return v.Duration }).(pulumi.StringOutput)
}

func (o JitSchedulingPolicyOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v JitSchedulingPolicy) string { return v.StartTime }).(pulumi.StringOutput)
}

type JitSchedulingPolicyPtrOutput struct{ *pulumi.OutputState }

func (JitSchedulingPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JitSchedulingPolicy)(nil)).Elem()
}

func (o JitSchedulingPolicyPtrOutput) ToJitSchedulingPolicyPtrOutput() JitSchedulingPolicyPtrOutput {
	return o
}

func (o JitSchedulingPolicyPtrOutput) ToJitSchedulingPolicyPtrOutputWithContext(ctx context.Context) JitSchedulingPolicyPtrOutput {
	return o
}

func (o JitSchedulingPolicyPtrOutput) Elem() JitSchedulingPolicyOutput {
	return o.ApplyT(func(v *JitSchedulingPolicy) JitSchedulingPolicy {
		if v != nil {
			return *v
		}
		var ret JitSchedulingPolicy
		return ret
	}).(JitSchedulingPolicyOutput)
}

func (o JitSchedulingPolicyPtrOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JitSchedulingPolicy) *string {
		if v == nil {
			return nil
		}
		return &v.Duration
	}).(pulumi.StringPtrOutput)
}

func (o JitSchedulingPolicyPtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JitSchedulingPolicy) *string {
		if v == nil {
			return nil
		}
		return &v.StartTime
	}).(pulumi.StringPtrOutput)
}

type JitSchedulingPolicyResponse struct {
	Duration  string `pulumi:"duration"`
	StartTime string `pulumi:"startTime"`
	Type      string `pulumi:"type"`
}





type JitSchedulingPolicyResponseInput interface {
	pulumi.Input

	ToJitSchedulingPolicyResponseOutput() JitSchedulingPolicyResponseOutput
	ToJitSchedulingPolicyResponseOutputWithContext(context.Context) JitSchedulingPolicyResponseOutput
}

type JitSchedulingPolicyResponseArgs struct {
	Duration  pulumi.StringInput `pulumi:"duration"`
	StartTime pulumi.StringInput `pulumi:"startTime"`
	Type      pulumi.StringInput `pulumi:"type"`
}

func (JitSchedulingPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitSchedulingPolicyResponse)(nil)).Elem()
}

func (i JitSchedulingPolicyResponseArgs) ToJitSchedulingPolicyResponseOutput() JitSchedulingPolicyResponseOutput {
	return i.ToJitSchedulingPolicyResponseOutputWithContext(context.Background())
}

func (i JitSchedulingPolicyResponseArgs) ToJitSchedulingPolicyResponseOutputWithContext(ctx context.Context) JitSchedulingPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitSchedulingPolicyResponseOutput)
}

func (i JitSchedulingPolicyResponseArgs) ToJitSchedulingPolicyResponsePtrOutput() JitSchedulingPolicyResponsePtrOutput {
	return i.ToJitSchedulingPolicyResponsePtrOutputWithContext(context.Background())
}

func (i JitSchedulingPolicyResponseArgs) ToJitSchedulingPolicyResponsePtrOutputWithContext(ctx context.Context) JitSchedulingPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitSchedulingPolicyResponseOutput).ToJitSchedulingPolicyResponsePtrOutputWithContext(ctx)
}









type JitSchedulingPolicyResponsePtrInput interface {
	pulumi.Input

	ToJitSchedulingPolicyResponsePtrOutput() JitSchedulingPolicyResponsePtrOutput
	ToJitSchedulingPolicyResponsePtrOutputWithContext(context.Context) JitSchedulingPolicyResponsePtrOutput
}

type jitSchedulingPolicyResponsePtrType JitSchedulingPolicyResponseArgs

func JitSchedulingPolicyResponsePtr(v *JitSchedulingPolicyResponseArgs) JitSchedulingPolicyResponsePtrInput {
	return (*jitSchedulingPolicyResponsePtrType)(v)
}

func (*jitSchedulingPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JitSchedulingPolicyResponse)(nil)).Elem()
}

func (i *jitSchedulingPolicyResponsePtrType) ToJitSchedulingPolicyResponsePtrOutput() JitSchedulingPolicyResponsePtrOutput {
	return i.ToJitSchedulingPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *jitSchedulingPolicyResponsePtrType) ToJitSchedulingPolicyResponsePtrOutputWithContext(ctx context.Context) JitSchedulingPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitSchedulingPolicyResponsePtrOutput)
}

type JitSchedulingPolicyResponseOutput struct{ *pulumi.OutputState }

func (JitSchedulingPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitSchedulingPolicyResponse)(nil)).Elem()
}

func (o JitSchedulingPolicyResponseOutput) ToJitSchedulingPolicyResponseOutput() JitSchedulingPolicyResponseOutput {
	return o
}

func (o JitSchedulingPolicyResponseOutput) ToJitSchedulingPolicyResponseOutputWithContext(ctx context.Context) JitSchedulingPolicyResponseOutput {
	return o
}

func (o JitSchedulingPolicyResponseOutput) ToJitSchedulingPolicyResponsePtrOutput() JitSchedulingPolicyResponsePtrOutput {
	return o.ToJitSchedulingPolicyResponsePtrOutputWithContext(context.Background())
}

func (o JitSchedulingPolicyResponseOutput) ToJitSchedulingPolicyResponsePtrOutputWithContext(ctx context.Context) JitSchedulingPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JitSchedulingPolicyResponse) *JitSchedulingPolicyResponse {
		return &v
	}).(JitSchedulingPolicyResponsePtrOutput)
}

func (o JitSchedulingPolicyResponseOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v JitSchedulingPolicyResponse) string { return v.Duration }).(pulumi.StringOutput)
}

func (o JitSchedulingPolicyResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v JitSchedulingPolicyResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o JitSchedulingPolicyResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v JitSchedulingPolicyResponse) string { return v.Type }).(pulumi.StringOutput)
}

type JitSchedulingPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (JitSchedulingPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JitSchedulingPolicyResponse)(nil)).Elem()
}

func (o JitSchedulingPolicyResponsePtrOutput) ToJitSchedulingPolicyResponsePtrOutput() JitSchedulingPolicyResponsePtrOutput {
	return o
}

func (o JitSchedulingPolicyResponsePtrOutput) ToJitSchedulingPolicyResponsePtrOutputWithContext(ctx context.Context) JitSchedulingPolicyResponsePtrOutput {
	return o
}

func (o JitSchedulingPolicyResponsePtrOutput) Elem() JitSchedulingPolicyResponseOutput {
	return o.ApplyT(func(v *JitSchedulingPolicyResponse) JitSchedulingPolicyResponse {
		if v != nil {
			return *v
		}
		var ret JitSchedulingPolicyResponse
		return ret
	}).(JitSchedulingPolicyResponseOutput)
}

func (o JitSchedulingPolicyResponsePtrOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JitSchedulingPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Duration
	}).(pulumi.StringPtrOutput)
}

func (o JitSchedulingPolicyResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JitSchedulingPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o JitSchedulingPolicyResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JitSchedulingPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type Plan struct {
	Name          string  `pulumi:"name"`
	Product       string  `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     string  `pulumi:"publisher"`
	Version       string  `pulumi:"version"`
}





type PlanInput interface {
	pulumi.Input

	ToPlanOutput() PlanOutput
	ToPlanOutputWithContext(context.Context) PlanOutput
}

type PlanArgs struct {
	Name          pulumi.StringInput    `pulumi:"name"`
	Product       pulumi.StringInput    `pulumi:"product"`
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	Publisher     pulumi.StringInput    `pulumi:"publisher"`
	Version       pulumi.StringInput    `pulumi:"version"`
}

func (PlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (i PlanArgs) ToPlanOutput() PlanOutput {
	return i.ToPlanOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput)
}

func (i PlanArgs) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput).ToPlanPtrOutputWithContext(ctx)
}









type PlanPtrInput interface {
	pulumi.Input

	ToPlanPtrOutput() PlanPtrOutput
	ToPlanPtrOutputWithContext(context.Context) PlanPtrOutput
}

type planPtrType PlanArgs

func PlanPtr(v *PlanArgs) PlanPtrInput {
	return (*planPtrType)(v)
}

func (*planPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (i *planPtrType) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i *planPtrType) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanPtrOutput)
}

type PlanOutput struct{ *pulumi.OutputState }

func (PlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (o PlanOutput) ToPlanOutput() PlanOutput {
	return o
}

func (o PlanOutput) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return o
}

func (o PlanOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o.ToPlanPtrOutputWithContext(context.Background())
}

func (o PlanOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Plan) *Plan {
		return &v
	}).(PlanPtrOutput)
}

func (o PlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Plan) string { return v.Name }).(pulumi.StringOutput)
}

func (o PlanOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v Plan) string { return v.Product }).(pulumi.StringOutput)
}

func (o PlanOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PlanOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v Plan) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o PlanOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v Plan) string { return v.Version }).(pulumi.StringOutput)
}

type PlanPtrOutput struct{ *pulumi.OutputState }

func (PlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (o PlanPtrOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) Elem() PlanOutput {
	return o.ApplyT(func(v *Plan) Plan {
		if v != nil {
			return *v
		}
		var ret Plan
		return ret
	}).(PlanOutput)
}

func (o PlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return &v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

type PlanResponse struct {
	Name          string  `pulumi:"name"`
	Product       string  `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     string  `pulumi:"publisher"`
	Version       string  `pulumi:"version"`
}





type PlanResponseInput interface {
	pulumi.Input

	ToPlanResponseOutput() PlanResponseOutput
	ToPlanResponseOutputWithContext(context.Context) PlanResponseOutput
}

type PlanResponseArgs struct {
	Name          pulumi.StringInput    `pulumi:"name"`
	Product       pulumi.StringInput    `pulumi:"product"`
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	Publisher     pulumi.StringInput    `pulumi:"publisher"`
	Version       pulumi.StringInput    `pulumi:"version"`
}

func (PlanResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanResponse)(nil)).Elem()
}

func (i PlanResponseArgs) ToPlanResponseOutput() PlanResponseOutput {
	return i.ToPlanResponseOutputWithContext(context.Background())
}

func (i PlanResponseArgs) ToPlanResponseOutputWithContext(ctx context.Context) PlanResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanResponseOutput)
}

func (i PlanResponseArgs) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return i.ToPlanResponsePtrOutputWithContext(context.Background())
}

func (i PlanResponseArgs) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanResponseOutput).ToPlanResponsePtrOutputWithContext(ctx)
}









type PlanResponsePtrInput interface {
	pulumi.Input

	ToPlanResponsePtrOutput() PlanResponsePtrOutput
	ToPlanResponsePtrOutputWithContext(context.Context) PlanResponsePtrOutput
}

type planResponsePtrType PlanResponseArgs

func PlanResponsePtr(v *PlanResponseArgs) PlanResponsePtrInput {
	return (*planResponsePtrType)(v)
}

func (*planResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanResponse)(nil)).Elem()
}

func (i *planResponsePtrType) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return i.ToPlanResponsePtrOutputWithContext(context.Background())
}

func (i *planResponsePtrType) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanResponsePtrOutput)
}

type PlanResponseOutput struct{ *pulumi.OutputState }

func (PlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanResponse)(nil)).Elem()
}

func (o PlanResponseOutput) ToPlanResponseOutput() PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) ToPlanResponseOutputWithContext(ctx context.Context) PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return o.ToPlanResponsePtrOutputWithContext(context.Background())
}

func (o PlanResponseOutput) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PlanResponse) *PlanResponse {
		return &v
	}).(PlanResponsePtrOutput)
}

func (o PlanResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Product }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PlanResponseOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Version }).(pulumi.StringOutput)
}

type PlanResponsePtrOutput struct{ *pulumi.OutputState }

func (PlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanResponse)(nil)).Elem()
}

func (o PlanResponsePtrOutput) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return o
}

func (o PlanResponsePtrOutput) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return o
}

func (o PlanResponsePtrOutput) Elem() PlanResponseOutput {
	return o.ApplyT(func(v *PlanResponse) PlanResponse {
		if v != nil {
			return *v
		}
		var ret PlanResponse
		return ret
	}).(PlanResponseOutput)
}

func (o PlanResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Model    *string `pulumi:"model"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Model    pulumi.StringPtrInput `pulumi:"model"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Model }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Model
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Model    *string `pulumi:"model"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Model    pulumi.StringPtrInput `pulumi:"model"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Model }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Model
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type UserAssignedResourceIdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
}





type UserAssignedResourceIdentityResponseInput interface {
	pulumi.Input

	ToUserAssignedResourceIdentityResponseOutput() UserAssignedResourceIdentityResponseOutput
	ToUserAssignedResourceIdentityResponseOutputWithContext(context.Context) UserAssignedResourceIdentityResponseOutput
}

type UserAssignedResourceIdentityResponseArgs struct {
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
}

func (UserAssignedResourceIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedResourceIdentityResponse)(nil)).Elem()
}

func (i UserAssignedResourceIdentityResponseArgs) ToUserAssignedResourceIdentityResponseOutput() UserAssignedResourceIdentityResponseOutput {
	return i.ToUserAssignedResourceIdentityResponseOutputWithContext(context.Background())
}

func (i UserAssignedResourceIdentityResponseArgs) ToUserAssignedResourceIdentityResponseOutputWithContext(ctx context.Context) UserAssignedResourceIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedResourceIdentityResponseOutput)
}





type UserAssignedResourceIdentityResponseMapInput interface {
	pulumi.Input

	ToUserAssignedResourceIdentityResponseMapOutput() UserAssignedResourceIdentityResponseMapOutput
	ToUserAssignedResourceIdentityResponseMapOutputWithContext(context.Context) UserAssignedResourceIdentityResponseMapOutput
}

type UserAssignedResourceIdentityResponseMap map[string]UserAssignedResourceIdentityResponseInput

func (UserAssignedResourceIdentityResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedResourceIdentityResponse)(nil)).Elem()
}

func (i UserAssignedResourceIdentityResponseMap) ToUserAssignedResourceIdentityResponseMapOutput() UserAssignedResourceIdentityResponseMapOutput {
	return i.ToUserAssignedResourceIdentityResponseMapOutputWithContext(context.Background())
}

func (i UserAssignedResourceIdentityResponseMap) ToUserAssignedResourceIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedResourceIdentityResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedResourceIdentityResponseMapOutput)
}

type UserAssignedResourceIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedResourceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedResourceIdentityResponse)(nil)).Elem()
}

func (o UserAssignedResourceIdentityResponseOutput) ToUserAssignedResourceIdentityResponseOutput() UserAssignedResourceIdentityResponseOutput {
	return o
}

func (o UserAssignedResourceIdentityResponseOutput) ToUserAssignedResourceIdentityResponseOutputWithContext(ctx context.Context) UserAssignedResourceIdentityResponseOutput {
	return o
}

func (o UserAssignedResourceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedResourceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o UserAssignedResourceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedResourceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

type UserAssignedResourceIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedResourceIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedResourceIdentityResponse)(nil)).Elem()
}

func (o UserAssignedResourceIdentityResponseMapOutput) ToUserAssignedResourceIdentityResponseMapOutput() UserAssignedResourceIdentityResponseMapOutput {
	return o
}

func (o UserAssignedResourceIdentityResponseMapOutput) ToUserAssignedResourceIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedResourceIdentityResponseMapOutput {
	return o
}

func (o UserAssignedResourceIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedResourceIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedResourceIdentityResponse {
		return vs[0].(map[string]UserAssignedResourceIdentityResponse)[vs[1].(string)]
	}).(UserAssignedResourceIdentityResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationArtifactResponseOutput{})
	pulumi.RegisterOutputType(ApplicationArtifactResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationAuthorizationOutput{})
	pulumi.RegisterOutputType(ApplicationAuthorizationArrayOutput{})
	pulumi.RegisterOutputType(ApplicationAuthorizationResponseOutput{})
	pulumi.RegisterOutputType(ApplicationAuthorizationResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationBillingDetailsDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ApplicationBillingDetailsDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationClientDetailsResponseOutput{})
	pulumi.RegisterOutputType(ApplicationClientDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationDefinitionArtifactOutput{})
	pulumi.RegisterOutputType(ApplicationDefinitionArtifactArrayOutput{})
	pulumi.RegisterOutputType(ApplicationDefinitionArtifactResponseOutput{})
	pulumi.RegisterOutputType(ApplicationDefinitionArtifactResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationDeploymentPolicyOutput{})
	pulumi.RegisterOutputType(ApplicationDeploymentPolicyPtrOutput{})
	pulumi.RegisterOutputType(ApplicationDeploymentPolicyResponseOutput{})
	pulumi.RegisterOutputType(ApplicationDeploymentPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationJitAccessPolicyOutput{})
	pulumi.RegisterOutputType(ApplicationJitAccessPolicyPtrOutput{})
	pulumi.RegisterOutputType(ApplicationJitAccessPolicyResponseOutput{})
	pulumi.RegisterOutputType(ApplicationJitAccessPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationManagementPolicyOutput{})
	pulumi.RegisterOutputType(ApplicationManagementPolicyPtrOutput{})
	pulumi.RegisterOutputType(ApplicationManagementPolicyResponseOutput{})
	pulumi.RegisterOutputType(ApplicationManagementPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationNotificationEndpointOutput{})
	pulumi.RegisterOutputType(ApplicationNotificationEndpointArrayOutput{})
	pulumi.RegisterOutputType(ApplicationNotificationEndpointResponseOutput{})
	pulumi.RegisterOutputType(ApplicationNotificationEndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationNotificationPolicyOutput{})
	pulumi.RegisterOutputType(ApplicationNotificationPolicyPtrOutput{})
	pulumi.RegisterOutputType(ApplicationNotificationPolicyResponseOutput{})
	pulumi.RegisterOutputType(ApplicationNotificationPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationPackageContactResponseOutput{})
	pulumi.RegisterOutputType(ApplicationPackageContactResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationPackageLockingPolicyDefinitionOutput{})
	pulumi.RegisterOutputType(ApplicationPackageLockingPolicyDefinitionPtrOutput{})
	pulumi.RegisterOutputType(ApplicationPackageLockingPolicyDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ApplicationPackageLockingPolicyDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationPackageSupportUrlsResponseOutput{})
	pulumi.RegisterOutputType(ApplicationPackageSupportUrlsResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationPolicyOutput{})
	pulumi.RegisterOutputType(ApplicationPolicyArrayOutput{})
	pulumi.RegisterOutputType(ApplicationPolicyResponseOutput{})
	pulumi.RegisterOutputType(ApplicationPolicyResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(JitApproverDefinitionOutput{})
	pulumi.RegisterOutputType(JitApproverDefinitionArrayOutput{})
	pulumi.RegisterOutputType(JitApproverDefinitionResponseOutput{})
	pulumi.RegisterOutputType(JitApproverDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(JitAuthorizationPoliciesOutput{})
	pulumi.RegisterOutputType(JitAuthorizationPoliciesArrayOutput{})
	pulumi.RegisterOutputType(JitAuthorizationPoliciesResponseOutput{})
	pulumi.RegisterOutputType(JitAuthorizationPoliciesResponseArrayOutput{})
	pulumi.RegisterOutputType(JitSchedulingPolicyOutput{})
	pulumi.RegisterOutputType(JitSchedulingPolicyPtrOutput{})
	pulumi.RegisterOutputType(JitSchedulingPolicyResponseOutput{})
	pulumi.RegisterOutputType(JitSchedulingPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(PlanOutput{})
	pulumi.RegisterOutputType(PlanPtrOutput{})
	pulumi.RegisterOutputType(PlanResponseOutput{})
	pulumi.RegisterOutputType(PlanResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(UserAssignedResourceIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedResourceIdentityResponseMapOutput{})
}
