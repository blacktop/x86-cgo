package disassemble

/*
#cgo CFLAGS: -I${SRCDIR}

#include "xed-decoded-inst.h"
#include "xed-error-enum.h"
#include "xed-operand-accessors.h"
#include "xed-state.h"
#include "xed-operand-values-interface.h"
#include "xed-print-info.h"

/// This is the main interface to the decoder.
///  @param xedd the decoded instruction of type #xed_decoded_inst_t . Mode/state sent in via xedd; See the #xed_state_t
///  @param itext the pointer to the array of instruction text bytes
///  @param bytes  the length of the itext input array. 1 to 15 bytes, anything more is ignored.
///  @return #xed_error_enum_t indicating success (#XED_ERROR_NONE) or failure. Note failure can be due to not
///  enough bytes in the input array.
///
/// The maximum instruction is 15B and XED will tell you how long the
/// actual instruction is via an API function call
/// xed_decoded_inst_get_length().  However, it is not always safe or
/// advisable for XED to read 15 bytes if the decode location is at the
/// boundary of some sort of protection limit. For example, if one is
/// decoding near the end of a page and the XED user does not want to cause
/// extra page faults, one might send in the number of bytes that would
/// stop at the page boundary. In this case, XED might not be able to
/// decode the instruction and would return an error. The XED user would
/// then have to decide if it was safe to touch the next page and try again
/// to decode with more bytes.  Also sometimes the user process does not
/// have read access to the next page and this allows the user to prevent
/// XED from causing process termination by limiting the memory range that
/// XED will access.
///
/// @ingroup DEC
XED_DLL_EXPORT xed_error_enum_t
xed_decode(xed_decoded_inst_t* xedd,
           const xed_uint8_t* itext,
           const unsigned int bytes);

/// Disassemble the decoded instruction using the specified syntax.
/// The output buffer must be at least 25 bytes long. Returns true if
/// disassembly proceeded without errors.
/// @param syntax a #xed_syntax_enum_t the specifies the disassembly format
/// @param xedd a #xed_decoded_inst_t for a decoded instruction
/// @param out_buffer a buffer to write the disassembly in to.
/// @param buffer_len maximum length of the disassembly buffer
/// @param runtime_instruction_address the address of the instruction being disassembled. If zero, the offset is printed for relative branches. If nonzero, XED attempts to print the target address for relative branches.
/// @param context A void* used only for the call back routine for symbolic disassembly if one is provided. Can be zero.
/// @param symbolic_callback A function pointer for obtaining symbolic disassembly. Can be zero.
/// @return Returns 0 if the disassembly fails, 1 otherwise.
///@ingroup PRINT
XED_DLL_EXPORT xed_bool_t
xed_format_context(xed_syntax_enum_t syntax,
                   const xed_decoded_inst_t* xedd,
                   char* out_buffer,
                   int  buffer_len,
                   xed_uint64_t runtime_instruction_address,
                   void* context,
                   xed_disassembly_callback_fn_t symbolic_callback);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type xed_syntax_enum_t uint32

const (
	XED_SYNTAX_INVALID xed_syntax_enum_t = iota
	XED_SYNTAX_XED                       ///< XED disassembly syntax
	XED_SYNTAX_ATT                       ///< ATT SYSV disassembly syntax
	XED_SYNTAX_INTEL                     ///< Intel disassembly syntax
	XED_SYNTAX_LAST
)

// GetOpCodeByteString returns the opcodes as a string of hex bytes
func GetOpCodeByteString(opcode []byte) string {
	return fmt.Sprintf("% x", opcode) // TODO: check if needs to be big endian etc
}

// Disassemble disassembles an instruction
func Disassemble(addr uint64, itext []byte, results *[1024]byte) (string, error) {

	var xedd C.xed_decoded_inst_t

	C.xed_decode(
		&xedd,                                   // xed_decoded_inst_t* xedd
		(*C.xed_uint8_t)(unsafe.Pointer(itext)), // const xed_uint8_t* itext
		15,                                      // const unsigned int bytes
	)

	C.xed_format_context(
		C.xed_syntax_enum_t(XED_SYNTAX_INTEL), // xed_syntax_enum_t syntax
		&xedd,                                 // const xed_decoded_inst_t* xedd
		(*C.char)(unsafe.Pointer(results)),    // char* out_buffer
		C.int(1024),                           // int  buffer_len
		C.xed_uint64_t(addr),                  // xed_uint64_t runtime_instruction_address
		0,                                     // void* context
		0,                                     // xed_disassembly_callback_fn_t symbolic_callbac
	)

	return C.GoString((*C.char)(unsafe.Pointer(results))), nil
}
