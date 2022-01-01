package dio7000

const (
	// CC Baud Rate Setting (CC)
	/*
	  03 -> 1200
	  04 -> 2400
	  05 -> 4800
	  06 -> 9600
	  07 -> 19200
	  08 -> 38400
	  09 -> 57600
	  0A -> 115200
	*/
	CC byte = 0x0A

	// TT Type Setting
	TT byte = 0x40

	// FF
	/* Data Format Setting
	*
	*  -----------------------------------
	*  |  7 |  6 | 5 | 4 | 3 | 2 | 1 | 0 |
	*  -----------------------------------
	*  | CU | CS |  reserved |    CD     |
	*  -----------------------------------
	*
	* Counter update CU ->
	*                          0: The counter is updated when there is a falling edge in the input signal.
	*                          1: The counter is updated when there is a rising edge in the input signal.
	*
	* Checksum setting CS ->
	*                          0: Disabled
	*                          1: Enabled
	*
	* Code CD ->
	*                          I-7050: 0 (read only)
	*                          I-7052: 2 (read only)
	*                          I-7053: 3 (read only)
	*                          I-7060: 1 (read only)
	*                          For other modules, the code value can be changed by %AANNTTCCFF command
	*                          and the default code value is 0.
	*       CU CS res CD
	* 0b -> 0  0  000 000
	 */
	FF byte = 0x00
)
