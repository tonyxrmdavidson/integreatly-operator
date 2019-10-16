# Integreatly Operator

A Kubernetes Operator based on the Operator SDK for installing and reconciling Integreatly products.

### Project status: _alpha_ 

This is a proof of concept/alpha version. Most functionality is present but it is highly likely there are bugs and improvements needed.

### Installed products
Currently the operator installs the following products:
- AMQ Online
- AMQ Streams
- Codeready
- Fuse
- Launcher
- Nexus
- RHSSO (both a cluster instance and a user instance)
- 3scale
- Integreatly solution explorer

## Prerequisites
- [operator-sdk](https://github.com/operator-framework/operator-sdk) version v0.10.0.
- [go](https://golang.org/dl/) version 1.12+
- [moq](https://github.com/matryer/moq)
- [oc](https://docs.okd.io/3.11/cli_reference/get_started_cli.html#cli-reference-get-started-cli) version v3.11+
- Access to an Openshift v4.1.0+ cluster
- A user with administrative privileges in the OpenShift cluster
- AWS account with permissions to create S3 buckets

After installation, the following commands must be run to avoid a known [issue](https://github.com/matryer/moq/issues/98) related to the Moq package:
```
go get -u .
go install github.com/matryer/moq
```

## Local Setup

Download the integreatly-operator project:
```sh
mkdir -p $GOPATH/src/github.com/integr8ly
cd $GOPATH/src/github.com/integr8ly
git clone https://github.com/integr8ly/integreatly-operator
cd integreatly-operator
```

Some products require certain credentials to be present in the namespace before installation can proceed: 
* 3scale requires AWS credentials for backups to an S3 bucket. The bucket should have all public access turned off
* RHSSO requires Github OAuth credentials to create a Github Identity Provider for Launcher (see [here](https://github.com/integr8ly/installation/#51-create-github-oauth-to-enable-github-authorization-for-launcher) for creating a Github OAuth app) and Codeready

**Note:** If these secrets aren't created, the integreatly preflight checks will fail

```sh
# The project name for the integreatly operator to watch 
export NAMESPACE="integreatly-test"

# 3scale requires AWS credentials for backups to S3
export AWS_ACCESS_KEY_ID=<access key>
export AWS_SECRET_ACCESS_KEY=<access secret>
export AWS_BUCKET=<bucket name>

# RHSSO requires Github OAuth credentials to setup a Github identity provider
# for Fabric8 Launcher and Codeready
export GH_CLIENT_ID=<client id>
export GH_CLIENT_SECRET=<client secret>

# Bootstrap the project
make cluster/prepare/local
```

An `Installation` custom resource can now be created which will kick of the installation of the integreatly products, once the operator is running:
```sh
# Create the installation custom resource
oc create -f deploy/crds/examples/installation.cr.yaml

# The operator can now be run locally
make code/run
```

### Logging in to SSO 

In the OpenShift UI, in `Projects > integreatly-rhsso > Networking > Routes`, select the `sso` route to open up the SSO login page.

The username is `admin`, and the password can be retrieved by running:
```sh
oc get dc sso -n integreatly-rhsso -o jsonpath='{.spec.template.spec.containers[0].env[?(@.name=="SSO_ADMIN_PASSWORD")].value}'
```

### Configuring Github OAuth
Log in to RHSSO (see above) and click `Identity Providers` in the left sidebar. In the Github identity provider, find the Redirect URI and paste this URL into the Homepage URL and Authorization callback URL fields of your Github OAuth app. 

## Deploying to a Cluster with OLM
Make sure to export the variables above (see [local setup](#local-setup)), then run:

```sh
make cluster/prepare
```

Within a few minutes, the Integreatly operator should be visible in the OperatorHub (`Catalog > OperatorHub`). To create a new subscription, click on the Install button, choose to install the operator in the created namespace and keep the approval strategy on automatic.

Once the subscription shows a status of `installed`, a new Integreatly `Installation` custom resource can be created which will begin to install the supported products.

In `Catalog > Developer Catalog`, choose the Integreatly Installation and click create. An example installation CR can be found below:

```yml
apiVersion: integreatly.org/v1alpha1
kind: Installation
metadata:
  name: example-installation
spec:
  type: workshop
  namespacePrefix: integreatly-
  routingSubdomain: example.com
  masterUrl: http://master.example.com
  selfSignedCerts: true
```
## Manual Deployment to a Cluster
### Create Namespace in Openshift Console
From Home/Projects
 
```sh
Create Project
```

Input name 			e.g.   integreatly

Input display name 	e.g.	integreatly

```sh
Create
```
### Install Operator in Openshift Console
Goto Catalog/‘Installed Operators’   

In top right corner Add/‘Import YAML’

Copy the source code from:
```sh
https://raw.githubusercontent.com/integr8ly/manifests/master/operator-source.yml
```
Should look like :
```sh
apiVersion: operators.coreos.com/v1
kind: OperatorSource
metadata:
  	name: integreatly-operators
  	namespace: openshift-marketplace
spec:
  	authorizationToken: {}
  	displayName: Integreatly Operators
  	endpoint: 'https://quay.io/cnr'
  	publisher: Integreatly Publisher
  	registryNamespace: integreatly
  	type: appregistry
```
Paste it into your console/ImportYAML
```sh
Create
```
Goto Catalog/OperatorHub

Select the integreatly Operator   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; *note :  you may need to refresh a couple of times before it becomes available.*

Select Install

Select ‘A specific namespace on the cluster’

From dropdown menu select integreatly (the example name from create project above)

```sh
Subscribe
```
### Create Custom Resource
From your terminal :
```sh
oc login 
cd ~/go/src/github.com/integr8ly/integreatly-operator/deploy/crds/examples
nano installation.cr.yaml
```
Edit the installation.cr.yaml with relevant details as required.
Example :
```sh
apiVersion: integreatly.org/v1alpha1
kind: Installation
metadata:
  name: example-installation
spec:
  type: managed
  namespacePrefix: integreatly-
  routingSubdomain: apps.<example>.cloudservices.rhmw.io
  masterUrl: https://console-openshift-console.apps.<example>.cloudservices.rhmw.io
  selfSignedCerts: true
```
Take note of metadata.name

Save file and exit

### Create the installation.cr.yaml in your integreatly namespace
```sh
oc create -f installation.cr.yaml  -n integreatly
```
Response should be

$ installation.integreatly.org/example-installation created
### Create AWS secrets
The appropriate keys can be obtained from ~/.aws/credentials
```sh
cd  ~/go/src/github.com/integr8ly/integreatly-operator
oc process -f deploy/s3-secrets.yaml -p INSTALLATION_NAMESPACE=integreatly -p AWS_ACCESS_KEY_ID=<your aws key> -p AWS_SECRET_ACCESS_KEY=<your aws secret key> -p AWS_BUCKET=<your aws bucket> -p AWS_REGION=eu-west-1 | oc apply -f -
```
### Create Github secrets
```sh
oc create secret generic github-oauth-secret --from-literal=clientId=<your github client oauth app client id> --from-literal=secret=<your github oauth app client secret> -n integreatly
```
### Monitor installaion
Select the integreatly namespace

Home/Search

From drop down menu select Installation

Select example-installation

Select YAML

You will see the YAML will update as the installation progresses.


## Set up dedicated admins 

To setup your cluster to have dedicated admins run the `setup/dedicated` target which installs the dedicated admin operator:
```sh
make setup/dedicated
```

If you want to remove the dedicated admin operator, run:
```sh
make clean/dedicated
```

## Tests

Running unit tests:
```sh
make test/unit
```

## Release

Update the operator version in the following files:

* Update [version/version.go](version/version.go) (`Version = "<version>"`)

* Update `TAG` in the [Makefile](Makefile) 

* Update the operator image version in [deploy/operator.yaml](deploy/operator.yaml)
(`image: quay.io/integreatly/integreatly-operator:v<version>`)


Commit changes and open pull request. When the PR is accepted, create a new release tag:

```sh
git tag v<version> && git push upstream v<version>
```