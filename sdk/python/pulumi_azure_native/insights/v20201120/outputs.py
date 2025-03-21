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

__all__ = [
    'WorkbookTemplateGalleryResponse',
    'WorkbookTemplateLocalizedGalleryResponse',
]

@pulumi.output_type
class WorkbookTemplateGalleryResponse(dict):
    """
    Gallery information for a workbook template.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "resourceType":
            suggest = "resource_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in WorkbookTemplateGalleryResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        WorkbookTemplateGalleryResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        WorkbookTemplateGalleryResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 category: Optional[str] = None,
                 name: Optional[str] = None,
                 order: Optional[int] = None,
                 resource_type: Optional[str] = None,
                 type: Optional[str] = None):
        """
        Gallery information for a workbook template.
        :param str category: Category for the gallery.
        :param str name: Name of the workbook template in the gallery.
        :param int order: Order of the template within the gallery.
        :param str resource_type: Azure resource type supported by the gallery.
        :param str type: Type of workbook supported by the workbook template.
        """
        if category is not None:
            pulumi.set(__self__, "category", category)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if order is not None:
            pulumi.set(__self__, "order", order)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def category(self) -> Optional[str]:
        """
        Category for the gallery.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the workbook template in the gallery.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def order(self) -> Optional[int]:
        """
        Order of the template within the gallery.
        """
        return pulumi.get(self, "order")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[str]:
        """
        Azure resource type supported by the gallery.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Type of workbook supported by the workbook template.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class WorkbookTemplateLocalizedGalleryResponse(dict):
    """
    Localized template data and gallery information.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "templateData":
            suggest = "template_data"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in WorkbookTemplateLocalizedGalleryResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        WorkbookTemplateLocalizedGalleryResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        WorkbookTemplateLocalizedGalleryResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 galleries: Optional[Sequence['outputs.WorkbookTemplateGalleryResponse']] = None,
                 template_data: Optional[Any] = None):
        """
        Localized template data and gallery information.
        :param Sequence['WorkbookTemplateGalleryResponse'] galleries: Workbook galleries supported by the template.
        :param Any template_data: Valid JSON object containing workbook template payload.
        """
        if galleries is not None:
            pulumi.set(__self__, "galleries", galleries)
        if template_data is not None:
            pulumi.set(__self__, "template_data", template_data)

    @property
    @pulumi.getter
    def galleries(self) -> Optional[Sequence['outputs.WorkbookTemplateGalleryResponse']]:
        """
        Workbook galleries supported by the template.
        """
        return pulumi.get(self, "galleries")

    @property
    @pulumi.getter(name="templateData")
    def template_data(self) -> Optional[Any]:
        """
        Valid JSON object containing workbook template payload.
        """
        return pulumi.get(self, "template_data")


