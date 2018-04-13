
<!--- This file is automatically generated by make gen-cli-docs; changes should be made in the go CLI command code (under cmd/kops) -->

## kops create cluster

Create a Kubernetes cluster.

### Synopsis


Create a kubernetes cluster using command line flags. This command creates cloud based resources such as networks and virtual machines. Once the infrastructure is in place Kubernetes is installed on the virtual machines. 

These operations are done in parallel and rely on eventual consistency.

```
kops create cluster
```

### Examples

```
  # Create a cluster in AWS
  kops create cluster --name=kubernetes-cluster.example.com \
  --state=s3://kops-state-1234 --zones=eu-west-1a \
  --node-count=2
  
  # Create a cluster in AWS that has HA masters.  This cluster
  # will be setup with an internal networking in a private VPC.
  # A bastion instance will be setup to provide instance access.
  
  export NODE_SIZE=${NODE_SIZE:-m4.large}
  export MASTER_SIZE=${MASTER_SIZE:-m4.large}
  export ZONES=${ZONES:-"us-east-1d,us-east-1b,us-east-1c"}
  export KOPS_STATE_STORE="s3://my-state-store"
  kops create cluster k8s-clusters.example.com \
  --node-count 3 \
  --zones $ZONES \
  --node-size $NODE_SIZE \
  --master-size $MASTER_SIZE \
  --master-zones $ZONES \
  --networking weave \
  --topology private \
  --bastion="true" \
  --yes
  
  # Create cluster in GCE.
  # This is an alpha feature.
  export KOPS_STATE_STORE="gs://mybucket-kops"
  export ZONES=${MASTER_ZONES:-"us-east1-b,us-east1-c,us-east1-d"}
  export KOPS_FEATURE_FLAGS=AlphaAllowGCE
  
  kops create cluster kubernetes-k8s-gce.example.com
  --zones $ZONES \
  --master-zones $ZONES \
  --node-count 3
  --project my-gce-project \
  --image "ubuntu-os-cloud/ubuntu-1604-xenial-v20170202" \
  --yes
  # Create manifest for a cluster in AWS
  kops create cluster --name=kubernetes-cluster.example.com \
  --state=s3://kops-state-1234 --zones=eu-west-1a \
  --node-count=2 --dry-run -oyaml
```

### Options

```
      --admin-access stringSlice             Restrict API access to this CIDR.  If not set, access will not be restricted by IP. (default [0.0.0.0/0])
      --api-loadbalancer-type string         Sets the API loadbalancer type to either 'public' or 'internal'
      --associate-public-ip                  Specify --associate-public-ip=[true|false] to enable/disable association of public IP for master ASG and nodes. Default is 'true'.
      --authorization string                 Authorization mode to use: AlwaysAllow or RBAC (default "RBAC")
      --bastion                              Pass the --bastion flag to enable a bastion instance group. Only applies to private topology.
      --channel string                       Channel for default versions and configuration to use (default "stable")
      --cloud string                         Cloud provider to use - gce, aws, vsphere
      --cloud-labels string                  A list of KV pairs used to tag all instance groups in AWS (eg "Owner=John Doe,Team=Some Team").
      --dns string                           DNS hosted zone to use: public|private. Default is 'public'. (default "Public")
      --dns-zone string                      DNS hosted zone to use (defaults to longest matching zone)
      --dry-run                              If true, only print the object that would be sent, without sending it. This flag can be used to create a cluster YAML or JSON manifest.
      --encrypt-etcd-storage                 Generate key in aws kms and use it for encrypt etcd volumes
      --image string                         Image to use for all instances.
      --kubernetes-version string            Version of kubernetes to run (defaults to version in channel)
      --master-count int32                   Set the number of masters.  Defaults to one master per master-zone
      --master-public-name string            Sets the public master public name
      --master-security-groups stringSlice   Add precreated additional security groups to masters.
      --master-size string                   Set instance size for masters
      --master-tenancy string                The tenancy of the master group on AWS. Can either be default or dedicated.
      --master-volume-size int32             Set instance volume size (in GB) for masters
      --master-zones stringSlice             Zones in which to run masters (must be an odd number)
      --model string                         Models to apply (separate multiple models with commas) (default "config,proto,cloudup")
      --network-cidr string                  Set to override the default network CIDR
      --networking string                    Networking mode to use.  kubenet (default), classic, external, kopeio-vxlan (or kopeio), weave, flannel-vxlan (or flannel), flannel-udp, calico, canal, kube-router, romana, amazon-vpc-routed-eni, cilium. (default "kubenet")
      --node-count int32                     Set the number of nodes
      --node-security-groups stringSlice     Add precreated additional security groups to nodes.
      --node-size string                     Set instance size for nodes
      --node-tenancy string                  The tenancy of the node group on AWS. Can be either default or dedicated.
      --node-volume-size int32               Set instance volume size (in GB) for nodes
      --out string                           Path to write any local output
  -o, --output string                        Output format. One of json|yaml. Used with the --dry-run flag.
      --project string                       Project to use (must be set on GCE)
      --ssh-access stringSlice               Restrict SSH access to this CIDR.  If not set, access will not be restricted by IP. (default [0.0.0.0/0])
      --ssh-public-key string                SSH public key to use (default "~/.ssh/id_rsa.pub")
      --subnets stringSlice                  Set to use shared subnets
      --target string                        Valid targets: direct, terraform, cloudformation. Set this flag to terraform if you want kops to generate terraform (default "direct")
  -t, --topology string                      Controls network topology for the cluster. public|private. Default is 'public'. (default "public")
      --utility-subnets stringSlice          Set to use shared utility subnets
      --vpc string                           Set to use a shared VPC
  -y, --yes                                  Specify --yes to immediately create the cluster
      --zones stringSlice                    Zones in which to run the cluster
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --config string                    config file (default is $HOME/.kops.yaml)
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files (default false)
      --name string                      Name of cluster. Overrides KOPS_CLUSTER_NAME environment variable
      --state string                     Location of state storage. Overrides KOPS_STATE_STORE environment variable
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO
* [kops create](kops_create.md)	 - Create a resource by command line, filename or stdin.
