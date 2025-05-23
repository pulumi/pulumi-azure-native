// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CosmosDB.Outputs
{

    /// <summary>
    /// The set of data plane operations permitted through this Role Definition.
    /// </summary>
    [OutputType]
    public sealed class PermissionResponse
    {
        /// <summary>
        /// An array of data actions that are allowed.
        /// </summary>
        public readonly ImmutableArray<string> DataActions;
        /// <summary>
        /// The id for the permission.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// An array of data actions that are denied.
        /// </summary>
        public readonly ImmutableArray<string> NotDataActions;

        [OutputConstructor]
        private PermissionResponse(
            ImmutableArray<string> dataActions,

            string? id,

            ImmutableArray<string> notDataActions)
        {
            DataActions = dataActions;
            Id = id;
            NotDataActions = notDataActions;
        }
    }
}
