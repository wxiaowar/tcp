package utils

import (
	"errors"
	"im/comet/proto"
	"im/pkg/log"
)

var (
	ErrRingEmpty = errors.New("ring buffer empty")
	ErrRingFull  = errors.New("ring buffer full")
)

type Ring struct {
	// read
	rp   uint64
	num  uint64
	mask uint64
	// TODO split cacheline, many cpu cache line size is 64
	// pad [40]byte
	// write
	wp   uint64
	data []proto.Proto
}

func NewRing(num int) *Ring {
	r := new(Ring)
	r.init(uint64(num))
	return r
}

func (r *Ring) Init(num int) {
	r.init(uint64(num))
}

func (r *Ring) init(num uint64) {
	// 2^N
	if num&(num-1) != 0 {
		for num&(num-1) != 0 {
			num &= (num - 1)
		}
		num = num << 1
	}
	r.data = make([]proto.Proto, num)
	r.num = num
	r.mask = r.num - 1
}

func (r *Ring) Get() (proto *proto.Proto, err error) {
	if r.rp == r.wp {
		return nil, ErrRingEmpty
	}
	proto = &r.data[r.rp&r.mask]
	return
}

func (r *Ring) GetAdv() {
	r.rp++
	log.Debug("ring rp: %d, idx: %d", r.rp, r.rp&r.mask)
}

func (r *Ring) Set() (proto *proto.Proto, err error) {
	if r.wp-r.rp >= r.num {
		return nil, ErrRingFull
	}
	proto = &r.data[r.wp&r.mask]
	return
}

func (r *Ring) SetAdv() {
	r.wp++
	log.Debug("ring wp: %d, idx: %d", r.wp, r.wp&r.mask)
}

func (r *Ring) Reset() {
	r.rp = 0
	r.wp = 0
	// prevent pad compiler optimization
	// r.pad = [40]byte{}
}
