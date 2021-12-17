


package labservices

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


func (val *AutoShutdownProfile) Defaults() *AutoShutdownProfile {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ShutdownOnDisconnect) {
		shutdownOnDisconnect_ := EnableState("Disabled")
		tmp.ShutdownOnDisconnect = &shutdownOnDisconnect_
	}
	if isZero(tmp.ShutdownOnIdle) {
		shutdownOnIdle_ := ShutdownOnIdleMode("None")
		tmp.ShutdownOnIdle = &shutdownOnIdle_
	}
	if isZero(tmp.ShutdownWhenNotConnected) {
		shutdownWhenNotConnected_ := EnableState("Disabled")
		tmp.ShutdownWhenNotConnected = &shutdownWhenNotConnected_
	}
	return &tmp
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


func (val *AutoShutdownProfileResponse) Defaults() *AutoShutdownProfileResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ShutdownOnDisconnect) {
		shutdownOnDisconnect_ := "Disabled"
		tmp.ShutdownOnDisconnect = &shutdownOnDisconnect_
	}
	if isZero(tmp.ShutdownOnIdle) {
		shutdownOnIdle_ := "None"
		tmp.ShutdownOnIdle = &shutdownOnIdle_
	}
	if isZero(tmp.ShutdownWhenNotConnected) {
		shutdownWhenNotConnected_ := "Disabled"
		tmp.ShutdownWhenNotConnected = &shutdownWhenNotConnected_
	}
	return &tmp
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


func (val *ConnectionProfile) Defaults() *ConnectionProfile {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ClientRdpAccess) {
		clientRdpAccess_ := ConnectionType("None")
		tmp.ClientRdpAccess = &clientRdpAccess_
	}
	if isZero(tmp.ClientSshAccess) {
		clientSshAccess_ := ConnectionType("None")
		tmp.ClientSshAccess = &clientSshAccess_
	}
	if isZero(tmp.WebRdpAccess) {
		webRdpAccess_ := ConnectionType("None")
		tmp.WebRdpAccess = &webRdpAccess_
	}
	if isZero(tmp.WebSshAccess) {
		webSshAccess_ := ConnectionType("None")
		tmp.WebSshAccess = &webSshAccess_
	}
	return &tmp
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


func (val *ConnectionProfileResponse) Defaults() *ConnectionProfileResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ClientRdpAccess) {
		clientRdpAccess_ := "None"
		tmp.ClientRdpAccess = &clientRdpAccess_
	}
	if isZero(tmp.ClientSshAccess) {
		clientSshAccess_ := "None"
		tmp.ClientSshAccess = &clientSshAccess_
	}
	if isZero(tmp.WebRdpAccess) {
		webRdpAccess_ := "None"
		tmp.WebRdpAccess = &webRdpAccess_
	}
	if isZero(tmp.WebSshAccess) {
		webSshAccess_ := "None"
		tmp.WebSshAccess = &webSshAccess_
	}
	return &tmp
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

type EnvironmentDetailsResponse struct {
	Description           string                        `pulumi:"description"`
	EnvironmentState      string                        `pulumi:"environmentState"`
	Id                    string                        `pulumi:"id"`
	LatestOperationResult LatestOperationResultResponse `pulumi:"latestOperationResult"`
	Name                  string                        `pulumi:"name"`
	PasswordLastReset     string                        `pulumi:"passwordLastReset"`
	ProvisioningState     string                        `pulumi:"provisioningState"`
	TotalUsage            string                        `pulumi:"totalUsage"`
	VirtualMachineDetails VirtualMachineDetailsResponse `pulumi:"virtualMachineDetails"`
}

type EnvironmentSizeResponse struct {
	MaxPrice         float64            `pulumi:"maxPrice"`
	MinMemory        float64            `pulumi:"minMemory"`
	MinNumberOfCores int                `pulumi:"minNumberOfCores"`
	Name             *string            `pulumi:"name"`
	VmSizes          []SizeInfoResponse `pulumi:"vmSizes"`
}

type EnvironmentSizeResponseOutput struct{ *pulumi.OutputState }

func (EnvironmentSizeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentSizeResponse)(nil)).Elem()
}

func (o EnvironmentSizeResponseOutput) ToEnvironmentSizeResponseOutput() EnvironmentSizeResponseOutput {
	return o
}

