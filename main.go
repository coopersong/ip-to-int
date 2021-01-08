package main

import (
    "flag"
    "fmt"
    "net"
    "strings"
)

import (
    "ip-to-int/ipv4"
    "ip-to-int/ipv6"
)

var help = flag.Bool("h", false, "help")
var ip = flag.String("a", "", "ip to convert")

func checkIPFormat(ip string) error {
    if net.ParseIP(ip) == nil {
        fmt.Errorf("invalid ip format")
    }
    return nil
}

func main() {
    flag.Parse()

    if *help {
        flag.PrintDefaults()
        return
    }

    if *ip == "" {
        flag.PrintDefaults()
        return
    }

    if err := checkIPFormat(*ip); err != nil {
        fmt.Println(err)
        return
    }

    if strings.Contains(*ip, ".") {
        // IPv4
        res, err := ipv4.IPv4ToInt(*ip)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println(res)
    } else {
        // IPv6
        res, err := ipv6.IPv6ToInt(*ip)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(res)
    }
}
