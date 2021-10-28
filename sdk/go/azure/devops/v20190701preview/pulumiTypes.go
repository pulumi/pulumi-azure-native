


package v20190701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Authorization struct {
	AuthorizationType string            `pulumi:"authorizationType"`
	Parameters        map[string]string `pulumi:"parameters"`
}





type AuthorizationInput interface {
	pulumi.Input

	ToAuthorizationOutput() AuthorizationOutput
	ToAuthorizationOutputWithContext(context.Context) AuthorizationOutput
}

type AuthorizationArgs struct {
	AuthorizationType pulumi.StringInput    `pulumi:"authorizationType"`
	Parameters        pulumi.StringMapInput `pulumi:"parameters"`
}

func (AuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Authorization)(nil)).Elem()
}

func (i AuthorizationArgs) ToAuthorizationOutput() AuthorizationOutput {
	return i.ToAuthorizationOutputWithContext(context.Background())
}

func (i AuthorizationArgs) ToAuthorizationOutputWithContext(ctx context.Context) AuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationOutput)
}

func (i AuthorizationArgs) ToAuthorizationPtrOutput() AuthorizationPtrOutput {
	return i.ToAuthorizationPtrOutputWithContext(context.Background())
}

func (i AuthorizationArgs) ToAuthorizationPtrOutputWithContext(ctx context.Context) AuthorizationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationOutput).ToAuthorizationPtrOutputWithContext(ctx)
}









type AuthorizationPtrInput interface {
	pulumi.Input

	ToAuthorizationPtrOutput() AuthorizationPtrOutput
	ToAuthorizationPtrOutputWithContext(context.Context) AuthorizationPtrOutput
}

type authorizationPtrType AuthorizationArgs

func AuthorizationPtr(v *AuthorizationArgs) AuthorizationPtrInput {
	return (*authorizationPtrType)(v)
}

func (*authorizationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Authorization)(nil)).Elem()
}

func (i *authorizationPtrType) ToAuthorizationPtrOutput() AuthorizationPtrOutput {
	return i.ToAuthorizationPtrOutputWithContext(context.Background())
}

func (i *authorizationPtrType) ToAuthorizationPtrOutputWithContext(ctx context.Context) AuthorizationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationPtrOutput)
}

type AuthorizationOutput struct{ *pulumi.OutputState }

func (AuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Authorization)(nil)).Elem()
}

func (o AuthorizationOutput) ToAuthorizationOutput() AuthorizationOutput {
	return o
}

func (o AuthorizationOutput) ToAuthorizationOutputWithContext(ctx context.Context) AuthorizationOutput {
	return o
}

func (o AuthorizationOutput) ToAuthorizationPtrOutput() AuthorizationPtrOutput {
	return o.ToAuthorizationPtrOutputWithContext(context.Background())
}

func (o AuthorizationOutput) ToAuthorizationPtrOutputWithContext(ctx context.Context) AuthorizationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Authorization) *Authorization {
		return &v
	}).(AuthorizationPtrOutput)
}

func (o AuthorizationOutput) AuthorizationType() pulumi.StringOutput {
	return o.ApplyT(func(v Authorization) string { return v.AuthorizationType }).(pulumi.StringOutput)
}

func (o AuthorizationOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v Authorization) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

type AuthorizationPtrOutput struct{ *pulumi.OutputState }

func (AuthorizationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Authorization)(nil)).Elem()
}

func (o AuthorizationPtrOutput) ToAuthorizationPtrOutput() AuthorizationPtrOutput {
	return o
}

func (o AuthorizationPtrOutput) ToAuthorizationPtrOutputWithContext(ctx context.Context) AuthorizationPtrOutput {
	return o
}

func (o AuthorizationPtrOutput) Elem() AuthorizationOutput {
	return o.ApplyT(func(v *Authorization) Authorization {
		if v != nil {
			return *v
		}
		var ret Authorization
		return ret
	}).(AuthorizationOutput)
}

func (o AuthorizationPtrOutput) AuthorizationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Authorization) *string {
		if v == nil {
			return nil
		}
		return &v.AuthorizationType
	}).(pulumi.StringPtrOutput)
}

func (o AuthorizationPtrOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Authorization) map[string]string {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.StringMapOutput)
}

type AuthorizationResponse struct {
	AuthorizationType string            `pulumi:"authorizationType"`
	Parameters        map[string]string `pulumi:"parameters"`
}





type AuthorizationResponseInput interface {
	pulumi.Input

	ToAuthorizationResponseOutput() AuthorizationResponseOutput
	ToAuthorizationResponseOutputWithContext(context.Context) AuthorizationResponseOutput
}

type AuthorizationResponseArgs struct {
	AuthorizationType pulumi.StringInput    `pulumi:"authorizationType"`
	Parameters        pulumi.StringMapInput `pulumi:"parameters"`
}

func (AuthorizationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationResponse)(nil)).Elem()
}

func (i AuthorizationResponseArgs) ToAuthorizationResponseOutput() AuthorizationResponseOutput {
	return i.ToAuthorizationResponseOutputWithContext(context.Background())
}

func (i AuthorizationResponseArgs) ToAuthorizationResponseOutputWithContext(ctx context.Context) AuthorizationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationResponseOutput)
}

func (i AuthorizationResponseArgs) ToAuthorizationResponsePtrOutput() AuthorizationResponsePtrOutput {
	return i.ToAuthorizationResponsePtrOutputWithContext(context.Background())
}

func (i AuthorizationResponseArgs) ToAuthorizationResponsePtrOutputWithContext(ctx context.Context) AuthorizationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationResponseOutput).ToAuthorizationResponsePtrOutputWithContext(ctx)
}









type AuthorizationResponsePtrInput interface {
	pulumi.Input

	ToAuthorizationResponsePtrOutput() AuthorizationResponsePtrOutput
	ToAuthorizationResponsePtrOutputWithContext(context.Context) AuthorizationResponsePtrOutput
}

type authorizationResponsePtrType AuthorizationResponseArgs

func AuthorizationResponsePtr(v *AuthorizationResponseArgs) AuthorizationResponsePtrInput {
	return (*authorizationResponsePtrType)(v)
}

func (*authorizationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationResponse)(nil)).Elem()
}

func (i *authorizationResponsePtrType) ToAuthorizationResponsePtrOutput() AuthorizationResponsePtrOutput {
	return i.ToAuthorizationResponsePtrOutputWithContext(context.Background())
}

func (i *authorizationResponsePtrType) ToAuthorizationResponsePtrOutputWithContext(ctx context.Context) AuthorizationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationResponsePtrOutput)
}

type AuthorizationResponseOutput struct{ *pulumi.OutputState }

func (AuthorizationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationResponse)(nil)).Elem()
}

func (o AuthorizationResponseOutput) ToAuthorizationResponseOutput() AuthorizationResponseOutput {
	return o
}

func (o AuthorizationResponseOutput) ToAuthorizationResponseOutputWithContext(ctx context.Context) AuthorizationResponseOutput {
	return o
}

func (o AuthorizationResponseOutput) ToAuthorizationResponsePtrOutput() AuthorizationResponsePtrOutput {
	return o.ToAuthorizationResponsePtrOutputWithContext(context.Background())
}

func (o AuthorizationResponseOutput) ToAuthorizationResponsePtrOutputWithContext(ctx context.Context) AuthorizationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthorizationResponse) *AuthorizationResponse {
		return &v
	}).(AuthorizationResponsePtrOutput)
}

