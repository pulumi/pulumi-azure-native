


package storsimple

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccessControlRecord(ctx *pulumi.Context, args *LookupAccessControlRecordArgs, opts ...pulumi.InvokeOption) (*LookupAccessControlRecordResult, error) {
	var rv LookupAccessControlRecordResult
	err := ctx.Invoke("azure-native:storsimple:getAccessControlRecord", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessControlRecordArgs struct {
	AccessControlRecordName string `pulumi:"accessControlRecordName"`
	ManagerName             string `pulumi:"managerName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupAccessControlRecordResult struct {
	Id            string  `pulumi:"id"`
	InitiatorName string  `pulumi:"initiatorName"`
	Kind          *string `pulumi:"kind"`
	Name          string  `pulumi:"name"`
	Type          string  `pulumi:"type"`
	VolumeCount   int     `pulumi:"volumeCount"`
}
