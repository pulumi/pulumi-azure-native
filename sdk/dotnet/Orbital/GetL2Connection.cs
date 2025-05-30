// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Orbital
{
    public static class GetL2Connection
    {
        /// <summary>
        /// Gets the specified L2 connection in a specified resource group.
        /// 
        /// Uses Azure REST API version 2024-03-01-preview.
        /// 
        /// Other available API versions: 2024-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native orbital [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetL2ConnectionResult> InvokeAsync(GetL2ConnectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetL2ConnectionResult>("azure-native:orbital:getL2Connection", args ?? new GetL2ConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified L2 connection in a specified resource group.
        /// 
        /// Uses Azure REST API version 2024-03-01-preview.
        /// 
        /// Other available API versions: 2024-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native orbital [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetL2ConnectionResult> Invoke(GetL2ConnectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetL2ConnectionResult>("azure-native:orbital:getL2Connection", args ?? new GetL2ConnectionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified L2 connection in a specified resource group.
        /// 
        /// Uses Azure REST API version 2024-03-01-preview.
        /// 
        /// Other available API versions: 2024-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native orbital [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetL2ConnectionResult> Invoke(GetL2ConnectionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetL2ConnectionResult>("azure-native:orbital:getL2Connection", args ?? new GetL2ConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetL2ConnectionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// L2 Connection name.
        /// </summary>
        [Input("l2ConnectionName", required: true)]
        public string L2ConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetL2ConnectionArgs()
        {
        }
        public static new GetL2ConnectionArgs Empty => new GetL2ConnectionArgs();
    }

    public sealed class GetL2ConnectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// L2 Connection name.
        /// </summary>
        [Input("l2ConnectionName", required: true)]
        public Input<string> L2ConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetL2ConnectionInvokeArgs()
        {
        }
        public static new GetL2ConnectionInvokeArgs Empty => new GetL2ConnectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetL2ConnectionResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Globally-unique identifier for this connection that is to be used as a circuit ID.
        /// </summary>
        public readonly string CircuitId;
        /// <summary>
        /// A reference to an Microsoft.Orbital/edgeSites resource to route traffic for.
        /// </summary>
        public readonly Outputs.L2ConnectionsPropertiesResponseEdgeSite EdgeSite;
        /// <summary>
        /// A reference to an Microsoft.Orbital/groundStations resource to route traffic for.
        /// </summary>
        public readonly Outputs.L2ConnectionsPropertiesResponseGroundStation GroundStation;
        /// <summary>
        /// The name of the partner router to establish a connection to within the ground station.
        /// </summary>
        public readonly Outputs.L2ConnectionsPropertiesResponseGroundStationPartnerRouter GroundStationPartnerRouter;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The VLAN ID for the L2 connection.
        /// </summary>
        public readonly int VlanId;

        [OutputConstructor]
        private GetL2ConnectionResult(
            string azureApiVersion,

            string circuitId,

            Outputs.L2ConnectionsPropertiesResponseEdgeSite edgeSite,

            Outputs.L2ConnectionsPropertiesResponseGroundStation groundStation,

            Outputs.L2ConnectionsPropertiesResponseGroundStationPartnerRouter groundStationPartnerRouter,

            string id,

            string location,

            string name,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            int vlanId)
        {
            AzureApiVersion = azureApiVersion;
            CircuitId = circuitId;
            EdgeSite = edgeSite;
            GroundStation = groundStation;
            GroundStationPartnerRouter = groundStationPartnerRouter;
            Id = id;
            Location = location;
            Name = name;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            VlanId = vlanId;
        }
    }
}
