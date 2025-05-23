// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Inputs
{

    /// <summary>
    /// The gcpOrganization data for the member account
    /// </summary>
    public sealed class GcpOrganizationalDataMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The GCP management project number from organizational onboarding
        /// </summary>
        [Input("managementProjectNumber")]
        public Input<string>? ManagementProjectNumber { get; set; }

        /// <summary>
        /// The multi cloud account's membership type in the organization
        /// Expected value is 'Member'.
        /// </summary>
        [Input("organizationMembershipType", required: true)]
        public Input<string> OrganizationMembershipType { get; set; } = null!;

        /// <summary>
        /// If the multi cloud account is not of membership type organization, this will be the ID of the project's parent
        /// </summary>
        [Input("parentHierarchyId")]
        public Input<string>? ParentHierarchyId { get; set; }

        public GcpOrganizationalDataMemberArgs()
        {
        }
        public static new GcpOrganizationalDataMemberArgs Empty => new GcpOrganizationalDataMemberArgs();
    }
}
