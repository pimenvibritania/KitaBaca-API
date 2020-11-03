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
