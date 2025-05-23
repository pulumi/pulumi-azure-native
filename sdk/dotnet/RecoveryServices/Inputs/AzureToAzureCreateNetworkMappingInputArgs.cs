// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.Inputs
{

    /// <summary>
    /// Create network mappings input properties/behavior specific to Azure to Azure Network mapping.
    /// </summary>
    public sealed class AzureToAzureCreateNetworkMappingInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The instance type.
        /// Expected value is 'AzureToAzure'.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// The primary azure vnet Id.
        /// </summary>
        [Input("primaryNetworkId", required: true)]
        public Input<string> PrimaryNetworkId { get; set; } = null!;

        public AzureToAzureCreateNetworkMappingInputArgs()
        {
        }
        public static new AzureToAzureCreateNetworkMappingInputArgs Empty => new AzureToAzureCreateNetworkMappingInputArgs();
    }
}
