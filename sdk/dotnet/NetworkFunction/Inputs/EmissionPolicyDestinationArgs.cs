// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetworkFunction.Inputs
{

    /// <summary>
    /// Emission policy destination properties.
    /// </summary>
    public sealed class EmissionPolicyDestinationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Emission destination type.
        /// </summary>
        [Input("destinationType")]
        public InputUnion<string, Pulumi.AzureNative.NetworkFunction.DestinationType>? DestinationType { get; set; }

        public EmissionPolicyDestinationArgs()
        {
        }
        public static new EmissionPolicyDestinationArgs Empty => new EmissionPolicyDestinationArgs();
    }
}
