// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of awsRdsExportTask
    /// </summary>
    public sealed class AwsRdsExportTaskPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("exportOnly")]
        private InputList<string>? _exportOnly;

        /// <summary>
        /// &lt;p&gt;The data exported from the snapshot or cluster.&lt;/p&gt; &lt;p&gt;Valid Values:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;database&lt;/code&gt; - Export all the data from a specified database.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;database.table&lt;/code&gt; &lt;i&gt;table-name&lt;/i&gt; - Export a table of the snapshot or cluster. This format is valid only for RDS for MySQL, RDS for MariaDB, and Aurora MySQL.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;database.schema&lt;/code&gt; &lt;i&gt;schema-name&lt;/i&gt; - Export a database schema of the snapshot or cluster. This format is valid only for RDS for PostgreSQL and Aurora PostgreSQL.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;database.schema.table&lt;/code&gt; &lt;i&gt;table-name&lt;/i&gt; - Export a table of the database schema. This format is valid only for RDS for PostgreSQL and Aurora PostgreSQL.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt;
        /// </summary>
        public InputList<string> ExportOnly
        {
            get => _exportOnly ?? (_exportOnly = new InputList<string>());
            set => _exportOnly = value;
        }

        /// <summary>
        /// &lt;p&gt;A unique identifier for the snapshot or cluster export task. This ID isn't an identifier for the Amazon S3 bucket where the data is exported.&lt;/p&gt;
        /// </summary>
        [Input("exportTaskIdentifier")]
        public Input<string>? ExportTaskIdentifier { get; set; }

        /// <summary>
        /// &lt;p&gt;The reason the export failed, if it failed.&lt;/p&gt;
        /// </summary>
        [Input("failureCause")]
        public Input<string>? FailureCause { get; set; }

        /// <summary>
        /// &lt;p&gt;The name of the IAM role that is used to write to Amazon S3 when exporting a snapshot or cluster.&lt;/p&gt;
        /// </summary>
        [Input("iamRoleArn")]
        public Input<string>? IamRoleArn { get; set; }

        /// <summary>
        /// &lt;p&gt;The key identifier of the Amazon Web Services KMS key that is used to encrypt the data when it's exported to Amazon S3. The KMS key identifier is its key ARN, key ID, alias ARN, or alias name. The IAM role used for the export must have encryption and decryption permissions to use this KMS key.&lt;/p&gt;
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// &lt;p&gt;The progress of the snapshot or cluster export task as a percentage.&lt;/p&gt;
        /// </summary>
        [Input("percentProgress")]
        public Input<int>? PercentProgress { get; set; }

        /// <summary>
        /// &lt;p&gt;The Amazon S3 bucket where the snapshot or cluster is exported to.&lt;/p&gt;
        /// </summary>
        [Input("s3Bucket")]
        public Input<string>? S3Bucket { get; set; }

        /// <summary>
        /// &lt;p&gt;The Amazon S3 bucket prefix that is the file name and path of the exported data.&lt;/p&gt;
        /// </summary>
        [Input("s3Prefix")]
        public Input<string>? S3Prefix { get; set; }

        /// <summary>
        /// &lt;p&gt;The time when the snapshot was created.&lt;/p&gt;
        /// </summary>
        [Input("snapshotTime")]
        public Input<string>? SnapshotTime { get; set; }

        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) of the snapshot or cluster exported to Amazon S3.&lt;/p&gt;
        /// </summary>
        [Input("sourceArn")]
        public Input<string>? SourceArn { get; set; }

        /// <summary>
        /// &lt;p&gt;The type of source for the export.&lt;/p&gt;
        /// </summary>
        [Input("sourceType")]
        public Input<Inputs.ExportSourceTypeEnumValueArgs>? SourceType { get; set; }

        /// <summary>
        /// &lt;p&gt;The progress status of the export task. The status can be one of the following:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;CANCELED&lt;/code&gt; &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;CANCELING&lt;/code&gt; &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;COMPLETE&lt;/code&gt; &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;FAILED&lt;/code&gt; &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;IN_PROGRESS&lt;/code&gt; &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;STARTING&lt;/code&gt; &lt;/p&gt; &lt;/li&gt; &lt;/ul&gt;
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// &lt;p&gt;The time when the snapshot or cluster export task ended.&lt;/p&gt;
        /// </summary>
        [Input("taskEndTime")]
        public Input<string>? TaskEndTime { get; set; }

        /// <summary>
        /// &lt;p&gt;The time when the snapshot or cluster export task started.&lt;/p&gt;
        /// </summary>
        [Input("taskStartTime")]
        public Input<string>? TaskStartTime { get; set; }

        /// <summary>
        /// &lt;p&gt;The total amount of data exported, in gigabytes.&lt;/p&gt;
        /// </summary>
        [Input("totalExtractedDataInGB")]
        public Input<int>? TotalExtractedDataInGB { get; set; }

        /// <summary>
        /// &lt;p&gt;A warning about the snapshot or cluster export task.&lt;/p&gt;
        /// </summary>
        [Input("warningMessage")]
        public Input<string>? WarningMessage { get; set; }

        public AwsRdsExportTaskPropertiesArgs()
        {
        }
        public static new AwsRdsExportTaskPropertiesArgs Empty => new AwsRdsExportTaskPropertiesArgs();
    }
}
