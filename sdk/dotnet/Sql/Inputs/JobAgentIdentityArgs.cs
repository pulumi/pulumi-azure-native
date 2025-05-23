// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql.Inputs
{

    /// <summary>
    /// Azure Active Directory identity configuration for a resource.
    /// </summary>
    public sealed class JobAgentIdentityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The job agent identity tenant id
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The job agent identity type
        /// </summary>
        [Input("type", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Sql.JobAgentIdentityType> Type { get; set; } = null!;

        [Input("userAssignedIdentities")]
        private InputList<string>? _userAssignedIdentities;

        /// <summary>
        /// The resource ids of the user assigned identities to use
        /// </summary>
        public InputList<string> UserAssignedIdentities
        {
            get => _userAssignedIdentities ?? (_userAssignedIdentities = new InputList<string>());
            set => _userAssignedIdentities = value;
        }

        public JobAgentIdentityArgs()
        {
        }
        public static new JobAgentIdentityArgs Empty => new JobAgentIdentityArgs();
    }
}
