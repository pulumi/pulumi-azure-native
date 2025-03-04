// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20210301Preview.Outputs
{

    /// <summary>
    /// Azure SQL Database datastore configuration.
    /// </summary>
    [OutputType]
    public sealed class AzureSqlDatabaseContentsResponse
    {
        /// <summary>
        /// Enum to determine the datastore contents type.
        /// Expected value is 'AzureSqlDatabase'.
        /// </summary>
        public readonly string ContentsType;
        /// <summary>
        /// [Required] Account credentials.
        /// </summary>
        public readonly object Credentials;
        /// <summary>
        /// [Required] Azure SQL database name.
        /// </summary>
        public readonly string DatabaseName;
        /// <summary>
        /// [Required] Azure cloud endpoint for the database.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// [Required] Azure SQL server port.
        /// </summary>
        public readonly int PortNumber;
        /// <summary>
        /// [Required] Azure SQL server name.
        /// </summary>
        public readonly string ServerName;

        [OutputConstructor]
        private AzureSqlDatabaseContentsResponse(
            string contentsType,

            object credentials,

            string databaseName,

            string endpoint,

            int portNumber,

            string serverName)
        {
            ContentsType = contentsType;
            Credentials = credentials;
            DatabaseName = databaseName;
            Endpoint = endpoint;
            PortNumber = portNumber;
            ServerName = serverName;
        }
    }
}
