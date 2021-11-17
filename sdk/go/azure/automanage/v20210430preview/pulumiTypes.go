


package v20210430preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigurationProfileAssignmentProperties struct {
	ConfigurationProfile *string                `pulumi:"configurationProfile"`
	ProfileOverrides     map[string]interface{} `pulumi:"profileOverrides"`
	TargetId             *string                `pulumi:"targetId"`
}





type ConfigurationProfileAssignmentPropertiesInput interface {
	pulumi.Input

	ToConfigurationProfileAssignmentPropertiesOutput() ConfigurationProfileAssignmentPropertiesOutput
	ToConfigurationProfileAssignmentPropertiesOutputWithContext(context.Context) ConfigurationProfileAssignmentPropertiesOutput
}

type ConfigurationProfileAssignmentPropertiesArgs struct {
	ConfigurationProfile pulumi.StringPtrInput `pulumi:"configurationProfile"`
	ProfileOverrides     pulumi.MapInput       `pulumi:"profileOverrides"`
	TargetId             pulumi.StringPtrInput `pulumi:"targetId"`
}

func (ConfigurationProfileAssignmentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileAssignmentProperties)(nil)).Elem()
}

func (i ConfigurationProfileAssignmentPropertiesArgs) ToConfigurationProfileAssignmentPropertiesOutput() ConfigurationProfileAssignmentPropertiesOutput {
	return i.ToConfigurationProfileAssignmentPropertiesOutputWithContext(context.Background())
}

func (i ConfigurationProfileAssignmentPropertiesArgs) ToConfigurationProfileAssignmentPropertiesOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentPropertiesOutput)
}

func (i ConfigurationProfileAssignmentPropertiesArgs) ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput {
	return i.ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (i ConfigurationProfileAssignmentPropertiesArgs) ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentPropertiesOutput).ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx)
}









type ConfigurationProfileAssignmentPropertiesPtrInput interface {
	pulumi.Input

	ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput
	ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput
}

type configurationProfileAssignmentPropertiesPtrType ConfigurationProfileAssignmentPropertiesArgs

func ConfigurationProfileAssignmentPropertiesPtr(v *ConfigurationProfileAssignmentPropertiesArgs) ConfigurationProfileAssignmentPropertiesPtrInput {
	return (*configurationProfileAssignmentPropertiesPtrType)(v)
}

func (*configurationProfileAssignmentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileAssignmentProperties)(nil)).Elem()
}

func (i *configurationProfileAssignmentPropertiesPtrType) ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput {
	return i.ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (i *configurationProfileAssignmentPropertiesPtrType) ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentPropertiesPtrOutput)
}

type ConfigurationProfileAssignmentPropertiesOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileAssignmentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileAssignmentProperties)(nil)).Elem()
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ToConfigurationProfileAssignmentPropertiesOutput() ConfigurationProfileAssignmentPropertiesOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ToConfigurationProfileAssignmentPropertiesOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput {
	return o.ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationProfileAssignmentProperties) *ConfigurationProfileAssignmentProperties {
		return &v
	}).(ConfigurationProfileAssignmentPropertiesPtrOutput)
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ConfigurationProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentProperties) *string { return v.ConfigurationProfile }).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ProfileOverrides() pulumi.MapOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentProperties) map[string]interface{} { return v.ProfileOverrides }).(pulumi.MapOutput)
}

func (o ConfigurationProfileAssignmentPropertiesOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentProperties) *string { return v.TargetId }).(pulumi.StringPtrOutput)
}

type ConfigurationProfileAssignmentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileAssignmentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileAssignmentProperties)(nil)).Elem()
}

func (o ConfigurationProfileAssignmentPropertiesPtrOutput) ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesPtrOutput) ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesPtrOutput) Elem() ConfigurationProfileAssignmentPropertiesOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentProperties) ConfigurationProfileAssignmentProperties {
		if v != nil {
			return *v
		}
		var ret ConfigurationProfileAssignmentProperties
		return ret
	}).(ConfigurationProfileAssignmentPropertiesOutput)
}

func (o ConfigurationProfileAssignmentPropertiesPtrOutput) ConfigurationProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.ConfigurationProfile
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfileAssignmentPropertiesPtrOutput) ProfileOverrides() pulumi.MapOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentProperties) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.ProfileOverrides
	}).(pulumi.MapOutput)
}

