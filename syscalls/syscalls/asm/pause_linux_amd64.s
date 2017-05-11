#include "textflag.h"

TEXT	·Pause(SB),NOSPLIT,$0
		MOVQ	$34, AX
		SYSCALL
