package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	execPath := "./tmp/__debug_bin.exe"

	fmt.Printf("=== INICIANDO DLV WRAPPER ===\n")
	fmt.Printf("Executável: %s\n", execPath)

	args := []string{
		"exec",
		execPath,
		"--listen=127.0.0.1:2345",
		"--headless=true",
		"--api-version=2",
		"--accept-multiclient",
		"--continue",
		"--log",
	}

	cmd := exec.Command("dlv", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Criar novo grupo de processos
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}

	// Capturar sinais de interrupção
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Iniciar o processo
	if err := cmd.Start(); err != nil {
		fmt.Printf("ERRO ao iniciar dlv: %v\n", err)
		os.Exit(1)
	}

	// Goroutine para aguardar sinais
	go func() {
		<-sigChan
		fmt.Println("\n=== RECEBIDO SINAL DE TÉRMINO ===")
		if cmd.Process != nil {
			// Enviar CTRL+BREAK para o grupo de processos no Windows
			dll, _ := syscall.LoadDLL("kernel32.dll")
			proc, _ := dll.FindProc("GenerateConsoleCtrlEvent")
			proc.Call(syscall.CTRL_BREAK_EVENT, uintptr(cmd.Process.Pid))

			cmd.Process.Kill()
		}
		os.Exit(0)
	}()

	// Aguardar o processo terminar
	if err := cmd.Wait(); err != nil {
		fmt.Printf("Processo dlv finalizado com erro: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("=== DLV WRAPPER FINALIZADO ===")
}
