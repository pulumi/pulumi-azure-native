


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackupVault(ctx *pulumi.Context, args *LookupBackupVaultArgs, opts ...pulumi.InvokeOption) (*LookupBackupVaultResult, error) {
	var rv LookupBackupVaultResult
	err := ctx.Invoke("azure-native:dataprotection/v20220901preview:getBackupVault", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupVaultArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VaultName         string `pulumi:"vaultName"`
}


type LookupBackupVaultResult struct {
	ETag       *string                     `pulumi:"eTag"`
	Id         string                      `pulumi:"id"`
	Identity   *DppIdentityDetailsResponse `pulumi:"identity"`
	Location   string                      `pulumi:"location"`
	Name       string                      `pulumi:"name"`
	Properties BackupVaultResponse         `pulumi:"properties"`
	SystemData SystemDataResponse          `pulumi:"systemData"`
	Tags       map[string]string           `pulumi:"tags"`
	Type       string                      `pulumi:"type"`
}

func LookupBackupVaultOutput(ctx *pulumi.Context, args LookupBackupVaultOutputArgs, opts ...pulumi.InvokeOption) LookupBackupVaultResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupVaultResult, error) {
			args := v.(LookupBackupVaultArgs)
			r, err := LookupBackupVault(ctx, &args, opts...)
			var s LookupBackupVaultResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackupVaultResultOutput)
}

type LookupBackupVaultOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VaultName         pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupBackupVaultOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupVaultArgs)(nil)).Elem()
}


type LookupBackupVaultResultOutput struct{ *pulumi.OutputState }

func (LookupBackupVaultResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupVaultResult)(nil)).Elem()
}

func (o LookupBackupVaultResultOutput) ToLookupBackupVaultResultOutput() LookupBackupVaultResultOutput {
	return o
}

func (o LookupBackupVaultResultOutput) ToLookupBackupVaultResultOutputWithContext(ctx context.Context) LookupBackupVaultResultOutput {
	return o
}

func (o LookupBackupVaultResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupBackupVaultResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBackupVaultResultOutput) Identity() DppIdentityDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) *DppIdentityDetailsResponse { return v.Identity }).(DppIdentityDetailsResponsePtrOutput)
}

func (o LookupBackupVaultResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupBackupVaultResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBackupVaultResultOutput) Properties() BackupVaultResponseOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) BackupVaultResponse { return v.Properties }).(BackupVaultResponseOutput)
}

func (o LookupBackupVaultResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupBackupVaultResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupBackupVaultResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupVaultResultOutput{})
}
