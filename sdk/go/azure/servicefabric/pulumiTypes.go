


package servicefabric

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationDeltaHealthPolicy struct {
	DefaultServiceTypeDeltaHealthPolicy *ServiceTypeDeltaHealthPolicy           `pulumi:"defaultServiceTypeDeltaHealthPolicy"`
	ServiceTypeDeltaHealthPolicies      map[string]ServiceTypeDeltaHealthPolicy `pulumi:"serviceTypeDeltaHealthPolicies"`
}


func (val *ApplicationDeltaHealthPolicy) Defaults() *ApplicationDeltaHealthPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DefaultServiceTypeDeltaHealthPolicy = tmp.DefaultServiceTypeDeltaHealthPolicy.Defaults()

	return &tmp
}





type ApplicationDeltaHealthPolicyInput interface {
	pulumi.Input

	ToApplicationDeltaHealthPolicyOutput() ApplicationDeltaHealthPolicyOutput
	ToApplicationDeltaHealthPolicyOutputWithContext(context.Context) ApplicationDeltaHealthPolicyOutput
}

type ApplicationDeltaHealthPolicyArgs struct {
	DefaultServiceTypeDeltaHealthPolicy ServiceTypeDeltaHealthPolicyPtrInput `pulumi:"defaultServiceTypeDeltaHealthPolicy"`
	ServiceTypeDeltaHealthPolicies      ServiceTypeDeltaHealthPolicyMapInput `pulumi:"serviceTypeDeltaHealthPolicies"`
}

func (ApplicationDeltaHealthPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationDeltaHealthPolicy)(nil)).Elem()
}

func (i ApplicationDeltaHealthPolicyArgs) ToApplicationDeltaHealthPolicyOutput() ApplicationDeltaHealthPolicyOutput {
	return i.ToApplicationDeltaHealthPolicyOutputWithContext(context.Background())
}

func (i ApplicationDeltaHealthPolicyArgs) ToApplicationDeltaHealthPolicyOutputWithContext(ctx context.Context) ApplicationDeltaHealthPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationDeltaHealthPolicyOutput)
}





type ApplicationDeltaHealthPolicyMapInput interface {
	pulumi.Input

	ToApplicationDeltaHealthPolicyMapOutput() ApplicationDeltaHealthPolicyMapOutput
	ToApplicationDeltaHealthPolicyMapOutputWithContext(context.Context) ApplicationDeltaHealthPolicyMapOutput
}

type ApplicationDeltaHealthPolicyMap map[string]ApplicationDeltaHealthPolicyInput

func (ApplicationDeltaHealthPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApplicationDeltaHealthPolicy)(nil)).Elem()
}

func (i ApplicationDeltaHealthPolicyMap) ToApplicationDeltaHealthPolicyMapOutput() ApplicationDeltaHealthPolicyMapOutput {
	return i.ToApplicationDeltaHealthPolicyMapOutputWithContext(context.Background())
}

func (i ApplicationDeltaHealthPolicyMap) ToApplicationDeltaHealthPolicyMapOutputWithContext(ctx context.Context) ApplicationDeltaHealthPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationDeltaHealthPolicyMapOutput)
}

type ApplicationDeltaHealthPolicyOutput struct{ *pulumi.OutputState }

func (ApplicationDeltaHealthPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationDeltaHealthPolicy)(nil)).Elem()
}

func (o ApplicationDeltaHealthPolicyOutput) ToApplicationDeltaHealthPolicyOutput() ApplicationDeltaHealthPolicyOutput {
	return o
}

func (o ApplicationDeltaHealthPolicyOutput) ToApplicationDeltaHealthPolicyOutputWithContext(ctx context.Context) ApplicationDeltaHealthPolicyOutput {
	return o
}

func (o ApplicationDeltaHealthPolicyOutput) DefaultServiceTypeDeltaHealthPolicy() ServiceTypeDeltaHealthPolicyPtrOutput {
	return o.ApplyT(func(v ApplicationDeltaHealthPolicy) *ServiceTypeDeltaHealthPolicy {
		return v.DefaultServiceTypeDeltaHealthPolicy
	}).(ServiceTypeDeltaHealthPolicyPtrOutput)
}

func (o ApplicationDeltaHealthPolicyOutput) ServiceTypeDeltaHealthPolicies() ServiceTypeDeltaHealthPolicyMapOutput {
	return o.ApplyT(func(v ApplicationDeltaHealthPolicy) map[string]ServiceTypeDeltaHealthPolicy {
		return v.ServiceTypeDeltaHealthPolicies
	}).(ServiceTypeDeltaHealthPolicyMapOutput)
}

type ApplicationDeltaHealthPolicyMapOutput struct{ *pulumi.OutputState }

func (ApplicationDeltaHealthPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApplicationDeltaHealthPolicy)(nil)).Elem()
}

func (o ApplicationDeltaHealthPolicyMapOutput) ToApplicationDeltaHealthPolicyMapOutput() ApplicationDeltaHealthPolicyMapOutput {
	return o
}

func (o ApplicationDeltaHealthPolicyMapOutput) ToApplicationDeltaHealthPolicyMapOutputWithContext(ctx context.Context) ApplicationDeltaHealthPolicyMapOutput {
	return o
}

func (o ApplicationDeltaHealthPolicyMapOutput) MapIndex(k pulumi.StringInput) ApplicationDeltaHealthPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApplicationDeltaHealthPolicy {
		return vs[0].(map[string]ApplicationDeltaHealthPolicy)[vs[1].(string)]
	}).(ApplicationDeltaHealthPolicyOutput)
}

type ApplicationDeltaHealthPolicyResponse struct {
	DefaultServiceTypeDeltaHealthPolicy *ServiceTypeDeltaHealthPolicyResponse           `pulumi:"defaultServiceTypeDeltaHealthPolicy"`
	ServiceTypeDeltaHealthPolicies      map[string]ServiceTypeDeltaHealthPolicyResponse `pulumi:"serviceTypeDeltaHealthPolicies"`
}


func (val *ApplicationDeltaHealthPolicyResponse) Defaults() *ApplicationDeltaHealthPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DefaultServiceTypeDeltaHealthPolicy = tmp.DefaultServiceTypeDeltaHealthPolicy.Defaults()

	return &tmp
}





type ApplicationDeltaHealthPolicyResponseInput interface {
	pulumi.Input

	ToApplicationDeltaHealthPolicyResponseOutput() ApplicationDeltaHealthPolicyResponseOutput
	ToApplicationDeltaHealthPolicyResponseOutputWithContext(context.Context) ApplicationDeltaHealthPolicyResponseOutput
}

type ApplicationDeltaHealthPolicyResponseArgs struct {
	DefaultServiceTypeDeltaHealthPolicy ServiceTypeDeltaHealthPolicyResponsePtrInput `pulumi:"defaultServiceTypeDeltaHealthPolicy"`
	ServiceTypeDeltaHealthPolicies      ServiceTypeDeltaHealthPolicyResponseMapInput `pulumi:"serviceTypeDeltaHealthPolicies"`
}

func (ApplicationDeltaHealthPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationDeltaHealthPolicyResponse)(nil)).Elem()
}

func (i ApplicationDeltaHealthPolicyResponseArgs) ToApplicationDeltaHealthPolicyResponseOutput() ApplicationDeltaHealthPolicyResponseOutput {
	return i.ToApplicationDeltaHealthPolicyResponseOutputWithContext(context.Background())
}

func (i ApplicationDeltaHealthPolicyResponseArgs) ToApplicationDeltaHealthPolicyResponseOutputWithContext(ctx context.Context) ApplicationDeltaHealthPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationDeltaHealthPolicyResponseOutput)
}





type ApplicationDeltaHealthPolicyResponseMapInput interface {
	pulumi.Input

	ToApplicationDeltaHealthPolicyResponseMapOutput() ApplicationDeltaHealthPolicyResponseMapOutput
	ToApplicationDeltaHealthPolicyResponseMapOutputWithContext(context.Context) ApplicationDeltaHealthPolicyResponseMapOutput
}

type ApplicationDeltaHealthPolicyResponseMap map[string]ApplicationDeltaHealthPolicyResponseInput

func (ApplicationDeltaHealthPolicyResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApplicationDeltaHealthPolicyResponse)(nil)).Elem()
}

func (i ApplicationDeltaHealthPolicyResponseMap) ToApplicationDeltaHealthPolicyResponseMapOutput() ApplicationDeltaHealthPolicyResponseMapOutput {
	return i.ToApplicationDeltaHealthPolicyResponseMapOutputWithContext(context.Background())
}

func (i ApplicationDeltaHealthPolicyResponseMap) ToApplicationDeltaHealthPolicyResponseMapOutputWithContext(ctx context.Context) ApplicationDeltaHealthPolicyResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationDeltaHealthPolicyResponseMapOutput)
}

type ApplicationDeltaHealthPolicyResponseOutput struct{ *pulumi.OutputState }

func (ApplicationDeltaHealthPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationDeltaHealthPolicyResponse)(nil)).Elem()
}

func (o ApplicationDeltaHealthPolicyResponseOutput) ToApplicationDeltaHealthPolicyResponseOutput() ApplicationDeltaHealthPolicyResponseOutput {
	return o
}

func (o ApplicationDeltaHealthPolicyResponseOutput) ToApplicationDeltaHealthPolicyResponseOutputWithContext(ctx context.Context) ApplicationDeltaHealthPolicyResponseOutput {
	return o
}

func (o ApplicationDeltaHealthPolicyResponseOutput) DefaultServiceTypeDeltaHealthPolicy() ServiceTypeDeltaHealthPolicyResponsePtrOutput {
	return o.ApplyT(func(v ApplicationDeltaHealthPolicyResponse) *ServiceTypeDeltaHealthPolicyResponse {
		return v.DefaultServiceTypeDeltaHealthPolicy
	}).(ServiceTypeDeltaHealthPolicyResponsePtrOutput)
}

func (o ApplicationDeltaHealthPolicyResponseOutput) ServiceTypeDeltaHealthPolicies() ServiceTypeDeltaHealthPolicyResponseMapOutput {
	return o.ApplyT(func(v ApplicationDeltaHealthPolicyResponse) map[string]ServiceTypeDeltaHealthPolicyResponse {
		return v.ServiceTypeDeltaHealthPolicies
	}).(ServiceTypeDeltaHealthPolicyResponseMapOutput)
}

type ApplicationDeltaHealthPolicyResponseMapOutput struct{ *pulumi.OutputState }

func (ApplicationDeltaHealthPolicyResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApplicationDeltaHealthPolicyResponse)(nil)).Elem()
}

func (o ApplicationDeltaHealthPolicyResponseMapOutput) ToApplicationDeltaHealthPolicyResponseMapOutput() ApplicationDeltaHealthPolicyResponseMapOutput {
	return o
}

func (o ApplicationDeltaHealthPolicyResponseMapOutput) ToApplicationDeltaHealthPolicyResponseMapOutputWithContext(ctx context.Context) ApplicationDeltaHealthPolicyResponseMapOutput {
	return o
}

func (o ApplicationDeltaHealthPolicyResponseMapOutput) MapIndex(k pulumi.StringInput) ApplicationDeltaHealthPolicyResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApplicationDeltaHealthPolicyResponse {
		return vs[0].(map[string]ApplicationDeltaHealthPolicyResponse)[vs[1].(string)]
	}).(ApplicationDeltaHealthPolicyResponseOutput)
}

type ApplicationHealthPolicy struct {
	DefaultServiceTypeHealthPolicy *ServiceTypeHealthPolicy           `pulumi:"defaultServiceTypeHealthPolicy"`
	ServiceTypeHealthPolicies      map[string]ServiceTypeHealthPolicy `pulumi:"serviceTypeHealthPolicies"`
}


func (val *ApplicationHealthPolicy) Defaults() *ApplicationHealthPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DefaultServiceTypeHealthPolicy = tmp.DefaultServiceTypeHealthPolicy.Defaults()

	return &tmp
}





type ApplicationHealthPolicyInput interface {
	pulumi.Input

	ToApplicationHealthPolicyOutput() ApplicationHealthPolicyOutput
	ToApplicationHealthPolicyOutputWithContext(context.Context) ApplicationHealthPolicyOutput
}

type ApplicationHealthPolicyArgs struct {
	DefaultServiceTypeHealthPolicy ServiceTypeHealthPolicyPtrInput `pulumi:"defaultServiceTypeHealthPolicy"`
	ServiceTypeHealthPolicies      ServiceTypeHealthPolicyMapInput `pulumi:"serviceTypeHealthPolicies"`
}

func (ApplicationHealthPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationHealthPolicy)(nil)).Elem()
}

func (i ApplicationHealthPolicyArgs) ToApplicationHealthPolicyOutput() ApplicationHealthPolicyOutput {
	return i.ToApplicationHealthPolicyOutputWithContext(context.Background())
}

func (i ApplicationHealthPolicyArgs) ToApplicationHealthPolicyOutputWithContext(ctx context.Context) ApplicationHealthPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationHealthPolicyOutput)
}





type ApplicationHealthPolicyMapInput interface {
	pulumi.Input

	ToApplicationHealthPolicyMapOutput() ApplicationHealthPolicyMapOutput
	ToApplicationHealthPolicyMapOutputWithContext(context.Context) ApplicationHealthPolicyMapOutput
}

type ApplicationHealthPolicyMap map[string]ApplicationHealthPolicyInput

func (ApplicationHealthPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApplicationHealthPolicy)(nil)).Elem()
}

func (i ApplicationHealthPolicyMap) ToApplicationHealthPolicyMapOutput() ApplicationHealthPolicyMapOutput {
	return i.ToApplicationHealthPolicyMapOutputWithContext(context.Background())
}

func (i ApplicationHealthPolicyMap) ToApplicationHealthPolicyMapOutputWithContext(ctx context.Context) ApplicationHealthPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationHealthPolicyMapOutput)
}

type ApplicationHealthPolicyOutput struct{ *pulumi.OutputState }

func (ApplicationHealthPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationHealthPolicy)(nil)).Elem()
}

func (o ApplicationHealthPolicyOutput) ToApplicationHealthPolicyOutput() ApplicationHealthPolicyOutput {
	return o
}

func (o ApplicationHealthPolicyOutput) ToApplicationHealthPolicyOutputWithContext(ctx context.Context) ApplicationHealthPolicyOutput {
	return o
}

func (o ApplicationHealthPolicyOutput) DefaultServiceTypeHealthPolicy() ServiceTypeHealthPolicyPtrOutput {
	return o.ApplyT(func(v ApplicationHealthPolicy) *ServiceTypeHealthPolicy { return v.DefaultServiceTypeHealthPolicy }).(ServiceTypeHealthPolicyPtrOutput)
}

func (o ApplicationHealthPolicyOutput) ServiceTypeHealthPolicies() ServiceTypeHealthPolicyMapOutput {
	return o.ApplyT(func(v ApplicationHealthPolicy) map[string]ServiceTypeHealthPolicy { return v.ServiceTypeHealthPolicies }).(ServiceTypeHealthPolicyMapOutput)
}

type ApplicationHealthPolicyMapOutput struct{ *pulumi.OutputState }

func (ApplicationHealthPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApplicationHealthPolicy)(nil)).Elem()
}

func (o ApplicationHealthPolicyMapOutput) ToApplicationHealthPolicyMapOutput() ApplicationHealthPolicyMapOutput {
	return o
}

func (o ApplicationHealthPolicyMapOutput) ToApplicationHealthPolicyMapOutputWithContext(ctx context.Context) ApplicationHealthPolicyMapOutput {
	return o
}

func (o ApplicationHealthPolicyMapOutput) MapIndex(k pulumi.StringInput) ApplicationHealthPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApplicationHealthPolicy {
		return vs[0].(map[string]ApplicationHealthPolicy)[vs[1].(string)]
	}).(ApplicationHealthPolicyOutput)
}

type ApplicationHealthPolicyResponse struct {
	DefaultServiceTypeHealthPolicy *ServiceTypeHealthPolicyResponse           `pulumi:"defaultServiceTypeHealthPolicy"`
	ServiceTypeHealthPolicies      map[string]ServiceTypeHealthPolicyResponse `pulumi:"serviceTypeHealthPolicies"`
}


func (val *ApplicationHealthPolicyResponse) Defaults() *ApplicationHealthPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DefaultServiceTypeHealthPolicy = tmp.DefaultServiceTypeHealthPolicy.Defaults()

	return &tmp
}





type ApplicationHealthPolicyResponseInput interface {
	pulumi.Input

	ToApplicationHealthPolicyResponseOutput() ApplicationHealthPolicyResponseOutput
	ToApplicationHealthPolicyResponseOutputWithContext(context.Context) ApplicationHealthPolicyResponseOutput
}

type ApplicationHealthPolicyResponseArgs struct {
	DefaultServiceTypeHealthPolicy ServiceTypeHealthPolicyResponsePtrInput `pulumi:"defaultServiceTypeHealthPolicy"`
	ServiceTypeHealthPolicies      ServiceTypeHealthPolicyResponseMapInput `pulumi:"serviceTypeHealthPolicies"`
}

func (ApplicationHealthPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationHealthPolicyResponse)(nil)).Elem()
}

func (i ApplicationHealthPolicyResponseArgs) ToApplicationHealthPolicyResponseOutput() ApplicationHealthPolicyResponseOutput {
	return i.ToApplicationHealthPolicyResponseOutputWithContext(context.Background())
}

func (i ApplicationHealthPolicyResponseArgs) ToApplicationHealthPolicyResponseOutputWithContext(ctx context.Context) ApplicationHealthPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationHealthPolicyResponseOutput)
}





type ApplicationHealthPolicyResponseMapInput interface {
	pulumi.Input

	ToApplicationHealthPolicyResponseMapOutput() ApplicationHealthPolicyResponseMapOutput
	ToApplicationHealthPolicyResponseMapOutputWithContext(context.Context) ApplicationHealthPolicyResponseMapOutput
}

type ApplicationHealthPolicyResponseMap map[string]ApplicationHealthPolicyResponseInput

func (ApplicationHealthPolicyResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApplicationHealthPolicyResponse)(nil)).Elem()
}

func (i ApplicationHealthPolicyResponseMap) ToApplicationHealthPolicyResponseMapOutput() ApplicationHealthPolicyResponseMapOutput {
	return i.ToApplicationHealthPolicyResponseMapOutputWithContext(context.Background())
}

func (i ApplicationHealthPolicyResponseMap) ToApplicationHealthPolicyResponseMapOutputWithContext(ctx context.Context) ApplicationHealthPolicyResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationHealthPolicyResponseMapOutput)
}

type ApplicationHealthPolicyResponseOutput struct{ *pulumi.OutputState }

func (ApplicationHealthPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationHealthPolicyResponse)(nil)).Elem()
}

func (o ApplicationHealthPolicyResponseOutput) ToApplicationHealthPolicyResponseOutput() ApplicationHealthPolicyResponseOutput {
	return o
}

func (o ApplicationHealthPolicyResponseOutput) ToApplicationHealthPolicyResponseOutputWithContext(ctx context.Context) ApplicationHealthPolicyResponseOutput {
	return o
}

func (o ApplicationHealthPolicyResponseOutput) DefaultServiceTypeHealthPolicy() ServiceTypeHealthPolicyResponsePtrOutput {
	return o.ApplyT(func(v ApplicationHealthPolicyResponse) *ServiceTypeHealthPolicyResponse {
		return v.DefaultServiceTypeHealthPolicy
	}).(ServiceTypeHealthPolicyResponsePtrOutput)
}

func (o ApplicationHealthPolicyResponseOutput) ServiceTypeHealthPolicies() ServiceTypeHealthPolicyResponseMapOutput {
	return o.ApplyT(func(v ApplicationHealthPolicyResponse) map[string]ServiceTypeHealthPolicyResponse {
		return v.ServiceTypeHealthPolicies
	}).(ServiceTypeHealthPolicyResponseMapOutput)
}

type ApplicationHealthPolicyResponseMapOutput struct{ *pulumi.OutputState }

func (ApplicationHealthPolicyResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApplicationHealthPolicyResponse)(nil)).Elem()
}

func (o ApplicationHealthPolicyResponseMapOutput) ToApplicationHealthPolicyResponseMapOutput() ApplicationHealthPolicyResponseMapOutput {
	return o
}

func (o ApplicationHealthPolicyResponseMapOutput) ToApplicationHealthPolicyResponseMapOutputWithContext(ctx context.Context) ApplicationHealthPolicyResponseMapOutput {
	return o
}

func (o ApplicationHealthPolicyResponseMapOutput) MapIndex(k pulumi.StringInput) ApplicationHealthPolicyResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApplicationHealthPolicyResponse {
		return vs[0].(map[string]ApplicationHealthPolicyResponse)[vs[1].(string)]
	}).(ApplicationHealthPolicyResponseOutput)
}

type ApplicationMetricDescription struct {
	MaximumCapacity          *float64 `pulumi:"maximumCapacity"`
	Name                     *string  `pulumi:"name"`
	ReservationCapacity      *float64 `pulumi:"reservationCapacity"`
	TotalApplicationCapacity *float64 `pulumi:"totalApplicationCapacity"`
}





type ApplicationMetricDescriptionInput interface {
	pulumi.Input

	ToApplicationMetricDescriptionOutput() ApplicationMetricDescriptionOutput
	ToApplicationMetricDescriptionOutputWithContext(context.Context) ApplicationMetricDescriptionOutput
}

type ApplicationMetricDescriptionArgs struct {
	MaximumCapacity          pulumi.Float64PtrInput `pulumi:"maximumCapacity"`
	Name                     pulumi.StringPtrInput  `pulumi:"name"`
	ReservationCapacity      pulumi.Float64PtrInput `pulumi:"reservationCapacity"`
	TotalApplicationCapacity pulumi.Float64PtrInput `pulumi:"totalApplicationCapacity"`
}

func (ApplicationMetricDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationMetricDescription)(nil)).Elem()
}

func (i ApplicationMetricDescriptionArgs) ToApplicationMetricDescriptionOutput() ApplicationMetricDescriptionOutput {
	return i.ToApplicationMetricDescriptionOutputWithContext(context.Background())
}

func (i ApplicationMetricDescriptionArgs) ToApplicationMetricDescriptionOutputWithContext(ctx context.Context) ApplicationMetricDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationMetricDescriptionOutput)
}





type ApplicationMetricDescriptionArrayInput interface {
	pulumi.Input

	ToApplicationMetricDescriptionArrayOutput() ApplicationMetricDescriptionArrayOutput
	ToApplicationMetricDescriptionArrayOutputWithContext(context.Context) ApplicationMetricDescriptionArrayOutput
}

type ApplicationMetricDescriptionArray []ApplicationMetricDescriptionInput

func (ApplicationMetricDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationMetricDescription)(nil)).Elem()
}

func (i ApplicationMetricDescriptionArray) ToApplicationMetricDescriptionArrayOutput() ApplicationMetricDescriptionArrayOutput {
	return i.ToApplicationMetricDescriptionArrayOutputWithContext(context.Background())
}

func (i ApplicationMetricDescriptionArray) ToApplicationMetricDescriptionArrayOutputWithContext(ctx context.Context) ApplicationMetricDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationMetricDescriptionArrayOutput)
}

type ApplicationMetricDescriptionOutput struct{ *pulumi.OutputState }

func (ApplicationMetricDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationMetricDescription)(nil)).Elem()
}

func (o ApplicationMetricDescriptionOutput) ToApplicationMetricDescriptionOutput() ApplicationMetricDescriptionOutput {
	return o
}

func (o ApplicationMetricDescriptionOutput) ToApplicationMetricDescriptionOutputWithContext(ctx context.Context) ApplicationMetricDescriptionOutput {
	return o
}

func (o ApplicationMetricDescriptionOutput) MaximumCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ApplicationMetricDescription) *float64 { return v.MaximumCapacity }).(pulumi.Float64PtrOutput)
}

func (o ApplicationMetricDescriptionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationMetricDescription) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationMetricDescriptionOutput) ReservationCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ApplicationMetricDescription) *float64 { return v.ReservationCapacity }).(pulumi.Float64PtrOutput)
}

func (o ApplicationMetricDescriptionOutput) TotalApplicationCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ApplicationMetricDescription) *float64 { return v.TotalApplicationCapacity }).(pulumi.Float64PtrOutput)
}

type ApplicationMetricDescriptionArrayOutput struct{ *pulumi.OutputState }

func (ApplicationMetricDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationMetricDescription)(nil)).Elem()
}

func (o ApplicationMetricDescriptionArrayOutput) ToApplicationMetricDescriptionArrayOutput() ApplicationMetricDescriptionArrayOutput {
	return o
}

func (o ApplicationMetricDescriptionArrayOutput) ToApplicationMetricDescriptionArrayOutputWithContext(ctx context.Context) ApplicationMetricDescriptionArrayOutput {
	return o
}

func (o ApplicationMetricDescriptionArrayOutput) Index(i pulumi.IntInput) ApplicationMetricDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationMetricDescription {
		return vs[0].([]ApplicationMetricDescription)[vs[1].(int)]
	}).(ApplicationMetricDescriptionOutput)
}

type ApplicationMetricDescriptionResponse struct {
	MaximumCapacity          *float64 `pulumi:"maximumCapacity"`
	Name                     *string  `pulumi:"name"`
	ReservationCapacity      *float64 `pulumi:"reservationCapacity"`
	TotalApplicationCapacity *float64 `pulumi:"totalApplicationCapacity"`
}





type ApplicationMetricDescriptionResponseInput interface {
	pulumi.Input

	ToApplicationMetricDescriptionResponseOutput() ApplicationMetricDescriptionResponseOutput
	ToApplicationMetricDescriptionResponseOutputWithContext(context.Context) ApplicationMetricDescriptionResponseOutput
}

type ApplicationMetricDescriptionResponseArgs struct {
	MaximumCapacity          pulumi.Float64PtrInput `pulumi:"maximumCapacity"`
	Name                     pulumi.StringPtrInput  `pulumi:"name"`
	ReservationCapacity      pulumi.Float64PtrInput `pulumi:"reservationCapacity"`
	TotalApplicationCapacity pulumi.Float64PtrInput `pulumi:"totalApplicationCapacity"`
}

func (ApplicationMetricDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationMetricDescriptionResponse)(nil)).Elem()
}

func (i ApplicationMetricDescriptionResponseArgs) ToApplicationMetricDescriptionResponseOutput() ApplicationMetricDescriptionResponseOutput {
	return i.ToApplicationMetricDescriptionResponseOutputWithContext(context.Background())
}

func (i ApplicationMetricDescriptionResponseArgs) ToApplicationMetricDescriptionResponseOutputWithContext(ctx context.Context) ApplicationMetricDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationMetricDescriptionResponseOutput)
}





type ApplicationMetricDescriptionResponseArrayInput interface {
	pulumi.Input

	ToApplicationMetricDescriptionResponseArrayOutput() ApplicationMetricDescriptionResponseArrayOutput
	ToApplicationMetricDescriptionResponseArrayOutputWithContext(context.Context) ApplicationMetricDescriptionResponseArrayOutput
}

type ApplicationMetricDescriptionResponseArray []ApplicationMetricDescriptionResponseInput

func (ApplicationMetricDescriptionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationMetricDescriptionResponse)(nil)).Elem()
}

func (i ApplicationMetricDescriptionResponseArray) ToApplicationMetricDescriptionResponseArrayOutput() ApplicationMetricDescriptionResponseArrayOutput {
	return i.ToApplicationMetricDescriptionResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationMetricDescriptionResponseArray) ToApplicationMetricDescriptionResponseArrayOutputWithContext(ctx context.Context) ApplicationMetricDescriptionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationMetricDescriptionResponseArrayOutput)
}

type ApplicationMetricDescriptionResponseOutput struct{ *pulumi.OutputState }

func (ApplicationMetricDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationMetricDescriptionResponse)(nil)).Elem()
}

func (o ApplicationMetricDescriptionResponseOutput) ToApplicationMetricDescriptionResponseOutput() ApplicationMetricDescriptionResponseOutput {
	return o
}

func (o ApplicationMetricDescriptionResponseOutput) ToApplicationMetricDescriptionResponseOutputWithContext(ctx context.Context) ApplicationMetricDescriptionResponseOutput {
	return o
}

func (o ApplicationMetricDescriptionResponseOutput) MaximumCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ApplicationMetricDescriptionResponse) *float64 { return v.MaximumCapacity }).(pulumi.Float64PtrOutput)
}

func (o ApplicationMetricDescriptionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationMetricDescriptionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationMetricDescriptionResponseOutput) ReservationCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ApplicationMetricDescriptionResponse) *float64 { return v.ReservationCapacity }).(pulumi.Float64PtrOutput)
}

func (o ApplicationMetricDescriptionResponseOutput) TotalApplicationCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ApplicationMetricDescriptionResponse) *float64 { return v.TotalApplicationCapacity }).(pulumi.Float64PtrOutput)
}

type ApplicationMetricDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationMetricDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationMetricDescriptionResponse)(nil)).Elem()
}

func (o ApplicationMetricDescriptionResponseArrayOutput) ToApplicationMetricDescriptionResponseArrayOutput() ApplicationMetricDescriptionResponseArrayOutput {
	return o
}

func (o ApplicationMetricDescriptionResponseArrayOutput) ToApplicationMetricDescriptionResponseArrayOutputWithContext(ctx context.Context) ApplicationMetricDescriptionResponseArrayOutput {
	return o
}

func (o ApplicationMetricDescriptionResponseArrayOutput) Index(i pulumi.IntInput) ApplicationMetricDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationMetricDescriptionResponse {
		return vs[0].([]ApplicationMetricDescriptionResponse)[vs[1].(int)]
	}).(ApplicationMetricDescriptionResponseOutput)
}

type ApplicationTypeVersionsCleanupPolicy struct {
	MaxUnusedVersionsToKeep float64 `pulumi:"maxUnusedVersionsToKeep"`
}





type ApplicationTypeVersionsCleanupPolicyInput interface {
	pulumi.Input

	ToApplicationTypeVersionsCleanupPolicyOutput() ApplicationTypeVersionsCleanupPolicyOutput
	ToApplicationTypeVersionsCleanupPolicyOutputWithContext(context.Context) ApplicationTypeVersionsCleanupPolicyOutput
}

type ApplicationTypeVersionsCleanupPolicyArgs struct {
	MaxUnusedVersionsToKeep pulumi.Float64Input `pulumi:"maxUnusedVersionsToKeep"`
}

func (ApplicationTypeVersionsCleanupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationTypeVersionsCleanupPolicy)(nil)).Elem()
}

func (i ApplicationTypeVersionsCleanupPolicyArgs) ToApplicationTypeVersionsCleanupPolicyOutput() ApplicationTypeVersionsCleanupPolicyOutput {
	return i.ToApplicationTypeVersionsCleanupPolicyOutputWithContext(context.Background())
}

func (i ApplicationTypeVersionsCleanupPolicyArgs) ToApplicationTypeVersionsCleanupPolicyOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTypeVersionsCleanupPolicyOutput)
}

func (i ApplicationTypeVersionsCleanupPolicyArgs) ToApplicationTypeVersionsCleanupPolicyPtrOutput() ApplicationTypeVersionsCleanupPolicyPtrOutput {
	return i.ToApplicationTypeVersionsCleanupPolicyPtrOutputWithContext(context.Background())
}

func (i ApplicationTypeVersionsCleanupPolicyArgs) ToApplicationTypeVersionsCleanupPolicyPtrOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTypeVersionsCleanupPolicyOutput).ToApplicationTypeVersionsCleanupPolicyPtrOutputWithContext(ctx)
}









type ApplicationTypeVersionsCleanupPolicyPtrInput interface {
	pulumi.Input

	ToApplicationTypeVersionsCleanupPolicyPtrOutput() ApplicationTypeVersionsCleanupPolicyPtrOutput
	ToApplicationTypeVersionsCleanupPolicyPtrOutputWithContext(context.Context) ApplicationTypeVersionsCleanupPolicyPtrOutput
}

type applicationTypeVersionsCleanupPolicyPtrType ApplicationTypeVersionsCleanupPolicyArgs

func ApplicationTypeVersionsCleanupPolicyPtr(v *ApplicationTypeVersionsCleanupPolicyArgs) ApplicationTypeVersionsCleanupPolicyPtrInput {
	return (*applicationTypeVersionsCleanupPolicyPtrType)(v)
}

func (*applicationTypeVersionsCleanupPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationTypeVersionsCleanupPolicy)(nil)).Elem()
}

func (i *applicationTypeVersionsCleanupPolicyPtrType) ToApplicationTypeVersionsCleanupPolicyPtrOutput() ApplicationTypeVersionsCleanupPolicyPtrOutput {
	return i.ToApplicationTypeVersionsCleanupPolicyPtrOutputWithContext(context.Background())
}

func (i *applicationTypeVersionsCleanupPolicyPtrType) ToApplicationTypeVersionsCleanupPolicyPtrOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTypeVersionsCleanupPolicyPtrOutput)
}

type ApplicationTypeVersionsCleanupPolicyOutput struct{ *pulumi.OutputState }

func (ApplicationTypeVersionsCleanupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationTypeVersionsCleanupPolicy)(nil)).Elem()
}

func (o ApplicationTypeVersionsCleanupPolicyOutput) ToApplicationTypeVersionsCleanupPolicyOutput() ApplicationTypeVersionsCleanupPolicyOutput {
	return o
}

func (o ApplicationTypeVersionsCleanupPolicyOutput) ToApplicationTypeVersionsCleanupPolicyOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyOutput {
	return o
}

func (o ApplicationTypeVersionsCleanupPolicyOutput) ToApplicationTypeVersionsCleanupPolicyPtrOutput() ApplicationTypeVersionsCleanupPolicyPtrOutput {
	return o.ToApplicationTypeVersionsCleanupPolicyPtrOutputWithContext(context.Background())
}

func (o ApplicationTypeVersionsCleanupPolicyOutput) ToApplicationTypeVersionsCleanupPolicyPtrOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationTypeVersionsCleanupPolicy) *ApplicationTypeVersionsCleanupPolicy {
		return &v
	}).(ApplicationTypeVersionsCleanupPolicyPtrOutput)
}

func (o ApplicationTypeVersionsCleanupPolicyOutput) MaxUnusedVersionsToKeep() pulumi.Float64Output {
	return o.ApplyT(func(v ApplicationTypeVersionsCleanupPolicy) float64 { return v.MaxUnusedVersionsToKeep }).(pulumi.Float64Output)
}

type ApplicationTypeVersionsCleanupPolicyPtrOutput struct{ *pulumi.OutputState }

func (ApplicationTypeVersionsCleanupPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationTypeVersionsCleanupPolicy)(nil)).Elem()
}

func (o ApplicationTypeVersionsCleanupPolicyPtrOutput) ToApplicationTypeVersionsCleanupPolicyPtrOutput() ApplicationTypeVersionsCleanupPolicyPtrOutput {
	return o
}

func (o ApplicationTypeVersionsCleanupPolicyPtrOutput) ToApplicationTypeVersionsCleanupPolicyPtrOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyPtrOutput {
	return o
}

func (o ApplicationTypeVersionsCleanupPolicyPtrOutput) Elem() ApplicationTypeVersionsCleanupPolicyOutput {
	return o.ApplyT(func(v *ApplicationTypeVersionsCleanupPolicy) ApplicationTypeVersionsCleanupPolicy {
		if v != nil {
			return *v
		}
		var ret ApplicationTypeVersionsCleanupPolicy
		return ret
	}).(ApplicationTypeVersionsCleanupPolicyOutput)
}

func (o ApplicationTypeVersionsCleanupPolicyPtrOutput) MaxUnusedVersionsToKeep() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ApplicationTypeVersionsCleanupPolicy) *float64 {
		if v == nil {
			return nil
		}
		return &v.MaxUnusedVersionsToKeep
	}).(pulumi.Float64PtrOutput)
}

type ApplicationTypeVersionsCleanupPolicyResponse struct {
	MaxUnusedVersionsToKeep float64 `pulumi:"maxUnusedVersionsToKeep"`
}





type ApplicationTypeVersionsCleanupPolicyResponseInput interface {
	pulumi.Input

	ToApplicationTypeVersionsCleanupPolicyResponseOutput() ApplicationTypeVersionsCleanupPolicyResponseOutput
	ToApplicationTypeVersionsCleanupPolicyResponseOutputWithContext(context.Context) ApplicationTypeVersionsCleanupPolicyResponseOutput
}

type ApplicationTypeVersionsCleanupPolicyResponseArgs struct {
	MaxUnusedVersionsToKeep pulumi.Float64Input `pulumi:"maxUnusedVersionsToKeep"`
}

func (ApplicationTypeVersionsCleanupPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationTypeVersionsCleanupPolicyResponse)(nil)).Elem()
}

func (i ApplicationTypeVersionsCleanupPolicyResponseArgs) ToApplicationTypeVersionsCleanupPolicyResponseOutput() ApplicationTypeVersionsCleanupPolicyResponseOutput {
	return i.ToApplicationTypeVersionsCleanupPolicyResponseOutputWithContext(context.Background())
}

func (i ApplicationTypeVersionsCleanupPolicyResponseArgs) ToApplicationTypeVersionsCleanupPolicyResponseOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTypeVersionsCleanupPolicyResponseOutput)
}

func (i ApplicationTypeVersionsCleanupPolicyResponseArgs) ToApplicationTypeVersionsCleanupPolicyResponsePtrOutput() ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return i.ToApplicationTypeVersionsCleanupPolicyResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationTypeVersionsCleanupPolicyResponseArgs) ToApplicationTypeVersionsCleanupPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTypeVersionsCleanupPolicyResponseOutput).ToApplicationTypeVersionsCleanupPolicyResponsePtrOutputWithContext(ctx)
}









type ApplicationTypeVersionsCleanupPolicyResponsePtrInput interface {
	pulumi.Input

	ToApplicationTypeVersionsCleanupPolicyResponsePtrOutput() ApplicationTypeVersionsCleanupPolicyResponsePtrOutput
	ToApplicationTypeVersionsCleanupPolicyResponsePtrOutputWithContext(context.Context) ApplicationTypeVersionsCleanupPolicyResponsePtrOutput
}

type applicationTypeVersionsCleanupPolicyResponsePtrType ApplicationTypeVersionsCleanupPolicyResponseArgs

func ApplicationTypeVersionsCleanupPolicyResponsePtr(v *ApplicationTypeVersionsCleanupPolicyResponseArgs) ApplicationTypeVersionsCleanupPolicyResponsePtrInput {
	return (*applicationTypeVersionsCleanupPolicyResponsePtrType)(v)
}

func (*applicationTypeVersionsCleanupPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationTypeVersionsCleanupPolicyResponse)(nil)).Elem()
}

func (i *applicationTypeVersionsCleanupPolicyResponsePtrType) ToApplicationTypeVersionsCleanupPolicyResponsePtrOutput() ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return i.ToApplicationTypeVersionsCleanupPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *applicationTypeVersionsCleanupPolicyResponsePtrType) ToApplicationTypeVersionsCleanupPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTypeVersionsCleanupPolicyResponsePtrOutput)
}

type ApplicationTypeVersionsCleanupPolicyResponseOutput struct{ *pulumi.OutputState }

func (ApplicationTypeVersionsCleanupPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationTypeVersionsCleanupPolicyResponse)(nil)).Elem()
}

func (o ApplicationTypeVersionsCleanupPolicyResponseOutput) ToApplicationTypeVersionsCleanupPolicyResponseOutput() ApplicationTypeVersionsCleanupPolicyResponseOutput {
	return o
}

func (o ApplicationTypeVersionsCleanupPolicyResponseOutput) ToApplicationTypeVersionsCleanupPolicyResponseOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyResponseOutput {
	return o
}

func (o ApplicationTypeVersionsCleanupPolicyResponseOutput) ToApplicationTypeVersionsCleanupPolicyResponsePtrOutput() ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return o.ToApplicationTypeVersionsCleanupPolicyResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationTypeVersionsCleanupPolicyResponseOutput) ToApplicationTypeVersionsCleanupPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationTypeVersionsCleanupPolicyResponse) *ApplicationTypeVersionsCleanupPolicyResponse {
		return &v
	}).(ApplicationTypeVersionsCleanupPolicyResponsePtrOutput)
}

func (o ApplicationTypeVersionsCleanupPolicyResponseOutput) MaxUnusedVersionsToKeep() pulumi.Float64Output {
	return o.ApplyT(func(v ApplicationTypeVersionsCleanupPolicyResponse) float64 { return v.MaxUnusedVersionsToKeep }).(pulumi.Float64Output)
}

type ApplicationTypeVersionsCleanupPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationTypeVersionsCleanupPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationTypeVersionsCleanupPolicyResponse)(nil)).Elem()
}

func (o ApplicationTypeVersionsCleanupPolicyResponsePtrOutput) ToApplicationTypeVersionsCleanupPolicyResponsePtrOutput() ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return o
}

func (o ApplicationTypeVersionsCleanupPolicyResponsePtrOutput) ToApplicationTypeVersionsCleanupPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return o
}

func (o ApplicationTypeVersionsCleanupPolicyResponsePtrOutput) Elem() ApplicationTypeVersionsCleanupPolicyResponseOutput {
	return o.ApplyT(func(v *ApplicationTypeVersionsCleanupPolicyResponse) ApplicationTypeVersionsCleanupPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationTypeVersionsCleanupPolicyResponse
		return ret
	}).(ApplicationTypeVersionsCleanupPolicyResponseOutput)
}

func (o ApplicationTypeVersionsCleanupPolicyResponsePtrOutput) MaxUnusedVersionsToKeep() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ApplicationTypeVersionsCleanupPolicyResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.MaxUnusedVersionsToKeep
	}).(pulumi.Float64PtrOutput)
}

type ApplicationUpgradePolicy struct {
	ApplicationHealthPolicy        *ArmApplicationHealthPolicy        `pulumi:"applicationHealthPolicy"`
	ForceRestart                   *bool                              `pulumi:"forceRestart"`
	RecreateApplication            *bool                              `pulumi:"recreateApplication"`
	RollingUpgradeMonitoringPolicy *ArmRollingUpgradeMonitoringPolicy `pulumi:"rollingUpgradeMonitoringPolicy"`
	UpgradeMode                    *string                            `pulumi:"upgradeMode"`
	UpgradeReplicaSetCheckTimeout  *string                            `pulumi:"upgradeReplicaSetCheckTimeout"`
}


func (val *ApplicationUpgradePolicy) Defaults() *ApplicationUpgradePolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ApplicationHealthPolicy = tmp.ApplicationHealthPolicy.Defaults()

	return &tmp
}





type ApplicationUpgradePolicyInput interface {
	pulumi.Input

	ToApplicationUpgradePolicyOutput() ApplicationUpgradePolicyOutput
	ToApplicationUpgradePolicyOutputWithContext(context.Context) ApplicationUpgradePolicyOutput
}

type ApplicationUpgradePolicyArgs struct {
	ApplicationHealthPolicy        ArmApplicationHealthPolicyPtrInput        `pulumi:"applicationHealthPolicy"`
	ForceRestart                   pulumi.BoolPtrInput                       `pulumi:"forceRestart"`
	RecreateApplication            pulumi.BoolPtrInput                       `pulumi:"recreateApplication"`
	RollingUpgradeMonitoringPolicy ArmRollingUpgradeMonitoringPolicyPtrInput `pulumi:"rollingUpgradeMonitoringPolicy"`
	UpgradeMode                    pulumi.StringPtrInput                     `pulumi:"upgradeMode"`
	UpgradeReplicaSetCheckTimeout  pulumi.StringPtrInput                     `pulumi:"upgradeReplicaSetCheckTimeout"`
}

func (ApplicationUpgradePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationUpgradePolicy)(nil)).Elem()
}

func (i ApplicationUpgradePolicyArgs) ToApplicationUpgradePolicyOutput() ApplicationUpgradePolicyOutput {
	return i.ToApplicationUpgradePolicyOutputWithContext(context.Background())
}

func (i ApplicationUpgradePolicyArgs) ToApplicationUpgradePolicyOutputWithContext(ctx context.Context) ApplicationUpgradePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUpgradePolicyOutput)
}

func (i ApplicationUpgradePolicyArgs) ToApplicationUpgradePolicyPtrOutput() ApplicationUpgradePolicyPtrOutput {
	return i.ToApplicationUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i ApplicationUpgradePolicyArgs) ToApplicationUpgradePolicyPtrOutputWithContext(ctx context.Context) ApplicationUpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUpgradePolicyOutput).ToApplicationUpgradePolicyPtrOutputWithContext(ctx)
}









type ApplicationUpgradePolicyPtrInput interface {
	pulumi.Input

	ToApplicationUpgradePolicyPtrOutput() ApplicationUpgradePolicyPtrOutput
	ToApplicationUpgradePolicyPtrOutputWithContext(context.Context) ApplicationUpgradePolicyPtrOutput
}

type applicationUpgradePolicyPtrType ApplicationUpgradePolicyArgs

func ApplicationUpgradePolicyPtr(v *ApplicationUpgradePolicyArgs) ApplicationUpgradePolicyPtrInput {
	return (*applicationUpgradePolicyPtrType)(v)
}

func (*applicationUpgradePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationUpgradePolicy)(nil)).Elem()
}

func (i *applicationUpgradePolicyPtrType) ToApplicationUpgradePolicyPtrOutput() ApplicationUpgradePolicyPtrOutput {
	return i.ToApplicationUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i *applicationUpgradePolicyPtrType) ToApplicationUpgradePolicyPtrOutputWithContext(ctx context.Context) ApplicationUpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUpgradePolicyPtrOutput)
}

type ApplicationUpgradePolicyOutput struct{ *pulumi.OutputState }

func (ApplicationUpgradePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationUpgradePolicy)(nil)).Elem()
}

func (o ApplicationUpgradePolicyOutput) ToApplicationUpgradePolicyOutput() ApplicationUpgradePolicyOutput {
	return o
}

func (o ApplicationUpgradePolicyOutput) ToApplicationUpgradePolicyOutputWithContext(ctx context.Context) ApplicationUpgradePolicyOutput {
	return o
}

func (o ApplicationUpgradePolicyOutput) ToApplicationUpgradePolicyPtrOutput() ApplicationUpgradePolicyPtrOutput {
	return o.ToApplicationUpgradePolicyPtrOutputWithContext(context.Background())
}

func (o ApplicationUpgradePolicyOutput) ToApplicationUpgradePolicyPtrOutputWithContext(ctx context.Context) ApplicationUpgradePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationUpgradePolicy) *ApplicationUpgradePolicy {
		return &v
	}).(ApplicationUpgradePolicyPtrOutput)
}

func (o ApplicationUpgradePolicyOutput) ApplicationHealthPolicy() ArmApplicationHealthPolicyPtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicy) *ArmApplicationHealthPolicy { return v.ApplicationHealthPolicy }).(ArmApplicationHealthPolicyPtrOutput)
}

func (o ApplicationUpgradePolicyOutput) ForceRestart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicy) *bool { return v.ForceRestart }).(pulumi.BoolPtrOutput)
}

func (o ApplicationUpgradePolicyOutput) RecreateApplication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicy) *bool { return v.RecreateApplication }).(pulumi.BoolPtrOutput)
}

func (o ApplicationUpgradePolicyOutput) RollingUpgradeMonitoringPolicy() ArmRollingUpgradeMonitoringPolicyPtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicy) *ArmRollingUpgradeMonitoringPolicy {
		return v.RollingUpgradeMonitoringPolicy
	}).(ArmRollingUpgradeMonitoringPolicyPtrOutput)
}

func (o ApplicationUpgradePolicyOutput) UpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicy) *string { return v.UpgradeMode }).(pulumi.StringPtrOutput)
}

func (o ApplicationUpgradePolicyOutput) UpgradeReplicaSetCheckTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicy) *string { return v.UpgradeReplicaSetCheckTimeout }).(pulumi.StringPtrOutput)
}

type ApplicationUpgradePolicyPtrOutput struct{ *pulumi.OutputState }

func (ApplicationUpgradePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationUpgradePolicy)(nil)).Elem()
}

func (o ApplicationUpgradePolicyPtrOutput) ToApplicationUpgradePolicyPtrOutput() ApplicationUpgradePolicyPtrOutput {
	return o
}

func (o ApplicationUpgradePolicyPtrOutput) ToApplicationUpgradePolicyPtrOutputWithContext(ctx context.Context) ApplicationUpgradePolicyPtrOutput {
	return o
}

func (o ApplicationUpgradePolicyPtrOutput) Elem() ApplicationUpgradePolicyOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicy) ApplicationUpgradePolicy {
		if v != nil {
			return *v
		}
		var ret ApplicationUpgradePolicy
		return ret
	}).(ApplicationUpgradePolicyOutput)
}

func (o ApplicationUpgradePolicyPtrOutput) ApplicationHealthPolicy() ArmApplicationHealthPolicyPtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicy) *ArmApplicationHealthPolicy {
		if v == nil {
			return nil
		}
		return v.ApplicationHealthPolicy
	}).(ArmApplicationHealthPolicyPtrOutput)
}

func (o ApplicationUpgradePolicyPtrOutput) ForceRestart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicy) *bool {
		if v == nil {
			return nil
		}
		return v.ForceRestart
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationUpgradePolicyPtrOutput) RecreateApplication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicy) *bool {
		if v == nil {
			return nil
		}
		return v.RecreateApplication
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationUpgradePolicyPtrOutput) RollingUpgradeMonitoringPolicy() ArmRollingUpgradeMonitoringPolicyPtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicy) *ArmRollingUpgradeMonitoringPolicy {
		if v == nil {
			return nil
		}
		return v.RollingUpgradeMonitoringPolicy
	}).(ArmRollingUpgradeMonitoringPolicyPtrOutput)
}

func (o ApplicationUpgradePolicyPtrOutput) UpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicy) *string {
		if v == nil {
			return nil
		}
		return v.UpgradeMode
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationUpgradePolicyPtrOutput) UpgradeReplicaSetCheckTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicy) *string {
		if v == nil {
			return nil
		}
		return v.UpgradeReplicaSetCheckTimeout
	}).(pulumi.StringPtrOutput)
}

type ApplicationUpgradePolicyResponse struct {
	ApplicationHealthPolicy        *ArmApplicationHealthPolicyResponse        `pulumi:"applicationHealthPolicy"`
	ForceRestart                   *bool                                      `pulumi:"forceRestart"`
	RecreateApplication            *bool                                      `pulumi:"recreateApplication"`
	RollingUpgradeMonitoringPolicy *ArmRollingUpgradeMonitoringPolicyResponse `pulumi:"rollingUpgradeMonitoringPolicy"`
	UpgradeMode                    *string                                    `pulumi:"upgradeMode"`
	UpgradeReplicaSetCheckTimeout  *string                                    `pulumi:"upgradeReplicaSetCheckTimeout"`
}


func (val *ApplicationUpgradePolicyResponse) Defaults() *ApplicationUpgradePolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ApplicationHealthPolicy = tmp.ApplicationHealthPolicy.Defaults()

	return &tmp
}





type ApplicationUpgradePolicyResponseInput interface {
	pulumi.Input

	ToApplicationUpgradePolicyResponseOutput() ApplicationUpgradePolicyResponseOutput
	ToApplicationUpgradePolicyResponseOutputWithContext(context.Context) ApplicationUpgradePolicyResponseOutput
}

type ApplicationUpgradePolicyResponseArgs struct {
	ApplicationHealthPolicy        ArmApplicationHealthPolicyResponsePtrInput        `pulumi:"applicationHealthPolicy"`
	ForceRestart                   pulumi.BoolPtrInput                               `pulumi:"forceRestart"`
	RecreateApplication            pulumi.BoolPtrInput                               `pulumi:"recreateApplication"`
	RollingUpgradeMonitoringPolicy ArmRollingUpgradeMonitoringPolicyResponsePtrInput `pulumi:"rollingUpgradeMonitoringPolicy"`
	UpgradeMode                    pulumi.StringPtrInput                             `pulumi:"upgradeMode"`
	UpgradeReplicaSetCheckTimeout  pulumi.StringPtrInput                             `pulumi:"upgradeReplicaSetCheckTimeout"`
}

func (ApplicationUpgradePolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationUpgradePolicyResponse)(nil)).Elem()
}

func (i ApplicationUpgradePolicyResponseArgs) ToApplicationUpgradePolicyResponseOutput() ApplicationUpgradePolicyResponseOutput {
	return i.ToApplicationUpgradePolicyResponseOutputWithContext(context.Background())
}

func (i ApplicationUpgradePolicyResponseArgs) ToApplicationUpgradePolicyResponseOutputWithContext(ctx context.Context) ApplicationUpgradePolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUpgradePolicyResponseOutput)
}

func (i ApplicationUpgradePolicyResponseArgs) ToApplicationUpgradePolicyResponsePtrOutput() ApplicationUpgradePolicyResponsePtrOutput {
	return i.ToApplicationUpgradePolicyResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationUpgradePolicyResponseArgs) ToApplicationUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationUpgradePolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUpgradePolicyResponseOutput).ToApplicationUpgradePolicyResponsePtrOutputWithContext(ctx)
}









type ApplicationUpgradePolicyResponsePtrInput interface {
	pulumi.Input

	ToApplicationUpgradePolicyResponsePtrOutput() ApplicationUpgradePolicyResponsePtrOutput
	ToApplicationUpgradePolicyResponsePtrOutputWithContext(context.Context) ApplicationUpgradePolicyResponsePtrOutput
}

type applicationUpgradePolicyResponsePtrType ApplicationUpgradePolicyResponseArgs

func ApplicationUpgradePolicyResponsePtr(v *ApplicationUpgradePolicyResponseArgs) ApplicationUpgradePolicyResponsePtrInput {
	return (*applicationUpgradePolicyResponsePtrType)(v)
}

func (*applicationUpgradePolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationUpgradePolicyResponse)(nil)).Elem()
}

func (i *applicationUpgradePolicyResponsePtrType) ToApplicationUpgradePolicyResponsePtrOutput() ApplicationUpgradePolicyResponsePtrOutput {
	return i.ToApplicationUpgradePolicyResponsePtrOutputWithContext(context.Background())
}

func (i *applicationUpgradePolicyResponsePtrType) ToApplicationUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationUpgradePolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUpgradePolicyResponsePtrOutput)
}

type ApplicationUpgradePolicyResponseOutput struct{ *pulumi.OutputState }

func (ApplicationUpgradePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationUpgradePolicyResponse)(nil)).Elem()
}

func (o ApplicationUpgradePolicyResponseOutput) ToApplicationUpgradePolicyResponseOutput() ApplicationUpgradePolicyResponseOutput {
	return o
}

func (o ApplicationUpgradePolicyResponseOutput) ToApplicationUpgradePolicyResponseOutputWithContext(ctx context.Context) ApplicationUpgradePolicyResponseOutput {
	return o
}

func (o ApplicationUpgradePolicyResponseOutput) ToApplicationUpgradePolicyResponsePtrOutput() ApplicationUpgradePolicyResponsePtrOutput {
	return o.ToApplicationUpgradePolicyResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationUpgradePolicyResponseOutput) ToApplicationUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationUpgradePolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationUpgradePolicyResponse) *ApplicationUpgradePolicyResponse {
		return &v
	}).(ApplicationUpgradePolicyResponsePtrOutput)
}

func (o ApplicationUpgradePolicyResponseOutput) ApplicationHealthPolicy() ArmApplicationHealthPolicyResponsePtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicyResponse) *ArmApplicationHealthPolicyResponse {
		return v.ApplicationHealthPolicy
	}).(ArmApplicationHealthPolicyResponsePtrOutput)
}

func (o ApplicationUpgradePolicyResponseOutput) ForceRestart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicyResponse) *bool { return v.ForceRestart }).(pulumi.BoolPtrOutput)
}

func (o ApplicationUpgradePolicyResponseOutput) RecreateApplication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicyResponse) *bool { return v.RecreateApplication }).(pulumi.BoolPtrOutput)
}

func (o ApplicationUpgradePolicyResponseOutput) RollingUpgradeMonitoringPolicy() ArmRollingUpgradeMonitoringPolicyResponsePtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicyResponse) *ArmRollingUpgradeMonitoringPolicyResponse {
		return v.RollingUpgradeMonitoringPolicy
	}).(ArmRollingUpgradeMonitoringPolicyResponsePtrOutput)
}

func (o ApplicationUpgradePolicyResponseOutput) UpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicyResponse) *string { return v.UpgradeMode }).(pulumi.StringPtrOutput)
}

func (o ApplicationUpgradePolicyResponseOutput) UpgradeReplicaSetCheckTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicyResponse) *string { return v.UpgradeReplicaSetCheckTimeout }).(pulumi.StringPtrOutput)
}

type ApplicationUpgradePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationUpgradePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationUpgradePolicyResponse)(nil)).Elem()
}

func (o ApplicationUpgradePolicyResponsePtrOutput) ToApplicationUpgradePolicyResponsePtrOutput() ApplicationUpgradePolicyResponsePtrOutput {
	return o
}

func (o ApplicationUpgradePolicyResponsePtrOutput) ToApplicationUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationUpgradePolicyResponsePtrOutput {
	return o
}

func (o ApplicationUpgradePolicyResponsePtrOutput) Elem() ApplicationUpgradePolicyResponseOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicyResponse) ApplicationUpgradePolicyResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationUpgradePolicyResponse
		return ret
	}).(ApplicationUpgradePolicyResponseOutput)
}

func (o ApplicationUpgradePolicyResponsePtrOutput) ApplicationHealthPolicy() ArmApplicationHealthPolicyResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicyResponse) *ArmApplicationHealthPolicyResponse {
		if v == nil {
			return nil
		}
		return v.ApplicationHealthPolicy
	}).(ArmApplicationHealthPolicyResponsePtrOutput)
}

func (o ApplicationUpgradePolicyResponsePtrOutput) ForceRestart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ForceRestart
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationUpgradePolicyResponsePtrOutput) RecreateApplication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RecreateApplication
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationUpgradePolicyResponsePtrOutput) RollingUpgradeMonitoringPolicy() ArmRollingUpgradeMonitoringPolicyResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicyResponse) *ArmRollingUpgradeMonitoringPolicyResponse {
		if v == nil {
			return nil
		}
		return v.RollingUpgradeMonitoringPolicy
	}).(ArmRollingUpgradeMonitoringPolicyResponsePtrOutput)
}

func (o ApplicationUpgradePolicyResponsePtrOutput) UpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.UpgradeMode
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationUpgradePolicyResponsePtrOutput) UpgradeReplicaSetCheckTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.UpgradeReplicaSetCheckTimeout
	}).(pulumi.StringPtrOutput)
}

type ApplicationUserAssignedIdentity struct {
	Name        string `pulumi:"name"`
	PrincipalId string `pulumi:"principalId"`
}





type ApplicationUserAssignedIdentityInput interface {
	pulumi.Input

	ToApplicationUserAssignedIdentityOutput() ApplicationUserAssignedIdentityOutput
	ToApplicationUserAssignedIdentityOutputWithContext(context.Context) ApplicationUserAssignedIdentityOutput
}

type ApplicationUserAssignedIdentityArgs struct {
	Name        pulumi.StringInput `pulumi:"name"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (ApplicationUserAssignedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationUserAssignedIdentity)(nil)).Elem()
}

func (i ApplicationUserAssignedIdentityArgs) ToApplicationUserAssignedIdentityOutput() ApplicationUserAssignedIdentityOutput {
	return i.ToApplicationUserAssignedIdentityOutputWithContext(context.Background())
}

func (i ApplicationUserAssignedIdentityArgs) ToApplicationUserAssignedIdentityOutputWithContext(ctx context.Context) ApplicationUserAssignedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUserAssignedIdentityOutput)
}





type ApplicationUserAssignedIdentityArrayInput interface {
	pulumi.Input

	ToApplicationUserAssignedIdentityArrayOutput() ApplicationUserAssignedIdentityArrayOutput
	ToApplicationUserAssignedIdentityArrayOutputWithContext(context.Context) ApplicationUserAssignedIdentityArrayOutput
}

type ApplicationUserAssignedIdentityArray []ApplicationUserAssignedIdentityInput

func (ApplicationUserAssignedIdentityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationUserAssignedIdentity)(nil)).Elem()
}

func (i ApplicationUserAssignedIdentityArray) ToApplicationUserAssignedIdentityArrayOutput() ApplicationUserAssignedIdentityArrayOutput {
	return i.ToApplicationUserAssignedIdentityArrayOutputWithContext(context.Background())
}

func (i ApplicationUserAssignedIdentityArray) ToApplicationUserAssignedIdentityArrayOutputWithContext(ctx context.Context) ApplicationUserAssignedIdentityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUserAssignedIdentityArrayOutput)
}

type ApplicationUserAssignedIdentityOutput struct{ *pulumi.OutputState }

func (ApplicationUserAssignedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationUserAssignedIdentity)(nil)).Elem()
}

func (o ApplicationUserAssignedIdentityOutput) ToApplicationUserAssignedIdentityOutput() ApplicationUserAssignedIdentityOutput {
	return o
}

func (o ApplicationUserAssignedIdentityOutput) ToApplicationUserAssignedIdentityOutputWithContext(ctx context.Context) ApplicationUserAssignedIdentityOutput {
	return o
}

func (o ApplicationUserAssignedIdentityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationUserAssignedIdentity) string { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationUserAssignedIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationUserAssignedIdentity) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type ApplicationUserAssignedIdentityArrayOutput struct{ *pulumi.OutputState }

func (ApplicationUserAssignedIdentityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationUserAssignedIdentity)(nil)).Elem()
}

func (o ApplicationUserAssignedIdentityArrayOutput) ToApplicationUserAssignedIdentityArrayOutput() ApplicationUserAssignedIdentityArrayOutput {
	return o
}

func (o ApplicationUserAssignedIdentityArrayOutput) ToApplicationUserAssignedIdentityArrayOutputWithContext(ctx context.Context) ApplicationUserAssignedIdentityArrayOutput {
	return o
}

func (o ApplicationUserAssignedIdentityArrayOutput) Index(i pulumi.IntInput) ApplicationUserAssignedIdentityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationUserAssignedIdentity {
		return vs[0].([]ApplicationUserAssignedIdentity)[vs[1].(int)]
	}).(ApplicationUserAssignedIdentityOutput)
}

type ApplicationUserAssignedIdentityResponse struct {
	Name        string `pulumi:"name"`
	PrincipalId string `pulumi:"principalId"`
}





type ApplicationUserAssignedIdentityResponseInput interface {
	pulumi.Input

	ToApplicationUserAssignedIdentityResponseOutput() ApplicationUserAssignedIdentityResponseOutput
	ToApplicationUserAssignedIdentityResponseOutputWithContext(context.Context) ApplicationUserAssignedIdentityResponseOutput
}

type ApplicationUserAssignedIdentityResponseArgs struct {
	Name        pulumi.StringInput `pulumi:"name"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (ApplicationUserAssignedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationUserAssignedIdentityResponse)(nil)).Elem()
}

func (i ApplicationUserAssignedIdentityResponseArgs) ToApplicationUserAssignedIdentityResponseOutput() ApplicationUserAssignedIdentityResponseOutput {
	return i.ToApplicationUserAssignedIdentityResponseOutputWithContext(context.Background())
}

func (i ApplicationUserAssignedIdentityResponseArgs) ToApplicationUserAssignedIdentityResponseOutputWithContext(ctx context.Context) ApplicationUserAssignedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUserAssignedIdentityResponseOutput)
}





type ApplicationUserAssignedIdentityResponseArrayInput interface {
	pulumi.Input

	ToApplicationUserAssignedIdentityResponseArrayOutput() ApplicationUserAssignedIdentityResponseArrayOutput
	ToApplicationUserAssignedIdentityResponseArrayOutputWithContext(context.Context) ApplicationUserAssignedIdentityResponseArrayOutput
}

type ApplicationUserAssignedIdentityResponseArray []ApplicationUserAssignedIdentityResponseInput

func (ApplicationUserAssignedIdentityResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationUserAssignedIdentityResponse)(nil)).Elem()
}

func (i ApplicationUserAssignedIdentityResponseArray) ToApplicationUserAssignedIdentityResponseArrayOutput() ApplicationUserAssignedIdentityResponseArrayOutput {
	return i.ToApplicationUserAssignedIdentityResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationUserAssignedIdentityResponseArray) ToApplicationUserAssignedIdentityResponseArrayOutputWithContext(ctx context.Context) ApplicationUserAssignedIdentityResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUserAssignedIdentityResponseArrayOutput)
}

type ApplicationUserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (ApplicationUserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationUserAssignedIdentityResponse)(nil)).Elem()
}

func (o ApplicationUserAssignedIdentityResponseOutput) ToApplicationUserAssignedIdentityResponseOutput() ApplicationUserAssignedIdentityResponseOutput {
	return o
}

func (o ApplicationUserAssignedIdentityResponseOutput) ToApplicationUserAssignedIdentityResponseOutputWithContext(ctx context.Context) ApplicationUserAssignedIdentityResponseOutput {
	return o
}

func (o ApplicationUserAssignedIdentityResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationUserAssignedIdentityResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationUserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationUserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type ApplicationUserAssignedIdentityResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationUserAssignedIdentityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationUserAssignedIdentityResponse)(nil)).Elem()
}

func (o ApplicationUserAssignedIdentityResponseArrayOutput) ToApplicationUserAssignedIdentityResponseArrayOutput() ApplicationUserAssignedIdentityResponseArrayOutput {
	return o
}

func (o ApplicationUserAssignedIdentityResponseArrayOutput) ToApplicationUserAssignedIdentityResponseArrayOutputWithContext(ctx context.Context) ApplicationUserAssignedIdentityResponseArrayOutput {
	return o
}

func (o ApplicationUserAssignedIdentityResponseArrayOutput) Index(i pulumi.IntInput) ApplicationUserAssignedIdentityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationUserAssignedIdentityResponse {
		return vs[0].([]ApplicationUserAssignedIdentityResponse)[vs[1].(int)]
	}).(ApplicationUserAssignedIdentityResponseOutput)
}

type ArmApplicationHealthPolicy struct {
	ConsiderWarningAsError                  *bool                                 `pulumi:"considerWarningAsError"`
	DefaultServiceTypeHealthPolicy          *ArmServiceTypeHealthPolicy           `pulumi:"defaultServiceTypeHealthPolicy"`
	MaxPercentUnhealthyDeployedApplications *int                                  `pulumi:"maxPercentUnhealthyDeployedApplications"`
	ServiceTypeHealthPolicyMap              map[string]ArmServiceTypeHealthPolicy `pulumi:"serviceTypeHealthPolicyMap"`
}


func (val *ArmApplicationHealthPolicy) Defaults() *ArmApplicationHealthPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ConsiderWarningAsError) {
		considerWarningAsError_ := false
		tmp.ConsiderWarningAsError = &considerWarningAsError_
	}
	tmp.DefaultServiceTypeHealthPolicy = tmp.DefaultServiceTypeHealthPolicy.Defaults()

	if isZero(tmp.MaxPercentUnhealthyDeployedApplications) {
		maxPercentUnhealthyDeployedApplications_ := 0
		tmp.MaxPercentUnhealthyDeployedApplications = &maxPercentUnhealthyDeployedApplications_
	}
	return &tmp
}





type ArmApplicationHealthPolicyInput interface {
	pulumi.Input

	ToArmApplicationHealthPolicyOutput() ArmApplicationHealthPolicyOutput
	ToArmApplicationHealthPolicyOutputWithContext(context.Context) ArmApplicationHealthPolicyOutput
}

type ArmApplicationHealthPolicyArgs struct {
	ConsiderWarningAsError                  pulumi.BoolPtrInput                `pulumi:"considerWarningAsError"`
	DefaultServiceTypeHealthPolicy          ArmServiceTypeHealthPolicyPtrInput `pulumi:"defaultServiceTypeHealthPolicy"`
	MaxPercentUnhealthyDeployedApplications pulumi.IntPtrInput                 `pulumi:"maxPercentUnhealthyDeployedApplications"`
	ServiceTypeHealthPolicyMap              ArmServiceTypeHealthPolicyMapInput `pulumi:"serviceTypeHealthPolicyMap"`
}

func (ArmApplicationHealthPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmApplicationHealthPolicy)(nil)).Elem()
}

func (i ArmApplicationHealthPolicyArgs) ToArmApplicationHealthPolicyOutput() ArmApplicationHealthPolicyOutput {
	return i.ToArmApplicationHealthPolicyOutputWithContext(context.Background())
}

func (i ArmApplicationHealthPolicyArgs) ToArmApplicationHealthPolicyOutputWithContext(ctx context.Context) ArmApplicationHealthPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmApplicationHealthPolicyOutput)
}

func (i ArmApplicationHealthPolicyArgs) ToArmApplicationHealthPolicyPtrOutput() ArmApplicationHealthPolicyPtrOutput {
	return i.ToArmApplicationHealthPolicyPtrOutputWithContext(context.Background())
}

func (i ArmApplicationHealthPolicyArgs) ToArmApplicationHealthPolicyPtrOutputWithContext(ctx context.Context) ArmApplicationHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmApplicationHealthPolicyOutput).ToArmApplicationHealthPolicyPtrOutputWithContext(ctx)
}









type ArmApplicationHealthPolicyPtrInput interface {
	pulumi.Input

	ToArmApplicationHealthPolicyPtrOutput() ArmApplicationHealthPolicyPtrOutput
	ToArmApplicationHealthPolicyPtrOutputWithContext(context.Context) ArmApplicationHealthPolicyPtrOutput
}

type armApplicationHealthPolicyPtrType ArmApplicationHealthPolicyArgs

func ArmApplicationHealthPolicyPtr(v *ArmApplicationHealthPolicyArgs) ArmApplicationHealthPolicyPtrInput {
	return (*armApplicationHealthPolicyPtrType)(v)
}

func (*armApplicationHealthPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmApplicationHealthPolicy)(nil)).Elem()
}

func (i *armApplicationHealthPolicyPtrType) ToArmApplicationHealthPolicyPtrOutput() ArmApplicationHealthPolicyPtrOutput {
	return i.ToArmApplicationHealthPolicyPtrOutputWithContext(context.Background())
}

func (i *armApplicationHealthPolicyPtrType) ToArmApplicationHealthPolicyPtrOutputWithContext(ctx context.Context) ArmApplicationHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmApplicationHealthPolicyPtrOutput)
}

type ArmApplicationHealthPolicyOutput struct{ *pulumi.OutputState }

func (ArmApplicationHealthPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmApplicationHealthPolicy)(nil)).Elem()
}

func (o ArmApplicationHealthPolicyOutput) ToArmApplicationHealthPolicyOutput() ArmApplicationHealthPolicyOutput {
	return o
}

func (o ArmApplicationHealthPolicyOutput) ToArmApplicationHealthPolicyOutputWithContext(ctx context.Context) ArmApplicationHealthPolicyOutput {
	return o
}

func (o ArmApplicationHealthPolicyOutput) ToArmApplicationHealthPolicyPtrOutput() ArmApplicationHealthPolicyPtrOutput {
	return o.ToArmApplicationHealthPolicyPtrOutputWithContext(context.Background())
}

func (o ArmApplicationHealthPolicyOutput) ToArmApplicationHealthPolicyPtrOutputWithContext(ctx context.Context) ArmApplicationHealthPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ArmApplicationHealthPolicy) *ArmApplicationHealthPolicy {
		return &v
	}).(ArmApplicationHealthPolicyPtrOutput)
}

func (o ArmApplicationHealthPolicyOutput) ConsiderWarningAsError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ArmApplicationHealthPolicy) *bool { return v.ConsiderWarningAsError }).(pulumi.BoolPtrOutput)
}

func (o ArmApplicationHealthPolicyOutput) DefaultServiceTypeHealthPolicy() ArmServiceTypeHealthPolicyPtrOutput {
	return o.ApplyT(func(v ArmApplicationHealthPolicy) *ArmServiceTypeHealthPolicy {
		return v.DefaultServiceTypeHealthPolicy
	}).(ArmServiceTypeHealthPolicyPtrOutput)
}

func (o ArmApplicationHealthPolicyOutput) MaxPercentUnhealthyDeployedApplications() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ArmApplicationHealthPolicy) *int { return v.MaxPercentUnhealthyDeployedApplications }).(pulumi.IntPtrOutput)
}

func (o ArmApplicationHealthPolicyOutput) ServiceTypeHealthPolicyMap() ArmServiceTypeHealthPolicyMapOutput {
	return o.ApplyT(func(v ArmApplicationHealthPolicy) map[string]ArmServiceTypeHealthPolicy {
		return v.ServiceTypeHealthPolicyMap
	}).(ArmServiceTypeHealthPolicyMapOutput)
}

type ArmApplicationHealthPolicyPtrOutput struct{ *pulumi.OutputState }

func (ArmApplicationHealthPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmApplicationHealthPolicy)(nil)).Elem()
}

func (o ArmApplicationHealthPolicyPtrOutput) ToArmApplicationHealthPolicyPtrOutput() ArmApplicationHealthPolicyPtrOutput {
	return o
}

func (o ArmApplicationHealthPolicyPtrOutput) ToArmApplicationHealthPolicyPtrOutputWithContext(ctx context.Context) ArmApplicationHealthPolicyPtrOutput {
	return o
}

func (o ArmApplicationHealthPolicyPtrOutput) Elem() ArmApplicationHealthPolicyOutput {
	return o.ApplyT(func(v *ArmApplicationHealthPolicy) ArmApplicationHealthPolicy {
		if v != nil {
			return *v
		}
		var ret ArmApplicationHealthPolicy
		return ret
	}).(ArmApplicationHealthPolicyOutput)
}

func (o ArmApplicationHealthPolicyPtrOutput) ConsiderWarningAsError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ArmApplicationHealthPolicy) *bool {
		if v == nil {
			return nil
		}
		return v.ConsiderWarningAsError
	}).(pulumi.BoolPtrOutput)
}

func (o ArmApplicationHealthPolicyPtrOutput) DefaultServiceTypeHealthPolicy() ArmServiceTypeHealthPolicyPtrOutput {
	return o.ApplyT(func(v *ArmApplicationHealthPolicy) *ArmServiceTypeHealthPolicy {
		if v == nil {
			return nil
		}
		return v.DefaultServiceTypeHealthPolicy
	}).(ArmServiceTypeHealthPolicyPtrOutput)
}

func (o ArmApplicationHealthPolicyPtrOutput) MaxPercentUnhealthyDeployedApplications() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ArmApplicationHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentUnhealthyDeployedApplications
	}).(pulumi.IntPtrOutput)
}

func (o ArmApplicationHealthPolicyPtrOutput) ServiceTypeHealthPolicyMap() ArmServiceTypeHealthPolicyMapOutput {
	return o.ApplyT(func(v *ArmApplicationHealthPolicy) map[string]ArmServiceTypeHealthPolicy {
		if v == nil {
			return nil
		}
		return v.ServiceTypeHealthPolicyMap
	}).(ArmServiceTypeHealthPolicyMapOutput)
}

type ArmApplicationHealthPolicyResponse struct {
	ConsiderWarningAsError                  *bool                                         `pulumi:"considerWarningAsError"`
	DefaultServiceTypeHealthPolicy          *ArmServiceTypeHealthPolicyResponse           `pulumi:"defaultServiceTypeHealthPolicy"`
	MaxPercentUnhealthyDeployedApplications *int                                          `pulumi:"maxPercentUnhealthyDeployedApplications"`
	ServiceTypeHealthPolicyMap              map[string]ArmServiceTypeHealthPolicyResponse `pulumi:"serviceTypeHealthPolicyMap"`
}


func (val *ArmApplicationHealthPolicyResponse) Defaults() *ArmApplicationHealthPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ConsiderWarningAsError) {
		considerWarningAsError_ := false
		tmp.ConsiderWarningAsError = &considerWarningAsError_
	}
	tmp.DefaultServiceTypeHealthPolicy = tmp.DefaultServiceTypeHealthPolicy.Defaults()

	if isZero(tmp.MaxPercentUnhealthyDeployedApplications) {
		maxPercentUnhealthyDeployedApplications_ := 0
		tmp.MaxPercentUnhealthyDeployedApplications = &maxPercentUnhealthyDeployedApplications_
	}
	return &tmp
}





type ArmApplicationHealthPolicyResponseInput interface {
	pulumi.Input

	ToArmApplicationHealthPolicyResponseOutput() ArmApplicationHealthPolicyResponseOutput
	ToArmApplicationHealthPolicyResponseOutputWithContext(context.Context) ArmApplicationHealthPolicyResponseOutput
}

type ArmApplicationHealthPolicyResponseArgs struct {
	ConsiderWarningAsError                  pulumi.BoolPtrInput                        `pulumi:"considerWarningAsError"`
	DefaultServiceTypeHealthPolicy          ArmServiceTypeHealthPolicyResponsePtrInput `pulumi:"defaultServiceTypeHealthPolicy"`
	MaxPercentUnhealthyDeployedApplications pulumi.IntPtrInput                         `pulumi:"maxPercentUnhealthyDeployedApplications"`
	ServiceTypeHealthPolicyMap              ArmServiceTypeHealthPolicyResponseMapInput `pulumi:"serviceTypeHealthPolicyMap"`
}

func (ArmApplicationHealthPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmApplicationHealthPolicyResponse)(nil)).Elem()
}

func (i ArmApplicationHealthPolicyResponseArgs) ToArmApplicationHealthPolicyResponseOutput() ArmApplicationHealthPolicyResponseOutput {
	return i.ToArmApplicationHealthPolicyResponseOutputWithContext(context.Background())
}

func (i ArmApplicationHealthPolicyResponseArgs) ToArmApplicationHealthPolicyResponseOutputWithContext(ctx context.Context) ArmApplicationHealthPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmApplicationHealthPolicyResponseOutput)
}

func (i ArmApplicationHealthPolicyResponseArgs) ToArmApplicationHealthPolicyResponsePtrOutput() ArmApplicationHealthPolicyResponsePtrOutput {
	return i.ToArmApplicationHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (i ArmApplicationHealthPolicyResponseArgs) ToArmApplicationHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ArmApplicationHealthPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmApplicationHealthPolicyResponseOutput).ToArmApplicationHealthPolicyResponsePtrOutputWithContext(ctx)
}









type ArmApplicationHealthPolicyResponsePtrInput interface {
	pulumi.Input

	ToArmApplicationHealthPolicyResponsePtrOutput() ArmApplicationHealthPolicyResponsePtrOutput
	ToArmApplicationHealthPolicyResponsePtrOutputWithContext(context.Context) ArmApplicationHealthPolicyResponsePtrOutput
}

type armApplicationHealthPolicyResponsePtrType ArmApplicationHealthPolicyResponseArgs

func ArmApplicationHealthPolicyResponsePtr(v *ArmApplicationHealthPolicyResponseArgs) ArmApplicationHealthPolicyResponsePtrInput {
	return (*armApplicationHealthPolicyResponsePtrType)(v)
}

func (*armApplicationHealthPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmApplicationHealthPolicyResponse)(nil)).Elem()
}

func (i *armApplicationHealthPolicyResponsePtrType) ToArmApplicationHealthPolicyResponsePtrOutput() ArmApplicationHealthPolicyResponsePtrOutput {
	return i.ToArmApplicationHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *armApplicationHealthPolicyResponsePtrType) ToArmApplicationHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ArmApplicationHealthPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmApplicationHealthPolicyResponsePtrOutput)
}

type ArmApplicationHealthPolicyResponseOutput struct{ *pulumi.OutputState }

func (ArmApplicationHealthPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmApplicationHealthPolicyResponse)(nil)).Elem()
}

func (o ArmApplicationHealthPolicyResponseOutput) ToArmApplicationHealthPolicyResponseOutput() ArmApplicationHealthPolicyResponseOutput {
	return o
}

func (o ArmApplicationHealthPolicyResponseOutput) ToArmApplicationHealthPolicyResponseOutputWithContext(ctx context.Context) ArmApplicationHealthPolicyResponseOutput {
	return o
}

func (o ArmApplicationHealthPolicyResponseOutput) ToArmApplicationHealthPolicyResponsePtrOutput() ArmApplicationHealthPolicyResponsePtrOutput {
	return o.ToArmApplicationHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (o ArmApplicationHealthPolicyResponseOutput) ToArmApplicationHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ArmApplicationHealthPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ArmApplicationHealthPolicyResponse) *ArmApplicationHealthPolicyResponse {
		return &v
	}).(ArmApplicationHealthPolicyResponsePtrOutput)
}

func (o ArmApplicationHealthPolicyResponseOutput) ConsiderWarningAsError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ArmApplicationHealthPolicyResponse) *bool { return v.ConsiderWarningAsError }).(pulumi.BoolPtrOutput)
}

func (o ArmApplicationHealthPolicyResponseOutput) DefaultServiceTypeHealthPolicy() ArmServiceTypeHealthPolicyResponsePtrOutput {
	return o.ApplyT(func(v ArmApplicationHealthPolicyResponse) *ArmServiceTypeHealthPolicyResponse {
		return v.DefaultServiceTypeHealthPolicy
	}).(ArmServiceTypeHealthPolicyResponsePtrOutput)
}

func (o ArmApplicationHealthPolicyResponseOutput) MaxPercentUnhealthyDeployedApplications() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ArmApplicationHealthPolicyResponse) *int { return v.MaxPercentUnhealthyDeployedApplications }).(pulumi.IntPtrOutput)
}

func (o ArmApplicationHealthPolicyResponseOutput) ServiceTypeHealthPolicyMap() ArmServiceTypeHealthPolicyResponseMapOutput {
	return o.ApplyT(func(v ArmApplicationHealthPolicyResponse) map[string]ArmServiceTypeHealthPolicyResponse {
		return v.ServiceTypeHealthPolicyMap
	}).(ArmServiceTypeHealthPolicyResponseMapOutput)
}

type ArmApplicationHealthPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ArmApplicationHealthPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmApplicationHealthPolicyResponse)(nil)).Elem()
}

func (o ArmApplicationHealthPolicyResponsePtrOutput) ToArmApplicationHealthPolicyResponsePtrOutput() ArmApplicationHealthPolicyResponsePtrOutput {
	return o
}

func (o ArmApplicationHealthPolicyResponsePtrOutput) ToArmApplicationHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ArmApplicationHealthPolicyResponsePtrOutput {
	return o
}

func (o ArmApplicationHealthPolicyResponsePtrOutput) Elem() ArmApplicationHealthPolicyResponseOutput {
	return o.ApplyT(func(v *ArmApplicationHealthPolicyResponse) ArmApplicationHealthPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ArmApplicationHealthPolicyResponse
		return ret
	}).(ArmApplicationHealthPolicyResponseOutput)
}

func (o ArmApplicationHealthPolicyResponsePtrOutput) ConsiderWarningAsError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ArmApplicationHealthPolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ConsiderWarningAsError
	}).(pulumi.BoolPtrOutput)
}

func (o ArmApplicationHealthPolicyResponsePtrOutput) DefaultServiceTypeHealthPolicy() ArmServiceTypeHealthPolicyResponsePtrOutput {
	return o.ApplyT(func(v *ArmApplicationHealthPolicyResponse) *ArmServiceTypeHealthPolicyResponse {
		if v == nil {
			return nil
		}
		return v.DefaultServiceTypeHealthPolicy
	}).(ArmServiceTypeHealthPolicyResponsePtrOutput)
}

func (o ArmApplicationHealthPolicyResponsePtrOutput) MaxPercentUnhealthyDeployedApplications() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ArmApplicationHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentUnhealthyDeployedApplications
	}).(pulumi.IntPtrOutput)
}

func (o ArmApplicationHealthPolicyResponsePtrOutput) ServiceTypeHealthPolicyMap() ArmServiceTypeHealthPolicyResponseMapOutput {
	return o.ApplyT(func(v *ArmApplicationHealthPolicyResponse) map[string]ArmServiceTypeHealthPolicyResponse {
		if v == nil {
			return nil
		}
		return v.ServiceTypeHealthPolicyMap
	}).(ArmServiceTypeHealthPolicyResponseMapOutput)
}

type ArmRollingUpgradeMonitoringPolicy struct {
	FailureAction             *string `pulumi:"failureAction"`
	HealthCheckRetryTimeout   *string `pulumi:"healthCheckRetryTimeout"`
	HealthCheckStableDuration *string `pulumi:"healthCheckStableDuration"`
	HealthCheckWaitDuration   *string `pulumi:"healthCheckWaitDuration"`
	UpgradeDomainTimeout      *string `pulumi:"upgradeDomainTimeout"`
	UpgradeTimeout            *string `pulumi:"upgradeTimeout"`
}





type ArmRollingUpgradeMonitoringPolicyInput interface {
	pulumi.Input

	ToArmRollingUpgradeMonitoringPolicyOutput() ArmRollingUpgradeMonitoringPolicyOutput
	ToArmRollingUpgradeMonitoringPolicyOutputWithContext(context.Context) ArmRollingUpgradeMonitoringPolicyOutput
}

type ArmRollingUpgradeMonitoringPolicyArgs struct {
	FailureAction             pulumi.StringPtrInput `pulumi:"failureAction"`
	HealthCheckRetryTimeout   pulumi.StringPtrInput `pulumi:"healthCheckRetryTimeout"`
	HealthCheckStableDuration pulumi.StringPtrInput `pulumi:"healthCheckStableDuration"`
	HealthCheckWaitDuration   pulumi.StringPtrInput `pulumi:"healthCheckWaitDuration"`
	UpgradeDomainTimeout      pulumi.StringPtrInput `pulumi:"upgradeDomainTimeout"`
	UpgradeTimeout            pulumi.StringPtrInput `pulumi:"upgradeTimeout"`
}

func (ArmRollingUpgradeMonitoringPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmRollingUpgradeMonitoringPolicy)(nil)).Elem()
}

func (i ArmRollingUpgradeMonitoringPolicyArgs) ToArmRollingUpgradeMonitoringPolicyOutput() ArmRollingUpgradeMonitoringPolicyOutput {
	return i.ToArmRollingUpgradeMonitoringPolicyOutputWithContext(context.Background())
}

func (i ArmRollingUpgradeMonitoringPolicyArgs) ToArmRollingUpgradeMonitoringPolicyOutputWithContext(ctx context.Context) ArmRollingUpgradeMonitoringPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmRollingUpgradeMonitoringPolicyOutput)
}

func (i ArmRollingUpgradeMonitoringPolicyArgs) ToArmRollingUpgradeMonitoringPolicyPtrOutput() ArmRollingUpgradeMonitoringPolicyPtrOutput {
	return i.ToArmRollingUpgradeMonitoringPolicyPtrOutputWithContext(context.Background())
}

func (i ArmRollingUpgradeMonitoringPolicyArgs) ToArmRollingUpgradeMonitoringPolicyPtrOutputWithContext(ctx context.Context) ArmRollingUpgradeMonitoringPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmRollingUpgradeMonitoringPolicyOutput).ToArmRollingUpgradeMonitoringPolicyPtrOutputWithContext(ctx)
}









type ArmRollingUpgradeMonitoringPolicyPtrInput interface {
	pulumi.Input

	ToArmRollingUpgradeMonitoringPolicyPtrOutput() ArmRollingUpgradeMonitoringPolicyPtrOutput
	ToArmRollingUpgradeMonitoringPolicyPtrOutputWithContext(context.Context) ArmRollingUpgradeMonitoringPolicyPtrOutput
}

type armRollingUpgradeMonitoringPolicyPtrType ArmRollingUpgradeMonitoringPolicyArgs

func ArmRollingUpgradeMonitoringPolicyPtr(v *ArmRollingUpgradeMonitoringPolicyArgs) ArmRollingUpgradeMonitoringPolicyPtrInput {
	return (*armRollingUpgradeMonitoringPolicyPtrType)(v)
}

func (*armRollingUpgradeMonitoringPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmRollingUpgradeMonitoringPolicy)(nil)).Elem()
}

func (i *armRollingUpgradeMonitoringPolicyPtrType) ToArmRollingUpgradeMonitoringPolicyPtrOutput() ArmRollingUpgradeMonitoringPolicyPtrOutput {
	return i.ToArmRollingUpgradeMonitoringPolicyPtrOutputWithContext(context.Background())
}

func (i *armRollingUpgradeMonitoringPolicyPtrType) ToArmRollingUpgradeMonitoringPolicyPtrOutputWithContext(ctx context.Context) ArmRollingUpgradeMonitoringPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmRollingUpgradeMonitoringPolicyPtrOutput)
}

type ArmRollingUpgradeMonitoringPolicyOutput struct{ *pulumi.OutputState }

func (ArmRollingUpgradeMonitoringPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmRollingUpgradeMonitoringPolicy)(nil)).Elem()
}

func (o ArmRollingUpgradeMonitoringPolicyOutput) ToArmRollingUpgradeMonitoringPolicyOutput() ArmRollingUpgradeMonitoringPolicyOutput {
	return o
}

func (o ArmRollingUpgradeMonitoringPolicyOutput) ToArmRollingUpgradeMonitoringPolicyOutputWithContext(ctx context.Context) ArmRollingUpgradeMonitoringPolicyOutput {
	return o
}

func (o ArmRollingUpgradeMonitoringPolicyOutput) ToArmRollingUpgradeMonitoringPolicyPtrOutput() ArmRollingUpgradeMonitoringPolicyPtrOutput {
	return o.ToArmRollingUpgradeMonitoringPolicyPtrOutputWithContext(context.Background())
}

func (o ArmRollingUpgradeMonitoringPolicyOutput) ToArmRollingUpgradeMonitoringPolicyPtrOutputWithContext(ctx context.Context) ArmRollingUpgradeMonitoringPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ArmRollingUpgradeMonitoringPolicy) *ArmRollingUpgradeMonitoringPolicy {
		return &v
	}).(ArmRollingUpgradeMonitoringPolicyPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyOutput) FailureAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmRollingUpgradeMonitoringPolicy) *string { return v.FailureAction }).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyOutput) HealthCheckRetryTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmRollingUpgradeMonitoringPolicy) *string { return v.HealthCheckRetryTimeout }).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyOutput) HealthCheckStableDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmRollingUpgradeMonitoringPolicy) *string { return v.HealthCheckStableDuration }).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyOutput) HealthCheckWaitDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmRollingUpgradeMonitoringPolicy) *string { return v.HealthCheckWaitDuration }).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyOutput) UpgradeDomainTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmRollingUpgradeMonitoringPolicy) *string { return v.UpgradeDomainTimeout }).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyOutput) UpgradeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmRollingUpgradeMonitoringPolicy) *string { return v.UpgradeTimeout }).(pulumi.StringPtrOutput)
}

type ArmRollingUpgradeMonitoringPolicyPtrOutput struct{ *pulumi.OutputState }

func (ArmRollingUpgradeMonitoringPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmRollingUpgradeMonitoringPolicy)(nil)).Elem()
}

func (o ArmRollingUpgradeMonitoringPolicyPtrOutput) ToArmRollingUpgradeMonitoringPolicyPtrOutput() ArmRollingUpgradeMonitoringPolicyPtrOutput {
	return o
}

func (o ArmRollingUpgradeMonitoringPolicyPtrOutput) ToArmRollingUpgradeMonitoringPolicyPtrOutputWithContext(ctx context.Context) ArmRollingUpgradeMonitoringPolicyPtrOutput {
	return o
}

func (o ArmRollingUpgradeMonitoringPolicyPtrOutput) Elem() ArmRollingUpgradeMonitoringPolicyOutput {
	return o.ApplyT(func(v *ArmRollingUpgradeMonitoringPolicy) ArmRollingUpgradeMonitoringPolicy {
		if v != nil {
			return *v
		}
		var ret ArmRollingUpgradeMonitoringPolicy
		return ret
	}).(ArmRollingUpgradeMonitoringPolicyOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyPtrOutput) FailureAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmRollingUpgradeMonitoringPolicy) *string {
		if v == nil {
			return nil
		}
		return v.FailureAction
	}).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyPtrOutput) HealthCheckRetryTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmRollingUpgradeMonitoringPolicy) *string {
		if v == nil {
			return nil
		}
		return v.HealthCheckRetryTimeout
	}).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyPtrOutput) HealthCheckStableDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmRollingUpgradeMonitoringPolicy) *string {
		if v == nil {
			return nil
		}
		return v.HealthCheckStableDuration
	}).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyPtrOutput) HealthCheckWaitDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmRollingUpgradeMonitoringPolicy) *string {
		if v == nil {
			return nil
		}
		return v.HealthCheckWaitDuration
	}).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyPtrOutput) UpgradeDomainTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmRollingUpgradeMonitoringPolicy) *string {
		if v == nil {
			return nil
		}
		return v.UpgradeDomainTimeout
	}).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyPtrOutput) UpgradeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmRollingUpgradeMonitoringPolicy) *string {
		if v == nil {
			return nil
		}
		return v.UpgradeTimeout
	}).(pulumi.StringPtrOutput)
}

type ArmRollingUpgradeMonitoringPolicyResponse struct {
	FailureAction             *string `pulumi:"failureAction"`
	HealthCheckRetryTimeout   *string `pulumi:"healthCheckRetryTimeout"`
	HealthCheckStableDuration *string `pulumi:"healthCheckStableDuration"`
	HealthCheckWaitDuration   *string `pulumi:"healthCheckWaitDuration"`
	UpgradeDomainTimeout      *string `pulumi:"upgradeDomainTimeout"`
	UpgradeTimeout            *string `pulumi:"upgradeTimeout"`
}





type ArmRollingUpgradeMonitoringPolicyResponseInput interface {
	pulumi.Input

	ToArmRollingUpgradeMonitoringPolicyResponseOutput() ArmRollingUpgradeMonitoringPolicyResponseOutput
	ToArmRollingUpgradeMonitoringPolicyResponseOutputWithContext(context.Context) ArmRollingUpgradeMonitoringPolicyResponseOutput
}

type ArmRollingUpgradeMonitoringPolicyResponseArgs struct {
	FailureAction             pulumi.StringPtrInput `pulumi:"failureAction"`
	HealthCheckRetryTimeout   pulumi.StringPtrInput `pulumi:"healthCheckRetryTimeout"`
	HealthCheckStableDuration pulumi.StringPtrInput `pulumi:"healthCheckStableDuration"`
	HealthCheckWaitDuration   pulumi.StringPtrInput `pulumi:"healthCheckWaitDuration"`
	UpgradeDomainTimeout      pulumi.StringPtrInput `pulumi:"upgradeDomainTimeout"`
	UpgradeTimeout            pulumi.StringPtrInput `pulumi:"upgradeTimeout"`
}

func (ArmRollingUpgradeMonitoringPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmRollingUpgradeMonitoringPolicyResponse)(nil)).Elem()
}

func (i ArmRollingUpgradeMonitoringPolicyResponseArgs) ToArmRollingUpgradeMonitoringPolicyResponseOutput() ArmRollingUpgradeMonitoringPolicyResponseOutput {
	return i.ToArmRollingUpgradeMonitoringPolicyResponseOutputWithContext(context.Background())
}

func (i ArmRollingUpgradeMonitoringPolicyResponseArgs) ToArmRollingUpgradeMonitoringPolicyResponseOutputWithContext(ctx context.Context) ArmRollingUpgradeMonitoringPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmRollingUpgradeMonitoringPolicyResponseOutput)
}

func (i ArmRollingUpgradeMonitoringPolicyResponseArgs) ToArmRollingUpgradeMonitoringPolicyResponsePtrOutput() ArmRollingUpgradeMonitoringPolicyResponsePtrOutput {
	return i.ToArmRollingUpgradeMonitoringPolicyResponsePtrOutputWithContext(context.Background())
}

func (i ArmRollingUpgradeMonitoringPolicyResponseArgs) ToArmRollingUpgradeMonitoringPolicyResponsePtrOutputWithContext(ctx context.Context) ArmRollingUpgradeMonitoringPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmRollingUpgradeMonitoringPolicyResponseOutput).ToArmRollingUpgradeMonitoringPolicyResponsePtrOutputWithContext(ctx)
}









type ArmRollingUpgradeMonitoringPolicyResponsePtrInput interface {
	pulumi.Input

	ToArmRollingUpgradeMonitoringPolicyResponsePtrOutput() ArmRollingUpgradeMonitoringPolicyResponsePtrOutput
	ToArmRollingUpgradeMonitoringPolicyResponsePtrOutputWithContext(context.Context) ArmRollingUpgradeMonitoringPolicyResponsePtrOutput
}

type armRollingUpgradeMonitoringPolicyResponsePtrType ArmRollingUpgradeMonitoringPolicyResponseArgs

func ArmRollingUpgradeMonitoringPolicyResponsePtr(v *ArmRollingUpgradeMonitoringPolicyResponseArgs) ArmRollingUpgradeMonitoringPolicyResponsePtrInput {
	return (*armRollingUpgradeMonitoringPolicyResponsePtrType)(v)
}

func (*armRollingUpgradeMonitoringPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmRollingUpgradeMonitoringPolicyResponse)(nil)).Elem()
}

func (i *armRollingUpgradeMonitoringPolicyResponsePtrType) ToArmRollingUpgradeMonitoringPolicyResponsePtrOutput() ArmRollingUpgradeMonitoringPolicyResponsePtrOutput {
	return i.ToArmRollingUpgradeMonitoringPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *armRollingUpgradeMonitoringPolicyResponsePtrType) ToArmRollingUpgradeMonitoringPolicyResponsePtrOutputWithContext(ctx context.Context) ArmRollingUpgradeMonitoringPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmRollingUpgradeMonitoringPolicyResponsePtrOutput)
}

type ArmRollingUpgradeMonitoringPolicyResponseOutput struct{ *pulumi.OutputState }

func (ArmRollingUpgradeMonitoringPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmRollingUpgradeMonitoringPolicyResponse)(nil)).Elem()
}

func (o ArmRollingUpgradeMonitoringPolicyResponseOutput) ToArmRollingUpgradeMonitoringPolicyResponseOutput() ArmRollingUpgradeMonitoringPolicyResponseOutput {
	return o
}

func (o ArmRollingUpgradeMonitoringPolicyResponseOutput) ToArmRollingUpgradeMonitoringPolicyResponseOutputWithContext(ctx context.Context) ArmRollingUpgradeMonitoringPolicyResponseOutput {
	return o
}

func (o ArmRollingUpgradeMonitoringPolicyResponseOutput) ToArmRollingUpgradeMonitoringPolicyResponsePtrOutput() ArmRollingUpgradeMonitoringPolicyResponsePtrOutput {
	return o.ToArmRollingUpgradeMonitoringPolicyResponsePtrOutputWithContext(context.Background())
}

func (o ArmRollingUpgradeMonitoringPolicyResponseOutput) ToArmRollingUpgradeMonitoringPolicyResponsePtrOutputWithContext(ctx context.Context) ArmRollingUpgradeMonitoringPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ArmRollingUpgradeMonitoringPolicyResponse) *ArmRollingUpgradeMonitoringPolicyResponse {
		return &v
	}).(ArmRollingUpgradeMonitoringPolicyResponsePtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyResponseOutput) FailureAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmRollingUpgradeMonitoringPolicyResponse) *string { return v.FailureAction }).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyResponseOutput) HealthCheckRetryTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmRollingUpgradeMonitoringPolicyResponse) *string { return v.HealthCheckRetryTimeout }).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyResponseOutput) HealthCheckStableDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmRollingUpgradeMonitoringPolicyResponse) *string { return v.HealthCheckStableDuration }).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyResponseOutput) HealthCheckWaitDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmRollingUpgradeMonitoringPolicyResponse) *string { return v.HealthCheckWaitDuration }).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyResponseOutput) UpgradeDomainTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmRollingUpgradeMonitoringPolicyResponse) *string { return v.UpgradeDomainTimeout }).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyResponseOutput) UpgradeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmRollingUpgradeMonitoringPolicyResponse) *string { return v.UpgradeTimeout }).(pulumi.StringPtrOutput)
}

type ArmRollingUpgradeMonitoringPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ArmRollingUpgradeMonitoringPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmRollingUpgradeMonitoringPolicyResponse)(nil)).Elem()
}

func (o ArmRollingUpgradeMonitoringPolicyResponsePtrOutput) ToArmRollingUpgradeMonitoringPolicyResponsePtrOutput() ArmRollingUpgradeMonitoringPolicyResponsePtrOutput {
	return o
}

func (o ArmRollingUpgradeMonitoringPolicyResponsePtrOutput) ToArmRollingUpgradeMonitoringPolicyResponsePtrOutputWithContext(ctx context.Context) ArmRollingUpgradeMonitoringPolicyResponsePtrOutput {
	return o
}

func (o ArmRollingUpgradeMonitoringPolicyResponsePtrOutput) Elem() ArmRollingUpgradeMonitoringPolicyResponseOutput {
	return o.ApplyT(func(v *ArmRollingUpgradeMonitoringPolicyResponse) ArmRollingUpgradeMonitoringPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ArmRollingUpgradeMonitoringPolicyResponse
		return ret
	}).(ArmRollingUpgradeMonitoringPolicyResponseOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyResponsePtrOutput) FailureAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmRollingUpgradeMonitoringPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.FailureAction
	}).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyResponsePtrOutput) HealthCheckRetryTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmRollingUpgradeMonitoringPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.HealthCheckRetryTimeout
	}).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyResponsePtrOutput) HealthCheckStableDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmRollingUpgradeMonitoringPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.HealthCheckStableDuration
	}).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyResponsePtrOutput) HealthCheckWaitDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmRollingUpgradeMonitoringPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.HealthCheckWaitDuration
	}).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyResponsePtrOutput) UpgradeDomainTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmRollingUpgradeMonitoringPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.UpgradeDomainTimeout
	}).(pulumi.StringPtrOutput)
}

func (o ArmRollingUpgradeMonitoringPolicyResponsePtrOutput) UpgradeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmRollingUpgradeMonitoringPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.UpgradeTimeout
	}).(pulumi.StringPtrOutput)
}

type ArmServiceTypeHealthPolicy struct {
	MaxPercentUnhealthyPartitionsPerService *int `pulumi:"maxPercentUnhealthyPartitionsPerService"`
	MaxPercentUnhealthyReplicasPerPartition *int `pulumi:"maxPercentUnhealthyReplicasPerPartition"`
	MaxPercentUnhealthyServices             *int `pulumi:"maxPercentUnhealthyServices"`
}


func (val *ArmServiceTypeHealthPolicy) Defaults() *ArmServiceTypeHealthPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxPercentUnhealthyPartitionsPerService) {
		maxPercentUnhealthyPartitionsPerService_ := 0
		tmp.MaxPercentUnhealthyPartitionsPerService = &maxPercentUnhealthyPartitionsPerService_
	}
	if isZero(tmp.MaxPercentUnhealthyReplicasPerPartition) {
		maxPercentUnhealthyReplicasPerPartition_ := 0
		tmp.MaxPercentUnhealthyReplicasPerPartition = &maxPercentUnhealthyReplicasPerPartition_
	}
	if isZero(tmp.MaxPercentUnhealthyServices) {
		maxPercentUnhealthyServices_ := 0
		tmp.MaxPercentUnhealthyServices = &maxPercentUnhealthyServices_
	}
	return &tmp
}





type ArmServiceTypeHealthPolicyInput interface {
	pulumi.Input

	ToArmServiceTypeHealthPolicyOutput() ArmServiceTypeHealthPolicyOutput
	ToArmServiceTypeHealthPolicyOutputWithContext(context.Context) ArmServiceTypeHealthPolicyOutput
}

type ArmServiceTypeHealthPolicyArgs struct {
	MaxPercentUnhealthyPartitionsPerService pulumi.IntPtrInput `pulumi:"maxPercentUnhealthyPartitionsPerService"`
	MaxPercentUnhealthyReplicasPerPartition pulumi.IntPtrInput `pulumi:"maxPercentUnhealthyReplicasPerPartition"`
	MaxPercentUnhealthyServices             pulumi.IntPtrInput `pulumi:"maxPercentUnhealthyServices"`
}

func (ArmServiceTypeHealthPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmServiceTypeHealthPolicy)(nil)).Elem()
}

func (i ArmServiceTypeHealthPolicyArgs) ToArmServiceTypeHealthPolicyOutput() ArmServiceTypeHealthPolicyOutput {
	return i.ToArmServiceTypeHealthPolicyOutputWithContext(context.Background())
}

func (i ArmServiceTypeHealthPolicyArgs) ToArmServiceTypeHealthPolicyOutputWithContext(ctx context.Context) ArmServiceTypeHealthPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmServiceTypeHealthPolicyOutput)
}

func (i ArmServiceTypeHealthPolicyArgs) ToArmServiceTypeHealthPolicyPtrOutput() ArmServiceTypeHealthPolicyPtrOutput {
	return i.ToArmServiceTypeHealthPolicyPtrOutputWithContext(context.Background())
}

func (i ArmServiceTypeHealthPolicyArgs) ToArmServiceTypeHealthPolicyPtrOutputWithContext(ctx context.Context) ArmServiceTypeHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmServiceTypeHealthPolicyOutput).ToArmServiceTypeHealthPolicyPtrOutputWithContext(ctx)
}









type ArmServiceTypeHealthPolicyPtrInput interface {
	pulumi.Input

	ToArmServiceTypeHealthPolicyPtrOutput() ArmServiceTypeHealthPolicyPtrOutput
	ToArmServiceTypeHealthPolicyPtrOutputWithContext(context.Context) ArmServiceTypeHealthPolicyPtrOutput
}

type armServiceTypeHealthPolicyPtrType ArmServiceTypeHealthPolicyArgs

func ArmServiceTypeHealthPolicyPtr(v *ArmServiceTypeHealthPolicyArgs) ArmServiceTypeHealthPolicyPtrInput {
	return (*armServiceTypeHealthPolicyPtrType)(v)
}

func (*armServiceTypeHealthPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmServiceTypeHealthPolicy)(nil)).Elem()
}

func (i *armServiceTypeHealthPolicyPtrType) ToArmServiceTypeHealthPolicyPtrOutput() ArmServiceTypeHealthPolicyPtrOutput {
	return i.ToArmServiceTypeHealthPolicyPtrOutputWithContext(context.Background())
}

func (i *armServiceTypeHealthPolicyPtrType) ToArmServiceTypeHealthPolicyPtrOutputWithContext(ctx context.Context) ArmServiceTypeHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmServiceTypeHealthPolicyPtrOutput)
}





type ArmServiceTypeHealthPolicyMapInput interface {
	pulumi.Input

	ToArmServiceTypeHealthPolicyMapOutput() ArmServiceTypeHealthPolicyMapOutput
	ToArmServiceTypeHealthPolicyMapOutputWithContext(context.Context) ArmServiceTypeHealthPolicyMapOutput
}

type ArmServiceTypeHealthPolicyMap map[string]ArmServiceTypeHealthPolicyInput

func (ArmServiceTypeHealthPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ArmServiceTypeHealthPolicy)(nil)).Elem()
}

func (i ArmServiceTypeHealthPolicyMap) ToArmServiceTypeHealthPolicyMapOutput() ArmServiceTypeHealthPolicyMapOutput {
	return i.ToArmServiceTypeHealthPolicyMapOutputWithContext(context.Background())
}

func (i ArmServiceTypeHealthPolicyMap) ToArmServiceTypeHealthPolicyMapOutputWithContext(ctx context.Context) ArmServiceTypeHealthPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmServiceTypeHealthPolicyMapOutput)
}

type ArmServiceTypeHealthPolicyOutput struct{ *pulumi.OutputState }

func (ArmServiceTypeHealthPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmServiceTypeHealthPolicy)(nil)).Elem()
}

func (o ArmServiceTypeHealthPolicyOutput) ToArmServiceTypeHealthPolicyOutput() ArmServiceTypeHealthPolicyOutput {
	return o
}

func (o ArmServiceTypeHealthPolicyOutput) ToArmServiceTypeHealthPolicyOutputWithContext(ctx context.Context) ArmServiceTypeHealthPolicyOutput {
	return o
}

func (o ArmServiceTypeHealthPolicyOutput) ToArmServiceTypeHealthPolicyPtrOutput() ArmServiceTypeHealthPolicyPtrOutput {
	return o.ToArmServiceTypeHealthPolicyPtrOutputWithContext(context.Background())
}

func (o ArmServiceTypeHealthPolicyOutput) ToArmServiceTypeHealthPolicyPtrOutputWithContext(ctx context.Context) ArmServiceTypeHealthPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ArmServiceTypeHealthPolicy) *ArmServiceTypeHealthPolicy {
		return &v
	}).(ArmServiceTypeHealthPolicyPtrOutput)
}

func (o ArmServiceTypeHealthPolicyOutput) MaxPercentUnhealthyPartitionsPerService() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ArmServiceTypeHealthPolicy) *int { return v.MaxPercentUnhealthyPartitionsPerService }).(pulumi.IntPtrOutput)
}

func (o ArmServiceTypeHealthPolicyOutput) MaxPercentUnhealthyReplicasPerPartition() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ArmServiceTypeHealthPolicy) *int { return v.MaxPercentUnhealthyReplicasPerPartition }).(pulumi.IntPtrOutput)
}

func (o ArmServiceTypeHealthPolicyOutput) MaxPercentUnhealthyServices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ArmServiceTypeHealthPolicy) *int { return v.MaxPercentUnhealthyServices }).(pulumi.IntPtrOutput)
}

type ArmServiceTypeHealthPolicyPtrOutput struct{ *pulumi.OutputState }

func (ArmServiceTypeHealthPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmServiceTypeHealthPolicy)(nil)).Elem()
}

func (o ArmServiceTypeHealthPolicyPtrOutput) ToArmServiceTypeHealthPolicyPtrOutput() ArmServiceTypeHealthPolicyPtrOutput {
	return o
}

func (o ArmServiceTypeHealthPolicyPtrOutput) ToArmServiceTypeHealthPolicyPtrOutputWithContext(ctx context.Context) ArmServiceTypeHealthPolicyPtrOutput {
	return o
}

func (o ArmServiceTypeHealthPolicyPtrOutput) Elem() ArmServiceTypeHealthPolicyOutput {
	return o.ApplyT(func(v *ArmServiceTypeHealthPolicy) ArmServiceTypeHealthPolicy {
		if v != nil {
			return *v
		}
		var ret ArmServiceTypeHealthPolicy
		return ret
	}).(ArmServiceTypeHealthPolicyOutput)
}

func (o ArmServiceTypeHealthPolicyPtrOutput) MaxPercentUnhealthyPartitionsPerService() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ArmServiceTypeHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentUnhealthyPartitionsPerService
	}).(pulumi.IntPtrOutput)
}

func (o ArmServiceTypeHealthPolicyPtrOutput) MaxPercentUnhealthyReplicasPerPartition() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ArmServiceTypeHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentUnhealthyReplicasPerPartition
	}).(pulumi.IntPtrOutput)
}

func (o ArmServiceTypeHealthPolicyPtrOutput) MaxPercentUnhealthyServices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ArmServiceTypeHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentUnhealthyServices
	}).(pulumi.IntPtrOutput)
}

type ArmServiceTypeHealthPolicyMapOutput struct{ *pulumi.OutputState }

func (ArmServiceTypeHealthPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ArmServiceTypeHealthPolicy)(nil)).Elem()
}

func (o ArmServiceTypeHealthPolicyMapOutput) ToArmServiceTypeHealthPolicyMapOutput() ArmServiceTypeHealthPolicyMapOutput {
	return o
}

func (o ArmServiceTypeHealthPolicyMapOutput) ToArmServiceTypeHealthPolicyMapOutputWithContext(ctx context.Context) ArmServiceTypeHealthPolicyMapOutput {
	return o
}

func (o ArmServiceTypeHealthPolicyMapOutput) MapIndex(k pulumi.StringInput) ArmServiceTypeHealthPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ArmServiceTypeHealthPolicy {
		return vs[0].(map[string]ArmServiceTypeHealthPolicy)[vs[1].(string)]
	}).(ArmServiceTypeHealthPolicyOutput)
}

type ArmServiceTypeHealthPolicyResponse struct {
	MaxPercentUnhealthyPartitionsPerService *int `pulumi:"maxPercentUnhealthyPartitionsPerService"`
	MaxPercentUnhealthyReplicasPerPartition *int `pulumi:"maxPercentUnhealthyReplicasPerPartition"`
	MaxPercentUnhealthyServices             *int `pulumi:"maxPercentUnhealthyServices"`
}


func (val *ArmServiceTypeHealthPolicyResponse) Defaults() *ArmServiceTypeHealthPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxPercentUnhealthyPartitionsPerService) {
		maxPercentUnhealthyPartitionsPerService_ := 0
		tmp.MaxPercentUnhealthyPartitionsPerService = &maxPercentUnhealthyPartitionsPerService_
	}
	if isZero(tmp.MaxPercentUnhealthyReplicasPerPartition) {
		maxPercentUnhealthyReplicasPerPartition_ := 0
		tmp.MaxPercentUnhealthyReplicasPerPartition = &maxPercentUnhealthyReplicasPerPartition_
	}
	if isZero(tmp.MaxPercentUnhealthyServices) {
		maxPercentUnhealthyServices_ := 0
		tmp.MaxPercentUnhealthyServices = &maxPercentUnhealthyServices_
	}
	return &tmp
}





type ArmServiceTypeHealthPolicyResponseInput interface {
	pulumi.Input

	ToArmServiceTypeHealthPolicyResponseOutput() ArmServiceTypeHealthPolicyResponseOutput
	ToArmServiceTypeHealthPolicyResponseOutputWithContext(context.Context) ArmServiceTypeHealthPolicyResponseOutput
}

type ArmServiceTypeHealthPolicyResponseArgs struct {
	MaxPercentUnhealthyPartitionsPerService pulumi.IntPtrInput `pulumi:"maxPercentUnhealthyPartitionsPerService"`
	MaxPercentUnhealthyReplicasPerPartition pulumi.IntPtrInput `pulumi:"maxPercentUnhealthyReplicasPerPartition"`
	MaxPercentUnhealthyServices             pulumi.IntPtrInput `pulumi:"maxPercentUnhealthyServices"`
}

func (ArmServiceTypeHealthPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmServiceTypeHealthPolicyResponse)(nil)).Elem()
}

func (i ArmServiceTypeHealthPolicyResponseArgs) ToArmServiceTypeHealthPolicyResponseOutput() ArmServiceTypeHealthPolicyResponseOutput {
	return i.ToArmServiceTypeHealthPolicyResponseOutputWithContext(context.Background())
}

func (i ArmServiceTypeHealthPolicyResponseArgs) ToArmServiceTypeHealthPolicyResponseOutputWithContext(ctx context.Context) ArmServiceTypeHealthPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmServiceTypeHealthPolicyResponseOutput)
}

func (i ArmServiceTypeHealthPolicyResponseArgs) ToArmServiceTypeHealthPolicyResponsePtrOutput() ArmServiceTypeHealthPolicyResponsePtrOutput {
	return i.ToArmServiceTypeHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (i ArmServiceTypeHealthPolicyResponseArgs) ToArmServiceTypeHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ArmServiceTypeHealthPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmServiceTypeHealthPolicyResponseOutput).ToArmServiceTypeHealthPolicyResponsePtrOutputWithContext(ctx)
}









type ArmServiceTypeHealthPolicyResponsePtrInput interface {
	pulumi.Input

	ToArmServiceTypeHealthPolicyResponsePtrOutput() ArmServiceTypeHealthPolicyResponsePtrOutput
	ToArmServiceTypeHealthPolicyResponsePtrOutputWithContext(context.Context) ArmServiceTypeHealthPolicyResponsePtrOutput
}

type armServiceTypeHealthPolicyResponsePtrType ArmServiceTypeHealthPolicyResponseArgs

func ArmServiceTypeHealthPolicyResponsePtr(v *ArmServiceTypeHealthPolicyResponseArgs) ArmServiceTypeHealthPolicyResponsePtrInput {
	return (*armServiceTypeHealthPolicyResponsePtrType)(v)
}

func (*armServiceTypeHealthPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmServiceTypeHealthPolicyResponse)(nil)).Elem()
}

func (i *armServiceTypeHealthPolicyResponsePtrType) ToArmServiceTypeHealthPolicyResponsePtrOutput() ArmServiceTypeHealthPolicyResponsePtrOutput {
	return i.ToArmServiceTypeHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *armServiceTypeHealthPolicyResponsePtrType) ToArmServiceTypeHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ArmServiceTypeHealthPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmServiceTypeHealthPolicyResponsePtrOutput)
}





type ArmServiceTypeHealthPolicyResponseMapInput interface {
	pulumi.Input

	ToArmServiceTypeHealthPolicyResponseMapOutput() ArmServiceTypeHealthPolicyResponseMapOutput
	ToArmServiceTypeHealthPolicyResponseMapOutputWithContext(context.Context) ArmServiceTypeHealthPolicyResponseMapOutput
}

type ArmServiceTypeHealthPolicyResponseMap map[string]ArmServiceTypeHealthPolicyResponseInput

func (ArmServiceTypeHealthPolicyResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ArmServiceTypeHealthPolicyResponse)(nil)).Elem()
}

func (i ArmServiceTypeHealthPolicyResponseMap) ToArmServiceTypeHealthPolicyResponseMapOutput() ArmServiceTypeHealthPolicyResponseMapOutput {
	return i.ToArmServiceTypeHealthPolicyResponseMapOutputWithContext(context.Background())
}

func (i ArmServiceTypeHealthPolicyResponseMap) ToArmServiceTypeHealthPolicyResponseMapOutputWithContext(ctx context.Context) ArmServiceTypeHealthPolicyResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmServiceTypeHealthPolicyResponseMapOutput)
}

type ArmServiceTypeHealthPolicyResponseOutput struct{ *pulumi.OutputState }

func (ArmServiceTypeHealthPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmServiceTypeHealthPolicyResponse)(nil)).Elem()
}

func (o ArmServiceTypeHealthPolicyResponseOutput) ToArmServiceTypeHealthPolicyResponseOutput() ArmServiceTypeHealthPolicyResponseOutput {
	return o
}

func (o ArmServiceTypeHealthPolicyResponseOutput) ToArmServiceTypeHealthPolicyResponseOutputWithContext(ctx context.Context) ArmServiceTypeHealthPolicyResponseOutput {
	return o
}

func (o ArmServiceTypeHealthPolicyResponseOutput) ToArmServiceTypeHealthPolicyResponsePtrOutput() ArmServiceTypeHealthPolicyResponsePtrOutput {
	return o.ToArmServiceTypeHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (o ArmServiceTypeHealthPolicyResponseOutput) ToArmServiceTypeHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ArmServiceTypeHealthPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ArmServiceTypeHealthPolicyResponse) *ArmServiceTypeHealthPolicyResponse {
		return &v
	}).(ArmServiceTypeHealthPolicyResponsePtrOutput)
}

func (o ArmServiceTypeHealthPolicyResponseOutput) MaxPercentUnhealthyPartitionsPerService() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ArmServiceTypeHealthPolicyResponse) *int { return v.MaxPercentUnhealthyPartitionsPerService }).(pulumi.IntPtrOutput)
}

func (o ArmServiceTypeHealthPolicyResponseOutput) MaxPercentUnhealthyReplicasPerPartition() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ArmServiceTypeHealthPolicyResponse) *int { return v.MaxPercentUnhealthyReplicasPerPartition }).(pulumi.IntPtrOutput)
}

func (o ArmServiceTypeHealthPolicyResponseOutput) MaxPercentUnhealthyServices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ArmServiceTypeHealthPolicyResponse) *int { return v.MaxPercentUnhealthyServices }).(pulumi.IntPtrOutput)
}

type ArmServiceTypeHealthPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ArmServiceTypeHealthPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmServiceTypeHealthPolicyResponse)(nil)).Elem()
}

func (o ArmServiceTypeHealthPolicyResponsePtrOutput) ToArmServiceTypeHealthPolicyResponsePtrOutput() ArmServiceTypeHealthPolicyResponsePtrOutput {
	return o
}

func (o ArmServiceTypeHealthPolicyResponsePtrOutput) ToArmServiceTypeHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ArmServiceTypeHealthPolicyResponsePtrOutput {
	return o
}

func (o ArmServiceTypeHealthPolicyResponsePtrOutput) Elem() ArmServiceTypeHealthPolicyResponseOutput {
	return o.ApplyT(func(v *ArmServiceTypeHealthPolicyResponse) ArmServiceTypeHealthPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ArmServiceTypeHealthPolicyResponse
		return ret
	}).(ArmServiceTypeHealthPolicyResponseOutput)
}

func (o ArmServiceTypeHealthPolicyResponsePtrOutput) MaxPercentUnhealthyPartitionsPerService() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ArmServiceTypeHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentUnhealthyPartitionsPerService
	}).(pulumi.IntPtrOutput)
}

func (o ArmServiceTypeHealthPolicyResponsePtrOutput) MaxPercentUnhealthyReplicasPerPartition() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ArmServiceTypeHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentUnhealthyReplicasPerPartition
	}).(pulumi.IntPtrOutput)
}

func (o ArmServiceTypeHealthPolicyResponsePtrOutput) MaxPercentUnhealthyServices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ArmServiceTypeHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentUnhealthyServices
	}).(pulumi.IntPtrOutput)
}

type ArmServiceTypeHealthPolicyResponseMapOutput struct{ *pulumi.OutputState }

func (ArmServiceTypeHealthPolicyResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ArmServiceTypeHealthPolicyResponse)(nil)).Elem()
}

func (o ArmServiceTypeHealthPolicyResponseMapOutput) ToArmServiceTypeHealthPolicyResponseMapOutput() ArmServiceTypeHealthPolicyResponseMapOutput {
	return o
}

func (o ArmServiceTypeHealthPolicyResponseMapOutput) ToArmServiceTypeHealthPolicyResponseMapOutputWithContext(ctx context.Context) ArmServiceTypeHealthPolicyResponseMapOutput {
	return o
}

func (o ArmServiceTypeHealthPolicyResponseMapOutput) MapIndex(k pulumi.StringInput) ArmServiceTypeHealthPolicyResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ArmServiceTypeHealthPolicyResponse {
		return vs[0].(map[string]ArmServiceTypeHealthPolicyResponse)[vs[1].(string)]
	}).(ArmServiceTypeHealthPolicyResponseOutput)
}

type AzureActiveDirectory struct {
	ClientApplication  *string `pulumi:"clientApplication"`
	ClusterApplication *string `pulumi:"clusterApplication"`
	TenantId           *string `pulumi:"tenantId"`
}





type AzureActiveDirectoryInput interface {
	pulumi.Input

	ToAzureActiveDirectoryOutput() AzureActiveDirectoryOutput
	ToAzureActiveDirectoryOutputWithContext(context.Context) AzureActiveDirectoryOutput
}

type AzureActiveDirectoryArgs struct {
	ClientApplication  pulumi.StringPtrInput `pulumi:"clientApplication"`
	ClusterApplication pulumi.StringPtrInput `pulumi:"clusterApplication"`
	TenantId           pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (AzureActiveDirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectory)(nil)).Elem()
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryOutput() AzureActiveDirectoryOutput {
	return i.ToAzureActiveDirectoryOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryOutputWithContext(ctx context.Context) AzureActiveDirectoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryOutput)
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return i.ToAzureActiveDirectoryPtrOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryOutput).ToAzureActiveDirectoryPtrOutputWithContext(ctx)
}









type AzureActiveDirectoryPtrInput interface {
	pulumi.Input

	ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput
	ToAzureActiveDirectoryPtrOutputWithContext(context.Context) AzureActiveDirectoryPtrOutput
}

type azureActiveDirectoryPtrType AzureActiveDirectoryArgs

func AzureActiveDirectoryPtr(v *AzureActiveDirectoryArgs) AzureActiveDirectoryPtrInput {
	return (*azureActiveDirectoryPtrType)(v)
}

func (*azureActiveDirectoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectory)(nil)).Elem()
}

func (i *azureActiveDirectoryPtrType) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return i.ToAzureActiveDirectoryPtrOutputWithContext(context.Background())
}

func (i *azureActiveDirectoryPtrType) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryPtrOutput)
}

type AzureActiveDirectoryOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectory)(nil)).Elem()
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryOutput() AzureActiveDirectoryOutput {
	return o
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryOutputWithContext(ctx context.Context) AzureActiveDirectoryOutput {
	return o
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return o.ToAzureActiveDirectoryPtrOutputWithContext(context.Background())
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureActiveDirectory) *AzureActiveDirectory {
		return &v
	}).(AzureActiveDirectoryPtrOutput)
}

func (o AzureActiveDirectoryOutput) ClientApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *string { return v.ClientApplication }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryOutput) ClusterApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *string { return v.ClusterApplication }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryPtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectory)(nil)).Elem()
}

func (o AzureActiveDirectoryPtrOutput) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return o
}

func (o AzureActiveDirectoryPtrOutput) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return o
}

func (o AzureActiveDirectoryPtrOutput) Elem() AzureActiveDirectoryOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) AzureActiveDirectory {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectory
		return ret
	}).(AzureActiveDirectoryOutput)
}

func (o AzureActiveDirectoryPtrOutput) ClientApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *string {
		if v == nil {
			return nil
		}
		return v.ClientApplication
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryPtrOutput) ClusterApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *string {
		if v == nil {
			return nil
		}
		return v.ClusterApplication
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryResponse struct {
	ClientApplication  *string `pulumi:"clientApplication"`
	ClusterApplication *string `pulumi:"clusterApplication"`
	TenantId           *string `pulumi:"tenantId"`
}





type AzureActiveDirectoryResponseInput interface {
	pulumi.Input

	ToAzureActiveDirectoryResponseOutput() AzureActiveDirectoryResponseOutput
	ToAzureActiveDirectoryResponseOutputWithContext(context.Context) AzureActiveDirectoryResponseOutput
}

type AzureActiveDirectoryResponseArgs struct {
	ClientApplication  pulumi.StringPtrInput `pulumi:"clientApplication"`
	ClusterApplication pulumi.StringPtrInput `pulumi:"clusterApplication"`
	TenantId           pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (AzureActiveDirectoryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryResponse)(nil)).Elem()
}

func (i AzureActiveDirectoryResponseArgs) ToAzureActiveDirectoryResponseOutput() AzureActiveDirectoryResponseOutput {
	return i.ToAzureActiveDirectoryResponseOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryResponseArgs) ToAzureActiveDirectoryResponseOutputWithContext(ctx context.Context) AzureActiveDirectoryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryResponseOutput)
}

func (i AzureActiveDirectoryResponseArgs) ToAzureActiveDirectoryResponsePtrOutput() AzureActiveDirectoryResponsePtrOutput {
	return i.ToAzureActiveDirectoryResponsePtrOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryResponseArgs) ToAzureActiveDirectoryResponsePtrOutputWithContext(ctx context.Context) AzureActiveDirectoryResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryResponseOutput).ToAzureActiveDirectoryResponsePtrOutputWithContext(ctx)
}









type AzureActiveDirectoryResponsePtrInput interface {
	pulumi.Input

	ToAzureActiveDirectoryResponsePtrOutput() AzureActiveDirectoryResponsePtrOutput
	ToAzureActiveDirectoryResponsePtrOutputWithContext(context.Context) AzureActiveDirectoryResponsePtrOutput
}

type azureActiveDirectoryResponsePtrType AzureActiveDirectoryResponseArgs

func AzureActiveDirectoryResponsePtr(v *AzureActiveDirectoryResponseArgs) AzureActiveDirectoryResponsePtrInput {
	return (*azureActiveDirectoryResponsePtrType)(v)
}

func (*azureActiveDirectoryResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryResponse)(nil)).Elem()
}

func (i *azureActiveDirectoryResponsePtrType) ToAzureActiveDirectoryResponsePtrOutput() AzureActiveDirectoryResponsePtrOutput {
	return i.ToAzureActiveDirectoryResponsePtrOutputWithContext(context.Background())
}

func (i *azureActiveDirectoryResponsePtrType) ToAzureActiveDirectoryResponsePtrOutputWithContext(ctx context.Context) AzureActiveDirectoryResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryResponsePtrOutput)
}

type AzureActiveDirectoryResponseOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryResponseOutput) ToAzureActiveDirectoryResponseOutput() AzureActiveDirectoryResponseOutput {
	return o
}

func (o AzureActiveDirectoryResponseOutput) ToAzureActiveDirectoryResponseOutputWithContext(ctx context.Context) AzureActiveDirectoryResponseOutput {
	return o
}

func (o AzureActiveDirectoryResponseOutput) ToAzureActiveDirectoryResponsePtrOutput() AzureActiveDirectoryResponsePtrOutput {
	return o.ToAzureActiveDirectoryResponsePtrOutputWithContext(context.Background())
}

func (o AzureActiveDirectoryResponseOutput) ToAzureActiveDirectoryResponsePtrOutputWithContext(ctx context.Context) AzureActiveDirectoryResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureActiveDirectoryResponse) *AzureActiveDirectoryResponse {
		return &v
	}).(AzureActiveDirectoryResponsePtrOutput)
}

func (o AzureActiveDirectoryResponseOutput) ClientApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *string { return v.ClientApplication }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryResponseOutput) ClusterApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *string { return v.ClusterApplication }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryResponsePtrOutput) ToAzureActiveDirectoryResponsePtrOutput() AzureActiveDirectoryResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryResponsePtrOutput) ToAzureActiveDirectoryResponsePtrOutputWithContext(ctx context.Context) AzureActiveDirectoryResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryResponsePtrOutput) Elem() AzureActiveDirectoryResponseOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) AzureActiveDirectoryResponse {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectoryResponse
		return ret
	}).(AzureActiveDirectoryResponseOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) ClientApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientApplication
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) ClusterApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClusterApplication
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type CertificateDescription struct {
	Thumbprint          string  `pulumi:"thumbprint"`
	ThumbprintSecondary *string `pulumi:"thumbprintSecondary"`
	X509StoreName       *string `pulumi:"x509StoreName"`
}





type CertificateDescriptionInput interface {
	pulumi.Input

	ToCertificateDescriptionOutput() CertificateDescriptionOutput
	ToCertificateDescriptionOutputWithContext(context.Context) CertificateDescriptionOutput
}

type CertificateDescriptionArgs struct {
	Thumbprint          pulumi.StringInput    `pulumi:"thumbprint"`
	ThumbprintSecondary pulumi.StringPtrInput `pulumi:"thumbprintSecondary"`
	X509StoreName       pulumi.StringPtrInput `pulumi:"x509StoreName"`
}

func (CertificateDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateDescription)(nil)).Elem()
}

func (i CertificateDescriptionArgs) ToCertificateDescriptionOutput() CertificateDescriptionOutput {
	return i.ToCertificateDescriptionOutputWithContext(context.Background())
}

func (i CertificateDescriptionArgs) ToCertificateDescriptionOutputWithContext(ctx context.Context) CertificateDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateDescriptionOutput)
}

func (i CertificateDescriptionArgs) ToCertificateDescriptionPtrOutput() CertificateDescriptionPtrOutput {
	return i.ToCertificateDescriptionPtrOutputWithContext(context.Background())
}

func (i CertificateDescriptionArgs) ToCertificateDescriptionPtrOutputWithContext(ctx context.Context) CertificateDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateDescriptionOutput).ToCertificateDescriptionPtrOutputWithContext(ctx)
}









type CertificateDescriptionPtrInput interface {
	pulumi.Input

	ToCertificateDescriptionPtrOutput() CertificateDescriptionPtrOutput
	ToCertificateDescriptionPtrOutputWithContext(context.Context) CertificateDescriptionPtrOutput
}

type certificateDescriptionPtrType CertificateDescriptionArgs

func CertificateDescriptionPtr(v *CertificateDescriptionArgs) CertificateDescriptionPtrInput {
	return (*certificateDescriptionPtrType)(v)
}

func (*certificateDescriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateDescription)(nil)).Elem()
}

func (i *certificateDescriptionPtrType) ToCertificateDescriptionPtrOutput() CertificateDescriptionPtrOutput {
	return i.ToCertificateDescriptionPtrOutputWithContext(context.Background())
}

func (i *certificateDescriptionPtrType) ToCertificateDescriptionPtrOutputWithContext(ctx context.Context) CertificateDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateDescriptionPtrOutput)
}

type CertificateDescriptionOutput struct{ *pulumi.OutputState }

func (CertificateDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateDescription)(nil)).Elem()
}

func (o CertificateDescriptionOutput) ToCertificateDescriptionOutput() CertificateDescriptionOutput {
	return o
}

func (o CertificateDescriptionOutput) ToCertificateDescriptionOutputWithContext(ctx context.Context) CertificateDescriptionOutput {
	return o
}

func (o CertificateDescriptionOutput) ToCertificateDescriptionPtrOutput() CertificateDescriptionPtrOutput {
	return o.ToCertificateDescriptionPtrOutputWithContext(context.Background())
}

func (o CertificateDescriptionOutput) ToCertificateDescriptionPtrOutputWithContext(ctx context.Context) CertificateDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificateDescription) *CertificateDescription {
		return &v
	}).(CertificateDescriptionPtrOutput)
}

func (o CertificateDescriptionOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDescription) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o CertificateDescriptionOutput) ThumbprintSecondary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateDescription) *string { return v.ThumbprintSecondary }).(pulumi.StringPtrOutput)
}

func (o CertificateDescriptionOutput) X509StoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateDescription) *string { return v.X509StoreName }).(pulumi.StringPtrOutput)
}

type CertificateDescriptionPtrOutput struct{ *pulumi.OutputState }

func (CertificateDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateDescription)(nil)).Elem()
}

func (o CertificateDescriptionPtrOutput) ToCertificateDescriptionPtrOutput() CertificateDescriptionPtrOutput {
	return o
}

func (o CertificateDescriptionPtrOutput) ToCertificateDescriptionPtrOutputWithContext(ctx context.Context) CertificateDescriptionPtrOutput {
	return o
}

func (o CertificateDescriptionPtrOutput) Elem() CertificateDescriptionOutput {
	return o.ApplyT(func(v *CertificateDescription) CertificateDescription {
		if v != nil {
			return *v
		}
		var ret CertificateDescription
		return ret
	}).(CertificateDescriptionOutput)
}

func (o CertificateDescriptionPtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDescription) *string {
		if v == nil {
			return nil
		}
		return &v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

func (o CertificateDescriptionPtrOutput) ThumbprintSecondary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDescription) *string {
		if v == nil {
			return nil
		}
		return v.ThumbprintSecondary
	}).(pulumi.StringPtrOutput)
}

func (o CertificateDescriptionPtrOutput) X509StoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDescription) *string {
		if v == nil {
			return nil
		}
		return v.X509StoreName
	}).(pulumi.StringPtrOutput)
}

type CertificateDescriptionResponse struct {
	Thumbprint          string  `pulumi:"thumbprint"`
	ThumbprintSecondary *string `pulumi:"thumbprintSecondary"`
	X509StoreName       *string `pulumi:"x509StoreName"`
}





type CertificateDescriptionResponseInput interface {
	pulumi.Input

	ToCertificateDescriptionResponseOutput() CertificateDescriptionResponseOutput
	ToCertificateDescriptionResponseOutputWithContext(context.Context) CertificateDescriptionResponseOutput
}

type CertificateDescriptionResponseArgs struct {
	Thumbprint          pulumi.StringInput    `pulumi:"thumbprint"`
	ThumbprintSecondary pulumi.StringPtrInput `pulumi:"thumbprintSecondary"`
	X509StoreName       pulumi.StringPtrInput `pulumi:"x509StoreName"`
}

func (CertificateDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateDescriptionResponse)(nil)).Elem()
}

func (i CertificateDescriptionResponseArgs) ToCertificateDescriptionResponseOutput() CertificateDescriptionResponseOutput {
	return i.ToCertificateDescriptionResponseOutputWithContext(context.Background())
}

func (i CertificateDescriptionResponseArgs) ToCertificateDescriptionResponseOutputWithContext(ctx context.Context) CertificateDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateDescriptionResponseOutput)
}

func (i CertificateDescriptionResponseArgs) ToCertificateDescriptionResponsePtrOutput() CertificateDescriptionResponsePtrOutput {
	return i.ToCertificateDescriptionResponsePtrOutputWithContext(context.Background())
}

func (i CertificateDescriptionResponseArgs) ToCertificateDescriptionResponsePtrOutputWithContext(ctx context.Context) CertificateDescriptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateDescriptionResponseOutput).ToCertificateDescriptionResponsePtrOutputWithContext(ctx)
}









type CertificateDescriptionResponsePtrInput interface {
	pulumi.Input

	ToCertificateDescriptionResponsePtrOutput() CertificateDescriptionResponsePtrOutput
	ToCertificateDescriptionResponsePtrOutputWithContext(context.Context) CertificateDescriptionResponsePtrOutput
}

type certificateDescriptionResponsePtrType CertificateDescriptionResponseArgs

func CertificateDescriptionResponsePtr(v *CertificateDescriptionResponseArgs) CertificateDescriptionResponsePtrInput {
	return (*certificateDescriptionResponsePtrType)(v)
}

func (*certificateDescriptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateDescriptionResponse)(nil)).Elem()
}

func (i *certificateDescriptionResponsePtrType) ToCertificateDescriptionResponsePtrOutput() CertificateDescriptionResponsePtrOutput {
	return i.ToCertificateDescriptionResponsePtrOutputWithContext(context.Background())
}

func (i *certificateDescriptionResponsePtrType) ToCertificateDescriptionResponsePtrOutputWithContext(ctx context.Context) CertificateDescriptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateDescriptionResponsePtrOutput)
}

type CertificateDescriptionResponseOutput struct{ *pulumi.OutputState }

func (CertificateDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateDescriptionResponse)(nil)).Elem()
}

func (o CertificateDescriptionResponseOutput) ToCertificateDescriptionResponseOutput() CertificateDescriptionResponseOutput {
	return o
}

func (o CertificateDescriptionResponseOutput) ToCertificateDescriptionResponseOutputWithContext(ctx context.Context) CertificateDescriptionResponseOutput {
	return o
}

func (o CertificateDescriptionResponseOutput) ToCertificateDescriptionResponsePtrOutput() CertificateDescriptionResponsePtrOutput {
	return o.ToCertificateDescriptionResponsePtrOutputWithContext(context.Background())
}

func (o CertificateDescriptionResponseOutput) ToCertificateDescriptionResponsePtrOutputWithContext(ctx context.Context) CertificateDescriptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificateDescriptionResponse) *CertificateDescriptionResponse {
		return &v
	}).(CertificateDescriptionResponsePtrOutput)
}

func (o CertificateDescriptionResponseOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDescriptionResponse) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o CertificateDescriptionResponseOutput) ThumbprintSecondary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateDescriptionResponse) *string { return v.ThumbprintSecondary }).(pulumi.StringPtrOutput)
}

func (o CertificateDescriptionResponseOutput) X509StoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateDescriptionResponse) *string { return v.X509StoreName }).(pulumi.StringPtrOutput)
}

type CertificateDescriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (CertificateDescriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateDescriptionResponse)(nil)).Elem()
}

func (o CertificateDescriptionResponsePtrOutput) ToCertificateDescriptionResponsePtrOutput() CertificateDescriptionResponsePtrOutput {
	return o
}

func (o CertificateDescriptionResponsePtrOutput) ToCertificateDescriptionResponsePtrOutputWithContext(ctx context.Context) CertificateDescriptionResponsePtrOutput {
	return o
}

func (o CertificateDescriptionResponsePtrOutput) Elem() CertificateDescriptionResponseOutput {
	return o.ApplyT(func(v *CertificateDescriptionResponse) CertificateDescriptionResponse {
		if v != nil {
			return *v
		}
		var ret CertificateDescriptionResponse
		return ret
	}).(CertificateDescriptionResponseOutput)
}

func (o CertificateDescriptionResponsePtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

func (o CertificateDescriptionResponsePtrOutput) ThumbprintSecondary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ThumbprintSecondary
	}).(pulumi.StringPtrOutput)
}

func (o CertificateDescriptionResponsePtrOutput) X509StoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.X509StoreName
	}).(pulumi.StringPtrOutput)
}

type ClientCertificate struct {
	CommonName       *string `pulumi:"commonName"`
	IsAdmin          bool    `pulumi:"isAdmin"`
	IssuerThumbprint *string `pulumi:"issuerThumbprint"`
	Thumbprint       *string `pulumi:"thumbprint"`
}





type ClientCertificateInput interface {
	pulumi.Input

	ToClientCertificateOutput() ClientCertificateOutput
	ToClientCertificateOutputWithContext(context.Context) ClientCertificateOutput
}

type ClientCertificateArgs struct {
	CommonName       pulumi.StringPtrInput `pulumi:"commonName"`
	IsAdmin          pulumi.BoolInput      `pulumi:"isAdmin"`
	IssuerThumbprint pulumi.StringPtrInput `pulumi:"issuerThumbprint"`
	Thumbprint       pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (ClientCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificate)(nil)).Elem()
}

func (i ClientCertificateArgs) ToClientCertificateOutput() ClientCertificateOutput {
	return i.ToClientCertificateOutputWithContext(context.Background())
}

func (i ClientCertificateArgs) ToClientCertificateOutputWithContext(ctx context.Context) ClientCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateOutput)
}





type ClientCertificateArrayInput interface {
	pulumi.Input

	ToClientCertificateArrayOutput() ClientCertificateArrayOutput
	ToClientCertificateArrayOutputWithContext(context.Context) ClientCertificateArrayOutput
}

type ClientCertificateArray []ClientCertificateInput

func (ClientCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificate)(nil)).Elem()
}

func (i ClientCertificateArray) ToClientCertificateArrayOutput() ClientCertificateArrayOutput {
	return i.ToClientCertificateArrayOutputWithContext(context.Background())
}

func (i ClientCertificateArray) ToClientCertificateArrayOutputWithContext(ctx context.Context) ClientCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateArrayOutput)
}

type ClientCertificateOutput struct{ *pulumi.OutputState }

func (ClientCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificate)(nil)).Elem()
}

func (o ClientCertificateOutput) ToClientCertificateOutput() ClientCertificateOutput {
	return o
}

func (o ClientCertificateOutput) ToClientCertificateOutputWithContext(ctx context.Context) ClientCertificateOutput {
	return o
}

func (o ClientCertificateOutput) CommonName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCertificate) *string { return v.CommonName }).(pulumi.StringPtrOutput)
}

func (o ClientCertificateOutput) IsAdmin() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientCertificate) bool { return v.IsAdmin }).(pulumi.BoolOutput)
}

func (o ClientCertificateOutput) IssuerThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCertificate) *string { return v.IssuerThumbprint }).(pulumi.StringPtrOutput)
}

func (o ClientCertificateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCertificate) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type ClientCertificateArrayOutput struct{ *pulumi.OutputState }

func (ClientCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificate)(nil)).Elem()
}

func (o ClientCertificateArrayOutput) ToClientCertificateArrayOutput() ClientCertificateArrayOutput {
	return o
}

func (o ClientCertificateArrayOutput) ToClientCertificateArrayOutputWithContext(ctx context.Context) ClientCertificateArrayOutput {
	return o
}

func (o ClientCertificateArrayOutput) Index(i pulumi.IntInput) ClientCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientCertificate {
		return vs[0].([]ClientCertificate)[vs[1].(int)]
	}).(ClientCertificateOutput)
}

type ClientCertificateCommonName struct {
	CertificateCommonName       string `pulumi:"certificateCommonName"`
	CertificateIssuerThumbprint string `pulumi:"certificateIssuerThumbprint"`
	IsAdmin                     bool   `pulumi:"isAdmin"`
}





type ClientCertificateCommonNameInput interface {
	pulumi.Input

	ToClientCertificateCommonNameOutput() ClientCertificateCommonNameOutput
	ToClientCertificateCommonNameOutputWithContext(context.Context) ClientCertificateCommonNameOutput
}

type ClientCertificateCommonNameArgs struct {
	CertificateCommonName       pulumi.StringInput `pulumi:"certificateCommonName"`
	CertificateIssuerThumbprint pulumi.StringInput `pulumi:"certificateIssuerThumbprint"`
	IsAdmin                     pulumi.BoolInput   `pulumi:"isAdmin"`
}

func (ClientCertificateCommonNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateCommonName)(nil)).Elem()
}

func (i ClientCertificateCommonNameArgs) ToClientCertificateCommonNameOutput() ClientCertificateCommonNameOutput {
	return i.ToClientCertificateCommonNameOutputWithContext(context.Background())
}

func (i ClientCertificateCommonNameArgs) ToClientCertificateCommonNameOutputWithContext(ctx context.Context) ClientCertificateCommonNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateCommonNameOutput)
}





type ClientCertificateCommonNameArrayInput interface {
	pulumi.Input

	ToClientCertificateCommonNameArrayOutput() ClientCertificateCommonNameArrayOutput
	ToClientCertificateCommonNameArrayOutputWithContext(context.Context) ClientCertificateCommonNameArrayOutput
}

type ClientCertificateCommonNameArray []ClientCertificateCommonNameInput

func (ClientCertificateCommonNameArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificateCommonName)(nil)).Elem()
}

func (i ClientCertificateCommonNameArray) ToClientCertificateCommonNameArrayOutput() ClientCertificateCommonNameArrayOutput {
	return i.ToClientCertificateCommonNameArrayOutputWithContext(context.Background())
}

func (i ClientCertificateCommonNameArray) ToClientCertificateCommonNameArrayOutputWithContext(ctx context.Context) ClientCertificateCommonNameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateCommonNameArrayOutput)
}

type ClientCertificateCommonNameOutput struct{ *pulumi.OutputState }

func (ClientCertificateCommonNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateCommonName)(nil)).Elem()
}

func (o ClientCertificateCommonNameOutput) ToClientCertificateCommonNameOutput() ClientCertificateCommonNameOutput {
	return o
}

func (o ClientCertificateCommonNameOutput) ToClientCertificateCommonNameOutputWithContext(ctx context.Context) ClientCertificateCommonNameOutput {
	return o
}

func (o ClientCertificateCommonNameOutput) CertificateCommonName() pulumi.StringOutput {
	return o.ApplyT(func(v ClientCertificateCommonName) string { return v.CertificateCommonName }).(pulumi.StringOutput)
}

func (o ClientCertificateCommonNameOutput) CertificateIssuerThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v ClientCertificateCommonName) string { return v.CertificateIssuerThumbprint }).(pulumi.StringOutput)
}

func (o ClientCertificateCommonNameOutput) IsAdmin() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientCertificateCommonName) bool { return v.IsAdmin }).(pulumi.BoolOutput)
}

type ClientCertificateCommonNameArrayOutput struct{ *pulumi.OutputState }

func (ClientCertificateCommonNameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificateCommonName)(nil)).Elem()
}

func (o ClientCertificateCommonNameArrayOutput) ToClientCertificateCommonNameArrayOutput() ClientCertificateCommonNameArrayOutput {
	return o
}

func (o ClientCertificateCommonNameArrayOutput) ToClientCertificateCommonNameArrayOutputWithContext(ctx context.Context) ClientCertificateCommonNameArrayOutput {
	return o
}

func (o ClientCertificateCommonNameArrayOutput) Index(i pulumi.IntInput) ClientCertificateCommonNameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientCertificateCommonName {
		return vs[0].([]ClientCertificateCommonName)[vs[1].(int)]
	}).(ClientCertificateCommonNameOutput)
}

type ClientCertificateCommonNameResponse struct {
	CertificateCommonName       string `pulumi:"certificateCommonName"`
	CertificateIssuerThumbprint string `pulumi:"certificateIssuerThumbprint"`
	IsAdmin                     bool   `pulumi:"isAdmin"`
}





type ClientCertificateCommonNameResponseInput interface {
	pulumi.Input

	ToClientCertificateCommonNameResponseOutput() ClientCertificateCommonNameResponseOutput
	ToClientCertificateCommonNameResponseOutputWithContext(context.Context) ClientCertificateCommonNameResponseOutput
}

type ClientCertificateCommonNameResponseArgs struct {
	CertificateCommonName       pulumi.StringInput `pulumi:"certificateCommonName"`
	CertificateIssuerThumbprint pulumi.StringInput `pulumi:"certificateIssuerThumbprint"`
	IsAdmin                     pulumi.BoolInput   `pulumi:"isAdmin"`
}

func (ClientCertificateCommonNameResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateCommonNameResponse)(nil)).Elem()
}

func (i ClientCertificateCommonNameResponseArgs) ToClientCertificateCommonNameResponseOutput() ClientCertificateCommonNameResponseOutput {
	return i.ToClientCertificateCommonNameResponseOutputWithContext(context.Background())
}

func (i ClientCertificateCommonNameResponseArgs) ToClientCertificateCommonNameResponseOutputWithContext(ctx context.Context) ClientCertificateCommonNameResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateCommonNameResponseOutput)
}





type ClientCertificateCommonNameResponseArrayInput interface {
	pulumi.Input

	ToClientCertificateCommonNameResponseArrayOutput() ClientCertificateCommonNameResponseArrayOutput
	ToClientCertificateCommonNameResponseArrayOutputWithContext(context.Context) ClientCertificateCommonNameResponseArrayOutput
}

type ClientCertificateCommonNameResponseArray []ClientCertificateCommonNameResponseInput

func (ClientCertificateCommonNameResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificateCommonNameResponse)(nil)).Elem()
}

func (i ClientCertificateCommonNameResponseArray) ToClientCertificateCommonNameResponseArrayOutput() ClientCertificateCommonNameResponseArrayOutput {
	return i.ToClientCertificateCommonNameResponseArrayOutputWithContext(context.Background())
}

func (i ClientCertificateCommonNameResponseArray) ToClientCertificateCommonNameResponseArrayOutputWithContext(ctx context.Context) ClientCertificateCommonNameResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateCommonNameResponseArrayOutput)
}

type ClientCertificateCommonNameResponseOutput struct{ *pulumi.OutputState }

func (ClientCertificateCommonNameResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateCommonNameResponse)(nil)).Elem()
}

func (o ClientCertificateCommonNameResponseOutput) ToClientCertificateCommonNameResponseOutput() ClientCertificateCommonNameResponseOutput {
	return o
}

func (o ClientCertificateCommonNameResponseOutput) ToClientCertificateCommonNameResponseOutputWithContext(ctx context.Context) ClientCertificateCommonNameResponseOutput {
	return o
}

func (o ClientCertificateCommonNameResponseOutput) CertificateCommonName() pulumi.StringOutput {
	return o.ApplyT(func(v ClientCertificateCommonNameResponse) string { return v.CertificateCommonName }).(pulumi.StringOutput)
}

func (o ClientCertificateCommonNameResponseOutput) CertificateIssuerThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v ClientCertificateCommonNameResponse) string { return v.CertificateIssuerThumbprint }).(pulumi.StringOutput)
}

func (o ClientCertificateCommonNameResponseOutput) IsAdmin() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientCertificateCommonNameResponse) bool { return v.IsAdmin }).(pulumi.BoolOutput)
}

type ClientCertificateCommonNameResponseArrayOutput struct{ *pulumi.OutputState }

func (ClientCertificateCommonNameResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificateCommonNameResponse)(nil)).Elem()
}

func (o ClientCertificateCommonNameResponseArrayOutput) ToClientCertificateCommonNameResponseArrayOutput() ClientCertificateCommonNameResponseArrayOutput {
	return o
}

func (o ClientCertificateCommonNameResponseArrayOutput) ToClientCertificateCommonNameResponseArrayOutputWithContext(ctx context.Context) ClientCertificateCommonNameResponseArrayOutput {
	return o
}

func (o ClientCertificateCommonNameResponseArrayOutput) Index(i pulumi.IntInput) ClientCertificateCommonNameResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientCertificateCommonNameResponse {
		return vs[0].([]ClientCertificateCommonNameResponse)[vs[1].(int)]
	}).(ClientCertificateCommonNameResponseOutput)
}

type ClientCertificateResponse struct {
	CommonName       *string `pulumi:"commonName"`
	IsAdmin          bool    `pulumi:"isAdmin"`
	IssuerThumbprint *string `pulumi:"issuerThumbprint"`
	Thumbprint       *string `pulumi:"thumbprint"`
}





type ClientCertificateResponseInput interface {
	pulumi.Input

	ToClientCertificateResponseOutput() ClientCertificateResponseOutput
	ToClientCertificateResponseOutputWithContext(context.Context) ClientCertificateResponseOutput
}

type ClientCertificateResponseArgs struct {
	CommonName       pulumi.StringPtrInput `pulumi:"commonName"`
	IsAdmin          pulumi.BoolInput      `pulumi:"isAdmin"`
	IssuerThumbprint pulumi.StringPtrInput `pulumi:"issuerThumbprint"`
	Thumbprint       pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (ClientCertificateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateResponse)(nil)).Elem()
}

func (i ClientCertificateResponseArgs) ToClientCertificateResponseOutput() ClientCertificateResponseOutput {
	return i.ToClientCertificateResponseOutputWithContext(context.Background())
}

func (i ClientCertificateResponseArgs) ToClientCertificateResponseOutputWithContext(ctx context.Context) ClientCertificateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateResponseOutput)
}





type ClientCertificateResponseArrayInput interface {
	pulumi.Input

	ToClientCertificateResponseArrayOutput() ClientCertificateResponseArrayOutput
	ToClientCertificateResponseArrayOutputWithContext(context.Context) ClientCertificateResponseArrayOutput
}

type ClientCertificateResponseArray []ClientCertificateResponseInput

func (ClientCertificateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificateResponse)(nil)).Elem()
}

func (i ClientCertificateResponseArray) ToClientCertificateResponseArrayOutput() ClientCertificateResponseArrayOutput {
	return i.ToClientCertificateResponseArrayOutputWithContext(context.Background())
}

func (i ClientCertificateResponseArray) ToClientCertificateResponseArrayOutputWithContext(ctx context.Context) ClientCertificateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateResponseArrayOutput)
}

type ClientCertificateResponseOutput struct{ *pulumi.OutputState }

func (ClientCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateResponse)(nil)).Elem()
}

func (o ClientCertificateResponseOutput) ToClientCertificateResponseOutput() ClientCertificateResponseOutput {
	return o
}

func (o ClientCertificateResponseOutput) ToClientCertificateResponseOutputWithContext(ctx context.Context) ClientCertificateResponseOutput {
	return o
}

func (o ClientCertificateResponseOutput) CommonName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCertificateResponse) *string { return v.CommonName }).(pulumi.StringPtrOutput)
}

func (o ClientCertificateResponseOutput) IsAdmin() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientCertificateResponse) bool { return v.IsAdmin }).(pulumi.BoolOutput)
}

func (o ClientCertificateResponseOutput) IssuerThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCertificateResponse) *string { return v.IssuerThumbprint }).(pulumi.StringPtrOutput)
}

func (o ClientCertificateResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCertificateResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type ClientCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (ClientCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificateResponse)(nil)).Elem()
}

func (o ClientCertificateResponseArrayOutput) ToClientCertificateResponseArrayOutput() ClientCertificateResponseArrayOutput {
	return o
}

func (o ClientCertificateResponseArrayOutput) ToClientCertificateResponseArrayOutputWithContext(ctx context.Context) ClientCertificateResponseArrayOutput {
	return o
}

func (o ClientCertificateResponseArrayOutput) Index(i pulumi.IntInput) ClientCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientCertificateResponse {
		return vs[0].([]ClientCertificateResponse)[vs[1].(int)]
	}).(ClientCertificateResponseOutput)
}

type ClientCertificateThumbprint struct {
	CertificateThumbprint string `pulumi:"certificateThumbprint"`
	IsAdmin               bool   `pulumi:"isAdmin"`
}





type ClientCertificateThumbprintInput interface {
	pulumi.Input

	ToClientCertificateThumbprintOutput() ClientCertificateThumbprintOutput
	ToClientCertificateThumbprintOutputWithContext(context.Context) ClientCertificateThumbprintOutput
}

type ClientCertificateThumbprintArgs struct {
	CertificateThumbprint pulumi.StringInput `pulumi:"certificateThumbprint"`
	IsAdmin               pulumi.BoolInput   `pulumi:"isAdmin"`
}

func (ClientCertificateThumbprintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateThumbprint)(nil)).Elem()
}

func (i ClientCertificateThumbprintArgs) ToClientCertificateThumbprintOutput() ClientCertificateThumbprintOutput {
	return i.ToClientCertificateThumbprintOutputWithContext(context.Background())
}

func (i ClientCertificateThumbprintArgs) ToClientCertificateThumbprintOutputWithContext(ctx context.Context) ClientCertificateThumbprintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateThumbprintOutput)
}





type ClientCertificateThumbprintArrayInput interface {
	pulumi.Input

	ToClientCertificateThumbprintArrayOutput() ClientCertificateThumbprintArrayOutput
	ToClientCertificateThumbprintArrayOutputWithContext(context.Context) ClientCertificateThumbprintArrayOutput
}

type ClientCertificateThumbprintArray []ClientCertificateThumbprintInput

func (ClientCertificateThumbprintArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificateThumbprint)(nil)).Elem()
}

func (i ClientCertificateThumbprintArray) ToClientCertificateThumbprintArrayOutput() ClientCertificateThumbprintArrayOutput {
	return i.ToClientCertificateThumbprintArrayOutputWithContext(context.Background())
}

func (i ClientCertificateThumbprintArray) ToClientCertificateThumbprintArrayOutputWithContext(ctx context.Context) ClientCertificateThumbprintArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateThumbprintArrayOutput)
}

type ClientCertificateThumbprintOutput struct{ *pulumi.OutputState }

func (ClientCertificateThumbprintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateThumbprint)(nil)).Elem()
}

func (o ClientCertificateThumbprintOutput) ToClientCertificateThumbprintOutput() ClientCertificateThumbprintOutput {
	return o
}

func (o ClientCertificateThumbprintOutput) ToClientCertificateThumbprintOutputWithContext(ctx context.Context) ClientCertificateThumbprintOutput {
	return o
}

func (o ClientCertificateThumbprintOutput) CertificateThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v ClientCertificateThumbprint) string { return v.CertificateThumbprint }).(pulumi.StringOutput)
}

func (o ClientCertificateThumbprintOutput) IsAdmin() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientCertificateThumbprint) bool { return v.IsAdmin }).(pulumi.BoolOutput)
}

type ClientCertificateThumbprintArrayOutput struct{ *pulumi.OutputState }

func (ClientCertificateThumbprintArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificateThumbprint)(nil)).Elem()
}

func (o ClientCertificateThumbprintArrayOutput) ToClientCertificateThumbprintArrayOutput() ClientCertificateThumbprintArrayOutput {
	return o
}

func (o ClientCertificateThumbprintArrayOutput) ToClientCertificateThumbprintArrayOutputWithContext(ctx context.Context) ClientCertificateThumbprintArrayOutput {
	return o
}

func (o ClientCertificateThumbprintArrayOutput) Index(i pulumi.IntInput) ClientCertificateThumbprintOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientCertificateThumbprint {
		return vs[0].([]ClientCertificateThumbprint)[vs[1].(int)]
	}).(ClientCertificateThumbprintOutput)
}

type ClientCertificateThumbprintResponse struct {
	CertificateThumbprint string `pulumi:"certificateThumbprint"`
	IsAdmin               bool   `pulumi:"isAdmin"`
}





type ClientCertificateThumbprintResponseInput interface {
	pulumi.Input

	ToClientCertificateThumbprintResponseOutput() ClientCertificateThumbprintResponseOutput
	ToClientCertificateThumbprintResponseOutputWithContext(context.Context) ClientCertificateThumbprintResponseOutput
}

type ClientCertificateThumbprintResponseArgs struct {
	CertificateThumbprint pulumi.StringInput `pulumi:"certificateThumbprint"`
	IsAdmin               pulumi.BoolInput   `pulumi:"isAdmin"`
}

func (ClientCertificateThumbprintResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateThumbprintResponse)(nil)).Elem()
}

func (i ClientCertificateThumbprintResponseArgs) ToClientCertificateThumbprintResponseOutput() ClientCertificateThumbprintResponseOutput {
	return i.ToClientCertificateThumbprintResponseOutputWithContext(context.Background())
}

func (i ClientCertificateThumbprintResponseArgs) ToClientCertificateThumbprintResponseOutputWithContext(ctx context.Context) ClientCertificateThumbprintResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateThumbprintResponseOutput)
}





type ClientCertificateThumbprintResponseArrayInput interface {
	pulumi.Input

	ToClientCertificateThumbprintResponseArrayOutput() ClientCertificateThumbprintResponseArrayOutput
	ToClientCertificateThumbprintResponseArrayOutputWithContext(context.Context) ClientCertificateThumbprintResponseArrayOutput
}

type ClientCertificateThumbprintResponseArray []ClientCertificateThumbprintResponseInput

func (ClientCertificateThumbprintResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificateThumbprintResponse)(nil)).Elem()
}

func (i ClientCertificateThumbprintResponseArray) ToClientCertificateThumbprintResponseArrayOutput() ClientCertificateThumbprintResponseArrayOutput {
	return i.ToClientCertificateThumbprintResponseArrayOutputWithContext(context.Background())
}

func (i ClientCertificateThumbprintResponseArray) ToClientCertificateThumbprintResponseArrayOutputWithContext(ctx context.Context) ClientCertificateThumbprintResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateThumbprintResponseArrayOutput)
}

type ClientCertificateThumbprintResponseOutput struct{ *pulumi.OutputState }

func (ClientCertificateThumbprintResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateThumbprintResponse)(nil)).Elem()
}

func (o ClientCertificateThumbprintResponseOutput) ToClientCertificateThumbprintResponseOutput() ClientCertificateThumbprintResponseOutput {
	return o
}

func (o ClientCertificateThumbprintResponseOutput) ToClientCertificateThumbprintResponseOutputWithContext(ctx context.Context) ClientCertificateThumbprintResponseOutput {
	return o
}

func (o ClientCertificateThumbprintResponseOutput) CertificateThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v ClientCertificateThumbprintResponse) string { return v.CertificateThumbprint }).(pulumi.StringOutput)
}

func (o ClientCertificateThumbprintResponseOutput) IsAdmin() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientCertificateThumbprintResponse) bool { return v.IsAdmin }).(pulumi.BoolOutput)
}

type ClientCertificateThumbprintResponseArrayOutput struct{ *pulumi.OutputState }

func (ClientCertificateThumbprintResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificateThumbprintResponse)(nil)).Elem()
}

func (o ClientCertificateThumbprintResponseArrayOutput) ToClientCertificateThumbprintResponseArrayOutput() ClientCertificateThumbprintResponseArrayOutput {
	return o
}

func (o ClientCertificateThumbprintResponseArrayOutput) ToClientCertificateThumbprintResponseArrayOutputWithContext(ctx context.Context) ClientCertificateThumbprintResponseArrayOutput {
	return o
}

func (o ClientCertificateThumbprintResponseArrayOutput) Index(i pulumi.IntInput) ClientCertificateThumbprintResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientCertificateThumbprintResponse {
		return vs[0].([]ClientCertificateThumbprintResponse)[vs[1].(int)]
	}).(ClientCertificateThumbprintResponseOutput)
}

type ClusterHealthPolicy struct {
	ApplicationHealthPolicies       map[string]ApplicationHealthPolicy `pulumi:"applicationHealthPolicies"`
	MaxPercentUnhealthyApplications *int                               `pulumi:"maxPercentUnhealthyApplications"`
	MaxPercentUnhealthyNodes        *int                               `pulumi:"maxPercentUnhealthyNodes"`
}


func (val *ClusterHealthPolicy) Defaults() *ClusterHealthPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxPercentUnhealthyApplications) {
		maxPercentUnhealthyApplications_ := 0
		tmp.MaxPercentUnhealthyApplications = &maxPercentUnhealthyApplications_
	}
	if isZero(tmp.MaxPercentUnhealthyNodes) {
		maxPercentUnhealthyNodes_ := 0
		tmp.MaxPercentUnhealthyNodes = &maxPercentUnhealthyNodes_
	}
	return &tmp
}





type ClusterHealthPolicyInput interface {
	pulumi.Input

	ToClusterHealthPolicyOutput() ClusterHealthPolicyOutput
	ToClusterHealthPolicyOutputWithContext(context.Context) ClusterHealthPolicyOutput
}

type ClusterHealthPolicyArgs struct {
	ApplicationHealthPolicies       ApplicationHealthPolicyMapInput `pulumi:"applicationHealthPolicies"`
	MaxPercentUnhealthyApplications pulumi.IntPtrInput              `pulumi:"maxPercentUnhealthyApplications"`
	MaxPercentUnhealthyNodes        pulumi.IntPtrInput              `pulumi:"maxPercentUnhealthyNodes"`
}

func (ClusterHealthPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterHealthPolicy)(nil)).Elem()
}

func (i ClusterHealthPolicyArgs) ToClusterHealthPolicyOutput() ClusterHealthPolicyOutput {
	return i.ToClusterHealthPolicyOutputWithContext(context.Background())
}

func (i ClusterHealthPolicyArgs) ToClusterHealthPolicyOutputWithContext(ctx context.Context) ClusterHealthPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterHealthPolicyOutput)
}

func (i ClusterHealthPolicyArgs) ToClusterHealthPolicyPtrOutput() ClusterHealthPolicyPtrOutput {
	return i.ToClusterHealthPolicyPtrOutputWithContext(context.Background())
}

func (i ClusterHealthPolicyArgs) ToClusterHealthPolicyPtrOutputWithContext(ctx context.Context) ClusterHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterHealthPolicyOutput).ToClusterHealthPolicyPtrOutputWithContext(ctx)
}









type ClusterHealthPolicyPtrInput interface {
	pulumi.Input

	ToClusterHealthPolicyPtrOutput() ClusterHealthPolicyPtrOutput
	ToClusterHealthPolicyPtrOutputWithContext(context.Context) ClusterHealthPolicyPtrOutput
}

type clusterHealthPolicyPtrType ClusterHealthPolicyArgs

func ClusterHealthPolicyPtr(v *ClusterHealthPolicyArgs) ClusterHealthPolicyPtrInput {
	return (*clusterHealthPolicyPtrType)(v)
}

func (*clusterHealthPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterHealthPolicy)(nil)).Elem()
}

func (i *clusterHealthPolicyPtrType) ToClusterHealthPolicyPtrOutput() ClusterHealthPolicyPtrOutput {
	return i.ToClusterHealthPolicyPtrOutputWithContext(context.Background())
}

func (i *clusterHealthPolicyPtrType) ToClusterHealthPolicyPtrOutputWithContext(ctx context.Context) ClusterHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterHealthPolicyPtrOutput)
}

type ClusterHealthPolicyOutput struct{ *pulumi.OutputState }

func (ClusterHealthPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterHealthPolicy)(nil)).Elem()
}

func (o ClusterHealthPolicyOutput) ToClusterHealthPolicyOutput() ClusterHealthPolicyOutput {
	return o
}

func (o ClusterHealthPolicyOutput) ToClusterHealthPolicyOutputWithContext(ctx context.Context) ClusterHealthPolicyOutput {
	return o
}

func (o ClusterHealthPolicyOutput) ToClusterHealthPolicyPtrOutput() ClusterHealthPolicyPtrOutput {
	return o.ToClusterHealthPolicyPtrOutputWithContext(context.Background())
}

func (o ClusterHealthPolicyOutput) ToClusterHealthPolicyPtrOutputWithContext(ctx context.Context) ClusterHealthPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterHealthPolicy) *ClusterHealthPolicy {
		return &v
	}).(ClusterHealthPolicyPtrOutput)
}

func (o ClusterHealthPolicyOutput) ApplicationHealthPolicies() ApplicationHealthPolicyMapOutput {
	return o.ApplyT(func(v ClusterHealthPolicy) map[string]ApplicationHealthPolicy { return v.ApplicationHealthPolicies }).(ApplicationHealthPolicyMapOutput)
}

func (o ClusterHealthPolicyOutput) MaxPercentUnhealthyApplications() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterHealthPolicy) *int { return v.MaxPercentUnhealthyApplications }).(pulumi.IntPtrOutput)
}

func (o ClusterHealthPolicyOutput) MaxPercentUnhealthyNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterHealthPolicy) *int { return v.MaxPercentUnhealthyNodes }).(pulumi.IntPtrOutput)
}

type ClusterHealthPolicyPtrOutput struct{ *pulumi.OutputState }

func (ClusterHealthPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterHealthPolicy)(nil)).Elem()
}

func (o ClusterHealthPolicyPtrOutput) ToClusterHealthPolicyPtrOutput() ClusterHealthPolicyPtrOutput {
	return o
}

func (o ClusterHealthPolicyPtrOutput) ToClusterHealthPolicyPtrOutputWithContext(ctx context.Context) ClusterHealthPolicyPtrOutput {
	return o
}

func (o ClusterHealthPolicyPtrOutput) Elem() ClusterHealthPolicyOutput {
	return o.ApplyT(func(v *ClusterHealthPolicy) ClusterHealthPolicy {
		if v != nil {
			return *v
		}
		var ret ClusterHealthPolicy
		return ret
	}).(ClusterHealthPolicyOutput)
}

func (o ClusterHealthPolicyPtrOutput) ApplicationHealthPolicies() ApplicationHealthPolicyMapOutput {
	return o.ApplyT(func(v *ClusterHealthPolicy) map[string]ApplicationHealthPolicy {
		if v == nil {
			return nil
		}
		return v.ApplicationHealthPolicies
	}).(ApplicationHealthPolicyMapOutput)
}

func (o ClusterHealthPolicyPtrOutput) MaxPercentUnhealthyApplications() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentUnhealthyApplications
	}).(pulumi.IntPtrOutput)
}

func (o ClusterHealthPolicyPtrOutput) MaxPercentUnhealthyNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentUnhealthyNodes
	}).(pulumi.IntPtrOutput)
}

type ClusterHealthPolicyResponse struct {
	ApplicationHealthPolicies       map[string]ApplicationHealthPolicyResponse `pulumi:"applicationHealthPolicies"`
	MaxPercentUnhealthyApplications *int                                       `pulumi:"maxPercentUnhealthyApplications"`
	MaxPercentUnhealthyNodes        *int                                       `pulumi:"maxPercentUnhealthyNodes"`
}


func (val *ClusterHealthPolicyResponse) Defaults() *ClusterHealthPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxPercentUnhealthyApplications) {
		maxPercentUnhealthyApplications_ := 0
		tmp.MaxPercentUnhealthyApplications = &maxPercentUnhealthyApplications_
	}
	if isZero(tmp.MaxPercentUnhealthyNodes) {
		maxPercentUnhealthyNodes_ := 0
		tmp.MaxPercentUnhealthyNodes = &maxPercentUnhealthyNodes_
	}
	return &tmp
}





type ClusterHealthPolicyResponseInput interface {
	pulumi.Input

	ToClusterHealthPolicyResponseOutput() ClusterHealthPolicyResponseOutput
	ToClusterHealthPolicyResponseOutputWithContext(context.Context) ClusterHealthPolicyResponseOutput
}

type ClusterHealthPolicyResponseArgs struct {
	ApplicationHealthPolicies       ApplicationHealthPolicyResponseMapInput `pulumi:"applicationHealthPolicies"`
	MaxPercentUnhealthyApplications pulumi.IntPtrInput                      `pulumi:"maxPercentUnhealthyApplications"`
	MaxPercentUnhealthyNodes        pulumi.IntPtrInput                      `pulumi:"maxPercentUnhealthyNodes"`
}

func (ClusterHealthPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterHealthPolicyResponse)(nil)).Elem()
}

func (i ClusterHealthPolicyResponseArgs) ToClusterHealthPolicyResponseOutput() ClusterHealthPolicyResponseOutput {
	return i.ToClusterHealthPolicyResponseOutputWithContext(context.Background())
}

func (i ClusterHealthPolicyResponseArgs) ToClusterHealthPolicyResponseOutputWithContext(ctx context.Context) ClusterHealthPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterHealthPolicyResponseOutput)
}

func (i ClusterHealthPolicyResponseArgs) ToClusterHealthPolicyResponsePtrOutput() ClusterHealthPolicyResponsePtrOutput {
	return i.ToClusterHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (i ClusterHealthPolicyResponseArgs) ToClusterHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ClusterHealthPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterHealthPolicyResponseOutput).ToClusterHealthPolicyResponsePtrOutputWithContext(ctx)
}









type ClusterHealthPolicyResponsePtrInput interface {
	pulumi.Input

	ToClusterHealthPolicyResponsePtrOutput() ClusterHealthPolicyResponsePtrOutput
	ToClusterHealthPolicyResponsePtrOutputWithContext(context.Context) ClusterHealthPolicyResponsePtrOutput
}

type clusterHealthPolicyResponsePtrType ClusterHealthPolicyResponseArgs

func ClusterHealthPolicyResponsePtr(v *ClusterHealthPolicyResponseArgs) ClusterHealthPolicyResponsePtrInput {
	return (*clusterHealthPolicyResponsePtrType)(v)
}

func (*clusterHealthPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterHealthPolicyResponse)(nil)).Elem()
}

func (i *clusterHealthPolicyResponsePtrType) ToClusterHealthPolicyResponsePtrOutput() ClusterHealthPolicyResponsePtrOutput {
	return i.ToClusterHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *clusterHealthPolicyResponsePtrType) ToClusterHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ClusterHealthPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterHealthPolicyResponsePtrOutput)
}

type ClusterHealthPolicyResponseOutput struct{ *pulumi.OutputState }

func (ClusterHealthPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterHealthPolicyResponse)(nil)).Elem()
}

func (o ClusterHealthPolicyResponseOutput) ToClusterHealthPolicyResponseOutput() ClusterHealthPolicyResponseOutput {
	return o
}

func (o ClusterHealthPolicyResponseOutput) ToClusterHealthPolicyResponseOutputWithContext(ctx context.Context) ClusterHealthPolicyResponseOutput {
	return o
}

func (o ClusterHealthPolicyResponseOutput) ToClusterHealthPolicyResponsePtrOutput() ClusterHealthPolicyResponsePtrOutput {
	return o.ToClusterHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (o ClusterHealthPolicyResponseOutput) ToClusterHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ClusterHealthPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterHealthPolicyResponse) *ClusterHealthPolicyResponse {
		return &v
	}).(ClusterHealthPolicyResponsePtrOutput)
}

func (o ClusterHealthPolicyResponseOutput) ApplicationHealthPolicies() ApplicationHealthPolicyResponseMapOutput {
	return o.ApplyT(func(v ClusterHealthPolicyResponse) map[string]ApplicationHealthPolicyResponse {
		return v.ApplicationHealthPolicies
	}).(ApplicationHealthPolicyResponseMapOutput)
}

func (o ClusterHealthPolicyResponseOutput) MaxPercentUnhealthyApplications() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterHealthPolicyResponse) *int { return v.MaxPercentUnhealthyApplications }).(pulumi.IntPtrOutput)
}

func (o ClusterHealthPolicyResponseOutput) MaxPercentUnhealthyNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterHealthPolicyResponse) *int { return v.MaxPercentUnhealthyNodes }).(pulumi.IntPtrOutput)
}

type ClusterHealthPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterHealthPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterHealthPolicyResponse)(nil)).Elem()
}

func (o ClusterHealthPolicyResponsePtrOutput) ToClusterHealthPolicyResponsePtrOutput() ClusterHealthPolicyResponsePtrOutput {
	return o
}

func (o ClusterHealthPolicyResponsePtrOutput) ToClusterHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ClusterHealthPolicyResponsePtrOutput {
	return o
}

func (o ClusterHealthPolicyResponsePtrOutput) Elem() ClusterHealthPolicyResponseOutput {
	return o.ApplyT(func(v *ClusterHealthPolicyResponse) ClusterHealthPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ClusterHealthPolicyResponse
		return ret
	}).(ClusterHealthPolicyResponseOutput)
}

func (o ClusterHealthPolicyResponsePtrOutput) ApplicationHealthPolicies() ApplicationHealthPolicyResponseMapOutput {
	return o.ApplyT(func(v *ClusterHealthPolicyResponse) map[string]ApplicationHealthPolicyResponse {
		if v == nil {
			return nil
		}
		return v.ApplicationHealthPolicies
	}).(ApplicationHealthPolicyResponseMapOutput)
}

func (o ClusterHealthPolicyResponsePtrOutput) MaxPercentUnhealthyApplications() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentUnhealthyApplications
	}).(pulumi.IntPtrOutput)
}

func (o ClusterHealthPolicyResponsePtrOutput) MaxPercentUnhealthyNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentUnhealthyNodes
	}).(pulumi.IntPtrOutput)
}

type ClusterUpgradeDeltaHealthPolicy struct {
	ApplicationDeltaHealthPolicies             map[string]ApplicationDeltaHealthPolicy `pulumi:"applicationDeltaHealthPolicies"`
	MaxPercentDeltaUnhealthyApplications       int                                     `pulumi:"maxPercentDeltaUnhealthyApplications"`
	MaxPercentDeltaUnhealthyNodes              int                                     `pulumi:"maxPercentDeltaUnhealthyNodes"`
	MaxPercentUpgradeDomainDeltaUnhealthyNodes int                                     `pulumi:"maxPercentUpgradeDomainDeltaUnhealthyNodes"`
}





type ClusterUpgradeDeltaHealthPolicyInput interface {
	pulumi.Input

	ToClusterUpgradeDeltaHealthPolicyOutput() ClusterUpgradeDeltaHealthPolicyOutput
	ToClusterUpgradeDeltaHealthPolicyOutputWithContext(context.Context) ClusterUpgradeDeltaHealthPolicyOutput
}

type ClusterUpgradeDeltaHealthPolicyArgs struct {
	ApplicationDeltaHealthPolicies             ApplicationDeltaHealthPolicyMapInput `pulumi:"applicationDeltaHealthPolicies"`
	MaxPercentDeltaUnhealthyApplications       pulumi.IntInput                      `pulumi:"maxPercentDeltaUnhealthyApplications"`
	MaxPercentDeltaUnhealthyNodes              pulumi.IntInput                      `pulumi:"maxPercentDeltaUnhealthyNodes"`
	MaxPercentUpgradeDomainDeltaUnhealthyNodes pulumi.IntInput                      `pulumi:"maxPercentUpgradeDomainDeltaUnhealthyNodes"`
}

func (ClusterUpgradeDeltaHealthPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterUpgradeDeltaHealthPolicy)(nil)).Elem()
}

func (i ClusterUpgradeDeltaHealthPolicyArgs) ToClusterUpgradeDeltaHealthPolicyOutput() ClusterUpgradeDeltaHealthPolicyOutput {
	return i.ToClusterUpgradeDeltaHealthPolicyOutputWithContext(context.Background())
}

func (i ClusterUpgradeDeltaHealthPolicyArgs) ToClusterUpgradeDeltaHealthPolicyOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterUpgradeDeltaHealthPolicyOutput)
}

func (i ClusterUpgradeDeltaHealthPolicyArgs) ToClusterUpgradeDeltaHealthPolicyPtrOutput() ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return i.ToClusterUpgradeDeltaHealthPolicyPtrOutputWithContext(context.Background())
}

func (i ClusterUpgradeDeltaHealthPolicyArgs) ToClusterUpgradeDeltaHealthPolicyPtrOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterUpgradeDeltaHealthPolicyOutput).ToClusterUpgradeDeltaHealthPolicyPtrOutputWithContext(ctx)
}









type ClusterUpgradeDeltaHealthPolicyPtrInput interface {
	pulumi.Input

	ToClusterUpgradeDeltaHealthPolicyPtrOutput() ClusterUpgradeDeltaHealthPolicyPtrOutput
	ToClusterUpgradeDeltaHealthPolicyPtrOutputWithContext(context.Context) ClusterUpgradeDeltaHealthPolicyPtrOutput
}

type clusterUpgradeDeltaHealthPolicyPtrType ClusterUpgradeDeltaHealthPolicyArgs

func ClusterUpgradeDeltaHealthPolicyPtr(v *ClusterUpgradeDeltaHealthPolicyArgs) ClusterUpgradeDeltaHealthPolicyPtrInput {
	return (*clusterUpgradeDeltaHealthPolicyPtrType)(v)
}

func (*clusterUpgradeDeltaHealthPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterUpgradeDeltaHealthPolicy)(nil)).Elem()
}

func (i *clusterUpgradeDeltaHealthPolicyPtrType) ToClusterUpgradeDeltaHealthPolicyPtrOutput() ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return i.ToClusterUpgradeDeltaHealthPolicyPtrOutputWithContext(context.Background())
}

func (i *clusterUpgradeDeltaHealthPolicyPtrType) ToClusterUpgradeDeltaHealthPolicyPtrOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterUpgradeDeltaHealthPolicyPtrOutput)
}

type ClusterUpgradeDeltaHealthPolicyOutput struct{ *pulumi.OutputState }

func (ClusterUpgradeDeltaHealthPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterUpgradeDeltaHealthPolicy)(nil)).Elem()
}

func (o ClusterUpgradeDeltaHealthPolicyOutput) ToClusterUpgradeDeltaHealthPolicyOutput() ClusterUpgradeDeltaHealthPolicyOutput {
	return o
}

func (o ClusterUpgradeDeltaHealthPolicyOutput) ToClusterUpgradeDeltaHealthPolicyOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyOutput {
	return o
}

func (o ClusterUpgradeDeltaHealthPolicyOutput) ToClusterUpgradeDeltaHealthPolicyPtrOutput() ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return o.ToClusterUpgradeDeltaHealthPolicyPtrOutputWithContext(context.Background())
}

func (o ClusterUpgradeDeltaHealthPolicyOutput) ToClusterUpgradeDeltaHealthPolicyPtrOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterUpgradeDeltaHealthPolicy) *ClusterUpgradeDeltaHealthPolicy {
		return &v
	}).(ClusterUpgradeDeltaHealthPolicyPtrOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyOutput) ApplicationDeltaHealthPolicies() ApplicationDeltaHealthPolicyMapOutput {
	return o.ApplyT(func(v ClusterUpgradeDeltaHealthPolicy) map[string]ApplicationDeltaHealthPolicy {
		return v.ApplicationDeltaHealthPolicies
	}).(ApplicationDeltaHealthPolicyMapOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyOutput) MaxPercentDeltaUnhealthyApplications() pulumi.IntOutput {
	return o.ApplyT(func(v ClusterUpgradeDeltaHealthPolicy) int { return v.MaxPercentDeltaUnhealthyApplications }).(pulumi.IntOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyOutput) MaxPercentDeltaUnhealthyNodes() pulumi.IntOutput {
	return o.ApplyT(func(v ClusterUpgradeDeltaHealthPolicy) int { return v.MaxPercentDeltaUnhealthyNodes }).(pulumi.IntOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyOutput) MaxPercentUpgradeDomainDeltaUnhealthyNodes() pulumi.IntOutput {
	return o.ApplyT(func(v ClusterUpgradeDeltaHealthPolicy) int { return v.MaxPercentUpgradeDomainDeltaUnhealthyNodes }).(pulumi.IntOutput)
}

type ClusterUpgradeDeltaHealthPolicyPtrOutput struct{ *pulumi.OutputState }

func (ClusterUpgradeDeltaHealthPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterUpgradeDeltaHealthPolicy)(nil)).Elem()
}

func (o ClusterUpgradeDeltaHealthPolicyPtrOutput) ToClusterUpgradeDeltaHealthPolicyPtrOutput() ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return o
}

func (o ClusterUpgradeDeltaHealthPolicyPtrOutput) ToClusterUpgradeDeltaHealthPolicyPtrOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return o
}

func (o ClusterUpgradeDeltaHealthPolicyPtrOutput) Elem() ClusterUpgradeDeltaHealthPolicyOutput {
	return o.ApplyT(func(v *ClusterUpgradeDeltaHealthPolicy) ClusterUpgradeDeltaHealthPolicy {
		if v != nil {
			return *v
		}
		var ret ClusterUpgradeDeltaHealthPolicy
		return ret
	}).(ClusterUpgradeDeltaHealthPolicyOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyPtrOutput) ApplicationDeltaHealthPolicies() ApplicationDeltaHealthPolicyMapOutput {
	return o.ApplyT(func(v *ClusterUpgradeDeltaHealthPolicy) map[string]ApplicationDeltaHealthPolicy {
		if v == nil {
			return nil
		}
		return v.ApplicationDeltaHealthPolicies
	}).(ApplicationDeltaHealthPolicyMapOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyPtrOutput) MaxPercentDeltaUnhealthyApplications() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradeDeltaHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentDeltaUnhealthyApplications
	}).(pulumi.IntPtrOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyPtrOutput) MaxPercentDeltaUnhealthyNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradeDeltaHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentDeltaUnhealthyNodes
	}).(pulumi.IntPtrOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyPtrOutput) MaxPercentUpgradeDomainDeltaUnhealthyNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradeDeltaHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentUpgradeDomainDeltaUnhealthyNodes
	}).(pulumi.IntPtrOutput)
}

type ClusterUpgradeDeltaHealthPolicyResponse struct {
	ApplicationDeltaHealthPolicies             map[string]ApplicationDeltaHealthPolicyResponse `pulumi:"applicationDeltaHealthPolicies"`
	MaxPercentDeltaUnhealthyApplications       int                                             `pulumi:"maxPercentDeltaUnhealthyApplications"`
	MaxPercentDeltaUnhealthyNodes              int                                             `pulumi:"maxPercentDeltaUnhealthyNodes"`
	MaxPercentUpgradeDomainDeltaUnhealthyNodes int                                             `pulumi:"maxPercentUpgradeDomainDeltaUnhealthyNodes"`
}





type ClusterUpgradeDeltaHealthPolicyResponseInput interface {
	pulumi.Input

	ToClusterUpgradeDeltaHealthPolicyResponseOutput() ClusterUpgradeDeltaHealthPolicyResponseOutput
	ToClusterUpgradeDeltaHealthPolicyResponseOutputWithContext(context.Context) ClusterUpgradeDeltaHealthPolicyResponseOutput
}

type ClusterUpgradeDeltaHealthPolicyResponseArgs struct {
	ApplicationDeltaHealthPolicies             ApplicationDeltaHealthPolicyResponseMapInput `pulumi:"applicationDeltaHealthPolicies"`
	MaxPercentDeltaUnhealthyApplications       pulumi.IntInput                              `pulumi:"maxPercentDeltaUnhealthyApplications"`
	MaxPercentDeltaUnhealthyNodes              pulumi.IntInput                              `pulumi:"maxPercentDeltaUnhealthyNodes"`
	MaxPercentUpgradeDomainDeltaUnhealthyNodes pulumi.IntInput                              `pulumi:"maxPercentUpgradeDomainDeltaUnhealthyNodes"`
}

func (ClusterUpgradeDeltaHealthPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterUpgradeDeltaHealthPolicyResponse)(nil)).Elem()
}

func (i ClusterUpgradeDeltaHealthPolicyResponseArgs) ToClusterUpgradeDeltaHealthPolicyResponseOutput() ClusterUpgradeDeltaHealthPolicyResponseOutput {
	return i.ToClusterUpgradeDeltaHealthPolicyResponseOutputWithContext(context.Background())
}

func (i ClusterUpgradeDeltaHealthPolicyResponseArgs) ToClusterUpgradeDeltaHealthPolicyResponseOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterUpgradeDeltaHealthPolicyResponseOutput)
}

func (i ClusterUpgradeDeltaHealthPolicyResponseArgs) ToClusterUpgradeDeltaHealthPolicyResponsePtrOutput() ClusterUpgradeDeltaHealthPolicyResponsePtrOutput {
	return i.ToClusterUpgradeDeltaHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (i ClusterUpgradeDeltaHealthPolicyResponseArgs) ToClusterUpgradeDeltaHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterUpgradeDeltaHealthPolicyResponseOutput).ToClusterUpgradeDeltaHealthPolicyResponsePtrOutputWithContext(ctx)
}









type ClusterUpgradeDeltaHealthPolicyResponsePtrInput interface {
	pulumi.Input

	ToClusterUpgradeDeltaHealthPolicyResponsePtrOutput() ClusterUpgradeDeltaHealthPolicyResponsePtrOutput
	ToClusterUpgradeDeltaHealthPolicyResponsePtrOutputWithContext(context.Context) ClusterUpgradeDeltaHealthPolicyResponsePtrOutput
}

type clusterUpgradeDeltaHealthPolicyResponsePtrType ClusterUpgradeDeltaHealthPolicyResponseArgs

func ClusterUpgradeDeltaHealthPolicyResponsePtr(v *ClusterUpgradeDeltaHealthPolicyResponseArgs) ClusterUpgradeDeltaHealthPolicyResponsePtrInput {
	return (*clusterUpgradeDeltaHealthPolicyResponsePtrType)(v)
}

func (*clusterUpgradeDeltaHealthPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterUpgradeDeltaHealthPolicyResponse)(nil)).Elem()
}

func (i *clusterUpgradeDeltaHealthPolicyResponsePtrType) ToClusterUpgradeDeltaHealthPolicyResponsePtrOutput() ClusterUpgradeDeltaHealthPolicyResponsePtrOutput {
	return i.ToClusterUpgradeDeltaHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *clusterUpgradeDeltaHealthPolicyResponsePtrType) ToClusterUpgradeDeltaHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterUpgradeDeltaHealthPolicyResponsePtrOutput)
}

type ClusterUpgradeDeltaHealthPolicyResponseOutput struct{ *pulumi.OutputState }

func (ClusterUpgradeDeltaHealthPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterUpgradeDeltaHealthPolicyResponse)(nil)).Elem()
}

func (o ClusterUpgradeDeltaHealthPolicyResponseOutput) ToClusterUpgradeDeltaHealthPolicyResponseOutput() ClusterUpgradeDeltaHealthPolicyResponseOutput {
	return o
}

func (o ClusterUpgradeDeltaHealthPolicyResponseOutput) ToClusterUpgradeDeltaHealthPolicyResponseOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyResponseOutput {
	return o
}

func (o ClusterUpgradeDeltaHealthPolicyResponseOutput) ToClusterUpgradeDeltaHealthPolicyResponsePtrOutput() ClusterUpgradeDeltaHealthPolicyResponsePtrOutput {
	return o.ToClusterUpgradeDeltaHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (o ClusterUpgradeDeltaHealthPolicyResponseOutput) ToClusterUpgradeDeltaHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterUpgradeDeltaHealthPolicyResponse) *ClusterUpgradeDeltaHealthPolicyResponse {
		return &v
	}).(ClusterUpgradeDeltaHealthPolicyResponsePtrOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyResponseOutput) ApplicationDeltaHealthPolicies() ApplicationDeltaHealthPolicyResponseMapOutput {
	return o.ApplyT(func(v ClusterUpgradeDeltaHealthPolicyResponse) map[string]ApplicationDeltaHealthPolicyResponse {
		return v.ApplicationDeltaHealthPolicies
	}).(ApplicationDeltaHealthPolicyResponseMapOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyResponseOutput) MaxPercentDeltaUnhealthyApplications() pulumi.IntOutput {
	return o.ApplyT(func(v ClusterUpgradeDeltaHealthPolicyResponse) int { return v.MaxPercentDeltaUnhealthyApplications }).(pulumi.IntOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyResponseOutput) MaxPercentDeltaUnhealthyNodes() pulumi.IntOutput {
	return o.ApplyT(func(v ClusterUpgradeDeltaHealthPolicyResponse) int { return v.MaxPercentDeltaUnhealthyNodes }).(pulumi.IntOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyResponseOutput) MaxPercentUpgradeDomainDeltaUnhealthyNodes() pulumi.IntOutput {
	return o.ApplyT(func(v ClusterUpgradeDeltaHealthPolicyResponse) int {
		return v.MaxPercentUpgradeDomainDeltaUnhealthyNodes
	}).(pulumi.IntOutput)
}

type ClusterUpgradeDeltaHealthPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterUpgradeDeltaHealthPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterUpgradeDeltaHealthPolicyResponse)(nil)).Elem()
}

func (o ClusterUpgradeDeltaHealthPolicyResponsePtrOutput) ToClusterUpgradeDeltaHealthPolicyResponsePtrOutput() ClusterUpgradeDeltaHealthPolicyResponsePtrOutput {
	return o
}

func (o ClusterUpgradeDeltaHealthPolicyResponsePtrOutput) ToClusterUpgradeDeltaHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyResponsePtrOutput {
	return o
}

func (o ClusterUpgradeDeltaHealthPolicyResponsePtrOutput) Elem() ClusterUpgradeDeltaHealthPolicyResponseOutput {
	return o.ApplyT(func(v *ClusterUpgradeDeltaHealthPolicyResponse) ClusterUpgradeDeltaHealthPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ClusterUpgradeDeltaHealthPolicyResponse
		return ret
	}).(ClusterUpgradeDeltaHealthPolicyResponseOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyResponsePtrOutput) ApplicationDeltaHealthPolicies() ApplicationDeltaHealthPolicyResponseMapOutput {
	return o.ApplyT(func(v *ClusterUpgradeDeltaHealthPolicyResponse) map[string]ApplicationDeltaHealthPolicyResponse {
		if v == nil {
			return nil
		}
		return v.ApplicationDeltaHealthPolicies
	}).(ApplicationDeltaHealthPolicyResponseMapOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyResponsePtrOutput) MaxPercentDeltaUnhealthyApplications() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradeDeltaHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentDeltaUnhealthyApplications
	}).(pulumi.IntPtrOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyResponsePtrOutput) MaxPercentDeltaUnhealthyNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradeDeltaHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentDeltaUnhealthyNodes
	}).(pulumi.IntPtrOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyResponsePtrOutput) MaxPercentUpgradeDomainDeltaUnhealthyNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradeDeltaHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentUpgradeDomainDeltaUnhealthyNodes
	}).(pulumi.IntPtrOutput)
}

type ClusterUpgradePolicy struct {
	DeltaHealthPolicy             *ClusterUpgradeDeltaHealthPolicy `pulumi:"deltaHealthPolicy"`
	ForceRestart                  *bool                            `pulumi:"forceRestart"`
	HealthCheckRetryTimeout       string                           `pulumi:"healthCheckRetryTimeout"`
	HealthCheckStableDuration     string                           `pulumi:"healthCheckStableDuration"`
	HealthCheckWaitDuration       string                           `pulumi:"healthCheckWaitDuration"`
	HealthPolicy                  ClusterHealthPolicy              `pulumi:"healthPolicy"`
	UpgradeDomainTimeout          string                           `pulumi:"upgradeDomainTimeout"`
	UpgradeReplicaSetCheckTimeout string                           `pulumi:"upgradeReplicaSetCheckTimeout"`
	UpgradeTimeout                string                           `pulumi:"upgradeTimeout"`
}


func (val *ClusterUpgradePolicy) Defaults() *ClusterUpgradePolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.HealthPolicy = *tmp.HealthPolicy.Defaults()

	return &tmp
}





type ClusterUpgradePolicyInput interface {
	pulumi.Input

	ToClusterUpgradePolicyOutput() ClusterUpgradePolicyOutput
	ToClusterUpgradePolicyOutputWithContext(context.Context) ClusterUpgradePolicyOutput
}

type ClusterUpgradePolicyArgs struct {
	DeltaHealthPolicy             ClusterUpgradeDeltaHealthPolicyPtrInput `pulumi:"deltaHealthPolicy"`
	ForceRestart                  pulumi.BoolPtrInput                     `pulumi:"forceRestart"`
	HealthCheckRetryTimeout       pulumi.StringInput                      `pulumi:"healthCheckRetryTimeout"`
	HealthCheckStableDuration     pulumi.StringInput                      `pulumi:"healthCheckStableDuration"`
	HealthCheckWaitDuration       pulumi.StringInput                      `pulumi:"healthCheckWaitDuration"`
	HealthPolicy                  ClusterHealthPolicyInput                `pulumi:"healthPolicy"`
	UpgradeDomainTimeout          pulumi.StringInput                      `pulumi:"upgradeDomainTimeout"`
	UpgradeReplicaSetCheckTimeout pulumi.StringInput                      `pulumi:"upgradeReplicaSetCheckTimeout"`
	UpgradeTimeout                pulumi.StringInput                      `pulumi:"upgradeTimeout"`
}

func (ClusterUpgradePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterUpgradePolicy)(nil)).Elem()
}

func (i ClusterUpgradePolicyArgs) ToClusterUpgradePolicyOutput() ClusterUpgradePolicyOutput {
	return i.ToClusterUpgradePolicyOutputWithContext(context.Background())
}

func (i ClusterUpgradePolicyArgs) ToClusterUpgradePolicyOutputWithContext(ctx context.Context) ClusterUpgradePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterUpgradePolicyOutput)
}

func (i ClusterUpgradePolicyArgs) ToClusterUpgradePolicyPtrOutput() ClusterUpgradePolicyPtrOutput {
	return i.ToClusterUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i ClusterUpgradePolicyArgs) ToClusterUpgradePolicyPtrOutputWithContext(ctx context.Context) ClusterUpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterUpgradePolicyOutput).ToClusterUpgradePolicyPtrOutputWithContext(ctx)
}









type ClusterUpgradePolicyPtrInput interface {
	pulumi.Input

	ToClusterUpgradePolicyPtrOutput() ClusterUpgradePolicyPtrOutput
	ToClusterUpgradePolicyPtrOutputWithContext(context.Context) ClusterUpgradePolicyPtrOutput
}

type clusterUpgradePolicyPtrType ClusterUpgradePolicyArgs

func ClusterUpgradePolicyPtr(v *ClusterUpgradePolicyArgs) ClusterUpgradePolicyPtrInput {
	return (*clusterUpgradePolicyPtrType)(v)
}

func (*clusterUpgradePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterUpgradePolicy)(nil)).Elem()
}

func (i *clusterUpgradePolicyPtrType) ToClusterUpgradePolicyPtrOutput() ClusterUpgradePolicyPtrOutput {
	return i.ToClusterUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i *clusterUpgradePolicyPtrType) ToClusterUpgradePolicyPtrOutputWithContext(ctx context.Context) ClusterUpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterUpgradePolicyPtrOutput)
}

type ClusterUpgradePolicyOutput struct{ *pulumi.OutputState }

func (ClusterUpgradePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterUpgradePolicy)(nil)).Elem()
}

func (o ClusterUpgradePolicyOutput) ToClusterUpgradePolicyOutput() ClusterUpgradePolicyOutput {
	return o
}

func (o ClusterUpgradePolicyOutput) ToClusterUpgradePolicyOutputWithContext(ctx context.Context) ClusterUpgradePolicyOutput {
	return o
}

func (o ClusterUpgradePolicyOutput) ToClusterUpgradePolicyPtrOutput() ClusterUpgradePolicyPtrOutput {
	return o.ToClusterUpgradePolicyPtrOutputWithContext(context.Background())
}

func (o ClusterUpgradePolicyOutput) ToClusterUpgradePolicyPtrOutputWithContext(ctx context.Context) ClusterUpgradePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterUpgradePolicy) *ClusterUpgradePolicy {
		return &v
	}).(ClusterUpgradePolicyPtrOutput)
}

func (o ClusterUpgradePolicyOutput) DeltaHealthPolicy() ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return o.ApplyT(func(v ClusterUpgradePolicy) *ClusterUpgradeDeltaHealthPolicy { return v.DeltaHealthPolicy }).(ClusterUpgradeDeltaHealthPolicyPtrOutput)
}

func (o ClusterUpgradePolicyOutput) ForceRestart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClusterUpgradePolicy) *bool { return v.ForceRestart }).(pulumi.BoolPtrOutput)
}

func (o ClusterUpgradePolicyOutput) HealthCheckRetryTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicy) string { return v.HealthCheckRetryTimeout }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyOutput) HealthCheckStableDuration() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicy) string { return v.HealthCheckStableDuration }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyOutput) HealthCheckWaitDuration() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicy) string { return v.HealthCheckWaitDuration }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyOutput) HealthPolicy() ClusterHealthPolicyOutput {
	return o.ApplyT(func(v ClusterUpgradePolicy) ClusterHealthPolicy { return v.HealthPolicy }).(ClusterHealthPolicyOutput)
}

func (o ClusterUpgradePolicyOutput) UpgradeDomainTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicy) string { return v.UpgradeDomainTimeout }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyOutput) UpgradeReplicaSetCheckTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicy) string { return v.UpgradeReplicaSetCheckTimeout }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyOutput) UpgradeTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicy) string { return v.UpgradeTimeout }).(pulumi.StringOutput)
}

type ClusterUpgradePolicyPtrOutput struct{ *pulumi.OutputState }

func (ClusterUpgradePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterUpgradePolicy)(nil)).Elem()
}

func (o ClusterUpgradePolicyPtrOutput) ToClusterUpgradePolicyPtrOutput() ClusterUpgradePolicyPtrOutput {
	return o
}

func (o ClusterUpgradePolicyPtrOutput) ToClusterUpgradePolicyPtrOutputWithContext(ctx context.Context) ClusterUpgradePolicyPtrOutput {
	return o
}

func (o ClusterUpgradePolicyPtrOutput) Elem() ClusterUpgradePolicyOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) ClusterUpgradePolicy {
		if v != nil {
			return *v
		}
		var ret ClusterUpgradePolicy
		return ret
	}).(ClusterUpgradePolicyOutput)
}

func (o ClusterUpgradePolicyPtrOutput) DeltaHealthPolicy() ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) *ClusterUpgradeDeltaHealthPolicy {
		if v == nil {
			return nil
		}
		return v.DeltaHealthPolicy
	}).(ClusterUpgradeDeltaHealthPolicyPtrOutput)
}

func (o ClusterUpgradePolicyPtrOutput) ForceRestart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) *bool {
		if v == nil {
			return nil
		}
		return v.ForceRestart
	}).(pulumi.BoolPtrOutput)
}

func (o ClusterUpgradePolicyPtrOutput) HealthCheckRetryTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) *string {
		if v == nil {
			return nil
		}
		return &v.HealthCheckRetryTimeout
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyPtrOutput) HealthCheckStableDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) *string {
		if v == nil {
			return nil
		}
		return &v.HealthCheckStableDuration
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyPtrOutput) HealthCheckWaitDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) *string {
		if v == nil {
			return nil
		}
		return &v.HealthCheckWaitDuration
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyPtrOutput) HealthPolicy() ClusterHealthPolicyPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) *ClusterHealthPolicy {
		if v == nil {
			return nil
		}
		return &v.HealthPolicy
	}).(ClusterHealthPolicyPtrOutput)
}

func (o ClusterUpgradePolicyPtrOutput) UpgradeDomainTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) *string {
		if v == nil {
			return nil
		}
		return &v.UpgradeDomainTimeout
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyPtrOutput) UpgradeReplicaSetCheckTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) *string {
		if v == nil {
			return nil
		}
		return &v.UpgradeReplicaSetCheckTimeout
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyPtrOutput) UpgradeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) *string {
		if v == nil {
			return nil
		}
		return &v.UpgradeTimeout
	}).(pulumi.StringPtrOutput)
}

type ClusterUpgradePolicyResponse struct {
	DeltaHealthPolicy             *ClusterUpgradeDeltaHealthPolicyResponse `pulumi:"deltaHealthPolicy"`
	ForceRestart                  *bool                                    `pulumi:"forceRestart"`
	HealthCheckRetryTimeout       string                                   `pulumi:"healthCheckRetryTimeout"`
	HealthCheckStableDuration     string                                   `pulumi:"healthCheckStableDuration"`
	HealthCheckWaitDuration       string                                   `pulumi:"healthCheckWaitDuration"`
	HealthPolicy                  ClusterHealthPolicyResponse              `pulumi:"healthPolicy"`
	UpgradeDomainTimeout          string                                   `pulumi:"upgradeDomainTimeout"`
	UpgradeReplicaSetCheckTimeout string                                   `pulumi:"upgradeReplicaSetCheckTimeout"`
	UpgradeTimeout                string                                   `pulumi:"upgradeTimeout"`
}


func (val *ClusterUpgradePolicyResponse) Defaults() *ClusterUpgradePolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.HealthPolicy = *tmp.HealthPolicy.Defaults()

	return &tmp
}





type ClusterUpgradePolicyResponseInput interface {
	pulumi.Input

	ToClusterUpgradePolicyResponseOutput() ClusterUpgradePolicyResponseOutput
	ToClusterUpgradePolicyResponseOutputWithContext(context.Context) ClusterUpgradePolicyResponseOutput
}

type ClusterUpgradePolicyResponseArgs struct {
	DeltaHealthPolicy             ClusterUpgradeDeltaHealthPolicyResponsePtrInput `pulumi:"deltaHealthPolicy"`
	ForceRestart                  pulumi.BoolPtrInput                             `pulumi:"forceRestart"`
	HealthCheckRetryTimeout       pulumi.StringInput                              `pulumi:"healthCheckRetryTimeout"`
	HealthCheckStableDuration     pulumi.StringInput                              `pulumi:"healthCheckStableDuration"`
	HealthCheckWaitDuration       pulumi.StringInput                              `pulumi:"healthCheckWaitDuration"`
	HealthPolicy                  ClusterHealthPolicyResponseInput                `pulumi:"healthPolicy"`
	UpgradeDomainTimeout          pulumi.StringInput                              `pulumi:"upgradeDomainTimeout"`
	UpgradeReplicaSetCheckTimeout pulumi.StringInput                              `pulumi:"upgradeReplicaSetCheckTimeout"`
	UpgradeTimeout                pulumi.StringInput                              `pulumi:"upgradeTimeout"`
}

func (ClusterUpgradePolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterUpgradePolicyResponse)(nil)).Elem()
}

func (i ClusterUpgradePolicyResponseArgs) ToClusterUpgradePolicyResponseOutput() ClusterUpgradePolicyResponseOutput {
	return i.ToClusterUpgradePolicyResponseOutputWithContext(context.Background())
}

func (i ClusterUpgradePolicyResponseArgs) ToClusterUpgradePolicyResponseOutputWithContext(ctx context.Context) ClusterUpgradePolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterUpgradePolicyResponseOutput)
}

func (i ClusterUpgradePolicyResponseArgs) ToClusterUpgradePolicyResponsePtrOutput() ClusterUpgradePolicyResponsePtrOutput {
	return i.ToClusterUpgradePolicyResponsePtrOutputWithContext(context.Background())
}

func (i ClusterUpgradePolicyResponseArgs) ToClusterUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) ClusterUpgradePolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterUpgradePolicyResponseOutput).ToClusterUpgradePolicyResponsePtrOutputWithContext(ctx)
}









type ClusterUpgradePolicyResponsePtrInput interface {
	pulumi.Input

	ToClusterUpgradePolicyResponsePtrOutput() ClusterUpgradePolicyResponsePtrOutput
	ToClusterUpgradePolicyResponsePtrOutputWithContext(context.Context) ClusterUpgradePolicyResponsePtrOutput
}

type clusterUpgradePolicyResponsePtrType ClusterUpgradePolicyResponseArgs

func ClusterUpgradePolicyResponsePtr(v *ClusterUpgradePolicyResponseArgs) ClusterUpgradePolicyResponsePtrInput {
	return (*clusterUpgradePolicyResponsePtrType)(v)
}

func (*clusterUpgradePolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterUpgradePolicyResponse)(nil)).Elem()
}

func (i *clusterUpgradePolicyResponsePtrType) ToClusterUpgradePolicyResponsePtrOutput() ClusterUpgradePolicyResponsePtrOutput {
	return i.ToClusterUpgradePolicyResponsePtrOutputWithContext(context.Background())
}

func (i *clusterUpgradePolicyResponsePtrType) ToClusterUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) ClusterUpgradePolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterUpgradePolicyResponsePtrOutput)
}

type ClusterUpgradePolicyResponseOutput struct{ *pulumi.OutputState }

func (ClusterUpgradePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterUpgradePolicyResponse)(nil)).Elem()
}

func (o ClusterUpgradePolicyResponseOutput) ToClusterUpgradePolicyResponseOutput() ClusterUpgradePolicyResponseOutput {
	return o
}

func (o ClusterUpgradePolicyResponseOutput) ToClusterUpgradePolicyResponseOutputWithContext(ctx context.Context) ClusterUpgradePolicyResponseOutput {
	return o
}

func (o ClusterUpgradePolicyResponseOutput) ToClusterUpgradePolicyResponsePtrOutput() ClusterUpgradePolicyResponsePtrOutput {
	return o.ToClusterUpgradePolicyResponsePtrOutputWithContext(context.Background())
}

func (o ClusterUpgradePolicyResponseOutput) ToClusterUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) ClusterUpgradePolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterUpgradePolicyResponse) *ClusterUpgradePolicyResponse {
		return &v
	}).(ClusterUpgradePolicyResponsePtrOutput)
}

func (o ClusterUpgradePolicyResponseOutput) DeltaHealthPolicy() ClusterUpgradeDeltaHealthPolicyResponsePtrOutput {
	return o.ApplyT(func(v ClusterUpgradePolicyResponse) *ClusterUpgradeDeltaHealthPolicyResponse {
		return v.DeltaHealthPolicy
	}).(ClusterUpgradeDeltaHealthPolicyResponsePtrOutput)
}

func (o ClusterUpgradePolicyResponseOutput) ForceRestart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClusterUpgradePolicyResponse) *bool { return v.ForceRestart }).(pulumi.BoolPtrOutput)
}

func (o ClusterUpgradePolicyResponseOutput) HealthCheckRetryTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicyResponse) string { return v.HealthCheckRetryTimeout }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyResponseOutput) HealthCheckStableDuration() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicyResponse) string { return v.HealthCheckStableDuration }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyResponseOutput) HealthCheckWaitDuration() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicyResponse) string { return v.HealthCheckWaitDuration }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyResponseOutput) HealthPolicy() ClusterHealthPolicyResponseOutput {
	return o.ApplyT(func(v ClusterUpgradePolicyResponse) ClusterHealthPolicyResponse { return v.HealthPolicy }).(ClusterHealthPolicyResponseOutput)
}

func (o ClusterUpgradePolicyResponseOutput) UpgradeDomainTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicyResponse) string { return v.UpgradeDomainTimeout }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyResponseOutput) UpgradeReplicaSetCheckTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicyResponse) string { return v.UpgradeReplicaSetCheckTimeout }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyResponseOutput) UpgradeTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicyResponse) string { return v.UpgradeTimeout }).(pulumi.StringOutput)
}

type ClusterUpgradePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterUpgradePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterUpgradePolicyResponse)(nil)).Elem()
}

func (o ClusterUpgradePolicyResponsePtrOutput) ToClusterUpgradePolicyResponsePtrOutput() ClusterUpgradePolicyResponsePtrOutput {
	return o
}

func (o ClusterUpgradePolicyResponsePtrOutput) ToClusterUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) ClusterUpgradePolicyResponsePtrOutput {
	return o
}

func (o ClusterUpgradePolicyResponsePtrOutput) Elem() ClusterUpgradePolicyResponseOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) ClusterUpgradePolicyResponse {
		if v != nil {
			return *v
		}
		var ret ClusterUpgradePolicyResponse
		return ret
	}).(ClusterUpgradePolicyResponseOutput)
}

func (o ClusterUpgradePolicyResponsePtrOutput) DeltaHealthPolicy() ClusterUpgradeDeltaHealthPolicyResponsePtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) *ClusterUpgradeDeltaHealthPolicyResponse {
		if v == nil {
			return nil
		}
		return v.DeltaHealthPolicy
	}).(ClusterUpgradeDeltaHealthPolicyResponsePtrOutput)
}

func (o ClusterUpgradePolicyResponsePtrOutput) ForceRestart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ForceRestart
	}).(pulumi.BoolPtrOutput)
}

func (o ClusterUpgradePolicyResponsePtrOutput) HealthCheckRetryTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HealthCheckRetryTimeout
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyResponsePtrOutput) HealthCheckStableDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HealthCheckStableDuration
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyResponsePtrOutput) HealthCheckWaitDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HealthCheckWaitDuration
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyResponsePtrOutput) HealthPolicy() ClusterHealthPolicyResponsePtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) *ClusterHealthPolicyResponse {
		if v == nil {
			return nil
		}
		return &v.HealthPolicy
	}).(ClusterHealthPolicyResponsePtrOutput)
}

func (o ClusterUpgradePolicyResponsePtrOutput) UpgradeDomainTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpgradeDomainTimeout
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyResponsePtrOutput) UpgradeReplicaSetCheckTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpgradeReplicaSetCheckTimeout
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyResponsePtrOutput) UpgradeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpgradeTimeout
	}).(pulumi.StringPtrOutput)
}

type ClusterVersionDetailsResponse struct {
	CodeVersion      *string `pulumi:"codeVersion"`
	Environment      *string `pulumi:"environment"`
	SupportExpiryUtc *string `pulumi:"supportExpiryUtc"`
}





type ClusterVersionDetailsResponseInput interface {
	pulumi.Input

	ToClusterVersionDetailsResponseOutput() ClusterVersionDetailsResponseOutput
	ToClusterVersionDetailsResponseOutputWithContext(context.Context) ClusterVersionDetailsResponseOutput
}

type ClusterVersionDetailsResponseArgs struct {
	CodeVersion      pulumi.StringPtrInput `pulumi:"codeVersion"`
	Environment      pulumi.StringPtrInput `pulumi:"environment"`
	SupportExpiryUtc pulumi.StringPtrInput `pulumi:"supportExpiryUtc"`
}

func (ClusterVersionDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterVersionDetailsResponse)(nil)).Elem()
}

func (i ClusterVersionDetailsResponseArgs) ToClusterVersionDetailsResponseOutput() ClusterVersionDetailsResponseOutput {
	return i.ToClusterVersionDetailsResponseOutputWithContext(context.Background())
}

func (i ClusterVersionDetailsResponseArgs) ToClusterVersionDetailsResponseOutputWithContext(ctx context.Context) ClusterVersionDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterVersionDetailsResponseOutput)
}





type ClusterVersionDetailsResponseArrayInput interface {
	pulumi.Input

	ToClusterVersionDetailsResponseArrayOutput() ClusterVersionDetailsResponseArrayOutput
	ToClusterVersionDetailsResponseArrayOutputWithContext(context.Context) ClusterVersionDetailsResponseArrayOutput
}

type ClusterVersionDetailsResponseArray []ClusterVersionDetailsResponseInput

func (ClusterVersionDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterVersionDetailsResponse)(nil)).Elem()
}

func (i ClusterVersionDetailsResponseArray) ToClusterVersionDetailsResponseArrayOutput() ClusterVersionDetailsResponseArrayOutput {
	return i.ToClusterVersionDetailsResponseArrayOutputWithContext(context.Background())
}

func (i ClusterVersionDetailsResponseArray) ToClusterVersionDetailsResponseArrayOutputWithContext(ctx context.Context) ClusterVersionDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterVersionDetailsResponseArrayOutput)
}

type ClusterVersionDetailsResponseOutput struct{ *pulumi.OutputState }

func (ClusterVersionDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterVersionDetailsResponse)(nil)).Elem()
}

func (o ClusterVersionDetailsResponseOutput) ToClusterVersionDetailsResponseOutput() ClusterVersionDetailsResponseOutput {
	return o
}

func (o ClusterVersionDetailsResponseOutput) ToClusterVersionDetailsResponseOutputWithContext(ctx context.Context) ClusterVersionDetailsResponseOutput {
	return o
}

func (o ClusterVersionDetailsResponseOutput) CodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterVersionDetailsResponse) *string { return v.CodeVersion }).(pulumi.StringPtrOutput)
}

func (o ClusterVersionDetailsResponseOutput) Environment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterVersionDetailsResponse) *string { return v.Environment }).(pulumi.StringPtrOutput)
}

func (o ClusterVersionDetailsResponseOutput) SupportExpiryUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterVersionDetailsResponse) *string { return v.SupportExpiryUtc }).(pulumi.StringPtrOutput)
}

type ClusterVersionDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (ClusterVersionDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterVersionDetailsResponse)(nil)).Elem()
}

func (o ClusterVersionDetailsResponseArrayOutput) ToClusterVersionDetailsResponseArrayOutput() ClusterVersionDetailsResponseArrayOutput {
	return o
}

func (o ClusterVersionDetailsResponseArrayOutput) ToClusterVersionDetailsResponseArrayOutputWithContext(ctx context.Context) ClusterVersionDetailsResponseArrayOutput {
	return o
}

func (o ClusterVersionDetailsResponseArrayOutput) Index(i pulumi.IntInput) ClusterVersionDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterVersionDetailsResponse {
		return vs[0].([]ClusterVersionDetailsResponse)[vs[1].(int)]
	}).(ClusterVersionDetailsResponseOutput)
}

type DiagnosticsStorageAccountConfig struct {
	BlobEndpoint             string  `pulumi:"blobEndpoint"`
	ProtectedAccountKeyName  string  `pulumi:"protectedAccountKeyName"`
	ProtectedAccountKeyName2 *string `pulumi:"protectedAccountKeyName2"`
	QueueEndpoint            string  `pulumi:"queueEndpoint"`
	StorageAccountName       string  `pulumi:"storageAccountName"`
	TableEndpoint            string  `pulumi:"tableEndpoint"`
}





type DiagnosticsStorageAccountConfigInput interface {
	pulumi.Input

	ToDiagnosticsStorageAccountConfigOutput() DiagnosticsStorageAccountConfigOutput
	ToDiagnosticsStorageAccountConfigOutputWithContext(context.Context) DiagnosticsStorageAccountConfigOutput
}

type DiagnosticsStorageAccountConfigArgs struct {
	BlobEndpoint             pulumi.StringInput    `pulumi:"blobEndpoint"`
	ProtectedAccountKeyName  pulumi.StringInput    `pulumi:"protectedAccountKeyName"`
	ProtectedAccountKeyName2 pulumi.StringPtrInput `pulumi:"protectedAccountKeyName2"`
	QueueEndpoint            pulumi.StringInput    `pulumi:"queueEndpoint"`
	StorageAccountName       pulumi.StringInput    `pulumi:"storageAccountName"`
	TableEndpoint            pulumi.StringInput    `pulumi:"tableEndpoint"`
}

func (DiagnosticsStorageAccountConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsStorageAccountConfig)(nil)).Elem()
}

func (i DiagnosticsStorageAccountConfigArgs) ToDiagnosticsStorageAccountConfigOutput() DiagnosticsStorageAccountConfigOutput {
	return i.ToDiagnosticsStorageAccountConfigOutputWithContext(context.Background())
}

func (i DiagnosticsStorageAccountConfigArgs) ToDiagnosticsStorageAccountConfigOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsStorageAccountConfigOutput)
}

func (i DiagnosticsStorageAccountConfigArgs) ToDiagnosticsStorageAccountConfigPtrOutput() DiagnosticsStorageAccountConfigPtrOutput {
	return i.ToDiagnosticsStorageAccountConfigPtrOutputWithContext(context.Background())
}

func (i DiagnosticsStorageAccountConfigArgs) ToDiagnosticsStorageAccountConfigPtrOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsStorageAccountConfigOutput).ToDiagnosticsStorageAccountConfigPtrOutputWithContext(ctx)
}









type DiagnosticsStorageAccountConfigPtrInput interface {
	pulumi.Input

	ToDiagnosticsStorageAccountConfigPtrOutput() DiagnosticsStorageAccountConfigPtrOutput
	ToDiagnosticsStorageAccountConfigPtrOutputWithContext(context.Context) DiagnosticsStorageAccountConfigPtrOutput
}

type diagnosticsStorageAccountConfigPtrType DiagnosticsStorageAccountConfigArgs

func DiagnosticsStorageAccountConfigPtr(v *DiagnosticsStorageAccountConfigArgs) DiagnosticsStorageAccountConfigPtrInput {
	return (*diagnosticsStorageAccountConfigPtrType)(v)
}

func (*diagnosticsStorageAccountConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsStorageAccountConfig)(nil)).Elem()
}

func (i *diagnosticsStorageAccountConfigPtrType) ToDiagnosticsStorageAccountConfigPtrOutput() DiagnosticsStorageAccountConfigPtrOutput {
	return i.ToDiagnosticsStorageAccountConfigPtrOutputWithContext(context.Background())
}

func (i *diagnosticsStorageAccountConfigPtrType) ToDiagnosticsStorageAccountConfigPtrOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsStorageAccountConfigPtrOutput)
}

type DiagnosticsStorageAccountConfigOutput struct{ *pulumi.OutputState }

func (DiagnosticsStorageAccountConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsStorageAccountConfig)(nil)).Elem()
}

func (o DiagnosticsStorageAccountConfigOutput) ToDiagnosticsStorageAccountConfigOutput() DiagnosticsStorageAccountConfigOutput {
	return o
}

func (o DiagnosticsStorageAccountConfigOutput) ToDiagnosticsStorageAccountConfigOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigOutput {
	return o
}

func (o DiagnosticsStorageAccountConfigOutput) ToDiagnosticsStorageAccountConfigPtrOutput() DiagnosticsStorageAccountConfigPtrOutput {
	return o.ToDiagnosticsStorageAccountConfigPtrOutputWithContext(context.Background())
}

func (o DiagnosticsStorageAccountConfigOutput) ToDiagnosticsStorageAccountConfigPtrOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiagnosticsStorageAccountConfig) *DiagnosticsStorageAccountConfig {
		return &v
	}).(DiagnosticsStorageAccountConfigPtrOutput)
}

func (o DiagnosticsStorageAccountConfigOutput) BlobEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfig) string { return v.BlobEndpoint }).(pulumi.StringOutput)
}

func (o DiagnosticsStorageAccountConfigOutput) ProtectedAccountKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfig) string { return v.ProtectedAccountKeyName }).(pulumi.StringOutput)
}

func (o DiagnosticsStorageAccountConfigOutput) ProtectedAccountKeyName2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfig) *string { return v.ProtectedAccountKeyName2 }).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigOutput) QueueEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfig) string { return v.QueueEndpoint }).(pulumi.StringOutput)
}

func (o DiagnosticsStorageAccountConfigOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfig) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o DiagnosticsStorageAccountConfigOutput) TableEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfig) string { return v.TableEndpoint }).(pulumi.StringOutput)
}

type DiagnosticsStorageAccountConfigPtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsStorageAccountConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsStorageAccountConfig)(nil)).Elem()
}

func (o DiagnosticsStorageAccountConfigPtrOutput) ToDiagnosticsStorageAccountConfigPtrOutput() DiagnosticsStorageAccountConfigPtrOutput {
	return o
}

func (o DiagnosticsStorageAccountConfigPtrOutput) ToDiagnosticsStorageAccountConfigPtrOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigPtrOutput {
	return o
}

func (o DiagnosticsStorageAccountConfigPtrOutput) Elem() DiagnosticsStorageAccountConfigOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfig) DiagnosticsStorageAccountConfig {
		if v != nil {
			return *v
		}
		var ret DiagnosticsStorageAccountConfig
		return ret
	}).(DiagnosticsStorageAccountConfigOutput)
}

func (o DiagnosticsStorageAccountConfigPtrOutput) BlobEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfig) *string {
		if v == nil {
			return nil
		}
		return &v.BlobEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigPtrOutput) ProtectedAccountKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfig) *string {
		if v == nil {
			return nil
		}
		return &v.ProtectedAccountKeyName
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigPtrOutput) ProtectedAccountKeyName2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfig) *string {
		if v == nil {
			return nil
		}
		return v.ProtectedAccountKeyName2
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigPtrOutput) QueueEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfig) *string {
		if v == nil {
			return nil
		}
		return &v.QueueEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigPtrOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfig) *string {
		if v == nil {
			return nil
		}
		return &v.StorageAccountName
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigPtrOutput) TableEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfig) *string {
		if v == nil {
			return nil
		}
		return &v.TableEndpoint
	}).(pulumi.StringPtrOutput)
}

type DiagnosticsStorageAccountConfigResponse struct {
	BlobEndpoint             string  `pulumi:"blobEndpoint"`
	ProtectedAccountKeyName  string  `pulumi:"protectedAccountKeyName"`
	ProtectedAccountKeyName2 *string `pulumi:"protectedAccountKeyName2"`
	QueueEndpoint            string  `pulumi:"queueEndpoint"`
	StorageAccountName       string  `pulumi:"storageAccountName"`
	TableEndpoint            string  `pulumi:"tableEndpoint"`
}





type DiagnosticsStorageAccountConfigResponseInput interface {
	pulumi.Input

	ToDiagnosticsStorageAccountConfigResponseOutput() DiagnosticsStorageAccountConfigResponseOutput
	ToDiagnosticsStorageAccountConfigResponseOutputWithContext(context.Context) DiagnosticsStorageAccountConfigResponseOutput
}

type DiagnosticsStorageAccountConfigResponseArgs struct {
	BlobEndpoint             pulumi.StringInput    `pulumi:"blobEndpoint"`
	ProtectedAccountKeyName  pulumi.StringInput    `pulumi:"protectedAccountKeyName"`
	ProtectedAccountKeyName2 pulumi.StringPtrInput `pulumi:"protectedAccountKeyName2"`
	QueueEndpoint            pulumi.StringInput    `pulumi:"queueEndpoint"`
	StorageAccountName       pulumi.StringInput    `pulumi:"storageAccountName"`
	TableEndpoint            pulumi.StringInput    `pulumi:"tableEndpoint"`
}

func (DiagnosticsStorageAccountConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsStorageAccountConfigResponse)(nil)).Elem()
}

func (i DiagnosticsStorageAccountConfigResponseArgs) ToDiagnosticsStorageAccountConfigResponseOutput() DiagnosticsStorageAccountConfigResponseOutput {
	return i.ToDiagnosticsStorageAccountConfigResponseOutputWithContext(context.Background())
}

func (i DiagnosticsStorageAccountConfigResponseArgs) ToDiagnosticsStorageAccountConfigResponseOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsStorageAccountConfigResponseOutput)
}

func (i DiagnosticsStorageAccountConfigResponseArgs) ToDiagnosticsStorageAccountConfigResponsePtrOutput() DiagnosticsStorageAccountConfigResponsePtrOutput {
	return i.ToDiagnosticsStorageAccountConfigResponsePtrOutputWithContext(context.Background())
}

func (i DiagnosticsStorageAccountConfigResponseArgs) ToDiagnosticsStorageAccountConfigResponsePtrOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsStorageAccountConfigResponseOutput).ToDiagnosticsStorageAccountConfigResponsePtrOutputWithContext(ctx)
}









type DiagnosticsStorageAccountConfigResponsePtrInput interface {
	pulumi.Input

	ToDiagnosticsStorageAccountConfigResponsePtrOutput() DiagnosticsStorageAccountConfigResponsePtrOutput
	ToDiagnosticsStorageAccountConfigResponsePtrOutputWithContext(context.Context) DiagnosticsStorageAccountConfigResponsePtrOutput
}

type diagnosticsStorageAccountConfigResponsePtrType DiagnosticsStorageAccountConfigResponseArgs

func DiagnosticsStorageAccountConfigResponsePtr(v *DiagnosticsStorageAccountConfigResponseArgs) DiagnosticsStorageAccountConfigResponsePtrInput {
	return (*diagnosticsStorageAccountConfigResponsePtrType)(v)
}

func (*diagnosticsStorageAccountConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsStorageAccountConfigResponse)(nil)).Elem()
}

func (i *diagnosticsStorageAccountConfigResponsePtrType) ToDiagnosticsStorageAccountConfigResponsePtrOutput() DiagnosticsStorageAccountConfigResponsePtrOutput {
	return i.ToDiagnosticsStorageAccountConfigResponsePtrOutputWithContext(context.Background())
}

func (i *diagnosticsStorageAccountConfigResponsePtrType) ToDiagnosticsStorageAccountConfigResponsePtrOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsStorageAccountConfigResponsePtrOutput)
}

type DiagnosticsStorageAccountConfigResponseOutput struct{ *pulumi.OutputState }

func (DiagnosticsStorageAccountConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsStorageAccountConfigResponse)(nil)).Elem()
}

func (o DiagnosticsStorageAccountConfigResponseOutput) ToDiagnosticsStorageAccountConfigResponseOutput() DiagnosticsStorageAccountConfigResponseOutput {
	return o
}

func (o DiagnosticsStorageAccountConfigResponseOutput) ToDiagnosticsStorageAccountConfigResponseOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigResponseOutput {
	return o
}

func (o DiagnosticsStorageAccountConfigResponseOutput) ToDiagnosticsStorageAccountConfigResponsePtrOutput() DiagnosticsStorageAccountConfigResponsePtrOutput {
	return o.ToDiagnosticsStorageAccountConfigResponsePtrOutputWithContext(context.Background())
}

func (o DiagnosticsStorageAccountConfigResponseOutput) ToDiagnosticsStorageAccountConfigResponsePtrOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiagnosticsStorageAccountConfigResponse) *DiagnosticsStorageAccountConfigResponse {
		return &v
	}).(DiagnosticsStorageAccountConfigResponsePtrOutput)
}

func (o DiagnosticsStorageAccountConfigResponseOutput) BlobEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfigResponse) string { return v.BlobEndpoint }).(pulumi.StringOutput)
}

func (o DiagnosticsStorageAccountConfigResponseOutput) ProtectedAccountKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfigResponse) string { return v.ProtectedAccountKeyName }).(pulumi.StringOutput)
}

func (o DiagnosticsStorageAccountConfigResponseOutput) ProtectedAccountKeyName2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfigResponse) *string { return v.ProtectedAccountKeyName2 }).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigResponseOutput) QueueEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfigResponse) string { return v.QueueEndpoint }).(pulumi.StringOutput)
}

func (o DiagnosticsStorageAccountConfigResponseOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfigResponse) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o DiagnosticsStorageAccountConfigResponseOutput) TableEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfigResponse) string { return v.TableEndpoint }).(pulumi.StringOutput)
}

type DiagnosticsStorageAccountConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsStorageAccountConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsStorageAccountConfigResponse)(nil)).Elem()
}

func (o DiagnosticsStorageAccountConfigResponsePtrOutput) ToDiagnosticsStorageAccountConfigResponsePtrOutput() DiagnosticsStorageAccountConfigResponsePtrOutput {
	return o
}

func (o DiagnosticsStorageAccountConfigResponsePtrOutput) ToDiagnosticsStorageAccountConfigResponsePtrOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigResponsePtrOutput {
	return o
}

func (o DiagnosticsStorageAccountConfigResponsePtrOutput) Elem() DiagnosticsStorageAccountConfigResponseOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfigResponse) DiagnosticsStorageAccountConfigResponse {
		if v != nil {
			return *v
		}
		var ret DiagnosticsStorageAccountConfigResponse
		return ret
	}).(DiagnosticsStorageAccountConfigResponseOutput)
}

func (o DiagnosticsStorageAccountConfigResponsePtrOutput) BlobEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.BlobEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigResponsePtrOutput) ProtectedAccountKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProtectedAccountKeyName
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigResponsePtrOutput) ProtectedAccountKeyName2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProtectedAccountKeyName2
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigResponsePtrOutput) QueueEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.QueueEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigResponsePtrOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StorageAccountName
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigResponsePtrOutput) TableEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TableEndpoint
	}).(pulumi.StringPtrOutput)
}

type EndpointRangeDescription struct {
	EndPort   int `pulumi:"endPort"`
	StartPort int `pulumi:"startPort"`
}





type EndpointRangeDescriptionInput interface {
	pulumi.Input

	ToEndpointRangeDescriptionOutput() EndpointRangeDescriptionOutput
	ToEndpointRangeDescriptionOutputWithContext(context.Context) EndpointRangeDescriptionOutput
}

type EndpointRangeDescriptionArgs struct {
	EndPort   pulumi.IntInput `pulumi:"endPort"`
	StartPort pulumi.IntInput `pulumi:"startPort"`
}

func (EndpointRangeDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointRangeDescription)(nil)).Elem()
}

func (i EndpointRangeDescriptionArgs) ToEndpointRangeDescriptionOutput() EndpointRangeDescriptionOutput {
	return i.ToEndpointRangeDescriptionOutputWithContext(context.Background())
}

func (i EndpointRangeDescriptionArgs) ToEndpointRangeDescriptionOutputWithContext(ctx context.Context) EndpointRangeDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRangeDescriptionOutput)
}

func (i EndpointRangeDescriptionArgs) ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput {
	return i.ToEndpointRangeDescriptionPtrOutputWithContext(context.Background())
}

func (i EndpointRangeDescriptionArgs) ToEndpointRangeDescriptionPtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRangeDescriptionOutput).ToEndpointRangeDescriptionPtrOutputWithContext(ctx)
}









type EndpointRangeDescriptionPtrInput interface {
	pulumi.Input

	ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput
	ToEndpointRangeDescriptionPtrOutputWithContext(context.Context) EndpointRangeDescriptionPtrOutput
}

type endpointRangeDescriptionPtrType EndpointRangeDescriptionArgs

func EndpointRangeDescriptionPtr(v *EndpointRangeDescriptionArgs) EndpointRangeDescriptionPtrInput {
	return (*endpointRangeDescriptionPtrType)(v)
}

func (*endpointRangeDescriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointRangeDescription)(nil)).Elem()
}

func (i *endpointRangeDescriptionPtrType) ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput {
	return i.ToEndpointRangeDescriptionPtrOutputWithContext(context.Background())
}

func (i *endpointRangeDescriptionPtrType) ToEndpointRangeDescriptionPtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRangeDescriptionPtrOutput)
}

type EndpointRangeDescriptionOutput struct{ *pulumi.OutputState }

func (EndpointRangeDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointRangeDescription)(nil)).Elem()
}

func (o EndpointRangeDescriptionOutput) ToEndpointRangeDescriptionOutput() EndpointRangeDescriptionOutput {
	return o
}

func (o EndpointRangeDescriptionOutput) ToEndpointRangeDescriptionOutputWithContext(ctx context.Context) EndpointRangeDescriptionOutput {
	return o
}

func (o EndpointRangeDescriptionOutput) ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput {
	return o.ToEndpointRangeDescriptionPtrOutputWithContext(context.Background())
}

func (o EndpointRangeDescriptionOutput) ToEndpointRangeDescriptionPtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointRangeDescription) *EndpointRangeDescription {
		return &v
	}).(EndpointRangeDescriptionPtrOutput)
}

func (o EndpointRangeDescriptionOutput) EndPort() pulumi.IntOutput {
	return o.ApplyT(func(v EndpointRangeDescription) int { return v.EndPort }).(pulumi.IntOutput)
}

func (o EndpointRangeDescriptionOutput) StartPort() pulumi.IntOutput {
	return o.ApplyT(func(v EndpointRangeDescription) int { return v.StartPort }).(pulumi.IntOutput)
}

type EndpointRangeDescriptionPtrOutput struct{ *pulumi.OutputState }

func (EndpointRangeDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointRangeDescription)(nil)).Elem()
}

func (o EndpointRangeDescriptionPtrOutput) ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput {
	return o
}

func (o EndpointRangeDescriptionPtrOutput) ToEndpointRangeDescriptionPtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionPtrOutput {
	return o
}

func (o EndpointRangeDescriptionPtrOutput) Elem() EndpointRangeDescriptionOutput {
	return o.ApplyT(func(v *EndpointRangeDescription) EndpointRangeDescription {
		if v != nil {
			return *v
		}
		var ret EndpointRangeDescription
		return ret
	}).(EndpointRangeDescriptionOutput)
}

func (o EndpointRangeDescriptionPtrOutput) EndPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EndpointRangeDescription) *int {
		if v == nil {
			return nil
		}
		return &v.EndPort
	}).(pulumi.IntPtrOutput)
}

func (o EndpointRangeDescriptionPtrOutput) StartPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EndpointRangeDescription) *int {
		if v == nil {
			return nil
		}
		return &v.StartPort
	}).(pulumi.IntPtrOutput)
}

type EndpointRangeDescriptionResponse struct {
	EndPort   int `pulumi:"endPort"`
	StartPort int `pulumi:"startPort"`
}





type EndpointRangeDescriptionResponseInput interface {
	pulumi.Input

	ToEndpointRangeDescriptionResponseOutput() EndpointRangeDescriptionResponseOutput
	ToEndpointRangeDescriptionResponseOutputWithContext(context.Context) EndpointRangeDescriptionResponseOutput
}

type EndpointRangeDescriptionResponseArgs struct {
	EndPort   pulumi.IntInput `pulumi:"endPort"`
	StartPort pulumi.IntInput `pulumi:"startPort"`
}

func (EndpointRangeDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointRangeDescriptionResponse)(nil)).Elem()
}

func (i EndpointRangeDescriptionResponseArgs) ToEndpointRangeDescriptionResponseOutput() EndpointRangeDescriptionResponseOutput {
	return i.ToEndpointRangeDescriptionResponseOutputWithContext(context.Background())
}

func (i EndpointRangeDescriptionResponseArgs) ToEndpointRangeDescriptionResponseOutputWithContext(ctx context.Context) EndpointRangeDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRangeDescriptionResponseOutput)
}

func (i EndpointRangeDescriptionResponseArgs) ToEndpointRangeDescriptionResponsePtrOutput() EndpointRangeDescriptionResponsePtrOutput {
	return i.ToEndpointRangeDescriptionResponsePtrOutputWithContext(context.Background())
}

func (i EndpointRangeDescriptionResponseArgs) ToEndpointRangeDescriptionResponsePtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRangeDescriptionResponseOutput).ToEndpointRangeDescriptionResponsePtrOutputWithContext(ctx)
}









type EndpointRangeDescriptionResponsePtrInput interface {
	pulumi.Input

	ToEndpointRangeDescriptionResponsePtrOutput() EndpointRangeDescriptionResponsePtrOutput
	ToEndpointRangeDescriptionResponsePtrOutputWithContext(context.Context) EndpointRangeDescriptionResponsePtrOutput
}

type endpointRangeDescriptionResponsePtrType EndpointRangeDescriptionResponseArgs

func EndpointRangeDescriptionResponsePtr(v *EndpointRangeDescriptionResponseArgs) EndpointRangeDescriptionResponsePtrInput {
	return (*endpointRangeDescriptionResponsePtrType)(v)
}

func (*endpointRangeDescriptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointRangeDescriptionResponse)(nil)).Elem()
}

func (i *endpointRangeDescriptionResponsePtrType) ToEndpointRangeDescriptionResponsePtrOutput() EndpointRangeDescriptionResponsePtrOutput {
	return i.ToEndpointRangeDescriptionResponsePtrOutputWithContext(context.Background())
}

func (i *endpointRangeDescriptionResponsePtrType) ToEndpointRangeDescriptionResponsePtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRangeDescriptionResponsePtrOutput)
}

type EndpointRangeDescriptionResponseOutput struct{ *pulumi.OutputState }

func (EndpointRangeDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointRangeDescriptionResponse)(nil)).Elem()
}

func (o EndpointRangeDescriptionResponseOutput) ToEndpointRangeDescriptionResponseOutput() EndpointRangeDescriptionResponseOutput {
	return o
}

func (o EndpointRangeDescriptionResponseOutput) ToEndpointRangeDescriptionResponseOutputWithContext(ctx context.Context) EndpointRangeDescriptionResponseOutput {
	return o
}

func (o EndpointRangeDescriptionResponseOutput) ToEndpointRangeDescriptionResponsePtrOutput() EndpointRangeDescriptionResponsePtrOutput {
	return o.ToEndpointRangeDescriptionResponsePtrOutputWithContext(context.Background())
}

func (o EndpointRangeDescriptionResponseOutput) ToEndpointRangeDescriptionResponsePtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointRangeDescriptionResponse) *EndpointRangeDescriptionResponse {
		return &v
	}).(EndpointRangeDescriptionResponsePtrOutput)
}

func (o EndpointRangeDescriptionResponseOutput) EndPort() pulumi.IntOutput {
	return o.ApplyT(func(v EndpointRangeDescriptionResponse) int { return v.EndPort }).(pulumi.IntOutput)
}

func (o EndpointRangeDescriptionResponseOutput) StartPort() pulumi.IntOutput {
	return o.ApplyT(func(v EndpointRangeDescriptionResponse) int { return v.StartPort }).(pulumi.IntOutput)
}

type EndpointRangeDescriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EndpointRangeDescriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointRangeDescriptionResponse)(nil)).Elem()
}

func (o EndpointRangeDescriptionResponsePtrOutput) ToEndpointRangeDescriptionResponsePtrOutput() EndpointRangeDescriptionResponsePtrOutput {
	return o
}

func (o EndpointRangeDescriptionResponsePtrOutput) ToEndpointRangeDescriptionResponsePtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionResponsePtrOutput {
	return o
}

func (o EndpointRangeDescriptionResponsePtrOutput) Elem() EndpointRangeDescriptionResponseOutput {
	return o.ApplyT(func(v *EndpointRangeDescriptionResponse) EndpointRangeDescriptionResponse {
		if v != nil {
			return *v
		}
		var ret EndpointRangeDescriptionResponse
		return ret
	}).(EndpointRangeDescriptionResponseOutput)
}

func (o EndpointRangeDescriptionResponsePtrOutput) EndPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EndpointRangeDescriptionResponse) *int {
		if v == nil {
			return nil
		}
		return &v.EndPort
	}).(pulumi.IntPtrOutput)
}

func (o EndpointRangeDescriptionResponsePtrOutput) StartPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EndpointRangeDescriptionResponse) *int {
		if v == nil {
			return nil
		}
		return &v.StartPort
	}).(pulumi.IntPtrOutput)
}

type LoadBalancingRule struct {
	BackendPort      int     `pulumi:"backendPort"`
	FrontendPort     int     `pulumi:"frontendPort"`
	ProbeProtocol    string  `pulumi:"probeProtocol"`
	ProbeRequestPath *string `pulumi:"probeRequestPath"`
	Protocol         string  `pulumi:"protocol"`
}





type LoadBalancingRuleInput interface {
	pulumi.Input

	ToLoadBalancingRuleOutput() LoadBalancingRuleOutput
	ToLoadBalancingRuleOutputWithContext(context.Context) LoadBalancingRuleOutput
}

type LoadBalancingRuleArgs struct {
	BackendPort      pulumi.IntInput       `pulumi:"backendPort"`
	FrontendPort     pulumi.IntInput       `pulumi:"frontendPort"`
	ProbeProtocol    pulumi.StringInput    `pulumi:"probeProtocol"`
	ProbeRequestPath pulumi.StringPtrInput `pulumi:"probeRequestPath"`
	Protocol         pulumi.StringInput    `pulumi:"protocol"`
}

func (LoadBalancingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingRule)(nil)).Elem()
}

func (i LoadBalancingRuleArgs) ToLoadBalancingRuleOutput() LoadBalancingRuleOutput {
	return i.ToLoadBalancingRuleOutputWithContext(context.Background())
}

func (i LoadBalancingRuleArgs) ToLoadBalancingRuleOutputWithContext(ctx context.Context) LoadBalancingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingRuleOutput)
}





type LoadBalancingRuleArrayInput interface {
	pulumi.Input

	ToLoadBalancingRuleArrayOutput() LoadBalancingRuleArrayOutput
	ToLoadBalancingRuleArrayOutputWithContext(context.Context) LoadBalancingRuleArrayOutput
}

type LoadBalancingRuleArray []LoadBalancingRuleInput

func (LoadBalancingRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingRule)(nil)).Elem()
}

func (i LoadBalancingRuleArray) ToLoadBalancingRuleArrayOutput() LoadBalancingRuleArrayOutput {
	return i.ToLoadBalancingRuleArrayOutputWithContext(context.Background())
}

func (i LoadBalancingRuleArray) ToLoadBalancingRuleArrayOutputWithContext(ctx context.Context) LoadBalancingRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingRuleArrayOutput)
}

type LoadBalancingRuleOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingRule)(nil)).Elem()
}

func (o LoadBalancingRuleOutput) ToLoadBalancingRuleOutput() LoadBalancingRuleOutput {
	return o
}

func (o LoadBalancingRuleOutput) ToLoadBalancingRuleOutputWithContext(ctx context.Context) LoadBalancingRuleOutput {
	return o
}

func (o LoadBalancingRuleOutput) BackendPort() pulumi.IntOutput {
	return o.ApplyT(func(v LoadBalancingRule) int { return v.BackendPort }).(pulumi.IntOutput)
}

func (o LoadBalancingRuleOutput) FrontendPort() pulumi.IntOutput {
	return o.ApplyT(func(v LoadBalancingRule) int { return v.FrontendPort }).(pulumi.IntOutput)
}

func (o LoadBalancingRuleOutput) ProbeProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingRule) string { return v.ProbeProtocol }).(pulumi.StringOutput)
}

func (o LoadBalancingRuleOutput) ProbeRequestPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *string { return v.ProbeRequestPath }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingRule) string { return v.Protocol }).(pulumi.StringOutput)
}

type LoadBalancingRuleArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingRule)(nil)).Elem()
}

func (o LoadBalancingRuleArrayOutput) ToLoadBalancingRuleArrayOutput() LoadBalancingRuleArrayOutput {
	return o
}

func (o LoadBalancingRuleArrayOutput) ToLoadBalancingRuleArrayOutputWithContext(ctx context.Context) LoadBalancingRuleArrayOutput {
	return o
}

func (o LoadBalancingRuleArrayOutput) Index(i pulumi.IntInput) LoadBalancingRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancingRule {
		return vs[0].([]LoadBalancingRule)[vs[1].(int)]
	}).(LoadBalancingRuleOutput)
}

type LoadBalancingRuleResponse struct {
	BackendPort      int     `pulumi:"backendPort"`
	FrontendPort     int     `pulumi:"frontendPort"`
	ProbeProtocol    string  `pulumi:"probeProtocol"`
	ProbeRequestPath *string `pulumi:"probeRequestPath"`
	Protocol         string  `pulumi:"protocol"`
}





type LoadBalancingRuleResponseInput interface {
	pulumi.Input

	ToLoadBalancingRuleResponseOutput() LoadBalancingRuleResponseOutput
	ToLoadBalancingRuleResponseOutputWithContext(context.Context) LoadBalancingRuleResponseOutput
}

type LoadBalancingRuleResponseArgs struct {
	BackendPort      pulumi.IntInput       `pulumi:"backendPort"`
	FrontendPort     pulumi.IntInput       `pulumi:"frontendPort"`
	ProbeProtocol    pulumi.StringInput    `pulumi:"probeProtocol"`
	ProbeRequestPath pulumi.StringPtrInput `pulumi:"probeRequestPath"`
	Protocol         pulumi.StringInput    `pulumi:"protocol"`
}

func (LoadBalancingRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingRuleResponse)(nil)).Elem()
}

func (i LoadBalancingRuleResponseArgs) ToLoadBalancingRuleResponseOutput() LoadBalancingRuleResponseOutput {
	return i.ToLoadBalancingRuleResponseOutputWithContext(context.Background())
}

func (i LoadBalancingRuleResponseArgs) ToLoadBalancingRuleResponseOutputWithContext(ctx context.Context) LoadBalancingRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingRuleResponseOutput)
}





type LoadBalancingRuleResponseArrayInput interface {
	pulumi.Input

	ToLoadBalancingRuleResponseArrayOutput() LoadBalancingRuleResponseArrayOutput
	ToLoadBalancingRuleResponseArrayOutputWithContext(context.Context) LoadBalancingRuleResponseArrayOutput
}

type LoadBalancingRuleResponseArray []LoadBalancingRuleResponseInput

func (LoadBalancingRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingRuleResponse)(nil)).Elem()
}

func (i LoadBalancingRuleResponseArray) ToLoadBalancingRuleResponseArrayOutput() LoadBalancingRuleResponseArrayOutput {
	return i.ToLoadBalancingRuleResponseArrayOutputWithContext(context.Background())
}

func (i LoadBalancingRuleResponseArray) ToLoadBalancingRuleResponseArrayOutputWithContext(ctx context.Context) LoadBalancingRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingRuleResponseArrayOutput)
}

type LoadBalancingRuleResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingRuleResponse)(nil)).Elem()
}

func (o LoadBalancingRuleResponseOutput) ToLoadBalancingRuleResponseOutput() LoadBalancingRuleResponseOutput {
	return o
}

func (o LoadBalancingRuleResponseOutput) ToLoadBalancingRuleResponseOutputWithContext(ctx context.Context) LoadBalancingRuleResponseOutput {
	return o
}

func (o LoadBalancingRuleResponseOutput) BackendPort() pulumi.IntOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) int { return v.BackendPort }).(pulumi.IntOutput)
}

func (o LoadBalancingRuleResponseOutput) FrontendPort() pulumi.IntOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) int { return v.FrontendPort }).(pulumi.IntOutput)
}

func (o LoadBalancingRuleResponseOutput) ProbeProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) string { return v.ProbeProtocol }).(pulumi.StringOutput)
}

func (o LoadBalancingRuleResponseOutput) ProbeRequestPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *string { return v.ProbeRequestPath }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

type LoadBalancingRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingRuleResponse)(nil)).Elem()
}

func (o LoadBalancingRuleResponseArrayOutput) ToLoadBalancingRuleResponseArrayOutput() LoadBalancingRuleResponseArrayOutput {
	return o
}

func (o LoadBalancingRuleResponseArrayOutput) ToLoadBalancingRuleResponseArrayOutputWithContext(ctx context.Context) LoadBalancingRuleResponseArrayOutput {
	return o
}

func (o LoadBalancingRuleResponseArrayOutput) Index(i pulumi.IntInput) LoadBalancingRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancingRuleResponse {
		return vs[0].([]LoadBalancingRuleResponse)[vs[1].(int)]
	}).(LoadBalancingRuleResponseOutput)
}

type ManagedIdentity struct {
	Type                   *ManagedIdentityType   `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ManagedIdentityInput interface {
	pulumi.Input

	ToManagedIdentityOutput() ManagedIdentityOutput
	ToManagedIdentityOutputWithContext(context.Context) ManagedIdentityOutput
}

type ManagedIdentityArgs struct {
	Type                   ManagedIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput             `pulumi:"userAssignedIdentities"`
}

func (ManagedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentity)(nil)).Elem()
}

func (i ManagedIdentityArgs) ToManagedIdentityOutput() ManagedIdentityOutput {
	return i.ToManagedIdentityOutputWithContext(context.Background())
}

func (i ManagedIdentityArgs) ToManagedIdentityOutputWithContext(ctx context.Context) ManagedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityOutput)
}

func (i ManagedIdentityArgs) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return i.ToManagedIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedIdentityArgs) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityOutput).ToManagedIdentityPtrOutputWithContext(ctx)
}









type ManagedIdentityPtrInput interface {
	pulumi.Input

	ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput
	ToManagedIdentityPtrOutputWithContext(context.Context) ManagedIdentityPtrOutput
}

type managedIdentityPtrType ManagedIdentityArgs

func ManagedIdentityPtr(v *ManagedIdentityArgs) ManagedIdentityPtrInput {
	return (*managedIdentityPtrType)(v)
}

func (*managedIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentity)(nil)).Elem()
}

func (i *managedIdentityPtrType) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return i.ToManagedIdentityPtrOutputWithContext(context.Background())
}

func (i *managedIdentityPtrType) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityPtrOutput)
}

type ManagedIdentityOutput struct{ *pulumi.OutputState }

func (ManagedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentity)(nil)).Elem()
}

func (o ManagedIdentityOutput) ToManagedIdentityOutput() ManagedIdentityOutput {
	return o
}

func (o ManagedIdentityOutput) ToManagedIdentityOutputWithContext(ctx context.Context) ManagedIdentityOutput {
	return o
}

func (o ManagedIdentityOutput) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return o.ToManagedIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedIdentityOutput) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentity) *ManagedIdentity {
		return &v
	}).(ManagedIdentityPtrOutput)
}

func (o ManagedIdentityOutput) Type() ManagedIdentityTypePtrOutput {
	return o.ApplyT(func(v ManagedIdentity) *ManagedIdentityType { return v.Type }).(ManagedIdentityTypePtrOutput)
}

func (o ManagedIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ManagedIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ManagedIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentity)(nil)).Elem()
}

func (o ManagedIdentityPtrOutput) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return o
}

func (o ManagedIdentityPtrOutput) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return o
}

func (o ManagedIdentityPtrOutput) Elem() ManagedIdentityOutput {
	return o.ApplyT(func(v *ManagedIdentity) ManagedIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedIdentity
		return ret
	}).(ManagedIdentityOutput)
}

func (o ManagedIdentityPtrOutput) Type() ManagedIdentityTypePtrOutput {
	return o.ApplyT(func(v *ManagedIdentity) *ManagedIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ManagedIdentityTypePtrOutput)
}

func (o ManagedIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ManagedIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ManagedIdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   *string                                 `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}





type ManagedIdentityResponseInput interface {
	pulumi.Input

	ToManagedIdentityResponseOutput() ManagedIdentityResponseOutput
	ToManagedIdentityResponseOutputWithContext(context.Context) ManagedIdentityResponseOutput
}

type ManagedIdentityResponseArgs struct {
	PrincipalId            pulumi.StringInput                   `pulumi:"principalId"`
	TenantId               pulumi.StringInput                   `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput                `pulumi:"type"`
	UserAssignedIdentities UserAssignedIdentityResponseMapInput `pulumi:"userAssignedIdentities"`
}

func (ManagedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityResponse)(nil)).Elem()
}

func (i ManagedIdentityResponseArgs) ToManagedIdentityResponseOutput() ManagedIdentityResponseOutput {
	return i.ToManagedIdentityResponseOutputWithContext(context.Background())
}

func (i ManagedIdentityResponseArgs) ToManagedIdentityResponseOutputWithContext(ctx context.Context) ManagedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityResponseOutput)
}

func (i ManagedIdentityResponseArgs) ToManagedIdentityResponsePtrOutput() ManagedIdentityResponsePtrOutput {
	return i.ToManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ManagedIdentityResponseArgs) ToManagedIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityResponseOutput).ToManagedIdentityResponsePtrOutputWithContext(ctx)
}









type ManagedIdentityResponsePtrInput interface {
	pulumi.Input

	ToManagedIdentityResponsePtrOutput() ManagedIdentityResponsePtrOutput
	ToManagedIdentityResponsePtrOutputWithContext(context.Context) ManagedIdentityResponsePtrOutput
}

type managedIdentityResponsePtrType ManagedIdentityResponseArgs

func ManagedIdentityResponsePtr(v *ManagedIdentityResponseArgs) ManagedIdentityResponsePtrInput {
	return (*managedIdentityResponsePtrType)(v)
}

func (*managedIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityResponse)(nil)).Elem()
}

func (i *managedIdentityResponsePtrType) ToManagedIdentityResponsePtrOutput() ManagedIdentityResponsePtrOutput {
	return i.ToManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *managedIdentityResponsePtrType) ToManagedIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityResponsePtrOutput)
}

type ManagedIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityResponse)(nil)).Elem()
}

func (o ManagedIdentityResponseOutput) ToManagedIdentityResponseOutput() ManagedIdentityResponseOutput {
	return o
}

func (o ManagedIdentityResponseOutput) ToManagedIdentityResponseOutputWithContext(ctx context.Context) ManagedIdentityResponseOutput {
	return o
}

func (o ManagedIdentityResponseOutput) ToManagedIdentityResponsePtrOutput() ManagedIdentityResponsePtrOutput {
	return o.ToManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ManagedIdentityResponseOutput) ToManagedIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentityResponse) *ManagedIdentityResponse {
		return &v
	}).(ManagedIdentityResponsePtrOutput)
}

func (o ManagedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ManagedIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityResponse)(nil)).Elem()
}

func (o ManagedIdentityResponsePtrOutput) ToManagedIdentityResponsePtrOutput() ManagedIdentityResponsePtrOutput {
	return o
}

func (o ManagedIdentityResponsePtrOutput) ToManagedIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityResponsePtrOutput {
	return o
}

func (o ManagedIdentityResponsePtrOutput) Elem() ManagedIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) ManagedIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedIdentityResponse
		return ret
	}).(ManagedIdentityResponseOutput)
}

func (o ManagedIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type NamedPartitionSchemeDescription struct {
	Count           int      `pulumi:"count"`
	Names           []string `pulumi:"names"`
	PartitionScheme string   `pulumi:"partitionScheme"`
}





type NamedPartitionSchemeDescriptionInput interface {
	pulumi.Input

	ToNamedPartitionSchemeDescriptionOutput() NamedPartitionSchemeDescriptionOutput
	ToNamedPartitionSchemeDescriptionOutputWithContext(context.Context) NamedPartitionSchemeDescriptionOutput
}

type NamedPartitionSchemeDescriptionArgs struct {
	Count           pulumi.IntInput         `pulumi:"count"`
	Names           pulumi.StringArrayInput `pulumi:"names"`
	PartitionScheme pulumi.StringInput      `pulumi:"partitionScheme"`
}

func (NamedPartitionSchemeDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NamedPartitionSchemeDescription)(nil)).Elem()
}

func (i NamedPartitionSchemeDescriptionArgs) ToNamedPartitionSchemeDescriptionOutput() NamedPartitionSchemeDescriptionOutput {
	return i.ToNamedPartitionSchemeDescriptionOutputWithContext(context.Background())
}

func (i NamedPartitionSchemeDescriptionArgs) ToNamedPartitionSchemeDescriptionOutputWithContext(ctx context.Context) NamedPartitionSchemeDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedPartitionSchemeDescriptionOutput)
}

type NamedPartitionSchemeDescriptionOutput struct{ *pulumi.OutputState }

func (NamedPartitionSchemeDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamedPartitionSchemeDescription)(nil)).Elem()
}

func (o NamedPartitionSchemeDescriptionOutput) ToNamedPartitionSchemeDescriptionOutput() NamedPartitionSchemeDescriptionOutput {
	return o
}

func (o NamedPartitionSchemeDescriptionOutput) ToNamedPartitionSchemeDescriptionOutputWithContext(ctx context.Context) NamedPartitionSchemeDescriptionOutput {
	return o
}

func (o NamedPartitionSchemeDescriptionOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v NamedPartitionSchemeDescription) int { return v.Count }).(pulumi.IntOutput)
}

func (o NamedPartitionSchemeDescriptionOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NamedPartitionSchemeDescription) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o NamedPartitionSchemeDescriptionOutput) PartitionScheme() pulumi.StringOutput {
	return o.ApplyT(func(v NamedPartitionSchemeDescription) string { return v.PartitionScheme }).(pulumi.StringOutput)
}

type NamedPartitionSchemeDescriptionResponse struct {
	Count           int      `pulumi:"count"`
	Names           []string `pulumi:"names"`
	PartitionScheme string   `pulumi:"partitionScheme"`
}





type NamedPartitionSchemeDescriptionResponseInput interface {
	pulumi.Input

	ToNamedPartitionSchemeDescriptionResponseOutput() NamedPartitionSchemeDescriptionResponseOutput
	ToNamedPartitionSchemeDescriptionResponseOutputWithContext(context.Context) NamedPartitionSchemeDescriptionResponseOutput
}

type NamedPartitionSchemeDescriptionResponseArgs struct {
	Count           pulumi.IntInput         `pulumi:"count"`
	Names           pulumi.StringArrayInput `pulumi:"names"`
	PartitionScheme pulumi.StringInput      `pulumi:"partitionScheme"`
}

func (NamedPartitionSchemeDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NamedPartitionSchemeDescriptionResponse)(nil)).Elem()
}

func (i NamedPartitionSchemeDescriptionResponseArgs) ToNamedPartitionSchemeDescriptionResponseOutput() NamedPartitionSchemeDescriptionResponseOutput {
	return i.ToNamedPartitionSchemeDescriptionResponseOutputWithContext(context.Background())
}

func (i NamedPartitionSchemeDescriptionResponseArgs) ToNamedPartitionSchemeDescriptionResponseOutputWithContext(ctx context.Context) NamedPartitionSchemeDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedPartitionSchemeDescriptionResponseOutput)
}

type NamedPartitionSchemeDescriptionResponseOutput struct{ *pulumi.OutputState }

func (NamedPartitionSchemeDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamedPartitionSchemeDescriptionResponse)(nil)).Elem()
}

func (o NamedPartitionSchemeDescriptionResponseOutput) ToNamedPartitionSchemeDescriptionResponseOutput() NamedPartitionSchemeDescriptionResponseOutput {
	return o
}

func (o NamedPartitionSchemeDescriptionResponseOutput) ToNamedPartitionSchemeDescriptionResponseOutputWithContext(ctx context.Context) NamedPartitionSchemeDescriptionResponseOutput {
	return o
}

func (o NamedPartitionSchemeDescriptionResponseOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v NamedPartitionSchemeDescriptionResponse) int { return v.Count }).(pulumi.IntOutput)
}

func (o NamedPartitionSchemeDescriptionResponseOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NamedPartitionSchemeDescriptionResponse) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o NamedPartitionSchemeDescriptionResponseOutput) PartitionScheme() pulumi.StringOutput {
	return o.ApplyT(func(v NamedPartitionSchemeDescriptionResponse) string { return v.PartitionScheme }).(pulumi.StringOutput)
}

type NodeTypeDescription struct {
	ApplicationPorts             *EndpointRangeDescription `pulumi:"applicationPorts"`
	Capacities                   map[string]string         `pulumi:"capacities"`
	ClientConnectionEndpointPort int                       `pulumi:"clientConnectionEndpointPort"`
	DurabilityLevel              *string                   `pulumi:"durabilityLevel"`
	EphemeralPorts               *EndpointRangeDescription `pulumi:"ephemeralPorts"`
	HttpGatewayEndpointPort      int                       `pulumi:"httpGatewayEndpointPort"`
	IsPrimary                    bool                      `pulumi:"isPrimary"`
	Name                         string                    `pulumi:"name"`
	PlacementProperties          map[string]string         `pulumi:"placementProperties"`
	ReverseProxyEndpointPort     *int                      `pulumi:"reverseProxyEndpointPort"`
	VmInstanceCount              int                       `pulumi:"vmInstanceCount"`
}





type NodeTypeDescriptionInput interface {
	pulumi.Input

	ToNodeTypeDescriptionOutput() NodeTypeDescriptionOutput
	ToNodeTypeDescriptionOutputWithContext(context.Context) NodeTypeDescriptionOutput
}

type NodeTypeDescriptionArgs struct {
	ApplicationPorts             EndpointRangeDescriptionPtrInput `pulumi:"applicationPorts"`
	Capacities                   pulumi.StringMapInput            `pulumi:"capacities"`
	ClientConnectionEndpointPort pulumi.IntInput                  `pulumi:"clientConnectionEndpointPort"`
	DurabilityLevel              pulumi.StringPtrInput            `pulumi:"durabilityLevel"`
	EphemeralPorts               EndpointRangeDescriptionPtrInput `pulumi:"ephemeralPorts"`
	HttpGatewayEndpointPort      pulumi.IntInput                  `pulumi:"httpGatewayEndpointPort"`
	IsPrimary                    pulumi.BoolInput                 `pulumi:"isPrimary"`
	Name                         pulumi.StringInput               `pulumi:"name"`
	PlacementProperties          pulumi.StringMapInput            `pulumi:"placementProperties"`
	ReverseProxyEndpointPort     pulumi.IntPtrInput               `pulumi:"reverseProxyEndpointPort"`
	VmInstanceCount              pulumi.IntInput                  `pulumi:"vmInstanceCount"`
}

func (NodeTypeDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeTypeDescription)(nil)).Elem()
}

func (i NodeTypeDescriptionArgs) ToNodeTypeDescriptionOutput() NodeTypeDescriptionOutput {
	return i.ToNodeTypeDescriptionOutputWithContext(context.Background())
}

func (i NodeTypeDescriptionArgs) ToNodeTypeDescriptionOutputWithContext(ctx context.Context) NodeTypeDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeTypeDescriptionOutput)
}





type NodeTypeDescriptionArrayInput interface {
	pulumi.Input

	ToNodeTypeDescriptionArrayOutput() NodeTypeDescriptionArrayOutput
	ToNodeTypeDescriptionArrayOutputWithContext(context.Context) NodeTypeDescriptionArrayOutput
}

type NodeTypeDescriptionArray []NodeTypeDescriptionInput

func (NodeTypeDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeTypeDescription)(nil)).Elem()
}

func (i NodeTypeDescriptionArray) ToNodeTypeDescriptionArrayOutput() NodeTypeDescriptionArrayOutput {
	return i.ToNodeTypeDescriptionArrayOutputWithContext(context.Background())
}

func (i NodeTypeDescriptionArray) ToNodeTypeDescriptionArrayOutputWithContext(ctx context.Context) NodeTypeDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeTypeDescriptionArrayOutput)
}

type NodeTypeDescriptionOutput struct{ *pulumi.OutputState }

func (NodeTypeDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeTypeDescription)(nil)).Elem()
}

func (o NodeTypeDescriptionOutput) ToNodeTypeDescriptionOutput() NodeTypeDescriptionOutput {
	return o
}

func (o NodeTypeDescriptionOutput) ToNodeTypeDescriptionOutputWithContext(ctx context.Context) NodeTypeDescriptionOutput {
	return o
}

func (o NodeTypeDescriptionOutput) ApplicationPorts() EndpointRangeDescriptionPtrOutput {
	return o.ApplyT(func(v NodeTypeDescription) *EndpointRangeDescription { return v.ApplicationPorts }).(EndpointRangeDescriptionPtrOutput)
}

func (o NodeTypeDescriptionOutput) Capacities() pulumi.StringMapOutput {
	return o.ApplyT(func(v NodeTypeDescription) map[string]string { return v.Capacities }).(pulumi.StringMapOutput)
}

func (o NodeTypeDescriptionOutput) ClientConnectionEndpointPort() pulumi.IntOutput {
	return o.ApplyT(func(v NodeTypeDescription) int { return v.ClientConnectionEndpointPort }).(pulumi.IntOutput)
}

func (o NodeTypeDescriptionOutput) DurabilityLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodeTypeDescription) *string { return v.DurabilityLevel }).(pulumi.StringPtrOutput)
}

func (o NodeTypeDescriptionOutput) EphemeralPorts() EndpointRangeDescriptionPtrOutput {
	return o.ApplyT(func(v NodeTypeDescription) *EndpointRangeDescription { return v.EphemeralPorts }).(EndpointRangeDescriptionPtrOutput)
}

func (o NodeTypeDescriptionOutput) HttpGatewayEndpointPort() pulumi.IntOutput {
	return o.ApplyT(func(v NodeTypeDescription) int { return v.HttpGatewayEndpointPort }).(pulumi.IntOutput)
}

func (o NodeTypeDescriptionOutput) IsPrimary() pulumi.BoolOutput {
	return o.ApplyT(func(v NodeTypeDescription) bool { return v.IsPrimary }).(pulumi.BoolOutput)
}

func (o NodeTypeDescriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v NodeTypeDescription) string { return v.Name }).(pulumi.StringOutput)
}

func (o NodeTypeDescriptionOutput) PlacementProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v NodeTypeDescription) map[string]string { return v.PlacementProperties }).(pulumi.StringMapOutput)
}

func (o NodeTypeDescriptionOutput) ReverseProxyEndpointPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NodeTypeDescription) *int { return v.ReverseProxyEndpointPort }).(pulumi.IntPtrOutput)
}

func (o NodeTypeDescriptionOutput) VmInstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodeTypeDescription) int { return v.VmInstanceCount }).(pulumi.IntOutput)
}

type NodeTypeDescriptionArrayOutput struct{ *pulumi.OutputState }

func (NodeTypeDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeTypeDescription)(nil)).Elem()
}

func (o NodeTypeDescriptionArrayOutput) ToNodeTypeDescriptionArrayOutput() NodeTypeDescriptionArrayOutput {
	return o
}

func (o NodeTypeDescriptionArrayOutput) ToNodeTypeDescriptionArrayOutputWithContext(ctx context.Context) NodeTypeDescriptionArrayOutput {
	return o
}

func (o NodeTypeDescriptionArrayOutput) Index(i pulumi.IntInput) NodeTypeDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NodeTypeDescription {
		return vs[0].([]NodeTypeDescription)[vs[1].(int)]
	}).(NodeTypeDescriptionOutput)
}

type NodeTypeDescriptionResponse struct {
	ApplicationPorts             *EndpointRangeDescriptionResponse `pulumi:"applicationPorts"`
	Capacities                   map[string]string                 `pulumi:"capacities"`
	ClientConnectionEndpointPort int                               `pulumi:"clientConnectionEndpointPort"`
	DurabilityLevel              *string                           `pulumi:"durabilityLevel"`
	EphemeralPorts               *EndpointRangeDescriptionResponse `pulumi:"ephemeralPorts"`
	HttpGatewayEndpointPort      int                               `pulumi:"httpGatewayEndpointPort"`
	IsPrimary                    bool                              `pulumi:"isPrimary"`
	Name                         string                            `pulumi:"name"`
	PlacementProperties          map[string]string                 `pulumi:"placementProperties"`
	ReverseProxyEndpointPort     *int                              `pulumi:"reverseProxyEndpointPort"`
	VmInstanceCount              int                               `pulumi:"vmInstanceCount"`
}





type NodeTypeDescriptionResponseInput interface {
	pulumi.Input

	ToNodeTypeDescriptionResponseOutput() NodeTypeDescriptionResponseOutput
	ToNodeTypeDescriptionResponseOutputWithContext(context.Context) NodeTypeDescriptionResponseOutput
}

type NodeTypeDescriptionResponseArgs struct {
	ApplicationPorts             EndpointRangeDescriptionResponsePtrInput `pulumi:"applicationPorts"`
	Capacities                   pulumi.StringMapInput                    `pulumi:"capacities"`
	ClientConnectionEndpointPort pulumi.IntInput                          `pulumi:"clientConnectionEndpointPort"`
	DurabilityLevel              pulumi.StringPtrInput                    `pulumi:"durabilityLevel"`
	EphemeralPorts               EndpointRangeDescriptionResponsePtrInput `pulumi:"ephemeralPorts"`
	HttpGatewayEndpointPort      pulumi.IntInput                          `pulumi:"httpGatewayEndpointPort"`
	IsPrimary                    pulumi.BoolInput                         `pulumi:"isPrimary"`
	Name                         pulumi.StringInput                       `pulumi:"name"`
	PlacementProperties          pulumi.StringMapInput                    `pulumi:"placementProperties"`
	ReverseProxyEndpointPort     pulumi.IntPtrInput                       `pulumi:"reverseProxyEndpointPort"`
	VmInstanceCount              pulumi.IntInput                          `pulumi:"vmInstanceCount"`
}

func (NodeTypeDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeTypeDescriptionResponse)(nil)).Elem()
}

func (i NodeTypeDescriptionResponseArgs) ToNodeTypeDescriptionResponseOutput() NodeTypeDescriptionResponseOutput {
	return i.ToNodeTypeDescriptionResponseOutputWithContext(context.Background())
}

func (i NodeTypeDescriptionResponseArgs) ToNodeTypeDescriptionResponseOutputWithContext(ctx context.Context) NodeTypeDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeTypeDescriptionResponseOutput)
}





type NodeTypeDescriptionResponseArrayInput interface {
	pulumi.Input

	ToNodeTypeDescriptionResponseArrayOutput() NodeTypeDescriptionResponseArrayOutput
	ToNodeTypeDescriptionResponseArrayOutputWithContext(context.Context) NodeTypeDescriptionResponseArrayOutput
}

type NodeTypeDescriptionResponseArray []NodeTypeDescriptionResponseInput

func (NodeTypeDescriptionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeTypeDescriptionResponse)(nil)).Elem()
}

func (i NodeTypeDescriptionResponseArray) ToNodeTypeDescriptionResponseArrayOutput() NodeTypeDescriptionResponseArrayOutput {
	return i.ToNodeTypeDescriptionResponseArrayOutputWithContext(context.Background())
}

func (i NodeTypeDescriptionResponseArray) ToNodeTypeDescriptionResponseArrayOutputWithContext(ctx context.Context) NodeTypeDescriptionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeTypeDescriptionResponseArrayOutput)
}

type NodeTypeDescriptionResponseOutput struct{ *pulumi.OutputState }

func (NodeTypeDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeTypeDescriptionResponse)(nil)).Elem()
}

func (o NodeTypeDescriptionResponseOutput) ToNodeTypeDescriptionResponseOutput() NodeTypeDescriptionResponseOutput {
	return o
}

func (o NodeTypeDescriptionResponseOutput) ToNodeTypeDescriptionResponseOutputWithContext(ctx context.Context) NodeTypeDescriptionResponseOutput {
	return o
}

func (o NodeTypeDescriptionResponseOutput) ApplicationPorts() EndpointRangeDescriptionResponsePtrOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) *EndpointRangeDescriptionResponse { return v.ApplicationPorts }).(EndpointRangeDescriptionResponsePtrOutput)
}

func (o NodeTypeDescriptionResponseOutput) Capacities() pulumi.StringMapOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) map[string]string { return v.Capacities }).(pulumi.StringMapOutput)
}

func (o NodeTypeDescriptionResponseOutput) ClientConnectionEndpointPort() pulumi.IntOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) int { return v.ClientConnectionEndpointPort }).(pulumi.IntOutput)
}

func (o NodeTypeDescriptionResponseOutput) DurabilityLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) *string { return v.DurabilityLevel }).(pulumi.StringPtrOutput)
}

func (o NodeTypeDescriptionResponseOutput) EphemeralPorts() EndpointRangeDescriptionResponsePtrOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) *EndpointRangeDescriptionResponse { return v.EphemeralPorts }).(EndpointRangeDescriptionResponsePtrOutput)
}

func (o NodeTypeDescriptionResponseOutput) HttpGatewayEndpointPort() pulumi.IntOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) int { return v.HttpGatewayEndpointPort }).(pulumi.IntOutput)
}

func (o NodeTypeDescriptionResponseOutput) IsPrimary() pulumi.BoolOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) bool { return v.IsPrimary }).(pulumi.BoolOutput)
}

func (o NodeTypeDescriptionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o NodeTypeDescriptionResponseOutput) PlacementProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) map[string]string { return v.PlacementProperties }).(pulumi.StringMapOutput)
}

func (o NodeTypeDescriptionResponseOutput) ReverseProxyEndpointPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) *int { return v.ReverseProxyEndpointPort }).(pulumi.IntPtrOutput)
}

func (o NodeTypeDescriptionResponseOutput) VmInstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) int { return v.VmInstanceCount }).(pulumi.IntOutput)
}

type NodeTypeDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (NodeTypeDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeTypeDescriptionResponse)(nil)).Elem()
}

func (o NodeTypeDescriptionResponseArrayOutput) ToNodeTypeDescriptionResponseArrayOutput() NodeTypeDescriptionResponseArrayOutput {
	return o
}

func (o NodeTypeDescriptionResponseArrayOutput) ToNodeTypeDescriptionResponseArrayOutputWithContext(ctx context.Context) NodeTypeDescriptionResponseArrayOutput {
	return o
}

func (o NodeTypeDescriptionResponseArrayOutput) Index(i pulumi.IntInput) NodeTypeDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NodeTypeDescriptionResponse {
		return vs[0].([]NodeTypeDescriptionResponse)[vs[1].(int)]
	}).(NodeTypeDescriptionResponseOutput)
}

type ServerCertificateCommonName struct {
	CertificateCommonName       string `pulumi:"certificateCommonName"`
	CertificateIssuerThumbprint string `pulumi:"certificateIssuerThumbprint"`
}





type ServerCertificateCommonNameInput interface {
	pulumi.Input

	ToServerCertificateCommonNameOutput() ServerCertificateCommonNameOutput
	ToServerCertificateCommonNameOutputWithContext(context.Context) ServerCertificateCommonNameOutput
}

type ServerCertificateCommonNameArgs struct {
	CertificateCommonName       pulumi.StringInput `pulumi:"certificateCommonName"`
	CertificateIssuerThumbprint pulumi.StringInput `pulumi:"certificateIssuerThumbprint"`
}

func (ServerCertificateCommonNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerCertificateCommonName)(nil)).Elem()
}

func (i ServerCertificateCommonNameArgs) ToServerCertificateCommonNameOutput() ServerCertificateCommonNameOutput {
	return i.ToServerCertificateCommonNameOutputWithContext(context.Background())
}

func (i ServerCertificateCommonNameArgs) ToServerCertificateCommonNameOutputWithContext(ctx context.Context) ServerCertificateCommonNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCertificateCommonNameOutput)
}





type ServerCertificateCommonNameArrayInput interface {
	pulumi.Input

	ToServerCertificateCommonNameArrayOutput() ServerCertificateCommonNameArrayOutput
	ToServerCertificateCommonNameArrayOutputWithContext(context.Context) ServerCertificateCommonNameArrayOutput
}

type ServerCertificateCommonNameArray []ServerCertificateCommonNameInput

func (ServerCertificateCommonNameArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerCertificateCommonName)(nil)).Elem()
}

func (i ServerCertificateCommonNameArray) ToServerCertificateCommonNameArrayOutput() ServerCertificateCommonNameArrayOutput {
	return i.ToServerCertificateCommonNameArrayOutputWithContext(context.Background())
}

func (i ServerCertificateCommonNameArray) ToServerCertificateCommonNameArrayOutputWithContext(ctx context.Context) ServerCertificateCommonNameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCertificateCommonNameArrayOutput)
}

type ServerCertificateCommonNameOutput struct{ *pulumi.OutputState }

func (ServerCertificateCommonNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerCertificateCommonName)(nil)).Elem()
}

func (o ServerCertificateCommonNameOutput) ToServerCertificateCommonNameOutput() ServerCertificateCommonNameOutput {
	return o
}

func (o ServerCertificateCommonNameOutput) ToServerCertificateCommonNameOutputWithContext(ctx context.Context) ServerCertificateCommonNameOutput {
	return o
}

func (o ServerCertificateCommonNameOutput) CertificateCommonName() pulumi.StringOutput {
	return o.ApplyT(func(v ServerCertificateCommonName) string { return v.CertificateCommonName }).(pulumi.StringOutput)
}

func (o ServerCertificateCommonNameOutput) CertificateIssuerThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v ServerCertificateCommonName) string { return v.CertificateIssuerThumbprint }).(pulumi.StringOutput)
}

type ServerCertificateCommonNameArrayOutput struct{ *pulumi.OutputState }

func (ServerCertificateCommonNameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerCertificateCommonName)(nil)).Elem()
}

func (o ServerCertificateCommonNameArrayOutput) ToServerCertificateCommonNameArrayOutput() ServerCertificateCommonNameArrayOutput {
	return o
}

func (o ServerCertificateCommonNameArrayOutput) ToServerCertificateCommonNameArrayOutputWithContext(ctx context.Context) ServerCertificateCommonNameArrayOutput {
	return o
}

func (o ServerCertificateCommonNameArrayOutput) Index(i pulumi.IntInput) ServerCertificateCommonNameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerCertificateCommonName {
		return vs[0].([]ServerCertificateCommonName)[vs[1].(int)]
	}).(ServerCertificateCommonNameOutput)
}

type ServerCertificateCommonNameResponse struct {
	CertificateCommonName       string `pulumi:"certificateCommonName"`
	CertificateIssuerThumbprint string `pulumi:"certificateIssuerThumbprint"`
}





type ServerCertificateCommonNameResponseInput interface {
	pulumi.Input

	ToServerCertificateCommonNameResponseOutput() ServerCertificateCommonNameResponseOutput
	ToServerCertificateCommonNameResponseOutputWithContext(context.Context) ServerCertificateCommonNameResponseOutput
}

type ServerCertificateCommonNameResponseArgs struct {
	CertificateCommonName       pulumi.StringInput `pulumi:"certificateCommonName"`
	CertificateIssuerThumbprint pulumi.StringInput `pulumi:"certificateIssuerThumbprint"`
}

func (ServerCertificateCommonNameResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerCertificateCommonNameResponse)(nil)).Elem()
}

func (i ServerCertificateCommonNameResponseArgs) ToServerCertificateCommonNameResponseOutput() ServerCertificateCommonNameResponseOutput {
	return i.ToServerCertificateCommonNameResponseOutputWithContext(context.Background())
}

func (i ServerCertificateCommonNameResponseArgs) ToServerCertificateCommonNameResponseOutputWithContext(ctx context.Context) ServerCertificateCommonNameResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCertificateCommonNameResponseOutput)
}





type ServerCertificateCommonNameResponseArrayInput interface {
	pulumi.Input

	ToServerCertificateCommonNameResponseArrayOutput() ServerCertificateCommonNameResponseArrayOutput
	ToServerCertificateCommonNameResponseArrayOutputWithContext(context.Context) ServerCertificateCommonNameResponseArrayOutput
}

type ServerCertificateCommonNameResponseArray []ServerCertificateCommonNameResponseInput

func (ServerCertificateCommonNameResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerCertificateCommonNameResponse)(nil)).Elem()
}

func (i ServerCertificateCommonNameResponseArray) ToServerCertificateCommonNameResponseArrayOutput() ServerCertificateCommonNameResponseArrayOutput {
	return i.ToServerCertificateCommonNameResponseArrayOutputWithContext(context.Background())
}

func (i ServerCertificateCommonNameResponseArray) ToServerCertificateCommonNameResponseArrayOutputWithContext(ctx context.Context) ServerCertificateCommonNameResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCertificateCommonNameResponseArrayOutput)
}

type ServerCertificateCommonNameResponseOutput struct{ *pulumi.OutputState }

func (ServerCertificateCommonNameResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerCertificateCommonNameResponse)(nil)).Elem()
}

func (o ServerCertificateCommonNameResponseOutput) ToServerCertificateCommonNameResponseOutput() ServerCertificateCommonNameResponseOutput {
	return o
}

func (o ServerCertificateCommonNameResponseOutput) ToServerCertificateCommonNameResponseOutputWithContext(ctx context.Context) ServerCertificateCommonNameResponseOutput {
	return o
}

func (o ServerCertificateCommonNameResponseOutput) CertificateCommonName() pulumi.StringOutput {
	return o.ApplyT(func(v ServerCertificateCommonNameResponse) string { return v.CertificateCommonName }).(pulumi.StringOutput)
}

func (o ServerCertificateCommonNameResponseOutput) CertificateIssuerThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v ServerCertificateCommonNameResponse) string { return v.CertificateIssuerThumbprint }).(pulumi.StringOutput)
}

type ServerCertificateCommonNameResponseArrayOutput struct{ *pulumi.OutputState }

func (ServerCertificateCommonNameResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerCertificateCommonNameResponse)(nil)).Elem()
}

func (o ServerCertificateCommonNameResponseArrayOutput) ToServerCertificateCommonNameResponseArrayOutput() ServerCertificateCommonNameResponseArrayOutput {
	return o
}

func (o ServerCertificateCommonNameResponseArrayOutput) ToServerCertificateCommonNameResponseArrayOutputWithContext(ctx context.Context) ServerCertificateCommonNameResponseArrayOutput {
	return o
}

func (o ServerCertificateCommonNameResponseArrayOutput) Index(i pulumi.IntInput) ServerCertificateCommonNameResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerCertificateCommonNameResponse {
		return vs[0].([]ServerCertificateCommonNameResponse)[vs[1].(int)]
	}).(ServerCertificateCommonNameResponseOutput)
}

type ServerCertificateCommonNames struct {
	CommonNames   []ServerCertificateCommonName `pulumi:"commonNames"`
	X509StoreName *string                       `pulumi:"x509StoreName"`
}





type ServerCertificateCommonNamesInput interface {
	pulumi.Input

	ToServerCertificateCommonNamesOutput() ServerCertificateCommonNamesOutput
	ToServerCertificateCommonNamesOutputWithContext(context.Context) ServerCertificateCommonNamesOutput
}

type ServerCertificateCommonNamesArgs struct {
	CommonNames   ServerCertificateCommonNameArrayInput `pulumi:"commonNames"`
	X509StoreName pulumi.StringPtrInput                 `pulumi:"x509StoreName"`
}

func (ServerCertificateCommonNamesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerCertificateCommonNames)(nil)).Elem()
}

func (i ServerCertificateCommonNamesArgs) ToServerCertificateCommonNamesOutput() ServerCertificateCommonNamesOutput {
	return i.ToServerCertificateCommonNamesOutputWithContext(context.Background())
}

func (i ServerCertificateCommonNamesArgs) ToServerCertificateCommonNamesOutputWithContext(ctx context.Context) ServerCertificateCommonNamesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCertificateCommonNamesOutput)
}

func (i ServerCertificateCommonNamesArgs) ToServerCertificateCommonNamesPtrOutput() ServerCertificateCommonNamesPtrOutput {
	return i.ToServerCertificateCommonNamesPtrOutputWithContext(context.Background())
}

func (i ServerCertificateCommonNamesArgs) ToServerCertificateCommonNamesPtrOutputWithContext(ctx context.Context) ServerCertificateCommonNamesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCertificateCommonNamesOutput).ToServerCertificateCommonNamesPtrOutputWithContext(ctx)
}









type ServerCertificateCommonNamesPtrInput interface {
	pulumi.Input

	ToServerCertificateCommonNamesPtrOutput() ServerCertificateCommonNamesPtrOutput
	ToServerCertificateCommonNamesPtrOutputWithContext(context.Context) ServerCertificateCommonNamesPtrOutput
}

type serverCertificateCommonNamesPtrType ServerCertificateCommonNamesArgs

func ServerCertificateCommonNamesPtr(v *ServerCertificateCommonNamesArgs) ServerCertificateCommonNamesPtrInput {
	return (*serverCertificateCommonNamesPtrType)(v)
}

func (*serverCertificateCommonNamesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerCertificateCommonNames)(nil)).Elem()
}

func (i *serverCertificateCommonNamesPtrType) ToServerCertificateCommonNamesPtrOutput() ServerCertificateCommonNamesPtrOutput {
	return i.ToServerCertificateCommonNamesPtrOutputWithContext(context.Background())
}

func (i *serverCertificateCommonNamesPtrType) ToServerCertificateCommonNamesPtrOutputWithContext(ctx context.Context) ServerCertificateCommonNamesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCertificateCommonNamesPtrOutput)
}

type ServerCertificateCommonNamesOutput struct{ *pulumi.OutputState }

func (ServerCertificateCommonNamesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerCertificateCommonNames)(nil)).Elem()
}

func (o ServerCertificateCommonNamesOutput) ToServerCertificateCommonNamesOutput() ServerCertificateCommonNamesOutput {
	return o
}

func (o ServerCertificateCommonNamesOutput) ToServerCertificateCommonNamesOutputWithContext(ctx context.Context) ServerCertificateCommonNamesOutput {
	return o
}

func (o ServerCertificateCommonNamesOutput) ToServerCertificateCommonNamesPtrOutput() ServerCertificateCommonNamesPtrOutput {
	return o.ToServerCertificateCommonNamesPtrOutputWithContext(context.Background())
}

func (o ServerCertificateCommonNamesOutput) ToServerCertificateCommonNamesPtrOutputWithContext(ctx context.Context) ServerCertificateCommonNamesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerCertificateCommonNames) *ServerCertificateCommonNames {
		return &v
	}).(ServerCertificateCommonNamesPtrOutput)
}

func (o ServerCertificateCommonNamesOutput) CommonNames() ServerCertificateCommonNameArrayOutput {
	return o.ApplyT(func(v ServerCertificateCommonNames) []ServerCertificateCommonName { return v.CommonNames }).(ServerCertificateCommonNameArrayOutput)
}

func (o ServerCertificateCommonNamesOutput) X509StoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerCertificateCommonNames) *string { return v.X509StoreName }).(pulumi.StringPtrOutput)
}

type ServerCertificateCommonNamesPtrOutput struct{ *pulumi.OutputState }

func (ServerCertificateCommonNamesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerCertificateCommonNames)(nil)).Elem()
}

func (o ServerCertificateCommonNamesPtrOutput) ToServerCertificateCommonNamesPtrOutput() ServerCertificateCommonNamesPtrOutput {
	return o
}

func (o ServerCertificateCommonNamesPtrOutput) ToServerCertificateCommonNamesPtrOutputWithContext(ctx context.Context) ServerCertificateCommonNamesPtrOutput {
	return o
}

func (o ServerCertificateCommonNamesPtrOutput) Elem() ServerCertificateCommonNamesOutput {
	return o.ApplyT(func(v *ServerCertificateCommonNames) ServerCertificateCommonNames {
		if v != nil {
			return *v
		}
		var ret ServerCertificateCommonNames
		return ret
	}).(ServerCertificateCommonNamesOutput)
}

func (o ServerCertificateCommonNamesPtrOutput) CommonNames() ServerCertificateCommonNameArrayOutput {
	return o.ApplyT(func(v *ServerCertificateCommonNames) []ServerCertificateCommonName {
		if v == nil {
			return nil
		}
		return v.CommonNames
	}).(ServerCertificateCommonNameArrayOutput)
}

func (o ServerCertificateCommonNamesPtrOutput) X509StoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerCertificateCommonNames) *string {
		if v == nil {
			return nil
		}
		return v.X509StoreName
	}).(pulumi.StringPtrOutput)
}

type ServerCertificateCommonNamesResponse struct {
	CommonNames   []ServerCertificateCommonNameResponse `pulumi:"commonNames"`
	X509StoreName *string                               `pulumi:"x509StoreName"`
}





type ServerCertificateCommonNamesResponseInput interface {
	pulumi.Input

	ToServerCertificateCommonNamesResponseOutput() ServerCertificateCommonNamesResponseOutput
	ToServerCertificateCommonNamesResponseOutputWithContext(context.Context) ServerCertificateCommonNamesResponseOutput
}

type ServerCertificateCommonNamesResponseArgs struct {
	CommonNames   ServerCertificateCommonNameResponseArrayInput `pulumi:"commonNames"`
	X509StoreName pulumi.StringPtrInput                         `pulumi:"x509StoreName"`
}

func (ServerCertificateCommonNamesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerCertificateCommonNamesResponse)(nil)).Elem()
}

func (i ServerCertificateCommonNamesResponseArgs) ToServerCertificateCommonNamesResponseOutput() ServerCertificateCommonNamesResponseOutput {
	return i.ToServerCertificateCommonNamesResponseOutputWithContext(context.Background())
}

func (i ServerCertificateCommonNamesResponseArgs) ToServerCertificateCommonNamesResponseOutputWithContext(ctx context.Context) ServerCertificateCommonNamesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCertificateCommonNamesResponseOutput)
}

func (i ServerCertificateCommonNamesResponseArgs) ToServerCertificateCommonNamesResponsePtrOutput() ServerCertificateCommonNamesResponsePtrOutput {
	return i.ToServerCertificateCommonNamesResponsePtrOutputWithContext(context.Background())
}

func (i ServerCertificateCommonNamesResponseArgs) ToServerCertificateCommonNamesResponsePtrOutputWithContext(ctx context.Context) ServerCertificateCommonNamesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCertificateCommonNamesResponseOutput).ToServerCertificateCommonNamesResponsePtrOutputWithContext(ctx)
}









type ServerCertificateCommonNamesResponsePtrInput interface {
	pulumi.Input

	ToServerCertificateCommonNamesResponsePtrOutput() ServerCertificateCommonNamesResponsePtrOutput
	ToServerCertificateCommonNamesResponsePtrOutputWithContext(context.Context) ServerCertificateCommonNamesResponsePtrOutput
}

type serverCertificateCommonNamesResponsePtrType ServerCertificateCommonNamesResponseArgs

func ServerCertificateCommonNamesResponsePtr(v *ServerCertificateCommonNamesResponseArgs) ServerCertificateCommonNamesResponsePtrInput {
	return (*serverCertificateCommonNamesResponsePtrType)(v)
}

func (*serverCertificateCommonNamesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerCertificateCommonNamesResponse)(nil)).Elem()
}

func (i *serverCertificateCommonNamesResponsePtrType) ToServerCertificateCommonNamesResponsePtrOutput() ServerCertificateCommonNamesResponsePtrOutput {
	return i.ToServerCertificateCommonNamesResponsePtrOutputWithContext(context.Background())
}

func (i *serverCertificateCommonNamesResponsePtrType) ToServerCertificateCommonNamesResponsePtrOutputWithContext(ctx context.Context) ServerCertificateCommonNamesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCertificateCommonNamesResponsePtrOutput)
}

type ServerCertificateCommonNamesResponseOutput struct{ *pulumi.OutputState }

func (ServerCertificateCommonNamesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerCertificateCommonNamesResponse)(nil)).Elem()
}

func (o ServerCertificateCommonNamesResponseOutput) ToServerCertificateCommonNamesResponseOutput() ServerCertificateCommonNamesResponseOutput {
	return o
}

func (o ServerCertificateCommonNamesResponseOutput) ToServerCertificateCommonNamesResponseOutputWithContext(ctx context.Context) ServerCertificateCommonNamesResponseOutput {
	return o
}

func (o ServerCertificateCommonNamesResponseOutput) ToServerCertificateCommonNamesResponsePtrOutput() ServerCertificateCommonNamesResponsePtrOutput {
	return o.ToServerCertificateCommonNamesResponsePtrOutputWithContext(context.Background())
}

func (o ServerCertificateCommonNamesResponseOutput) ToServerCertificateCommonNamesResponsePtrOutputWithContext(ctx context.Context) ServerCertificateCommonNamesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerCertificateCommonNamesResponse) *ServerCertificateCommonNamesResponse {
		return &v
	}).(ServerCertificateCommonNamesResponsePtrOutput)
}

func (o ServerCertificateCommonNamesResponseOutput) CommonNames() ServerCertificateCommonNameResponseArrayOutput {
	return o.ApplyT(func(v ServerCertificateCommonNamesResponse) []ServerCertificateCommonNameResponse {
		return v.CommonNames
	}).(ServerCertificateCommonNameResponseArrayOutput)
}

func (o ServerCertificateCommonNamesResponseOutput) X509StoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerCertificateCommonNamesResponse) *string { return v.X509StoreName }).(pulumi.StringPtrOutput)
}

type ServerCertificateCommonNamesResponsePtrOutput struct{ *pulumi.OutputState }

func (ServerCertificateCommonNamesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerCertificateCommonNamesResponse)(nil)).Elem()
}

func (o ServerCertificateCommonNamesResponsePtrOutput) ToServerCertificateCommonNamesResponsePtrOutput() ServerCertificateCommonNamesResponsePtrOutput {
	return o
}

func (o ServerCertificateCommonNamesResponsePtrOutput) ToServerCertificateCommonNamesResponsePtrOutputWithContext(ctx context.Context) ServerCertificateCommonNamesResponsePtrOutput {
	return o
}

func (o ServerCertificateCommonNamesResponsePtrOutput) Elem() ServerCertificateCommonNamesResponseOutput {
	return o.ApplyT(func(v *ServerCertificateCommonNamesResponse) ServerCertificateCommonNamesResponse {
		if v != nil {
			return *v
		}
		var ret ServerCertificateCommonNamesResponse
		return ret
	}).(ServerCertificateCommonNamesResponseOutput)
}

func (o ServerCertificateCommonNamesResponsePtrOutput) CommonNames() ServerCertificateCommonNameResponseArrayOutput {
	return o.ApplyT(func(v *ServerCertificateCommonNamesResponse) []ServerCertificateCommonNameResponse {
		if v == nil {
			return nil
		}
		return v.CommonNames
	}).(ServerCertificateCommonNameResponseArrayOutput)
}

func (o ServerCertificateCommonNamesResponsePtrOutput) X509StoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerCertificateCommonNamesResponse) *string {
		if v == nil {
			return nil
		}
		return v.X509StoreName
	}).(pulumi.StringPtrOutput)
}

type ServiceCorrelationDescription struct {
	Scheme      string `pulumi:"scheme"`
	ServiceName string `pulumi:"serviceName"`
}





type ServiceCorrelationDescriptionInput interface {
	pulumi.Input

	ToServiceCorrelationDescriptionOutput() ServiceCorrelationDescriptionOutput
	ToServiceCorrelationDescriptionOutputWithContext(context.Context) ServiceCorrelationDescriptionOutput
}

type ServiceCorrelationDescriptionArgs struct {
	Scheme      pulumi.StringInput `pulumi:"scheme"`
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (ServiceCorrelationDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorrelationDescription)(nil)).Elem()
}

func (i ServiceCorrelationDescriptionArgs) ToServiceCorrelationDescriptionOutput() ServiceCorrelationDescriptionOutput {
	return i.ToServiceCorrelationDescriptionOutputWithContext(context.Background())
}

func (i ServiceCorrelationDescriptionArgs) ToServiceCorrelationDescriptionOutputWithContext(ctx context.Context) ServiceCorrelationDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorrelationDescriptionOutput)
}





type ServiceCorrelationDescriptionArrayInput interface {
	pulumi.Input

	ToServiceCorrelationDescriptionArrayOutput() ServiceCorrelationDescriptionArrayOutput
	ToServiceCorrelationDescriptionArrayOutputWithContext(context.Context) ServiceCorrelationDescriptionArrayOutput
}

type ServiceCorrelationDescriptionArray []ServiceCorrelationDescriptionInput

func (ServiceCorrelationDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceCorrelationDescription)(nil)).Elem()
}

func (i ServiceCorrelationDescriptionArray) ToServiceCorrelationDescriptionArrayOutput() ServiceCorrelationDescriptionArrayOutput {
	return i.ToServiceCorrelationDescriptionArrayOutputWithContext(context.Background())
}

func (i ServiceCorrelationDescriptionArray) ToServiceCorrelationDescriptionArrayOutputWithContext(ctx context.Context) ServiceCorrelationDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorrelationDescriptionArrayOutput)
}

type ServiceCorrelationDescriptionOutput struct{ *pulumi.OutputState }

func (ServiceCorrelationDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorrelationDescription)(nil)).Elem()
}

func (o ServiceCorrelationDescriptionOutput) ToServiceCorrelationDescriptionOutput() ServiceCorrelationDescriptionOutput {
	return o
}

func (o ServiceCorrelationDescriptionOutput) ToServiceCorrelationDescriptionOutputWithContext(ctx context.Context) ServiceCorrelationDescriptionOutput {
	return o
}

func (o ServiceCorrelationDescriptionOutput) Scheme() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceCorrelationDescription) string { return v.Scheme }).(pulumi.StringOutput)
}

func (o ServiceCorrelationDescriptionOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceCorrelationDescription) string { return v.ServiceName }).(pulumi.StringOutput)
}

type ServiceCorrelationDescriptionArrayOutput struct{ *pulumi.OutputState }

func (ServiceCorrelationDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceCorrelationDescription)(nil)).Elem()
}

func (o ServiceCorrelationDescriptionArrayOutput) ToServiceCorrelationDescriptionArrayOutput() ServiceCorrelationDescriptionArrayOutput {
	return o
}

func (o ServiceCorrelationDescriptionArrayOutput) ToServiceCorrelationDescriptionArrayOutputWithContext(ctx context.Context) ServiceCorrelationDescriptionArrayOutput {
	return o
}

func (o ServiceCorrelationDescriptionArrayOutput) Index(i pulumi.IntInput) ServiceCorrelationDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceCorrelationDescription {
		return vs[0].([]ServiceCorrelationDescription)[vs[1].(int)]
	}).(ServiceCorrelationDescriptionOutput)
}

type ServiceCorrelationDescriptionResponse struct {
	Scheme      string `pulumi:"scheme"`
	ServiceName string `pulumi:"serviceName"`
}





type ServiceCorrelationDescriptionResponseInput interface {
	pulumi.Input

	ToServiceCorrelationDescriptionResponseOutput() ServiceCorrelationDescriptionResponseOutput
	ToServiceCorrelationDescriptionResponseOutputWithContext(context.Context) ServiceCorrelationDescriptionResponseOutput
}

type ServiceCorrelationDescriptionResponseArgs struct {
	Scheme      pulumi.StringInput `pulumi:"scheme"`
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (ServiceCorrelationDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorrelationDescriptionResponse)(nil)).Elem()
}

func (i ServiceCorrelationDescriptionResponseArgs) ToServiceCorrelationDescriptionResponseOutput() ServiceCorrelationDescriptionResponseOutput {
	return i.ToServiceCorrelationDescriptionResponseOutputWithContext(context.Background())
}

func (i ServiceCorrelationDescriptionResponseArgs) ToServiceCorrelationDescriptionResponseOutputWithContext(ctx context.Context) ServiceCorrelationDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorrelationDescriptionResponseOutput)
}





type ServiceCorrelationDescriptionResponseArrayInput interface {
	pulumi.Input

	ToServiceCorrelationDescriptionResponseArrayOutput() ServiceCorrelationDescriptionResponseArrayOutput
	ToServiceCorrelationDescriptionResponseArrayOutputWithContext(context.Context) ServiceCorrelationDescriptionResponseArrayOutput
}

type ServiceCorrelationDescriptionResponseArray []ServiceCorrelationDescriptionResponseInput

func (ServiceCorrelationDescriptionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceCorrelationDescriptionResponse)(nil)).Elem()
}

func (i ServiceCorrelationDescriptionResponseArray) ToServiceCorrelationDescriptionResponseArrayOutput() ServiceCorrelationDescriptionResponseArrayOutput {
	return i.ToServiceCorrelationDescriptionResponseArrayOutputWithContext(context.Background())
}

func (i ServiceCorrelationDescriptionResponseArray) ToServiceCorrelationDescriptionResponseArrayOutputWithContext(ctx context.Context) ServiceCorrelationDescriptionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorrelationDescriptionResponseArrayOutput)
}

type ServiceCorrelationDescriptionResponseOutput struct{ *pulumi.OutputState }

func (ServiceCorrelationDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorrelationDescriptionResponse)(nil)).Elem()
}

func (o ServiceCorrelationDescriptionResponseOutput) ToServiceCorrelationDescriptionResponseOutput() ServiceCorrelationDescriptionResponseOutput {
	return o
}

func (o ServiceCorrelationDescriptionResponseOutput) ToServiceCorrelationDescriptionResponseOutputWithContext(ctx context.Context) ServiceCorrelationDescriptionResponseOutput {
	return o
}

func (o ServiceCorrelationDescriptionResponseOutput) Scheme() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceCorrelationDescriptionResponse) string { return v.Scheme }).(pulumi.StringOutput)
}

func (o ServiceCorrelationDescriptionResponseOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceCorrelationDescriptionResponse) string { return v.ServiceName }).(pulumi.StringOutput)
}

type ServiceCorrelationDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceCorrelationDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceCorrelationDescriptionResponse)(nil)).Elem()
}

func (o ServiceCorrelationDescriptionResponseArrayOutput) ToServiceCorrelationDescriptionResponseArrayOutput() ServiceCorrelationDescriptionResponseArrayOutput {
	return o
}

func (o ServiceCorrelationDescriptionResponseArrayOutput) ToServiceCorrelationDescriptionResponseArrayOutputWithContext(ctx context.Context) ServiceCorrelationDescriptionResponseArrayOutput {
	return o
}

func (o ServiceCorrelationDescriptionResponseArrayOutput) Index(i pulumi.IntInput) ServiceCorrelationDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceCorrelationDescriptionResponse {
		return vs[0].([]ServiceCorrelationDescriptionResponse)[vs[1].(int)]
	}).(ServiceCorrelationDescriptionResponseOutput)
}

type ServiceLoadMetricDescription struct {
	DefaultLoad          *int    `pulumi:"defaultLoad"`
	Name                 string  `pulumi:"name"`
	PrimaryDefaultLoad   *int    `pulumi:"primaryDefaultLoad"`
	SecondaryDefaultLoad *int    `pulumi:"secondaryDefaultLoad"`
	Weight               *string `pulumi:"weight"`
}





type ServiceLoadMetricDescriptionInput interface {
	pulumi.Input

	ToServiceLoadMetricDescriptionOutput() ServiceLoadMetricDescriptionOutput
	ToServiceLoadMetricDescriptionOutputWithContext(context.Context) ServiceLoadMetricDescriptionOutput
}

type ServiceLoadMetricDescriptionArgs struct {
	DefaultLoad          pulumi.IntPtrInput    `pulumi:"defaultLoad"`
	Name                 pulumi.StringInput    `pulumi:"name"`
	PrimaryDefaultLoad   pulumi.IntPtrInput    `pulumi:"primaryDefaultLoad"`
	SecondaryDefaultLoad pulumi.IntPtrInput    `pulumi:"secondaryDefaultLoad"`
	Weight               pulumi.StringPtrInput `pulumi:"weight"`
}

func (ServiceLoadMetricDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceLoadMetricDescription)(nil)).Elem()
}

func (i ServiceLoadMetricDescriptionArgs) ToServiceLoadMetricDescriptionOutput() ServiceLoadMetricDescriptionOutput {
	return i.ToServiceLoadMetricDescriptionOutputWithContext(context.Background())
}

func (i ServiceLoadMetricDescriptionArgs) ToServiceLoadMetricDescriptionOutputWithContext(ctx context.Context) ServiceLoadMetricDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceLoadMetricDescriptionOutput)
}





type ServiceLoadMetricDescriptionArrayInput interface {
	pulumi.Input

	ToServiceLoadMetricDescriptionArrayOutput() ServiceLoadMetricDescriptionArrayOutput
	ToServiceLoadMetricDescriptionArrayOutputWithContext(context.Context) ServiceLoadMetricDescriptionArrayOutput
}

type ServiceLoadMetricDescriptionArray []ServiceLoadMetricDescriptionInput

func (ServiceLoadMetricDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceLoadMetricDescription)(nil)).Elem()
}

func (i ServiceLoadMetricDescriptionArray) ToServiceLoadMetricDescriptionArrayOutput() ServiceLoadMetricDescriptionArrayOutput {
	return i.ToServiceLoadMetricDescriptionArrayOutputWithContext(context.Background())
}

func (i ServiceLoadMetricDescriptionArray) ToServiceLoadMetricDescriptionArrayOutputWithContext(ctx context.Context) ServiceLoadMetricDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceLoadMetricDescriptionArrayOutput)
}

type ServiceLoadMetricDescriptionOutput struct{ *pulumi.OutputState }

func (ServiceLoadMetricDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceLoadMetricDescription)(nil)).Elem()
}

func (o ServiceLoadMetricDescriptionOutput) ToServiceLoadMetricDescriptionOutput() ServiceLoadMetricDescriptionOutput {
	return o
}

func (o ServiceLoadMetricDescriptionOutput) ToServiceLoadMetricDescriptionOutputWithContext(ctx context.Context) ServiceLoadMetricDescriptionOutput {
	return o
}

func (o ServiceLoadMetricDescriptionOutput) DefaultLoad() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceLoadMetricDescription) *int { return v.DefaultLoad }).(pulumi.IntPtrOutput)
}

func (o ServiceLoadMetricDescriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceLoadMetricDescription) string { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceLoadMetricDescriptionOutput) PrimaryDefaultLoad() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceLoadMetricDescription) *int { return v.PrimaryDefaultLoad }).(pulumi.IntPtrOutput)
}

func (o ServiceLoadMetricDescriptionOutput) SecondaryDefaultLoad() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceLoadMetricDescription) *int { return v.SecondaryDefaultLoad }).(pulumi.IntPtrOutput)
}

func (o ServiceLoadMetricDescriptionOutput) Weight() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceLoadMetricDescription) *string { return v.Weight }).(pulumi.StringPtrOutput)
}

type ServiceLoadMetricDescriptionArrayOutput struct{ *pulumi.OutputState }

func (ServiceLoadMetricDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceLoadMetricDescription)(nil)).Elem()
}

func (o ServiceLoadMetricDescriptionArrayOutput) ToServiceLoadMetricDescriptionArrayOutput() ServiceLoadMetricDescriptionArrayOutput {
	return o
}

func (o ServiceLoadMetricDescriptionArrayOutput) ToServiceLoadMetricDescriptionArrayOutputWithContext(ctx context.Context) ServiceLoadMetricDescriptionArrayOutput {
	return o
}

func (o ServiceLoadMetricDescriptionArrayOutput) Index(i pulumi.IntInput) ServiceLoadMetricDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceLoadMetricDescription {
		return vs[0].([]ServiceLoadMetricDescription)[vs[1].(int)]
	}).(ServiceLoadMetricDescriptionOutput)
}

type ServiceLoadMetricDescriptionResponse struct {
	DefaultLoad          *int    `pulumi:"defaultLoad"`
	Name                 string  `pulumi:"name"`
	PrimaryDefaultLoad   *int    `pulumi:"primaryDefaultLoad"`
	SecondaryDefaultLoad *int    `pulumi:"secondaryDefaultLoad"`
	Weight               *string `pulumi:"weight"`
}





type ServiceLoadMetricDescriptionResponseInput interface {
	pulumi.Input

	ToServiceLoadMetricDescriptionResponseOutput() ServiceLoadMetricDescriptionResponseOutput
	ToServiceLoadMetricDescriptionResponseOutputWithContext(context.Context) ServiceLoadMetricDescriptionResponseOutput
}

type ServiceLoadMetricDescriptionResponseArgs struct {
	DefaultLoad          pulumi.IntPtrInput    `pulumi:"defaultLoad"`
	Name                 pulumi.StringInput    `pulumi:"name"`
	PrimaryDefaultLoad   pulumi.IntPtrInput    `pulumi:"primaryDefaultLoad"`
	SecondaryDefaultLoad pulumi.IntPtrInput    `pulumi:"secondaryDefaultLoad"`
	Weight               pulumi.StringPtrInput `pulumi:"weight"`
}

func (ServiceLoadMetricDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceLoadMetricDescriptionResponse)(nil)).Elem()
}

func (i ServiceLoadMetricDescriptionResponseArgs) ToServiceLoadMetricDescriptionResponseOutput() ServiceLoadMetricDescriptionResponseOutput {
	return i.ToServiceLoadMetricDescriptionResponseOutputWithContext(context.Background())
}

func (i ServiceLoadMetricDescriptionResponseArgs) ToServiceLoadMetricDescriptionResponseOutputWithContext(ctx context.Context) ServiceLoadMetricDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceLoadMetricDescriptionResponseOutput)
}





type ServiceLoadMetricDescriptionResponseArrayInput interface {
	pulumi.Input

	ToServiceLoadMetricDescriptionResponseArrayOutput() ServiceLoadMetricDescriptionResponseArrayOutput
	ToServiceLoadMetricDescriptionResponseArrayOutputWithContext(context.Context) ServiceLoadMetricDescriptionResponseArrayOutput
}

type ServiceLoadMetricDescriptionResponseArray []ServiceLoadMetricDescriptionResponseInput

func (ServiceLoadMetricDescriptionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceLoadMetricDescriptionResponse)(nil)).Elem()
}

func (i ServiceLoadMetricDescriptionResponseArray) ToServiceLoadMetricDescriptionResponseArrayOutput() ServiceLoadMetricDescriptionResponseArrayOutput {
	return i.ToServiceLoadMetricDescriptionResponseArrayOutputWithContext(context.Background())
}

func (i ServiceLoadMetricDescriptionResponseArray) ToServiceLoadMetricDescriptionResponseArrayOutputWithContext(ctx context.Context) ServiceLoadMetricDescriptionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceLoadMetricDescriptionResponseArrayOutput)
}

type ServiceLoadMetricDescriptionResponseOutput struct{ *pulumi.OutputState }

func (ServiceLoadMetricDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceLoadMetricDescriptionResponse)(nil)).Elem()
}

func (o ServiceLoadMetricDescriptionResponseOutput) ToServiceLoadMetricDescriptionResponseOutput() ServiceLoadMetricDescriptionResponseOutput {
	return o
}

func (o ServiceLoadMetricDescriptionResponseOutput) ToServiceLoadMetricDescriptionResponseOutputWithContext(ctx context.Context) ServiceLoadMetricDescriptionResponseOutput {
	return o
}

func (o ServiceLoadMetricDescriptionResponseOutput) DefaultLoad() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceLoadMetricDescriptionResponse) *int { return v.DefaultLoad }).(pulumi.IntPtrOutput)
}

func (o ServiceLoadMetricDescriptionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceLoadMetricDescriptionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceLoadMetricDescriptionResponseOutput) PrimaryDefaultLoad() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceLoadMetricDescriptionResponse) *int { return v.PrimaryDefaultLoad }).(pulumi.IntPtrOutput)
}

func (o ServiceLoadMetricDescriptionResponseOutput) SecondaryDefaultLoad() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceLoadMetricDescriptionResponse) *int { return v.SecondaryDefaultLoad }).(pulumi.IntPtrOutput)
}

func (o ServiceLoadMetricDescriptionResponseOutput) Weight() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceLoadMetricDescriptionResponse) *string { return v.Weight }).(pulumi.StringPtrOutput)
}

type ServiceLoadMetricDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceLoadMetricDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceLoadMetricDescriptionResponse)(nil)).Elem()
}

func (o ServiceLoadMetricDescriptionResponseArrayOutput) ToServiceLoadMetricDescriptionResponseArrayOutput() ServiceLoadMetricDescriptionResponseArrayOutput {
	return o
}

func (o ServiceLoadMetricDescriptionResponseArrayOutput) ToServiceLoadMetricDescriptionResponseArrayOutputWithContext(ctx context.Context) ServiceLoadMetricDescriptionResponseArrayOutput {
	return o
}

func (o ServiceLoadMetricDescriptionResponseArrayOutput) Index(i pulumi.IntInput) ServiceLoadMetricDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceLoadMetricDescriptionResponse {
		return vs[0].([]ServiceLoadMetricDescriptionResponse)[vs[1].(int)]
	}).(ServiceLoadMetricDescriptionResponseOutput)
}

type ServicePlacementPolicyDescription struct {
	Type string `pulumi:"type"`
}





type ServicePlacementPolicyDescriptionInput interface {
	pulumi.Input

	ToServicePlacementPolicyDescriptionOutput() ServicePlacementPolicyDescriptionOutput
	ToServicePlacementPolicyDescriptionOutputWithContext(context.Context) ServicePlacementPolicyDescriptionOutput
}

type ServicePlacementPolicyDescriptionArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
}

func (ServicePlacementPolicyDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementPolicyDescription)(nil)).Elem()
}

func (i ServicePlacementPolicyDescriptionArgs) ToServicePlacementPolicyDescriptionOutput() ServicePlacementPolicyDescriptionOutput {
	return i.ToServicePlacementPolicyDescriptionOutputWithContext(context.Background())
}

func (i ServicePlacementPolicyDescriptionArgs) ToServicePlacementPolicyDescriptionOutputWithContext(ctx context.Context) ServicePlacementPolicyDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePlacementPolicyDescriptionOutput)
}





type ServicePlacementPolicyDescriptionArrayInput interface {
	pulumi.Input

	ToServicePlacementPolicyDescriptionArrayOutput() ServicePlacementPolicyDescriptionArrayOutput
	ToServicePlacementPolicyDescriptionArrayOutputWithContext(context.Context) ServicePlacementPolicyDescriptionArrayOutput
}

type ServicePlacementPolicyDescriptionArray []ServicePlacementPolicyDescriptionInput

func (ServicePlacementPolicyDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServicePlacementPolicyDescription)(nil)).Elem()
}

func (i ServicePlacementPolicyDescriptionArray) ToServicePlacementPolicyDescriptionArrayOutput() ServicePlacementPolicyDescriptionArrayOutput {
	return i.ToServicePlacementPolicyDescriptionArrayOutputWithContext(context.Background())
}

func (i ServicePlacementPolicyDescriptionArray) ToServicePlacementPolicyDescriptionArrayOutputWithContext(ctx context.Context) ServicePlacementPolicyDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePlacementPolicyDescriptionArrayOutput)
}

type ServicePlacementPolicyDescriptionOutput struct{ *pulumi.OutputState }

func (ServicePlacementPolicyDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementPolicyDescription)(nil)).Elem()
}

func (o ServicePlacementPolicyDescriptionOutput) ToServicePlacementPolicyDescriptionOutput() ServicePlacementPolicyDescriptionOutput {
	return o
}

func (o ServicePlacementPolicyDescriptionOutput) ToServicePlacementPolicyDescriptionOutputWithContext(ctx context.Context) ServicePlacementPolicyDescriptionOutput {
	return o
}

func (o ServicePlacementPolicyDescriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementPolicyDescription) string { return v.Type }).(pulumi.StringOutput)
}

type ServicePlacementPolicyDescriptionArrayOutput struct{ *pulumi.OutputState }

func (ServicePlacementPolicyDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServicePlacementPolicyDescription)(nil)).Elem()
}

func (o ServicePlacementPolicyDescriptionArrayOutput) ToServicePlacementPolicyDescriptionArrayOutput() ServicePlacementPolicyDescriptionArrayOutput {
	return o
}

func (o ServicePlacementPolicyDescriptionArrayOutput) ToServicePlacementPolicyDescriptionArrayOutputWithContext(ctx context.Context) ServicePlacementPolicyDescriptionArrayOutput {
	return o
}

func (o ServicePlacementPolicyDescriptionArrayOutput) Index(i pulumi.IntInput) ServicePlacementPolicyDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServicePlacementPolicyDescription {
		return vs[0].([]ServicePlacementPolicyDescription)[vs[1].(int)]
	}).(ServicePlacementPolicyDescriptionOutput)
}

type ServicePlacementPolicyDescriptionResponse struct {
	Type string `pulumi:"type"`
}





type ServicePlacementPolicyDescriptionResponseInput interface {
	pulumi.Input

	ToServicePlacementPolicyDescriptionResponseOutput() ServicePlacementPolicyDescriptionResponseOutput
	ToServicePlacementPolicyDescriptionResponseOutputWithContext(context.Context) ServicePlacementPolicyDescriptionResponseOutput
}

type ServicePlacementPolicyDescriptionResponseArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
}

func (ServicePlacementPolicyDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementPolicyDescriptionResponse)(nil)).Elem()
}

func (i ServicePlacementPolicyDescriptionResponseArgs) ToServicePlacementPolicyDescriptionResponseOutput() ServicePlacementPolicyDescriptionResponseOutput {
	return i.ToServicePlacementPolicyDescriptionResponseOutputWithContext(context.Background())
}

func (i ServicePlacementPolicyDescriptionResponseArgs) ToServicePlacementPolicyDescriptionResponseOutputWithContext(ctx context.Context) ServicePlacementPolicyDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePlacementPolicyDescriptionResponseOutput)
}





type ServicePlacementPolicyDescriptionResponseArrayInput interface {
	pulumi.Input

	ToServicePlacementPolicyDescriptionResponseArrayOutput() ServicePlacementPolicyDescriptionResponseArrayOutput
	ToServicePlacementPolicyDescriptionResponseArrayOutputWithContext(context.Context) ServicePlacementPolicyDescriptionResponseArrayOutput
}

type ServicePlacementPolicyDescriptionResponseArray []ServicePlacementPolicyDescriptionResponseInput

func (ServicePlacementPolicyDescriptionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServicePlacementPolicyDescriptionResponse)(nil)).Elem()
}

func (i ServicePlacementPolicyDescriptionResponseArray) ToServicePlacementPolicyDescriptionResponseArrayOutput() ServicePlacementPolicyDescriptionResponseArrayOutput {
	return i.ToServicePlacementPolicyDescriptionResponseArrayOutputWithContext(context.Background())
}

func (i ServicePlacementPolicyDescriptionResponseArray) ToServicePlacementPolicyDescriptionResponseArrayOutputWithContext(ctx context.Context) ServicePlacementPolicyDescriptionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePlacementPolicyDescriptionResponseArrayOutput)
}

type ServicePlacementPolicyDescriptionResponseOutput struct{ *pulumi.OutputState }

func (ServicePlacementPolicyDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementPolicyDescriptionResponse)(nil)).Elem()
}

func (o ServicePlacementPolicyDescriptionResponseOutput) ToServicePlacementPolicyDescriptionResponseOutput() ServicePlacementPolicyDescriptionResponseOutput {
	return o
}

func (o ServicePlacementPolicyDescriptionResponseOutput) ToServicePlacementPolicyDescriptionResponseOutputWithContext(ctx context.Context) ServicePlacementPolicyDescriptionResponseOutput {
	return o
}

func (o ServicePlacementPolicyDescriptionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementPolicyDescriptionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ServicePlacementPolicyDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (ServicePlacementPolicyDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServicePlacementPolicyDescriptionResponse)(nil)).Elem()
}

func (o ServicePlacementPolicyDescriptionResponseArrayOutput) ToServicePlacementPolicyDescriptionResponseArrayOutput() ServicePlacementPolicyDescriptionResponseArrayOutput {
	return o
}

func (o ServicePlacementPolicyDescriptionResponseArrayOutput) ToServicePlacementPolicyDescriptionResponseArrayOutputWithContext(ctx context.Context) ServicePlacementPolicyDescriptionResponseArrayOutput {
	return o
}

func (o ServicePlacementPolicyDescriptionResponseArrayOutput) Index(i pulumi.IntInput) ServicePlacementPolicyDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServicePlacementPolicyDescriptionResponse {
		return vs[0].([]ServicePlacementPolicyDescriptionResponse)[vs[1].(int)]
	}).(ServicePlacementPolicyDescriptionResponseOutput)
}

type ServiceTypeDeltaHealthPolicy struct {
	MaxPercentDeltaUnhealthyServices *int `pulumi:"maxPercentDeltaUnhealthyServices"`
}


func (val *ServiceTypeDeltaHealthPolicy) Defaults() *ServiceTypeDeltaHealthPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxPercentDeltaUnhealthyServices) {
		maxPercentDeltaUnhealthyServices_ := 0
		tmp.MaxPercentDeltaUnhealthyServices = &maxPercentDeltaUnhealthyServices_
	}
	return &tmp
}





type ServiceTypeDeltaHealthPolicyInput interface {
	pulumi.Input

	ToServiceTypeDeltaHealthPolicyOutput() ServiceTypeDeltaHealthPolicyOutput
	ToServiceTypeDeltaHealthPolicyOutputWithContext(context.Context) ServiceTypeDeltaHealthPolicyOutput
}

type ServiceTypeDeltaHealthPolicyArgs struct {
	MaxPercentDeltaUnhealthyServices pulumi.IntPtrInput `pulumi:"maxPercentDeltaUnhealthyServices"`
}

func (ServiceTypeDeltaHealthPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTypeDeltaHealthPolicy)(nil)).Elem()
}

func (i ServiceTypeDeltaHealthPolicyArgs) ToServiceTypeDeltaHealthPolicyOutput() ServiceTypeDeltaHealthPolicyOutput {
	return i.ToServiceTypeDeltaHealthPolicyOutputWithContext(context.Background())
}

func (i ServiceTypeDeltaHealthPolicyArgs) ToServiceTypeDeltaHealthPolicyOutputWithContext(ctx context.Context) ServiceTypeDeltaHealthPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeDeltaHealthPolicyOutput)
}

func (i ServiceTypeDeltaHealthPolicyArgs) ToServiceTypeDeltaHealthPolicyPtrOutput() ServiceTypeDeltaHealthPolicyPtrOutput {
	return i.ToServiceTypeDeltaHealthPolicyPtrOutputWithContext(context.Background())
}

func (i ServiceTypeDeltaHealthPolicyArgs) ToServiceTypeDeltaHealthPolicyPtrOutputWithContext(ctx context.Context) ServiceTypeDeltaHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeDeltaHealthPolicyOutput).ToServiceTypeDeltaHealthPolicyPtrOutputWithContext(ctx)
}









type ServiceTypeDeltaHealthPolicyPtrInput interface {
	pulumi.Input

	ToServiceTypeDeltaHealthPolicyPtrOutput() ServiceTypeDeltaHealthPolicyPtrOutput
	ToServiceTypeDeltaHealthPolicyPtrOutputWithContext(context.Context) ServiceTypeDeltaHealthPolicyPtrOutput
}

type serviceTypeDeltaHealthPolicyPtrType ServiceTypeDeltaHealthPolicyArgs

func ServiceTypeDeltaHealthPolicyPtr(v *ServiceTypeDeltaHealthPolicyArgs) ServiceTypeDeltaHealthPolicyPtrInput {
	return (*serviceTypeDeltaHealthPolicyPtrType)(v)
}

func (*serviceTypeDeltaHealthPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTypeDeltaHealthPolicy)(nil)).Elem()
}

func (i *serviceTypeDeltaHealthPolicyPtrType) ToServiceTypeDeltaHealthPolicyPtrOutput() ServiceTypeDeltaHealthPolicyPtrOutput {
	return i.ToServiceTypeDeltaHealthPolicyPtrOutputWithContext(context.Background())
}

func (i *serviceTypeDeltaHealthPolicyPtrType) ToServiceTypeDeltaHealthPolicyPtrOutputWithContext(ctx context.Context) ServiceTypeDeltaHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeDeltaHealthPolicyPtrOutput)
}





type ServiceTypeDeltaHealthPolicyMapInput interface {
	pulumi.Input

	ToServiceTypeDeltaHealthPolicyMapOutput() ServiceTypeDeltaHealthPolicyMapOutput
	ToServiceTypeDeltaHealthPolicyMapOutputWithContext(context.Context) ServiceTypeDeltaHealthPolicyMapOutput
}

type ServiceTypeDeltaHealthPolicyMap map[string]ServiceTypeDeltaHealthPolicyInput

func (ServiceTypeDeltaHealthPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceTypeDeltaHealthPolicy)(nil)).Elem()
}

func (i ServiceTypeDeltaHealthPolicyMap) ToServiceTypeDeltaHealthPolicyMapOutput() ServiceTypeDeltaHealthPolicyMapOutput {
	return i.ToServiceTypeDeltaHealthPolicyMapOutputWithContext(context.Background())
}

func (i ServiceTypeDeltaHealthPolicyMap) ToServiceTypeDeltaHealthPolicyMapOutputWithContext(ctx context.Context) ServiceTypeDeltaHealthPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeDeltaHealthPolicyMapOutput)
}

type ServiceTypeDeltaHealthPolicyOutput struct{ *pulumi.OutputState }

func (ServiceTypeDeltaHealthPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTypeDeltaHealthPolicy)(nil)).Elem()
}

func (o ServiceTypeDeltaHealthPolicyOutput) ToServiceTypeDeltaHealthPolicyOutput() ServiceTypeDeltaHealthPolicyOutput {
	return o
}

func (o ServiceTypeDeltaHealthPolicyOutput) ToServiceTypeDeltaHealthPolicyOutputWithContext(ctx context.Context) ServiceTypeDeltaHealthPolicyOutput {
	return o
}

func (o ServiceTypeDeltaHealthPolicyOutput) ToServiceTypeDeltaHealthPolicyPtrOutput() ServiceTypeDeltaHealthPolicyPtrOutput {
	return o.ToServiceTypeDeltaHealthPolicyPtrOutputWithContext(context.Background())
}

func (o ServiceTypeDeltaHealthPolicyOutput) ToServiceTypeDeltaHealthPolicyPtrOutputWithContext(ctx context.Context) ServiceTypeDeltaHealthPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceTypeDeltaHealthPolicy) *ServiceTypeDeltaHealthPolicy {
		return &v
	}).(ServiceTypeDeltaHealthPolicyPtrOutput)
}

func (o ServiceTypeDeltaHealthPolicyOutput) MaxPercentDeltaUnhealthyServices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceTypeDeltaHealthPolicy) *int { return v.MaxPercentDeltaUnhealthyServices }).(pulumi.IntPtrOutput)
}

type ServiceTypeDeltaHealthPolicyPtrOutput struct{ *pulumi.OutputState }

func (ServiceTypeDeltaHealthPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTypeDeltaHealthPolicy)(nil)).Elem()
}

func (o ServiceTypeDeltaHealthPolicyPtrOutput) ToServiceTypeDeltaHealthPolicyPtrOutput() ServiceTypeDeltaHealthPolicyPtrOutput {
	return o
}

func (o ServiceTypeDeltaHealthPolicyPtrOutput) ToServiceTypeDeltaHealthPolicyPtrOutputWithContext(ctx context.Context) ServiceTypeDeltaHealthPolicyPtrOutput {
	return o
}

func (o ServiceTypeDeltaHealthPolicyPtrOutput) Elem() ServiceTypeDeltaHealthPolicyOutput {
	return o.ApplyT(func(v *ServiceTypeDeltaHealthPolicy) ServiceTypeDeltaHealthPolicy {
		if v != nil {
			return *v
		}
		var ret ServiceTypeDeltaHealthPolicy
		return ret
	}).(ServiceTypeDeltaHealthPolicyOutput)
}

func (o ServiceTypeDeltaHealthPolicyPtrOutput) MaxPercentDeltaUnhealthyServices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceTypeDeltaHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentDeltaUnhealthyServices
	}).(pulumi.IntPtrOutput)
}

type ServiceTypeDeltaHealthPolicyMapOutput struct{ *pulumi.OutputState }

func (ServiceTypeDeltaHealthPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceTypeDeltaHealthPolicy)(nil)).Elem()
}

func (o ServiceTypeDeltaHealthPolicyMapOutput) ToServiceTypeDeltaHealthPolicyMapOutput() ServiceTypeDeltaHealthPolicyMapOutput {
	return o
}

func (o ServiceTypeDeltaHealthPolicyMapOutput) ToServiceTypeDeltaHealthPolicyMapOutputWithContext(ctx context.Context) ServiceTypeDeltaHealthPolicyMapOutput {
	return o
}

func (o ServiceTypeDeltaHealthPolicyMapOutput) MapIndex(k pulumi.StringInput) ServiceTypeDeltaHealthPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServiceTypeDeltaHealthPolicy {
		return vs[0].(map[string]ServiceTypeDeltaHealthPolicy)[vs[1].(string)]
	}).(ServiceTypeDeltaHealthPolicyOutput)
}

type ServiceTypeDeltaHealthPolicyResponse struct {
	MaxPercentDeltaUnhealthyServices *int `pulumi:"maxPercentDeltaUnhealthyServices"`
}


func (val *ServiceTypeDeltaHealthPolicyResponse) Defaults() *ServiceTypeDeltaHealthPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxPercentDeltaUnhealthyServices) {
		maxPercentDeltaUnhealthyServices_ := 0
		tmp.MaxPercentDeltaUnhealthyServices = &maxPercentDeltaUnhealthyServices_
	}
	return &tmp
}





type ServiceTypeDeltaHealthPolicyResponseInput interface {
	pulumi.Input

	ToServiceTypeDeltaHealthPolicyResponseOutput() ServiceTypeDeltaHealthPolicyResponseOutput
	ToServiceTypeDeltaHealthPolicyResponseOutputWithContext(context.Context) ServiceTypeDeltaHealthPolicyResponseOutput
}

type ServiceTypeDeltaHealthPolicyResponseArgs struct {
	MaxPercentDeltaUnhealthyServices pulumi.IntPtrInput `pulumi:"maxPercentDeltaUnhealthyServices"`
}

func (ServiceTypeDeltaHealthPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTypeDeltaHealthPolicyResponse)(nil)).Elem()
}

func (i ServiceTypeDeltaHealthPolicyResponseArgs) ToServiceTypeDeltaHealthPolicyResponseOutput() ServiceTypeDeltaHealthPolicyResponseOutput {
	return i.ToServiceTypeDeltaHealthPolicyResponseOutputWithContext(context.Background())
}

func (i ServiceTypeDeltaHealthPolicyResponseArgs) ToServiceTypeDeltaHealthPolicyResponseOutputWithContext(ctx context.Context) ServiceTypeDeltaHealthPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeDeltaHealthPolicyResponseOutput)
}

func (i ServiceTypeDeltaHealthPolicyResponseArgs) ToServiceTypeDeltaHealthPolicyResponsePtrOutput() ServiceTypeDeltaHealthPolicyResponsePtrOutput {
	return i.ToServiceTypeDeltaHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (i ServiceTypeDeltaHealthPolicyResponseArgs) ToServiceTypeDeltaHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ServiceTypeDeltaHealthPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeDeltaHealthPolicyResponseOutput).ToServiceTypeDeltaHealthPolicyResponsePtrOutputWithContext(ctx)
}









type ServiceTypeDeltaHealthPolicyResponsePtrInput interface {
	pulumi.Input

	ToServiceTypeDeltaHealthPolicyResponsePtrOutput() ServiceTypeDeltaHealthPolicyResponsePtrOutput
	ToServiceTypeDeltaHealthPolicyResponsePtrOutputWithContext(context.Context) ServiceTypeDeltaHealthPolicyResponsePtrOutput
}

type serviceTypeDeltaHealthPolicyResponsePtrType ServiceTypeDeltaHealthPolicyResponseArgs

func ServiceTypeDeltaHealthPolicyResponsePtr(v *ServiceTypeDeltaHealthPolicyResponseArgs) ServiceTypeDeltaHealthPolicyResponsePtrInput {
	return (*serviceTypeDeltaHealthPolicyResponsePtrType)(v)
}

func (*serviceTypeDeltaHealthPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTypeDeltaHealthPolicyResponse)(nil)).Elem()
}

func (i *serviceTypeDeltaHealthPolicyResponsePtrType) ToServiceTypeDeltaHealthPolicyResponsePtrOutput() ServiceTypeDeltaHealthPolicyResponsePtrOutput {
	return i.ToServiceTypeDeltaHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *serviceTypeDeltaHealthPolicyResponsePtrType) ToServiceTypeDeltaHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ServiceTypeDeltaHealthPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeDeltaHealthPolicyResponsePtrOutput)
}





type ServiceTypeDeltaHealthPolicyResponseMapInput interface {
	pulumi.Input

	ToServiceTypeDeltaHealthPolicyResponseMapOutput() ServiceTypeDeltaHealthPolicyResponseMapOutput
	ToServiceTypeDeltaHealthPolicyResponseMapOutputWithContext(context.Context) ServiceTypeDeltaHealthPolicyResponseMapOutput
}

type ServiceTypeDeltaHealthPolicyResponseMap map[string]ServiceTypeDeltaHealthPolicyResponseInput

func (ServiceTypeDeltaHealthPolicyResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceTypeDeltaHealthPolicyResponse)(nil)).Elem()
}

func (i ServiceTypeDeltaHealthPolicyResponseMap) ToServiceTypeDeltaHealthPolicyResponseMapOutput() ServiceTypeDeltaHealthPolicyResponseMapOutput {
	return i.ToServiceTypeDeltaHealthPolicyResponseMapOutputWithContext(context.Background())
}

func (i ServiceTypeDeltaHealthPolicyResponseMap) ToServiceTypeDeltaHealthPolicyResponseMapOutputWithContext(ctx context.Context) ServiceTypeDeltaHealthPolicyResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeDeltaHealthPolicyResponseMapOutput)
}

type ServiceTypeDeltaHealthPolicyResponseOutput struct{ *pulumi.OutputState }

func (ServiceTypeDeltaHealthPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTypeDeltaHealthPolicyResponse)(nil)).Elem()
}

func (o ServiceTypeDeltaHealthPolicyResponseOutput) ToServiceTypeDeltaHealthPolicyResponseOutput() ServiceTypeDeltaHealthPolicyResponseOutput {
	return o
}

func (o ServiceTypeDeltaHealthPolicyResponseOutput) ToServiceTypeDeltaHealthPolicyResponseOutputWithContext(ctx context.Context) ServiceTypeDeltaHealthPolicyResponseOutput {
	return o
}

func (o ServiceTypeDeltaHealthPolicyResponseOutput) ToServiceTypeDeltaHealthPolicyResponsePtrOutput() ServiceTypeDeltaHealthPolicyResponsePtrOutput {
	return o.ToServiceTypeDeltaHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (o ServiceTypeDeltaHealthPolicyResponseOutput) ToServiceTypeDeltaHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ServiceTypeDeltaHealthPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceTypeDeltaHealthPolicyResponse) *ServiceTypeDeltaHealthPolicyResponse {
		return &v
	}).(ServiceTypeDeltaHealthPolicyResponsePtrOutput)
}

func (o ServiceTypeDeltaHealthPolicyResponseOutput) MaxPercentDeltaUnhealthyServices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceTypeDeltaHealthPolicyResponse) *int { return v.MaxPercentDeltaUnhealthyServices }).(pulumi.IntPtrOutput)
}

type ServiceTypeDeltaHealthPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceTypeDeltaHealthPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTypeDeltaHealthPolicyResponse)(nil)).Elem()
}

func (o ServiceTypeDeltaHealthPolicyResponsePtrOutput) ToServiceTypeDeltaHealthPolicyResponsePtrOutput() ServiceTypeDeltaHealthPolicyResponsePtrOutput {
	return o
}

func (o ServiceTypeDeltaHealthPolicyResponsePtrOutput) ToServiceTypeDeltaHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ServiceTypeDeltaHealthPolicyResponsePtrOutput {
	return o
}

func (o ServiceTypeDeltaHealthPolicyResponsePtrOutput) Elem() ServiceTypeDeltaHealthPolicyResponseOutput {
	return o.ApplyT(func(v *ServiceTypeDeltaHealthPolicyResponse) ServiceTypeDeltaHealthPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ServiceTypeDeltaHealthPolicyResponse
		return ret
	}).(ServiceTypeDeltaHealthPolicyResponseOutput)
}

func (o ServiceTypeDeltaHealthPolicyResponsePtrOutput) MaxPercentDeltaUnhealthyServices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceTypeDeltaHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentDeltaUnhealthyServices
	}).(pulumi.IntPtrOutput)
}

type ServiceTypeDeltaHealthPolicyResponseMapOutput struct{ *pulumi.OutputState }

func (ServiceTypeDeltaHealthPolicyResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceTypeDeltaHealthPolicyResponse)(nil)).Elem()
}

func (o ServiceTypeDeltaHealthPolicyResponseMapOutput) ToServiceTypeDeltaHealthPolicyResponseMapOutput() ServiceTypeDeltaHealthPolicyResponseMapOutput {
	return o
}

func (o ServiceTypeDeltaHealthPolicyResponseMapOutput) ToServiceTypeDeltaHealthPolicyResponseMapOutputWithContext(ctx context.Context) ServiceTypeDeltaHealthPolicyResponseMapOutput {
	return o
}

func (o ServiceTypeDeltaHealthPolicyResponseMapOutput) MapIndex(k pulumi.StringInput) ServiceTypeDeltaHealthPolicyResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServiceTypeDeltaHealthPolicyResponse {
		return vs[0].(map[string]ServiceTypeDeltaHealthPolicyResponse)[vs[1].(string)]
	}).(ServiceTypeDeltaHealthPolicyResponseOutput)
}

type ServiceTypeHealthPolicy struct {
	MaxPercentUnhealthyServices *int `pulumi:"maxPercentUnhealthyServices"`
}


func (val *ServiceTypeHealthPolicy) Defaults() *ServiceTypeHealthPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxPercentUnhealthyServices) {
		maxPercentUnhealthyServices_ := 0
		tmp.MaxPercentUnhealthyServices = &maxPercentUnhealthyServices_
	}
	return &tmp
}





type ServiceTypeHealthPolicyInput interface {
	pulumi.Input

	ToServiceTypeHealthPolicyOutput() ServiceTypeHealthPolicyOutput
	ToServiceTypeHealthPolicyOutputWithContext(context.Context) ServiceTypeHealthPolicyOutput
}

type ServiceTypeHealthPolicyArgs struct {
	MaxPercentUnhealthyServices pulumi.IntPtrInput `pulumi:"maxPercentUnhealthyServices"`
}

func (ServiceTypeHealthPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTypeHealthPolicy)(nil)).Elem()
}

func (i ServiceTypeHealthPolicyArgs) ToServiceTypeHealthPolicyOutput() ServiceTypeHealthPolicyOutput {
	return i.ToServiceTypeHealthPolicyOutputWithContext(context.Background())
}

func (i ServiceTypeHealthPolicyArgs) ToServiceTypeHealthPolicyOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeHealthPolicyOutput)
}

func (i ServiceTypeHealthPolicyArgs) ToServiceTypeHealthPolicyPtrOutput() ServiceTypeHealthPolicyPtrOutput {
	return i.ToServiceTypeHealthPolicyPtrOutputWithContext(context.Background())
}

func (i ServiceTypeHealthPolicyArgs) ToServiceTypeHealthPolicyPtrOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeHealthPolicyOutput).ToServiceTypeHealthPolicyPtrOutputWithContext(ctx)
}









type ServiceTypeHealthPolicyPtrInput interface {
	pulumi.Input

	ToServiceTypeHealthPolicyPtrOutput() ServiceTypeHealthPolicyPtrOutput
	ToServiceTypeHealthPolicyPtrOutputWithContext(context.Context) ServiceTypeHealthPolicyPtrOutput
}

type serviceTypeHealthPolicyPtrType ServiceTypeHealthPolicyArgs

func ServiceTypeHealthPolicyPtr(v *ServiceTypeHealthPolicyArgs) ServiceTypeHealthPolicyPtrInput {
	return (*serviceTypeHealthPolicyPtrType)(v)
}

func (*serviceTypeHealthPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTypeHealthPolicy)(nil)).Elem()
}

func (i *serviceTypeHealthPolicyPtrType) ToServiceTypeHealthPolicyPtrOutput() ServiceTypeHealthPolicyPtrOutput {
	return i.ToServiceTypeHealthPolicyPtrOutputWithContext(context.Background())
}

func (i *serviceTypeHealthPolicyPtrType) ToServiceTypeHealthPolicyPtrOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeHealthPolicyPtrOutput)
}





type ServiceTypeHealthPolicyMapInput interface {
	pulumi.Input

	ToServiceTypeHealthPolicyMapOutput() ServiceTypeHealthPolicyMapOutput
	ToServiceTypeHealthPolicyMapOutputWithContext(context.Context) ServiceTypeHealthPolicyMapOutput
}

type ServiceTypeHealthPolicyMap map[string]ServiceTypeHealthPolicyInput

func (ServiceTypeHealthPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceTypeHealthPolicy)(nil)).Elem()
}

func (i ServiceTypeHealthPolicyMap) ToServiceTypeHealthPolicyMapOutput() ServiceTypeHealthPolicyMapOutput {
	return i.ToServiceTypeHealthPolicyMapOutputWithContext(context.Background())
}

func (i ServiceTypeHealthPolicyMap) ToServiceTypeHealthPolicyMapOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeHealthPolicyMapOutput)
}

type ServiceTypeHealthPolicyOutput struct{ *pulumi.OutputState }

func (ServiceTypeHealthPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTypeHealthPolicy)(nil)).Elem()
}

func (o ServiceTypeHealthPolicyOutput) ToServiceTypeHealthPolicyOutput() ServiceTypeHealthPolicyOutput {
	return o
}

func (o ServiceTypeHealthPolicyOutput) ToServiceTypeHealthPolicyOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyOutput {
	return o
}

func (o ServiceTypeHealthPolicyOutput) ToServiceTypeHealthPolicyPtrOutput() ServiceTypeHealthPolicyPtrOutput {
	return o.ToServiceTypeHealthPolicyPtrOutputWithContext(context.Background())
}

func (o ServiceTypeHealthPolicyOutput) ToServiceTypeHealthPolicyPtrOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceTypeHealthPolicy) *ServiceTypeHealthPolicy {
		return &v
	}).(ServiceTypeHealthPolicyPtrOutput)
}

func (o ServiceTypeHealthPolicyOutput) MaxPercentUnhealthyServices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceTypeHealthPolicy) *int { return v.MaxPercentUnhealthyServices }).(pulumi.IntPtrOutput)
}

type ServiceTypeHealthPolicyPtrOutput struct{ *pulumi.OutputState }

func (ServiceTypeHealthPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTypeHealthPolicy)(nil)).Elem()
}

func (o ServiceTypeHealthPolicyPtrOutput) ToServiceTypeHealthPolicyPtrOutput() ServiceTypeHealthPolicyPtrOutput {
	return o
}

func (o ServiceTypeHealthPolicyPtrOutput) ToServiceTypeHealthPolicyPtrOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyPtrOutput {
	return o
}

func (o ServiceTypeHealthPolicyPtrOutput) Elem() ServiceTypeHealthPolicyOutput {
	return o.ApplyT(func(v *ServiceTypeHealthPolicy) ServiceTypeHealthPolicy {
		if v != nil {
			return *v
		}
		var ret ServiceTypeHealthPolicy
		return ret
	}).(ServiceTypeHealthPolicyOutput)
}

func (o ServiceTypeHealthPolicyPtrOutput) MaxPercentUnhealthyServices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceTypeHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentUnhealthyServices
	}).(pulumi.IntPtrOutput)
}

type ServiceTypeHealthPolicyMapOutput struct{ *pulumi.OutputState }

func (ServiceTypeHealthPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceTypeHealthPolicy)(nil)).Elem()
}

func (o ServiceTypeHealthPolicyMapOutput) ToServiceTypeHealthPolicyMapOutput() ServiceTypeHealthPolicyMapOutput {
	return o
}

func (o ServiceTypeHealthPolicyMapOutput) ToServiceTypeHealthPolicyMapOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyMapOutput {
	return o
}

func (o ServiceTypeHealthPolicyMapOutput) MapIndex(k pulumi.StringInput) ServiceTypeHealthPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServiceTypeHealthPolicy {
		return vs[0].(map[string]ServiceTypeHealthPolicy)[vs[1].(string)]
	}).(ServiceTypeHealthPolicyOutput)
}

type ServiceTypeHealthPolicyResponse struct {
	MaxPercentUnhealthyServices *int `pulumi:"maxPercentUnhealthyServices"`
}


func (val *ServiceTypeHealthPolicyResponse) Defaults() *ServiceTypeHealthPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxPercentUnhealthyServices) {
		maxPercentUnhealthyServices_ := 0
		tmp.MaxPercentUnhealthyServices = &maxPercentUnhealthyServices_
	}
	return &tmp
}





type ServiceTypeHealthPolicyResponseInput interface {
	pulumi.Input

	ToServiceTypeHealthPolicyResponseOutput() ServiceTypeHealthPolicyResponseOutput
	ToServiceTypeHealthPolicyResponseOutputWithContext(context.Context) ServiceTypeHealthPolicyResponseOutput
}

type ServiceTypeHealthPolicyResponseArgs struct {
	MaxPercentUnhealthyServices pulumi.IntPtrInput `pulumi:"maxPercentUnhealthyServices"`
}

func (ServiceTypeHealthPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTypeHealthPolicyResponse)(nil)).Elem()
}

func (i ServiceTypeHealthPolicyResponseArgs) ToServiceTypeHealthPolicyResponseOutput() ServiceTypeHealthPolicyResponseOutput {
	return i.ToServiceTypeHealthPolicyResponseOutputWithContext(context.Background())
}

func (i ServiceTypeHealthPolicyResponseArgs) ToServiceTypeHealthPolicyResponseOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeHealthPolicyResponseOutput)
}

func (i ServiceTypeHealthPolicyResponseArgs) ToServiceTypeHealthPolicyResponsePtrOutput() ServiceTypeHealthPolicyResponsePtrOutput {
	return i.ToServiceTypeHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (i ServiceTypeHealthPolicyResponseArgs) ToServiceTypeHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeHealthPolicyResponseOutput).ToServiceTypeHealthPolicyResponsePtrOutputWithContext(ctx)
}









type ServiceTypeHealthPolicyResponsePtrInput interface {
	pulumi.Input

	ToServiceTypeHealthPolicyResponsePtrOutput() ServiceTypeHealthPolicyResponsePtrOutput
	ToServiceTypeHealthPolicyResponsePtrOutputWithContext(context.Context) ServiceTypeHealthPolicyResponsePtrOutput
}

type serviceTypeHealthPolicyResponsePtrType ServiceTypeHealthPolicyResponseArgs

func ServiceTypeHealthPolicyResponsePtr(v *ServiceTypeHealthPolicyResponseArgs) ServiceTypeHealthPolicyResponsePtrInput {
	return (*serviceTypeHealthPolicyResponsePtrType)(v)
}

func (*serviceTypeHealthPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTypeHealthPolicyResponse)(nil)).Elem()
}

func (i *serviceTypeHealthPolicyResponsePtrType) ToServiceTypeHealthPolicyResponsePtrOutput() ServiceTypeHealthPolicyResponsePtrOutput {
	return i.ToServiceTypeHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *serviceTypeHealthPolicyResponsePtrType) ToServiceTypeHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeHealthPolicyResponsePtrOutput)
}





type ServiceTypeHealthPolicyResponseMapInput interface {
	pulumi.Input

	ToServiceTypeHealthPolicyResponseMapOutput() ServiceTypeHealthPolicyResponseMapOutput
	ToServiceTypeHealthPolicyResponseMapOutputWithContext(context.Context) ServiceTypeHealthPolicyResponseMapOutput
}

type ServiceTypeHealthPolicyResponseMap map[string]ServiceTypeHealthPolicyResponseInput

func (ServiceTypeHealthPolicyResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceTypeHealthPolicyResponse)(nil)).Elem()
}

func (i ServiceTypeHealthPolicyResponseMap) ToServiceTypeHealthPolicyResponseMapOutput() ServiceTypeHealthPolicyResponseMapOutput {
	return i.ToServiceTypeHealthPolicyResponseMapOutputWithContext(context.Background())
}

func (i ServiceTypeHealthPolicyResponseMap) ToServiceTypeHealthPolicyResponseMapOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeHealthPolicyResponseMapOutput)
}

type ServiceTypeHealthPolicyResponseOutput struct{ *pulumi.OutputState }

func (ServiceTypeHealthPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTypeHealthPolicyResponse)(nil)).Elem()
}

func (o ServiceTypeHealthPolicyResponseOutput) ToServiceTypeHealthPolicyResponseOutput() ServiceTypeHealthPolicyResponseOutput {
	return o
}

func (o ServiceTypeHealthPolicyResponseOutput) ToServiceTypeHealthPolicyResponseOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyResponseOutput {
	return o
}

func (o ServiceTypeHealthPolicyResponseOutput) ToServiceTypeHealthPolicyResponsePtrOutput() ServiceTypeHealthPolicyResponsePtrOutput {
	return o.ToServiceTypeHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (o ServiceTypeHealthPolicyResponseOutput) ToServiceTypeHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceTypeHealthPolicyResponse) *ServiceTypeHealthPolicyResponse {
		return &v
	}).(ServiceTypeHealthPolicyResponsePtrOutput)
}

func (o ServiceTypeHealthPolicyResponseOutput) MaxPercentUnhealthyServices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceTypeHealthPolicyResponse) *int { return v.MaxPercentUnhealthyServices }).(pulumi.IntPtrOutput)
}

type ServiceTypeHealthPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceTypeHealthPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTypeHealthPolicyResponse)(nil)).Elem()
}

func (o ServiceTypeHealthPolicyResponsePtrOutput) ToServiceTypeHealthPolicyResponsePtrOutput() ServiceTypeHealthPolicyResponsePtrOutput {
	return o
}

func (o ServiceTypeHealthPolicyResponsePtrOutput) ToServiceTypeHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyResponsePtrOutput {
	return o
}

func (o ServiceTypeHealthPolicyResponsePtrOutput) Elem() ServiceTypeHealthPolicyResponseOutput {
	return o.ApplyT(func(v *ServiceTypeHealthPolicyResponse) ServiceTypeHealthPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ServiceTypeHealthPolicyResponse
		return ret
	}).(ServiceTypeHealthPolicyResponseOutput)
}

func (o ServiceTypeHealthPolicyResponsePtrOutput) MaxPercentUnhealthyServices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceTypeHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentUnhealthyServices
	}).(pulumi.IntPtrOutput)
}

type ServiceTypeHealthPolicyResponseMapOutput struct{ *pulumi.OutputState }

func (ServiceTypeHealthPolicyResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceTypeHealthPolicyResponse)(nil)).Elem()
}

func (o ServiceTypeHealthPolicyResponseMapOutput) ToServiceTypeHealthPolicyResponseMapOutput() ServiceTypeHealthPolicyResponseMapOutput {
	return o
}

func (o ServiceTypeHealthPolicyResponseMapOutput) ToServiceTypeHealthPolicyResponseMapOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyResponseMapOutput {
	return o
}

func (o ServiceTypeHealthPolicyResponseMapOutput) MapIndex(k pulumi.StringInput) ServiceTypeHealthPolicyResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServiceTypeHealthPolicyResponse {
		return vs[0].(map[string]ServiceTypeHealthPolicyResponse)[vs[1].(string)]
	}).(ServiceTypeHealthPolicyResponseOutput)
}

type SettingsParameterDescription struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type SettingsParameterDescriptionInput interface {
	pulumi.Input

	ToSettingsParameterDescriptionOutput() SettingsParameterDescriptionOutput
	ToSettingsParameterDescriptionOutputWithContext(context.Context) SettingsParameterDescriptionOutput
}

type SettingsParameterDescriptionArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (SettingsParameterDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsParameterDescription)(nil)).Elem()
}

func (i SettingsParameterDescriptionArgs) ToSettingsParameterDescriptionOutput() SettingsParameterDescriptionOutput {
	return i.ToSettingsParameterDescriptionOutputWithContext(context.Background())
}

func (i SettingsParameterDescriptionArgs) ToSettingsParameterDescriptionOutputWithContext(ctx context.Context) SettingsParameterDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsParameterDescriptionOutput)
}





type SettingsParameterDescriptionArrayInput interface {
	pulumi.Input

	ToSettingsParameterDescriptionArrayOutput() SettingsParameterDescriptionArrayOutput
	ToSettingsParameterDescriptionArrayOutputWithContext(context.Context) SettingsParameterDescriptionArrayOutput
}

type SettingsParameterDescriptionArray []SettingsParameterDescriptionInput

func (SettingsParameterDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsParameterDescription)(nil)).Elem()
}

func (i SettingsParameterDescriptionArray) ToSettingsParameterDescriptionArrayOutput() SettingsParameterDescriptionArrayOutput {
	return i.ToSettingsParameterDescriptionArrayOutputWithContext(context.Background())
}

func (i SettingsParameterDescriptionArray) ToSettingsParameterDescriptionArrayOutputWithContext(ctx context.Context) SettingsParameterDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsParameterDescriptionArrayOutput)
}

type SettingsParameterDescriptionOutput struct{ *pulumi.OutputState }

func (SettingsParameterDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsParameterDescription)(nil)).Elem()
}

func (o SettingsParameterDescriptionOutput) ToSettingsParameterDescriptionOutput() SettingsParameterDescriptionOutput {
	return o
}

func (o SettingsParameterDescriptionOutput) ToSettingsParameterDescriptionOutputWithContext(ctx context.Context) SettingsParameterDescriptionOutput {
	return o
}

func (o SettingsParameterDescriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsParameterDescription) string { return v.Name }).(pulumi.StringOutput)
}

func (o SettingsParameterDescriptionOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsParameterDescription) string { return v.Value }).(pulumi.StringOutput)
}

type SettingsParameterDescriptionArrayOutput struct{ *pulumi.OutputState }

func (SettingsParameterDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsParameterDescription)(nil)).Elem()
}

func (o SettingsParameterDescriptionArrayOutput) ToSettingsParameterDescriptionArrayOutput() SettingsParameterDescriptionArrayOutput {
	return o
}

func (o SettingsParameterDescriptionArrayOutput) ToSettingsParameterDescriptionArrayOutputWithContext(ctx context.Context) SettingsParameterDescriptionArrayOutput {
	return o
}

func (o SettingsParameterDescriptionArrayOutput) Index(i pulumi.IntInput) SettingsParameterDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingsParameterDescription {
		return vs[0].([]SettingsParameterDescription)[vs[1].(int)]
	}).(SettingsParameterDescriptionOutput)
}

type SettingsParameterDescriptionResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type SettingsParameterDescriptionResponseInput interface {
	pulumi.Input

	ToSettingsParameterDescriptionResponseOutput() SettingsParameterDescriptionResponseOutput
	ToSettingsParameterDescriptionResponseOutputWithContext(context.Context) SettingsParameterDescriptionResponseOutput
}

type SettingsParameterDescriptionResponseArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (SettingsParameterDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsParameterDescriptionResponse)(nil)).Elem()
}

func (i SettingsParameterDescriptionResponseArgs) ToSettingsParameterDescriptionResponseOutput() SettingsParameterDescriptionResponseOutput {
	return i.ToSettingsParameterDescriptionResponseOutputWithContext(context.Background())
}

func (i SettingsParameterDescriptionResponseArgs) ToSettingsParameterDescriptionResponseOutputWithContext(ctx context.Context) SettingsParameterDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsParameterDescriptionResponseOutput)
}





type SettingsParameterDescriptionResponseArrayInput interface {
	pulumi.Input

	ToSettingsParameterDescriptionResponseArrayOutput() SettingsParameterDescriptionResponseArrayOutput
	ToSettingsParameterDescriptionResponseArrayOutputWithContext(context.Context) SettingsParameterDescriptionResponseArrayOutput
}

type SettingsParameterDescriptionResponseArray []SettingsParameterDescriptionResponseInput

func (SettingsParameterDescriptionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsParameterDescriptionResponse)(nil)).Elem()
}

func (i SettingsParameterDescriptionResponseArray) ToSettingsParameterDescriptionResponseArrayOutput() SettingsParameterDescriptionResponseArrayOutput {
	return i.ToSettingsParameterDescriptionResponseArrayOutputWithContext(context.Background())
}

func (i SettingsParameterDescriptionResponseArray) ToSettingsParameterDescriptionResponseArrayOutputWithContext(ctx context.Context) SettingsParameterDescriptionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsParameterDescriptionResponseArrayOutput)
}

type SettingsParameterDescriptionResponseOutput struct{ *pulumi.OutputState }

func (SettingsParameterDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsParameterDescriptionResponse)(nil)).Elem()
}

func (o SettingsParameterDescriptionResponseOutput) ToSettingsParameterDescriptionResponseOutput() SettingsParameterDescriptionResponseOutput {
	return o
}

func (o SettingsParameterDescriptionResponseOutput) ToSettingsParameterDescriptionResponseOutputWithContext(ctx context.Context) SettingsParameterDescriptionResponseOutput {
	return o
}

func (o SettingsParameterDescriptionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsParameterDescriptionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SettingsParameterDescriptionResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsParameterDescriptionResponse) string { return v.Value }).(pulumi.StringOutput)
}

type SettingsParameterDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (SettingsParameterDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsParameterDescriptionResponse)(nil)).Elem()
}

func (o SettingsParameterDescriptionResponseArrayOutput) ToSettingsParameterDescriptionResponseArrayOutput() SettingsParameterDescriptionResponseArrayOutput {
	return o
}

func (o SettingsParameterDescriptionResponseArrayOutput) ToSettingsParameterDescriptionResponseArrayOutputWithContext(ctx context.Context) SettingsParameterDescriptionResponseArrayOutput {
	return o
}

func (o SettingsParameterDescriptionResponseArrayOutput) Index(i pulumi.IntInput) SettingsParameterDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingsParameterDescriptionResponse {
		return vs[0].([]SettingsParameterDescriptionResponse)[vs[1].(int)]
	}).(SettingsParameterDescriptionResponseOutput)
}

type SettingsSectionDescription struct {
	Name       string                         `pulumi:"name"`
	Parameters []SettingsParameterDescription `pulumi:"parameters"`
}





type SettingsSectionDescriptionInput interface {
	pulumi.Input

	ToSettingsSectionDescriptionOutput() SettingsSectionDescriptionOutput
	ToSettingsSectionDescriptionOutputWithContext(context.Context) SettingsSectionDescriptionOutput
}

type SettingsSectionDescriptionArgs struct {
	Name       pulumi.StringInput                     `pulumi:"name"`
	Parameters SettingsParameterDescriptionArrayInput `pulumi:"parameters"`
}

func (SettingsSectionDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsSectionDescription)(nil)).Elem()
}

func (i SettingsSectionDescriptionArgs) ToSettingsSectionDescriptionOutput() SettingsSectionDescriptionOutput {
	return i.ToSettingsSectionDescriptionOutputWithContext(context.Background())
}

func (i SettingsSectionDescriptionArgs) ToSettingsSectionDescriptionOutputWithContext(ctx context.Context) SettingsSectionDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsSectionDescriptionOutput)
}





type SettingsSectionDescriptionArrayInput interface {
	pulumi.Input

	ToSettingsSectionDescriptionArrayOutput() SettingsSectionDescriptionArrayOutput
	ToSettingsSectionDescriptionArrayOutputWithContext(context.Context) SettingsSectionDescriptionArrayOutput
}

type SettingsSectionDescriptionArray []SettingsSectionDescriptionInput

func (SettingsSectionDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsSectionDescription)(nil)).Elem()
}

func (i SettingsSectionDescriptionArray) ToSettingsSectionDescriptionArrayOutput() SettingsSectionDescriptionArrayOutput {
	return i.ToSettingsSectionDescriptionArrayOutputWithContext(context.Background())
}

func (i SettingsSectionDescriptionArray) ToSettingsSectionDescriptionArrayOutputWithContext(ctx context.Context) SettingsSectionDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsSectionDescriptionArrayOutput)
}

type SettingsSectionDescriptionOutput struct{ *pulumi.OutputState }

func (SettingsSectionDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsSectionDescription)(nil)).Elem()
}

func (o SettingsSectionDescriptionOutput) ToSettingsSectionDescriptionOutput() SettingsSectionDescriptionOutput {
	return o
}

func (o SettingsSectionDescriptionOutput) ToSettingsSectionDescriptionOutputWithContext(ctx context.Context) SettingsSectionDescriptionOutput {
	return o
}

func (o SettingsSectionDescriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsSectionDescription) string { return v.Name }).(pulumi.StringOutput)
}

func (o SettingsSectionDescriptionOutput) Parameters() SettingsParameterDescriptionArrayOutput {
	return o.ApplyT(func(v SettingsSectionDescription) []SettingsParameterDescription { return v.Parameters }).(SettingsParameterDescriptionArrayOutput)
}

type SettingsSectionDescriptionArrayOutput struct{ *pulumi.OutputState }

func (SettingsSectionDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsSectionDescription)(nil)).Elem()
}

func (o SettingsSectionDescriptionArrayOutput) ToSettingsSectionDescriptionArrayOutput() SettingsSectionDescriptionArrayOutput {
	return o
}

func (o SettingsSectionDescriptionArrayOutput) ToSettingsSectionDescriptionArrayOutputWithContext(ctx context.Context) SettingsSectionDescriptionArrayOutput {
	return o
}

func (o SettingsSectionDescriptionArrayOutput) Index(i pulumi.IntInput) SettingsSectionDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingsSectionDescription {
		return vs[0].([]SettingsSectionDescription)[vs[1].(int)]
	}).(SettingsSectionDescriptionOutput)
}

type SettingsSectionDescriptionResponse struct {
	Name       string                                 `pulumi:"name"`
	Parameters []SettingsParameterDescriptionResponse `pulumi:"parameters"`
}





type SettingsSectionDescriptionResponseInput interface {
	pulumi.Input

	ToSettingsSectionDescriptionResponseOutput() SettingsSectionDescriptionResponseOutput
	ToSettingsSectionDescriptionResponseOutputWithContext(context.Context) SettingsSectionDescriptionResponseOutput
}

type SettingsSectionDescriptionResponseArgs struct {
	Name       pulumi.StringInput                             `pulumi:"name"`
	Parameters SettingsParameterDescriptionResponseArrayInput `pulumi:"parameters"`
}

func (SettingsSectionDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsSectionDescriptionResponse)(nil)).Elem()
}

func (i SettingsSectionDescriptionResponseArgs) ToSettingsSectionDescriptionResponseOutput() SettingsSectionDescriptionResponseOutput {
	return i.ToSettingsSectionDescriptionResponseOutputWithContext(context.Background())
}

func (i SettingsSectionDescriptionResponseArgs) ToSettingsSectionDescriptionResponseOutputWithContext(ctx context.Context) SettingsSectionDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsSectionDescriptionResponseOutput)
}





type SettingsSectionDescriptionResponseArrayInput interface {
	pulumi.Input

	ToSettingsSectionDescriptionResponseArrayOutput() SettingsSectionDescriptionResponseArrayOutput
	ToSettingsSectionDescriptionResponseArrayOutputWithContext(context.Context) SettingsSectionDescriptionResponseArrayOutput
}

type SettingsSectionDescriptionResponseArray []SettingsSectionDescriptionResponseInput

func (SettingsSectionDescriptionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsSectionDescriptionResponse)(nil)).Elem()
}

func (i SettingsSectionDescriptionResponseArray) ToSettingsSectionDescriptionResponseArrayOutput() SettingsSectionDescriptionResponseArrayOutput {
	return i.ToSettingsSectionDescriptionResponseArrayOutputWithContext(context.Background())
}

func (i SettingsSectionDescriptionResponseArray) ToSettingsSectionDescriptionResponseArrayOutputWithContext(ctx context.Context) SettingsSectionDescriptionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsSectionDescriptionResponseArrayOutput)
}

type SettingsSectionDescriptionResponseOutput struct{ *pulumi.OutputState }

func (SettingsSectionDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsSectionDescriptionResponse)(nil)).Elem()
}

func (o SettingsSectionDescriptionResponseOutput) ToSettingsSectionDescriptionResponseOutput() SettingsSectionDescriptionResponseOutput {
	return o
}

func (o SettingsSectionDescriptionResponseOutput) ToSettingsSectionDescriptionResponseOutputWithContext(ctx context.Context) SettingsSectionDescriptionResponseOutput {
	return o
}

func (o SettingsSectionDescriptionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsSectionDescriptionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SettingsSectionDescriptionResponseOutput) Parameters() SettingsParameterDescriptionResponseArrayOutput {
	return o.ApplyT(func(v SettingsSectionDescriptionResponse) []SettingsParameterDescriptionResponse { return v.Parameters }).(SettingsParameterDescriptionResponseArrayOutput)
}

type SettingsSectionDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (SettingsSectionDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsSectionDescriptionResponse)(nil)).Elem()
}

func (o SettingsSectionDescriptionResponseArrayOutput) ToSettingsSectionDescriptionResponseArrayOutput() SettingsSectionDescriptionResponseArrayOutput {
	return o
}

func (o SettingsSectionDescriptionResponseArrayOutput) ToSettingsSectionDescriptionResponseArrayOutputWithContext(ctx context.Context) SettingsSectionDescriptionResponseArrayOutput {
	return o
}

func (o SettingsSectionDescriptionResponseArrayOutput) Index(i pulumi.IntInput) SettingsSectionDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingsSectionDescriptionResponse {
		return vs[0].([]SettingsSectionDescriptionResponse)[vs[1].(int)]
	}).(SettingsSectionDescriptionResponseOutput)
}

type SingletonPartitionSchemeDescription struct {
	PartitionScheme string `pulumi:"partitionScheme"`
}





type SingletonPartitionSchemeDescriptionInput interface {
	pulumi.Input

	ToSingletonPartitionSchemeDescriptionOutput() SingletonPartitionSchemeDescriptionOutput
	ToSingletonPartitionSchemeDescriptionOutputWithContext(context.Context) SingletonPartitionSchemeDescriptionOutput
}

type SingletonPartitionSchemeDescriptionArgs struct {
	PartitionScheme pulumi.StringInput `pulumi:"partitionScheme"`
}

func (SingletonPartitionSchemeDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SingletonPartitionSchemeDescription)(nil)).Elem()
}

func (i SingletonPartitionSchemeDescriptionArgs) ToSingletonPartitionSchemeDescriptionOutput() SingletonPartitionSchemeDescriptionOutput {
	return i.ToSingletonPartitionSchemeDescriptionOutputWithContext(context.Background())
}

func (i SingletonPartitionSchemeDescriptionArgs) ToSingletonPartitionSchemeDescriptionOutputWithContext(ctx context.Context) SingletonPartitionSchemeDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SingletonPartitionSchemeDescriptionOutput)
}

type SingletonPartitionSchemeDescriptionOutput struct{ *pulumi.OutputState }

func (SingletonPartitionSchemeDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SingletonPartitionSchemeDescription)(nil)).Elem()
}

func (o SingletonPartitionSchemeDescriptionOutput) ToSingletonPartitionSchemeDescriptionOutput() SingletonPartitionSchemeDescriptionOutput {
	return o
}

func (o SingletonPartitionSchemeDescriptionOutput) ToSingletonPartitionSchemeDescriptionOutputWithContext(ctx context.Context) SingletonPartitionSchemeDescriptionOutput {
	return o
}

func (o SingletonPartitionSchemeDescriptionOutput) PartitionScheme() pulumi.StringOutput {
	return o.ApplyT(func(v SingletonPartitionSchemeDescription) string { return v.PartitionScheme }).(pulumi.StringOutput)
}

type SingletonPartitionSchemeDescriptionResponse struct {
	PartitionScheme string `pulumi:"partitionScheme"`
}





type SingletonPartitionSchemeDescriptionResponseInput interface {
	pulumi.Input

	ToSingletonPartitionSchemeDescriptionResponseOutput() SingletonPartitionSchemeDescriptionResponseOutput
	ToSingletonPartitionSchemeDescriptionResponseOutputWithContext(context.Context) SingletonPartitionSchemeDescriptionResponseOutput
}

type SingletonPartitionSchemeDescriptionResponseArgs struct {
	PartitionScheme pulumi.StringInput `pulumi:"partitionScheme"`
}

func (SingletonPartitionSchemeDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SingletonPartitionSchemeDescriptionResponse)(nil)).Elem()
}

func (i SingletonPartitionSchemeDescriptionResponseArgs) ToSingletonPartitionSchemeDescriptionResponseOutput() SingletonPartitionSchemeDescriptionResponseOutput {
	return i.ToSingletonPartitionSchemeDescriptionResponseOutputWithContext(context.Background())
}

func (i SingletonPartitionSchemeDescriptionResponseArgs) ToSingletonPartitionSchemeDescriptionResponseOutputWithContext(ctx context.Context) SingletonPartitionSchemeDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SingletonPartitionSchemeDescriptionResponseOutput)
}

type SingletonPartitionSchemeDescriptionResponseOutput struct{ *pulumi.OutputState }

func (SingletonPartitionSchemeDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SingletonPartitionSchemeDescriptionResponse)(nil)).Elem()
}

func (o SingletonPartitionSchemeDescriptionResponseOutput) ToSingletonPartitionSchemeDescriptionResponseOutput() SingletonPartitionSchemeDescriptionResponseOutput {
	return o
}

func (o SingletonPartitionSchemeDescriptionResponseOutput) ToSingletonPartitionSchemeDescriptionResponseOutputWithContext(ctx context.Context) SingletonPartitionSchemeDescriptionResponseOutput {
	return o
}

func (o SingletonPartitionSchemeDescriptionResponseOutput) PartitionScheme() pulumi.StringOutput {
	return o.ApplyT(func(v SingletonPartitionSchemeDescriptionResponse) string { return v.PartitionScheme }).(pulumi.StringOutput)
}

type Sku struct {
	Name string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
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

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
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

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
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

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
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

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SubResource struct {
	Id *string `pulumi:"id"`
}





type SubResourceInput interface {
	pulumi.Input

	ToSubResourceOutput() SubResourceOutput
	ToSubResourceOutputWithContext(context.Context) SubResourceOutput
}

type SubResourceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (i SubResourceArgs) ToSubResourceOutput() SubResourceOutput {
	return i.ToSubResourceOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput)
}

type SubResourceOutput struct{ *pulumi.OutputState }

func (SubResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (o SubResourceOutput) ToSubResourceOutput() SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return o
}

func (o SubResourceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResource) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourceResponse struct {
	Id *string `pulumi:"id"`
}





type SubResourceResponseInput interface {
	pulumi.Input

	ToSubResourceResponseOutput() SubResourceResponseOutput
	ToSubResourceResponseOutputWithContext(context.Context) SubResourceResponseOutput
}

type SubResourceResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (i SubResourceResponseArgs) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return i.ToSubResourceResponseOutputWithContext(context.Background())
}

func (i SubResourceResponseArgs) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceResponseOutput)
}

type SubResourceResponseOutput struct{ *pulumi.OutputState }

func (SubResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResourceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type UniformInt64RangePartitionSchemeDescription struct {
	Count           int    `pulumi:"count"`
	HighKey         string `pulumi:"highKey"`
	LowKey          string `pulumi:"lowKey"`
	PartitionScheme string `pulumi:"partitionScheme"`
}





type UniformInt64RangePartitionSchemeDescriptionInput interface {
	pulumi.Input

	ToUniformInt64RangePartitionSchemeDescriptionOutput() UniformInt64RangePartitionSchemeDescriptionOutput
	ToUniformInt64RangePartitionSchemeDescriptionOutputWithContext(context.Context) UniformInt64RangePartitionSchemeDescriptionOutput
}

type UniformInt64RangePartitionSchemeDescriptionArgs struct {
	Count           pulumi.IntInput    `pulumi:"count"`
	HighKey         pulumi.StringInput `pulumi:"highKey"`
	LowKey          pulumi.StringInput `pulumi:"lowKey"`
	PartitionScheme pulumi.StringInput `pulumi:"partitionScheme"`
}

func (UniformInt64RangePartitionSchemeDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UniformInt64RangePartitionSchemeDescription)(nil)).Elem()
}

func (i UniformInt64RangePartitionSchemeDescriptionArgs) ToUniformInt64RangePartitionSchemeDescriptionOutput() UniformInt64RangePartitionSchemeDescriptionOutput {
	return i.ToUniformInt64RangePartitionSchemeDescriptionOutputWithContext(context.Background())
}

func (i UniformInt64RangePartitionSchemeDescriptionArgs) ToUniformInt64RangePartitionSchemeDescriptionOutputWithContext(ctx context.Context) UniformInt64RangePartitionSchemeDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniformInt64RangePartitionSchemeDescriptionOutput)
}

type UniformInt64RangePartitionSchemeDescriptionOutput struct{ *pulumi.OutputState }

func (UniformInt64RangePartitionSchemeDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UniformInt64RangePartitionSchemeDescription)(nil)).Elem()
}

func (o UniformInt64RangePartitionSchemeDescriptionOutput) ToUniformInt64RangePartitionSchemeDescriptionOutput() UniformInt64RangePartitionSchemeDescriptionOutput {
	return o
}

func (o UniformInt64RangePartitionSchemeDescriptionOutput) ToUniformInt64RangePartitionSchemeDescriptionOutputWithContext(ctx context.Context) UniformInt64RangePartitionSchemeDescriptionOutput {
	return o
}

func (o UniformInt64RangePartitionSchemeDescriptionOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v UniformInt64RangePartitionSchemeDescription) int { return v.Count }).(pulumi.IntOutput)
}

func (o UniformInt64RangePartitionSchemeDescriptionOutput) HighKey() pulumi.StringOutput {
	return o.ApplyT(func(v UniformInt64RangePartitionSchemeDescription) string { return v.HighKey }).(pulumi.StringOutput)
}

func (o UniformInt64RangePartitionSchemeDescriptionOutput) LowKey() pulumi.StringOutput {
	return o.ApplyT(func(v UniformInt64RangePartitionSchemeDescription) string { return v.LowKey }).(pulumi.StringOutput)
}

func (o UniformInt64RangePartitionSchemeDescriptionOutput) PartitionScheme() pulumi.StringOutput {
	return o.ApplyT(func(v UniformInt64RangePartitionSchemeDescription) string { return v.PartitionScheme }).(pulumi.StringOutput)
}

type UniformInt64RangePartitionSchemeDescriptionResponse struct {
	Count           int    `pulumi:"count"`
	HighKey         string `pulumi:"highKey"`
	LowKey          string `pulumi:"lowKey"`
	PartitionScheme string `pulumi:"partitionScheme"`
}





type UniformInt64RangePartitionSchemeDescriptionResponseInput interface {
	pulumi.Input

	ToUniformInt64RangePartitionSchemeDescriptionResponseOutput() UniformInt64RangePartitionSchemeDescriptionResponseOutput
	ToUniformInt64RangePartitionSchemeDescriptionResponseOutputWithContext(context.Context) UniformInt64RangePartitionSchemeDescriptionResponseOutput
}

type UniformInt64RangePartitionSchemeDescriptionResponseArgs struct {
	Count           pulumi.IntInput    `pulumi:"count"`
	HighKey         pulumi.StringInput `pulumi:"highKey"`
	LowKey          pulumi.StringInput `pulumi:"lowKey"`
	PartitionScheme pulumi.StringInput `pulumi:"partitionScheme"`
}

func (UniformInt64RangePartitionSchemeDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UniformInt64RangePartitionSchemeDescriptionResponse)(nil)).Elem()
}

func (i UniformInt64RangePartitionSchemeDescriptionResponseArgs) ToUniformInt64RangePartitionSchemeDescriptionResponseOutput() UniformInt64RangePartitionSchemeDescriptionResponseOutput {
	return i.ToUniformInt64RangePartitionSchemeDescriptionResponseOutputWithContext(context.Background())
}

func (i UniformInt64RangePartitionSchemeDescriptionResponseArgs) ToUniformInt64RangePartitionSchemeDescriptionResponseOutputWithContext(ctx context.Context) UniformInt64RangePartitionSchemeDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniformInt64RangePartitionSchemeDescriptionResponseOutput)
}

type UniformInt64RangePartitionSchemeDescriptionResponseOutput struct{ *pulumi.OutputState }

func (UniformInt64RangePartitionSchemeDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UniformInt64RangePartitionSchemeDescriptionResponse)(nil)).Elem()
}

func (o UniformInt64RangePartitionSchemeDescriptionResponseOutput) ToUniformInt64RangePartitionSchemeDescriptionResponseOutput() UniformInt64RangePartitionSchemeDescriptionResponseOutput {
	return o
}

func (o UniformInt64RangePartitionSchemeDescriptionResponseOutput) ToUniformInt64RangePartitionSchemeDescriptionResponseOutputWithContext(ctx context.Context) UniformInt64RangePartitionSchemeDescriptionResponseOutput {
	return o
}

func (o UniformInt64RangePartitionSchemeDescriptionResponseOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v UniformInt64RangePartitionSchemeDescriptionResponse) int { return v.Count }).(pulumi.IntOutput)
}

func (o UniformInt64RangePartitionSchemeDescriptionResponseOutput) HighKey() pulumi.StringOutput {
	return o.ApplyT(func(v UniformInt64RangePartitionSchemeDescriptionResponse) string { return v.HighKey }).(pulumi.StringOutput)
}

func (o UniformInt64RangePartitionSchemeDescriptionResponseOutput) LowKey() pulumi.StringOutput {
	return o.ApplyT(func(v UniformInt64RangePartitionSchemeDescriptionResponse) string { return v.LowKey }).(pulumi.StringOutput)
}

func (o UniformInt64RangePartitionSchemeDescriptionResponseOutput) PartitionScheme() pulumi.StringOutput {
	return o.ApplyT(func(v UniformInt64RangePartitionSchemeDescriptionResponse) string { return v.PartitionScheme }).(pulumi.StringOutput)
}

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}





type UserAssignedIdentityResponseInput interface {
	pulumi.Input

	ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput
	ToUserAssignedIdentityResponseOutputWithContext(context.Context) UserAssignedIdentityResponseOutput
}

type UserAssignedIdentityResponseArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (UserAssignedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (i UserAssignedIdentityResponseArgs) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return i.ToUserAssignedIdentityResponseOutputWithContext(context.Background())
}

func (i UserAssignedIdentityResponseArgs) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityResponseOutput)
}





type UserAssignedIdentityResponseMapInput interface {
	pulumi.Input

	ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput
	ToUserAssignedIdentityResponseMapOutputWithContext(context.Context) UserAssignedIdentityResponseMapOutput
}

type UserAssignedIdentityResponseMap map[string]UserAssignedIdentityResponseInput

func (UserAssignedIdentityResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (i UserAssignedIdentityResponseMap) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return i.ToUserAssignedIdentityResponseMapOutputWithContext(context.Background())
}

func (i UserAssignedIdentityResponseMap) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityResponseMapOutput)
}

type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

type VMSSExtension struct {
	AutoUpgradeMinorVersion  *bool       `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag           *string     `pulumi:"forceUpdateTag"`
	Name                     string      `pulumi:"name"`
	ProtectedSettings        interface{} `pulumi:"protectedSettings"`
	ProvisionAfterExtensions []string    `pulumi:"provisionAfterExtensions"`
	Publisher                string      `pulumi:"publisher"`
	Settings                 interface{} `pulumi:"settings"`
	Type                     string      `pulumi:"type"`
	TypeHandlerVersion       string      `pulumi:"typeHandlerVersion"`
}





type VMSSExtensionInput interface {
	pulumi.Input

	ToVMSSExtensionOutput() VMSSExtensionOutput
	ToVMSSExtensionOutputWithContext(context.Context) VMSSExtensionOutput
}

type VMSSExtensionArgs struct {
	AutoUpgradeMinorVersion  pulumi.BoolPtrInput     `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag           pulumi.StringPtrInput   `pulumi:"forceUpdateTag"`
	Name                     pulumi.StringInput      `pulumi:"name"`
	ProtectedSettings        pulumi.Input            `pulumi:"protectedSettings"`
	ProvisionAfterExtensions pulumi.StringArrayInput `pulumi:"provisionAfterExtensions"`
	Publisher                pulumi.StringInput      `pulumi:"publisher"`
	Settings                 pulumi.Input            `pulumi:"settings"`
	Type                     pulumi.StringInput      `pulumi:"type"`
	TypeHandlerVersion       pulumi.StringInput      `pulumi:"typeHandlerVersion"`
}

func (VMSSExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSSExtension)(nil)).Elem()
}

func (i VMSSExtensionArgs) ToVMSSExtensionOutput() VMSSExtensionOutput {
	return i.ToVMSSExtensionOutputWithContext(context.Background())
}

func (i VMSSExtensionArgs) ToVMSSExtensionOutputWithContext(ctx context.Context) VMSSExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSSExtensionOutput)
}





type VMSSExtensionArrayInput interface {
	pulumi.Input

	ToVMSSExtensionArrayOutput() VMSSExtensionArrayOutput
	ToVMSSExtensionArrayOutputWithContext(context.Context) VMSSExtensionArrayOutput
}

type VMSSExtensionArray []VMSSExtensionInput

func (VMSSExtensionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMSSExtension)(nil)).Elem()
}

func (i VMSSExtensionArray) ToVMSSExtensionArrayOutput() VMSSExtensionArrayOutput {
	return i.ToVMSSExtensionArrayOutputWithContext(context.Background())
}

func (i VMSSExtensionArray) ToVMSSExtensionArrayOutputWithContext(ctx context.Context) VMSSExtensionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSSExtensionArrayOutput)
}

type VMSSExtensionOutput struct{ *pulumi.OutputState }

func (VMSSExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSSExtension)(nil)).Elem()
}

func (o VMSSExtensionOutput) ToVMSSExtensionOutput() VMSSExtensionOutput {
	return o
}

func (o VMSSExtensionOutput) ToVMSSExtensionOutputWithContext(ctx context.Context) VMSSExtensionOutput {
	return o
}

func (o VMSSExtensionOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VMSSExtension) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o VMSSExtensionOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMSSExtension) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o VMSSExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtension) string { return v.Name }).(pulumi.StringOutput)
}

func (o VMSSExtensionOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v VMSSExtension) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o VMSSExtensionOutput) ProvisionAfterExtensions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VMSSExtension) []string { return v.ProvisionAfterExtensions }).(pulumi.StringArrayOutput)
}

func (o VMSSExtensionOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtension) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o VMSSExtensionOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v VMSSExtension) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o VMSSExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtension) string { return v.Type }).(pulumi.StringOutput)
}

func (o VMSSExtensionOutput) TypeHandlerVersion() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtension) string { return v.TypeHandlerVersion }).(pulumi.StringOutput)
}

type VMSSExtensionArrayOutput struct{ *pulumi.OutputState }

func (VMSSExtensionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMSSExtension)(nil)).Elem()
}

func (o VMSSExtensionArrayOutput) ToVMSSExtensionArrayOutput() VMSSExtensionArrayOutput {
	return o
}

func (o VMSSExtensionArrayOutput) ToVMSSExtensionArrayOutputWithContext(ctx context.Context) VMSSExtensionArrayOutput {
	return o
}

func (o VMSSExtensionArrayOutput) Index(i pulumi.IntInput) VMSSExtensionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VMSSExtension {
		return vs[0].([]VMSSExtension)[vs[1].(int)]
	}).(VMSSExtensionOutput)
}

type VMSSExtensionResponse struct {
	AutoUpgradeMinorVersion  *bool       `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag           *string     `pulumi:"forceUpdateTag"`
	Name                     string      `pulumi:"name"`
	ProtectedSettings        interface{} `pulumi:"protectedSettings"`
	ProvisionAfterExtensions []string    `pulumi:"provisionAfterExtensions"`
	ProvisioningState        string      `pulumi:"provisioningState"`
	Publisher                string      `pulumi:"publisher"`
	Settings                 interface{} `pulumi:"settings"`
	Type                     string      `pulumi:"type"`
	TypeHandlerVersion       string      `pulumi:"typeHandlerVersion"`
}





type VMSSExtensionResponseInput interface {
	pulumi.Input

	ToVMSSExtensionResponseOutput() VMSSExtensionResponseOutput
	ToVMSSExtensionResponseOutputWithContext(context.Context) VMSSExtensionResponseOutput
}

type VMSSExtensionResponseArgs struct {
	AutoUpgradeMinorVersion  pulumi.BoolPtrInput     `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag           pulumi.StringPtrInput   `pulumi:"forceUpdateTag"`
	Name                     pulumi.StringInput      `pulumi:"name"`
	ProtectedSettings        pulumi.Input            `pulumi:"protectedSettings"`
	ProvisionAfterExtensions pulumi.StringArrayInput `pulumi:"provisionAfterExtensions"`
	ProvisioningState        pulumi.StringInput      `pulumi:"provisioningState"`
	Publisher                pulumi.StringInput      `pulumi:"publisher"`
	Settings                 pulumi.Input            `pulumi:"settings"`
	Type                     pulumi.StringInput      `pulumi:"type"`
	TypeHandlerVersion       pulumi.StringInput      `pulumi:"typeHandlerVersion"`
}

func (VMSSExtensionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSSExtensionResponse)(nil)).Elem()
}

func (i VMSSExtensionResponseArgs) ToVMSSExtensionResponseOutput() VMSSExtensionResponseOutput {
	return i.ToVMSSExtensionResponseOutputWithContext(context.Background())
}

func (i VMSSExtensionResponseArgs) ToVMSSExtensionResponseOutputWithContext(ctx context.Context) VMSSExtensionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSSExtensionResponseOutput)
}





type VMSSExtensionResponseArrayInput interface {
	pulumi.Input

	ToVMSSExtensionResponseArrayOutput() VMSSExtensionResponseArrayOutput
	ToVMSSExtensionResponseArrayOutputWithContext(context.Context) VMSSExtensionResponseArrayOutput
}

type VMSSExtensionResponseArray []VMSSExtensionResponseInput

func (VMSSExtensionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMSSExtensionResponse)(nil)).Elem()
}

func (i VMSSExtensionResponseArray) ToVMSSExtensionResponseArrayOutput() VMSSExtensionResponseArrayOutput {
	return i.ToVMSSExtensionResponseArrayOutputWithContext(context.Background())
}

func (i VMSSExtensionResponseArray) ToVMSSExtensionResponseArrayOutputWithContext(ctx context.Context) VMSSExtensionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSSExtensionResponseArrayOutput)
}

type VMSSExtensionResponseOutput struct{ *pulumi.OutputState }

func (VMSSExtensionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSSExtensionResponse)(nil)).Elem()
}

func (o VMSSExtensionResponseOutput) ToVMSSExtensionResponseOutput() VMSSExtensionResponseOutput {
	return o
}

func (o VMSSExtensionResponseOutput) ToVMSSExtensionResponseOutputWithContext(ctx context.Context) VMSSExtensionResponseOutput {
	return o
}

func (o VMSSExtensionResponseOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o VMSSExtensionResponseOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o VMSSExtensionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VMSSExtensionResponseOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o VMSSExtensionResponseOutput) ProvisionAfterExtensions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) []string { return v.ProvisionAfterExtensions }).(pulumi.StringArrayOutput)
}

func (o VMSSExtensionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VMSSExtensionResponseOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o VMSSExtensionResponseOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o VMSSExtensionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VMSSExtensionResponseOutput) TypeHandlerVersion() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) string { return v.TypeHandlerVersion }).(pulumi.StringOutput)
}

type VMSSExtensionResponseArrayOutput struct{ *pulumi.OutputState }

func (VMSSExtensionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMSSExtensionResponse)(nil)).Elem()
}

func (o VMSSExtensionResponseArrayOutput) ToVMSSExtensionResponseArrayOutput() VMSSExtensionResponseArrayOutput {
	return o
}

func (o VMSSExtensionResponseArrayOutput) ToVMSSExtensionResponseArrayOutputWithContext(ctx context.Context) VMSSExtensionResponseArrayOutput {
	return o
}

func (o VMSSExtensionResponseArrayOutput) Index(i pulumi.IntInput) VMSSExtensionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VMSSExtensionResponse {
		return vs[0].([]VMSSExtensionResponse)[vs[1].(int)]
	}).(VMSSExtensionResponseOutput)
}

type VaultCertificate struct {
	CertificateStore string `pulumi:"certificateStore"`
	CertificateUrl   string `pulumi:"certificateUrl"`
}





type VaultCertificateInput interface {
	pulumi.Input

	ToVaultCertificateOutput() VaultCertificateOutput
	ToVaultCertificateOutputWithContext(context.Context) VaultCertificateOutput
}

type VaultCertificateArgs struct {
	CertificateStore pulumi.StringInput `pulumi:"certificateStore"`
	CertificateUrl   pulumi.StringInput `pulumi:"certificateUrl"`
}

func (VaultCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultCertificate)(nil)).Elem()
}

func (i VaultCertificateArgs) ToVaultCertificateOutput() VaultCertificateOutput {
	return i.ToVaultCertificateOutputWithContext(context.Background())
}

func (i VaultCertificateArgs) ToVaultCertificateOutputWithContext(ctx context.Context) VaultCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultCertificateOutput)
}





type VaultCertificateArrayInput interface {
	pulumi.Input

	ToVaultCertificateArrayOutput() VaultCertificateArrayOutput
	ToVaultCertificateArrayOutputWithContext(context.Context) VaultCertificateArrayOutput
}

type VaultCertificateArray []VaultCertificateInput

func (VaultCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultCertificate)(nil)).Elem()
}

func (i VaultCertificateArray) ToVaultCertificateArrayOutput() VaultCertificateArrayOutput {
	return i.ToVaultCertificateArrayOutputWithContext(context.Background())
}

func (i VaultCertificateArray) ToVaultCertificateArrayOutputWithContext(ctx context.Context) VaultCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultCertificateArrayOutput)
}

type VaultCertificateOutput struct{ *pulumi.OutputState }

func (VaultCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultCertificate)(nil)).Elem()
}

func (o VaultCertificateOutput) ToVaultCertificateOutput() VaultCertificateOutput {
	return o
}

func (o VaultCertificateOutput) ToVaultCertificateOutputWithContext(ctx context.Context) VaultCertificateOutput {
	return o
}

func (o VaultCertificateOutput) CertificateStore() pulumi.StringOutput {
	return o.ApplyT(func(v VaultCertificate) string { return v.CertificateStore }).(pulumi.StringOutput)
}

func (o VaultCertificateOutput) CertificateUrl() pulumi.StringOutput {
	return o.ApplyT(func(v VaultCertificate) string { return v.CertificateUrl }).(pulumi.StringOutput)
}

type VaultCertificateArrayOutput struct{ *pulumi.OutputState }

func (VaultCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultCertificate)(nil)).Elem()
}

func (o VaultCertificateArrayOutput) ToVaultCertificateArrayOutput() VaultCertificateArrayOutput {
	return o
}

func (o VaultCertificateArrayOutput) ToVaultCertificateArrayOutputWithContext(ctx context.Context) VaultCertificateArrayOutput {
	return o
}

func (o VaultCertificateArrayOutput) Index(i pulumi.IntInput) VaultCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultCertificate {
		return vs[0].([]VaultCertificate)[vs[1].(int)]
	}).(VaultCertificateOutput)
}

type VaultCertificateResponse struct {
	CertificateStore string `pulumi:"certificateStore"`
	CertificateUrl   string `pulumi:"certificateUrl"`
}





type VaultCertificateResponseInput interface {
	pulumi.Input

	ToVaultCertificateResponseOutput() VaultCertificateResponseOutput
	ToVaultCertificateResponseOutputWithContext(context.Context) VaultCertificateResponseOutput
}

type VaultCertificateResponseArgs struct {
	CertificateStore pulumi.StringInput `pulumi:"certificateStore"`
	CertificateUrl   pulumi.StringInput `pulumi:"certificateUrl"`
}

func (VaultCertificateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultCertificateResponse)(nil)).Elem()
}

func (i VaultCertificateResponseArgs) ToVaultCertificateResponseOutput() VaultCertificateResponseOutput {
	return i.ToVaultCertificateResponseOutputWithContext(context.Background())
}

func (i VaultCertificateResponseArgs) ToVaultCertificateResponseOutputWithContext(ctx context.Context) VaultCertificateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultCertificateResponseOutput)
}





type VaultCertificateResponseArrayInput interface {
	pulumi.Input

	ToVaultCertificateResponseArrayOutput() VaultCertificateResponseArrayOutput
	ToVaultCertificateResponseArrayOutputWithContext(context.Context) VaultCertificateResponseArrayOutput
}

type VaultCertificateResponseArray []VaultCertificateResponseInput

func (VaultCertificateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultCertificateResponse)(nil)).Elem()
}

func (i VaultCertificateResponseArray) ToVaultCertificateResponseArrayOutput() VaultCertificateResponseArrayOutput {
	return i.ToVaultCertificateResponseArrayOutputWithContext(context.Background())
}

func (i VaultCertificateResponseArray) ToVaultCertificateResponseArrayOutputWithContext(ctx context.Context) VaultCertificateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultCertificateResponseArrayOutput)
}

type VaultCertificateResponseOutput struct{ *pulumi.OutputState }

func (VaultCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultCertificateResponse)(nil)).Elem()
}

func (o VaultCertificateResponseOutput) ToVaultCertificateResponseOutput() VaultCertificateResponseOutput {
	return o
}

func (o VaultCertificateResponseOutput) ToVaultCertificateResponseOutputWithContext(ctx context.Context) VaultCertificateResponseOutput {
	return o
}

func (o VaultCertificateResponseOutput) CertificateStore() pulumi.StringOutput {
	return o.ApplyT(func(v VaultCertificateResponse) string { return v.CertificateStore }).(pulumi.StringOutput)
}

func (o VaultCertificateResponseOutput) CertificateUrl() pulumi.StringOutput {
	return o.ApplyT(func(v VaultCertificateResponse) string { return v.CertificateUrl }).(pulumi.StringOutput)
}

type VaultCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VaultCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultCertificateResponse)(nil)).Elem()
}

func (o VaultCertificateResponseArrayOutput) ToVaultCertificateResponseArrayOutput() VaultCertificateResponseArrayOutput {
	return o
}

func (o VaultCertificateResponseArrayOutput) ToVaultCertificateResponseArrayOutputWithContext(ctx context.Context) VaultCertificateResponseArrayOutput {
	return o
}

func (o VaultCertificateResponseArrayOutput) Index(i pulumi.IntInput) VaultCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultCertificateResponse {
		return vs[0].([]VaultCertificateResponse)[vs[1].(int)]
	}).(VaultCertificateResponseOutput)
}

type VaultSecretGroup struct {
	SourceVault       SubResource        `pulumi:"sourceVault"`
	VaultCertificates []VaultCertificate `pulumi:"vaultCertificates"`
}





type VaultSecretGroupInput interface {
	pulumi.Input

	ToVaultSecretGroupOutput() VaultSecretGroupOutput
	ToVaultSecretGroupOutputWithContext(context.Context) VaultSecretGroupOutput
}

type VaultSecretGroupArgs struct {
	SourceVault       SubResourceInput           `pulumi:"sourceVault"`
	VaultCertificates VaultCertificateArrayInput `pulumi:"vaultCertificates"`
}

func (VaultSecretGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultSecretGroup)(nil)).Elem()
}

func (i VaultSecretGroupArgs) ToVaultSecretGroupOutput() VaultSecretGroupOutput {
	return i.ToVaultSecretGroupOutputWithContext(context.Background())
}

func (i VaultSecretGroupArgs) ToVaultSecretGroupOutputWithContext(ctx context.Context) VaultSecretGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultSecretGroupOutput)
}





type VaultSecretGroupArrayInput interface {
	pulumi.Input

	ToVaultSecretGroupArrayOutput() VaultSecretGroupArrayOutput
	ToVaultSecretGroupArrayOutputWithContext(context.Context) VaultSecretGroupArrayOutput
}

type VaultSecretGroupArray []VaultSecretGroupInput

func (VaultSecretGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultSecretGroup)(nil)).Elem()
}

func (i VaultSecretGroupArray) ToVaultSecretGroupArrayOutput() VaultSecretGroupArrayOutput {
	return i.ToVaultSecretGroupArrayOutputWithContext(context.Background())
}

func (i VaultSecretGroupArray) ToVaultSecretGroupArrayOutputWithContext(ctx context.Context) VaultSecretGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultSecretGroupArrayOutput)
}

type VaultSecretGroupOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultSecretGroup)(nil)).Elem()
}

func (o VaultSecretGroupOutput) ToVaultSecretGroupOutput() VaultSecretGroupOutput {
	return o
}

func (o VaultSecretGroupOutput) ToVaultSecretGroupOutputWithContext(ctx context.Context) VaultSecretGroupOutput {
	return o
}

func (o VaultSecretGroupOutput) SourceVault() SubResourceOutput {
	return o.ApplyT(func(v VaultSecretGroup) SubResource { return v.SourceVault }).(SubResourceOutput)
}

func (o VaultSecretGroupOutput) VaultCertificates() VaultCertificateArrayOutput {
	return o.ApplyT(func(v VaultSecretGroup) []VaultCertificate { return v.VaultCertificates }).(VaultCertificateArrayOutput)
}

type VaultSecretGroupArrayOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultSecretGroup)(nil)).Elem()
}

func (o VaultSecretGroupArrayOutput) ToVaultSecretGroupArrayOutput() VaultSecretGroupArrayOutput {
	return o
}

func (o VaultSecretGroupArrayOutput) ToVaultSecretGroupArrayOutputWithContext(ctx context.Context) VaultSecretGroupArrayOutput {
	return o
}

func (o VaultSecretGroupArrayOutput) Index(i pulumi.IntInput) VaultSecretGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultSecretGroup {
		return vs[0].([]VaultSecretGroup)[vs[1].(int)]
	}).(VaultSecretGroupOutput)
}

type VaultSecretGroupResponse struct {
	SourceVault       SubResourceResponse        `pulumi:"sourceVault"`
	VaultCertificates []VaultCertificateResponse `pulumi:"vaultCertificates"`
}





type VaultSecretGroupResponseInput interface {
	pulumi.Input

	ToVaultSecretGroupResponseOutput() VaultSecretGroupResponseOutput
	ToVaultSecretGroupResponseOutputWithContext(context.Context) VaultSecretGroupResponseOutput
}

type VaultSecretGroupResponseArgs struct {
	SourceVault       SubResourceResponseInput           `pulumi:"sourceVault"`
	VaultCertificates VaultCertificateResponseArrayInput `pulumi:"vaultCertificates"`
}

func (VaultSecretGroupResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultSecretGroupResponse)(nil)).Elem()
}

func (i VaultSecretGroupResponseArgs) ToVaultSecretGroupResponseOutput() VaultSecretGroupResponseOutput {
	return i.ToVaultSecretGroupResponseOutputWithContext(context.Background())
}

func (i VaultSecretGroupResponseArgs) ToVaultSecretGroupResponseOutputWithContext(ctx context.Context) VaultSecretGroupResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultSecretGroupResponseOutput)
}





type VaultSecretGroupResponseArrayInput interface {
	pulumi.Input

	ToVaultSecretGroupResponseArrayOutput() VaultSecretGroupResponseArrayOutput
	ToVaultSecretGroupResponseArrayOutputWithContext(context.Context) VaultSecretGroupResponseArrayOutput
}

type VaultSecretGroupResponseArray []VaultSecretGroupResponseInput

func (VaultSecretGroupResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultSecretGroupResponse)(nil)).Elem()
}

func (i VaultSecretGroupResponseArray) ToVaultSecretGroupResponseArrayOutput() VaultSecretGroupResponseArrayOutput {
	return i.ToVaultSecretGroupResponseArrayOutputWithContext(context.Background())
}

func (i VaultSecretGroupResponseArray) ToVaultSecretGroupResponseArrayOutputWithContext(ctx context.Context) VaultSecretGroupResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultSecretGroupResponseArrayOutput)
}

type VaultSecretGroupResponseOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultSecretGroupResponse)(nil)).Elem()
}

func (o VaultSecretGroupResponseOutput) ToVaultSecretGroupResponseOutput() VaultSecretGroupResponseOutput {
	return o
}

func (o VaultSecretGroupResponseOutput) ToVaultSecretGroupResponseOutputWithContext(ctx context.Context) VaultSecretGroupResponseOutput {
	return o
}

func (o VaultSecretGroupResponseOutput) SourceVault() SubResourceResponseOutput {
	return o.ApplyT(func(v VaultSecretGroupResponse) SubResourceResponse { return v.SourceVault }).(SubResourceResponseOutput)
}

func (o VaultSecretGroupResponseOutput) VaultCertificates() VaultCertificateResponseArrayOutput {
	return o.ApplyT(func(v VaultSecretGroupResponse) []VaultCertificateResponse { return v.VaultCertificates }).(VaultCertificateResponseArrayOutput)
}

type VaultSecretGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultSecretGroupResponse)(nil)).Elem()
}

func (o VaultSecretGroupResponseArrayOutput) ToVaultSecretGroupResponseArrayOutput() VaultSecretGroupResponseArrayOutput {
	return o
}

func (o VaultSecretGroupResponseArrayOutput) ToVaultSecretGroupResponseArrayOutputWithContext(ctx context.Context) VaultSecretGroupResponseArrayOutput {
	return o
}

func (o VaultSecretGroupResponseArrayOutput) Index(i pulumi.IntInput) VaultSecretGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultSecretGroupResponse {
		return vs[0].([]VaultSecretGroupResponse)[vs[1].(int)]
	}).(VaultSecretGroupResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationDeltaHealthPolicyOutput{})
	pulumi.RegisterOutputType(ApplicationDeltaHealthPolicyMapOutput{})
	pulumi.RegisterOutputType(ApplicationDeltaHealthPolicyResponseOutput{})
	pulumi.RegisterOutputType(ApplicationDeltaHealthPolicyResponseMapOutput{})
	pulumi.RegisterOutputType(ApplicationHealthPolicyOutput{})
	pulumi.RegisterOutputType(ApplicationHealthPolicyMapOutput{})
	pulumi.RegisterOutputType(ApplicationHealthPolicyResponseOutput{})
	pulumi.RegisterOutputType(ApplicationHealthPolicyResponseMapOutput{})
	pulumi.RegisterOutputType(ApplicationMetricDescriptionOutput{})
	pulumi.RegisterOutputType(ApplicationMetricDescriptionArrayOutput{})
	pulumi.RegisterOutputType(ApplicationMetricDescriptionResponseOutput{})
	pulumi.RegisterOutputType(ApplicationMetricDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationTypeVersionsCleanupPolicyOutput{})
	pulumi.RegisterOutputType(ApplicationTypeVersionsCleanupPolicyPtrOutput{})
	pulumi.RegisterOutputType(ApplicationTypeVersionsCleanupPolicyResponseOutput{})
	pulumi.RegisterOutputType(ApplicationTypeVersionsCleanupPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationUpgradePolicyOutput{})
	pulumi.RegisterOutputType(ApplicationUpgradePolicyPtrOutput{})
	pulumi.RegisterOutputType(ApplicationUpgradePolicyResponseOutput{})
	pulumi.RegisterOutputType(ApplicationUpgradePolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationUserAssignedIdentityOutput{})
	pulumi.RegisterOutputType(ApplicationUserAssignedIdentityArrayOutput{})
	pulumi.RegisterOutputType(ApplicationUserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(ApplicationUserAssignedIdentityResponseArrayOutput{})
	pulumi.RegisterOutputType(ArmApplicationHealthPolicyOutput{})
	pulumi.RegisterOutputType(ArmApplicationHealthPolicyPtrOutput{})
	pulumi.RegisterOutputType(ArmApplicationHealthPolicyResponseOutput{})
	pulumi.RegisterOutputType(ArmApplicationHealthPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ArmRollingUpgradeMonitoringPolicyOutput{})
	pulumi.RegisterOutputType(ArmRollingUpgradeMonitoringPolicyPtrOutput{})
	pulumi.RegisterOutputType(ArmRollingUpgradeMonitoringPolicyResponseOutput{})
	pulumi.RegisterOutputType(ArmRollingUpgradeMonitoringPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ArmServiceTypeHealthPolicyOutput{})
	pulumi.RegisterOutputType(ArmServiceTypeHealthPolicyPtrOutput{})
	pulumi.RegisterOutputType(ArmServiceTypeHealthPolicyMapOutput{})
	pulumi.RegisterOutputType(ArmServiceTypeHealthPolicyResponseOutput{})
	pulumi.RegisterOutputType(ArmServiceTypeHealthPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ArmServiceTypeHealthPolicyResponseMapOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryPtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryResponseOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryResponsePtrOutput{})
	pulumi.RegisterOutputType(CertificateDescriptionOutput{})
	pulumi.RegisterOutputType(CertificateDescriptionPtrOutput{})
	pulumi.RegisterOutputType(CertificateDescriptionResponseOutput{})
	pulumi.RegisterOutputType(CertificateDescriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(ClientCertificateOutput{})
	pulumi.RegisterOutputType(ClientCertificateArrayOutput{})
	pulumi.RegisterOutputType(ClientCertificateCommonNameOutput{})
	pulumi.RegisterOutputType(ClientCertificateCommonNameArrayOutput{})
	pulumi.RegisterOutputType(ClientCertificateCommonNameResponseOutput{})
	pulumi.RegisterOutputType(ClientCertificateCommonNameResponseArrayOutput{})
	pulumi.RegisterOutputType(ClientCertificateResponseOutput{})
	pulumi.RegisterOutputType(ClientCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(ClientCertificateThumbprintOutput{})
	pulumi.RegisterOutputType(ClientCertificateThumbprintArrayOutput{})
	pulumi.RegisterOutputType(ClientCertificateThumbprintResponseOutput{})
	pulumi.RegisterOutputType(ClientCertificateThumbprintResponseArrayOutput{})
	pulumi.RegisterOutputType(ClusterHealthPolicyOutput{})
	pulumi.RegisterOutputType(ClusterHealthPolicyPtrOutput{})
	pulumi.RegisterOutputType(ClusterHealthPolicyResponseOutput{})
	pulumi.RegisterOutputType(ClusterHealthPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ClusterUpgradeDeltaHealthPolicyOutput{})
	pulumi.RegisterOutputType(ClusterUpgradeDeltaHealthPolicyPtrOutput{})
	pulumi.RegisterOutputType(ClusterUpgradeDeltaHealthPolicyResponseOutput{})
	pulumi.RegisterOutputType(ClusterUpgradeDeltaHealthPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ClusterUpgradePolicyOutput{})
	pulumi.RegisterOutputType(ClusterUpgradePolicyPtrOutput{})
	pulumi.RegisterOutputType(ClusterUpgradePolicyResponseOutput{})
	pulumi.RegisterOutputType(ClusterUpgradePolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ClusterVersionDetailsResponseOutput{})
	pulumi.RegisterOutputType(ClusterVersionDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(DiagnosticsStorageAccountConfigOutput{})
	pulumi.RegisterOutputType(DiagnosticsStorageAccountConfigPtrOutput{})
	pulumi.RegisterOutputType(DiagnosticsStorageAccountConfigResponseOutput{})
	pulumi.RegisterOutputType(DiagnosticsStorageAccountConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(EndpointRangeDescriptionOutput{})
	pulumi.RegisterOutputType(EndpointRangeDescriptionPtrOutput{})
	pulumi.RegisterOutputType(EndpointRangeDescriptionResponseOutput{})
	pulumi.RegisterOutputType(EndpointRangeDescriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedIdentityOutput{})
	pulumi.RegisterOutputType(ManagedIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(NamedPartitionSchemeDescriptionOutput{})
	pulumi.RegisterOutputType(NamedPartitionSchemeDescriptionResponseOutput{})
	pulumi.RegisterOutputType(NodeTypeDescriptionOutput{})
	pulumi.RegisterOutputType(NodeTypeDescriptionArrayOutput{})
	pulumi.RegisterOutputType(NodeTypeDescriptionResponseOutput{})
	pulumi.RegisterOutputType(NodeTypeDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(ServerCertificateCommonNameOutput{})
	pulumi.RegisterOutputType(ServerCertificateCommonNameArrayOutput{})
	pulumi.RegisterOutputType(ServerCertificateCommonNameResponseOutput{})
	pulumi.RegisterOutputType(ServerCertificateCommonNameResponseArrayOutput{})
	pulumi.RegisterOutputType(ServerCertificateCommonNamesOutput{})
	pulumi.RegisterOutputType(ServerCertificateCommonNamesPtrOutput{})
	pulumi.RegisterOutputType(ServerCertificateCommonNamesResponseOutput{})
	pulumi.RegisterOutputType(ServerCertificateCommonNamesResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceCorrelationDescriptionOutput{})
	pulumi.RegisterOutputType(ServiceCorrelationDescriptionArrayOutput{})
	pulumi.RegisterOutputType(ServiceCorrelationDescriptionResponseOutput{})
	pulumi.RegisterOutputType(ServiceCorrelationDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceLoadMetricDescriptionOutput{})
	pulumi.RegisterOutputType(ServiceLoadMetricDescriptionArrayOutput{})
	pulumi.RegisterOutputType(ServiceLoadMetricDescriptionResponseOutput{})
	pulumi.RegisterOutputType(ServiceLoadMetricDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(ServicePlacementPolicyDescriptionOutput{})
	pulumi.RegisterOutputType(ServicePlacementPolicyDescriptionArrayOutput{})
	pulumi.RegisterOutputType(ServicePlacementPolicyDescriptionResponseOutput{})
	pulumi.RegisterOutputType(ServicePlacementPolicyDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceTypeDeltaHealthPolicyOutput{})
	pulumi.RegisterOutputType(ServiceTypeDeltaHealthPolicyPtrOutput{})
	pulumi.RegisterOutputType(ServiceTypeDeltaHealthPolicyMapOutput{})
	pulumi.RegisterOutputType(ServiceTypeDeltaHealthPolicyResponseOutput{})
	pulumi.RegisterOutputType(ServiceTypeDeltaHealthPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceTypeDeltaHealthPolicyResponseMapOutput{})
	pulumi.RegisterOutputType(ServiceTypeHealthPolicyOutput{})
	pulumi.RegisterOutputType(ServiceTypeHealthPolicyPtrOutput{})
	pulumi.RegisterOutputType(ServiceTypeHealthPolicyMapOutput{})
	pulumi.RegisterOutputType(ServiceTypeHealthPolicyResponseOutput{})
	pulumi.RegisterOutputType(ServiceTypeHealthPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceTypeHealthPolicyResponseMapOutput{})
	pulumi.RegisterOutputType(SettingsParameterDescriptionOutput{})
	pulumi.RegisterOutputType(SettingsParameterDescriptionArrayOutput{})
	pulumi.RegisterOutputType(SettingsParameterDescriptionResponseOutput{})
	pulumi.RegisterOutputType(SettingsParameterDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(SettingsSectionDescriptionOutput{})
	pulumi.RegisterOutputType(SettingsSectionDescriptionArrayOutput{})
	pulumi.RegisterOutputType(SettingsSectionDescriptionResponseOutput{})
	pulumi.RegisterOutputType(SettingsSectionDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(SingletonPartitionSchemeDescriptionOutput{})
	pulumi.RegisterOutputType(SingletonPartitionSchemeDescriptionResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(UniformInt64RangePartitionSchemeDescriptionOutput{})
	pulumi.RegisterOutputType(UniformInt64RangePartitionSchemeDescriptionResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(VMSSExtensionOutput{})
	pulumi.RegisterOutputType(VMSSExtensionArrayOutput{})
	pulumi.RegisterOutputType(VMSSExtensionResponseOutput{})
	pulumi.RegisterOutputType(VMSSExtensionResponseArrayOutput{})
	pulumi.RegisterOutputType(VaultCertificateOutput{})
	pulumi.RegisterOutputType(VaultCertificateArrayOutput{})
	pulumi.RegisterOutputType(VaultCertificateResponseOutput{})
	pulumi.RegisterOutputType(VaultCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupArrayOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupResponseOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupResponseArrayOutput{})
}
