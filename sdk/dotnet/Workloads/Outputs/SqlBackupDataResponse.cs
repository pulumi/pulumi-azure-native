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
    /// Defines the SQL Backup data for a virtual instance for SAP.
    /// </summary>
    [OutputType]
    public sealed class SqlBackupDataResponse
    {
        /// <summary>
        /// Defines the policy properties for database backup.
        /// </summary>
        public readonly Outputs.DBBackupPolicyPropertiesResponse BackupPolicy;
        /// <summary>
        /// The type of backup, VM, SQL or HANA.
        /// Expected value is 'SQL'.
        /// </summary>
        public readonly string BackupType;
        /// <summary>
        /// The properties of the recovery services vault used for backup.
        /// </summary>
        public readonly Union<Outputs.ExistingRecoveryServicesVaultResponse, Outputs.NewRecoveryServicesVaultResponse> RecoveryServicesVault;

        [OutputConstructor]
        private SqlBackupDataResponse(
            Outputs.DBBackupPolicyPropertiesResponse backupPolicy,

            string backupType,

            Union<Outputs.ExistingRecoveryServicesVaultResponse, Outputs.NewRecoveryServicesVaultResponse> recoveryServicesVault)
        {
            BackupPolicy = backupPolicy;
            BackupType = backupType;
            RecoveryServicesVault = recoveryServicesVault;
        }
    }
}
