


package v20220904

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type APIServerProfile struct {
	Ip         *string `pulumi:"ip"`
	Url        *string `pulumi:"url"`
	Visibility *string `pulumi:"visibility"`
}





type APIServerProfileInput interface {
	pulumi.Input

	ToAPIServerProfileOutput() APIServerProfileOutput
	ToAPIServerProfileOutputWithContext(context.Context) APIServerProfileOutput
}

type APIServerProfileArgs struct {
	Ip         pulumi.StringPtrInput `pulumi:"ip"`
	Url        pulumi.StringPtrInput `pulumi:"url"`
	Visibility pulumi.StringPtrInput `pulumi:"visibility"`
}

func (APIServerProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*APIServerProfile)(nil)).Elem()
}

func (i APIServerProfileArgs) ToAPIServerProfileOutput() APIServerProfileOutput {
	return i.ToAPIServerProfileOutputWithContext(context.Background())
}

func (i APIServerProfileArgs) ToAPIServerProfileOutputWithContext(ctx context.Context) APIServerProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIServerProfileOutput)
}

func (i APIServerProfileArgs) ToAPIServerProfilePtrOutput() APIServerProfilePtrOutput {
	return i.ToAPIServerProfilePtrOutputWithContext(context.Background())
}

func (i APIServerProfileArgs) ToAPIServerProfilePtrOutputWithContext(ctx context.Context) APIServerProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIServerProfileOutput).ToAPIServerProfilePtrOutputWithContext(ctx)
}









type APIServerProfilePtrInput interface {
	pulumi.Input

	ToAPIServerProfilePtrOutput() APIServerProfilePtrOutput
	ToAPIServerProfilePtrOutputWithContext(context.Context) APIServerProfilePtrOutput
}

type apiserverProfilePtrType APIServerProfileArgs

func APIServerProfilePtr(v *APIServerProfileArgs) APIServerProfilePtrInput {
	return (*apiserverProfilePtrType)(v)
}

func (*apiserverProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**APIServerProfile)(nil)).Elem()
}

func (i *apiserverProfilePtrType) ToAPIServerProfilePtrOutput() APIServerProfilePtrOutput {
	return i.ToAPIServerProfilePtrOutputWithContext(context.Background())
}

func (i *apiserverProfilePtrType) ToAPIServerProfilePtrOutputWithContext(ctx context.Context) APIServerProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIServerProfilePtrOutput)
}

type APIServerProfileOutput struct{ *pulumi.OutputState }

func (APIServerProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*APIServerProfile)(nil)).Elem()
}

func (o APIServerProfileOutput) ToAPIServerProfileOutput() APIServerProfileOutput {
	return o
}

func (o APIServerProfileOutput) ToAPIServerProfileOutputWithContext(ctx context.Context) APIServerProfileOutput {
	return o
}

func (o APIServerProfileOutput) ToAPIServerProfilePtrOutput() APIServerProfilePtrOutput {
	return o.ToAPIServerProfilePtrOutputWithContext(context.Background())
}

func (o APIServerProfileOutput) ToAPIServerProfilePtrOutputWithContext(ctx context.Context) APIServerProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v APIServerProfile) *APIServerProfile {
		return &v
	}).(APIServerProfilePtrOutput)
}

func (o APIServerProfileOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v APIServerProfile) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o APIServerProfileOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v APIServerProfile) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o APIServerProfileOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v APIServerProfile) *string { return v.Visibility }).(pulumi.StringPtrOutput)
}

type APIServerProfilePtrOutput struct{ *pulumi.OutputState }

func (APIServerProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**APIServerProfile)(nil)).Elem()
}

func (o APIServerProfilePtrOutput) ToAPIServerProfilePtrOutput() APIServerProfilePtrOutput {
	return o
}

func (o APIServerProfilePtrOutput) ToAPIServerProfilePtrOutputWithContext(ctx context.Context) APIServerProfilePtrOutput {
	return o
}

func (o APIServerProfilePtrOutput) Elem() APIServerProfileOutput {
	return o.ApplyT(func(v *APIServerProfile) APIServerProfile {
		if v != nil {
			return *v
		}
		var ret APIServerProfile
		return ret
	}).(APIServerProfileOutput)
}

func (o APIServerProfilePtrOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APIServerProfile) *string {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(pulumi.StringPtrOutput)
}

func (o APIServerProfilePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APIServerProfile) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

func (o APIServerProfilePtrOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APIServerProfile) *string {
		if v == nil {
			return nil
		}
		return v.Visibility
	}).(pulumi.StringPtrOutput)
}

type APIServerProfileResponse struct {
	Ip         *string `pulumi:"ip"`
	Url        *string `pulumi:"url"`
	Visibility *string `pulumi:"visibility"`
}

type APIServerProfileResponseOutput struct{ *pulumi.OutputState }

func (APIServerProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*APIServerProfileResponse)(nil)).Elem()
}

func (o APIServerProfileResponseOutput) ToAPIServerProfileResponseOutput() APIServerProfileResponseOutput {
	return o
}

func (o APIServerProfileResponseOutput) ToAPIServerProfileResponseOutputWithContext(ctx context.Context) APIServerProfileResponseOutput {
	return o
}

func (o APIServerProfileResponseOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v APIServerProfileResponse) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o APIServerProfileResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v APIServerProfileResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o APIServerProfileResponseOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v APIServerProfileResponse) *string { return v.Visibility }).(pulumi.StringPtrOutput)
}

type APIServerProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (APIServerProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**APIServerProfileResponse)(nil)).Elem()
}

func (o APIServerProfileResponsePtrOutput) ToAPIServerProfileResponsePtrOutput() APIServerProfileResponsePtrOutput {
	return o
}

func (o APIServerProfileResponsePtrOutput) ToAPIServerProfileResponsePtrOutputWithContext(ctx context.Context) APIServerProfileResponsePtrOutput {
	return o
}

