
# file

### Makefile
```
CC = gcc
CFLAGS = -W -Wall
TARGET = hello
OBJECTS = hello.o

all : $(TARGET)

$(TARGET) : $(OBJECTS)
	$(CC) $(CFLAGS) -o $@ $^

clean : 
	rm  *.o $(TARGET)
```
### hello compile
```
#include<stdio.h>
int main (int argc, char *agrv[]){
    printf("hello, world\n");
    return 0;
}
```

```
$ ls -l
-rwxr-xr-x 1 good dev 8304  5월 10 22:32 hello
-rw-r--r-- 1 good dev   98  5월 10 22:18 hello.c
-rw-r--r-- 1 good dev 1560  5월 10 22:32 hello.o
```
### ELF type
```
$ file hello
hello: ELF 64-bit LSB shared object, x86-64, version 1 (SYSV), dynamically linked, interpreter /lib64/ld-linux-x86-64.so.2, for GNU/Linux 3.2.0, BuildID[sha1]=4c2f1013baa212fae367413fac48e0eda97e6848, not stripped
```

### dynamic link library
```
#include<stdio.h>
int main (int argc, char *agrv[]){
    printf("hello, world\n");
    return 0;
}
```
### objdump 
```
$ objdump -d hello

hello:     file format elf64-x86-64


Disassembly of section .init:

00000000000004e8 <_init>:
 4e8:   48 83 ec 08             sub    $0x8,%rsp
 4ec:   48 8b 05 f5 0a 20 00    mov    0x200af5(%rip),%rax        # 200fe8 <__gmon_start__>
 4f3:   48 85 c0                test   %rax,%rax
 4f6:   74 02                   je     4fa <_init+0x12>
 4f8:   ff d0                   callq  *%rax
 4fa:   48 83 c4 08             add    $0x8,%rsp
 4fe:   c3                      retq   

Disassembly of section .plt:

0000000000000500 <.plt>:
 500:   ff 35 ba 0a 20 00       pushq  0x200aba(%rip)        # 200fc0 <_GLOBAL_OFFSET_TABLE_+0x8>
 506:   ff 25 bc 0a 20 00       jmpq   *0x200abc(%rip)        # 200fc8 <_GLOBAL_OFFSET_TABLE_+0x10>
 50c:   0f 1f 40 00             nopl   0x0(%rax)

0000000000000510 <puts@plt>:
 510:   ff 25 ba 0a 20 00       jmpq   *0x200aba(%rip)        # 200fd0 <puts@GLIBC_2.2.5>
 516:   68 00 00 00 00          pushq  $0x0
 51b:   e9 e0 ff ff ff          jmpq   500 <.plt>

Disassembly of section .plt.got:

0000000000000520 <__cxa_finalize@plt>:
 520:   ff 25 d2 0a 20 00       jmpq   *0x200ad2(%rip)        # 200ff8 <__cxa_finalize@GLIBC_2.2.5>
 526:   66 90                   xchg   %ax,%ax

Disassembly of section .text:

0000000000000530 <_start>:
 530:   31 ed                   xor    %ebp,%ebp
 532:   49 89 d1                mov    %rdx,%r9
 535:   5e                      pop    %rsi
 536:   48 89 e2                mov    %rsp,%rdx
 539:   48 83 e4 f0             and    $0xfffffffffffffff0,%rsp
 53d:   50                      push   %rax
 53e:   54                      push   %rsp
 53f:   4c 8d 05 8a 01 00 00    lea    0x18a(%rip),%r8        # 6d0 <__libc_csu_fini>
 546:   48 8d 0d 13 01 00 00    lea    0x113(%rip),%rcx        # 660 <__libc_csu_init>
 54d:   48 8d 3d e6 00 00 00    lea    0xe6(%rip),%rdi        # 63a <main>
 554:   ff 15 86 0a 20 00       callq  *0x200a86(%rip)        # 200fe0 <__libc_start_main@GLIBC_2.2.5>
 55a:   f4                      hlt    
 55b:   0f 1f 44 00 00          nopl   0x0(%rax,%rax,1)

0000000000000560 <deregister_tm_clones>:
 560:   48 8d 3d a9 0a 20 00    lea    0x200aa9(%rip),%rdi        # 201010 <__TMC_END__>
 567:   55                      push   %rbp
 568:   48 8d 05 a1 0a 20 00    lea    0x200aa1(%rip),%rax        # 201010 <__TMC_END__>
 56f:   48 39 f8                cmp    %rdi,%rax
 572:   48 89 e5                mov    %rsp,%rbp
 575:   74 19                   je     590 <deregister_tm_clones+0x30>
 577:   48 8b 05 5a 0a 20 00    mov    0x200a5a(%rip),%rax        # 200fd8 <_ITM_deregisterTMCloneTable>
 57e:   48 85 c0                test   %rax,%rax
 581:   74 0d                   je     590 <deregister_tm_clones+0x30>
 583:   5d                      pop    %rbp
 584:   ff e0                   jmpq   *%rax
 586:   66 2e 0f 1f 84 00 00    nopw   %cs:0x0(%rax,%rax,1)
 58d:   00 00 00 
 590:   5d                      pop    %rbp
 591:   c3                      retq   
 592:   0f 1f 40 00             nopl   0x0(%rax)
 596:   66 2e 0f 1f 84 00 00    nopw   %cs:0x0(%rax,%rax,1)
 59d:   00 00 00 

00000000000005a0 <register_tm_clones>:
 5a0:   48 8d 3d 69 0a 20 00    lea    0x200a69(%rip),%rdi        # 201010 <__TMC_END__>
 5a7:   48 8d 35 62 0a 20 00    lea    0x200a62(%rip),%rsi        # 201010 <__TMC_END__>
 5ae:   55                      push   %rbp
 5af:   48 29 fe                sub    %rdi,%rsi
 5b2:   48 89 e5                mov    %rsp,%rbp
 5b5:   48 c1 fe 03             sar    $0x3,%rsi
 5b9:   48 89 f0                mov    %rsi,%rax
 5bc:   48 c1 e8 3f             shr    $0x3f,%rax
 5c0:   48 01 c6                add    %rax,%rsi
 5c3:   48 d1 fe                sar    %rsi
 5c6:   74 18                   je     5e0 <register_tm_clones+0x40>
 5c8:   48 8b 05 21 0a 20 00    mov    0x200a21(%rip),%rax        # 200ff0 <_ITM_registerTMCloneTable>
 5cf:   48 85 c0                test   %rax,%rax
 5d2:   74 0c                   je     5e0 <register_tm_clones+0x40>
 5d4:   5d                      pop    %rbp
 5d5:   ff e0                   jmpq   *%rax
 5d7:   66 0f 1f 84 00 00 00    nopw   0x0(%rax,%rax,1)
 5de:   00 00 
 5e0:   5d                      pop    %rbp
 5e1:   c3                      retq   
 5e2:   0f 1f 40 00             nopl   0x0(%rax)
 5e6:   66 2e 0f 1f 84 00 00    nopw   %cs:0x0(%rax,%rax,1)
 5ed:   00 00 00 

00000000000005f0 <__do_global_dtors_aux>:
 5f0:   80 3d 19 0a 20 00 00    cmpb   $0x0,0x200a19(%rip)        # 201010 <__TMC_END__>
 5f7:   75 2f                   jne    628 <__do_global_dtors_aux+0x38>
 5f9:   48 83 3d f7 09 20 00    cmpq   $0x0,0x2009f7(%rip)        # 200ff8 <__cxa_finalize@GLIBC_2.2.5>
 600:   00 
 601:   55                      push   %rbp
 602:   48 89 e5                mov    %rsp,%rbp
 605:   74 0c                   je     613 <__do_global_dtors_aux+0x23>
 607:   48 8b 3d fa 09 20 00    mov    0x2009fa(%rip),%rdi        # 201008 <__dso_handle>
 60e:   e8 0d ff ff ff          callq  520 <__cxa_finalize@plt>
 613:   e8 48 ff ff ff          callq  560 <deregister_tm_clones>
 618:   c6 05 f1 09 20 00 01    movb   $0x1,0x2009f1(%rip)        # 201010 <__TMC_END__>
 61f:   5d                      pop    %rbp
 620:   c3                      retq   
 621:   0f 1f 80 00 00 00 00    nopl   0x0(%rax)
 628:   f3 c3                   repz retq 
 62a:   66 0f 1f 44 00 00       nopw   0x0(%rax,%rax,1)

0000000000000630 <frame_dummy>:
 630:   55                      push   %rbp
 631:   48 89 e5                mov    %rsp,%rbp
 634:   5d                      pop    %rbp
 635:   e9 66 ff ff ff          jmpq   5a0 <register_tm_clones>

000000000000063a <main>:
 63a:   55                      push   %rbp
 63b:   48 89 e5                mov    %rsp,%rbp
 63e:   48 83 ec 10             sub    $0x10,%rsp
 642:   89 7d fc                mov    %edi,-0x4(%rbp)
 645:   48 89 75 f0             mov    %rsi,-0x10(%rbp)
 649:   48 8d 3d 94 00 00 00    lea    0x94(%rip),%rdi        # 6e4 <_IO_stdin_used+0x4>
 650:   e8 bb fe ff ff          callq  510 <puts@plt>
 655:   b8 00 00 00 00          mov    $0x0,%eax
 65a:   c9                      leaveq 
 65b:   c3                      retq   
 65c:   0f 1f 40 00             nopl   0x0(%rax)

0000000000000660 <__libc_csu_init>:
 660:   41 57                   push   %r15
 662:   41 56                   push   %r14
 664:   49 89 d7                mov    %rdx,%r15
 667:   41 55                   push   %r13
 669:   41 54                   push   %r12
 66b:   4c 8d 25 46 07 20 00    lea    0x200746(%rip),%r12        # 200db8 <__frame_dummy_init_array_entry>
 672:   55                      push   %rbp
 673:   48 8d 2d 46 07 20 00    lea    0x200746(%rip),%rbp        # 200dc0 <__init_array_end>
 67a:   53                      push   %rbx
 67b:   41 89 fd                mov    %edi,%r13d
 67e:   49 89 f6                mov    %rsi,%r14
 681:   4c 29 e5                sub    %r12,%rbp
 684:   48 83 ec 08             sub    $0x8,%rsp
 688:   48 c1 fd 03             sar    $0x3,%rbp
 68c:   e8 57 fe ff ff          callq  4e8 <_init>
 691:   48 85 ed                test   %rbp,%rbp
 694:   74 20                   je     6b6 <__libc_csu_init+0x56>
 696:   31 db                   xor    %ebx,%ebx
 698:   0f 1f 84 00 00 00 00    nopl   0x0(%rax,%rax,1)
 69f:   00 
 6a0:   4c 89 fa                mov    %r15,%rdx
 6a3:   4c 89 f6                mov    %r14,%rsi
 6a6:   44 89 ef                mov    %r13d,%edi
 6a9:   41 ff 14 dc             callq  *(%r12,%rbx,8)
 6ad:   48 83 c3 01             add    $0x1,%rbx
 6b1:   48 39 dd                cmp    %rbx,%rbp
 6b4:   75 ea                   jne    6a0 <__libc_csu_init+0x40>
 6b6:   48 83 c4 08             add    $0x8,%rsp
 6ba:   5b                      pop    %rbx
 6bb:   5d                      pop    %rbp
 6bc:   41 5c                   pop    %r12
 6be:   41 5d                   pop    %r13
 6c0:   41 5e                   pop    %r14
 6c2:   41 5f                   pop    %r15
 6c4:   c3                      retq   
 6c5:   90                      nop
 6c6:   66 2e 0f 1f 84 00 00    nopw   %cs:0x0(%rax,%rax,1)
 6cd:   00 00 00 

00000000000006d0 <__libc_csu_fini>:
 6d0:   f3 c3                   repz retq 

Disassembly of section .fini:

00000000000006d4 <_fini>:
 6d4:   48 83 ec 08             sub    $0x8,%rsp
 6d8:   48 83 c4 08             add    $0x8,%rsp
 6dc:   c3                      retq   
 ```
