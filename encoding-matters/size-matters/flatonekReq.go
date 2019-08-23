// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package main

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type OneKReq struct {
	_tab flatbuffers.Table
}

func GetRootAsOneKReq(buf []byte, offset flatbuffers.UOffsetT) *OneKReq {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &OneKReq{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *OneKReq) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *OneKReq) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *OneKReq) Id() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *OneKReq) MutateId(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *OneKReq) Data(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *OneKReq) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *OneKReq) MutateData(j int, n int8) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func OneKReqStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func OneKReqAddId(builder *flatbuffers.Builder, id uint64) {
	builder.PrependUint64Slot(0, id, 0)
}
func OneKReqAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(data), 0)
}
func OneKReqStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func OneKReqEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}