//go:build tinygo.wasm || re2_cgo

package cre2

/*
#include <stdlib.h>
#include <stdbool.h>

*/
import "C"
import "unsafe"

func New(patternPtr unsafe.Pointer, patternLen int, opts unsafe.Pointer) unsafe.Pointer {
	return C.cre2_new(patternPtr, C.int(patternLen), opts)
}

func Delete(ptr unsafe.Pointer) {
	C.cre2_delete(ptr)
}

func ErrorCode(rePtr unsafe.Pointer) int {
	return int(C.cre2_error_code(rePtr))
}

func ErrorArg(rePtr unsafe.Pointer, argPtr unsafe.Pointer) {
	C.cre2_error_arg(rePtr, argPtr)
}

func FindAndConsume(rePtr unsafe.Pointer, textPtr unsafe.Pointer, matchPtr unsafe.Pointer, nMatch int) bool {
	return C.cre2_find_and_consume_re(rePtr, textPtr, matchPtr, C.int(nMatch)) > 0
}

func GlobalReplace(rePtr unsafe.Pointer, textAndTargetPtr unsafe.Pointer, rewritePtr unsafe.Pointer) bool {
	return C.cre2_global_replace_re(rePtr, textAndTargetPtr, rewritePtr) > 0
}

func Match(rePtr unsafe.Pointer, textPtr unsafe.Pointer, textLen int, startPos int, endPos int, anchor int, matchArr unsafe.Pointer, nMatch int) bool {
	return C.cre2_match(rePtr, textPtr, C.int(textLen), C.int(startPos), C.int(endPos), C.int(anchor), matchArr, C.int(nMatch)) > 0
}

func NamedGroupsIterNew(rePtr unsafe.Pointer) unsafe.Pointer {
	return C.cre2_named_groups_iter_new(rePtr)
}

func NamedGroupsIterNext(iterPtr unsafe.Pointer, namePtr *unsafe.Pointer, indexPtr *int) bool {
	cIndex := C.int(0)
	res := C.cre2_named_groups_iter_next(iterPtr, namePtr, &cIndex)
	*indexPtr = int(cIndex)
	return bool(res)
}

func NamedGroupsIterDelete(iterPtr unsafe.Pointer) {
	C.cre2_named_groups_iter_delete(iterPtr)
}

func NumCapturingGroups(rePtr unsafe.Pointer) int {
	return int(C.cre2_num_capturing_groups(rePtr))
}

func NewOpt() unsafe.Pointer {
	return C.cre2_opt_new()
}

func DeleteOpt(opt unsafe.Pointer) {
	C.cre2_opt_delete(opt)
}

func OptSetLogErrors(opt unsafe.Pointer, flag bool) {
	C.cre2_opt_set_log_errors(opt, cFlag(flag))
}

func OptSetLongestMatch(opt unsafe.Pointer, flag bool) {
	C.cre2_opt_set_longest_match(opt, cFlag(flag))
}

func OptSetPosixSyntax(opt unsafe.Pointer, flag bool) {
	C.cre2_opt_set_posix_syntax(opt, cFlag(flag))
}

func OptSetCaseSensitive(opt unsafe.Pointer, flag bool) {
	C.cre2_opt_set_case_sensitive(opt, cFlag(flag))
}

func OptSetLatin1Encoding(opt unsafe.Pointer) {
	C.cre2_opt_set_latin1_encoding(opt)
}

func Malloc(size int) unsafe.Pointer {
	return C.malloc(C.ulong(size))
}

func Free(ptr unsafe.Pointer) {
	C.free(ptr)
}

func CopyCBytes(sPtr unsafe.Pointer, sLen int) []byte {
	return C.GoBytes(sPtr, C.int(sLen))
}

func CopyCString(sPtr unsafe.Pointer) string {
	return C.GoString((*C.char)(sPtr))
}

func CopyCStringN(sPtr unsafe.Pointer, n int) string {
	return C.GoStringN((*C.char)(sPtr), C.int(n))
}

func cFlag(flag bool) C.int {
	if flag {
		return C.int(1)
	}
	return C.int(0)
}
