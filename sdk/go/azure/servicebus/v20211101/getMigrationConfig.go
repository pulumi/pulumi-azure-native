


package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMigrationConfig(ctx *pulumi.Context, args *LookupMigrationConfigArgs, opts ...pulumi.InvokeOption) (*LookupMigrationConfigResult, error) {
	var rv LookupMigrationConfigResult
	err := ctx.Invoke("azure-native:servicebus/v20211101:getMigrationConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMigrationConfigArgs struct {
	ConfigName        string `pulumi:"configName"`
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupMigrationConfigResult struct {
	Id                                string             `pulumi:"id"`
	Location                          string             `pulumi:"location"`
	MigrationState                    string             `pulumi:"migrationState"`
	Name                              string             `pulumi:"name"`
	PendingReplicationOperationsCount float64            `pulumi:"pendingReplicationOperationsCount"`
	PostMigrationName                 string             `pulumi:"postMigrationName"`
	ProvisioningState                 string             `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse `pulumi:"systemData"`
	TargetNamespace                   string             `pulumi:"targetNamespace"`
	Type                              string             `pulumi:"type"`
}

func LookupMigrationConfigOutput(ctx *pulumi.Context, args LookupMigrationConfigOutputArgs, opts ...pulumi.InvokeOption) LookupMigrationConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMigrationConfigResult, error) {
			args := v.(LookupMigrationConfigArgs)
			r, err := LookupMigrationConfig(ctx, &args, opts...)
			var s LookupMigrationConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMigrationConfigResultOutput)
}

type LookupMigrationConfigOutputArgs struct {
	ConfigName        pulumi.StringInput `pulumi:"configName"`
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMigrationConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMigrationConfigArgs)(nil)).Elem()
}


type LookupMigrationConfigResultOutput struct{ *pulumi.OutputState }

func (LookupMigrationConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMigrationConfigResult)(nil)).Elem()
}

func (o LookupMigrationConfigResultOutput) ToLookupMigrationConfigResultOutput() LookupMigrationConfigResultOutput {
	return o
}

func (o LookupMigrationConfigResultOutput) ToLookupMigrationConfigResultOutputWithContext(ctx context.Context) LookupMigrationConfigResultOutput {
	return o
}

func (o LookupMigrationConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMigrationConfigResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationConfigResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupMigrationConfigResultOutput) MigrationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationConfigResult) string { return v.MigrationState }).(pulumi.StringOutput)
}

func (o LookupMigrationConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMigrationConfigResultOutput) PendingReplicationOperationsCount() pulumi.Float64Output {
	return o.ApplyT(func(v LookupMigrationConfigResult) float64 { return v.PendingReplicationOperationsCount }).(pulumi.Float64Output)
}

func (o LookupMigrationConfigResultOutput) PostMigrationName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationConfigResult) string { return v.PostMigrationName }).(pulumi.StringOutput)
}

func (o LookupMigrationConfigResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationConfigResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupMigrationConfigResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMigrationConfigResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMigrationConfigResultOutput) TargetNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationConfigResult) string { return v.TargetNamespace }).(pulumi.StringOutput)
}

func (o LookupMigrationConfigResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationConfigResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMigrationConfigResultOutput{})
}
