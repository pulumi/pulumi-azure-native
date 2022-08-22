


package cdn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecret(ctx *pulumi.Context, args *LookupSecretArgs, opts ...pulumi.InvokeOption) (*LookupSecretResult, error) {
	var rv LookupSecretResult
	err := ctx.Invoke("azure-native:cdn:getSecret", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecretArgs struct {
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SecretName        string `pulumi:"secretName"`
}


type LookupSecretResult struct {
	DeploymentStatus  string             `pulumi:"deploymentStatus"`
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	Parameters        interface{}        `pulumi:"parameters"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}

func LookupSecretOutput(ctx *pulumi.Context, args LookupSecretOutputArgs, opts ...pulumi.InvokeOption) LookupSecretResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecretResult, error) {
			args := v.(LookupSecretArgs)
			r, err := LookupSecret(ctx, &args, opts...)
			var s LookupSecretResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecretResultOutput)
}

type LookupSecretOutputArgs struct {
	ProfileName       pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SecretName        pulumi.StringInput `pulumi:"secretName"`
}

func (LookupSecretOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretArgs)(nil)).Elem()
}


type LookupSecretResultOutput struct{ *pulumi.OutputState }

func (LookupSecretResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretResult)(nil)).Elem()
}

func (o LookupSecretResultOutput) ToLookupSecretResultOutput() LookupSecretResultOutput {
	return o
}

func (o LookupSecretResultOutput) ToLookupSecretResultOutputWithContext(ctx context.Context) LookupSecretResultOutput {
	return o
}

func (o LookupSecretResultOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.DeploymentStatus }).(pulumi.StringOutput)
}

func (o LookupSecretResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSecretResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSecretResultOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupSecretResult) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o LookupSecretResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSecretResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSecretResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSecretResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecretResultOutput{})
}
