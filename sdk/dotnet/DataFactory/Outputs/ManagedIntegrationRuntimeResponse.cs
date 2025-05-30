// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Outputs
{

    /// <summary>
    /// Managed integration runtime, including managed elastic and managed dedicated integration runtimes.
    /// </summary>
    [OutputType]
    public sealed class ManagedIntegrationRuntimeResponse
    {
        /// <summary>
        /// The compute resource for managed integration runtime.
        /// </summary>
        public readonly Outputs.IntegrationRuntimeComputePropertiesResponse? ComputeProperties;
        /// <summary>
        /// The name of virtual network to which Azure-SSIS integration runtime will join
        /// </summary>
        public readonly Outputs.IntegrationRuntimeCustomerVirtualNetworkResponse? CustomerVirtualNetwork;
        /// <summary>
        /// Integration runtime description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Managed Virtual Network reference.
        /// </summary>
        public readonly Outputs.ManagedVirtualNetworkReferenceResponse? ManagedVirtualNetwork;
        /// <summary>
        /// SSIS properties for managed integration runtime.
        /// </summary>
        public readonly Outputs.IntegrationRuntimeSsisPropertiesResponse? SsisProperties;
        /// <summary>
        /// Integration runtime state, only valid for managed dedicated integration runtime.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The type of integration runtime.
        /// Expected value is 'Managed'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ManagedIntegrationRuntimeResponse(
            Outputs.IntegrationRuntimeComputePropertiesResponse? computeProperties,

            Outputs.IntegrationRuntimeCustomerVirtualNetworkResponse? customerVirtualNetwork,

            string? description,

            Outputs.ManagedVirtualNetworkReferenceResponse? managedVirtualNetwork,

            Outputs.IntegrationRuntimeSsisPropertiesResponse? ssisProperties,

            string state,

            string type)
        {
            ComputeProperties = computeProperties;
            CustomerVirtualNetwork = customerVirtualNetwork;
            Description = description;
            ManagedVirtualNetwork = managedVirtualNetwork;
            SsisProperties = ssisProperties;
            State = state;
            Type = type;
        }
    }
}
