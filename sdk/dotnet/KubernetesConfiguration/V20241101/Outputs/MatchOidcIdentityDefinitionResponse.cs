// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.KubernetesConfiguration.V20241101.Outputs
{

    /// <summary>
    /// MatchOIDCIdentity defines the criteria for matching the identity while verifying an OCI artifact.
    /// </summary>
    [OutputType]
    public sealed class MatchOidcIdentityDefinitionResponse
    {
        /// <summary>
        /// The regex pattern to match against to verify the OIDC issuer.
        /// </summary>
        public readonly string? Issuer;
        /// <summary>
        /// The regex pattern to match against to verify the identity subject.
        /// </summary>
        public readonly string? Subject;

        [OutputConstructor]
        private MatchOidcIdentityDefinitionResponse(
            string? issuer,

            string? subject)
        {
            Issuer = issuer;
            Subject = subject;
        }
    }
}
