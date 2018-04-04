#ifndef NDN_DPDK_STRATEGY_API_STRUCT_H
#define NDN_DPDK_STRATEGY_API_STRUCT_H

/// \file

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <stdlib.h>

/** \brief Indicate why the strategy program is invoked.
 */
typedef enum SgEvent {
  SGEVT_NONE = 0,
  SGEVT_TIMER = 1,    ///< timer expires
  SGEVT_INTEREST = 2, ///< Interest arrives
} SgEvent;

typedef uint16_t FaceId;

/** \brief Context of strategy invocation.
 */
typedef struct SgCtx
{
  SgEvent eventKind;
  FaceId* nexthops;
  uint8_t nNexthops;
} SgCtx;

#endif // NDN_DPDK_STRATEGY_API_STRUCT_H
