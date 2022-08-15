


package v20191101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiKeyAuthentication struct {
	In    RestAuthLocation `pulumi:"in"`
	Name  string           `pulumi:"name"`
	Type  string           `pulumi:"type"`
	Value string           `pulumi:"value"`
}

type ApiKeyAuthenticationResponse struct {
	In    string `pulumi:"in"`
	Name  string `pulumi:"name"`
	Type  string `pulumi:"type"`
	Value string `pulumi:"value"`
}

type CloudErrorBodyResponse struct {
	Code    string                   `pulumi:"code"`
	Details []CloudErrorBodyResponse `pulumi:"details"`
	Message string                   `pulumi:"message"`
	Target  *string                  `pulumi:"target"`
}

type CloudErrorBodyResponseOutput struct{ *pulumi.OutputState }

func (CloudErrorBodyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudErrorBodyResponse)(nil)).Elem()
}

func (o CloudErrorBodyResponseOutput) ToCloudErrorBodyResponseOutput() CloudErrorBodyResponseOutput {
	return o
}

func (o CloudErrorBodyResponseOutput) ToCloudErrorBodyResponseOutputWithContext(ctx context.Context) CloudErrorBodyResponseOutput {
	return o
}

func (o CloudErrorBodyResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v CloudErrorBodyResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o CloudErrorBodyResponseOutput) Details() CloudErrorBodyResponseArrayOutput {
	return o.ApplyT(func(v CloudErrorBodyResponse) []CloudErrorBodyResponse { return v.Details }).(CloudErrorBodyResponseArrayOutput)
}

func (o CloudErrorBodyResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v CloudErrorBodyResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o CloudErrorBodyResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudErrorBodyResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type CloudErrorBodyResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudErrorBodyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudErrorBodyResponse)(nil)).Elem()
}

func (o CloudErrorBodyResponsePtrOutput) ToCloudErrorBodyResponsePtrOutput() CloudErrorBodyResponsePtrOutput {
	return o
}

func (o CloudErrorBodyResponsePtrOutput) ToCloudErrorBodyResponsePtrOutputWithContext(ctx context.Context) CloudErrorBodyResponsePtrOutput {
	return o
}

func (o CloudErrorBodyResponsePtrOutput) Elem() CloudErrorBodyResponseOutput {
	return o.ApplyT(func(v *CloudErrorBodyResponse) CloudErrorBodyResponse {
		if v != nil {
			return *v
		}
		var ret CloudErrorBodyResponse
		return ret
	}).(CloudErrorBodyResponseOutput)
}

func (o CloudErrorBodyResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudErrorBodyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o CloudErrorBodyResponsePtrOutput) Details() CloudErrorBodyResponseArrayOutput {
	return o.ApplyT(func(v *CloudErrorBodyResponse) []CloudErrorBodyResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(CloudErrorBodyResponseArrayOutput)
}

func (o CloudErrorBodyResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudErrorBodyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o CloudErrorBodyResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudErrorBodyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

type CloudErrorBodyResponseArrayOutput struct{ *pulumi.OutputState }

func (CloudErrorBodyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudErrorBodyResponse)(nil)).Elem()
}

func (o CloudErrorBodyResponseArrayOutput) ToCloudErrorBodyResponseArrayOutput() CloudErrorBodyResponseArrayOutput {
	return o
}

func (o CloudErrorBodyResponseArrayOutput) ToCloudErrorBodyResponseArrayOutputWithContext(ctx context.Context) CloudErrorBodyResponseArrayOutput {
	return o
}

func (o CloudErrorBodyResponseArrayOutput) Index(i pulumi.IntInput) CloudErrorBodyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CloudErrorBodyResponse {
		return vs[0].([]CloudErrorBodyResponse)[vs[1].(int)]
	}).(CloudErrorBodyResponseOutput)
}

type HealthCheckStepProperties struct {
	Attributes RestHealthCheckStepAttributes `pulumi:"attributes"`
	StepType   string                        `pulumi:"stepType"`
}

