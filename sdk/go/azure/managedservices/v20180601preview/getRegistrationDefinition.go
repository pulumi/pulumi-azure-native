


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupRegistrationDefinition(ctx *pulumi.Context, args *LookupRegistrationDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupRegistrationDefinitionResult, error) {
	var rv LookupRegistrationDefinitionResult
	err := ctx.Invoke("azure-native:managedservices/v20180601preview:getRegistrationDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegistrationDefinitionArgs struct {
	RegistrationDefinitionId string `pulumi:"registrationDefinitionId"`
	Scope                    string `pulumi:"scope"`
}


type LookupRegistrationDefinitionResult struct {
	Id         string                                   `pulumi:"id"`
	Name       string                                   `pulumi:"name"`
	Plan       *PlanResponse                            `pulumi:"plan"`
	Properties RegistrationDefinitionPropertiesResponse `pulumi:"properties"`
	Type       string                                   `pulumi:"type"`
}

func LookupRegistrationDefinitionOutput(ctx *pulumi.Context, args LookupRegistrationDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupRegistrationDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistrationDefinitionResult, error) {
			args := v.(LookupRegistrationDefinitionArgs)
			r, err := LookupRegistrationDefinition(ctx, &args, opts...)
			var s LookupRegistrationDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistrationDefinitionResultOutput)
}

type LookupRegistrationDefinitionOutputArgs struct {
	RegistrationDefinitionId pulumi.StringInput `pulumi:"registrationDefinitionId"`
	Scope                    pulumi.StringInput `pulumi:"scope"`
}

func (LookupRegistrationDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistrationDefinitionArgs)(nil)).Elem()
}


type LookupRegistrationDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupRegistrationDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistrationDefinitionResult)(nil)).Elem()
}

func (o LookupRegistrationDefinitionResultOutput) ToLookupRegistrationDefinitionResultOutput() LookupRegistrationDefinitionResultOutput {
	return o
}

func (o LookupRegistrationDefinitionResultOutput) ToLookupRegistrationDefinitionResultOutputWithContext(ctx context.Context) LookupRegistrationDefinitionResultOutput {
	return o
}

func (o LookupRegistrationDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegistrationDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegistrationDefinitionResultOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v LookupRegistrationDefinitionResult) *PlanResponse { return v.Plan }).(PlanResponsePtrOutput)
}

func (o LookupRegistrationDefinitionResultOutput) Properties() RegistrationDefinitionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupRegistrationDefinitionResult) RegistrationDefinitionPropertiesResponse {
		return v.Properties
	}).(RegistrationDefinitionPropertiesResponseOutput)
}

func (o LookupRegistrationDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistrationDefinitionResultOutput{})
}
