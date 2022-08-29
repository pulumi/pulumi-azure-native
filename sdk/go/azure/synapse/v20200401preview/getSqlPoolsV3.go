


package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSqlPoolsV3(ctx *pulumi.Context, args *LookupSqlPoolsV3Args, opts ...pulumi.InvokeOption) (*LookupSqlPoolsV3Result, error) {
	var rv LookupSqlPoolsV3Result
	err := ctx.Invoke("azure-native:synapse/v20200401preview:getSqlPoolsV3", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlPoolsV3Args struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SqlPoolName       string `pulumi:"sqlPoolName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupSqlPoolsV3Result struct {
	AutoPauseTimer                *int               `pulumi:"autoPauseTimer"`
	AutoResume                    *bool              `pulumi:"autoResume"`
	CurrentServiceObjectiveName   string             `pulumi:"currentServiceObjectiveName"`
	Id                            string             `pulumi:"id"`
	Kind                          string             `pulumi:"kind"`
	Location                      string             `pulumi:"location"`
	MaxServiceObjectiveName       *string            `pulumi:"maxServiceObjectiveName"`
	Name                          string             `pulumi:"name"`
	RequestedServiceObjectiveName string             `pulumi:"requestedServiceObjectiveName"`
	Sku                           *SkuV3Response     `pulumi:"sku"`
	SqlPoolGuid                   string             `pulumi:"sqlPoolGuid"`
	Status                        string             `pulumi:"status"`
	SystemData                    SystemDataResponse `pulumi:"systemData"`
	Tags                          map[string]string  `pulumi:"tags"`
	Type                          string             `pulumi:"type"`
}

func LookupSqlPoolsV3Output(ctx *pulumi.Context, args LookupSqlPoolsV3OutputArgs, opts ...pulumi.InvokeOption) LookupSqlPoolsV3ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlPoolsV3Result, error) {
			args := v.(LookupSqlPoolsV3Args)
			r, err := LookupSqlPoolsV3(ctx, &args, opts...)
			var s LookupSqlPoolsV3Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlPoolsV3ResultOutput)
}

type LookupSqlPoolsV3OutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SqlPoolName       pulumi.StringInput `pulumi:"sqlPoolName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupSqlPoolsV3OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolsV3Args)(nil)).Elem()
}


type LookupSqlPoolsV3ResultOutput struct{ *pulumi.OutputState }

func (LookupSqlPoolsV3ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolsV3Result)(nil)).Elem()
}

func (o LookupSqlPoolsV3ResultOutput) ToLookupSqlPoolsV3ResultOutput() LookupSqlPoolsV3ResultOutput {
	return o
}

func (o LookupSqlPoolsV3ResultOutput) ToLookupSqlPoolsV3ResultOutputWithContext(ctx context.Context) LookupSqlPoolsV3ResultOutput {
	return o
}

func (o LookupSqlPoolsV3ResultOutput) AutoPauseTimer() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolsV3Result) *int { return v.AutoPauseTimer }).(pulumi.IntPtrOutput)
}

func (o LookupSqlPoolsV3ResultOutput) AutoResume() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolsV3Result) *bool { return v.AutoResume }).(pulumi.BoolPtrOutput)
}

func (o LookupSqlPoolsV3ResultOutput) CurrentServiceObjectiveName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolsV3Result) string { return v.CurrentServiceObjectiveName }).(pulumi.StringOutput)
}

func (o LookupSqlPoolsV3ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolsV3Result) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlPoolsV3ResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolsV3Result) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupSqlPoolsV3ResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolsV3Result) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSqlPoolsV3ResultOutput) MaxServiceObjectiveName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolsV3Result) *string { return v.MaxServiceObjectiveName }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolsV3ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolsV3Result) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlPoolsV3ResultOutput) RequestedServiceObjectiveName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolsV3Result) string { return v.RequestedServiceObjectiveName }).(pulumi.StringOutput)
}

func (o LookupSqlPoolsV3ResultOutput) Sku() SkuV3ResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlPoolsV3Result) *SkuV3Response { return v.Sku }).(SkuV3ResponsePtrOutput)
}

func (o LookupSqlPoolsV3ResultOutput) SqlPoolGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolsV3Result) string { return v.SqlPoolGuid }).(pulumi.StringOutput)
}

func (o LookupSqlPoolsV3ResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolsV3Result) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupSqlPoolsV3ResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSqlPoolsV3Result) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSqlPoolsV3ResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlPoolsV3Result) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSqlPoolsV3ResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolsV3Result) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlPoolsV3ResultOutput{})
}
