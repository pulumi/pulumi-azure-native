


package v20210404preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSpacecraft(ctx *pulumi.Context, args *LookupSpacecraftArgs, opts ...pulumi.InvokeOption) (*LookupSpacecraftResult, error) {
	var rv LookupSpacecraftResult
	err := ctx.Invoke("azure-native:orbital/v20210404preview:getSpacecraft", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSpacecraftArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SpacecraftName    string `pulumi:"spacecraftName"`
}


type LookupSpacecraftResult struct {
	AuthorizationStatus         string                   `pulumi:"authorizationStatus"`
	AuthorizationStatusExtended string                   `pulumi:"authorizationStatusExtended"`
	Etag                        string                   `pulumi:"etag"`
	Id                          string                   `pulumi:"id"`
	Links                       []SpacecraftLinkResponse `pulumi:"links"`
	Location                    string                   `pulumi:"location"`
	Name                        string                   `pulumi:"name"`
	NoradId                     string                   `pulumi:"noradId"`
	SystemData                  SystemDataResponse       `pulumi:"systemData"`
	Tags                        map[string]string        `pulumi:"tags"`
	TitleLine                   *string                  `pulumi:"titleLine"`
	TleLine1                    *string                  `pulumi:"tleLine1"`
	TleLine2                    *string                  `pulumi:"tleLine2"`
	Type                        string                   `pulumi:"type"`
}

func LookupSpacecraftOutput(ctx *pulumi.Context, args LookupSpacecraftOutputArgs, opts ...pulumi.InvokeOption) LookupSpacecraftResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSpacecraftResult, error) {
			args := v.(LookupSpacecraftArgs)
			r, err := LookupSpacecraft(ctx, &args, opts...)
			var s LookupSpacecraftResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSpacecraftResultOutput)
}

type LookupSpacecraftOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SpacecraftName    pulumi.StringInput `pulumi:"spacecraftName"`
}

func (LookupSpacecraftOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSpacecraftArgs)(nil)).Elem()
}


type LookupSpacecraftResultOutput struct{ *pulumi.OutputState }

func (LookupSpacecraftResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSpacecraftResult)(nil)).Elem()
}

func (o LookupSpacecraftResultOutput) ToLookupSpacecraftResultOutput() LookupSpacecraftResultOutput {
	return o
}

func (o LookupSpacecraftResultOutput) ToLookupSpacecraftResultOutputWithContext(ctx context.Context) LookupSpacecraftResultOutput {
	return o
}

func (o LookupSpacecraftResultOutput) AuthorizationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpacecraftResult) string { return v.AuthorizationStatus }).(pulumi.StringOutput)
}

func (o LookupSpacecraftResultOutput) AuthorizationStatusExtended() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpacecraftResult) string { return v.AuthorizationStatusExtended }).(pulumi.StringOutput)
}

func (o LookupSpacecraftResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpacecraftResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupSpacecraftResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpacecraftResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSpacecraftResultOutput) Links() SpacecraftLinkResponseArrayOutput {
	return o.ApplyT(func(v LookupSpacecraftResult) []SpacecraftLinkResponse { return v.Links }).(SpacecraftLinkResponseArrayOutput)
}

func (o LookupSpacecraftResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpacecraftResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSpacecraftResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpacecraftResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSpacecraftResultOutput) NoradId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpacecraftResult) string { return v.NoradId }).(pulumi.StringOutput)
}

func (o LookupSpacecraftResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSpacecraftResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSpacecraftResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSpacecraftResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSpacecraftResultOutput) TitleLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSpacecraftResult) *string { return v.TitleLine }).(pulumi.StringPtrOutput)
}

func (o LookupSpacecraftResultOutput) TleLine1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSpacecraftResult) *string { return v.TleLine1 }).(pulumi.StringPtrOutput)
}

func (o LookupSpacecraftResultOutput) TleLine2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSpacecraftResult) *string { return v.TleLine2 }).(pulumi.StringPtrOutput)
}

func (o LookupSpacecraftResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpacecraftResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSpacecraftResultOutput{})
}
