


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AutoShutdownProfile struct {
	DisconnectDelay          *string             `pulumi:"disconnectDelay"`
	IdleDelay                *string             `pulumi:"idleDelay"`
	NoConnectDelay           *string             `pulumi:"noConnectDelay"`
	ShutdownOnDisconnect     *EnableState        `pulumi:"shutdownOnDisconnect"`
	ShutdownOnIdle           *ShutdownOnIdleMode `pulumi:"shutdownOnIdle"`
	ShutdownWhenNotConnected *EnableState        `pulumi:"shutdownWhenNotConnected"`
}





type AutoShutdownProfileInput interface {
	pulumi.Input

	ToAutoShutdownProfileOutput() AutoShutdownProfileOutput
	ToAutoShutdownProfileOutputWithContext(context.Context) AutoShutdownProfileOutput
}

type AutoShutdownProfileArgs struct {
	DisconnectDelay          pulumi.StringPtrInput      `pulumi:"disconnectDelay"`
	IdleDelay                pulumi.StringPtrInput      `pulumi:"idleDelay"`
	NoConnectDelay           pulumi.StringPtrInput      `pulumi:"noConnectDelay"`
	ShutdownOnDisconnect     EnableStatePtrInput        `pulumi:"shutdownOnDisconnect"`
	ShutdownOnIdle           ShutdownOnIdleModePtrInput `pulumi:"shutdownOnIdle"`
	ShutdownWhenNotConnected EnableStatePtrInput        `pulumi:"shutdownWhenNotConnected"`
}

func (AutoShutdownProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoShutdownProfile)(nil)).Elem()
}

func (i AutoShutdownProfileArgs) ToAutoShutdownProfileOutput() AutoShutdownProfileOutput {
	return i.ToAutoShutdownProfileOutputWithContext(context.Background())
}

func (i AutoShutdownProfileArgs) ToAutoShutdownProfileOutputWithContext(ctx context.Context) AutoShutdownProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoShutdownProfileOutput)
}

func (i AutoShutdownProfileArgs) ToAutoShutdownProfilePtrOutput() AutoShutdownProfilePtrOutput {
	return i.ToAutoShutdownProfilePtrOutputWithContext(context.Background())
}

func (i AutoShutdownProfileArgs) ToAutoShutdownProfilePtrOutputWithContext(ctx context.Context) AutoShutdownProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoShutdownProfileOutput).ToAutoShutdownProfilePtrOutputWithContext(ctx)
}









type AutoShutdownProfilePtrInput interface {
	pulumi.Input

	ToAutoShutdownProfilePtrOutput() AutoShutdownProfilePtrOutput
	ToAutoShutdownProfilePtrOutputWithContext(context.Context) AutoShutdownProfilePtrOutput
}

type autoShutdownProfilePtrType AutoShutdownProfileArgs

func AutoShutdownProfilePtr(v *AutoShutdownProfileArgs) AutoShutdownProfilePtrInput {
	return (*autoShutdownProfilePtrType)(v)
}

func (*autoShutdownProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoShutdownProfile)(nil)).Elem()
}

func (i *autoShutdownProfilePtrType) ToAutoShutdownProfilePtrOutput() AutoShutdownProfilePtrOutput {
	return i.ToAutoShutdownProfilePtrOutputWithContext(context.Background())
}

func (i *autoShutdownProfilePtrType) ToAutoShutdownProfilePtrOutputWithContext(ctx context.Context) AutoShutdownProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoShutdownProfilePtrOutput)
}

type AutoShutdownProfileOutput struct{ *pulumi.OutputState }

func (AutoShutdownProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoShutdownProfile)(nil)).Elem()
}

func (o AutoShutdownProfileOutput) ToAutoShutdownProfileOutput() AutoShutdownProfileOutput {
	return o
}

func (o AutoShutdownProfileOutput) ToAutoShutdownProfileOutputWithContext(ctx context.Context) AutoShutdownProfileOutput {
	return o
}

func (o AutoShutdownProfileOutput) ToAutoShutdownProfilePtrOutput() AutoShutdownProfilePtrOutput {
	return o.ToAutoShutdownProfilePtrOutputWithContext(context.Background())
}

func (o AutoShutdownProfileOutput) ToAutoShutdownProfilePtrOutputWithContext(ctx context.Context) AutoShutdownProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoShutdownProfile) *AutoShutdownProfile {
		return &v
	}).(AutoShutdownProfilePtrOutput)
}

func (o AutoShutdownProfileOutput) DisconnectDelay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoShutdownProfile) *string { return v.DisconnectDelay }).(pulumi.StringPtrOutput)
}

func (o AutoShutdownProfileOutput) IdleDelay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoShutdownProfile) *string { return v.IdleDelay }).(pulumi.StringPtrOutput)
}

func (o AutoShutdownProfileOutput) NoConnectDelay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoShutdownProfile) *string { return v.NoConnectDelay }).(pulumi.StringPtrOutput)
}

func (o AutoShutdownProfileOutput) ShutdownOnDisconnect() EnableStatePtrOutput {
	return o.ApplyT(func(v AutoShutdownProfile) *EnableState { return v.ShutdownOnDisconnect }).(EnableStatePtrOutput)
}

func (o AutoShutdownProfileOutput) ShutdownOnIdle() ShutdownOnIdleModePtrOutput {
	return o.ApplyT(func(v AutoShutdownProfile) *ShutdownOnIdleMode { return v.ShutdownOnIdle }).(ShutdownOnIdleModePtrOutput)
}

func (o AutoShutdownProfileOutput) ShutdownWhenNotConnected() EnableStatePtrOutput {
	return o.ApplyT(func(v AutoShutdownProfile) *EnableState { return v.ShutdownWhenNotConnected }).(EnableStatePtrOutput)
}

type AutoShutdownProfilePtrOutput struct{ *pulumi.OutputState }

func (AutoShutdownProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoShutdownProfile)(nil)).Elem()
}

func (o AutoShutdownProfilePtrOutput) ToAutoShutdownProfilePtrOutput() AutoShutdownProfilePtrOutput {
	return o
}

func (o AutoShutdownProfilePtrOutput) ToAutoShutdownProfilePtrOutputWithContext(ctx context.Context) AutoShutdownProfilePtrOutput {
	return o
}

func (o AutoShutdownProfilePtrOutput) Elem() AutoShutdownProfileOutput {
	return o.ApplyT(func(v *AutoShutdownProfile) AutoShutdownProfile {
		if v != nil {
			return *v
		}
		var ret AutoShutdownProfile
		return ret
	}).(AutoShutdownProfileOutput)
}

func (o AutoShutdownProfilePtrOutput) DisconnectDelay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoShutdownProfile) *string {
		if v == nil {
			return nil
		}
		return v.DisconnectDelay
	}).(pulumi.StringPtrOutput)
}

func (o AutoShutdownProfilePtrOutput) IdleDelay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoShutdownProfile) *string {
		if v == nil {
			return nil
		}
		return v.IdleDelay
	}).(pulumi.StringPtrOutput)
}

func (o AutoShutdownProfilePtrOutput) NoConnectDelay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoShutdownProfile) *string {
		if v == nil {
			return nil
		}
		return v.NoConnectDelay
	}).(pulumi.StringPtrOutput)
}

func (o AutoShutdownProfilePtrOutput) ShutdownOnDisconnect() EnableStatePtrOutput {
	return o.ApplyT(func(v *AutoShutdownProfile) *EnableState {
		if v == nil {
			return nil
		}
		return v.ShutdownOnDisconnect
	}).(EnableStatePtrOutput)
}

func (o AutoShutdownProfilePtrOutput) ShutdownOnIdle() ShutdownOnIdleModePtrOutput {
	return o.ApplyT(func(v *AutoShutdownProfile) *ShutdownOnIdleMode {
		if v == nil {
			return nil
		}
		return v.ShutdownOnIdle
	}).(ShutdownOnIdleModePtrOutput)
}

func (o AutoShutdownProfilePtrOutput) ShutdownWhenNotConnected() EnableStatePtrOutput {
	return o.ApplyT(func(v *AutoShutdownProfile) *EnableState {
		if v == nil {
			return nil
		}
		return v.ShutdownWhenNotConnected
	}).(EnableStatePtrOutput)
}

type AutoShutdownProfileResponse struct {
	DisconnectDelay          *string `pulumi:"disconnectDelay"`
	IdleDelay                *string `pulumi:"idleDelay"`
	NoConnectDelay           *string `pulumi:"noConnectDelay"`
	ShutdownOnDisconnect     *string `pulumi:"shutdownOnDisconnect"`
	ShutdownOnIdle           *string `pulumi:"shutdownOnIdle"`
	ShutdownWhenNotConnected *string `pulumi:"shutdownWhenNotConnected"`
}





type AutoShutdownProfileResponseInput interface {
	pulumi.Input

	ToAutoShutdownProfileResponseOutput() AutoShutdownProfileResponseOutput
	ToAutoShutdownProfileResponseOutputWithContext(context.Context) AutoShutdownProfileResponseOutput
}

type AutoShutdownProfileResponseArgs struct {
	DisconnectDelay          pulumi.StringPtrInput `pulumi:"disconnectDelay"`
	IdleDelay                pulumi.StringPtrInput `pulumi:"idleDelay"`
	NoConnectDelay           pulumi.StringPtrInput `pulumi:"noConnectDelay"`
	ShutdownOnDisconnect     pulumi.StringPtrInput `pulumi:"shutdownOnDisconnect"`
	ShutdownOnIdle           pulumi.StringPtrInput `pulumi:"shutdownOnIdle"`
	ShutdownWhenNotConnected pulumi.StringPtrInput `pulumi:"shutdownWhenNotConnected"`
}

func (AutoShutdownProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoShutdownProfileResponse)(nil)).Elem()
}

func (i AutoShutdownProfileResponseArgs) ToAutoShutdownProfileResponseOutput() AutoShutdownProfileResponseOutput {
	return i.ToAutoShutdownProfileResponseOutputWithContext(context.Background())
}

func (i AutoShutdownProfileResponseArgs) ToAutoShutdownProfileResponseOutputWithContext(ctx context.Context) AutoShutdownProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoShutdownProfileResponseOutput)
}

func (i AutoShutdownProfileResponseArgs) ToAutoShutdownProfileResponsePtrOutput() AutoShutdownProfileResponsePtrOutput {
	return i.ToAutoShutdownProfileResponsePtrOutputWithContext(context.Background())
}

func (i AutoShutdownProfileResponseArgs) ToAutoShutdownProfileResponsePtrOutputWithContext(ctx context.Context) AutoShutdownProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoShutdownProfileResponseOutput).ToAutoShutdownProfileResponsePtrOutputWithContext(ctx)
}









type AutoShutdownProfileResponsePtrInput interface {
	pulumi.Input

	ToAutoShutdownProfileResponsePtrOutput() AutoShutdownProfileResponsePtrOutput
	ToAutoShutdownProfileResponsePtrOutputWithContext(context.Context) AutoShutdownProfileResponsePtrOutput
}

type autoShutdownProfileResponsePtrType AutoShutdownProfileResponseArgs

func AutoShutdownProfileResponsePtr(v *AutoShutdownProfileResponseArgs) AutoShutdownProfileResponsePtrInput {
	return (*autoShutdownProfileResponsePtrType)(v)
}

func (*autoShutdownProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoShutdownProfileResponse)(nil)).Elem()
}

func (i *autoShutdownProfileResponsePtrType) ToAutoShutdownProfileResponsePtrOutput() AutoShutdownProfileResponsePtrOutput {
	return i.ToAutoShutdownProfileResponsePtrOutputWithContext(context.Background())
}

func (i *autoShutdownProfileResponsePtrType) ToAutoShutdownProfileResponsePtrOutputWithContext(ctx context.Context) AutoShutdownProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoShutdownProfileResponsePtrOutput)
}

type AutoShutdownProfileResponseOutput struct{ *pulumi.OutputState }

func (AutoShutdownProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoShutdownProfileResponse)(nil)).Elem()
}

func (o AutoShutdownProfileResponseOutput) ToAutoShutdownProfileResponseOutput() AutoShutdownProfileResponseOutput {
	return o
}

func (o AutoShutdownProfileResponseOutput) ToAutoShutdownProfileResponseOutputWithContext(ctx context.Context) AutoShutdownProfileResponseOutput {
	return o
}

func (o AutoShutdownProfileResponseOutput) ToAutoShutdownProfileResponsePtrOutput() AutoShutdownProfileResponsePtrOutput {
	return o.ToAutoShutdownProfileResponsePtrOutputWithContext(context.Background())
}

func (o AutoShutdownProfileResponseOutput) ToAutoShutdownProfileResponsePtrOutputWithContext(ctx context.Context) AutoShutdownProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoShutdownProfileResponse) *AutoShutdownProfileResponse {
		return &v
	}).(AutoShutdownProfileResponsePtrOutput)
}

func (o AutoShutdownProfileResponseOutput) DisconnectDelay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoShutdownProfileResponse) *string { return v.DisconnectDelay }).(pulumi.StringPtrOutput)
}

func (o AutoShutdownProfileResponseOutput) IdleDelay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoShutdownProfileResponse) *string { return v.IdleDelay }).(pulumi.StringPtrOutput)
}

func (o AutoShutdownProfileResponseOutput) NoConnectDelay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoShutdownProfileResponse) *string { return v.NoConnectDelay }).(pulumi.StringPtrOutput)
}

func (o AutoShutdownProfileResponseOutput) ShutdownOnDisconnect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoShutdownProfileResponse) *string { return v.ShutdownOnDisconnect }).(pulumi.StringPtrOutput)
}

func (o AutoShutdownProfileResponseOutput) ShutdownOnIdle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoShutdownProfileResponse) *string { return v.ShutdownOnIdle }).(pulumi.StringPtrOutput)
}

func (o AutoShutdownProfileResponseOutput) ShutdownWhenNotConnected() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoShutdownProfileResponse) *string { return v.ShutdownWhenNotConnected }).(pulumi.StringPtrOutput)
}

type AutoShutdownProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoShutdownProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoShutdownProfileResponse)(nil)).Elem()
}

func (o AutoShutdownProfileResponsePtrOutput) ToAutoShutdownProfileResponsePtrOutput() AutoShutdownProfileResponsePtrOutput {
	return o
}

func (o AutoShutdownProfileResponsePtrOutput) ToAutoShutdownProfileResponsePtrOutputWithContext(ctx context.Context) AutoShutdownProfileResponsePtrOutput {
	return o
}

func (o AutoShutdownProfileResponsePtrOutput) Elem() AutoShutdownProfileResponseOutput {
	return o.ApplyT(func(v *AutoShutdownProfileResponse) AutoShutdownProfileResponse {
		if v != nil {
			return *v
		}
		var ret AutoShutdownProfileResponse
		return ret
	}).(AutoShutdownProfileResponseOutput)
}

func (o AutoShutdownProfileResponsePtrOutput) DisconnectDelay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoShutdownProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisconnectDelay
	}).(pulumi.StringPtrOutput)
}

func (o AutoShutdownProfileResponsePtrOutput) IdleDelay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoShutdownProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.IdleDelay
	}).(pulumi.StringPtrOutput)
}

func (o AutoShutdownProfileResponsePtrOutput) NoConnectDelay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoShutdownProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.NoConnectDelay
	}).(pulumi.StringPtrOutput)
}

func (o AutoShutdownProfileResponsePtrOutput) ShutdownOnDisconnect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoShutdownProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ShutdownOnDisconnect
	}).(pulumi.StringPtrOutput)
}

func (o AutoShutdownProfileResponsePtrOutput) ShutdownOnIdle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoShutdownProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ShutdownOnIdle
	}).(pulumi.StringPtrOutput)
}

func (o AutoShutdownProfileResponsePtrOutput) ShutdownWhenNotConnected() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoShutdownProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ShutdownWhenNotConnected
	}).(pulumi.StringPtrOutput)
}

type ConnectionProfile struct {
	ClientRdpAccess *ConnectionType `pulumi:"clientRdpAccess"`
	ClientSshAccess *ConnectionType `pulumi:"clientSshAccess"`
	WebRdpAccess    *ConnectionType `pulumi:"webRdpAccess"`
	WebSshAccess    *ConnectionType `pulumi:"webSshAccess"`
}





type ConnectionProfileInput interface {
	pulumi.Input

	ToConnectionProfileOutput() ConnectionProfileOutput
	ToConnectionProfileOutputWithContext(context.Context) ConnectionProfileOutput
}

type ConnectionProfileArgs struct {
	ClientRdpAccess ConnectionTypePtrInput `pulumi:"clientRdpAccess"`
	ClientSshAccess ConnectionTypePtrInput `pulumi:"clientSshAccess"`
	WebRdpAccess    ConnectionTypePtrInput `pulumi:"webRdpAccess"`
	WebSshAccess    ConnectionTypePtrInput `pulumi:"webSshAccess"`
}

