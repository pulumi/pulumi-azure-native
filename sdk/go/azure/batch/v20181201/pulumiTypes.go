


package v20181201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationPackageReference struct {
	Id      string  `pulumi:"id"`
	Version *string `pulumi:"version"`
}





type ApplicationPackageReferenceInput interface {
	pulumi.Input

	ToApplicationPackageReferenceOutput() ApplicationPackageReferenceOutput
	ToApplicationPackageReferenceOutputWithContext(context.Context) ApplicationPackageReferenceOutput
}

type ApplicationPackageReferenceArgs struct {
	Id      pulumi.StringInput    `pulumi:"id"`
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (ApplicationPackageReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackageReference)(nil)).Elem()
}

func (i ApplicationPackageReferenceArgs) ToApplicationPackageReferenceOutput() ApplicationPackageReferenceOutput {
	return i.ToApplicationPackageReferenceOutputWithContext(context.Background())
}

func (i ApplicationPackageReferenceArgs) ToApplicationPackageReferenceOutputWithContext(ctx context.Context) ApplicationPackageReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPackageReferenceOutput)
}





type ApplicationPackageReferenceArrayInput interface {
	pulumi.Input

	ToApplicationPackageReferenceArrayOutput() ApplicationPackageReferenceArrayOutput
	ToApplicationPackageReferenceArrayOutputWithContext(context.Context) ApplicationPackageReferenceArrayOutput
}

type ApplicationPackageReferenceArray []ApplicationPackageReferenceInput

func (ApplicationPackageReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationPackageReference)(nil)).Elem()
}

func (i ApplicationPackageReferenceArray) ToApplicationPackageReferenceArrayOutput() ApplicationPackageReferenceArrayOutput {
	return i.ToApplicationPackageReferenceArrayOutputWithContext(context.Background())
}

func (i ApplicationPackageReferenceArray) ToApplicationPackageReferenceArrayOutputWithContext(ctx context.Context) ApplicationPackageReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPackageReferenceArrayOutput)
}

type ApplicationPackageReferenceOutput struct{ *pulumi.OutputState }

func (ApplicationPackageReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackageReference)(nil)).Elem()
}

func (o ApplicationPackageReferenceOutput) ToApplicationPackageReferenceOutput() ApplicationPackageReferenceOutput {
	return o
}

func (o ApplicationPackageReferenceOutput) ToApplicationPackageReferenceOutputWithContext(ctx context.Context) ApplicationPackageReferenceOutput {
	return o
}

func (o ApplicationPackageReferenceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageReference) string { return v.Id }).(pulumi.StringOutput)
}

func (o ApplicationPackageReferenceOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageReference) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ApplicationPackageReferenceArrayOutput struct{ *pulumi.OutputState }

func (ApplicationPackageReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationPackageReference)(nil)).Elem()
}

func (o ApplicationPackageReferenceArrayOutput) ToApplicationPackageReferenceArrayOutput() ApplicationPackageReferenceArrayOutput {
	return o
}

func (o ApplicationPackageReferenceArrayOutput) ToApplicationPackageReferenceArrayOutputWithContext(ctx context.Context) ApplicationPackageReferenceArrayOutput {
	return o
}

func (o ApplicationPackageReferenceArrayOutput) Index(i pulumi.IntInput) ApplicationPackageReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationPackageReference {
		return vs[0].([]ApplicationPackageReference)[vs[1].(int)]
	}).(ApplicationPackageReferenceOutput)
}

type ApplicationPackageReferenceResponse struct {
	Id      string  `pulumi:"id"`
	Version *string `pulumi:"version"`
}

type ApplicationPackageReferenceResponseOutput struct{ *pulumi.OutputState }

func (ApplicationPackageReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackageReferenceResponse)(nil)).Elem()
}

func (o ApplicationPackageReferenceResponseOutput) ToApplicationPackageReferenceResponseOutput() ApplicationPackageReferenceResponseOutput {
	return o
}

func (o ApplicationPackageReferenceResponseOutput) ToApplicationPackageReferenceResponseOutputWithContext(ctx context.Context) ApplicationPackageReferenceResponseOutput {
	return o
}

func (o ApplicationPackageReferenceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageReferenceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ApplicationPackageReferenceResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageReferenceResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ApplicationPackageReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationPackageReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationPackageReferenceResponse)(nil)).Elem()
}

func (o ApplicationPackageReferenceResponseArrayOutput) ToApplicationPackageReferenceResponseArrayOutput() ApplicationPackageReferenceResponseArrayOutput {
	return o
}

func (o ApplicationPackageReferenceResponseArrayOutput) ToApplicationPackageReferenceResponseArrayOutputWithContext(ctx context.Context) ApplicationPackageReferenceResponseArrayOutput {
	return o
}

func (o ApplicationPackageReferenceResponseArrayOutput) Index(i pulumi.IntInput) ApplicationPackageReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationPackageReferenceResponse {
		return vs[0].([]ApplicationPackageReferenceResponse)[vs[1].(int)]
	}).(ApplicationPackageReferenceResponseOutput)
}

type AutoScaleRunErrorResponse struct {
	Code    string                      `pulumi:"code"`
	Details []AutoScaleRunErrorResponse `pulumi:"details"`
	Message string                      `pulumi:"message"`
}

type AutoScaleRunErrorResponseOutput struct{ *pulumi.OutputState }

func (AutoScaleRunErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleRunErrorResponse)(nil)).Elem()
}

func (o AutoScaleRunErrorResponseOutput) ToAutoScaleRunErrorResponseOutput() AutoScaleRunErrorResponseOutput {
	return o
}

func (o AutoScaleRunErrorResponseOutput) ToAutoScaleRunErrorResponseOutputWithContext(ctx context.Context) AutoScaleRunErrorResponseOutput {
	return o
}

func (o AutoScaleRunErrorResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v AutoScaleRunErrorResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o AutoScaleRunErrorResponseOutput) Details() AutoScaleRunErrorResponseArrayOutput {
	return o.ApplyT(func(v AutoScaleRunErrorResponse) []AutoScaleRunErrorResponse { return v.Details }).(AutoScaleRunErrorResponseArrayOutput)
}

func (o AutoScaleRunErrorResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v AutoScaleRunErrorResponse) string { return v.Message }).(pulumi.StringOutput)
}

type AutoScaleRunErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoScaleRunErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScaleRunErrorResponse)(nil)).Elem()
}

func (o AutoScaleRunErrorResponsePtrOutput) ToAutoScaleRunErrorResponsePtrOutput() AutoScaleRunErrorResponsePtrOutput {
	return o
}

func (o AutoScaleRunErrorResponsePtrOutput) ToAutoScaleRunErrorResponsePtrOutputWithContext(ctx context.Context) AutoScaleRunErrorResponsePtrOutput {
	return o
}

func (o AutoScaleRunErrorResponsePtrOutput) Elem() AutoScaleRunErrorResponseOutput {
	return o.ApplyT(func(v *AutoScaleRunErrorResponse) AutoScaleRunErrorResponse {
		if v != nil {
			return *v
		}
		var ret AutoScaleRunErrorResponse
		return ret
	}).(AutoScaleRunErrorResponseOutput)
}

func (o AutoScaleRunErrorResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScaleRunErrorResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o AutoScaleRunErrorResponsePtrOutput) Details() AutoScaleRunErrorResponseArrayOutput {
	return o.ApplyT(func(v *AutoScaleRunErrorResponse) []AutoScaleRunErrorResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(AutoScaleRunErrorResponseArrayOutput)
}

func (o AutoScaleRunErrorResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScaleRunErrorResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

type AutoScaleRunErrorResponseArrayOutput struct{ *pulumi.OutputState }

func (AutoScaleRunErrorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutoScaleRunErrorResponse)(nil)).Elem()
}

func (o AutoScaleRunErrorResponseArrayOutput) ToAutoScaleRunErrorResponseArrayOutput() AutoScaleRunErrorResponseArrayOutput {
	return o
}

func (o AutoScaleRunErrorResponseArrayOutput) ToAutoScaleRunErrorResponseArrayOutputWithContext(ctx context.Context) AutoScaleRunErrorResponseArrayOutput {
	return o
}

func (o AutoScaleRunErrorResponseArrayOutput) Index(i pulumi.IntInput) AutoScaleRunErrorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutoScaleRunErrorResponse {
		return vs[0].([]AutoScaleRunErrorResponse)[vs[1].(int)]
	}).(AutoScaleRunErrorResponseOutput)
}

type AutoScaleRunResponse struct {
	Error          *AutoScaleRunErrorResponse `pulumi:"error"`
	EvaluationTime string                     `pulumi:"evaluationTime"`
	Results        *string                    `pulumi:"results"`
}

type AutoScaleRunResponseOutput struct{ *pulumi.OutputState }

func (AutoScaleRunResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleRunResponse)(nil)).Elem()
}

func (o AutoScaleRunResponseOutput) ToAutoScaleRunResponseOutput() AutoScaleRunResponseOutput {
	return o
}

func (o AutoScaleRunResponseOutput) ToAutoScaleRunResponseOutputWithContext(ctx context.Context) AutoScaleRunResponseOutput {
	return o
}

func (o AutoScaleRunResponseOutput) Error() AutoScaleRunErrorResponsePtrOutput {
	return o.ApplyT(func(v AutoScaleRunResponse) *AutoScaleRunErrorResponse { return v.Error }).(AutoScaleRunErrorResponsePtrOutput)
}

func (o AutoScaleRunResponseOutput) EvaluationTime() pulumi.StringOutput {
	return o.ApplyT(func(v AutoScaleRunResponse) string { return v.EvaluationTime }).(pulumi.StringOutput)
}

func (o AutoScaleRunResponseOutput) Results() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoScaleRunResponse) *string { return v.Results }).(pulumi.StringPtrOutput)
}

type AutoScaleSettings struct {
	EvaluationInterval *string `pulumi:"evaluationInterval"`
	Formula            string  `pulumi:"formula"`
}





type AutoScaleSettingsInput interface {
	pulumi.Input

	ToAutoScaleSettingsOutput() AutoScaleSettingsOutput
	ToAutoScaleSettingsOutputWithContext(context.Context) AutoScaleSettingsOutput
}

type AutoScaleSettingsArgs struct {
	EvaluationInterval pulumi.StringPtrInput `pulumi:"evaluationInterval"`
	Formula            pulumi.StringInput    `pulumi:"formula"`
}

func (AutoScaleSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleSettings)(nil)).Elem()
}

func (i AutoScaleSettingsArgs) ToAutoScaleSettingsOutput() AutoScaleSettingsOutput {
	return i.ToAutoScaleSettingsOutputWithContext(context.Background())
}

func (i AutoScaleSettingsArgs) ToAutoScaleSettingsOutputWithContext(ctx context.Context) AutoScaleSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScaleSettingsOutput)
}

func (i AutoScaleSettingsArgs) ToAutoScaleSettingsPtrOutput() AutoScaleSettingsPtrOutput {
	return i.ToAutoScaleSettingsPtrOutputWithContext(context.Background())
}

func (i AutoScaleSettingsArgs) ToAutoScaleSettingsPtrOutputWithContext(ctx context.Context) AutoScaleSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScaleSettingsOutput).ToAutoScaleSettingsPtrOutputWithContext(ctx)
}









type AutoScaleSettingsPtrInput interface {
	pulumi.Input

	ToAutoScaleSettingsPtrOutput() AutoScaleSettingsPtrOutput
	ToAutoScaleSettingsPtrOutputWithContext(context.Context) AutoScaleSettingsPtrOutput
}

type autoScaleSettingsPtrType AutoScaleSettingsArgs

func AutoScaleSettingsPtr(v *AutoScaleSettingsArgs) AutoScaleSettingsPtrInput {
	return (*autoScaleSettingsPtrType)(v)
}

func (*autoScaleSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScaleSettings)(nil)).Elem()
}

func (i *autoScaleSettingsPtrType) ToAutoScaleSettingsPtrOutput() AutoScaleSettingsPtrOutput {
	return i.ToAutoScaleSettingsPtrOutputWithContext(context.Background())
}

func (i *autoScaleSettingsPtrType) ToAutoScaleSettingsPtrOutputWithContext(ctx context.Context) AutoScaleSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScaleSettingsPtrOutput)
}

type AutoScaleSettingsOutput struct{ *pulumi.OutputState }

func (AutoScaleSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleSettings)(nil)).Elem()
}

func (o AutoScaleSettingsOutput) ToAutoScaleSettingsOutput() AutoScaleSettingsOutput {
	return o
}

func (o AutoScaleSettingsOutput) ToAutoScaleSettingsOutputWithContext(ctx context.Context) AutoScaleSettingsOutput {
	return o
}

func (o AutoScaleSettingsOutput) ToAutoScaleSettingsPtrOutput() AutoScaleSettingsPtrOutput {
	return o.ToAutoScaleSettingsPtrOutputWithContext(context.Background())
}

func (o AutoScaleSettingsOutput) ToAutoScaleSettingsPtrOutputWithContext(ctx context.Context) AutoScaleSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoScaleSettings) *AutoScaleSettings {
		return &v
	}).(AutoScaleSettingsPtrOutput)
}

func (o AutoScaleSettingsOutput) EvaluationInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoScaleSettings) *string { return v.EvaluationInterval }).(pulumi.StringPtrOutput)
}

func (o AutoScaleSettingsOutput) Formula() pulumi.StringOutput {
	return o.ApplyT(func(v AutoScaleSettings) string { return v.Formula }).(pulumi.StringOutput)
}

type AutoScaleSettingsPtrOutput struct{ *pulumi.OutputState }

func (AutoScaleSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScaleSettings)(nil)).Elem()
}

func (o AutoScaleSettingsPtrOutput) ToAutoScaleSettingsPtrOutput() AutoScaleSettingsPtrOutput {
	return o
}

func (o AutoScaleSettingsPtrOutput) ToAutoScaleSettingsPtrOutputWithContext(ctx context.Context) AutoScaleSettingsPtrOutput {
	return o
}

func (o AutoScaleSettingsPtrOutput) Elem() AutoScaleSettingsOutput {
	return o.ApplyT(func(v *AutoScaleSettings) AutoScaleSettings {
		if v != nil {
			return *v
		}
		var ret AutoScaleSettings
		return ret
	}).(AutoScaleSettingsOutput)
}

func (o AutoScaleSettingsPtrOutput) EvaluationInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScaleSettings) *string {
		if v == nil {
			return nil
		}
		return v.EvaluationInterval
	}).(pulumi.StringPtrOutput)
}

func (o AutoScaleSettingsPtrOutput) Formula() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScaleSettings) *string {
		if v == nil {
			return nil
		}
		return &v.Formula
	}).(pulumi.StringPtrOutput)
}

type AutoScaleSettingsResponse struct {
	EvaluationInterval *string `pulumi:"evaluationInterval"`
	Formula            string  `pulumi:"formula"`
}

type AutoScaleSettingsResponseOutput struct{ *pulumi.OutputState }

func (AutoScaleSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleSettingsResponse)(nil)).Elem()
}

func (o AutoScaleSettingsResponseOutput) ToAutoScaleSettingsResponseOutput() AutoScaleSettingsResponseOutput {
	return o
}

func (o AutoScaleSettingsResponseOutput) ToAutoScaleSettingsResponseOutputWithContext(ctx context.Context) AutoScaleSettingsResponseOutput {
	return o
}

func (o AutoScaleSettingsResponseOutput) EvaluationInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoScaleSettingsResponse) *string { return v.EvaluationInterval }).(pulumi.StringPtrOutput)
}

func (o AutoScaleSettingsResponseOutput) Formula() pulumi.StringOutput {
	return o.ApplyT(func(v AutoScaleSettingsResponse) string { return v.Formula }).(pulumi.StringOutput)
}

type AutoScaleSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoScaleSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScaleSettingsResponse)(nil)).Elem()
}

func (o AutoScaleSettingsResponsePtrOutput) ToAutoScaleSettingsResponsePtrOutput() AutoScaleSettingsResponsePtrOutput {
	return o
}

func (o AutoScaleSettingsResponsePtrOutput) ToAutoScaleSettingsResponsePtrOutputWithContext(ctx context.Context) AutoScaleSettingsResponsePtrOutput {
	return o
}

func (o AutoScaleSettingsResponsePtrOutput) Elem() AutoScaleSettingsResponseOutput {
	return o.ApplyT(func(v *AutoScaleSettingsResponse) AutoScaleSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AutoScaleSettingsResponse
		return ret
	}).(AutoScaleSettingsResponseOutput)
}

func (o AutoScaleSettingsResponsePtrOutput) EvaluationInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScaleSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.EvaluationInterval
	}).(pulumi.StringPtrOutput)
}

func (o AutoScaleSettingsResponsePtrOutput) Formula() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScaleSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Formula
	}).(pulumi.StringPtrOutput)
}

type AutoStorageBaseProperties struct {
	StorageAccountId string `pulumi:"storageAccountId"`
}





type AutoStorageBasePropertiesInput interface {
	pulumi.Input

	ToAutoStorageBasePropertiesOutput() AutoStorageBasePropertiesOutput
	ToAutoStorageBasePropertiesOutputWithContext(context.Context) AutoStorageBasePropertiesOutput
}

type AutoStorageBasePropertiesArgs struct {
	StorageAccountId pulumi.StringInput `pulumi:"storageAccountId"`
}

func (AutoStorageBasePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoStorageBaseProperties)(nil)).Elem()
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesOutput() AutoStorageBasePropertiesOutput {
	return i.ToAutoStorageBasePropertiesOutputWithContext(context.Background())
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesOutputWithContext(ctx context.Context) AutoStorageBasePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoStorageBasePropertiesOutput)
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return i.ToAutoStorageBasePropertiesPtrOutputWithContext(context.Background())
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoStorageBasePropertiesOutput).ToAutoStorageBasePropertiesPtrOutputWithContext(ctx)
}









type AutoStorageBasePropertiesPtrInput interface {
	pulumi.Input

	ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput
	ToAutoStorageBasePropertiesPtrOutputWithContext(context.Context) AutoStorageBasePropertiesPtrOutput
}

type autoStorageBasePropertiesPtrType AutoStorageBasePropertiesArgs

func AutoStorageBasePropertiesPtr(v *AutoStorageBasePropertiesArgs) AutoStorageBasePropertiesPtrInput {
	return (*autoStorageBasePropertiesPtrType)(v)
}

func (*autoStorageBasePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoStorageBaseProperties)(nil)).Elem()
}

func (i *autoStorageBasePropertiesPtrType) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return i.ToAutoStorageBasePropertiesPtrOutputWithContext(context.Background())
}

func (i *autoStorageBasePropertiesPtrType) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoStorageBasePropertiesPtrOutput)
}

type AutoStorageBasePropertiesOutput struct{ *pulumi.OutputState }

func (AutoStorageBasePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoStorageBaseProperties)(nil)).Elem()
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesOutput() AutoStorageBasePropertiesOutput {
	return o
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesOutputWithContext(ctx context.Context) AutoStorageBasePropertiesOutput {
	return o
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return o.ToAutoStorageBasePropertiesPtrOutputWithContext(context.Background())
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoStorageBaseProperties) *AutoStorageBaseProperties {
		return &v
	}).(AutoStorageBasePropertiesPtrOutput)
}

func (o AutoStorageBasePropertiesOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v AutoStorageBaseProperties) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

type AutoStorageBasePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AutoStorageBasePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoStorageBaseProperties)(nil)).Elem()
}

func (o AutoStorageBasePropertiesPtrOutput) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return o
}

func (o AutoStorageBasePropertiesPtrOutput) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return o
}

func (o AutoStorageBasePropertiesPtrOutput) Elem() AutoStorageBasePropertiesOutput {
	return o.ApplyT(func(v *AutoStorageBaseProperties) AutoStorageBaseProperties {
		if v != nil {
			return *v
		}
		var ret AutoStorageBaseProperties
		return ret
	}).(AutoStorageBasePropertiesOutput)
}

func (o AutoStorageBasePropertiesPtrOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoStorageBaseProperties) *string {
		if v == nil {
			return nil
		}
		return &v.StorageAccountId
	}).(pulumi.StringPtrOutput)
}

type AutoStoragePropertiesResponse struct {
	LastKeySync      string `pulumi:"lastKeySync"`
	StorageAccountId string `pulumi:"storageAccountId"`
}

type AutoStoragePropertiesResponseOutput struct{ *pulumi.OutputState }

func (AutoStoragePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoStoragePropertiesResponse)(nil)).Elem()
}

func (o AutoStoragePropertiesResponseOutput) ToAutoStoragePropertiesResponseOutput() AutoStoragePropertiesResponseOutput {
	return o
}

func (o AutoStoragePropertiesResponseOutput) ToAutoStoragePropertiesResponseOutputWithContext(ctx context.Context) AutoStoragePropertiesResponseOutput {
	return o
}

func (o AutoStoragePropertiesResponseOutput) LastKeySync() pulumi.StringOutput {
	return o.ApplyT(func(v AutoStoragePropertiesResponse) string { return v.LastKeySync }).(pulumi.StringOutput)
}

func (o AutoStoragePropertiesResponseOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v AutoStoragePropertiesResponse) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

type AutoUserSpecification struct {
	ElevationLevel *ElevationLevel `pulumi:"elevationLevel"`
	Scope          *AutoUserScope  `pulumi:"scope"`
}





type AutoUserSpecificationInput interface {
	pulumi.Input

	ToAutoUserSpecificationOutput() AutoUserSpecificationOutput
	ToAutoUserSpecificationOutputWithContext(context.Context) AutoUserSpecificationOutput
}

type AutoUserSpecificationArgs struct {
	ElevationLevel ElevationLevelPtrInput `pulumi:"elevationLevel"`
	Scope          AutoUserScopePtrInput  `pulumi:"scope"`
}

func (AutoUserSpecificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoUserSpecification)(nil)).Elem()
}

func (i AutoUserSpecificationArgs) ToAutoUserSpecificationOutput() AutoUserSpecificationOutput {
	return i.ToAutoUserSpecificationOutputWithContext(context.Background())
}

func (i AutoUserSpecificationArgs) ToAutoUserSpecificationOutputWithContext(ctx context.Context) AutoUserSpecificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoUserSpecificationOutput)
}

func (i AutoUserSpecificationArgs) ToAutoUserSpecificationPtrOutput() AutoUserSpecificationPtrOutput {
	return i.ToAutoUserSpecificationPtrOutputWithContext(context.Background())
}

func (i AutoUserSpecificationArgs) ToAutoUserSpecificationPtrOutputWithContext(ctx context.Context) AutoUserSpecificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoUserSpecificationOutput).ToAutoUserSpecificationPtrOutputWithContext(ctx)
}









type AutoUserSpecificationPtrInput interface {
	pulumi.Input

	ToAutoUserSpecificationPtrOutput() AutoUserSpecificationPtrOutput
	ToAutoUserSpecificationPtrOutputWithContext(context.Context) AutoUserSpecificationPtrOutput
}

type autoUserSpecificationPtrType AutoUserSpecificationArgs

func AutoUserSpecificationPtr(v *AutoUserSpecificationArgs) AutoUserSpecificationPtrInput {
	return (*autoUserSpecificationPtrType)(v)
}

func (*autoUserSpecificationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoUserSpecification)(nil)).Elem()
}

func (i *autoUserSpecificationPtrType) ToAutoUserSpecificationPtrOutput() AutoUserSpecificationPtrOutput {
	return i.ToAutoUserSpecificationPtrOutputWithContext(context.Background())
}

func (i *autoUserSpecificationPtrType) ToAutoUserSpecificationPtrOutputWithContext(ctx context.Context) AutoUserSpecificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoUserSpecificationPtrOutput)
}

type AutoUserSpecificationOutput struct{ *pulumi.OutputState }

func (AutoUserSpecificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoUserSpecification)(nil)).Elem()
}

func (o AutoUserSpecificationOutput) ToAutoUserSpecificationOutput() AutoUserSpecificationOutput {
	return o
}

func (o AutoUserSpecificationOutput) ToAutoUserSpecificationOutputWithContext(ctx context.Context) AutoUserSpecificationOutput {
	return o
}

func (o AutoUserSpecificationOutput) ToAutoUserSpecificationPtrOutput() AutoUserSpecificationPtrOutput {
	return o.ToAutoUserSpecificationPtrOutputWithContext(context.Background())
}

func (o AutoUserSpecificationOutput) ToAutoUserSpecificationPtrOutputWithContext(ctx context.Context) AutoUserSpecificationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoUserSpecification) *AutoUserSpecification {
		return &v
	}).(AutoUserSpecificationPtrOutput)
}

func (o AutoUserSpecificationOutput) ElevationLevel() ElevationLevelPtrOutput {
	return o.ApplyT(func(v AutoUserSpecification) *ElevationLevel { return v.ElevationLevel }).(ElevationLevelPtrOutput)
}

func (o AutoUserSpecificationOutput) Scope() AutoUserScopePtrOutput {
	return o.ApplyT(func(v AutoUserSpecification) *AutoUserScope { return v.Scope }).(AutoUserScopePtrOutput)
}

type AutoUserSpecificationPtrOutput struct{ *pulumi.OutputState }

func (AutoUserSpecificationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoUserSpecification)(nil)).Elem()
}

func (o AutoUserSpecificationPtrOutput) ToAutoUserSpecificationPtrOutput() AutoUserSpecificationPtrOutput {
	return o
}

func (o AutoUserSpecificationPtrOutput) ToAutoUserSpecificationPtrOutputWithContext(ctx context.Context) AutoUserSpecificationPtrOutput {
	return o
}

func (o AutoUserSpecificationPtrOutput) Elem() AutoUserSpecificationOutput {
	return o.ApplyT(func(v *AutoUserSpecification) AutoUserSpecification {
		if v != nil {
			return *v
		}
		var ret AutoUserSpecification
		return ret
	}).(AutoUserSpecificationOutput)
}

func (o AutoUserSpecificationPtrOutput) ElevationLevel() ElevationLevelPtrOutput {
	return o.ApplyT(func(v *AutoUserSpecification) *ElevationLevel {
		if v == nil {
			return nil
		}
		return v.ElevationLevel
	}).(ElevationLevelPtrOutput)
}

func (o AutoUserSpecificationPtrOutput) Scope() AutoUserScopePtrOutput {
	return o.ApplyT(func(v *AutoUserSpecification) *AutoUserScope {
		if v == nil {
			return nil
		}
		return v.Scope
	}).(AutoUserScopePtrOutput)
}

type AutoUserSpecificationResponse struct {
	ElevationLevel *string `pulumi:"elevationLevel"`
	Scope          *string `pulumi:"scope"`
}

type AutoUserSpecificationResponseOutput struct{ *pulumi.OutputState }

func (AutoUserSpecificationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoUserSpecificationResponse)(nil)).Elem()
}

func (o AutoUserSpecificationResponseOutput) ToAutoUserSpecificationResponseOutput() AutoUserSpecificationResponseOutput {
	return o
}

func (o AutoUserSpecificationResponseOutput) ToAutoUserSpecificationResponseOutputWithContext(ctx context.Context) AutoUserSpecificationResponseOutput {
	return o
}

func (o AutoUserSpecificationResponseOutput) ElevationLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoUserSpecificationResponse) *string { return v.ElevationLevel }).(pulumi.StringPtrOutput)
}

func (o AutoUserSpecificationResponseOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoUserSpecificationResponse) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

type AutoUserSpecificationResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoUserSpecificationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoUserSpecificationResponse)(nil)).Elem()
}

func (o AutoUserSpecificationResponsePtrOutput) ToAutoUserSpecificationResponsePtrOutput() AutoUserSpecificationResponsePtrOutput {
	return o
}

func (o AutoUserSpecificationResponsePtrOutput) ToAutoUserSpecificationResponsePtrOutputWithContext(ctx context.Context) AutoUserSpecificationResponsePtrOutput {
	return o
}

func (o AutoUserSpecificationResponsePtrOutput) Elem() AutoUserSpecificationResponseOutput {
	return o.ApplyT(func(v *AutoUserSpecificationResponse) AutoUserSpecificationResponse {
		if v != nil {
			return *v
		}
		var ret AutoUserSpecificationResponse
		return ret
	}).(AutoUserSpecificationResponseOutput)
}

