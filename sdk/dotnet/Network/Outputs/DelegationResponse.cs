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
    /// Details the service to which the subnet is delegated.
    /// </summary>
    [OutputType]
    public sealed class DelegationResponse
    {
        /// <summary>
        /// The actions permitted to the service upon delegation.
        /// </summary>
        public readonly ImmutableArray<string> Actions;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the resource that is unique within a subnet. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The provisioning state of the service delegation resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The name of the service to whom the subnet should be delegated (e.g. Microsoft.Sql/servers).
        /// </summary>
        public readonly string? ServiceName;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private DelegationResponse(
            ImmutableArray<string> actions,

            string etag,

            string? id,

            string? name,

            string provisioningState,

            string? serviceName,

            string? type)
        {
            Actions = actions;
            Etag = etag;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            ServiceName = serviceName;
            Type = type;
        }
    }
}
