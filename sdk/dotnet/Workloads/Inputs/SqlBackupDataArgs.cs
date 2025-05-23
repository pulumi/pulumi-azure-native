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
    /// Defines the SQL Backup data for a virtual instance for SAP.
    /// </summary>
    public sealed class SqlBackupDataArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines the policy properties for database backup.
        /// </summary>
        [Input("backupPolicy", required: true)]
        public Input<Inputs.DBBackupPolicyPropertiesArgs> BackupPolicy { get; set; } = null!;

        /// <summary>
        /// The type of backup, VM, SQL or HANA.
        /// Expected value is 'SQL'.
        /// </summary>
        [Input("backupType", required: true)]
        public Input<string> BackupType { get; set; } = null!;

        /// <summary>
        /// The properties of the recovery services vault used for backup.
        /// </summary>
        [Input("recoveryServicesVault", required: true)]
        public InputUnion<Inputs.ExistingRecoveryServicesVaultArgs, Inputs.NewRecoveryServicesVaultArgs> RecoveryServicesVault { get; set; } = null!;

        public SqlBackupDataArgs()
        {
        }
        public static new SqlBackupDataArgs Empty => new SqlBackupDataArgs();
    }
}
