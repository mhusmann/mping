// Small tool to scan living IPs in my network
// using external tools: fping and host, 2014-01-26 mh
package main

import (
  "fmt"
  "os/exec"
  "log"
  "bytes"
  "strings"
  "regexp"
)

func main() {
  startIp := "192.168.178.1"
  endIp := "192.168.178.100"
  // regex to eliminate the output IP address
  repl := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)
  re := regexp.MustCompile(`\s`)

  fmt.Printf("Prüfe alle Ips im Bereich von %s bis %s auf Erreichbarkeit\n\n",
              startIp, endIp)

  out0, err0 := exec.Command("fping", "-gae", startIp,
                             endIp + " 2>/dev/null").Output()
  
  ips := bytes.Split(out0, []byte{'\n'})
  fmt.Printf("%d Geräte gefunden\n", len(ips)-1)
  for _, ip := range(ips) {
    ipAndTime := re.Split(string(ip), 2)
    out1, _ := exec.Command("host", ipAndTime[0]).Output()
    // remove ip address from string
    out2 := repl.ReplaceAllString(string(out1), "")
    fmt.Printf("%-15s --> %s", ip, strings.Replace(string(out2),
               ".in-addr.arpa domain name pointer",
               "", 1))
  }

  if err0 != nil {
    log.Fatal(" --> nicht alle hosts erreichbar")
  }
}
