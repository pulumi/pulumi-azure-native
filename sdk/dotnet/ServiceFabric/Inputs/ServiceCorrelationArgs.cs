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
    /// Creates a particular correlation between services.
    /// </summary>
    public sealed class ServiceCorrelationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ServiceCorrelationScheme which describes the relationship between this service and the service specified via ServiceName.
        /// </summary>
        [Input("scheme", required: true)]
        public InputUnion<string, Pulumi.AzureNative.ServiceFabric.ServiceCorrelationScheme> Scheme { get; set; } = null!;

        /// <summary>
        /// The Arm Resource ID of the service that the correlation relationship is established with.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public ServiceCorrelationArgs()
        {
        }
        public static new ServiceCorrelationArgs Empty => new ServiceCorrelationArgs();
    }
}
