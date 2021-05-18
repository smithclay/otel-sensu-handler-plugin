[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/smithclay/otel-sensu-handler-plugin)
![goreleaser](https://github.com/smithclay/otel-sensu-handler-plugin/workflows/goreleaser/badge.svg)
[![Go Test](https://github.com/smithclay/otel-sensu-handler-plugin/workflows/Go%20Test/badge.svg)](https://github.com/smithclay/otel-sensu-handler-plugin/actions?query=workflow%3A%22Go+Test%22)
[![goreleaser](https://github.com/smithclay/otel-sensu-handler-plugin/workflows/goreleaser/badge.svg)](https://github.com/smithclay/otel-sensu-handler-plugin/actions?query=workflow%3Agoreleaser)

# otel-sensu-handler-plugin

## Overview
Sensu handler plugin that emits OpenTelemetry metrics. Experimental/work-in-progress.

## Functionality

After successfully creating a project from this template, update the `Config` struct with any
configuration options for the plugin, map those values as plugin options in the variable `options`,
and customize the `checkArgs` and `executeHandler` functions in [main.go][7].

When writing or updating a plugin's README from this template, review the Sensu Community
[plugin README style guide][3] for content suggestions and guidance. Remove everything
prior to `# otel-sensu-handler-plugin` from the generated README file, and add additional context about the
plugin per the style guide.

## Releases with Github Actions

To release a version of your project, simply tag the target sha with a semver release without a `v`
prefix (ex. `1.0.0`). This will trigger the [GitHub action][5] workflow to [build and release][4]
the plugin with goreleaser. Register the asset with [Bonsai][8] to share it with the community!

***

# otel-sensu-handler-plugin

## Table of Contents
- [Overview](#overview)
- [Files](#files)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Handler definition](#handler-definition)
  - [Annotations](#annotations)
- [Installation from source](#installation-from-source)
- [Additional notes](#additional-notes)
- [Contributing](#contributing)

## Overview

The otel-sensu-handler-plugin is a [Sensu Handler][6] that emits OpenTelemety metrics.

## Files

## Usage examples

## Configuration

### Asset registration

[Sensu Assets][10] are the best way to make use of this plugin. If you're not using an asset, please
consider doing so! If you're using sensuctl 5.13 with Sensu Backend 5.13 or later, you can use the
following command to add the asset:

```
sensuctl asset add smithclay/otel-sensu-handler-plugin
```

If you're using an earlier version of sensuctl, you can find the asset on the [Bonsai Asset Index][https://bonsai.sensu.io/assets/smithclay/otel-sensu-handler-plugin].

### Handler definition

```yml
---
type: Handler
api_version: core/v2
metadata:
  name: otel-sensu-handler-plugin
  namespace: default
spec:
  command: otel-sensu-handler-plugin
  type: pipe
  runtime_assets:
  - smithclay/otel-sensu-handler-plugin
```

#### Proxy Support

This handler supports the use of the environment variables HTTP_PROXY,
HTTPS_PROXY, and NO_PROXY (or the lowercase versions thereof). HTTPS_PROXY takes
precedence over HTTP_PROXY for https requests.  The environment values may be
either a complete URL or a "host[:port]", in which case the "http" scheme is assumed.

### Annotations

All arguments for this handler are tunable on a per entity or check basis based on annotations.  The
annotations keyspace for this handler is `sensu.io/plugins/otel-sensu-handler-plugin/config`.

#### Examples

To change the example argument for a particular check, for that checks's metadata add the following:

```yml
type: CheckConfig
api_version: core/v2
metadata:
  annotations:
    sensu.io/plugins/otel-sensu-handler-plugin/config/example-argument: "Example change"
[...]
```

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an Asset. If you would
like to compile and install the plugin from source or contribute to it, download the latest version
or create an executable script from this source.

From the local path of the otel-sensu-handler-plugin repository:

```
go build
```

## Additional notes

## Contributing

For more information about contributing to this plugin, see [Contributing][1].

[1]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[2]: https://github.com/sensu-community/sensu-plugin-sdk
[3]: https://github.com/sensu-plugins/community/blob/master/PLUGIN_STYLEGUIDE.md
[4]: https://github.com/sensu-community/handler-plugin-template/blob/master/.github/workflows/release.yml
[5]: https://github.com/sensu-community/handler-plugin-template/actions
[6]: https://docs.sensu.io/sensu-go/latest/reference/handlers/
[7]: https://github.com/sensu-community/handler-plugin-template/blob/master/main.go
[8]: https://bonsai.sensu.io/
[9]: https://github.com/sensu-community/sensu-plugin-tool
[10]: https://docs.sensu.io/sensu-go/latest/reference/assets/
