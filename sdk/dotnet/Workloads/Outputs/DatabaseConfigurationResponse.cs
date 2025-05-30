// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads.Outputs
{

    /// <summary>
    /// Gets or sets the database configuration.
    /// </summary>
    [OutputType]
    public sealed class DatabaseConfigurationResponse
    {
        /// <summary>
        /// The database type.
        /// </summary>
        public readonly string? DatabaseType;
        /// <summary>
        /// Gets or sets the disk configuration.
        /// </summary>
        public readonly Outputs.DiskConfigurationResponse? DiskConfiguration;
        /// <summary>
        /// The number of database VMs.
        /// </summary>
        public readonly double InstanceCount;
        /// <summary>
        /// The subnet id.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// Gets or sets the virtual machine configuration.
        /// </summary>
        public readonly Outputs.VirtualMachineConfigurationResponse VirtualMachineConfiguration;

        [OutputConstructor]
        private DatabaseConfigurationResponse(
            string? databaseType,

            Outputs.DiskConfigurationResponse? diskConfiguration,

            double instanceCount,

            string subnetId,

            Outputs.VirtualMachineConfigurationResponse virtualMachineConfiguration)
        {
            DatabaseType = databaseType;
            DiskConfiguration = diskConfiguration;
            InstanceCount = instanceCount;
            SubnetId = subnetId;
            VirtualMachineConfiguration = virtualMachineConfiguration;
        }
    }
}
