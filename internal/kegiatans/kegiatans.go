package kegiatans

import (
	database "github.com/lutfiahdewi/graphql-go/internal/pkg/db/mssql"
	"github.com/lutfiahdewi/graphql-go/internal/users"

	"log"
)

type Kegiatan struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
	User   *users.User
}

func GetAll() []Kegiatan {
	// Only concern with one table, kegiatans
	// stmt, err := database.Db.Prepare("select id, title, status from Kegiatans")
	stmt, err := database.Db.Prepare("select K.id,K.title, K.status, K.UserID, U.Username from Kegiatans K inner join Users U on K.UserID = U.ID")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var kegiatans []Kegiatan
	var username string
	var id string
	for rows.Next() {
		var kegiatan Kegiatan
		err := rows.Scan(&kegiatan.ID, &kegiatan.Title, &kegiatan.Status, &id, &username)
		if err != nil {
			log.Fatal(err)
		}
		kegiatan.User = &users.User{
			ID:       id,
			Username: username,
		} // changed
		kegiatans = append(kegiatans, kegiatan)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return kegiatans
}

func (kegiatan Kegiatan) Create() int64 {
	statement, err := database.Db.Prepare("INSERT INTO [dbo].[kegiatans] ([Title],[Status],[UserID]) VALUES(@p1,@p2,@p3); select ID = convert(bigint, SCOPE_IDENTITY())")
	print(statement)
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	var lastInsertId int64
	rows, err := statement.Query(kegiatan.Title, kegiatan.Status, kegiatan.User.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&lastInsertId)
		log.Printf("LastInsertId from SCOPE_IDENTITY(): %d\n", lastInsertId)
	}
	/*err = statement.QueryRow(kegiatan.Title, kegiatan.Status, 4).Scan(&lastInsertId)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("LastInsertId from output inserted: %d\n", lastInsertId)*/

	return lastInsertId
}

func Delete(id int64) error {
	statement, err := database.Db.Prepare("DELETE FROM Kegiatans WHERE ID=(@p1)")
	print(statement)
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}

	res, err := statement.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Panic("Error:", err.Error())
		return err
	}
	log.Printf("Affected delete : %d", rows)
	return nil
}

func (kegiatan Kegiatan) Update(newTitle string, newStatus string) error {
	statement, err := database.Db.Prepare("UPDATE FROM Kegiatans SET title=(@p1), title=(@p2) WHERE ID=(@p3)")
	print(statement)
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}

	res, err := statement.Exec(newTitle, newStatus, kegiatan.ID)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Panic("Error:", err.Error())
		return err
	}
	log.Printf("Affected update : %d", rows)
	return nil
}
