// +build ignore

package ndni

/*
#include "../csrc/ndn/interest.h"
*/
import "C"

// Template to encode an Interest.
type InterestTemplate C.InterestTemplate