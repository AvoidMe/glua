// Code generated by "stringer -type=Opcode"; DO NOT EDIT.

package compiler

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OP_MOVE-0]
	_ = x[OP_LOADI-1]
	_ = x[OP_LOADF-2]
	_ = x[OP_LOADK-3]
	_ = x[OP_LOADKX-4]
	_ = x[OP_LOADFALSE-5]
	_ = x[OP_LFALSESKIP-6]
	_ = x[OP_LOADTRUE-7]
	_ = x[OP_LOADNIL-8]
	_ = x[OP_GETUPVAL-9]
	_ = x[OP_SETUPVAL-10]
	_ = x[OP_GETTABUP-11]
	_ = x[OP_GETTABLE-12]
	_ = x[OP_GETI-13]
	_ = x[OP_GETFIELD-14]
	_ = x[OP_SETTABUP-15]
	_ = x[OP_SETTABLE-16]
	_ = x[OP_SETI-17]
	_ = x[OP_SETFIELD-18]
	_ = x[OP_NEWTABLE-19]
	_ = x[OP_SELF-20]
	_ = x[OP_ADDI-21]
	_ = x[OP_ADDK-22]
	_ = x[OP_SUBK-23]
	_ = x[OP_MULK-24]
	_ = x[OP_MODK-25]
	_ = x[OP_POWK-26]
	_ = x[OP_DIVK-27]
	_ = x[OP_IDIVK-28]
	_ = x[OP_BANDK-29]
	_ = x[OP_BORK-30]
	_ = x[OP_BXORK-31]
	_ = x[OP_SHRI-32]
	_ = x[OP_SHLI-33]
	_ = x[OP_ADD-34]
	_ = x[OP_SUB-35]
	_ = x[OP_MUL-36]
	_ = x[OP_MOD-37]
	_ = x[OP_POW-38]
	_ = x[OP_DIV-39]
	_ = x[OP_IDIV-40]
	_ = x[OP_BAND-41]
	_ = x[OP_BOR-42]
	_ = x[OP_BXOR-43]
	_ = x[OP_SHL-44]
	_ = x[OP_SHR-45]
	_ = x[OP_MMBIN-46]
	_ = x[OP_MMBINI-47]
	_ = x[OP_MMBINK-48]
	_ = x[OP_UNM-49]
	_ = x[OP_BNOT-50]
	_ = x[OP_NOT-51]
	_ = x[OP_LEN-52]
	_ = x[OP_CONCAT-53]
	_ = x[OP_CLOSE-54]
	_ = x[OP_TBC-55]
	_ = x[OP_JMP-56]
	_ = x[OP_EQ-57]
	_ = x[OP_LT-58]
	_ = x[OP_LE-59]
	_ = x[OP_EQK-60]
	_ = x[OP_EQI-61]
	_ = x[OP_LTI-62]
	_ = x[OP_LEI-63]
	_ = x[OP_GTI-64]
	_ = x[OP_GEI-65]
	_ = x[OP_TEST-66]
	_ = x[OP_TESTSET-67]
	_ = x[OP_CALL-68]
	_ = x[OP_TAILCALL-69]
	_ = x[OP_RETURN-70]
	_ = x[OP_RETURN0-71]
	_ = x[OP_RETURN1-72]
	_ = x[OP_FORLOOP-73]
	_ = x[OP_FORPREP-74]
	_ = x[OP_TFORPREP-75]
	_ = x[OP_TFORCALL-76]
	_ = x[OP_TFORLOOP-77]
	_ = x[OP_SETLIST-78]
	_ = x[OP_CLOSURE-79]
	_ = x[OP_VARARG-80]
	_ = x[OP_VARARGPREP-81]
	_ = x[OP_EXTRAARG-82]
}

const _Opcode_name = "OP_MOVEOP_LOADIOP_LOADFOP_LOADKOP_LOADKXOP_LOADFALSEOP_LFALSESKIPOP_LOADTRUEOP_LOADNILOP_GETUPVALOP_SETUPVALOP_GETTABUPOP_GETTABLEOP_GETIOP_GETFIELDOP_SETTABUPOP_SETTABLEOP_SETIOP_SETFIELDOP_NEWTABLEOP_SELFOP_ADDIOP_ADDKOP_SUBKOP_MULKOP_MODKOP_POWKOP_DIVKOP_IDIVKOP_BANDKOP_BORKOP_BXORKOP_SHRIOP_SHLIOP_ADDOP_SUBOP_MULOP_MODOP_POWOP_DIVOP_IDIVOP_BANDOP_BOROP_BXOROP_SHLOP_SHROP_MMBINOP_MMBINIOP_MMBINKOP_UNMOP_BNOTOP_NOTOP_LENOP_CONCATOP_CLOSEOP_TBCOP_JMPOP_EQOP_LTOP_LEOP_EQKOP_EQIOP_LTIOP_LEIOP_GTIOP_GEIOP_TESTOP_TESTSETOP_CALLOP_TAILCALLOP_RETURNOP_RETURN0OP_RETURN1OP_FORLOOPOP_FORPREPOP_TFORPREPOP_TFORCALLOP_TFORLOOPOP_SETLISTOP_CLOSUREOP_VARARGOP_VARARGPREPOP_EXTRAARG"

var _Opcode_index = [...]uint16{0, 7, 15, 23, 31, 40, 52, 65, 76, 86, 97, 108, 119, 130, 137, 148, 159, 170, 177, 188, 199, 206, 213, 220, 227, 234, 241, 248, 255, 263, 271, 278, 286, 293, 300, 306, 312, 318, 324, 330, 336, 343, 350, 356, 363, 369, 375, 383, 392, 401, 407, 414, 420, 426, 435, 443, 449, 455, 460, 465, 470, 476, 482, 488, 494, 500, 506, 513, 523, 530, 541, 550, 560, 570, 580, 590, 601, 612, 623, 633, 643, 652, 665, 676}

func (i Opcode) String() string {
	if i < 0 || i >= Opcode(len(_Opcode_index)-1) {
		return "Opcode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Opcode_name[_Opcode_index[i]:_Opcode_index[i+1]]
}