func (o AuthorizationResponseOutput) AuthorizationType() pulumi.StringOutput {
	return o.ApplyT(func(v AuthorizationResponse) string { return v.AuthorizationType }).(pulumi.StringOutput)
}

func (o AuthorizationResponseOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v AuthorizationResponse) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

type AuthorizationResponsePtrOutput struct{ *pulumi.OutputState }

func (AuthorizationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationResponse)(nil)).Elem()
}

func (o AuthorizationResponsePtrOutput) ToAuthorizationResponsePtrOutput() AuthorizationResponsePtrOutput {
	return o
}

func (o AuthorizationResponsePtrOutput) ToAuthorizationResponsePtrOutputWithContext(ctx context.Context) AuthorizationResponsePtrOutput {
	return o
}

func (o AuthorizationResponsePtrOutput) Elem() AuthorizationResponseOutput {
	return o.ApplyT(func(v *AuthorizationResponse) AuthorizationResponse {
		if v != nil {
			return *v
		}
		var ret AuthorizationResponse
		return ret
	}).(AuthorizationResponseOutput)
}

func (o AuthorizationResponsePtrOutput) AuthorizationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AuthorizationType
	}).(pulumi.StringPtrOutput)
}

func (o AuthorizationResponsePtrOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AuthorizationResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.StringMapOutput)
}

type BootstrapConfiguration struct {
	Repository *CodeRepository  `pulumi:"repository"`
	Template   PipelineTemplate `pulumi:"template"`
}





type BootstrapConfigurationInput interface {
	pulumi.Input

	ToBootstrapConfigurationOutput() BootstrapConfigurationOutput
	ToBootstrapConfigurationOutputWithContext(context.Context) BootstrapConfigurationOutput
}

type BootstrapConfigurationArgs struct {
	Repository CodeRepositoryPtrInput `pulumi:"repository"`
	Template   PipelineTemplateInput  `pulumi:"template"`
}

func (BootstrapConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BootstrapConfiguration)(nil)).Elem()
}

func (i BootstrapConfigurationArgs) ToBootstrapConfigurationOutput() BootstrapConfigurationOutput {
	return i.ToBootstrapConfigurationOutputWithContext(context.Background())
}

func (i BootstrapConfigurationArgs) ToBootstrapConfigurationOutputWithContext(ctx context.Context) BootstrapConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootstrapConfigurationOutput)
}

func (i BootstrapConfigurationArgs) ToBootstrapConfigurationPtrOutput() BootstrapConfigurationPtrOutput {
	return i.ToBootstrapConfigurationPtrOutputWithContext(context.Background())
}

func (i BootstrapConfigurationArgs) ToBootstrapConfigurationPtrOutputWithContext(ctx context.Context) BootstrapConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootstrapConfigurationOutput).ToBootstrapConfigurationPtrOutputWithContext(ctx)
}









type BootstrapConfigurationPtrInput interface {
	pulumi.Input

	ToBootstrapConfigurationPtrOutput() BootstrapConfigurationPtrOutput
	ToBootstrapConfigurationPtrOutputWithContext(context.Context) BootstrapConfigurationPtrOutput
}

type bootstrapConfigurationPtrType BootstrapConfigurationArgs

func BootstrapConfigurationPtr(v *BootstrapConfigurationArgs) BootstrapConfigurationPtrInput {
	return (*bootstrapConfigurationPtrType)(v)
}

func (*bootstrapConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BootstrapConfiguration)(nil)).Elem()
}

func (i *bootstrapConfigurationPtrType) ToBootstrapConfigurationPtrOutput() BootstrapConfigurationPtrOutput {
	return i.ToBootstrapConfigurationPtrOutputWithContext(context.Background())
}

func (i *bootstrapConfigurationPtrType) ToBootstrapConfigurationPtrOutputWithContext(ctx context.Context) BootstrapConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootstrapConfigurationPtrOutput)
}

type BootstrapConfigurationOutput struct{ *pulumi.OutputState }

func (BootstrapConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BootstrapConfiguration)(nil)).Elem()
}

func (o BootstrapConfigurationOutput) ToBootstrapConfigurationOutput() BootstrapConfigurationOutput {
	return o
}

func (o BootstrapConfigurationOutput) ToBootstrapConfigurationOutputWithContext(ctx context.Context) BootstrapConfigurationOutput {
	return o
}

func (o BootstrapConfigurationOutput) ToBootstrapConfigurationPtrOutput() BootstrapConfigurationPtrOutput {
	return o.ToBootstrapConfigurationPtrOutputWithContext(context.Background())
}

func (o BootstrapConfigurationOutput) ToBootstrapConfigurationPtrOutputWithContext(ctx context.Context) BootstrapConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BootstrapConfiguration) *BootstrapConfiguration {
		return &v
	}).(BootstrapConfigurationPtrOutput)
}

func (o BootstrapConfigurationOutput) Repository() CodeRepositoryPtrOutput {
	return o.ApplyT(func(v BootstrapConfiguration) *CodeRepository { return v.Repository }).(CodeRepositoryPtrOutput)
}

func (o BootstrapConfigurationOutput) Template() PipelineTemplateOutput {
	return o.ApplyT(func(v BootstrapConfiguration) PipelineTemplate { return v.Template }).(PipelineTemplateOutput)
}

type BootstrapConfigurationPtrOutput struct{ *pulumi.OutputState }

func (BootstrapConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BootstrapConfiguration)(nil)).Elem()
}

func (o BootstrapConfigurationPtrOutput) ToBootstrapConfigurationPtrOutput() BootstrapConfigurationPtrOutput {
	return o
}

func (o BootstrapConfigurationPtrOutput) ToBootstrapConfigurationPtrOutputWithContext(ctx context.Context) BootstrapConfigurationPtrOutput {
	return o
}

func (o BootstrapConfigurationPtrOutput) Elem() BootstrapConfigurationOutput {
	return o.ApplyT(func(v *BootstrapConfiguration) BootstrapConfiguration {
		if v != nil {
			return *v
		}
		var ret BootstrapConfiguration
		return ret
	}).(BootstrapConfigurationOutput)
}

func (o BootstrapConfigurationPtrOutput) Repository() CodeRepositoryPtrOutput {
	return o.ApplyT(func(v *BootstrapConfiguration) *CodeRepository {
		if v == nil {
			return nil
		}
		return v.Repository
	}).(CodeRepositoryPtrOutput)
}

func (o BootstrapConfigurationPtrOutput) Template() PipelineTemplatePtrOutput {
	return o.ApplyT(func(v *BootstrapConfiguration) *PipelineTemplate {
		if v == nil {
			return nil
		}
		return &v.Template
	}).(PipelineTemplatePtrOutput)
}

type BootstrapConfigurationResponse struct {
	Repository *CodeRepositoryResponse  `pulumi:"repository"`
	Template   PipelineTemplateResponse `pulumi:"template"`
}





type BootstrapConfigurationResponseInput interface {
	pulumi.Input

	ToBootstrapConfigurationResponseOutput() BootstrapConfigurationResponseOutput
	ToBootstrapConfigurationResponseOutputWithContext(context.Context) BootstrapConfigurationResponseOutput
}

type BootstrapConfigurationResponseArgs struct {
	Repository CodeRepositoryResponsePtrInput `pulumi:"repository"`
	Template   PipelineTemplateResponseInput  `pulumi:"template"`
}

func (BootstrapConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BootstrapConfigurationResponse)(nil)).Elem()
}

