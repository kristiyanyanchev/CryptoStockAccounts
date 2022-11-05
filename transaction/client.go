package transaction

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "mouse.db.elephantsql.com"
	port     = 5432
	user     = "tfalxqpw"
	password = "P4Ewg8JM_QgiH6pmWhvQpIevj_XmsHvf"
	dbname   = "tfalxqpw"
)

type Transaction struct {
	Id     int
	UserId int
	Type   string
	Symbol string
	Amount float64
	Price  float64
	Date   time.Time
}

type Database struct {
	instance *sql.DB
}

func (db *Database) Init() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	var err error

	db.instance, err = sql.Open("postgres", psqlconn)
	CheckError(err)

	// check db
	err = db.instance.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func (db *Database) InsertData(data Transaction) {
	query := `insert into "Transactions"("userId", "type", "symbol", "amount", "price", "date") values('` + strconv.Itoa(data.UserId) + `', '` + data.Type + `', '` + data.Symbol + `', '` + strconv.FormatFloat(data.Amount, 'E', -1, 64) + `', '` + strconv.FormatFloat(data.Price, 'E', -1, 64) + `', '` + data.Date.Format("2006-1-2") + `')`
	_, e := db.instance.Exec(query)
	CheckError(e)
}

func (db *Database) GetAllTransactions() []Transaction {
	query := `SELECT * FROM "Transactions"`
	rows, err := db.instance.Query(query)
	CheckError(err)

	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var transaction Transaction
		err = rows.Scan(&transaction.Id, &transaction.UserId, &transaction.Type, &transaction.Symbol, &transaction.Amount, &transaction.Price, &transaction.Date)
		transactions = append(transactions, transaction)
		CheckError(err)
	}

	CheckError(err)
	return transactions
}

func (db *Database) Close() {
	defer db.instance.Close()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
