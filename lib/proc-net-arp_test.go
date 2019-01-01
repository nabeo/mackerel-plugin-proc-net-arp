package mpprocnetarp

import (
  "strings"
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestGraphDefinition(t *testing.T) {
  var r ArpPlugin
  graphdef := r.GraphDefinition()
  assert.Len(t, graphdef, 1)
}

func TestParse(t *testing.T) {
  var r ArpPlugin
  s1 := `IP address       HW type     Flags       HW address            Mask     Device
192.168.0.1      0x1         0x2         00:00:5E:00:53:00     *        eth0
192.168.0.2      0x1         0x2         00:00:5E:00:53:01     *        eth0
192.168.0.3      0x1         0x2         00:00:5E:00:53:02     *        eth0
192.168.0.4      0x1         0x2         00:00:5E:00:53:03     *        eth0
192.168.0.5      0x1         0x2         00:00:5E:00:53:04     *        eth0
192.168.0.6      0x1         0x2         00:00:5E:00:53:05     *        eth0
`

  stubData1 := strings.NewReader(s1)
  a, err := r.Parse(stubData1)
  assert.Nil(t, err)
  assert.EqualValues(t, 6, a["size"])
}
