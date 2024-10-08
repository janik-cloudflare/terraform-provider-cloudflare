---
page_title: "cloudflare_bot_management Resource - Cloudflare"
subcategory: ""
description: |-
  Provides a resource to configure Bot Management.
  Specifically, this resource can be used to manage:
  Bot Fight ModeSuper Bot Fight ModeBot Management for Enterprise
---

# cloudflare_bot_management (Resource)

Provides a resource to configure Bot Management.

Specifically, this resource can be used to manage:

- **Bot Fight Mode**
- **Super Bot Fight Mode**
- **Bot Management for Enterprise**

## Example Usage

```terraform
resource "cloudflare_bot_management" "example" {
  zone_id                         = "0da42c8d2132a9ddaf714f9e7c920711"
  enable_js                       = true
  sbfm_definitely_automated       = "block"
  sbfm_likely_automated           = "managed_challenge"
  sbfm_verified_bots              = "allow"
  sbfm_static_resource_protection = false
  optimize_wordpress              = true
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `zone_id` (String) The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**

### Optional

- `ai_bots_protection` (String) Enable rule to block AI Scrapers and Crawlers.
- `auto_update_model` (Boolean) Automatically update to the newest bot detection models created by Cloudflare as they are released. [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes).
- `enable_js` (Boolean) Use lightweight, invisible JavaScript detections to improve Bot Management. [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
- `fight_mode` (Boolean) Whether to enable Bot Fight Mode.
- `optimize_wordpress` (Boolean) Whether to optimize Super Bot Fight Mode protections for Wordpress.
- `sbfm_definitely_automated` (String) Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
- `sbfm_likely_automated` (String) Super Bot Fight Mode (SBFM) action to take on likely automated requests.
- `sbfm_static_resource_protection` (Boolean) Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if static resources on your application need bot protection. Note: Static resource protection can also result in legitimate traffic being blocked.
- `sbfm_verified_bots` (String) Super Bot Fight Mode (SBFM) action to take on verified bots requests.
- `suppress_session_score` (Boolean) Whether to disable tracking the highest bot score for a session in the Bot Management cookie.

### Read-Only

- `id` (String) The ID of this resource.
- `using_latest_model` (Boolean) A read-only field that indicates whether the zone currently is running the latest ML model.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_bot_management.example <zone_id>
```
