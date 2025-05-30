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
    public sealed class PrivilegeResponse
    {
        /// <summary>
        /// An array of actions that are allowed.
        /// </summary>
        public readonly ImmutableArray<string> Actions;
        /// <summary>
        /// An Azure Cosmos DB Mongo DB Resource.
        /// </summary>
        public readonly Outputs.PrivilegeResponseResource? Resource;

        [OutputConstructor]
        private PrivilegeResponse(
            ImmutableArray<string> actions,

            Outputs.PrivilegeResponseResource? resource)
        {
            Actions = actions;
            Resource = resource;
        }
    }
}
