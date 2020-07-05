# nezha web-terminal
websocket to k8s container.

# usage
```
web-terminal.exe -h
Usage of web-terminal.exe:
  -kubecfg string
        kubeconfig绝对路径 (default "C:\\Users\\Sahara\\.kube\\config")
  -port string
        监听端口号 (default "8000")

```

### demo
```
web-terminal.exe -kubecfg config -port 8088
```
### config file demo
```
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUJWekNCL3FBREFnRUNBZ0VBTUFvR0NDcUdTTTQ5QkFNQ01DTXhJVEFmQmdOVkJBTU1HR3N6Y3kxelpYSjIKWlhJdFkyRkFNVFU0T1RFd05UZ3dPVEFlRncweU1EQTFNVEF4TURFMk5EbGFGdzB6TURBMU1EZ3hNREUyTkRsYQpNQ014SVRBZkJnTlZCQU1NR0dzemN5MXpaWEoyWlhJdFkyRkFNVFU0T1RFd05UZ3dPVEJaTUJNR0J5cUdTTTQ5CkFnRUdDQ3FHU000OUF3RUhBMElBQkN2TEt6Wmt0b21KdVhJMExUK0FCYlRpVnFLcGwvVmZaTURyYVRnODFxbjcKLzd1dmVtU2owMW1yY0w1RE11UnM5Mnludk51b0FNYS9FSDZLQWxBVVJDR2pJekFoTUE0R0ExVWREd0VCL3dRRQpBd0lDcERBUEJnTlZIUk1CQWY4RUJUQURBUUgvTUFvR0NDcUdTTTQ5QkFNQ0EwZ0FNRVVDSVFEUjZndVUwNXpKCndpMHhaZGF2RWRaNlRodFhNZnRFSmtJbmVFYk1tK25HU3dJZ0lVNUk5VUdtVVdzOE1kUEVBZ3c0a0RFUzFmSzAKWTdKMkJtY01RVE1LZHY4PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==
    server: https://192.168.1.111:6443
  name: default
contexts:
- context:
    cluster: default
    user: default
  name: default
current-context: default
kind: Config
preferences: {}
users:
- name: default
  user:
    password: 0fb4d10271daced60b408bac822db3a9
    username: admin
```
