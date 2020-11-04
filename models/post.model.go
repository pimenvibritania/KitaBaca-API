package models

import (
	"net/http"

	"github.com/pimenvibritania/rest-api/db"
)

type Post struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func FetchAllPost() (Response, error) {
	var obj Post
	var arrobj []Post
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM post"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Title, &obj.Description)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func StorePost(title string, description string) (Response, error) {

	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO post VALUES (DEFAULT,$1,$2) RETURNING id" //returning id , mengembalikan nilai id terakhir di eksekusi

	smst, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	// menampung id terakhir dari query eksekusi
	var id int64

	// eksekusi query dan (Scan) mengambil response return id dari pgsql
	err = smst.QueryRow(title, description).Scan(&id)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": id,
	}

	return res, nil

}
