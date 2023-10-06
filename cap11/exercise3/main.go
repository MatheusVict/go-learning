package main

import "fmt"

type veiculo struct {
	portas string
	cor    string
}

type caminhonente struct {
	veiculo
	quatroRodas bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {

	vei := veiculo{cor: "fadjskl", portas: "çdfjkaslfas"}

	cam := caminhonente{quatroRodas: false, veiculo: vei}
	sed := sedan{modeloLuxo: false, veiculo: vei}

	fmt.Println(cam, sed)

}
