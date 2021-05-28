package CORE

import (
	"fmt"
)

// ---------------------------- Library Function ---------------------------- //

// Memory Page Boundary cross detection
func MemPageBoundary(original_addr, new_addr uint16) byte {

	var extra_cycle byte = 0

	// Page Boundary Cross detected
	if original_addr>>8 != new_addr>>8 { // Get the High byte only to compare

		extra_cycle = 1

		if Debug {
			fmt.Printf("\tMemory Page Boundary Cross detected! Add 1 cycle.\tPC High byte: %02X\tBranch High byte: %02X\n", original_addr>>8, new_addr>>8)
		}

		// NO Page Boundary Cross detected
	} else {

		extra_cycle = 0

		if Debug {
			fmt.Printf("\tNo Memory Page Boundary Cross detected.\tPC High byte: %02X\tBranch High byte: %02X\n", original_addr>>8, new_addr>>8)
		}
	}

	return extra_cycle
}

// Decode Two's Complement
func DecodeTwoComplement(num byte) int8 {

	var sum int8 = 0

	for i := 0; i < 8; i++ {
		// Sum each bit and sum the value of the bit power of i (<<i)
		sum += (int8(num) >> i & 0x01) << i
	}

	return sum
}

// Decode opcode for debug messages
func debug_decode_opc(bytes uint16) string {

	var opc_string string

	// Decode opcode and operators
	for i := 0; i < int(bytes); i++ {
		if i == 1 {
			opc_string += fmt.Sprintf(" %02X", Memory[PC+uint16(i)])
		} else {
			opc_string += fmt.Sprintf("%02X", Memory[PC+uint16(i)])
		}
	}

	// Insert number of bytes into the string
	if bytes == 1 {
		opc_string += " [1 byte]"
	} else if bytes == 2 {
		opc_string += " [2 bytes]"
	} else {
		opc_string += " [3 bytes]"
	}

	return opc_string
}

// // BCD - Binary Coded Decimal
// func BCD(number byte) byte {

// 	var tmp_hundreds, tmp_tens, tmp_ones, bcd byte

// 	// Split the Decimal Value
// 	tmp_hundreds = number / 100    // Hundreds
// 	tmp_tens = (number / 10) % 10  // Tens
// 	tmp_ones = (number % 100) % 10 // Ones

// 	fmt.Printf("H: %d\tT: %d\tO: %d\n", tmp_hundreds, tmp_tens, tmp_ones)

// 	// Combine in one decimal number
// 	bcd = (tmp_hundreds * 100) + (tmp_tens * 10) + tmp_ones

// 	return bcd
// }

// Memory Bus - Used by INC, STA, STY and STX to update memory and sinalize TIA about the actions
func memUpdate(memAddr uint16, value byte) {

	Memory[memAddr] = value

}
