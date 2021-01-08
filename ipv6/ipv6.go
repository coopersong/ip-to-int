package ipv6

import (
    "math/big"
    "strings"
)

func IPv6ToInt(ip string) (*big.Int, error) {
    ip = strings.ToLower(ip)
    res := big.NewInt(0)
    sections := strings.Split(ip, ":")
    // 判断是否存在连续多个冒号缩写为两个冒号的缩写形式
    if len(sections) < 8 {
        // 省略冒号的个数
        omitCount := 8 - len(sections)
        startIndex := strings.Index(ip, "::")
        prefix := ip[:startIndex]
        postfix := ip[startIndex:]
        for cnt := 0; cnt < omitCount; cnt++ {
            prefix += ":"
        }
        ip = prefix + postfix
    }
    sections = strings.Split(ip, ":")
    outerBase := big.NewInt(1)
    for i := len(sections) - 1; i >= 0; i -- {
        if sections[i] == "" {
            sections[i] = "0"
        }
        innerBase := outerBase
        for j := len(sections[i]) - 1; j >= 0; j-- {
            now := big.NewInt(hexToInt(sections[i][j]))
            res = res.Add(res, now.Mul(now, innerBase))
            innerBase = innerBase.Mul(innerBase, big.NewInt(16))
        }
        outerBase = outerBase.Mul(outerBase, big.NewInt(16 * 16 * 16 * 16))
    }
    return res, nil
}