func (o EnvironmentSizeResponseOutput) ToEnvironmentSizeResponseOutputWithContext(ctx context.Context) EnvironmentSizeResponseOutput {
	return o
}

func (o EnvironmentSizeResponseOutput) MaxPrice() pulumi.Float64Output {
	return o.ApplyT(func(v EnvironmentSizeResponse) float64 { return v.MaxPrice }).(pulumi.Float64Output)
}

func (o EnvironmentSizeResponseOutput) MinMemory() pulumi.Float64Output {
	return o.ApplyT(func(v EnvironmentSizeResponse) float64 { return v.MinMemory }).(pulumi.Float64Output)
}

func (o EnvironmentSizeResponseOutput) MinNumberOfCores() pulumi.IntOutput {
	return o.ApplyT(func(v EnvironmentSizeResponse) int { return v.MinNumberOfCores }).(pulumi.IntOutput)
}

func (o EnvironmentSizeResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentSizeResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EnvironmentSizeResponseOutput) VmSizes() SizeInfoResponseArrayOutput {
	return o.ApplyT(func(v EnvironmentSizeResponse) []SizeInfoResponse { return v.VmSizes }).(SizeInfoResponseArrayOutput)
}

type EnvironmentSizeResponseArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentSizeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentSizeResponse)(nil)).Elem()
}

func (o EnvironmentSizeResponseArrayOutput) ToEnvironmentSizeResponseArrayOutput() EnvironmentSizeResponseArrayOutput {
	return o
}

func (o EnvironmentSizeResponseArrayOutput) ToEnvironmentSizeResponseArrayOutputWithContext(ctx context.Context) EnvironmentSizeResponseArrayOutput {
	return o
}

func (o EnvironmentSizeResponseArrayOutput) Index(i pulumi.IntInput) EnvironmentSizeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentSizeResponse {
		return vs[0].([]EnvironmentSizeResponse)[vs[1].(int)]
	}).(EnvironmentSizeResponseOutput)
}

type GalleryImageReferenceResponse struct {
	Offer     *string `pulumi:"offer"`
	OsType    *string `pulumi:"osType"`
	Publisher *string `pulumi:"publisher"`
	Sku       *string `pulumi:"sku"`
	Version   *string `pulumi:"version"`
}

type GalleryImageReferenceResponseOutput struct{ *pulumi.OutputState }

func (GalleryImageReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageReferenceResponse)(nil)).Elem()
}

func (o GalleryImageReferenceResponseOutput) ToGalleryImageReferenceResponseOutput() GalleryImageReferenceResponseOutput {
	return o
}

func (o GalleryImageReferenceResponseOutput) ToGalleryImageReferenceResponseOutputWithContext(ctx context.Context) GalleryImageReferenceResponseOutput {
	return o
}

func (o GalleryImageReferenceResponseOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponseOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type LabDetailsResponse struct {
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	ProvisioningState *string `pulumi:"provisioningState"`
	UsageQuota        string  `pulumi:"usageQuota"`
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

type LatestOperationResultResponse struct {
	ErrorCode    string `pulumi:"errorCode"`
	ErrorMessage string `pulumi:"errorMessage"`
	HttpMethod   string `pulumi:"httpMethod"`
	OperationUrl string `pulumi:"operationUrl"`
	RequestUri   string `pulumi:"requestUri"`
	Status       string `pulumi:"status"`
}

type LatestOperationResultResponseOutput struct{ *pulumi.OutputState }

func (LatestOperationResultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LatestOperationResultResponse)(nil)).Elem()
}

func (o LatestOperationResultResponseOutput) ToLatestOperationResultResponseOutput() LatestOperationResultResponseOutput {
	return o
}

func (o LatestOperationResultResponseOutput) ToLatestOperationResultResponseOutputWithContext(ctx context.Context) LatestOperationResultResponseOutput {
	return o
}

func (o LatestOperationResultResponseOutput) ErrorCode() pulumi.StringOutput {
	return o.ApplyT(func(v LatestOperationResultResponse) string { return v.ErrorCode }).(pulumi.StringOutput)
}

func (o LatestOperationResultResponseOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LatestOperationResultResponse) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o LatestOperationResultResponseOutput) HttpMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LatestOperationResultResponse) string { return v.HttpMethod }).(pulumi.StringOutput)
}

