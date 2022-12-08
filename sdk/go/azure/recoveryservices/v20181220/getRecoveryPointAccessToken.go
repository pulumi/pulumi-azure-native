


package v20181220

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRecoveryPointAccessToken(ctx *pulumi.Context, args *GetRecoveryPointAccessTokenArgs, opts ...pulumi.InvokeOption) (*GetRecoveryPointAccessTokenResult, error) {
	var rv GetRecoveryPointAccessTokenResult
	err := ctx.Invoke("azure-native:recoveryservices/v20181220:getRecoveryPointAccessToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetRecoveryPointAccessTokenArgs struct {
	ContainerName     string            `pulumi:"containerName"`
	ETag              *string           `pulumi:"eTag"`
	FabricName        string            `pulumi:"fabricName"`
	Location          *string           `pulumi:"location"`
	Properties        *AADProperties    `pulumi:"properties"`
	ProtectedItemName string            `pulumi:"protectedItemName"`
	RecoveryPointId   string            `pulumi:"recoveryPointId"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	VaultName         string            `pulumi:"vaultName"`
}

type GetRecoveryPointAccessTokenResult struct {
	ETag       *string                        `pulumi:"eTag"`
	Id         string                         `pulumi:"id"`
	Location   *string                        `pulumi:"location"`
	Name       string                         `pulumi:"name"`
	Properties WorkloadCrrAccessTokenResponse `pulumi:"properties"`
	Tags       map[string]string              `pulumi:"tags"`
	Type       string                         `pulumi:"type"`
}

func GetRecoveryPointAccessTokenOutput(ctx *pulumi.Context, args GetRecoveryPointAccessTokenOutputArgs, opts ...pulumi.InvokeOption) GetRecoveryPointAccessTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRecoveryPointAccessTokenResult, error) {
			args := v.(GetRecoveryPointAccessTokenArgs)
			r, err := GetRecoveryPointAccessToken(ctx, &args, opts...)
			var s GetRecoveryPointAccessTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRecoveryPointAccessTokenResultOutput)
}

type GetRecoveryPointAccessTokenOutputArgs struct {
	ContainerName     pulumi.StringInput    `pulumi:"containerName"`
	ETag              pulumi.StringPtrInput `pulumi:"eTag"`
	FabricName        pulumi.StringInput    `pulumi:"fabricName"`
	Location          pulumi.StringPtrInput `pulumi:"location"`
	Properties        AADPropertiesPtrInput `pulumi:"properties"`
	ProtectedItemName pulumi.StringInput    `pulumi:"protectedItemName"`
	RecoveryPointId   pulumi.StringInput    `pulumi:"recoveryPointId"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	Tags              pulumi.StringMapInput `pulumi:"tags"`
	VaultName         pulumi.StringInput    `pulumi:"vaultName"`
}

func (GetRecoveryPointAccessTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRecoveryPointAccessTokenArgs)(nil)).Elem()
}

type GetRecoveryPointAccessTokenResultOutput struct{ *pulumi.OutputState }

func (GetRecoveryPointAccessTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRecoveryPointAccessTokenResult)(nil)).Elem()
}

func (o GetRecoveryPointAccessTokenResultOutput) ToGetRecoveryPointAccessTokenResultOutput() GetRecoveryPointAccessTokenResultOutput {
	return o
}

func (o GetRecoveryPointAccessTokenResultOutput) ToGetRecoveryPointAccessTokenResultOutputWithContext(ctx context.Context) GetRecoveryPointAccessTokenResultOutput {
	return o
}

func (o GetRecoveryPointAccessTokenResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecoveryPointAccessTokenResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o GetRecoveryPointAccessTokenResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRecoveryPointAccessTokenResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRecoveryPointAccessTokenResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecoveryPointAccessTokenResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o GetRecoveryPointAccessTokenResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetRecoveryPointAccessTokenResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetRecoveryPointAccessTokenResultOutput) Properties() WorkloadCrrAccessTokenResponseOutput {
	return o.ApplyT(func(v GetRecoveryPointAccessTokenResult) WorkloadCrrAccessTokenResponse { return v.Properties }).(WorkloadCrrAccessTokenResponseOutput)
}

func (o GetRecoveryPointAccessTokenResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetRecoveryPointAccessTokenResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetRecoveryPointAccessTokenResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetRecoveryPointAccessTokenResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRecoveryPointAccessTokenResultOutput{})
}
