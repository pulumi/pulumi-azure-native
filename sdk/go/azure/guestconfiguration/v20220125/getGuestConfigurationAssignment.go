


package v20220125

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGuestConfigurationAssignment(ctx *pulumi.Context, args *LookupGuestConfigurationAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupGuestConfigurationAssignmentResult, error) {
	var rv LookupGuestConfigurationAssignmentResult
	err := ctx.Invoke("azure-native:guestconfiguration/v20220125:getGuestConfigurationAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupGuestConfigurationAssignmentArgs struct {
	GuestConfigurationAssignmentName string `pulumi:"guestConfigurationAssignmentName"`
	ResourceGroupName                string `pulumi:"resourceGroupName"`
	VmName                           string `pulumi:"vmName"`
}


type LookupGuestConfigurationAssignmentResult struct {
	Id         string                                         `pulumi:"id"`
	Location   *string                                        `pulumi:"location"`
	Name       *string                                        `pulumi:"name"`
	Properties GuestConfigurationAssignmentPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                             `pulumi:"systemData"`
	Type       string                                         `pulumi:"type"`
}


func (val *LookupGuestConfigurationAssignmentResult) Defaults() *LookupGuestConfigurationAssignmentResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupGuestConfigurationAssignmentOutput(ctx *pulumi.Context, args LookupGuestConfigurationAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupGuestConfigurationAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGuestConfigurationAssignmentResult, error) {
			args := v.(LookupGuestConfigurationAssignmentArgs)
			r, err := LookupGuestConfigurationAssignment(ctx, &args, opts...)
			var s LookupGuestConfigurationAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGuestConfigurationAssignmentResultOutput)
}

type LookupGuestConfigurationAssignmentOutputArgs struct {
	GuestConfigurationAssignmentName pulumi.StringInput `pulumi:"guestConfigurationAssignmentName"`
	ResourceGroupName                pulumi.StringInput `pulumi:"resourceGroupName"`
	VmName                           pulumi.StringInput `pulumi:"vmName"`
}

func (LookupGuestConfigurationAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestConfigurationAssignmentArgs)(nil)).Elem()
}


type LookupGuestConfigurationAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupGuestConfigurationAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestConfigurationAssignmentResult)(nil)).Elem()
}

func (o LookupGuestConfigurationAssignmentResultOutput) ToLookupGuestConfigurationAssignmentResultOutput() LookupGuestConfigurationAssignmentResultOutput {
	return o
}

func (o LookupGuestConfigurationAssignmentResultOutput) ToLookupGuestConfigurationAssignmentResultOutputWithContext(ctx context.Context) LookupGuestConfigurationAssignmentResultOutput {
	return o
}

func (o LookupGuestConfigurationAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestConfigurationAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGuestConfigurationAssignmentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGuestConfigurationAssignmentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupGuestConfigurationAssignmentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGuestConfigurationAssignmentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupGuestConfigurationAssignmentResultOutput) Properties() GuestConfigurationAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v LookupGuestConfigurationAssignmentResult) GuestConfigurationAssignmentPropertiesResponse {
		return v.Properties
	}).(GuestConfigurationAssignmentPropertiesResponseOutput)
}

func (o LookupGuestConfigurationAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGuestConfigurationAssignmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupGuestConfigurationAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestConfigurationAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGuestConfigurationAssignmentResultOutput{})
}