type HealthCheckStepPropertiesResponse struct {
	Attributes RestHealthCheckStepAttributesResponse `pulumi:"attributes"`
	StepType   string                                `pulumi:"stepType"`
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

func (o IdentityResponsePtrOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IdentityResponse) []string {
		if v == nil {
			return nil
		}
		return v.IdentityIds
	}).(pulumi.StringArrayOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type MessageResponse struct {
	Message   string `pulumi:"message"`
	TimeStamp string `pulumi:"timeStamp"`
}

type MessageResponseOutput struct{ *pulumi.OutputState }

func (MessageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MessageResponse)(nil)).Elem()
}

func (o MessageResponseOutput) ToMessageResponseOutput() MessageResponseOutput {
	return o
}

func (o MessageResponseOutput) ToMessageResponseOutputWithContext(ctx context.Context) MessageResponseOutput {
	return o
}

func (o MessageResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v MessageResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o MessageResponseOutput) TimeStamp() pulumi.StringOutput {
	return o.ApplyT(func(v MessageResponse) string { return v.TimeStamp }).(pulumi.StringOutput)
}

type MessageResponseArrayOutput struct{ *pulumi.OutputState }

func (MessageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MessageResponse)(nil)).Elem()
}

func (o MessageResponseArrayOutput) ToMessageResponseArrayOutput() MessageResponseArrayOutput {
	return o
}

func (o MessageResponseArrayOutput) ToMessageResponseArrayOutputWithContext(ctx context.Context) MessageResponseArrayOutput {
	return o
}

func (o MessageResponseArrayOutput) Index(i pulumi.IntInput) MessageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MessageResponse {
		return vs[0].([]MessageResponse)[vs[1].(int)]
	}).(MessageResponseOutput)
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

type ResourceOperationResponseOutput struct{ *pulumi.OutputState }

func (ResourceOperationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceOperationResponse)(nil)).Elem()
}

func (o ResourceOperationResponseOutput) ToResourceOperationResponseOutput() ResourceOperationResponseOutput {
	return o
}

func (o ResourceOperationResponseOutput) ToResourceOperationResponseOutputWithContext(ctx context.Context) ResourceOperationResponseOutput {
	return o
}

func (o ResourceOperationResponseOutput) OperationId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceOperationResponse) string { return v.OperationId }).(pulumi.StringOutput)
}

func (o ResourceOperationResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceOperationResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ResourceOperationResponseOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceOperationResponse) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o ResourceOperationResponseOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceOperationResponse) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o ResourceOperationResponseOutput) StatusCode() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceOperationResponse) string { return v.StatusCode }).(pulumi.StringOutput)
}

func (o ResourceOperationResponseOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceOperationResponse) string { return v.StatusMessage }).(pulumi.StringOutput)
}

type ResourceOperationResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceOperationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceOperationResponse)(nil)).Elem()
}

func (o ResourceOperationResponseArrayOutput) ToResourceOperationResponseArrayOutput() ResourceOperationResponseArrayOutput {
	return o
}

func (o ResourceOperationResponseArrayOutput) ToResourceOperationResponseArrayOutputWithContext(ctx context.Context) ResourceOperationResponseArrayOutput {
	return o
}

func (o ResourceOperationResponseArrayOutput) Index(i pulumi.IntInput) ResourceOperationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceOperationResponse {
		return vs[0].([]ResourceOperationResponse)[vs[1].(int)]
	}).(ResourceOperationResponseOutput)
}

type RestHealthCheck struct {
	Name     string        `pulumi:"name"`
	Request  RestRequest   `pulumi:"request"`
	Response *RestResponse `pulumi:"response"`
}

type RestHealthCheckResponse struct {
	Name     string                `pulumi:"name"`
	Request  RestRequestResponse   `pulumi:"request"`
	Response *RestResponseResponse `pulumi:"response"`
}

type RestHealthCheckStepAttributes struct {
	HealthChecks         []RestHealthCheck `pulumi:"healthChecks"`
	HealthyStateDuration string            `pulumi:"healthyStateDuration"`
	MaxElasticDuration   *string           `pulumi:"maxElasticDuration"`
	Type                 string            `pulumi:"type"`
	WaitDuration         *string           `pulumi:"waitDuration"`
}

