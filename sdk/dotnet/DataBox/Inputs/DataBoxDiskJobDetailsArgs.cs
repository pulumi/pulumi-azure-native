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
    /// DataBox Disk Job Details.
    /// </summary>
    public sealed class DataBoxDiskJobDetailsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Contact details for notification and shipping.
        /// </summary>
        [Input("contactDetails", required: true)]
        public Input<Inputs.ContactDetailsArgs> ContactDetails { get; set; } = null!;

        [Input("dataExportDetails")]
        private InputList<Inputs.DataExportDetailsArgs>? _dataExportDetails;

        /// <summary>
        /// Details of the data to be exported from azure.
        /// </summary>
        public InputList<Inputs.DataExportDetailsArgs> DataExportDetails
        {
            get => _dataExportDetails ?? (_dataExportDetails = new InputList<Inputs.DataExportDetailsArgs>());
            set => _dataExportDetails = value;
        }

        [Input("dataImportDetails")]
        private InputList<Inputs.DataImportDetailsArgs>? _dataImportDetails;

        /// <summary>
        /// Details of the data to be imported into azure.
        /// </summary>
        public InputList<Inputs.DataImportDetailsArgs> DataImportDetails
        {
            get => _dataImportDetails ?? (_dataImportDetails = new InputList<Inputs.DataImportDetailsArgs>());
            set => _dataImportDetails = value;
        }

        /// <summary>
        /// The expected size of the data, which needs to be transferred in this job, in terabytes.
        /// </summary>
        [Input("expectedDataSizeInTeraBytes")]
        public Input<int>? ExpectedDataSizeInTeraBytes { get; set; }

        /// <summary>
        /// Indicates the type of job details.
        /// Expected value is 'DataBoxDisk'.
        /// </summary>
        [Input("jobDetailsType", required: true)]
        public Input<string> JobDetailsType { get; set; } = null!;

        /// <summary>
        /// Details about which key encryption type is being used.
        /// </summary>
        [Input("keyEncryptionKey")]
        public Input<Inputs.KeyEncryptionKeyArgs>? KeyEncryptionKey { get; set; }

        /// <summary>
        /// User entered passkey for DataBox Disk job.
        /// </summary>
        [Input("passkey")]
        public Input<string>? Passkey { get; set; }

        /// <summary>
        /// Preferences for the order.
        /// </summary>
        [Input("preferences")]
        public Input<Inputs.PreferencesArgs>? Preferences { get; set; }

        [Input("preferredDisks")]
        private InputMap<int>? _preferredDisks;

        /// <summary>
        /// User preference on what size disks are needed for the job. The map is from the disk size in TB to the count. Eg. {2,5} means 5 disks of 2 TB size. Key is string but will be checked against an int.
        /// </summary>
        public InputMap<int> PreferredDisks
        {
            get => _preferredDisks ?? (_preferredDisks = new InputMap<int>());
            set => _preferredDisks = value;
        }

        /// <summary>
        /// Optional Reverse Shipping details for order.
        /// </summary>
        [Input("reverseShippingDetails")]
        public Input<Inputs.ReverseShippingDetailsArgs>? ReverseShippingDetails { get; set; }

        /// <summary>
        /// Shipping address of the customer.
        /// </summary>
        [Input("shippingAddress")]
        public Input<Inputs.ShippingAddressArgs>? ShippingAddress { get; set; }

        public DataBoxDiskJobDetailsArgs()
        {
        }
        public static new DataBoxDiskJobDetailsArgs Empty => new DataBoxDiskJobDetailsArgs();
    }
}
