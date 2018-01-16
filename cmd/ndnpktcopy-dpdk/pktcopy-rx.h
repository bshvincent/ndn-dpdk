#ifndef NDN_DPDK_CMD_NDNPKTCOPY_RX_H
#define NDN_DPDK_CMD_NDNPKTCOPY_RX_H

#include "../../iface/face.h"

#define PKTCOPYRX_MAXTX 16

typedef struct PktcopyRx
{
  Face* face;
  struct rte_mempool* mpIndirect;
  struct rte_ring* txRings[PKTCOPYRX_MAXTX];
  int nTxRings;
} PktcopyRx;

void PktcopyRx_AddTxRing(PktcopyRx* pcrx, struct rte_ring* r);

void PktcopyRx_Run(PktcopyRx* pcrx);

#endif // NDN_DPDK_CMD_NDNPKTCOPY_RX_H
