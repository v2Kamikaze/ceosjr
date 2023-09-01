package main

import (
	"fmt"
	"log"
	"os"

	"github.com/v2Kamikaze/ceosjr/src/list"
	"github.com/v2Kamikaze/ceosjr/src/matrix"
)

func main() {
	matrices()
	lists()
}

func matrices() {
	m := matrix.New()

	if len(os.Args) < 2 {
		fmt.Println("Passe o nome do arquivo que contém a matriz de inteiros como argumento. Ex: ./out ./teste.txt ou go run main.go ./teste.txt")
		os.Exit(1)
	}

	if err := m.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	}

	fmt.Println("####################################  Matrizes ####################################")
	fmt.Println("Matriz original: ")
	m.Print()
	m.ReplaceRepeated()
	fmt.Println("Matriz sem repetições de elementos: ")
	m.Print()
}

func lists() {
	fmt.Println("############################## Manipulação de listas ##############################")
	l := []int{1, 2, 3, 4, 56, 4, 7, 74, 99}

	fmt.Printf("O terceiro maior da lista %+v é: %d\n\n", l, list.ThirdGreatest(l))
	l1, l2 := []int{1, 1, 2, 3, 7, 4}, []int{6, 85, 76, 4, 5, 5}
	fmt.Printf("Listas: %+v e %+v\n\n", l1, l2)
	if is, l1, l2 := list.ListPartition(l1, l2); !is {
		fmt.Printf("Impossível reorganizar l1 e l2 de modo que todos os elementos de l1 sejam menores que de l2: %+v e %+v\n\n", l1, l2)
	} else {
		fmt.Printf("Ao reorganizar l1 e l2, todos os elementos em l1 são menores que os de l2: %+v e %+v\n\n", l1, l2)
	}

	l1, l2 = []int{1, 2, 5, 5, 5, 5, 5, 3}, []int{3, 2, 5, 5, 5, 5, 5, 1}
	if is := list.Permutations(l1, l2); !is {
		fmt.Printf("A lista l1 %+v não é uma permutação da lista l2 %+v\n\n", l1, l2)
	} else {
		fmt.Printf("A lista l1 %+v é uma permutação da lista l2 %+v\n\n", l1, l2)
	}
}
