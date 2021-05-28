package CORE

import (
	"fmt"
)

// BIT  Test Bits in Memory with Accumulator
//
//      bits 7 and 6 of operand are transfered to bit 7 and 6 of SR (N,V);
//      the zeroflag is set to the result of operand AND accumulator.
//
//      A AND M, M7 -> N, M6 -> V        N Z C I D V
//                                      M7 + - - - M6
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      BIT oper      24    2     3
//      absolute      BIT oper      2C    3     4

func opc_BIT(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Print Opcode Debug Message
		opc_BIT_DebugMsg(bytes, mode, memAddr)

		// Memory Address bit 7 (A) -> N (Negative)
		if Debug {
			fmt.Printf("\tFlag N: %d -> ", P[7])
		}
		P[7] = Memory[memAddr] >> 7 & 0x1
		if Debug {
			fmt.Printf("%d\n", P[7])
		}

		// Memory Address bit 6 (A) -> V (oVerflow)
		if Debug {
			fmt.Printf("\tFlag V: %d -> ", P[6])
		}
		P[6] = Memory[memAddr] >> 6 & 0x1
		if Debug {
			fmt.Printf("%d\n", P[6])
		}

		flags_Z(A & Memory[memAddr])

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}
}

func opc_BIT_DebugMsg(bytes uint16, mode string, memAddr uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tBIT  Test Bits in Memory with Accumulator.\tA (%08b) AND Memory[0x%04X] (%08b) = %08b \tM7 -> N, M6 -> V\n", opc_string, mode, A, memAddr, Memory[memAddr], A&Memory[memAddr])
		fmt.Println(dbg_show_message)
	}
}
