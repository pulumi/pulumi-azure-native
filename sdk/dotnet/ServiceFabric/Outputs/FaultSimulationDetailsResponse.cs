// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabric.Outputs
{

    /// <summary>
    /// Details for Fault Simulation.
    /// </summary>
    [OutputType]
    public sealed class FaultSimulationDetailsResponse
    {
        /// <summary>
        /// unique identifier for the cluster resource.
        /// </summary>
        public readonly string? ClusterId;
        /// <summary>
        /// List of node type simulations associated with the cluster fault simulation.
        /// </summary>
        public readonly ImmutableArray<Outputs.NodeTypeFaultSimulationResponse> NodeTypeFaultSimulation;
        /// <summary>
        /// unique identifier for the operation associated with the fault simulation.
        /// </summary>
        public readonly string? OperationId;
        /// <summary>
        /// Fault simulation parameters.
        /// </summary>
        public readonly Outputs.ZoneFaultSimulationContentResponse? Parameters;

        [OutputConstructor]
        private FaultSimulationDetailsResponse(
            string? clusterId,

            ImmutableArray<Outputs.NodeTypeFaultSimulationResponse> nodeTypeFaultSimulation,

            string? operationId,

            Outputs.ZoneFaultSimulationContentResponse? parameters)
        {
            ClusterId = clusterId;
            NodeTypeFaultSimulation = nodeTypeFaultSimulation;
            OperationId = operationId;
            Parameters = parameters;
        }
    }
}