func (o APIServerProfileResponsePtrOutput) Elem() APIServerProfileResponseOutput {
	return o.ApplyT(func(v *APIServerProfileResponse) APIServerProfileResponse {
		if v != nil {
			return *v
		}
		var ret APIServerProfileResponse
		return ret
	}).(APIServerProfileResponseOutput)
}

func (o APIServerProfileResponsePtrOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APIServerProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(pulumi.StringPtrOutput)
}

func (o APIServerProfileResponsePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APIServerProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

func (o APIServerProfileResponsePtrOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APIServerProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Visibility
	}).(pulumi.StringPtrOutput)
}

type ClusterProfile struct {
	Domain               *string `pulumi:"domain"`
	FipsValidatedModules *string `pulumi:"fipsValidatedModules"`
	PullSecret           *string `pulumi:"pullSecret"`
	ResourceGroupId      *string `pulumi:"resourceGroupId"`
	Version              *string `pulumi:"version"`
}





type ClusterProfileInput interface {
	pulumi.Input

	ToClusterProfileOutput() ClusterProfileOutput
	ToClusterProfileOutputWithContext(context.Context) ClusterProfileOutput
}

type ClusterProfileArgs struct {
	Domain               pulumi.StringPtrInput `pulumi:"domain"`
	FipsValidatedModules pulumi.StringPtrInput `pulumi:"fipsValidatedModules"`
	PullSecret           pulumi.StringPtrInput `pulumi:"pullSecret"`
	ResourceGroupId      pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	Version              pulumi.StringPtrInput `pulumi:"version"`
}

func (ClusterProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterProfile)(nil)).Elem()
}

func (i ClusterProfileArgs) ToClusterProfileOutput() ClusterProfileOutput {
	return i.ToClusterProfileOutputWithContext(context.Background())
}

func (i ClusterProfileArgs) ToClusterProfileOutputWithContext(ctx context.Context) ClusterProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterProfileOutput)
}

func (i ClusterProfileArgs) ToClusterProfilePtrOutput() ClusterProfilePtrOutput {
	return i.ToClusterProfilePtrOutputWithContext(context.Background())
}

func (i ClusterProfileArgs) ToClusterProfilePtrOutputWithContext(ctx context.Context) ClusterProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterProfileOutput).ToClusterProfilePtrOutputWithContext(ctx)
}









type ClusterProfilePtrInput interface {
	pulumi.Input

	ToClusterProfilePtrOutput() ClusterProfilePtrOutput
	ToClusterProfilePtrOutputWithContext(context.Context) ClusterProfilePtrOutput
}

type clusterProfilePtrType ClusterProfileArgs

func ClusterProfilePtr(v *ClusterProfileArgs) ClusterProfilePtrInput {
	return (*clusterProfilePtrType)(v)
}

func (*clusterProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterProfile)(nil)).Elem()
}

func (i *clusterProfilePtrType) ToClusterProfilePtrOutput() ClusterProfilePtrOutput {
	return i.ToClusterProfilePtrOutputWithContext(context.Background())
}

func (i *clusterProfilePtrType) ToClusterProfilePtrOutputWithContext(ctx context.Context) ClusterProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterProfilePtrOutput)
}

type ClusterProfileOutput struct{ *pulumi.OutputState }

func (ClusterProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterProfile)(nil)).Elem()
}

func (o ClusterProfileOutput) ToClusterProfileOutput() ClusterProfileOutput {
	return o
}

func (o ClusterProfileOutput) ToClusterProfileOutputWithContext(ctx context.Context) ClusterProfileOutput {
	return o
}

func (o ClusterProfileOutput) ToClusterProfilePtrOutput() ClusterProfilePtrOutput {
	return o.ToClusterProfilePtrOutputWithContext(context.Background())
}

func (o ClusterProfileOutput) ToClusterProfilePtrOutputWithContext(ctx context.Context) ClusterProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterProfile) *ClusterProfile {
		return &v
	}).(ClusterProfilePtrOutput)
}

func (o ClusterProfileOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterProfile) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o ClusterProfileOutput) FipsValidatedModules() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterProfile) *string { return v.FipsValidatedModules }).(pulumi.StringPtrOutput)
}

func (o ClusterProfileOutput) PullSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterProfile) *string { return v.PullSecret }).(pulumi.StringPtrOutput)
}

func (o ClusterProfileOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterProfile) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o ClusterProfileOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterProfile) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ClusterProfilePtrOutput struct{ *pulumi.OutputState }

func (ClusterProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterProfile)(nil)).Elem()
}

func (o ClusterProfilePtrOutput) ToClusterProfilePtrOutput() ClusterProfilePtrOutput {
	return o
}

func (o ClusterProfilePtrOutput) ToClusterProfilePtrOutputWithContext(ctx context.Context) ClusterProfilePtrOutput {
	return o
}

func (o ClusterProfilePtrOutput) Elem() ClusterProfileOutput {
	return o.ApplyT(func(v *ClusterProfile) ClusterProfile {
		if v != nil {
			return *v
		}
		var ret ClusterProfile
		return ret
	}).(ClusterProfileOutput)
}

func (o ClusterProfilePtrOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterProfile) *string {
		if v == nil {
			return nil
		}
		return v.Domain
	}).(pulumi.StringPtrOutput)
}

func (o ClusterProfilePtrOutput) FipsValidatedModules() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterProfile) *string {
		if v == nil {
			return nil
		}
		return v.FipsValidatedModules
	}).(pulumi.StringPtrOutput)
}

func (o ClusterProfilePtrOutput) PullSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterProfile) *string {
		if v == nil {
			return nil
		}
		return v.PullSecret
	}).(pulumi.StringPtrOutput)
}

func (o ClusterProfilePtrOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterProfile) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroupId
	}).(pulumi.StringPtrOutput)
}

func (o ClusterProfilePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterProfile) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ClusterProfileResponse struct {
	Domain               *string `pulumi:"domain"`
	FipsValidatedModules *string `pulumi:"fipsValidatedModules"`
	PullSecret           *string `pulumi:"pullSecret"`
	ResourceGroupId      *string `pulumi:"resourceGroupId"`
	Version              *string `pulumi:"version"`
}

