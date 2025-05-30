// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Outputs
{

    [OutputType]
    public sealed class ResourceTypeRegistrationPropertiesResponseManagement
    {
        public readonly string? IncidentContactEmail;
        public readonly string? IncidentRoutingService;
        public readonly string? IncidentRoutingTeam;
        public readonly ImmutableArray<string> ManifestOwners;
        public readonly string? ResourceAccessPolicy;
        public readonly ImmutableArray<object> ResourceAccessRoles;
        public readonly ImmutableArray<string> SchemaOwners;
        public readonly ImmutableArray<Outputs.ServiceTreeInfoResponse> ServiceTreeInfos;

        [OutputConstructor]
        private ResourceTypeRegistrationPropertiesResponseManagement(
            string? incidentContactEmail,

            string? incidentRoutingService,

            string? incidentRoutingTeam,

            ImmutableArray<string> manifestOwners,

            string? resourceAccessPolicy,

            ImmutableArray<object> resourceAccessRoles,

            ImmutableArray<string> schemaOwners,

            ImmutableArray<Outputs.ServiceTreeInfoResponse> serviceTreeInfos)
        {
            IncidentContactEmail = incidentContactEmail;
            IncidentRoutingService = incidentRoutingService;
            IncidentRoutingTeam = incidentRoutingTeam;
            ManifestOwners = manifestOwners;
            ResourceAccessPolicy = resourceAccessPolicy;
            ResourceAccessRoles = resourceAccessRoles;
            SchemaOwners = schemaOwners;
            ServiceTreeInfos = serviceTreeInfos;
        }
    }
}