### objdum -s -j .rodata 
o6e4 .. 
```
$ objdump -s -j .rodata hello

hello:     file format elf64-x86-64

Contents of section .rodata:
 06e0 01000200 68656c6c 6f2c2077 6f726c64  ....hello, world
 06f0 00
 ```

 ### hello.c

unistd.h  posix api access 
 ```
 #include <unistd.h>
int main (int argc, char *agrv[]){
    write(1,"hello, world\n",14);
    return 0;
}
 ```



### objdump from syscall
```
$ objdump -d hello

hello:     file format elf64-x86-64


Disassembly of section .init:

00000000000004f0 <_init>:
 4f0:   48 83 ec 08             sub    $0x8,%rsp
 4f4:   48 8b 05 ed 0a 20 00    mov    0x200aed(%rip),%rax        # 200fe8 <__gmon_start__>
 4fb:   48 85 c0                test   %rax,%rax
 4fe:   74 02                   je     502 <_init+0x12>
 500:   ff d0                   callq  *%rax
 502:   48 83 c4 08             add    $0x8,%rsp
 506:   c3                      retq   

Disassembly of section .plt:

0000000000000510 <.plt>:
 510:   ff 35 aa 0a 20 00       pushq  0x200aaa(%rip)        # 200fc0 <_GLOBAL_OFFSET_TABLE_+0x8>
 516:   ff 25 ac 0a 20 00       jmpq   *0x200aac(%rip)        # 200fc8 <_GLOBAL_OFFSET_TABLE_+0x10>
 51c:   0f 1f 40 00             nopl   0x0(%rax)

0000000000000520 <syscall@plt>:
 520:   ff 25 aa 0a 20 00       jmpq   *0x200aaa(%rip)        # 200fd0 <syscall@GLIBC_2.2.5>
 526:   68 00 00 00 00          pushq  $0x0
 52b:   e9 e0 ff ff ff          jmpq   510 <.plt>

Disassembly of section .plt.got:

0000000000000530 <__cxa_finalize@plt>:
 530:   ff 25 c2 0a 20 00       jmpq   *0x200ac2(%rip)        # 200ff8 <__cxa_finalize@GLIBC_2.2.5>
 536:   66 90                   xchg   %ax,%ax

Disassembly of section .text:

0000000000000540 <_start>:
 540:   31 ed                   xor    %ebp,%ebp
 542:   49 89 d1                mov    %rdx,%r9
 545:   5e                      pop    %rsi
 546:   48 89 e2                mov    %rsp,%rdx
 549:   48 83 e4 f0             and    $0xfffffffffffffff0,%rsp
 54d:   50                      push   %rax
 54e:   54                      push   %rsp
 54f:   4c 8d 05 9a 01 00 00    lea    0x19a(%rip),%r8        # 6f0 <__libc_csu_fini>
 556:   48 8d 0d 23 01 00 00    lea    0x123(%rip),%rcx        # 680 <__libc_csu_init>
 55d:   48 8d 3d e6 00 00 00    lea    0xe6(%rip),%rdi        # 64a <main>
 564:   ff 15 76 0a 20 00       callq  *0x200a76(%rip)        # 200fe0 <__libc_start_main@GLIBC_2.2.5>
 56a:   f4                      hlt    
 56b:   0f 1f 44 00 00          nopl   0x0(%rax,%rax,1)

0000000000000570 <deregister_tm_clones>:
 570:   48 8d 3d 99 0a 20 00    lea    0x200a99(%rip),%rdi        # 201010 <__TMC_END__>
 577:   55                      push   %rbp
 578:   48 8d 05 91 0a 20 00    lea    0x200a91(%rip),%rax        # 201010 <__TMC_END__>
 57f:   48 39 f8                cmp    %rdi,%rax
 582:   48 89 e5                mov    %rsp,%rbp
 585:   74 19                   je     5a0 <deregister_tm_clones+0x30>
 587:   48 8b 05 4a 0a 20 00    mov    0x200a4a(%rip),%rax        # 200fd8 <_ITM_deregisterTMCloneTable>
 58e:   48 85 c0                test   %rax,%rax
 591:   74 0d                   je     5a0 <deregister_tm_clones+0x30>
 593:   5d                      pop    %rbp
 594:   ff e0                   jmpq   *%rax
 596:   66 2e 0f 1f 84 00 00    nopw   %cs:0x0(%rax,%rax,1)
 59d:   00 00 00 
 5a0:   5d                      pop    %rbp
 5a1:   c3                      retq   
 5a2:   0f 1f 40 00             nopl   0x0(%rax)
 5a6:   66 2e 0f 1f 84 00 00    nopw   %cs:0x0(%rax,%rax,1)
 5ad:   00 00 00 

00000000000005b0 <register_tm_clones>:
 5b0:   48 8d 3d 59 0a 20 00    lea    0x200a59(%rip),%rdi        # 201010 <__TMC_END__>
 5b7:   48 8d 35 52 0a 20 00    lea    0x200a52(%rip),%rsi        # 201010 <__TMC_END__>
 5be:   55                      push   %rbp
 5bf:   48 29 fe                sub    %rdi,%rsi
 5c2:   48 89 e5                mov    %rsp,%rbp
 5c5:   48 c1 fe 03             sar    $0x3,%rsi
 5c9:   48 89 f0                mov    %rsi,%rax
 5cc:   48 c1 e8 3f             shr    $0x3f,%rax
 5d0:   48 01 c6                add    %rax,%rsi
 5d3:   48 d1 fe                sar    %rsi
 5d6:   74 18                   je     5f0 <register_tm_clones+0x40>
 5d8:   48 8b 05 11 0a 20 00    mov    0x200a11(%rip),%rax        # 200ff0 <_ITM_registerTMCloneTable>
 5df:   48 85 c0                test   %rax,%rax
 5e2:   74 0c                   je     5f0 <register_tm_clones+0x40>
 5e4:   5d                      pop    %rbp
 5e5:   ff e0                   jmpq   *%rax
 5e7:   66 0f 1f 84 00 00 00    nopw   0x0(%rax,%rax,1)
 5ee:   00 00 
 5f0:   5d                      pop    %rbp
 5f1:   c3                      retq   
 5f2:   0f 1f 40 00             nopl   0x0(%rax)
 5f6:   66 2e 0f 1f 84 00 00    nopw   %cs:0x0(%rax,%rax,1)
 5fd:   00 00 00 

0000000000000600 <__do_global_dtors_aux>:
 600:   80 3d 09 0a 20 00 00    cmpb   $0x0,0x200a09(%rip)        # 201010 <__TMC_END__>
 607:   75 2f                   jne    638 <__do_global_dtors_aux+0x38>
 609:   48 83 3d e7 09 20 00    cmpq   $0x0,0x2009e7(%rip)        # 200ff8 <__cxa_finalize@GLIBC_2.2.5>
 610:   00 
 611:   55                      push   %rbp
 612:   48 89 e5                mov    %rsp,%rbp
 615:   74 0c                   je     623 <__do_global_dtors_aux+0x23>
 617:   48 8b 3d ea 09 20 00    mov    0x2009ea(%rip),%rdi        # 201008 <__dso_handle>
 61e:   e8 0d ff ff ff          callq  530 <__cxa_finalize@plt>
 623:   e8 48 ff ff ff          callq  570 <deregister_tm_clones>
 628:   c6 05 e1 09 20 00 01    movb   $0x1,0x2009e1(%rip)        # 201010 <__TMC_END__>
 62f:   5d                      pop    %rbp
 630:   c3                      retq   
 631:   0f 1f 80 00 00 00 00    nopl   0x0(%rax)
 638:   f3 c3                   repz retq 
 63a:   66 0f 1f 44 00 00       nopw   0x0(%rax,%rax,1)

0000000000000640 <frame_dummy>:
 640:   55                      push   %rbp
 641:   48 89 e5                mov    %rsp,%rbp
 644:   5d                      pop    %rbp
 645:   e9 66 ff ff ff          jmpq   5b0 <register_tm_clones>

000000000000064a <main>:
 64a:   55                      push   %rbp
 64b:   48 89 e5                mov    %rsp,%rbp
 64e:   48 83 ec 10             sub    $0x10,%rsp
 652:   89 7d fc                mov    %edi,-0x4(%rbp)
 655:   48 89 75 f0             mov    %rsi,-0x10(%rbp)
 659:   b9 0e 00 00 00          mov    $0xe,%ecx
 65e:   48 8d 15 9f 00 00 00    lea    0x9f(%rip),%rdx        # 704 <_IO_stdin_used+0x4>
 665:   be 01 00 00 00          mov    $0x1,%esi
 66a:   bf 01 00 00 00          mov    $0x1,%edi
 66f:   b8 00 00 00 00          mov    $0x0,%eax
 674:   e8 a7 fe ff ff          callq  520 <syscall@plt>
 679:   b8 00 00 00 00          mov    $0x0,%eax
 67e:   c9                      leaveq 
 67f:   c3                      retq   


00000000000006f0 <__libc_csu_fini>:
 6f0:   f3 c3                   repz retq 

Disassembly of section .fini:

00000000000006f4 <_fini>:
 6f4:   48 83 ec 08             sub    $0x8,%rsp
 6f8:   48 83 c4 08             add    $0x8,%rsp
 6fc:   c3                      retq 
 ```


 ### asm
```
$ gcc -S hello.c
-rw-r--r-- 1 good dev    582  5월 10 23:19 hello.s
```



