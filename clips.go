package clips

/*
#cgo LDFLAGS: -lm
#cgo CFLAGS: -w
#include "clips.h"
#include <string.h>
#include <stdio.h>

const char *DOPToStr(DATA_OBJECT *obj) {
    return DOPToString(obj);
}

int GetTypePFunc(DATA_OBJECT *obj) {
    return obj->type;
}

DATA_OBJECT makeDataObject() {
    DATA_OBJECT obj;
    return obj;
}

void ConvFactListToString(void *env, DATA_OBJECT *obj) {
    char *result = "";

    if (GetpType(obj) != MULTIFIELD) {
        return;
    }

    long end = GetpDOEnd(obj);
    void *multifieldPtr = GetpValue(obj);
    long i = GetpDOBegin(obj);
    //for (long i = GetpDOBegin(obj); i <= end; ++i)
    //{
        if (GetMFType(multifieldPtr, i) == FACT_ADDRESS) {
            void *factPtr = GetMFValue(multifieldPtr, i);
            printf("%i -- ", factPtr);

            factPtr = ValueToPointer(factPtr);
            printf("%i\n", factPtr);

            // void *factPtr = ValueToPointer(GetMFValue(multifieldPtr, i));
            char *fact = malloc(sizeof(char)*500);

            // struct fact* theFact = factPtr;

            // struct multiField* theProp = &(theFact->theProposition);

            // printf("%i\n", theFact->factIndex);

            // ERROR HERE ____________________________________
            EnvGetFactPPForm(env, &fact, 500, factPtr);
            // -----------------------------------------------
            // printf("%s\n", fact);

            // strcat(result, &fact);
        }
    //}

    // return result;
}

void GetFactListAsString(void *env) {
    DATA_OBJECT list;
    EnvEval(env, "(facts)", &list);

    printf("%i\n", GetType(list));
    printf("##%s##\n", ValueToString(GetValue(list)));
    // DATA_OBJECT* factList = malloc(sizeof(DATA_OBJECT));
    // EnvGetFactList(env, factList, NULL);

    // long end = GetpDOEnd(factList);
    // void *multifieldPtr = GetpValue(factList);
    // long i = GetpDOBegin(factList);

    // char *fact = malloc(sizeof(char)*500);
    // void *factPtr = GetMFValue(multifieldPtr, i);

    // EnvGetFactPPForm(env, &fact, 500, factPtr);
}

*/
import "C"
import "unsafe"

type EnvironmentPointer unsafe.Pointer

func CreateEnvironment() EnvironmentPointer {
    env, _ := C.CreateEnvironment()

    return EnvironmentPointer(env)
}

func Load(env EnvironmentPointer, filename string) {
    sbytes := []byte(filename)
    ccp := (*C.char)(unsafe.Pointer(&sbytes[0]))

    C.EnvLoad(unsafe.Pointer(env), ccp)
}

func Reset(env EnvironmentPointer) {
    C.EnvReset(unsafe.Pointer(env))
}

func AssertString(env EnvironmentPointer, fact string) {
    // sbytes := []byte(fact)
    // ccp := (*C.char)(unsafe.Pointer(&sbytes[0]))
    ccp := C.CString(fact)

    C.EnvAssertString(unsafe.Pointer(env), ccp)
}

func Run(env EnvironmentPointer, steps int) {
    ll := C.longlong(steps)

    C.EnvRun(unsafe.Pointer(env), ll)
}

func Build(env EnvironmentPointer, cs string) {
    sbytes := []byte(cs)
    ccp := (*C.char)(unsafe.Pointer(&sbytes[0]))

    C.EnvBuild(unsafe.Pointer(env), ccp)
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

func SetStrategy(env EnvironmentPointer, strat strategy) {
    C.EnvSetStrategy(unsafe.Pointer(env), C.int(strat))
}

type ActivationPointer unsafe.Pointer

func GetNextActivation(env EnvironmentPointer) ActivationPointer {
    ret := ActivationPointer(C.EnvGetNextActivation(unsafe.Pointer(env), nil))
    return ret
}

func GetActivationPPForm(env EnvironmentPointer, activation ActivationPointer, size int) string {
    var buffer []byte
    buffer = make([]byte, size, size);
    ccp := (*C.char)(unsafe.Pointer(&buffer[0]))

    C.EnvGetActivationPPForm(unsafe.Pointer(env), ccp, C.size_t(size), unsafe.Pointer(activation))

    return string(buffer[:])
}

func GetActivationName(env EnvironmentPointer, activation ActivationPointer) string {
    ccp := C.EnvGetActivationName(unsafe.Pointer(env), unsafe.Pointer(activation))

    return C.GoString(ccp)
}

type GlobalPointer unsafe.Pointer

func FindGlobal(env EnvironmentPointer, name string) GlobalPointer {
    ptr := C.EnvFindDefglobal(unsafe.Pointer(env), C.CString(name))

    return GlobalPointer(ptr)
}

func GetGlobalValueForm(env EnvironmentPointer, ptr GlobalPointer, size int) string {
    var buffer []byte
    buffer = make([]byte, size, size);
    ccp := (*C.char)(unsafe.Pointer(&buffer[0]))

    C.EnvGetDefglobalValueForm(unsafe.Pointer(env), ccp, C.size_t(size), unsafe.Pointer(ptr))

    return C.GoString(ccp)
}

func Clear(env EnvironmentPointer) {
    C.EnvClear(unsafe.Pointer(env))
}

type FactPointer unsafe.Pointer

func GetFactList(env EnvironmentPointer) string {
    //factList := new(C.DATA_OBJECT)

    // C.EnvGetFactList(unsafe.Pointer(env), factList, nil)

    // C.ConvFactListToString(unsafe.Pointer(env), factList)
    C.GetFactListAsString(unsafe.Pointer(env));
    return "";

    // return C.GoString(C.ConvFactListToString(&factList))

    // fmt.Printf(strconv.Itoa(int(C.GetTypePFunc(&factList))))
    // a := C.DOPToStr(factList)

    // return C.GoString(a)
}

