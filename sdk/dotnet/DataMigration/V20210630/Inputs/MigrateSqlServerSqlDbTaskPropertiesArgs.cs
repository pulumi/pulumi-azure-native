// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataMigration.V20210630.Inputs
{

    /// <summary>
    /// Properties for the task that migrates on-prem SQL Server databases to Azure SQL Database
    /// </summary>
    public sealed class MigrateSqlServerSqlDbTaskPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("clientData")]
        private InputMap<string>? _clientData;

        /// <summary>
        /// Key value pairs of client data to attach meta data information to task
        /// </summary>
        public InputMap<string> ClientData
        {
            get => _clientData ?? (_clientData = new InputMap<string>());
            set => _clientData = value;
        }

        /// <summary>
        /// Task input
        /// </summary>
        [Input("input")]
        public Input<Inputs.MigrateSqlServerSqlDbTaskInputArgs>? Input { get; set; }

        /// <summary>
        /// Task type.
        /// Expected value is 'Migrate.SqlServer.SqlDb'.
        /// </summary>
        [Input("taskType", required: true)]
        public Input<string> TaskType { get; set; } = null!;

        public MigrateSqlServerSqlDbTaskPropertiesArgs()
        {
        }
        public static new MigrateSqlServerSqlDbTaskPropertiesArgs Empty => new MigrateSqlServerSqlDbTaskPropertiesArgs();
    }
}
