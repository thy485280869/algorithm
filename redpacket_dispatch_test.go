package algorithm

import (
	"fmt"
	"testing"
)

func TestRedPacket(t *testing.T) {
	res, err := RedPacket(100, 3)
	if err != nil {
		t.Errorf(err.Error())
	}
	sum := 0
	for i := 0; i < len(res); i++ {
		fmt.Println(fmt.Sprintf("第%d个人抢到的红包为%d(%.2f元)", i+1, res[i], float64(res[i])/100))
		sum += res[i]
	}
	if sum != 100 {
		t.Errorf("")
	}

	if res, err = RedPacket(100, 110); err == nil || err.Error() != ErrMoneyTooLittle {
		t.Errorf(err.Error())
	}
	if res, err = RedPacket(100, 10); err != nil {
		t.Errorf(err.Error())
	}

	if res, err = RedPacket(1000, 501); err == nil || err.Error() != ErrReaPacketTooMuch {
		t.Errorf(err.Error())
	}
	if res, err = RedPacket(500, 500); err != nil {
		t.Errorf(err.Error())
	}
	if res, err = RedPacket(600, 500); err != nil {
		t.Errorf(err.Error())
	}

}
