package CORE

import "fmt"

// STX  Store Index X in Memory
//
//      X -> M                           N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      STX oper      86    2     3
//      zeropage,Y    STX oper,Y    96    2     4
//      absolute      STX oper      8E    3     4

func opc_STX(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Write data to Memory (adress in Memory Bus) and update the value in Data BUS
		var memData byte = dataBUS_Write(memAddr, X)

		// Print Opcode Debug Message
		opc_STX_DebugMsg(bytes, mode, memAddr, memData)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_STX_DebugMsg(bytes uint16, mode string, memAddr uint16, memData byte) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tSTX  Store Index X in Memory.\tMemory[0x%02X] = X (%d)\n", opc_string, mode, memAddr, memData)
		fmt.Println(dbg_show_message)
	}
}