func (o AutoUserSpecificationResponsePtrOutput) ElevationLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoUserSpecificationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ElevationLevel
	}).(pulumi.StringPtrOutput)
}

func (o AutoUserSpecificationResponsePtrOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoUserSpecificationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Scope
	}).(pulumi.StringPtrOutput)
}

type CertificateReference struct {
	Id            string                    `pulumi:"id"`
	StoreLocation *CertificateStoreLocation `pulumi:"storeLocation"`
	StoreName     *string                   `pulumi:"storeName"`
	Visibility    []CertificateVisibility   `pulumi:"visibility"`
}





type CertificateReferenceInput interface {
	pulumi.Input

	ToCertificateReferenceOutput() CertificateReferenceOutput
	ToCertificateReferenceOutputWithContext(context.Context) CertificateReferenceOutput
}

type CertificateReferenceArgs struct {
	Id            pulumi.StringInput               `pulumi:"id"`
	StoreLocation CertificateStoreLocationPtrInput `pulumi:"storeLocation"`
	StoreName     pulumi.StringPtrInput            `pulumi:"storeName"`
	Visibility    CertificateVisibilityArrayInput  `pulumi:"visibility"`
}

func (CertificateReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateReference)(nil)).Elem()
}

func (i CertificateReferenceArgs) ToCertificateReferenceOutput() CertificateReferenceOutput {
	return i.ToCertificateReferenceOutputWithContext(context.Background())
}

func (i CertificateReferenceArgs) ToCertificateReferenceOutputWithContext(ctx context.Context) CertificateReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateReferenceOutput)
}





type CertificateReferenceArrayInput interface {
	pulumi.Input

	ToCertificateReferenceArrayOutput() CertificateReferenceArrayOutput
	ToCertificateReferenceArrayOutputWithContext(context.Context) CertificateReferenceArrayOutput
}

type CertificateReferenceArray []CertificateReferenceInput

func (CertificateReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateReference)(nil)).Elem()
}

func (i CertificateReferenceArray) ToCertificateReferenceArrayOutput() CertificateReferenceArrayOutput {
	return i.ToCertificateReferenceArrayOutputWithContext(context.Background())
}

func (i CertificateReferenceArray) ToCertificateReferenceArrayOutputWithContext(ctx context.Context) CertificateReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateReferenceArrayOutput)
}

type CertificateReferenceOutput struct{ *pulumi.OutputState }

func (CertificateReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateReference)(nil)).Elem()
}

func (o CertificateReferenceOutput) ToCertificateReferenceOutput() CertificateReferenceOutput {
	return o
}

func (o CertificateReferenceOutput) ToCertificateReferenceOutputWithContext(ctx context.Context) CertificateReferenceOutput {
	return o
}

func (o CertificateReferenceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateReference) string { return v.Id }).(pulumi.StringOutput)
}

func (o CertificateReferenceOutput) StoreLocation() CertificateStoreLocationPtrOutput {
	return o.ApplyT(func(v CertificateReference) *CertificateStoreLocation { return v.StoreLocation }).(CertificateStoreLocationPtrOutput)
}

func (o CertificateReferenceOutput) StoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateReference) *string { return v.StoreName }).(pulumi.StringPtrOutput)
}

func (o CertificateReferenceOutput) Visibility() CertificateVisibilityArrayOutput {
	return o.ApplyT(func(v CertificateReference) []CertificateVisibility { return v.Visibility }).(CertificateVisibilityArrayOutput)
}

type CertificateReferenceArrayOutput struct{ *pulumi.OutputState }

func (CertificateReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateReference)(nil)).Elem()
}

func (o CertificateReferenceArrayOutput) ToCertificateReferenceArrayOutput() CertificateReferenceArrayOutput {
	return o
}

func (o CertificateReferenceArrayOutput) ToCertificateReferenceArrayOutputWithContext(ctx context.Context) CertificateReferenceArrayOutput {
	return o
}

func (o CertificateReferenceArrayOutput) Index(i pulumi.IntInput) CertificateReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CertificateReference {
		return vs[0].([]CertificateReference)[vs[1].(int)]
	}).(CertificateReferenceOutput)
}

type CertificateReferenceResponse struct {
	Id            string   `pulumi:"id"`
	StoreLocation *string  `pulumi:"storeLocation"`
	StoreName     *string  `pulumi:"storeName"`
	Visibility    []string `pulumi:"visibility"`
}

type CertificateReferenceResponseOutput struct{ *pulumi.OutputState }

func (CertificateReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateReferenceResponse)(nil)).Elem()
}

func (o CertificateReferenceResponseOutput) ToCertificateReferenceResponseOutput() CertificateReferenceResponseOutput {
	return o
}

func (o CertificateReferenceResponseOutput) ToCertificateReferenceResponseOutputWithContext(ctx context.Context) CertificateReferenceResponseOutput {
	return o
}

func (o CertificateReferenceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateReferenceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o CertificateReferenceResponseOutput) StoreLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateReferenceResponse) *string { return v.StoreLocation }).(pulumi.StringPtrOutput)
}

func (o CertificateReferenceResponseOutput) StoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateReferenceResponse) *string { return v.StoreName }).(pulumi.StringPtrOutput)
}

func (o CertificateReferenceResponseOutput) Visibility() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CertificateReferenceResponse) []string { return v.Visibility }).(pulumi.StringArrayOutput)
}

type CertificateReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (CertificateReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateReferenceResponse)(nil)).Elem()
}

func (o CertificateReferenceResponseArrayOutput) ToCertificateReferenceResponseArrayOutput() CertificateReferenceResponseArrayOutput {
	return o
}

func (o CertificateReferenceResponseArrayOutput) ToCertificateReferenceResponseArrayOutputWithContext(ctx context.Context) CertificateReferenceResponseArrayOutput {
	return o
}

func (o CertificateReferenceResponseArrayOutput) Index(i pulumi.IntInput) CertificateReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CertificateReferenceResponse {
		return vs[0].([]CertificateReferenceResponse)[vs[1].(int)]
	}).(CertificateReferenceResponseOutput)
}

type CloudServiceConfiguration struct {
	OsFamily  string  `pulumi:"osFamily"`
	OsVersion *string `pulumi:"osVersion"`
}





type CloudServiceConfigurationInput interface {
	pulumi.Input

	ToCloudServiceConfigurationOutput() CloudServiceConfigurationOutput
	ToCloudServiceConfigurationOutputWithContext(context.Context) CloudServiceConfigurationOutput
}

type CloudServiceConfigurationArgs struct {
	OsFamily  pulumi.StringInput    `pulumi:"osFamily"`
	OsVersion pulumi.StringPtrInput `pulumi:"osVersion"`
}

func (CloudServiceConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceConfiguration)(nil)).Elem()
}

func (i CloudServiceConfigurationArgs) ToCloudServiceConfigurationOutput() CloudServiceConfigurationOutput {
	return i.ToCloudServiceConfigurationOutputWithContext(context.Background())
}

func (i CloudServiceConfigurationArgs) ToCloudServiceConfigurationOutputWithContext(ctx context.Context) CloudServiceConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceConfigurationOutput)
}

func (i CloudServiceConfigurationArgs) ToCloudServiceConfigurationPtrOutput() CloudServiceConfigurationPtrOutput {
	return i.ToCloudServiceConfigurationPtrOutputWithContext(context.Background())
}

func (i CloudServiceConfigurationArgs) ToCloudServiceConfigurationPtrOutputWithContext(ctx context.Context) CloudServiceConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceConfigurationOutput).ToCloudServiceConfigurationPtrOutputWithContext(ctx)
}









type CloudServiceConfigurationPtrInput interface {
	pulumi.Input

	ToCloudServiceConfigurationPtrOutput() CloudServiceConfigurationPtrOutput
	ToCloudServiceConfigurationPtrOutputWithContext(context.Context) CloudServiceConfigurationPtrOutput
}

type cloudServiceConfigurationPtrType CloudServiceConfigurationArgs

func CloudServiceConfigurationPtr(v *CloudServiceConfigurationArgs) CloudServiceConfigurationPtrInput {
	return (*cloudServiceConfigurationPtrType)(v)
}

func (*cloudServiceConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceConfiguration)(nil)).Elem()
}

func (i *cloudServiceConfigurationPtrType) ToCloudServiceConfigurationPtrOutput() CloudServiceConfigurationPtrOutput {
	return i.ToCloudServiceConfigurationPtrOutputWithContext(context.Background())
}

func (i *cloudServiceConfigurationPtrType) ToCloudServiceConfigurationPtrOutputWithContext(ctx context.Context) CloudServiceConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceConfigurationPtrOutput)
}

type CloudServiceConfigurationOutput struct{ *pulumi.OutputState }

func (CloudServiceConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceConfiguration)(nil)).Elem()
}

func (o CloudServiceConfigurationOutput) ToCloudServiceConfigurationOutput() CloudServiceConfigurationOutput {
	return o
}

func (o CloudServiceConfigurationOutput) ToCloudServiceConfigurationOutputWithContext(ctx context.Context) CloudServiceConfigurationOutput {
	return o
}

func (o CloudServiceConfigurationOutput) ToCloudServiceConfigurationPtrOutput() CloudServiceConfigurationPtrOutput {
	return o.ToCloudServiceConfigurationPtrOutputWithContext(context.Background())
}

func (o CloudServiceConfigurationOutput) ToCloudServiceConfigurationPtrOutputWithContext(ctx context.Context) CloudServiceConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudServiceConfiguration) *CloudServiceConfiguration {
		return &v
	}).(CloudServiceConfigurationPtrOutput)
}

func (o CloudServiceConfigurationOutput) OsFamily() pulumi.StringOutput {
	return o.ApplyT(func(v CloudServiceConfiguration) string { return v.OsFamily }).(pulumi.StringOutput)
}

func (o CloudServiceConfigurationOutput) OsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceConfiguration) *string { return v.OsVersion }).(pulumi.StringPtrOutput)
}

type CloudServiceConfigurationPtrOutput struct{ *pulumi.OutputState }

func (CloudServiceConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceConfiguration)(nil)).Elem()
}

func (o CloudServiceConfigurationPtrOutput) ToCloudServiceConfigurationPtrOutput() CloudServiceConfigurationPtrOutput {
	return o
}

func (o CloudServiceConfigurationPtrOutput) ToCloudServiceConfigurationPtrOutputWithContext(ctx context.Context) CloudServiceConfigurationPtrOutput {
	return o
}

func (o CloudServiceConfigurationPtrOutput) Elem() CloudServiceConfigurationOutput {
	return o.ApplyT(func(v *CloudServiceConfiguration) CloudServiceConfiguration {
		if v != nil {
			return *v
		}
		var ret CloudServiceConfiguration
		return ret
	}).(CloudServiceConfigurationOutput)
}

func (o CloudServiceConfigurationPtrOutput) OsFamily() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.OsFamily
	}).(pulumi.StringPtrOutput)
}

func (o CloudServiceConfigurationPtrOutput) OsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.OsVersion
	}).(pulumi.StringPtrOutput)
}

type CloudServiceConfigurationResponse struct {
	OsFamily  string  `pulumi:"osFamily"`
	OsVersion *string `pulumi:"osVersion"`
}

type CloudServiceConfigurationResponseOutput struct{ *pulumi.OutputState }

func (CloudServiceConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceConfigurationResponse)(nil)).Elem()
}

func (o CloudServiceConfigurationResponseOutput) ToCloudServiceConfigurationResponseOutput() CloudServiceConfigurationResponseOutput {
	return o
}

func (o CloudServiceConfigurationResponseOutput) ToCloudServiceConfigurationResponseOutputWithContext(ctx context.Context) CloudServiceConfigurationResponseOutput {
	return o
}

func (o CloudServiceConfigurationResponseOutput) OsFamily() pulumi.StringOutput {
	return o.ApplyT(func(v CloudServiceConfigurationResponse) string { return v.OsFamily }).(pulumi.StringOutput)
}

func (o CloudServiceConfigurationResponseOutput) OsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceConfigurationResponse) *string { return v.OsVersion }).(pulumi.StringPtrOutput)
}

type CloudServiceConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudServiceConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceConfigurationResponse)(nil)).Elem()
}

func (o CloudServiceConfigurationResponsePtrOutput) ToCloudServiceConfigurationResponsePtrOutput() CloudServiceConfigurationResponsePtrOutput {
	return o
}

func (o CloudServiceConfigurationResponsePtrOutput) ToCloudServiceConfigurationResponsePtrOutputWithContext(ctx context.Context) CloudServiceConfigurationResponsePtrOutput {
	return o
}

func (o CloudServiceConfigurationResponsePtrOutput) Elem() CloudServiceConfigurationResponseOutput {
	return o.ApplyT(func(v *CloudServiceConfigurationResponse) CloudServiceConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret CloudServiceConfigurationResponse
		return ret
	}).(CloudServiceConfigurationResponseOutput)
}

func (o CloudServiceConfigurationResponsePtrOutput) OsFamily() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OsFamily
	}).(pulumi.StringPtrOutput)
}

func (o CloudServiceConfigurationResponsePtrOutput) OsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsVersion
	}).(pulumi.StringPtrOutput)
}

type ContainerConfiguration struct {
	ContainerImageNames []string            `pulumi:"containerImageNames"`
	ContainerRegistries []ContainerRegistry `pulumi:"containerRegistries"`
	Type                ContainerType       `pulumi:"type"`
}





type ContainerConfigurationInput interface {
	pulumi.Input

	ToContainerConfigurationOutput() ContainerConfigurationOutput
	ToContainerConfigurationOutputWithContext(context.Context) ContainerConfigurationOutput
}

type ContainerConfigurationArgs struct {
	ContainerImageNames pulumi.StringArrayInput     `pulumi:"containerImageNames"`
	ContainerRegistries ContainerRegistryArrayInput `pulumi:"containerRegistries"`
	Type                ContainerTypeInput          `pulumi:"type"`
}

func (ContainerConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerConfiguration)(nil)).Elem()
}

func (i ContainerConfigurationArgs) ToContainerConfigurationOutput() ContainerConfigurationOutput {
	return i.ToContainerConfigurationOutputWithContext(context.Background())
}

func (i ContainerConfigurationArgs) ToContainerConfigurationOutputWithContext(ctx context.Context) ContainerConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerConfigurationOutput)
}

func (i ContainerConfigurationArgs) ToContainerConfigurationPtrOutput() ContainerConfigurationPtrOutput {
	return i.ToContainerConfigurationPtrOutputWithContext(context.Background())
}

func (i ContainerConfigurationArgs) ToContainerConfigurationPtrOutputWithContext(ctx context.Context) ContainerConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerConfigurationOutput).ToContainerConfigurationPtrOutputWithContext(ctx)
}









type ContainerConfigurationPtrInput interface {
	pulumi.Input

	ToContainerConfigurationPtrOutput() ContainerConfigurationPtrOutput
	ToContainerConfigurationPtrOutputWithContext(context.Context) ContainerConfigurationPtrOutput
}

type containerConfigurationPtrType ContainerConfigurationArgs

func ContainerConfigurationPtr(v *ContainerConfigurationArgs) ContainerConfigurationPtrInput {
	return (*containerConfigurationPtrType)(v)
}

func (*containerConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerConfiguration)(nil)).Elem()
}

func (i *containerConfigurationPtrType) ToContainerConfigurationPtrOutput() ContainerConfigurationPtrOutput {
	return i.ToContainerConfigurationPtrOutputWithContext(context.Background())
}

func (i *containerConfigurationPtrType) ToContainerConfigurationPtrOutputWithContext(ctx context.Context) ContainerConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerConfigurationPtrOutput)
}

type ContainerConfigurationOutput struct{ *pulumi.OutputState }

func (ContainerConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerConfiguration)(nil)).Elem()
}

func (o ContainerConfigurationOutput) ToContainerConfigurationOutput() ContainerConfigurationOutput {
	return o
}

func (o ContainerConfigurationOutput) ToContainerConfigurationOutputWithContext(ctx context.Context) ContainerConfigurationOutput {
	return o
}

func (o ContainerConfigurationOutput) ToContainerConfigurationPtrOutput() ContainerConfigurationPtrOutput {
	return o.ToContainerConfigurationPtrOutputWithContext(context.Background())
}

func (o ContainerConfigurationOutput) ToContainerConfigurationPtrOutputWithContext(ctx context.Context) ContainerConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerConfiguration) *ContainerConfiguration {
		return &v
	}).(ContainerConfigurationPtrOutput)
}

func (o ContainerConfigurationOutput) ContainerImageNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContainerConfiguration) []string { return v.ContainerImageNames }).(pulumi.StringArrayOutput)
}

func (o ContainerConfigurationOutput) ContainerRegistries() ContainerRegistryArrayOutput {
	return o.ApplyT(func(v ContainerConfiguration) []ContainerRegistry { return v.ContainerRegistries }).(ContainerRegistryArrayOutput)
}

func (o ContainerConfigurationOutput) Type() ContainerTypeOutput {
	return o.ApplyT(func(v ContainerConfiguration) ContainerType { return v.Type }).(ContainerTypeOutput)
}

type ContainerConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ContainerConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerConfiguration)(nil)).Elem()
}

func (o ContainerConfigurationPtrOutput) ToContainerConfigurationPtrOutput() ContainerConfigurationPtrOutput {
	return o
}

func (o ContainerConfigurationPtrOutput) ToContainerConfigurationPtrOutputWithContext(ctx context.Context) ContainerConfigurationPtrOutput {
	return o
}

func (o ContainerConfigurationPtrOutput) Elem() ContainerConfigurationOutput {
	return o.ApplyT(func(v *ContainerConfiguration) ContainerConfiguration {
		if v != nil {
			return *v
		}
		var ret ContainerConfiguration
		return ret
	}).(ContainerConfigurationOutput)
}

func (o ContainerConfigurationPtrOutput) ContainerImageNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.ContainerImageNames
	}).(pulumi.StringArrayOutput)
}

func (o ContainerConfigurationPtrOutput) ContainerRegistries() ContainerRegistryArrayOutput {
	return o.ApplyT(func(v *ContainerConfiguration) []ContainerRegistry {
		if v == nil {
			return nil
		}
		return v.ContainerRegistries
	}).(ContainerRegistryArrayOutput)
}

func (o ContainerConfigurationPtrOutput) Type() ContainerTypePtrOutput {
	return o.ApplyT(func(v *ContainerConfiguration) *ContainerType {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(ContainerTypePtrOutput)
}

type ContainerConfigurationResponse struct {
	ContainerImageNames []string                    `pulumi:"containerImageNames"`
	ContainerRegistries []ContainerRegistryResponse `pulumi:"containerRegistries"`
	Type                string                      `pulumi:"type"`
}

type ContainerConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ContainerConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerConfigurationResponse)(nil)).Elem()
}

func (o ContainerConfigurationResponseOutput) ToContainerConfigurationResponseOutput() ContainerConfigurationResponseOutput {
	return o
}

func (o ContainerConfigurationResponseOutput) ToContainerConfigurationResponseOutputWithContext(ctx context.Context) ContainerConfigurationResponseOutput {
	return o
}

func (o ContainerConfigurationResponseOutput) ContainerImageNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContainerConfigurationResponse) []string { return v.ContainerImageNames }).(pulumi.StringArrayOutput)
}

func (o ContainerConfigurationResponseOutput) ContainerRegistries() ContainerRegistryResponseArrayOutput {
	return o.ApplyT(func(v ContainerConfigurationResponse) []ContainerRegistryResponse { return v.ContainerRegistries }).(ContainerRegistryResponseArrayOutput)
}

func (o ContainerConfigurationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerConfigurationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ContainerConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerConfigurationResponse)(nil)).Elem()
}

func (o ContainerConfigurationResponsePtrOutput) ToContainerConfigurationResponsePtrOutput() ContainerConfigurationResponsePtrOutput {
	return o
}

func (o ContainerConfigurationResponsePtrOutput) ToContainerConfigurationResponsePtrOutputWithContext(ctx context.Context) ContainerConfigurationResponsePtrOutput {
	return o
}

func (o ContainerConfigurationResponsePtrOutput) Elem() ContainerConfigurationResponseOutput {
	return o.ApplyT(func(v *ContainerConfigurationResponse) ContainerConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ContainerConfigurationResponse
		return ret
	}).(ContainerConfigurationResponseOutput)
}

func (o ContainerConfigurationResponsePtrOutput) ContainerImageNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.ContainerImageNames
	}).(pulumi.StringArrayOutput)
}

func (o ContainerConfigurationResponsePtrOutput) ContainerRegistries() ContainerRegistryResponseArrayOutput {
	return o.ApplyT(func(v *ContainerConfigurationResponse) []ContainerRegistryResponse {
		if v == nil {
			return nil
		}
		return v.ContainerRegistries
	}).(ContainerRegistryResponseArrayOutput)
}

func (o ContainerConfigurationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ContainerRegistry struct {
	Password       string  `pulumi:"password"`
	RegistryServer *string `pulumi:"registryServer"`
	UserName       string  `pulumi:"userName"`
}





type ContainerRegistryInput interface {
	pulumi.Input

	ToContainerRegistryOutput() ContainerRegistryOutput
	ToContainerRegistryOutputWithContext(context.Context) ContainerRegistryOutput
}

type ContainerRegistryArgs struct {
	Password       pulumi.StringInput    `pulumi:"password"`
	RegistryServer pulumi.StringPtrInput `pulumi:"registryServer"`
	UserName       pulumi.StringInput    `pulumi:"userName"`
}

func (ContainerRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerRegistry)(nil)).Elem()
}

func (i ContainerRegistryArgs) ToContainerRegistryOutput() ContainerRegistryOutput {
	return i.ToContainerRegistryOutputWithContext(context.Background())
}

func (i ContainerRegistryArgs) ToContainerRegistryOutputWithContext(ctx context.Context) ContainerRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryOutput)
}

func (i ContainerRegistryArgs) ToContainerRegistryPtrOutput() ContainerRegistryPtrOutput {
	return i.ToContainerRegistryPtrOutputWithContext(context.Background())
}

func (i ContainerRegistryArgs) ToContainerRegistryPtrOutputWithContext(ctx context.Context) ContainerRegistryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryOutput).ToContainerRegistryPtrOutputWithContext(ctx)
}









type ContainerRegistryPtrInput interface {
	pulumi.Input

	ToContainerRegistryPtrOutput() ContainerRegistryPtrOutput
	ToContainerRegistryPtrOutputWithContext(context.Context) ContainerRegistryPtrOutput
}

type containerRegistryPtrType ContainerRegistryArgs

func ContainerRegistryPtr(v *ContainerRegistryArgs) ContainerRegistryPtrInput {
	return (*containerRegistryPtrType)(v)
}

func (*containerRegistryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistry)(nil)).Elem()
}

func (i *containerRegistryPtrType) ToContainerRegistryPtrOutput() ContainerRegistryPtrOutput {
	return i.ToContainerRegistryPtrOutputWithContext(context.Background())
}

func (i *containerRegistryPtrType) ToContainerRegistryPtrOutputWithContext(ctx context.Context) ContainerRegistryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryPtrOutput)
}





type ContainerRegistryArrayInput interface {
	pulumi.Input

	ToContainerRegistryArrayOutput() ContainerRegistryArrayOutput
	ToContainerRegistryArrayOutputWithContext(context.Context) ContainerRegistryArrayOutput
}

type ContainerRegistryArray []ContainerRegistryInput

func (ContainerRegistryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerRegistry)(nil)).Elem()
}

func (i ContainerRegistryArray) ToContainerRegistryArrayOutput() ContainerRegistryArrayOutput {
	return i.ToContainerRegistryArrayOutputWithContext(context.Background())
}

func (i ContainerRegistryArray) ToContainerRegistryArrayOutputWithContext(ctx context.Context) ContainerRegistryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryArrayOutput)
}

type ContainerRegistryOutput struct{ *pulumi.OutputState }

func (ContainerRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerRegistry)(nil)).Elem()
}

func (o ContainerRegistryOutput) ToContainerRegistryOutput() ContainerRegistryOutput {
	return o
}

func (o ContainerRegistryOutput) ToContainerRegistryOutputWithContext(ctx context.Context) ContainerRegistryOutput {
	return o
}

func (o ContainerRegistryOutput) ToContainerRegistryPtrOutput() ContainerRegistryPtrOutput {
	return o.ToContainerRegistryPtrOutputWithContext(context.Background())
}

func (o ContainerRegistryOutput) ToContainerRegistryPtrOutputWithContext(ctx context.Context) ContainerRegistryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerRegistry) *ContainerRegistry {
		return &v
	}).(ContainerRegistryPtrOutput)
}

func (o ContainerRegistryOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerRegistry) string { return v.Password }).(pulumi.StringOutput)
}

func (o ContainerRegistryOutput) RegistryServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerRegistry) *string { return v.RegistryServer }).(pulumi.StringPtrOutput)
}

func (o ContainerRegistryOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerRegistry) string { return v.UserName }).(pulumi.StringOutput)
}

type ContainerRegistryPtrOutput struct{ *pulumi.OutputState }

func (ContainerRegistryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistry)(nil)).Elem()
}

func (o ContainerRegistryPtrOutput) ToContainerRegistryPtrOutput() ContainerRegistryPtrOutput {
	return o
}

func (o ContainerRegistryPtrOutput) ToContainerRegistryPtrOutputWithContext(ctx context.Context) ContainerRegistryPtrOutput {
	return o
}

func (o ContainerRegistryPtrOutput) Elem() ContainerRegistryOutput {
	return o.ApplyT(func(v *ContainerRegistry) ContainerRegistry {
		if v != nil {
			return *v
		}
		var ret ContainerRegistry
		return ret
	}).(ContainerRegistryOutput)
}

func (o ContainerRegistryPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRegistry) *string {
		if v == nil {
			return nil
		}
		return &v.Password
	}).(pulumi.StringPtrOutput)
}

func (o ContainerRegistryPtrOutput) RegistryServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRegistry) *string {
		if v == nil {
			return nil
		}
		return v.RegistryServer
	}).(pulumi.StringPtrOutput)
}

func (o ContainerRegistryPtrOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRegistry) *string {
		if v == nil {
			return nil
		}
		return &v.UserName
	}).(pulumi.StringPtrOutput)
}

type ContainerRegistryArrayOutput struct{ *pulumi.OutputState }

func (ContainerRegistryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerRegistry)(nil)).Elem()
}

func (o ContainerRegistryArrayOutput) ToContainerRegistryArrayOutput() ContainerRegistryArrayOutput {
	return o
}

func (o ContainerRegistryArrayOutput) ToContainerRegistryArrayOutputWithContext(ctx context.Context) ContainerRegistryArrayOutput {
	return o
}

func (o ContainerRegistryArrayOutput) Index(i pulumi.IntInput) ContainerRegistryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerRegistry {
		return vs[0].([]ContainerRegistry)[vs[1].(int)]
	}).(ContainerRegistryOutput)
}

type ContainerRegistryResponse struct {
	Password       string  `pulumi:"password"`
	RegistryServer *string `pulumi:"registryServer"`
	UserName       string  `pulumi:"userName"`
}

type ContainerRegistryResponseOutput struct{ *pulumi.OutputState }

func (ContainerRegistryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerRegistryResponse)(nil)).Elem()
}

func (o ContainerRegistryResponseOutput) ToContainerRegistryResponseOutput() ContainerRegistryResponseOutput {
	return o
}

func (o ContainerRegistryResponseOutput) ToContainerRegistryResponseOutputWithContext(ctx context.Context) ContainerRegistryResponseOutput {
	return o
}

func (o ContainerRegistryResponseOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerRegistryResponse) string { return v.Password }).(pulumi.StringOutput)
}

func (o ContainerRegistryResponseOutput) RegistryServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerRegistryResponse) *string { return v.RegistryServer }).(pulumi.StringPtrOutput)
}

func (o ContainerRegistryResponseOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerRegistryResponse) string { return v.UserName }).(pulumi.StringOutput)
}

type ContainerRegistryResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerRegistryResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryResponse)(nil)).Elem()
}

