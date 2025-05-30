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
    /// Input for the task that validates Oracle database connection
    /// </summary>
    public sealed class ConnectToSourceOracleSyncTaskInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Information for connecting to Oracle source
        /// </summary>
        [Input("sourceConnectionInfo", required: true)]
        public Input<Inputs.OracleConnectionInfoArgs> SourceConnectionInfo { get; set; } = null!;

        public ConnectToSourceOracleSyncTaskInputArgs()
        {
        }
        public static new ConnectToSourceOracleSyncTaskInputArgs Empty => new ConnectToSourceOracleSyncTaskInputArgs();
    }
}
