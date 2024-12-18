/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (https://www.swig.org).
 * Version 4.3.0
 *
 * Do not make changes to this file unless you know what you are doing - modify
 * the SWIG interface file instead.
 * ----------------------------------------------------------------------------- */

// source: /client/ubiquityclient.i

package Ubiquity

/*
#cgo LDFLAGS: -L. -lubiquityclientwrapper -lubiquity_client_static -lbsd -lcurl -luuid -lstdc++ -lm -ldl

#define intgo swig_intgo
typedef void *swig_voidp;

#include <stddef.h>
#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef long long swig_type_1;
typedef long long swig_type_2;
typedef _gostring_ swig_type_3;
typedef _gostring_ swig_type_4;
typedef long long swig_type_5;
typedef _gostring_ swig_type_6;
typedef _gostring_ swig_type_7;
typedef long long swig_type_8;
typedef _gostring_ swig_type_9;
typedef long long swig_type_10;
typedef _gostring_ swig_type_11;
typedef long long swig_type_12;
typedef _gostring_ swig_type_13;
typedef long long swig_type_14;
typedef _gostring_ swig_type_15;
typedef _gostring_ swig_type_16;
typedef _gostring_ swig_type_17;
typedef _gostring_ swig_type_18;
typedef _gostring_ swig_type_19;
typedef _gostring_ swig_type_20;
typedef _gostring_ swig_type_21;
typedef long long swig_type_22;
typedef _gostring_ swig_type_23;
extern void _wrap_Swig_free_Ubiquity_efad76b24322b4d4(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_Ubiquity_efad76b24322b4d4(swig_intgo arg1);
extern uintptr_t _wrap_new_ClientHandle_Ubiquity_efad76b24322b4d4(void);
extern void _wrap_ClientHandle_close_Ubiquity_efad76b24322b4d4(uintptr_t arg1);
extern swig_type_1 _wrap_ClientHandle_getCreativeId_Ubiquity_efad76b24322b4d4(uintptr_t arg1);
extern swig_type_2 _wrap_ClientHandle_getQueryId_Ubiquity_efad76b24322b4d4(uintptr_t arg1);
extern swig_type_3 _wrap_ClientHandle_getModelId_Ubiquity_efad76b24322b4d4(uintptr_t arg1);
extern swig_intgo _wrap_ClientHandle_getStatusCode_Ubiquity_efad76b24322b4d4(uintptr_t arg1);
extern swig_type_4 _wrap_ClientHandle_getStatusString_Ubiquity_efad76b24322b4d4(uintptr_t arg1);
extern float _wrap_ClientHandle_getPrice_Ubiquity_efad76b24322b4d4(uintptr_t arg1);
extern swig_intgo _wrap_ClientHandle_getPredictionType_Ubiquity_efad76b24322b4d4(uintptr_t arg1);
extern float _wrap_ClientHandle_getProbability_Ubiquity_efad76b24322b4d4(uintptr_t arg1);
extern swig_intgo _wrap_ClientHandle_poll_Ubiquity_efad76b24322b4d4(uintptr_t arg1);
extern swig_intgo _wrap_ClientHandle_predict_Ubiquity_efad76b24322b4d4(uintptr_t arg1);
extern void _wrap_ClientHandle_putBoolean_Ubiquity_efad76b24322b4d4(uintptr_t arg1, char arg2);
extern void _wrap_ClientHandle_putCategorical_Ubiquity_efad76b24322b4d4(uintptr_t arg1, char arg2);
extern void _wrap_ClientHandle_putCreative_Ubiquity_efad76b24322b4d4(uintptr_t arg1, swig_type_5 arg2);
extern void _wrap_ClientHandle_putNominal_Ubiquity_efad76b24322b4d4(uintptr_t arg1, swig_type_6 arg2);
extern void _wrap_ClientHandle_putNumeric_Ubiquity_efad76b24322b4d4(uintptr_t arg1, float arg2);
extern void _wrap_ClientHandle_startBidShading_Ubiquity_efad76b24322b4d4(uintptr_t arg1, swig_type_7 arg2, swig_type_8 arg3, float arg4, float arg5);
extern void _wrap_ClientHandle_startLineItem_Ubiquity_efad76b24322b4d4(uintptr_t arg1, swig_type_9 arg2, swig_type_10 arg3, float arg4);
extern void _wrap_ClientHandle_startProbability_Ubiquity_efad76b24322b4d4(uintptr_t arg1, swig_type_11 arg2, swig_type_12 arg3);
extern swig_intgo _wrap_ClientHandle_train_Ubiquity_efad76b24322b4d4(uintptr_t arg1, swig_type_13 arg2, swig_type_14 arg3, swig_intgo arg4);
extern uintptr_t _wrap_getClientHandle__SWIG_0_Ubiquity_efad76b24322b4d4(void);
extern void _wrap_delete_ClientHandle_Ubiquity_efad76b24322b4d4(uintptr_t arg1);
extern swig_type_15 _wrap_getStatusString_Ubiquity_efad76b24322b4d4(swig_intgo arg1);
extern void _wrap_setAeronDirectory_Ubiquity_efad76b24322b4d4(swig_type_16 arg1);
extern void _wrap_setCredential_Ubiquity_efad76b24322b4d4(swig_type_17 arg1);
extern void _wrap_setApiAddress_Ubiquity_efad76b24322b4d4(swig_type_18 arg1);
extern void _wrap_setResourceStringInCache_Ubiquity_efad76b24322b4d4(swig_type_19 arg1, swig_type_20 arg2);
extern uintptr_t _wrap_getClientHandle__SWIG_1_Ubiquity_efad76b24322b4d4(void);
extern swig_intgo _wrap_openClient_Ubiquity_efad76b24322b4d4(swig_type_21 arg1);
extern void _wrap_closeClient_Ubiquity_efad76b24322b4d4(void);
extern swig_type_22 _wrap_xxh3_64bits_Ubiquity_efad76b24322b4d4(swig_type_23 arg1);
#undef intgo
*/
import "C"
import "math"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


func getSwigcptr(v interface { Swigcptr() uintptr }) uintptr {
	if v == nil {
		return 0
	}
	return v.Swigcptr()
}


type _ sync.Mutex


type swig_gostring struct { p uintptr; n int }
func swigCopyString(s string) string {
  p := *(*swig_gostring)(unsafe.Pointer(&s))
  r := string((*[0x7fffffff]byte)(unsafe.Pointer(p.p))[:p.n])
  Swig_free(p.p)
  return r
}

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_Ubiquity_efad76b24322b4d4(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrClientHandle uintptr

func (p SwigcptrClientHandle) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrClientHandle) SwigIsClientHandle() {
}

func NewClientHandle() (_swig_ret ClientHandle) {
	var swig_r ClientHandle
	swig_r = (ClientHandle)(SwigcptrClientHandle(C._wrap_new_ClientHandle_Ubiquity_efad76b24322b4d4()))
	return swig_r
}

func (arg1 SwigcptrClientHandle) Close() {
	_swig_i_0 := arg1
	C._wrap_ClientHandle_close_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrClientHandle) GetCreativeId() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_ClientHandle_getCreativeId_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrClientHandle) GetQueryId() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_ClientHandle_getQueryId_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrClientHandle) GetModelId() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_ClientHandle_getModelId_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrClientHandle) GetStatusCode() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_ClientHandle_getStatusCode_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrClientHandle) GetStatusString() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_ClientHandle_getStatusString_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrClientHandle) GetPrice() (_swig_ret float32) {
	var swig_r float32
	_swig_i_0 := arg1
	swig_r = (float32)(C._wrap_ClientHandle_getPrice_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrClientHandle) GetPredictionType() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_ClientHandle_getPredictionType_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrClientHandle) GetProbability() (_swig_ret float32) {
	var swig_r float32
	_swig_i_0 := arg1
	swig_r = (float32)(C._wrap_ClientHandle_getProbability_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrClientHandle) Poll() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_ClientHandle_poll_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrClientHandle) Predict() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_ClientHandle_predict_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrClientHandle) PutBoolean(arg2 int8) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_ClientHandle_putBoolean_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0), C.char(_swig_i_1))
}

