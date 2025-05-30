// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabric.Inputs
{

    /// <summary>
    /// Specifies a metric to load balance a service during runtime.
    /// </summary>
    public sealed class ServiceLoadMetricArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Used only for Stateless services. The default amount of load, as a number, that this service creates for this metric.
        /// </summary>
        [Input("defaultLoad")]
        public Input<int>? DefaultLoad { get; set; }

        /// <summary>
        /// The name of the metric. If the service chooses to report load during runtime, the load metric name should match the name that is specified in Name exactly. Note that metric names are case sensitive.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Used only for Stateful services. The default amount of load, as a number, that this service creates for this metric when it is a Primary replica.
        /// </summary>
        [Input("primaryDefaultLoad")]
        public Input<int>? PrimaryDefaultLoad { get; set; }

        /// <summary>
        /// Used only for Stateful services. The default amount of load, as a number, that this service creates for this metric when it is a Secondary replica.
        /// </summary>
        [Input("secondaryDefaultLoad")]
        public Input<int>? SecondaryDefaultLoad { get; set; }

        /// <summary>
        /// The service load metric relative weight, compared to other metrics configured for this service, as a number.
        /// </summary>
        [Input("weight")]
        public InputUnion<string, Pulumi.AzureNative.ServiceFabric.ServiceLoadMetricWeight>? Weight { get; set; }

        public ServiceLoadMetricArgs()
        {
        }
        public static new ServiceLoadMetricArgs Empty => new ServiceLoadMetricArgs();
    }
}
