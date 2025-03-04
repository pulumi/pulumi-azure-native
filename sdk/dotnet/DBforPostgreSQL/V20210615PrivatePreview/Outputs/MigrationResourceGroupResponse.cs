// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforPostgreSQL.V20210615PrivatePreview.Outputs
{

    /// <summary>
    /// Migration resource group.
    /// </summary>
    [OutputType]
    public sealed class MigrationResourceGroupResponse
    {
        public readonly string? ResourceId;
        public readonly string? SubnetResourceId;

        [OutputConstructor]
        private MigrationResourceGroupResponse(
            string? resourceId,

            string? subnetResourceId)
        {
            ResourceId = resourceId;
            SubnetResourceId = subnetResourceId;
        }
    }
}
