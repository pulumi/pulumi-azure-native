# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['MonitorArgs', 'Monitor']

@pulumi.input_type
class MonitorArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 app_location: Optional[pulumi.Input[builtins.str]] = None,
                 app_service_plan_configuration: Optional[pulumi.Input['AppServicePlanConfigurationArgs']] = None,
                 identity: Optional[pulumi.Input['ManagedServiceIdentityArgs']] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 log_analytics_workspace_arm_id: Optional[pulumi.Input[builtins.str]] = None,
                 managed_resource_group_configuration: Optional[pulumi.Input['ManagedResourceGroupConfigurationArgs']] = None,
                 monitor_name: Optional[pulumi.Input[builtins.str]] = None,
                 monitor_subnet: Optional[pulumi.Input[builtins.str]] = None,
                 routing_preference: Optional[pulumi.Input[Union[builtins.str, 'RoutingPreference']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 zone_redundancy_preference: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Monitor resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] app_location: The SAP monitor resources will be deployed in the SAP monitoring region. The subnet region should be same as the SAP monitoring region.
        :param pulumi.Input['AppServicePlanConfigurationArgs'] app_service_plan_configuration: App service plan configuration
        :param pulumi.Input['ManagedServiceIdentityArgs'] identity: The managed service identities assigned to this resource.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.str] log_analytics_workspace_arm_id: The ARM ID of the Log Analytics Workspace that is used for SAP monitoring.
        :param pulumi.Input['ManagedResourceGroupConfigurationArgs'] managed_resource_group_configuration: Managed resource group configuration
        :param pulumi.Input[builtins.str] monitor_name: Name of the SAP monitor resource.
        :param pulumi.Input[builtins.str] monitor_subnet: The subnet which the SAP monitor will be deployed in
        :param pulumi.Input[Union[builtins.str, 'RoutingPreference']] routing_preference: Sets the routing preference of the SAP monitor. By default only RFC1918 traffic is routed to the customer VNET.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[builtins.str] zone_redundancy_preference: Sets the preference for zone redundancy on resources created for the SAP monitor. By default resources will be created which do not support zone redundancy.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if app_location is not None:
            pulumi.set(__self__, "app_location", app_location)
        if app_service_plan_configuration is not None:
            pulumi.set(__self__, "app_service_plan_configuration", app_service_plan_configuration)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if log_analytics_workspace_arm_id is not None:
            pulumi.set(__self__, "log_analytics_workspace_arm_id", log_analytics_workspace_arm_id)
        if managed_resource_group_configuration is not None:
            pulumi.set(__self__, "managed_resource_group_configuration", managed_resource_group_configuration)
        if monitor_name is not None:
            pulumi.set(__self__, "monitor_name", monitor_name)
        if monitor_subnet is not None:
            pulumi.set(__self__, "monitor_subnet", monitor_subnet)
        if routing_preference is not None:
            pulumi.set(__self__, "routing_preference", routing_preference)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone_redundancy_preference is not None:
            pulumi.set(__self__, "zone_redundancy_preference", zone_redundancy_preference)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group. The name is case insensitive.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="appLocation")
    def app_location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The SAP monitor resources will be deployed in the SAP monitoring region. The subnet region should be same as the SAP monitoring region.
        """
        return pulumi.get(self, "app_location")

    @app_location.setter
    def app_location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "app_location", value)

    @property
    @pulumi.getter(name="appServicePlanConfiguration")
    def app_service_plan_configuration(self) -> Optional[pulumi.Input['AppServicePlanConfigurationArgs']]:
        """
        App service plan configuration
        """
        return pulumi.get(self, "app_service_plan_configuration")

    @app_service_plan_configuration.setter
    def app_service_plan_configuration(self, value: Optional[pulumi.Input['AppServicePlanConfigurationArgs']]):
        pulumi.set(self, "app_service_plan_configuration", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['ManagedServiceIdentityArgs']]:
        """
        The managed service identities assigned to this resource.
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['ManagedServiceIdentityArgs']]):
        pulumi.set(self, "identity", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="logAnalyticsWorkspaceArmId")
    def log_analytics_workspace_arm_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ARM ID of the Log Analytics Workspace that is used for SAP monitoring.
        """
        return pulumi.get(self, "log_analytics_workspace_arm_id")

    @log_analytics_workspace_arm_id.setter
    def log_analytics_workspace_arm_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "log_analytics_workspace_arm_id", value)

    @property
    @pulumi.getter(name="managedResourceGroupConfiguration")
    def managed_resource_group_configuration(self) -> Optional[pulumi.Input['ManagedResourceGroupConfigurationArgs']]:
        """
        Managed resource group configuration
        """
        return pulumi.get(self, "managed_resource_group_configuration")

    @managed_resource_group_configuration.setter
    def managed_resource_group_configuration(self, value: Optional[pulumi.Input['ManagedResourceGroupConfigurationArgs']]):
        pulumi.set(self, "managed_resource_group_configuration", value)

    @property
    @pulumi.getter(name="monitorName")
    def monitor_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the SAP monitor resource.
        """
        return pulumi.get(self, "monitor_name")

    @monitor_name.setter
    def monitor_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "monitor_name", value)

    @property
    @pulumi.getter(name="monitorSubnet")
    def monitor_subnet(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The subnet which the SAP monitor will be deployed in
        """
        return pulumi.get(self, "monitor_subnet")

    @monitor_subnet.setter
    def monitor_subnet(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "monitor_subnet", value)

    @property
    @pulumi.getter(name="routingPreference")
    def routing_preference(self) -> Optional[pulumi.Input[Union[builtins.str, 'RoutingPreference']]]:
        """
        Sets the routing preference of the SAP monitor. By default only RFC1918 traffic is routed to the customer VNET.
        """
        return pulumi.get(self, "routing_preference")

    @routing_preference.setter
    def routing_preference(self, value: Optional[pulumi.Input[Union[builtins.str, 'RoutingPreference']]]):
        pulumi.set(self, "routing_preference", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="zoneRedundancyPreference")
    def zone_redundancy_preference(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Sets the preference for zone redundancy on resources created for the SAP monitor. By default resources will be created which do not support zone redundancy.
        """
        return pulumi.get(self, "zone_redundancy_preference")

    @zone_redundancy_preference.setter
    def zone_redundancy_preference(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "zone_redundancy_preference", value)


@pulumi.type_token("azure-native:workloads:Monitor")
class Monitor(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_location: Optional[pulumi.Input[builtins.str]] = None,
                 app_service_plan_configuration: Optional[pulumi.Input[Union['AppServicePlanConfigurationArgs', 'AppServicePlanConfigurationArgsDict']]] = None,
                 identity: Optional[pulumi.Input[Union['ManagedServiceIdentityArgs', 'ManagedServiceIdentityArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 log_analytics_workspace_arm_id: Optional[pulumi.Input[builtins.str]] = None,
                 managed_resource_group_configuration: Optional[pulumi.Input[Union['ManagedResourceGroupConfigurationArgs', 'ManagedResourceGroupConfigurationArgsDict']]] = None,
                 monitor_name: Optional[pulumi.Input[builtins.str]] = None,
                 monitor_subnet: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 routing_preference: Optional[pulumi.Input[Union[builtins.str, 'RoutingPreference']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 zone_redundancy_preference: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        SAP monitor info on Azure (ARM properties and SAP monitor properties)

        Uses Azure REST API version 2024-02-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-12-01-preview.

        Other available API versions: 2023-04-01, 2023-10-01-preview, 2023-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native workloads [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] app_location: The SAP monitor resources will be deployed in the SAP monitoring region. The subnet region should be same as the SAP monitoring region.
        :param pulumi.Input[Union['AppServicePlanConfigurationArgs', 'AppServicePlanConfigurationArgsDict']] app_service_plan_configuration: App service plan configuration
        :param pulumi.Input[Union['ManagedServiceIdentityArgs', 'ManagedServiceIdentityArgsDict']] identity: The managed service identities assigned to this resource.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.str] log_analytics_workspace_arm_id: The ARM ID of the Log Analytics Workspace that is used for SAP monitoring.
        :param pulumi.Input[Union['ManagedResourceGroupConfigurationArgs', 'ManagedResourceGroupConfigurationArgsDict']] managed_resource_group_configuration: Managed resource group configuration
        :param pulumi.Input[builtins.str] monitor_name: Name of the SAP monitor resource.
        :param pulumi.Input[builtins.str] monitor_subnet: The subnet which the SAP monitor will be deployed in
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Union[builtins.str, 'RoutingPreference']] routing_preference: Sets the routing preference of the SAP monitor. By default only RFC1918 traffic is routed to the customer VNET.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[builtins.str] zone_redundancy_preference: Sets the preference for zone redundancy on resources created for the SAP monitor. By default resources will be created which do not support zone redundancy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MonitorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        SAP monitor info on Azure (ARM properties and SAP monitor properties)

        Uses Azure REST API version 2024-02-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-12-01-preview.

        Other available API versions: 2023-04-01, 2023-10-01-preview, 2023-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native workloads [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param MonitorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MonitorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_location: Optional[pulumi.Input[builtins.str]] = None,
                 app_service_plan_configuration: Optional[pulumi.Input[Union['AppServicePlanConfigurationArgs', 'AppServicePlanConfigurationArgsDict']]] = None,
                 identity: Optional[pulumi.Input[Union['ManagedServiceIdentityArgs', 'ManagedServiceIdentityArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 log_analytics_workspace_arm_id: Optional[pulumi.Input[builtins.str]] = None,
                 managed_resource_group_configuration: Optional[pulumi.Input[Union['ManagedResourceGroupConfigurationArgs', 'ManagedResourceGroupConfigurationArgsDict']]] = None,
                 monitor_name: Optional[pulumi.Input[builtins.str]] = None,
                 monitor_subnet: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 routing_preference: Optional[pulumi.Input[Union[builtins.str, 'RoutingPreference']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 zone_redundancy_preference: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MonitorArgs.__new__(MonitorArgs)

            __props__.__dict__["app_location"] = app_location
            __props__.__dict__["app_service_plan_configuration"] = app_service_plan_configuration
            __props__.__dict__["identity"] = identity
            __props__.__dict__["location"] = location
            __props__.__dict__["log_analytics_workspace_arm_id"] = log_analytics_workspace_arm_id
            __props__.__dict__["managed_resource_group_configuration"] = managed_resource_group_configuration
            __props__.__dict__["monitor_name"] = monitor_name
            __props__.__dict__["monitor_subnet"] = monitor_subnet
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["routing_preference"] = routing_preference
            __props__.__dict__["tags"] = tags
            __props__.__dict__["zone_redundancy_preference"] = zone_redundancy_preference
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["errors"] = None
            __props__.__dict__["msi_arm_id"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["storage_account_arm_id"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:workloads/v20211201preview:Monitor"), pulumi.Alias(type_="azure-native:workloads/v20221101preview:Monitor"), pulumi.Alias(type_="azure-native:workloads/v20230401:Monitor"), pulumi.Alias(type_="azure-native:workloads/v20231001preview:Monitor"), pulumi.Alias(type_="azure-native:workloads/v20231201preview:Monitor"), pulumi.Alias(type_="azure-native:workloads/v20240201preview:Monitor")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Monitor, __self__).__init__(
            'azure-native:workloads:Monitor',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Monitor':
        """
        Get an existing Monitor resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = MonitorArgs.__new__(MonitorArgs)

        __props__.__dict__["app_location"] = None
        __props__.__dict__["app_service_plan_configuration"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["errors"] = None
        __props__.__dict__["identity"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["log_analytics_workspace_arm_id"] = None
        __props__.__dict__["managed_resource_group_configuration"] = None
        __props__.__dict__["monitor_subnet"] = None
        __props__.__dict__["msi_arm_id"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["routing_preference"] = None
        __props__.__dict__["storage_account_arm_id"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["zone_redundancy_preference"] = None
        return Monitor(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appLocation")
    def app_location(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The SAP monitor resources will be deployed in the SAP monitoring region. The subnet region should be same as the SAP monitoring region.
        """
        return pulumi.get(self, "app_location")

    @property
    @pulumi.getter(name="appServicePlanConfiguration")
    def app_service_plan_configuration(self) -> pulumi.Output[Optional['outputs.AppServicePlanConfigurationResponse']]:
        """
        App service plan configuration
        """
        return pulumi.get(self, "app_service_plan_configuration")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def errors(self) -> pulumi.Output['outputs.ErrorDetailResponse']:
        """
        Defines the SAP monitor errors.
        """
        return pulumi.get(self, "errors")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.ManagedServiceIdentityResponse']]:
        """
        The managed service identities assigned to this resource.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="logAnalyticsWorkspaceArmId")
    def log_analytics_workspace_arm_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The ARM ID of the Log Analytics Workspace that is used for SAP monitoring.
        """
        return pulumi.get(self, "log_analytics_workspace_arm_id")

    @property
    @pulumi.getter(name="managedResourceGroupConfiguration")
    def managed_resource_group_configuration(self) -> pulumi.Output[Optional['outputs.ManagedResourceGroupConfigurationResponse']]:
        """
        Managed resource group configuration
        """
        return pulumi.get(self, "managed_resource_group_configuration")

    @property
    @pulumi.getter(name="monitorSubnet")
    def monitor_subnet(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The subnet which the SAP monitor will be deployed in
        """
        return pulumi.get(self, "monitor_subnet")

    @property
    @pulumi.getter(name="msiArmId")
    def msi_arm_id(self) -> pulumi.Output[builtins.str]:
        """
        The ARM ID of the MSI used for SAP monitoring.
        """
        return pulumi.get(self, "msi_arm_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        State of provisioning of the SAP monitor.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="routingPreference")
    def routing_preference(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Sets the routing preference of the SAP monitor. By default only RFC1918 traffic is routed to the customer VNET.
        """
        return pulumi.get(self, "routing_preference")

    @property
    @pulumi.getter(name="storageAccountArmId")
    def storage_account_arm_id(self) -> pulumi.Output[builtins.str]:
        """
        The ARM ID of the Storage account used for SAP monitoring.
        """
        return pulumi.get(self, "storage_account_arm_id")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="zoneRedundancyPreference")
    def zone_redundancy_preference(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Sets the preference for zone redundancy on resources created for the SAP monitor. By default resources will be created which do not support zone redundancy.
        """
        return pulumi.get(self, "zone_redundancy_preference")

