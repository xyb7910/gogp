package sql

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type sqlTestSuite struct {
	suite.Suite

	// 配置字段
	driver string
	dsn    string

	// 初始化字段
	db *sql.DB
}

//

func (s *sqlTestSuite) SetupSuite() {
	db, err := sql.Open(s.driver, s.dsn)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = s.db.ExecContext(ctx,
		`CREATE TABLE IF 
    NOT EXISTS user 
	(id int PRIMARY KEY AUTO_INCREMENT, 
	first_name varchar(255) NOT NULL, 
	last_name varchar(255) NOT NULL, 
	age int)`)
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *sqlTestSuite) TestCRUD() {
	t := s.T()
	_, err := sql.Open("mysql", "root:123456789@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//res, err := db.ExecContext(ctx, "INSERT INTO user (first_name, last_name, age) VALUES (?, ?, ?)", "吴", "慧颖", 18)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//affected, err := res.RowsAffected()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//if affected != 1 {
	//	t.Fatal(err)
	//}

	//rows, err := db.QueryContext(ctx, "SELECT id, first_name, last_name, age FROM user LIMIT 1")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//for rows.Next() {
	//	tm := User{}
	//	err = rows.Scan(&tm.Id, &tm.FirstName, &tm.LastName, &tm.Age)
	//	if err != nil {
	//		rows.Close()
	//		t.Fatal(err)
	//	}
	//	assert.Equal(t, "吴", tm.FirstName)
	//	assert.Equal(t, "慧颖", tm.LastName)
	//}
	//rows.Close()

	//res, err := db.ExecContext(ctx, "UPDATE user SET age = ? WHERE id = ?", 20, 1)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//affected, err := res.RowsAffected()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//if affected != 1 {
	//	t.Fatal(err)
	//}

	row := s.db.QueryRowContext(ctx, "SELECT id, first_name, last_name, age FROM user WHERE id = ?", 1)
	if row.Err() != nil {
		t.Fatal(row.Err())
	}
	tm := &User{}
	err = row.Scan(&tm.Id, &tm.FirstName, &tm.LastName, &tm.Age)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "吴", tm.FirstName)
}

func (s *sqlTestSuite) TestTx() {
	t := s.T()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := tx.ExecContext(ctx, "INSERT INTO user (first_name, last_name, age) VALUES (?, ?, ?)", "李", "斯曼", 18)
	if err != nil {
		t.Fatal(err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		t.Fatal(err)
	}
	if affected != 1 {
		t.Fatal(err)
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
	}
}

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func TestMysql(t *testing.T) {
	suite.Run(t, &sqlTestSuite{
		driver: "mysql",
		dsn:    "root:123456789@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local",
	})
}
