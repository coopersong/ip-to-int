package ipv4

import (
    "math/big"
    "strconv"
    "strings"
)

func IPv4ToInt(ip string) (*big.Int, error) {
    res := big.NewInt(0)
    var sections []string
    // IPv4
    sections = strings.Split(ip, ".")
    base := big.NewInt(1)
    for i := len(sections) - 1; i >= 0; i-- {
        decimalInt, err := strconv.Atoi(sections[i])
        if err != nil {
            return big.NewInt(0), err
        }
        decimal := big.NewInt(int64(decimalInt))
        res = res.Add(res, decimal.Mul(decimal, base))
        base = base.Mul(base, big.NewInt(256))
    }
    return res, nil
}