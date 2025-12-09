// 示例程序，展示如何简要地使用
// sql 包。
package main

import (
	"database/sql"

	_ "github.com/goinaction/code/chapter3/dbdriver/postgres"
)

// main 是应用程序的入口点。
func main() {
	sql.Open("postgres", "mydb")
}
