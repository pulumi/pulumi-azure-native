// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Outputs
{

    /// <summary>
    /// ExpressRouteLink child resource definition.
    /// </summary>
    [OutputType]
    public sealed class ExpressRouteLinkResponse
    {
        /// <summary>
        /// Administrative state of the physical port.
        /// </summary>
        public readonly string? AdminState;
        /// <summary>
        /// Cololocation for ExpressRoute Hybrid Direct.
        /// </summary>
        public readonly string ColoLocation;
        /// <summary>
        /// Physical fiber port type.
        /// </summary>
        public readonly string ConnectorType;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Name of Azure router interface.
        /// </summary>
        public readonly string InterfaceName;
        /// <summary>
        /// MacSec configuration.
        /// </summary>
        public readonly Outputs.ExpressRouteLinkMacSecConfigResponse? MacSecConfig;
        /// <summary>
        /// Name of child port resource that is unique among child port resources of the parent.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Mapping between physical port to patch panel port.
        /// </summary>
        public readonly string PatchPanelId;
        /// <summary>
        /// The provisioning state of the express route link resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Mapping of physical patch panel to rack.
        /// </summary>
        public readonly string RackId;
        /// <summary>
        /// Name of Azure router associated with physical port.
        /// </summary>
        public readonly string RouterName;

        [OutputConstructor]
        private ExpressRouteLinkResponse(
            string? adminState,

            string coloLocation,

            string connectorType,

            string etag,

            string? id,

            string interfaceName,

            Outputs.ExpressRouteLinkMacSecConfigResponse? macSecConfig,

            string? name,

            string patchPanelId,

            string provisioningState,

            string rackId,

            string routerName)
        {
            AdminState = adminState;
            ColoLocation = coloLocation;
            ConnectorType = connectorType;
            Etag = etag;
            Id = id;
            InterfaceName = interfaceName;
            MacSecConfig = macSecConfig;
            Name = name;
            PatchPanelId = patchPanelId;
            ProvisioningState = provisioningState;
            RackId = rackId;
            RouterName = routerName;
        }
    }
}
