package hash

import (
	"fmt"
	"math/big"
	"sort"
	"testing"
	"time"
)

func TestMD5String(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "shuai.yang",
			args: args{
				s: "shuai.yang\n",
			},
			want: "0f016edeb148f9a7c826c4dd87418b44",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5String(tt.args.s); got != tt.want {
				t.Errorf("MD5String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMd5ToBigInt(t *testing.T) {

	nums := []big.Int{}
	index := map[string]int{}
	for i := 0; i < 20; i++ {
		s := "shuai.yang+1234" + fmt.Sprint(i)
		nums = append(nums, Md5ToBigInt(s))
		index[nums[i].String()] = i
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i].Cmp(&nums[j]) <= 0
	})
	for i, v := range nums {
		key := v.String()
		fmt.Println(i, key, index[key])
	}
	fmt.Println("==========")
	for i := 100; i < 200; i++ {
		target := Md5ToBigInt(fmt.Sprintf("%d+%s", i, time.Now()))
		x := sort.Search(len(nums), func(i int) bool {
			return nums[i].Cmp(&target) >= 0
		})
		if x < 0 {
			x = 0
		}
		if x == len(nums) {
			x--
		}
		xs := nums[x].String()
		fmt.Printf("lucky number: %s rank: %d hit number: %s raw index:%d\n", target.String(), x, xs, index[xs])
	}
}

func TestBigInt2String(t *testing.T) {
	s := "shuai.yang+1234"
	s2 := MD5String(s)
	num := Md5ToBigInt(s)
	s3 := num.Text(16)
	t.Log(s2 == s3, s, s2, s3)
}
