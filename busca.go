package main

import (
	"fmt"
	"time"
)

func Busca(testeSlice []int, valor int) int {
	for i, v := range testeSlice {
		if v == valor {
			return i
		}
	}
	return -1
}

func BuscaBinaria(testeSlice []int, valor int) int {
	inicio := 0
	fim := len(testeSlice) - 1
	for inicio <= fim {
		meio := (inicio + fim) / 2
		if testeSlice[meio] == valor {
			return meio
		} else if testeSlice[meio] < valor {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}
	return -1
}

func main() {
	testeSlice := []int{}
	var escolha int
	var numero int

	fmt.Println("Escolha uma das opções abaixo: ")
	fmt.Println("1 - Teste com 100 elementos.\n2 - Teste com 100.000 elementos.\n3 - Teste com 100.000.000 elementos.\n4 - Finalizar")
	fmt.Scanf("%d\n", &escolha)

	for escolha > 0 && escolha < 5 {
		if escolha == 1 {
			for i := 0; i < 100; i++ {
				testeSlice = append(testeSlice, i)
			}
			fmt.Println("Informe o valor que deseja buscar: ")
			fmt.Scanf("%d\n", &numero)
			fmt.Println("Busca: ", Busca(testeSlice, numero))
			fmt.Println("Busca Binária: ", BuscaBinaria(testeSlice, numero))
		} else if escolha == 2 {
			for i := 0; i < 100000; i++ {
				testeSlice = append(testeSlice, i)
			}
			fmt.Println("Informe o valor que deseja buscar: ")
			fmt.Scanf("%d\n", &numero)
			fmt.Println("Busca: ", Busca(testeSlice, numero))
			fmt.Println("Busca Binária: ", BuscaBinaria(testeSlice, numero))
		} else if escolha == 3 {
			for i := 0; i < 100000000; i++ {
				testeSlice = append(testeSlice, i)
			}
			fmt.Println("Informe o valor que deseja buscar: ")
			fmt.Scanf("%d\n", &numero)
			t0 := time.Now()
			fmt.Println("Busca sequenciada: ", Busca(testeSlice, numero))
			t1 := time.Now()
			fmt.Println("Demorou: ", t1.Sub(t0))
			t3 := time.Now()
			fmt.Println("Busca Binária: ", BuscaBinaria(testeSlice, numero))
			t4 := time.Now()
			fmt.Println("Demorou: ", t4.Sub(t3))

		} else if escolha == 4 {
			fmt.Println("Finalizado.")
			break
		} else {
			fmt.Println("Opção inválida.")
		}
	}
}