func (o LatestOperationResultResponseOutput) OperationUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LatestOperationResultResponse) string { return v.OperationUrl }).(pulumi.StringOutput)
}

func (o LatestOperationResultResponseOutput) RequestUri() pulumi.StringOutput {
	return o.ApplyT(func(v LatestOperationResultResponse) string { return v.RequestUri }).(pulumi.StringOutput)
}

func (o LatestOperationResultResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LatestOperationResultResponse) string { return v.Status }).(pulumi.StringOutput)
}

type NetworkInterfaceResponse struct {
	PrivateIpAddress string `pulumi:"privateIpAddress"`
	RdpAuthority     string `pulumi:"rdpAuthority"`
	SshAuthority     string `pulumi:"sshAuthority"`
	Username         string `pulumi:"username"`
}

type NetworkInterfaceResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceResponse)(nil)).Elem()
}

func (o NetworkInterfaceResponseOutput) ToNetworkInterfaceResponseOutput() NetworkInterfaceResponseOutput {
	return o
}

func (o NetworkInterfaceResponseOutput) ToNetworkInterfaceResponseOutputWithContext(ctx context.Context) NetworkInterfaceResponseOutput {
	return o
}

func (o NetworkInterfaceResponseOutput) PrivateIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) string { return v.PrivateIpAddress }).(pulumi.StringOutput)
}

func (o NetworkInterfaceResponseOutput) RdpAuthority() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) string { return v.RdpAuthority }).(pulumi.StringOutput)
}

func (o NetworkInterfaceResponseOutput) SshAuthority() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) string { return v.SshAuthority }).(pulumi.StringOutput)
}

func (o NetworkInterfaceResponseOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) string { return v.Username }).(pulumi.StringOutput)
}

type OperationBatchStatusResponseItemResponse struct {
	OperationUrl string `pulumi:"operationUrl"`
	Status       string `pulumi:"status"`
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

type ReferenceVm struct {
	Password *string `pulumi:"password"`
	UserName string  `pulumi:"userName"`
}





type ReferenceVmInput interface {
	pulumi.Input

	ToReferenceVmOutput() ReferenceVmOutput
	ToReferenceVmOutputWithContext(context.Context) ReferenceVmOutput
}

type ReferenceVmArgs struct {
	Password pulumi.StringPtrInput `pulumi:"password"`
	UserName pulumi.StringInput    `pulumi:"userName"`
}

func (ReferenceVmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceVm)(nil)).Elem()
}

func (i ReferenceVmArgs) ToReferenceVmOutput() ReferenceVmOutput {
	return i.ToReferenceVmOutputWithContext(context.Background())
}

func (i ReferenceVmArgs) ToReferenceVmOutputWithContext(ctx context.Context) ReferenceVmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceVmOutput)
}

type ReferenceVmOutput struct{ *pulumi.OutputState }

func (ReferenceVmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceVm)(nil)).Elem()
}

func (o ReferenceVmOutput) ToReferenceVmOutput() ReferenceVmOutput {
	return o
}

func (o ReferenceVmOutput) ToReferenceVmOutputWithContext(ctx context.Context) ReferenceVmOutput {
	return o
}

func (o ReferenceVmOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReferenceVm) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ReferenceVmOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v ReferenceVm) string { return v.UserName }).(pulumi.StringOutput)
}

type ReferenceVmResponse struct {
	Password       *string                `pulumi:"password"`
	UserName       string                 `pulumi:"userName"`
	VmResourceId   string                 `pulumi:"vmResourceId"`
	VmStateDetails VmStateDetailsResponse `pulumi:"vmStateDetails"`
}

type ReferenceVmResponseOutput struct{ *pulumi.OutputState }

func (ReferenceVmResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceVmResponse)(nil)).Elem()
}

func (o ReferenceVmResponseOutput) ToReferenceVmResponseOutput() ReferenceVmResponseOutput {
	return o
}

func (o ReferenceVmResponseOutput) ToReferenceVmResponseOutputWithContext(ctx context.Context) ReferenceVmResponseOutput {
	return o
}

func (o ReferenceVmResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReferenceVmResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ReferenceVmResponseOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v ReferenceVmResponse) string { return v.UserName }).(pulumi.StringOutput)
}