func (o ConfigurationProfileAssignmentPropertiesPtrOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.TargetId
	}).(pulumi.StringPtrOutput)
}

type ConfigurationProfileAssignmentPropertiesResponse struct {
	ConfigurationProfile *string                `pulumi:"configurationProfile"`
	ProfileOverrides     map[string]interface{} `pulumi:"profileOverrides"`
	Status               string                 `pulumi:"status"`
	TargetId             *string                `pulumi:"targetId"`
}





type ConfigurationProfileAssignmentPropertiesResponseInput interface {
	pulumi.Input

	ToConfigurationProfileAssignmentPropertiesResponseOutput() ConfigurationProfileAssignmentPropertiesResponseOutput
	ToConfigurationProfileAssignmentPropertiesResponseOutputWithContext(context.Context) ConfigurationProfileAssignmentPropertiesResponseOutput
}

type ConfigurationProfileAssignmentPropertiesResponseArgs struct {
	ConfigurationProfile pulumi.StringPtrInput `pulumi:"configurationProfile"`
	ProfileOverrides     pulumi.MapInput       `pulumi:"profileOverrides"`
	Status               pulumi.StringInput    `pulumi:"status"`
	TargetId             pulumi.StringPtrInput `pulumi:"targetId"`
}

func (ConfigurationProfileAssignmentPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileAssignmentPropertiesResponse)(nil)).Elem()
}

func (i ConfigurationProfileAssignmentPropertiesResponseArgs) ToConfigurationProfileAssignmentPropertiesResponseOutput() ConfigurationProfileAssignmentPropertiesResponseOutput {
	return i.ToConfigurationProfileAssignmentPropertiesResponseOutputWithContext(context.Background())
}

func (i ConfigurationProfileAssignmentPropertiesResponseArgs) ToConfigurationProfileAssignmentPropertiesResponseOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentPropertiesResponseOutput)
}

func (i ConfigurationProfileAssignmentPropertiesResponseArgs) ToConfigurationProfileAssignmentPropertiesResponsePtrOutput() ConfigurationProfileAssignmentPropertiesResponsePtrOutput {
	return i.ToConfigurationProfileAssignmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ConfigurationProfileAssignmentPropertiesResponseArgs) ToConfigurationProfileAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentPropertiesResponseOutput).ToConfigurationProfileAssignmentPropertiesResponsePtrOutputWithContext(ctx)
}









type ConfigurationProfileAssignmentPropertiesResponsePtrInput interface {
	pulumi.Input

	ToConfigurationProfileAssignmentPropertiesResponsePtrOutput() ConfigurationProfileAssignmentPropertiesResponsePtrOutput
	ToConfigurationProfileAssignmentPropertiesResponsePtrOutputWithContext(context.Context) ConfigurationProfileAssignmentPropertiesResponsePtrOutput
}

type configurationProfileAssignmentPropertiesResponsePtrType ConfigurationProfileAssignmentPropertiesResponseArgs

func ConfigurationProfileAssignmentPropertiesResponsePtr(v *ConfigurationProfileAssignmentPropertiesResponseArgs) ConfigurationProfileAssignmentPropertiesResponsePtrInput {
	return (*configurationProfileAssignmentPropertiesResponsePtrType)(v)
}

func (*configurationProfileAssignmentPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileAssignmentPropertiesResponse)(nil)).Elem()
}

func (i *configurationProfileAssignmentPropertiesResponsePtrType) ToConfigurationProfileAssignmentPropertiesResponsePtrOutput() ConfigurationProfileAssignmentPropertiesResponsePtrOutput {
	return i.ToConfigurationProfileAssignmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *configurationProfileAssignmentPropertiesResponsePtrType) ToConfigurationProfileAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentPropertiesResponsePtrOutput)
}

type ConfigurationProfileAssignmentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileAssignmentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileAssignmentPropertiesResponse)(nil)).Elem()
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ToConfigurationProfileAssignmentPropertiesResponseOutput() ConfigurationProfileAssignmentPropertiesResponseOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ToConfigurationProfileAssignmentPropertiesResponseOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesResponseOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ToConfigurationProfileAssignmentPropertiesResponsePtrOutput() ConfigurationProfileAssignmentPropertiesResponsePtrOutput {
	return o.ToConfigurationProfileAssignmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ToConfigurationProfileAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationProfileAssignmentPropertiesResponse) *ConfigurationProfileAssignmentPropertiesResponse {
		return &v
	}).(ConfigurationProfileAssignmentPropertiesResponsePtrOutput)
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ConfigurationProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) *string { return v.ConfigurationProfile }).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ProfileOverrides() pulumi.MapOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) map[string]interface{} {
		return v.ProfileOverrides
	}).(pulumi.MapOutput)
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) *string { return v.TargetId }).(pulumi.StringPtrOutput)
}

type ConfigurationProfileAssignmentPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileAssignmentPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileAssignmentPropertiesResponse)(nil)).Elem()
}

func (o ConfigurationProfileAssignmentPropertiesResponsePtrOutput) ToConfigurationProfileAssignmentPropertiesResponsePtrOutput() ConfigurationProfileAssignmentPropertiesResponsePtrOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesResponsePtrOutput) ToConfigurationProfileAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesResponsePtrOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesResponsePtrOutput) Elem() ConfigurationProfileAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentPropertiesResponse) ConfigurationProfileAssignmentPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ConfigurationProfileAssignmentPropertiesResponse
		return ret
	}).(ConfigurationProfileAssignmentPropertiesResponseOutput)
}

func (o ConfigurationProfileAssignmentPropertiesResponsePtrOutput) ConfigurationProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConfigurationProfile
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfileAssignmentPropertiesResponsePtrOutput) ProfileOverrides() pulumi.MapOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentPropertiesResponse) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.ProfileOverrides
	}).(pulumi.MapOutput)
}

func (o ConfigurationProfileAssignmentPropertiesResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfileAssignmentPropertiesResponsePtrOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetId
	}).(pulumi.StringPtrOutput)
}

type ConfigurationProfileProperties struct {
	Configuration interface{}   `pulumi:"configuration"`
	Overrides     []interface{} `pulumi:"overrides"`
}





type ConfigurationProfilePropertiesInput interface {
	pulumi.Input

	ToConfigurationProfilePropertiesOutput() ConfigurationProfilePropertiesOutput
	ToConfigurationProfilePropertiesOutputWithContext(context.Context) ConfigurationProfilePropertiesOutput
}

type ConfigurationProfilePropertiesArgs struct {
	Configuration pulumi.Input      `pulumi:"configuration"`
	Overrides     pulumi.ArrayInput `pulumi:"overrides"`
}

func (ConfigurationProfilePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileProperties)(nil)).Elem()
}

func (i ConfigurationProfilePropertiesArgs) ToConfigurationProfilePropertiesOutput() ConfigurationProfilePropertiesOutput {
	return i.ToConfigurationProfilePropertiesOutputWithContext(context.Background())
}

func (i ConfigurationProfilePropertiesArgs) ToConfigurationProfilePropertiesOutputWithContext(ctx context.Context) ConfigurationProfilePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePropertiesOutput)
}

func (i ConfigurationProfilePropertiesArgs) ToConfigurationProfilePropertiesPtrOutput() ConfigurationProfilePropertiesPtrOutput {
	return i.ToConfigurationProfilePropertiesPtrOutputWithContext(context.Background())
}

func (i ConfigurationProfilePropertiesArgs) ToConfigurationProfilePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfilePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePropertiesOutput).ToConfigurationProfilePropertiesPtrOutputWithContext(ctx)
}









type ConfigurationProfilePropertiesPtrInput interface {
	pulumi.Input

	ToConfigurationProfilePropertiesPtrOutput() ConfigurationProfilePropertiesPtrOutput
	ToConfigurationProfilePropertiesPtrOutputWithContext(context.Context) ConfigurationProfilePropertiesPtrOutput
}

type configurationProfilePropertiesPtrType ConfigurationProfilePropertiesArgs

func ConfigurationProfilePropertiesPtr(v *ConfigurationProfilePropertiesArgs) ConfigurationProfilePropertiesPtrInput {
	return (*configurationProfilePropertiesPtrType)(v)
}

func (*configurationProfilePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileProperties)(nil)).Elem()
}

func (i *configurationProfilePropertiesPtrType) ToConfigurationProfilePropertiesPtrOutput() ConfigurationProfilePropertiesPtrOutput {
	return i.ToConfigurationProfilePropertiesPtrOutputWithContext(context.Background())
}