func (ConnectionProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionProfile)(nil)).Elem()
}

func (i ConnectionProfileArgs) ToConnectionProfileOutput() ConnectionProfileOutput {
	return i.ToConnectionProfileOutputWithContext(context.Background())
}

func (i ConnectionProfileArgs) ToConnectionProfileOutputWithContext(ctx context.Context) ConnectionProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionProfileOutput)
}

func (i ConnectionProfileArgs) ToConnectionProfilePtrOutput() ConnectionProfilePtrOutput {
	return i.ToConnectionProfilePtrOutputWithContext(context.Background())
}

func (i ConnectionProfileArgs) ToConnectionProfilePtrOutputWithContext(ctx context.Context) ConnectionProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionProfileOutput).ToConnectionProfilePtrOutputWithContext(ctx)
}









type ConnectionProfilePtrInput interface {
	pulumi.Input

	ToConnectionProfilePtrOutput() ConnectionProfilePtrOutput
	ToConnectionProfilePtrOutputWithContext(context.Context) ConnectionProfilePtrOutput
}

type connectionProfilePtrType ConnectionProfileArgs

func ConnectionProfilePtr(v *ConnectionProfileArgs) ConnectionProfilePtrInput {
	return (*connectionProfilePtrType)(v)
}

func (*connectionProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionProfile)(nil)).Elem()
}

func (i *connectionProfilePtrType) ToConnectionProfilePtrOutput() ConnectionProfilePtrOutput {
	return i.ToConnectionProfilePtrOutputWithContext(context.Background())
}

func (i *connectionProfilePtrType) ToConnectionProfilePtrOutputWithContext(ctx context.Context) ConnectionProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionProfilePtrOutput)
}

type ConnectionProfileOutput struct{ *pulumi.OutputState }

func (ConnectionProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionProfile)(nil)).Elem()
}

func (o ConnectionProfileOutput) ToConnectionProfileOutput() ConnectionProfileOutput {
	return o
}

func (o ConnectionProfileOutput) ToConnectionProfileOutputWithContext(ctx context.Context) ConnectionProfileOutput {
	return o
}

func (o ConnectionProfileOutput) ToConnectionProfilePtrOutput() ConnectionProfilePtrOutput {
	return o.ToConnectionProfilePtrOutputWithContext(context.Background())
}

func (o ConnectionProfileOutput) ToConnectionProfilePtrOutputWithContext(ctx context.Context) ConnectionProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionProfile) *ConnectionProfile {
		return &v
	}).(ConnectionProfilePtrOutput)
}

func (o ConnectionProfileOutput) ClientRdpAccess() ConnectionTypePtrOutput {
	return o.ApplyT(func(v ConnectionProfile) *ConnectionType { return v.ClientRdpAccess }).(ConnectionTypePtrOutput)
}

func (o ConnectionProfileOutput) ClientSshAccess() ConnectionTypePtrOutput {
	return o.ApplyT(func(v ConnectionProfile) *ConnectionType { return v.ClientSshAccess }).(ConnectionTypePtrOutput)
}

func (o ConnectionProfileOutput) WebRdpAccess() ConnectionTypePtrOutput {
	return o.ApplyT(func(v ConnectionProfile) *ConnectionType { return v.WebRdpAccess }).(ConnectionTypePtrOutput)
}

func (o ConnectionProfileOutput) WebSshAccess() ConnectionTypePtrOutput {
	return o.ApplyT(func(v ConnectionProfile) *ConnectionType { return v.WebSshAccess }).(ConnectionTypePtrOutput)
}

type ConnectionProfilePtrOutput struct{ *pulumi.OutputState }

func (ConnectionProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionProfile)(nil)).Elem()
}

func (o ConnectionProfilePtrOutput) ToConnectionProfilePtrOutput() ConnectionProfilePtrOutput {
	return o
}

func (o ConnectionProfilePtrOutput) ToConnectionProfilePtrOutputWithContext(ctx context.Context) ConnectionProfilePtrOutput {
	return o
}

func (o ConnectionProfilePtrOutput) Elem() ConnectionProfileOutput {
	return o.ApplyT(func(v *ConnectionProfile) ConnectionProfile {
		if v != nil {
			return *v
		}
		var ret ConnectionProfile
		return ret
	}).(ConnectionProfileOutput)
}

func (o ConnectionProfilePtrOutput) ClientRdpAccess() ConnectionTypePtrOutput {
	return o.ApplyT(func(v *ConnectionProfile) *ConnectionType {
		if v == nil {
			return nil
		}
		return v.ClientRdpAccess
	}).(ConnectionTypePtrOutput)
}

func (o ConnectionProfilePtrOutput) ClientSshAccess() ConnectionTypePtrOutput {
	return o.ApplyT(func(v *ConnectionProfile) *ConnectionType {
		if v == nil {
			return nil
		}
		return v.ClientSshAccess
	}).(ConnectionTypePtrOutput)
}

func (o ConnectionProfilePtrOutput) WebRdpAccess() ConnectionTypePtrOutput {
	return o.ApplyT(func(v *ConnectionProfile) *ConnectionType {
		if v == nil {
			return nil
		}
		return v.WebRdpAccess
	}).(ConnectionTypePtrOutput)
}

func (o ConnectionProfilePtrOutput) WebSshAccess() ConnectionTypePtrOutput {
	return o.ApplyT(func(v *ConnectionProfile) *ConnectionType {
		if v == nil {
			return nil
		}
		return v.WebSshAccess
	}).(ConnectionTypePtrOutput)
}

type ConnectionProfileResponse struct {
	ClientRdpAccess *string `pulumi:"clientRdpAccess"`
	ClientSshAccess *string `pulumi:"clientSshAccess"`
	WebRdpAccess    *string `pulumi:"webRdpAccess"`
	WebSshAccess    *string `pulumi:"webSshAccess"`
}





type ConnectionProfileResponseInput interface {
	pulumi.Input

	ToConnectionProfileResponseOutput() ConnectionProfileResponseOutput
	ToConnectionProfileResponseOutputWithContext(context.Context) ConnectionProfileResponseOutput
}

type ConnectionProfileResponseArgs struct {
	ClientRdpAccess pulumi.StringPtrInput `pulumi:"clientRdpAccess"`
	ClientSshAccess pulumi.StringPtrInput `pulumi:"clientSshAccess"`
	WebRdpAccess    pulumi.StringPtrInput `pulumi:"webRdpAccess"`
	WebSshAccess    pulumi.StringPtrInput `pulumi:"webSshAccess"`
}

func (ConnectionProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionProfileResponse)(nil)).Elem()
}

func (i ConnectionProfileResponseArgs) ToConnectionProfileResponseOutput() ConnectionProfileResponseOutput {
	return i.ToConnectionProfileResponseOutputWithContext(context.Background())
}

func (i ConnectionProfileResponseArgs) ToConnectionProfileResponseOutputWithContext(ctx context.Context) ConnectionProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionProfileResponseOutput)
}

func (i ConnectionProfileResponseArgs) ToConnectionProfileResponsePtrOutput() ConnectionProfileResponsePtrOutput {
	return i.ToConnectionProfileResponsePtrOutputWithContext(context.Background())
}

func (i ConnectionProfileResponseArgs) ToConnectionProfileResponsePtrOutputWithContext(ctx context.Context) ConnectionProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionProfileResponseOutput).ToConnectionProfileResponsePtrOutputWithContext(ctx)
}









type ConnectionProfileResponsePtrInput interface {
	pulumi.Input

	ToConnectionProfileResponsePtrOutput() ConnectionProfileResponsePtrOutput
	ToConnectionProfileResponsePtrOutputWithContext(context.Context) ConnectionProfileResponsePtrOutput
}

type connectionProfileResponsePtrType ConnectionProfileResponseArgs

func ConnectionProfileResponsePtr(v *ConnectionProfileResponseArgs) ConnectionProfileResponsePtrInput {
	return (*connectionProfileResponsePtrType)(v)
}

func (*connectionProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionProfileResponse)(nil)).Elem()
}

func (i *connectionProfileResponsePtrType) ToConnectionProfileResponsePtrOutput() ConnectionProfileResponsePtrOutput {
	return i.ToConnectionProfileResponsePtrOutputWithContext(context.Background())
}

func (i *connectionProfileResponsePtrType) ToConnectionProfileResponsePtrOutputWithContext(ctx context.Context) ConnectionProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionProfileResponsePtrOutput)
}

type ConnectionProfileResponseOutput struct{ *pulumi.OutputState }

func (ConnectionProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionProfileResponse)(nil)).Elem()
}

func (o ConnectionProfileResponseOutput) ToConnectionProfileResponseOutput() ConnectionProfileResponseOutput {
	return o
}

func (o ConnectionProfileResponseOutput) ToConnectionProfileResponseOutputWithContext(ctx context.Context) ConnectionProfileResponseOutput {
	return o
}

func (o ConnectionProfileResponseOutput) ToConnectionProfileResponsePtrOutput() ConnectionProfileResponsePtrOutput {
	return o.ToConnectionProfileResponsePtrOutputWithContext(context.Background())
}

func (o ConnectionProfileResponseOutput) ToConnectionProfileResponsePtrOutputWithContext(ctx context.Context) ConnectionProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionProfileResponse) *ConnectionProfileResponse {
		return &v
	}).(ConnectionProfileResponsePtrOutput)
}

func (o ConnectionProfileResponseOutput) ClientRdpAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionProfileResponse) *string { return v.ClientRdpAccess }).(pulumi.StringPtrOutput)
}

func (o ConnectionProfileResponseOutput) ClientSshAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionProfileResponse) *string { return v.ClientSshAccess }).(pulumi.StringPtrOutput)
}

func (o ConnectionProfileResponseOutput) WebRdpAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionProfileResponse) *string { return v.WebRdpAccess }).(pulumi.StringPtrOutput)
}

func (o ConnectionProfileResponseOutput) WebSshAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionProfileResponse) *string { return v.WebSshAccess }).(pulumi.StringPtrOutput)
}

type ConnectionProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectionProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionProfileResponse)(nil)).Elem()
}

func (o ConnectionProfileResponsePtrOutput) ToConnectionProfileResponsePtrOutput() ConnectionProfileResponsePtrOutput {
	return o
}

func (o ConnectionProfileResponsePtrOutput) ToConnectionProfileResponsePtrOutputWithContext(ctx context.Context) ConnectionProfileResponsePtrOutput {
	return o
}

func (o ConnectionProfileResponsePtrOutput) Elem() ConnectionProfileResponseOutput {
	return o.ApplyT(func(v *ConnectionProfileResponse) ConnectionProfileResponse {
		if v != nil {
			return *v
		}
		var ret ConnectionProfileResponse
		return ret
	}).(ConnectionProfileResponseOutput)
}

func (o ConnectionProfileResponsePtrOutput) ClientRdpAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientRdpAccess
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionProfileResponsePtrOutput) ClientSshAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSshAccess
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionProfileResponsePtrOutput) WebRdpAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.WebRdpAccess
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionProfileResponsePtrOutput) WebSshAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.WebSshAccess
	}).(pulumi.StringPtrOutput)
}

type Credentials struct {
	Password *string `pulumi:"password"`
	Username string  `pulumi:"username"`
}





type CredentialsInput interface {
	pulumi.Input

	ToCredentialsOutput() CredentialsOutput
	ToCredentialsOutputWithContext(context.Context) CredentialsOutput
}

type CredentialsArgs struct {
	Password pulumi.StringPtrInput `pulumi:"password"`
	Username pulumi.StringInput    `pulumi:"username"`
}

func (CredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Credentials)(nil)).Elem()
}

func (i CredentialsArgs) ToCredentialsOutput() CredentialsOutput {
	return i.ToCredentialsOutputWithContext(context.Background())
}

func (i CredentialsArgs) ToCredentialsOutputWithContext(ctx context.Context) CredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialsOutput)
}

func (i CredentialsArgs) ToCredentialsPtrOutput() CredentialsPtrOutput {
	return i.ToCredentialsPtrOutputWithContext(context.Background())
}

func (i CredentialsArgs) ToCredentialsPtrOutputWithContext(ctx context.Context) CredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialsOutput).ToCredentialsPtrOutputWithContext(ctx)
}









type CredentialsPtrInput interface {
	pulumi.Input

	ToCredentialsPtrOutput() CredentialsPtrOutput
	ToCredentialsPtrOutputWithContext(context.Context) CredentialsPtrOutput
}

type credentialsPtrType CredentialsArgs

func CredentialsPtr(v *CredentialsArgs) CredentialsPtrInput {
	return (*credentialsPtrType)(v)
}

func (*credentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Credentials)(nil)).Elem()
}

func (i *credentialsPtrType) ToCredentialsPtrOutput() CredentialsPtrOutput {
	return i.ToCredentialsPtrOutputWithContext(context.Background())
}

func (i *credentialsPtrType) ToCredentialsPtrOutputWithContext(ctx context.Context) CredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialsPtrOutput)
}

type CredentialsOutput struct{ *pulumi.OutputState }

func (CredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Credentials)(nil)).Elem()
}

func (o CredentialsOutput) ToCredentialsOutput() CredentialsOutput {
	return o
}

func (o CredentialsOutput) ToCredentialsOutputWithContext(ctx context.Context) CredentialsOutput {
	return o
}

func (o CredentialsOutput) ToCredentialsPtrOutput() CredentialsPtrOutput {
	return o.ToCredentialsPtrOutputWithContext(context.Background())
}

func (o CredentialsOutput) ToCredentialsPtrOutputWithContext(ctx context.Context) CredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Credentials) *Credentials {
		return &v
	}).(CredentialsPtrOutput)
}

func (o CredentialsOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Credentials) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o CredentialsOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v Credentials) string { return v.Username }).(pulumi.StringOutput)
}

type CredentialsPtrOutput struct{ *pulumi.OutputState }

func (CredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Credentials)(nil)).Elem()
}

func (o CredentialsPtrOutput) ToCredentialsPtrOutput() CredentialsPtrOutput {
	return o
}

func (o CredentialsPtrOutput) ToCredentialsPtrOutputWithContext(ctx context.Context) CredentialsPtrOutput {
	return o
}

func (o CredentialsPtrOutput) Elem() CredentialsOutput {
	return o.ApplyT(func(v *Credentials) Credentials {
		if v != nil {
			return *v
		}
		var ret Credentials
		return ret
	}).(CredentialsOutput)
}

func (o CredentialsPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Credentials) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o CredentialsPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Credentials) *string {
		if v == nil {
			return nil
		}
		return &v.Username
	}).(pulumi.StringPtrOutput)
}

type CredentialsResponse struct {
	Username string `pulumi:"username"`
}





type CredentialsResponseInput interface {
	pulumi.Input

	ToCredentialsResponseOutput() CredentialsResponseOutput
	ToCredentialsResponseOutputWithContext(context.Context) CredentialsResponseOutput
}

type CredentialsResponseArgs struct {
	Username pulumi.StringInput `pulumi:"username"`
}

func (CredentialsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CredentialsResponse)(nil)).Elem()
}

func (i CredentialsResponseArgs) ToCredentialsResponseOutput() CredentialsResponseOutput {
	return i.ToCredentialsResponseOutputWithContext(context.Background())
}

func (i CredentialsResponseArgs) ToCredentialsResponseOutputWithContext(ctx context.Context) CredentialsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialsResponseOutput)
}

func (i CredentialsResponseArgs) ToCredentialsResponsePtrOutput() CredentialsResponsePtrOutput {
	return i.ToCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i CredentialsResponseArgs) ToCredentialsResponsePtrOutputWithContext(ctx context.Context) CredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialsResponseOutput).ToCredentialsResponsePtrOutputWithContext(ctx)
}









type CredentialsResponsePtrInput interface {
	pulumi.Input

	ToCredentialsResponsePtrOutput() CredentialsResponsePtrOutput
	ToCredentialsResponsePtrOutputWithContext(context.Context) CredentialsResponsePtrOutput
}

type credentialsResponsePtrType CredentialsResponseArgs

func CredentialsResponsePtr(v *CredentialsResponseArgs) CredentialsResponsePtrInput {
	return (*credentialsResponsePtrType)(v)
}

func (*credentialsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CredentialsResponse)(nil)).Elem()
}

func (i *credentialsResponsePtrType) ToCredentialsResponsePtrOutput() CredentialsResponsePtrOutput {
	return i.ToCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i *credentialsResponsePtrType) ToCredentialsResponsePtrOutputWithContext(ctx context.Context) CredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialsResponsePtrOutput)
}

