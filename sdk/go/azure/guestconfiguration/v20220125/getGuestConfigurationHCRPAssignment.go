


package v20220125

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGuestConfigurationHCRPAssignment(ctx *pulumi.Context, args *LookupGuestConfigurationHCRPAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupGuestConfigurationHCRPAssignmentResult, error) {
	var rv LookupGuestConfigurationHCRPAssignmentResult
	err := ctx.Invoke("azure-native:guestconfiguration/v20220125:getGuestConfigurationHCRPAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupGuestConfigurationHCRPAssignmentArgs struct {
	GuestConfigurationAssignmentName string `pulumi:"guestConfigurationAssignmentName"`
	MachineName                      string `pulumi:"machineName"`
	ResourceGroupName                string `pulumi:"resourceGroupName"`
}


type LookupGuestConfigurationHCRPAssignmentResult struct {
	Id         string                                         `pulumi:"id"`
	Location   *string                                        `pulumi:"location"`
	Name       *string                                        `pulumi:"name"`
	Properties GuestConfigurationAssignmentPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                             `pulumi:"systemData"`
	Type       string                                         `pulumi:"type"`
}


func (val *LookupGuestConfigurationHCRPAssignmentResult) Defaults() *LookupGuestConfigurationHCRPAssignmentResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupGuestConfigurationHCRPAssignmentOutput(ctx *pulumi.Context, args LookupGuestConfigurationHCRPAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupGuestConfigurationHCRPAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGuestConfigurationHCRPAssignmentResult, error) {
			args := v.(LookupGuestConfigurationHCRPAssignmentArgs)
			r, err := LookupGuestConfigurationHCRPAssignment(ctx, &args, opts...)
			var s LookupGuestConfigurationHCRPAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGuestConfigurationHCRPAssignmentResultOutput)
}

type LookupGuestConfigurationHCRPAssignmentOutputArgs struct {
	GuestConfigurationAssignmentName pulumi.StringInput `pulumi:"guestConfigurationAssignmentName"`
	MachineName                      pulumi.StringInput `pulumi:"machineName"`
	ResourceGroupName                pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGuestConfigurationHCRPAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestConfigurationHCRPAssignmentArgs)(nil)).Elem()
}


type LookupGuestConfigurationHCRPAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupGuestConfigurationHCRPAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestConfigurationHCRPAssignmentResult)(nil)).Elem()
}

func (o LookupGuestConfigurationHCRPAssignmentResultOutput) ToLookupGuestConfigurationHCRPAssignmentResultOutput() LookupGuestConfigurationHCRPAssignmentResultOutput {
	return o
}

func (o LookupGuestConfigurationHCRPAssignmentResultOutput) ToLookupGuestConfigurationHCRPAssignmentResultOutputWithContext(ctx context.Context) LookupGuestConfigurationHCRPAssignmentResultOutput {
	return o
}

func (o LookupGuestConfigurationHCRPAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestConfigurationHCRPAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGuestConfigurationHCRPAssignmentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGuestConfigurationHCRPAssignmentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupGuestConfigurationHCRPAssignmentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGuestConfigurationHCRPAssignmentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupGuestConfigurationHCRPAssignmentResultOutput) Properties() GuestConfigurationAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v LookupGuestConfigurationHCRPAssignmentResult) GuestConfigurationAssignmentPropertiesResponse {
		return v.Properties
	}).(GuestConfigurationAssignmentPropertiesResponseOutput)
}

func (o LookupGuestConfigurationHCRPAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGuestConfigurationHCRPAssignmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupGuestConfigurationHCRPAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestConfigurationHCRPAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGuestConfigurationHCRPAssignmentResultOutput{})
}
