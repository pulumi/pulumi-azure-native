


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

func (o BootstrapConfigurationOutput) Repository() CodeRepositoryPtrOutput {
	return o.ApplyT(func(v BootstrapConfiguration) *CodeRepository { return v.Repository }).(CodeRepositoryPtrOutput)
}

func (o BootstrapConfigurationOutput) Template() PipelineTemplateOutput {
	return o.ApplyT(func(v BootstrapConfiguration) PipelineTemplate { return v.Template }).(PipelineTemplateOutput)
}

type BootstrapConfigurationResponse struct {
	Repository *CodeRepositoryResponse  `pulumi:"repository"`
	Template   PipelineTemplateResponse `pulumi:"template"`
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

func (o BootstrapConfigurationResponseOutput) Repository() CodeRepositoryResponsePtrOutput {
	return o.ApplyT(func(v BootstrapConfigurationResponse) *CodeRepositoryResponse { return v.Repository }).(CodeRepositoryResponsePtrOutput)
}

func (o BootstrapConfigurationResponseOutput) Template() PipelineTemplateResponseOutput {
	return o.ApplyT(func(v BootstrapConfigurationResponse) PipelineTemplateResponse { return v.Template }).(PipelineTemplateResponseOutput)
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

func (o OrganizationReferenceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v OrganizationReference) string { return v.Name }).(pulumi.StringOutput)
}

type OrganizationReferenceResponse struct {
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
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

func (o OrganizationReferenceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v OrganizationReferenceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o OrganizationReferenceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v OrganizationReferenceResponse) string { return v.Name }).(pulumi.StringOutput)
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

func (o PipelineTemplateOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PipelineTemplate) string { return v.Id }).(pulumi.StringOutput)
}

func (o PipelineTemplateOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v PipelineTemplate) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

type PipelineTemplateResponse struct {
	Id         string            `pulumi:"id"`
	Parameters map[string]string `pulumi:"parameters"`
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

func (o PipelineTemplateResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PipelineTemplateResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PipelineTemplateResponseOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v PipelineTemplateResponse) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
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

func (o ProjectReferenceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ProjectReference) string { return v.Name }).(pulumi.StringOutput)
}

type ProjectReferenceResponse struct {
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
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

func (o ProjectReferenceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ProjectReferenceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ProjectReferenceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ProjectReferenceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorizationOutput{})
	pulumi.RegisterOutputType(AuthorizationPtrOutput{})
	pulumi.RegisterOutputType(AuthorizationResponseOutput{})
	pulumi.RegisterOutputType(AuthorizationResponsePtrOutput{})
	pulumi.RegisterOutputType(BootstrapConfigurationOutput{})
	pulumi.RegisterOutputType(BootstrapConfigurationResponseOutput{})
	pulumi.RegisterOutputType(CodeRepositoryOutput{})
	pulumi.RegisterOutputType(CodeRepositoryPtrOutput{})
	pulumi.RegisterOutputType(CodeRepositoryResponseOutput{})
	pulumi.RegisterOutputType(CodeRepositoryResponsePtrOutput{})
	pulumi.RegisterOutputType(OrganizationReferenceOutput{})
	pulumi.RegisterOutputType(OrganizationReferenceResponseOutput{})
	pulumi.RegisterOutputType(PipelineTemplateOutput{})
	pulumi.RegisterOutputType(PipelineTemplateResponseOutput{})
	pulumi.RegisterOutputType(ProjectReferenceOutput{})
	pulumi.RegisterOutputType(ProjectReferenceResponseOutput{})
}
