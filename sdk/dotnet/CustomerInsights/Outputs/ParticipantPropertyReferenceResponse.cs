// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CustomerInsights.Outputs
{

    /// <summary>
    /// The participant property reference.
    /// </summary>
    [OutputType]
    public sealed class ParticipantPropertyReferenceResponse
    {
        /// <summary>
        /// The source property that maps to the target property.
        /// </summary>
        public readonly string SourcePropertyName;
        /// <summary>
        /// The target property that maps to the source property.
        /// </summary>
        public readonly string TargetPropertyName;

        [OutputConstructor]
        private ParticipantPropertyReferenceResponse(
            string sourcePropertyName,

            string targetPropertyName)
        {
            SourcePropertyName = sourcePropertyName;
            TargetPropertyName = targetPropertyName;
        }
    }
}
