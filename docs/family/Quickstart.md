---
title: Quickstart
weight: 1
---

# Quickstart

This guide walks through the process to install Upbound Universal Crossplane and install the GCP Beta provider-family.

To use GCP Beta provider-family, install Upbound Universal Crossplane into your Kubernetes cluster, install the `Provider`, apply a `ProviderConfig`, and create a *managed resource* in GCP via Kubernetes.

## Install the Up command-line
Download and install the Upbound `up` command-line.

```shell
curl -sL "https://cli.upbound.io" | sh
mv up /usr/local/bin/
```

Verify the version of `up` with `up --version`

```shell
$ up --version
v0.35.0
```
_Note_: official providers only support `up` command-line versions v0.13.0 or later.

## Install Universal Crossplane
Install Upbound Universal Crossplane with the Up command-line.

```shell
$ up uxp install
UXP 1.18.0-up.1 installed
```

_Note_: Official provider-families only support crossplane version 1.12.1 or UXP version 1.12.1-up.1 or later.

Verify the UXP pods are running with `kubectl get pods -n upbound-system`

```shell
$ kubectl get pods -n upbound-system
NAME                                       READY   STATUS    RESTARTS   AGE
crossplane-649f76c8db-vm5r5                1/1     Running   0          80s
crossplane-rbac-manager-645fdf89d6-rvj6q   1/1     Running   0          80s
```

## Install the GCP Beta provider-family

Install the official provider-family into the Kubernetes cluster with a Kubernetes configuration file.
For instance, let's install the `provider-gcp-beta-cloudplatform`

_Note_: The first provider installed of a family also installs an extra provider-family Provider.
The provider-family provider manages the ProviderConfig for all other providers in the same family.

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-gcp-beta-cloudplatform
spec:
  package: xpkg.upbound.io/upbound/provider-gcp-beta-cloudplatform:<version>
