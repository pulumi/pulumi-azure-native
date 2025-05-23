// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.Outputs
{

    /// <summary>
    /// Backend pool service information
    /// </summary>
    [OutputType]
    public sealed class BackendPoolItemResponse
    {
        /// <summary>
        /// The unique ARM id of the backend entity. The ARM id should refer to an already existing backend entity.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The priority of the backend entity in the backend pool. Must be between 0 and 100. It can be also null if the value not specified.
        /// </summary>
        public readonly int? Priority;
        /// <summary>
        /// The weight of the backend entity in the backend pool. Must be between 0 and 100. It can be also null if the value not specified.
        /// </summary>
        public readonly int? Weight;

        [OutputConstructor]
        private BackendPoolItemResponse(
            string id,

            int? priority,

            int? weight)
        {
            Id = id;
            Priority = priority;
            Weight = weight;
        }
    }
}