func (o ContainerRegistryResponsePtrOutput) ToContainerRegistryResponsePtrOutput() ContainerRegistryResponsePtrOutput {
	return o
}

func (o ContainerRegistryResponsePtrOutput) ToContainerRegistryResponsePtrOutputWithContext(ctx context.Context) ContainerRegistryResponsePtrOutput {
	return o
}

func (o ContainerRegistryResponsePtrOutput) Elem() ContainerRegistryResponseOutput {
	return o.ApplyT(func(v *ContainerRegistryResponse) ContainerRegistryResponse {
		if v != nil {
			return *v
		}
		var ret ContainerRegistryResponse
		return ret
	}).(ContainerRegistryResponseOutput)
}

func (o ContainerRegistryResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRegistryResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Password
	}).(pulumi.StringPtrOutput)
}

func (o ContainerRegistryResponsePtrOutput) RegistryServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRegistryResponse) *string {
		if v == nil {
			return nil
		}
		return v.RegistryServer
	}).(pulumi.StringPtrOutput)
}

func (o ContainerRegistryResponsePtrOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRegistryResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UserName
	}).(pulumi.StringPtrOutput)
}

type ContainerRegistryResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerRegistryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerRegistryResponse)(nil)).Elem()
}

func (o ContainerRegistryResponseArrayOutput) ToContainerRegistryResponseArrayOutput() ContainerRegistryResponseArrayOutput {
	return o
}

func (o ContainerRegistryResponseArrayOutput) ToContainerRegistryResponseArrayOutputWithContext(ctx context.Context) ContainerRegistryResponseArrayOutput {
	return o
}

func (o ContainerRegistryResponseArrayOutput) Index(i pulumi.IntInput) ContainerRegistryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerRegistryResponse {
		return vs[0].([]ContainerRegistryResponse)[vs[1].(int)]
	}).(ContainerRegistryResponseOutput)
}

type DataDisk struct {
	Caching            *CachingType        `pulumi:"caching"`
	DiskSizeGB         int                 `pulumi:"diskSizeGB"`
	Lun                int                 `pulumi:"lun"`
	StorageAccountType *StorageAccountType `pulumi:"storageAccountType"`
}





type DataDiskInput interface {
	pulumi.Input

	ToDataDiskOutput() DataDiskOutput
	ToDataDiskOutputWithContext(context.Context) DataDiskOutput
}

type DataDiskArgs struct {
	Caching            CachingTypePtrInput        `pulumi:"caching"`
	DiskSizeGB         pulumi.IntInput            `pulumi:"diskSizeGB"`
	Lun                pulumi.IntInput            `pulumi:"lun"`
	StorageAccountType StorageAccountTypePtrInput `pulumi:"storageAccountType"`
}

func (DataDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDisk)(nil)).Elem()
}

func (i DataDiskArgs) ToDataDiskOutput() DataDiskOutput {
	return i.ToDataDiskOutputWithContext(context.Background())
}

func (i DataDiskArgs) ToDataDiskOutputWithContext(ctx context.Context) DataDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskOutput)
}





type DataDiskArrayInput interface {
	pulumi.Input

	ToDataDiskArrayOutput() DataDiskArrayOutput
	ToDataDiskArrayOutputWithContext(context.Context) DataDiskArrayOutput
}

type DataDiskArray []DataDiskInput

func (DataDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDisk)(nil)).Elem()
}

func (i DataDiskArray) ToDataDiskArrayOutput() DataDiskArrayOutput {
	return i.ToDataDiskArrayOutputWithContext(context.Background())
}

func (i DataDiskArray) ToDataDiskArrayOutputWithContext(ctx context.Context) DataDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskArrayOutput)
}

type DataDiskOutput struct{ *pulumi.OutputState }

func (DataDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDisk)(nil)).Elem()
}

func (o DataDiskOutput) ToDataDiskOutput() DataDiskOutput {
	return o
}

func (o DataDiskOutput) ToDataDiskOutputWithContext(ctx context.Context) DataDiskOutput {
	return o
}

func (o DataDiskOutput) Caching() CachingTypePtrOutput {
	return o.ApplyT(func(v DataDisk) *CachingType { return v.Caching }).(CachingTypePtrOutput)
}

func (o DataDiskOutput) DiskSizeGB() pulumi.IntOutput {
	return o.ApplyT(func(v DataDisk) int { return v.DiskSizeGB }).(pulumi.IntOutput)
}

func (o DataDiskOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v DataDisk) int { return v.Lun }).(pulumi.IntOutput)
}

func (o DataDiskOutput) StorageAccountType() StorageAccountTypePtrOutput {
	return o.ApplyT(func(v DataDisk) *StorageAccountType { return v.StorageAccountType }).(StorageAccountTypePtrOutput)
}

type DataDiskArrayOutput struct{ *pulumi.OutputState }

func (DataDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDisk)(nil)).Elem()
}

func (o DataDiskArrayOutput) ToDataDiskArrayOutput() DataDiskArrayOutput {
	return o
}

func (o DataDiskArrayOutput) ToDataDiskArrayOutputWithContext(ctx context.Context) DataDiskArrayOutput {
	return o
}

func (o DataDiskArrayOutput) Index(i pulumi.IntInput) DataDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDisk {
		return vs[0].([]DataDisk)[vs[1].(int)]
	}).(DataDiskOutput)
}

type DataDiskResponse struct {
	Caching            *string `pulumi:"caching"`
	DiskSizeGB         int     `pulumi:"diskSizeGB"`
	Lun                int     `pulumi:"lun"`
	StorageAccountType *string `pulumi:"storageAccountType"`
}

type DataDiskResponseOutput struct{ *pulumi.OutputState }

func (DataDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDiskResponse)(nil)).Elem()
}

func (o DataDiskResponseOutput) ToDataDiskResponseOutput() DataDiskResponseOutput {
	return o
}

func (o DataDiskResponseOutput) ToDataDiskResponseOutputWithContext(ctx context.Context) DataDiskResponseOutput {
	return o
}

func (o DataDiskResponseOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *string { return v.Caching }).(pulumi.StringPtrOutput)
}

func (o DataDiskResponseOutput) DiskSizeGB() pulumi.IntOutput {
	return o.ApplyT(func(v DataDiskResponse) int { return v.DiskSizeGB }).(pulumi.IntOutput)
}

func (o DataDiskResponseOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v DataDiskResponse) int { return v.Lun }).(pulumi.IntOutput)
}

func (o DataDiskResponseOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type DataDiskResponseArrayOutput struct{ *pulumi.OutputState }

func (DataDiskResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDiskResponse)(nil)).Elem()
}

func (o DataDiskResponseArrayOutput) ToDataDiskResponseArrayOutput() DataDiskResponseArrayOutput {
	return o
}

func (o DataDiskResponseArrayOutput) ToDataDiskResponseArrayOutputWithContext(ctx context.Context) DataDiskResponseArrayOutput {
	return o
}

func (o DataDiskResponseArrayOutput) Index(i pulumi.IntInput) DataDiskResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDiskResponse {
		return vs[0].([]DataDiskResponse)[vs[1].(int)]
	}).(DataDiskResponseOutput)
}

type DeleteCertificateErrorResponse struct {
	Code    string                           `pulumi:"code"`
	Details []DeleteCertificateErrorResponse `pulumi:"details"`
	Message string                           `pulumi:"message"`
	Target  *string                          `pulumi:"target"`
}

type DeleteCertificateErrorResponseOutput struct{ *pulumi.OutputState }

func (DeleteCertificateErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeleteCertificateErrorResponse)(nil)).Elem()
}

func (o DeleteCertificateErrorResponseOutput) ToDeleteCertificateErrorResponseOutput() DeleteCertificateErrorResponseOutput {
	return o
}

func (o DeleteCertificateErrorResponseOutput) ToDeleteCertificateErrorResponseOutputWithContext(ctx context.Context) DeleteCertificateErrorResponseOutput {
	return o
}

func (o DeleteCertificateErrorResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v DeleteCertificateErrorResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o DeleteCertificateErrorResponseOutput) Details() DeleteCertificateErrorResponseArrayOutput {
	return o.ApplyT(func(v DeleteCertificateErrorResponse) []DeleteCertificateErrorResponse { return v.Details }).(DeleteCertificateErrorResponseArrayOutput)
}

func (o DeleteCertificateErrorResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v DeleteCertificateErrorResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o DeleteCertificateErrorResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeleteCertificateErrorResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type DeleteCertificateErrorResponseArrayOutput struct{ *pulumi.OutputState }

func (DeleteCertificateErrorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeleteCertificateErrorResponse)(nil)).Elem()
}

func (o DeleteCertificateErrorResponseArrayOutput) ToDeleteCertificateErrorResponseArrayOutput() DeleteCertificateErrorResponseArrayOutput {
	return o
}

func (o DeleteCertificateErrorResponseArrayOutput) ToDeleteCertificateErrorResponseArrayOutputWithContext(ctx context.Context) DeleteCertificateErrorResponseArrayOutput {
	return o
}

func (o DeleteCertificateErrorResponseArrayOutput) Index(i pulumi.IntInput) DeleteCertificateErrorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeleteCertificateErrorResponse {
		return vs[0].([]DeleteCertificateErrorResponse)[vs[1].(int)]
	}).(DeleteCertificateErrorResponseOutput)
}

type DeploymentConfiguration struct {
	CloudServiceConfiguration   *CloudServiceConfiguration   `pulumi:"cloudServiceConfiguration"`
	VirtualMachineConfiguration *VirtualMachineConfiguration `pulumi:"virtualMachineConfiguration"`
}





type DeploymentConfigurationInput interface {
	pulumi.Input

	ToDeploymentConfigurationOutput() DeploymentConfigurationOutput
	ToDeploymentConfigurationOutputWithContext(context.Context) DeploymentConfigurationOutput
}

type DeploymentConfigurationArgs struct {
	CloudServiceConfiguration   CloudServiceConfigurationPtrInput   `pulumi:"cloudServiceConfiguration"`
	VirtualMachineConfiguration VirtualMachineConfigurationPtrInput `pulumi:"virtualMachineConfiguration"`
}

func (DeploymentConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentConfiguration)(nil)).Elem()
}

func (i DeploymentConfigurationArgs) ToDeploymentConfigurationOutput() DeploymentConfigurationOutput {
	return i.ToDeploymentConfigurationOutputWithContext(context.Background())
}

func (i DeploymentConfigurationArgs) ToDeploymentConfigurationOutputWithContext(ctx context.Context) DeploymentConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentConfigurationOutput)
}

func (i DeploymentConfigurationArgs) ToDeploymentConfigurationPtrOutput() DeploymentConfigurationPtrOutput {
	return i.ToDeploymentConfigurationPtrOutputWithContext(context.Background())
}

func (i DeploymentConfigurationArgs) ToDeploymentConfigurationPtrOutputWithContext(ctx context.Context) DeploymentConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentConfigurationOutput).ToDeploymentConfigurationPtrOutputWithContext(ctx)
}









type DeploymentConfigurationPtrInput interface {
	pulumi.Input

	ToDeploymentConfigurationPtrOutput() DeploymentConfigurationPtrOutput
	ToDeploymentConfigurationPtrOutputWithContext(context.Context) DeploymentConfigurationPtrOutput
}

type deploymentConfigurationPtrType DeploymentConfigurationArgs

func DeploymentConfigurationPtr(v *DeploymentConfigurationArgs) DeploymentConfigurationPtrInput {
	return (*deploymentConfigurationPtrType)(v)
}

func (*deploymentConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentConfiguration)(nil)).Elem()
}

func (i *deploymentConfigurationPtrType) ToDeploymentConfigurationPtrOutput() DeploymentConfigurationPtrOutput {
	return i.ToDeploymentConfigurationPtrOutputWithContext(context.Background())
}

func (i *deploymentConfigurationPtrType) ToDeploymentConfigurationPtrOutputWithContext(ctx context.Context) DeploymentConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentConfigurationPtrOutput)
}

type DeploymentConfigurationOutput struct{ *pulumi.OutputState }

func (DeploymentConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentConfiguration)(nil)).Elem()
}

func (o DeploymentConfigurationOutput) ToDeploymentConfigurationOutput() DeploymentConfigurationOutput {
	return o
}

func (o DeploymentConfigurationOutput) ToDeploymentConfigurationOutputWithContext(ctx context.Context) DeploymentConfigurationOutput {
	return o
}

func (o DeploymentConfigurationOutput) ToDeploymentConfigurationPtrOutput() DeploymentConfigurationPtrOutput {
	return o.ToDeploymentConfigurationPtrOutputWithContext(context.Background())
}

func (o DeploymentConfigurationOutput) ToDeploymentConfigurationPtrOutputWithContext(ctx context.Context) DeploymentConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentConfiguration) *DeploymentConfiguration {
		return &v
	}).(DeploymentConfigurationPtrOutput)
}

func (o DeploymentConfigurationOutput) CloudServiceConfiguration() CloudServiceConfigurationPtrOutput {
	return o.ApplyT(func(v DeploymentConfiguration) *CloudServiceConfiguration { return v.CloudServiceConfiguration }).(CloudServiceConfigurationPtrOutput)
}

func (o DeploymentConfigurationOutput) VirtualMachineConfiguration() VirtualMachineConfigurationPtrOutput {
	return o.ApplyT(func(v DeploymentConfiguration) *VirtualMachineConfiguration { return v.VirtualMachineConfiguration }).(VirtualMachineConfigurationPtrOutput)
}

type DeploymentConfigurationPtrOutput struct{ *pulumi.OutputState }

func (DeploymentConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentConfiguration)(nil)).Elem()
}

func (o DeploymentConfigurationPtrOutput) ToDeploymentConfigurationPtrOutput() DeploymentConfigurationPtrOutput {
	return o
}

func (o DeploymentConfigurationPtrOutput) ToDeploymentConfigurationPtrOutputWithContext(ctx context.Context) DeploymentConfigurationPtrOutput {
	return o
}

func (o DeploymentConfigurationPtrOutput) Elem() DeploymentConfigurationOutput {
	return o.ApplyT(func(v *DeploymentConfiguration) DeploymentConfiguration {
		if v != nil {
			return *v
		}
		var ret DeploymentConfiguration
		return ret
	}).(DeploymentConfigurationOutput)
}

func (o DeploymentConfigurationPtrOutput) CloudServiceConfiguration() CloudServiceConfigurationPtrOutput {
	return o.ApplyT(func(v *DeploymentConfiguration) *CloudServiceConfiguration {
		if v == nil {
			return nil
		}
		return v.CloudServiceConfiguration
	}).(CloudServiceConfigurationPtrOutput)
}

func (o DeploymentConfigurationPtrOutput) VirtualMachineConfiguration() VirtualMachineConfigurationPtrOutput {
	return o.ApplyT(func(v *DeploymentConfiguration) *VirtualMachineConfiguration {
		if v == nil {
			return nil
		}
		return v.VirtualMachineConfiguration
	}).(VirtualMachineConfigurationPtrOutput)
}

type DeploymentConfigurationResponse struct {
	CloudServiceConfiguration   *CloudServiceConfigurationResponse   `pulumi:"cloudServiceConfiguration"`
	VirtualMachineConfiguration *VirtualMachineConfigurationResponse `pulumi:"virtualMachineConfiguration"`
}

type DeploymentConfigurationResponseOutput struct{ *pulumi.OutputState }

func (DeploymentConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentConfigurationResponse)(nil)).Elem()
}

func (o DeploymentConfigurationResponseOutput) ToDeploymentConfigurationResponseOutput() DeploymentConfigurationResponseOutput {
	return o
}

func (o DeploymentConfigurationResponseOutput) ToDeploymentConfigurationResponseOutputWithContext(ctx context.Context) DeploymentConfigurationResponseOutput {
	return o
}

func (o DeploymentConfigurationResponseOutput) CloudServiceConfiguration() CloudServiceConfigurationResponsePtrOutput {
	return o.ApplyT(func(v DeploymentConfigurationResponse) *CloudServiceConfigurationResponse {
		return v.CloudServiceConfiguration
	}).(CloudServiceConfigurationResponsePtrOutput)
}

func (o DeploymentConfigurationResponseOutput) VirtualMachineConfiguration() VirtualMachineConfigurationResponsePtrOutput {
	return o.ApplyT(func(v DeploymentConfigurationResponse) *VirtualMachineConfigurationResponse {
		return v.VirtualMachineConfiguration
	}).(VirtualMachineConfigurationResponsePtrOutput)
}

type DeploymentConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (DeploymentConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentConfigurationResponse)(nil)).Elem()
}

func (o DeploymentConfigurationResponsePtrOutput) ToDeploymentConfigurationResponsePtrOutput() DeploymentConfigurationResponsePtrOutput {
	return o
}

func (o DeploymentConfigurationResponsePtrOutput) ToDeploymentConfigurationResponsePtrOutputWithContext(ctx context.Context) DeploymentConfigurationResponsePtrOutput {
	return o
}

func (o DeploymentConfigurationResponsePtrOutput) Elem() DeploymentConfigurationResponseOutput {
	return o.ApplyT(func(v *DeploymentConfigurationResponse) DeploymentConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret DeploymentConfigurationResponse
		return ret
	}).(DeploymentConfigurationResponseOutput)
}

func (o DeploymentConfigurationResponsePtrOutput) CloudServiceConfiguration() CloudServiceConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentConfigurationResponse) *CloudServiceConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.CloudServiceConfiguration
	}).(CloudServiceConfigurationResponsePtrOutput)
}

func (o DeploymentConfigurationResponsePtrOutput) VirtualMachineConfiguration() VirtualMachineConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentConfigurationResponse) *VirtualMachineConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.VirtualMachineConfiguration
	}).(VirtualMachineConfigurationResponsePtrOutput)
}

type EnvironmentSetting struct {
	Name  string  `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type EnvironmentSettingInput interface {
	pulumi.Input

	ToEnvironmentSettingOutput() EnvironmentSettingOutput
	ToEnvironmentSettingOutputWithContext(context.Context) EnvironmentSettingOutput
}

type EnvironmentSettingArgs struct {
	Name  pulumi.StringInput    `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (EnvironmentSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentSetting)(nil)).Elem()
}

func (i EnvironmentSettingArgs) ToEnvironmentSettingOutput() EnvironmentSettingOutput {
	return i.ToEnvironmentSettingOutputWithContext(context.Background())
}

func (i EnvironmentSettingArgs) ToEnvironmentSettingOutputWithContext(ctx context.Context) EnvironmentSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentSettingOutput)
}





type EnvironmentSettingArrayInput interface {
	pulumi.Input

	ToEnvironmentSettingArrayOutput() EnvironmentSettingArrayOutput
	ToEnvironmentSettingArrayOutputWithContext(context.Context) EnvironmentSettingArrayOutput
}

type EnvironmentSettingArray []EnvironmentSettingInput

func (EnvironmentSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentSetting)(nil)).Elem()
}

func (i EnvironmentSettingArray) ToEnvironmentSettingArrayOutput() EnvironmentSettingArrayOutput {
	return i.ToEnvironmentSettingArrayOutputWithContext(context.Background())
}

func (i EnvironmentSettingArray) ToEnvironmentSettingArrayOutputWithContext(ctx context.Context) EnvironmentSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentSettingArrayOutput)
}

type EnvironmentSettingOutput struct{ *pulumi.OutputState }

func (EnvironmentSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentSetting)(nil)).Elem()
}

func (o EnvironmentSettingOutput) ToEnvironmentSettingOutput() EnvironmentSettingOutput {
	return o
}

func (o EnvironmentSettingOutput) ToEnvironmentSettingOutputWithContext(ctx context.Context) EnvironmentSettingOutput {
	return o
}

func (o EnvironmentSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EnvironmentSetting) string { return v.Name }).(pulumi.StringOutput)
}

func (o EnvironmentSettingOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentSetting) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type EnvironmentSettingArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentSetting)(nil)).Elem()
}

func (o EnvironmentSettingArrayOutput) ToEnvironmentSettingArrayOutput() EnvironmentSettingArrayOutput {
	return o
}

func (o EnvironmentSettingArrayOutput) ToEnvironmentSettingArrayOutputWithContext(ctx context.Context) EnvironmentSettingArrayOutput {
	return o
}

func (o EnvironmentSettingArrayOutput) Index(i pulumi.IntInput) EnvironmentSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentSetting {
		return vs[0].([]EnvironmentSetting)[vs[1].(int)]
	}).(EnvironmentSettingOutput)
}