```

Apply this configuration with `kubectl apply -f`.

After installing the provider, verify the install with `kubectl get providers`.   

```shell
$ kubectl get providers
NAME                               INSTALLED   HEALTHY   PACKAGE                                                          AGE
provider-gcp-beta-cloudplatform    True        True      xpkg.upbound.io/upbound/provider-gcp-beta-cloudplatform:v0.3.0   37s
upbound-provider-family-gcp-beta   True        True      xpkg.upbound.io/upbound/provider-family-gcp-beta:v0.3.0          31s
```

It may take up to 5 minutes to report `HEALTHY`.

If you are going to use your own registry please check [Install Providers in an offline environment](https://docs.upbound.io/providers/provider-families/#installing-a-provider-family:~:text=services%20to%20install.-,Install%20Providers%20in%20an%20offline%20environment,-View%20the%20installed).

## Create a Kubernetes secret
The `provider-gcp-beta-cloudplatform` requires credentials to create and manage GCP resources.

### Generate a GCP JSON key file
Create a JSON key file containing the GCP account credentials. GCP provides documentation on [how to create a key file](https://cloud.google.com/iam/docs/creating-managing-service-account-keys).

Here is an example key file:  
```json
{
  "type": "service_account",
  "project_id": "caramel-goat-354919",
  "private_key_id": "e97e40a4a27661f12345678f4bd92139324dbf46",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCwA+6MWRhmcPB3\nF/irb5MDPYAT6BWr7Vu/16U8FbCHk7xtsAWYjKXKHu5mGzum4F781sM0aMCeitlv\n+jr2y7Ny23S9uP5W2kfnD/lfj0EjCdfoaN3m7W0j4DrriJviV6ESeSdb0Ehg+iEW\ngNrkb/ljigYgsSLMuemby5lvJVINUazXJtGUEZew+iAOnI4/j/IrDXPCYVNo5z+b\neiMsDYWfccenWGOQf1hkbVWyKqzsInxu8NQef3tNhoUXNOn+/kgarOA5VTYvFUPr\n2l1P9TxzrcYuL8XK++HjVj5mcNaWXNN+jnFpxjMIJOiDJOZoAo0X7tuCJFXtAZbH\n9P61GjhbAgMBAAECggEARXo31kRw4jbgZFIdASa4hAXpoXHx4/x8Q9yOR4pUNR/2\nt+FMRCv4YTEWb01+nV9hfzISuYRDzBEIxS+jyLkda0/+48i69HOTAD0I9VRppLgE\ne97e40a4a27661f12345678f4bd92139324dbf46+2H7ulQDtbEgfcWpNMQcL2JiFq+WS\neh3H0gHSFFIWGnAM/xofrlhGsN64palZmbt2YiKXcHPT+WgLbD45mT5j9oMYxBJf\nPkUUX5QibSSBQyvNqCgRKHSnsY9yAkoNTbPnEV0clQ4FmSccogyS9uPEocQDefuY\nY7gpwSzjXpaw7tP5scK3NtWmmssi+dwDadfLrKF7oQKBgQDjIZ+jwAggCp7AYB/S\n6dznl5/G28Mw6CIM6kPgFnJ8P/C/Yi2y/OPKFKhMs2ecQI8lJfcvvpU/z+kZizcG\nr/7iRMR/SX8n1eqS8XfWKeBzIdwQmiKyRg2AKelGKljuVtI8sXKv9t6cm8RkWKuZ\n9uVroTCPWGpIrh2EMxLeOrlm0QKBgQDGYxoBvl5GfrOzjhYOa5GBgGYYPdE7kNny\nhpHE9CrPZFIcb5nGMlBCOfV+bqA9ALCXKFCr0eHhTjk9HjHfloxuxDmz34vC0xXG\ncegqfV9GNKZPDctysAlCWW/dMYw4+tzAgoG9Qm13Iyfi2Ikll7vfeMX7fH1cnJs0\nnYpN9LYPawKBgQCwMi09QoMLGDH+2pLVc0ZDAoSYJ3NMRUfk7Paqp784VAHW9bqt\n1zB+W3gTyDjgJdTl5IXVK+tsDUWu4yhUr8LylJY6iDF0HaZTR67HHMVZizLETk4M\nLfvbKKgmHkPO4NtG6gEmMESRCOVZUtAMKFPhIrIhAV2x9CBBpb1FWBjrgQKBgQCj\nkP3WRjDQipJ7DkEdLo9PaJ/EiOND60/m6BCzhGTvjVUt4M22XbFSiRrhXTB8W189\noZ2xrGBCNQ54V7bjE+tBQEQbC8rdnNAtR6kVrzyoU6xzLXp6Wq2nqLnUc4+bQypT\nBscVVfmO6stt+v5Iomvh+l+x05hAjVZh8Sog0AxzdQKBgQCMgMTXt0ZBs0ScrG9v\np5CGa18KC+S3oUOjK/qyACmCqhtd+hKHIxHx3/FQPBWb4rDJRsZHH7C6URR1pHzJ\nmhCWgKGsvYrXkNxtiyPXwnU7PNP9JNuCWa45dr/vE/uxcbccK4JnWJ8+Kk/9LEX0\nmjtDm7wtLVlTswYhP6AP69RoMQ==\n-----END PRIVATE KEY-----\n",
  "client_email": "my-sa-313@caramel-goat-354919.iam.gserviceaccount.com",
  "client_id": "103735491955093092925",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/my-sa-313%40caramel-goat-354919.iam.gserviceaccount.com"
}
```

Save this JSON file as `gcp-credentials.json`.

### Create a Kubernetes secret with GCP credentials
Use `kubectl create secret -n upbound-system` to generate the Kubernetes secret object inside the Universal Crossplane cluster.

`kubectl create secret generic gcp-secret -n upbound-system --from-file=creds=./gcp-credentials.json`

View the secret with `kubectl describe secret`
```shell
$ kubectl describe secret gcp-secret -n upbound-system
Name:         gcp-secret
Namespace:    upbound-system
Labels:       <none>
Annotations:  <none>

Type:  Opaque

Data
====
creds:  2334 bytes
```
## Create a ProviderConfig
Create a `ProviderConfig` Kubernetes configuration file to attach the GCP credentials to the installed official `provider-gcp-beta-cloudplatform`.

**Note:** the `ProviderConfig` must contain the correct GCP project ID. The project ID must match the `project_id` from the JSON key file.

```yaml
apiVersion: gcp-beta.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  projectID: <PROJECT_ID>
  credentials:
    source: Secret
    secretRef:
      namespace: upbound-system
      name: gcp-secret
      key: creds
