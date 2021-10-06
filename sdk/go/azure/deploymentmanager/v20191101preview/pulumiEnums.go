


package v20191101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DeploymentMode string

const (
	DeploymentModeIncremental = DeploymentMode("Incremental")
	DeploymentModeComplete    = DeploymentMode("Complete")
)

func (DeploymentMode) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentMode)(nil)).Elem()
}

func (e DeploymentMode) ToDeploymentModeOutput() DeploymentModeOutput {
	return pulumi.ToOutput(e).(DeploymentModeOutput)
}

func (e DeploymentMode) ToDeploymentModeOutputWithContext(ctx context.Context) DeploymentModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DeploymentModeOutput)
}

func (e DeploymentMode) ToDeploymentModePtrOutput() DeploymentModePtrOutput {
	return e.ToDeploymentModePtrOutputWithContext(context.Background())
}

func (e DeploymentMode) ToDeploymentModePtrOutputWithContext(ctx context.Context) DeploymentModePtrOutput {
	return DeploymentMode(e).ToDeploymentModeOutputWithContext(ctx).ToDeploymentModePtrOutputWithContext(ctx)
}

func (e DeploymentMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeploymentMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeploymentMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DeploymentMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DeploymentModeOutput struct{ *pulumi.OutputState }

func (DeploymentModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentMode)(nil)).Elem()
}

func (o DeploymentModeOutput) ToDeploymentModeOutput() DeploymentModeOutput {
	return o
}

func (o DeploymentModeOutput) ToDeploymentModeOutputWithContext(ctx context.Context) DeploymentModeOutput {
	return o
}

func (o DeploymentModeOutput) ToDeploymentModePtrOutput() DeploymentModePtrOutput {
	return o.ToDeploymentModePtrOutputWithContext(context.Background())
}

func (o DeploymentModeOutput) ToDeploymentModePtrOutputWithContext(ctx context.Context) DeploymentModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentMode) *DeploymentMode {
		return &v
	}).(DeploymentModePtrOutput)
}

func (o DeploymentModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DeploymentModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeploymentMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DeploymentModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeploymentModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeploymentMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DeploymentModePtrOutput struct{ *pulumi.OutputState }

func (DeploymentModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentMode)(nil)).Elem()
}

func (o DeploymentModePtrOutput) ToDeploymentModePtrOutput() DeploymentModePtrOutput {
	return o
}

func (o DeploymentModePtrOutput) ToDeploymentModePtrOutputWithContext(ctx context.Context) DeploymentModePtrOutput {
	return o
}

func (o DeploymentModePtrOutput) Elem() DeploymentModeOutput {
	return o.ApplyT(func(v *DeploymentMode) DeploymentMode {
		if v != nil {
			return *v
		}
		var ret DeploymentMode
		return ret
	}).(DeploymentModeOutput)
}

func (o DeploymentModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeploymentModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DeploymentMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DeploymentModeInput interface {
	pulumi.Input

	ToDeploymentModeOutput() DeploymentModeOutput
	ToDeploymentModeOutputWithContext(context.Context) DeploymentModeOutput
}

var deploymentModePtrType = reflect.TypeOf((**DeploymentMode)(nil)).Elem()

type DeploymentModePtrInput interface {
	pulumi.Input

	ToDeploymentModePtrOutput() DeploymentModePtrOutput
	ToDeploymentModePtrOutputWithContext(context.Context) DeploymentModePtrOutput
}

type deploymentModePtr string

func DeploymentModePtr(v string) DeploymentModePtrInput {
	return (*deploymentModePtr)(&v)
}

func (*deploymentModePtr) ElementType() reflect.Type {
	return deploymentModePtrType
}

func (in *deploymentModePtr) ToDeploymentModePtrOutput() DeploymentModePtrOutput {
	return pulumi.ToOutput(in).(DeploymentModePtrOutput)
}

func (in *deploymentModePtr) ToDeploymentModePtrOutputWithContext(ctx context.Context) DeploymentModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DeploymentModePtrOutput)
}

type RestAuthLocation string

const (
	RestAuthLocationQuery  = RestAuthLocation("Query")
	RestAuthLocationHeader = RestAuthLocation("Header")
)

