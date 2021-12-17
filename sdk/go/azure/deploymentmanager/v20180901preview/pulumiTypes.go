


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudErrorBodyResponse struct {
	Code    string                   `pulumi:"code"`
	Details []CloudErrorBodyResponse `pulumi:"details"`
	Message string                   `pulumi:"message"`
	Target  *string                  `pulumi:"target"`
}

type Identity struct {
	IdentityIds []string `pulumi:"identityIds"`
	Type        string   `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	IdentityIds pulumi.StringArrayInput `pulumi:"identityIds"`
	Type        pulumi.StringInput      `pulumi:"type"`
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

func (o IdentityOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Identity) []string { return v.IdentityIds }).(pulumi.StringArrayOutput)
}

func (o IdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v Identity) string { return v.Type }).(pulumi.StringOutput)
}

type IdentityResponse struct {
	IdentityIds []string `pulumi:"identityIds"`
	Type        string   `pulumi:"type"`
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

func (o IdentityResponseOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IdentityResponse) []string { return v.IdentityIds }).(pulumi.StringArrayOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type MessageResponse struct {
	Message   string `pulumi:"message"`
	TimeStamp string `pulumi:"timeStamp"`
}

type PrePostStep struct {
	StepId string `pulumi:"stepId"`
}





type PrePostStepInput interface {
	pulumi.Input

	ToPrePostStepOutput() PrePostStepOutput
	ToPrePostStepOutputWithContext(context.Context) PrePostStepOutput
}

type PrePostStepArgs struct {
	StepId pulumi.StringInput `pulumi:"stepId"`
}

func (PrePostStepArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrePostStep)(nil)).Elem()
}

func (i PrePostStepArgs) ToPrePostStepOutput() PrePostStepOutput {
	return i.ToPrePostStepOutputWithContext(context.Background())
}

func (i PrePostStepArgs) ToPrePostStepOutputWithContext(ctx context.Context) PrePostStepOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrePostStepOutput)
}





type PrePostStepArrayInput interface {
	pulumi.Input

	ToPrePostStepArrayOutput() PrePostStepArrayOutput
	ToPrePostStepArrayOutputWithContext(context.Context) PrePostStepArrayOutput
}

type PrePostStepArray []PrePostStepInput

func (PrePostStepArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrePostStep)(nil)).Elem()
}

func (i PrePostStepArray) ToPrePostStepArrayOutput() PrePostStepArrayOutput {
	return i.ToPrePostStepArrayOutputWithContext(context.Background())
}

func (i PrePostStepArray) ToPrePostStepArrayOutputWithContext(ctx context.Context) PrePostStepArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrePostStepArrayOutput)
}

type PrePostStepOutput struct{ *pulumi.OutputState }

func (PrePostStepOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrePostStep)(nil)).Elem()
}

func (o PrePostStepOutput) ToPrePostStepOutput() PrePostStepOutput {
	return o
}

func (o PrePostStepOutput) ToPrePostStepOutputWithContext(ctx context.Context) PrePostStepOutput {
	return o
}

func (o PrePostStepOutput) StepId() pulumi.StringOutput {
	return o.ApplyT(func(v PrePostStep) string { return v.StepId }).(pulumi.StringOutput)
}

type PrePostStepArrayOutput struct{ *pulumi.OutputState }

func (PrePostStepArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrePostStep)(nil)).Elem()
}

func (o PrePostStepArrayOutput) ToPrePostStepArrayOutput() PrePostStepArrayOutput {
	return o
}

func (o PrePostStepArrayOutput) ToPrePostStepArrayOutputWithContext(ctx context.Context) PrePostStepArrayOutput {
	return o
}

func (o PrePostStepArrayOutput) Index(i pulumi.IntInput) PrePostStepOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrePostStep {
		return vs[0].([]PrePostStep)[vs[1].(int)]
	}).(PrePostStepOutput)
}

type PrePostStepResponse struct {
	StepId string `pulumi:"stepId"`
}

type PrePostStepResponseOutput struct{ *pulumi.OutputState }

func (PrePostStepResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrePostStepResponse)(nil)).Elem()
}

func (o PrePostStepResponseOutput) ToPrePostStepResponseOutput() PrePostStepResponseOutput {
	return o
}

func (o PrePostStepResponseOutput) ToPrePostStepResponseOutputWithContext(ctx context.Context) PrePostStepResponseOutput {
	return o
}

func (o PrePostStepResponseOutput) StepId() pulumi.StringOutput {
	return o.ApplyT(func(v PrePostStepResponse) string { return v.StepId }).(pulumi.StringOutput)
}

type PrePostStepResponseArrayOutput struct{ *pulumi.OutputState }

func (PrePostStepResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrePostStepResponse)(nil)).Elem()
}

func (o PrePostStepResponseArrayOutput) ToPrePostStepResponseArrayOutput() PrePostStepResponseArrayOutput {
	return o
}

func (o PrePostStepResponseArrayOutput) ToPrePostStepResponseArrayOutputWithContext(ctx context.Context) PrePostStepResponseArrayOutput {
	return o
}

func (o PrePostStepResponseArrayOutput) Index(i pulumi.IntInput) PrePostStepResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrePostStepResponse {
		return vs[0].([]PrePostStepResponse)[vs[1].(int)]
	}).(PrePostStepResponseOutput)
}

type ResourceOperationResponse struct {
	OperationId       string  `pulumi:"operationId"`
	ProvisioningState string  `pulumi:"provisioningState"`
	ResourceName      *string `pulumi:"resourceName"`
	ResourceType      *string `pulumi:"resourceType"`
	StatusCode        string  `pulumi:"statusCode"`
	StatusMessage     string  `pulumi:"statusMessage"`
}

type RolloutOperationInfoResponse struct {
	EndTime              string                 `pulumi:"endTime"`
	Error                CloudErrorBodyResponse `pulumi:"error"`
	RetryAttempt         int                    `pulumi:"retryAttempt"`
	SkipSucceededOnRetry bool                   `pulumi:"skipSucceededOnRetry"`
	StartTime            string                 `pulumi:"startTime"`
}

type RolloutStepResponse struct {
	Messages           []MessageResponse           `pulumi:"messages"`
	Name               string                      `pulumi:"name"`
	OperationInfo      StepOperationInfoResponse   `pulumi:"operationInfo"`
	ResourceOperations []ResourceOperationResponse `pulumi:"resourceOperations"`
	Status             string                      `pulumi:"status"`
	StepGroup          *string                     `pulumi:"stepGroup"`
}

type SasAuthentication struct {
	SasUri string `pulumi:"sasUri"`
	Type   string `pulumi:"type"`
}





type SasAuthenticationInput interface {
	pulumi.Input

	ToSasAuthenticationOutput() SasAuthenticationOutput
	ToSasAuthenticationOutputWithContext(context.Context) SasAuthenticationOutput
}

type SasAuthenticationArgs struct {
	SasUri pulumi.StringInput `pulumi:"sasUri"`
	Type   pulumi.StringInput `pulumi:"type"`
}

func (SasAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SasAuthentication)(nil)).Elem()
}

func (i SasAuthenticationArgs) ToSasAuthenticationOutput() SasAuthenticationOutput {
	return i.ToSasAuthenticationOutputWithContext(context.Background())
}

func (i SasAuthenticationArgs) ToSasAuthenticationOutputWithContext(ctx context.Context) SasAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SasAuthenticationOutput)
}

type SasAuthenticationOutput struct{ *pulumi.OutputState }

func (SasAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SasAuthentication)(nil)).Elem()
}

func (o SasAuthenticationOutput) ToSasAuthenticationOutput() SasAuthenticationOutput {
	return o
}

func (o SasAuthenticationOutput) ToSasAuthenticationOutputWithContext(ctx context.Context) SasAuthenticationOutput {
	return o
}

func (o SasAuthenticationOutput) SasUri() pulumi.StringOutput {
	return o.ApplyT(func(v SasAuthentication) string { return v.SasUri }).(pulumi.StringOutput)
}

func (o SasAuthenticationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SasAuthentication) string { return v.Type }).(pulumi.StringOutput)
}

type SasAuthenticationResponse struct {
	SasUri string `pulumi:"sasUri"`
	Type   string `pulumi:"type"`
}

type SasAuthenticationResponseOutput struct{ *pulumi.OutputState }

func (SasAuthenticationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SasAuthenticationResponse)(nil)).Elem()
}

func (o SasAuthenticationResponseOutput) ToSasAuthenticationResponseOutput() SasAuthenticationResponseOutput {
	return o
}

func (o SasAuthenticationResponseOutput) ToSasAuthenticationResponseOutputWithContext(ctx context.Context) SasAuthenticationResponseOutput {
	return o
}

func (o SasAuthenticationResponseOutput) SasUri() pulumi.StringOutput {
	return o.ApplyT(func(v SasAuthenticationResponse) string { return v.SasUri }).(pulumi.StringOutput)
}

func (o SasAuthenticationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SasAuthenticationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ServiceResponse struct {
	Name                 *string               `pulumi:"name"`
	ServiceUnits         []ServiceUnitResponse `pulumi:"serviceUnits"`
	TargetLocation       string                `pulumi:"targetLocation"`
	TargetSubscriptionId string                `pulumi:"targetSubscriptionId"`
}

type ServiceUnitArtifacts struct {
	ParametersArtifactSourceRelativePath *string `pulumi:"parametersArtifactSourceRelativePath"`
	ParametersUri                        *string `pulumi:"parametersUri"`
	TemplateArtifactSourceRelativePath   *string `pulumi:"templateArtifactSourceRelativePath"`
	TemplateUri                          *string `pulumi:"templateUri"`
}





type ServiceUnitArtifactsInput interface {
	pulumi.Input

	ToServiceUnitArtifactsOutput() ServiceUnitArtifactsOutput
	ToServiceUnitArtifactsOutputWithContext(context.Context) ServiceUnitArtifactsOutput
}

type ServiceUnitArtifactsArgs struct {
	ParametersArtifactSourceRelativePath pulumi.StringPtrInput `pulumi:"parametersArtifactSourceRelativePath"`
	ParametersUri                        pulumi.StringPtrInput `pulumi:"parametersUri"`
	TemplateArtifactSourceRelativePath   pulumi.StringPtrInput `pulumi:"templateArtifactSourceRelativePath"`
	TemplateUri                          pulumi.StringPtrInput `pulumi:"templateUri"`
}

func (ServiceUnitArtifactsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceUnitArtifacts)(nil)).Elem()
}

func (i ServiceUnitArtifactsArgs) ToServiceUnitArtifactsOutput() ServiceUnitArtifactsOutput {
	return i.ToServiceUnitArtifactsOutputWithContext(context.Background())
}

func (i ServiceUnitArtifactsArgs) ToServiceUnitArtifactsOutputWithContext(ctx context.Context) ServiceUnitArtifactsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceUnitArtifactsOutput)
}

func (i ServiceUnitArtifactsArgs) ToServiceUnitArtifactsPtrOutput() ServiceUnitArtifactsPtrOutput {
	return i.ToServiceUnitArtifactsPtrOutputWithContext(context.Background())
}

func (i ServiceUnitArtifactsArgs) ToServiceUnitArtifactsPtrOutputWithContext(ctx context.Context) ServiceUnitArtifactsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceUnitArtifactsOutput).ToServiceUnitArtifactsPtrOutputWithContext(ctx)
}









type ServiceUnitArtifactsPtrInput interface {
	pulumi.Input

	ToServiceUnitArtifactsPtrOutput() ServiceUnitArtifactsPtrOutput
	ToServiceUnitArtifactsPtrOutputWithContext(context.Context) ServiceUnitArtifactsPtrOutput
}

type serviceUnitArtifactsPtrType ServiceUnitArtifactsArgs

func ServiceUnitArtifactsPtr(v *ServiceUnitArtifactsArgs) ServiceUnitArtifactsPtrInput {
	return (*serviceUnitArtifactsPtrType)(v)
}

func (*serviceUnitArtifactsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceUnitArtifacts)(nil)).Elem()
}

func (i *serviceUnitArtifactsPtrType) ToServiceUnitArtifactsPtrOutput() ServiceUnitArtifactsPtrOutput {
	return i.ToServiceUnitArtifactsPtrOutputWithContext(context.Background())
}

func (i *serviceUnitArtifactsPtrType) ToServiceUnitArtifactsPtrOutputWithContext(ctx context.Context) ServiceUnitArtifactsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceUnitArtifactsPtrOutput)
}

type ServiceUnitArtifactsOutput struct{ *pulumi.OutputState }

func (ServiceUnitArtifactsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceUnitArtifacts)(nil)).Elem()
}

func (o ServiceUnitArtifactsOutput) ToServiceUnitArtifactsOutput() ServiceUnitArtifactsOutput {
	return o
}

func (o ServiceUnitArtifactsOutput) ToServiceUnitArtifactsOutputWithContext(ctx context.Context) ServiceUnitArtifactsOutput {
	return o
}

func (o ServiceUnitArtifactsOutput) ToServiceUnitArtifactsPtrOutput() ServiceUnitArtifactsPtrOutput {
	return o.ToServiceUnitArtifactsPtrOutputWithContext(context.Background())
}

func (o ServiceUnitArtifactsOutput) ToServiceUnitArtifactsPtrOutputWithContext(ctx context.Context) ServiceUnitArtifactsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceUnitArtifacts) *ServiceUnitArtifacts {
		return &v
	}).(ServiceUnitArtifactsPtrOutput)
}

func (o ServiceUnitArtifactsOutput) ParametersArtifactSourceRelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceUnitArtifacts) *string { return v.ParametersArtifactSourceRelativePath }).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsOutput) ParametersUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceUnitArtifacts) *string { return v.ParametersUri }).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsOutput) TemplateArtifactSourceRelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceUnitArtifacts) *string { return v.TemplateArtifactSourceRelativePath }).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsOutput) TemplateUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceUnitArtifacts) *string { return v.TemplateUri }).(pulumi.StringPtrOutput)
}

type ServiceUnitArtifactsPtrOutput struct{ *pulumi.OutputState }

func (ServiceUnitArtifactsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceUnitArtifacts)(nil)).Elem()
}

func (o ServiceUnitArtifactsPtrOutput) ToServiceUnitArtifactsPtrOutput() ServiceUnitArtifactsPtrOutput {
	return o
}

func (o ServiceUnitArtifactsPtrOutput) ToServiceUnitArtifactsPtrOutputWithContext(ctx context.Context) ServiceUnitArtifactsPtrOutput {
	return o
}

func (o ServiceUnitArtifactsPtrOutput) Elem() ServiceUnitArtifactsOutput {
	return o.ApplyT(func(v *ServiceUnitArtifacts) ServiceUnitArtifacts {
		if v != nil {
			return *v
		}
		var ret ServiceUnitArtifacts
		return ret
	}).(ServiceUnitArtifactsOutput)
}

func (o ServiceUnitArtifactsPtrOutput) ParametersArtifactSourceRelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceUnitArtifacts) *string {
		if v == nil {
			return nil
		}
		return v.ParametersArtifactSourceRelativePath
	}).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsPtrOutput) ParametersUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceUnitArtifacts) *string {
		if v == nil {
			return nil
		}
		return v.ParametersUri
	}).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsPtrOutput) TemplateArtifactSourceRelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceUnitArtifacts) *string {
		if v == nil {
			return nil
		}
		return v.TemplateArtifactSourceRelativePath
	}).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsPtrOutput) TemplateUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceUnitArtifacts) *string {
		if v == nil {
			return nil
		}
		return v.TemplateUri
	}).(pulumi.StringPtrOutput)
}

type ServiceUnitArtifactsResponse struct {
	ParametersArtifactSourceRelativePath *string `pulumi:"parametersArtifactSourceRelativePath"`
	ParametersUri                        *string `pulumi:"parametersUri"`
	TemplateArtifactSourceRelativePath   *string `pulumi:"templateArtifactSourceRelativePath"`
	TemplateUri                          *string `pulumi:"templateUri"`
}

type ServiceUnitArtifactsResponseOutput struct{ *pulumi.OutputState }

func (ServiceUnitArtifactsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceUnitArtifactsResponse)(nil)).Elem()
}

func (o ServiceUnitArtifactsResponseOutput) ToServiceUnitArtifactsResponseOutput() ServiceUnitArtifactsResponseOutput {
	return o
}

func (o ServiceUnitArtifactsResponseOutput) ToServiceUnitArtifactsResponseOutputWithContext(ctx context.Context) ServiceUnitArtifactsResponseOutput {
	return o
}

func (o ServiceUnitArtifactsResponseOutput) ParametersArtifactSourceRelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceUnitArtifactsResponse) *string { return v.ParametersArtifactSourceRelativePath }).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsResponseOutput) ParametersUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceUnitArtifactsResponse) *string { return v.ParametersUri }).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsResponseOutput) TemplateArtifactSourceRelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceUnitArtifactsResponse) *string { return v.TemplateArtifactSourceRelativePath }).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsResponseOutput) TemplateUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceUnitArtifactsResponse) *string { return v.TemplateUri }).(pulumi.StringPtrOutput)
}

type ServiceUnitArtifactsResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceUnitArtifactsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceUnitArtifactsResponse)(nil)).Elem()
}

func (o ServiceUnitArtifactsResponsePtrOutput) ToServiceUnitArtifactsResponsePtrOutput() ServiceUnitArtifactsResponsePtrOutput {
	return o
}

func (o ServiceUnitArtifactsResponsePtrOutput) ToServiceUnitArtifactsResponsePtrOutputWithContext(ctx context.Context) ServiceUnitArtifactsResponsePtrOutput {
	return o
}

func (o ServiceUnitArtifactsResponsePtrOutput) Elem() ServiceUnitArtifactsResponseOutput {
	return o.ApplyT(func(v *ServiceUnitArtifactsResponse) ServiceUnitArtifactsResponse {
		if v != nil {
			return *v
		}
		var ret ServiceUnitArtifactsResponse
		return ret
	}).(ServiceUnitArtifactsResponseOutput)
}

func (o ServiceUnitArtifactsResponsePtrOutput) ParametersArtifactSourceRelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceUnitArtifactsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ParametersArtifactSourceRelativePath
	}).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsResponsePtrOutput) ParametersUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceUnitArtifactsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ParametersUri
	}).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsResponsePtrOutput) TemplateArtifactSourceRelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceUnitArtifactsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TemplateArtifactSourceRelativePath
	}).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsResponsePtrOutput) TemplateUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceUnitArtifactsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TemplateUri
	}).(pulumi.StringPtrOutput)
}

type ServiceUnitResponse struct {
	Artifacts           *ServiceUnitArtifactsResponse `pulumi:"artifacts"`
	DeploymentMode      string                        `pulumi:"deploymentMode"`
	Name                *string                       `pulumi:"name"`
	Steps               []RolloutStepResponse         `pulumi:"steps"`
	TargetResourceGroup string                        `pulumi:"targetResourceGroup"`
}

type StepType struct {
	DependsOnStepGroups []string      `pulumi:"dependsOnStepGroups"`
	DeploymentTargetId  string        `pulumi:"deploymentTargetId"`
	Name                string        `pulumi:"name"`
	PostDeploymentSteps []PrePostStep `pulumi:"postDeploymentSteps"`
	PreDeploymentSteps  []PrePostStep `pulumi:"preDeploymentSteps"`
}





type StepTypeInput interface {
	pulumi.Input

	ToStepTypeOutput() StepTypeOutput
	ToStepTypeOutputWithContext(context.Context) StepTypeOutput
}

type StepTypeArgs struct {
	DependsOnStepGroups pulumi.StringArrayInput `pulumi:"dependsOnStepGroups"`
	DeploymentTargetId  pulumi.StringInput      `pulumi:"deploymentTargetId"`
	Name                pulumi.StringInput      `pulumi:"name"`
	PostDeploymentSteps PrePostStepArrayInput   `pulumi:"postDeploymentSteps"`
	PreDeploymentSteps  PrePostStepArrayInput   `pulumi:"preDeploymentSteps"`
}

func (StepTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StepType)(nil)).Elem()
}

func (i StepTypeArgs) ToStepTypeOutput() StepTypeOutput {
	return i.ToStepTypeOutputWithContext(context.Background())
}

func (i StepTypeArgs) ToStepTypeOutputWithContext(ctx context.Context) StepTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StepTypeOutput)
}





type StepTypeArrayInput interface {
	pulumi.Input

	ToStepTypeArrayOutput() StepTypeArrayOutput
	ToStepTypeArrayOutputWithContext(context.Context) StepTypeArrayOutput
}

type StepTypeArray []StepTypeInput

func (StepTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StepType)(nil)).Elem()
}

func (i StepTypeArray) ToStepTypeArrayOutput() StepTypeArrayOutput {
	return i.ToStepTypeArrayOutputWithContext(context.Background())
}

func (i StepTypeArray) ToStepTypeArrayOutputWithContext(ctx context.Context) StepTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StepTypeArrayOutput)
}

type StepTypeOutput struct{ *pulumi.OutputState }

func (StepTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StepType)(nil)).Elem()
}

func (o StepTypeOutput) ToStepTypeOutput() StepTypeOutput {
	return o
}

func (o StepTypeOutput) ToStepTypeOutputWithContext(ctx context.Context) StepTypeOutput {
	return o
}

func (o StepTypeOutput) DependsOnStepGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StepType) []string { return v.DependsOnStepGroups }).(pulumi.StringArrayOutput)
}

func (o StepTypeOutput) DeploymentTargetId() pulumi.StringOutput {
	return o.ApplyT(func(v StepType) string { return v.DeploymentTargetId }).(pulumi.StringOutput)
}

func (o StepTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v StepType) string { return v.Name }).(pulumi.StringOutput)
}

func (o StepTypeOutput) PostDeploymentSteps() PrePostStepArrayOutput {
	return o.ApplyT(func(v StepType) []PrePostStep { return v.PostDeploymentSteps }).(PrePostStepArrayOutput)
}

func (o StepTypeOutput) PreDeploymentSteps() PrePostStepArrayOutput {
	return o.ApplyT(func(v StepType) []PrePostStep { return v.PreDeploymentSteps }).(PrePostStepArrayOutput)
}

type StepTypeArrayOutput struct{ *pulumi.OutputState }

func (StepTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StepType)(nil)).Elem()
}

func (o StepTypeArrayOutput) ToStepTypeArrayOutput() StepTypeArrayOutput {
	return o
}

func (o StepTypeArrayOutput) ToStepTypeArrayOutputWithContext(ctx context.Context) StepTypeArrayOutput {
	return o
}

func (o StepTypeArrayOutput) Index(i pulumi.IntInput) StepTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StepType {
		return vs[0].([]StepType)[vs[1].(int)]
	}).(StepTypeOutput)
}

type StepOperationInfoResponse struct {
	CorrelationId   string                  `pulumi:"correlationId"`
	DeploymentName  string                  `pulumi:"deploymentName"`
	EndTime         string                  `pulumi:"endTime"`
	Error           *CloudErrorBodyResponse `pulumi:"error"`
	LastUpdatedTime string                  `pulumi:"lastUpdatedTime"`
	StartTime       string                  `pulumi:"startTime"`
}

type StepResponse struct {
	DependsOnStepGroups []string              `pulumi:"dependsOnStepGroups"`
	DeploymentTargetId  string                `pulumi:"deploymentTargetId"`
	Name                string                `pulumi:"name"`
	PostDeploymentSteps []PrePostStepResponse `pulumi:"postDeploymentSteps"`
	PreDeploymentSteps  []PrePostStepResponse `pulumi:"preDeploymentSteps"`
}

type StepResponseOutput struct{ *pulumi.OutputState }

func (StepResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StepResponse)(nil)).Elem()
}

func (o StepResponseOutput) ToStepResponseOutput() StepResponseOutput {
	return o
}

func (o StepResponseOutput) ToStepResponseOutputWithContext(ctx context.Context) StepResponseOutput {
	return o
}

func (o StepResponseOutput) DependsOnStepGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StepResponse) []string { return v.DependsOnStepGroups }).(pulumi.StringArrayOutput)
}

func (o StepResponseOutput) DeploymentTargetId() pulumi.StringOutput {
	return o.ApplyT(func(v StepResponse) string { return v.DeploymentTargetId }).(pulumi.StringOutput)
}

func (o StepResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v StepResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o StepResponseOutput) PostDeploymentSteps() PrePostStepResponseArrayOutput {
	return o.ApplyT(func(v StepResponse) []PrePostStepResponse { return v.PostDeploymentSteps }).(PrePostStepResponseArrayOutput)
}

func (o StepResponseOutput) PreDeploymentSteps() PrePostStepResponseArrayOutput {
	return o.ApplyT(func(v StepResponse) []PrePostStepResponse { return v.PreDeploymentSteps }).(PrePostStepResponseArrayOutput)
}

type StepResponseArrayOutput struct{ *pulumi.OutputState }

func (StepResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StepResponse)(nil)).Elem()
}

func (o StepResponseArrayOutput) ToStepResponseArrayOutput() StepResponseArrayOutput {
	return o
}

func (o StepResponseArrayOutput) ToStepResponseArrayOutputWithContext(ctx context.Context) StepResponseArrayOutput {
	return o
}

func (o StepResponseArrayOutput) Index(i pulumi.IntInput) StepResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StepResponse {
		return vs[0].([]StepResponse)[vs[1].(int)]
	}).(StepResponseOutput)
}

type WaitStepAttributes struct {
	Duration string `pulumi:"duration"`
}





type WaitStepAttributesInput interface {
	pulumi.Input

	ToWaitStepAttributesOutput() WaitStepAttributesOutput
	ToWaitStepAttributesOutputWithContext(context.Context) WaitStepAttributesOutput
}

type WaitStepAttributesArgs struct {
	Duration pulumi.StringInput `pulumi:"duration"`
}

func (WaitStepAttributesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WaitStepAttributes)(nil)).Elem()
}

func (i WaitStepAttributesArgs) ToWaitStepAttributesOutput() WaitStepAttributesOutput {
	return i.ToWaitStepAttributesOutputWithContext(context.Background())
}

func (i WaitStepAttributesArgs) ToWaitStepAttributesOutputWithContext(ctx context.Context) WaitStepAttributesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WaitStepAttributesOutput)
}

func (i WaitStepAttributesArgs) ToWaitStepAttributesPtrOutput() WaitStepAttributesPtrOutput {
	return i.ToWaitStepAttributesPtrOutputWithContext(context.Background())
}

func (i WaitStepAttributesArgs) ToWaitStepAttributesPtrOutputWithContext(ctx context.Context) WaitStepAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WaitStepAttributesOutput).ToWaitStepAttributesPtrOutputWithContext(ctx)
}









type WaitStepAttributesPtrInput interface {
	pulumi.Input

	ToWaitStepAttributesPtrOutput() WaitStepAttributesPtrOutput
	ToWaitStepAttributesPtrOutputWithContext(context.Context) WaitStepAttributesPtrOutput
}

type waitStepAttributesPtrType WaitStepAttributesArgs

func WaitStepAttributesPtr(v *WaitStepAttributesArgs) WaitStepAttributesPtrInput {
	return (*waitStepAttributesPtrType)(v)
}

func (*waitStepAttributesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WaitStepAttributes)(nil)).Elem()
}

func (i *waitStepAttributesPtrType) ToWaitStepAttributesPtrOutput() WaitStepAttributesPtrOutput {
	return i.ToWaitStepAttributesPtrOutputWithContext(context.Background())
}

func (i *waitStepAttributesPtrType) ToWaitStepAttributesPtrOutputWithContext(ctx context.Context) WaitStepAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WaitStepAttributesPtrOutput)
}

type WaitStepAttributesOutput struct{ *pulumi.OutputState }

func (WaitStepAttributesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WaitStepAttributes)(nil)).Elem()
}

func (o WaitStepAttributesOutput) ToWaitStepAttributesOutput() WaitStepAttributesOutput {
	return o
}

func (o WaitStepAttributesOutput) ToWaitStepAttributesOutputWithContext(ctx context.Context) WaitStepAttributesOutput {
	return o
}

func (o WaitStepAttributesOutput) ToWaitStepAttributesPtrOutput() WaitStepAttributesPtrOutput {
	return o.ToWaitStepAttributesPtrOutputWithContext(context.Background())
}

func (o WaitStepAttributesOutput) ToWaitStepAttributesPtrOutputWithContext(ctx context.Context) WaitStepAttributesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WaitStepAttributes) *WaitStepAttributes {
		return &v
	}).(WaitStepAttributesPtrOutput)
}

func (o WaitStepAttributesOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v WaitStepAttributes) string { return v.Duration }).(pulumi.StringOutput)
}

type WaitStepAttributesPtrOutput struct{ *pulumi.OutputState }

func (WaitStepAttributesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WaitStepAttributes)(nil)).Elem()
}

func (o WaitStepAttributesPtrOutput) ToWaitStepAttributesPtrOutput() WaitStepAttributesPtrOutput {
	return o
}

func (o WaitStepAttributesPtrOutput) ToWaitStepAttributesPtrOutputWithContext(ctx context.Context) WaitStepAttributesPtrOutput {
	return o
}

func (o WaitStepAttributesPtrOutput) Elem() WaitStepAttributesOutput {
	return o.ApplyT(func(v *WaitStepAttributes) WaitStepAttributes {
		if v != nil {
			return *v
		}
		var ret WaitStepAttributes
		return ret
	}).(WaitStepAttributesOutput)
}

func (o WaitStepAttributesPtrOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WaitStepAttributes) *string {
		if v == nil {
			return nil
		}
		return &v.Duration
	}).(pulumi.StringPtrOutput)
}

type WaitStepAttributesResponse struct {
	Duration string `pulumi:"duration"`
}

type WaitStepAttributesResponseOutput struct{ *pulumi.OutputState }

func (WaitStepAttributesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WaitStepAttributesResponse)(nil)).Elem()
}

func (o WaitStepAttributesResponseOutput) ToWaitStepAttributesResponseOutput() WaitStepAttributesResponseOutput {
	return o
}

func (o WaitStepAttributesResponseOutput) ToWaitStepAttributesResponseOutputWithContext(ctx context.Context) WaitStepAttributesResponseOutput {
	return o
}

func (o WaitStepAttributesResponseOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v WaitStepAttributesResponse) string { return v.Duration }).(pulumi.StringOutput)
}

type WaitStepAttributesResponsePtrOutput struct{ *pulumi.OutputState }

func (WaitStepAttributesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WaitStepAttributesResponse)(nil)).Elem()
}

func (o WaitStepAttributesResponsePtrOutput) ToWaitStepAttributesResponsePtrOutput() WaitStepAttributesResponsePtrOutput {
	return o
}

func (o WaitStepAttributesResponsePtrOutput) ToWaitStepAttributesResponsePtrOutputWithContext(ctx context.Context) WaitStepAttributesResponsePtrOutput {
	return o
}

func (o WaitStepAttributesResponsePtrOutput) Elem() WaitStepAttributesResponseOutput {
	return o.ApplyT(func(v *WaitStepAttributesResponse) WaitStepAttributesResponse {
		if v != nil {
			return *v
		}
		var ret WaitStepAttributesResponse
		return ret
	}).(WaitStepAttributesResponseOutput)
}

func (o WaitStepAttributesResponsePtrOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WaitStepAttributesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Duration
	}).(pulumi.StringPtrOutput)
}

type WaitStepProperties struct {
	Attributes *WaitStepAttributes `pulumi:"attributes"`
	StepType   string              `pulumi:"stepType"`
}





type WaitStepPropertiesInput interface {
	pulumi.Input

	ToWaitStepPropertiesOutput() WaitStepPropertiesOutput
	ToWaitStepPropertiesOutputWithContext(context.Context) WaitStepPropertiesOutput
}

type WaitStepPropertiesArgs struct {
	Attributes WaitStepAttributesPtrInput `pulumi:"attributes"`
	StepType   pulumi.StringInput         `pulumi:"stepType"`
}

func (WaitStepPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WaitStepProperties)(nil)).Elem()
}

func (i WaitStepPropertiesArgs) ToWaitStepPropertiesOutput() WaitStepPropertiesOutput {
	return i.ToWaitStepPropertiesOutputWithContext(context.Background())
}

func (i WaitStepPropertiesArgs) ToWaitStepPropertiesOutputWithContext(ctx context.Context) WaitStepPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WaitStepPropertiesOutput)
}

type WaitStepPropertiesOutput struct{ *pulumi.OutputState }

func (WaitStepPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WaitStepProperties)(nil)).Elem()
}

func (o WaitStepPropertiesOutput) ToWaitStepPropertiesOutput() WaitStepPropertiesOutput {
	return o
}

func (o WaitStepPropertiesOutput) ToWaitStepPropertiesOutputWithContext(ctx context.Context) WaitStepPropertiesOutput {
	return o
}

func (o WaitStepPropertiesOutput) Attributes() WaitStepAttributesPtrOutput {
	return o.ApplyT(func(v WaitStepProperties) *WaitStepAttributes { return v.Attributes }).(WaitStepAttributesPtrOutput)
}

func (o WaitStepPropertiesOutput) StepType() pulumi.StringOutput {
	return o.ApplyT(func(v WaitStepProperties) string { return v.StepType }).(pulumi.StringOutput)
}

type WaitStepPropertiesResponse struct {
	Attributes *WaitStepAttributesResponse `pulumi:"attributes"`
	StepType   string                      `pulumi:"stepType"`
}

type WaitStepPropertiesResponseOutput struct{ *pulumi.OutputState }

func (WaitStepPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WaitStepPropertiesResponse)(nil)).Elem()
}

func (o WaitStepPropertiesResponseOutput) ToWaitStepPropertiesResponseOutput() WaitStepPropertiesResponseOutput {
	return o
}

func (o WaitStepPropertiesResponseOutput) ToWaitStepPropertiesResponseOutputWithContext(ctx context.Context) WaitStepPropertiesResponseOutput {
	return o
}

func (o WaitStepPropertiesResponseOutput) Attributes() WaitStepAttributesResponsePtrOutput {
	return o.ApplyT(func(v WaitStepPropertiesResponse) *WaitStepAttributesResponse { return v.Attributes }).(WaitStepAttributesResponsePtrOutput)
}

func (o WaitStepPropertiesResponseOutput) StepType() pulumi.StringOutput {
	return o.ApplyT(func(v WaitStepPropertiesResponse) string { return v.StepType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(PrePostStepOutput{})
	pulumi.RegisterOutputType(PrePostStepArrayOutput{})
	pulumi.RegisterOutputType(PrePostStepResponseOutput{})
	pulumi.RegisterOutputType(PrePostStepResponseArrayOutput{})
	pulumi.RegisterOutputType(SasAuthenticationOutput{})
	pulumi.RegisterOutputType(SasAuthenticationResponseOutput{})
	pulumi.RegisterOutputType(ServiceUnitArtifactsOutput{})
	pulumi.RegisterOutputType(ServiceUnitArtifactsPtrOutput{})
	pulumi.RegisterOutputType(ServiceUnitArtifactsResponseOutput{})
	pulumi.RegisterOutputType(ServiceUnitArtifactsResponsePtrOutput{})
	pulumi.RegisterOutputType(StepTypeOutput{})
	pulumi.RegisterOutputType(StepTypeArrayOutput{})
	pulumi.RegisterOutputType(StepResponseOutput{})
	pulumi.RegisterOutputType(StepResponseArrayOutput{})
	pulumi.RegisterOutputType(WaitStepAttributesOutput{})
	pulumi.RegisterOutputType(WaitStepAttributesPtrOutput{})
	pulumi.RegisterOutputType(WaitStepAttributesResponseOutput{})
	pulumi.RegisterOutputType(WaitStepAttributesResponsePtrOutput{})
	pulumi.RegisterOutputType(WaitStepPropertiesOutput{})
	pulumi.RegisterOutputType(WaitStepPropertiesResponseOutput{})
}
