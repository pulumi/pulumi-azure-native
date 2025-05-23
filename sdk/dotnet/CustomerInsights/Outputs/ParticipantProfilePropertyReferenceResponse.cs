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
    /// The participant profile property reference.
    /// </summary>
    [OutputType]
    public sealed class ParticipantProfilePropertyReferenceResponse
    {
        /// <summary>
        /// The source interaction property that maps to the target profile property.
        /// </summary>
        public readonly string InteractionPropertyName;
        /// <summary>
        /// The target profile property that maps to the source interaction property.
        /// </summary>
        public readonly string ProfilePropertyName;

        [OutputConstructor]
        private ParticipantProfilePropertyReferenceResponse(
            string interactionPropertyName,

            string profilePropertyName)
        {
            InteractionPropertyName = interactionPropertyName;
            ProfilePropertyName = profilePropertyName;
        }
    }
}
