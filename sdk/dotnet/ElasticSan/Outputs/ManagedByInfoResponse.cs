// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ElasticSan.Outputs
{

    /// <summary>
    /// Parent resource information.
    /// </summary>
    [OutputType]
    public sealed class ManagedByInfoResponse
    {
        /// <summary>
        /// Resource ID of the resource managing the volume, this is a restricted field and can only be set for internal use.
        /// </summary>
        public readonly string? ResourceId;

        [OutputConstructor]
        private ManagedByInfoResponse(string? resourceId)
        {
            ResourceId = resourceId;
        }
    }
}
