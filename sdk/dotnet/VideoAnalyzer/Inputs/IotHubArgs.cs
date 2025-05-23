// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VideoAnalyzer.Inputs
{

    /// <summary>
    /// The IoT Hub details.
    /// </summary>
    public sealed class IotHubArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IoT Hub resource identifier.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The IoT Hub identity.
        /// </summary>
        [Input("identity", required: true)]
        public Input<Inputs.ResourceIdentityArgs> Identity { get; set; } = null!;

        public IotHubArgs()
        {
        }
        public static new IotHubArgs Empty => new IotHubArgs();
    }
}
