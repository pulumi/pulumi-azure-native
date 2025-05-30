// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.PolicyInsights.Outputs
{

    /// <summary>
    /// A piece of evidence supporting the compliance state set in the attestation.
    /// </summary>
    [OutputType]
    public sealed class AttestationEvidenceResponse
    {
        /// <summary>
        /// The description for this piece of evidence.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The URI location of the evidence.
        /// </summary>
        public readonly string? SourceUri;

        [OutputConstructor]
        private AttestationEvidenceResponse(
            string? description,

            string? sourceUri)
        {
            Description = description;
            SourceUri = sourceUri;
        }
    }
}
