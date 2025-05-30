// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProgrammableConnectivity.Outputs
{

    /// <summary>
    /// Details about the SaaS offer purchased from the marketplace.
    /// </summary>
    [OutputType]
    public sealed class SaasPropertiesResponse
    {
        /// <summary>
        /// Resource ID of the SaaS offer purchased from the marketplace.
        /// </summary>
        public readonly string? SaasResourceId;
        /// <summary>
        /// Subscription ID of the SaaS offer purchased from the marketplace.
        /// </summary>
        public readonly string? SaasSubscriptionId;

        [OutputConstructor]
        private SaasPropertiesResponse(
            string? saasResourceId,

            string? saasSubscriptionId)
        {
            SaasResourceId = saasResourceId;
            SaasSubscriptionId = saasSubscriptionId;
        }
    }
}