func (RestAuthLocation) ElementType() reflect.Type {
	return reflect.TypeOf((*RestAuthLocation)(nil)).Elem()
}

func (e RestAuthLocation) ToRestAuthLocationOutput() RestAuthLocationOutput {
	return pulumi.ToOutput(e).(RestAuthLocationOutput)
}

func (e RestAuthLocation) ToRestAuthLocationOutputWithContext(ctx context.Context) RestAuthLocationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RestAuthLocationOutput)
}

func (e RestAuthLocation) ToRestAuthLocationPtrOutput() RestAuthLocationPtrOutput {
	return e.ToRestAuthLocationPtrOutputWithContext(context.Background())
}

func (e RestAuthLocation) ToRestAuthLocationPtrOutputWithContext(ctx context.Context) RestAuthLocationPtrOutput {
	return RestAuthLocation(e).ToRestAuthLocationOutputWithContext(ctx).ToRestAuthLocationPtrOutputWithContext(ctx)
}

func (e RestAuthLocation) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RestAuthLocation) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RestAuthLocation) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RestAuthLocation) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RestAuthLocationOutput struct{ *pulumi.OutputState }

func (RestAuthLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestAuthLocation)(nil)).Elem()
}

func (o RestAuthLocationOutput) ToRestAuthLocationOutput() RestAuthLocationOutput {
	return o
}

func (o RestAuthLocationOutput) ToRestAuthLocationOutputWithContext(ctx context.Context) RestAuthLocationOutput {
	return o
}

func (o RestAuthLocationOutput) ToRestAuthLocationPtrOutput() RestAuthLocationPtrOutput {
	return o.ToRestAuthLocationPtrOutputWithContext(context.Background())
}

func (o RestAuthLocationOutput) ToRestAuthLocationPtrOutputWithContext(ctx context.Context) RestAuthLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RestAuthLocation) *RestAuthLocation {
		return &v
	}).(RestAuthLocationPtrOutput)
}

func (o RestAuthLocationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RestAuthLocationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RestAuthLocation) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RestAuthLocationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RestAuthLocationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RestAuthLocation) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RestAuthLocationPtrOutput struct{ *pulumi.OutputState }

func (RestAuthLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestAuthLocation)(nil)).Elem()
}

func (o RestAuthLocationPtrOutput) ToRestAuthLocationPtrOutput() RestAuthLocationPtrOutput {
	return o
}

func (o RestAuthLocationPtrOutput) ToRestAuthLocationPtrOutputWithContext(ctx context.Context) RestAuthLocationPtrOutput {
	return o
}

func (o RestAuthLocationPtrOutput) Elem() RestAuthLocationOutput {
	return o.ApplyT(func(v *RestAuthLocation) RestAuthLocation {
		if v != nil {
			return *v
		}
		var ret RestAuthLocation
		return ret
	}).(RestAuthLocationOutput)
}

func (o RestAuthLocationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RestAuthLocationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RestAuthLocation) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RestAuthLocationInput interface {
	pulumi.Input

	ToRestAuthLocationOutput() RestAuthLocationOutput
	ToRestAuthLocationOutputWithContext(context.Context) RestAuthLocationOutput
}

var restAuthLocationPtrType = reflect.TypeOf((**RestAuthLocation)(nil)).Elem()

type RestAuthLocationPtrInput interface {
	pulumi.Input

	ToRestAuthLocationPtrOutput() RestAuthLocationPtrOutput
	ToRestAuthLocationPtrOutputWithContext(context.Context) RestAuthLocationPtrOutput
}

type restAuthLocationPtr string

func RestAuthLocationPtr(v string) RestAuthLocationPtrInput {
	return (*restAuthLocationPtr)(&v)
}

func (*restAuthLocationPtr) ElementType() reflect.Type {
	return restAuthLocationPtrType
}

func (in *restAuthLocationPtr) ToRestAuthLocationPtrOutput() RestAuthLocationPtrOutput {
	return pulumi.ToOutput(in).(RestAuthLocationPtrOutput)
}

func (in *restAuthLocationPtr) ToRestAuthLocationPtrOutputWithContext(ctx context.Context) RestAuthLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RestAuthLocationPtrOutput)
}

type RestAuthType string

