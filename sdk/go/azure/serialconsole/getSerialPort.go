


package serialconsole

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSerialPort(ctx *pulumi.Context, args *LookupSerialPortArgs, opts ...pulumi.InvokeOption) (*LookupSerialPortResult, error) {
	var rv LookupSerialPortResult
	err := ctx.Invoke("azure-native:serialconsole:getSerialPort", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSerialPortArgs struct {
	ParentResource            string `pulumi:"parentResource"`
	ParentResourceType        string `pulumi:"parentResourceType"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	ResourceProviderNamespace string `pulumi:"resourceProviderNamespace"`
	SerialPort                string `pulumi:"serialPort"`
}


type LookupSerialPortResult struct {
	Id    string  `pulumi:"id"`
	Name  string  `pulumi:"name"`
	State *string `pulumi:"state"`
	Type  string  `pulumi:"type"`
}
