// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataProtection.V20231101.Inputs
{

    /// <summary>
    /// Storage setting
    /// </summary>
    public sealed class StorageSettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the type of the datastore.
        /// </summary>
        [Input("datastoreType")]
        public InputUnion<string, Pulumi.AzureNative.DataProtection.V20231101.StorageSettingStoreTypes>? DatastoreType { get; set; }

        /// <summary>
        /// Gets or sets the type.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNative.DataProtection.V20231101.StorageSettingTypes>? Type { get; set; }

        public StorageSettingArgs()
        {
        }
        public static new StorageSettingArgs Empty => new StorageSettingArgs();
    }
}
