package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := app.NewAppResiliency(ctx, "appResiliency", &app.AppResiliencyArgs{
			AppName: pulumi.String("testcontainerApp0"),
			CircuitBreakerPolicy: &app.CircuitBreakerPolicyArgs{
				ConsecutiveErrors:  pulumi.Int(5),
				IntervalInSeconds:  pulumi.Int(10),
				MaxEjectionPercent: pulumi.Int(50),
			},
			HttpConnectionPool: &app.HttpConnectionPoolArgs{
				Http1MaxPendingRequests: pulumi.Int(1024),
				Http2MaxRequests:        pulumi.Int(1024),
			},
			HttpRetryPolicy: app.HttpRetryPolicyResponse{
				Errors: pulumi.StringArray{
					pulumi.String("5xx"),
					pulumi.String("connect-failure"),
					pulumi.String("reset"),
					pulumi.String("retriable-headers"),
					pulumi.String("retriable-status-codes"),
				},
				Headers: app.HeaderMatchArray{
					&app.HeaderMatchArgs{
						Header:      pulumi.String("X-Content-Type"),
						PrefixMatch: pulumi.String("GOATS"),
					},
				},
				HttpStatusCodes: pulumi.IntArray{
					pulumi.Int(502),
					pulumi.Int(503),
				},
				InitialDelayInMilliseconds: pulumi.Float64(1000),
				MaxIntervalInMilliseconds:  pulumi.Float64(10000),
				MaxRetries:                 pulumi.Int(5),
			},
			Name:              pulumi.String("resiliency-policy-1"),
			ResourceGroupName: pulumi.String("rg"),
			TcpConnectionPool: &app.TcpConnectionPoolArgs{
				MaxConnections: pulumi.Int(100),
			},
			TcpRetryPolicy: &app.TcpRetryPolicyArgs{
				MaxConnectAttempts: pulumi.Int(3),
			},
			TimeoutPolicy: &app.TimeoutPolicyArgs{
				ConnectionTimeoutInSeconds: pulumi.Int(5),
				ResponseTimeoutInSeconds:   pulumi.Int(15),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
