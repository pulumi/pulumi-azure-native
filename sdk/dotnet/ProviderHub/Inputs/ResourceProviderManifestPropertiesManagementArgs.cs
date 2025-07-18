// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Inputs
{

    /// <summary>
    /// The resource provider management.
    /// </summary>
    public sealed class ResourceProviderManifestPropertiesManagementArgs : global::Pulumi.ResourceArgs
    {
        [Input("authorizationOwners")]
        private InputList<string>? _authorizationOwners;

        /// <summary>
        /// The authorization owners.
        /// </summary>
        public InputList<string> AuthorizationOwners
        {
            get => _authorizationOwners ?? (_authorizationOwners = new InputList<string>());
            set => _authorizationOwners = value;
        }

        [Input("canaryManifestOwners")]
        private InputList<string>? _canaryManifestOwners;

        /// <summary>
        /// List of manifest owners for canary.
        /// </summary>
        public InputList<string> CanaryManifestOwners
        {
            get => _canaryManifestOwners ?? (_canaryManifestOwners = new InputList<string>());
            set => _canaryManifestOwners = value;
        }

        /// <summary>
        /// Options for error response messages.
        /// </summary>
        [Input("errorResponseMessageOptions")]
        public Input<Inputs.ResourceProviderManagementErrorResponseMessageOptionsArgs>? ErrorResponseMessageOptions { get; set; }

        /// <summary>
        /// Metadata for expedited rollout.
        /// </summary>
        [Input("expeditedRolloutMetadata")]
        public Input<Inputs.ResourceProviderManagementExpeditedRolloutMetadataArgs>? ExpeditedRolloutMetadata { get; set; }

        [Input("expeditedRolloutSubmitters")]
        private InputList<string>? _expeditedRolloutSubmitters;

        /// <summary>
        /// List of expedited rollout submitters.
        /// </summary>
        public InputList<string> ExpeditedRolloutSubmitters
        {
            get => _expeditedRolloutSubmitters ?? (_expeditedRolloutSubmitters = new InputList<string>());
            set => _expeditedRolloutSubmitters = value;
        }

        /// <summary>
        /// The incident contact email.
        /// </summary>
        [Input("incidentContactEmail")]
        public Input<string>? IncidentContactEmail { get; set; }

        /// <summary>
        /// The incident routing service.
        /// </summary>
        [Input("incidentRoutingService")]
        public Input<string>? IncidentRoutingService { get; set; }

        /// <summary>
        /// The incident routing team.
        /// </summary>
        [Input("incidentRoutingTeam")]
        public Input<string>? IncidentRoutingTeam { get; set; }

        [Input("manifestOwners")]
        private InputList<string>? _manifestOwners;

        /// <summary>
        /// The manifest owners.
        /// </summary>
        public InputList<string> ManifestOwners
        {
            get => _manifestOwners ?? (_manifestOwners = new InputList<string>());
            set => _manifestOwners = value;
        }

        /// <summary>
        /// The profit center code for the subscription.
        /// </summary>
        [Input("pcCode")]
        public Input<string>? PcCode { get; set; }

        /// <summary>
        /// The profit center program id for the subscription.
        /// </summary>
        [Input("profitCenterProgramId")]
        public Input<string>? ProfitCenterProgramId { get; set; }

        /// <summary>
        /// The resource access policy.
        /// </summary>
        [Input("resourceAccessPolicy")]
        public InputUnion<string, Pulumi.AzureNative.ProviderHub.ResourceAccessPolicy>? ResourceAccessPolicy { get; set; }

        [Input("resourceAccessRoles")]
        private InputList<Inputs.ResourceAccessRoleArgs>? _resourceAccessRoles;

        /// <summary>
        /// The resource access roles.
        /// </summary>
        public InputList<Inputs.ResourceAccessRoleArgs> ResourceAccessRoles
        {
            get => _resourceAccessRoles ?? (_resourceAccessRoles = new InputList<Inputs.ResourceAccessRoleArgs>());
            set => _resourceAccessRoles = value;
        }

        [Input("schemaOwners")]
        private InputList<string>? _schemaOwners;

        /// <summary>
        /// The schema owners.
        /// </summary>
        public InputList<string> SchemaOwners
        {
            get => _schemaOwners ?? (_schemaOwners = new InputList<string>());
            set => _schemaOwners = value;
        }

        [Input("serviceTreeInfos")]
        private InputList<Inputs.ServiceTreeInfoArgs>? _serviceTreeInfos;

        /// <summary>
        /// The service tree infos.
        /// </summary>
        public InputList<Inputs.ServiceTreeInfoArgs> ServiceTreeInfos
        {
            get => _serviceTreeInfos ?? (_serviceTreeInfos = new InputList<Inputs.ServiceTreeInfoArgs>());
            set => _serviceTreeInfos = value;
        }

        public ResourceProviderManifestPropertiesManagementArgs()
        {
        }
        public static new ResourceProviderManifestPropertiesManagementArgs Empty => new ResourceProviderManifestPropertiesManagementArgs();
    }
}
