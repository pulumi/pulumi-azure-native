// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MobileNetwork.Outputs
{

    /// <summary>
    /// Reference to a data network resource.
    /// </summary>
    [OutputType]
    public sealed class DataNetworkResourceIdResponse
    {
        /// <summary>
        /// Data network resource ID.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private DataNetworkResourceIdResponse(string id)
        {
            Id = id;
        }
    }
}
