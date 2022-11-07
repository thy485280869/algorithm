package algorithm

import (
	"errors"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

const (
	RED_PACKET_LIMIT    = 500 // 红包数量最大限制
	ErrMoneyTooLittle   = "single redpacket money must >= 0.01(yuan)"
	ErrReaPacketTooMuch = "redpacket num greater than limit!"
)

// redPacket : m分钱随机分成n份，返回每份的金额(分)
// 传入的m，n必须保证 m>n
func redPacket(m, n int) []int {
	splitMap := make(map[int]struct{})
	splitSlice := make([]int, 0, n-1)
	for i := 0; i < n-1; i++ {
		t := rand.Intn(m)
		if _, ok := splitMap[t]; ok || t == 0 { // t==0时会出现0元钱，所以要重新选
			i--
			continue
		}
		splitMap[t] = struct{}{}
		splitSlice = append(splitSlice, t) // 最终存入的t一定在[1,m)范围内，且无重复
	}
	sort.Ints(splitSlice)
	res := make([]int, 0, n)
	now := 0
	for i := 0; i < n-1; i++ {
		res = append(res, splitSlice[i]-now)
		now = splitSlice[i]
	}
	res = append(res, m-now)
	return res
}

// RedPacket : m分钱随机分成n份，返回每份的金额(分)
func RedPacket(m, n int) ([]int, error) {
	if m < n { // single redpacket money must >= 0.01(yuan)
		return nil, errors.New(ErrMoneyTooLittle)
	} else if n > RED_PACKET_LIMIT {
		return nil, errors.New(ErrReaPacketTooMuch)
	} else if m == n {
		res := make([]int, n)
		for i := 0; i < len(res); i++ {
			res[i] = 1
		}
		return res, nil
	} else {
		return redPacket(m, n), nil
	}
}
