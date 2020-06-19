#ifndef NDN_DPDK_NDN_ENCODE_DATA_H
#define NDN_DPDK_NDN_ENCODE_DATA_H

/// \file

#include "name.h"

void
EncodeData_(struct rte_mbuf* m,
            uint16_t namePrefixL,
            const uint8_t* namePrefixV,
            uint16_t nameSuffixL,
            const uint8_t* nameSuffixV,
            uint32_t freshnessPeriod,
            uint16_t contentL,
            const uint8_t* contentV);

/** \brief Encode a Data.
 *  \param m output mbuf, must be empty and is the only segment, must have
 *           \c EncodeData_GetHeadroom() in headroom and
 *           <tt>EncodeData_GetTailroom(namePrefix.length + nameSuffix.length,
 *           contentL)</tt> in tailroom; headroom for Ethernet and NDNLP
 *           headers may be included if needed.
 *  \param contentV the payload, will be copied.
 */
static inline void
EncodeData(struct rte_mbuf* m,
           LName namePrefix,
           LName nameSuffix,
           uint32_t freshnessPeriod,
           uint16_t contentL,
           const uint8_t* contentV)
{
  EncodeData_(m,
              namePrefix.length,
              namePrefix.value,
              nameSuffix.length,
              nameSuffix.value,
              freshnessPeriod,
              contentL,
              contentV);
}

/** \brief Data encoder optimized for NdnpingServer.
 *
 *  The main difference from \c EncodeData() is that, DataGen puts everything
 *  except name prefix into a template, and then creates two-segment packets
 *  where the second segment references the template. It's faster for traffic
 *  generator use case, but does not allow changing Content payload.
 */
typedef struct DataGen
{
} DataGen;

DataGen*
MakeDataGen_(struct rte_mbuf* m,
             uint16_t nameSuffixL,
             const uint8_t* nameSuffixV,
             uint32_t freshnessPeriod,
             uint16_t contentL,
             const uint8_t* contentV);

/** \brief Prepare DataGen template.
 *  \param m template mbuf, must be empty and is the only segment, must have
 *           <tt>DataGen_GetTailroom1(nameSuffix.length, contentL)</tt> in
 *           tailroom. DataGen takes ownership of this mbuf.
 */
DataGen*
MakeDataGen(struct rte_mbuf* m,
            LName nameSuffix,
            uint32_t freshnessPeriod,
            uint16_t contentL,
            const uint8_t* contentV);

void
DataGen_Close(DataGen* gen);

void
DataGen_Encode_(DataGen* gen,
                struct rte_mbuf* seg0,
                struct rte_mbuf* seg1,
                uint16_t namePrefixL,
                const uint8_t* namePrefixV);

/** \brief Encode Data with DataGen template.
 *  \param seg0 segment 0 mbuf, must be empty and is the only segment, must
 *              have \c DataGen_GetHeadroom0() in headroom and
 *              <tt>DataGen_GetTailroom0(namePrefix.length)</tt> in tailroom.
 *              This becomes the encoded Data packet.
 *  \param seg1 segment 1 indirect mbuf. This is chained onto \p seg0 .
 */
static inline void
DataGen_Encode(DataGen* gen,
               struct rte_mbuf* seg0,
               struct rte_mbuf* seg1,
               LName namePrefix)
{
  DataGen_Encode_(gen, seg0, seg1, namePrefix.length, namePrefix.value);
}

#endif // NDN_DPDK_NDN_ENCODE_DATA_H
