# [WIP] x86-cgo 🚧

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/blacktop/x86-cgo/Go)
![GitHub all releases](https://img.shields.io/github/downloads/blacktop/x86-cgo/total)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/blacktop/x86-cgo)](https://github.com/blacktop/x86-cgo/releases/latest)
![GitHub](https://img.shields.io/github/license/blacktop/x86-cgo?color=blue)

> Golang bindings for the Binary Ninja [x86/x64 Disassembler](https://github.com/Vector35/arch-x86).

## Getting Started

```
go get github.com/blacktop/x86-cgo
```

## Status

```bash
❯ make build
 > Building locally
CGO_ENABLED=1 go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o disass.1.0.1 ./cmd/disass
# github.com/blacktop/x86-cgo/disassemble
Undefined symbols for architecture x86_64:
  "_xed_chip_features", referenced from:
      _xed_get_chip_features in libxed_macos.a(xed-chip-features.o)
      _xed_isa_set_is_valid_for_chip in libxed_macos.a(xed-isa-set.o)
  "_xed_chip_supports_avx512", referenced from:
      _xed_instruction_length_decode in libxed_macos.a(xed-ild.o)
  "_xed_convert_table", referenced from:
      _xed_print_operand_decorations in libxed_macos.a(xed-disas.o)
  "_xed_gpr_reg_class_array", referenced from:
      _xed_gpr_reg_class in libxed_macos.a(xed-reg-class.o)
  "_xed_largest_enclosing_register_array", referenced from:
      _xed_get_largest_enclosing_register in libxed_macos.a(xed-reg-class.o)
  "_xed_largest_enclosing_register_array_32", referenced from:
      _xed_get_largest_enclosing_register32 in libxed_macos.a(xed-reg-class.o)
  "_xed_pointer_name", referenced from:
      _xed_format_generic in libxed_macos.a(xed-disas.o)
      _emit_agen_and_mem in libxed_macos.a(xed-operand-values-interface.o)
  "_xed_pointer_name_suffix", referenced from:
      _xed_format_generic in libxed_macos.a(xed-disas.o)
  "_xed_reg_class_array", referenced from:
      _xed_reg_class in libxed_macos.a(xed-reg-class.o)
  "_xed_reg_width_bits", referenced from:
      _xed_decoded_inst_operand_length_bits in libxed_macos.a(xed-decoded-inst.o)
      _xed_decoded_inst_operand_element_size_bits in libxed_macos.a(xed-decoded-inst.o)
      _xed_get_register_width_bits in libxed_macos.a(xed-reg-class.o)
      _xed_get_register_width_bits64 in libxed_macos.a(xed-reg-class.o)
  "_xed_width_bits", referenced from:
      _xed_decoded_inst_compute_memory_operand_length in libxed_macos.a(xed-decoded-inst.o)
      _xed_decoded_inst_operand_length_bits in libxed_macos.a(xed-decoded-inst.o)
      _xed_decoded_inst_operand_element_size_bits in libxed_macos.a(xed-decoded-inst.o)
      _xed_operand_width_bits in libxed_macos.a(xed-inst.o)
ld: symbol(s) not found for architecture x86_64
clang: error: linker command failed with exit code 1 (use -v to see invocation)
make: *** [build] Error 2
```

## License

MIT Copyright (c) 2021 blacktop
