


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSensor(ctx *pulumi.Context, args *LookupSensorArgs, opts ...pulumi.InvokeOption) (*LookupSensorResult, error) {
	var rv LookupSensorResult
	err := ctx.Invoke("azure-native:iotsecurity/v20210201preview:getSensor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSensorArgs struct {
	Scope      string `pulumi:"scope"`
	SensorName string `pulumi:"sensorName"`
}


type LookupSensorResult struct {
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
