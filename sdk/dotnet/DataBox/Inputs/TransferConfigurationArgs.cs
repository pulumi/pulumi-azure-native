// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBox.Inputs
{

    /// <summary>
    /// Configuration for defining the transfer of data.
    /// </summary>
    public sealed class TransferConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Map of filter type and the details to transfer all data. This field is required only if the TransferConfigurationType is given as TransferAll
        /// </summary>
        [Input("transferAllDetails")]
        public Input<Inputs.TransferConfigurationTransferAllDetailsArgs>? TransferAllDetails { get; set; }

        /// <summary>
        /// Type of the configuration for transfer.
        /// </summary>
        [Input("transferConfigurationType", required: true)]
        public InputUnion<string, Pulumi.AzureNative.DataBox.TransferConfigurationType> TransferConfigurationType { get; set; } = null!;

        /// <summary>
        /// Map of filter type and the details to filter. This field is required only if the TransferConfigurationType is given as TransferUsingFilter.
        /// </summary>
        [Input("transferFilterDetails")]
        public Input<Inputs.TransferConfigurationTransferFilterDetailsArgs>? TransferFilterDetails { get; set; }

        public TransferConfigurationArgs()
        {
        }
        public static new TransferConfigurationArgs Empty => new TransferConfigurationArgs();
    }
}
