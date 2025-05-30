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
    /// Gets or sets the database configuration.
    /// </summary>
    public sealed class DatabaseConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The database type.
        /// </summary>
        [Input("databaseType")]
        public InputUnion<string, Pulumi.AzureNative.Workloads.SAPDatabaseType>? DatabaseType { get; set; }

        /// <summary>
        /// Gets or sets the disk configuration.
        /// </summary>
        [Input("diskConfiguration")]
        public Input<Inputs.DiskConfigurationArgs>? DiskConfiguration { get; set; }

        /// <summary>
        /// The number of database VMs.
        /// </summary>
        [Input("instanceCount", required: true)]
        public Input<double> InstanceCount { get; set; } = null!;

        /// <summary>
        /// The subnet id.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        /// <summary>
        /// Gets or sets the virtual machine configuration.
        /// </summary>
        [Input("virtualMachineConfiguration", required: true)]
        public Input<Inputs.VirtualMachineConfigurationArgs> VirtualMachineConfiguration { get; set; } = null!;

        public DatabaseConfigurationArgs()
        {
        }
        public static new DatabaseConfigurationArgs Empty => new DatabaseConfigurationArgs();
    }
}
