// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices
{
    public static class GetMachineLearningDatastore
    {
        /// <summary>
        /// Get a Datastore by name.
        /// 
        /// Uses Azure REST API version 2020-05-01-preview.
        /// </summary>
        public static Task<GetMachineLearningDatastoreResult> InvokeAsync(GetMachineLearningDatastoreArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMachineLearningDatastoreResult>("azure-native:machinelearningservices:getMachineLearningDatastore", args ?? new GetMachineLearningDatastoreArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Datastore by name.
        /// 
        /// Uses Azure REST API version 2020-05-01-preview.
        /// </summary>
        public static Output<GetMachineLearningDatastoreResult> Invoke(GetMachineLearningDatastoreInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMachineLearningDatastoreResult>("azure-native:machinelearningservices:getMachineLearningDatastore", args ?? new GetMachineLearningDatastoreInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Datastore by name.
        /// 
        /// Uses Azure REST API version 2020-05-01-preview.
        /// </summary>
        public static Output<GetMachineLearningDatastoreResult> Invoke(GetMachineLearningDatastoreInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMachineLearningDatastoreResult>("azure-native:machinelearningservices:getMachineLearningDatastore", args ?? new GetMachineLearningDatastoreInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMachineLearningDatastoreArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Datastore name.
        /// </summary>
        [Input("datastoreName", required: true)]
        public string DatastoreName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group in which workspace is located.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetMachineLearningDatastoreArgs()
        {
        }
        public static new GetMachineLearningDatastoreArgs Empty => new GetMachineLearningDatastoreArgs();
    }

    public sealed class GetMachineLearningDatastoreInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Datastore name.
        /// </summary>
        [Input("datastoreName", required: true)]
        public Input<string> DatastoreName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group in which workspace is located.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetMachineLearningDatastoreInvokeArgs()
        {
        }
        public static new GetMachineLearningDatastoreInvokeArgs Empty => new GetMachineLearningDatastoreInvokeArgs();
    }


    [OutputType]
    public sealed class GetMachineLearningDatastoreResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Specifies the resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The identity of the resource.
        /// </summary>
        public readonly Outputs.IdentityResponse? Identity;
        /// <summary>
        /// Specifies the location of the resource.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Specifies the name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Datastore properties
        /// </summary>
        public readonly Outputs.DatastoreResponse Properties;
        /// <summary>
        /// The sku of the workspace.
        /// </summary>
        public readonly Outputs.SkuResponse? Sku;
        /// <summary>
        /// Contains resource tags defined as key/value pairs.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Specifies the type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetMachineLearningDatastoreResult(
            string azureApiVersion,

            string id,

            Outputs.IdentityResponse? identity,

            string? location,

            string name,

            Outputs.DatastoreResponse properties,

            Outputs.SkuResponse? sku,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Identity = identity;
            Location = location;
            Name = name;
            Properties = properties;
            Sku = sku;
            Tags = tags;
            Type = type;
        }
    }
}