const (
	RestAuthTypeApiKey          = RestAuthType("ApiKey")
	RestAuthTypeRolloutIdentity = RestAuthType("RolloutIdentity")
)

func (RestAuthType) ElementType() reflect.Type {
	return reflect.TypeOf((*RestAuthType)(nil)).Elem()
}

func (e RestAuthType) ToRestAuthTypeOutput() RestAuthTypeOutput {
	return pulumi.ToOutput(e).(RestAuthTypeOutput)
}

func (e RestAuthType) ToRestAuthTypeOutputWithContext(ctx context.Context) RestAuthTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RestAuthTypeOutput)
}

func (e RestAuthType) ToRestAuthTypePtrOutput() RestAuthTypePtrOutput {
	return e.ToRestAuthTypePtrOutputWithContext(context.Background())
}

func (e RestAuthType) ToRestAuthTypePtrOutputWithContext(ctx context.Context) RestAuthTypePtrOutput {
	return RestAuthType(e).ToRestAuthTypeOutputWithContext(ctx).ToRestAuthTypePtrOutputWithContext(ctx)
}

func (e RestAuthType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RestAuthType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RestAuthType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RestAuthType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RestAuthTypeOutput struct{ *pulumi.OutputState }

func (RestAuthTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestAuthType)(nil)).Elem()
}

func (o RestAuthTypeOutput) ToRestAuthTypeOutput() RestAuthTypeOutput {
	return o
}

func (o RestAuthTypeOutput) ToRestAuthTypeOutputWithContext(ctx context.Context) RestAuthTypeOutput {
	return o
}

func (o RestAuthTypeOutput) ToRestAuthTypePtrOutput() RestAuthTypePtrOutput {
	return o.ToRestAuthTypePtrOutputWithContext(context.Background())
}

func (o RestAuthTypeOutput) ToRestAuthTypePtrOutputWithContext(ctx context.Context) RestAuthTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RestAuthType) *RestAuthType {
		return &v
	}).(RestAuthTypePtrOutput)
}

func (o RestAuthTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RestAuthTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RestAuthType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RestAuthTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RestAuthTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RestAuthType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RestAuthTypePtrOutput struct{ *pulumi.OutputState }

func (RestAuthTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestAuthType)(nil)).Elem()
}

func (o RestAuthTypePtrOutput) ToRestAuthTypePtrOutput() RestAuthTypePtrOutput {
	return o
}

func (o RestAuthTypePtrOutput) ToRestAuthTypePtrOutputWithContext(ctx context.Context) RestAuthTypePtrOutput {
	return o
}

func (o RestAuthTypePtrOutput) Elem() RestAuthTypeOutput {
	return o.ApplyT(func(v *RestAuthType) RestAuthType {
		if v != nil {
			return *v
		}
		var ret RestAuthType
		return ret
	}).(RestAuthTypeOutput)
}

func (o RestAuthTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RestAuthTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RestAuthType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RestAuthTypeInput interface {
	pulumi.Input

	ToRestAuthTypeOutput() RestAuthTypeOutput
	ToRestAuthTypeOutputWithContext(context.Context) RestAuthTypeOutput
}

var restAuthTypePtrType = reflect.TypeOf((**RestAuthType)(nil)).Elem()

type RestAuthTypePtrInput interface {
	pulumi.Input

	ToRestAuthTypePtrOutput() RestAuthTypePtrOutput
	ToRestAuthTypePtrOutputWithContext(context.Context) RestAuthTypePtrOutput
}

type restAuthTypePtr string

func RestAuthTypePtr(v string) RestAuthTypePtrInput {
	return (*restAuthTypePtr)(&v)
}

func (*restAuthTypePtr) ElementType() reflect.Type {
	return restAuthTypePtrType
}

func (in *restAuthTypePtr) ToRestAuthTypePtrOutput() RestAuthTypePtrOutput {
	return pulumi.ToOutput(in).(RestAuthTypePtrOutput)
}

func (in *restAuthTypePtr) ToRestAuthTypePtrOutputWithContext(ctx context.Context) RestAuthTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RestAuthTypePtrOutput)
}

type RestMatchQuantifier string

const (
	RestMatchQuantifierAll = RestMatchQuantifier("All")
	RestMatchQuantifierAny = RestMatchQuantifier("Any")
)

