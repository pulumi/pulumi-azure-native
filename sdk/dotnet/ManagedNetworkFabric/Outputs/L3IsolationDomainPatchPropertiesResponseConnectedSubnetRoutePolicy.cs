// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedNetworkFabric.Outputs
{

    /// <summary>
    /// Connected Subnet RoutePolicy
    /// </summary>
    [OutputType]
    public sealed class L3IsolationDomainPatchPropertiesResponseConnectedSubnetRoutePolicy
    {
        /// <summary>
        /// Enabled/Disabled connected subnet route policy. Ex: Enabled | Disabled.
        /// </summary>
        public readonly string AdministrativeState;
        /// <summary>
        /// exportRoutePolicyId value.
        /// </summary>
        public readonly string? ExportRoutePolicyId;

        [OutputConstructor]
        private L3IsolationDomainPatchPropertiesResponseConnectedSubnetRoutePolicy(
            string administrativeState,

            string? exportRoutePolicyId)
        {
            AdministrativeState = administrativeState;
            ExportRoutePolicyId = exportRoutePolicyId;
        }
    }
}
