package main

/*
#include <stdio.h>
#include <Windows.h>
#include <Processthreadsapi.h>
#pragma comment(lib, "Advapi32.lib")


static void add_mitigations(HANDLE hProc)
{

	PROCESS_MITIGATION_BINARY_SIGNATURE_POLICY signature = { 0 };
	GetProcessMitigationPolicy(hProc, ProcessSignaturePolicy, &signature, sizeof(signature));

	printf("ProcessSignaturePolicy:\n");
	printf("   MicrosoftSignedOnly                        %u\n", signature.MicrosoftSignedOnly);
	signature.MicrosoftSignedOnly = 1;


	if (!SetProcessMitigationPolicy(ProcessSignaturePolicy, &signature, sizeof(signature))) {
		printf("[!] ProcessSignaturePolicy failed\n");
		return;
	}
	printf("ProcessSignaturePolicy:\n");
	printf("   MicrosoftSignedOnly                        %u\n", signature.MicrosoftSignedOnly);
}

int test()
{
	HANDLE hProcess = GetCurrentProcess();
	add_mitigations(hProcess);
	//	getchar();
	return 0;
}
*/
import "C"
import (
	"bufio"
	"fmt"
	"os"
)

func init(){
	C.test()
}

func main() {

	fmt.Println("Gotcha!")
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

}