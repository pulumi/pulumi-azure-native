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

__all__ = ['ConnectorArgs', 'Connector']

@pulumi.input_type
class ConnectorArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[builtins.str],
                 environment_id: pulumi.Input[builtins.str],
                 organization_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 connector_basic_info: Optional[pulumi.Input['ConnectorInfoBaseArgs']] = None,
                 connector_name: Optional[pulumi.Input[builtins.str]] = None,
                 connector_service_type_info: Optional[pulumi.Input[Union['AzureBlobStorageSinkConnectorServiceInfoArgs', 'AzureBlobStorageSourceConnectorServiceInfoArgs', 'AzureCosmosDBSinkConnectorServiceInfoArgs', 'AzureCosmosDBSourceConnectorServiceInfoArgs', 'AzureSynapseAnalyticsSinkConnectorServiceInfoArgs']]] = None,
                 partner_connector_info: Optional[pulumi.Input[Union['KafkaAzureBlobStorageSinkConnectorInfoArgs', 'KafkaAzureBlobStorageSourceConnectorInfoArgs', 'KafkaAzureCosmosDBSinkConnectorInfoArgs', 'KafkaAzureCosmosDBSourceConnectorInfoArgs', 'KafkaAzureSynapseAnalyticsSinkConnectorInfoArgs']]] = None):
        """
        The set of arguments for constructing a Connector resource.
        :param pulumi.Input[builtins.str] cluster_id: Confluent kafka or schema registry cluster id
        :param pulumi.Input[builtins.str] environment_id: Confluent environment id
        :param pulumi.Input[builtins.str] organization_name: Organization resource name
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input['ConnectorInfoBaseArgs'] connector_basic_info: Connector Info Base
        :param pulumi.Input[builtins.str] connector_name: Confluent connector name
        :param pulumi.Input[Union['AzureBlobStorageSinkConnectorServiceInfoArgs', 'AzureBlobStorageSourceConnectorServiceInfoArgs', 'AzureCosmosDBSinkConnectorServiceInfoArgs', 'AzureCosmosDBSourceConnectorServiceInfoArgs', 'AzureSynapseAnalyticsSinkConnectorServiceInfoArgs']] connector_service_type_info: Connector Service type info base properties.
        :param pulumi.Input[Union['KafkaAzureBlobStorageSinkConnectorInfoArgs', 'KafkaAzureBlobStorageSourceConnectorInfoArgs', 'KafkaAzureCosmosDBSinkConnectorInfoArgs', 'KafkaAzureCosmosDBSourceConnectorInfoArgs', 'KafkaAzureSynapseAnalyticsSinkConnectorInfoArgs']] partner_connector_info: The connection information consumed by applications.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "environment_id", environment_id)
        pulumi.set(__self__, "organization_name", organization_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if connector_basic_info is not None:
            pulumi.set(__self__, "connector_basic_info", connector_basic_info)
        if connector_name is not None:
            pulumi.set(__self__, "connector_name", connector_name)
        if connector_service_type_info is not None:
            pulumi.set(__self__, "connector_service_type_info", connector_service_type_info)
        if partner_connector_info is not None:
            pulumi.set(__self__, "partner_connector_info", partner_connector_info)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[builtins.str]:
        """
        Confluent kafka or schema registry cluster id
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Input[builtins.str]:
        """
        Confluent environment id
        """
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "environment_id", value)

    @property
    @pulumi.getter(name="organizationName")
    def organization_name(self) -> pulumi.Input[builtins.str]:
        """
        Organization resource name
        """
        return pulumi.get(self, "organization_name")

    @organization_name.setter
    def organization_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "organization_name", value)

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
    @pulumi.getter(name="connectorBasicInfo")
    def connector_basic_info(self) -> Optional[pulumi.Input['ConnectorInfoBaseArgs']]:
        """
        Connector Info Base
        """
        return pulumi.get(self, "connector_basic_info")

    @connector_basic_info.setter
    def connector_basic_info(self, value: Optional[pulumi.Input['ConnectorInfoBaseArgs']]):
        pulumi.set(self, "connector_basic_info", value)

    @property
    @pulumi.getter(name="connectorName")
    def connector_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Confluent connector name
        """
        return pulumi.get(self, "connector_name")

    @connector_name.setter
    def connector_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "connector_name", value)

    @property
    @pulumi.getter(name="connectorServiceTypeInfo")
    def connector_service_type_info(self) -> Optional[pulumi.Input[Union['AzureBlobStorageSinkConnectorServiceInfoArgs', 'AzureBlobStorageSourceConnectorServiceInfoArgs', 'AzureCosmosDBSinkConnectorServiceInfoArgs', 'AzureCosmosDBSourceConnectorServiceInfoArgs', 'AzureSynapseAnalyticsSinkConnectorServiceInfoArgs']]]:
        """
        Connector Service type info base properties.
        """
        return pulumi.get(self, "connector_service_type_info")

    @connector_service_type_info.setter
    def connector_service_type_info(self, value: Optional[pulumi.Input[Union['AzureBlobStorageSinkConnectorServiceInfoArgs', 'AzureBlobStorageSourceConnectorServiceInfoArgs', 'AzureCosmosDBSinkConnectorServiceInfoArgs', 'AzureCosmosDBSourceConnectorServiceInfoArgs', 'AzureSynapseAnalyticsSinkConnectorServiceInfoArgs']]]):
        pulumi.set(self, "connector_service_type_info", value)

    @property
    @pulumi.getter(name="partnerConnectorInfo")
    def partner_connector_info(self) -> Optional[pulumi.Input[Union['KafkaAzureBlobStorageSinkConnectorInfoArgs', 'KafkaAzureBlobStorageSourceConnectorInfoArgs', 'KafkaAzureCosmosDBSinkConnectorInfoArgs', 'KafkaAzureCosmosDBSourceConnectorInfoArgs', 'KafkaAzureSynapseAnalyticsSinkConnectorInfoArgs']]]:
        """
        The connection information consumed by applications.
        """
        return pulumi.get(self, "partner_connector_info")

    @partner_connector_info.setter
    def partner_connector_info(self, value: Optional[pulumi.Input[Union['KafkaAzureBlobStorageSinkConnectorInfoArgs', 'KafkaAzureBlobStorageSourceConnectorInfoArgs', 'KafkaAzureCosmosDBSinkConnectorInfoArgs', 'KafkaAzureCosmosDBSourceConnectorInfoArgs', 'KafkaAzureSynapseAnalyticsSinkConnectorInfoArgs']]]):
        pulumi.set(self, "partner_connector_info", value)


@pulumi.type_token("azure-native:confluent:Connector")
class Connector(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 connector_basic_info: Optional[pulumi.Input[Union['ConnectorInfoBaseArgs', 'ConnectorInfoBaseArgsDict']]] = None,
                 connector_name: Optional[pulumi.Input[builtins.str]] = None,
                 connector_service_type_info: Optional[pulumi.Input[Union[Union['AzureBlobStorageSinkConnectorServiceInfoArgs', 'AzureBlobStorageSinkConnectorServiceInfoArgsDict'], Union['AzureBlobStorageSourceConnectorServiceInfoArgs', 'AzureBlobStorageSourceConnectorServiceInfoArgsDict'], Union['AzureCosmosDBSinkConnectorServiceInfoArgs', 'AzureCosmosDBSinkConnectorServiceInfoArgsDict'], Union['AzureCosmosDBSourceConnectorServiceInfoArgs', 'AzureCosmosDBSourceConnectorServiceInfoArgsDict'], Union['AzureSynapseAnalyticsSinkConnectorServiceInfoArgs', 'AzureSynapseAnalyticsSinkConnectorServiceInfoArgsDict']]]] = None,
                 environment_id: Optional[pulumi.Input[builtins.str]] = None,
                 organization_name: Optional[pulumi.Input[builtins.str]] = None,
                 partner_connector_info: Optional[pulumi.Input[Union[Union['KafkaAzureBlobStorageSinkConnectorInfoArgs', 'KafkaAzureBlobStorageSinkConnectorInfoArgsDict'], Union['KafkaAzureBlobStorageSourceConnectorInfoArgs', 'KafkaAzureBlobStorageSourceConnectorInfoArgsDict'], Union['KafkaAzureCosmosDBSinkConnectorInfoArgs', 'KafkaAzureCosmosDBSinkConnectorInfoArgsDict'], Union['KafkaAzureCosmosDBSourceConnectorInfoArgs', 'KafkaAzureCosmosDBSourceConnectorInfoArgsDict'], Union['KafkaAzureSynapseAnalyticsSinkConnectorInfoArgs', 'KafkaAzureSynapseAnalyticsSinkConnectorInfoArgsDict']]]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Details of connector record

        Uses Azure REST API version 2024-07-01. In version 2.x of the Azure Native provider, it used API version 2024-07-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] cluster_id: Confluent kafka or schema registry cluster id
        :param pulumi.Input[Union['ConnectorInfoBaseArgs', 'ConnectorInfoBaseArgsDict']] connector_basic_info: Connector Info Base
        :param pulumi.Input[builtins.str] connector_name: Confluent connector name
        :param pulumi.Input[Union[Union['AzureBlobStorageSinkConnectorServiceInfoArgs', 'AzureBlobStorageSinkConnectorServiceInfoArgsDict'], Union['AzureBlobStorageSourceConnectorServiceInfoArgs', 'AzureBlobStorageSourceConnectorServiceInfoArgsDict'], Union['AzureCosmosDBSinkConnectorServiceInfoArgs', 'AzureCosmosDBSinkConnectorServiceInfoArgsDict'], Union['AzureCosmosDBSourceConnectorServiceInfoArgs', 'AzureCosmosDBSourceConnectorServiceInfoArgsDict'], Union['AzureSynapseAnalyticsSinkConnectorServiceInfoArgs', 'AzureSynapseAnalyticsSinkConnectorServiceInfoArgsDict']]] connector_service_type_info: Connector Service type info base properties.
        :param pulumi.Input[builtins.str] environment_id: Confluent environment id
        :param pulumi.Input[builtins.str] organization_name: Organization resource name
        :param pulumi.Input[Union[Union['KafkaAzureBlobStorageSinkConnectorInfoArgs', 'KafkaAzureBlobStorageSinkConnectorInfoArgsDict'], Union['KafkaAzureBlobStorageSourceConnectorInfoArgs', 'KafkaAzureBlobStorageSourceConnectorInfoArgsDict'], Union['KafkaAzureCosmosDBSinkConnectorInfoArgs', 'KafkaAzureCosmosDBSinkConnectorInfoArgsDict'], Union['KafkaAzureCosmosDBSourceConnectorInfoArgs', 'KafkaAzureCosmosDBSourceConnectorInfoArgsDict'], Union['KafkaAzureSynapseAnalyticsSinkConnectorInfoArgs', 'KafkaAzureSynapseAnalyticsSinkConnectorInfoArgsDict']]] partner_connector_info: The connection information consumed by applications.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConnectorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Details of connector record

        Uses Azure REST API version 2024-07-01. In version 2.x of the Azure Native provider, it used API version 2024-07-01.

        :param str resource_name: The name of the resource.
        :param ConnectorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConnectorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 connector_basic_info: Optional[pulumi.Input[Union['ConnectorInfoBaseArgs', 'ConnectorInfoBaseArgsDict']]] = None,
                 connector_name: Optional[pulumi.Input[builtins.str]] = None,
                 connector_service_type_info: Optional[pulumi.Input[Union[Union['AzureBlobStorageSinkConnectorServiceInfoArgs', 'AzureBlobStorageSinkConnectorServiceInfoArgsDict'], Union['AzureBlobStorageSourceConnectorServiceInfoArgs', 'AzureBlobStorageSourceConnectorServiceInfoArgsDict'], Union['AzureCosmosDBSinkConnectorServiceInfoArgs', 'AzureCosmosDBSinkConnectorServiceInfoArgsDict'], Union['AzureCosmosDBSourceConnectorServiceInfoArgs', 'AzureCosmosDBSourceConnectorServiceInfoArgsDict'], Union['AzureSynapseAnalyticsSinkConnectorServiceInfoArgs', 'AzureSynapseAnalyticsSinkConnectorServiceInfoArgsDict']]]] = None,
                 environment_id: Optional[pulumi.Input[builtins.str]] = None,
                 organization_name: Optional[pulumi.Input[builtins.str]] = None,
                 partner_connector_info: Optional[pulumi.Input[Union[Union['KafkaAzureBlobStorageSinkConnectorInfoArgs', 'KafkaAzureBlobStorageSinkConnectorInfoArgsDict'], Union['KafkaAzureBlobStorageSourceConnectorInfoArgs', 'KafkaAzureBlobStorageSourceConnectorInfoArgsDict'], Union['KafkaAzureCosmosDBSinkConnectorInfoArgs', 'KafkaAzureCosmosDBSinkConnectorInfoArgsDict'], Union['KafkaAzureCosmosDBSourceConnectorInfoArgs', 'KafkaAzureCosmosDBSourceConnectorInfoArgsDict'], Union['KafkaAzureSynapseAnalyticsSinkConnectorInfoArgs', 'KafkaAzureSynapseAnalyticsSinkConnectorInfoArgsDict']]]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConnectorArgs.__new__(ConnectorArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["connector_basic_info"] = connector_basic_info
            __props__.__dict__["connector_name"] = connector_name
            __props__.__dict__["connector_service_type_info"] = connector_service_type_info
            if environment_id is None and not opts.urn:
                raise TypeError("Missing required property 'environment_id'")
            __props__.__dict__["environment_id"] = environment_id
            if organization_name is None and not opts.urn:
                raise TypeError("Missing required property 'organization_name'")
            __props__.__dict__["organization_name"] = organization_name
            __props__.__dict__["partner_connector_info"] = partner_connector_info
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:confluent/v20240701:Connector")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Connector, __self__).__init__(
            'azure-native:confluent:Connector',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Connector':
        """
        Get an existing Connector resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ConnectorArgs.__new__(ConnectorArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["connector_basic_info"] = None
        __props__.__dict__["connector_service_type_info"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["partner_connector_info"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return Connector(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="connectorBasicInfo")
    def connector_basic_info(self) -> pulumi.Output[Optional['outputs.ConnectorInfoBaseResponse']]:
        """
        Connector Info Base
        """
        return pulumi.get(self, "connector_basic_info")

    @property
    @pulumi.getter(name="connectorServiceTypeInfo")
    def connector_service_type_info(self) -> pulumi.Output[Optional[Any]]:
        """
        Connector Service type info base properties.
        """
        return pulumi.get(self, "connector_service_type_info")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="partnerConnectorInfo")
    def partner_connector_info(self) -> pulumi.Output[Optional[Any]]:
        """
        The connection information consumed by applications.
        """
        return pulumi.get(self, "partner_connector_info")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

