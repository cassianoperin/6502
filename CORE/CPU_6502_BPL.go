package CORE

import "fmt"

// BPL  Branch on Result Plus
//
//      branch on N = 0                  N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      relative      BPL oper      10    2     2**

func opc_BPL(value int8, bytes uint16, opc_cycles byte) {

	// If Positive
	if P[7] == 0 {

		// Show current opcode cycle
		if Debug {
			fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + 1 cycle for branch + %d extra cycles for branch in different page)\n", counter_F_Cycle, opc_cycle_count, opc_cycles+opc_cycle_extra+1, opc_cycles, opc_cycle_extra)
		}

		// Just increment the Opcode cycle Counter
		if opc_cycle_count < opc_cycles+1+opc_cycle_extra {
			opc_cycle_count++

			// After spending the cycles needed, execute the opcode
		} else {
			if Debug {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: Relative]\tBPL  Branch on Result POSITIVE.\tNEGATIVE flag DISABLED, JUMP TO %04X\n", opcode, Memory[PC+1], PC+2+uint16(value))
				fmt.Println(dbg_show_message)
			}

			// PC + the number of bytes to jump on carry clear
			PC += uint16(value)

			// Increment PC
			PC += bytes

			// Reset Opcode Cycle counter
			opc_cycle_count = 1

			// Reset Opcode Extra Cycle counter
			opc_cycle_extra = 0
		}

		// If not positive
	} else {

		// Show current opcode cycle
		if Debug {
			fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
		}

		// Just increment the Opcode cycle Counter
		if opc_cycle_count < opc_cycles {
			opc_cycle_count++

			// After spending the cycles needed, execute the opcode
		} else {
			if Debug {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes]\tBPL  Branch on Result POSITIVE.\t\tNEGATIVE flag enabled, PC+=2\n", opcode, Memory[PC+1])
				fmt.Println(dbg_show_message)
			}

			// Increment PC
			PC += bytes

			// Reset Opcode Cycle counter
			opc_cycle_count = 1
		}

	}

}
