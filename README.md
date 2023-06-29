[![Actions Status](https://github.com/Hellthrashers/pulumi-airflow/workflows/release/badge.svg)](https://github.com/Hellthrashers/pulumi-airflow/actions)

# Apache airflow provider

This code defines a Pulumi provider resource for the Airflow package.

## Installation

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/airflow

## Provider and Arguments

The `Provider` resource accepts the following arguments:

- `baseEndpoint`: (Required) The base endpoint URL for the API.
- `oauth2Token`: (Optional) The oauth token to use for API authentication.
- `password`: (Optional) The password to use for API basic authentication.
- `username`: (Optional) The username to use for API basic authentication.

These arguments can be passed as an object of type `ProviderArgs` to the `Provider` constructor.

## Example set provider with Node.js (JavaScript/TypeScript)

### Google Composer 2 Example (OAUTH2 access token)
If you use Composer version 2 changes how the API is accessed by the provider, you can just use the composer airflow web UI endpoint, and you can use an Oauth2 access token.

With service account impersionation

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as airflow from "@pulumi/airflow";

const myProvider = new airflow.Provider("my-provider", {
    baseEndpoint: "https://my-airflow-instance/api/v1",
    oauth2Token: "my-token",
});

// Use `myProvider` to create resources in your Pulumi program.
```


## Example Usage for set resoruce connection with Node.js (JavaScript/TypeScript)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as airflow from "@pulumi/airflow";

const example = new airflow.Connection("example", {
    connType: "example",
    connectionId: "example",
});
```

## Import

Connections can be imported using the connection key. terraform

```sh
 $ pulumi import airflow:index/connection:Connection default example
```