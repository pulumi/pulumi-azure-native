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
    /// PrivateDnsZoneConfig resource.
    /// </summary>
    [OutputType]
    public sealed class PrivateDnsZoneConfigResponse
    {
        /// <summary>
        /// Name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The resource id of the private dns zone.
        /// </summary>
        public readonly string? PrivateDnsZoneId;
        /// <summary>
        /// A collection of information regarding a recordSet, holding information to identify private resources.
        /// </summary>
        public readonly ImmutableArray<Outputs.RecordSetResponse> RecordSets;

        [OutputConstructor]
        private PrivateDnsZoneConfigResponse(
            string? name,

            string? privateDnsZoneId,

            ImmutableArray<Outputs.RecordSetResponse> recordSets)
        {
            Name = name;
            PrivateDnsZoneId = privateDnsZoneId;
            RecordSets = recordSets;
        }
    }
}