type RestHealthCheckStepAttributesResponse struct {
	HealthChecks         []RestHealthCheckResponse `pulumi:"healthChecks"`
	HealthyStateDuration string                    `pulumi:"healthyStateDuration"`
	MaxElasticDuration   *string                   `pulumi:"maxElasticDuration"`
	Type                 string                    `pulumi:"type"`
	WaitDuration         *string                   `pulumi:"waitDuration"`
}

type RestRequest struct {
	Authentication interface{}       `pulumi:"authentication"`
	Method         RestRequestMethod `pulumi:"method"`
	Uri            string            `pulumi:"uri"`
}

type RestRequestResponse struct {
	Authentication interface{} `pulumi:"authentication"`
	Method         string      `pulumi:"method"`
	Uri            string      `pulumi:"uri"`
}

type RestResponse struct {
	Regex              *RestResponseRegex `pulumi:"regex"`
	SuccessStatusCodes []string           `pulumi:"successStatusCodes"`
}

type RestResponseRegex struct {
	MatchQuantifier *RestMatchQuantifier `pulumi:"matchQuantifier"`
	Matches         []string             `pulumi:"matches"`
}

type RestResponseResponse struct {
	Regex              *RestResponseResponseRegex `pulumi:"regex"`
	SuccessStatusCodes []string                   `pulumi:"successStatusCodes"`
}

type RestResponseResponseRegex struct {
	MatchQuantifier *string  `pulumi:"matchQuantifier"`
	Matches         []string `pulumi:"matches"`
}

type RolloutIdentityAuthentication struct {
	Type string `pulumi:"type"`
}

type RolloutIdentityAuthenticationResponse struct {
	Type string `pulumi:"type"`
}

type RolloutOperationInfoResponse struct {
	EndTime              string                 `pulumi:"endTime"`
	Error                CloudErrorBodyResponse `pulumi:"error"`
	RetryAttempt         int                    `pulumi:"retryAttempt"`
	SkipSucceededOnRetry bool                   `pulumi:"skipSucceededOnRetry"`
	StartTime            string                 `pulumi:"startTime"`
}

type RolloutOperationInfoResponseOutput struct{ *pulumi.OutputState }

func (RolloutOperationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RolloutOperationInfoResponse)(nil)).Elem()
}

func (o RolloutOperationInfoResponseOutput) ToRolloutOperationInfoResponseOutput() RolloutOperationInfoResponseOutput {
	return o
}

func (o RolloutOperationInfoResponseOutput) ToRolloutOperationInfoResponseOutputWithContext(ctx context.Context) RolloutOperationInfoResponseOutput {
	return o
}

func (o RolloutOperationInfoResponseOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v RolloutOperationInfoResponse) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o RolloutOperationInfoResponseOutput) Error() CloudErrorBodyResponseOutput {
	return o.ApplyT(func(v RolloutOperationInfoResponse) CloudErrorBodyResponse { return v.Error }).(CloudErrorBodyResponseOutput)
}

func (o RolloutOperationInfoResponseOutput) RetryAttempt() pulumi.IntOutput {
	return o.ApplyT(func(v RolloutOperationInfoResponse) int { return v.RetryAttempt }).(pulumi.IntOutput)
}

func (o RolloutOperationInfoResponseOutput) SkipSucceededOnRetry() pulumi.BoolOutput {
	return o.ApplyT(func(v RolloutOperationInfoResponse) bool { return v.SkipSucceededOnRetry }).(pulumi.BoolOutput)
}

func (o RolloutOperationInfoResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v RolloutOperationInfoResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

type RolloutStepResponse struct {
	Messages           []MessageResponse           `pulumi:"messages"`
	Name               string                      `pulumi:"name"`
	OperationInfo      StepOperationInfoResponse   `pulumi:"operationInfo"`
	ResourceOperations []ResourceOperationResponse `pulumi:"resourceOperations"`
	Status             string                      `pulumi:"status"`
	StepGroup          *string                     `pulumi:"stepGroup"`
}

type RolloutStepResponseOutput struct{ *pulumi.OutputState }

func (RolloutStepResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RolloutStepResponse)(nil)).Elem()
}