func (RestMatchQuantifier) ElementType() reflect.Type {
	return reflect.TypeOf((*RestMatchQuantifier)(nil)).Elem()
}

func (e RestMatchQuantifier) ToRestMatchQuantifierOutput() RestMatchQuantifierOutput {
	return pulumi.ToOutput(e).(RestMatchQuantifierOutput)
}

func (e RestMatchQuantifier) ToRestMatchQuantifierOutputWithContext(ctx context.Context) RestMatchQuantifierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RestMatchQuantifierOutput)
}

func (e RestMatchQuantifier) ToRestMatchQuantifierPtrOutput() RestMatchQuantifierPtrOutput {
	return e.ToRestMatchQuantifierPtrOutputWithContext(context.Background())
}

func (e RestMatchQuantifier) ToRestMatchQuantifierPtrOutputWithContext(ctx context.Context) RestMatchQuantifierPtrOutput {
	return RestMatchQuantifier(e).ToRestMatchQuantifierOutputWithContext(ctx).ToRestMatchQuantifierPtrOutputWithContext(ctx)
}

func (e RestMatchQuantifier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RestMatchQuantifier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RestMatchQuantifier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RestMatchQuantifier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RestMatchQuantifierOutput struct{ *pulumi.OutputState }

func (RestMatchQuantifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestMatchQuantifier)(nil)).Elem()
}

func (o RestMatchQuantifierOutput) ToRestMatchQuantifierOutput() RestMatchQuantifierOutput {
	return o
}

func (o RestMatchQuantifierOutput) ToRestMatchQuantifierOutputWithContext(ctx context.Context) RestMatchQuantifierOutput {
	return o
}

func (o RestMatchQuantifierOutput) ToRestMatchQuantifierPtrOutput() RestMatchQuantifierPtrOutput {
	return o.ToRestMatchQuantifierPtrOutputWithContext(context.Background())
}

func (o RestMatchQuantifierOutput) ToRestMatchQuantifierPtrOutputWithContext(ctx context.Context) RestMatchQuantifierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RestMatchQuantifier) *RestMatchQuantifier {
		return &v
	}).(RestMatchQuantifierPtrOutput)
}

func (o RestMatchQuantifierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RestMatchQuantifierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RestMatchQuantifier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RestMatchQuantifierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RestMatchQuantifierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RestMatchQuantifier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RestMatchQuantifierPtrOutput struct{ *pulumi.OutputState }

func (RestMatchQuantifierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestMatchQuantifier)(nil)).Elem()
}

func (o RestMatchQuantifierPtrOutput) ToRestMatchQuantifierPtrOutput() RestMatchQuantifierPtrOutput {
	return o
}

func (o RestMatchQuantifierPtrOutput) ToRestMatchQuantifierPtrOutputWithContext(ctx context.Context) RestMatchQuantifierPtrOutput {
	return o
}

func (o RestMatchQuantifierPtrOutput) Elem() RestMatchQuantifierOutput {
	return o.ApplyT(func(v *RestMatchQuantifier) RestMatchQuantifier {
		if v != nil {
			return *v
		}
		var ret RestMatchQuantifier
		return ret
	}).(RestMatchQuantifierOutput)
}

func (o RestMatchQuantifierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RestMatchQuantifierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RestMatchQuantifier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RestMatchQuantifierInput interface {
	pulumi.Input

	ToRestMatchQuantifierOutput() RestMatchQuantifierOutput
	ToRestMatchQuantifierOutputWithContext(context.Context) RestMatchQuantifierOutput
}

var restMatchQuantifierPtrType = reflect.TypeOf((**RestMatchQuantifier)(nil)).Elem()

type RestMatchQuantifierPtrInput interface {
	pulumi.Input

	ToRestMatchQuantifierPtrOutput() RestMatchQuantifierPtrOutput
	ToRestMatchQuantifierPtrOutputWithContext(context.Context) RestMatchQuantifierPtrOutput
}

type restMatchQuantifierPtr string

func RestMatchQuantifierPtr(v string) RestMatchQuantifierPtrInput {
	return (*restMatchQuantifierPtr)(&v)
}

func (*restMatchQuantifierPtr) ElementType() reflect.Type {
	return restMatchQuantifierPtrType
}

func (in *restMatchQuantifierPtr) ToRestMatchQuantifierPtrOutput() RestMatchQuantifierPtrOutput {
	return pulumi.ToOutput(in).(RestMatchQuantifierPtrOutput)
}

func (in *restMatchQuantifierPtr) ToRestMatchQuantifierPtrOutputWithContext(ctx context.Context) RestMatchQuantifierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RestMatchQuantifierPtrOutput)
}