type ClusterProfileResponseOutput struct{ *pulumi.OutputState }

func (ClusterProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterProfileResponse)(nil)).Elem()
}

func (o ClusterProfileResponseOutput) ToClusterProfileResponseOutput() ClusterProfileResponseOutput {
	return o
}

func (o ClusterProfileResponseOutput) ToClusterProfileResponseOutputWithContext(ctx context.Context) ClusterProfileResponseOutput {
	return o
}

func (o ClusterProfileResponseOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterProfileResponse) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o ClusterProfileResponseOutput) FipsValidatedModules() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterProfileResponse) *string { return v.FipsValidatedModules }).(pulumi.StringPtrOutput)
}

func (o ClusterProfileResponseOutput) PullSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterProfileResponse) *string { return v.PullSecret }).(pulumi.StringPtrOutput)
}

func (o ClusterProfileResponseOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterProfileResponse) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o ClusterProfileResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterProfileResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ClusterProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterProfileResponse)(nil)).Elem()
}

func (o ClusterProfileResponsePtrOutput) ToClusterProfileResponsePtrOutput() ClusterProfileResponsePtrOutput {
	return o
}

func (o ClusterProfileResponsePtrOutput) ToClusterProfileResponsePtrOutputWithContext(ctx context.Context) ClusterProfileResponsePtrOutput {
	return o
}

func (o ClusterProfileResponsePtrOutput) Elem() ClusterProfileResponseOutput {
	return o.ApplyT(func(v *ClusterProfileResponse) ClusterProfileResponse {
		if v != nil {
			return *v
		}
		var ret ClusterProfileResponse
		return ret
	}).(ClusterProfileResponseOutput)
}

func (o ClusterProfileResponsePtrOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Domain
	}).(pulumi.StringPtrOutput)
}

func (o ClusterProfileResponsePtrOutput) FipsValidatedModules() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.FipsValidatedModules
	}).(pulumi.StringPtrOutput)
}

func (o ClusterProfileResponsePtrOutput) PullSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.PullSecret
	}).(pulumi.StringPtrOutput)
}

func (o ClusterProfileResponsePtrOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroupId
	}).(pulumi.StringPtrOutput)
}

func (o ClusterProfileResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ConsoleProfile struct {
	Url *string `pulumi:"url"`
}





type ConsoleProfileInput interface {
	pulumi.Input

	ToConsoleProfileOutput() ConsoleProfileOutput
	ToConsoleProfileOutputWithContext(context.Context) ConsoleProfileOutput
}

type ConsoleProfileArgs struct {
	Url pulumi.StringPtrInput `pulumi:"url"`
}

func (ConsoleProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsoleProfile)(nil)).Elem()
}

func (i ConsoleProfileArgs) ToConsoleProfileOutput() ConsoleProfileOutput {
	return i.ToConsoleProfileOutputWithContext(context.Background())
}

func (i ConsoleProfileArgs) ToConsoleProfileOutputWithContext(ctx context.Context) ConsoleProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsoleProfileOutput)
}

func (i ConsoleProfileArgs) ToConsoleProfilePtrOutput() ConsoleProfilePtrOutput {
	return i.ToConsoleProfilePtrOutputWithContext(context.Background())
}

func (i ConsoleProfileArgs) ToConsoleProfilePtrOutputWithContext(ctx context.Context) ConsoleProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsoleProfileOutput).ToConsoleProfilePtrOutputWithContext(ctx)
}









type ConsoleProfilePtrInput interface {
	pulumi.Input

	ToConsoleProfilePtrOutput() ConsoleProfilePtrOutput
	ToConsoleProfilePtrOutputWithContext(context.Context) ConsoleProfilePtrOutput
}

type consoleProfilePtrType ConsoleProfileArgs

func ConsoleProfilePtr(v *ConsoleProfileArgs) ConsoleProfilePtrInput {
	return (*consoleProfilePtrType)(v)
}

func (*consoleProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsoleProfile)(nil)).Elem()
}

func (i *consoleProfilePtrType) ToConsoleProfilePtrOutput() ConsoleProfilePtrOutput {
	return i.ToConsoleProfilePtrOutputWithContext(context.Background())
}

func (i *consoleProfilePtrType) ToConsoleProfilePtrOutputWithContext(ctx context.Context) ConsoleProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsoleProfilePtrOutput)
}

type ConsoleProfileOutput struct{ *pulumi.OutputState }

func (ConsoleProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsoleProfile)(nil)).Elem()
}

func (o ConsoleProfileOutput) ToConsoleProfileOutput() ConsoleProfileOutput {
	return o
}

func (o ConsoleProfileOutput) ToConsoleProfileOutputWithContext(ctx context.Context) ConsoleProfileOutput {
	return o
}

func (o ConsoleProfileOutput) ToConsoleProfilePtrOutput() ConsoleProfilePtrOutput {
	return o.ToConsoleProfilePtrOutputWithContext(context.Background())
}

func (o ConsoleProfileOutput) ToConsoleProfilePtrOutputWithContext(ctx context.Context) ConsoleProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConsoleProfile) *ConsoleProfile {
		return &v
	}).(ConsoleProfilePtrOutput)
}

func (o ConsoleProfileOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsoleProfile) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type ConsoleProfilePtrOutput struct{ *pulumi.OutputState }

func (ConsoleProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsoleProfile)(nil)).Elem()
}

func (o ConsoleProfilePtrOutput) ToConsoleProfilePtrOutput() ConsoleProfilePtrOutput {
	return o
}

func (o ConsoleProfilePtrOutput) ToConsoleProfilePtrOutputWithContext(ctx context.Context) ConsoleProfilePtrOutput {
	return o
}

func (o ConsoleProfilePtrOutput) Elem() ConsoleProfileOutput {
	return o.ApplyT(func(v *ConsoleProfile) ConsoleProfile {
		if v != nil {
			return *v
		}
		var ret ConsoleProfile
		return ret
	}).(ConsoleProfileOutput)
}

