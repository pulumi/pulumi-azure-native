// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads.Inputs
{

    /// <summary>
    /// The resource names object for shared storage.
    /// </summary>
    public sealed class SharedStorageResourceNamesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The full name of the shared storage account. If it is not provided, it will be defaulted to {SID}nfs{guid of 15 chars}.
        /// </summary>
        [Input("sharedStorageAccountName")]
        public Input<string>? SharedStorageAccountName { get; set; }

        /// <summary>
        /// The full name of private end point for the shared storage account. If it is not provided, it will be defaulted to {storageAccountName}_pe
        /// </summary>
        [Input("sharedStorageAccountPrivateEndPointName")]
        public Input<string>? SharedStorageAccountPrivateEndPointName { get; set; }

        public SharedStorageResourceNamesArgs()
        {
        }
        public static new SharedStorageResourceNamesArgs Empty => new SharedStorageResourceNamesArgs();
    }
}
