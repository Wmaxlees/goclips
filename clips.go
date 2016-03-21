package clips

// #cgo LDFLAGS: -lm
// #include "clips.h"
import "C"
import "unsafe"

var env unsafe.Pointer

func CreateEnvironment() {
    env, _ = C.CreateEnvironment()
}

func Load(filename string) {
    sbytes := []byte(filename)
    ccp := (*C.char)(unsafe.Pointer(&sbytes[0]))

    C.EnvLoad(env, ccp)
}

func Reset() {
    vp := (*C.void)(env)
    C.EnvReset(vp)
}

func AssertString(fact string) {
    sbytes := []byte(fact)
    ccp := (*C.char)(unsafe.Pointer(&sbytes[0]))

    C.EnvAssertString(env, ccp)
}

func Run(steps int) {
    ll := C.longlong(steps)

    C.EnvRun(env, ll)
}
