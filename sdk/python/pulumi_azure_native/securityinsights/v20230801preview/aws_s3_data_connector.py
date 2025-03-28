# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
from ... import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['AwsS3DataConnectorArgs', 'AwsS3DataConnector']

@pulumi.input_type
class AwsS3DataConnectorArgs:
    def __init__(__self__, *,
                 data_types: pulumi.Input['AwsS3DataConnectorDataTypesArgs'],
                 destination_table: pulumi.Input[str],
                 kind: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 role_arn: pulumi.Input[str],
                 sqs_urls: pulumi.Input[Sequence[pulumi.Input[str]]],
                 workspace_name: pulumi.Input[str],
                 data_connector_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AwsS3DataConnector resource.
        :param pulumi.Input['AwsS3DataConnectorDataTypesArgs'] data_types: The available data types for the connector.
        :param pulumi.Input[str] destination_table: The logs destination table name in LogAnalytics.
        :param pulumi.Input[str] kind: The kind of the data connector
               Expected value is 'AmazonWebServicesS3'.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[str] role_arn: The Aws Role Arn that is used to access the Aws account.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] sqs_urls: The AWS sqs urls for the connector.
        :param pulumi.Input[str] workspace_name: The name of the workspace.
        :param pulumi.Input[str] data_connector_id: Connector ID
        """
        pulumi.set(__self__, "data_types", data_types)
        pulumi.set(__self__, "destination_table", destination_table)
        pulumi.set(__self__, "kind", 'AmazonWebServicesS3')
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "role_arn", role_arn)
        pulumi.set(__self__, "sqs_urls", sqs_urls)
        pulumi.set(__self__, "workspace_name", workspace_name)
        if data_connector_id is not None:
            pulumi.set(__self__, "data_connector_id", data_connector_id)

    @property
    @pulumi.getter(name="dataTypes")
    def data_types(self) -> pulumi.Input['AwsS3DataConnectorDataTypesArgs']:
        """
        The available data types for the connector.
        """
        return pulumi.get(self, "data_types")

    @data_types.setter
    def data_types(self, value: pulumi.Input['AwsS3DataConnectorDataTypesArgs']):
        pulumi.set(self, "data_types", value)

    @property
    @pulumi.getter(name="destinationTable")
    def destination_table(self) -> pulumi.Input[str]:
        """
        The logs destination table name in LogAnalytics.
        """
        return pulumi.get(self, "destination_table")

    @destination_table.setter
    def destination_table(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_table", value)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[str]:
        """
        The kind of the data connector
        Expected value is 'AmazonWebServicesS3'.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: pulumi.Input[str]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group. The name is case insensitive.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        """
        The Aws Role Arn that is used to access the Aws account.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="sqsUrls")
    def sqs_urls(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The AWS sqs urls for the connector.
        """
        return pulumi.get(self, "sqs_urls")

    @sqs_urls.setter
    def sqs_urls(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "sqs_urls", value)

    @property
    @pulumi.getter(name="workspaceName")
    def workspace_name(self) -> pulumi.Input[str]:
        """
        The name of the workspace.
        """
        return pulumi.get(self, "workspace_name")

    @workspace_name.setter
    def workspace_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "workspace_name", value)

    @property
    @pulumi.getter(name="dataConnectorId")
    def data_connector_id(self) -> Optional[pulumi.Input[str]]:
        """
        Connector ID
        """
        return pulumi.get(self, "data_connector_id")

    @data_connector_id.setter
    def data_connector_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_connector_id", value)


