package createface

import (
	"ndn-dpdk/iface"
	"ndn-dpdk/iface/ethface"
)

func handleFaceClosing(id iface.FaceId) {
	if !isInitialized {
		return
	}
	face := iface.Get(id)
	switch id.GetKind() {
	case iface.FaceKind_Mock, iface.FaceKind_Socket:
		stopSmRxtx(face)
	case iface.FaceKind_Eth:
		stopEthFaceRxtx(face.(*ethface.EthFace))
	}
}

func handleFaceClosed(id iface.FaceId) {
	if !isInitialized || id.GetKind() != iface.FaceKind_Eth {
		return
	}
	for _, port := range ethface.ListPorts() {
		if port.CountFaces() == 0 {
			stopEthPortRxtx(port)
		}
	}
}

var (
	theFaceClosingEvt = iface.OnFaceClosing(handleFaceClosing)
	theFaceClosedEvt  = iface.OnFaceClosed(handleFaceClosed)
)