// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.WeightsAndBiases.Outputs
{

    /// <summary>
    /// Properties specific to Single Sign On Resource
    /// </summary>
    [OutputType]
    public sealed class LiftrBaseSingleSignOnPropertiesV2Response
    {
        /// <summary>
        /// List of AAD domains fetched from Microsoft Graph for user.
        /// </summary>
        public readonly ImmutableArray<string> AadDomains;
        /// <summary>
        /// AAD enterprise application Id used to setup SSO
        /// </summary>
        public readonly string? EnterpriseAppId;
        /// <summary>
        /// State of the Single Sign On for the resource
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// Type of Single Sign-On mechanism being used
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// URL for SSO to be used by the partner to redirect the user to their system
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private LiftrBaseSingleSignOnPropertiesV2Response(
            ImmutableArray<string> aadDomains,

            string? enterpriseAppId,

            string? state,

            string type,

            string? url)
        {
            AadDomains = aadDomains;
            EnterpriseAppId = enterpriseAppId;
            State = state;
            Type = type;
            Url = url;
        }
    }
}
