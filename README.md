mackerel-plugin-proc-net-arp
============================

mackerel-agent plugin for /proc/net/arp

## Synopsis

```shell
mackerel-plugin-proc-net-arp [-target=<path to /proc/net/arp>] [-tempfile=<tempfile>]
```

## Example of mackerel-agent.conf

```ascii
[plugin.metrics.ipvs]
command = "mackerel-plugin-proc-net-arp"
```
