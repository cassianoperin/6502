package CORE

import (
	"fmt"
)

// STA  Store Accumulator in Memory
//
//      A -> M                           N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      STA oper      85    2     3
//      zeropage,X    STA oper,X    95    2     4
//      absolute      STA oper      8D    3     4
//      absolute,X    STA oper,X    9D    3     5
//      absolute,Y    STA oper,Y    99    3     5
//      (indirect,X)  STA (oper,X)  81    2     6
//      (indirect),Y  STA (oper),Y  91    2     6

func opc_STA(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Update Memory[memAddr] with value of A and notify TIA about the update
		memUpdate(memAddr, A)

		// Print Opcode Debug Message
		opc_STA_DebugMsg(bytes, mode, memAddr)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}

func opc_STA_DebugMsg(bytes uint16, mode string, memAddr uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tSTA  Store Accumulator in Memory.\tMemory[0x%02X] = A (0x%02X)\n", opc_string, mode, memAddr, Memory[memAddr])
		fmt.Println(dbg_show_message)
	}
}
