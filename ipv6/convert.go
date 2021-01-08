package ipv6

func hexToInt(hex byte) int64 {
    if hex >= '0' && hex <= '9' {
        return int64(hex - '0')
    } else {
        return 10 + int64(hex - 'a')
    }
}