func (i *configurationProfilePropertiesPtrType) ToConfigurationProfilePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfilePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePropertiesPtrOutput)
}

type ConfigurationProfilePropertiesOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileProperties)(nil)).Elem()
}

func (o ConfigurationProfilePropertiesOutput) ToConfigurationProfilePropertiesOutput() ConfigurationProfilePropertiesOutput {
	return o
}

func (o ConfigurationProfilePropertiesOutput) ToConfigurationProfilePropertiesOutputWithContext(ctx context.Context) ConfigurationProfilePropertiesOutput {
	return o
}

func (o ConfigurationProfilePropertiesOutput) ToConfigurationProfilePropertiesPtrOutput() ConfigurationProfilePropertiesPtrOutput {
	return o.ToConfigurationProfilePropertiesPtrOutputWithContext(context.Background())
}

func (o ConfigurationProfilePropertiesOutput) ToConfigurationProfilePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfilePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationProfileProperties) *ConfigurationProfileProperties {
		return &v
	}).(ConfigurationProfilePropertiesPtrOutput)
}

func (o ConfigurationProfilePropertiesOutput) Configuration() pulumi.AnyOutput {
	return o.ApplyT(func(v ConfigurationProfileProperties) interface{} { return v.Configuration }).(pulumi.AnyOutput)
}

func (o ConfigurationProfilePropertiesOutput) Overrides() pulumi.ArrayOutput {
	return o.ApplyT(func(v ConfigurationProfileProperties) []interface{} { return v.Overrides }).(pulumi.ArrayOutput)
}

type ConfigurationProfilePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileProperties)(nil)).Elem()
}

func (o ConfigurationProfilePropertiesPtrOutput) ToConfigurationProfilePropertiesPtrOutput() ConfigurationProfilePropertiesPtrOutput {
	return o
}

func (o ConfigurationProfilePropertiesPtrOutput) ToConfigurationProfilePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfilePropertiesPtrOutput {
	return o
}

func (o ConfigurationProfilePropertiesPtrOutput) Elem() ConfigurationProfilePropertiesOutput {
	return o.ApplyT(func(v *ConfigurationProfileProperties) ConfigurationProfileProperties {
		if v != nil {
			return *v
		}
		var ret ConfigurationProfileProperties
		return ret
	}).(ConfigurationProfilePropertiesOutput)
}

func (o ConfigurationProfilePropertiesPtrOutput) Configuration() pulumi.AnyOutput {
	return o.ApplyT(func(v *ConfigurationProfileProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(pulumi.AnyOutput)
}

func (o ConfigurationProfilePropertiesPtrOutput) Overrides() pulumi.ArrayOutput {
	return o.ApplyT(func(v *ConfigurationProfileProperties) []interface{} {
		if v == nil {
			return nil
		}
		return v.Overrides
	}).(pulumi.ArrayOutput)
}

type ConfigurationProfilePropertiesResponse struct {
	Configuration interface{}   `pulumi:"configuration"`
	Overrides     []interface{} `pulumi:"overrides"`
}





type ConfigurationProfilePropertiesResponseInput interface {
	pulumi.Input

	ToConfigurationProfilePropertiesResponseOutput() ConfigurationProfilePropertiesResponseOutput
	ToConfigurationProfilePropertiesResponseOutputWithContext(context.Context) ConfigurationProfilePropertiesResponseOutput
}

type ConfigurationProfilePropertiesResponseArgs struct {
	Configuration pulumi.Input      `pulumi:"configuration"`
	Overrides     pulumi.ArrayInput `pulumi:"overrides"`
}

func (ConfigurationProfilePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfilePropertiesResponse)(nil)).Elem()
}

func (i ConfigurationProfilePropertiesResponseArgs) ToConfigurationProfilePropertiesResponseOutput() ConfigurationProfilePropertiesResponseOutput {
	return i.ToConfigurationProfilePropertiesResponseOutputWithContext(context.Background())
}

func (i ConfigurationProfilePropertiesResponseArgs) ToConfigurationProfilePropertiesResponseOutputWithContext(ctx context.Context) ConfigurationProfilePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePropertiesResponseOutput)
}