func (i BootstrapConfigurationResponseArgs) ToBootstrapConfigurationResponseOutput() BootstrapConfigurationResponseOutput {
	return i.ToBootstrapConfigurationResponseOutputWithContext(context.Background())
}

func (i BootstrapConfigurationResponseArgs) ToBootstrapConfigurationResponseOutputWithContext(ctx context.Context) BootstrapConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootstrapConfigurationResponseOutput)
}

func (i BootstrapConfigurationResponseArgs) ToBootstrapConfigurationResponsePtrOutput() BootstrapConfigurationResponsePtrOutput {
	return i.ToBootstrapConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i BootstrapConfigurationResponseArgs) ToBootstrapConfigurationResponsePtrOutputWithContext(ctx context.Context) BootstrapConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootstrapConfigurationResponseOutput).ToBootstrapConfigurationResponsePtrOutputWithContext(ctx)
}









type BootstrapConfigurationResponsePtrInput interface {
	pulumi.Input

	ToBootstrapConfigurationResponsePtrOutput() BootstrapConfigurationResponsePtrOutput
	ToBootstrapConfigurationResponsePtrOutputWithContext(context.Context) BootstrapConfigurationResponsePtrOutput
}

type bootstrapConfigurationResponsePtrType BootstrapConfigurationResponseArgs

func BootstrapConfigurationResponsePtr(v *BootstrapConfigurationResponseArgs) BootstrapConfigurationResponsePtrInput {
	return (*bootstrapConfigurationResponsePtrType)(v)
}

func (*bootstrapConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BootstrapConfigurationResponse)(nil)).Elem()
}

func (i *bootstrapConfigurationResponsePtrType) ToBootstrapConfigurationResponsePtrOutput() BootstrapConfigurationResponsePtrOutput {
	return i.ToBootstrapConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *bootstrapConfigurationResponsePtrType) ToBootstrapConfigurationResponsePtrOutputWithContext(ctx context.Context) BootstrapConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootstrapConfigurationResponsePtrOutput)
}

type BootstrapConfigurationResponseOutput struct{ *pulumi.OutputState }

func (BootstrapConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BootstrapConfigurationResponse)(nil)).Elem()
}

func (o BootstrapConfigurationResponseOutput) ToBootstrapConfigurationResponseOutput() BootstrapConfigurationResponseOutput {
	return o
}

func (o BootstrapConfigurationResponseOutput) ToBootstrapConfigurationResponseOutputWithContext(ctx context.Context) BootstrapConfigurationResponseOutput {
	return o
}

func (o BootstrapConfigurationResponseOutput) ToBootstrapConfigurationResponsePtrOutput() BootstrapConfigurationResponsePtrOutput {
	return o.ToBootstrapConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o BootstrapConfigurationResponseOutput) ToBootstrapConfigurationResponsePtrOutputWithContext(ctx context.Context) BootstrapConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BootstrapConfigurationResponse) *BootstrapConfigurationResponse {
		return &v
	}).(BootstrapConfigurationResponsePtrOutput)
}

func (o BootstrapConfigurationResponseOutput) Repository() CodeRepositoryResponsePtrOutput {
	return o.ApplyT(func(v BootstrapConfigurationResponse) *CodeRepositoryResponse { return v.Repository }).(CodeRepositoryResponsePtrOutput)
}

func (o BootstrapConfigurationResponseOutput) Template() PipelineTemplateResponseOutput {
	return o.ApplyT(func(v BootstrapConfigurationResponse) PipelineTemplateResponse { return v.Template }).(PipelineTemplateResponseOutput)
}

type BootstrapConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (BootstrapConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BootstrapConfigurationResponse)(nil)).Elem()
}

func (o BootstrapConfigurationResponsePtrOutput) ToBootstrapConfigurationResponsePtrOutput() BootstrapConfigurationResponsePtrOutput {
	return o
}

func (o BootstrapConfigurationResponsePtrOutput) ToBootstrapConfigurationResponsePtrOutputWithContext(ctx context.Context) BootstrapConfigurationResponsePtrOutput {
	return o
}

func (o BootstrapConfigurationResponsePtrOutput) Elem() BootstrapConfigurationResponseOutput {
	return o.ApplyT(func(v *BootstrapConfigurationResponse) BootstrapConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret BootstrapConfigurationResponse
		return ret
	}).(BootstrapConfigurationResponseOutput)
}

func (o BootstrapConfigurationResponsePtrOutput) Repository() CodeRepositoryResponsePtrOutput {
	return o.ApplyT(func(v *BootstrapConfigurationResponse) *CodeRepositoryResponse {
		if v == nil {
			return nil
		}
		return v.Repository
	}).(CodeRepositoryResponsePtrOutput)
}

func (o BootstrapConfigurationResponsePtrOutput) Template() PipelineTemplateResponsePtrOutput {
	return o.ApplyT(func(v *BootstrapConfigurationResponse) *PipelineTemplateResponse {
		if v == nil {
			return nil
		}
		return &v.Template
	}).(PipelineTemplateResponsePtrOutput)
}

type CodeRepository struct {
	Authorization  *Authorization    `pulumi:"authorization"`
	DefaultBranch  string            `pulumi:"defaultBranch"`
	Id             string            `pulumi:"id"`
	Properties     map[string]string `pulumi:"properties"`
	RepositoryType string            `pulumi:"repositoryType"`
}





type CodeRepositoryInput interface {
	pulumi.Input

	ToCodeRepositoryOutput() CodeRepositoryOutput
	ToCodeRepositoryOutputWithContext(context.Context) CodeRepositoryOutput
}

type CodeRepositoryArgs struct {
	Authorization  AuthorizationPtrInput `pulumi:"authorization"`
	DefaultBranch  pulumi.StringInput    `pulumi:"defaultBranch"`
	Id             pulumi.StringInput    `pulumi:"id"`
	Properties     pulumi.StringMapInput `pulumi:"properties"`
	RepositoryType pulumi.StringInput    `pulumi:"repositoryType"`
}

func (CodeRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodeRepository)(nil)).Elem()
}

func (i CodeRepositoryArgs) ToCodeRepositoryOutput() CodeRepositoryOutput {
	return i.ToCodeRepositoryOutputWithContext(context.Background())
}

func (i CodeRepositoryArgs) ToCodeRepositoryOutputWithContext(ctx context.Context) CodeRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeRepositoryOutput)
}

func (i CodeRepositoryArgs) ToCodeRepositoryPtrOutput() CodeRepositoryPtrOutput {
	return i.ToCodeRepositoryPtrOutputWithContext(context.Background())
}

func (i CodeRepositoryArgs) ToCodeRepositoryPtrOutputWithContext(ctx context.Context) CodeRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeRepositoryOutput).ToCodeRepositoryPtrOutputWithContext(ctx)
}









type CodeRepositoryPtrInput interface {
	pulumi.Input

	ToCodeRepositoryPtrOutput() CodeRepositoryPtrOutput
	ToCodeRepositoryPtrOutputWithContext(context.Context) CodeRepositoryPtrOutput
}

type codeRepositoryPtrType CodeRepositoryArgs

func CodeRepositoryPtr(v *CodeRepositoryArgs) CodeRepositoryPtrInput {
	return (*codeRepositoryPtrType)(v)
}

func (*codeRepositoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CodeRepository)(nil)).Elem()
}

