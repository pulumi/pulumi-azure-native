// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridContainerService.Outputs
{

    [OutputType]
    public sealed class StorageSpacesPropertiesResponseVmwareStorageProfile
    {
        /// <summary>
        /// Name of the datacenter in VSphere
        /// </summary>
        public readonly string? Datacenter;
        /// <summary>
        /// Name of the datastore in VSphere
        /// </summary>
        public readonly string? Datastore;
        /// <summary>
        /// Name of the folder in VSphere
        /// </summary>
        public readonly string? Folder;
        /// <summary>
        /// Name of the resource pool in VSphere
        /// </summary>
        public readonly string? ResourcePool;

        [OutputConstructor]
        private StorageSpacesPropertiesResponseVmwareStorageProfile(
            string? datacenter,

            string? datastore,

            string? folder,

            string? resourcePool)
        {
            Datacenter = datacenter;
            Datastore = datastore;
            Folder = folder;
            ResourcePool = resourcePool;
        }
    }
}