func (i ConfigurationProfilePropertiesResponseArgs) ToConfigurationProfilePropertiesResponsePtrOutput() ConfigurationProfilePropertiesResponsePtrOutput {
	return i.ToConfigurationProfilePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ConfigurationProfilePropertiesResponseArgs) ToConfigurationProfilePropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfilePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePropertiesResponseOutput).ToConfigurationProfilePropertiesResponsePtrOutputWithContext(ctx)
}









type ConfigurationProfilePropertiesResponsePtrInput interface {
	pulumi.Input

	ToConfigurationProfilePropertiesResponsePtrOutput() ConfigurationProfilePropertiesResponsePtrOutput
	ToConfigurationProfilePropertiesResponsePtrOutputWithContext(context.Context) ConfigurationProfilePropertiesResponsePtrOutput
}

type configurationProfilePropertiesResponsePtrType ConfigurationProfilePropertiesResponseArgs

func ConfigurationProfilePropertiesResponsePtr(v *ConfigurationProfilePropertiesResponseArgs) ConfigurationProfilePropertiesResponsePtrInput {
	return (*configurationProfilePropertiesResponsePtrType)(v)
}

func (*configurationProfilePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfilePropertiesResponse)(nil)).Elem()
}

func (i *configurationProfilePropertiesResponsePtrType) ToConfigurationProfilePropertiesResponsePtrOutput() ConfigurationProfilePropertiesResponsePtrOutput {
	return i.ToConfigurationProfilePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *configurationProfilePropertiesResponsePtrType) ToConfigurationProfilePropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfilePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePropertiesResponsePtrOutput)
}

type ConfigurationProfilePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfilePropertiesResponse)(nil)).Elem()
}

func (o ConfigurationProfilePropertiesResponseOutput) ToConfigurationProfilePropertiesResponseOutput() ConfigurationProfilePropertiesResponseOutput {
	return o
}

func (o ConfigurationProfilePropertiesResponseOutput) ToConfigurationProfilePropertiesResponseOutputWithContext(ctx context.Context) ConfigurationProfilePropertiesResponseOutput {
	return o
}

func (o ConfigurationProfilePropertiesResponseOutput) ToConfigurationProfilePropertiesResponsePtrOutput() ConfigurationProfilePropertiesResponsePtrOutput {
	return o.ToConfigurationProfilePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ConfigurationProfilePropertiesResponseOutput) ToConfigurationProfilePropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfilePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationProfilePropertiesResponse) *ConfigurationProfilePropertiesResponse {
		return &v
	}).(ConfigurationProfilePropertiesResponsePtrOutput)
}

func (o ConfigurationProfilePropertiesResponseOutput) Configuration() pulumi.AnyOutput {
	return o.ApplyT(func(v ConfigurationProfilePropertiesResponse) interface{} { return v.Configuration }).(pulumi.AnyOutput)
}

func (o ConfigurationProfilePropertiesResponseOutput) Overrides() pulumi.ArrayOutput {
	return o.ApplyT(func(v ConfigurationProfilePropertiesResponse) []interface{} { return v.Overrides }).(pulumi.ArrayOutput)
}

type ConfigurationProfilePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfilePropertiesResponse)(nil)).Elem()
}

func (o ConfigurationProfilePropertiesResponsePtrOutput) ToConfigurationProfilePropertiesResponsePtrOutput() ConfigurationProfilePropertiesResponsePtrOutput {
	return o
}

func (o ConfigurationProfilePropertiesResponsePtrOutput) ToConfigurationProfilePropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfilePropertiesResponsePtrOutput {
	return o
}

func (o ConfigurationProfilePropertiesResponsePtrOutput) Elem() ConfigurationProfilePropertiesResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfilePropertiesResponse) ConfigurationProfilePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ConfigurationProfilePropertiesResponse
		return ret
	}).(ConfigurationProfilePropertiesResponseOutput)
}

func (o ConfigurationProfilePropertiesResponsePtrOutput) Configuration() pulumi.AnyOutput {
	return o.ApplyT(func(v *ConfigurationProfilePropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(pulumi.AnyOutput)
}

func (o ConfigurationProfilePropertiesResponsePtrOutput) Overrides() pulumi.ArrayOutput {
	return o.ApplyT(func(v *ConfigurationProfilePropertiesResponse) []interface{} {
		if v == nil {
			return nil
		}
		return v.Overrides
	}).(pulumi.ArrayOutput)
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
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentPropertiesOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePropertiesOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