func (i *codeRepositoryPtrType) ToCodeRepositoryPtrOutput() CodeRepositoryPtrOutput {
	return i.ToCodeRepositoryPtrOutputWithContext(context.Background())
}

func (i *codeRepositoryPtrType) ToCodeRepositoryPtrOutputWithContext(ctx context.Context) CodeRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeRepositoryPtrOutput)
}

type CodeRepositoryOutput struct{ *pulumi.OutputState }

func (CodeRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodeRepository)(nil)).Elem()
}

func (o CodeRepositoryOutput) ToCodeRepositoryOutput() CodeRepositoryOutput {
	return o
}

func (o CodeRepositoryOutput) ToCodeRepositoryOutputWithContext(ctx context.Context) CodeRepositoryOutput {
	return o
}

func (o CodeRepositoryOutput) ToCodeRepositoryPtrOutput() CodeRepositoryPtrOutput {
	return o.ToCodeRepositoryPtrOutputWithContext(context.Background())
}

func (o CodeRepositoryOutput) ToCodeRepositoryPtrOutputWithContext(ctx context.Context) CodeRepositoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CodeRepository) *CodeRepository {
		return &v
	}).(CodeRepositoryPtrOutput)
}

func (o CodeRepositoryOutput) Authorization() AuthorizationPtrOutput {
	return o.ApplyT(func(v CodeRepository) *Authorization { return v.Authorization }).(AuthorizationPtrOutput)
}

func (o CodeRepositoryOutput) DefaultBranch() pulumi.StringOutput {
	return o.ApplyT(func(v CodeRepository) string { return v.DefaultBranch }).(pulumi.StringOutput)
}

func (o CodeRepositoryOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CodeRepository) string { return v.Id }).(pulumi.StringOutput)
}

func (o CodeRepositoryOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v CodeRepository) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o CodeRepositoryOutput) RepositoryType() pulumi.StringOutput {
	return o.ApplyT(func(v CodeRepository) string { return v.RepositoryType }).(pulumi.StringOutput)
}

type CodeRepositoryPtrOutput struct{ *pulumi.OutputState }

func (CodeRepositoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodeRepository)(nil)).Elem()
}

func (o CodeRepositoryPtrOutput) ToCodeRepositoryPtrOutput() CodeRepositoryPtrOutput {
	return o
}

func (o CodeRepositoryPtrOutput) ToCodeRepositoryPtrOutputWithContext(ctx context.Context) CodeRepositoryPtrOutput {
	return o
}

func (o CodeRepositoryPtrOutput) Elem() CodeRepositoryOutput {
	return o.ApplyT(func(v *CodeRepository) CodeRepository {
		if v != nil {
			return *v
		}
		var ret CodeRepository
		return ret
	}).(CodeRepositoryOutput)
}

func (o CodeRepositoryPtrOutput) Authorization() AuthorizationPtrOutput {
	return o.ApplyT(func(v *CodeRepository) *Authorization {
		if v == nil {
			return nil
		}
		return v.Authorization
	}).(AuthorizationPtrOutput)
}

func (o CodeRepositoryPtrOutput) DefaultBranch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodeRepository) *string {
		if v == nil {
			return nil
		}
		return &v.DefaultBranch
	}).(pulumi.StringPtrOutput)
}

func (o CodeRepositoryPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodeRepository) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o CodeRepositoryPtrOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CodeRepository) map[string]string {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.StringMapOutput)
}

func (o CodeRepositoryPtrOutput) RepositoryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodeRepository) *string {
		if v == nil {
			return nil
		}
		return &v.RepositoryType
	}).(pulumi.StringPtrOutput)
}

type CodeRepositoryResponse struct {
	Authorization  *AuthorizationResponse `pulumi:"authorization"`
	DefaultBranch  string                 `pulumi:"defaultBranch"`
	Id             string                 `pulumi:"id"`
	Properties     map[string]string      `pulumi:"properties"`
	RepositoryType string                 `pulumi:"repositoryType"`
}





type CodeRepositoryResponseInput interface {
	pulumi.Input

	ToCodeRepositoryResponseOutput() CodeRepositoryResponseOutput
	ToCodeRepositoryResponseOutputWithContext(context.Context) CodeRepositoryResponseOutput
}

type CodeRepositoryResponseArgs struct {
	Authorization  AuthorizationResponsePtrInput `pulumi:"authorization"`
	DefaultBranch  pulumi.StringInput            `pulumi:"defaultBranch"`
	Id             pulumi.StringInput            `pulumi:"id"`
	Properties     pulumi.StringMapInput         `pulumi:"properties"`
	RepositoryType pulumi.StringInput            `pulumi:"repositoryType"`
}

func (CodeRepositoryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodeRepositoryResponse)(nil)).Elem()
}

func (i CodeRepositoryResponseArgs) ToCodeRepositoryResponseOutput() CodeRepositoryResponseOutput {
	return i.ToCodeRepositoryResponseOutputWithContext(context.Background())
}

func (i CodeRepositoryResponseArgs) ToCodeRepositoryResponseOutputWithContext(ctx context.Context) CodeRepositoryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeRepositoryResponseOutput)
}

func (i CodeRepositoryResponseArgs) ToCodeRepositoryResponsePtrOutput() CodeRepositoryResponsePtrOutput {
	return i.ToCodeRepositoryResponsePtrOutputWithContext(context.Background())
}

func (i CodeRepositoryResponseArgs) ToCodeRepositoryResponsePtrOutputWithContext(ctx context.Context) CodeRepositoryResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeRepositoryResponseOutput).ToCodeRepositoryResponsePtrOutputWithContext(ctx)
}









type CodeRepositoryResponsePtrInput interface {
	pulumi.Input

	ToCodeRepositoryResponsePtrOutput() CodeRepositoryResponsePtrOutput
	ToCodeRepositoryResponsePtrOutputWithContext(context.Context) CodeRepositoryResponsePtrOutput
}

type codeRepositoryResponsePtrType CodeRepositoryResponseArgs

func CodeRepositoryResponsePtr(v *CodeRepositoryResponseArgs) CodeRepositoryResponsePtrInput {
	return (*codeRepositoryResponsePtrType)(v)
}

func (*codeRepositoryResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CodeRepositoryResponse)(nil)).Elem()
}

func (i *codeRepositoryResponsePtrType) ToCodeRepositoryResponsePtrOutput() CodeRepositoryResponsePtrOutput {
	return i.ToCodeRepositoryResponsePtrOutputWithContext(context.Background())
}

func (i *codeRepositoryResponsePtrType) ToCodeRepositoryResponsePtrOutputWithContext(ctx context.Context) CodeRepositoryResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeRepositoryResponsePtrOutput)
}

type CodeRepositoryResponseOutput struct{ *pulumi.OutputState }

func (CodeRepositoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodeRepositoryResponse)(nil)).Elem()
}

func (o CodeRepositoryResponseOutput) ToCodeRepositoryResponseOutput() CodeRepositoryResponseOutput {
	return o
}

func (o CodeRepositoryResponseOutput) ToCodeRepositoryResponseOutputWithContext(ctx context.Context) CodeRepositoryResponseOutput {
	return o
}

func (o CodeRepositoryResponseOutput) ToCodeRepositoryResponsePtrOutput() CodeRepositoryResponsePtrOutput {
	return o.ToCodeRepositoryResponsePtrOutputWithContext(context.Background())
}

