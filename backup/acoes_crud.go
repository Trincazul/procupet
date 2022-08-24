
// Função Insert, insere valores no banco de dados
func Insert(w http.ResponseWriter, r *http.Request) {

	//Abre a conexão com banco de dados usando a função: dbConn()
	db := dbConn()

	// Verifica o METHOD do fomrulário passado
	if r.Method == "POST" {

		// Pega os campos do formulário
		nome := r.FormValue("name")
		raca := r.FormValue("raca")
		localperd := r.FormValue("localperd")
		tipani := r.FormValue("tipani")

		// Prepara a SQL e verifica errors
		insForm, err := db.Prepare("INSERT INTO animalper(name, raca, localperd, tipani) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}

		// Insere valores do formulario com a SQL tratada e verifica errors
		insForm.Exec(nome, raca, localperd, tipani)

		// Exibe um log com os valores digitados no formulário
		log.Println("INSERT: Nome: " + nome + " | Raça: " + raca + " | Local Perdido: " + localperd + " | Tipo de Animal: " + tipani)
	}

	// Encerra a conexão do dbConn()
	defer db.Close()

	//Retorna a HOME
	http.Redirect(w, r, "/", 301)
}

// Função Update, atualiza valores no banco de dados
func Update(w http.ResponseWriter, r *http.Request) {

	// Abre a conexão com o banco de dados usando a função: dbConn()
	db := dbConn()

	// Verifica o METHOD do formulário passado
	if r.Method == "POST" {

		// Pega os campos do formulário
		nome := r.FormValue("name")
		raca := r.FormValue("raca")
		localperd := r.FormValue("localperd")
		tipani := r.FormValue("tipani")
		id := r.FormValue("uid")

		// Prepara a SQL e verifica errors
		insForm, err := db.Prepare("UPDATE animalper SET nome=?, raca=?, localperd=?, tipani=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}

		// Insere valores do formulário com a SQL tratada e verifica erros
		insForm.Exec(nome, raca, localperd, tipani, id)

		// Exibe um log com os valores digitados no formulario
		log.Println("UPDATE: Nome: " + nome + " | Raça: " + raca + " | Local Perdido: " + localperd + " | Tipo de Animal: " + tipani)
	}

	// Encerra a conexão do dbConn()
	defer db.Close()

	// Retorna a HOME
	http.Redirect(w, r, "/", 301)
}

// Função Delete, deleta valores no banco de dados
func Delete(w http.ResponseWriter, r *http.Request) {

	// Abre conexão com banco de dados usando a função: dbConn()
	db := dbConn()

	nId := r.URL.Query().Get("id")

	// Prepara a SQL e verifica errors
	delForm, err := db.Prepare("DELETE FROM animalper WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	// Insere valores do form com a SQL tratada e verifica errors
	delForm.Exec(nId)

	// Exibe um log com os valores digitados no form
	log.Println("DELETE")

	// Encerra a conexão do dbConn()
	defer db.Close()

	// Retorna a HOME
	http.Redirect(w, r, "/", 301)
}
