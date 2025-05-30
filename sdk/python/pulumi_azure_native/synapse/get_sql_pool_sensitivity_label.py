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

__all__ = [
    'GetSqlPoolSensitivityLabelResult',
    'AwaitableGetSqlPoolSensitivityLabelResult',
    'get_sql_pool_sensitivity_label',
    'get_sql_pool_sensitivity_label_output',
]

@pulumi.output_type
class GetSqlPoolSensitivityLabelResult:
    """
    A sensitivity label.
    """
    def __init__(__self__, azure_api_version=None, column_name=None, id=None, information_type=None, information_type_id=None, is_disabled=None, label_id=None, label_name=None, managed_by=None, name=None, rank=None, schema_name=None, table_name=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if column_name and not isinstance(column_name, str):
            raise TypeError("Expected argument 'column_name' to be a str")
        pulumi.set(__self__, "column_name", column_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if information_type and not isinstance(information_type, str):
            raise TypeError("Expected argument 'information_type' to be a str")
        pulumi.set(__self__, "information_type", information_type)
        if information_type_id and not isinstance(information_type_id, str):
            raise TypeError("Expected argument 'information_type_id' to be a str")
        pulumi.set(__self__, "information_type_id", information_type_id)
        if is_disabled and not isinstance(is_disabled, bool):
            raise TypeError("Expected argument 'is_disabled' to be a bool")
        pulumi.set(__self__, "is_disabled", is_disabled)
        if label_id and not isinstance(label_id, str):
            raise TypeError("Expected argument 'label_id' to be a str")
        pulumi.set(__self__, "label_id", label_id)
        if label_name and not isinstance(label_name, str):
            raise TypeError("Expected argument 'label_name' to be a str")
        pulumi.set(__self__, "label_name", label_name)
        if managed_by and not isinstance(managed_by, str):
            raise TypeError("Expected argument 'managed_by' to be a str")
        pulumi.set(__self__, "managed_by", managed_by)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if rank and not isinstance(rank, str):
            raise TypeError("Expected argument 'rank' to be a str")
        pulumi.set(__self__, "rank", rank)
        if schema_name and not isinstance(schema_name, str):
            raise TypeError("Expected argument 'schema_name' to be a str")
        pulumi.set(__self__, "schema_name", schema_name)
        if table_name and not isinstance(table_name, str):
            raise TypeError("Expected argument 'table_name' to be a str")
        pulumi.set(__self__, "table_name", table_name)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="columnName")
    def column_name(self) -> builtins.str:
        """
        The column name.
        """
        return pulumi.get(self, "column_name")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="informationType")
    def information_type(self) -> Optional[builtins.str]:
        """
        The information type.
        """
        return pulumi.get(self, "information_type")

    @property
    @pulumi.getter(name="informationTypeId")
    def information_type_id(self) -> Optional[builtins.str]:
        """
        The information type ID.
        """
        return pulumi.get(self, "information_type_id")

    @property
    @pulumi.getter(name="isDisabled")
    def is_disabled(self) -> builtins.bool:
        """
        Is sensitivity recommendation disabled. Applicable for recommended sensitivity label only. Specifies whether the sensitivity recommendation on this column is disabled (dismissed) or not.
        """
        return pulumi.get(self, "is_disabled")

    @property
    @pulumi.getter(name="labelId")
    def label_id(self) -> Optional[builtins.str]:
        """
        The label ID.
        """
        return pulumi.get(self, "label_id")

    @property
    @pulumi.getter(name="labelName")
    def label_name(self) -> Optional[builtins.str]:
        """
        The label name.
        """
        return pulumi.get(self, "label_name")

    @property
    @pulumi.getter(name="managedBy")
    def managed_by(self) -> builtins.str:
        """
        managed by
        """
        return pulumi.get(self, "managed_by")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rank(self) -> Optional[builtins.str]:
        return pulumi.get(self, "rank")

    @property
    @pulumi.getter(name="schemaName")
    def schema_name(self) -> builtins.str:
        """
        The schema name.
        """
        return pulumi.get(self, "schema_name")

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> builtins.str:
        """
        The table name.
        """
        return pulumi.get(self, "table_name")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetSqlPoolSensitivityLabelResult(GetSqlPoolSensitivityLabelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSqlPoolSensitivityLabelResult(
            azure_api_version=self.azure_api_version,
            column_name=self.column_name,
            id=self.id,
            information_type=self.information_type,
            information_type_id=self.information_type_id,
            is_disabled=self.is_disabled,
            label_id=self.label_id,
            label_name=self.label_name,
            managed_by=self.managed_by,
            name=self.name,
            rank=self.rank,
            schema_name=self.schema_name,
            table_name=self.table_name,
            type=self.type)


def get_sql_pool_sensitivity_label(column_name: Optional[builtins.str] = None,
                                   resource_group_name: Optional[builtins.str] = None,
                                   schema_name: Optional[builtins.str] = None,
                                   sensitivity_label_source: Optional[builtins.str] = None,
                                   sql_pool_name: Optional[builtins.str] = None,
                                   table_name: Optional[builtins.str] = None,
                                   workspace_name: Optional[builtins.str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSqlPoolSensitivityLabelResult:
    """
    Gets the sensitivity label of a given column

    Uses Azure REST API version 2021-06-01.

    Other available API versions: 2021-04-01-preview, 2021-05-01, 2021-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native synapse [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str column_name: The name of the column.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str schema_name: The name of the schema.
    :param builtins.str sensitivity_label_source: The source of the sensitivity label.
    :param builtins.str sql_pool_name: SQL pool name
    :param builtins.str table_name: The name of the table.
    :param builtins.str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['columnName'] = column_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['schemaName'] = schema_name
    __args__['sensitivityLabelSource'] = sensitivity_label_source
    __args__['sqlPoolName'] = sql_pool_name
    __args__['tableName'] = table_name
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:synapse:getSqlPoolSensitivityLabel', __args__, opts=opts, typ=GetSqlPoolSensitivityLabelResult).value

    return AwaitableGetSqlPoolSensitivityLabelResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        column_name=pulumi.get(__ret__, 'column_name'),
        id=pulumi.get(__ret__, 'id'),
        information_type=pulumi.get(__ret__, 'information_type'),
        information_type_id=pulumi.get(__ret__, 'information_type_id'),
        is_disabled=pulumi.get(__ret__, 'is_disabled'),
        label_id=pulumi.get(__ret__, 'label_id'),
        label_name=pulumi.get(__ret__, 'label_name'),
        managed_by=pulumi.get(__ret__, 'managed_by'),
        name=pulumi.get(__ret__, 'name'),
        rank=pulumi.get(__ret__, 'rank'),
        schema_name=pulumi.get(__ret__, 'schema_name'),
        table_name=pulumi.get(__ret__, 'table_name'),
        type=pulumi.get(__ret__, 'type'))
def get_sql_pool_sensitivity_label_output(column_name: Optional[pulumi.Input[builtins.str]] = None,
                                          resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                          schema_name: Optional[pulumi.Input[builtins.str]] = None,
                                          sensitivity_label_source: Optional[pulumi.Input[builtins.str]] = None,
                                          sql_pool_name: Optional[pulumi.Input[builtins.str]] = None,
                                          table_name: Optional[pulumi.Input[builtins.str]] = None,
                                          workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSqlPoolSensitivityLabelResult]:
    """
    Gets the sensitivity label of a given column

    Uses Azure REST API version 2021-06-01.

    Other available API versions: 2021-04-01-preview, 2021-05-01, 2021-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native synapse [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str column_name: The name of the column.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str schema_name: The name of the schema.
    :param builtins.str sensitivity_label_source: The source of the sensitivity label.
    :param builtins.str sql_pool_name: SQL pool name
    :param builtins.str table_name: The name of the table.
    :param builtins.str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['columnName'] = column_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['schemaName'] = schema_name
    __args__['sensitivityLabelSource'] = sensitivity_label_source
    __args__['sqlPoolName'] = sql_pool_name
    __args__['tableName'] = table_name
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:synapse:getSqlPoolSensitivityLabel', __args__, opts=opts, typ=GetSqlPoolSensitivityLabelResult)
    return __ret__.apply(lambda __response__: GetSqlPoolSensitivityLabelResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        column_name=pulumi.get(__response__, 'column_name'),
        id=pulumi.get(__response__, 'id'),
        information_type=pulumi.get(__response__, 'information_type'),
        information_type_id=pulumi.get(__response__, 'information_type_id'),
        is_disabled=pulumi.get(__response__, 'is_disabled'),
        label_id=pulumi.get(__response__, 'label_id'),
        label_name=pulumi.get(__response__, 'label_name'),
        managed_by=pulumi.get(__response__, 'managed_by'),
        name=pulumi.get(__response__, 'name'),
        rank=pulumi.get(__response__, 'rank'),
        schema_name=pulumi.get(__response__, 'schema_name'),
        table_name=pulumi.get(__response__, 'table_name'),
        type=pulumi.get(__response__, 'type')))
