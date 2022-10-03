#include <sys/syscall.h>
#include <unistd.h>
int main (int argc, char *agrv[]){
    syscall(SYS_write,1,"hello, world\n",14);
    return 0;
}