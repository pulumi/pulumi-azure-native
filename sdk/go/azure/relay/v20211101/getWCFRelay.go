


package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWCFRelay(ctx *pulumi.Context, args *LookupWCFRelayArgs, opts ...pulumi.InvokeOption) (*LookupWCFRelayResult, error) {
	var rv LookupWCFRelayResult
	err := ctx.Invoke("azure-native:relay/v20211101:getWCFRelay", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWCFRelayArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	RelayName         string `pulumi:"relayName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWCFRelayResult struct {
	CreatedAt                   string             `pulumi:"createdAt"`
	Id                          string             `pulumi:"id"`
	IsDynamic                   bool               `pulumi:"isDynamic"`
	ListenerCount               int                `pulumi:"listenerCount"`
	Location                    string             `pulumi:"location"`
	Name                        string             `pulumi:"name"`
	RelayType                   *string            `pulumi:"relayType"`
	RequiresClientAuthorization *bool              `pulumi:"requiresClientAuthorization"`
	RequiresTransportSecurity   *bool              `pulumi:"requiresTransportSecurity"`
	SystemData                  SystemDataResponse `pulumi:"systemData"`
	Type                        string             `pulumi:"type"`
	UpdatedAt                   string             `pulumi:"updatedAt"`
	UserMetadata                *string            `pulumi:"userMetadata"`
}

func LookupWCFRelayOutput(ctx *pulumi.Context, args LookupWCFRelayOutputArgs, opts ...pulumi.InvokeOption) LookupWCFRelayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWCFRelayResult, error) {
			args := v.(LookupWCFRelayArgs)
			r, err := LookupWCFRelay(ctx, &args, opts...)
			var s LookupWCFRelayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWCFRelayResultOutput)
}

type LookupWCFRelayOutputArgs struct {
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	RelayName         pulumi.StringInput `pulumi:"relayName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWCFRelayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWCFRelayArgs)(nil)).Elem()
}


type LookupWCFRelayResultOutput struct{ *pulumi.OutputState }

func (LookupWCFRelayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWCFRelayResult)(nil)).Elem()
}

func (o LookupWCFRelayResultOutput) ToLookupWCFRelayResultOutput() LookupWCFRelayResultOutput {
	return o
}

func (o LookupWCFRelayResultOutput) ToLookupWCFRelayResultOutputWithContext(ctx context.Context) LookupWCFRelayResultOutput {
	return o
}

func (o LookupWCFRelayResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWCFRelayResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupWCFRelayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWCFRelayResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWCFRelayResultOutput) IsDynamic() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWCFRelayResult) bool { return v.IsDynamic }).(pulumi.BoolOutput)
}

func (o LookupWCFRelayResultOutput) ListenerCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWCFRelayResult) int { return v.ListenerCount }).(pulumi.IntOutput)
}

func (o LookupWCFRelayResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWCFRelayResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupWCFRelayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWCFRelayResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWCFRelayResultOutput) RelayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWCFRelayResult) *string { return v.RelayType }).(pulumi.StringPtrOutput)
}

func (o LookupWCFRelayResultOutput) RequiresClientAuthorization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWCFRelayResult) *bool { return v.RequiresClientAuthorization }).(pulumi.BoolPtrOutput)
}

func (o LookupWCFRelayResultOutput) RequiresTransportSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWCFRelayResult) *bool { return v.RequiresTransportSecurity }).(pulumi.BoolPtrOutput)
}

func (o LookupWCFRelayResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWCFRelayResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupWCFRelayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWCFRelayResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWCFRelayResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWCFRelayResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupWCFRelayResultOutput) UserMetadata() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWCFRelayResult) *string { return v.UserMetadata }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWCFRelayResultOutput{})
}
