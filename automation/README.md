# automation

CLI tools for maintaining the Pulumi Azure Native provider.

## Running

```bash
cd automation
go run .
```

## Tools

### Version Tracking Updates

Compares the tracking versions in `versions/v3-spec.yaml` against the latest available API versions in `versions/az-provider-list.json`, so you can see at a glance which modules are behind.

| Indicator | Meaning |
|-----------|---------|
| `↑ bump`  | A newer stable API version is available |
| `✓`       | Tracking version matches the latest |
| `?`       | Module has no tracking version or is not in the provider list |

**Keys:** `j`/`k` scroll · `/` filter by module name · `Esc` clear filter · `q` back to menu
