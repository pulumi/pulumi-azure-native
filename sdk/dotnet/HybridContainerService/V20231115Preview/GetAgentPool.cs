// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridContainerService.V20231115Preview
{
    public static class GetAgentPool
    {
        /// <summary>
        /// Gets the agent pool in the Hybrid AKS provisioned cluster instance
        /// </summary>
        public static Task<GetAgentPoolResult> InvokeAsync(GetAgentPoolArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAgentPoolResult>("azure-native:hybridcontainerservice/v20231115preview:getAgentPool", args ?? new GetAgentPoolArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the agent pool in the Hybrid AKS provisioned cluster instance
        /// </summary>
        public static Output<GetAgentPoolResult> Invoke(GetAgentPoolInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAgentPoolResult>("azure-native:hybridcontainerservice/v20231115preview:getAgentPool", args ?? new GetAgentPoolInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the agent pool in the Hybrid AKS provisioned cluster instance
        /// </summary>
        public static Output<GetAgentPoolResult> Invoke(GetAgentPoolInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAgentPoolResult>("azure-native:hybridcontainerservice/v20231115preview:getAgentPool", args ?? new GetAgentPoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAgentPoolArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Parameter for the name of the agent pool in the provisioned cluster
        /// </summary>
        [Input("agentPoolName", required: true)]
        public string AgentPoolName { get; set; } = null!;

        /// <summary>
        /// The fully qualified Azure Resource manager identifier of the connected cluster resource.
        /// </summary>
        [Input("connectedClusterResourceUri", required: true)]
        public string ConnectedClusterResourceUri { get; set; } = null!;

        public GetAgentPoolArgs()
        {
        }
        public static new GetAgentPoolArgs Empty => new GetAgentPoolArgs();
    }

    public sealed class GetAgentPoolInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Parameter for the name of the agent pool in the provisioned cluster
        /// </summary>
        [Input("agentPoolName", required: true)]
        public Input<string> AgentPoolName { get; set; } = null!;

        /// <summary>
        /// The fully qualified Azure Resource manager identifier of the connected cluster resource.
        /// </summary>
        [Input("connectedClusterResourceUri", required: true)]
        public Input<string> ConnectedClusterResourceUri { get; set; } = null!;

        public GetAgentPoolInvokeArgs()
        {
        }
        public static new GetAgentPoolInvokeArgs Empty => new GetAgentPoolInvokeArgs();
    }


    [OutputType]
    public sealed class GetAgentPoolResult
    {
        /// <summary>
        /// AvailabilityZones - The list of Availability zones to use for nodes. Datacenter racks modelled as zones
        /// </summary>
        public readonly ImmutableArray<string> AvailabilityZones;
        /// <summary>
        /// Count - Number of agents to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
        /// </summary>
        public readonly int? Count;
        /// <summary>
        /// Extended Location definition
        /// </summary>
        public readonly Outputs.ExtendedLocationResponse? ExtendedLocation;
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The resource location
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource Name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The version of node image
        /// </summary>
        public readonly string? NodeImageVersion;
        /// <summary>
        /// Specifies the OS SKU used by the agent pool. The default is CBLMariner if OSType is Linux. The default is Windows2019 when OSType is Windows.
        /// </summary>
        public readonly string? OsSKU;
        /// <summary>
        /// OsType - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux. Possible values include: 'Linux', 'Windows'
        /// </summary>
        public readonly string? OsType;
        /// <summary>
        /// Provisioning state of the resource
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Defines the observed state of the agent pool
        /// </summary>
        public readonly Outputs.AgentPoolProvisioningStatusResponseStatus? Status;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource Type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// VmSize - The size of the agent pool VMs.
        /// </summary>
        public readonly string? VmSize;

        [OutputConstructor]
        private GetAgentPoolResult(
            ImmutableArray<string> availabilityZones,

            int? count,

            Outputs.ExtendedLocationResponse? extendedLocation,

            string id,

            string? location,

            string name,

            string? nodeImageVersion,

            string? osSKU,

            string? osType,

            string provisioningState,

            Outputs.AgentPoolProvisioningStatusResponseStatus? status,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            string? vmSize)
        {
            AvailabilityZones = availabilityZones;
            Count = count;
            ExtendedLocation = extendedLocation;
            Id = id;
            Location = location;
            Name = name;
            NodeImageVersion = nodeImageVersion;
            OsSKU = osSKU;
            OsType = osType;
            ProvisioningState = provisioningState;
            Status = status;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            VmSize = vmSize;
        }
    }
}