func (o ConsoleProfilePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsoleProfile) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type ConsoleProfileResponse struct {
	Url *string `pulumi:"url"`
}

type ConsoleProfileResponseOutput struct{ *pulumi.OutputState }

func (ConsoleProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsoleProfileResponse)(nil)).Elem()
}

func (o ConsoleProfileResponseOutput) ToConsoleProfileResponseOutput() ConsoleProfileResponseOutput {
	return o
}

func (o ConsoleProfileResponseOutput) ToConsoleProfileResponseOutputWithContext(ctx context.Context) ConsoleProfileResponseOutput {
	return o
}

func (o ConsoleProfileResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsoleProfileResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type ConsoleProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ConsoleProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsoleProfileResponse)(nil)).Elem()
}

func (o ConsoleProfileResponsePtrOutput) ToConsoleProfileResponsePtrOutput() ConsoleProfileResponsePtrOutput {
	return o
}

func (o ConsoleProfileResponsePtrOutput) ToConsoleProfileResponsePtrOutputWithContext(ctx context.Context) ConsoleProfileResponsePtrOutput {
	return o
}

func (o ConsoleProfileResponsePtrOutput) Elem() ConsoleProfileResponseOutput {
	return o.ApplyT(func(v *ConsoleProfileResponse) ConsoleProfileResponse {
		if v != nil {
			return *v
		}
		var ret ConsoleProfileResponse
		return ret
	}).(ConsoleProfileResponseOutput)
}

