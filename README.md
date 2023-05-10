## kubectl plugin: kubectl-cluster

A simple kubectl plugin: supports specifying cluster names when using kubectl

If you have a multi-cluster config file in the ~/.kube directory, you can use it

### Installation

```bash
🤖 ~ git clone https://github.com/Hargeek/kubectl-cluster.git
🤖 ~ cd kubectl-cluster
🤖 ~ make build && make install
```

### Usage

```bash
🤖 ~ kubectl cluster -m <cluster-name> <kubectl command>
```

#### e.g.

```bash
🤖 ~ kubectl cluster -m dev -n default get pods
NAME                     READY   STATUS    RESTARTS      AGE
nginx-5dd948d984-lw68t   1/1     Running   1 (46h ago)   13d
```