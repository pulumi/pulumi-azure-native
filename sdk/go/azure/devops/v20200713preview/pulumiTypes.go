


package v20200713preview

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
	SourceRepository *CodeRepository  `pulumi:"sourceRepository"`
	Template         PipelineTemplate `pulumi:"template"`
}





type BootstrapConfigurationInput interface {
	pulumi.Input

	ToBootstrapConfigurationOutput() BootstrapConfigurationOutput
	ToBootstrapConfigurationOutputWithContext(context.Context) BootstrapConfigurationOutput
}

type BootstrapConfigurationArgs struct {
	SourceRepository CodeRepositoryPtrInput `pulumi:"sourceRepository"`
	Template         PipelineTemplateInput  `pulumi:"template"`
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

func (o BootstrapConfigurationOutput) SourceRepository() CodeRepositoryPtrOutput {
	return o.ApplyT(func(v BootstrapConfiguration) *CodeRepository { return v.SourceRepository }).(CodeRepositoryPtrOutput)
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

func (o BootstrapConfigurationPtrOutput) SourceRepository() CodeRepositoryPtrOutput {
	return o.ApplyT(func(v *BootstrapConfiguration) *CodeRepository {
		if v == nil {
			return nil
		}
		return v.SourceRepository
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
	SourceRepository *CodeRepositoryResponse  `pulumi:"sourceRepository"`
	Template         PipelineTemplateResponse `pulumi:"template"`
}





type BootstrapConfigurationResponseInput interface {
	pulumi.Input

	ToBootstrapConfigurationResponseOutput() BootstrapConfigurationResponseOutput
	ToBootstrapConfigurationResponseOutputWithContext(context.Context) BootstrapConfigurationResponseOutput
}

type BootstrapConfigurationResponseArgs struct {
	SourceRepository CodeRepositoryResponsePtrInput `pulumi:"sourceRepository"`
	Template         PipelineTemplateResponseInput  `pulumi:"template"`
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

func (o BootstrapConfigurationResponseOutput) SourceRepository() CodeRepositoryResponsePtrOutput {
	return o.ApplyT(func(v BootstrapConfigurationResponse) *CodeRepositoryResponse { return v.SourceRepository }).(CodeRepositoryResponsePtrOutput)
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

func (o BootstrapConfigurationResponsePtrOutput) SourceRepository() CodeRepositoryResponsePtrOutput {
	return o.ApplyT(func(v *BootstrapConfigurationResponse) *CodeRepositoryResponse {
		if v == nil {
			return nil
		}
		return v.SourceRepository
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

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthorizationInput)(nil)).Elem(), AuthorizationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthorizationPtrInput)(nil)).Elem(), AuthorizationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthorizationResponseInput)(nil)).Elem(), AuthorizationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthorizationResponsePtrInput)(nil)).Elem(), AuthorizationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BootstrapConfigurationInput)(nil)).Elem(), BootstrapConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BootstrapConfigurationPtrInput)(nil)).Elem(), BootstrapConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BootstrapConfigurationResponseInput)(nil)).Elem(), BootstrapConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BootstrapConfigurationResponsePtrInput)(nil)).Elem(), BootstrapConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CodeRepositoryInput)(nil)).Elem(), CodeRepositoryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CodeRepositoryPtrInput)(nil)).Elem(), CodeRepositoryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CodeRepositoryResponseInput)(nil)).Elem(), CodeRepositoryResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CodeRepositoryResponsePtrInput)(nil)).Elem(), CodeRepositoryResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineTemplateInput)(nil)).Elem(), PipelineTemplateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineTemplatePtrInput)(nil)).Elem(), PipelineTemplateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineTemplateResponseInput)(nil)).Elem(), PipelineTemplateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineTemplateResponsePtrInput)(nil)).Elem(), PipelineTemplateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
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
	pulumi.RegisterOutputType(PipelineTemplateOutput{})
	pulumi.RegisterOutputType(PipelineTemplatePtrOutput{})
	pulumi.RegisterOutputType(PipelineTemplateResponseOutput{})
	pulumi.RegisterOutputType(PipelineTemplateResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
