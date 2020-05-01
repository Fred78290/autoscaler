# Grpc bridge cloud provider

This cloud provider allow to externalize the real cloud provider, it provide a grpc bridge. All apis calls are forwarded thru grpc to the real cloud provider.

This method permit to develop your own cloud provider without modify cluster-autoscaler code.

## When to use it

By default cluster-autoscaler embed cloud provider and this one talk to the hypervisor (Azure, aws....) but if you want to use an another provider, you need to fork the original project and add your code.

Now using the grpc cloud provider, you write your cloud provider outside the cluster-auto scaler code.

## Examples of externalized cloud provider

Some externalized cloud provider could be found here:

vSphere: <https://github.com/Fred78290/kubernetes-vmware-autoscaler>

Multipass: <https://github.com/Fred78290/kubernetes-multipass-autoscaler>

## Run it

Before start cluster-autscaler, be sure that your externalized grpc cloud provider is running, else the cluster-autoscaler wont start.

cluster-autoscaler have same command line arguments but you must supply mandatory arguments

| Parameter | Description |
| --- | --- |
| `cloud-provider` | The type of provider. Must be grpc  |
| `cloud-config`  | The path to the grpc config  |
| `nodes`  | The grpc node definition |
| `node-autoprovisioning-enabled` | Tell to your provider to create a node group |

The syntax of **nodes** parameter is

min nodes:max nodes:include existing nodes/node group name|node label;node label

Node label could be omited.

And fill the node groups in container command by `--nodes`, e.g.

```yaml
- --nodes=0:3:true/ca-grpc-ng-01|monitor=true;database=true
```

or multiple node groups:

```yaml
- --nodes=0:3:true/ca-grpc-ng-01|monitor=true;database=true
- --nodes=0:3:true/ca-grpc-ng-02|monitor=true;database=true
```

The cloud config is like this

```json
{
    "address": "10.0.0.158:5200",
    "secret": "vmware",
    "timeout": 300,
    "config": {
        "kubeAdmAddress": "10.0.0.200:6443",
        "kubeAdmToken": "5cn3x8.2bgs3x8tivo07aj1",
        "kubeAdmCACert": "sha256:a14f7dd6b8f46ebb1810a9ef734f7048e2146bb897469ff0ff08d09008883703",
        "kubeAdmExtraArguments": [
            "--ignore-preflight-errors=All"
        ]
    }
}
```

If you provide the argument `--node-autoprovisioning-enabled`, it will tell to your provider to create the node group.

### Example to launch cluster-autoscaler

```bash
cluster-autoscaler \
    --v=4 \
    --stderrthreshold=info \
    --cloud-provider=grpc \
    --cloud-config=$HOME/go/src/github.com/Fred78290/kubernetes-vmware-autoscaler/masterkube/config/grpc-config.json \
    --kubeconfig=$HOME/go/src/github.com/Fred78290/kubernetes-vmware-autoscaler/masterkube/cluster/config \
    --nodes="0:3:true/ca-grpc-ng-01|monitor=true;database=true" \
    --max-nodes-total=3 \
    --cores-total=0:16 \
    --memory-total=0:24 \
    --node-autoprovisioning-enabled \
    --max-autoprovisioned-node-group-count=1 \
    --unremovable-node-recheck-timeout=2m \
    --scale-down-enabled=true \
    --scale-down-delay-after-add=1m \
    --scale-down-delay-after-delete=1m \
    --scale-down-delay-after-failure=1m \
    --scale-down-unneeded-time=1m \
    --scale-down-unready-time=1m
```
