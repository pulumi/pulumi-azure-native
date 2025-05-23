// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridData
{
    public static class GetJobDefinition
    {
        /// <summary>
        /// This method gets job definition object by name.
        /// 
        /// Uses Azure REST API version 2019-06-01.
        /// </summary>
        public static Task<GetJobDefinitionResult> InvokeAsync(GetJobDefinitionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetJobDefinitionResult>("azure-native:hybriddata:getJobDefinition", args ?? new GetJobDefinitionArgs(), options.WithDefaults());

        /// <summary>
        /// This method gets job definition object by name.
        /// 
        /// Uses Azure REST API version 2019-06-01.
        /// </summary>
        public static Output<GetJobDefinitionResult> Invoke(GetJobDefinitionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetJobDefinitionResult>("azure-native:hybriddata:getJobDefinition", args ?? new GetJobDefinitionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// This method gets job definition object by name.
        /// 
        /// Uses Azure REST API version 2019-06-01.
        /// </summary>
        public static Output<GetJobDefinitionResult> Invoke(GetJobDefinitionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetJobDefinitionResult>("azure-native:hybriddata:getJobDefinition", args ?? new GetJobDefinitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetJobDefinitionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
        /// </summary>
        [Input("dataManagerName", required: true)]
        public string DataManagerName { get; set; } = null!;

        /// <summary>
        /// The data service name of the job definition
        /// </summary>
        [Input("dataServiceName", required: true)]
        public string DataServiceName { get; set; } = null!;

        /// <summary>
        /// The job definition name that is being queried.
        /// </summary>
        [Input("jobDefinitionName", required: true)]
        public string JobDefinitionName { get; set; } = null!;

        /// <summary>
        /// The Resource Group Name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetJobDefinitionArgs()
        {
        }
        public static new GetJobDefinitionArgs Empty => new GetJobDefinitionArgs();
    }

    public sealed class GetJobDefinitionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
        /// </summary>
        [Input("dataManagerName", required: true)]
        public Input<string> DataManagerName { get; set; } = null!;

        /// <summary>
        /// The data service name of the job definition
        /// </summary>
        [Input("dataServiceName", required: true)]
        public Input<string> DataServiceName { get; set; } = null!;

        /// <summary>
        /// The job definition name that is being queried.
        /// </summary>
        [Input("jobDefinitionName", required: true)]
        public Input<string> JobDefinitionName { get; set; } = null!;

        /// <summary>
        /// The Resource Group Name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetJobDefinitionInvokeArgs()
        {
        }
        public static new GetJobDefinitionInvokeArgs Empty => new GetJobDefinitionInvokeArgs();
    }


    [OutputType]
    public sealed class GetJobDefinitionResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
        /// </summary>
        public readonly ImmutableArray<Outputs.CustomerSecretResponse> CustomerSecrets;
        /// <summary>
        /// A generic json used differently by each data service type.
        /// </summary>
        public readonly object? DataServiceInput;
        /// <summary>
        /// Data Sink Id associated to the job definition.
        /// </summary>
        public readonly string DataSinkId;
        /// <summary>
        /// Data Source Id associated to the job definition.
        /// </summary>
        public readonly string DataSourceId;
        /// <summary>
        /// Id of the object.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Last modified time of the job definition.
        /// </summary>
        public readonly string? LastModifiedTime;
        /// <summary>
        /// Name of the object.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// This is the preferred geo location for the job to run.
        /// </summary>
        public readonly string? RunLocation;
        /// <summary>
        /// Schedule for running the job definition
        /// </summary>
        public readonly ImmutableArray<Outputs.ScheduleResponse> Schedules;
        /// <summary>
        /// State of the job definition.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Type of the object.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Enum to detect if user confirmation is required. If not passed will default to NotRequired.
        /// </summary>
        public readonly string? UserConfirmation;

        [OutputConstructor]
        private GetJobDefinitionResult(
            string azureApiVersion,

            ImmutableArray<Outputs.CustomerSecretResponse> customerSecrets,

            object? dataServiceInput,

            string dataSinkId,

            string dataSourceId,

            string id,

            string? lastModifiedTime,

            string name,

            string? runLocation,

            ImmutableArray<Outputs.ScheduleResponse> schedules,

            string state,

            string type,

            string? userConfirmation)
        {
            AzureApiVersion = azureApiVersion;
            CustomerSecrets = customerSecrets;
            DataServiceInput = dataServiceInput;
            DataSinkId = dataSinkId;
            DataSourceId = dataSourceId;
            Id = id;
            LastModifiedTime = lastModifiedTime;
            Name = name;
            RunLocation = runLocation;
            Schedules = schedules;
            State = state;
            Type = type;
            UserConfirmation = userConfirmation;
        }
    }
}