func (o ReferenceVmResponseOutput) VmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ReferenceVmResponse) string { return v.VmResourceId }).(pulumi.StringOutput)
}

func (o ReferenceVmResponseOutput) VmStateDetails() VmStateDetailsResponseOutput {
	return o.ApplyT(func(v ReferenceVmResponse) VmStateDetailsResponse { return v.VmStateDetails }).(VmStateDetailsResponseOutput)
}

type RegionalAvailabilityResponse struct {
	Region             *string                    `pulumi:"region"`
	SizeAvailabilities []SizeAvailabilityResponse `pulumi:"sizeAvailabilities"`
}

type ResourceSet struct {
	ResourceSettingId *string `pulumi:"resourceSettingId"`
	VmResourceId      *string `pulumi:"vmResourceId"`
}





type ResourceSetInput interface {
	pulumi.Input

	ToResourceSetOutput() ResourceSetOutput
	ToResourceSetOutputWithContext(context.Context) ResourceSetOutput
}

type ResourceSetArgs struct {
	ResourceSettingId pulumi.StringPtrInput `pulumi:"resourceSettingId"`
	VmResourceId      pulumi.StringPtrInput `pulumi:"vmResourceId"`
}

func (ResourceSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSet)(nil)).Elem()
}

func (i ResourceSetArgs) ToResourceSetOutput() ResourceSetOutput {
	return i.ToResourceSetOutputWithContext(context.Background())
}

func (i ResourceSetArgs) ToResourceSetOutputWithContext(ctx context.Context) ResourceSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSetOutput)
}

func (i ResourceSetArgs) ToResourceSetPtrOutput() ResourceSetPtrOutput {
	return i.ToResourceSetPtrOutputWithContext(context.Background())
}

func (i ResourceSetArgs) ToResourceSetPtrOutputWithContext(ctx context.Context) ResourceSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSetOutput).ToResourceSetPtrOutputWithContext(ctx)
}









type ResourceSetPtrInput interface {
	pulumi.Input

	ToResourceSetPtrOutput() ResourceSetPtrOutput
	ToResourceSetPtrOutputWithContext(context.Context) ResourceSetPtrOutput
}

type resourceSetPtrType ResourceSetArgs

func ResourceSetPtr(v *ResourceSetArgs) ResourceSetPtrInput {
	return (*resourceSetPtrType)(v)
}

func (*resourceSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSet)(nil)).Elem()
}

func (i *resourceSetPtrType) ToResourceSetPtrOutput() ResourceSetPtrOutput {
	return i.ToResourceSetPtrOutputWithContext(context.Background())
}

func (i *resourceSetPtrType) ToResourceSetPtrOutputWithContext(ctx context.Context) ResourceSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSetPtrOutput)
}

type ResourceSetOutput struct{ *pulumi.OutputState }

func (ResourceSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSet)(nil)).Elem()
}

func (o ResourceSetOutput) ToResourceSetOutput() ResourceSetOutput {
	return o
}

func (o ResourceSetOutput) ToResourceSetOutputWithContext(ctx context.Context) ResourceSetOutput {
	return o
}

func (o ResourceSetOutput) ToResourceSetPtrOutput() ResourceSetPtrOutput {
	return o.ToResourceSetPtrOutputWithContext(context.Background())
}

func (o ResourceSetOutput) ToResourceSetPtrOutputWithContext(ctx context.Context) ResourceSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSet) *ResourceSet {
		return &v
	}).(ResourceSetPtrOutput)
}

func (o ResourceSetOutput) ResourceSettingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSet) *string { return v.ResourceSettingId }).(pulumi.StringPtrOutput)
}

func (o ResourceSetOutput) VmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSet) *string { return v.VmResourceId }).(pulumi.StringPtrOutput)
}

type ResourceSetPtrOutput struct{ *pulumi.OutputState }

func (ResourceSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSet)(nil)).Elem()
}

func (o ResourceSetPtrOutput) ToResourceSetPtrOutput() ResourceSetPtrOutput {
	return o
}

func (o ResourceSetPtrOutput) ToResourceSetPtrOutputWithContext(ctx context.Context) ResourceSetPtrOutput {
	return o
}

func (o ResourceSetPtrOutput) Elem() ResourceSetOutput {
	return o.ApplyT(func(v *ResourceSet) ResourceSet {
		if v != nil {
			return *v
		}
		var ret ResourceSet
		return ret
	}).(ResourceSetOutput)
}

