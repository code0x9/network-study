## about

network study

## lecture

Networking labs on ARM: https://youtu.be/_BTa-CiTpvI?si=DtdWVWI-UTPXixph
Running Containlerab on macOS and Windows with Devcontainers: https://youtu.be/Xue1pLiO0qQ?si=zHrmeuNOIYsHGgEk

## ARM64-native Network OS

- nokia SR Linux: https://www.nokia.com/ip-networks/service-router-linux-NOS/
- arista cEOS
- cisco
- juniper networks: cRPD

### others

- FRR: https://quay.io/repository/frrouting/frr?tab=info
- VyOS
- sonic-vs
- cumulus CVX

## kubernetes

- open virtual network(OVN) -> OVN-K
  - pure, simple
- open vSwitch -> KUBE-OVN
  - complex, VPC per namespace

## computing

kubevirt: https://github.com/kubevirt/kubevirt

- use case: https://youtu.be/K88rBh9o0SY?si=nc_u6enSF8VYu8-q
  -> it needs routable IP, injected via multus CNI for seamless VM experience

## dev setup

- vscode extension
  - Dev Containers: https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers

## gNMI: <https://gnmic.openconfig.net>

```shell
brew install gnmic
```

```shell
gnmic -a 172.20.20.2:57400 --skip-verify -u admin -p 'NokiaSrl1!' capabilities 
```

```shell
gnmic -a 172.20.20.2:57400 --skip-verify -u admin -p 'NokiaSrl1!' --encoding JSON_IETF get --path '/interface[name=ethernet-1/1]/subinterface[index=0]/ipv4/address'
```

```shell
brew install wireshark wireshark-app
ln -s /Applications/Wireshark.app/Contents/MacOS/Wireshark /usr/local/bin/wireshark
```

```shell
## check generated root password from `docker logs`
ssh root@localhost -p 8022 'tcpdump -i eth0 -w - "not port 22"' | /Applications/Wireshark.app/Contents/MacOS/Wireshark -k -i -
```

## SONiC
https://containerlab.dev/manual/kinds/sonic-vs/

1. get sonic image: https://containerlab.dev/manual/kinds/sonic-vs/#getting-sonic-images
2. load image: `docker load docker-sonic-vs.gz`

## ref

- https://containerlab.dev/
  - Containerlab on macOS: https://containerlab.dev/macos/
  - Lab catalog: https://clabs.netdevops.me/
- https://www.openconfig.net/
- Network OS
  - Nokia SR Linux: https://learn.srlinux.dev/
