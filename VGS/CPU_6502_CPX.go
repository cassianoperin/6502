package VGS

import "fmt"

// CPX  Compare Memory and Index X
//
//      X - M                            N Z C I D V
//                                       + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     CPX #oper     E0    2     2
//      zeropage      CPX oper    	E4    2	    3
//      absolute      CPX oper      EC    3     4

func opc_CPX(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		tmp := X - Memory[memAddr]

		if Debug {

			if tmp == 0 {

				if bytes == 2 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCPX  Compare Memory and Index X.\tX(%d) - Memory[%02X](%d) = (%d) EQUAL\n", opcode, Memory[PC+1], mode, X, PC+1, Memory[memAddr], tmp)
					fmt.Println(dbg_show_message)
				} else if bytes == 3 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tCPX  Compare Memory and Index X.\tX(%d) - Memory[%02X](%d) = (%d) EQUAL\n", opcode, Memory[PC+2], Memory[PC+1], mode, X, PC+1, Memory[memAddr], tmp)
					fmt.Println(dbg_show_message)
				}

			} else {

				if bytes == 2 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCPX  Compare Memory and Index X.\tX(%d) - Memory[%02X](%d) = (%d) NOT EQUAL\n", opcode, Memory[PC+1], mode, X, PC+1, Memory[memAddr], tmp)
					fmt.Println(dbg_show_message)
				} else if bytes == 3 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tCPX  Compare Memory and Index X.\tX(%d) - Memory[%02X](%d) = (%d) NOT EQUAL\n", opcode, Memory[PC+2], Memory[PC+1], mode, X, PC+1, Memory[memAddr], tmp)
					fmt.Println(dbg_show_message)
				}

			}

		}

		// Set if X = M
		flags_Z(tmp)
		// Set if bit 7 of the result is set
		flags_N(tmp)
		// Set if X >= M
		flags_C_CPX_CPY_CMP(X, Memory[memAddr])

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}