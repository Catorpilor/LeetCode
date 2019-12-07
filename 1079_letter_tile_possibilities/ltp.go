package ltp

func numOfPossibilities(tiles string) int {
    res := make(map[string]bool)
    helper(res, []byte(tiles), 0)
    return len(res)
}

func helper(res map[string]bool, st []byte, idx int) {
    if idx == len(st) {
        return
    }
    c := st[idx]
    curKeys := make([]string, 0, len(res))
    for k := range res {
        curKeys = append(curKeys, k)
    }
    for _, k := range curKeys {
        // fmt.Printf("key: %s, c:%s\n", k, string(c))
        n := len(k)
        sk := make([]byte, n+1)
        copy(sk, k)
        sk[n] = c
        for i := 0; i < n; i++ {
            sk[i], sk[n] = sk[n], sk[i]
            res[string(sk)] = true
            sk[i], sk[n] = sk[n], sk[i]
        }
        res[string(sk)] = true
    }
    res[string(c)] = true
    helper(res, st, idx+1)
}
