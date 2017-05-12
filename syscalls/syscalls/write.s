// gcc -nostdlib -static write.s

.data

msg:
	.string "Hello Gophers!\n"
	.set len, . - msg - 1 

.text

.globl _start

_start:
	movq	$1, %rax
	movq	$1, %rdi
	movq	$msg, %rsi
	movq	$len, %rdx
	syscall

	movq	%rax, %rdi
	movq	$60, %rax
	syscall

