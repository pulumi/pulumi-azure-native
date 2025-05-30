// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridContainerService.Inputs
{

    /// <summary>
    /// InfraStorageProfile - List of infra storage profiles for the provisioned cluster
    /// </summary>
    public sealed class CloudProviderProfileInfraStorageProfileArgs : global::Pulumi.ResourceArgs
    {
        [Input("storageSpaceIds")]
        private InputList<string>? _storageSpaceIds;

        /// <summary>
        /// Reference to azure resource corresponding to the new HybridAKSStorage object e.g. /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridContainerService/storageSpaces/{storageSpaceName}
        /// </summary>
        public InputList<string> StorageSpaceIds
        {
            get => _storageSpaceIds ?? (_storageSpaceIds = new InputList<string>());
            set => _storageSpaceIds = value;
        }

        public CloudProviderProfileInfraStorageProfileArgs()
        {
        }
        public static new CloudProviderProfileInfraStorageProfileArgs Empty => new CloudProviderProfileInfraStorageProfileArgs();
    }
}
