


package v20221027

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListApplianceKeys(ctx *pulumi.Context, args *ListApplianceKeysArgs, opts ...pulumi.InvokeOption) (*ListApplianceKeysResult, error) {
	var rv ListApplianceKeysResult
	err := ctx.Invoke("azure-native:resourceconnector/v20221027:listApplianceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListApplianceKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListApplianceKeysResult struct {
	ArtifactProfiles map[string]ArtifactProfileResponse      `pulumi:"artifactProfiles"`
	Kubeconfigs      []ApplianceCredentialKubeconfigResponse `pulumi:"kubeconfigs"`
	SshKeys          map[string]SSHKeyResponse               `pulumi:"sshKeys"`
}

func ListApplianceKeysOutput(ctx *pulumi.Context, args ListApplianceKeysOutputArgs, opts ...pulumi.InvokeOption) ListApplianceKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListApplianceKeysResult, error) {
			args := v.(ListApplianceKeysArgs)
			r, err := ListApplianceKeys(ctx, &args, opts...)
			var s ListApplianceKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListApplianceKeysResultOutput)
}

type ListApplianceKeysOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (ListApplianceKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListApplianceKeysArgs)(nil)).Elem()
}


type ListApplianceKeysResultOutput struct{ *pulumi.OutputState }

func (ListApplianceKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListApplianceKeysResult)(nil)).Elem()
}

func (o ListApplianceKeysResultOutput) ToListApplianceKeysResultOutput() ListApplianceKeysResultOutput {
	return o
}

func (o ListApplianceKeysResultOutput) ToListApplianceKeysResultOutputWithContext(ctx context.Context) ListApplianceKeysResultOutput {
	return o
}

func (o ListApplianceKeysResultOutput) ArtifactProfiles() ArtifactProfileResponseMapOutput {
	return o.ApplyT(func(v ListApplianceKeysResult) map[string]ArtifactProfileResponse { return v.ArtifactProfiles }).(ArtifactProfileResponseMapOutput)
}

func (o ListApplianceKeysResultOutput) Kubeconfigs() ApplianceCredentialKubeconfigResponseArrayOutput {
	return o.ApplyT(func(v ListApplianceKeysResult) []ApplianceCredentialKubeconfigResponse { return v.Kubeconfigs }).(ApplianceCredentialKubeconfigResponseArrayOutput)
}

func (o ListApplianceKeysResultOutput) SshKeys() SSHKeyResponseMapOutput {
	return o.ApplyT(func(v ListApplianceKeysResult) map[string]SSHKeyResponse { return v.SshKeys }).(SSHKeyResponseMapOutput)
}

func init() {
	pulumi.RegisterOutputType(ListApplianceKeysResultOutput{})
}
