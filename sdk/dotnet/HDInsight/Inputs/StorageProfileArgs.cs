// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.Inputs
{

    /// <summary>
    /// The storage profile.
    /// </summary>
    public sealed class StorageProfileArgs : global::Pulumi.ResourceArgs
    {
        [Input("storageaccounts")]
        private InputList<Inputs.StorageAccountArgs>? _storageaccounts;

        /// <summary>
        /// The list of storage accounts in the cluster.
        /// </summary>
        public InputList<Inputs.StorageAccountArgs> Storageaccounts
        {
            get => _storageaccounts ?? (_storageaccounts = new InputList<Inputs.StorageAccountArgs>());
            set => _storageaccounts = value;
        }

        public StorageProfileArgs()
        {
        }
        public static new StorageProfileArgs Empty => new StorageProfileArgs();
    }
}
