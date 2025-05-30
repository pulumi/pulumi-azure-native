// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevOpsInfrastructure.Outputs
{

    /// <summary>
    /// GitHub organization profile
    /// </summary>
    [OutputType]
    public sealed class GitHubOrganizationProfileResponse
    {
        /// <summary>
        /// Discriminator property for OrganizationProfile.
        /// Expected value is 'GitHub'.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The list of GitHub organizations/repositories the pool should be present in.
        /// </summary>
        public readonly ImmutableArray<Outputs.GitHubOrganizationResponse> Organizations;

        [OutputConstructor]
        private GitHubOrganizationProfileResponse(
            string kind,

            ImmutableArray<Outputs.GitHubOrganizationResponse> organizations)
        {
            Kind = kind;
            Organizations = organizations;
        }
    }
}