type CredentialsResponseOutput struct{ *pulumi.OutputState }

func (CredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CredentialsResponse)(nil)).Elem()
}

func (o CredentialsResponseOutput) ToCredentialsResponseOutput() CredentialsResponseOutput {
	return o
}

func (o CredentialsResponseOutput) ToCredentialsResponseOutputWithContext(ctx context.Context) CredentialsResponseOutput {
	return o
}

func (o CredentialsResponseOutput) ToCredentialsResponsePtrOutput() CredentialsResponsePtrOutput {
	return o.ToCredentialsResponsePtrOutputWithContext(context.Background())
}

func (o CredentialsResponseOutput) ToCredentialsResponsePtrOutputWithContext(ctx context.Context) CredentialsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CredentialsResponse) *CredentialsResponse {
		return &v
	}).(CredentialsResponsePtrOutput)
}

func (o CredentialsResponseOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v CredentialsResponse) string { return v.Username }).(pulumi.StringOutput)
}

type CredentialsResponsePtrOutput struct{ *pulumi.OutputState }

func (CredentialsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CredentialsResponse)(nil)).Elem()
}

func (o CredentialsResponsePtrOutput) ToCredentialsResponsePtrOutput() CredentialsResponsePtrOutput {
	return o
}

func (o CredentialsResponsePtrOutput) ToCredentialsResponsePtrOutputWithContext(ctx context.Context) CredentialsResponsePtrOutput {
	return o
}

func (o CredentialsResponsePtrOutput) Elem() CredentialsResponseOutput {
	return o.ApplyT(func(v *CredentialsResponse) CredentialsResponse {
		if v != nil {
			return *v
		}
		var ret CredentialsResponse
		return ret
	}).(CredentialsResponseOutput)
}

func (o CredentialsResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Username
	}).(pulumi.StringPtrOutput)
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
	ExactVersion string  `pulumi:"exactVersion"`
	Id           *string `pulumi:"id"`
	Offer        *string `pulumi:"offer"`
	Publisher    *string `pulumi:"publisher"`
	Sku          *string `pulumi:"sku"`
	Version      *string `pulumi:"version"`
}





type ImageReferenceResponseInput interface {
	pulumi.Input

	ToImageReferenceResponseOutput() ImageReferenceResponseOutput
	ToImageReferenceResponseOutputWithContext(context.Context) ImageReferenceResponseOutput
}

type ImageReferenceResponseArgs struct {
	ExactVersion pulumi.StringInput    `pulumi:"exactVersion"`
	Id           pulumi.StringPtrInput `pulumi:"id"`
	Offer        pulumi.StringPtrInput `pulumi:"offer"`
	Publisher    pulumi.StringPtrInput `pulumi:"publisher"`
	Sku          pulumi.StringPtrInput `pulumi:"sku"`
	Version      pulumi.StringPtrInput `pulumi:"version"`
}

func (ImageReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReferenceResponse)(nil)).Elem()
}

func (i ImageReferenceResponseArgs) ToImageReferenceResponseOutput() ImageReferenceResponseOutput {
	return i.ToImageReferenceResponseOutputWithContext(context.Background())
}

func (i ImageReferenceResponseArgs) ToImageReferenceResponseOutputWithContext(ctx context.Context) ImageReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferenceResponseOutput)
}

func (i ImageReferenceResponseArgs) ToImageReferenceResponsePtrOutput() ImageReferenceResponsePtrOutput {
	return i.ToImageReferenceResponsePtrOutputWithContext(context.Background())
}

func (i ImageReferenceResponseArgs) ToImageReferenceResponsePtrOutputWithContext(ctx context.Context) ImageReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferenceResponseOutput).ToImageReferenceResponsePtrOutputWithContext(ctx)
}









type ImageReferenceResponsePtrInput interface {
	pulumi.Input

	ToImageReferenceResponsePtrOutput() ImageReferenceResponsePtrOutput
	ToImageReferenceResponsePtrOutputWithContext(context.Context) ImageReferenceResponsePtrOutput
}

type imageReferenceResponsePtrType ImageReferenceResponseArgs

func ImageReferenceResponsePtr(v *ImageReferenceResponseArgs) ImageReferenceResponsePtrInput {
	return (*imageReferenceResponsePtrType)(v)
}

func (*imageReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageReferenceResponse)(nil)).Elem()
}

func (i *imageReferenceResponsePtrType) ToImageReferenceResponsePtrOutput() ImageReferenceResponsePtrOutput {
	return i.ToImageReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *imageReferenceResponsePtrType) ToImageReferenceResponsePtrOutputWithContext(ctx context.Context) ImageReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferenceResponsePtrOutput)
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

func (o ImageReferenceResponseOutput) ToImageReferenceResponsePtrOutput() ImageReferenceResponsePtrOutput {
	return o.ToImageReferenceResponsePtrOutputWithContext(context.Background())
}

func (o ImageReferenceResponseOutput) ToImageReferenceResponsePtrOutputWithContext(ctx context.Context) ImageReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageReferenceResponse) *ImageReferenceResponse {
		return &v
	}).(ImageReferenceResponsePtrOutput)
}

func (o ImageReferenceResponseOutput) ExactVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ImageReferenceResponse) string { return v.ExactVersion }).(pulumi.StringOutput)
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

func (o ImageReferenceResponsePtrOutput) ExactVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ExactVersion
	}).(pulumi.StringPtrOutput)
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

type LabNetworkProfile struct {
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	PublicIpId     *string `pulumi:"publicIpId"`
	SubnetId       *string `pulumi:"subnetId"`
}





type LabNetworkProfileInput interface {
	pulumi.Input

	ToLabNetworkProfileOutput() LabNetworkProfileOutput
	ToLabNetworkProfileOutputWithContext(context.Context) LabNetworkProfileOutput
}

type LabNetworkProfileArgs struct {
	LoadBalancerId pulumi.StringPtrInput `pulumi:"loadBalancerId"`
	PublicIpId     pulumi.StringPtrInput `pulumi:"publicIpId"`
	SubnetId       pulumi.StringPtrInput `pulumi:"subnetId"`
}

func (LabNetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LabNetworkProfile)(nil)).Elem()
}

func (i LabNetworkProfileArgs) ToLabNetworkProfileOutput() LabNetworkProfileOutput {
	return i.ToLabNetworkProfileOutputWithContext(context.Background())
}

func (i LabNetworkProfileArgs) ToLabNetworkProfileOutputWithContext(ctx context.Context) LabNetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabNetworkProfileOutput)
}

func (i LabNetworkProfileArgs) ToLabNetworkProfilePtrOutput() LabNetworkProfilePtrOutput {
	return i.ToLabNetworkProfilePtrOutputWithContext(context.Background())
}

func (i LabNetworkProfileArgs) ToLabNetworkProfilePtrOutputWithContext(ctx context.Context) LabNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabNetworkProfileOutput).ToLabNetworkProfilePtrOutputWithContext(ctx)
}









type LabNetworkProfilePtrInput interface {
	pulumi.Input

	ToLabNetworkProfilePtrOutput() LabNetworkProfilePtrOutput
	ToLabNetworkProfilePtrOutputWithContext(context.Context) LabNetworkProfilePtrOutput
}

type labNetworkProfilePtrType LabNetworkProfileArgs

func LabNetworkProfilePtr(v *LabNetworkProfileArgs) LabNetworkProfilePtrInput {
	return (*labNetworkProfilePtrType)(v)
}

func (*labNetworkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LabNetworkProfile)(nil)).Elem()
}

func (i *labNetworkProfilePtrType) ToLabNetworkProfilePtrOutput() LabNetworkProfilePtrOutput {
	return i.ToLabNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *labNetworkProfilePtrType) ToLabNetworkProfilePtrOutputWithContext(ctx context.Context) LabNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabNetworkProfilePtrOutput)
}

type LabNetworkProfileOutput struct{ *pulumi.OutputState }

func (LabNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabNetworkProfile)(nil)).Elem()
}

func (o LabNetworkProfileOutput) ToLabNetworkProfileOutput() LabNetworkProfileOutput {
	return o
}

func (o LabNetworkProfileOutput) ToLabNetworkProfileOutputWithContext(ctx context.Context) LabNetworkProfileOutput {
	return o
}

func (o LabNetworkProfileOutput) ToLabNetworkProfilePtrOutput() LabNetworkProfilePtrOutput {
	return o.ToLabNetworkProfilePtrOutputWithContext(context.Background())
}

func (o LabNetworkProfileOutput) ToLabNetworkProfilePtrOutputWithContext(ctx context.Context) LabNetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LabNetworkProfile) *LabNetworkProfile {
		return &v
	}).(LabNetworkProfilePtrOutput)
}

func (o LabNetworkProfileOutput) LoadBalancerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabNetworkProfile) *string { return v.LoadBalancerId }).(pulumi.StringPtrOutput)
}

func (o LabNetworkProfileOutput) PublicIpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabNetworkProfile) *string { return v.PublicIpId }).(pulumi.StringPtrOutput)
}

func (o LabNetworkProfileOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabNetworkProfile) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

type LabNetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (LabNetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabNetworkProfile)(nil)).Elem()
}

func (o LabNetworkProfilePtrOutput) ToLabNetworkProfilePtrOutput() LabNetworkProfilePtrOutput {
	return o
}

func (o LabNetworkProfilePtrOutput) ToLabNetworkProfilePtrOutputWithContext(ctx context.Context) LabNetworkProfilePtrOutput {
	return o
}

func (o LabNetworkProfilePtrOutput) Elem() LabNetworkProfileOutput {
	return o.ApplyT(func(v *LabNetworkProfile) LabNetworkProfile {
		if v != nil {
			return *v
		}
		var ret LabNetworkProfile
		return ret
	}).(LabNetworkProfileOutput)
}

func (o LabNetworkProfilePtrOutput) LoadBalancerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancerId
	}).(pulumi.StringPtrOutput)
}

func (o LabNetworkProfilePtrOutput) PublicIpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.PublicIpId
	}).(pulumi.StringPtrOutput)
}

func (o LabNetworkProfilePtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.SubnetId
	}).(pulumi.StringPtrOutput)
}

type LabNetworkProfileResponse struct {
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	PublicIpId     *string `pulumi:"publicIpId"`
	SubnetId       *string `pulumi:"subnetId"`
}





type LabNetworkProfileResponseInput interface {
	pulumi.Input

	ToLabNetworkProfileResponseOutput() LabNetworkProfileResponseOutput
	ToLabNetworkProfileResponseOutputWithContext(context.Context) LabNetworkProfileResponseOutput
}

type LabNetworkProfileResponseArgs struct {
	LoadBalancerId pulumi.StringPtrInput `pulumi:"loadBalancerId"`
	PublicIpId     pulumi.StringPtrInput `pulumi:"publicIpId"`
	SubnetId       pulumi.StringPtrInput `pulumi:"subnetId"`
}

func (LabNetworkProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LabNetworkProfileResponse)(nil)).Elem()
}

func (i LabNetworkProfileResponseArgs) ToLabNetworkProfileResponseOutput() LabNetworkProfileResponseOutput {
	return i.ToLabNetworkProfileResponseOutputWithContext(context.Background())
}

func (i LabNetworkProfileResponseArgs) ToLabNetworkProfileResponseOutputWithContext(ctx context.Context) LabNetworkProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabNetworkProfileResponseOutput)
}

func (i LabNetworkProfileResponseArgs) ToLabNetworkProfileResponsePtrOutput() LabNetworkProfileResponsePtrOutput {
	return i.ToLabNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (i LabNetworkProfileResponseArgs) ToLabNetworkProfileResponsePtrOutputWithContext(ctx context.Context) LabNetworkProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabNetworkProfileResponseOutput).ToLabNetworkProfileResponsePtrOutputWithContext(ctx)
}









type LabNetworkProfileResponsePtrInput interface {
	pulumi.Input

	ToLabNetworkProfileResponsePtrOutput() LabNetworkProfileResponsePtrOutput
	ToLabNetworkProfileResponsePtrOutputWithContext(context.Context) LabNetworkProfileResponsePtrOutput
}

type labNetworkProfileResponsePtrType LabNetworkProfileResponseArgs

func LabNetworkProfileResponsePtr(v *LabNetworkProfileResponseArgs) LabNetworkProfileResponsePtrInput {
	return (*labNetworkProfileResponsePtrType)(v)
}

func (*labNetworkProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LabNetworkProfileResponse)(nil)).Elem()
}

func (i *labNetworkProfileResponsePtrType) ToLabNetworkProfileResponsePtrOutput() LabNetworkProfileResponsePtrOutput {
	return i.ToLabNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (i *labNetworkProfileResponsePtrType) ToLabNetworkProfileResponsePtrOutputWithContext(ctx context.Context) LabNetworkProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabNetworkProfileResponsePtrOutput)
}

type LabNetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (LabNetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabNetworkProfileResponse)(nil)).Elem()
}

func (o LabNetworkProfileResponseOutput) ToLabNetworkProfileResponseOutput() LabNetworkProfileResponseOutput {
	return o
}

func (o LabNetworkProfileResponseOutput) ToLabNetworkProfileResponseOutputWithContext(ctx context.Context) LabNetworkProfileResponseOutput {
	return o
}

func (o LabNetworkProfileResponseOutput) ToLabNetworkProfileResponsePtrOutput() LabNetworkProfileResponsePtrOutput {
	return o.ToLabNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (o LabNetworkProfileResponseOutput) ToLabNetworkProfileResponsePtrOutputWithContext(ctx context.Context) LabNetworkProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LabNetworkProfileResponse) *LabNetworkProfileResponse {
		return &v
	}).(LabNetworkProfileResponsePtrOutput)
}

func (o LabNetworkProfileResponseOutput) LoadBalancerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabNetworkProfileResponse) *string { return v.LoadBalancerId }).(pulumi.StringPtrOutput)
}

func (o LabNetworkProfileResponseOutput) PublicIpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabNetworkProfileResponse) *string { return v.PublicIpId }).(pulumi.StringPtrOutput)
}

func (o LabNetworkProfileResponseOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabNetworkProfileResponse) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

type LabNetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (LabNetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabNetworkProfileResponse)(nil)).Elem()
}

func (o LabNetworkProfileResponsePtrOutput) ToLabNetworkProfileResponsePtrOutput() LabNetworkProfileResponsePtrOutput {
	return o
}

func (o LabNetworkProfileResponsePtrOutput) ToLabNetworkProfileResponsePtrOutputWithContext(ctx context.Context) LabNetworkProfileResponsePtrOutput {
	return o
}

func (o LabNetworkProfileResponsePtrOutput) Elem() LabNetworkProfileResponseOutput {
	return o.ApplyT(func(v *LabNetworkProfileResponse) LabNetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret LabNetworkProfileResponse
		return ret
	}).(LabNetworkProfileResponseOutput)
}

func (o LabNetworkProfileResponsePtrOutput) LoadBalancerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancerId
	}).(pulumi.StringPtrOutput)
}

func (o LabNetworkProfileResponsePtrOutput) PublicIpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicIpId
	}).(pulumi.StringPtrOutput)
}

func (o LabNetworkProfileResponsePtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubnetId
	}).(pulumi.StringPtrOutput)
}

type LabPlanNetworkProfile struct {
	SubnetId *string `pulumi:"subnetId"`
}





type LabPlanNetworkProfileInput interface {
	pulumi.Input

	ToLabPlanNetworkProfileOutput() LabPlanNetworkProfileOutput
	ToLabPlanNetworkProfileOutputWithContext(context.Context) LabPlanNetworkProfileOutput
}

type LabPlanNetworkProfileArgs struct {
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
}

func (LabPlanNetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LabPlanNetworkProfile)(nil)).Elem()
}

func (i LabPlanNetworkProfileArgs) ToLabPlanNetworkProfileOutput() LabPlanNetworkProfileOutput {
	return i.ToLabPlanNetworkProfileOutputWithContext(context.Background())
}

func (i LabPlanNetworkProfileArgs) ToLabPlanNetworkProfileOutputWithContext(ctx context.Context) LabPlanNetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabPlanNetworkProfileOutput)
}

func (i LabPlanNetworkProfileArgs) ToLabPlanNetworkProfilePtrOutput() LabPlanNetworkProfilePtrOutput {
	return i.ToLabPlanNetworkProfilePtrOutputWithContext(context.Background())
}

func (i LabPlanNetworkProfileArgs) ToLabPlanNetworkProfilePtrOutputWithContext(ctx context.Context) LabPlanNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabPlanNetworkProfileOutput).ToLabPlanNetworkProfilePtrOutputWithContext(ctx)
}









type LabPlanNetworkProfilePtrInput interface {
	pulumi.Input

	ToLabPlanNetworkProfilePtrOutput() LabPlanNetworkProfilePtrOutput
	ToLabPlanNetworkProfilePtrOutputWithContext(context.Context) LabPlanNetworkProfilePtrOutput
}

