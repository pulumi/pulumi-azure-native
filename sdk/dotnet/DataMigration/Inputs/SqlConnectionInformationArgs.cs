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
    /// Source SQL Connection
    /// </summary>
    public sealed class SqlConnectionInformationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication type.
        /// </summary>
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        /// <summary>
        /// Data source.
        /// </summary>
        [Input("dataSource")]
        public Input<string>? DataSource { get; set; }

        /// <summary>
        /// Whether to encrypt connection or not.
        /// </summary>
        [Input("encryptConnection")]
        public Input<bool>? EncryptConnection { get; set; }

        /// <summary>
        /// Password to connect to source SQL.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Whether to trust server certificate or not.
        /// </summary>
        [Input("trustServerCertificate")]
        public Input<bool>? TrustServerCertificate { get; set; }

        /// <summary>
        /// User name to connect to source SQL.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public SqlConnectionInformationArgs()
        {
        }
        public static new SqlConnectionInformationArgs Empty => new SqlConnectionInformationArgs();
    }
}