func (o CodeRepositoryResponseOutput) ToCodeRepositoryResponsePtrOutputWithContext(ctx context.Context) CodeRepositoryResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CodeRepositoryResponse) *CodeRepositoryResponse {
		return &v
	}).(CodeRepositoryResponsePtrOutput)
}

func (o CodeRepositoryResponseOutput) Authorization() AuthorizationResponsePtrOutput {
	return o.ApplyT(func(v CodeRepositoryResponse) *AuthorizationResponse { return v.Authorization }).(AuthorizationResponsePtrOutput)
}

func (o CodeRepositoryResponseOutput) DefaultBranch() pulumi.StringOutput {
	return o.ApplyT(func(v CodeRepositoryResponse) string { return v.DefaultBranch }).(pulumi.StringOutput)
}

func (o CodeRepositoryResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CodeRepositoryResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o CodeRepositoryResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v CodeRepositoryResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o CodeRepositoryResponseOutput) RepositoryType() pulumi.StringOutput {
	return o.ApplyT(func(v CodeRepositoryResponse) string { return v.RepositoryType }).(pulumi.StringOutput)
}

type CodeRepositoryResponsePtrOutput struct{ *pulumi.OutputState }

func (CodeRepositoryResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodeRepositoryResponse)(nil)).Elem()
}

func (o CodeRepositoryResponsePtrOutput) ToCodeRepositoryResponsePtrOutput() CodeRepositoryResponsePtrOutput {
	return o
}

func (o CodeRepositoryResponsePtrOutput) ToCodeRepositoryResponsePtrOutputWithContext(ctx context.Context) CodeRepositoryResponsePtrOutput {
	return o
}

func (o CodeRepositoryResponsePtrOutput) Elem() CodeRepositoryResponseOutput {
	return o.ApplyT(func(v *CodeRepositoryResponse) CodeRepositoryResponse {
		if v != nil {
			return *v
		}
		var ret CodeRepositoryResponse
		return ret
	}).(CodeRepositoryResponseOutput)
}

func (o CodeRepositoryResponsePtrOutput) Authorization() AuthorizationResponsePtrOutput {
	return o.ApplyT(func(v *CodeRepositoryResponse) *AuthorizationResponse {
		if v == nil {
			return nil
		}
		return v.Authorization
	}).(AuthorizationResponsePtrOutput)
}

func (o CodeRepositoryResponsePtrOutput) DefaultBranch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodeRepositoryResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DefaultBranch
	}).(pulumi.StringPtrOutput)
}

func (o CodeRepositoryResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodeRepositoryResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o CodeRepositoryResponsePtrOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CodeRepositoryResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.StringMapOutput)
}

func (o CodeRepositoryResponsePtrOutput) RepositoryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodeRepositoryResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RepositoryType
	}).(pulumi.StringPtrOutput)
}

type OrganizationReference struct {
	Name string `pulumi:"name"`
}





type OrganizationReferenceInput interface {
	pulumi.Input

	ToOrganizationReferenceOutput() OrganizationReferenceOutput
	ToOrganizationReferenceOutputWithContext(context.Context) OrganizationReferenceOutput
}

type OrganizationReferenceArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (OrganizationReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationReference)(nil)).Elem()
}

func (i OrganizationReferenceArgs) ToOrganizationReferenceOutput() OrganizationReferenceOutput {
	return i.ToOrganizationReferenceOutputWithContext(context.Background())
}

func (i OrganizationReferenceArgs) ToOrganizationReferenceOutputWithContext(ctx context.Context) OrganizationReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationReferenceOutput)
}

func (i OrganizationReferenceArgs) ToOrganizationReferencePtrOutput() OrganizationReferencePtrOutput {
	return i.ToOrganizationReferencePtrOutputWithContext(context.Background())
}

func (i OrganizationReferenceArgs) ToOrganizationReferencePtrOutputWithContext(ctx context.Context) OrganizationReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationReferenceOutput).ToOrganizationReferencePtrOutputWithContext(ctx)
}









type OrganizationReferencePtrInput interface {
	pulumi.Input

	ToOrganizationReferencePtrOutput() OrganizationReferencePtrOutput
	ToOrganizationReferencePtrOutputWithContext(context.Context) OrganizationReferencePtrOutput
}

type organizationReferencePtrType OrganizationReferenceArgs

func OrganizationReferencePtr(v *OrganizationReferenceArgs) OrganizationReferencePtrInput {
	return (*organizationReferencePtrType)(v)
}

func (*organizationReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationReference)(nil)).Elem()
}

func (i *organizationReferencePtrType) ToOrganizationReferencePtrOutput() OrganizationReferencePtrOutput {
	return i.ToOrganizationReferencePtrOutputWithContext(context.Background())
}

func (i *organizationReferencePtrType) ToOrganizationReferencePtrOutputWithContext(ctx context.Context) OrganizationReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationReferencePtrOutput)
}

type OrganizationReferenceOutput struct{ *pulumi.OutputState }

func (OrganizationReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationReference)(nil)).Elem()
}

func (o OrganizationReferenceOutput) ToOrganizationReferenceOutput() OrganizationReferenceOutput {
	return o
}

func (o OrganizationReferenceOutput) ToOrganizationReferenceOutputWithContext(ctx context.Context) OrganizationReferenceOutput {
	return o
}

func (o OrganizationReferenceOutput) ToOrganizationReferencePtrOutput() OrganizationReferencePtrOutput {
	return o.ToOrganizationReferencePtrOutputWithContext(context.Background())
}

func (o OrganizationReferenceOutput) ToOrganizationReferencePtrOutputWithContext(ctx context.Context) OrganizationReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OrganizationReference) *OrganizationReference {
		return &v
	}).(OrganizationReferencePtrOutput)
}

func (o OrganizationReferenceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v OrganizationReference) string { return v.Name }).(pulumi.StringOutput)
}

type OrganizationReferencePtrOutput struct{ *pulumi.OutputState }

func (OrganizationReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationReference)(nil)).Elem()
}

func (o OrganizationReferencePtrOutput) ToOrganizationReferencePtrOutput() OrganizationReferencePtrOutput {
	return o
}

func (o OrganizationReferencePtrOutput) ToOrganizationReferencePtrOutputWithContext(ctx context.Context) OrganizationReferencePtrOutput {
	return o
}

func (o OrganizationReferencePtrOutput) Elem() OrganizationReferenceOutput {
	return o.ApplyT(func(v *OrganizationReference) OrganizationReference {
		if v != nil {
			return *v
		}
		var ret OrganizationReference
		return ret
	}).(OrganizationReferenceOutput)
}

func (o OrganizationReferencePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationReference) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type OrganizationReferenceResponse struct {
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}





type OrganizationReferenceResponseInput interface {
	pulumi.Input

	ToOrganizationReferenceResponseOutput() OrganizationReferenceResponseOutput
	ToOrganizationReferenceResponseOutputWithContext(context.Context) OrganizationReferenceResponseOutput
}

type OrganizationReferenceResponseArgs struct {
	Id   pulumi.StringInput `pulumi:"id"`
	Name pulumi.StringInput `pulumi:"name"`
}

func (OrganizationReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationReferenceResponse)(nil)).Elem()
}

func (i OrganizationReferenceResponseArgs) ToOrganizationReferenceResponseOutput() OrganizationReferenceResponseOutput {
	return i.ToOrganizationReferenceResponseOutputWithContext(context.Background())
}