type labPlanNetworkProfilePtrType LabPlanNetworkProfileArgs

func LabPlanNetworkProfilePtr(v *LabPlanNetworkProfileArgs) LabPlanNetworkProfilePtrInput {
	return (*labPlanNetworkProfilePtrType)(v)
}

func (*labPlanNetworkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LabPlanNetworkProfile)(nil)).Elem()
}

func (i *labPlanNetworkProfilePtrType) ToLabPlanNetworkProfilePtrOutput() LabPlanNetworkProfilePtrOutput {
	return i.ToLabPlanNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *labPlanNetworkProfilePtrType) ToLabPlanNetworkProfilePtrOutputWithContext(ctx context.Context) LabPlanNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabPlanNetworkProfilePtrOutput)
}

type LabPlanNetworkProfileOutput struct{ *pulumi.OutputState }

func (LabPlanNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabPlanNetworkProfile)(nil)).Elem()
}

func (o LabPlanNetworkProfileOutput) ToLabPlanNetworkProfileOutput() LabPlanNetworkProfileOutput {
	return o
}

func (o LabPlanNetworkProfileOutput) ToLabPlanNetworkProfileOutputWithContext(ctx context.Context) LabPlanNetworkProfileOutput {
	return o
}

func (o LabPlanNetworkProfileOutput) ToLabPlanNetworkProfilePtrOutput() LabPlanNetworkProfilePtrOutput {
	return o.ToLabPlanNetworkProfilePtrOutputWithContext(context.Background())
}

func (o LabPlanNetworkProfileOutput) ToLabPlanNetworkProfilePtrOutputWithContext(ctx context.Context) LabPlanNetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LabPlanNetworkProfile) *LabPlanNetworkProfile {
		return &v
	}).(LabPlanNetworkProfilePtrOutput)
}

func (o LabPlanNetworkProfileOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabPlanNetworkProfile) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

type LabPlanNetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (LabPlanNetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabPlanNetworkProfile)(nil)).Elem()
}

func (o LabPlanNetworkProfilePtrOutput) ToLabPlanNetworkProfilePtrOutput() LabPlanNetworkProfilePtrOutput {
	return o
}

func (o LabPlanNetworkProfilePtrOutput) ToLabPlanNetworkProfilePtrOutputWithContext(ctx context.Context) LabPlanNetworkProfilePtrOutput {
	return o
}

func (o LabPlanNetworkProfilePtrOutput) Elem() LabPlanNetworkProfileOutput {
	return o.ApplyT(func(v *LabPlanNetworkProfile) LabPlanNetworkProfile {
		if v != nil {
			return *v
		}
		var ret LabPlanNetworkProfile
		return ret
	}).(LabPlanNetworkProfileOutput)
}

func (o LabPlanNetworkProfilePtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabPlanNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.SubnetId
	}).(pulumi.StringPtrOutput)
}

type LabPlanNetworkProfileResponse struct {
	SubnetId *string `pulumi:"subnetId"`
}





type LabPlanNetworkProfileResponseInput interface {
	pulumi.Input

	ToLabPlanNetworkProfileResponseOutput() LabPlanNetworkProfileResponseOutput
	ToLabPlanNetworkProfileResponseOutputWithContext(context.Context) LabPlanNetworkProfileResponseOutput
}

type LabPlanNetworkProfileResponseArgs struct {
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
}

func (LabPlanNetworkProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LabPlanNetworkProfileResponse)(nil)).Elem()
}

func (i LabPlanNetworkProfileResponseArgs) ToLabPlanNetworkProfileResponseOutput() LabPlanNetworkProfileResponseOutput {
	return i.ToLabPlanNetworkProfileResponseOutputWithContext(context.Background())
}

func (i LabPlanNetworkProfileResponseArgs) ToLabPlanNetworkProfileResponseOutputWithContext(ctx context.Context) LabPlanNetworkProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabPlanNetworkProfileResponseOutput)
}

func (i LabPlanNetworkProfileResponseArgs) ToLabPlanNetworkProfileResponsePtrOutput() LabPlanNetworkProfileResponsePtrOutput {
	return i.ToLabPlanNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (i LabPlanNetworkProfileResponseArgs) ToLabPlanNetworkProfileResponsePtrOutputWithContext(ctx context.Context) LabPlanNetworkProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabPlanNetworkProfileResponseOutput).ToLabPlanNetworkProfileResponsePtrOutputWithContext(ctx)
}









type LabPlanNetworkProfileResponsePtrInput interface {
	pulumi.Input

	ToLabPlanNetworkProfileResponsePtrOutput() LabPlanNetworkProfileResponsePtrOutput
	ToLabPlanNetworkProfileResponsePtrOutputWithContext(context.Context) LabPlanNetworkProfileResponsePtrOutput
}

type labPlanNetworkProfileResponsePtrType LabPlanNetworkProfileResponseArgs

func LabPlanNetworkProfileResponsePtr(v *LabPlanNetworkProfileResponseArgs) LabPlanNetworkProfileResponsePtrInput {
	return (*labPlanNetworkProfileResponsePtrType)(v)
}

func (*labPlanNetworkProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LabPlanNetworkProfileResponse)(nil)).Elem()
}

func (i *labPlanNetworkProfileResponsePtrType) ToLabPlanNetworkProfileResponsePtrOutput() LabPlanNetworkProfileResponsePtrOutput {
	return i.ToLabPlanNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (i *labPlanNetworkProfileResponsePtrType) ToLabPlanNetworkProfileResponsePtrOutputWithContext(ctx context.Context) LabPlanNetworkProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabPlanNetworkProfileResponsePtrOutput)
}

type LabPlanNetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (LabPlanNetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabPlanNetworkProfileResponse)(nil)).Elem()
}

func (o LabPlanNetworkProfileResponseOutput) ToLabPlanNetworkProfileResponseOutput() LabPlanNetworkProfileResponseOutput {
	return o
}

func (o LabPlanNetworkProfileResponseOutput) ToLabPlanNetworkProfileResponseOutputWithContext(ctx context.Context) LabPlanNetworkProfileResponseOutput {
	return o
}

func (o LabPlanNetworkProfileResponseOutput) ToLabPlanNetworkProfileResponsePtrOutput() LabPlanNetworkProfileResponsePtrOutput {
	return o.ToLabPlanNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (o LabPlanNetworkProfileResponseOutput) ToLabPlanNetworkProfileResponsePtrOutputWithContext(ctx context.Context) LabPlanNetworkProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LabPlanNetworkProfileResponse) *LabPlanNetworkProfileResponse {
		return &v
	}).(LabPlanNetworkProfileResponsePtrOutput)
}

func (o LabPlanNetworkProfileResponseOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabPlanNetworkProfileResponse) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

type LabPlanNetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (LabPlanNetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabPlanNetworkProfileResponse)(nil)).Elem()
}

func (o LabPlanNetworkProfileResponsePtrOutput) ToLabPlanNetworkProfileResponsePtrOutput() LabPlanNetworkProfileResponsePtrOutput {
	return o
}

func (o LabPlanNetworkProfileResponsePtrOutput) ToLabPlanNetworkProfileResponsePtrOutputWithContext(ctx context.Context) LabPlanNetworkProfileResponsePtrOutput {
	return o
}

func (o LabPlanNetworkProfileResponsePtrOutput) Elem() LabPlanNetworkProfileResponseOutput {
	return o.ApplyT(func(v *LabPlanNetworkProfileResponse) LabPlanNetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret LabPlanNetworkProfileResponse
		return ret
	}).(LabPlanNetworkProfileResponseOutput)
}

func (o LabPlanNetworkProfileResponsePtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabPlanNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubnetId
	}).(pulumi.StringPtrOutput)
}

type RecurrencePattern struct {
	ExpirationDate string              `pulumi:"expirationDate"`
	Frequency      RecurrenceFrequency `pulumi:"frequency"`
	Interval       *int                `pulumi:"interval"`
	WeekDays       []WeekDay           `pulumi:"weekDays"`
}





type RecurrencePatternInput interface {
	pulumi.Input

	ToRecurrencePatternOutput() RecurrencePatternOutput
	ToRecurrencePatternOutputWithContext(context.Context) RecurrencePatternOutput
}

type RecurrencePatternArgs struct {
	ExpirationDate pulumi.StringInput       `pulumi:"expirationDate"`
	Frequency      RecurrenceFrequencyInput `pulumi:"frequency"`
	Interval       pulumi.IntPtrInput       `pulumi:"interval"`
	WeekDays       WeekDayArrayInput        `pulumi:"weekDays"`
}

func (RecurrencePatternArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrencePattern)(nil)).Elem()
}

func (i RecurrencePatternArgs) ToRecurrencePatternOutput() RecurrencePatternOutput {
	return i.ToRecurrencePatternOutputWithContext(context.Background())
}

func (i RecurrencePatternArgs) ToRecurrencePatternOutputWithContext(ctx context.Context) RecurrencePatternOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrencePatternOutput)
}

func (i RecurrencePatternArgs) ToRecurrencePatternPtrOutput() RecurrencePatternPtrOutput {
	return i.ToRecurrencePatternPtrOutputWithContext(context.Background())
}

func (i RecurrencePatternArgs) ToRecurrencePatternPtrOutputWithContext(ctx context.Context) RecurrencePatternPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrencePatternOutput).ToRecurrencePatternPtrOutputWithContext(ctx)
}









type RecurrencePatternPtrInput interface {
	pulumi.Input

	ToRecurrencePatternPtrOutput() RecurrencePatternPtrOutput
	ToRecurrencePatternPtrOutputWithContext(context.Context) RecurrencePatternPtrOutput
}

type recurrencePatternPtrType RecurrencePatternArgs

func RecurrencePatternPtr(v *RecurrencePatternArgs) RecurrencePatternPtrInput {
	return (*recurrencePatternPtrType)(v)
}

func (*recurrencePatternPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RecurrencePattern)(nil)).Elem()
}

func (i *recurrencePatternPtrType) ToRecurrencePatternPtrOutput() RecurrencePatternPtrOutput {
	return i.ToRecurrencePatternPtrOutputWithContext(context.Background())
}

func (i *recurrencePatternPtrType) ToRecurrencePatternPtrOutputWithContext(ctx context.Context) RecurrencePatternPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrencePatternPtrOutput)
}

type RecurrencePatternOutput struct{ *pulumi.OutputState }

func (RecurrencePatternOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrencePattern)(nil)).Elem()
}

func (o RecurrencePatternOutput) ToRecurrencePatternOutput() RecurrencePatternOutput {
	return o
}

func (o RecurrencePatternOutput) ToRecurrencePatternOutputWithContext(ctx context.Context) RecurrencePatternOutput {
	return o
}

func (o RecurrencePatternOutput) ToRecurrencePatternPtrOutput() RecurrencePatternPtrOutput {
	return o.ToRecurrencePatternPtrOutputWithContext(context.Background())
}

func (o RecurrencePatternOutput) ToRecurrencePatternPtrOutputWithContext(ctx context.Context) RecurrencePatternPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecurrencePattern) *RecurrencePattern {
		return &v
	}).(RecurrencePatternPtrOutput)
}

func (o RecurrencePatternOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v RecurrencePattern) string { return v.ExpirationDate }).(pulumi.StringOutput)
}

func (o RecurrencePatternOutput) Frequency() RecurrenceFrequencyOutput {
	return o.ApplyT(func(v RecurrencePattern) RecurrenceFrequency { return v.Frequency }).(RecurrenceFrequencyOutput)
}

func (o RecurrencePatternOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RecurrencePattern) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o RecurrencePatternOutput) WeekDays() WeekDayArrayOutput {
	return o.ApplyT(func(v RecurrencePattern) []WeekDay { return v.WeekDays }).(WeekDayArrayOutput)
}

type RecurrencePatternPtrOutput struct{ *pulumi.OutputState }

func (RecurrencePatternPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecurrencePattern)(nil)).Elem()
}

func (o RecurrencePatternPtrOutput) ToRecurrencePatternPtrOutput() RecurrencePatternPtrOutput {
	return o
}

func (o RecurrencePatternPtrOutput) ToRecurrencePatternPtrOutputWithContext(ctx context.Context) RecurrencePatternPtrOutput {
	return o
}

func (o RecurrencePatternPtrOutput) Elem() RecurrencePatternOutput {
	return o.ApplyT(func(v *RecurrencePattern) RecurrencePattern {
		if v != nil {
			return *v
		}
		var ret RecurrencePattern
		return ret
	}).(RecurrencePatternOutput)
}

func (o RecurrencePatternPtrOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecurrencePattern) *string {
		if v == nil {
			return nil
		}
		return &v.ExpirationDate
	}).(pulumi.StringPtrOutput)
}

func (o RecurrencePatternPtrOutput) Frequency() RecurrenceFrequencyPtrOutput {
	return o.ApplyT(func(v *RecurrencePattern) *RecurrenceFrequency {
		if v == nil {
			return nil
		}
		return &v.Frequency
	}).(RecurrenceFrequencyPtrOutput)
}

func (o RecurrencePatternPtrOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RecurrencePattern) *int {
		if v == nil {
			return nil
		}
		return v.Interval
	}).(pulumi.IntPtrOutput)
}

func (o RecurrencePatternPtrOutput) WeekDays() WeekDayArrayOutput {
	return o.ApplyT(func(v *RecurrencePattern) []WeekDay {
		if v == nil {
			return nil
		}
		return v.WeekDays
	}).(WeekDayArrayOutput)
}

type RecurrencePatternResponse struct {
	ExpirationDate string   `pulumi:"expirationDate"`
	Frequency      string   `pulumi:"frequency"`
	Interval       *int     `pulumi:"interval"`
	WeekDays       []string `pulumi:"weekDays"`
}





type RecurrencePatternResponseInput interface {
	pulumi.Input

	ToRecurrencePatternResponseOutput() RecurrencePatternResponseOutput
	ToRecurrencePatternResponseOutputWithContext(context.Context) RecurrencePatternResponseOutput
}

type RecurrencePatternResponseArgs struct {
	ExpirationDate pulumi.StringInput      `pulumi:"expirationDate"`
	Frequency      pulumi.StringInput      `pulumi:"frequency"`
	Interval       pulumi.IntPtrInput      `pulumi:"interval"`
	WeekDays       pulumi.StringArrayInput `pulumi:"weekDays"`
}

func (RecurrencePatternResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrencePatternResponse)(nil)).Elem()
}

func (i RecurrencePatternResponseArgs) ToRecurrencePatternResponseOutput() RecurrencePatternResponseOutput {
	return i.ToRecurrencePatternResponseOutputWithContext(context.Background())
}

func (i RecurrencePatternResponseArgs) ToRecurrencePatternResponseOutputWithContext(ctx context.Context) RecurrencePatternResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrencePatternResponseOutput)
}

func (i RecurrencePatternResponseArgs) ToRecurrencePatternResponsePtrOutput() RecurrencePatternResponsePtrOutput {
	return i.ToRecurrencePatternResponsePtrOutputWithContext(context.Background())
}

func (i RecurrencePatternResponseArgs) ToRecurrencePatternResponsePtrOutputWithContext(ctx context.Context) RecurrencePatternResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrencePatternResponseOutput).ToRecurrencePatternResponsePtrOutputWithContext(ctx)
}









type RecurrencePatternResponsePtrInput interface {
	pulumi.Input

	ToRecurrencePatternResponsePtrOutput() RecurrencePatternResponsePtrOutput
	ToRecurrencePatternResponsePtrOutputWithContext(context.Context) RecurrencePatternResponsePtrOutput
}

type recurrencePatternResponsePtrType RecurrencePatternResponseArgs

func RecurrencePatternResponsePtr(v *RecurrencePatternResponseArgs) RecurrencePatternResponsePtrInput {
	return (*recurrencePatternResponsePtrType)(v)
}

func (*recurrencePatternResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RecurrencePatternResponse)(nil)).Elem()
}

func (i *recurrencePatternResponsePtrType) ToRecurrencePatternResponsePtrOutput() RecurrencePatternResponsePtrOutput {
	return i.ToRecurrencePatternResponsePtrOutputWithContext(context.Background())
}

func (i *recurrencePatternResponsePtrType) ToRecurrencePatternResponsePtrOutputWithContext(ctx context.Context) RecurrencePatternResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrencePatternResponsePtrOutput)
}

type RecurrencePatternResponseOutput struct{ *pulumi.OutputState }

func (RecurrencePatternResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrencePatternResponse)(nil)).Elem()
}

func (o RecurrencePatternResponseOutput) ToRecurrencePatternResponseOutput() RecurrencePatternResponseOutput {
	return o
}

