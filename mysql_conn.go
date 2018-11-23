package comm

//以下代码一定要放到项目 中
// import (
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/go-xorm/xorm"
// )

// type MySqlConfig struct {
// 	UserName         string
// 	Password         string
// 	Port             string
// 	Host             string
// 	DBName           string
// 	MaxIdleConns     int
// 	MaxOpenConn      int
// 	ShowSqlOnConsole bool
// }

// //var Engine *xorm.Engine

// func (db MySqlConfig) MySqlConn() (*xorm.Engine, error) {
// 	//var err error
// 	log.Println("sql--->", db.UserName+":'"+db.Password+"'@tcp("+db.Host+":"+db.Port+")/"+db.DBName+"?charset=utf8")
// 	Engine, err := xorm.NewEngine("mysql", db.UserName+":"+db.Password+"@tcp("+db.Host+":"+db.Port+")/"+db.DBName+"?charset=utf8")
// 	//b, err := sql.Open("mysql", db.UserName+":"+db.Password+"@tcp("+db.Host+":"+db.Port+")/"+db.DBName+"?charset=utf8")
// 	if err != nil {
// 		log.Println("mySql Conn error:", err)
// 		return nil, err
// 	}
// 	//defer Engine.Close()

// 	Engine.SetMaxIdleConns(db.MaxIdleConns)
// 	Engine.SetMaxOpenConns(db.MaxOpenConn)
// 	Engine.ShowSQL(db.ShowSqlOnConsole)

// 	return Engine, nil
// }

// // func MysqlPing() error {
// // 	return Engine.Ping()
// // }
