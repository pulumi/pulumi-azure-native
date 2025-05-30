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
    /// Reference to an Azure Stack HCI cluster resource.
    /// </summary>
    [OutputType]
    public sealed class AzureStackHCIClusterResourceIdResponse
    {
        /// <summary>
        /// Azure Stack HCI cluster resource ID.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private AzureStackHCIClusterResourceIdResponse(string id)
        {
            Id = id;
        }
    }
}