func (o ResourceSetPtrOutput) ResourceSettingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSet) *string {
		if v == nil {
			return nil
		}
		return v.ResourceSettingId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSetPtrOutput) VmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSet) *string {
		if v == nil {
			return nil
		}
		return v.VmResourceId
	}).(pulumi.StringPtrOutput)
}

type ResourceSetResponse struct {
	ResourceSettingId *string `pulumi:"resourceSettingId"`
	VmResourceId      *string `pulumi:"vmResourceId"`
}

type ResourceSetResponseOutput struct{ *pulumi.OutputState }

func (ResourceSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSetResponse)(nil)).Elem()
}

func (o ResourceSetResponseOutput) ToResourceSetResponseOutput() ResourceSetResponseOutput {
	return o
}

func (o ResourceSetResponseOutput) ToResourceSetResponseOutputWithContext(ctx context.Context) ResourceSetResponseOutput {
	return o
}

func (o ResourceSetResponseOutput) ResourceSettingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSetResponse) *string { return v.ResourceSettingId }).(pulumi.StringPtrOutput)
}

func (o ResourceSetResponseOutput) VmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSetResponse) *string { return v.VmResourceId }).(pulumi.StringPtrOutput)
}

type ResourceSetResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceSetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSetResponse)(nil)).Elem()
}

func (o ResourceSetResponsePtrOutput) ToResourceSetResponsePtrOutput() ResourceSetResponsePtrOutput {
	return o
}

func (o ResourceSetResponsePtrOutput) ToResourceSetResponsePtrOutputWithContext(ctx context.Context) ResourceSetResponsePtrOutput {
	return o
}

func (o ResourceSetResponsePtrOutput) Elem() ResourceSetResponseOutput {
	return o.ApplyT(func(v *ResourceSetResponse) ResourceSetResponse {
		if v != nil {
			return *v
		}
		var ret ResourceSetResponse
		return ret
	}).(ResourceSetResponseOutput)
}

func (o ResourceSetResponsePtrOutput) ResourceSettingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceSettingId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSetResponsePtrOutput) VmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.VmResourceId
	}).(pulumi.StringPtrOutput)
}

type ResourceSettings struct {
	GalleryImageResourceId *string     `pulumi:"galleryImageResourceId"`
	ReferenceVm            ReferenceVm `pulumi:"referenceVm"`
	Size                   *string     `pulumi:"size"`
}





type ResourceSettingsInput interface {
	pulumi.Input

	ToResourceSettingsOutput() ResourceSettingsOutput
	ToResourceSettingsOutputWithContext(context.Context) ResourceSettingsOutput
}

type ResourceSettingsArgs struct {
	GalleryImageResourceId pulumi.StringPtrInput `pulumi:"galleryImageResourceId"`
	ReferenceVm            ReferenceVmInput      `pulumi:"referenceVm"`
	Size                   pulumi.StringPtrInput `pulumi:"size"`
}

func (ResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSettings)(nil)).Elem()
}

func (i ResourceSettingsArgs) ToResourceSettingsOutput() ResourceSettingsOutput {
	return i.ToResourceSettingsOutputWithContext(context.Background())
}

func (i ResourceSettingsArgs) ToResourceSettingsOutputWithContext(ctx context.Context) ResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSettingsOutput)
}

type ResourceSettingsOutput struct{ *pulumi.OutputState }

func (ResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSettings)(nil)).Elem()
}

func (o ResourceSettingsOutput) ToResourceSettingsOutput() ResourceSettingsOutput {
	return o
}

func (o ResourceSettingsOutput) ToResourceSettingsOutputWithContext(ctx context.Context) ResourceSettingsOutput {
	return o
}

func (o ResourceSettingsOutput) GalleryImageResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSettings) *string { return v.GalleryImageResourceId }).(pulumi.StringPtrOutput)
}

func (o ResourceSettingsOutput) ReferenceVm() ReferenceVmOutput {
	return o.ApplyT(func(v ResourceSettings) ReferenceVm { return v.ReferenceVm }).(ReferenceVmOutput)
}

func (o ResourceSettingsOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSettings) *string { return v.Size }).(pulumi.StringPtrOutput)
}

