// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    /// <summary>
    /// Details for each region the registry is in
    /// </summary>
    public sealed class RegistryRegionArmDetailsArgs : global::Pulumi.ResourceArgs
    {
        [Input("acrDetails")]
        private InputList<Inputs.AcrDetailsArgs>? _acrDetails;

        /// <summary>
        /// List of ACR accounts
        /// </summary>
        public InputList<Inputs.AcrDetailsArgs> AcrDetails
        {
            get => _acrDetails ?? (_acrDetails = new InputList<Inputs.AcrDetailsArgs>());
            set => _acrDetails = value;
        }

        /// <summary>
        /// The location where the registry exists
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("storageAccountDetails")]
        private InputList<Inputs.StorageAccountDetailsArgs>? _storageAccountDetails;

        /// <summary>
        /// List of storage accounts
        /// </summary>
        public InputList<Inputs.StorageAccountDetailsArgs> StorageAccountDetails
        {
            get => _storageAccountDetails ?? (_storageAccountDetails = new InputList<Inputs.StorageAccountDetailsArgs>());
            set => _storageAccountDetails = value;
        }

        public RegistryRegionArmDetailsArgs()
        {
        }
        public static new RegistryRegionArmDetailsArgs Empty => new RegistryRegionArmDetailsArgs();
    }
}
