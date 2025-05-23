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
    /// Describe the properties of a assignment attestation
    /// </summary>
    public sealed class AttestationEvidenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the evidence
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The source url of the evidence
        /// </summary>
        [Input("sourceUrl")]
        public Input<string>? SourceUrl { get; set; }

        public AttestationEvidenceArgs()
        {
        }
        public static new AttestationEvidenceArgs Empty => new AttestationEvidenceArgs();
    }
}