func (i OrganizationReferenceResponseArgs) ToOrganizationReferenceResponseOutputWithContext(ctx context.Context) OrganizationReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationReferenceResponseOutput)
}

func (i OrganizationReferenceResponseArgs) ToOrganizationReferenceResponsePtrOutput() OrganizationReferenceResponsePtrOutput {
	return i.ToOrganizationReferenceResponsePtrOutputWithContext(context.Background())
}

func (i OrganizationReferenceResponseArgs) ToOrganizationReferenceResponsePtrOutputWithContext(ctx context.Context) OrganizationReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationReferenceResponseOutput).ToOrganizationReferenceResponsePtrOutputWithContext(ctx)
}









type OrganizationReferenceResponsePtrInput interface {
	pulumi.Input

	ToOrganizationReferenceResponsePtrOutput() OrganizationReferenceResponsePtrOutput
	ToOrganizationReferenceResponsePtrOutputWithContext(context.Context) OrganizationReferenceResponsePtrOutput
}

type organizationReferenceResponsePtrType OrganizationReferenceResponseArgs

func OrganizationReferenceResponsePtr(v *OrganizationReferenceResponseArgs) OrganizationReferenceResponsePtrInput {
	return (*organizationReferenceResponsePtrType)(v)
}

func (*organizationReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationReferenceResponse)(nil)).Elem()
}

func (i *organizationReferenceResponsePtrType) ToOrganizationReferenceResponsePtrOutput() OrganizationReferenceResponsePtrOutput {
	return i.ToOrganizationReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *organizationReferenceResponsePtrType) ToOrganizationReferenceResponsePtrOutputWithContext(ctx context.Context) OrganizationReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationReferenceResponsePtrOutput)
}

type OrganizationReferenceResponseOutput struct{ *pulumi.OutputState }

func (OrganizationReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationReferenceResponse)(nil)).Elem()
}

func (o OrganizationReferenceResponseOutput) ToOrganizationReferenceResponseOutput() OrganizationReferenceResponseOutput {
	return o
}

func (o OrganizationReferenceResponseOutput) ToOrganizationReferenceResponseOutputWithContext(ctx context.Context) OrganizationReferenceResponseOutput {
	return o
}

func (o OrganizationReferenceResponseOutput) ToOrganizationReferenceResponsePtrOutput() OrganizationReferenceResponsePtrOutput {
	return o.ToOrganizationReferenceResponsePtrOutputWithContext(context.Background())
}

func (o OrganizationReferenceResponseOutput) ToOrganizationReferenceResponsePtrOutputWithContext(ctx context.Context) OrganizationReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OrganizationReferenceResponse) *OrganizationReferenceResponse {
		return &v
	}).(OrganizationReferenceResponsePtrOutput)
}

func (o OrganizationReferenceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v OrganizationReferenceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o OrganizationReferenceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v OrganizationReferenceResponse) string { return v.Name }).(pulumi.StringOutput)
}

type OrganizationReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (OrganizationReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationReferenceResponse)(nil)).Elem()
}

func (o OrganizationReferenceResponsePtrOutput) ToOrganizationReferenceResponsePtrOutput() OrganizationReferenceResponsePtrOutput {
	return o
}

func (o OrganizationReferenceResponsePtrOutput) ToOrganizationReferenceResponsePtrOutputWithContext(ctx context.Context) OrganizationReferenceResponsePtrOutput {
	return o
}

func (o OrganizationReferenceResponsePtrOutput) Elem() OrganizationReferenceResponseOutput {
	return o.ApplyT(func(v *OrganizationReferenceResponse) OrganizationReferenceResponse {
		if v != nil {
			return *v
		}
		var ret OrganizationReferenceResponse
		return ret
	}).(OrganizationReferenceResponseOutput)
}

func (o OrganizationReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o OrganizationReferenceResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type PipelineTemplate struct {
	Id         string            `pulumi:"id"`
	Parameters map[string]string `pulumi:"parameters"`
}





type PipelineTemplateInput interface {
	pulumi.Input

	ToPipelineTemplateOutput() PipelineTemplateOutput
	ToPipelineTemplateOutputWithContext(context.Context) PipelineTemplateOutput
}

type PipelineTemplateArgs struct {
	Id         pulumi.StringInput    `pulumi:"id"`
	Parameters pulumi.StringMapInput `pulumi:"parameters"`
}

func (PipelineTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineTemplate)(nil)).Elem()
}

func (i PipelineTemplateArgs) ToPipelineTemplateOutput() PipelineTemplateOutput {
	return i.ToPipelineTemplateOutputWithContext(context.Background())
}

func (i PipelineTemplateArgs) ToPipelineTemplateOutputWithContext(ctx context.Context) PipelineTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTemplateOutput)
}

func (i PipelineTemplateArgs) ToPipelineTemplatePtrOutput() PipelineTemplatePtrOutput {
	return i.ToPipelineTemplatePtrOutputWithContext(context.Background())
}

func (i PipelineTemplateArgs) ToPipelineTemplatePtrOutputWithContext(ctx context.Context) PipelineTemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTemplateOutput).ToPipelineTemplatePtrOutputWithContext(ctx)
}









type PipelineTemplatePtrInput interface {
	pulumi.Input

	ToPipelineTemplatePtrOutput() PipelineTemplatePtrOutput
	ToPipelineTemplatePtrOutputWithContext(context.Context) PipelineTemplatePtrOutput
}

type pipelineTemplatePtrType PipelineTemplateArgs

func PipelineTemplatePtr(v *PipelineTemplateArgs) PipelineTemplatePtrInput {
	return (*pipelineTemplatePtrType)(v)
}

func (*pipelineTemplatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineTemplate)(nil)).Elem()
}

func (i *pipelineTemplatePtrType) ToPipelineTemplatePtrOutput() PipelineTemplatePtrOutput {
	return i.ToPipelineTemplatePtrOutputWithContext(context.Background())
}

func (i *pipelineTemplatePtrType) ToPipelineTemplatePtrOutputWithContext(ctx context.Context) PipelineTemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTemplatePtrOutput)
}

type PipelineTemplateOutput struct{ *pulumi.OutputState }

func (PipelineTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineTemplate)(nil)).Elem()
}

func (o PipelineTemplateOutput) ToPipelineTemplateOutput() PipelineTemplateOutput {
	return o
}

func (o PipelineTemplateOutput) ToPipelineTemplateOutputWithContext(ctx context.Context) PipelineTemplateOutput {
	return o
}

func (o PipelineTemplateOutput) ToPipelineTemplatePtrOutput() PipelineTemplatePtrOutput {
	return o.ToPipelineTemplatePtrOutputWithContext(context.Background())
}

func (o PipelineTemplateOutput) ToPipelineTemplatePtrOutputWithContext(ctx context.Context) PipelineTemplatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineTemplate) *PipelineTemplate {
		return &v
	}).(PipelineTemplatePtrOutput)
}

func (o PipelineTemplateOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PipelineTemplate) string { return v.Id }).(pulumi.StringOutput)
}

func (o PipelineTemplateOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v PipelineTemplate) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

type PipelineTemplatePtrOutput struct{ *pulumi.OutputState }

func (PipelineTemplatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineTemplate)(nil)).Elem()
}

func (o PipelineTemplatePtrOutput) ToPipelineTemplatePtrOutput() PipelineTemplatePtrOutput {
	return o
}