type ResourceSettingsResponse struct {
	Cores                  int                 `pulumi:"cores"`
	GalleryImageResourceId *string             `pulumi:"galleryImageResourceId"`
	Id                     string              `pulumi:"id"`
	ImageName              string              `pulumi:"imageName"`
	ReferenceVm            ReferenceVmResponse `pulumi:"referenceVm"`
	Size                   *string             `pulumi:"size"`
}

type ResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (ResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSettingsResponse)(nil)).Elem()
}

func (o ResourceSettingsResponseOutput) ToResourceSettingsResponseOutput() ResourceSettingsResponseOutput {
	return o
}

func (o ResourceSettingsResponseOutput) ToResourceSettingsResponseOutputWithContext(ctx context.Context) ResourceSettingsResponseOutput {
	return o
}

func (o ResourceSettingsResponseOutput) Cores() pulumi.IntOutput {
	return o.ApplyT(func(v ResourceSettingsResponse) int { return v.Cores }).(pulumi.IntOutput)
}

func (o ResourceSettingsResponseOutput) GalleryImageResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSettingsResponse) *string { return v.GalleryImageResourceId }).(pulumi.StringPtrOutput)
}

func (o ResourceSettingsResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSettingsResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ResourceSettingsResponseOutput) ImageName() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSettingsResponse) string { return v.ImageName }).(pulumi.StringOutput)
}

func (o ResourceSettingsResponseOutput) ReferenceVm() ReferenceVmResponseOutput {
	return o.ApplyT(func(v ResourceSettingsResponse) ReferenceVmResponse { return v.ReferenceVm }).(ReferenceVmResponseOutput)
}

func (o ResourceSettingsResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSettingsResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

type SizeAvailabilityResponse struct {
	IsAvailable  *bool   `pulumi:"isAvailable"`
	SizeCategory *string `pulumi:"sizeCategory"`
}

type SizeConfigurationPropertiesResponse struct {
	EnvironmentSizes []EnvironmentSizeResponse `pulumi:"environmentSizes"`
}

type SizeConfigurationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SizeConfigurationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SizeConfigurationPropertiesResponse)(nil)).Elem()
}

func (o SizeConfigurationPropertiesResponseOutput) ToSizeConfigurationPropertiesResponseOutput() SizeConfigurationPropertiesResponseOutput {
	return o
}

func (o SizeConfigurationPropertiesResponseOutput) ToSizeConfigurationPropertiesResponseOutputWithContext(ctx context.Context) SizeConfigurationPropertiesResponseOutput {
	return o
}

func (o SizeConfigurationPropertiesResponseOutput) EnvironmentSizes() EnvironmentSizeResponseArrayOutput {
	return o.ApplyT(func(v SizeConfigurationPropertiesResponse) []EnvironmentSizeResponse { return v.EnvironmentSizes }).(EnvironmentSizeResponseArrayOutput)
}

type SizeInfoResponse struct {
	ComputeSize   *string  `pulumi:"computeSize"`
	Memory        *float64 `pulumi:"memory"`
	NumberOfCores *int     `pulumi:"numberOfCores"`
	Price         *float64 `pulumi:"price"`
}

type SizeInfoResponseOutput struct{ *pulumi.OutputState }

func (SizeInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SizeInfoResponse)(nil)).Elem()
}

func (o SizeInfoResponseOutput) ToSizeInfoResponseOutput() SizeInfoResponseOutput {
	return o
}

func (o SizeInfoResponseOutput) ToSizeInfoResponseOutputWithContext(ctx context.Context) SizeInfoResponseOutput {
	return o
}

func (o SizeInfoResponseOutput) ComputeSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SizeInfoResponse) *string { return v.ComputeSize }).(pulumi.StringPtrOutput)
}

func (o SizeInfoResponseOutput) Memory() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SizeInfoResponse) *float64 { return v.Memory }).(pulumi.Float64PtrOutput)
}

func (o SizeInfoResponseOutput) NumberOfCores() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SizeInfoResponse) *int { return v.NumberOfCores }).(pulumi.IntPtrOutput)
}

func (o SizeInfoResponseOutput) Price() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SizeInfoResponse) *float64 { return v.Price }).(pulumi.Float64PtrOutput)
}

type SizeInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (SizeInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SizeInfoResponse)(nil)).Elem()
}

