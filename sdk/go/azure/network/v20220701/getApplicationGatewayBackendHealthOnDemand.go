


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetApplicationGatewayBackendHealthOnDemand(ctx *pulumi.Context, args *GetApplicationGatewayBackendHealthOnDemandArgs, opts ...pulumi.InvokeOption) (*GetApplicationGatewayBackendHealthOnDemandResult, error) {
	var rv GetApplicationGatewayBackendHealthOnDemandResult
	err := ctx.Invoke("azure-native:network/v20220701:getApplicationGatewayBackendHealthOnDemand", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetApplicationGatewayBackendHealthOnDemandArgs struct {
	ApplicationGatewayName              string                                      `pulumi:"applicationGatewayName"`
	BackendAddressPool                  *SubResource                                `pulumi:"backendAddressPool"`
	BackendHttpSettings                 *SubResource                                `pulumi:"backendHttpSettings"`
	Expand                              *string                                     `pulumi:"expand"`
	Host                                *string                                     `pulumi:"host"`
	Match                               *ApplicationGatewayProbeHealthResponseMatch `pulumi:"match"`
	Path                                *string                                     `pulumi:"path"`
	PickHostNameFromBackendHttpSettings *bool                                       `pulumi:"pickHostNameFromBackendHttpSettings"`
	Protocol                            *string                                     `pulumi:"protocol"`
	ResourceGroupName                   string                                      `pulumi:"resourceGroupName"`
	Timeout                             *int                                        `pulumi:"timeout"`
}


type GetApplicationGatewayBackendHealthOnDemandResult struct {
	BackendAddressPool        *ApplicationGatewayBackendAddressPoolResponse        `pulumi:"backendAddressPool"`
	BackendHealthHttpSettings *ApplicationGatewayBackendHealthHttpSettingsResponse `pulumi:"backendHealthHttpSettings"`
}

func GetApplicationGatewayBackendHealthOnDemandOutput(ctx *pulumi.Context, args GetApplicationGatewayBackendHealthOnDemandOutputArgs, opts ...pulumi.InvokeOption) GetApplicationGatewayBackendHealthOnDemandResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetApplicationGatewayBackendHealthOnDemandResult, error) {
			args := v.(GetApplicationGatewayBackendHealthOnDemandArgs)
			r, err := GetApplicationGatewayBackendHealthOnDemand(ctx, &args, opts...)
			var s GetApplicationGatewayBackendHealthOnDemandResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetApplicationGatewayBackendHealthOnDemandResultOutput)
}

type GetApplicationGatewayBackendHealthOnDemandOutputArgs struct {
	ApplicationGatewayName              pulumi.StringInput                                 `pulumi:"applicationGatewayName"`
	BackendAddressPool                  SubResourcePtrInput                                `pulumi:"backendAddressPool"`
	BackendHttpSettings                 SubResourcePtrInput                                `pulumi:"backendHttpSettings"`
	Expand                              pulumi.StringPtrInput                              `pulumi:"expand"`
	Host                                pulumi.StringPtrInput                              `pulumi:"host"`
	Match                               ApplicationGatewayProbeHealthResponseMatchPtrInput `pulumi:"match"`
	Path                                pulumi.StringPtrInput                              `pulumi:"path"`
	PickHostNameFromBackendHttpSettings pulumi.BoolPtrInput                                `pulumi:"pickHostNameFromBackendHttpSettings"`
	Protocol                            pulumi.StringPtrInput                              `pulumi:"protocol"`
	ResourceGroupName                   pulumi.StringInput                                 `pulumi:"resourceGroupName"`
	Timeout                             pulumi.IntPtrInput                                 `pulumi:"timeout"`
}

func (GetApplicationGatewayBackendHealthOnDemandOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationGatewayBackendHealthOnDemandArgs)(nil)).Elem()
}


type GetApplicationGatewayBackendHealthOnDemandResultOutput struct{ *pulumi.OutputState }

func (GetApplicationGatewayBackendHealthOnDemandResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationGatewayBackendHealthOnDemandResult)(nil)).Elem()
}

func (o GetApplicationGatewayBackendHealthOnDemandResultOutput) ToGetApplicationGatewayBackendHealthOnDemandResultOutput() GetApplicationGatewayBackendHealthOnDemandResultOutput {
	return o
}

func (o GetApplicationGatewayBackendHealthOnDemandResultOutput) ToGetApplicationGatewayBackendHealthOnDemandResultOutputWithContext(ctx context.Context) GetApplicationGatewayBackendHealthOnDemandResultOutput {
	return o
}

func (o GetApplicationGatewayBackendHealthOnDemandResultOutput) BackendAddressPool() ApplicationGatewayBackendAddressPoolResponsePtrOutput {
	return o.ApplyT(func(v GetApplicationGatewayBackendHealthOnDemandResult) *ApplicationGatewayBackendAddressPoolResponse {
		return v.BackendAddressPool
	}).(ApplicationGatewayBackendAddressPoolResponsePtrOutput)
}

func (o GetApplicationGatewayBackendHealthOnDemandResultOutput) BackendHealthHttpSettings() ApplicationGatewayBackendHealthHttpSettingsResponsePtrOutput {
	return o.ApplyT(func(v GetApplicationGatewayBackendHealthOnDemandResult) *ApplicationGatewayBackendHealthHttpSettingsResponse {
		return v.BackendHealthHttpSettings
	}).(ApplicationGatewayBackendHealthHttpSettingsResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApplicationGatewayBackendHealthOnDemandResultOutput{})
}