func (o RolloutStepResponseOutput) ToRolloutStepResponseOutput() RolloutStepResponseOutput {
	return o
}

func (o RolloutStepResponseOutput) ToRolloutStepResponseOutputWithContext(ctx context.Context) RolloutStepResponseOutput {
	return o
}

func (o RolloutStepResponseOutput) Messages() MessageResponseArrayOutput {
	return o.ApplyT(func(v RolloutStepResponse) []MessageResponse { return v.Messages }).(MessageResponseArrayOutput)
}

func (o RolloutStepResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RolloutStepResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RolloutStepResponseOutput) OperationInfo() StepOperationInfoResponseOutput {
	return o.ApplyT(func(v RolloutStepResponse) StepOperationInfoResponse { return v.OperationInfo }).(StepOperationInfoResponseOutput)
}

func (o RolloutStepResponseOutput) ResourceOperations() ResourceOperationResponseArrayOutput {
	return o.ApplyT(func(v RolloutStepResponse) []ResourceOperationResponse { return v.ResourceOperations }).(ResourceOperationResponseArrayOutput)
}

func (o RolloutStepResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v RolloutStepResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o RolloutStepResponseOutput) StepGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RolloutStepResponse) *string { return v.StepGroup }).(pulumi.StringPtrOutput)
}

type RolloutStepResponseArrayOutput struct{ *pulumi.OutputState }

func (RolloutStepResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RolloutStepResponse)(nil)).Elem()
}

func (o RolloutStepResponseArrayOutput) ToRolloutStepResponseArrayOutput() RolloutStepResponseArrayOutput {
	return o
}

func (o RolloutStepResponseArrayOutput) ToRolloutStepResponseArrayOutputWithContext(ctx context.Context) RolloutStepResponseArrayOutput {
	return o
}

func (o RolloutStepResponseArrayOutput) Index(i pulumi.IntInput) RolloutStepResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RolloutStepResponse {
		return vs[0].([]RolloutStepResponse)[vs[1].(int)]
	}).(RolloutStepResponseOutput)
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

type ServiceResponseOutput struct{ *pulumi.OutputState }

func (ServiceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceResponse)(nil)).Elem()
}

func (o ServiceResponseOutput) ToServiceResponseOutput() ServiceResponseOutput {
	return o
}

func (o ServiceResponseOutput) ToServiceResponseOutputWithContext(ctx context.Context) ServiceResponseOutput {
	return o
}

func (o ServiceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServiceResponseOutput) ServiceUnits() ServiceUnitResponseArrayOutput {
	return o.ApplyT(func(v ServiceResponse) []ServiceUnitResponse { return v.ServiceUnits }).(ServiceUnitResponseArrayOutput)
}

func (o ServiceResponseOutput) TargetLocation() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResponse) string { return v.TargetLocation }).(pulumi.StringOutput)
}

func (o ServiceResponseOutput) TargetSubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResponse) string { return v.TargetSubscriptionId }).(pulumi.StringOutput)
}

type ServiceResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceResponse)(nil)).Elem()
}

func (o ServiceResponseArrayOutput) ToServiceResponseArrayOutput() ServiceResponseArrayOutput {
	return o
}

func (o ServiceResponseArrayOutput) ToServiceResponseArrayOutputWithContext(ctx context.Context) ServiceResponseArrayOutput {
	return o
}

func (o ServiceResponseArrayOutput) Index(i pulumi.IntInput) ServiceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceResponse {
		return vs[0].([]ServiceResponse)[vs[1].(int)]
	}).(ServiceResponseOutput)
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

type ServiceUnitResponseOutput struct{ *pulumi.OutputState }

func (ServiceUnitResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceUnitResponse)(nil)).Elem()
}

func (o ServiceUnitResponseOutput) ToServiceUnitResponseOutput() ServiceUnitResponseOutput {
	return o
}

func (o ServiceUnitResponseOutput) ToServiceUnitResponseOutputWithContext(ctx context.Context) ServiceUnitResponseOutput {
	return o
}