type RestRequestMethod string

const (
	RestRequestMethodGET  = RestRequestMethod("GET")
	RestRequestMethodPOST = RestRequestMethod("POST")
)

func (RestRequestMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*RestRequestMethod)(nil)).Elem()
}

func (e RestRequestMethod) ToRestRequestMethodOutput() RestRequestMethodOutput {
	return pulumi.ToOutput(e).(RestRequestMethodOutput)
}

func (e RestRequestMethod) ToRestRequestMethodOutputWithContext(ctx context.Context) RestRequestMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RestRequestMethodOutput)
}

func (e RestRequestMethod) ToRestRequestMethodPtrOutput() RestRequestMethodPtrOutput {
	return e.ToRestRequestMethodPtrOutputWithContext(context.Background())
}

func (e RestRequestMethod) ToRestRequestMethodPtrOutputWithContext(ctx context.Context) RestRequestMethodPtrOutput {
	return RestRequestMethod(e).ToRestRequestMethodOutputWithContext(ctx).ToRestRequestMethodPtrOutputWithContext(ctx)
}

func (e RestRequestMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RestRequestMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RestRequestMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RestRequestMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RestRequestMethodOutput struct{ *pulumi.OutputState }

func (RestRequestMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestRequestMethod)(nil)).Elem()
}

func (o RestRequestMethodOutput) ToRestRequestMethodOutput() RestRequestMethodOutput {
	return o
}

func (o RestRequestMethodOutput) ToRestRequestMethodOutputWithContext(ctx context.Context) RestRequestMethodOutput {
	return o
}

func (o RestRequestMethodOutput) ToRestRequestMethodPtrOutput() RestRequestMethodPtrOutput {
	return o.ToRestRequestMethodPtrOutputWithContext(context.Background())
}

func (o RestRequestMethodOutput) ToRestRequestMethodPtrOutputWithContext(ctx context.Context) RestRequestMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RestRequestMethod) *RestRequestMethod {
		return &v
	}).(RestRequestMethodPtrOutput)
}

func (o RestRequestMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RestRequestMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RestRequestMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RestRequestMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RestRequestMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RestRequestMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RestRequestMethodPtrOutput struct{ *pulumi.OutputState }

func (RestRequestMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestRequestMethod)(nil)).Elem()
}

func (o RestRequestMethodPtrOutput) ToRestRequestMethodPtrOutput() RestRequestMethodPtrOutput {
	return o
}

func (o RestRequestMethodPtrOutput) ToRestRequestMethodPtrOutputWithContext(ctx context.Context) RestRequestMethodPtrOutput {
	return o
}

func (o RestRequestMethodPtrOutput) Elem() RestRequestMethodOutput {
	return o.ApplyT(func(v *RestRequestMethod) RestRequestMethod {
		if v != nil {
			return *v
		}
		var ret RestRequestMethod
		return ret
	}).(RestRequestMethodOutput)
}

func (o RestRequestMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RestRequestMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RestRequestMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RestRequestMethodInput interface {
	pulumi.Input

	ToRestRequestMethodOutput() RestRequestMethodOutput
	ToRestRequestMethodOutputWithContext(context.Context) RestRequestMethodOutput
}

var restRequestMethodPtrType = reflect.TypeOf((**RestRequestMethod)(nil)).Elem()

type RestRequestMethodPtrInput interface {
	pulumi.Input

	ToRestRequestMethodPtrOutput() RestRequestMethodPtrOutput
	ToRestRequestMethodPtrOutputWithContext(context.Context) RestRequestMethodPtrOutput
}

type restRequestMethodPtr string

func RestRequestMethodPtr(v string) RestRequestMethodPtrInput {
	return (*restRequestMethodPtr)(&v)
}

func (*restRequestMethodPtr) ElementType() reflect.Type {
	return restRequestMethodPtrType
}

func (in *restRequestMethodPtr) ToRestRequestMethodPtrOutput() RestRequestMethodPtrOutput {
	return pulumi.ToOutput(in).(RestRequestMethodPtrOutput)
}

func (in *restRequestMethodPtr) ToRestRequestMethodPtrOutputWithContext(ctx context.Context) RestRequestMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RestRequestMethodPtrOutput)
}

