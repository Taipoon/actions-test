package main_test

import (
	"database/sql"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"testing"
)

var db *sql.DB

type Record struct {
	ID    int
	At    interface{}
	Name  string
	Value int
}

func TestHakaruHandler(t *testing.T) {

	t.Run("データベースにデータを登録できる", func(t *testing.T) {
		endpoint := "http://127.0.0.1:8081/"
		p := "hakaru"

		u, err := url.Parse(endpoint)
		if err != nil {
			t.Error("URLが正しくありません:", err, endpoint)
		}
		u.Path = path.Join(u.Path, p)

		q := u.Query()
		q.Set("name", "GoUnitTest")
		q.Set("value", strconv.Itoa(1))
		u.RawQuery = q.Encode()

		req, err := http.Get(u.String())
		if err != nil {
			t.Error("HTTP GETリクエストの送信に失敗しました:", err)
		}
		defer req.Body.Close()
	})

	t.Run("データベースから登録したデータを取得できる", func(t *testing.T) {
		t.Skip("あとから修正")
		record := Record{}
		row := db.QueryRow("SELECT `id`, `at`, `name`, `value` FROM `eventlog` WHERE `id` = 1")

		err := row.Scan(&record.ID, &record.At, &record.Name, &record.Value)
		if err != nil {
			t.Error("データ取得に失敗しました:", err)
		}

		if !(record.ID == 1 &&
			record.Name == "GoUnitTest" &&
			record.Value == 1) {
			t.Error("期待していた値と異なります")
		}
	})

}

func TestMain(m *testing.M) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:13306)/hakaru")
	if err != nil {
		log.Fatalln("テスト用のデータベースへ接続できませんでした:", err)
	}
	err = SetupDummyTestTable(db)
	if err != nil {
		log.Fatalln("テストテーブルの作成に失敗しました:", err)
	}
	m.Run()
}

func SetupDummyTestTable(db *sql.DB) error {
	dropTableQuery := "DROP TABLE IF EXISTS `eventlog`"
	createTableQuery := "CREATE TABLE IF NOT EXISTS `eventlog` (`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,`at` datetime DEFAULT NULL,`name` varchar(255) NOT NULL,`value` int(10) unsigned, PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;"
	_, err := db.Exec(dropTableQuery)
	if err != nil {
		return err
	}
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return err
	}
	return nil
}