type EnvironmentSettingResponse struct {
	Name  string  `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type EnvironmentSettingResponseOutput struct{ *pulumi.OutputState }

func (EnvironmentSettingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentSettingResponse)(nil)).Elem()
}

func (o EnvironmentSettingResponseOutput) ToEnvironmentSettingResponseOutput() EnvironmentSettingResponseOutput {
	return o
}

func (o EnvironmentSettingResponseOutput) ToEnvironmentSettingResponseOutputWithContext(ctx context.Context) EnvironmentSettingResponseOutput {
	return o
}

func (o EnvironmentSettingResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EnvironmentSettingResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o EnvironmentSettingResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentSettingResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type EnvironmentSettingResponseArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentSettingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentSettingResponse)(nil)).Elem()
}

func (o EnvironmentSettingResponseArrayOutput) ToEnvironmentSettingResponseArrayOutput() EnvironmentSettingResponseArrayOutput {
	return o
}

func (o EnvironmentSettingResponseArrayOutput) ToEnvironmentSettingResponseArrayOutputWithContext(ctx context.Context) EnvironmentSettingResponseArrayOutput {
	return o
}

func (o EnvironmentSettingResponseArrayOutput) Index(i pulumi.IntInput) EnvironmentSettingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentSettingResponse {
		return vs[0].([]EnvironmentSettingResponse)[vs[1].(int)]
	}).(EnvironmentSettingResponseOutput)
}

type FixedScaleSettings struct {
	NodeDeallocationOption *ComputeNodeDeallocationOption `pulumi:"nodeDeallocationOption"`
	ResizeTimeout          *string                        `pulumi:"resizeTimeout"`
	TargetDedicatedNodes   *int                           `pulumi:"targetDedicatedNodes"`
	TargetLowPriorityNodes *int                           `pulumi:"targetLowPriorityNodes"`
}





type FixedScaleSettingsInput interface {
	pulumi.Input

	ToFixedScaleSettingsOutput() FixedScaleSettingsOutput
	ToFixedScaleSettingsOutputWithContext(context.Context) FixedScaleSettingsOutput
}

type FixedScaleSettingsArgs struct {
	NodeDeallocationOption ComputeNodeDeallocationOptionPtrInput `pulumi:"nodeDeallocationOption"`
	ResizeTimeout          pulumi.StringPtrInput                 `pulumi:"resizeTimeout"`
	TargetDedicatedNodes   pulumi.IntPtrInput                    `pulumi:"targetDedicatedNodes"`
	TargetLowPriorityNodes pulumi.IntPtrInput                    `pulumi:"targetLowPriorityNodes"`
}

func (FixedScaleSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FixedScaleSettings)(nil)).Elem()
}

func (i FixedScaleSettingsArgs) ToFixedScaleSettingsOutput() FixedScaleSettingsOutput {
	return i.ToFixedScaleSettingsOutputWithContext(context.Background())
}

func (i FixedScaleSettingsArgs) ToFixedScaleSettingsOutputWithContext(ctx context.Context) FixedScaleSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FixedScaleSettingsOutput)
}

func (i FixedScaleSettingsArgs) ToFixedScaleSettingsPtrOutput() FixedScaleSettingsPtrOutput {
	return i.ToFixedScaleSettingsPtrOutputWithContext(context.Background())
}

func (i FixedScaleSettingsArgs) ToFixedScaleSettingsPtrOutputWithContext(ctx context.Context) FixedScaleSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FixedScaleSettingsOutput).ToFixedScaleSettingsPtrOutputWithContext(ctx)
}









type FixedScaleSettingsPtrInput interface {
	pulumi.Input

	ToFixedScaleSettingsPtrOutput() FixedScaleSettingsPtrOutput
	ToFixedScaleSettingsPtrOutputWithContext(context.Context) FixedScaleSettingsPtrOutput
}

type fixedScaleSettingsPtrType FixedScaleSettingsArgs

func FixedScaleSettingsPtr(v *FixedScaleSettingsArgs) FixedScaleSettingsPtrInput {
	return (*fixedScaleSettingsPtrType)(v)
}

func (*fixedScaleSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FixedScaleSettings)(nil)).Elem()
}

func (i *fixedScaleSettingsPtrType) ToFixedScaleSettingsPtrOutput() FixedScaleSettingsPtrOutput {
	return i.ToFixedScaleSettingsPtrOutputWithContext(context.Background())
}

func (i *fixedScaleSettingsPtrType) ToFixedScaleSettingsPtrOutputWithContext(ctx context.Context) FixedScaleSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FixedScaleSettingsPtrOutput)
}

type FixedScaleSettingsOutput struct{ *pulumi.OutputState }

func (FixedScaleSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FixedScaleSettings)(nil)).Elem()
}

func (o FixedScaleSettingsOutput) ToFixedScaleSettingsOutput() FixedScaleSettingsOutput {
	return o
}

func (o FixedScaleSettingsOutput) ToFixedScaleSettingsOutputWithContext(ctx context.Context) FixedScaleSettingsOutput {
	return o
}

func (o FixedScaleSettingsOutput) ToFixedScaleSettingsPtrOutput() FixedScaleSettingsPtrOutput {
	return o.ToFixedScaleSettingsPtrOutputWithContext(context.Background())
}

func (o FixedScaleSettingsOutput) ToFixedScaleSettingsPtrOutputWithContext(ctx context.Context) FixedScaleSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FixedScaleSettings) *FixedScaleSettings {
		return &v
	}).(FixedScaleSettingsPtrOutput)
}

func (o FixedScaleSettingsOutput) NodeDeallocationOption() ComputeNodeDeallocationOptionPtrOutput {
	return o.ApplyT(func(v FixedScaleSettings) *ComputeNodeDeallocationOption { return v.NodeDeallocationOption }).(ComputeNodeDeallocationOptionPtrOutput)
}

func (o FixedScaleSettingsOutput) ResizeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FixedScaleSettings) *string { return v.ResizeTimeout }).(pulumi.StringPtrOutput)
}

func (o FixedScaleSettingsOutput) TargetDedicatedNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FixedScaleSettings) *int { return v.TargetDedicatedNodes }).(pulumi.IntPtrOutput)
}

func (o FixedScaleSettingsOutput) TargetLowPriorityNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FixedScaleSettings) *int { return v.TargetLowPriorityNodes }).(pulumi.IntPtrOutput)
}

type FixedScaleSettingsPtrOutput struct{ *pulumi.OutputState }

func (FixedScaleSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FixedScaleSettings)(nil)).Elem()
}

func (o FixedScaleSettingsPtrOutput) ToFixedScaleSettingsPtrOutput() FixedScaleSettingsPtrOutput {
	return o
}

func (o FixedScaleSettingsPtrOutput) ToFixedScaleSettingsPtrOutputWithContext(ctx context.Context) FixedScaleSettingsPtrOutput {
	return o
}

func (o FixedScaleSettingsPtrOutput) Elem() FixedScaleSettingsOutput {
	return o.ApplyT(func(v *FixedScaleSettings) FixedScaleSettings {
		if v != nil {
			return *v
		}
		var ret FixedScaleSettings
		return ret
	}).(FixedScaleSettingsOutput)
}

func (o FixedScaleSettingsPtrOutput) NodeDeallocationOption() ComputeNodeDeallocationOptionPtrOutput {
	return o.ApplyT(func(v *FixedScaleSettings) *ComputeNodeDeallocationOption {
		if v == nil {
			return nil
		}
		return v.NodeDeallocationOption
	}).(ComputeNodeDeallocationOptionPtrOutput)
}

func (o FixedScaleSettingsPtrOutput) ResizeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FixedScaleSettings) *string {
		if v == nil {
			return nil
		}
		return v.ResizeTimeout
	}).(pulumi.StringPtrOutput)
}

func (o FixedScaleSettingsPtrOutput) TargetDedicatedNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FixedScaleSettings) *int {
		if v == nil {
			return nil
		}
		return v.TargetDedicatedNodes
	}).(pulumi.IntPtrOutput)
}

func (o FixedScaleSettingsPtrOutput) TargetLowPriorityNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FixedScaleSettings) *int {
		if v == nil {
			return nil
		}
		return v.TargetLowPriorityNodes
	}).(pulumi.IntPtrOutput)
}

type FixedScaleSettingsResponse struct {
	NodeDeallocationOption *string `pulumi:"nodeDeallocationOption"`
	ResizeTimeout          *string `pulumi:"resizeTimeout"`
	TargetDedicatedNodes   *int    `pulumi:"targetDedicatedNodes"`
	TargetLowPriorityNodes *int    `pulumi:"targetLowPriorityNodes"`
}

type FixedScaleSettingsResponseOutput struct{ *pulumi.OutputState }

func (FixedScaleSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FixedScaleSettingsResponse)(nil)).Elem()
}

func (o FixedScaleSettingsResponseOutput) ToFixedScaleSettingsResponseOutput() FixedScaleSettingsResponseOutput {
	return o
}

func (o FixedScaleSettingsResponseOutput) ToFixedScaleSettingsResponseOutputWithContext(ctx context.Context) FixedScaleSettingsResponseOutput {
	return o
}

func (o FixedScaleSettingsResponseOutput) NodeDeallocationOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FixedScaleSettingsResponse) *string { return v.NodeDeallocationOption }).(pulumi.StringPtrOutput)
}

func (o FixedScaleSettingsResponseOutput) ResizeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FixedScaleSettingsResponse) *string { return v.ResizeTimeout }).(pulumi.StringPtrOutput)
}

func (o FixedScaleSettingsResponseOutput) TargetDedicatedNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FixedScaleSettingsResponse) *int { return v.TargetDedicatedNodes }).(pulumi.IntPtrOutput)
}

func (o FixedScaleSettingsResponseOutput) TargetLowPriorityNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FixedScaleSettingsResponse) *int { return v.TargetLowPriorityNodes }).(pulumi.IntPtrOutput)
}

type FixedScaleSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (FixedScaleSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FixedScaleSettingsResponse)(nil)).Elem()
}

func (o FixedScaleSettingsResponsePtrOutput) ToFixedScaleSettingsResponsePtrOutput() FixedScaleSettingsResponsePtrOutput {
	return o
}

func (o FixedScaleSettingsResponsePtrOutput) ToFixedScaleSettingsResponsePtrOutputWithContext(ctx context.Context) FixedScaleSettingsResponsePtrOutput {
	return o
}

func (o FixedScaleSettingsResponsePtrOutput) Elem() FixedScaleSettingsResponseOutput {
	return o.ApplyT(func(v *FixedScaleSettingsResponse) FixedScaleSettingsResponse {
		if v != nil {
			return *v
		}
		var ret FixedScaleSettingsResponse
		return ret
	}).(FixedScaleSettingsResponseOutput)
}

func (o FixedScaleSettingsResponsePtrOutput) NodeDeallocationOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FixedScaleSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.NodeDeallocationOption
	}).(pulumi.StringPtrOutput)
}

func (o FixedScaleSettingsResponsePtrOutput) ResizeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FixedScaleSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResizeTimeout
	}).(pulumi.StringPtrOutput)
}

func (o FixedScaleSettingsResponsePtrOutput) TargetDedicatedNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FixedScaleSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.TargetDedicatedNodes
	}).(pulumi.IntPtrOutput)
}

func (o FixedScaleSettingsResponsePtrOutput) TargetLowPriorityNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FixedScaleSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.TargetLowPriorityNodes
	}).(pulumi.IntPtrOutput)
}

type ImageReference struct {
	Id        *string `pulumi:"id"`
	Offer     *string `pulumi:"offer"`
	Publisher *string `pulumi:"publisher"`
	Sku       *string `pulumi:"sku"`
	Version   *string `pulumi:"version"`
}





type ImageReferenceInput interface {
	pulumi.Input

	ToImageReferenceOutput() ImageReferenceOutput
	ToImageReferenceOutputWithContext(context.Context) ImageReferenceOutput
}

type ImageReferenceArgs struct {
	Id        pulumi.StringPtrInput `pulumi:"id"`
	Offer     pulumi.StringPtrInput `pulumi:"offer"`
	Publisher pulumi.StringPtrInput `pulumi:"publisher"`
	Sku       pulumi.StringPtrInput `pulumi:"sku"`
	Version   pulumi.StringPtrInput `pulumi:"version"`
}

func (ImageReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReference)(nil)).Elem()
}

func (i ImageReferenceArgs) ToImageReferenceOutput() ImageReferenceOutput {
	return i.ToImageReferenceOutputWithContext(context.Background())
}

func (i ImageReferenceArgs) ToImageReferenceOutputWithContext(ctx context.Context) ImageReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferenceOutput)
}

func (i ImageReferenceArgs) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return i.ToImageReferencePtrOutputWithContext(context.Background())
}

func (i ImageReferenceArgs) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferenceOutput).ToImageReferencePtrOutputWithContext(ctx)
}









type ImageReferencePtrInput interface {
	pulumi.Input

	ToImageReferencePtrOutput() ImageReferencePtrOutput
	ToImageReferencePtrOutputWithContext(context.Context) ImageReferencePtrOutput
}

type imageReferencePtrType ImageReferenceArgs

func ImageReferencePtr(v *ImageReferenceArgs) ImageReferencePtrInput {
	return (*imageReferencePtrType)(v)
}

func (*imageReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageReference)(nil)).Elem()
}

func (i *imageReferencePtrType) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return i.ToImageReferencePtrOutputWithContext(context.Background())
}

func (i *imageReferencePtrType) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferencePtrOutput)
}

type ImageReferenceOutput struct{ *pulumi.OutputState }

func (ImageReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReference)(nil)).Elem()
}

func (o ImageReferenceOutput) ToImageReferenceOutput() ImageReferenceOutput {
	return o
}

func (o ImageReferenceOutput) ToImageReferenceOutputWithContext(ctx context.Context) ImageReferenceOutput {
	return o
}

func (o ImageReferenceOutput) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return o.ToImageReferencePtrOutputWithContext(context.Background())
}

func (o ImageReferenceOutput) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageReference) *ImageReference {
		return &v
	}).(ImageReferencePtrOutput)
}

func (o ImageReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ImageReferencePtrOutput struct{ *pulumi.OutputState }

func (ImageReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageReference)(nil)).Elem()
}

func (o ImageReferencePtrOutput) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return o
}

func (o ImageReferencePtrOutput) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return o
}

func (o ImageReferencePtrOutput) Elem() ImageReferenceOutput {
	return o.ApplyT(func(v *ImageReference) ImageReference {
		if v != nil {
			return *v
		}
		var ret ImageReference
		return ret
	}).(ImageReferenceOutput)
}

func (o ImageReferencePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ImageReferenceResponse struct {
	Id        *string `pulumi:"id"`
	Offer     *string `pulumi:"offer"`
	Publisher *string `pulumi:"publisher"`
	Sku       *string `pulumi:"sku"`
	Version   *string `pulumi:"version"`
}

type ImageReferenceResponseOutput struct{ *pulumi.OutputState }

func (ImageReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReferenceResponse)(nil)).Elem()
}

func (o ImageReferenceResponseOutput) ToImageReferenceResponseOutput() ImageReferenceResponseOutput {
	return o
}

func (o ImageReferenceResponseOutput) ToImageReferenceResponseOutputWithContext(ctx context.Context) ImageReferenceResponseOutput {
	return o
}

func (o ImageReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ImageReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageReferenceResponse)(nil)).Elem()
}

func (o ImageReferenceResponsePtrOutput) ToImageReferenceResponsePtrOutput() ImageReferenceResponsePtrOutput {
	return o
}

func (o ImageReferenceResponsePtrOutput) ToImageReferenceResponsePtrOutputWithContext(ctx context.Context) ImageReferenceResponsePtrOutput {
	return o
}

func (o ImageReferenceResponsePtrOutput) Elem() ImageReferenceResponseOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) ImageReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ImageReferenceResponse
		return ret
	}).(ImageReferenceResponseOutput)
}

func (o ImageReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type InboundNatPool struct {
	BackendPort               int                        `pulumi:"backendPort"`
	FrontendPortRangeEnd      int                        `pulumi:"frontendPortRangeEnd"`
	FrontendPortRangeStart    int                        `pulumi:"frontendPortRangeStart"`
	Name                      string                     `pulumi:"name"`
	NetworkSecurityGroupRules []NetworkSecurityGroupRule `pulumi:"networkSecurityGroupRules"`
	Protocol                  InboundEndpointProtocol    `pulumi:"protocol"`
}





type InboundNatPoolInput interface {
	pulumi.Input

	ToInboundNatPoolOutput() InboundNatPoolOutput
	ToInboundNatPoolOutputWithContext(context.Context) InboundNatPoolOutput
}

type InboundNatPoolArgs struct {
	BackendPort               pulumi.IntInput                    `pulumi:"backendPort"`
	FrontendPortRangeEnd      pulumi.IntInput                    `pulumi:"frontendPortRangeEnd"`
	FrontendPortRangeStart    pulumi.IntInput                    `pulumi:"frontendPortRangeStart"`
	Name                      pulumi.StringInput                 `pulumi:"name"`
	NetworkSecurityGroupRules NetworkSecurityGroupRuleArrayInput `pulumi:"networkSecurityGroupRules"`
	Protocol                  InboundEndpointProtocolInput       `pulumi:"protocol"`
}

func (InboundNatPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatPool)(nil)).Elem()
}

func (i InboundNatPoolArgs) ToInboundNatPoolOutput() InboundNatPoolOutput {
	return i.ToInboundNatPoolOutputWithContext(context.Background())
}

func (i InboundNatPoolArgs) ToInboundNatPoolOutputWithContext(ctx context.Context) InboundNatPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatPoolOutput)
}





type InboundNatPoolArrayInput interface {
	pulumi.Input

	ToInboundNatPoolArrayOutput() InboundNatPoolArrayOutput
	ToInboundNatPoolArrayOutputWithContext(context.Context) InboundNatPoolArrayOutput
}

type InboundNatPoolArray []InboundNatPoolInput

func (InboundNatPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatPool)(nil)).Elem()
}

func (i InboundNatPoolArray) ToInboundNatPoolArrayOutput() InboundNatPoolArrayOutput {
	return i.ToInboundNatPoolArrayOutputWithContext(context.Background())
}

func (i InboundNatPoolArray) ToInboundNatPoolArrayOutputWithContext(ctx context.Context) InboundNatPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatPoolArrayOutput)
}

type InboundNatPoolOutput struct{ *pulumi.OutputState }

func (InboundNatPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatPool)(nil)).Elem()
}

func (o InboundNatPoolOutput) ToInboundNatPoolOutput() InboundNatPoolOutput {
	return o
}

func (o InboundNatPoolOutput) ToInboundNatPoolOutputWithContext(ctx context.Context) InboundNatPoolOutput {
	return o
}

func (o InboundNatPoolOutput) BackendPort() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatPool) int { return v.BackendPort }).(pulumi.IntOutput)
}

func (o InboundNatPoolOutput) FrontendPortRangeEnd() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatPool) int { return v.FrontendPortRangeEnd }).(pulumi.IntOutput)
}

func (o InboundNatPoolOutput) FrontendPortRangeStart() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatPool) int { return v.FrontendPortRangeStart }).(pulumi.IntOutput)
}

func (o InboundNatPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v InboundNatPool) string { return v.Name }).(pulumi.StringOutput)
}

func (o InboundNatPoolOutput) NetworkSecurityGroupRules() NetworkSecurityGroupRuleArrayOutput {
	return o.ApplyT(func(v InboundNatPool) []NetworkSecurityGroupRule { return v.NetworkSecurityGroupRules }).(NetworkSecurityGroupRuleArrayOutput)
}

func (o InboundNatPoolOutput) Protocol() InboundEndpointProtocolOutput {
	return o.ApplyT(func(v InboundNatPool) InboundEndpointProtocol { return v.Protocol }).(InboundEndpointProtocolOutput)
}

type InboundNatPoolArrayOutput struct{ *pulumi.OutputState }

func (InboundNatPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatPool)(nil)).Elem()
}

func (o InboundNatPoolArrayOutput) ToInboundNatPoolArrayOutput() InboundNatPoolArrayOutput {
	return o
}

func (o InboundNatPoolArrayOutput) ToInboundNatPoolArrayOutputWithContext(ctx context.Context) InboundNatPoolArrayOutput {
	return o
}

func (o InboundNatPoolArrayOutput) Index(i pulumi.IntInput) InboundNatPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundNatPool {
		return vs[0].([]InboundNatPool)[vs[1].(int)]
	}).(InboundNatPoolOutput)
}

type InboundNatPoolResponse struct {
	BackendPort               int                                `pulumi:"backendPort"`
	FrontendPortRangeEnd      int                                `pulumi:"frontendPortRangeEnd"`
	FrontendPortRangeStart    int                                `pulumi:"frontendPortRangeStart"`
	Name                      string                             `pulumi:"name"`
	NetworkSecurityGroupRules []NetworkSecurityGroupRuleResponse `pulumi:"networkSecurityGroupRules"`
	Protocol                  string                             `pulumi:"protocol"`
}

type InboundNatPoolResponseOutput struct{ *pulumi.OutputState }

func (InboundNatPoolResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatPoolResponse)(nil)).Elem()
}

func (o InboundNatPoolResponseOutput) ToInboundNatPoolResponseOutput() InboundNatPoolResponseOutput {
	return o
}

func (o InboundNatPoolResponseOutput) ToInboundNatPoolResponseOutputWithContext(ctx context.Context) InboundNatPoolResponseOutput {
	return o
}

func (o InboundNatPoolResponseOutput) BackendPort() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) int { return v.BackendPort }).(pulumi.IntOutput)
}

func (o InboundNatPoolResponseOutput) FrontendPortRangeEnd() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) int { return v.FrontendPortRangeEnd }).(pulumi.IntOutput)
}

func (o InboundNatPoolResponseOutput) FrontendPortRangeStart() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) int { return v.FrontendPortRangeStart }).(pulumi.IntOutput)
}

func (o InboundNatPoolResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o InboundNatPoolResponseOutput) NetworkSecurityGroupRules() NetworkSecurityGroupRuleResponseArrayOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) []NetworkSecurityGroupRuleResponse { return v.NetworkSecurityGroupRules }).(NetworkSecurityGroupRuleResponseArrayOutput)
}

func (o InboundNatPoolResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

type InboundNatPoolResponseArrayOutput struct{ *pulumi.OutputState }

func (InboundNatPoolResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatPoolResponse)(nil)).Elem()
}

func (o InboundNatPoolResponseArrayOutput) ToInboundNatPoolResponseArrayOutput() InboundNatPoolResponseArrayOutput {
	return o
}

func (o InboundNatPoolResponseArrayOutput) ToInboundNatPoolResponseArrayOutputWithContext(ctx context.Context) InboundNatPoolResponseArrayOutput {
	return o
}

func (o InboundNatPoolResponseArrayOutput) Index(i pulumi.IntInput) InboundNatPoolResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundNatPoolResponse {
		return vs[0].([]InboundNatPoolResponse)[vs[1].(int)]
	}).(InboundNatPoolResponseOutput)
}

type KeyVaultReference struct {
	Id  string `pulumi:"id"`
	Url string `pulumi:"url"`
}





type KeyVaultReferenceInput interface {
	pulumi.Input

	ToKeyVaultReferenceOutput() KeyVaultReferenceOutput
	ToKeyVaultReferenceOutputWithContext(context.Context) KeyVaultReferenceOutput
}

type KeyVaultReferenceArgs struct {
	Id  pulumi.StringInput `pulumi:"id"`
	Url pulumi.StringInput `pulumi:"url"`
}

func (KeyVaultReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultReference)(nil)).Elem()
}

func (i KeyVaultReferenceArgs) ToKeyVaultReferenceOutput() KeyVaultReferenceOutput {
	return i.ToKeyVaultReferenceOutputWithContext(context.Background())
}

func (i KeyVaultReferenceArgs) ToKeyVaultReferenceOutputWithContext(ctx context.Context) KeyVaultReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultReferenceOutput)
}

func (i KeyVaultReferenceArgs) ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput {
	return i.ToKeyVaultReferencePtrOutputWithContext(context.Background())
}

func (i KeyVaultReferenceArgs) ToKeyVaultReferencePtrOutputWithContext(ctx context.Context) KeyVaultReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultReferenceOutput).ToKeyVaultReferencePtrOutputWithContext(ctx)
}









type KeyVaultReferencePtrInput interface {
	pulumi.Input

	ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput
	ToKeyVaultReferencePtrOutputWithContext(context.Context) KeyVaultReferencePtrOutput
}

type keyVaultReferencePtrType KeyVaultReferenceArgs

func KeyVaultReferencePtr(v *KeyVaultReferenceArgs) KeyVaultReferencePtrInput {
	return (*keyVaultReferencePtrType)(v)
}

func (*keyVaultReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultReference)(nil)).Elem()
}

func (i *keyVaultReferencePtrType) ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput {
	return i.ToKeyVaultReferencePtrOutputWithContext(context.Background())
}

func (i *keyVaultReferencePtrType) ToKeyVaultReferencePtrOutputWithContext(ctx context.Context) KeyVaultReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultReferencePtrOutput)
}

type KeyVaultReferenceOutput struct{ *pulumi.OutputState }

func (KeyVaultReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultReference)(nil)).Elem()
}

func (o KeyVaultReferenceOutput) ToKeyVaultReferenceOutput() KeyVaultReferenceOutput {
	return o
}

func (o KeyVaultReferenceOutput) ToKeyVaultReferenceOutputWithContext(ctx context.Context) KeyVaultReferenceOutput {
	return o
}

func (o KeyVaultReferenceOutput) ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput {
	return o.ToKeyVaultReferencePtrOutputWithContext(context.Background())
}

func (o KeyVaultReferenceOutput) ToKeyVaultReferencePtrOutputWithContext(ctx context.Context) KeyVaultReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultReference) *KeyVaultReference {
		return &v
	}).(KeyVaultReferencePtrOutput)
}

func (o KeyVaultReferenceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultReference) string { return v.Id }).(pulumi.StringOutput)
}

func (o KeyVaultReferenceOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultReference) string { return v.Url }).(pulumi.StringOutput)
}

type KeyVaultReferencePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultReference)(nil)).Elem()
}

func (o KeyVaultReferencePtrOutput) ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput {
	return o
}

func (o KeyVaultReferencePtrOutput) ToKeyVaultReferencePtrOutputWithContext(ctx context.Context) KeyVaultReferencePtrOutput {
	return o
}

func (o KeyVaultReferencePtrOutput) Elem() KeyVaultReferenceOutput {
	return o.ApplyT(func(v *KeyVaultReference) KeyVaultReference {
		if v != nil {
			return *v
		}
		var ret KeyVaultReference
		return ret
	}).(KeyVaultReferenceOutput)
}

func (o KeyVaultReferencePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultReference) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultReferencePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultReference) *string {
		if v == nil {
			return nil
		}
		return &v.Url
	}).(pulumi.StringPtrOutput)
}

type KeyVaultReferenceResponse struct {
	Id  string `pulumi:"id"`
	Url string `pulumi:"url"`
}

type KeyVaultReferenceResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultReferenceResponse)(nil)).Elem()
}

func (o KeyVaultReferenceResponseOutput) ToKeyVaultReferenceResponseOutput() KeyVaultReferenceResponseOutput {
	return o
}

func (o KeyVaultReferenceResponseOutput) ToKeyVaultReferenceResponseOutputWithContext(ctx context.Context) KeyVaultReferenceResponseOutput {
	return o
}

func (o KeyVaultReferenceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultReferenceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o KeyVaultReferenceResponseOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultReferenceResponse) string { return v.Url }).(pulumi.StringOutput)
}

type LinuxUserConfiguration struct {
	Gid           *int    `pulumi:"gid"`
	SshPrivateKey *string `pulumi:"sshPrivateKey"`
	Uid           *int    `pulumi:"uid"`
}





type LinuxUserConfigurationInput interface {
	pulumi.Input

	ToLinuxUserConfigurationOutput() LinuxUserConfigurationOutput
	ToLinuxUserConfigurationOutputWithContext(context.Context) LinuxUserConfigurationOutput
}

type LinuxUserConfigurationArgs struct {
	Gid           pulumi.IntPtrInput    `pulumi:"gid"`
	SshPrivateKey pulumi.StringPtrInput `pulumi:"sshPrivateKey"`
	Uid           pulumi.IntPtrInput    `pulumi:"uid"`
}

func (LinuxUserConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxUserConfiguration)(nil)).Elem()
}

func (i LinuxUserConfigurationArgs) ToLinuxUserConfigurationOutput() LinuxUserConfigurationOutput {
	return i.ToLinuxUserConfigurationOutputWithContext(context.Background())
}

func (i LinuxUserConfigurationArgs) ToLinuxUserConfigurationOutputWithContext(ctx context.Context) LinuxUserConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxUserConfigurationOutput)
}

func (i LinuxUserConfigurationArgs) ToLinuxUserConfigurationPtrOutput() LinuxUserConfigurationPtrOutput {
	return i.ToLinuxUserConfigurationPtrOutputWithContext(context.Background())
}

func (i LinuxUserConfigurationArgs) ToLinuxUserConfigurationPtrOutputWithContext(ctx context.Context) LinuxUserConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxUserConfigurationOutput).ToLinuxUserConfigurationPtrOutputWithContext(ctx)
}









type LinuxUserConfigurationPtrInput interface {
	pulumi.Input

	ToLinuxUserConfigurationPtrOutput() LinuxUserConfigurationPtrOutput
	ToLinuxUserConfigurationPtrOutputWithContext(context.Context) LinuxUserConfigurationPtrOutput
}

type linuxUserConfigurationPtrType LinuxUserConfigurationArgs

func LinuxUserConfigurationPtr(v *LinuxUserConfigurationArgs) LinuxUserConfigurationPtrInput {
	return (*linuxUserConfigurationPtrType)(v)
}

func (*linuxUserConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxUserConfiguration)(nil)).Elem()
}

func (i *linuxUserConfigurationPtrType) ToLinuxUserConfigurationPtrOutput() LinuxUserConfigurationPtrOutput {
	return i.ToLinuxUserConfigurationPtrOutputWithContext(context.Background())
}

func (i *linuxUserConfigurationPtrType) ToLinuxUserConfigurationPtrOutputWithContext(ctx context.Context) LinuxUserConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxUserConfigurationPtrOutput)
}

type LinuxUserConfigurationOutput struct{ *pulumi.OutputState }

func (LinuxUserConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxUserConfiguration)(nil)).Elem()
}

func (o LinuxUserConfigurationOutput) ToLinuxUserConfigurationOutput() LinuxUserConfigurationOutput {
	return o
}

func (o LinuxUserConfigurationOutput) ToLinuxUserConfigurationOutputWithContext(ctx context.Context) LinuxUserConfigurationOutput {
	return o
}

func (o LinuxUserConfigurationOutput) ToLinuxUserConfigurationPtrOutput() LinuxUserConfigurationPtrOutput {
	return o.ToLinuxUserConfigurationPtrOutputWithContext(context.Background())
}

func (o LinuxUserConfigurationOutput) ToLinuxUserConfigurationPtrOutputWithContext(ctx context.Context) LinuxUserConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinuxUserConfiguration) *LinuxUserConfiguration {
		return &v
	}).(LinuxUserConfigurationPtrOutput)
}

func (o LinuxUserConfigurationOutput) Gid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LinuxUserConfiguration) *int { return v.Gid }).(pulumi.IntPtrOutput)
}

func (o LinuxUserConfigurationOutput) SshPrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxUserConfiguration) *string { return v.SshPrivateKey }).(pulumi.StringPtrOutput)
}

func (o LinuxUserConfigurationOutput) Uid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LinuxUserConfiguration) *int { return v.Uid }).(pulumi.IntPtrOutput)
}

type LinuxUserConfigurationPtrOutput struct{ *pulumi.OutputState }

func (LinuxUserConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxUserConfiguration)(nil)).Elem()
}

func (o LinuxUserConfigurationPtrOutput) ToLinuxUserConfigurationPtrOutput() LinuxUserConfigurationPtrOutput {
	return o
}

func (o LinuxUserConfigurationPtrOutput) ToLinuxUserConfigurationPtrOutputWithContext(ctx context.Context) LinuxUserConfigurationPtrOutput {
	return o
}

func (o LinuxUserConfigurationPtrOutput) Elem() LinuxUserConfigurationOutput {
	return o.ApplyT(func(v *LinuxUserConfiguration) LinuxUserConfiguration {
		if v != nil {
			return *v
		}
		var ret LinuxUserConfiguration
		return ret
	}).(LinuxUserConfigurationOutput)
}

func (o LinuxUserConfigurationPtrOutput) Gid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LinuxUserConfiguration) *int {
		if v == nil {
			return nil
		}
		return v.Gid
	}).(pulumi.IntPtrOutput)
}

func (o LinuxUserConfigurationPtrOutput) SshPrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxUserConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.SshPrivateKey
	}).(pulumi.StringPtrOutput)
}

func (o LinuxUserConfigurationPtrOutput) Uid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LinuxUserConfiguration) *int {
		if v == nil {
			return nil
		}
		return v.Uid
	}).(pulumi.IntPtrOutput)
}

type LinuxUserConfigurationResponse struct {
	Gid           *int    `pulumi:"gid"`
	SshPrivateKey *string `pulumi:"sshPrivateKey"`
	Uid           *int    `pulumi:"uid"`
}

type LinuxUserConfigurationResponseOutput struct{ *pulumi.OutputState }

func (LinuxUserConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxUserConfigurationResponse)(nil)).Elem()
}

func (o LinuxUserConfigurationResponseOutput) ToLinuxUserConfigurationResponseOutput() LinuxUserConfigurationResponseOutput {
	return o
}

func (o LinuxUserConfigurationResponseOutput) ToLinuxUserConfigurationResponseOutputWithContext(ctx context.Context) LinuxUserConfigurationResponseOutput {
	return o
}

func (o LinuxUserConfigurationResponseOutput) Gid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LinuxUserConfigurationResponse) *int { return v.Gid }).(pulumi.IntPtrOutput)
}

func (o LinuxUserConfigurationResponseOutput) SshPrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxUserConfigurationResponse) *string { return v.SshPrivateKey }).(pulumi.StringPtrOutput)
}

func (o LinuxUserConfigurationResponseOutput) Uid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LinuxUserConfigurationResponse) *int { return v.Uid }).(pulumi.IntPtrOutput)
}

type LinuxUserConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (LinuxUserConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxUserConfigurationResponse)(nil)).Elem()
}

func (o LinuxUserConfigurationResponsePtrOutput) ToLinuxUserConfigurationResponsePtrOutput() LinuxUserConfigurationResponsePtrOutput {
	return o
}

func (o LinuxUserConfigurationResponsePtrOutput) ToLinuxUserConfigurationResponsePtrOutputWithContext(ctx context.Context) LinuxUserConfigurationResponsePtrOutput {
	return o
}

func (o LinuxUserConfigurationResponsePtrOutput) Elem() LinuxUserConfigurationResponseOutput {
	return o.ApplyT(func(v *LinuxUserConfigurationResponse) LinuxUserConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret LinuxUserConfigurationResponse
		return ret
	}).(LinuxUserConfigurationResponseOutput)
}

func (o LinuxUserConfigurationResponsePtrOutput) Gid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LinuxUserConfigurationResponse) *int {
		if v == nil {
			return nil
		}
		return v.Gid
	}).(pulumi.IntPtrOutput)
}

func (o LinuxUserConfigurationResponsePtrOutput) SshPrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxUserConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.SshPrivateKey
	}).(pulumi.StringPtrOutput)
}

func (o LinuxUserConfigurationResponsePtrOutput) Uid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LinuxUserConfigurationResponse) *int {
		if v == nil {
			return nil
		}
		return v.Uid
	}).(pulumi.IntPtrOutput)
}

type MetadataItem struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type MetadataItemInput interface {
	pulumi.Input

	ToMetadataItemOutput() MetadataItemOutput
	ToMetadataItemOutputWithContext(context.Context) MetadataItemOutput
}

type MetadataItemArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (MetadataItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataItem)(nil)).Elem()
}

func (i MetadataItemArgs) ToMetadataItemOutput() MetadataItemOutput {
	return i.ToMetadataItemOutputWithContext(context.Background())
}

func (i MetadataItemArgs) ToMetadataItemOutputWithContext(ctx context.Context) MetadataItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataItemOutput)
}





type MetadataItemArrayInput interface {
	pulumi.Input

	ToMetadataItemArrayOutput() MetadataItemArrayOutput
	ToMetadataItemArrayOutputWithContext(context.Context) MetadataItemArrayOutput
}

type MetadataItemArray []MetadataItemInput

func (MetadataItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetadataItem)(nil)).Elem()
}

func (i MetadataItemArray) ToMetadataItemArrayOutput() MetadataItemArrayOutput {
	return i.ToMetadataItemArrayOutputWithContext(context.Background())
}

func (i MetadataItemArray) ToMetadataItemArrayOutputWithContext(ctx context.Context) MetadataItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataItemArrayOutput)
}

type MetadataItemOutput struct{ *pulumi.OutputState }

func (MetadataItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataItem)(nil)).Elem()
}

func (o MetadataItemOutput) ToMetadataItemOutput() MetadataItemOutput {
	return o
}

func (o MetadataItemOutput) ToMetadataItemOutputWithContext(ctx context.Context) MetadataItemOutput {
	return o
}

func (o MetadataItemOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MetadataItem) string { return v.Name }).(pulumi.StringOutput)
}

func (o MetadataItemOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v MetadataItem) string { return v.Value }).(pulumi.StringOutput)
}

type MetadataItemArrayOutput struct{ *pulumi.OutputState }

func (MetadataItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetadataItem)(nil)).Elem()
}

func (o MetadataItemArrayOutput) ToMetadataItemArrayOutput() MetadataItemArrayOutput {
	return o
}

func (o MetadataItemArrayOutput) ToMetadataItemArrayOutputWithContext(ctx context.Context) MetadataItemArrayOutput {
	return o
}

func (o MetadataItemArrayOutput) Index(i pulumi.IntInput) MetadataItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetadataItem {
		return vs[0].([]MetadataItem)[vs[1].(int)]
	}).(MetadataItemOutput)
}

type MetadataItemResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type MetadataItemResponseOutput struct{ *pulumi.OutputState }

func (MetadataItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataItemResponse)(nil)).Elem()
}

func (o MetadataItemResponseOutput) ToMetadataItemResponseOutput() MetadataItemResponseOutput {
	return o
}

func (o MetadataItemResponseOutput) ToMetadataItemResponseOutputWithContext(ctx context.Context) MetadataItemResponseOutput {
	return o
}

func (o MetadataItemResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MetadataItemResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o MetadataItemResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v MetadataItemResponse) string { return v.Value }).(pulumi.StringOutput)
}

type MetadataItemResponseArrayOutput struct{ *pulumi.OutputState }

func (MetadataItemResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetadataItemResponse)(nil)).Elem()
}

func (o MetadataItemResponseArrayOutput) ToMetadataItemResponseArrayOutput() MetadataItemResponseArrayOutput {
	return o
}

func (o MetadataItemResponseArrayOutput) ToMetadataItemResponseArrayOutputWithContext(ctx context.Context) MetadataItemResponseArrayOutput {
	return o
}

func (o MetadataItemResponseArrayOutput) Index(i pulumi.IntInput) MetadataItemResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetadataItemResponse {
		return vs[0].([]MetadataItemResponse)[vs[1].(int)]
	}).(MetadataItemResponseOutput)
}

type NetworkConfiguration struct {
	EndpointConfiguration *PoolEndpointConfiguration `pulumi:"endpointConfiguration"`
	SubnetId              *string                    `pulumi:"subnetId"`
}





type NetworkConfigurationInput interface {
	pulumi.Input

	ToNetworkConfigurationOutput() NetworkConfigurationOutput
	ToNetworkConfigurationOutputWithContext(context.Context) NetworkConfigurationOutput
}

type NetworkConfigurationArgs struct {
	EndpointConfiguration PoolEndpointConfigurationPtrInput `pulumi:"endpointConfiguration"`
	SubnetId              pulumi.StringPtrInput             `pulumi:"subnetId"`
}

func (NetworkConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkConfiguration)(nil)).Elem()
}

func (i NetworkConfigurationArgs) ToNetworkConfigurationOutput() NetworkConfigurationOutput {
	return i.ToNetworkConfigurationOutputWithContext(context.Background())
}

func (i NetworkConfigurationArgs) ToNetworkConfigurationOutputWithContext(ctx context.Context) NetworkConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkConfigurationOutput)
}

func (i NetworkConfigurationArgs) ToNetworkConfigurationPtrOutput() NetworkConfigurationPtrOutput {
	return i.ToNetworkConfigurationPtrOutputWithContext(context.Background())
}

func (i NetworkConfigurationArgs) ToNetworkConfigurationPtrOutputWithContext(ctx context.Context) NetworkConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkConfigurationOutput).ToNetworkConfigurationPtrOutputWithContext(ctx)
}









type NetworkConfigurationPtrInput interface {
	pulumi.Input

	ToNetworkConfigurationPtrOutput() NetworkConfigurationPtrOutput
	ToNetworkConfigurationPtrOutputWithContext(context.Context) NetworkConfigurationPtrOutput
}

type networkConfigurationPtrType NetworkConfigurationArgs

func NetworkConfigurationPtr(v *NetworkConfigurationArgs) NetworkConfigurationPtrInput {
	return (*networkConfigurationPtrType)(v)
}

func (*networkConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkConfiguration)(nil)).Elem()
}

func (i *networkConfigurationPtrType) ToNetworkConfigurationPtrOutput() NetworkConfigurationPtrOutput {
	return i.ToNetworkConfigurationPtrOutputWithContext(context.Background())
}

func (i *networkConfigurationPtrType) ToNetworkConfigurationPtrOutputWithContext(ctx context.Context) NetworkConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkConfigurationPtrOutput)
}

type NetworkConfigurationOutput struct{ *pulumi.OutputState }

func (NetworkConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkConfiguration)(nil)).Elem()
}

func (o NetworkConfigurationOutput) ToNetworkConfigurationOutput() NetworkConfigurationOutput {
	return o
}

func (o NetworkConfigurationOutput) ToNetworkConfigurationOutputWithContext(ctx context.Context) NetworkConfigurationOutput {
	return o
}

func (o NetworkConfigurationOutput) ToNetworkConfigurationPtrOutput() NetworkConfigurationPtrOutput {
	return o.ToNetworkConfigurationPtrOutputWithContext(context.Background())
}

func (o NetworkConfigurationOutput) ToNetworkConfigurationPtrOutputWithContext(ctx context.Context) NetworkConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkConfiguration) *NetworkConfiguration {
		return &v
	}).(NetworkConfigurationPtrOutput)
}

func (o NetworkConfigurationOutput) EndpointConfiguration() PoolEndpointConfigurationPtrOutput {
	return o.ApplyT(func(v NetworkConfiguration) *PoolEndpointConfiguration { return v.EndpointConfiguration }).(PoolEndpointConfigurationPtrOutput)
}

func (o NetworkConfigurationOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkConfiguration) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

type NetworkConfigurationPtrOutput struct{ *pulumi.OutputState }

func (NetworkConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkConfiguration)(nil)).Elem()
}

func (o NetworkConfigurationPtrOutput) ToNetworkConfigurationPtrOutput() NetworkConfigurationPtrOutput {
	return o
}

func (o NetworkConfigurationPtrOutput) ToNetworkConfigurationPtrOutputWithContext(ctx context.Context) NetworkConfigurationPtrOutput {
	return o
}

func (o NetworkConfigurationPtrOutput) Elem() NetworkConfigurationOutput {
	return o.ApplyT(func(v *NetworkConfiguration) NetworkConfiguration {
		if v != nil {
			return *v
		}
		var ret NetworkConfiguration
		return ret
	}).(NetworkConfigurationOutput)
}

func (o NetworkConfigurationPtrOutput) EndpointConfiguration() PoolEndpointConfigurationPtrOutput {
	return o.ApplyT(func(v *NetworkConfiguration) *PoolEndpointConfiguration {
		if v == nil {
			return nil
		}
		return v.EndpointConfiguration
	}).(PoolEndpointConfigurationPtrOutput)
}

func (o NetworkConfigurationPtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.SubnetId
	}).(pulumi.StringPtrOutput)
}

type NetworkConfigurationResponse struct {
	EndpointConfiguration *PoolEndpointConfigurationResponse `pulumi:"endpointConfiguration"`
	SubnetId              *string                            `pulumi:"subnetId"`
}

type NetworkConfigurationResponseOutput struct{ *pulumi.OutputState }

func (NetworkConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkConfigurationResponse)(nil)).Elem()
}

func (o NetworkConfigurationResponseOutput) ToNetworkConfigurationResponseOutput() NetworkConfigurationResponseOutput {
	return o
}

func (o NetworkConfigurationResponseOutput) ToNetworkConfigurationResponseOutputWithContext(ctx context.Context) NetworkConfigurationResponseOutput {
	return o
}

func (o NetworkConfigurationResponseOutput) EndpointConfiguration() PoolEndpointConfigurationResponsePtrOutput {
	return o.ApplyT(func(v NetworkConfigurationResponse) *PoolEndpointConfigurationResponse {
		return v.EndpointConfiguration
	}).(PoolEndpointConfigurationResponsePtrOutput)
}

func (o NetworkConfigurationResponseOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkConfigurationResponse) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

type NetworkConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkConfigurationResponse)(nil)).Elem()
}

func (o NetworkConfigurationResponsePtrOutput) ToNetworkConfigurationResponsePtrOutput() NetworkConfigurationResponsePtrOutput {
	return o
}

func (o NetworkConfigurationResponsePtrOutput) ToNetworkConfigurationResponsePtrOutputWithContext(ctx context.Context) NetworkConfigurationResponsePtrOutput {
	return o
}

func (o NetworkConfigurationResponsePtrOutput) Elem() NetworkConfigurationResponseOutput {
	return o.ApplyT(func(v *NetworkConfigurationResponse) NetworkConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret NetworkConfigurationResponse
		return ret
	}).(NetworkConfigurationResponseOutput)
}

func (o NetworkConfigurationResponsePtrOutput) EndpointConfiguration() PoolEndpointConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *NetworkConfigurationResponse) *PoolEndpointConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.EndpointConfiguration
	}).(PoolEndpointConfigurationResponsePtrOutput)
}

func (o NetworkConfigurationResponsePtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubnetId
	}).(pulumi.StringPtrOutput)
}

type NetworkSecurityGroupRule struct {
	Access              NetworkSecurityGroupRuleAccess `pulumi:"access"`
	Priority            int                            `pulumi:"priority"`
	SourceAddressPrefix string                         `pulumi:"sourceAddressPrefix"`
}





type NetworkSecurityGroupRuleInput interface {
	pulumi.Input

	ToNetworkSecurityGroupRuleOutput() NetworkSecurityGroupRuleOutput
	ToNetworkSecurityGroupRuleOutputWithContext(context.Context) NetworkSecurityGroupRuleOutput
}

type NetworkSecurityGroupRuleArgs struct {
	Access              NetworkSecurityGroupRuleAccessInput `pulumi:"access"`
	Priority            pulumi.IntInput                     `pulumi:"priority"`
	SourceAddressPrefix pulumi.StringInput                  `pulumi:"sourceAddressPrefix"`
}

func (NetworkSecurityGroupRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupRule)(nil)).Elem()
}

func (i NetworkSecurityGroupRuleArgs) ToNetworkSecurityGroupRuleOutput() NetworkSecurityGroupRuleOutput {
	return i.ToNetworkSecurityGroupRuleOutputWithContext(context.Background())
}

func (i NetworkSecurityGroupRuleArgs) ToNetworkSecurityGroupRuleOutputWithContext(ctx context.Context) NetworkSecurityGroupRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityGroupRuleOutput)
}





type NetworkSecurityGroupRuleArrayInput interface {
	pulumi.Input

	ToNetworkSecurityGroupRuleArrayOutput() NetworkSecurityGroupRuleArrayOutput
	ToNetworkSecurityGroupRuleArrayOutputWithContext(context.Context) NetworkSecurityGroupRuleArrayOutput
}

type NetworkSecurityGroupRuleArray []NetworkSecurityGroupRuleInput

func (NetworkSecurityGroupRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkSecurityGroupRule)(nil)).Elem()
}

func (i NetworkSecurityGroupRuleArray) ToNetworkSecurityGroupRuleArrayOutput() NetworkSecurityGroupRuleArrayOutput {
	return i.ToNetworkSecurityGroupRuleArrayOutputWithContext(context.Background())
}

func (i NetworkSecurityGroupRuleArray) ToNetworkSecurityGroupRuleArrayOutputWithContext(ctx context.Context) NetworkSecurityGroupRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityGroupRuleArrayOutput)
}

type NetworkSecurityGroupRuleOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupRule)(nil)).Elem()
}

func (o NetworkSecurityGroupRuleOutput) ToNetworkSecurityGroupRuleOutput() NetworkSecurityGroupRuleOutput {
	return o
}

func (o NetworkSecurityGroupRuleOutput) ToNetworkSecurityGroupRuleOutputWithContext(ctx context.Context) NetworkSecurityGroupRuleOutput {
	return o
}

func (o NetworkSecurityGroupRuleOutput) Access() NetworkSecurityGroupRuleAccessOutput {
	return o.ApplyT(func(v NetworkSecurityGroupRule) NetworkSecurityGroupRuleAccess { return v.Access }).(NetworkSecurityGroupRuleAccessOutput)
}

func (o NetworkSecurityGroupRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v NetworkSecurityGroupRule) int { return v.Priority }).(pulumi.IntOutput)
}

func (o NetworkSecurityGroupRuleOutput) SourceAddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityGroupRule) string { return v.SourceAddressPrefix }).(pulumi.StringOutput)
}

type NetworkSecurityGroupRuleArrayOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkSecurityGroupRule)(nil)).Elem()
}

func (o NetworkSecurityGroupRuleArrayOutput) ToNetworkSecurityGroupRuleArrayOutput() NetworkSecurityGroupRuleArrayOutput {
	return o
}

func (o NetworkSecurityGroupRuleArrayOutput) ToNetworkSecurityGroupRuleArrayOutputWithContext(ctx context.Context) NetworkSecurityGroupRuleArrayOutput {
	return o
}

func (o NetworkSecurityGroupRuleArrayOutput) Index(i pulumi.IntInput) NetworkSecurityGroupRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkSecurityGroupRule {
		return vs[0].([]NetworkSecurityGroupRule)[vs[1].(int)]
	}).(NetworkSecurityGroupRuleOutput)
}

type NetworkSecurityGroupRuleResponse struct {
	Access              string `pulumi:"access"`
	Priority            int    `pulumi:"priority"`
	SourceAddressPrefix string `pulumi:"sourceAddressPrefix"`
}

type NetworkSecurityGroupRuleResponseOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupRuleResponse)(nil)).Elem()
}

func (o NetworkSecurityGroupRuleResponseOutput) ToNetworkSecurityGroupRuleResponseOutput() NetworkSecurityGroupRuleResponseOutput {
	return o
}

func (o NetworkSecurityGroupRuleResponseOutput) ToNetworkSecurityGroupRuleResponseOutputWithContext(ctx context.Context) NetworkSecurityGroupRuleResponseOutput {
	return o
}

func (o NetworkSecurityGroupRuleResponseOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityGroupRuleResponse) string { return v.Access }).(pulumi.StringOutput)
}

func (o NetworkSecurityGroupRuleResponseOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v NetworkSecurityGroupRuleResponse) int { return v.Priority }).(pulumi.IntOutput)
}

func (o NetworkSecurityGroupRuleResponseOutput) SourceAddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityGroupRuleResponse) string { return v.SourceAddressPrefix }).(pulumi.StringOutput)
}

type NetworkSecurityGroupRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkSecurityGroupRuleResponse)(nil)).Elem()
}

func (o NetworkSecurityGroupRuleResponseArrayOutput) ToNetworkSecurityGroupRuleResponseArrayOutput() NetworkSecurityGroupRuleResponseArrayOutput {
	return o
}

func (o NetworkSecurityGroupRuleResponseArrayOutput) ToNetworkSecurityGroupRuleResponseArrayOutputWithContext(ctx context.Context) NetworkSecurityGroupRuleResponseArrayOutput {
	return o
}

func (o NetworkSecurityGroupRuleResponseArrayOutput) Index(i pulumi.IntInput) NetworkSecurityGroupRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkSecurityGroupRuleResponse {
		return vs[0].([]NetworkSecurityGroupRuleResponse)[vs[1].(int)]
	}).(NetworkSecurityGroupRuleResponseOutput)
}

type PoolEndpointConfiguration struct {
	InboundNatPools []InboundNatPool `pulumi:"inboundNatPools"`
}





type PoolEndpointConfigurationInput interface {
	pulumi.Input

	ToPoolEndpointConfigurationOutput() PoolEndpointConfigurationOutput
	ToPoolEndpointConfigurationOutputWithContext(context.Context) PoolEndpointConfigurationOutput
}

type PoolEndpointConfigurationArgs struct {
	InboundNatPools InboundNatPoolArrayInput `pulumi:"inboundNatPools"`
}

func (PoolEndpointConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PoolEndpointConfiguration)(nil)).Elem()
}

func (i PoolEndpointConfigurationArgs) ToPoolEndpointConfigurationOutput() PoolEndpointConfigurationOutput {
	return i.ToPoolEndpointConfigurationOutputWithContext(context.Background())
}

func (i PoolEndpointConfigurationArgs) ToPoolEndpointConfigurationOutputWithContext(ctx context.Context) PoolEndpointConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolEndpointConfigurationOutput)
}

func (i PoolEndpointConfigurationArgs) ToPoolEndpointConfigurationPtrOutput() PoolEndpointConfigurationPtrOutput {
	return i.ToPoolEndpointConfigurationPtrOutputWithContext(context.Background())
}

func (i PoolEndpointConfigurationArgs) ToPoolEndpointConfigurationPtrOutputWithContext(ctx context.Context) PoolEndpointConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolEndpointConfigurationOutput).ToPoolEndpointConfigurationPtrOutputWithContext(ctx)
}









type PoolEndpointConfigurationPtrInput interface {
	pulumi.Input

	ToPoolEndpointConfigurationPtrOutput() PoolEndpointConfigurationPtrOutput
	ToPoolEndpointConfigurationPtrOutputWithContext(context.Context) PoolEndpointConfigurationPtrOutput
}

type poolEndpointConfigurationPtrType PoolEndpointConfigurationArgs

func PoolEndpointConfigurationPtr(v *PoolEndpointConfigurationArgs) PoolEndpointConfigurationPtrInput {
	return (*poolEndpointConfigurationPtrType)(v)
}

func (*poolEndpointConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PoolEndpointConfiguration)(nil)).Elem()
}

func (i *poolEndpointConfigurationPtrType) ToPoolEndpointConfigurationPtrOutput() PoolEndpointConfigurationPtrOutput {
	return i.ToPoolEndpointConfigurationPtrOutputWithContext(context.Background())
}

func (i *poolEndpointConfigurationPtrType) ToPoolEndpointConfigurationPtrOutputWithContext(ctx context.Context) PoolEndpointConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolEndpointConfigurationPtrOutput)
}

type PoolEndpointConfigurationOutput struct{ *pulumi.OutputState }

func (PoolEndpointConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PoolEndpointConfiguration)(nil)).Elem()
}

func (o PoolEndpointConfigurationOutput) ToPoolEndpointConfigurationOutput() PoolEndpointConfigurationOutput {
	return o
}

func (o PoolEndpointConfigurationOutput) ToPoolEndpointConfigurationOutputWithContext(ctx context.Context) PoolEndpointConfigurationOutput {
	return o
}

func (o PoolEndpointConfigurationOutput) ToPoolEndpointConfigurationPtrOutput() PoolEndpointConfigurationPtrOutput {
	return o.ToPoolEndpointConfigurationPtrOutputWithContext(context.Background())
}

func (o PoolEndpointConfigurationOutput) ToPoolEndpointConfigurationPtrOutputWithContext(ctx context.Context) PoolEndpointConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PoolEndpointConfiguration) *PoolEndpointConfiguration {
		return &v
	}).(PoolEndpointConfigurationPtrOutput)
}

func (o PoolEndpointConfigurationOutput) InboundNatPools() InboundNatPoolArrayOutput {
	return o.ApplyT(func(v PoolEndpointConfiguration) []InboundNatPool { return v.InboundNatPools }).(InboundNatPoolArrayOutput)
}

type PoolEndpointConfigurationPtrOutput struct{ *pulumi.OutputState }

func (PoolEndpointConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PoolEndpointConfiguration)(nil)).Elem()
}

func (o PoolEndpointConfigurationPtrOutput) ToPoolEndpointConfigurationPtrOutput() PoolEndpointConfigurationPtrOutput {
	return o
}

func (o PoolEndpointConfigurationPtrOutput) ToPoolEndpointConfigurationPtrOutputWithContext(ctx context.Context) PoolEndpointConfigurationPtrOutput {
	return o
}

func (o PoolEndpointConfigurationPtrOutput) Elem() PoolEndpointConfigurationOutput {
	return o.ApplyT(func(v *PoolEndpointConfiguration) PoolEndpointConfiguration {
		if v != nil {
			return *v
		}
		var ret PoolEndpointConfiguration
		return ret
	}).(PoolEndpointConfigurationOutput)
}

func (o PoolEndpointConfigurationPtrOutput) InboundNatPools() InboundNatPoolArrayOutput {
	return o.ApplyT(func(v *PoolEndpointConfiguration) []InboundNatPool {
		if v == nil {
			return nil
		}
		return v.InboundNatPools
	}).(InboundNatPoolArrayOutput)
}

type PoolEndpointConfigurationResponse struct {
	InboundNatPools []InboundNatPoolResponse `pulumi:"inboundNatPools"`
}

type PoolEndpointConfigurationResponseOutput struct{ *pulumi.OutputState }

func (PoolEndpointConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PoolEndpointConfigurationResponse)(nil)).Elem()
}

func (o PoolEndpointConfigurationResponseOutput) ToPoolEndpointConfigurationResponseOutput() PoolEndpointConfigurationResponseOutput {
	return o
}

func (o PoolEndpointConfigurationResponseOutput) ToPoolEndpointConfigurationResponseOutputWithContext(ctx context.Context) PoolEndpointConfigurationResponseOutput {
	return o
}

func (o PoolEndpointConfigurationResponseOutput) InboundNatPools() InboundNatPoolResponseArrayOutput {
	return o.ApplyT(func(v PoolEndpointConfigurationResponse) []InboundNatPoolResponse { return v.InboundNatPools }).(InboundNatPoolResponseArrayOutput)
}

type PoolEndpointConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (PoolEndpointConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PoolEndpointConfigurationResponse)(nil)).Elem()
}

func (o PoolEndpointConfigurationResponsePtrOutput) ToPoolEndpointConfigurationResponsePtrOutput() PoolEndpointConfigurationResponsePtrOutput {
	return o
}

func (o PoolEndpointConfigurationResponsePtrOutput) ToPoolEndpointConfigurationResponsePtrOutputWithContext(ctx context.Context) PoolEndpointConfigurationResponsePtrOutput {
	return o
}

func (o PoolEndpointConfigurationResponsePtrOutput) Elem() PoolEndpointConfigurationResponseOutput {
	return o.ApplyT(func(v *PoolEndpointConfigurationResponse) PoolEndpointConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret PoolEndpointConfigurationResponse
		return ret
	}).(PoolEndpointConfigurationResponseOutput)
}

func (o PoolEndpointConfigurationResponsePtrOutput) InboundNatPools() InboundNatPoolResponseArrayOutput {
	return o.ApplyT(func(v *PoolEndpointConfigurationResponse) []InboundNatPoolResponse {
		if v == nil {
			return nil
		}
		return v.InboundNatPools
	}).(InboundNatPoolResponseArrayOutput)
}

type ResizeErrorResponse struct {
	Code    string                `pulumi:"code"`
	Details []ResizeErrorResponse `pulumi:"details"`
	Message string                `pulumi:"message"`
}

type ResizeErrorResponseOutput struct{ *pulumi.OutputState }

func (ResizeErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResizeErrorResponse)(nil)).Elem()
}

func (o ResizeErrorResponseOutput) ToResizeErrorResponseOutput() ResizeErrorResponseOutput {
	return o
}

func (o ResizeErrorResponseOutput) ToResizeErrorResponseOutputWithContext(ctx context.Context) ResizeErrorResponseOutput {
	return o
}

func (o ResizeErrorResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ResizeErrorResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ResizeErrorResponseOutput) Details() ResizeErrorResponseArrayOutput {
	return o.ApplyT(func(v ResizeErrorResponse) []ResizeErrorResponse { return v.Details }).(ResizeErrorResponseArrayOutput)
}

func (o ResizeErrorResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ResizeErrorResponse) string { return v.Message }).(pulumi.StringOutput)
}

type ResizeErrorResponseArrayOutput struct{ *pulumi.OutputState }

func (ResizeErrorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResizeErrorResponse)(nil)).Elem()
}

func (o ResizeErrorResponseArrayOutput) ToResizeErrorResponseArrayOutput() ResizeErrorResponseArrayOutput {
	return o
}

func (o ResizeErrorResponseArrayOutput) ToResizeErrorResponseArrayOutputWithContext(ctx context.Context) ResizeErrorResponseArrayOutput {
	return o
}

func (o ResizeErrorResponseArrayOutput) Index(i pulumi.IntInput) ResizeErrorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResizeErrorResponse {
		return vs[0].([]ResizeErrorResponse)[vs[1].(int)]
	}).(ResizeErrorResponseOutput)
}

type ResizeOperationStatusResponse struct {
	Errors                 []ResizeErrorResponse `pulumi:"errors"`
	NodeDeallocationOption *string               `pulumi:"nodeDeallocationOption"`
	ResizeTimeout          *string               `pulumi:"resizeTimeout"`
	StartTime              *string               `pulumi:"startTime"`
	TargetDedicatedNodes   *int                  `pulumi:"targetDedicatedNodes"`
	TargetLowPriorityNodes *int                  `pulumi:"targetLowPriorityNodes"`
}

type ResizeOperationStatusResponseOutput struct{ *pulumi.OutputState }

func (ResizeOperationStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResizeOperationStatusResponse)(nil)).Elem()
}

func (o ResizeOperationStatusResponseOutput) ToResizeOperationStatusResponseOutput() ResizeOperationStatusResponseOutput {
	return o
}

func (o ResizeOperationStatusResponseOutput) ToResizeOperationStatusResponseOutputWithContext(ctx context.Context) ResizeOperationStatusResponseOutput {
	return o
}

func (o ResizeOperationStatusResponseOutput) Errors() ResizeErrorResponseArrayOutput {
	return o.ApplyT(func(v ResizeOperationStatusResponse) []ResizeErrorResponse { return v.Errors }).(ResizeErrorResponseArrayOutput)
}

func (o ResizeOperationStatusResponseOutput) NodeDeallocationOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResizeOperationStatusResponse) *string { return v.NodeDeallocationOption }).(pulumi.StringPtrOutput)
}

func (o ResizeOperationStatusResponseOutput) ResizeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResizeOperationStatusResponse) *string { return v.ResizeTimeout }).(pulumi.StringPtrOutput)
}

func (o ResizeOperationStatusResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResizeOperationStatusResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o ResizeOperationStatusResponseOutput) TargetDedicatedNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResizeOperationStatusResponse) *int { return v.TargetDedicatedNodes }).(pulumi.IntPtrOutput)
}

func (o ResizeOperationStatusResponseOutput) TargetLowPriorityNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResizeOperationStatusResponse) *int { return v.TargetLowPriorityNodes }).(pulumi.IntPtrOutput)
}

type ResourceFile struct {
	AutoStorageContainerName *string `pulumi:"autoStorageContainerName"`
	BlobPrefix               *string `pulumi:"blobPrefix"`
	FileMode                 *string `pulumi:"fileMode"`
	FilePath                 *string `pulumi:"filePath"`
	HttpUrl                  *string `pulumi:"httpUrl"`
	StorageContainerUrl      *string `pulumi:"storageContainerUrl"`
}





type ResourceFileInput interface {
	pulumi.Input

	ToResourceFileOutput() ResourceFileOutput
	ToResourceFileOutputWithContext(context.Context) ResourceFileOutput
}

type ResourceFileArgs struct {
	AutoStorageContainerName pulumi.StringPtrInput `pulumi:"autoStorageContainerName"`
	BlobPrefix               pulumi.StringPtrInput `pulumi:"blobPrefix"`
	FileMode                 pulumi.StringPtrInput `pulumi:"fileMode"`
	FilePath                 pulumi.StringPtrInput `pulumi:"filePath"`
	HttpUrl                  pulumi.StringPtrInput `pulumi:"httpUrl"`
	StorageContainerUrl      pulumi.StringPtrInput `pulumi:"storageContainerUrl"`
}

func (ResourceFileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceFile)(nil)).Elem()
}

func (i ResourceFileArgs) ToResourceFileOutput() ResourceFileOutput {
	return i.ToResourceFileOutputWithContext(context.Background())
}

func (i ResourceFileArgs) ToResourceFileOutputWithContext(ctx context.Context) ResourceFileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceFileOutput)
}





type ResourceFileArrayInput interface {
	pulumi.Input

	ToResourceFileArrayOutput() ResourceFileArrayOutput
	ToResourceFileArrayOutputWithContext(context.Context) ResourceFileArrayOutput
}

type ResourceFileArray []ResourceFileInput

func (ResourceFileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceFile)(nil)).Elem()
}

func (i ResourceFileArray) ToResourceFileArrayOutput() ResourceFileArrayOutput {
	return i.ToResourceFileArrayOutputWithContext(context.Background())
}

func (i ResourceFileArray) ToResourceFileArrayOutputWithContext(ctx context.Context) ResourceFileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceFileArrayOutput)
}

type ResourceFileOutput struct{ *pulumi.OutputState }

func (ResourceFileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceFile)(nil)).Elem()
}

func (o ResourceFileOutput) ToResourceFileOutput() ResourceFileOutput {
	return o
}

func (o ResourceFileOutput) ToResourceFileOutputWithContext(ctx context.Context) ResourceFileOutput {
	return o
}

func (o ResourceFileOutput) AutoStorageContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceFile) *string { return v.AutoStorageContainerName }).(pulumi.StringPtrOutput)
}

func (o ResourceFileOutput) BlobPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceFile) *string { return v.BlobPrefix }).(pulumi.StringPtrOutput)
}

func (o ResourceFileOutput) FileMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceFile) *string { return v.FileMode }).(pulumi.StringPtrOutput)
}

func (o ResourceFileOutput) FilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceFile) *string { return v.FilePath }).(pulumi.StringPtrOutput)
}

func (o ResourceFileOutput) HttpUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceFile) *string { return v.HttpUrl }).(pulumi.StringPtrOutput)
}

func (o ResourceFileOutput) StorageContainerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceFile) *string { return v.StorageContainerUrl }).(pulumi.StringPtrOutput)
}

type ResourceFileArrayOutput struct{ *pulumi.OutputState }

func (ResourceFileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceFile)(nil)).Elem()
}

func (o ResourceFileArrayOutput) ToResourceFileArrayOutput() ResourceFileArrayOutput {
	return o
}

func (o ResourceFileArrayOutput) ToResourceFileArrayOutputWithContext(ctx context.Context) ResourceFileArrayOutput {
	return o
}

func (o ResourceFileArrayOutput) Index(i pulumi.IntInput) ResourceFileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceFile {
		return vs[0].([]ResourceFile)[vs[1].(int)]
	}).(ResourceFileOutput)
}

type ResourceFileResponse struct {
	AutoStorageContainerName *string `pulumi:"autoStorageContainerName"`
	BlobPrefix               *string `pulumi:"blobPrefix"`
	FileMode                 *string `pulumi:"fileMode"`
	FilePath                 *string `pulumi:"filePath"`
	HttpUrl                  *string `pulumi:"httpUrl"`
	StorageContainerUrl      *string `pulumi:"storageContainerUrl"`
}

type ResourceFileResponseOutput struct{ *pulumi.OutputState }

func (ResourceFileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceFileResponse)(nil)).Elem()
}

func (o ResourceFileResponseOutput) ToResourceFileResponseOutput() ResourceFileResponseOutput {
	return o
}

func (o ResourceFileResponseOutput) ToResourceFileResponseOutputWithContext(ctx context.Context) ResourceFileResponseOutput {
	return o
}

func (o ResourceFileResponseOutput) AutoStorageContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceFileResponse) *string { return v.AutoStorageContainerName }).(pulumi.StringPtrOutput)
}

func (o ResourceFileResponseOutput) BlobPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceFileResponse) *string { return v.BlobPrefix }).(pulumi.StringPtrOutput)
}

func (o ResourceFileResponseOutput) FileMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceFileResponse) *string { return v.FileMode }).(pulumi.StringPtrOutput)
}

func (o ResourceFileResponseOutput) FilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceFileResponse) *string { return v.FilePath }).(pulumi.StringPtrOutput)
}

func (o ResourceFileResponseOutput) HttpUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceFileResponse) *string { return v.HttpUrl }).(pulumi.StringPtrOutput)
}

func (o ResourceFileResponseOutput) StorageContainerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceFileResponse) *string { return v.StorageContainerUrl }).(pulumi.StringPtrOutput)
}

type ResourceFileResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceFileResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceFileResponse)(nil)).Elem()
}

func (o ResourceFileResponseArrayOutput) ToResourceFileResponseArrayOutput() ResourceFileResponseArrayOutput {
	return o
}

func (o ResourceFileResponseArrayOutput) ToResourceFileResponseArrayOutputWithContext(ctx context.Context) ResourceFileResponseArrayOutput {
	return o
}

func (o ResourceFileResponseArrayOutput) Index(i pulumi.IntInput) ResourceFileResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceFileResponse {
		return vs[0].([]ResourceFileResponse)[vs[1].(int)]
	}).(ResourceFileResponseOutput)
}

type ScaleSettings struct {
	AutoScale  *AutoScaleSettings  `pulumi:"autoScale"`
	FixedScale *FixedScaleSettings `pulumi:"fixedScale"`
}





type ScaleSettingsInput interface {
	pulumi.Input

	ToScaleSettingsOutput() ScaleSettingsOutput
	ToScaleSettingsOutputWithContext(context.Context) ScaleSettingsOutput
}

type ScaleSettingsArgs struct {
	AutoScale  AutoScaleSettingsPtrInput  `pulumi:"autoScale"`
	FixedScale FixedScaleSettingsPtrInput `pulumi:"fixedScale"`
}

func (ScaleSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleSettings)(nil)).Elem()
}

func (i ScaleSettingsArgs) ToScaleSettingsOutput() ScaleSettingsOutput {
	return i.ToScaleSettingsOutputWithContext(context.Background())
}

func (i ScaleSettingsArgs) ToScaleSettingsOutputWithContext(ctx context.Context) ScaleSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleSettingsOutput)
}

func (i ScaleSettingsArgs) ToScaleSettingsPtrOutput() ScaleSettingsPtrOutput {
	return i.ToScaleSettingsPtrOutputWithContext(context.Background())
}

func (i ScaleSettingsArgs) ToScaleSettingsPtrOutputWithContext(ctx context.Context) ScaleSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleSettingsOutput).ToScaleSettingsPtrOutputWithContext(ctx)
}









type ScaleSettingsPtrInput interface {
	pulumi.Input

	ToScaleSettingsPtrOutput() ScaleSettingsPtrOutput
	ToScaleSettingsPtrOutputWithContext(context.Context) ScaleSettingsPtrOutput
}

type scaleSettingsPtrType ScaleSettingsArgs

func ScaleSettingsPtr(v *ScaleSettingsArgs) ScaleSettingsPtrInput {
	return (*scaleSettingsPtrType)(v)
}

func (*scaleSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleSettings)(nil)).Elem()
}

func (i *scaleSettingsPtrType) ToScaleSettingsPtrOutput() ScaleSettingsPtrOutput {
	return i.ToScaleSettingsPtrOutputWithContext(context.Background())
}

func (i *scaleSettingsPtrType) ToScaleSettingsPtrOutputWithContext(ctx context.Context) ScaleSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleSettingsPtrOutput)
}

type ScaleSettingsOutput struct{ *pulumi.OutputState }

func (ScaleSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleSettings)(nil)).Elem()
}

func (o ScaleSettingsOutput) ToScaleSettingsOutput() ScaleSettingsOutput {
	return o
}

func (o ScaleSettingsOutput) ToScaleSettingsOutputWithContext(ctx context.Context) ScaleSettingsOutput {
	return o
}

func (o ScaleSettingsOutput) ToScaleSettingsPtrOutput() ScaleSettingsPtrOutput {
	return o.ToScaleSettingsPtrOutputWithContext(context.Background())
}

func (o ScaleSettingsOutput) ToScaleSettingsPtrOutputWithContext(ctx context.Context) ScaleSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScaleSettings) *ScaleSettings {
		return &v
	}).(ScaleSettingsPtrOutput)
}

func (o ScaleSettingsOutput) AutoScale() AutoScaleSettingsPtrOutput {
	return o.ApplyT(func(v ScaleSettings) *AutoScaleSettings { return v.AutoScale }).(AutoScaleSettingsPtrOutput)
}

func (o ScaleSettingsOutput) FixedScale() FixedScaleSettingsPtrOutput {
	return o.ApplyT(func(v ScaleSettings) *FixedScaleSettings { return v.FixedScale }).(FixedScaleSettingsPtrOutput)
}

type ScaleSettingsPtrOutput struct{ *pulumi.OutputState }

func (ScaleSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleSettings)(nil)).Elem()
}

func (o ScaleSettingsPtrOutput) ToScaleSettingsPtrOutput() ScaleSettingsPtrOutput {
	return o
}

func (o ScaleSettingsPtrOutput) ToScaleSettingsPtrOutputWithContext(ctx context.Context) ScaleSettingsPtrOutput {
	return o
}

func (o ScaleSettingsPtrOutput) Elem() ScaleSettingsOutput {
	return o.ApplyT(func(v *ScaleSettings) ScaleSettings {
		if v != nil {
			return *v
		}
		var ret ScaleSettings
		return ret
	}).(ScaleSettingsOutput)
}

func (o ScaleSettingsPtrOutput) AutoScale() AutoScaleSettingsPtrOutput {
	return o.ApplyT(func(v *ScaleSettings) *AutoScaleSettings {
		if v == nil {
			return nil
		}
		return v.AutoScale
	}).(AutoScaleSettingsPtrOutput)
}

func (o ScaleSettingsPtrOutput) FixedScale() FixedScaleSettingsPtrOutput {
	return o.ApplyT(func(v *ScaleSettings) *FixedScaleSettings {
		if v == nil {
			return nil
		}
		return v.FixedScale
	}).(FixedScaleSettingsPtrOutput)
}

type ScaleSettingsResponse struct {
	AutoScale  *AutoScaleSettingsResponse  `pulumi:"autoScale"`
	FixedScale *FixedScaleSettingsResponse `pulumi:"fixedScale"`
}

type ScaleSettingsResponseOutput struct{ *pulumi.OutputState }

func (ScaleSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleSettingsResponse)(nil)).Elem()
}

func (o ScaleSettingsResponseOutput) ToScaleSettingsResponseOutput() ScaleSettingsResponseOutput {
	return o
}

func (o ScaleSettingsResponseOutput) ToScaleSettingsResponseOutputWithContext(ctx context.Context) ScaleSettingsResponseOutput {
	return o
}

func (o ScaleSettingsResponseOutput) AutoScale() AutoScaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v ScaleSettingsResponse) *AutoScaleSettingsResponse { return v.AutoScale }).(AutoScaleSettingsResponsePtrOutput)
}

func (o ScaleSettingsResponseOutput) FixedScale() FixedScaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v ScaleSettingsResponse) *FixedScaleSettingsResponse { return v.FixedScale }).(FixedScaleSettingsResponsePtrOutput)
}

type ScaleSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ScaleSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleSettingsResponse)(nil)).Elem()
}

func (o ScaleSettingsResponsePtrOutput) ToScaleSettingsResponsePtrOutput() ScaleSettingsResponsePtrOutput {
	return o
}

func (o ScaleSettingsResponsePtrOutput) ToScaleSettingsResponsePtrOutputWithContext(ctx context.Context) ScaleSettingsResponsePtrOutput {
	return o
}

func (o ScaleSettingsResponsePtrOutput) Elem() ScaleSettingsResponseOutput {
	return o.ApplyT(func(v *ScaleSettingsResponse) ScaleSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ScaleSettingsResponse
		return ret
	}).(ScaleSettingsResponseOutput)
}

func (o ScaleSettingsResponsePtrOutput) AutoScale() AutoScaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ScaleSettingsResponse) *AutoScaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutoScale
	}).(AutoScaleSettingsResponsePtrOutput)
}

func (o ScaleSettingsResponsePtrOutput) FixedScale() FixedScaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ScaleSettingsResponse) *FixedScaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.FixedScale
	}).(FixedScaleSettingsResponsePtrOutput)
}

type StartTask struct {
	CommandLine         *string                `pulumi:"commandLine"`
	ContainerSettings   *TaskContainerSettings `pulumi:"containerSettings"`
	EnvironmentSettings []EnvironmentSetting   `pulumi:"environmentSettings"`
	MaxTaskRetryCount   *int                   `pulumi:"maxTaskRetryCount"`
	ResourceFiles       []ResourceFile         `pulumi:"resourceFiles"`
	UserIdentity        *UserIdentity          `pulumi:"userIdentity"`
	WaitForSuccess      *bool                  `pulumi:"waitForSuccess"`
}





type StartTaskInput interface {
	pulumi.Input

	ToStartTaskOutput() StartTaskOutput
	ToStartTaskOutputWithContext(context.Context) StartTaskOutput
}

type StartTaskArgs struct {
	CommandLine         pulumi.StringPtrInput         `pulumi:"commandLine"`
	ContainerSettings   TaskContainerSettingsPtrInput `pulumi:"containerSettings"`
	EnvironmentSettings EnvironmentSettingArrayInput  `pulumi:"environmentSettings"`
	MaxTaskRetryCount   pulumi.IntPtrInput            `pulumi:"maxTaskRetryCount"`
	ResourceFiles       ResourceFileArrayInput        `pulumi:"resourceFiles"`
	UserIdentity        UserIdentityPtrInput          `pulumi:"userIdentity"`
	WaitForSuccess      pulumi.BoolPtrInput           `pulumi:"waitForSuccess"`
}

func (StartTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StartTask)(nil)).Elem()
}

func (i StartTaskArgs) ToStartTaskOutput() StartTaskOutput {
	return i.ToStartTaskOutputWithContext(context.Background())
}

func (i StartTaskArgs) ToStartTaskOutputWithContext(ctx context.Context) StartTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartTaskOutput)
}

func (i StartTaskArgs) ToStartTaskPtrOutput() StartTaskPtrOutput {
	return i.ToStartTaskPtrOutputWithContext(context.Background())
}

func (i StartTaskArgs) ToStartTaskPtrOutputWithContext(ctx context.Context) StartTaskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartTaskOutput).ToStartTaskPtrOutputWithContext(ctx)
}









type StartTaskPtrInput interface {
	pulumi.Input

	ToStartTaskPtrOutput() StartTaskPtrOutput
	ToStartTaskPtrOutputWithContext(context.Context) StartTaskPtrOutput
}

type startTaskPtrType StartTaskArgs

func StartTaskPtr(v *StartTaskArgs) StartTaskPtrInput {
	return (*startTaskPtrType)(v)
}

func (*startTaskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StartTask)(nil)).Elem()
}

func (i *startTaskPtrType) ToStartTaskPtrOutput() StartTaskPtrOutput {
	return i.ToStartTaskPtrOutputWithContext(context.Background())
}

func (i *startTaskPtrType) ToStartTaskPtrOutputWithContext(ctx context.Context) StartTaskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartTaskPtrOutput)
}

type StartTaskOutput struct{ *pulumi.OutputState }

func (StartTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StartTask)(nil)).Elem()
}

func (o StartTaskOutput) ToStartTaskOutput() StartTaskOutput {
	return o
}

func (o StartTaskOutput) ToStartTaskOutputWithContext(ctx context.Context) StartTaskOutput {
	return o
}

func (o StartTaskOutput) ToStartTaskPtrOutput() StartTaskPtrOutput {
	return o.ToStartTaskPtrOutputWithContext(context.Background())
}

func (o StartTaskOutput) ToStartTaskPtrOutputWithContext(ctx context.Context) StartTaskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StartTask) *StartTask {
		return &v
	}).(StartTaskPtrOutput)
}

func (o StartTaskOutput) CommandLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StartTask) *string { return v.CommandLine }).(pulumi.StringPtrOutput)
}

func (o StartTaskOutput) ContainerSettings() TaskContainerSettingsPtrOutput {
	return o.ApplyT(func(v StartTask) *TaskContainerSettings { return v.ContainerSettings }).(TaskContainerSettingsPtrOutput)
}

func (o StartTaskOutput) EnvironmentSettings() EnvironmentSettingArrayOutput {
	return o.ApplyT(func(v StartTask) []EnvironmentSetting { return v.EnvironmentSettings }).(EnvironmentSettingArrayOutput)
}

func (o StartTaskOutput) MaxTaskRetryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StartTask) *int { return v.MaxTaskRetryCount }).(pulumi.IntPtrOutput)
}

func (o StartTaskOutput) ResourceFiles() ResourceFileArrayOutput {
	return o.ApplyT(func(v StartTask) []ResourceFile { return v.ResourceFiles }).(ResourceFileArrayOutput)
}

func (o StartTaskOutput) UserIdentity() UserIdentityPtrOutput {
	return o.ApplyT(func(v StartTask) *UserIdentity { return v.UserIdentity }).(UserIdentityPtrOutput)
}

func (o StartTaskOutput) WaitForSuccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StartTask) *bool { return v.WaitForSuccess }).(pulumi.BoolPtrOutput)
}

type StartTaskPtrOutput struct{ *pulumi.OutputState }

func (StartTaskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StartTask)(nil)).Elem()
}

func (o StartTaskPtrOutput) ToStartTaskPtrOutput() StartTaskPtrOutput {
	return o
}

func (o StartTaskPtrOutput) ToStartTaskPtrOutputWithContext(ctx context.Context) StartTaskPtrOutput {
	return o
}

func (o StartTaskPtrOutput) Elem() StartTaskOutput {
	return o.ApplyT(func(v *StartTask) StartTask {
		if v != nil {
			return *v
		}
		var ret StartTask
		return ret
	}).(StartTaskOutput)
}

func (o StartTaskPtrOutput) CommandLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StartTask) *string {
		if v == nil {
			return nil
		}
		return v.CommandLine
	}).(pulumi.StringPtrOutput)
}

func (o StartTaskPtrOutput) ContainerSettings() TaskContainerSettingsPtrOutput {
	return o.ApplyT(func(v *StartTask) *TaskContainerSettings {
		if v == nil {
			return nil
		}
		return v.ContainerSettings
	}).(TaskContainerSettingsPtrOutput)
}

func (o StartTaskPtrOutput) EnvironmentSettings() EnvironmentSettingArrayOutput {
	return o.ApplyT(func(v *StartTask) []EnvironmentSetting {
		if v == nil {
			return nil
		}
		return v.EnvironmentSettings
	}).(EnvironmentSettingArrayOutput)
}

func (o StartTaskPtrOutput) MaxTaskRetryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StartTask) *int {
		if v == nil {
			return nil
		}
		return v.MaxTaskRetryCount
	}).(pulumi.IntPtrOutput)
}

func (o StartTaskPtrOutput) ResourceFiles() ResourceFileArrayOutput {
	return o.ApplyT(func(v *StartTask) []ResourceFile {
		if v == nil {
			return nil
		}
		return v.ResourceFiles
	}).(ResourceFileArrayOutput)
}

func (o StartTaskPtrOutput) UserIdentity() UserIdentityPtrOutput {
	return o.ApplyT(func(v *StartTask) *UserIdentity {
		if v == nil {
			return nil
		}
		return v.UserIdentity
	}).(UserIdentityPtrOutput)
}

func (o StartTaskPtrOutput) WaitForSuccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StartTask) *bool {
		if v == nil {
			return nil
		}
		return v.WaitForSuccess
	}).(pulumi.BoolPtrOutput)
}

type StartTaskResponse struct {
	CommandLine         *string                        `pulumi:"commandLine"`
	ContainerSettings   *TaskContainerSettingsResponse `pulumi:"containerSettings"`
	EnvironmentSettings []EnvironmentSettingResponse   `pulumi:"environmentSettings"`
	MaxTaskRetryCount   *int                           `pulumi:"maxTaskRetryCount"`
	ResourceFiles       []ResourceFileResponse         `pulumi:"resourceFiles"`
	UserIdentity        *UserIdentityResponse          `pulumi:"userIdentity"`
	WaitForSuccess      *bool                          `pulumi:"waitForSuccess"`
}

type StartTaskResponseOutput struct{ *pulumi.OutputState }

func (StartTaskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StartTaskResponse)(nil)).Elem()
}

func (o StartTaskResponseOutput) ToStartTaskResponseOutput() StartTaskResponseOutput {
	return o
}

func (o StartTaskResponseOutput) ToStartTaskResponseOutputWithContext(ctx context.Context) StartTaskResponseOutput {
	return o
}

func (o StartTaskResponseOutput) CommandLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StartTaskResponse) *string { return v.CommandLine }).(pulumi.StringPtrOutput)
}

func (o StartTaskResponseOutput) ContainerSettings() TaskContainerSettingsResponsePtrOutput {
	return o.ApplyT(func(v StartTaskResponse) *TaskContainerSettingsResponse { return v.ContainerSettings }).(TaskContainerSettingsResponsePtrOutput)
}

func (o StartTaskResponseOutput) EnvironmentSettings() EnvironmentSettingResponseArrayOutput {
	return o.ApplyT(func(v StartTaskResponse) []EnvironmentSettingResponse { return v.EnvironmentSettings }).(EnvironmentSettingResponseArrayOutput)
}

func (o StartTaskResponseOutput) MaxTaskRetryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StartTaskResponse) *int { return v.MaxTaskRetryCount }).(pulumi.IntPtrOutput)
}

func (o StartTaskResponseOutput) ResourceFiles() ResourceFileResponseArrayOutput {
	return o.ApplyT(func(v StartTaskResponse) []ResourceFileResponse { return v.ResourceFiles }).(ResourceFileResponseArrayOutput)
}

func (o StartTaskResponseOutput) UserIdentity() UserIdentityResponsePtrOutput {
	return o.ApplyT(func(v StartTaskResponse) *UserIdentityResponse { return v.UserIdentity }).(UserIdentityResponsePtrOutput)
}

func (o StartTaskResponseOutput) WaitForSuccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StartTaskResponse) *bool { return v.WaitForSuccess }).(pulumi.BoolPtrOutput)
}

type StartTaskResponsePtrOutput struct{ *pulumi.OutputState }

func (StartTaskResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StartTaskResponse)(nil)).Elem()
}

func (o StartTaskResponsePtrOutput) ToStartTaskResponsePtrOutput() StartTaskResponsePtrOutput {
	return o
}

func (o StartTaskResponsePtrOutput) ToStartTaskResponsePtrOutputWithContext(ctx context.Context) StartTaskResponsePtrOutput {
	return o
}

func (o StartTaskResponsePtrOutput) Elem() StartTaskResponseOutput {
	return o.ApplyT(func(v *StartTaskResponse) StartTaskResponse {
		if v != nil {
			return *v
		}
		var ret StartTaskResponse
		return ret
	}).(StartTaskResponseOutput)
}

func (o StartTaskResponsePtrOutput) CommandLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StartTaskResponse) *string {
		if v == nil {
			return nil
		}
		return v.CommandLine
	}).(pulumi.StringPtrOutput)
}

func (o StartTaskResponsePtrOutput) ContainerSettings() TaskContainerSettingsResponsePtrOutput {
	return o.ApplyT(func(v *StartTaskResponse) *TaskContainerSettingsResponse {
		if v == nil {
			return nil
		}
		return v.ContainerSettings
	}).(TaskContainerSettingsResponsePtrOutput)
}

func (o StartTaskResponsePtrOutput) EnvironmentSettings() EnvironmentSettingResponseArrayOutput {
	return o.ApplyT(func(v *StartTaskResponse) []EnvironmentSettingResponse {
		if v == nil {
			return nil
		}
		return v.EnvironmentSettings
	}).(EnvironmentSettingResponseArrayOutput)
}

func (o StartTaskResponsePtrOutput) MaxTaskRetryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StartTaskResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxTaskRetryCount
	}).(pulumi.IntPtrOutput)
}

func (o StartTaskResponsePtrOutput) ResourceFiles() ResourceFileResponseArrayOutput {
	return o.ApplyT(func(v *StartTaskResponse) []ResourceFileResponse {
		if v == nil {
			return nil
		}
		return v.ResourceFiles
	}).(ResourceFileResponseArrayOutput)
}

func (o StartTaskResponsePtrOutput) UserIdentity() UserIdentityResponsePtrOutput {
	return o.ApplyT(func(v *StartTaskResponse) *UserIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserIdentity
	}).(UserIdentityResponsePtrOutput)
}

func (o StartTaskResponsePtrOutput) WaitForSuccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StartTaskResponse) *bool {
		if v == nil {
			return nil
		}
		return v.WaitForSuccess
	}).(pulumi.BoolPtrOutput)
}

type TaskContainerSettings struct {
	ContainerRunOptions *string            `pulumi:"containerRunOptions"`
	ImageName           string             `pulumi:"imageName"`
	Registry            *ContainerRegistry `pulumi:"registry"`
}





type TaskContainerSettingsInput interface {
	pulumi.Input

	ToTaskContainerSettingsOutput() TaskContainerSettingsOutput
	ToTaskContainerSettingsOutputWithContext(context.Context) TaskContainerSettingsOutput
}

type TaskContainerSettingsArgs struct {
	ContainerRunOptions pulumi.StringPtrInput     `pulumi:"containerRunOptions"`
	ImageName           pulumi.StringInput        `pulumi:"imageName"`
	Registry            ContainerRegistryPtrInput `pulumi:"registry"`
}

func (TaskContainerSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskContainerSettings)(nil)).Elem()
}

func (i TaskContainerSettingsArgs) ToTaskContainerSettingsOutput() TaskContainerSettingsOutput {
	return i.ToTaskContainerSettingsOutputWithContext(context.Background())
}

func (i TaskContainerSettingsArgs) ToTaskContainerSettingsOutputWithContext(ctx context.Context) TaskContainerSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskContainerSettingsOutput)
}

func (i TaskContainerSettingsArgs) ToTaskContainerSettingsPtrOutput() TaskContainerSettingsPtrOutput {
	return i.ToTaskContainerSettingsPtrOutputWithContext(context.Background())
}

func (i TaskContainerSettingsArgs) ToTaskContainerSettingsPtrOutputWithContext(ctx context.Context) TaskContainerSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskContainerSettingsOutput).ToTaskContainerSettingsPtrOutputWithContext(ctx)
}









type TaskContainerSettingsPtrInput interface {
	pulumi.Input

	ToTaskContainerSettingsPtrOutput() TaskContainerSettingsPtrOutput
	ToTaskContainerSettingsPtrOutputWithContext(context.Context) TaskContainerSettingsPtrOutput
}

type taskContainerSettingsPtrType TaskContainerSettingsArgs

func TaskContainerSettingsPtr(v *TaskContainerSettingsArgs) TaskContainerSettingsPtrInput {
	return (*taskContainerSettingsPtrType)(v)
}

func (*taskContainerSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskContainerSettings)(nil)).Elem()
}

func (i *taskContainerSettingsPtrType) ToTaskContainerSettingsPtrOutput() TaskContainerSettingsPtrOutput {
	return i.ToTaskContainerSettingsPtrOutputWithContext(context.Background())
}

func (i *taskContainerSettingsPtrType) ToTaskContainerSettingsPtrOutputWithContext(ctx context.Context) TaskContainerSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskContainerSettingsPtrOutput)
}

type TaskContainerSettingsOutput struct{ *pulumi.OutputState }

func (TaskContainerSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskContainerSettings)(nil)).Elem()
}

func (o TaskContainerSettingsOutput) ToTaskContainerSettingsOutput() TaskContainerSettingsOutput {
	return o
}

func (o TaskContainerSettingsOutput) ToTaskContainerSettingsOutputWithContext(ctx context.Context) TaskContainerSettingsOutput {
	return o
}

func (o TaskContainerSettingsOutput) ToTaskContainerSettingsPtrOutput() TaskContainerSettingsPtrOutput {
	return o.ToTaskContainerSettingsPtrOutputWithContext(context.Background())
}

func (o TaskContainerSettingsOutput) ToTaskContainerSettingsPtrOutputWithContext(ctx context.Context) TaskContainerSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TaskContainerSettings) *TaskContainerSettings {
		return &v
	}).(TaskContainerSettingsPtrOutput)
}

func (o TaskContainerSettingsOutput) ContainerRunOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskContainerSettings) *string { return v.ContainerRunOptions }).(pulumi.StringPtrOutput)
}

func (o TaskContainerSettingsOutput) ImageName() pulumi.StringOutput {
	return o.ApplyT(func(v TaskContainerSettings) string { return v.ImageName }).(pulumi.StringOutput)
}

func (o TaskContainerSettingsOutput) Registry() ContainerRegistryPtrOutput {
	return o.ApplyT(func(v TaskContainerSettings) *ContainerRegistry { return v.Registry }).(ContainerRegistryPtrOutput)
}

type TaskContainerSettingsPtrOutput struct{ *pulumi.OutputState }

func (TaskContainerSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskContainerSettings)(nil)).Elem()
}

func (o TaskContainerSettingsPtrOutput) ToTaskContainerSettingsPtrOutput() TaskContainerSettingsPtrOutput {
	return o
}

func (o TaskContainerSettingsPtrOutput) ToTaskContainerSettingsPtrOutputWithContext(ctx context.Context) TaskContainerSettingsPtrOutput {
	return o
}

func (o TaskContainerSettingsPtrOutput) Elem() TaskContainerSettingsOutput {
	return o.ApplyT(func(v *TaskContainerSettings) TaskContainerSettings {
		if v != nil {
			return *v
		}
		var ret TaskContainerSettings
		return ret
	}).(TaskContainerSettingsOutput)
}

func (o TaskContainerSettingsPtrOutput) ContainerRunOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskContainerSettings) *string {
		if v == nil {
			return nil
		}
		return v.ContainerRunOptions
	}).(pulumi.StringPtrOutput)
}

func (o TaskContainerSettingsPtrOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskContainerSettings) *string {
		if v == nil {
			return nil
		}
		return &v.ImageName
	}).(pulumi.StringPtrOutput)
}

func (o TaskContainerSettingsPtrOutput) Registry() ContainerRegistryPtrOutput {
	return o.ApplyT(func(v *TaskContainerSettings) *ContainerRegistry {
		if v == nil {
			return nil
		}
		return v.Registry
	}).(ContainerRegistryPtrOutput)
}

type TaskContainerSettingsResponse struct {
	ContainerRunOptions *string                    `pulumi:"containerRunOptions"`
	ImageName           string                     `pulumi:"imageName"`
	Registry            *ContainerRegistryResponse `pulumi:"registry"`
}

type TaskContainerSettingsResponseOutput struct{ *pulumi.OutputState }

func (TaskContainerSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskContainerSettingsResponse)(nil)).Elem()
}

func (o TaskContainerSettingsResponseOutput) ToTaskContainerSettingsResponseOutput() TaskContainerSettingsResponseOutput {
	return o
}

func (o TaskContainerSettingsResponseOutput) ToTaskContainerSettingsResponseOutputWithContext(ctx context.Context) TaskContainerSettingsResponseOutput {
	return o
}

func (o TaskContainerSettingsResponseOutput) ContainerRunOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskContainerSettingsResponse) *string { return v.ContainerRunOptions }).(pulumi.StringPtrOutput)
}

func (o TaskContainerSettingsResponseOutput) ImageName() pulumi.StringOutput {
	return o.ApplyT(func(v TaskContainerSettingsResponse) string { return v.ImageName }).(pulumi.StringOutput)
}

func (o TaskContainerSettingsResponseOutput) Registry() ContainerRegistryResponsePtrOutput {
	return o.ApplyT(func(v TaskContainerSettingsResponse) *ContainerRegistryResponse { return v.Registry }).(ContainerRegistryResponsePtrOutput)
}

type TaskContainerSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (TaskContainerSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskContainerSettingsResponse)(nil)).Elem()
}

func (o TaskContainerSettingsResponsePtrOutput) ToTaskContainerSettingsResponsePtrOutput() TaskContainerSettingsResponsePtrOutput {
	return o
}

func (o TaskContainerSettingsResponsePtrOutput) ToTaskContainerSettingsResponsePtrOutputWithContext(ctx context.Context) TaskContainerSettingsResponsePtrOutput {
	return o
}

func (o TaskContainerSettingsResponsePtrOutput) Elem() TaskContainerSettingsResponseOutput {
	return o.ApplyT(func(v *TaskContainerSettingsResponse) TaskContainerSettingsResponse {
		if v != nil {
			return *v
		}
		var ret TaskContainerSettingsResponse
		return ret
	}).(TaskContainerSettingsResponseOutput)
}

func (o TaskContainerSettingsResponsePtrOutput) ContainerRunOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskContainerSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContainerRunOptions
	}).(pulumi.StringPtrOutput)
}

func (o TaskContainerSettingsResponsePtrOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskContainerSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ImageName
	}).(pulumi.StringPtrOutput)
}

func (o TaskContainerSettingsResponsePtrOutput) Registry() ContainerRegistryResponsePtrOutput {
	return o.ApplyT(func(v *TaskContainerSettingsResponse) *ContainerRegistryResponse {
		if v == nil {
			return nil
		}
		return v.Registry
	}).(ContainerRegistryResponsePtrOutput)
}

type TaskSchedulingPolicy struct {
	NodeFillType ComputeNodeFillType `pulumi:"nodeFillType"`
}





type TaskSchedulingPolicyInput interface {
	pulumi.Input

	ToTaskSchedulingPolicyOutput() TaskSchedulingPolicyOutput
	ToTaskSchedulingPolicyOutputWithContext(context.Context) TaskSchedulingPolicyOutput
}

type TaskSchedulingPolicyArgs struct {
	NodeFillType ComputeNodeFillTypeInput `pulumi:"nodeFillType"`
}

func (TaskSchedulingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskSchedulingPolicy)(nil)).Elem()
}

func (i TaskSchedulingPolicyArgs) ToTaskSchedulingPolicyOutput() TaskSchedulingPolicyOutput {
	return i.ToTaskSchedulingPolicyOutputWithContext(context.Background())
}

func (i TaskSchedulingPolicyArgs) ToTaskSchedulingPolicyOutputWithContext(ctx context.Context) TaskSchedulingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskSchedulingPolicyOutput)
}

func (i TaskSchedulingPolicyArgs) ToTaskSchedulingPolicyPtrOutput() TaskSchedulingPolicyPtrOutput {
	return i.ToTaskSchedulingPolicyPtrOutputWithContext(context.Background())
}

func (i TaskSchedulingPolicyArgs) ToTaskSchedulingPolicyPtrOutputWithContext(ctx context.Context) TaskSchedulingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskSchedulingPolicyOutput).ToTaskSchedulingPolicyPtrOutputWithContext(ctx)
}









type TaskSchedulingPolicyPtrInput interface {
	pulumi.Input

	ToTaskSchedulingPolicyPtrOutput() TaskSchedulingPolicyPtrOutput
	ToTaskSchedulingPolicyPtrOutputWithContext(context.Context) TaskSchedulingPolicyPtrOutput
}

type taskSchedulingPolicyPtrType TaskSchedulingPolicyArgs

func TaskSchedulingPolicyPtr(v *TaskSchedulingPolicyArgs) TaskSchedulingPolicyPtrInput {
	return (*taskSchedulingPolicyPtrType)(v)
}

func (*taskSchedulingPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskSchedulingPolicy)(nil)).Elem()
}

func (i *taskSchedulingPolicyPtrType) ToTaskSchedulingPolicyPtrOutput() TaskSchedulingPolicyPtrOutput {
	return i.ToTaskSchedulingPolicyPtrOutputWithContext(context.Background())
}

func (i *taskSchedulingPolicyPtrType) ToTaskSchedulingPolicyPtrOutputWithContext(ctx context.Context) TaskSchedulingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskSchedulingPolicyPtrOutput)
}

type TaskSchedulingPolicyOutput struct{ *pulumi.OutputState }

func (TaskSchedulingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskSchedulingPolicy)(nil)).Elem()
}

func (o TaskSchedulingPolicyOutput) ToTaskSchedulingPolicyOutput() TaskSchedulingPolicyOutput {
	return o
}

func (o TaskSchedulingPolicyOutput) ToTaskSchedulingPolicyOutputWithContext(ctx context.Context) TaskSchedulingPolicyOutput {
	return o
}

func (o TaskSchedulingPolicyOutput) ToTaskSchedulingPolicyPtrOutput() TaskSchedulingPolicyPtrOutput {
	return o.ToTaskSchedulingPolicyPtrOutputWithContext(context.Background())
}

func (o TaskSchedulingPolicyOutput) ToTaskSchedulingPolicyPtrOutputWithContext(ctx context.Context) TaskSchedulingPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TaskSchedulingPolicy) *TaskSchedulingPolicy {
		return &v
	}).(TaskSchedulingPolicyPtrOutput)
}

func (o TaskSchedulingPolicyOutput) NodeFillType() ComputeNodeFillTypeOutput {
	return o.ApplyT(func(v TaskSchedulingPolicy) ComputeNodeFillType { return v.NodeFillType }).(ComputeNodeFillTypeOutput)
}

type TaskSchedulingPolicyPtrOutput struct{ *pulumi.OutputState }

func (TaskSchedulingPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskSchedulingPolicy)(nil)).Elem()
}

func (o TaskSchedulingPolicyPtrOutput) ToTaskSchedulingPolicyPtrOutput() TaskSchedulingPolicyPtrOutput {
	return o
}

func (o TaskSchedulingPolicyPtrOutput) ToTaskSchedulingPolicyPtrOutputWithContext(ctx context.Context) TaskSchedulingPolicyPtrOutput {
	return o
}

func (o TaskSchedulingPolicyPtrOutput) Elem() TaskSchedulingPolicyOutput {
	return o.ApplyT(func(v *TaskSchedulingPolicy) TaskSchedulingPolicy {
		if v != nil {
			return *v
		}
		var ret TaskSchedulingPolicy
		return ret
	}).(TaskSchedulingPolicyOutput)
}

func (o TaskSchedulingPolicyPtrOutput) NodeFillType() ComputeNodeFillTypePtrOutput {
	return o.ApplyT(func(v *TaskSchedulingPolicy) *ComputeNodeFillType {
		if v == nil {
			return nil
		}
		return &v.NodeFillType
	}).(ComputeNodeFillTypePtrOutput)
}

type TaskSchedulingPolicyResponse struct {
	NodeFillType string `pulumi:"nodeFillType"`
}

type TaskSchedulingPolicyResponseOutput struct{ *pulumi.OutputState }

func (TaskSchedulingPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskSchedulingPolicyResponse)(nil)).Elem()
}

func (o TaskSchedulingPolicyResponseOutput) ToTaskSchedulingPolicyResponseOutput() TaskSchedulingPolicyResponseOutput {
	return o
}

func (o TaskSchedulingPolicyResponseOutput) ToTaskSchedulingPolicyResponseOutputWithContext(ctx context.Context) TaskSchedulingPolicyResponseOutput {
	return o
}

func (o TaskSchedulingPolicyResponseOutput) NodeFillType() pulumi.StringOutput {
	return o.ApplyT(func(v TaskSchedulingPolicyResponse) string { return v.NodeFillType }).(pulumi.StringOutput)
}

type TaskSchedulingPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (TaskSchedulingPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskSchedulingPolicyResponse)(nil)).Elem()
}

func (o TaskSchedulingPolicyResponsePtrOutput) ToTaskSchedulingPolicyResponsePtrOutput() TaskSchedulingPolicyResponsePtrOutput {
	return o
}

func (o TaskSchedulingPolicyResponsePtrOutput) ToTaskSchedulingPolicyResponsePtrOutputWithContext(ctx context.Context) TaskSchedulingPolicyResponsePtrOutput {
	return o
}

func (o TaskSchedulingPolicyResponsePtrOutput) Elem() TaskSchedulingPolicyResponseOutput {
	return o.ApplyT(func(v *TaskSchedulingPolicyResponse) TaskSchedulingPolicyResponse {
		if v != nil {
			return *v
		}
		var ret TaskSchedulingPolicyResponse
		return ret
	}).(TaskSchedulingPolicyResponseOutput)
}

func (o TaskSchedulingPolicyResponsePtrOutput) NodeFillType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskSchedulingPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.NodeFillType
	}).(pulumi.StringPtrOutput)
}

type UserAccount struct {
	ElevationLevel           *ElevationLevel           `pulumi:"elevationLevel"`
	LinuxUserConfiguration   *LinuxUserConfiguration   `pulumi:"linuxUserConfiguration"`
	Name                     string                    `pulumi:"name"`
	Password                 string                    `pulumi:"password"`
	WindowsUserConfiguration *WindowsUserConfiguration `pulumi:"windowsUserConfiguration"`
}





type UserAccountInput interface {
	pulumi.Input

	ToUserAccountOutput() UserAccountOutput
	ToUserAccountOutputWithContext(context.Context) UserAccountOutput
}

type UserAccountArgs struct {
	ElevationLevel           ElevationLevelPtrInput           `pulumi:"elevationLevel"`
	LinuxUserConfiguration   LinuxUserConfigurationPtrInput   `pulumi:"linuxUserConfiguration"`
	Name                     pulumi.StringInput               `pulumi:"name"`
	Password                 pulumi.StringInput               `pulumi:"password"`
	WindowsUserConfiguration WindowsUserConfigurationPtrInput `pulumi:"windowsUserConfiguration"`
}

func (UserAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAccount)(nil)).Elem()
}

func (i UserAccountArgs) ToUserAccountOutput() UserAccountOutput {
	return i.ToUserAccountOutputWithContext(context.Background())
}

func (i UserAccountArgs) ToUserAccountOutputWithContext(ctx context.Context) UserAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAccountOutput)
}





type UserAccountArrayInput interface {
	pulumi.Input

	ToUserAccountArrayOutput() UserAccountArrayOutput
	ToUserAccountArrayOutputWithContext(context.Context) UserAccountArrayOutput
}

type UserAccountArray []UserAccountInput

func (UserAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserAccount)(nil)).Elem()
}

func (i UserAccountArray) ToUserAccountArrayOutput() UserAccountArrayOutput {
	return i.ToUserAccountArrayOutputWithContext(context.Background())
}

func (i UserAccountArray) ToUserAccountArrayOutputWithContext(ctx context.Context) UserAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAccountArrayOutput)
}

type UserAccountOutput struct{ *pulumi.OutputState }

func (UserAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAccount)(nil)).Elem()
}

func (o UserAccountOutput) ToUserAccountOutput() UserAccountOutput {
	return o
}

func (o UserAccountOutput) ToUserAccountOutputWithContext(ctx context.Context) UserAccountOutput {
	return o
}

func (o UserAccountOutput) ElevationLevel() ElevationLevelPtrOutput {
	return o.ApplyT(func(v UserAccount) *ElevationLevel { return v.ElevationLevel }).(ElevationLevelPtrOutput)
}

func (o UserAccountOutput) LinuxUserConfiguration() LinuxUserConfigurationPtrOutput {
	return o.ApplyT(func(v UserAccount) *LinuxUserConfiguration { return v.LinuxUserConfiguration }).(LinuxUserConfigurationPtrOutput)
}

func (o UserAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v UserAccount) string { return v.Name }).(pulumi.StringOutput)
}

func (o UserAccountOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v UserAccount) string { return v.Password }).(pulumi.StringOutput)
}

func (o UserAccountOutput) WindowsUserConfiguration() WindowsUserConfigurationPtrOutput {
	return o.ApplyT(func(v UserAccount) *WindowsUserConfiguration { return v.WindowsUserConfiguration }).(WindowsUserConfigurationPtrOutput)
}

type UserAccountArrayOutput struct{ *pulumi.OutputState }

func (UserAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserAccount)(nil)).Elem()
}

func (o UserAccountArrayOutput) ToUserAccountArrayOutput() UserAccountArrayOutput {
	return o
}

func (o UserAccountArrayOutput) ToUserAccountArrayOutputWithContext(ctx context.Context) UserAccountArrayOutput {
	return o
}

func (o UserAccountArrayOutput) Index(i pulumi.IntInput) UserAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserAccount {
		return vs[0].([]UserAccount)[vs[1].(int)]
	}).(UserAccountOutput)
}

type UserAccountResponse struct {
	ElevationLevel           *string                           `pulumi:"elevationLevel"`
	LinuxUserConfiguration   *LinuxUserConfigurationResponse   `pulumi:"linuxUserConfiguration"`
	Name                     string                            `pulumi:"name"`
	Password                 string                            `pulumi:"password"`
	WindowsUserConfiguration *WindowsUserConfigurationResponse `pulumi:"windowsUserConfiguration"`
}

type UserAccountResponseOutput struct{ *pulumi.OutputState }

func (UserAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAccountResponse)(nil)).Elem()
}

func (o UserAccountResponseOutput) ToUserAccountResponseOutput() UserAccountResponseOutput {
	return o
}

func (o UserAccountResponseOutput) ToUserAccountResponseOutputWithContext(ctx context.Context) UserAccountResponseOutput {
	return o
}

func (o UserAccountResponseOutput) ElevationLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAccountResponse) *string { return v.ElevationLevel }).(pulumi.StringPtrOutput)
}

func (o UserAccountResponseOutput) LinuxUserConfiguration() LinuxUserConfigurationResponsePtrOutput {
	return o.ApplyT(func(v UserAccountResponse) *LinuxUserConfigurationResponse { return v.LinuxUserConfiguration }).(LinuxUserConfigurationResponsePtrOutput)
}

func (o UserAccountResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v UserAccountResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o UserAccountResponseOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v UserAccountResponse) string { return v.Password }).(pulumi.StringOutput)
}

func (o UserAccountResponseOutput) WindowsUserConfiguration() WindowsUserConfigurationResponsePtrOutput {
	return o.ApplyT(func(v UserAccountResponse) *WindowsUserConfigurationResponse { return v.WindowsUserConfiguration }).(WindowsUserConfigurationResponsePtrOutput)
}

type UserAccountResponseArrayOutput struct{ *pulumi.OutputState }

func (UserAccountResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserAccountResponse)(nil)).Elem()
}

func (o UserAccountResponseArrayOutput) ToUserAccountResponseArrayOutput() UserAccountResponseArrayOutput {
	return o
}

func (o UserAccountResponseArrayOutput) ToUserAccountResponseArrayOutputWithContext(ctx context.Context) UserAccountResponseArrayOutput {
	return o
}

func (o UserAccountResponseArrayOutput) Index(i pulumi.IntInput) UserAccountResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserAccountResponse {
		return vs[0].([]UserAccountResponse)[vs[1].(int)]
	}).(UserAccountResponseOutput)
}

type UserIdentity struct {
	AutoUser *AutoUserSpecification `pulumi:"autoUser"`
	UserName *string                `pulumi:"userName"`
}





type UserIdentityInput interface {
	pulumi.Input

	ToUserIdentityOutput() UserIdentityOutput
	ToUserIdentityOutputWithContext(context.Context) UserIdentityOutput
}

type UserIdentityArgs struct {
	AutoUser AutoUserSpecificationPtrInput `pulumi:"autoUser"`
	UserName pulumi.StringPtrInput         `pulumi:"userName"`
}

func (UserIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentity)(nil)).Elem()
}

func (i UserIdentityArgs) ToUserIdentityOutput() UserIdentityOutput {
	return i.ToUserIdentityOutputWithContext(context.Background())
}

func (i UserIdentityArgs) ToUserIdentityOutputWithContext(ctx context.Context) UserIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityOutput)
}

func (i UserIdentityArgs) ToUserIdentityPtrOutput() UserIdentityPtrOutput {
	return i.ToUserIdentityPtrOutputWithContext(context.Background())
}

func (i UserIdentityArgs) ToUserIdentityPtrOutputWithContext(ctx context.Context) UserIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityOutput).ToUserIdentityPtrOutputWithContext(ctx)
}









type UserIdentityPtrInput interface {
	pulumi.Input

	ToUserIdentityPtrOutput() UserIdentityPtrOutput
	ToUserIdentityPtrOutputWithContext(context.Context) UserIdentityPtrOutput
}

type userIdentityPtrType UserIdentityArgs

func UserIdentityPtr(v *UserIdentityArgs) UserIdentityPtrInput {
	return (*userIdentityPtrType)(v)
}

func (*userIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserIdentity)(nil)).Elem()
}

func (i *userIdentityPtrType) ToUserIdentityPtrOutput() UserIdentityPtrOutput {
	return i.ToUserIdentityPtrOutputWithContext(context.Background())
}

func (i *userIdentityPtrType) ToUserIdentityPtrOutputWithContext(ctx context.Context) UserIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityPtrOutput)
}

type UserIdentityOutput struct{ *pulumi.OutputState }

func (UserIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentity)(nil)).Elem()
}

func (o UserIdentityOutput) ToUserIdentityOutput() UserIdentityOutput {
	return o
}

func (o UserIdentityOutput) ToUserIdentityOutputWithContext(ctx context.Context) UserIdentityOutput {
	return o
}

func (o UserIdentityOutput) ToUserIdentityPtrOutput() UserIdentityPtrOutput {
	return o.ToUserIdentityPtrOutputWithContext(context.Background())
}

func (o UserIdentityOutput) ToUserIdentityPtrOutputWithContext(ctx context.Context) UserIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserIdentity) *UserIdentity {
		return &v
	}).(UserIdentityPtrOutput)
}

func (o UserIdentityOutput) AutoUser() AutoUserSpecificationPtrOutput {
	return o.ApplyT(func(v UserIdentity) *AutoUserSpecification { return v.AutoUser }).(AutoUserSpecificationPtrOutput)
}

func (o UserIdentityOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentity) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

type UserIdentityPtrOutput struct{ *pulumi.OutputState }

func (UserIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserIdentity)(nil)).Elem()
}

func (o UserIdentityPtrOutput) ToUserIdentityPtrOutput() UserIdentityPtrOutput {
	return o
}

func (o UserIdentityPtrOutput) ToUserIdentityPtrOutputWithContext(ctx context.Context) UserIdentityPtrOutput {
	return o
}

func (o UserIdentityPtrOutput) Elem() UserIdentityOutput {
	return o.ApplyT(func(v *UserIdentity) UserIdentity {
		if v != nil {
			return *v
		}
		var ret UserIdentity
		return ret
	}).(UserIdentityOutput)
}

func (o UserIdentityPtrOutput) AutoUser() AutoUserSpecificationPtrOutput {
	return o.ApplyT(func(v *UserIdentity) *AutoUserSpecification {
		if v == nil {
			return nil
		}
		return v.AutoUser
	}).(AutoUserSpecificationPtrOutput)
}

func (o UserIdentityPtrOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserIdentity) *string {
		if v == nil {
			return nil
		}
		return v.UserName
	}).(pulumi.StringPtrOutput)
}

type UserIdentityResponse struct {
	AutoUser *AutoUserSpecificationResponse `pulumi:"autoUser"`
	UserName *string                        `pulumi:"userName"`
}

type UserIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityResponse)(nil)).Elem()
}

func (o UserIdentityResponseOutput) ToUserIdentityResponseOutput() UserIdentityResponseOutput {
	return o
}

func (o UserIdentityResponseOutput) ToUserIdentityResponseOutputWithContext(ctx context.Context) UserIdentityResponseOutput {
	return o
}

func (o UserIdentityResponseOutput) AutoUser() AutoUserSpecificationResponsePtrOutput {
	return o.ApplyT(func(v UserIdentityResponse) *AutoUserSpecificationResponse { return v.AutoUser }).(AutoUserSpecificationResponsePtrOutput)
}

func (o UserIdentityResponseOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityResponse) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

type UserIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (UserIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserIdentityResponse)(nil)).Elem()
}

func (o UserIdentityResponsePtrOutput) ToUserIdentityResponsePtrOutput() UserIdentityResponsePtrOutput {
	return o
}

func (o UserIdentityResponsePtrOutput) ToUserIdentityResponsePtrOutputWithContext(ctx context.Context) UserIdentityResponsePtrOutput {
	return o
}

func (o UserIdentityResponsePtrOutput) Elem() UserIdentityResponseOutput {
	return o.ApplyT(func(v *UserIdentityResponse) UserIdentityResponse {
		if v != nil {
			return *v
		}
		var ret UserIdentityResponse
		return ret
	}).(UserIdentityResponseOutput)
}

func (o UserIdentityResponsePtrOutput) AutoUser() AutoUserSpecificationResponsePtrOutput {
	return o.ApplyT(func(v *UserIdentityResponse) *AutoUserSpecificationResponse {
		if v == nil {
			return nil
		}
		return v.AutoUser
	}).(AutoUserSpecificationResponsePtrOutput)
}

func (o UserIdentityResponsePtrOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserName
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineConfiguration struct {
	ContainerConfiguration *ContainerConfiguration `pulumi:"containerConfiguration"`
	DataDisks              []DataDisk              `pulumi:"dataDisks"`
	ImageReference         ImageReference          `pulumi:"imageReference"`
	LicenseType            *string                 `pulumi:"licenseType"`
	NodeAgentSkuId         string                  `pulumi:"nodeAgentSkuId"`
	WindowsConfiguration   *WindowsConfiguration   `pulumi:"windowsConfiguration"`
}





type VirtualMachineConfigurationInput interface {
	pulumi.Input

	ToVirtualMachineConfigurationOutput() VirtualMachineConfigurationOutput
	ToVirtualMachineConfigurationOutputWithContext(context.Context) VirtualMachineConfigurationOutput
}

type VirtualMachineConfigurationArgs struct {
	ContainerConfiguration ContainerConfigurationPtrInput `pulumi:"containerConfiguration"`
	DataDisks              DataDiskArrayInput             `pulumi:"dataDisks"`
	ImageReference         ImageReferenceInput            `pulumi:"imageReference"`
	LicenseType            pulumi.StringPtrInput          `pulumi:"licenseType"`
	NodeAgentSkuId         pulumi.StringInput             `pulumi:"nodeAgentSkuId"`
	WindowsConfiguration   WindowsConfigurationPtrInput   `pulumi:"windowsConfiguration"`
}

func (VirtualMachineConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineConfiguration)(nil)).Elem()
}

func (i VirtualMachineConfigurationArgs) ToVirtualMachineConfigurationOutput() VirtualMachineConfigurationOutput {
	return i.ToVirtualMachineConfigurationOutputWithContext(context.Background())
}

func (i VirtualMachineConfigurationArgs) ToVirtualMachineConfigurationOutputWithContext(ctx context.Context) VirtualMachineConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineConfigurationOutput)
}

func (i VirtualMachineConfigurationArgs) ToVirtualMachineConfigurationPtrOutput() VirtualMachineConfigurationPtrOutput {
	return i.ToVirtualMachineConfigurationPtrOutputWithContext(context.Background())
}

func (i VirtualMachineConfigurationArgs) ToVirtualMachineConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineConfigurationOutput).ToVirtualMachineConfigurationPtrOutputWithContext(ctx)
}









type VirtualMachineConfigurationPtrInput interface {
	pulumi.Input

	ToVirtualMachineConfigurationPtrOutput() VirtualMachineConfigurationPtrOutput
	ToVirtualMachineConfigurationPtrOutputWithContext(context.Context) VirtualMachineConfigurationPtrOutput
}

type virtualMachineConfigurationPtrType VirtualMachineConfigurationArgs

func VirtualMachineConfigurationPtr(v *VirtualMachineConfigurationArgs) VirtualMachineConfigurationPtrInput {
	return (*virtualMachineConfigurationPtrType)(v)
}

func (*virtualMachineConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineConfiguration)(nil)).Elem()
}

func (i *virtualMachineConfigurationPtrType) ToVirtualMachineConfigurationPtrOutput() VirtualMachineConfigurationPtrOutput {
	return i.ToVirtualMachineConfigurationPtrOutputWithContext(context.Background())
}

func (i *virtualMachineConfigurationPtrType) ToVirtualMachineConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineConfigurationPtrOutput)
}

type VirtualMachineConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualMachineConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineConfiguration)(nil)).Elem()
}

func (o VirtualMachineConfigurationOutput) ToVirtualMachineConfigurationOutput() VirtualMachineConfigurationOutput {
	return o
}

func (o VirtualMachineConfigurationOutput) ToVirtualMachineConfigurationOutputWithContext(ctx context.Context) VirtualMachineConfigurationOutput {
	return o
}

func (o VirtualMachineConfigurationOutput) ToVirtualMachineConfigurationPtrOutput() VirtualMachineConfigurationPtrOutput {
	return o.ToVirtualMachineConfigurationPtrOutputWithContext(context.Background())
}

func (o VirtualMachineConfigurationOutput) ToVirtualMachineConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineConfiguration) *VirtualMachineConfiguration {
		return &v
	}).(VirtualMachineConfigurationPtrOutput)
}

func (o VirtualMachineConfigurationOutput) ContainerConfiguration() ContainerConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualMachineConfiguration) *ContainerConfiguration { return v.ContainerConfiguration }).(ContainerConfigurationPtrOutput)
}

func (o VirtualMachineConfigurationOutput) DataDisks() DataDiskArrayOutput {
	return o.ApplyT(func(v VirtualMachineConfiguration) []DataDisk { return v.DataDisks }).(DataDiskArrayOutput)
}

func (o VirtualMachineConfigurationOutput) ImageReference() ImageReferenceOutput {
	return o.ApplyT(func(v VirtualMachineConfiguration) ImageReference { return v.ImageReference }).(ImageReferenceOutput)
}

func (o VirtualMachineConfigurationOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineConfiguration) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineConfigurationOutput) NodeAgentSkuId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineConfiguration) string { return v.NodeAgentSkuId }).(pulumi.StringOutput)
}

func (o VirtualMachineConfigurationOutput) WindowsConfiguration() WindowsConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualMachineConfiguration) *WindowsConfiguration { return v.WindowsConfiguration }).(WindowsConfigurationPtrOutput)
}

type VirtualMachineConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineConfiguration)(nil)).Elem()
}

func (o VirtualMachineConfigurationPtrOutput) ToVirtualMachineConfigurationPtrOutput() VirtualMachineConfigurationPtrOutput {
	return o
}

func (o VirtualMachineConfigurationPtrOutput) ToVirtualMachineConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineConfigurationPtrOutput {
	return o
}

func (o VirtualMachineConfigurationPtrOutput) Elem() VirtualMachineConfigurationOutput {
	return o.ApplyT(func(v *VirtualMachineConfiguration) VirtualMachineConfiguration {
		if v != nil {
			return *v
		}
		var ret VirtualMachineConfiguration
		return ret
	}).(VirtualMachineConfigurationOutput)
}

func (o VirtualMachineConfigurationPtrOutput) ContainerConfiguration() ContainerConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualMachineConfiguration) *ContainerConfiguration {
		if v == nil {
			return nil
		}
		return v.ContainerConfiguration
	}).(ContainerConfigurationPtrOutput)
}

func (o VirtualMachineConfigurationPtrOutput) DataDisks() DataDiskArrayOutput {
	return o.ApplyT(func(v *VirtualMachineConfiguration) []DataDisk {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(DataDiskArrayOutput)
}

func (o VirtualMachineConfigurationPtrOutput) ImageReference() ImageReferencePtrOutput {
	return o.ApplyT(func(v *VirtualMachineConfiguration) *ImageReference {
		if v == nil {
			return nil
		}
		return &v.ImageReference
	}).(ImageReferencePtrOutput)
}

func (o VirtualMachineConfigurationPtrOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.LicenseType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineConfigurationPtrOutput) NodeAgentSkuId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.NodeAgentSkuId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineConfigurationPtrOutput) WindowsConfiguration() WindowsConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualMachineConfiguration) *WindowsConfiguration {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(WindowsConfigurationPtrOutput)
}

type VirtualMachineConfigurationResponse struct {
	ContainerConfiguration *ContainerConfigurationResponse `pulumi:"containerConfiguration"`
	DataDisks              []DataDiskResponse              `pulumi:"dataDisks"`
	ImageReference         ImageReferenceResponse          `pulumi:"imageReference"`
	LicenseType            *string                         `pulumi:"licenseType"`
	NodeAgentSkuId         string                          `pulumi:"nodeAgentSkuId"`
	WindowsConfiguration   *WindowsConfigurationResponse   `pulumi:"windowsConfiguration"`
}

type VirtualMachineConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineConfigurationResponseOutput) ToVirtualMachineConfigurationResponseOutput() VirtualMachineConfigurationResponseOutput {
	return o
}

func (o VirtualMachineConfigurationResponseOutput) ToVirtualMachineConfigurationResponseOutputWithContext(ctx context.Context) VirtualMachineConfigurationResponseOutput {
	return o
}

func (o VirtualMachineConfigurationResponseOutput) ContainerConfiguration() ContainerConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineConfigurationResponse) *ContainerConfigurationResponse {
		return v.ContainerConfiguration
	}).(ContainerConfigurationResponsePtrOutput)
}

func (o VirtualMachineConfigurationResponseOutput) DataDisks() DataDiskResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineConfigurationResponse) []DataDiskResponse { return v.DataDisks }).(DataDiskResponseArrayOutput)
}

func (o VirtualMachineConfigurationResponseOutput) ImageReference() ImageReferenceResponseOutput {
	return o.ApplyT(func(v VirtualMachineConfigurationResponse) ImageReferenceResponse { return v.ImageReference }).(ImageReferenceResponseOutput)
}

func (o VirtualMachineConfigurationResponseOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineConfigurationResponse) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineConfigurationResponseOutput) NodeAgentSkuId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineConfigurationResponse) string { return v.NodeAgentSkuId }).(pulumi.StringOutput)
}

func (o VirtualMachineConfigurationResponseOutput) WindowsConfiguration() WindowsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineConfigurationResponse) *WindowsConfigurationResponse {
		return v.WindowsConfiguration
	}).(WindowsConfigurationResponsePtrOutput)
}

type VirtualMachineConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineConfigurationResponsePtrOutput) ToVirtualMachineConfigurationResponsePtrOutput() VirtualMachineConfigurationResponsePtrOutput {
	return o
}

func (o VirtualMachineConfigurationResponsePtrOutput) ToVirtualMachineConfigurationResponsePtrOutputWithContext(ctx context.Context) VirtualMachineConfigurationResponsePtrOutput {
	return o
}

func (o VirtualMachineConfigurationResponsePtrOutput) Elem() VirtualMachineConfigurationResponseOutput {
	return o.ApplyT(func(v *VirtualMachineConfigurationResponse) VirtualMachineConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineConfigurationResponse
		return ret
	}).(VirtualMachineConfigurationResponseOutput)
}

func (o VirtualMachineConfigurationResponsePtrOutput) ContainerConfiguration() ContainerConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineConfigurationResponse) *ContainerConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.ContainerConfiguration
	}).(ContainerConfigurationResponsePtrOutput)
}

func (o VirtualMachineConfigurationResponsePtrOutput) DataDisks() DataDiskResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineConfigurationResponse) []DataDiskResponse {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(DataDiskResponseArrayOutput)
}

func (o VirtualMachineConfigurationResponsePtrOutput) ImageReference() ImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineConfigurationResponse) *ImageReferenceResponse {
		if v == nil {
			return nil
		}
		return &v.ImageReference
	}).(ImageReferenceResponsePtrOutput)
}

func (o VirtualMachineConfigurationResponsePtrOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.LicenseType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineConfigurationResponsePtrOutput) NodeAgentSkuId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.NodeAgentSkuId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineConfigurationResponsePtrOutput) WindowsConfiguration() WindowsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineConfigurationResponse) *WindowsConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(WindowsConfigurationResponsePtrOutput)
}

type WindowsConfiguration struct {
	EnableAutomaticUpdates *bool `pulumi:"enableAutomaticUpdates"`
}





type WindowsConfigurationInput interface {
	pulumi.Input

	ToWindowsConfigurationOutput() WindowsConfigurationOutput
	ToWindowsConfigurationOutputWithContext(context.Context) WindowsConfigurationOutput
}

type WindowsConfigurationArgs struct {
	EnableAutomaticUpdates pulumi.BoolPtrInput `pulumi:"enableAutomaticUpdates"`
}

func (WindowsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsConfiguration)(nil)).Elem()
}

func (i WindowsConfigurationArgs) ToWindowsConfigurationOutput() WindowsConfigurationOutput {
	return i.ToWindowsConfigurationOutputWithContext(context.Background())
}

func (i WindowsConfigurationArgs) ToWindowsConfigurationOutputWithContext(ctx context.Context) WindowsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsConfigurationOutput)
}

func (i WindowsConfigurationArgs) ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput {
	return i.ToWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (i WindowsConfigurationArgs) ToWindowsConfigurationPtrOutputWithContext(ctx context.Context) WindowsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsConfigurationOutput).ToWindowsConfigurationPtrOutputWithContext(ctx)
}









type WindowsConfigurationPtrInput interface {
	pulumi.Input

	ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput
	ToWindowsConfigurationPtrOutputWithContext(context.Context) WindowsConfigurationPtrOutput
}

type windowsConfigurationPtrType WindowsConfigurationArgs

func WindowsConfigurationPtr(v *WindowsConfigurationArgs) WindowsConfigurationPtrInput {
	return (*windowsConfigurationPtrType)(v)
}

func (*windowsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsConfiguration)(nil)).Elem()
}

func (i *windowsConfigurationPtrType) ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput {
	return i.ToWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (i *windowsConfigurationPtrType) ToWindowsConfigurationPtrOutputWithContext(ctx context.Context) WindowsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsConfigurationPtrOutput)
}

type WindowsConfigurationOutput struct{ *pulumi.OutputState }

func (WindowsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsConfiguration)(nil)).Elem()
}

func (o WindowsConfigurationOutput) ToWindowsConfigurationOutput() WindowsConfigurationOutput {
	return o
}

func (o WindowsConfigurationOutput) ToWindowsConfigurationOutputWithContext(ctx context.Context) WindowsConfigurationOutput {
	return o
}

func (o WindowsConfigurationOutput) ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput {
	return o.ToWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (o WindowsConfigurationOutput) ToWindowsConfigurationPtrOutputWithContext(ctx context.Context) WindowsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WindowsConfiguration) *WindowsConfiguration {
		return &v
	}).(WindowsConfigurationPtrOutput)
}

func (o WindowsConfigurationOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WindowsConfiguration) *bool { return v.EnableAutomaticUpdates }).(pulumi.BoolPtrOutput)
}

type WindowsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (WindowsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsConfiguration)(nil)).Elem()
}

func (o WindowsConfigurationPtrOutput) ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput {
	return o
}

func (o WindowsConfigurationPtrOutput) ToWindowsConfigurationPtrOutputWithContext(ctx context.Context) WindowsConfigurationPtrOutput {
	return o
}

func (o WindowsConfigurationPtrOutput) Elem() WindowsConfigurationOutput {
	return o.ApplyT(func(v *WindowsConfiguration) WindowsConfiguration {
		if v != nil {
			return *v
		}
		var ret WindowsConfiguration
		return ret
	}).(WindowsConfigurationOutput)
}

func (o WindowsConfigurationPtrOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WindowsConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAutomaticUpdates
	}).(pulumi.BoolPtrOutput)
}

type WindowsConfigurationResponse struct {
	EnableAutomaticUpdates *bool `pulumi:"enableAutomaticUpdates"`
}

type WindowsConfigurationResponseOutput struct{ *pulumi.OutputState }

func (WindowsConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsConfigurationResponse)(nil)).Elem()
}

func (o WindowsConfigurationResponseOutput) ToWindowsConfigurationResponseOutput() WindowsConfigurationResponseOutput {
	return o
}

func (o WindowsConfigurationResponseOutput) ToWindowsConfigurationResponseOutputWithContext(ctx context.Context) WindowsConfigurationResponseOutput {
	return o
}

func (o WindowsConfigurationResponseOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WindowsConfigurationResponse) *bool { return v.EnableAutomaticUpdates }).(pulumi.BoolPtrOutput)
}

type WindowsConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (WindowsConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsConfigurationResponse)(nil)).Elem()
}

func (o WindowsConfigurationResponsePtrOutput) ToWindowsConfigurationResponsePtrOutput() WindowsConfigurationResponsePtrOutput {
	return o
}

func (o WindowsConfigurationResponsePtrOutput) ToWindowsConfigurationResponsePtrOutputWithContext(ctx context.Context) WindowsConfigurationResponsePtrOutput {
	return o
}

func (o WindowsConfigurationResponsePtrOutput) Elem() WindowsConfigurationResponseOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) WindowsConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret WindowsConfigurationResponse
		return ret
	}).(WindowsConfigurationResponseOutput)
}

func (o WindowsConfigurationResponsePtrOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAutomaticUpdates
	}).(pulumi.BoolPtrOutput)
}

type WindowsUserConfiguration struct {
	LoginMode *LoginMode `pulumi:"loginMode"`
}





type WindowsUserConfigurationInput interface {
	pulumi.Input

	ToWindowsUserConfigurationOutput() WindowsUserConfigurationOutput
	ToWindowsUserConfigurationOutputWithContext(context.Context) WindowsUserConfigurationOutput
}

type WindowsUserConfigurationArgs struct {
	LoginMode LoginModePtrInput `pulumi:"loginMode"`
}

func (WindowsUserConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsUserConfiguration)(nil)).Elem()
}

func (i WindowsUserConfigurationArgs) ToWindowsUserConfigurationOutput() WindowsUserConfigurationOutput {
	return i.ToWindowsUserConfigurationOutputWithContext(context.Background())
}

func (i WindowsUserConfigurationArgs) ToWindowsUserConfigurationOutputWithContext(ctx context.Context) WindowsUserConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsUserConfigurationOutput)
}

func (i WindowsUserConfigurationArgs) ToWindowsUserConfigurationPtrOutput() WindowsUserConfigurationPtrOutput {
	return i.ToWindowsUserConfigurationPtrOutputWithContext(context.Background())
}

func (i WindowsUserConfigurationArgs) ToWindowsUserConfigurationPtrOutputWithContext(ctx context.Context) WindowsUserConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsUserConfigurationOutput).ToWindowsUserConfigurationPtrOutputWithContext(ctx)
}









type WindowsUserConfigurationPtrInput interface {
	pulumi.Input

	ToWindowsUserConfigurationPtrOutput() WindowsUserConfigurationPtrOutput
	ToWindowsUserConfigurationPtrOutputWithContext(context.Context) WindowsUserConfigurationPtrOutput
}

type windowsUserConfigurationPtrType WindowsUserConfigurationArgs

func WindowsUserConfigurationPtr(v *WindowsUserConfigurationArgs) WindowsUserConfigurationPtrInput {
	return (*windowsUserConfigurationPtrType)(v)
}

func (*windowsUserConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsUserConfiguration)(nil)).Elem()
}

func (i *windowsUserConfigurationPtrType) ToWindowsUserConfigurationPtrOutput() WindowsUserConfigurationPtrOutput {
	return i.ToWindowsUserConfigurationPtrOutputWithContext(context.Background())
}

func (i *windowsUserConfigurationPtrType) ToWindowsUserConfigurationPtrOutputWithContext(ctx context.Context) WindowsUserConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsUserConfigurationPtrOutput)
}

type WindowsUserConfigurationOutput struct{ *pulumi.OutputState }

func (WindowsUserConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsUserConfiguration)(nil)).Elem()
}

func (o WindowsUserConfigurationOutput) ToWindowsUserConfigurationOutput() WindowsUserConfigurationOutput {
	return o
}

func (o WindowsUserConfigurationOutput) ToWindowsUserConfigurationOutputWithContext(ctx context.Context) WindowsUserConfigurationOutput {
	return o
}

func (o WindowsUserConfigurationOutput) ToWindowsUserConfigurationPtrOutput() WindowsUserConfigurationPtrOutput {
	return o.ToWindowsUserConfigurationPtrOutputWithContext(context.Background())
}

func (o WindowsUserConfigurationOutput) ToWindowsUserConfigurationPtrOutputWithContext(ctx context.Context) WindowsUserConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WindowsUserConfiguration) *WindowsUserConfiguration {
		return &v
	}).(WindowsUserConfigurationPtrOutput)
}

func (o WindowsUserConfigurationOutput) LoginMode() LoginModePtrOutput {
	return o.ApplyT(func(v WindowsUserConfiguration) *LoginMode { return v.LoginMode }).(LoginModePtrOutput)
}

type WindowsUserConfigurationPtrOutput struct{ *pulumi.OutputState }

func (WindowsUserConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsUserConfiguration)(nil)).Elem()
}

func (o WindowsUserConfigurationPtrOutput) ToWindowsUserConfigurationPtrOutput() WindowsUserConfigurationPtrOutput {
	return o
}

func (o WindowsUserConfigurationPtrOutput) ToWindowsUserConfigurationPtrOutputWithContext(ctx context.Context) WindowsUserConfigurationPtrOutput {
	return o
}

func (o WindowsUserConfigurationPtrOutput) Elem() WindowsUserConfigurationOutput {
	return o.ApplyT(func(v *WindowsUserConfiguration) WindowsUserConfiguration {
		if v != nil {
			return *v
		}
		var ret WindowsUserConfiguration
		return ret
	}).(WindowsUserConfigurationOutput)
}

func (o WindowsUserConfigurationPtrOutput) LoginMode() LoginModePtrOutput {
	return o.ApplyT(func(v *WindowsUserConfiguration) *LoginMode {
		if v == nil {
			return nil
		}
		return v.LoginMode
	}).(LoginModePtrOutput)
}

type WindowsUserConfigurationResponse struct {
	LoginMode *string `pulumi:"loginMode"`
}

type WindowsUserConfigurationResponseOutput struct{ *pulumi.OutputState }

func (WindowsUserConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsUserConfigurationResponse)(nil)).Elem()
}

func (o WindowsUserConfigurationResponseOutput) ToWindowsUserConfigurationResponseOutput() WindowsUserConfigurationResponseOutput {
	return o
}

func (o WindowsUserConfigurationResponseOutput) ToWindowsUserConfigurationResponseOutputWithContext(ctx context.Context) WindowsUserConfigurationResponseOutput {
	return o
}

func (o WindowsUserConfigurationResponseOutput) LoginMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsUserConfigurationResponse) *string { return v.LoginMode }).(pulumi.StringPtrOutput)
}

type WindowsUserConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (WindowsUserConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsUserConfigurationResponse)(nil)).Elem()
}

func (o WindowsUserConfigurationResponsePtrOutput) ToWindowsUserConfigurationResponsePtrOutput() WindowsUserConfigurationResponsePtrOutput {
	return o
}

func (o WindowsUserConfigurationResponsePtrOutput) ToWindowsUserConfigurationResponsePtrOutputWithContext(ctx context.Context) WindowsUserConfigurationResponsePtrOutput {
	return o
}

func (o WindowsUserConfigurationResponsePtrOutput) Elem() WindowsUserConfigurationResponseOutput {
	return o.ApplyT(func(v *WindowsUserConfigurationResponse) WindowsUserConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret WindowsUserConfigurationResponse
		return ret
	}).(WindowsUserConfigurationResponseOutput)
}

func (o WindowsUserConfigurationResponsePtrOutput) LoginMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsUserConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.LoginMode
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationPackageReferenceOutput{})
	pulumi.RegisterOutputType(ApplicationPackageReferenceArrayOutput{})
	pulumi.RegisterOutputType(ApplicationPackageReferenceResponseOutput{})
	pulumi.RegisterOutputType(ApplicationPackageReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(AutoScaleRunErrorResponseOutput{})
	pulumi.RegisterOutputType(AutoScaleRunErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoScaleRunErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(AutoScaleRunResponseOutput{})
	pulumi.RegisterOutputType(AutoScaleSettingsOutput{})
	pulumi.RegisterOutputType(AutoScaleSettingsPtrOutput{})
	pulumi.RegisterOutputType(AutoScaleSettingsResponseOutput{})
	pulumi.RegisterOutputType(AutoScaleSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoStorageBasePropertiesOutput{})
	pulumi.RegisterOutputType(AutoStorageBasePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AutoStoragePropertiesResponseOutput{})
	pulumi.RegisterOutputType(AutoUserSpecificationOutput{})
	pulumi.RegisterOutputType(AutoUserSpecificationPtrOutput{})
	pulumi.RegisterOutputType(AutoUserSpecificationResponseOutput{})
	pulumi.RegisterOutputType(AutoUserSpecificationResponsePtrOutput{})
	pulumi.RegisterOutputType(CertificateReferenceOutput{})
	pulumi.RegisterOutputType(CertificateReferenceArrayOutput{})
	pulumi.RegisterOutputType(CertificateReferenceResponseOutput{})
	pulumi.RegisterOutputType(CertificateReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(CloudServiceConfigurationOutput{})
	pulumi.RegisterOutputType(CloudServiceConfigurationPtrOutput{})
	pulumi.RegisterOutputType(CloudServiceConfigurationResponseOutput{})
	pulumi.RegisterOutputType(CloudServiceConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerConfigurationOutput{})
	pulumi.RegisterOutputType(ContainerConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ContainerConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ContainerConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerRegistryOutput{})
	pulumi.RegisterOutputType(ContainerRegistryPtrOutput{})
	pulumi.RegisterOutputType(ContainerRegistryArrayOutput{})
	pulumi.RegisterOutputType(ContainerRegistryResponseOutput{})
	pulumi.RegisterOutputType(ContainerRegistryResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerRegistryResponseArrayOutput{})
	pulumi.RegisterOutputType(DataDiskOutput{})
	pulumi.RegisterOutputType(DataDiskArrayOutput{})
	pulumi.RegisterOutputType(DataDiskResponseOutput{})
	pulumi.RegisterOutputType(DataDiskResponseArrayOutput{})
	pulumi.RegisterOutputType(DeleteCertificateErrorResponseOutput{})
	pulumi.RegisterOutputType(DeleteCertificateErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(DeploymentConfigurationOutput{})
	pulumi.RegisterOutputType(DeploymentConfigurationPtrOutput{})
	pulumi.RegisterOutputType(DeploymentConfigurationResponseOutput{})
	pulumi.RegisterOutputType(DeploymentConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(EnvironmentSettingOutput{})
	pulumi.RegisterOutputType(EnvironmentSettingArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentSettingResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentSettingResponseArrayOutput{})
	pulumi.RegisterOutputType(FixedScaleSettingsOutput{})
	pulumi.RegisterOutputType(FixedScaleSettingsPtrOutput{})
	pulumi.RegisterOutputType(FixedScaleSettingsResponseOutput{})
	pulumi.RegisterOutputType(FixedScaleSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageReferenceOutput{})
	pulumi.RegisterOutputType(ImageReferencePtrOutput{})
	pulumi.RegisterOutputType(ImageReferenceResponseOutput{})
	pulumi.RegisterOutputType(ImageReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(InboundNatPoolOutput{})
	pulumi.RegisterOutputType(InboundNatPoolArrayOutput{})
	pulumi.RegisterOutputType(InboundNatPoolResponseOutput{})
	pulumi.RegisterOutputType(InboundNatPoolResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultReferenceOutput{})
	pulumi.RegisterOutputType(KeyVaultReferencePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultReferenceResponseOutput{})
	pulumi.RegisterOutputType(LinuxUserConfigurationOutput{})
	pulumi.RegisterOutputType(LinuxUserConfigurationPtrOutput{})
	pulumi.RegisterOutputType(LinuxUserConfigurationResponseOutput{})
	pulumi.RegisterOutputType(LinuxUserConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(MetadataItemOutput{})
	pulumi.RegisterOutputType(MetadataItemArrayOutput{})
	pulumi.RegisterOutputType(MetadataItemResponseOutput{})
	pulumi.RegisterOutputType(MetadataItemResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkConfigurationOutput{})
	pulumi.RegisterOutputType(NetworkConfigurationPtrOutput{})
	pulumi.RegisterOutputType(NetworkConfigurationResponseOutput{})
	pulumi.RegisterOutputType(NetworkConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkSecurityGroupRuleOutput{})
	pulumi.RegisterOutputType(NetworkSecurityGroupRuleArrayOutput{})
	pulumi.RegisterOutputType(NetworkSecurityGroupRuleResponseOutput{})
	pulumi.RegisterOutputType(NetworkSecurityGroupRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(PoolEndpointConfigurationOutput{})
	pulumi.RegisterOutputType(PoolEndpointConfigurationPtrOutput{})
	pulumi.RegisterOutputType(PoolEndpointConfigurationResponseOutput{})
	pulumi.RegisterOutputType(PoolEndpointConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ResizeErrorResponseOutput{})
	pulumi.RegisterOutputType(ResizeErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(ResizeOperationStatusResponseOutput{})
	pulumi.RegisterOutputType(ResourceFileOutput{})
	pulumi.RegisterOutputType(ResourceFileArrayOutput{})
	pulumi.RegisterOutputType(ResourceFileResponseOutput{})
	pulumi.RegisterOutputType(ResourceFileResponseArrayOutput{})
	pulumi.RegisterOutputType(ScaleSettingsOutput{})
	pulumi.RegisterOutputType(ScaleSettingsPtrOutput{})
	pulumi.RegisterOutputType(ScaleSettingsResponseOutput{})
	pulumi.RegisterOutputType(ScaleSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(StartTaskOutput{})
	pulumi.RegisterOutputType(StartTaskPtrOutput{})
	pulumi.RegisterOutputType(StartTaskResponseOutput{})
	pulumi.RegisterOutputType(StartTaskResponsePtrOutput{})
	pulumi.RegisterOutputType(TaskContainerSettingsOutput{})
	pulumi.RegisterOutputType(TaskContainerSettingsPtrOutput{})
	pulumi.RegisterOutputType(TaskContainerSettingsResponseOutput{})
	pulumi.RegisterOutputType(TaskContainerSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(TaskSchedulingPolicyOutput{})
	pulumi.RegisterOutputType(TaskSchedulingPolicyPtrOutput{})
	pulumi.RegisterOutputType(TaskSchedulingPolicyResponseOutput{})
	pulumi.RegisterOutputType(TaskSchedulingPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(UserAccountOutput{})
	pulumi.RegisterOutputType(UserAccountArrayOutput{})
	pulumi.RegisterOutputType(UserAccountResponseOutput{})
	pulumi.RegisterOutputType(UserAccountResponseArrayOutput{})
	pulumi.RegisterOutputType(UserIdentityOutput{})
	pulumi.RegisterOutputType(UserIdentityPtrOutput{})
	pulumi.RegisterOutputType(UserIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualMachineConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(WindowsConfigurationOutput{})
	pulumi.RegisterOutputType(WindowsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(WindowsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(WindowsConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(WindowsUserConfigurationOutput{})
	pulumi.RegisterOutputType(WindowsUserConfigurationPtrOutput{})
	pulumi.RegisterOutputType(WindowsUserConfigurationResponseOutput{})
	pulumi.RegisterOutputType(WindowsUserConfigurationResponsePtrOutput{})
}