func (o RecurrencePatternResponseOutput) ToRecurrencePatternResponseOutputWithContext(ctx context.Context) RecurrencePatternResponseOutput {
	return o
}

func (o RecurrencePatternResponseOutput) ToRecurrencePatternResponsePtrOutput() RecurrencePatternResponsePtrOutput {
	return o.ToRecurrencePatternResponsePtrOutputWithContext(context.Background())
}

func (o RecurrencePatternResponseOutput) ToRecurrencePatternResponsePtrOutputWithContext(ctx context.Context) RecurrencePatternResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecurrencePatternResponse) *RecurrencePatternResponse {
		return &v
	}).(RecurrencePatternResponsePtrOutput)
}

func (o RecurrencePatternResponseOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v RecurrencePatternResponse) string { return v.ExpirationDate }).(pulumi.StringOutput)
}

func (o RecurrencePatternResponseOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v RecurrencePatternResponse) string { return v.Frequency }).(pulumi.StringOutput)
}

func (o RecurrencePatternResponseOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RecurrencePatternResponse) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o RecurrencePatternResponseOutput) WeekDays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RecurrencePatternResponse) []string { return v.WeekDays }).(pulumi.StringArrayOutput)
}

type RecurrencePatternResponsePtrOutput struct{ *pulumi.OutputState }

func (RecurrencePatternResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecurrencePatternResponse)(nil)).Elem()
}

func (o RecurrencePatternResponsePtrOutput) ToRecurrencePatternResponsePtrOutput() RecurrencePatternResponsePtrOutput {
	return o
}

func (o RecurrencePatternResponsePtrOutput) ToRecurrencePatternResponsePtrOutputWithContext(ctx context.Context) RecurrencePatternResponsePtrOutput {
	return o
}

func (o RecurrencePatternResponsePtrOutput) Elem() RecurrencePatternResponseOutput {
	return o.ApplyT(func(v *RecurrencePatternResponse) RecurrencePatternResponse {
		if v != nil {
			return *v
		}
		var ret RecurrencePatternResponse
		return ret
	}).(RecurrencePatternResponseOutput)
}

func (o RecurrencePatternResponsePtrOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecurrencePatternResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ExpirationDate
	}).(pulumi.StringPtrOutput)
}

func (o RecurrencePatternResponsePtrOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecurrencePatternResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Frequency
	}).(pulumi.StringPtrOutput)
}

func (o RecurrencePatternResponsePtrOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RecurrencePatternResponse) *int {
		if v == nil {
			return nil
		}
		return v.Interval
	}).(pulumi.IntPtrOutput)
}

func (o RecurrencePatternResponsePtrOutput) WeekDays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RecurrencePatternResponse) []string {
		if v == nil {
			return nil
		}
		return v.WeekDays
	}).(pulumi.StringArrayOutput)
}

type RosterProfile struct {
	ActiveDirectoryGroupId *string `pulumi:"activeDirectoryGroupId"`
	LmsInstance            *string `pulumi:"lmsInstance"`
	LtiClientId            *string `pulumi:"ltiClientId"`
	LtiContextId           *string `pulumi:"ltiContextId"`
	LtiRosterEndpoint      *string `pulumi:"ltiRosterEndpoint"`
}





type RosterProfileInput interface {
	pulumi.Input

	ToRosterProfileOutput() RosterProfileOutput
	ToRosterProfileOutputWithContext(context.Context) RosterProfileOutput
}

type RosterProfileArgs struct {
	ActiveDirectoryGroupId pulumi.StringPtrInput `pulumi:"activeDirectoryGroupId"`
	LmsInstance            pulumi.StringPtrInput `pulumi:"lmsInstance"`
	LtiClientId            pulumi.StringPtrInput `pulumi:"ltiClientId"`
	LtiContextId           pulumi.StringPtrInput `pulumi:"ltiContextId"`
	LtiRosterEndpoint      pulumi.StringPtrInput `pulumi:"ltiRosterEndpoint"`
}

func (RosterProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RosterProfile)(nil)).Elem()
}

func (i RosterProfileArgs) ToRosterProfileOutput() RosterProfileOutput {
	return i.ToRosterProfileOutputWithContext(context.Background())
}

func (i RosterProfileArgs) ToRosterProfileOutputWithContext(ctx context.Context) RosterProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RosterProfileOutput)
}

func (i RosterProfileArgs) ToRosterProfilePtrOutput() RosterProfilePtrOutput {
	return i.ToRosterProfilePtrOutputWithContext(context.Background())
}

func (i RosterProfileArgs) ToRosterProfilePtrOutputWithContext(ctx context.Context) RosterProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RosterProfileOutput).ToRosterProfilePtrOutputWithContext(ctx)
}









type RosterProfilePtrInput interface {
	pulumi.Input

	ToRosterProfilePtrOutput() RosterProfilePtrOutput
	ToRosterProfilePtrOutputWithContext(context.Context) RosterProfilePtrOutput
}

type rosterProfilePtrType RosterProfileArgs

func RosterProfilePtr(v *RosterProfileArgs) RosterProfilePtrInput {
	return (*rosterProfilePtrType)(v)
}

func (*rosterProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RosterProfile)(nil)).Elem()
}

func (i *rosterProfilePtrType) ToRosterProfilePtrOutput() RosterProfilePtrOutput {
	return i.ToRosterProfilePtrOutputWithContext(context.Background())
}

func (i *rosterProfilePtrType) ToRosterProfilePtrOutputWithContext(ctx context.Context) RosterProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RosterProfilePtrOutput)
}

type RosterProfileOutput struct{ *pulumi.OutputState }

func (RosterProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RosterProfile)(nil)).Elem()
}

func (o RosterProfileOutput) ToRosterProfileOutput() RosterProfileOutput {
	return o
}

func (o RosterProfileOutput) ToRosterProfileOutputWithContext(ctx context.Context) RosterProfileOutput {
	return o
}

func (o RosterProfileOutput) ToRosterProfilePtrOutput() RosterProfilePtrOutput {
	return o.ToRosterProfilePtrOutputWithContext(context.Background())
}

func (o RosterProfileOutput) ToRosterProfilePtrOutputWithContext(ctx context.Context) RosterProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RosterProfile) *RosterProfile {
		return &v
	}).(RosterProfilePtrOutput)
}

func (o RosterProfileOutput) ActiveDirectoryGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RosterProfile) *string { return v.ActiveDirectoryGroupId }).(pulumi.StringPtrOutput)
}

func (o RosterProfileOutput) LmsInstance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RosterProfile) *string { return v.LmsInstance }).(pulumi.StringPtrOutput)
}

func (o RosterProfileOutput) LtiClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RosterProfile) *string { return v.LtiClientId }).(pulumi.StringPtrOutput)
}

func (o RosterProfileOutput) LtiContextId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RosterProfile) *string { return v.LtiContextId }).(pulumi.StringPtrOutput)
}

func (o RosterProfileOutput) LtiRosterEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RosterProfile) *string { return v.LtiRosterEndpoint }).(pulumi.StringPtrOutput)
}

type RosterProfilePtrOutput struct{ *pulumi.OutputState }

func (RosterProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RosterProfile)(nil)).Elem()
}

func (o RosterProfilePtrOutput) ToRosterProfilePtrOutput() RosterProfilePtrOutput {
	return o
}

func (o RosterProfilePtrOutput) ToRosterProfilePtrOutputWithContext(ctx context.Context) RosterProfilePtrOutput {
	return o
}

func (o RosterProfilePtrOutput) Elem() RosterProfileOutput {
	return o.ApplyT(func(v *RosterProfile) RosterProfile {
		if v != nil {
			return *v
		}
		var ret RosterProfile
		return ret
	}).(RosterProfileOutput)
}

func (o RosterProfilePtrOutput) ActiveDirectoryGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RosterProfile) *string {
		if v == nil {
			return nil
		}
		return v.ActiveDirectoryGroupId
	}).(pulumi.StringPtrOutput)
}

func (o RosterProfilePtrOutput) LmsInstance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RosterProfile) *string {
		if v == nil {
			return nil
		}
		return v.LmsInstance
	}).(pulumi.StringPtrOutput)
}

func (o RosterProfilePtrOutput) LtiClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RosterProfile) *string {
		if v == nil {
			return nil
		}
		return v.LtiClientId
	}).(pulumi.StringPtrOutput)
}

func (o RosterProfilePtrOutput) LtiContextId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RosterProfile) *string {
		if v == nil {
			return nil
		}
		return v.LtiContextId
	}).(pulumi.StringPtrOutput)
}

func (o RosterProfilePtrOutput) LtiRosterEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RosterProfile) *string {
		if v == nil {
			return nil
		}
		return v.LtiRosterEndpoint
	}).(pulumi.StringPtrOutput)
}

type RosterProfileResponse struct {
	ActiveDirectoryGroupId *string `pulumi:"activeDirectoryGroupId"`
	LmsInstance            *string `pulumi:"lmsInstance"`
	LtiClientId            *string `pulumi:"ltiClientId"`
	LtiContextId           *string `pulumi:"ltiContextId"`
	LtiRosterEndpoint      *string `pulumi:"ltiRosterEndpoint"`
}





type RosterProfileResponseInput interface {
	pulumi.Input

	ToRosterProfileResponseOutput() RosterProfileResponseOutput
	ToRosterProfileResponseOutputWithContext(context.Context) RosterProfileResponseOutput
}

type RosterProfileResponseArgs struct {
	ActiveDirectoryGroupId pulumi.StringPtrInput `pulumi:"activeDirectoryGroupId"`
	LmsInstance            pulumi.StringPtrInput `pulumi:"lmsInstance"`
	LtiClientId            pulumi.StringPtrInput `pulumi:"ltiClientId"`
	LtiContextId           pulumi.StringPtrInput `pulumi:"ltiContextId"`
	LtiRosterEndpoint      pulumi.StringPtrInput `pulumi:"ltiRosterEndpoint"`
}

func (RosterProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RosterProfileResponse)(nil)).Elem()
}

func (i RosterProfileResponseArgs) ToRosterProfileResponseOutput() RosterProfileResponseOutput {
	return i.ToRosterProfileResponseOutputWithContext(context.Background())
}

func (i RosterProfileResponseArgs) ToRosterProfileResponseOutputWithContext(ctx context.Context) RosterProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RosterProfileResponseOutput)
}

func (i RosterProfileResponseArgs) ToRosterProfileResponsePtrOutput() RosterProfileResponsePtrOutput {
	return i.ToRosterProfileResponsePtrOutputWithContext(context.Background())
}

func (i RosterProfileResponseArgs) ToRosterProfileResponsePtrOutputWithContext(ctx context.Context) RosterProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RosterProfileResponseOutput).ToRosterProfileResponsePtrOutputWithContext(ctx)
}









type RosterProfileResponsePtrInput interface {
	pulumi.Input

	ToRosterProfileResponsePtrOutput() RosterProfileResponsePtrOutput
	ToRosterProfileResponsePtrOutputWithContext(context.Context) RosterProfileResponsePtrOutput
}

type rosterProfileResponsePtrType RosterProfileResponseArgs

func RosterProfileResponsePtr(v *RosterProfileResponseArgs) RosterProfileResponsePtrInput {
	return (*rosterProfileResponsePtrType)(v)
}

func (*rosterProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RosterProfileResponse)(nil)).Elem()
}

func (i *rosterProfileResponsePtrType) ToRosterProfileResponsePtrOutput() RosterProfileResponsePtrOutput {
	return i.ToRosterProfileResponsePtrOutputWithContext(context.Background())
}

func (i *rosterProfileResponsePtrType) ToRosterProfileResponsePtrOutputWithContext(ctx context.Context) RosterProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RosterProfileResponsePtrOutput)
}

type RosterProfileResponseOutput struct{ *pulumi.OutputState }

func (RosterProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RosterProfileResponse)(nil)).Elem()
}

func (o RosterProfileResponseOutput) ToRosterProfileResponseOutput() RosterProfileResponseOutput {
	return o
}

func (o RosterProfileResponseOutput) ToRosterProfileResponseOutputWithContext(ctx context.Context) RosterProfileResponseOutput {
	return o
}

func (o RosterProfileResponseOutput) ToRosterProfileResponsePtrOutput() RosterProfileResponsePtrOutput {
	return o.ToRosterProfileResponsePtrOutputWithContext(context.Background())
}

func (o RosterProfileResponseOutput) ToRosterProfileResponsePtrOutputWithContext(ctx context.Context) RosterProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RosterProfileResponse) *RosterProfileResponse {
		return &v
	}).(RosterProfileResponsePtrOutput)
}

func (o RosterProfileResponseOutput) ActiveDirectoryGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RosterProfileResponse) *string { return v.ActiveDirectoryGroupId }).(pulumi.StringPtrOutput)
}

func (o RosterProfileResponseOutput) LmsInstance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RosterProfileResponse) *string { return v.LmsInstance }).(pulumi.StringPtrOutput)
}

func (o RosterProfileResponseOutput) LtiClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RosterProfileResponse) *string { return v.LtiClientId }).(pulumi.StringPtrOutput)
}

func (o RosterProfileResponseOutput) LtiContextId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RosterProfileResponse) *string { return v.LtiContextId }).(pulumi.StringPtrOutput)
}

func (o RosterProfileResponseOutput) LtiRosterEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RosterProfileResponse) *string { return v.LtiRosterEndpoint }).(pulumi.StringPtrOutput)
}

type RosterProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (RosterProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RosterProfileResponse)(nil)).Elem()
}

func (o RosterProfileResponsePtrOutput) ToRosterProfileResponsePtrOutput() RosterProfileResponsePtrOutput {
	return o
}

func (o RosterProfileResponsePtrOutput) ToRosterProfileResponsePtrOutputWithContext(ctx context.Context) RosterProfileResponsePtrOutput {
	return o
}

func (o RosterProfileResponsePtrOutput) Elem() RosterProfileResponseOutput {
	return o.ApplyT(func(v *RosterProfileResponse) RosterProfileResponse {
		if v != nil {
			return *v
		}
		var ret RosterProfileResponse
		return ret
	}).(RosterProfileResponseOutput)
}

func (o RosterProfileResponsePtrOutput) ActiveDirectoryGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RosterProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActiveDirectoryGroupId
	}).(pulumi.StringPtrOutput)
}

func (o RosterProfileResponsePtrOutput) LmsInstance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RosterProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.LmsInstance
	}).(pulumi.StringPtrOutput)
}

func (o RosterProfileResponsePtrOutput) LtiClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RosterProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.LtiClientId
	}).(pulumi.StringPtrOutput)
}

func (o RosterProfileResponsePtrOutput) LtiContextId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RosterProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.LtiContextId
	}).(pulumi.StringPtrOutput)
}

func (o RosterProfileResponsePtrOutput) LtiRosterEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RosterProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.LtiRosterEndpoint
	}).(pulumi.StringPtrOutput)
}

type SecurityProfile struct {
	OpenAccess *EnableState `pulumi:"openAccess"`
}





type SecurityProfileInput interface {
	pulumi.Input

	ToSecurityProfileOutput() SecurityProfileOutput
	ToSecurityProfileOutputWithContext(context.Context) SecurityProfileOutput
}

type SecurityProfileArgs struct {
	OpenAccess EnableStatePtrInput `pulumi:"openAccess"`
}

func (SecurityProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityProfile)(nil)).Elem()
}

func (i SecurityProfileArgs) ToSecurityProfileOutput() SecurityProfileOutput {
	return i.ToSecurityProfileOutputWithContext(context.Background())
}

func (i SecurityProfileArgs) ToSecurityProfileOutputWithContext(ctx context.Context) SecurityProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityProfileOutput)
}

func (i SecurityProfileArgs) ToSecurityProfilePtrOutput() SecurityProfilePtrOutput {
	return i.ToSecurityProfilePtrOutputWithContext(context.Background())
}

func (i SecurityProfileArgs) ToSecurityProfilePtrOutputWithContext(ctx context.Context) SecurityProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityProfileOutput).ToSecurityProfilePtrOutputWithContext(ctx)
}









type SecurityProfilePtrInput interface {
	pulumi.Input

	ToSecurityProfilePtrOutput() SecurityProfilePtrOutput
	ToSecurityProfilePtrOutputWithContext(context.Context) SecurityProfilePtrOutput
}

type securityProfilePtrType SecurityProfileArgs

func SecurityProfilePtr(v *SecurityProfileArgs) SecurityProfilePtrInput {
	return (*securityProfilePtrType)(v)
}

func (*securityProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityProfile)(nil)).Elem()
}

func (i *securityProfilePtrType) ToSecurityProfilePtrOutput() SecurityProfilePtrOutput {
	return i.ToSecurityProfilePtrOutputWithContext(context.Background())
}

