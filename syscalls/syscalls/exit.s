// gcc -nostdlib exit.s

.text

.globl _start 

_start:
	movq	$1, %rax
	movq	$42, %rbx
	int		$0x80