func (o ServiceUnitResponseOutput) Artifacts() ServiceUnitArtifactsResponsePtrOutput {
	return o.ApplyT(func(v ServiceUnitResponse) *ServiceUnitArtifactsResponse { return v.Artifacts }).(ServiceUnitArtifactsResponsePtrOutput)
}

func (o ServiceUnitResponseOutput) DeploymentMode() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceUnitResponse) string { return v.DeploymentMode }).(pulumi.StringOutput)
}

func (o ServiceUnitResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceUnitResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServiceUnitResponseOutput) Steps() RolloutStepResponseArrayOutput {
	return o.ApplyT(func(v ServiceUnitResponse) []RolloutStepResponse { return v.Steps }).(RolloutStepResponseArrayOutput)
}

func (o ServiceUnitResponseOutput) TargetResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceUnitResponse) string { return v.TargetResourceGroup }).(pulumi.StringOutput)
}

type ServiceUnitResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceUnitResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceUnitResponse)(nil)).Elem()
}

func (o ServiceUnitResponseArrayOutput) ToServiceUnitResponseArrayOutput() ServiceUnitResponseArrayOutput {
	return o
}

func (o ServiceUnitResponseArrayOutput) ToServiceUnitResponseArrayOutputWithContext(ctx context.Context) ServiceUnitResponseArrayOutput {
	return o
}

func (o ServiceUnitResponseArrayOutput) Index(i pulumi.IntInput) ServiceUnitResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceUnitResponse {
		return vs[0].([]ServiceUnitResponse)[vs[1].(int)]
	}).(ServiceUnitResponseOutput)
}

type StepGroup struct {
	DependsOnStepGroups []string      `pulumi:"dependsOnStepGroups"`
	DeploymentTargetId  string        `pulumi:"deploymentTargetId"`
	Name                string        `pulumi:"name"`
	PostDeploymentSteps []PrePostStep `pulumi:"postDeploymentSteps"`
	PreDeploymentSteps  []PrePostStep `pulumi:"preDeploymentSteps"`
}





type StepGroupInput interface {
	pulumi.Input

	ToStepGroupOutput() StepGroupOutput
	ToStepGroupOutputWithContext(context.Context) StepGroupOutput
}

type StepGroupArgs struct {
	DependsOnStepGroups pulumi.StringArrayInput `pulumi:"dependsOnStepGroups"`
	DeploymentTargetId  pulumi.StringInput      `pulumi:"deploymentTargetId"`
	Name                pulumi.StringInput      `pulumi:"name"`
	PostDeploymentSteps PrePostStepArrayInput   `pulumi:"postDeploymentSteps"`
	PreDeploymentSteps  PrePostStepArrayInput   `pulumi:"preDeploymentSteps"`
}

func (StepGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StepGroup)(nil)).Elem()
}

func (i StepGroupArgs) ToStepGroupOutput() StepGroupOutput {
	return i.ToStepGroupOutputWithContext(context.Background())
}

func (i StepGroupArgs) ToStepGroupOutputWithContext(ctx context.Context) StepGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StepGroupOutput)
}





type StepGroupArrayInput interface {
	pulumi.Input

	ToStepGroupArrayOutput() StepGroupArrayOutput
	ToStepGroupArrayOutputWithContext(context.Context) StepGroupArrayOutput
}

type StepGroupArray []StepGroupInput

func (StepGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StepGroup)(nil)).Elem()
}

func (i StepGroupArray) ToStepGroupArrayOutput() StepGroupArrayOutput {
	return i.ToStepGroupArrayOutputWithContext(context.Background())
}

func (i StepGroupArray) ToStepGroupArrayOutputWithContext(ctx context.Context) StepGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StepGroupArrayOutput)
}

type StepGroupOutput struct{ *pulumi.OutputState }

func (StepGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StepGroup)(nil)).Elem()
}

func (o StepGroupOutput) ToStepGroupOutput() StepGroupOutput {
	return o
}

func (o StepGroupOutput) ToStepGroupOutputWithContext(ctx context.Context) StepGroupOutput {
	return o
}

func (o StepGroupOutput) DependsOnStepGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StepGroup) []string { return v.DependsOnStepGroups }).(pulumi.StringArrayOutput)
}