func (o PipelineTemplatePtrOutput) ToPipelineTemplatePtrOutputWithContext(ctx context.Context) PipelineTemplatePtrOutput {
	return o
}

func (o PipelineTemplatePtrOutput) Elem() PipelineTemplateOutput {
	return o.ApplyT(func(v *PipelineTemplate) PipelineTemplate {
		if v != nil {
			return *v
		}
		var ret PipelineTemplate
		return ret
	}).(PipelineTemplateOutput)
}

func (o PipelineTemplatePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineTemplate) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o PipelineTemplatePtrOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PipelineTemplate) map[string]string {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.StringMapOutput)
}

type PipelineTemplateResponse struct {
	Id         string            `pulumi:"id"`
	Parameters map[string]string `pulumi:"parameters"`
}





type PipelineTemplateResponseInput interface {
	pulumi.Input

	ToPipelineTemplateResponseOutput() PipelineTemplateResponseOutput
	ToPipelineTemplateResponseOutputWithContext(context.Context) PipelineTemplateResponseOutput
}

type PipelineTemplateResponseArgs struct {
	Id         pulumi.StringInput    `pulumi:"id"`
	Parameters pulumi.StringMapInput `pulumi:"parameters"`
}

func (PipelineTemplateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineTemplateResponse)(nil)).Elem()
}

func (i PipelineTemplateResponseArgs) ToPipelineTemplateResponseOutput() PipelineTemplateResponseOutput {
	return i.ToPipelineTemplateResponseOutputWithContext(context.Background())
}

func (i PipelineTemplateResponseArgs) ToPipelineTemplateResponseOutputWithContext(ctx context.Context) PipelineTemplateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTemplateResponseOutput)
}

func (i PipelineTemplateResponseArgs) ToPipelineTemplateResponsePtrOutput() PipelineTemplateResponsePtrOutput {
	return i.ToPipelineTemplateResponsePtrOutputWithContext(context.Background())
}

func (i PipelineTemplateResponseArgs) ToPipelineTemplateResponsePtrOutputWithContext(ctx context.Context) PipelineTemplateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTemplateResponseOutput).ToPipelineTemplateResponsePtrOutputWithContext(ctx)
}









type PipelineTemplateResponsePtrInput interface {
	pulumi.Input

	ToPipelineTemplateResponsePtrOutput() PipelineTemplateResponsePtrOutput
	ToPipelineTemplateResponsePtrOutputWithContext(context.Context) PipelineTemplateResponsePtrOutput
}

type pipelineTemplateResponsePtrType PipelineTemplateResponseArgs

func PipelineTemplateResponsePtr(v *PipelineTemplateResponseArgs) PipelineTemplateResponsePtrInput {
	return (*pipelineTemplateResponsePtrType)(v)
}

func (*pipelineTemplateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineTemplateResponse)(nil)).Elem()
}

func (i *pipelineTemplateResponsePtrType) ToPipelineTemplateResponsePtrOutput() PipelineTemplateResponsePtrOutput {
	return i.ToPipelineTemplateResponsePtrOutputWithContext(context.Background())
}

func (i *pipelineTemplateResponsePtrType) ToPipelineTemplateResponsePtrOutputWithContext(ctx context.Context) PipelineTemplateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTemplateResponsePtrOutput)
}

type PipelineTemplateResponseOutput struct{ *pulumi.OutputState }

func (PipelineTemplateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineTemplateResponse)(nil)).Elem()
}

func (o PipelineTemplateResponseOutput) ToPipelineTemplateResponseOutput() PipelineTemplateResponseOutput {
	return o
}

func (o PipelineTemplateResponseOutput) ToPipelineTemplateResponseOutputWithContext(ctx context.Context) PipelineTemplateResponseOutput {
	return o
}

func (o PipelineTemplateResponseOutput) ToPipelineTemplateResponsePtrOutput() PipelineTemplateResponsePtrOutput {
	return o.ToPipelineTemplateResponsePtrOutputWithContext(context.Background())
}

func (o PipelineTemplateResponseOutput) ToPipelineTemplateResponsePtrOutputWithContext(ctx context.Context) PipelineTemplateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineTemplateResponse) *PipelineTemplateResponse {
		return &v
	}).(PipelineTemplateResponsePtrOutput)
}

func (o PipelineTemplateResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PipelineTemplateResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PipelineTemplateResponseOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v PipelineTemplateResponse) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

type PipelineTemplateResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineTemplateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineTemplateResponse)(nil)).Elem()
}

func (o PipelineTemplateResponsePtrOutput) ToPipelineTemplateResponsePtrOutput() PipelineTemplateResponsePtrOutput {
	return o
}

func (o PipelineTemplateResponsePtrOutput) ToPipelineTemplateResponsePtrOutputWithContext(ctx context.Context) PipelineTemplateResponsePtrOutput {
	return o
}

func (o PipelineTemplateResponsePtrOutput) Elem() PipelineTemplateResponseOutput {
	return o.ApplyT(func(v *PipelineTemplateResponse) PipelineTemplateResponse {
		if v != nil {
			return *v
		}
		var ret PipelineTemplateResponse
		return ret
	}).(PipelineTemplateResponseOutput)
}

func (o PipelineTemplateResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineTemplateResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o PipelineTemplateResponsePtrOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PipelineTemplateResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.StringMapOutput)
}

type ProjectReference struct {
	Name string `pulumi:"name"`
}





type ProjectReferenceInput interface {
	pulumi.Input

	ToProjectReferenceOutput() ProjectReferenceOutput
	ToProjectReferenceOutputWithContext(context.Context) ProjectReferenceOutput
}

type ProjectReferenceArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (ProjectReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectReference)(nil)).Elem()
}

func (i ProjectReferenceArgs) ToProjectReferenceOutput() ProjectReferenceOutput {
	return i.ToProjectReferenceOutputWithContext(context.Background())
}

func (i ProjectReferenceArgs) ToProjectReferenceOutputWithContext(ctx context.Context) ProjectReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectReferenceOutput)
}

func (i ProjectReferenceArgs) ToProjectReferencePtrOutput() ProjectReferencePtrOutput {
	return i.ToProjectReferencePtrOutputWithContext(context.Background())
}

func (i ProjectReferenceArgs) ToProjectReferencePtrOutputWithContext(ctx context.Context) ProjectReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectReferenceOutput).ToProjectReferencePtrOutputWithContext(ctx)
}









type ProjectReferencePtrInput interface {
	pulumi.Input

	ToProjectReferencePtrOutput() ProjectReferencePtrOutput
	ToProjectReferencePtrOutputWithContext(context.Context) ProjectReferencePtrOutput
}

type projectReferencePtrType ProjectReferenceArgs

func ProjectReferencePtr(v *ProjectReferenceArgs) ProjectReferencePtrInput {
	return (*projectReferencePtrType)(v)
}

func (*projectReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectReference)(nil)).Elem()
}

func (i *projectReferencePtrType) ToProjectReferencePtrOutput() ProjectReferencePtrOutput {
	return i.ToProjectReferencePtrOutputWithContext(context.Background())
}

func (i *projectReferencePtrType) ToProjectReferencePtrOutputWithContext(ctx context.Context) ProjectReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectReferencePtrOutput)
}

type ProjectReferenceOutput struct{ *pulumi.OutputState }

func (ProjectReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectReference)(nil)).Elem()
}

