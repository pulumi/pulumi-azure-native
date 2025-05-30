// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataMigration.Inputs
{

    /// <summary>
    /// Properties for task that migrates SSIS packages from SQL Server databases to Azure SQL Database Managed Instance.
    /// </summary>
    public sealed class MigrateSsisTaskPropertiesArgs : global::Pulumi.ResourceArgs
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
        public Input<Inputs.MigrateSsisTaskInputArgs>? Input { get; set; }

        /// <summary>
        /// Task type.
        /// Expected value is 'Migrate.Ssis'.
        /// </summary>
        [Input("taskType", required: true)]
        public Input<string> TaskType { get; set; } = null!;

        public MigrateSsisTaskPropertiesArgs()
        {
        }
        public static new MigrateSsisTaskPropertiesArgs Empty => new MigrateSsisTaskPropertiesArgs();
    }
}
