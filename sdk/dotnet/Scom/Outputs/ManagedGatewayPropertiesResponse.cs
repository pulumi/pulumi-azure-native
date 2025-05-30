// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Scom.Outputs
{

    /// <summary>
    /// The properties of a gateway resource.
    /// </summary>
    [OutputType]
    public sealed class ManagedGatewayPropertiesResponse
    {
        /// <summary>
        /// ComputerName of the gateway to be monitored.
        /// </summary>
        public readonly string? ComputerName;
        /// <summary>
        /// The connection status of the gateway resource.
        /// </summary>
        public readonly string ConnectionStatus;
        /// <summary>
        /// The domain name associated with the gateway to be monitored.
        /// </summary>
        public readonly string? DomainName;
        /// <summary>
        /// The health status of the gateway resource.
        /// </summary>
        public readonly string HealthStatus;
        /// <summary>
        /// Install type of gateway resource.
        /// </summary>
        public readonly string InstallType;
        /// <summary>
        /// The management server endpoint to which the gateway is directed.
        /// </summary>
        public readonly string ManagementServerEndpoint;
        public readonly string ProvisioningState;
        /// <summary>
        /// ArmId of the gateway to be monitored.
        /// </summary>
        public readonly string? ResourceId;
        /// <summary>
        /// Location of the gateway to be monitored.
        /// </summary>
        public readonly string? ResourceLocation;
        /// <summary>
        /// The version of the gateway resource.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private ManagedGatewayPropertiesResponse(
            string? computerName,

            string connectionStatus,

            string? domainName,

            string healthStatus,

            string installType,

            string managementServerEndpoint,

            string provisioningState,

            string? resourceId,

            string? resourceLocation,

            string version)
        {
            ComputerName = computerName;
            ConnectionStatus = connectionStatus;
            DomainName = domainName;
            HealthStatus = healthStatus;
            InstallType = installType;
            ManagementServerEndpoint = managementServerEndpoint;
            ProvisioningState = provisioningState;
            ResourceId = resourceId;
            ResourceLocation = resourceLocation;
            Version = version;
        }
    }
}