func (i *securityProfilePtrType) ToSecurityProfilePtrOutputWithContext(ctx context.Context) SecurityProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityProfilePtrOutput)
}

type SecurityProfileOutput struct{ *pulumi.OutputState }

func (SecurityProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityProfile)(nil)).Elem()
}

func (o SecurityProfileOutput) ToSecurityProfileOutput() SecurityProfileOutput {
	return o
}

func (o SecurityProfileOutput) ToSecurityProfileOutputWithContext(ctx context.Context) SecurityProfileOutput {
	return o
}

func (o SecurityProfileOutput) ToSecurityProfilePtrOutput() SecurityProfilePtrOutput {
	return o.ToSecurityProfilePtrOutputWithContext(context.Background())
}

func (o SecurityProfileOutput) ToSecurityProfilePtrOutputWithContext(ctx context.Context) SecurityProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityProfile) *SecurityProfile {
		return &v
	}).(SecurityProfilePtrOutput)
}

func (o SecurityProfileOutput) OpenAccess() EnableStatePtrOutput {
	return o.ApplyT(func(v SecurityProfile) *EnableState { return v.OpenAccess }).(EnableStatePtrOutput)
}

type SecurityProfilePtrOutput struct{ *pulumi.OutputState }

func (SecurityProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityProfile)(nil)).Elem()
}

func (o SecurityProfilePtrOutput) ToSecurityProfilePtrOutput() SecurityProfilePtrOutput {
	return o
}

func (o SecurityProfilePtrOutput) ToSecurityProfilePtrOutputWithContext(ctx context.Context) SecurityProfilePtrOutput {
	return o
}

func (o SecurityProfilePtrOutput) Elem() SecurityProfileOutput {
	return o.ApplyT(func(v *SecurityProfile) SecurityProfile {
		if v != nil {
			return *v
		}
		var ret SecurityProfile
		return ret
	}).(SecurityProfileOutput)
}

func (o SecurityProfilePtrOutput) OpenAccess() EnableStatePtrOutput {
	return o.ApplyT(func(v *SecurityProfile) *EnableState {
		if v == nil {
			return nil
		}
		return v.OpenAccess
	}).(EnableStatePtrOutput)
}

type SecurityProfileResponse struct {
	OpenAccess       *string `pulumi:"openAccess"`
	RegistrationCode string  `pulumi:"registrationCode"`
}





type SecurityProfileResponseInput interface {
	pulumi.Input

	ToSecurityProfileResponseOutput() SecurityProfileResponseOutput
	ToSecurityProfileResponseOutputWithContext(context.Context) SecurityProfileResponseOutput
}

type SecurityProfileResponseArgs struct {
	OpenAccess       pulumi.StringPtrInput `pulumi:"openAccess"`
	RegistrationCode pulumi.StringInput    `pulumi:"registrationCode"`
}

func (SecurityProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityProfileResponse)(nil)).Elem()
}

func (i SecurityProfileResponseArgs) ToSecurityProfileResponseOutput() SecurityProfileResponseOutput {
	return i.ToSecurityProfileResponseOutputWithContext(context.Background())
}

func (i SecurityProfileResponseArgs) ToSecurityProfileResponseOutputWithContext(ctx context.Context) SecurityProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityProfileResponseOutput)
}

func (i SecurityProfileResponseArgs) ToSecurityProfileResponsePtrOutput() SecurityProfileResponsePtrOutput {
	return i.ToSecurityProfileResponsePtrOutputWithContext(context.Background())
}

func (i SecurityProfileResponseArgs) ToSecurityProfileResponsePtrOutputWithContext(ctx context.Context) SecurityProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityProfileResponseOutput).ToSecurityProfileResponsePtrOutputWithContext(ctx)
}









type SecurityProfileResponsePtrInput interface {
	pulumi.Input

	ToSecurityProfileResponsePtrOutput() SecurityProfileResponsePtrOutput
	ToSecurityProfileResponsePtrOutputWithContext(context.Context) SecurityProfileResponsePtrOutput
}

type securityProfileResponsePtrType SecurityProfileResponseArgs

func SecurityProfileResponsePtr(v *SecurityProfileResponseArgs) SecurityProfileResponsePtrInput {
	return (*securityProfileResponsePtrType)(v)
}

func (*securityProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityProfileResponse)(nil)).Elem()
}

func (i *securityProfileResponsePtrType) ToSecurityProfileResponsePtrOutput() SecurityProfileResponsePtrOutput {
	return i.ToSecurityProfileResponsePtrOutputWithContext(context.Background())
}

func (i *securityProfileResponsePtrType) ToSecurityProfileResponsePtrOutputWithContext(ctx context.Context) SecurityProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityProfileResponsePtrOutput)
}

type SecurityProfileResponseOutput struct{ *pulumi.OutputState }

func (SecurityProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityProfileResponse)(nil)).Elem()
}

func (o SecurityProfileResponseOutput) ToSecurityProfileResponseOutput() SecurityProfileResponseOutput {
	return o
}

func (o SecurityProfileResponseOutput) ToSecurityProfileResponseOutputWithContext(ctx context.Context) SecurityProfileResponseOutput {
	return o
}

func (o SecurityProfileResponseOutput) ToSecurityProfileResponsePtrOutput() SecurityProfileResponsePtrOutput {
	return o.ToSecurityProfileResponsePtrOutputWithContext(context.Background())
}

func (o SecurityProfileResponseOutput) ToSecurityProfileResponsePtrOutputWithContext(ctx context.Context) SecurityProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityProfileResponse) *SecurityProfileResponse {
		return &v
	}).(SecurityProfileResponsePtrOutput)
}

func (o SecurityProfileResponseOutput) OpenAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityProfileResponse) *string { return v.OpenAccess }).(pulumi.StringPtrOutput)
}

func (o SecurityProfileResponseOutput) RegistrationCode() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityProfileResponse) string { return v.RegistrationCode }).(pulumi.StringOutput)
}

type SecurityProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (SecurityProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityProfileResponse)(nil)).Elem()
}

func (o SecurityProfileResponsePtrOutput) ToSecurityProfileResponsePtrOutput() SecurityProfileResponsePtrOutput {
	return o
}

func (o SecurityProfileResponsePtrOutput) ToSecurityProfileResponsePtrOutputWithContext(ctx context.Context) SecurityProfileResponsePtrOutput {
	return o
}

func (o SecurityProfileResponsePtrOutput) Elem() SecurityProfileResponseOutput {
	return o.ApplyT(func(v *SecurityProfileResponse) SecurityProfileResponse {
		if v != nil {
			return *v
		}
		var ret SecurityProfileResponse
		return ret
	}).(SecurityProfileResponseOutput)
}

func (o SecurityProfileResponsePtrOutput) OpenAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.OpenAccess
	}).(pulumi.StringPtrOutput)
}

