// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforMySQL.V20231230.Inputs
{

    /// <summary>
    /// Storage Profile properties of a server
    /// </summary>
    public sealed class StorageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable Storage Auto Grow or not.
        /// </summary>
        [Input("autoGrow")]
        public InputUnion<string, Pulumi.AzureNative.DBforMySQL.V20231230.EnableStatusEnum>? AutoGrow { get; set; }

        /// <summary>
        /// Enable IO Auto Scaling or not.
        /// </summary>
        [Input("autoIoScaling")]
        public InputUnion<string, Pulumi.AzureNative.DBforMySQL.V20231230.EnableStatusEnum>? AutoIoScaling { get; set; }

        /// <summary>
        /// Storage IOPS for a server.
        /// </summary>
        [Input("iops")]
        public Input<int>? Iops { get; set; }

        /// <summary>
        /// Enable Log On Disk or not.
        /// </summary>
        [Input("logOnDisk")]
        public InputUnion<string, Pulumi.AzureNative.DBforMySQL.V20231230.EnableStatusEnum>? LogOnDisk { get; set; }

        /// <summary>
        /// Max storage size allowed for a server.
        /// </summary>
        [Input("storageSizeGB")]
        public Input<int>? StorageSizeGB { get; set; }

        public StorageArgs()
        {
            AutoGrow = "Disabled";
            AutoIoScaling = "Enabled";
            LogOnDisk = "Disabled";
        }
        public static new StorageArgs Empty => new StorageArgs();
    }
}
