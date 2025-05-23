// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.Outputs
{

    [OutputType]
    public sealed class RangerAdminSpecResponseDatabase
    {
        /// <summary>
        /// The database URL
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// The database name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Reference for the database password
        /// </summary>
        public readonly string? PasswordSecretRef;
        /// <summary>
        /// The name of the database user
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private RangerAdminSpecResponseDatabase(
            string host,

            string name,

            string? passwordSecretRef,

            string? username)
        {
            Host = host;
            Name = name;
            PasswordSecretRef = passwordSecretRef;
            Username = username;
        }
    }
}
