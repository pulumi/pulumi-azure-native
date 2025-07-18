// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql
{
    public static class GetBackupShortTermRetentionPolicy
    {
        /// <summary>
        /// Gets a database's short term retention policy.
        /// 
        /// Uses Azure REST API version 2023-08-01.
        /// 
        /// Other available API versions: 2017-10-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetBackupShortTermRetentionPolicyResult> InvokeAsync(GetBackupShortTermRetentionPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBackupShortTermRetentionPolicyResult>("azure-native:sql:getBackupShortTermRetentionPolicy", args ?? new GetBackupShortTermRetentionPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a database's short term retention policy.
        /// 
        /// Uses Azure REST API version 2023-08-01.
        /// 
        /// Other available API versions: 2017-10-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetBackupShortTermRetentionPolicyResult> Invoke(GetBackupShortTermRetentionPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBackupShortTermRetentionPolicyResult>("azure-native:sql:getBackupShortTermRetentionPolicy", args ?? new GetBackupShortTermRetentionPolicyInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a database's short term retention policy.
        /// 
        /// Uses Azure REST API version 2023-08-01.
        /// 
        /// Other available API versions: 2017-10-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetBackupShortTermRetentionPolicyResult> Invoke(GetBackupShortTermRetentionPolicyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetBackupShortTermRetentionPolicyResult>("azure-native:sql:getBackupShortTermRetentionPolicy", args ?? new GetBackupShortTermRetentionPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackupShortTermRetentionPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("databaseName", required: true)]
        public string DatabaseName { get; set; } = null!;

        /// <summary>
        /// The policy name. Should always be "default".
        /// </summary>
        [Input("policyName", required: true)]
        public string PolicyName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public string ServerName { get; set; } = null!;

        public GetBackupShortTermRetentionPolicyArgs()
        {
        }
        public static new GetBackupShortTermRetentionPolicyArgs Empty => new GetBackupShortTermRetentionPolicyArgs();
    }

    public sealed class GetBackupShortTermRetentionPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The policy name. Should always be "default".
        /// </summary>
        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        public GetBackupShortTermRetentionPolicyInvokeArgs()
        {
        }
        public static new GetBackupShortTermRetentionPolicyInvokeArgs Empty => new GetBackupShortTermRetentionPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetBackupShortTermRetentionPolicyResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The differential backup interval in hours. This is how many interval hours between each differential backup will be supported. This is only applicable to live databases but not dropped databases.
        /// </summary>
        public readonly int? DiffBackupIntervalInHours;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The backup retention period in days. This is how many days Point-in-Time Restore will be supported.
        /// </summary>
        public readonly int? RetentionDays;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetBackupShortTermRetentionPolicyResult(
            string azureApiVersion,

            int? diffBackupIntervalInHours,

            string id,

            string name,

            int? retentionDays,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            DiffBackupIntervalInHours = diffBackupIntervalInHours;
            Id = id;
            Name = name;
            RetentionDays = retentionDays;
            Type = type;
        }
    }
}
