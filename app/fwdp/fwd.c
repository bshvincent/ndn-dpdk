#include "fwd.h"

#include "../../core/logger.h"

INIT_ZF_LOG(FwFwd);

#define FW_FWD_BURST_SIZE 16

void
FwFwd_Close(FwFwd* fwd)
{
  Packet* npkt;
  while (rte_ring_dequeue(fwd->queue, (void**)&npkt) == 0) {
    rte_pktmbuf_free(Packet_ToMbuf(npkt));
  }
  rte_ring_free(fwd->queue);
  rte_free(fwd);
}

typedef void (*FwFwd_RxFunc)(FwFwd* fwd, Packet* npkt);
static const FwFwd_RxFunc FwFwd_RxFuncs[L3PktType_MAX] = {
  NULL, FwFwd_RxInterest, FwFwd_RxData, FwFwd_RxNack,
};

void
FwFwd_Run(FwFwd* fwd)
{
  assert((void*)fwd->pit == (void*)fwd->cs);
  ZF_LOGI("fwdId=%" PRIu8 " fwd=%p queue=%p pit+cs=%p", fwd->id, fwd,
          fwd->queue, fwd->pcct);

  Packet* npkts[FW_FWD_BURST_SIZE];
  while (!fwd->stop) {
    rcu_quiescent_state();
    MinSched_Trigger(Pit_GetPriv(fwd->pit)->timeoutSched);

    unsigned count = rte_ring_dequeue_burst(fwd->queue, (void**)npkts,
                                            FW_FWD_BURST_SIZE, NULL);
    TscTime now = rte_get_tsc_cycles();
    for (unsigned i = 0; i < count; ++i) {
      Packet* npkt = npkts[i];
      TscDuration timeSinceRx = now - Packet_ToMbuf(npkt)->timestamp;
      RunningStat_Push(&fwd->latencyStat, timeSinceRx);

      L3PktType l3type = Packet_GetL3PktType(npkt);
      assert(l3type != L3PktType_None && l3type < L3PktType_MAX);
      FwFwd_RxFunc rxFunc = FwFwd_RxFuncs[l3type];
      (*rxFunc)(fwd, npkt);
    }
  }

  ZF_LOGI("fwdId=%" PRIu8 " STOP", fwd->id);
}
