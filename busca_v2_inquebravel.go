package main

import (
	"fmt"
	"time"
)

func buscaSequencia(quant int) {

	var numero int
	elementos := []int{}

	fmt.Println("\nUma busca Sequenciada será realizada.\nInforme o numero que deseja buscar: ")
	fmt.Scanf("%d\n", &numero)
	if numero >= quant || numero <= 0 {
		fmt.Println("\nNumero não encontrado, lembre-se, a contagem de elementos começa em 0.")
		buscaSequencia(quant)
	} else {
		for i := 0; i < quant; i++ {
			elementos = append(elementos, i)
		}
		for v := range elementos {
			if v == numero {
				fmt.Println("O numero ", numero, " foi encontrado na posição ", v+1, "\n")
			}
		}
	}
}

func buscaBinaria(quant int) {

	var numero int
	elementos := []int{}

	fmt.Println("\nUma busca binaria será realizada.\nInforme o numero que deseja buscar: ")
	fmt.Scanf("%d\n", &numero)

	if numero > quant-1 || numero < 0 {
		fmt.Println("\nNumero não encontrado, lembre-se, a contagem de elementos começa em 0.")
		buscaBinaria(quant)
	} else {
		for i := 0; i < quant; i++ {
			elementos = append(elementos, i)
		}
		inicio := 0
		fim := len(elementos) - 1
		for inicio <= fim {
			meio := (inicio + fim) / 2
			if elementos[meio] == numero {
				fmt.Println("O numero ", numero, " foi encontrado na posição ", meio+1)
				break
			} else if elementos[meio] < numero {
				inicio = meio + 1
			} else {
				fim = meio - 1
			}
		}
	}

}

func teste100() {

	var busca int
	elementos := []int{}

	for i := 0; i < 100; i++ {
		elementos = append(elementos, i)
	}

	fmt.Println("Informe o numero que deseja buscar: ")
	fmt.Scanf("%d\n", &busca)
	if busca >= 100 || busca < 0 {
		fmt.Println("Numero não encontrado, lembre-se, a contagem de elementos começa em 0.")
		teste100()
	} else {
		t0 := time.Now()
		for i, v := range elementos {
			if v == busca {
				fmt.Println("O numero ", busca, " foi encontrado na posição ", i+1)
			}
		}
		t1 := time.Now()
		fmt.Println("Tempo de execução: ", t1.Sub(t0))

		inicio := 0
		fim := len(elementos) - 1
		t2 := time.Now()
		for inicio <= fim {
			meio := (inicio + fim) / 2
			if elementos[meio] == busca {
				fmt.Println("O numero ", busca, " foi encontrado na posição ", meio+1)
				break
			} else if elementos[meio] < busca {
				inicio = meio + 1
			} else {
				fim = meio - 1
			}
		}
		t3 := time.Now()
		fmt.Println("Tempo de execução: ", t3.Sub(t2))

	}

}

func teste100000() {

	var busca int
	elementos := []int{}

	for i := 0; i < 100000; i++ {
		elementos = append(elementos, i)
	}

	fmt.Println("Informe o numero que deseja buscar: ")
	fmt.Scanf("%d\n", &busca)

	if busca >= 100000 || busca < 0 {
		fmt.Println("Numero não encontrado, lembre-se, a contagem de elementos começa em 0.")
		teste100000()
	} else {
		t0 := time.Now()
		for i, v := range elementos {
			if v == busca {
				fmt.Println("O numero ", busca, " foi encontrado na posição ", i+1)
			}
		}
		t1 := time.Now()
		fmt.Println("Tempo de execução: ", t1.Sub(t0))

		inicio := 0
		fim := len(elementos) - 1
		t2 := time.Now()
		for inicio <= fim {
			meio := (inicio + fim) / 2
			if elementos[meio] == busca {
				fmt.Println("O numero ", busca, " foi encontrado na posição ", meio+1)
				break
			} else if elementos[meio] < busca {
				inicio = meio + 1
			} else {
				fim = meio - 1
			}
		}
		t3 := time.Now()
		fmt.Println("Tempo de execução: ", t3.Sub(t2))

	}

}

func teste100000000() {

	var busca int
	elementos := []int{}

	for i := 0; i < 100000000; i++ {
		elementos = append(elementos, i)
	}

	fmt.Println("Informe o numero que deseja buscar: ")
	fmt.Scanf("%d\n", &busca)

	if busca >= 100000000 || busca < 0 {
		fmt.Println("Numero não encontrado, lembre-se, a contagem de elementos começa em 0.")
		teste100000000()
	} else {
		t0 := time.Now()
		for i, v := range elementos {
			if v == busca {
				fmt.Println("O numero ", busca, " foi encontrado na posição ", i)
			}
		}
		t1 := time.Now()
		fmt.Println("Tempo de execução: ", t1.Sub(t0))

		inicio := 0
		fim := len(elementos) - 1
		t2 := time.Now()
		for inicio <= fim {
			meio := (inicio + fim) / 2
			if elementos[meio] == busca {
				fmt.Println("O numero ", busca, " foi encontrado na posição ", meio)
				break
			} else if elementos[meio] < busca {
				inicio = meio + 1
			} else {
				fim = meio - 1
			}
		}
		t3 := time.Now()
		fmt.Println("Tempo de execução: ", t3.Sub(t2))
	}

}

func main() {

	var escolha = 1

	for escolha > 0 && escolha <= 5 {
		fmt.Println("Escolha uma opção: ")
		fmt.Println("1 - Busca\n2 - Teste 100\n3 - Teste 100000\n4 - Teste 100000000\n5 - Sair")
		fmt.Scanf("%d\n", &escolha)

		if escolha == 1 {
			var quant int
			fmt.Println("Informe a quantidade de elementos:")
			fmt.Scanf("%d\n", &quant)
			if quant > 0 && quant <= 128 {
				buscaSequencia(quant)
			}
			if quant > 128 {
				buscaBinaria(quant)
			}
		} else if escolha == 2 {
			teste100()
		} else if escolha == 3 {
			teste100000()
		} else if escolha == 4 {
			teste100000000()
		} else if escolha == 5 {
			break
		} else {
			fmt.Println("Opção invalida!\n")
			main()
		}
	}

}