func (o StepGroupOutput) DeploymentTargetId() pulumi.StringOutput {
	return o.ApplyT(func(v StepGroup) string { return v.DeploymentTargetId }).(pulumi.StringOutput)
}

func (o StepGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v StepGroup) string { return v.Name }).(pulumi.StringOutput)
}

func (o StepGroupOutput) PostDeploymentSteps() PrePostStepArrayOutput {
	return o.ApplyT(func(v StepGroup) []PrePostStep { return v.PostDeploymentSteps }).(PrePostStepArrayOutput)
}

func (o StepGroupOutput) PreDeploymentSteps() PrePostStepArrayOutput {
	return o.ApplyT(func(v StepGroup) []PrePostStep { return v.PreDeploymentSteps }).(PrePostStepArrayOutput)
}

type StepGroupArrayOutput struct{ *pulumi.OutputState }

func (StepGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StepGroup)(nil)).Elem()
}

func (o StepGroupArrayOutput) ToStepGroupArrayOutput() StepGroupArrayOutput {
	return o
}

func (o StepGroupArrayOutput) ToStepGroupArrayOutputWithContext(ctx context.Context) StepGroupArrayOutput {
	return o
}

func (o StepGroupArrayOutput) Index(i pulumi.IntInput) StepGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StepGroup {
		return vs[0].([]StepGroup)[vs[1].(int)]
	}).(StepGroupOutput)
}

type StepGroupResponse struct {
	DependsOnStepGroups []string              `pulumi:"dependsOnStepGroups"`
	DeploymentTargetId  string                `pulumi:"deploymentTargetId"`
	Name                string                `pulumi:"name"`
	PostDeploymentSteps []PrePostStepResponse `pulumi:"postDeploymentSteps"`
	PreDeploymentSteps  []PrePostStepResponse `pulumi:"preDeploymentSteps"`
}

type StepGroupResponseOutput struct{ *pulumi.OutputState }

func (StepGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StepGroupResponse)(nil)).Elem()
}

func (o StepGroupResponseOutput) ToStepGroupResponseOutput() StepGroupResponseOutput {
	return o
}

func (o StepGroupResponseOutput) ToStepGroupResponseOutputWithContext(ctx context.Context) StepGroupResponseOutput {
	return o
}

func (o StepGroupResponseOutput) DependsOnStepGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StepGroupResponse) []string { return v.DependsOnStepGroups }).(pulumi.StringArrayOutput)
}

func (o StepGroupResponseOutput) DeploymentTargetId() pulumi.StringOutput {
	return o.ApplyT(func(v StepGroupResponse) string { return v.DeploymentTargetId }).(pulumi.StringOutput)
}

func (o StepGroupResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v StepGroupResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o StepGroupResponseOutput) PostDeploymentSteps() PrePostStepResponseArrayOutput {
	return o.ApplyT(func(v StepGroupResponse) []PrePostStepResponse { return v.PostDeploymentSteps }).(PrePostStepResponseArrayOutput)
}

func (o StepGroupResponseOutput) PreDeploymentSteps() PrePostStepResponseArrayOutput {
	return o.ApplyT(func(v StepGroupResponse) []PrePostStepResponse { return v.PreDeploymentSteps }).(PrePostStepResponseArrayOutput)
}

type StepGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (StepGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StepGroupResponse)(nil)).Elem()
}

func (o StepGroupResponseArrayOutput) ToStepGroupResponseArrayOutput() StepGroupResponseArrayOutput {
	return o
}

func (o StepGroupResponseArrayOutput) ToStepGroupResponseArrayOutputWithContext(ctx context.Context) StepGroupResponseArrayOutput {
	return o
}

func (o StepGroupResponseArrayOutput) Index(i pulumi.IntInput) StepGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StepGroupResponse {
		return vs[0].([]StepGroupResponse)[vs[1].(int)]
	}).(StepGroupResponseOutput)
}

