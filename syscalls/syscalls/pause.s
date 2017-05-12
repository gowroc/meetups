// gcc -nostdlib pause.s

.text

.globl _start

_start:
	movq	$34, %rax
	syscall

	movq	$60, %rax
	movq	$0, %rdi
	syscall