func (arg1 SwigcptrClientHandle) PutCategorical(arg2 int8) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_ClientHandle_putCategorical_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0), C.char(_swig_i_1))
}

func (arg1 SwigcptrClientHandle) PutCreative(arg2 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_ClientHandle_putCreative_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0), C.swig_type_5(_swig_i_1))
}

func (arg1 SwigcptrClientHandle) PutNominal(arg2 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_ClientHandle_putNominal_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0), *(*C.swig_type_6)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrClientHandle) PutNumeric(arg2 float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_ClientHandle_putNumeric_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0), C.float(_swig_i_1))
}

func (arg1 SwigcptrClientHandle) StartBidShading(arg2 string, arg3 int64, arg4 float32, arg5 float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	C._wrap_ClientHandle_startBidShading_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0), *(*C.swig_type_7)(unsafe.Pointer(&_swig_i_1)), C.swig_type_8(_swig_i_2), C.float(_swig_i_3), C.float(_swig_i_4))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrClientHandle) StartLineItem(arg2 string, arg3 int64, arg4 float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	C._wrap_ClientHandle_startLineItem_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0), *(*C.swig_type_9)(unsafe.Pointer(&_swig_i_1)), C.swig_type_10(_swig_i_2), C.float(_swig_i_3))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrClientHandle) StartProbability(arg2 string, arg3 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_ClientHandle_startProbability_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0), *(*C.swig_type_11)(unsafe.Pointer(&_swig_i_1)), C.swig_type_12(_swig_i_2))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrClientHandle) Train(arg2 string, arg3 int64, arg4 int) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	swig_r = (int)(C._wrap_ClientHandle_train_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0), *(*C.swig_type_13)(unsafe.Pointer(&_swig_i_1)), C.swig_type_14(_swig_i_2), C.swig_intgo(_swig_i_3)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func GetClientHandle__SWIG_0() (_swig_ret ClientHandle) {
	var swig_r ClientHandle
	swig_r = (ClientHandle)(SwigcptrClientHandle(C._wrap_getClientHandle__SWIG_0_Ubiquity_efad76b24322b4d4()))
	return swig_r
}

