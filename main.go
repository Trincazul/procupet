package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Animalper struct {
	Id        int
	Nome      string
	Raca      string
	Localperd string
	Tipani    string
}

type Animalenc struct {
	Id       int
	Defini   string
	Raca     string
	Localenc string
	Tipani   string
}

func dbConn() (db *sql.DB) {

	db, err := sql.Open("mysql", "root:animal@tcp(127.0.0.1:3306)/procupet")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Insert(w http.ResponseWriter, r *http.Request) {

	//Abre a conexão com banco de dados usando a função: dbConn()
	db := dbConn()

	// Verifica o METHOD do fomrulário passado
	if r.Method == "POST" {

		// Pega os campos do formulário
		nome := r.FormValue("nome")
		raca := r.FormValue("raca")
		localperd := r.FormValue("localperd")
		tipani := r.FormValue("tipani")

		// Prepara a SQL e verifica errors
		insForm, err := db.Prepare("INSERT INTO animalper(nome, raca, localperd, tipani) VALUES(?,?,?,?)")
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
		nome := r.FormValue("nome")
		raca := r.FormValue("raca")
		localperd := r.FormValue("localperd")
		tipani := r.FormValue("tipani")
		id := r.FormValue("id")

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
		var Id int
		var Nome string
		var Raca string
		var Localperd string
		var Tipani string

		// Faz o Scan do SELECT
		err = selDB.Scan(&Id, &Nome, &Raca, &Localperd, &Tipani)
		if err != nil {
			panic(err.Error())
		}

		// Envia os resultados para a struct
		n.Id = Id
		n.Nome = Nome
		n.Raca = Raca
		n.Localperd = Localperd
		n.Tipani = Tipani

		// Junta a Struct com Array
		res = append(res, n)
	}

	// Abre a página Index e exibe todos os registrados na tela
	tmpl.ExecuteTemplate(w, "Index", res)

	// Fecha a conexão
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	// Pega o ID do parametro da URL
	nId := r.URL.Query().Get("id")

	// Usa o ID para fazer a consulta e tratar erros
	selDB, err := db.Query("SELECT * FROM animalper WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	// Monta a strcut para ser utilizada no template
	n := Animalper{}

	// Realiza a estrutura de repetição pegando todos os valores do banco
	for selDB.Next() {
		// Armazena os valores em variaveis
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
	}

	// Mostra o template
	tmpl.ExecuteTemplate(w, "Show", n)

	// Fecha a conexão
	defer db.Close()

}

// Função New apenas exibe o formulário para inserir novos dados
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

// Função Edit, edita os dados
func Edit(w http.ResponseWriter, r *http.Request) {
	// Abre a conexão com banco de dados
	db := dbConn()

	// Pega o ID do parametro da URL
	nId := r.URL.Query().Get("id")

	selDB, err := db.Query("SELECT * FROM animalper WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	// Monta a struct para ser utilizada no template
	n := Animalper{}

	// Realiza a estrutura de repetição pegando todos os valores do banco
	for selDB.Next() {
		//Armazena os valores em variaveis
		var Id int
		var Nome string
		var Raca string
		var Localperd string
		var Tipani string

		// Faz o Scan do SELECT
		err = selDB.Scan(&Id, &Nome, &Raca, &Localperd, &Tipani)
		if err != nil {
			panic(err.Error())
		}

		// Envia os resultados para a struct
		n.Id = Id
		n.Nome = Nome
		n.Raca = Raca
		n.Localperd = Localperd
		n.Tipani = Tipani
	}

	// Mostra o template com formulário preenchido para edição
	tmpl.ExecuteTemplate(w, "Edit", n)

	// Fecha a conexão com o banco de dados
	defer db.Close()
}

func main() {
	log.Println("Servidor rodando em: http://localhost:9000")

	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)

	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

	http.ListenAndServe(":9000", nil)

}

var tmpl = template.Must(template.ParseGlob("tmpl/*"))
