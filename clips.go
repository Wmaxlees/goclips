package clips

/*
#cgo LDFLAGS: -lm
#cgo CFLAGS: -Wunused-result
#include "clips.h"

void* env = NULL;
*/
import "C"
import "unsafe"

func CreateEnvironment() {
    C.env, _ = C.CreateEnvironment()
}

func Load(filename string) {
    sbytes := []byte(filename)
    ccp := (*C.char)(unsafe.Pointer(&sbytes[0]))

    C.EnvLoad(C.env, ccp)
}

func Reset() {
    C.EnvReset(C.env)
}

func AssertString(fact string) {
    sbytes := []byte(fact)
    ccp := (*C.char)(unsafe.Pointer(&sbytes[0]))

    C.EnvAssertString(C.env, ccp)
}

func Run(steps int) {
    ll := C.longlong(steps)

    C.EnvRun(C.env, ll)
}

func Build(cs string) {
    sbytes := []byte(cs)
    ccp := (*C.char)(unsafe.Pointer(&sbytes[0]))

    C.EnvBuild(C.env, ccp)
}

type strategy int

const (
    DEPTH_STRATEGY strategy = C.DEPTH_STRATEGY
    BREADTH_STRATEGY strategy = C.BREADTH_STRATEGY
    LEX_STRATEGY strategy = C.LEX_STRATEGY
    MEA_STRATEGY strategy = C.MEA_STRATEGY
    COMPLEXITY_STRATEGY strategy = C.COMPLEXITY_STRATEGY
    SIMPLICITY_STRATEGY strategy = C.SIMPLICITY_STRATEGY
    RANDOM_STRATEGY strategy = C.RANDOM_STRATEGY
)

func SetStrategy(strat strategy) {
    C.EnvSetStrategy(C.env, C.int(strat))
}