type StepOperationInfoResponse struct {
	CorrelationId   string                  `pulumi:"correlationId"`
	DeploymentName  string                  `pulumi:"deploymentName"`
	EndTime         string                  `pulumi:"endTime"`
	Error           *CloudErrorBodyResponse `pulumi:"error"`
	LastUpdatedTime string                  `pulumi:"lastUpdatedTime"`
	StartTime       string                  `pulumi:"startTime"`
}

type StepOperationInfoResponseOutput struct{ *pulumi.OutputState }

func (StepOperationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StepOperationInfoResponse)(nil)).Elem()
}

func (o StepOperationInfoResponseOutput) ToStepOperationInfoResponseOutput() StepOperationInfoResponseOutput {
	return o
}

func (o StepOperationInfoResponseOutput) ToStepOperationInfoResponseOutputWithContext(ctx context.Context) StepOperationInfoResponseOutput {
	return o
}

func (o StepOperationInfoResponseOutput) CorrelationId() pulumi.StringOutput {
	return o.ApplyT(func(v StepOperationInfoResponse) string { return v.CorrelationId }).(pulumi.StringOutput)
}

func (o StepOperationInfoResponseOutput) DeploymentName() pulumi.StringOutput {
	return o.ApplyT(func(v StepOperationInfoResponse) string { return v.DeploymentName }).(pulumi.StringOutput)
}

func (o StepOperationInfoResponseOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v StepOperationInfoResponse) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o StepOperationInfoResponseOutput) Error() CloudErrorBodyResponsePtrOutput {
	return o.ApplyT(func(v StepOperationInfoResponse) *CloudErrorBodyResponse { return v.Error }).(CloudErrorBodyResponsePtrOutput)
}

func (o StepOperationInfoResponseOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v StepOperationInfoResponse) string { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

func (o StepOperationInfoResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v StepOperationInfoResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

type WaitStepAttributes struct {
	Duration string `pulumi:"duration"`
}

type WaitStepAttributesResponse struct {
	Duration string `pulumi:"duration"`
}

type WaitStepProperties struct {
	Attributes WaitStepAttributes `pulumi:"attributes"`
	StepType   string             `pulumi:"stepType"`
}

type WaitStepPropertiesResponse struct {
	Attributes WaitStepAttributesResponse `pulumi:"attributes"`
	StepType   string                     `pulumi:"stepType"`
}

func init() {
	pulumi.RegisterOutputType(CloudErrorBodyResponseOutput{})
	pulumi.RegisterOutputType(CloudErrorBodyResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudErrorBodyResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(MessageResponseOutput{})
	pulumi.RegisterOutputType(MessageResponseArrayOutput{})
	pulumi.RegisterOutputType(PrePostStepOutput{})
	pulumi.RegisterOutputType(PrePostStepArrayOutput{})
	pulumi.RegisterOutputType(PrePostStepResponseOutput{})
	pulumi.RegisterOutputType(PrePostStepResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceOperationResponseOutput{})
	pulumi.RegisterOutputType(ResourceOperationResponseArrayOutput{})
	pulumi.RegisterOutputType(RolloutOperationInfoResponseOutput{})
	pulumi.RegisterOutputType(RolloutStepResponseOutput{})
	pulumi.RegisterOutputType(RolloutStepResponseArrayOutput{})
	pulumi.RegisterOutputType(SasAuthenticationOutput{})
	pulumi.RegisterOutputType(SasAuthenticationResponseOutput{})
	pulumi.RegisterOutputType(ServiceResponseOutput{})
	pulumi.RegisterOutputType(ServiceResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceUnitArtifactsOutput{})
	pulumi.RegisterOutputType(ServiceUnitArtifactsPtrOutput{})
	pulumi.RegisterOutputType(ServiceUnitArtifactsResponseOutput{})
	pulumi.RegisterOutputType(ServiceUnitArtifactsResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceUnitResponseOutput{})
	pulumi.RegisterOutputType(ServiceUnitResponseArrayOutput{})
	pulumi.RegisterOutputType(StepGroupOutput{})
	pulumi.RegisterOutputType(StepGroupArrayOutput{})
	pulumi.RegisterOutputType(StepGroupResponseOutput{})
	pulumi.RegisterOutputType(StepGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(StepOperationInfoResponseOutput{})
}
