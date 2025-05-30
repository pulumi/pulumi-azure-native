// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Storage.Outputs
{

    /// <summary>
    /// The LegalHold property of a blob container.
    /// </summary>
    [OutputType]
    public sealed class LegalHoldPropertiesResponse
    {
        /// <summary>
        /// The hasLegalHold public property is set to true by SRP if there are at least one existing tag. The hasLegalHold public property is set to false by SRP if all existing legal hold tags are cleared out. There can be a maximum of 1000 blob containers with hasLegalHold=true for a given account.
        /// </summary>
        public readonly bool HasLegalHold;
        /// <summary>
        /// Protected append blob writes history.
        /// </summary>
        public readonly Outputs.ProtectedAppendWritesHistoryResponse? ProtectedAppendWritesHistory;
        /// <summary>
        /// The list of LegalHold tags of a blob container.
        /// </summary>
        public readonly ImmutableArray<Outputs.TagPropertyResponse> Tags;

        [OutputConstructor]
        private LegalHoldPropertiesResponse(
            bool hasLegalHold,

            Outputs.ProtectedAppendWritesHistoryResponse? protectedAppendWritesHistory,

            ImmutableArray<Outputs.TagPropertyResponse> tags)
        {
            HasLegalHold = hasLegalHold;
            ProtectedAppendWritesHistory = protectedAppendWritesHistory;
            Tags = tags;
        }
    }
}
