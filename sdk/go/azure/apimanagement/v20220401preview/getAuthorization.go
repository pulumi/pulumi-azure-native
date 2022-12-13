


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAuthorization(ctx *pulumi.Context, args *LookupAuthorizationArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizationResult, error) {
	var rv LookupAuthorizationResult
	err := ctx.Invoke("azure-native:apimanagement/v20220401preview:getAuthorization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAuthorizationArgs struct {
	AuthorizationId         string `pulumi:"authorizationId"`
	AuthorizationProviderId string `pulumi:"authorizationProviderId"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	ServiceName             string `pulumi:"serviceName"`
}


type LookupAuthorizationResult struct {
	AuthorizationType *string                     `pulumi:"authorizationType"`
	Error             *AuthorizationErrorResponse `pulumi:"error"`
	Id                string                      `pulumi:"id"`
	Name              string                      `pulumi:"name"`
	OAuth2GrantType   *string                     `pulumi:"oAuth2GrantType"`
	Parameters        map[string]string           `pulumi:"parameters"`
	Status            *string                     `pulumi:"status"`
	Type              string                      `pulumi:"type"`
}

func LookupAuthorizationOutput(ctx *pulumi.Context, args LookupAuthorizationOutputArgs, opts ...pulumi.InvokeOption) LookupAuthorizationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAuthorizationResult, error) {
			args := v.(LookupAuthorizationArgs)
			r, err := LookupAuthorization(ctx, &args, opts...)
			var s LookupAuthorizationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAuthorizationResultOutput)
}

type LookupAuthorizationOutputArgs struct {
	AuthorizationId         pulumi.StringInput `pulumi:"authorizationId"`
	AuthorizationProviderId pulumi.StringInput `pulumi:"authorizationProviderId"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName             pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupAuthorizationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationArgs)(nil)).Elem()
}


type LookupAuthorizationResultOutput struct{ *pulumi.OutputState }

func (LookupAuthorizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationResult)(nil)).Elem()
}

func (o LookupAuthorizationResultOutput) ToLookupAuthorizationResultOutput() LookupAuthorizationResultOutput {
	return o
}

func (o LookupAuthorizationResultOutput) ToLookupAuthorizationResultOutputWithContext(ctx context.Context) LookupAuthorizationResultOutput {
	return o
}

func (o LookupAuthorizationResultOutput) AuthorizationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) *string { return v.AuthorizationType }).(pulumi.StringPtrOutput)
}

func (o LookupAuthorizationResultOutput) Error() AuthorizationErrorResponsePtrOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) *AuthorizationErrorResponse { return v.Error }).(AuthorizationErrorResponsePtrOutput)
}

func (o LookupAuthorizationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAuthorizationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAuthorizationResultOutput) OAuth2GrantType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) *string { return v.OAuth2GrantType }).(pulumi.StringPtrOutput)
}

func (o LookupAuthorizationResultOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o LookupAuthorizationResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupAuthorizationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthorizationResultOutput{})
}