func (o SizeInfoResponseArrayOutput) ToSizeInfoResponseArrayOutput() SizeInfoResponseArrayOutput {
	return o
}

func (o SizeInfoResponseArrayOutput) ToSizeInfoResponseArrayOutputWithContext(ctx context.Context) SizeInfoResponseArrayOutput {
	return o
}

func (o SizeInfoResponseArrayOutput) Index(i pulumi.IntInput) SizeInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SizeInfoResponse {
		return vs[0].([]SizeInfoResponse)[vs[1].(int)]
	}).(SizeInfoResponseOutput)
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

type VirtualMachineDetailsResponse struct {
	LastKnownPowerState string `pulumi:"lastKnownPowerState"`
	PrivateIpAddress    string `pulumi:"privateIpAddress"`
	ProvisioningState   string `pulumi:"provisioningState"`
	RdpAuthority        string `pulumi:"rdpAuthority"`
	SshAuthority        string `pulumi:"sshAuthority"`
	UserName            string `pulumi:"userName"`
}

type VmStateDetailsResponse struct {
	LastKnownPowerState string `pulumi:"lastKnownPowerState"`
	PowerState          string `pulumi:"powerState"`
	RdpAuthority        string `pulumi:"rdpAuthority"`
	SshAuthority        string `pulumi:"sshAuthority"`
}

type VmStateDetailsResponseOutput struct{ *pulumi.OutputState }

func (VmStateDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmStateDetailsResponse)(nil)).Elem()
}

func (o VmStateDetailsResponseOutput) ToVmStateDetailsResponseOutput() VmStateDetailsResponseOutput {
	return o
}

func (o VmStateDetailsResponseOutput) ToVmStateDetailsResponseOutputWithContext(ctx context.Context) VmStateDetailsResponseOutput {
	return o
}

func (o VmStateDetailsResponseOutput) LastKnownPowerState() pulumi.StringOutput {
	return o.ApplyT(func(v VmStateDetailsResponse) string { return v.LastKnownPowerState }).(pulumi.StringOutput)
}

func (o VmStateDetailsResponseOutput) PowerState() pulumi.StringOutput {
	return o.ApplyT(func(v VmStateDetailsResponse) string { return v.PowerState }).(pulumi.StringOutput)
}

func (o VmStateDetailsResponseOutput) RdpAuthority() pulumi.StringOutput {
	return o.ApplyT(func(v VmStateDetailsResponse) string { return v.RdpAuthority }).(pulumi.StringOutput)
}

func (o VmStateDetailsResponseOutput) SshAuthority() pulumi.StringOutput {
	return o.ApplyT(func(v VmStateDetailsResponse) string { return v.SshAuthority }).(pulumi.StringOutput)
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
	pulumi.RegisterOutputType(EnvironmentSizeResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentSizeResponseArrayOutput{})
	pulumi.RegisterOutputType(GalleryImageReferenceResponseOutput{})
	pulumi.RegisterOutputType(LabPlanNetworkProfileOutput{})
	pulumi.RegisterOutputType(LabPlanNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(LabPlanNetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(LabPlanNetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(LatestOperationResultResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceResponseOutput{})
	pulumi.RegisterOutputType(RecurrencePatternOutput{})
	pulumi.RegisterOutputType(RecurrencePatternPtrOutput{})
	pulumi.RegisterOutputType(RecurrencePatternResponseOutput{})
	pulumi.RegisterOutputType(RecurrencePatternResponsePtrOutput{})
	pulumi.RegisterOutputType(ReferenceVmOutput{})
	pulumi.RegisterOutputType(ReferenceVmResponseOutput{})
	pulumi.RegisterOutputType(ResourceSetOutput{})
	pulumi.RegisterOutputType(ResourceSetPtrOutput{})
	pulumi.RegisterOutputType(ResourceSetResponseOutput{})
	pulumi.RegisterOutputType(ResourceSetResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceSettingsOutput{})
	pulumi.RegisterOutputType(ResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(SizeConfigurationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SizeInfoResponseOutput{})
	pulumi.RegisterOutputType(SizeInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(SupportInfoOutput{})
	pulumi.RegisterOutputType(SupportInfoPtrOutput{})
	pulumi.RegisterOutputType(SupportInfoResponseOutput{})
	pulumi.RegisterOutputType(SupportInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(VmStateDetailsResponseOutput{})
}
