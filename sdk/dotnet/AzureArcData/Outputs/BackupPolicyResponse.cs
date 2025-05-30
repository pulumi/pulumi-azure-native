// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureArcData.Outputs
{

    /// <summary>
    /// The backup profile for the SQL server.
    /// </summary>
    [OutputType]
    public sealed class BackupPolicyResponse
    {
        /// <summary>
        /// The differential backup interval in hours.
        /// </summary>
        public readonly int? DifferentialBackupHours;
        /// <summary>
        /// The value indicating days between full backups.
        /// </summary>
        public readonly int? FullBackupDays;
        /// <summary>
        /// The retention period for all the databases in this managed instance.
        /// </summary>
        public readonly int? RetentionPeriodDays;
        /// <summary>
        /// The value indicating minutes between transaction log backups.
        /// </summary>
        public readonly int? TransactionLogBackupMinutes;

        [OutputConstructor]
        private BackupPolicyResponse(
            int? differentialBackupHours,

            int? fullBackupDays,

            int? retentionPeriodDays,

            int? transactionLogBackupMinutes)
        {
            DifferentialBackupHours = differentialBackupHours;
            FullBackupDays = fullBackupDays;
            RetentionPeriodDays = retentionPeriodDays;
            TransactionLogBackupMinutes = transactionLogBackupMinutes;
        }
    }
}
