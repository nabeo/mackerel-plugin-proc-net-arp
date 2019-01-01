package mpprocnetarp

import (
  "os"
  "io"
  "bufio"
  "strings"

  mp "github.com/mackerelio/go-mackerel-plugin"
)

// ArpPlugin struct
type ArpPlugin struct {
  Prefix string
}

// define graph
var graphdef = map[string]mp.Graphs{
  "proc.net.arp": {
    Label: "ARP Table",
    Unit: mp.UnitInteger,
    Metrics: []mp.Metrics {
      {Name: "size", Label: "ARP Table size" },
    },
  },
}

// GraphDefinition : interface for go-mackerel-plugin interface
func (r ArpPlugin) GraphDefinition() map[string]mp.Graphs {
  return graphdef
}

// FetchMetrics : interface for go-mackerel-plugin interface
func (r ArpPlugin) FetchMetrics() (map[string]float64, error) {
  file, err := os.Open("/proc/net/arp")
  if err != nil {
    return nil, err
  }
  defer file.Close()

  return r.Parse(file)
}

// Parse : Parse /proc/net/arp
func (r ArpPlugin) Parse(stat io.Reader) (map[string]float64, error) {
  data := make(map[string]float64)
  scanner := bufio.NewScanner(stat)
  for scanner.Scan() {
    fields := strings.Fields(scanner.Text())
    // <IP Address> <HW Type> <Flags> <HW address> <Mask> <Device>
    if len(fields) != 6 {
      continue
    }
    // Skip header line
    if fields[0] == "IP" {
      continue
    }
    data["size"]++
  }
  return data, nil
}

// Do : Do plugin
func Do() {
  var Arp ArpPlugin

  helper := mp.NewMackerelPlugin(Arp)

  helper.Run()
}
