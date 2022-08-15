


package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHybridConnection(ctx *pulumi.Context, args *LookupHybridConnectionArgs, opts ...pulumi.InvokeOption) (*LookupHybridConnectionResult, error) {
	var rv LookupHybridConnectionResult
	err := ctx.Invoke("azure-native:relay/v20211101:getHybridConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHybridConnectionArgs struct {
	HybridConnectionName string `pulumi:"hybridConnectionName"`
	NamespaceName        string `pulumi:"namespaceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupHybridConnectionResult struct {
	CreatedAt                   string             `pulumi:"createdAt"`
	Id                          string             `pulumi:"id"`
	ListenerCount               int                `pulumi:"listenerCount"`
	Location                    string             `pulumi:"location"`
	Name                        string             `pulumi:"name"`
	RequiresClientAuthorization *bool              `pulumi:"requiresClientAuthorization"`
	SystemData                  SystemDataResponse `pulumi:"systemData"`
	Type                        string             `pulumi:"type"`
	UpdatedAt                   string             `pulumi:"updatedAt"`
	UserMetadata                *string            `pulumi:"userMetadata"`
}

func LookupHybridConnectionOutput(ctx *pulumi.Context, args LookupHybridConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupHybridConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHybridConnectionResult, error) {
			args := v.(LookupHybridConnectionArgs)
			r, err := LookupHybridConnection(ctx, &args, opts...)
			var s LookupHybridConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHybridConnectionResultOutput)
}

type LookupHybridConnectionOutputArgs struct {
	HybridConnectionName pulumi.StringInput `pulumi:"hybridConnectionName"`
	NamespaceName        pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupHybridConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridConnectionArgs)(nil)).Elem()
}


type LookupHybridConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupHybridConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridConnectionResult)(nil)).Elem()
}

func (o LookupHybridConnectionResultOutput) ToLookupHybridConnectionResultOutput() LookupHybridConnectionResultOutput {
	return o
}

func (o LookupHybridConnectionResultOutput) ToLookupHybridConnectionResultOutputWithContext(ctx context.Context) LookupHybridConnectionResultOutput {
	return o
}

func (o LookupHybridConnectionResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridConnectionResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupHybridConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHybridConnectionResultOutput) ListenerCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHybridConnectionResult) int { return v.ListenerCount }).(pulumi.IntOutput)
}

func (o LookupHybridConnectionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridConnectionResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupHybridConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupHybridConnectionResultOutput) RequiresClientAuthorization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupHybridConnectionResult) *bool { return v.RequiresClientAuthorization }).(pulumi.BoolPtrOutput)
}

func (o LookupHybridConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupHybridConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupHybridConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupHybridConnectionResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridConnectionResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupHybridConnectionResultOutput) UserMetadata() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHybridConnectionResult) *string { return v.UserMetadata }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHybridConnectionResultOutput{})
}
