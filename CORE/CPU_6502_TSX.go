package CORE

import "fmt"

// TSX  Transfer Stack Pointer to Index X
//
//      SP -> X                          N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       TSX           BA    1     2

func opc_TSX(bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		X = SP

		// Print Opcode Debug Message
		opc_TSX_DebugMsg(bytes)

		flags_Z(X)
		flags_N(X)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}

func opc_TSX_DebugMsg(bytes uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Implied]\tTSX  Transfer Stack Pointer to Index X.\tX = SP (%d)\n", opc_string, SP)
		fmt.Println(dbg_show_message)
	}
}
