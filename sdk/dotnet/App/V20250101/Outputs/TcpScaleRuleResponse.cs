// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.V20250101.Outputs
{

    /// <summary>
    /// Container App container Tcp scaling rule.
    /// </summary>
    [OutputType]
    public sealed class TcpScaleRuleResponse
    {
        /// <summary>
        /// Authentication secrets for the tcp scale rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.ScaleRuleAuthResponse> Auth;
        /// <summary>
        /// The resource ID of a user-assigned managed identity that is assigned to the Container App, or 'system' for system-assigned identity.
        /// </summary>
        public readonly string? Identity;
        /// <summary>
        /// Metadata properties to describe tcp scale rule.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Metadata;

        [OutputConstructor]
        private TcpScaleRuleResponse(
            ImmutableArray<Outputs.ScaleRuleAuthResponse> auth,

            string? identity,

            ImmutableDictionary<string, string>? metadata)
        {
            Auth = auth;
            Identity = identity;
            Metadata = metadata;
        }
    }
}
