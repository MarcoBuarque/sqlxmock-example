# sqlxmock-example

# Run test
 make test > test.log 
# Infos about sqlxmock
### ExpectCommit
utilizado para DB.Commit

### Expect ExpectExec
utilizado para DB.Exec

### ExpectQuery
utilizado para  DB.QueryRow, DB.Select e DB.Get

### ExpectRollback
utilizado para DB.Rollback
### ExpectPrepare
utilizado para DB.Prepare
### ExpectBegin
utilizado para DB.MustBegin

### AnyArgs
utilizado quando existe algum argumento sobre o qual não temos controle
por exemplo utilizar o time.now() no createdAt

### ExpectationsWereMet
verificar se todas as expects foram satisfeitas, retornando erro caso ainda exista alguma expect na fila 

## MatchExpectationsInOrder
True: Todos os expects serão executados na ordem em que foram criados (default)
False: Utilizar quando os testes são executados paralelamente  

# Setup DB
Esse setup não é necessário para rodar os testes
## create DB
https://dev.to/vapordev/se-conectando-ao-postgresql-usando-golang-381h

## migrate
https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md