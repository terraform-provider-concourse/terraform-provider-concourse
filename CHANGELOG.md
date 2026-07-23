## Changelog

### 8.0.2

`concourse_team` resource: auth entries with an unrecognised first token (e.g.
`"saml:platform-owners"` instead of `"group:saml:platform-owners"`) were
previously silently dropped, causing team auth to be configured with an empty
group list. The provider now returns a clear error at apply time and validates
entries at plan time.

**SAML groups** — use `"group:saml:<group-name>"`:
```hcl
resource "concourse_team" "platform" {
  team_name = "platform"
  owners    = ["group:saml:platform-owners"]
  members   = ["group:saml:platform-members"]
}
```

### 8.0.0

`concourse_pipeline` resource now supports supplying (concourse)
template variables through the `vars` argument. Technically this is a
breaking change if any of your pipelines happen to have any
double-parentheses (`(( ... ))`) references that aren't intended to
be interpreted by concourse.
