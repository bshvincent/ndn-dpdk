package cptrtest

/*
int callbackCArg = 0;

int callbackC(void* arg) {
	return 1 + *(int*)arg;
}
*/
import "C"
import (
	"unsafe"

	"github.com/usnistgov/ndn-dpdk/core/cptr"
)

func makeCFunction(arg int) cptr.Function {
	C.callbackCArg = C.int(arg)
	return cptr.CFunction(C.callbackC, unsafe.Pointer(&C.callbackCArg))
}
