// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch
{
    public static class GetApplicationPackage
    {
        /// <summary>
        /// Gets information about the specified application package.
        /// 
        /// Uses Azure REST API version 2024-07-01.
        /// 
        /// Other available API versions: 2023-05-01, 2023-11-01, 2024-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native batch [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetApplicationPackageResult> InvokeAsync(GetApplicationPackageArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApplicationPackageResult>("azure-native:batch:getApplicationPackage", args ?? new GetApplicationPackageArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about the specified application package.
        /// 
        /// Uses Azure REST API version 2024-07-01.
        /// 
        /// Other available API versions: 2023-05-01, 2023-11-01, 2024-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native batch [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetApplicationPackageResult> Invoke(GetApplicationPackageInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationPackageResult>("azure-native:batch:getApplicationPackage", args ?? new GetApplicationPackageInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about the specified application package.
        /// 
        /// Uses Azure REST API version 2024-07-01.
        /// 
        /// Other available API versions: 2023-05-01, 2023-11-01, 2024-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native batch [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetApplicationPackageResult> Invoke(GetApplicationPackageInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationPackageResult>("azure-native:batch:getApplicationPackage", args ?? new GetApplicationPackageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationPackageArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Batch account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the application. This must be unique within the account.
        /// </summary>
        [Input("applicationName", required: true)]
        public string ApplicationName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the Batch account.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The version of the application.
        /// </summary>
        [Input("versionName", required: true)]
        public string VersionName { get; set; } = null!;

        public GetApplicationPackageArgs()
        {
        }
        public static new GetApplicationPackageArgs Empty => new GetApplicationPackageArgs();
    }

    public sealed class GetApplicationPackageInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Batch account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the application. This must be unique within the account.
        /// </summary>
        [Input("applicationName", required: true)]
        public Input<string> ApplicationName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the Batch account.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The version of the application.
        /// </summary>
        [Input("versionName", required: true)]
        public Input<string> VersionName { get; set; } = null!;

        public GetApplicationPackageInvokeArgs()
        {
        }
        public static new GetApplicationPackageInvokeArgs Empty => new GetApplicationPackageInvokeArgs();
    }


    [OutputType]
    public sealed class GetApplicationPackageResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The ETag of the resource, used for concurrency statements.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The format of the application package, if the package is active.
        /// </summary>
        public readonly string Format;
        /// <summary>
        /// The ID of the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The time at which the package was last activated, if the package is active.
        /// </summary>
        public readonly string LastActivationTime;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current state of the application package.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The URL for the application package in Azure Storage.
        /// </summary>
        public readonly string StorageUrl;
        /// <summary>
        /// The UTC time at which the Azure Storage URL will expire.
        /// </summary>
        public readonly string StorageUrlExpiry;
        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetApplicationPackageResult(
            string azureApiVersion,

            string etag,

            string format,

            string id,

            string lastActivationTime,

            string name,

            string state,

            string storageUrl,

            string storageUrlExpiry,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Etag = etag;
            Format = format;
            Id = id;
            LastActivationTime = lastActivationTime;
            Name = name;
            State = state;
            StorageUrl = storageUrl;
            StorageUrlExpiry = storageUrlExpiry;
            Tags = tags;
            Type = type;
        }
    }
}
