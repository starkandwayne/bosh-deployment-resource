- client: Username or UAA client
- client-secret: Password or UAA client secret

# BOSH Deployment Resource

A resource that will deploy releases and stemcells.

## Source Configuration

* `deployment`: *Required.* The name of the deployment.
* `target`: *Required.* The address of the BOSH director which will be used for
  the deployment.
* `client`: *Required.* The UAA client ID for the BOSH director.
* `client_secret`: *Required.* The UAA client secret for the BOSH director.
* `ca_cert`: *Required.* CA certificate used to validate SSL connections to Director and UAA.

### Example

``` yaml
- name: staging
  type: bosh-deployment
  source:
    deployment: staging-deployment-name
    target: https://bosh.example.com:25555
    client: admin
    client_secret: admin
    ca_cert: "-----BEGIN CERTIFICATE-----\n-----END CERTIFICATE-----"
```

## Behaviour

### `in`: Deploy a BOSH deployment

This will download the deployment manifest. It will place two files in the target directory:

- `manifest.yml`: The deployment manifest
- `target`: The hostname for the director

_Note_: Only the most recent version is fetchable

### `out`: Deploy a BOSH deployment

This will deploy the deployment provided.

#### Parameters

* `manifest`: *Required.* Path to a BOSH deployment manifest file.

``` yaml
- put: staging
  params:
    manifest: path/to/manifest.yml
```