class AwsS3DataConnector(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_connector_id: Optional[pulumi.Input[str]] = None,
                 data_types: Optional[pulumi.Input[Union['AwsS3DataConnectorDataTypesArgs', 'AwsS3DataConnectorDataTypesArgsDict']]] = None,
                 destination_table: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 sqs_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 workspace_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Represents Amazon Web Services S3 data connector.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_connector_id: Connector ID
        :param pulumi.Input[Union['AwsS3DataConnectorDataTypesArgs', 'AwsS3DataConnectorDataTypesArgsDict']] data_types: The available data types for the connector.
        :param pulumi.Input[str] destination_table: The logs destination table name in LogAnalytics.
        :param pulumi.Input[str] kind: The kind of the data connector
               Expected value is 'AmazonWebServicesS3'.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[str] role_arn: The Aws Role Arn that is used to access the Aws account.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] sqs_urls: The AWS sqs urls for the connector.
        :param pulumi.Input[str] workspace_name: The name of the workspace.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AwsS3DataConnectorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents Amazon Web Services S3 data connector.

        :param str resource_name: The name of the resource.
        :param AwsS3DataConnectorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AwsS3DataConnectorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_connector_id: Optional[pulumi.Input[str]] = None,
                 data_types: Optional[pulumi.Input[Union['AwsS3DataConnectorDataTypesArgs', 'AwsS3DataConnectorDataTypesArgsDict']]] = None,
                 destination_table: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 sqs_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 workspace_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AwsS3DataConnectorArgs.__new__(AwsS3DataConnectorArgs)

            __props__.__dict__["data_connector_id"] = data_connector_id
            if data_types is None and not opts.urn:
                raise TypeError("Missing required property 'data_types'")
            __props__.__dict__["data_types"] = data_types
            if destination_table is None and not opts.urn:
                raise TypeError("Missing required property 'destination_table'")
            __props__.__dict__["destination_table"] = destination_table
            if kind is None and not opts.urn:
                raise TypeError("Missing required property 'kind'")
            __props__.__dict__["kind"] = 'AmazonWebServicesS3'
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__.__dict__["role_arn"] = role_arn
            if sqs_urls is None and not opts.urn:
                raise TypeError("Missing required property 'sqs_urls'")
            __props__.__dict__["sqs_urls"] = sqs_urls
            if workspace_name is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_name'")
            __props__.__dict__["workspace_name"] = workspace_name
            __props__.__dict__["etag"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:securityinsights/v20190101preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20200101:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20210301preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20210901preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20211001:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20211001preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20220101preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20220401preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20220501preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20220601preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20220701preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20220801:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20220801preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20220901preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20221001preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20221101:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20221101preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20221201preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20230201:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20230201preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20230301preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20230401preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20230501preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20230601preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20230701preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20230901preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20231001preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20231101:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20231201preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20240101preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20240301:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20240401preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20240901:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20241001preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20250101preview:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights/v20250301:AwsS3DataConnector"), pulumi.Alias(type_="azure-native:securityinsights:AwsS3DataConnector")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AwsS3DataConnector, __self__).__init__(
            'azure-native:securityinsights/v20230801preview:AwsS3DataConnector',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AwsS3DataConnector':
        """
        Get an existing AwsS3DataConnector resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AwsS3DataConnectorArgs.__new__(AwsS3DataConnectorArgs)

        __props__.__dict__["data_types"] = None
        __props__.__dict__["destination_table"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["role_arn"] = None
        __props__.__dict__["sqs_urls"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return AwsS3DataConnector(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dataTypes")
    def data_types(self) -> pulumi.Output['outputs.AwsS3DataConnectorDataTypesResponse']:
        """
        The available data types for the connector.
        """
        return pulumi.get(self, "data_types")

    @property
    @pulumi.getter(name="destinationTable")
    def destination_table(self) -> pulumi.Output[str]:
        """
        The logs destination table name in LogAnalytics.
        """
        return pulumi.get(self, "destination_table")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[Optional[str]]:
        """
        Etag of the azure resource
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        The kind of the data connector
        Expected value is 'AmazonWebServicesS3'.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        """
        The Aws Role Arn that is used to access the Aws account.
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="sqsUrls")
    def sqs_urls(self) -> pulumi.Output[Sequence[str]]:
        """
        The AWS sqs urls for the connector.
        """
        return pulumi.get(self, "sqs_urls")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

