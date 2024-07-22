package userinput_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/RaulCatalinas/HuskyBC/internal/constants"
	userinput "github.com/RaulCatalinas/HuskyBC/internal/user-input"
)

// TestPrintListPackageManeger prueba la función PrintListPackageManeger
func TestPrintListPackageManeger(t *testing.T) {
	// Guardamos la salida original de os.Stdout
	oldStdout := os.Stdout

	// Creamos un pipe para capturar la salida
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Llamamos a la función que queremos probar
	userinput.PrintListPackageManeger()

	// Cerramos el writer y restauramos os.Stdout
	w.Close()
	os.Stdout = oldStdout

	// Leemos la salida capturada
	var buf bytes.Buffer
	buf.ReadFrom(r)

	// Comprobamos que la salida es la esperada
	expectedOutput := "0. npm\n1. yarn\n2. pnpm\n3. bun\n"
	if buf.String() != expectedOutput {
		t.Errorf("Expected output %q, but got %q", expectedOutput, buf.String())
	}
	// Imprimimos la salida del código principal para referencia
	t.Logf("Output of PrintListPackageManeger: %q", buf.String())
}

// TestGetPackageManager prueba la función GetPackageManager
func TestGetPackageManager(t *testing.T) {
	// Guardamos la salida original de os.Stdout y la entrada original de os.Stdin
	oldStdout := os.Stdout
	oldStdin := os.Stdin

	// Creamos un pipe para capturar la salida
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut

	// Creamos un pipe para simular la entrada del usuario
	rIn, wIn, _ := os.Pipe()
	os.Stdin = rIn

	// Escribimos la entrada simulada
	wIn.WriteString("2\n")
	wIn.Close()

	// Llamamos a la función que queremos probar
	selectedPackageManager := userinput.GetPackageManager()

	// Cerramos el writer de salida y restauramos os.Stdout y os.Stdin
	wOut.Close()
	os.Stdout = oldStdout
	os.Stdin = oldStdin

	// Leemos la salida capturada
	var bufOut bytes.Buffer
	bufOut.ReadFrom(rOut)

	// Comparamos la salida con la salida esperada
	expectedOutput := "Which package manager do you wanna use?\n0. npm\n1. yarn\n2. pnpm\n3. bun\n"
	if bufOut.String() != expectedOutput {
		t.Errorf("Expected output %q, but got %q", expectedOutput, bufOut.String())
	}

	// Imprimimos la salida del código principal para referencia
	t.Logf("Output of GetPackageManager: %q", bufOut.String())

	expectedPackageManager := constants.PACKAGE_MANAGERS[2]
	if selectedPackageManager != expectedPackageManager {
		t.Errorf("Expected package manager %q, but got %q", expectedPackageManager, selectedPackageManager)
	}
}
