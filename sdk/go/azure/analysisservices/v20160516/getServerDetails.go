


package v20160516

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupServerDetails(ctx *pulumi.Context, args *LookupServerDetailsArgs, opts ...pulumi.InvokeOption) (*LookupServerDetailsResult, error) {
	var rv LookupServerDetailsResult
	err := ctx.Invoke("azure-native:analysisservices/v20160516:getServerDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupServerDetailsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupServerDetailsResult struct {
	AsAdministrators       *ServerAdministratorsResponse `pulumi:"asAdministrators"`
	BackupBlobContainerUri *string                       `pulumi:"backupBlobContainerUri"`
	Id                     string                        `pulumi:"id"`
	Location               string                        `pulumi:"location"`
	ManagedMode            *int                          `pulumi:"managedMode"`
	Name                   string                        `pulumi:"name"`
	ProvisioningState      string                        `pulumi:"provisioningState"`
	ServerFullName         string                        `pulumi:"serverFullName"`
	ServerMonitorMode      *int                          `pulumi:"serverMonitorMode"`
	Sku                    ResourceSkuResponse           `pulumi:"sku"`
	State                  string                        `pulumi:"state"`
	Tags                   map[string]string             `pulumi:"tags"`
	Type                   string                        `pulumi:"type"`
}


func (val *LookupServerDetailsResult) Defaults() *LookupServerDetailsResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ManagedMode) {
		managedMode_ := 1
		tmp.ManagedMode = &managedMode_
	}
	if isZero(tmp.ServerMonitorMode) {
		serverMonitorMode_ := 1
		tmp.ServerMonitorMode = &serverMonitorMode_
	}
	tmp.Sku = *tmp.Sku.Defaults()

	return &tmp
}

func LookupServerDetailsOutput(ctx *pulumi.Context, args LookupServerDetailsOutputArgs, opts ...pulumi.InvokeOption) LookupServerDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerDetailsResult, error) {
			args := v.(LookupServerDetailsArgs)
			r, err := LookupServerDetails(ctx, &args, opts...)
			var s LookupServerDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerDetailsResultOutput)
}

type LookupServerDetailsOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerDetailsArgs)(nil)).Elem()
}


type LookupServerDetailsResultOutput struct{ *pulumi.OutputState }

func (LookupServerDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerDetailsResult)(nil)).Elem()
}

func (o LookupServerDetailsResultOutput) ToLookupServerDetailsResultOutput() LookupServerDetailsResultOutput {
	return o
}

func (o LookupServerDetailsResultOutput) ToLookupServerDetailsResultOutputWithContext(ctx context.Context) LookupServerDetailsResultOutput {
	return o
}

func (o LookupServerDetailsResultOutput) AsAdministrators() ServerAdministratorsResponsePtrOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) *ServerAdministratorsResponse { return v.AsAdministrators }).(ServerAdministratorsResponsePtrOutput)
}

func (o LookupServerDetailsResultOutput) BackupBlobContainerUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) *string { return v.BackupBlobContainerUri }).(pulumi.StringPtrOutput)
}

func (o LookupServerDetailsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerDetailsResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupServerDetailsResultOutput) ManagedMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) *int { return v.ManagedMode }).(pulumi.IntPtrOutput)
}

func (o LookupServerDetailsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerDetailsResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupServerDetailsResultOutput) ServerFullName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) string { return v.ServerFullName }).(pulumi.StringOutput)
}

func (o LookupServerDetailsResultOutput) ServerMonitorMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) *int { return v.ServerMonitorMode }).(pulumi.IntPtrOutput)
}

func (o LookupServerDetailsResultOutput) Sku() ResourceSkuResponseOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) ResourceSkuResponse { return v.Sku }).(ResourceSkuResponseOutput)
}

func (o LookupServerDetailsResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupServerDetailsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupServerDetailsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerDetailsResultOutput{})
}
