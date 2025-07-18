// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Storage
{
    public static class GetTable
    {
        /// <summary>
        /// Gets the table with the specified table name, under the specified account if it exists.
        /// 
        /// Uses Azure REST API version 2024-01-01.
        /// 
        /// Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetTableResult> InvokeAsync(GetTableArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTableResult>("azure-native:storage:getTable", args ?? new GetTableArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the table with the specified table name, under the specified account if it exists.
        /// 
        /// Uses Azure REST API version 2024-01-01.
        /// 
        /// Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetTableResult> Invoke(GetTableInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTableResult>("azure-native:storage:getTable", args ?? new GetTableInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the table with the specified table name, under the specified account if it exists.
        /// 
        /// Uses Azure REST API version 2024-01-01.
        /// 
        /// Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetTableResult> Invoke(GetTableInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTableResult>("azure-native:storage:getTable", args ?? new GetTableInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTableArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A table name must be unique within a storage account and must be between 3 and 63 characters.The name must comprise of only alphanumeric characters and it cannot begin with a numeric character.
        /// </summary>
        [Input("tableName", required: true)]
        public string TableName { get; set; } = null!;

        public GetTableArgs()
        {
        }
        public static new GetTableArgs Empty => new GetTableArgs();
    }

    public sealed class GetTableInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A table name must be unique within a storage account and must be between 3 and 63 characters.The name must comprise of only alphanumeric characters and it cannot begin with a numeric character.
        /// </summary>
        [Input("tableName", required: true)]
        public Input<string> TableName { get; set; } = null!;

        public GetTableInvokeArgs()
        {
        }
        public static new GetTableInvokeArgs Empty => new GetTableInvokeArgs();
    }


    [OutputType]
    public sealed class GetTableResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of stored access policies specified on the table.
        /// </summary>
        public readonly ImmutableArray<Outputs.TableSignedIdentifierResponse> SignedIdentifiers;
        /// <summary>
        /// Table name under the specified account
        /// </summary>
        public readonly string TableName;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetTableResult(
            string azureApiVersion,

            string id,

            string name,

            ImmutableArray<Outputs.TableSignedIdentifierResponse> signedIdentifiers,

            string tableName,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            SignedIdentifiers = signedIdentifiers;
            TableName = tableName;
            Type = type;
        }
    }
}
