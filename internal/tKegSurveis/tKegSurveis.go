package tKegSurveis

import (
	database "github.com/lutfiahdewi/graphql-go/internal/pkg/db/mssql"

	"log"
	"time"
)

/*Available function:
1. GetAll : getting all rows
2. Get() : get one/more row
3. Create : add a row
4. Update : update a row
5. Delete : delete a row*/

type TKegSurvei struct {
	ID               int       `json:"id"`
	Survei_kd        string    `json:"survei_kd"`
	Keg_kd           string    `json:"keg_kd"`
	Status           int       `json:"status"`
	Tgl_buka         time.Time `json:"tgl_buka"`
	Tgl_rek_mulai    time.Time `json:"tgl_rek_mulai"`
	Tgl_rek_selesai  time.Time `json:"tgl_rek_selesai"`
	Tgl_mulai        time.Time `json:"tgl_mulai"`
	Tgl_selesai      time.Time `json:"tgl_selesai"`
	Is_rekrutmen     int    `json:"is_rekrutmen"`
	Is_multi         int    `json:"is_multi"`
	Is_confirm       bool      `json:"Is_confirm"`
	Is_add_indicator bool      `json:"Is_add_indicator"`
	Created_by string `json:"created_by"` //ubah jd graph terhubung tUser
	Created_at string `json:"created_at"`
	Updated_by string `json:"uppdated_by"` //ubah jd graph terhubung tUser
	Updated_at string `json:"updated_at"`
}

func GetAll() []TKegSurvei {
	stmt, err := database.Db.Prepare("SELECT * FROM tKegSurvei") //spesifikasi kolom2nye
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var kegSurveis []TKegSurvei
	for rows.Next() {
		var kegSurvei TKegSurvei
		err := rows.Scan(&kegSurvei.ID, &kegSurvei.Status, &kegSurvei.Tgl_buka) //krg lengkap
		if err != nil {
			log.Fatal(err)
		}
		kegSurveis = append(kegSurveis, kegSurvei)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return kegSurveis
}

func (tKegSurvei TKegSurvei) Create() (lastInsertId int64) {
	statement, err := database.Db.Prepare("INSERT INTO [dbo].[tKegSurvei] ([survei_kd],[keg_kd],[status],[tgl_buka],[tgl_rek_mulai],[tgl_rek_selesai],[tgl_mulai],[tgl_selesai],[is_rekrutmen],[is_multi],[is_confirm],[is_add_indicator],[created_by]) 	 VALUES(@p1,@p2,@p3,@p4,@p5,@p6,@p7,@p8,@p9,@p10,@p11,@p12,@p13); select ID = convert(bigint, SCOPE_IDENTITY())")
	print(statement)
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	rows, err := statement.Query(
		tKegSurvei.Survei_kd,
		tKegSurvei.Keg_kd,
		tKegSurvei.Status,
		(tKegSurvei.Tgl_buka).Format(time.DateTime),
		// nil,
		(tKegSurvei.Tgl_rek_mulai).Format(time.DateTime),
		(tKegSurvei.Tgl_rek_selesai).Format(time.DateTime),
		(tKegSurvei.Tgl_mulai).Format(time.DateTime),
		(tKegSurvei.Tgl_selesai).Format(time.DateTime),
		tKegSurvei.Is_rekrutmen,
		tKegSurvei.Is_multi,
		tKegSurvei.Is_confirm,
		tKegSurvei.Is_add_indicator,
		tKegSurvei.Created_by,
	)
	// log.Printf("MSSQL execute: %s", statement)
	if err != nil {
		log.Fatal("Execution failed:", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&lastInsertId)
		log.Printf("LastInsertId from SCOPE_IDENTITY(): %d\n", lastInsertId)
	}

	return
}

func Update() (TKegSurvei, error) {
	statement, err := database.Db.Prepare("")
	print(statement)
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	var kegSurvei TKegSurvei
	kegSurvei.Updated_at = time.Now().Format(time.DateTime)
	res, err := statement.Exec(kegSurvei.Updated_at)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Panic("Error:", err.Error())
		return kegSurvei, err
	}
	log.Printf("Affected update : %d", rows)
	return kegSurvei, nil

}
