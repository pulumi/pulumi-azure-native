// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureArcData.Inputs
{

    /// <summary>
    /// The backup profile for the SQL server.
    /// </summary>
    public sealed class BackupPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The differential backup interval in hours.
        /// </summary>
        [Input("differentialBackupHours")]
        public Input<int>? DifferentialBackupHours { get; set; }

        /// <summary>
        /// The value indicating days between full backups.
        /// </summary>
        [Input("fullBackupDays")]
        public Input<int>? FullBackupDays { get; set; }

        /// <summary>
        /// The retention period for all the databases in this managed instance.
        /// </summary>
        [Input("retentionPeriodDays")]
        public Input<int>? RetentionPeriodDays { get; set; }

        /// <summary>
        /// The value indicating minutes between transaction log backups.
        /// </summary>
        [Input("transactionLogBackupMinutes")]
        public Input<int>? TransactionLogBackupMinutes { get; set; }

        public BackupPolicyArgs()
        {
        }
        public static new BackupPolicyArgs Empty => new BackupPolicyArgs();
    }
}
