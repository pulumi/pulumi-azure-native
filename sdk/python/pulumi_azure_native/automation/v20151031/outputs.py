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

__all__ = [
    'RunbookAssociationPropertyResponse',
]

@pulumi.output_type
class RunbookAssociationPropertyResponse(dict):
    """
    The runbook property associated with the entity.
    """
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        The runbook property associated with the entity.
        :param str name: Gets or sets the name of the runbook.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Gets or sets the name of the runbook.
        """
        return pulumi.get(self, "name")


