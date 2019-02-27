// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Panel struct {
	_tab flatbuffers.Table
}

func GetRootAsPanel(buf []byte, offset flatbuffers.UOffsetT) *Panel {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Panel{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Panel) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Panel) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Panel) SupplementId() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Panel) MutateSupplementId(n uint16) bool {
	return rcv._tab.MutateUint16Slot(4, n)
}

func PanelStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func PanelAddSupplementId(builder *flatbuffers.Builder, supplementId uint16) {
	builder.PrependUint16Slot(0, supplementId, 0)
}
func PanelEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}