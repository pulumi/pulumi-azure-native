


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIscsiTarget(ctx *pulumi.Context, args *LookupIscsiTargetArgs, opts ...pulumi.InvokeOption) (*LookupIscsiTargetResult, error) {
	var rv LookupIscsiTargetResult
	err := ctx.Invoke("azure-native:storagepool/v20210801:getIscsiTarget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIscsiTargetArgs struct {
	DiskPoolName      string `pulumi:"diskPoolName"`
	IscsiTargetName   string `pulumi:"iscsiTargetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupIscsiTargetResult struct {
	AclMode           string                 `pulumi:"aclMode"`
	Endpoints         []string               `pulumi:"endpoints"`
	Id                string                 `pulumi:"id"`
	Luns              []IscsiLunResponse     `pulumi:"luns"`
	ManagedBy         string                 `pulumi:"managedBy"`
	ManagedByExtended []string               `pulumi:"managedByExtended"`
	Name              string                 `pulumi:"name"`
	Port              *int                   `pulumi:"port"`
	ProvisioningState string                 `pulumi:"provisioningState"`
	Sessions          []string               `pulumi:"sessions"`
	StaticAcls        []AclResponse          `pulumi:"staticAcls"`
	Status            string                 `pulumi:"status"`
	SystemData        SystemMetadataResponse `pulumi:"systemData"`
	TargetIqn         string                 `pulumi:"targetIqn"`
	Type              string                 `pulumi:"type"`
}

func LookupIscsiTargetOutput(ctx *pulumi.Context, args LookupIscsiTargetOutputArgs, opts ...pulumi.InvokeOption) LookupIscsiTargetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIscsiTargetResult, error) {
			args := v.(LookupIscsiTargetArgs)
			r, err := LookupIscsiTarget(ctx, &args, opts...)
			var s LookupIscsiTargetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIscsiTargetResultOutput)
}

type LookupIscsiTargetOutputArgs struct {
	DiskPoolName      pulumi.StringInput `pulumi:"diskPoolName"`
	IscsiTargetName   pulumi.StringInput `pulumi:"iscsiTargetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIscsiTargetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIscsiTargetArgs)(nil)).Elem()
}


type LookupIscsiTargetResultOutput struct{ *pulumi.OutputState }

func (LookupIscsiTargetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIscsiTargetResult)(nil)).Elem()
}

func (o LookupIscsiTargetResultOutput) ToLookupIscsiTargetResultOutput() LookupIscsiTargetResultOutput {
	return o
}

func (o LookupIscsiTargetResultOutput) ToLookupIscsiTargetResultOutputWithContext(ctx context.Context) LookupIscsiTargetResultOutput {
	return o
}

func (o LookupIscsiTargetResultOutput) AclMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) string { return v.AclMode }).(pulumi.StringOutput)
}

func (o LookupIscsiTargetResultOutput) Endpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) []string { return v.Endpoints }).(pulumi.StringArrayOutput)
}

func (o LookupIscsiTargetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIscsiTargetResultOutput) Luns() IscsiLunResponseArrayOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) []IscsiLunResponse { return v.Luns }).(IscsiLunResponseArrayOutput)
}

func (o LookupIscsiTargetResultOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) string { return v.ManagedBy }).(pulumi.StringOutput)
}

func (o LookupIscsiTargetResultOutput) ManagedByExtended() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) []string { return v.ManagedByExtended }).(pulumi.StringArrayOutput)
}

func (o LookupIscsiTargetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIscsiTargetResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o LookupIscsiTargetResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupIscsiTargetResultOutput) Sessions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) []string { return v.Sessions }).(pulumi.StringArrayOutput)
}

func (o LookupIscsiTargetResultOutput) StaticAcls() AclResponseArrayOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) []AclResponse { return v.StaticAcls }).(AclResponseArrayOutput)
}

func (o LookupIscsiTargetResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupIscsiTargetResultOutput) SystemData() SystemMetadataResponseOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) SystemMetadataResponse { return v.SystemData }).(SystemMetadataResponseOutput)
}

func (o LookupIscsiTargetResultOutput) TargetIqn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) string { return v.TargetIqn }).(pulumi.StringOutput)
}

func (o LookupIscsiTargetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIscsiTargetResultOutput{})
}