func (o ProjectReferenceOutput) ToProjectReferenceOutput() ProjectReferenceOutput {
	return o
}

func (o ProjectReferenceOutput) ToProjectReferenceOutputWithContext(ctx context.Context) ProjectReferenceOutput {
	return o
}

func (o ProjectReferenceOutput) ToProjectReferencePtrOutput() ProjectReferencePtrOutput {
	return o.ToProjectReferencePtrOutputWithContext(context.Background())
}

func (o ProjectReferenceOutput) ToProjectReferencePtrOutputWithContext(ctx context.Context) ProjectReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProjectReference) *ProjectReference {
		return &v
	}).(ProjectReferencePtrOutput)
}

func (o ProjectReferenceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ProjectReference) string { return v.Name }).(pulumi.StringOutput)
}

type ProjectReferencePtrOutput struct{ *pulumi.OutputState }

func (ProjectReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectReference)(nil)).Elem()
}

func (o ProjectReferencePtrOutput) ToProjectReferencePtrOutput() ProjectReferencePtrOutput {
	return o
}

func (o ProjectReferencePtrOutput) ToProjectReferencePtrOutputWithContext(ctx context.Context) ProjectReferencePtrOutput {
	return o
}

func (o ProjectReferencePtrOutput) Elem() ProjectReferenceOutput {
	return o.ApplyT(func(v *ProjectReference) ProjectReference {
		if v != nil {
			return *v
		}
		var ret ProjectReference
		return ret
	}).(ProjectReferenceOutput)
}

func (o ProjectReferencePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectReference) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type ProjectReferenceResponse struct {
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}





type ProjectReferenceResponseInput interface {
	pulumi.Input

	ToProjectReferenceResponseOutput() ProjectReferenceResponseOutput
	ToProjectReferenceResponseOutputWithContext(context.Context) ProjectReferenceResponseOutput
}

type ProjectReferenceResponseArgs struct {
	Id   pulumi.StringInput `pulumi:"id"`
	Name pulumi.StringInput `pulumi:"name"`
}

func (ProjectReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectReferenceResponse)(nil)).Elem()
}

func (i ProjectReferenceResponseArgs) ToProjectReferenceResponseOutput() ProjectReferenceResponseOutput {
	return i.ToProjectReferenceResponseOutputWithContext(context.Background())
}

func (i ProjectReferenceResponseArgs) ToProjectReferenceResponseOutputWithContext(ctx context.Context) ProjectReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectReferenceResponseOutput)
}

func (i ProjectReferenceResponseArgs) ToProjectReferenceResponsePtrOutput() ProjectReferenceResponsePtrOutput {
	return i.ToProjectReferenceResponsePtrOutputWithContext(context.Background())
}

func (i ProjectReferenceResponseArgs) ToProjectReferenceResponsePtrOutputWithContext(ctx context.Context) ProjectReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectReferenceResponseOutput).ToProjectReferenceResponsePtrOutputWithContext(ctx)
}









type ProjectReferenceResponsePtrInput interface {
	pulumi.Input

	ToProjectReferenceResponsePtrOutput() ProjectReferenceResponsePtrOutput
	ToProjectReferenceResponsePtrOutputWithContext(context.Context) ProjectReferenceResponsePtrOutput
}

type projectReferenceResponsePtrType ProjectReferenceResponseArgs

func ProjectReferenceResponsePtr(v *ProjectReferenceResponseArgs) ProjectReferenceResponsePtrInput {
	return (*projectReferenceResponsePtrType)(v)
}

func (*projectReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectReferenceResponse)(nil)).Elem()
}

func (i *projectReferenceResponsePtrType) ToProjectReferenceResponsePtrOutput() ProjectReferenceResponsePtrOutput {
	return i.ToProjectReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *projectReferenceResponsePtrType) ToProjectReferenceResponsePtrOutputWithContext(ctx context.Context) ProjectReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectReferenceResponsePtrOutput)
}

type ProjectReferenceResponseOutput struct{ *pulumi.OutputState }

func (ProjectReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectReferenceResponse)(nil)).Elem()
}

func (o ProjectReferenceResponseOutput) ToProjectReferenceResponseOutput() ProjectReferenceResponseOutput {
	return o
}

func (o ProjectReferenceResponseOutput) ToProjectReferenceResponseOutputWithContext(ctx context.Context) ProjectReferenceResponseOutput {
	return o
}

func (o ProjectReferenceResponseOutput) ToProjectReferenceResponsePtrOutput() ProjectReferenceResponsePtrOutput {
	return o.ToProjectReferenceResponsePtrOutputWithContext(context.Background())
}

func (o ProjectReferenceResponseOutput) ToProjectReferenceResponsePtrOutputWithContext(ctx context.Context) ProjectReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProjectReferenceResponse) *ProjectReferenceResponse {
		return &v
	}).(ProjectReferenceResponsePtrOutput)
}

func (o ProjectReferenceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ProjectReferenceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ProjectReferenceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ProjectReferenceResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ProjectReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ProjectReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectReferenceResponse)(nil)).Elem()
}

func (o ProjectReferenceResponsePtrOutput) ToProjectReferenceResponsePtrOutput() ProjectReferenceResponsePtrOutput {
	return o
}

func (o ProjectReferenceResponsePtrOutput) ToProjectReferenceResponsePtrOutputWithContext(ctx context.Context) ProjectReferenceResponsePtrOutput {
	return o
}

func (o ProjectReferenceResponsePtrOutput) Elem() ProjectReferenceResponseOutput {
	return o.ApplyT(func(v *ProjectReferenceResponse) ProjectReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ProjectReferenceResponse
		return ret
	}).(ProjectReferenceResponseOutput)
}

func (o ProjectReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ProjectReferenceResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorizationOutput{})
	pulumi.RegisterOutputType(AuthorizationPtrOutput{})
	pulumi.RegisterOutputType(AuthorizationResponseOutput{})
	pulumi.RegisterOutputType(AuthorizationResponsePtrOutput{})
	pulumi.RegisterOutputType(BootstrapConfigurationOutput{})
	pulumi.RegisterOutputType(BootstrapConfigurationPtrOutput{})
	pulumi.RegisterOutputType(BootstrapConfigurationResponseOutput{})
	pulumi.RegisterOutputType(BootstrapConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(CodeRepositoryOutput{})
	pulumi.RegisterOutputType(CodeRepositoryPtrOutput{})
	pulumi.RegisterOutputType(CodeRepositoryResponseOutput{})
	pulumi.RegisterOutputType(CodeRepositoryResponsePtrOutput{})
	pulumi.RegisterOutputType(OrganizationReferenceOutput{})
	pulumi.RegisterOutputType(OrganizationReferencePtrOutput{})
	pulumi.RegisterOutputType(OrganizationReferenceResponseOutput{})
	pulumi.RegisterOutputType(OrganizationReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineTemplateOutput{})
	pulumi.RegisterOutputType(PipelineTemplatePtrOutput{})
	pulumi.RegisterOutputType(PipelineTemplateResponseOutput{})
	pulumi.RegisterOutputType(PipelineTemplateResponsePtrOutput{})
	pulumi.RegisterOutputType(ProjectReferenceOutput{})
	pulumi.RegisterOutputType(ProjectReferencePtrOutput{})
	pulumi.RegisterOutputType(ProjectReferenceResponseOutput{})
	pulumi.RegisterOutputType(ProjectReferenceResponsePtrOutput{})
}
