// +build ignore

package eal

/*
#include "../../csrc/core/common.h"
#include <rte_pci.h>
*/
import "C"

// PCI address.
type PciAddress C.struct_rte_pci_addr
