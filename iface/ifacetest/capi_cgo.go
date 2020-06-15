package ifacetest

/*
#include "../../csrc/iface/face.h"
*/
import "C"
import (
	"ndn-dpdk/core/cptr"
	"ndn-dpdk/iface"
	"ndn-dpdk/ndn"
)

func Face_IsDown(faceId iface.FaceId) bool {
	return bool(C.Face_IsDown(C.FaceId(faceId)))
}

func Face_TxBurst(faceId iface.FaceId, pkts []*ndn.Packet) {
	ptr, count := cptr.ParseCptrArray(pkts)
	C.Face_TxBurst(C.FaceId(faceId), (**C.Packet)(ptr), C.uint16_t(count))
}