func (o ConsoleProfileResponsePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsoleProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type IngressProfile struct {
	Ip         *string `pulumi:"ip"`
	Name       *string `pulumi:"name"`
	Visibility *string `pulumi:"visibility"`
}





type IngressProfileInput interface {
	pulumi.Input

	ToIngressProfileOutput() IngressProfileOutput
	ToIngressProfileOutputWithContext(context.Context) IngressProfileOutput
}

type IngressProfileArgs struct {
	Ip         pulumi.StringPtrInput `pulumi:"ip"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Visibility pulumi.StringPtrInput `pulumi:"visibility"`
}

func (IngressProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressProfile)(nil)).Elem()
}

func (i IngressProfileArgs) ToIngressProfileOutput() IngressProfileOutput {
	return i.ToIngressProfileOutputWithContext(context.Background())
}

func (i IngressProfileArgs) ToIngressProfileOutputWithContext(ctx context.Context) IngressProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressProfileOutput)
}





type IngressProfileArrayInput interface {
	pulumi.Input

	ToIngressProfileArrayOutput() IngressProfileArrayOutput
	ToIngressProfileArrayOutputWithContext(context.Context) IngressProfileArrayOutput
}

type IngressProfileArray []IngressProfileInput

func (IngressProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IngressProfile)(nil)).Elem()
}

func (i IngressProfileArray) ToIngressProfileArrayOutput() IngressProfileArrayOutput {
	return i.ToIngressProfileArrayOutputWithContext(context.Background())
}

func (i IngressProfileArray) ToIngressProfileArrayOutputWithContext(ctx context.Context) IngressProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressProfileArrayOutput)
}

type IngressProfileOutput struct{ *pulumi.OutputState }

func (IngressProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressProfile)(nil)).Elem()
}

func (o IngressProfileOutput) ToIngressProfileOutput() IngressProfileOutput {
	return o
}

func (o IngressProfileOutput) ToIngressProfileOutputWithContext(ctx context.Context) IngressProfileOutput {
	return o
}

func (o IngressProfileOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngressProfile) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o IngressProfileOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngressProfile) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IngressProfileOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngressProfile) *string { return v.Visibility }).(pulumi.StringPtrOutput)
}

type IngressProfileArrayOutput struct{ *pulumi.OutputState }

func (IngressProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IngressProfile)(nil)).Elem()
}

func (o IngressProfileArrayOutput) ToIngressProfileArrayOutput() IngressProfileArrayOutput {
	return o
}

func (o IngressProfileArrayOutput) ToIngressProfileArrayOutputWithContext(ctx context.Context) IngressProfileArrayOutput {
	return o
}

func (o IngressProfileArrayOutput) Index(i pulumi.IntInput) IngressProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IngressProfile {
		return vs[0].([]IngressProfile)[vs[1].(int)]
	}).(IngressProfileOutput)
}

type IngressProfileResponse struct {
	Ip         *string `pulumi:"ip"`
	Name       *string `pulumi:"name"`
	Visibility *string `pulumi:"visibility"`
}

type IngressProfileResponseOutput struct{ *pulumi.OutputState }

func (IngressProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressProfileResponse)(nil)).Elem()
}

func (o IngressProfileResponseOutput) ToIngressProfileResponseOutput() IngressProfileResponseOutput {
	return o
}

func (o IngressProfileResponseOutput) ToIngressProfileResponseOutputWithContext(ctx context.Context) IngressProfileResponseOutput {
	return o
}

func (o IngressProfileResponseOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngressProfileResponse) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o IngressProfileResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngressProfileResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IngressProfileResponseOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngressProfileResponse) *string { return v.Visibility }).(pulumi.StringPtrOutput)
}

type IngressProfileResponseArrayOutput struct{ *pulumi.OutputState }

func (IngressProfileResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IngressProfileResponse)(nil)).Elem()
}

func (o IngressProfileResponseArrayOutput) ToIngressProfileResponseArrayOutput() IngressProfileResponseArrayOutput {
	return o
}

func (o IngressProfileResponseArrayOutput) ToIngressProfileResponseArrayOutputWithContext(ctx context.Context) IngressProfileResponseArrayOutput {
	return o
}

func (o IngressProfileResponseArrayOutput) Index(i pulumi.IntInput) IngressProfileResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IngressProfileResponse {
		return vs[0].([]IngressProfileResponse)[vs[1].(int)]
	}).(IngressProfileResponseOutput)
}

type MasterProfile struct {
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	EncryptionAtHost    *string `pulumi:"encryptionAtHost"`
	SubnetId            *string `pulumi:"subnetId"`
	VmSize              *string `pulumi:"vmSize"`
}





type MasterProfileInput interface {
	pulumi.Input

	ToMasterProfileOutput() MasterProfileOutput
	ToMasterProfileOutputWithContext(context.Context) MasterProfileOutput
}

type MasterProfileArgs struct {
	DiskEncryptionSetId pulumi.StringPtrInput `pulumi:"diskEncryptionSetId"`
	EncryptionAtHost    pulumi.StringPtrInput `pulumi:"encryptionAtHost"`
	SubnetId            pulumi.StringPtrInput `pulumi:"subnetId"`
	VmSize              pulumi.StringPtrInput `pulumi:"vmSize"`
}

func (MasterProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MasterProfile)(nil)).Elem()
}

func (i MasterProfileArgs) ToMasterProfileOutput() MasterProfileOutput {
	return i.ToMasterProfileOutputWithContext(context.Background())
}

func (i MasterProfileArgs) ToMasterProfileOutputWithContext(ctx context.Context) MasterProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MasterProfileOutput)
}

func (i MasterProfileArgs) ToMasterProfilePtrOutput() MasterProfilePtrOutput {
	return i.ToMasterProfilePtrOutputWithContext(context.Background())
}

func (i MasterProfileArgs) ToMasterProfilePtrOutputWithContext(ctx context.Context) MasterProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MasterProfileOutput).ToMasterProfilePtrOutputWithContext(ctx)
}









type MasterProfilePtrInput interface {
	pulumi.Input

	ToMasterProfilePtrOutput() MasterProfilePtrOutput
	ToMasterProfilePtrOutputWithContext(context.Context) MasterProfilePtrOutput
}

type masterProfilePtrType MasterProfileArgs

func MasterProfilePtr(v *MasterProfileArgs) MasterProfilePtrInput {
	return (*masterProfilePtrType)(v)
}

func (*masterProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MasterProfile)(nil)).Elem()
}

func (i *masterProfilePtrType) ToMasterProfilePtrOutput() MasterProfilePtrOutput {
	return i.ToMasterProfilePtrOutputWithContext(context.Background())
}

func (i *masterProfilePtrType) ToMasterProfilePtrOutputWithContext(ctx context.Context) MasterProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MasterProfilePtrOutput)
}

type MasterProfileOutput struct{ *pulumi.OutputState }

func (MasterProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MasterProfile)(nil)).Elem()
}

func (o MasterProfileOutput) ToMasterProfileOutput() MasterProfileOutput {
	return o
}

func (o MasterProfileOutput) ToMasterProfileOutputWithContext(ctx context.Context) MasterProfileOutput {
	return o
}

func (o MasterProfileOutput) ToMasterProfilePtrOutput() MasterProfilePtrOutput {
	return o.ToMasterProfilePtrOutputWithContext(context.Background())
}

func (o MasterProfileOutput) ToMasterProfilePtrOutputWithContext(ctx context.Context) MasterProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MasterProfile) *MasterProfile {
		return &v
	}).(MasterProfilePtrOutput)
}

func (o MasterProfileOutput) DiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MasterProfile) *string { return v.DiskEncryptionSetId }).(pulumi.StringPtrOutput)
}

func (o MasterProfileOutput) EncryptionAtHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MasterProfile) *string { return v.EncryptionAtHost }).(pulumi.StringPtrOutput)
}

func (o MasterProfileOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MasterProfile) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o MasterProfileOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MasterProfile) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type MasterProfilePtrOutput struct{ *pulumi.OutputState }

func (MasterProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MasterProfile)(nil)).Elem()
}

func (o MasterProfilePtrOutput) ToMasterProfilePtrOutput() MasterProfilePtrOutput {
	return o
}

func (o MasterProfilePtrOutput) ToMasterProfilePtrOutputWithContext(ctx context.Context) MasterProfilePtrOutput {
	return o
}

func (o MasterProfilePtrOutput) Elem() MasterProfileOutput {
	return o.ApplyT(func(v *MasterProfile) MasterProfile {
		if v != nil {
			return *v
		}
		var ret MasterProfile
		return ret
	}).(MasterProfileOutput)
}

func (o MasterProfilePtrOutput) DiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MasterProfile) *string {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionSetId
	}).(pulumi.StringPtrOutput)
}

func (o MasterProfilePtrOutput) EncryptionAtHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MasterProfile) *string {
		if v == nil {
			return nil
		}
		return v.EncryptionAtHost
	}).(pulumi.StringPtrOutput)
}

func (o MasterProfilePtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MasterProfile) *string {
		if v == nil {
			return nil
		}
		return v.SubnetId
	}).(pulumi.StringPtrOutput)
}

func (o MasterProfilePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MasterProfile) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type MasterProfileResponse struct {
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	EncryptionAtHost    *string `pulumi:"encryptionAtHost"`
	SubnetId            *string `pulumi:"subnetId"`
	VmSize              *string `pulumi:"vmSize"`
}

type MasterProfileResponseOutput struct{ *pulumi.OutputState }

func (MasterProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MasterProfileResponse)(nil)).Elem()
}

func (o MasterProfileResponseOutput) ToMasterProfileResponseOutput() MasterProfileResponseOutput {
	return o
}

func (o MasterProfileResponseOutput) ToMasterProfileResponseOutputWithContext(ctx context.Context) MasterProfileResponseOutput {
	return o
}

func (o MasterProfileResponseOutput) DiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MasterProfileResponse) *string { return v.DiskEncryptionSetId }).(pulumi.StringPtrOutput)
}

func (o MasterProfileResponseOutput) EncryptionAtHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MasterProfileResponse) *string { return v.EncryptionAtHost }).(pulumi.StringPtrOutput)
}

func (o MasterProfileResponseOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MasterProfileResponse) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o MasterProfileResponseOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MasterProfileResponse) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type MasterProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (MasterProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MasterProfileResponse)(nil)).Elem()
}

func (o MasterProfileResponsePtrOutput) ToMasterProfileResponsePtrOutput() MasterProfileResponsePtrOutput {
	return o
}

func (o MasterProfileResponsePtrOutput) ToMasterProfileResponsePtrOutputWithContext(ctx context.Context) MasterProfileResponsePtrOutput {
	return o
}

func (o MasterProfileResponsePtrOutput) Elem() MasterProfileResponseOutput {
	return o.ApplyT(func(v *MasterProfileResponse) MasterProfileResponse {
		if v != nil {
			return *v
		}
		var ret MasterProfileResponse
		return ret
	}).(MasterProfileResponseOutput)
}

func (o MasterProfileResponsePtrOutput) DiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MasterProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionSetId
	}).(pulumi.StringPtrOutput)
}

func (o MasterProfileResponsePtrOutput) EncryptionAtHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MasterProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.EncryptionAtHost
	}).(pulumi.StringPtrOutput)
}

func (o MasterProfileResponsePtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MasterProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubnetId
	}).(pulumi.StringPtrOutput)
}

func (o MasterProfileResponsePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MasterProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type NetworkProfile struct {
	PodCidr     *string `pulumi:"podCidr"`
	ServiceCidr *string `pulumi:"serviceCidr"`
}





type NetworkProfileInput interface {
	pulumi.Input

	ToNetworkProfileOutput() NetworkProfileOutput
	ToNetworkProfileOutputWithContext(context.Context) NetworkProfileOutput
}

type NetworkProfileArgs struct {
	PodCidr     pulumi.StringPtrInput `pulumi:"podCidr"`
	ServiceCidr pulumi.StringPtrInput `pulumi:"serviceCidr"`
}

func (NetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (i NetworkProfileArgs) ToNetworkProfileOutput() NetworkProfileOutput {
	return i.ToNetworkProfileOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput)
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput).ToNetworkProfilePtrOutputWithContext(ctx)
}









type NetworkProfilePtrInput interface {
	pulumi.Input

	ToNetworkProfilePtrOutput() NetworkProfilePtrOutput
	ToNetworkProfilePtrOutputWithContext(context.Context) NetworkProfilePtrOutput
}

type networkProfilePtrType NetworkProfileArgs

func NetworkProfilePtr(v *NetworkProfileArgs) NetworkProfilePtrInput {
	return (*networkProfilePtrType)(v)
}

func (*networkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfilePtrOutput)
}

type NetworkProfileOutput struct{ *pulumi.OutputState }

func (NetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (o NetworkProfileOutput) ToNetworkProfileOutput() NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkProfile) *NetworkProfile {
		return &v
	}).(NetworkProfilePtrOutput)
}

func (o NetworkProfileOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.PodCidr }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.ServiceCidr }).(pulumi.StringPtrOutput)
}

type NetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) Elem() NetworkProfileOutput {
	return o.ApplyT(func(v *NetworkProfile) NetworkProfile {
		if v != nil {
			return *v
		}
		var ret NetworkProfile
		return ret
	}).(NetworkProfileOutput)
}

func (o NetworkProfilePtrOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.PodCidr
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCidr
	}).(pulumi.StringPtrOutput)
}

type NetworkProfileResponse struct {
	PodCidr     *string `pulumi:"podCidr"`
	ServiceCidr *string `pulumi:"serviceCidr"`
}

type NetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutput() NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutputWithContext(ctx context.Context) NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.PodCidr }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.ServiceCidr }).(pulumi.StringPtrOutput)
}

type NetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutputWithContext(ctx context.Context) NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) Elem() NetworkProfileResponseOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) NetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret NetworkProfileResponse
		return ret
	}).(NetworkProfileResponseOutput)
}

func (o NetworkProfileResponsePtrOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.PodCidr
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCidr
	}).(pulumi.StringPtrOutput)
}

type ServicePrincipalProfile struct {
	ClientId     *string `pulumi:"clientId"`
	ClientSecret *string `pulumi:"clientSecret"`
}





type ServicePrincipalProfileInput interface {
	pulumi.Input

	ToServicePrincipalProfileOutput() ServicePrincipalProfileOutput
	ToServicePrincipalProfileOutputWithContext(context.Context) ServicePrincipalProfileOutput
}

type ServicePrincipalProfileArgs struct {
	ClientId     pulumi.StringPtrInput `pulumi:"clientId"`
	ClientSecret pulumi.StringPtrInput `pulumi:"clientSecret"`
}

func (ServicePrincipalProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalProfile)(nil)).Elem()
}

func (i ServicePrincipalProfileArgs) ToServicePrincipalProfileOutput() ServicePrincipalProfileOutput {
	return i.ToServicePrincipalProfileOutputWithContext(context.Background())
}

func (i ServicePrincipalProfileArgs) ToServicePrincipalProfileOutputWithContext(ctx context.Context) ServicePrincipalProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalProfileOutput)
}

func (i ServicePrincipalProfileArgs) ToServicePrincipalProfilePtrOutput() ServicePrincipalProfilePtrOutput {
	return i.ToServicePrincipalProfilePtrOutputWithContext(context.Background())
}

func (i ServicePrincipalProfileArgs) ToServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ServicePrincipalProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalProfileOutput).ToServicePrincipalProfilePtrOutputWithContext(ctx)
}









type ServicePrincipalProfilePtrInput interface {
	pulumi.Input

	ToServicePrincipalProfilePtrOutput() ServicePrincipalProfilePtrOutput
	ToServicePrincipalProfilePtrOutputWithContext(context.Context) ServicePrincipalProfilePtrOutput
}

type servicePrincipalProfilePtrType ServicePrincipalProfileArgs

func ServicePrincipalProfilePtr(v *ServicePrincipalProfileArgs) ServicePrincipalProfilePtrInput {
	return (*servicePrincipalProfilePtrType)(v)
}

func (*servicePrincipalProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalProfile)(nil)).Elem()
}

func (i *servicePrincipalProfilePtrType) ToServicePrincipalProfilePtrOutput() ServicePrincipalProfilePtrOutput {
	return i.ToServicePrincipalProfilePtrOutputWithContext(context.Background())
}

func (i *servicePrincipalProfilePtrType) ToServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ServicePrincipalProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalProfilePtrOutput)
}

type ServicePrincipalProfileOutput struct{ *pulumi.OutputState }

func (ServicePrincipalProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalProfile)(nil)).Elem()
}

func (o ServicePrincipalProfileOutput) ToServicePrincipalProfileOutput() ServicePrincipalProfileOutput {
	return o
}

func (o ServicePrincipalProfileOutput) ToServicePrincipalProfileOutputWithContext(ctx context.Context) ServicePrincipalProfileOutput {
	return o
}

func (o ServicePrincipalProfileOutput) ToServicePrincipalProfilePtrOutput() ServicePrincipalProfilePtrOutput {
	return o.ToServicePrincipalProfilePtrOutputWithContext(context.Background())
}

func (o ServicePrincipalProfileOutput) ToServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ServicePrincipalProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicePrincipalProfile) *ServicePrincipalProfile {
		return &v
	}).(ServicePrincipalProfilePtrOutput)
}

func (o ServicePrincipalProfileOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalProfile) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalProfileOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalProfile) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

type ServicePrincipalProfilePtrOutput struct{ *pulumi.OutputState }

func (ServicePrincipalProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalProfile)(nil)).Elem()
}

func (o ServicePrincipalProfilePtrOutput) ToServicePrincipalProfilePtrOutput() ServicePrincipalProfilePtrOutput {
	return o
}

func (o ServicePrincipalProfilePtrOutput) ToServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ServicePrincipalProfilePtrOutput {
	return o
}

func (o ServicePrincipalProfilePtrOutput) Elem() ServicePrincipalProfileOutput {
	return o.ApplyT(func(v *ServicePrincipalProfile) ServicePrincipalProfile {
		if v != nil {
			return *v
		}
		var ret ServicePrincipalProfile
		return ret
	}).(ServicePrincipalProfileOutput)
}

func (o ServicePrincipalProfilePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalProfile) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalProfilePtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalProfile) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

type ServicePrincipalProfileResponse struct {
	ClientId     *string `pulumi:"clientId"`
	ClientSecret *string `pulumi:"clientSecret"`
}

type ServicePrincipalProfileResponseOutput struct{ *pulumi.OutputState }

func (ServicePrincipalProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalProfileResponse)(nil)).Elem()
}

func (o ServicePrincipalProfileResponseOutput) ToServicePrincipalProfileResponseOutput() ServicePrincipalProfileResponseOutput {
	return o
}

func (o ServicePrincipalProfileResponseOutput) ToServicePrincipalProfileResponseOutputWithContext(ctx context.Context) ServicePrincipalProfileResponseOutput {
	return o
}

func (o ServicePrincipalProfileResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalProfileResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalProfileResponseOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalProfileResponse) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

type ServicePrincipalProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ServicePrincipalProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalProfileResponse)(nil)).Elem()
}

func (o ServicePrincipalProfileResponsePtrOutput) ToServicePrincipalProfileResponsePtrOutput() ServicePrincipalProfileResponsePtrOutput {
	return o
}

func (o ServicePrincipalProfileResponsePtrOutput) ToServicePrincipalProfileResponsePtrOutputWithContext(ctx context.Context) ServicePrincipalProfileResponsePtrOutput {
	return o
}

func (o ServicePrincipalProfileResponsePtrOutput) Elem() ServicePrincipalProfileResponseOutput {
	return o.ApplyT(func(v *ServicePrincipalProfileResponse) ServicePrincipalProfileResponse {
		if v != nil {
			return *v
		}
		var ret ServicePrincipalProfileResponse
		return ret
	}).(ServicePrincipalProfileResponseOutput)
}

func (o ServicePrincipalProfileResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalProfileResponsePtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
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

type WorkerProfile struct {
	Count               *int    `pulumi:"count"`
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	DiskSizeGB          *int    `pulumi:"diskSizeGB"`
	EncryptionAtHost    *string `pulumi:"encryptionAtHost"`
	Name                *string `pulumi:"name"`
	SubnetId            *string `pulumi:"subnetId"`
	VmSize              *string `pulumi:"vmSize"`
}





type WorkerProfileInput interface {
	pulumi.Input

	ToWorkerProfileOutput() WorkerProfileOutput
	ToWorkerProfileOutputWithContext(context.Context) WorkerProfileOutput
}

type WorkerProfileArgs struct {
	Count               pulumi.IntPtrInput    `pulumi:"count"`
	DiskEncryptionSetId pulumi.StringPtrInput `pulumi:"diskEncryptionSetId"`
	DiskSizeGB          pulumi.IntPtrInput    `pulumi:"diskSizeGB"`
	EncryptionAtHost    pulumi.StringPtrInput `pulumi:"encryptionAtHost"`
	Name                pulumi.StringPtrInput `pulumi:"name"`
	SubnetId            pulumi.StringPtrInput `pulumi:"subnetId"`
	VmSize              pulumi.StringPtrInput `pulumi:"vmSize"`
}

func (WorkerProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkerProfile)(nil)).Elem()
}

func (i WorkerProfileArgs) ToWorkerProfileOutput() WorkerProfileOutput {
	return i.ToWorkerProfileOutputWithContext(context.Background())
}

func (i WorkerProfileArgs) ToWorkerProfileOutputWithContext(ctx context.Context) WorkerProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerProfileOutput)
}





type WorkerProfileArrayInput interface {
	pulumi.Input

	ToWorkerProfileArrayOutput() WorkerProfileArrayOutput
	ToWorkerProfileArrayOutputWithContext(context.Context) WorkerProfileArrayOutput
}

type WorkerProfileArray []WorkerProfileInput

func (WorkerProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkerProfile)(nil)).Elem()
}

func (i WorkerProfileArray) ToWorkerProfileArrayOutput() WorkerProfileArrayOutput {
	return i.ToWorkerProfileArrayOutputWithContext(context.Background())
}

func (i WorkerProfileArray) ToWorkerProfileArrayOutputWithContext(ctx context.Context) WorkerProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerProfileArrayOutput)
}

type WorkerProfileOutput struct{ *pulumi.OutputState }

func (WorkerProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkerProfile)(nil)).Elem()
}

func (o WorkerProfileOutput) ToWorkerProfileOutput() WorkerProfileOutput {
	return o
}

func (o WorkerProfileOutput) ToWorkerProfileOutputWithContext(ctx context.Context) WorkerProfileOutput {
	return o
}

func (o WorkerProfileOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkerProfile) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o WorkerProfileOutput) DiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerProfile) *string { return v.DiskEncryptionSetId }).(pulumi.StringPtrOutput)
}

func (o WorkerProfileOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkerProfile) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o WorkerProfileOutput) EncryptionAtHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerProfile) *string { return v.EncryptionAtHost }).(pulumi.StringPtrOutput)
}

func (o WorkerProfileOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerProfile) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WorkerProfileOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerProfile) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o WorkerProfileOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerProfile) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type WorkerProfileArrayOutput struct{ *pulumi.OutputState }

func (WorkerProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkerProfile)(nil)).Elem()
}

func (o WorkerProfileArrayOutput) ToWorkerProfileArrayOutput() WorkerProfileArrayOutput {
	return o
}

func (o WorkerProfileArrayOutput) ToWorkerProfileArrayOutputWithContext(ctx context.Context) WorkerProfileArrayOutput {
	return o
}

func (o WorkerProfileArrayOutput) Index(i pulumi.IntInput) WorkerProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkerProfile {
		return vs[0].([]WorkerProfile)[vs[1].(int)]
	}).(WorkerProfileOutput)
}

type WorkerProfileResponse struct {
	Count               *int    `pulumi:"count"`
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	DiskSizeGB          *int    `pulumi:"diskSizeGB"`
	EncryptionAtHost    *string `pulumi:"encryptionAtHost"`
	Name                *string `pulumi:"name"`
	SubnetId            *string `pulumi:"subnetId"`
	VmSize              *string `pulumi:"vmSize"`
}

type WorkerProfileResponseOutput struct{ *pulumi.OutputState }

func (WorkerProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkerProfileResponse)(nil)).Elem()
}

func (o WorkerProfileResponseOutput) ToWorkerProfileResponseOutput() WorkerProfileResponseOutput {
	return o
}

func (o WorkerProfileResponseOutput) ToWorkerProfileResponseOutputWithContext(ctx context.Context) WorkerProfileResponseOutput {
	return o
}

func (o WorkerProfileResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkerProfileResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o WorkerProfileResponseOutput) DiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerProfileResponse) *string { return v.DiskEncryptionSetId }).(pulumi.StringPtrOutput)
}

func (o WorkerProfileResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkerProfileResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o WorkerProfileResponseOutput) EncryptionAtHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerProfileResponse) *string { return v.EncryptionAtHost }).(pulumi.StringPtrOutput)
}

func (o WorkerProfileResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerProfileResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WorkerProfileResponseOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerProfileResponse) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o WorkerProfileResponseOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerProfileResponse) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type WorkerProfileResponseArrayOutput struct{ *pulumi.OutputState }

func (WorkerProfileResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkerProfileResponse)(nil)).Elem()
}

func (o WorkerProfileResponseArrayOutput) ToWorkerProfileResponseArrayOutput() WorkerProfileResponseArrayOutput {
	return o
}

func (o WorkerProfileResponseArrayOutput) ToWorkerProfileResponseArrayOutputWithContext(ctx context.Context) WorkerProfileResponseArrayOutput {
	return o
}

func (o WorkerProfileResponseArrayOutput) Index(i pulumi.IntInput) WorkerProfileResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkerProfileResponse {
		return vs[0].([]WorkerProfileResponse)[vs[1].(int)]
	}).(WorkerProfileResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(APIServerProfileOutput{})
	pulumi.RegisterOutputType(APIServerProfilePtrOutput{})
	pulumi.RegisterOutputType(APIServerProfileResponseOutput{})
	pulumi.RegisterOutputType(APIServerProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ClusterProfileOutput{})
	pulumi.RegisterOutputType(ClusterProfilePtrOutput{})
	pulumi.RegisterOutputType(ClusterProfileResponseOutput{})
	pulumi.RegisterOutputType(ClusterProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ConsoleProfileOutput{})
	pulumi.RegisterOutputType(ConsoleProfilePtrOutput{})
	pulumi.RegisterOutputType(ConsoleProfileResponseOutput{})
	pulumi.RegisterOutputType(ConsoleProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(IngressProfileOutput{})
	pulumi.RegisterOutputType(IngressProfileArrayOutput{})
	pulumi.RegisterOutputType(IngressProfileResponseOutput{})
	pulumi.RegisterOutputType(IngressProfileResponseArrayOutput{})
	pulumi.RegisterOutputType(MasterProfileOutput{})
	pulumi.RegisterOutputType(MasterProfilePtrOutput{})
	pulumi.RegisterOutputType(MasterProfileResponseOutput{})
	pulumi.RegisterOutputType(MasterProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkProfileOutput{})
	pulumi.RegisterOutputType(NetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ServicePrincipalProfileOutput{})
	pulumi.RegisterOutputType(ServicePrincipalProfilePtrOutput{})
	pulumi.RegisterOutputType(ServicePrincipalProfileResponseOutput{})
	pulumi.RegisterOutputType(ServicePrincipalProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(WorkerProfileOutput{})
	pulumi.RegisterOutputType(WorkerProfileArrayOutput{})
	pulumi.RegisterOutputType(WorkerProfileResponseOutput{})
	pulumi.RegisterOutputType(WorkerProfileResponseArrayOutput{})
}