type StepType string

const (
	StepTypeWait        = StepType("Wait")
	StepTypeHealthCheck = StepType("HealthCheck")
)

func (StepType) ElementType() reflect.Type {
	return reflect.TypeOf((*StepType)(nil)).Elem()
}

func (e StepType) ToStepTypeOutput() StepTypeOutput {
	return pulumi.ToOutput(e).(StepTypeOutput)
}

func (e StepType) ToStepTypeOutputWithContext(ctx context.Context) StepTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StepTypeOutput)
}

func (e StepType) ToStepTypePtrOutput() StepTypePtrOutput {
	return e.ToStepTypePtrOutputWithContext(context.Background())
}

func (e StepType) ToStepTypePtrOutputWithContext(ctx context.Context) StepTypePtrOutput {
	return StepType(e).ToStepTypeOutputWithContext(ctx).ToStepTypePtrOutputWithContext(ctx)
}

func (e StepType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StepType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StepType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StepType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
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

func (o StepTypeOutput) ToStepTypePtrOutput() StepTypePtrOutput {
	return o.ToStepTypePtrOutputWithContext(context.Background())
}

func (o StepTypeOutput) ToStepTypePtrOutputWithContext(ctx context.Context) StepTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StepType) *StepType {
		return &v
	}).(StepTypePtrOutput)
}

func (o StepTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StepTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StepType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StepTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StepTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StepType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StepTypePtrOutput struct{ *pulumi.OutputState }

func (StepTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StepType)(nil)).Elem()
}

func (o StepTypePtrOutput) ToStepTypePtrOutput() StepTypePtrOutput {
	return o
}

func (o StepTypePtrOutput) ToStepTypePtrOutputWithContext(ctx context.Context) StepTypePtrOutput {
	return o
}

func (o StepTypePtrOutput) Elem() StepTypeOutput {
	return o.ApplyT(func(v *StepType) StepType {
		if v != nil {
			return *v
		}
		var ret StepType
		return ret
	}).(StepTypeOutput)
}

func (o StepTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StepTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StepType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StepTypeInput interface {
	pulumi.Input

	ToStepTypeOutput() StepTypeOutput
	ToStepTypeOutputWithContext(context.Context) StepTypeOutput
}

var stepTypePtrType = reflect.TypeOf((**StepType)(nil)).Elem()

type StepTypePtrInput interface {
	pulumi.Input

	ToStepTypePtrOutput() StepTypePtrOutput
	ToStepTypePtrOutputWithContext(context.Context) StepTypePtrOutput
}

type stepTypePtr string

func StepTypePtr(v string) StepTypePtrInput {
	return (*stepTypePtr)(&v)
}

func (*stepTypePtr) ElementType() reflect.Type {
	return stepTypePtrType
}

func (in *stepTypePtr) ToStepTypePtrOutput() StepTypePtrOutput {
	return pulumi.ToOutput(in).(StepTypePtrOutput)
}

func (in *stepTypePtr) ToStepTypePtrOutputWithContext(ctx context.Context) StepTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StepTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentModeOutput{})
	pulumi.RegisterOutputType(DeploymentModePtrOutput{})
	pulumi.RegisterOutputType(RestAuthLocationOutput{})
	pulumi.RegisterOutputType(RestAuthLocationPtrOutput{})
	pulumi.RegisterOutputType(RestAuthTypeOutput{})
	pulumi.RegisterOutputType(RestAuthTypePtrOutput{})
	pulumi.RegisterOutputType(RestMatchQuantifierOutput{})
	pulumi.RegisterOutputType(RestMatchQuantifierPtrOutput{})
	pulumi.RegisterOutputType(RestRequestMethodOutput{})
	pulumi.RegisterOutputType(RestRequestMethodPtrOutput{})
	pulumi.RegisterOutputType(StepTypeOutput{})
	pulumi.RegisterOutputType(StepTypePtrOutput{})
}
