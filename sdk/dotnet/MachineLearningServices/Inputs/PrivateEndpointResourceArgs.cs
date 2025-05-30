// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    /// <summary>
    /// The PE network resource that is linked to this PE connection.
    /// </summary>
    public sealed class PrivateEndpointResourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The subnetId that the private endpoint is connected to.
        /// </summary>
        [Input("subnetArmId")]
        public Input<string>? SubnetArmId { get; set; }

        public PrivateEndpointResourceArgs()
        {
        }
        public static new PrivateEndpointResourceArgs Empty => new PrivateEndpointResourceArgs();
    }
}
