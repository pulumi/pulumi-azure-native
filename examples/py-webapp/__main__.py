from pulumi_azure_native import resources
from pulumi_azure_native import web

rg = resources.ResourceGroup("rg")

plan = web.AppServicePlan("plan",
    resource_group_name = rg.name,
    sku = web.SkuDescriptionArgs(
        name = "S1",
        capacity = 1
    ),
)

app = web.WebApp("app",
    resource_group_name = rg.name,
    server_farm_id = plan.id,
    kind = "app",
    # Set some random properties in siteConfig that we know are not the default
    # and will not be returned by "GET /sites/{siteName}".
    site_config = web.SiteConfigArgs(
        default_documents=["pulumi.html"],
        ip_security_restrictions = [web.IpSecurityRestrictionArgs(
            ip_address = "198.51.100.0/22",
            action = "Allow",
            priority = 100,
            name = "pulumi"
        )]
    ),
)

slot = web.WebAppSlot("app",
    name = app.name,
    slot = "staging",
    resource_group_name = rg.name,
    server_farm_id = plan.id,
    kind = "app",
    # Set some random properties in siteConfig that we know are not the default
    # and will not be returned by "GET /sites/{siteName}".
    site_config = web.SiteConfigArgs(
        default_documents=["pulumi.html"],
        ip_security_restrictions = [web.IpSecurityRestrictionArgs(
            ip_address = "198.51.100.0/22",
            action = "Allow",
            priority = 100,
            name = "pulumi"
        )]
    ),
)