func DeleteClientHandle(arg1 ClientHandle) {
	_swig_i_0 := getSwigcptr(arg1)
	C._wrap_delete_ClientHandle_Ubiquity_efad76b24322b4d4(C.uintptr_t(_swig_i_0))
}

type ClientHandle interface {
	Swigcptr() uintptr
	SwigIsClientHandle()
	Close()
	GetCreativeId() (_swig_ret int64)
	GetQueryId() (_swig_ret int64)
	GetModelId() (_swig_ret string)
	GetStatusCode() (_swig_ret int)
	GetStatusString() (_swig_ret string)
	GetPrice() (_swig_ret float32)
	GetPredictionType() (_swig_ret int)
	GetProbability() (_swig_ret float32)
	Poll() (_swig_ret int)
	Predict() (_swig_ret int)
	PutBoolean(arg2 int8)
	PutCategorical(arg2 int8)
	PutCreative(arg2 int64)
	PutNominal(arg2 string)
	PutNumeric(arg2 float32)
	StartBidShading(arg2 string, arg3 int64, arg4 float32, arg5 float32)
	StartLineItem(arg2 string, arg3 int64, arg4 float32)
	StartProbability(arg2 string, arg3 int64)
	Train(arg2 string, arg3 int64, arg4 int) (_swig_ret int)
}

func GetStatusString(arg1 int) (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_getStatusString_Ubiquity_efad76b24322b4d4(C.swig_intgo(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func SetAeronDirectory(arg1 string) {
	_swig_i_0 := arg1
	C._wrap_setAeronDirectory_Ubiquity_efad76b24322b4d4(*(*C.swig_type_16)(unsafe.Pointer(&_swig_i_0)))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
}

func SetCredential(arg1 string) {
	_swig_i_0 := arg1
	C._wrap_setCredential_Ubiquity_efad76b24322b4d4(*(*C.swig_type_17)(unsafe.Pointer(&_swig_i_0)))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
}

func SetApiAddress(arg1 string) {
	_swig_i_0 := arg1
	C._wrap_setApiAddress_Ubiquity_efad76b24322b4d4(*(*C.swig_type_18)(unsafe.Pointer(&_swig_i_0)))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
}

func SetResourceStringInCache(arg1 string, arg2 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_setResourceStringInCache_Ubiquity_efad76b24322b4d4(*(*C.swig_type_19)(unsafe.Pointer(&_swig_i_0)), *(*C.swig_type_20)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func GetClientHandle__SWIG_1() (_swig_ret ClientHandle) {
	var swig_r ClientHandle
	swig_r = (ClientHandle)(SwigcptrClientHandle(C._wrap_getClientHandle__SWIG_1_Ubiquity_efad76b24322b4d4()))
	return swig_r
}

func GetClientHandle(a ...interface{}) ClientHandle {
	argc := len(a)
	if argc == 0 {
		return GetClientHandle__SWIG_0()
	}
	if argc == 0 {
		return GetClientHandle__SWIG_1()
	}
	panic("No match for overloaded function call")
}

func OpenClient(arg1 string) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_openClient_Ubiquity_efad76b24322b4d4(*(*C.swig_type_21)(unsafe.Pointer(&_swig_i_0))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func CloseClient() {
	C._wrap_closeClient_Ubiquity_efad76b24322b4d4()
}

func Xxh3_64bits(arg1 string) (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_xxh3_64bits_Ubiquity_efad76b24322b4d4(*(*C.swig_type_23)(unsafe.Pointer(&_swig_i_0))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}



const MISSING_BOOLEAN int8 = -1
const MISSING_CATEGORICAL int8 = -1
const MISSING_NOMINAL string = "_MISSING_"
var MISSING_NUMERIC float32 = float32(math.NaN())

const UNDEFINED int = 0

const SUCCESS int = 1
const NO_DATA int = 2
const BACK_PRESSURED int = 3
const NO_COMPATIBLE_NODE int = 4

const ARCHITECTURE_NOT_SUPPORTED int = 0x20001
const AERON_INITIALIZATION_FAILED int = 0x20002
const INVALID_RESOURCE_ID int = 0x20003
const INVALID_APP_FOR_MODEL int = 0x20004
const MISMATCHED_FEATURE_TYPE int = 0x20005
const INVALID_FEEDBACK_LEVEL int = 0x20006
const MISSING_CREATIVES int = 0x20007
const TOO_MANY_CREATIVES int = 0x20008
const MISPLACED_CREATIVE_ID int = 0x20009
const NO_API_RESPONSE int = 0x2000A
const FORBIDDEN int = 0x2000B
const UNKNOWN_RESOURCE int = 0x2000C
const OTHER_FETCH_ERROR int = 0x2000D
const INVALID_CLUSTER_DESCRIPTION int = 0x2000E
const SCHEMA_UNAVAILABLE int = 0x2000F
const INVALID_HANDLE int = 0x20010

const APP_BID_SHADING int = 0
const APP_PROBABILITY int = 1
const APP_LINE_ITEM int = 2
const APP_UNDEFINED int = 3