```

Apply this configuration with `kubectl apply -f`.

**Note:** the `ProviderConfig` value `spec.secretRef.name` must match the `name` of the secret in `kubectl get secrets -n upbound-system` and `spec.secretRef.key` must match the value in the `Data` section of the secret.

Verify the `ProviderConfig` with `kubectl describe providerconfigs`. 

```yaml
$ kubectl describe providerconfigs
Name:         default
Namespace:
API Version:  gcp-beta.upbound.io/v1beta1
Kind:         ProviderConfig
# Output truncated
Spec:
  Credentials:
    Secret Ref:
      Key:        creds
      Name:       gcp-secret
      Namespace:  upbound-system
    Source:       Secret
  Project ID:     caramel-goat-354919
```

## Create a managed resource
Create a managed resource to verify the `provider-gcp-beta-cloudplatform` is functioning.

This example creates a GCP cloudplatform ServiceAccount, which requires a globally unique name.

Create a `ServiceAccount` configuration file.

```yaml
apiVersion: cloudplatform.gcp-beta.upbound.io/v1beta1
kind: ServiceAccount
metadata:
  annotations:
    meta.upbound.io/example-id: cloudplatform/v1beta1/serviceaccount
  labels:
    testing.upbound.io/example-name: service_account
  name: service-account
spec:
  forProvider:
    displayName: Upbound GCP Beta Example Service Account
  providerConfigRef:
    name: defaultw
```

**Note:** the `spec.providerConfigRef.name` must match the `ProviderConfig` `metadata.name` value.

Apply this configuration with `kubectl apply -f`.

Use `kubectl get managed` to verify service account creation.

```shell
NAME                                                               SYNCED   READY   EXTERNAL-NAME     AGE
serviceaccount.cloudplatform.gcp-beta.upbound.io/service-account   True     True    service-account   37s
```

Provider created the service account when the values `READY` and `SYNCED` are `True`.

If the `READY` or `SYNCED` are blank or `False` use `kubectl describe` to understand why.

Here is an example of a failure because the `ServiceAccount` `spec.providerConfigRef.name` value doesn't match the `ProviderConfig` `metadata.name`.

```shell
> kubectl describe serviceaccount.cloudplatform.gcp-beta.upbound.io/service-account
Name:         service-account
Namespace:
Labels:       testing.upbound.io/example-name=service_account
Annotations:  crossplane.io/external-name: service-account
              meta.upbound.io/example-id: cloudplatform/v1beta1/serviceaccount
API Version:  cloudplatform.gcp-beta.upbound.io/v1beta1
Kind:         ServiceAccount
Metadata:
  Creation Timestamp:  2025-01-06T17:57:29Z
  Generation:          2
  Resource Version:    21974
  UID:                 f7280fd6-8073-4c6a-8d78-5e9552ae6913
Spec:
  Deletion Policy:  Delete
  For Provider:
    Display Name:  Upbound GCP Beta Example Service Account
  Init Provider:
  Management Policies:
    *
  Provider Config Ref:
    Name:  defaultw
Status:
  At Provider:
  Conditions:
    Last Transition Time:  2025-01-06T17:57:29Z
    Message:               connect failed: cannot initialize the Terraform plugin SDK async external client: cannot get terraform setup: cannot get referenced ProviderConfig: ProviderConfig.gcp-beta.upbound.io "default" not found
    Reason:                ReconcileError
    Status:                False
    Type:                  Synced
Events:
  Type     Reason                   Age               From                                                                    Message
  ----     ------                   ----              ----                                                                    -------
  Warning  CannotConnectToProvider  6s (x4 over 11s)  managed/cloudplatform.gcp-beta.upbound.io/v1beta1, kind=serviceaccount  cannot initialize the Terraform plugin SDK async external client: cannot get terraform setup: cannot get referenced ProviderConfig: ProviderConfig.gcp-beta.upbound.io "default" not found
```

The output indicates the `ServiceAccount` is using `ProviderConfig` named `default`. The applied `ProviderConfig` is `my-config`.

```shell
$ kubectl get providerconfigs
NAME        AGE
my-config   56s
```

## Delete the managed resource
Remove the managed resource by using `kubectl delete -f` with the same `ServiceAccount` object file. Verify the removal of the bucket with `kubectl get managed`.

```shell
$ kubectl get managed
No resources found
```
