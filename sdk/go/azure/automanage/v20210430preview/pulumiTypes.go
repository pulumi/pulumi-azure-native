


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

func (o ConfigurationProfilePropertiesResponseOutput) Configuration() pulumi.AnyOutput {
	return o.ApplyT(func(v ConfigurationProfilePropertiesResponse) interface{} { return v.Configuration }).(pulumi.AnyOutput)
}

func (o ConfigurationProfilePropertiesResponseOutput) Overrides() pulumi.ArrayOutput {
	return o.ApplyT(func(v ConfigurationProfilePropertiesResponse) []interface{} { return v.Overrides }).(pulumi.ArrayOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
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

func init() {
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentPropertiesOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePropertiesOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
