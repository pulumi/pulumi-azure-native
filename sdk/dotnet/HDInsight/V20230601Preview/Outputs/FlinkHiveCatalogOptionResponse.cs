// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.V20230601Preview.Outputs
{

    /// <summary>
    /// Hive Catalog Option for Flink cluster.
    /// </summary>
    [OutputType]
    public sealed class FlinkHiveCatalogOptionResponse
    {
        /// <summary>
        /// Secret reference name from secretsProfile.secrets containing password for database connection.
        /// </summary>
        public readonly string MetastoreDbConnectionPasswordSecret;
        /// <summary>
        /// Connection string for hive metastore database.
        /// </summary>
        public readonly string MetastoreDbConnectionURL;
        /// <summary>
        /// User name for database connection.
        /// </summary>
        public readonly string MetastoreDbConnectionUserName;

        [OutputConstructor]
        private FlinkHiveCatalogOptionResponse(
            string metastoreDbConnectionPasswordSecret,

            string metastoreDbConnectionURL,

            string metastoreDbConnectionUserName)
        {
            MetastoreDbConnectionPasswordSecret = metastoreDbConnectionPasswordSecret;
            MetastoreDbConnectionURL = metastoreDbConnectionURL;
            MetastoreDbConnectionUserName = metastoreDbConnectionUserName;
        }
    }
}
