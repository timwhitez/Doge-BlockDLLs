package main

import (
	"github.com/D00MFist/Go4aRun/pkg/sliversyscalls/syscalls"
	"golang.org/x/sys/windows"

	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"
)

func CreateProcess(startupInfo syscalls.StartupInfoEx){
	target := "C:\\Windows\\System32\\notepad.exe"
	commandLine, err := syscall.UTF16PtrFromString(target)

	if err != nil {
		panic(err)
	}

	var procInfo windows.ProcessInformation
	startupInfo.Cb = uint32(unsafe.Sizeof(startupInfo))
	creationFlags := windows.EXTENDED_STARTUPINFO_PRESENT
	if err = syscalls.CreateProcess(
		nil,
		commandLine,
		nil,
		nil,
		true,
		uint32(creationFlags),
		nil,
		nil,
		&startupInfo,
		&procInfo);
		err != nil {
		log.Printf("CreateProcess failed: %v\n", err)
	}

	return
}

func main(){
	procThreadAttributeSize := uintptr(0)
	_ = syscalls.InitializeProcThreadAttributeList(nil, 2, 0, &procThreadAttributeSize)
	procHeap, _ := syscalls.GetProcessHeap()
	attributeList, _ := syscalls.HeapAlloc(procHeap, 0, procThreadAttributeSize)
	defer syscalls.HeapFree(procHeap, 0, attributeList)
	var startupInfo syscalls.StartupInfoEx
	startupInfo.AttributeList = (*syscalls.PROC_THREAD_ATTRIBUTE_LIST)(unsafe.Pointer(attributeList))
	_ = syscalls.InitializeProcThreadAttributeList(startupInfo.AttributeList, 2, 0, &procThreadAttributeSize)
	mitigate := 0x20007 //"PROC_THREAD_ATTRIBUTE_MITIGATION_POLICY"
	//Options for Block Dlls
	nonms := uintptr(0x100000000000|0x1000000000)     //"PROCESS_CREATION_MITIGATION_POLICY_BLOCK_NON_MICROSOFT_BINARIES_ALWAYS_ON"|"PROCESS_CREATION_MITIGATION_POLICY_PROHIBIT_DYNAMIC_CODE_ALWAYS_ON"
	onlystore := uintptr(0x300000000000|0x1000000000) //"BLOCK_NON_MICROSOFT_BINARIES_ALLOW_STORE"
	block := "nonms"

	if block == "nonms" {
		_ = syscalls.UpdateProcThreadAttribute(startupInfo.AttributeList, 0, uintptr(mitigate), &nonms, unsafe.Sizeof(nonms), 0, nil)
	} else if block == "onlystore" {
		_ = syscalls.UpdateProcThreadAttribute(startupInfo.AttributeList, 0, uintptr(mitigate), &onlystore, unsafe.Sizeof(onlystore), 0, nil)
	} else {
		fmt.Println("wrong block mode")
	}
	var si syscalls.StartupInfoEx
	si.AttributeList = startupInfo.AttributeList
	CreateProcess(si)
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
