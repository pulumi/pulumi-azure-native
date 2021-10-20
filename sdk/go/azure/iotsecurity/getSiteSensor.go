


package iotsecurity

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSiteSensor(ctx *pulumi.Context, args *LookupSiteSensorArgs, opts ...pulumi.InvokeOption) (*LookupSiteSensorResult, error) {
	var rv LookupSiteSensorResult
	err := ctx.Invoke("azure-native:iotsecurity:getSiteSensor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteSensorArgs struct {
	IotDefenderLocation string `pulumi:"iotDefenderLocation"`
	SensorName          string `pulumi:"sensorName"`
	SiteName            string `pulumi:"siteName"`
}


type LookupSiteSensorResult struct {
	ConnectivityTime   string             `pulumi:"connectivityTime"`
	DynamicLearning    bool               `pulumi:"dynamicLearning"`
	Id                 string             `pulumi:"id"`
	LearningMode       bool               `pulumi:"learningMode"`
	Name               string             `pulumi:"name"`
	SensorStatus       string             `pulumi:"sensorStatus"`
	SensorType         *string            `pulumi:"sensorType"`
	SensorVersion      string             `pulumi:"sensorVersion"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	TiAutomaticUpdates *bool              `pulumi:"tiAutomaticUpdates"`
	TiStatus           string             `pulumi:"tiStatus"`
	TiVersion          string             `pulumi:"tiVersion"`
	Type               string             `pulumi:"type"`
	Zone               *string            `pulumi:"zone"`
}
