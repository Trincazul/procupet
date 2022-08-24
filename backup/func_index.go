

// Função usada para renderizar o arquivo Index
func Index(w http.ResponseWriter, r *http.Request) {
	// Abre a conexão com o banco de dados utilizando a função dbConn()
	db := dbConn()
	// Realiza a consulta com banco de dados e trata erros
	selDB, err := db.Query("SELECT * FROM animalper ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	// Monta a struct para ser utilizada no template
	n := Animalper{}

	// Monta um array para guardar os valores da struct
	res := []Animalper{}

	// Realiza a estrutura de repetição pegando todos os valores do banco
	for selDB.Next() {
		// Armazena os valores em variáveis
		var id int
		var nome string
		var raca string
		var localperd string
		var tipani string

		// Faz o Scan do SELECT
		err = selDB.Scan(&id, &nome, &raca, &localperd, &tipani)
		if err != nil {
			panic(err.Error())
		}

		// Envia os resultados para a struct
		n.Id = id
		n.Nome = nome
		n.Raca = raca
		n.Localperd = localperd
		n.Tipani = tipani

		// Junta a Struct com Array
		res = append(res, n)
	}

	var tmpl = template.Must(template.ParseGlob("tmpl/*"))
	// Abre a página Index e exibe todos os registrados na tela
	tmpl.ExecuteTemplate(w, "Index", res)

	// Fecha a conexão
	defer db.Close()
}