func (o SecurityProfileResponsePtrOutput) RegistrationCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RegistrationCode
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Capacity *int     `pulumi:"capacity"`
	Family   *string  `pulumi:"family"`
	Name     string   `pulumi:"name"`
	Size     *string  `pulumi:"size"`
	Tier     *SkuTier `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     SkuTierPtrInput       `pulumi:"tier"`
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

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() SkuTierPtrOutput {
	return o.ApplyT(func(v Sku) *SkuTier { return v.Tier }).(SkuTierPtrOutput)
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

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() SkuTierPtrOutput {
	return o.ApplyT(func(v *Sku) *SkuTier {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(SkuTierPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
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

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
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

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SupportInfo struct {
	Email        *string `pulumi:"email"`
	Instructions *string `pulumi:"instructions"`
	Phone        *string `pulumi:"phone"`
	Url          *string `pulumi:"url"`
}





type SupportInfoInput interface {
	pulumi.Input

	ToSupportInfoOutput() SupportInfoOutput
	ToSupportInfoOutputWithContext(context.Context) SupportInfoOutput
}

type SupportInfoArgs struct {
	Email        pulumi.StringPtrInput `pulumi:"email"`
	Instructions pulumi.StringPtrInput `pulumi:"instructions"`
	Phone        pulumi.StringPtrInput `pulumi:"phone"`
	Url          pulumi.StringPtrInput `pulumi:"url"`
}

func (SupportInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportInfo)(nil)).Elem()
}

func (i SupportInfoArgs) ToSupportInfoOutput() SupportInfoOutput {
	return i.ToSupportInfoOutputWithContext(context.Background())
}

func (i SupportInfoArgs) ToSupportInfoOutputWithContext(ctx context.Context) SupportInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SupportInfoOutput)
}

func (i SupportInfoArgs) ToSupportInfoPtrOutput() SupportInfoPtrOutput {
	return i.ToSupportInfoPtrOutputWithContext(context.Background())
}

func (i SupportInfoArgs) ToSupportInfoPtrOutputWithContext(ctx context.Context) SupportInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SupportInfoOutput).ToSupportInfoPtrOutputWithContext(ctx)
}









type SupportInfoPtrInput interface {
	pulumi.Input

	ToSupportInfoPtrOutput() SupportInfoPtrOutput
	ToSupportInfoPtrOutputWithContext(context.Context) SupportInfoPtrOutput
}

type supportInfoPtrType SupportInfoArgs

func SupportInfoPtr(v *SupportInfoArgs) SupportInfoPtrInput {
	return (*supportInfoPtrType)(v)
}

func (*supportInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SupportInfo)(nil)).Elem()
}

func (i *supportInfoPtrType) ToSupportInfoPtrOutput() SupportInfoPtrOutput {
	return i.ToSupportInfoPtrOutputWithContext(context.Background())
}

func (i *supportInfoPtrType) ToSupportInfoPtrOutputWithContext(ctx context.Context) SupportInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SupportInfoPtrOutput)
}

type SupportInfoOutput struct{ *pulumi.OutputState }

func (SupportInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportInfo)(nil)).Elem()
}

func (o SupportInfoOutput) ToSupportInfoOutput() SupportInfoOutput {
	return o
}

func (o SupportInfoOutput) ToSupportInfoOutputWithContext(ctx context.Context) SupportInfoOutput {
	return o
}

func (o SupportInfoOutput) ToSupportInfoPtrOutput() SupportInfoPtrOutput {
	return o.ToSupportInfoPtrOutputWithContext(context.Background())
}

func (o SupportInfoOutput) ToSupportInfoPtrOutputWithContext(ctx context.Context) SupportInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SupportInfo) *SupportInfo {
		return &v
	}).(SupportInfoPtrOutput)
}

func (o SupportInfoOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SupportInfo) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o SupportInfoOutput) Instructions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SupportInfo) *string { return v.Instructions }).(pulumi.StringPtrOutput)
}

func (o SupportInfoOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SupportInfo) *string { return v.Phone }).(pulumi.StringPtrOutput)
}

func (o SupportInfoOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SupportInfo) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type SupportInfoPtrOutput struct{ *pulumi.OutputState }

func (SupportInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SupportInfo)(nil)).Elem()
}

func (o SupportInfoPtrOutput) ToSupportInfoPtrOutput() SupportInfoPtrOutput {
	return o
}

func (o SupportInfoPtrOutput) ToSupportInfoPtrOutputWithContext(ctx context.Context) SupportInfoPtrOutput {
	return o
}

func (o SupportInfoPtrOutput) Elem() SupportInfoOutput {
	return o.ApplyT(func(v *SupportInfo) SupportInfo {
		if v != nil {
			return *v
		}
		var ret SupportInfo
		return ret
	}).(SupportInfoOutput)
}

func (o SupportInfoPtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SupportInfo) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o SupportInfoPtrOutput) Instructions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SupportInfo) *string {
		if v == nil {
			return nil
		}
		return v.Instructions
	}).(pulumi.StringPtrOutput)
}

func (o SupportInfoPtrOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SupportInfo) *string {
		if v == nil {
			return nil
		}
		return v.Phone
	}).(pulumi.StringPtrOutput)
}

func (o SupportInfoPtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SupportInfo) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type SupportInfoResponse struct {
	Email        *string `pulumi:"email"`
	Instructions *string `pulumi:"instructions"`
	Phone        *string `pulumi:"phone"`
	Url          *string `pulumi:"url"`
}





type SupportInfoResponseInput interface {
	pulumi.Input

	ToSupportInfoResponseOutput() SupportInfoResponseOutput
	ToSupportInfoResponseOutputWithContext(context.Context) SupportInfoResponseOutput
}

type SupportInfoResponseArgs struct {
	Email        pulumi.StringPtrInput `pulumi:"email"`
	Instructions pulumi.StringPtrInput `pulumi:"instructions"`
	Phone        pulumi.StringPtrInput `pulumi:"phone"`
	Url          pulumi.StringPtrInput `pulumi:"url"`
}

func (SupportInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportInfoResponse)(nil)).Elem()
}

func (i SupportInfoResponseArgs) ToSupportInfoResponseOutput() SupportInfoResponseOutput {
	return i.ToSupportInfoResponseOutputWithContext(context.Background())
}

func (i SupportInfoResponseArgs) ToSupportInfoResponseOutputWithContext(ctx context.Context) SupportInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SupportInfoResponseOutput)
}

func (i SupportInfoResponseArgs) ToSupportInfoResponsePtrOutput() SupportInfoResponsePtrOutput {
	return i.ToSupportInfoResponsePtrOutputWithContext(context.Background())
}

func (i SupportInfoResponseArgs) ToSupportInfoResponsePtrOutputWithContext(ctx context.Context) SupportInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SupportInfoResponseOutput).ToSupportInfoResponsePtrOutputWithContext(ctx)
}









type SupportInfoResponsePtrInput interface {
	pulumi.Input

	ToSupportInfoResponsePtrOutput() SupportInfoResponsePtrOutput
	ToSupportInfoResponsePtrOutputWithContext(context.Context) SupportInfoResponsePtrOutput
}

type supportInfoResponsePtrType SupportInfoResponseArgs

func SupportInfoResponsePtr(v *SupportInfoResponseArgs) SupportInfoResponsePtrInput {
	return (*supportInfoResponsePtrType)(v)
}

func (*supportInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SupportInfoResponse)(nil)).Elem()
}

func (i *supportInfoResponsePtrType) ToSupportInfoResponsePtrOutput() SupportInfoResponsePtrOutput {
	return i.ToSupportInfoResponsePtrOutputWithContext(context.Background())
}

func (i *supportInfoResponsePtrType) ToSupportInfoResponsePtrOutputWithContext(ctx context.Context) SupportInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SupportInfoResponsePtrOutput)
}

type SupportInfoResponseOutput struct{ *pulumi.OutputState }

func (SupportInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportInfoResponse)(nil)).Elem()
}

func (o SupportInfoResponseOutput) ToSupportInfoResponseOutput() SupportInfoResponseOutput {
	return o
}

func (o SupportInfoResponseOutput) ToSupportInfoResponseOutputWithContext(ctx context.Context) SupportInfoResponseOutput {
	return o
}

func (o SupportInfoResponseOutput) ToSupportInfoResponsePtrOutput() SupportInfoResponsePtrOutput {
	return o.ToSupportInfoResponsePtrOutputWithContext(context.Background())
}

func (o SupportInfoResponseOutput) ToSupportInfoResponsePtrOutputWithContext(ctx context.Context) SupportInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SupportInfoResponse) *SupportInfoResponse {
		return &v
	}).(SupportInfoResponsePtrOutput)
}

func (o SupportInfoResponseOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SupportInfoResponse) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o SupportInfoResponseOutput) Instructions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SupportInfoResponse) *string { return v.Instructions }).(pulumi.StringPtrOutput)
}

func (o SupportInfoResponseOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SupportInfoResponse) *string { return v.Phone }).(pulumi.StringPtrOutput)
}

func (o SupportInfoResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SupportInfoResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type SupportInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (SupportInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SupportInfoResponse)(nil)).Elem()
}

func (o SupportInfoResponsePtrOutput) ToSupportInfoResponsePtrOutput() SupportInfoResponsePtrOutput {
	return o
}

func (o SupportInfoResponsePtrOutput) ToSupportInfoResponsePtrOutputWithContext(ctx context.Context) SupportInfoResponsePtrOutput {
	return o
}

func (o SupportInfoResponsePtrOutput) Elem() SupportInfoResponseOutput {
	return o.ApplyT(func(v *SupportInfoResponse) SupportInfoResponse {
		if v != nil {
			return *v
		}
		var ret SupportInfoResponse
		return ret
	}).(SupportInfoResponseOutput)
}

func (o SupportInfoResponsePtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SupportInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o SupportInfoResponsePtrOutput) Instructions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SupportInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Instructions
	}).(pulumi.StringPtrOutput)
}

func (o SupportInfoResponsePtrOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SupportInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Phone
	}).(pulumi.StringPtrOutput)
}

func (o SupportInfoResponsePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SupportInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Url
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

type VirtualMachineAdditionalCapabilities struct {
	InstallGpuDrivers *EnableState `pulumi:"installGpuDrivers"`
}





type VirtualMachineAdditionalCapabilitiesInput interface {
	pulumi.Input

	ToVirtualMachineAdditionalCapabilitiesOutput() VirtualMachineAdditionalCapabilitiesOutput
	ToVirtualMachineAdditionalCapabilitiesOutputWithContext(context.Context) VirtualMachineAdditionalCapabilitiesOutput
}

type VirtualMachineAdditionalCapabilitiesArgs struct {
	InstallGpuDrivers EnableStatePtrInput `pulumi:"installGpuDrivers"`
}

func (VirtualMachineAdditionalCapabilitiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineAdditionalCapabilities)(nil)).Elem()
}

func (i VirtualMachineAdditionalCapabilitiesArgs) ToVirtualMachineAdditionalCapabilitiesOutput() VirtualMachineAdditionalCapabilitiesOutput {
	return i.ToVirtualMachineAdditionalCapabilitiesOutputWithContext(context.Background())
}

func (i VirtualMachineAdditionalCapabilitiesArgs) ToVirtualMachineAdditionalCapabilitiesOutputWithContext(ctx context.Context) VirtualMachineAdditionalCapabilitiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineAdditionalCapabilitiesOutput)
}

func (i VirtualMachineAdditionalCapabilitiesArgs) ToVirtualMachineAdditionalCapabilitiesPtrOutput() VirtualMachineAdditionalCapabilitiesPtrOutput {
	return i.ToVirtualMachineAdditionalCapabilitiesPtrOutputWithContext(context.Background())
}

func (i VirtualMachineAdditionalCapabilitiesArgs) ToVirtualMachineAdditionalCapabilitiesPtrOutputWithContext(ctx context.Context) VirtualMachineAdditionalCapabilitiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineAdditionalCapabilitiesOutput).ToVirtualMachineAdditionalCapabilitiesPtrOutputWithContext(ctx)
}









type VirtualMachineAdditionalCapabilitiesPtrInput interface {
	pulumi.Input

	ToVirtualMachineAdditionalCapabilitiesPtrOutput() VirtualMachineAdditionalCapabilitiesPtrOutput
	ToVirtualMachineAdditionalCapabilitiesPtrOutputWithContext(context.Context) VirtualMachineAdditionalCapabilitiesPtrOutput
}

type virtualMachineAdditionalCapabilitiesPtrType VirtualMachineAdditionalCapabilitiesArgs

func VirtualMachineAdditionalCapabilitiesPtr(v *VirtualMachineAdditionalCapabilitiesArgs) VirtualMachineAdditionalCapabilitiesPtrInput {
	return (*virtualMachineAdditionalCapabilitiesPtrType)(v)
}

func (*virtualMachineAdditionalCapabilitiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineAdditionalCapabilities)(nil)).Elem()
}

func (i *virtualMachineAdditionalCapabilitiesPtrType) ToVirtualMachineAdditionalCapabilitiesPtrOutput() VirtualMachineAdditionalCapabilitiesPtrOutput {
	return i.ToVirtualMachineAdditionalCapabilitiesPtrOutputWithContext(context.Background())
}

func (i *virtualMachineAdditionalCapabilitiesPtrType) ToVirtualMachineAdditionalCapabilitiesPtrOutputWithContext(ctx context.Context) VirtualMachineAdditionalCapabilitiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineAdditionalCapabilitiesPtrOutput)
}

type VirtualMachineAdditionalCapabilitiesOutput struct{ *pulumi.OutputState }

func (VirtualMachineAdditionalCapabilitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineAdditionalCapabilities)(nil)).Elem()
}

func (o VirtualMachineAdditionalCapabilitiesOutput) ToVirtualMachineAdditionalCapabilitiesOutput() VirtualMachineAdditionalCapabilitiesOutput {
	return o
}

func (o VirtualMachineAdditionalCapabilitiesOutput) ToVirtualMachineAdditionalCapabilitiesOutputWithContext(ctx context.Context) VirtualMachineAdditionalCapabilitiesOutput {
	return o
}

func (o VirtualMachineAdditionalCapabilitiesOutput) ToVirtualMachineAdditionalCapabilitiesPtrOutput() VirtualMachineAdditionalCapabilitiesPtrOutput {
	return o.ToVirtualMachineAdditionalCapabilitiesPtrOutputWithContext(context.Background())
}

func (o VirtualMachineAdditionalCapabilitiesOutput) ToVirtualMachineAdditionalCapabilitiesPtrOutputWithContext(ctx context.Context) VirtualMachineAdditionalCapabilitiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineAdditionalCapabilities) *VirtualMachineAdditionalCapabilities {
		return &v
	}).(VirtualMachineAdditionalCapabilitiesPtrOutput)
}

func (o VirtualMachineAdditionalCapabilitiesOutput) InstallGpuDrivers() EnableStatePtrOutput {
	return o.ApplyT(func(v VirtualMachineAdditionalCapabilities) *EnableState { return v.InstallGpuDrivers }).(EnableStatePtrOutput)
}

type VirtualMachineAdditionalCapabilitiesPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineAdditionalCapabilitiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineAdditionalCapabilities)(nil)).Elem()
}

func (o VirtualMachineAdditionalCapabilitiesPtrOutput) ToVirtualMachineAdditionalCapabilitiesPtrOutput() VirtualMachineAdditionalCapabilitiesPtrOutput {
	return o
}

func (o VirtualMachineAdditionalCapabilitiesPtrOutput) ToVirtualMachineAdditionalCapabilitiesPtrOutputWithContext(ctx context.Context) VirtualMachineAdditionalCapabilitiesPtrOutput {
	return o
}

func (o VirtualMachineAdditionalCapabilitiesPtrOutput) Elem() VirtualMachineAdditionalCapabilitiesOutput {
	return o.ApplyT(func(v *VirtualMachineAdditionalCapabilities) VirtualMachineAdditionalCapabilities {
		if v != nil {
			return *v
		}
		var ret VirtualMachineAdditionalCapabilities
		return ret
	}).(VirtualMachineAdditionalCapabilitiesOutput)
}

func (o VirtualMachineAdditionalCapabilitiesPtrOutput) InstallGpuDrivers() EnableStatePtrOutput {
	return o.ApplyT(func(v *VirtualMachineAdditionalCapabilities) *EnableState {
		if v == nil {
			return nil
		}
		return v.InstallGpuDrivers
	}).(EnableStatePtrOutput)
}

type VirtualMachineAdditionalCapabilitiesResponse struct {
	InstallGpuDrivers *string `pulumi:"installGpuDrivers"`
}





type VirtualMachineAdditionalCapabilitiesResponseInput interface {
	pulumi.Input

	ToVirtualMachineAdditionalCapabilitiesResponseOutput() VirtualMachineAdditionalCapabilitiesResponseOutput
	ToVirtualMachineAdditionalCapabilitiesResponseOutputWithContext(context.Context) VirtualMachineAdditionalCapabilitiesResponseOutput
}

type VirtualMachineAdditionalCapabilitiesResponseArgs struct {
	InstallGpuDrivers pulumi.StringPtrInput `pulumi:"installGpuDrivers"`
}

func (VirtualMachineAdditionalCapabilitiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineAdditionalCapabilitiesResponse)(nil)).Elem()
}

func (i VirtualMachineAdditionalCapabilitiesResponseArgs) ToVirtualMachineAdditionalCapabilitiesResponseOutput() VirtualMachineAdditionalCapabilitiesResponseOutput {
	return i.ToVirtualMachineAdditionalCapabilitiesResponseOutputWithContext(context.Background())
}

func (i VirtualMachineAdditionalCapabilitiesResponseArgs) ToVirtualMachineAdditionalCapabilitiesResponseOutputWithContext(ctx context.Context) VirtualMachineAdditionalCapabilitiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineAdditionalCapabilitiesResponseOutput)
}

func (i VirtualMachineAdditionalCapabilitiesResponseArgs) ToVirtualMachineAdditionalCapabilitiesResponsePtrOutput() VirtualMachineAdditionalCapabilitiesResponsePtrOutput {
	return i.ToVirtualMachineAdditionalCapabilitiesResponsePtrOutputWithContext(context.Background())
}

func (i VirtualMachineAdditionalCapabilitiesResponseArgs) ToVirtualMachineAdditionalCapabilitiesResponsePtrOutputWithContext(ctx context.Context) VirtualMachineAdditionalCapabilitiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineAdditionalCapabilitiesResponseOutput).ToVirtualMachineAdditionalCapabilitiesResponsePtrOutputWithContext(ctx)
}









type VirtualMachineAdditionalCapabilitiesResponsePtrInput interface {
	pulumi.Input

	ToVirtualMachineAdditionalCapabilitiesResponsePtrOutput() VirtualMachineAdditionalCapabilitiesResponsePtrOutput
	ToVirtualMachineAdditionalCapabilitiesResponsePtrOutputWithContext(context.Context) VirtualMachineAdditionalCapabilitiesResponsePtrOutput
}

type virtualMachineAdditionalCapabilitiesResponsePtrType VirtualMachineAdditionalCapabilitiesResponseArgs

func VirtualMachineAdditionalCapabilitiesResponsePtr(v *VirtualMachineAdditionalCapabilitiesResponseArgs) VirtualMachineAdditionalCapabilitiesResponsePtrInput {
	return (*virtualMachineAdditionalCapabilitiesResponsePtrType)(v)
}

func (*virtualMachineAdditionalCapabilitiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineAdditionalCapabilitiesResponse)(nil)).Elem()
}

func (i *virtualMachineAdditionalCapabilitiesResponsePtrType) ToVirtualMachineAdditionalCapabilitiesResponsePtrOutput() VirtualMachineAdditionalCapabilitiesResponsePtrOutput {
	return i.ToVirtualMachineAdditionalCapabilitiesResponsePtrOutputWithContext(context.Background())
}

func (i *virtualMachineAdditionalCapabilitiesResponsePtrType) ToVirtualMachineAdditionalCapabilitiesResponsePtrOutputWithContext(ctx context.Context) VirtualMachineAdditionalCapabilitiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineAdditionalCapabilitiesResponsePtrOutput)
}

type VirtualMachineAdditionalCapabilitiesResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineAdditionalCapabilitiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineAdditionalCapabilitiesResponse)(nil)).Elem()
}

func (o VirtualMachineAdditionalCapabilitiesResponseOutput) ToVirtualMachineAdditionalCapabilitiesResponseOutput() VirtualMachineAdditionalCapabilitiesResponseOutput {
	return o
}

func (o VirtualMachineAdditionalCapabilitiesResponseOutput) ToVirtualMachineAdditionalCapabilitiesResponseOutputWithContext(ctx context.Context) VirtualMachineAdditionalCapabilitiesResponseOutput {
	return o
}

func (o VirtualMachineAdditionalCapabilitiesResponseOutput) ToVirtualMachineAdditionalCapabilitiesResponsePtrOutput() VirtualMachineAdditionalCapabilitiesResponsePtrOutput {
	return o.ToVirtualMachineAdditionalCapabilitiesResponsePtrOutputWithContext(context.Background())
}

func (o VirtualMachineAdditionalCapabilitiesResponseOutput) ToVirtualMachineAdditionalCapabilitiesResponsePtrOutputWithContext(ctx context.Context) VirtualMachineAdditionalCapabilitiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineAdditionalCapabilitiesResponse) *VirtualMachineAdditionalCapabilitiesResponse {
		return &v
	}).(VirtualMachineAdditionalCapabilitiesResponsePtrOutput)
}

func (o VirtualMachineAdditionalCapabilitiesResponseOutput) InstallGpuDrivers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineAdditionalCapabilitiesResponse) *string { return v.InstallGpuDrivers }).(pulumi.StringPtrOutput)
}

type VirtualMachineAdditionalCapabilitiesResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineAdditionalCapabilitiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineAdditionalCapabilitiesResponse)(nil)).Elem()
}

func (o VirtualMachineAdditionalCapabilitiesResponsePtrOutput) ToVirtualMachineAdditionalCapabilitiesResponsePtrOutput() VirtualMachineAdditionalCapabilitiesResponsePtrOutput {
	return o
}

func (o VirtualMachineAdditionalCapabilitiesResponsePtrOutput) ToVirtualMachineAdditionalCapabilitiesResponsePtrOutputWithContext(ctx context.Context) VirtualMachineAdditionalCapabilitiesResponsePtrOutput {
	return o
}

func (o VirtualMachineAdditionalCapabilitiesResponsePtrOutput) Elem() VirtualMachineAdditionalCapabilitiesResponseOutput {
	return o.ApplyT(func(v *VirtualMachineAdditionalCapabilitiesResponse) VirtualMachineAdditionalCapabilitiesResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineAdditionalCapabilitiesResponse
		return ret
	}).(VirtualMachineAdditionalCapabilitiesResponseOutput)
}

func (o VirtualMachineAdditionalCapabilitiesResponsePtrOutput) InstallGpuDrivers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineAdditionalCapabilitiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.InstallGpuDrivers
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineProfile struct {
	AdditionalCapabilities *VirtualMachineAdditionalCapabilities `pulumi:"additionalCapabilities"`
	AdminUser              Credentials                           `pulumi:"adminUser"`
	CreateOption           CreateOption                          `pulumi:"createOption"`
	ImageReference         ImageReference                        `pulumi:"imageReference"`
	NonAdminUser           *Credentials                          `pulumi:"nonAdminUser"`
	Sku                    Sku                                   `pulumi:"sku"`
	UsageQuota             string                                `pulumi:"usageQuota"`
	UseSharedPassword      *EnableState                          `pulumi:"useSharedPassword"`
}





type VirtualMachineProfileInput interface {
	pulumi.Input

	ToVirtualMachineProfileOutput() VirtualMachineProfileOutput
	ToVirtualMachineProfileOutputWithContext(context.Context) VirtualMachineProfileOutput
}

type VirtualMachineProfileArgs struct {
	AdditionalCapabilities VirtualMachineAdditionalCapabilitiesPtrInput `pulumi:"additionalCapabilities"`
	AdminUser              CredentialsInput                             `pulumi:"adminUser"`
	CreateOption           CreateOptionInput                            `pulumi:"createOption"`
	ImageReference         ImageReferenceInput                          `pulumi:"imageReference"`
	NonAdminUser           CredentialsPtrInput                          `pulumi:"nonAdminUser"`
	Sku                    SkuInput                                     `pulumi:"sku"`
	UsageQuota             pulumi.StringInput                           `pulumi:"usageQuota"`
	UseSharedPassword      EnableStatePtrInput                          `pulumi:"useSharedPassword"`
}

func (VirtualMachineProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineProfile)(nil)).Elem()
}

func (i VirtualMachineProfileArgs) ToVirtualMachineProfileOutput() VirtualMachineProfileOutput {
	return i.ToVirtualMachineProfileOutputWithContext(context.Background())
}

func (i VirtualMachineProfileArgs) ToVirtualMachineProfileOutputWithContext(ctx context.Context) VirtualMachineProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineProfileOutput)
}

func (i VirtualMachineProfileArgs) ToVirtualMachineProfilePtrOutput() VirtualMachineProfilePtrOutput {
	return i.ToVirtualMachineProfilePtrOutputWithContext(context.Background())
}

func (i VirtualMachineProfileArgs) ToVirtualMachineProfilePtrOutputWithContext(ctx context.Context) VirtualMachineProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineProfileOutput).ToVirtualMachineProfilePtrOutputWithContext(ctx)
}









type VirtualMachineProfilePtrInput interface {
	pulumi.Input

	ToVirtualMachineProfilePtrOutput() VirtualMachineProfilePtrOutput
	ToVirtualMachineProfilePtrOutputWithContext(context.Context) VirtualMachineProfilePtrOutput
}

type virtualMachineProfilePtrType VirtualMachineProfileArgs

func VirtualMachineProfilePtr(v *VirtualMachineProfileArgs) VirtualMachineProfilePtrInput {
	return (*virtualMachineProfilePtrType)(v)
}

func (*virtualMachineProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineProfile)(nil)).Elem()
}

func (i *virtualMachineProfilePtrType) ToVirtualMachineProfilePtrOutput() VirtualMachineProfilePtrOutput {
	return i.ToVirtualMachineProfilePtrOutputWithContext(context.Background())
}

func (i *virtualMachineProfilePtrType) ToVirtualMachineProfilePtrOutputWithContext(ctx context.Context) VirtualMachineProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineProfilePtrOutput)
}

type VirtualMachineProfileOutput struct{ *pulumi.OutputState }

func (VirtualMachineProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineProfile)(nil)).Elem()
}

func (o VirtualMachineProfileOutput) ToVirtualMachineProfileOutput() VirtualMachineProfileOutput {
	return o
}

func (o VirtualMachineProfileOutput) ToVirtualMachineProfileOutputWithContext(ctx context.Context) VirtualMachineProfileOutput {
	return o
}

func (o VirtualMachineProfileOutput) ToVirtualMachineProfilePtrOutput() VirtualMachineProfilePtrOutput {
	return o.ToVirtualMachineProfilePtrOutputWithContext(context.Background())
}

func (o VirtualMachineProfileOutput) ToVirtualMachineProfilePtrOutputWithContext(ctx context.Context) VirtualMachineProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineProfile) *VirtualMachineProfile {
		return &v
	}).(VirtualMachineProfilePtrOutput)
}

func (o VirtualMachineProfileOutput) AdditionalCapabilities() VirtualMachineAdditionalCapabilitiesPtrOutput {
	return o.ApplyT(func(v VirtualMachineProfile) *VirtualMachineAdditionalCapabilities { return v.AdditionalCapabilities }).(VirtualMachineAdditionalCapabilitiesPtrOutput)
}

func (o VirtualMachineProfileOutput) AdminUser() CredentialsOutput {
	return o.ApplyT(func(v VirtualMachineProfile) Credentials { return v.AdminUser }).(CredentialsOutput)
}

func (o VirtualMachineProfileOutput) CreateOption() CreateOptionOutput {
	return o.ApplyT(func(v VirtualMachineProfile) CreateOption { return v.CreateOption }).(CreateOptionOutput)
}

func (o VirtualMachineProfileOutput) ImageReference() ImageReferenceOutput {
	return o.ApplyT(func(v VirtualMachineProfile) ImageReference { return v.ImageReference }).(ImageReferenceOutput)
}

func (o VirtualMachineProfileOutput) NonAdminUser() CredentialsPtrOutput {
	return o.ApplyT(func(v VirtualMachineProfile) *Credentials { return v.NonAdminUser }).(CredentialsPtrOutput)
}

func (o VirtualMachineProfileOutput) Sku() SkuOutput {
	return o.ApplyT(func(v VirtualMachineProfile) Sku { return v.Sku }).(SkuOutput)
}

func (o VirtualMachineProfileOutput) UsageQuota() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineProfile) string { return v.UsageQuota }).(pulumi.StringOutput)
}

func (o VirtualMachineProfileOutput) UseSharedPassword() EnableStatePtrOutput {
	return o.ApplyT(func(v VirtualMachineProfile) *EnableState { return v.UseSharedPassword }).(EnableStatePtrOutput)
}

type VirtualMachineProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineProfile)(nil)).Elem()
}

func (o VirtualMachineProfilePtrOutput) ToVirtualMachineProfilePtrOutput() VirtualMachineProfilePtrOutput {
	return o
}

func (o VirtualMachineProfilePtrOutput) ToVirtualMachineProfilePtrOutputWithContext(ctx context.Context) VirtualMachineProfilePtrOutput {
	return o
}

func (o VirtualMachineProfilePtrOutput) Elem() VirtualMachineProfileOutput {
	return o.ApplyT(func(v *VirtualMachineProfile) VirtualMachineProfile {
		if v != nil {
			return *v
		}
		var ret VirtualMachineProfile
		return ret
	}).(VirtualMachineProfileOutput)
}

func (o VirtualMachineProfilePtrOutput) AdditionalCapabilities() VirtualMachineAdditionalCapabilitiesPtrOutput {
	return o.ApplyT(func(v *VirtualMachineProfile) *VirtualMachineAdditionalCapabilities {
		if v == nil {
			return nil
		}
		return v.AdditionalCapabilities
	}).(VirtualMachineAdditionalCapabilitiesPtrOutput)
}

func (o VirtualMachineProfilePtrOutput) AdminUser() CredentialsPtrOutput {
	return o.ApplyT(func(v *VirtualMachineProfile) *Credentials {
		if v == nil {
			return nil
		}
		return &v.AdminUser
	}).(CredentialsPtrOutput)
}

func (o VirtualMachineProfilePtrOutput) CreateOption() CreateOptionPtrOutput {
	return o.ApplyT(func(v *VirtualMachineProfile) *CreateOption {
		if v == nil {
			return nil
		}
		return &v.CreateOption
	}).(CreateOptionPtrOutput)
}

func (o VirtualMachineProfilePtrOutput) ImageReference() ImageReferencePtrOutput {
	return o.ApplyT(func(v *VirtualMachineProfile) *ImageReference {
		if v == nil {
			return nil
		}
		return &v.ImageReference
	}).(ImageReferencePtrOutput)
}

func (o VirtualMachineProfilePtrOutput) NonAdminUser() CredentialsPtrOutput {
	return o.ApplyT(func(v *VirtualMachineProfile) *Credentials {
		if v == nil {
			return nil
		}
		return v.NonAdminUser
	}).(CredentialsPtrOutput)
}

func (o VirtualMachineProfilePtrOutput) Sku() SkuPtrOutput {
	return o.ApplyT(func(v *VirtualMachineProfile) *Sku {
		if v == nil {
			return nil
		}
		return &v.Sku
	}).(SkuPtrOutput)
}

func (o VirtualMachineProfilePtrOutput) UsageQuota() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineProfile) *string {
		if v == nil {
			return nil
		}
		return &v.UsageQuota
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineProfilePtrOutput) UseSharedPassword() EnableStatePtrOutput {
	return o.ApplyT(func(v *VirtualMachineProfile) *EnableState {
		if v == nil {
			return nil
		}
		return v.UseSharedPassword
	}).(EnableStatePtrOutput)
}

type VirtualMachineProfileResponse struct {
	AdditionalCapabilities *VirtualMachineAdditionalCapabilitiesResponse `pulumi:"additionalCapabilities"`
	AdminUser              CredentialsResponse                           `pulumi:"adminUser"`
	CreateOption           string                                        `pulumi:"createOption"`
	ImageReference         ImageReferenceResponse                        `pulumi:"imageReference"`
	NonAdminUser           *CredentialsResponse                          `pulumi:"nonAdminUser"`
	OsType                 string                                        `pulumi:"osType"`
	Sku                    SkuResponse                                   `pulumi:"sku"`
	UsageQuota             string                                        `pulumi:"usageQuota"`
	UseSharedPassword      *string                                       `pulumi:"useSharedPassword"`
}





type VirtualMachineProfileResponseInput interface {
	pulumi.Input

	ToVirtualMachineProfileResponseOutput() VirtualMachineProfileResponseOutput
	ToVirtualMachineProfileResponseOutputWithContext(context.Context) VirtualMachineProfileResponseOutput
}

type VirtualMachineProfileResponseArgs struct {
	AdditionalCapabilities VirtualMachineAdditionalCapabilitiesResponsePtrInput `pulumi:"additionalCapabilities"`
	AdminUser              CredentialsResponseInput                             `pulumi:"adminUser"`
	CreateOption           pulumi.StringInput                                   `pulumi:"createOption"`
	ImageReference         ImageReferenceResponseInput                          `pulumi:"imageReference"`
	NonAdminUser           CredentialsResponsePtrInput                          `pulumi:"nonAdminUser"`
	OsType                 pulumi.StringInput                                   `pulumi:"osType"`
	Sku                    SkuResponseInput                                     `pulumi:"sku"`
	UsageQuota             pulumi.StringInput                                   `pulumi:"usageQuota"`
	UseSharedPassword      pulumi.StringPtrInput                                `pulumi:"useSharedPassword"`
}

func (VirtualMachineProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineProfileResponse)(nil)).Elem()
}

func (i VirtualMachineProfileResponseArgs) ToVirtualMachineProfileResponseOutput() VirtualMachineProfileResponseOutput {
	return i.ToVirtualMachineProfileResponseOutputWithContext(context.Background())
}

func (i VirtualMachineProfileResponseArgs) ToVirtualMachineProfileResponseOutputWithContext(ctx context.Context) VirtualMachineProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineProfileResponseOutput)
}

func (i VirtualMachineProfileResponseArgs) ToVirtualMachineProfileResponsePtrOutput() VirtualMachineProfileResponsePtrOutput {
	return i.ToVirtualMachineProfileResponsePtrOutputWithContext(context.Background())
}

func (i VirtualMachineProfileResponseArgs) ToVirtualMachineProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineProfileResponseOutput).ToVirtualMachineProfileResponsePtrOutputWithContext(ctx)
}









type VirtualMachineProfileResponsePtrInput interface {
	pulumi.Input

	ToVirtualMachineProfileResponsePtrOutput() VirtualMachineProfileResponsePtrOutput
	ToVirtualMachineProfileResponsePtrOutputWithContext(context.Context) VirtualMachineProfileResponsePtrOutput
}

type virtualMachineProfileResponsePtrType VirtualMachineProfileResponseArgs

func VirtualMachineProfileResponsePtr(v *VirtualMachineProfileResponseArgs) VirtualMachineProfileResponsePtrInput {
	return (*virtualMachineProfileResponsePtrType)(v)
}

func (*virtualMachineProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineProfileResponse)(nil)).Elem()
}

func (i *virtualMachineProfileResponsePtrType) ToVirtualMachineProfileResponsePtrOutput() VirtualMachineProfileResponsePtrOutput {
	return i.ToVirtualMachineProfileResponsePtrOutputWithContext(context.Background())
}

func (i *virtualMachineProfileResponsePtrType) ToVirtualMachineProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineProfileResponsePtrOutput)
}

type VirtualMachineProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineProfileResponse)(nil)).Elem()
}

func (o VirtualMachineProfileResponseOutput) ToVirtualMachineProfileResponseOutput() VirtualMachineProfileResponseOutput {
	return o
}

func (o VirtualMachineProfileResponseOutput) ToVirtualMachineProfileResponseOutputWithContext(ctx context.Context) VirtualMachineProfileResponseOutput {
	return o
}

func (o VirtualMachineProfileResponseOutput) ToVirtualMachineProfileResponsePtrOutput() VirtualMachineProfileResponsePtrOutput {
	return o.ToVirtualMachineProfileResponsePtrOutputWithContext(context.Background())
}

func (o VirtualMachineProfileResponseOutput) ToVirtualMachineProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineProfileResponse) *VirtualMachineProfileResponse {
		return &v
	}).(VirtualMachineProfileResponsePtrOutput)
}

func (o VirtualMachineProfileResponseOutput) AdditionalCapabilities() VirtualMachineAdditionalCapabilitiesResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineProfileResponse) *VirtualMachineAdditionalCapabilitiesResponse {
		return v.AdditionalCapabilities
	}).(VirtualMachineAdditionalCapabilitiesResponsePtrOutput)
}

func (o VirtualMachineProfileResponseOutput) AdminUser() CredentialsResponseOutput {
	return o.ApplyT(func(v VirtualMachineProfileResponse) CredentialsResponse { return v.AdminUser }).(CredentialsResponseOutput)
}

func (o VirtualMachineProfileResponseOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineProfileResponse) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o VirtualMachineProfileResponseOutput) ImageReference() ImageReferenceResponseOutput {
	return o.ApplyT(func(v VirtualMachineProfileResponse) ImageReferenceResponse { return v.ImageReference }).(ImageReferenceResponseOutput)
}

func (o VirtualMachineProfileResponseOutput) NonAdminUser() CredentialsResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineProfileResponse) *CredentialsResponse { return v.NonAdminUser }).(CredentialsResponsePtrOutput)
}

func (o VirtualMachineProfileResponseOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineProfileResponse) string { return v.OsType }).(pulumi.StringOutput)
}

func (o VirtualMachineProfileResponseOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v VirtualMachineProfileResponse) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o VirtualMachineProfileResponseOutput) UsageQuota() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineProfileResponse) string { return v.UsageQuota }).(pulumi.StringOutput)
}

func (o VirtualMachineProfileResponseOutput) UseSharedPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineProfileResponse) *string { return v.UseSharedPassword }).(pulumi.StringPtrOutput)
}

type VirtualMachineProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineProfileResponse)(nil)).Elem()
}

func (o VirtualMachineProfileResponsePtrOutput) ToVirtualMachineProfileResponsePtrOutput() VirtualMachineProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineProfileResponsePtrOutput) ToVirtualMachineProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineProfileResponsePtrOutput) Elem() VirtualMachineProfileResponseOutput {
	return o.ApplyT(func(v *VirtualMachineProfileResponse) VirtualMachineProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineProfileResponse
		return ret
	}).(VirtualMachineProfileResponseOutput)
}

func (o VirtualMachineProfileResponsePtrOutput) AdditionalCapabilities() VirtualMachineAdditionalCapabilitiesResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineProfileResponse) *VirtualMachineAdditionalCapabilitiesResponse {
		if v == nil {
			return nil
		}
		return v.AdditionalCapabilities
	}).(VirtualMachineAdditionalCapabilitiesResponsePtrOutput)
}

func (o VirtualMachineProfileResponsePtrOutput) AdminUser() CredentialsResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineProfileResponse) *CredentialsResponse {
		if v == nil {
			return nil
		}
		return &v.AdminUser
	}).(CredentialsResponsePtrOutput)
}

func (o VirtualMachineProfileResponsePtrOutput) CreateOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreateOption
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineProfileResponsePtrOutput) ImageReference() ImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineProfileResponse) *ImageReferenceResponse {
		if v == nil {
			return nil
		}
		return &v.ImageReference
	}).(ImageReferenceResponsePtrOutput)
}

func (o VirtualMachineProfileResponsePtrOutput) NonAdminUser() CredentialsResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineProfileResponse) *CredentialsResponse {
		if v == nil {
			return nil
		}
		return v.NonAdminUser
	}).(CredentialsResponsePtrOutput)
}

func (o VirtualMachineProfileResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineProfileResponsePtrOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineProfileResponse) *SkuResponse {
		if v == nil {
			return nil
		}
		return &v.Sku
	}).(SkuResponsePtrOutput)
}

func (o VirtualMachineProfileResponsePtrOutput) UsageQuota() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UsageQuota
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineProfileResponsePtrOutput) UseSharedPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.UseSharedPassword
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AutoShutdownProfileOutput{})
	pulumi.RegisterOutputType(AutoShutdownProfilePtrOutput{})
	pulumi.RegisterOutputType(AutoShutdownProfileResponseOutput{})
	pulumi.RegisterOutputType(AutoShutdownProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectionProfileOutput{})
	pulumi.RegisterOutputType(ConnectionProfilePtrOutput{})
	pulumi.RegisterOutputType(ConnectionProfileResponseOutput{})
	pulumi.RegisterOutputType(ConnectionProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(CredentialsOutput{})
	pulumi.RegisterOutputType(CredentialsPtrOutput{})
	pulumi.RegisterOutputType(CredentialsResponseOutput{})
	pulumi.RegisterOutputType(CredentialsResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageReferenceOutput{})
	pulumi.RegisterOutputType(ImageReferencePtrOutput{})
	pulumi.RegisterOutputType(ImageReferenceResponseOutput{})
	pulumi.RegisterOutputType(ImageReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(LabNetworkProfileOutput{})
	pulumi.RegisterOutputType(LabNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(LabNetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(LabNetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(LabPlanNetworkProfileOutput{})
	pulumi.RegisterOutputType(LabPlanNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(LabPlanNetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(LabPlanNetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(RecurrencePatternOutput{})
	pulumi.RegisterOutputType(RecurrencePatternPtrOutput{})
	pulumi.RegisterOutputType(RecurrencePatternResponseOutput{})
	pulumi.RegisterOutputType(RecurrencePatternResponsePtrOutput{})
	pulumi.RegisterOutputType(RosterProfileOutput{})
	pulumi.RegisterOutputType(RosterProfilePtrOutput{})
	pulumi.RegisterOutputType(RosterProfileResponseOutput{})
	pulumi.RegisterOutputType(RosterProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SecurityProfileOutput{})
	pulumi.RegisterOutputType(SecurityProfilePtrOutput{})
	pulumi.RegisterOutputType(SecurityProfileResponseOutput{})
	pulumi.RegisterOutputType(SecurityProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SupportInfoOutput{})
	pulumi.RegisterOutputType(SupportInfoPtrOutput{})
	pulumi.RegisterOutputType(SupportInfoResponseOutput{})
	pulumi.RegisterOutputType(SupportInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineAdditionalCapabilitiesOutput{})
	pulumi.RegisterOutputType(VirtualMachineAdditionalCapabilitiesPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineAdditionalCapabilitiesResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineAdditionalCapabilitiesResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineProfileOutput{})
	pulumi.RegisterOutputType(VirtualMachineProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineProfileResponsePtrOutput{})
}
