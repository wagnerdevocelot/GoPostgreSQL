# Se conectando ao PostgreSQL usando Golang

![](https://cdn-images-1.medium.com/max/800/1*rNIlHxSCv9vXoTKMOQJibg.png)

## Nem tudo precisa ser na CLI - Instalando o PostgreSQL e o pgAdmin4

O inicio é bem simples e vai depender mais do SO, estou usando o _postgreSQL_ pois é um dos databases mais utilizados, mas você pode utilizar outro banco relacional com pequenas modificações.

Não se trata de um tutorial sobre SQL em si e sim de fazer a conexão do seu banco de dados com o _Golang_.

Por esse motivo eu vou instalar e administrar o _Postgres_ da forma mais simples possível, que é através de um client, pois não sou de ferro. Estou usando o _Manjaro_ no momento e vou instalar o _Postgres_ através do gerenciador de aplicações e também vou instalar o _pgAdmin_ para gerenciamento de servidores e bancos de dados visual no PostgreSQL.

Acredito que essa seja uma abordagem interessante para iniciantes que não tem experiência com banco de dados, existe um folclore muito grande a respeito de UPDATE sem WHERE por aí e apesar das histórias parecerem engraçadas depois de um tempo o medo é real hahaha.

![](https://cdn-images-1.medium.com/max/800/1*HolNkoEh4g3UtGsxAu-Usw.png)

![](https://cdn-images-1.medium.com/max/800/1*k9_nGErbDBQ-XbRh85cwmQ.png)

Na loja de apps do seu SO é muito provável que você encontre os dois mas de qualquer forma vou deixar alguns links.

Aqui vou deixar um link pra quem quiser [baixar o PostgreSQL no ubuntu](https://www.geeksforgeeks.org/install-postgresql-on-linux/).

Aqui o link de download do [pgAdmin4](https://www.pgadmin.org/download/)

Terminando de instalar os dois, quando abrir o pgAdmin ele vai subir o client no seu browser e em seguida pedir pra você definir uma senha, vou definir a minha como “vapordev123”.

![](https://cdn-images-1.medium.com/max/800/1*kN3kL0PdDnbe_DfZfPn6Ng.png)

Clicando com o botão direito em cima de **_servers_** podemos criar um, vou chamar o meu de “AHAB”.

![](https://cdn-images-1.medium.com/max/800/1*QBJY5e497gIavteQOJ1iSg.png)

Abaixo uma mensagem de erro nos avisa de que precisamos definir mais informações.

![](https://cdn-images-1.medium.com/max/800/1*GVAxjAHLAbmpGSji149POg.png)

Na aba de **_connection_** vamos definir o nosso **_host_** que será o “localhost” e a **_porta_** que será “5432” o restante vamos manter, clicamos em **_save_** e temos nosso servidor.

![](https://cdn-images-1.medium.com/max/800/1*eKSgUxuNPsdkDYxgSIQUcQ.png)

![](https://cdn-images-1.medium.com/max/800/1*GyUdLS5vb2i14hSwLe3j8w.png)

Vou criar agora um banco de dados, clico com o botão direito em cima do server **_AHAB_**, create/database.

Aqui só precisamos definir o nome do **_banco_**, que será “MobyDick”, e agora **_save_**!. O Banco de dados tá pronto, agora vamos criar a tabela que vamos manipular usando Go.

![](https://cdn-images-1.medium.com/max/800/1*9g-ok78SOYkzWrSTOTF11g.png)

Acompanhe na imagem:  
- Azul: Clique em cima para selecionar **_MobyDick_**.  
- Vermelho: Clique no **_Query Tool_** e abrira um console para escrevermos.  
- Verde: Nossa tabela se chama “article”, terá um “id”, “title” e “body”.  
- Rosa: Clique no botão de execução para criar a tabela.

![](https://cdn-images-1.medium.com/max/800/1*AoPdsBfD8wpndXd_a72wjg.png)

```text
CREATE TABLE article (
	id SERIAL PRIMARY KEY,
	title VARCHAR(255),
	body TEXT

);
```

![](https://cdn-images-1.medium.com/max/800/1*j5CRLVeEJeNhNDCfd8G0_A.png)

Clicando com o botão direito sobre **_MobyDick_** clique em **_refresh_** e agora você pode verificar a tabela em:

MobyDick/Schemas/public/Tables, e **_article_** estará lá.

Os bancos de dados tem seus próprios tipos, para ter uma ideia simples da tabela que criamos, **_id_** é apenas um identificador, ele serve também para quando tivermos uma nova tabela criar uma relação entre elas passando o id como uma chave estrangeira.

Nosso **_title_** recebeu um **_VARCHAR(255)_** que é um dado de caractere como _string_ e tem uma limitação de _255_ caracteres.

**_Body_** recebeu o tipo **_TEXT_** pois esse é um dado de _caractere_ que não tem limite, que o torna o candidato mais adequado para ser o corpo de um artigo.

Aqui mesmo no _pgAdmin_ vamos inserir o primeiro artigo na nossa tabela antes de ir ao código _Golang_.

Clique no **_Query Tool_** e use o seguinte código e execute.

```text
INSERT INTO article (id, title, body) VALUES (1, 'Golang + PostgreSQL', 'But I must explain to you how all this mistaken idea');
```

Esse é um artigo em que no body eu coloquei um “lorem ipsum”. Mas a estrutura desse **_insert_** é essa:

```text
INSERT INTO article (id,title,body) VALUES (O que você quer inserir separado por vírgulas);
```

![](https://cdn-images-1.medium.com/max/800/1*XJmv61OBTz7xKItulBL8yw.png)

Depois que conseguir damos um **_refresh_** no **_DB_** e vamos agora fazer um **_select_** para verificar nossos dados.

```text
SELECT * FROM article
```

Isso deve retornar na parte de baixo a tabela e dados inseridos. Estou usando cada vez menos imagens pois imagino que você deva estar se habituando a página de administração do DB e também para evitar repetição.

## Database Driver Config


A estrutura das pastas é assim:

```text
gopostgres {
	dbconfig {
		driverConfig.go
	}
	main.go
}
```	

Você pode escolher o nome que for melhor, mas eu fiz uma pasta diferente para o arquivo *driverConfig.go* pois vou usá-lo como pacote e importar no *main.go* mais a frente.

```go
package dbconfig

import "fmt"

type Article struct {
	ID    int
	Title string
	Body  []byte
}
```

A struct que será a nossa representação da tabela *Article* aqui no código, caso você tenha visto o tutorial o [WEBtrhough](https://github.com/wagnerdevocelot/gowiki) a estrutura é parecida, foi proposital pra que esse artigo seja um complemento.

```go
const PostgresDriver = "postgres"

const User = "postgres"

const Host = "localhost"

const Port = "5432"

const Password = "vapordev123"

const DbName = "MobyDick"

const TableName = "article"
```

Uma série de constantes, com strings que já são familiares, todas as informações que definimos no banco de dados usaremos aqui para que a conexão possa ser feita.

A única constante definida aqui em função do código é a *PostgresDriver*, vamos usar um driver de terceiros pois não existe opção na biblioteca padrão.

O driver é especifico de acordo com o banco de dados que você pretende utilizar, e com as suas ambições, alguns drivers vem com um toolkit para aplicações, aqui eu quero apenas fazer a conexão e usar a biblioteca padrão de _SQL_ do _Go_ para fazer as operações necessárias.

```go
var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
```

Essa é a criação da string completa com todas as informações necessárias para conectar e fazer operações.

A única coisa que não cheguei a mencionar for essa ultima parte da string "_sslmode_" quando criamos nosso banco de dados não fizemos modificação no *SSL*, então aqui o valor na string é "_disable_" pois esse é o valor padrão.

A função *Sprintf* envia strings formatadas, então iremos usar *DataSourceName* como argumento mais adiante.

## main.go

O arquivo *main.go* fica fora da pasta *dbconfig* por isso precisamos importar aqui *./dbconfig* e usamos uma variável para chamar os itens presentes no arquivo *driverConfig.go*.

Importamos também "*fmt*", o pacote *SQL*, e o *github.com/lib/pq* que é o driver do *postgres* temos um underscore na importação do driver pois não usaremos funções dele aqui, só iremos enviar a string de conexão.

Como já disse as libs de drivers são de terceiros então você precisará usar _go get_ para fazer download do pacote.

```text
go get github.com/lib/pq
```

Duas variáveis com package level scope *db* que aponta para o struct *DB* presente no pacote *SQL* e a *err* do tipo error que será usada em conjunto com a função *checkErr* para checagem de erros adiante.

```go
package main

import (
	"database/sql"
	"fmt"

	dbConfig "./dbconfig"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error
```
Uma função bem simples para tratar erros durante o uso das queries.

```go
func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
```

Na *main func* sinalizamos o acesso e você já pode ver como usamos dados importados do *driverConfig.go* usando o nome definido aqui na importação e chamando o nome original presente na constante.

Atribuímos as variáveis o valor da função de *Open* do *sql* que recebe como parâmetro, o *driver* e a *dataSourceName* com todas as constantes definidas como senha, host e etc.

Fazemos *error handling* para checar tudo ocorreu como esperado, pode ser usada também a função *Ping()* que bate no host do banco pra verificar a conexão.

O *defer* é um statement do _Golang_, ele serve pra fazer o adiamento de algo, abaixo de *db.Close* tem as funções que para manipular a nossa tabela, usar um defer aqui quer dizer: "Fecha a conexão com o banco de dados, mas só depois que terminar de executar essas chamadas de função".

```go
func main() {

	fmt.Printf("Accessing %s ... ", dbConfig.DbName)

	db, err = sql.Open(dbConfig.PostgresDriver, dbConfig.DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}

	defer db.Close()

	sqlSelect()
	sqlSelectID()
	sqlInsert()
	sqlUpdate()
	sqlDelete()
}
```
Para testar a conexão só precisamos comentar as chamadas de funções com "//" e também não podemos esquecer de deixar aberto o client do PostgreSQL pois precisamos bater no localhost para obter as respostas.

Output:
```text
Accessing MobyDick ... Connected!
```

## SELECT

A estrutura original do _SELECT_ que a maioria deve conhecer:

```text
SELECT id, title, body FROM article
```

Ou poderia ser:

```text
SELECT * FROM article
```

Significa, "Postgre, me da todas as informações dessa tabela aqui", isso é chamado de _statement_ no _SQL_ que você pode traduzir e chamar de "declaração", o _SELECT_ seria apenas uma dessas declarações.

```go
func sqlSelect() {

	sqlStatement, err := db.Query("SELECT id, title, body FROM " + dbConfig.TableName)
	checkErr(err)

	for sqlStatement.Next() {

		var article dbConfig.Article

		err = sqlStatement.Scan(&article.ID, &article.Title, &article.Body)
		checkErr(err)

		fmt.Printf("%d\t%s\t%s \n", article.ID, article.Title, article.Body)
	}
}
```

Temos um *error handling*.

A função *db.Query* retorna "_rows_" que são as linhas do sql, essas linhas precisam ser scaneadas uma a uma com auxilio de outras duas funções.

Usamos o *for* com *sqlStatement.Next* o Next prepara as linhas do _DB_ uma a uma para serem lidas pelo *sqlStatement.Scan*.

Instanciámos *article* para passar os valores do _DB_ para a struct, o método *Scan* recebe os valores da struct e transforma os tipos do _PostgreSQL_ como VARCHAR, ID, TEXT em tipos do Go.

Depois de recuperados os dados printamos eles no console.
Essa é a primeira linha que adicionei na tabela usando o _pgAdmin_.

Output:
```text
Accessing MobyDick ... Connected!

1	Golang + PostgreSQL	But I must explain to you how all this mistaken idea 
```

## SELECT BY ID

Fazemos um _SELECT_ para buscar apenas uma linha da tabela pelo seu *id*.

```text
SELECT id, title, body FROM article WHERE id 1
```

"Postgre, me retorna id, title e body da tabela article onde o id é 1"

O *Sprintf* para enviar o *sqlStatement* para *db.QueryRow* que é usado quando se precisa retornar apenas uma linha da tabela, podendo receber mais argumentos além do _statement_, nesse caso "1" que é o id da linha que devemos retornar.

Em seguida a mesma estrutura de *Scan* e *Print* que usamos para resgatar os valores e imprimir o output.

```go
func sqlSelectID() {

	var article dbConfig.Article

	sqlStatement := fmt.Sprintf("SELECT id, title, body FROM %s where id = $1", dbConfig.TableName)

	err = db.QueryRow(sqlStatement, 1).Scan(&article.ID, &article.Title, &article.Body)
	checkErr(err)

	fmt.Printf("%d\t%s\t%s \n", article.ID, article.Title, article.Body)
}
```

```text
Accessing MobyDick ... Connected!

1	Golang + PostgreSQL	But I must explain to you how all this mistaken idea
```
Temos o mesmo resultado pois não fizemos alterações com a tabela e esse _select_ usa o *id* "1" como parâmetro, vou mostrar mais um output usando "2" como *id* pra mostrar como fica a saída de erro.

```text
Accessing MobyDick ... Connected!
panic: sql: no rows in result set

goroutine 1 [running]:
main.checkErr(...)
	/home/vapordev/workspace/GoPostgres/main.go:114
main.sqlSelectID()
	/home/vapordev/workspace/GoPostgres/main.go:58 +0x34b
main.main()
	/home/vapordev/workspace/GoPostgres/main.go:29 +0x1aa
exit status 2
```

Ele conecta com o _DB_ mas não retorna a linha pois ela não existe.
Podemos ver tbm as linhas com o rastro do erro, linha *114* é onde está a nossa função *checkErr()* então toda fonte de erro começa nela, em seguida temos linha *58* que é exatamente onde o erro ocorreu e linha *29* que é de onde foi feita a chamada da função.

## INSERT

Um _statement_ que insere informações na tabela, esse e os demais tem uma estrutura muito parecida pois nenhum precisará retornar uma tabela só precisaremos verificar se a ação foi executada com sucesso.

Criamos o statement:

```text
INSERT INTO article VALUES ("o que você quer inserir separado por virgulas")
```

O *db.Prepare* cria uma instrução para queries ou execuções posteriores, várias queries ou execuções podem ser executadas de forma concorrente.

O *db.Prepare* atribuído na variável *insert* é usado com o método *Exec* que executa uma instrução preparada com os argumentos e retorna um Resultado. Nessa etapa é onde inserimos os dados que foram interpolados no _statement_ ($1=id, $2=title, $3=body).

*RowsAffected* retorna o número de linhas afetadas pela operação. Nem todos os bancos suportam essa função, no caso o *RowsAffected* faz parte de uma interface result que tem outra função chamada *LastInsertId()* que nesse caso não funciona com o driver do postgres então vamos usar *RowsAffected* mesmo.

```go
func sqlInsert() {

	sqlStatement := fmt.Sprintf("INSERT INTO %s VALUES ($1,$2, $3)", dbConfig.TableName)

	insert, err := db.Prepare(sqlStatement)
	checkErr(err)

	result, err := insert.Exec(5, "Maps in Golang", "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium")
	checkErr(err)

	affect, err := result.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}
```
A saida mostrando que a conecxão foi bem sucedida e que uma linha foi afetada, a partir daqui não temos retorno de linhas da tabela, mas se quisermos verificar o conteúdo do banco basta dar um novo _select_.

Output:
```text
Accessing MobyDick ... Connected!
1
```

A saida após feito o insert e utilizado _select_ novamente.

Output:
```text
Accessing MobyDick ... Connected!
1	Golang + PostgreSQL	But I must explain to you how all this mistaken idea 
5	Maps in Golang	Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium 
```

## UPDATE

A partir daqui as coisas ficam entediantes pois o que muda são os padrões de _statement_, mas o código tem pouca diferença.

Novamente preparamos uma string com o _Statement_ da vez que é um _UPDATE_.

```text
UPDATE article SET body WHERE id 5
```

"Vamos fazer uma modificação em article no campo body do elemento com o id numero 5"

```go
func sqlUpdate() {

	sqlStatement := fmt.Sprintf("update %s set body=$1 where id=$2", dbConfig.TableName)

	update, err := db.Prepare(sqlStatement)
	checkErr(err)

	result, err := update.Exec("But I must explain to you how all this mistaken idea", 5)
	checkErr(err)

	affect, err := result.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}
```

Temos a mesma estrutura agora porém com apenas dois argumentos no *updade.Exec* o *body* e o *id* da linha que queremos modificar na tabela. 

O identificador no final das contas ajuda muito nas buscas quando precisamos editar algo pois é mais fácil de usar como referência.

A saida após o _UPDATE_ na linha que acabamos de inserir na tabela, nesse caso mudamos o *body*.

Output:
```text
Accessing MobyDick ... Connected!
1
```

Aqui vou dar mais um _SELECT_ pra mostrar o resultado da tabela após as alterações, você pode chamar a função de *sqlSelect()* que foi criada, dentro das outras funções pra verificar a tabela após a modificação se quiser.

Output:
```text
Accessing MobyDick ... Connected!
1	Golang + PostgreSQL	But I must explain to you how all this mistaken idea 
5	Maps in Golang	But I must explain to you how all this mistaken idea 
```

## DELETE

O _DELETE_ é ainda mais simples pois só precisamos do identificador para passar como referência do que queremos deletar.

```text
DELETE FROM article WHERE id=1
```

"Deleta da tabela article o item de id 1"

```go
func sqlDelete() {

	sqlStatement := fmt.Sprintf("delete from %s where id=$1", dbConfig.TableName)

	delete, err := db.Prepare(sqlStatement)
	checkErr(err)

	result, err := delete.Exec(5)
	checkErr(err)

	affect, err := result.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}
```

Aqui temos também o retorno de linhas afetadas e vou chamar o _SELECT_ novamente para mostrar o estado atual da tabela.

Output:
```text
Accessing MobyDick ... Connected!
1
```

Chamando o _SELECT_ após as alterações.

Output:
```text
Accessing MobyDick ... Connected!
1	Golang + PostgreSQL	But I must explain to you how all this mistaken idea 
```

Sobrou apenas a linha original.

Existem outras alterações com SQL que podem ser feitas com duas ou mais tabelas como JOINS, esse não é o caso então podemos deixar essa para uma proxima.