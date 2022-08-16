


package v20190901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegistrationAssignment(ctx *pulumi.Context, args *LookupRegistrationAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupRegistrationAssignmentResult, error) {
	var rv LookupRegistrationAssignmentResult
	err := ctx.Invoke("azure-native:managedservices/v20190901:getRegistrationAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegistrationAssignmentArgs struct {
	ExpandRegistrationDefinition *bool  `pulumi:"expandRegistrationDefinition"`
	RegistrationAssignmentId     string `pulumi:"registrationAssignmentId"`
	Scope                        string `pulumi:"scope"`
}


type LookupRegistrationAssignmentResult struct {
	Id         string                                   `pulumi:"id"`
	Name       string                                   `pulumi:"name"`
	Properties RegistrationAssignmentPropertiesResponse `pulumi:"properties"`
	Type       string                                   `pulumi:"type"`
}

func LookupRegistrationAssignmentOutput(ctx *pulumi.Context, args LookupRegistrationAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupRegistrationAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistrationAssignmentResult, error) {
			args := v.(LookupRegistrationAssignmentArgs)
			r, err := LookupRegistrationAssignment(ctx, &args, opts...)
			var s LookupRegistrationAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistrationAssignmentResultOutput)
}

type LookupRegistrationAssignmentOutputArgs struct {
	ExpandRegistrationDefinition pulumi.BoolPtrInput `pulumi:"expandRegistrationDefinition"`
	RegistrationAssignmentId     pulumi.StringInput  `pulumi:"registrationAssignmentId"`
	Scope                        pulumi.StringInput  `pulumi:"scope"`
}

func (LookupRegistrationAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistrationAssignmentArgs)(nil)).Elem()
}


type LookupRegistrationAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupRegistrationAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistrationAssignmentResult)(nil)).Elem()
}

func (o LookupRegistrationAssignmentResultOutput) ToLookupRegistrationAssignmentResultOutput() LookupRegistrationAssignmentResultOutput {
	return o
}

func (o LookupRegistrationAssignmentResultOutput) ToLookupRegistrationAssignmentResultOutputWithContext(ctx context.Context) LookupRegistrationAssignmentResultOutput {
	return o
}

func (o LookupRegistrationAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegistrationAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegistrationAssignmentResultOutput) Properties() RegistrationAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v LookupRegistrationAssignmentResult) RegistrationAssignmentPropertiesResponse {
		return v.Properties
	}).(RegistrationAssignmentPropertiesResponseOutput)
}

func (o LookupRegistrationAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistrationAssignmentResultOutput{})
}
