package main

import "fmt"

func main() {
	map1 := map[string][]string{
		"henrique_matheus": []string{"arroz", "feijão", "macarrão"},
		"Tú":               []string{"pedir carinho", "comer", "dormir"},
		"Nata":             []string{"aperriar", "chorar", "atormentar"},
		"SHirley":          []string{"arroz", "feijão", "macarrão"},
	}

	map1["arroz"] = []string{"feijão"}

	delete(map1, "Nata")

	for s, _ := range map1 {
		fmt.Println(s, map1[s])
	}
}
