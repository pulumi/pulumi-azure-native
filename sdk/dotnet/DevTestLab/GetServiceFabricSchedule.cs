// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevTestLab
{
    public static class GetServiceFabricSchedule
    {
        /// <summary>
        /// Get schedule.
        /// 
        /// Uses Azure REST API version 2018-09-15.
        /// </summary>
        public static Task<GetServiceFabricScheduleResult> InvokeAsync(GetServiceFabricScheduleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServiceFabricScheduleResult>("azure-native:devtestlab:getServiceFabricSchedule", args ?? new GetServiceFabricScheduleArgs(), options.WithDefaults());

        /// <summary>
        /// Get schedule.
        /// 
        /// Uses Azure REST API version 2018-09-15.
        /// </summary>
        public static Output<GetServiceFabricScheduleResult> Invoke(GetServiceFabricScheduleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceFabricScheduleResult>("azure-native:devtestlab:getServiceFabricSchedule", args ?? new GetServiceFabricScheduleInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get schedule.
        /// 
        /// Uses Azure REST API version 2018-09-15.
        /// </summary>
        public static Output<GetServiceFabricScheduleResult> Invoke(GetServiceFabricScheduleInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceFabricScheduleResult>("azure-native:devtestlab:getServiceFabricSchedule", args ?? new GetServiceFabricScheduleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceFabricScheduleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the $expand query. Example: 'properties($select=status)'
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// labs
        /// </summary>
        [Input("labName", required: true)]
        public string LabName { get; set; } = null!;

        /// <summary>
        /// The name of the Schedule
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// servicefabrics
        /// </summary>
        [Input("serviceFabricName", required: true)]
        public string ServiceFabricName { get; set; } = null!;

        /// <summary>
        /// users
        /// </summary>
        [Input("userName", required: true)]
        public string UserName { get; set; } = null!;

        public GetServiceFabricScheduleArgs()
        {
        }
        public static new GetServiceFabricScheduleArgs Empty => new GetServiceFabricScheduleArgs();
    }

    public sealed class GetServiceFabricScheduleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the $expand query. Example: 'properties($select=status)'
        /// </summary>
        [Input("expand")]
        public Input<string>? Expand { get; set; }

        /// <summary>
        /// labs
        /// </summary>
        [Input("labName", required: true)]
        public Input<string> LabName { get; set; } = null!;

        /// <summary>
        /// The name of the Schedule
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// servicefabrics
        /// </summary>
        [Input("serviceFabricName", required: true)]
        public Input<string> ServiceFabricName { get; set; } = null!;

        /// <summary>
        /// users
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public GetServiceFabricScheduleInvokeArgs()
        {
        }
        public static new GetServiceFabricScheduleInvokeArgs Empty => new GetServiceFabricScheduleInvokeArgs();
    }


    [OutputType]
    public sealed class GetServiceFabricScheduleResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The creation date of the schedule.
        /// </summary>
        public readonly string CreatedDate;
        /// <summary>
        /// If the schedule will occur once each day of the week, specify the daily recurrence.
        /// </summary>
        public readonly Outputs.DayDetailsResponse? DailyRecurrence;
        /// <summary>
        /// If the schedule will occur multiple times a day, specify the hourly recurrence.
        /// </summary>
        public readonly Outputs.HourDetailsResponse? HourlyRecurrence;
        /// <summary>
        /// The identifier of the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The location of the resource.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Notification settings.
        /// </summary>
        public readonly Outputs.NotificationSettingsResponse? NotificationSettings;
        /// <summary>
        /// The provisioning status of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The status of the schedule (i.e. Enabled, Disabled)
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The resource ID to which the schedule belongs
        /// </summary>
        public readonly string? TargetResourceId;
        /// <summary>
        /// The task type of the schedule (e.g. LabVmsShutdownTask, LabVmAutoStart).
        /// </summary>
        public readonly string? TaskType;
        /// <summary>
        /// The time zone ID (e.g. China Standard Time, Greenland Standard Time, Pacific Standard time, etc.). The possible values for this property can be found in `IReadOnlyCollection&lt;string&gt; TimeZoneConverter.TZConvert.KnownWindowsTimeZoneIds` (https://github.com/mattjohnsonpint/TimeZoneConverter/blob/main/README.md)
        /// </summary>
        public readonly string? TimeZoneId;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The unique immutable identifier of a resource (Guid).
        /// </summary>
        public readonly string UniqueIdentifier;
        /// <summary>
        /// If the schedule will occur only some days of the week, specify the weekly recurrence.
        /// </summary>
        public readonly Outputs.WeekDetailsResponse? WeeklyRecurrence;

        [OutputConstructor]
        private GetServiceFabricScheduleResult(
            string azureApiVersion,

            string createdDate,

            Outputs.DayDetailsResponse? dailyRecurrence,

            Outputs.HourDetailsResponse? hourlyRecurrence,

            string id,

            string? location,

            string name,

            Outputs.NotificationSettingsResponse? notificationSettings,

            string provisioningState,

            string? status,

            ImmutableDictionary<string, string>? tags,

            string? targetResourceId,

            string? taskType,

            string? timeZoneId,

            string type,

            string uniqueIdentifier,

            Outputs.WeekDetailsResponse? weeklyRecurrence)
        {
            AzureApiVersion = azureApiVersion;
            CreatedDate = createdDate;
            DailyRecurrence = dailyRecurrence;
            HourlyRecurrence = hourlyRecurrence;
            Id = id;
            Location = location;
            Name = name;
            NotificationSettings = notificationSettings;
            ProvisioningState = provisioningState;
            Status = status;
            Tags = tags;
            TargetResourceId = targetResourceId;
            TaskType = taskType;
            TimeZoneId = timeZoneId;
            Type = type;
            UniqueIdentifier = uniqueIdentifier;
            WeeklyRecurrence = weeklyRecurrence;
        }
    }
}
