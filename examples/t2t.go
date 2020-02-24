package main

import (
	"fmt"
	"os"

	"github.com/gohouse/converter"
)

var _go_path_ = os.Getenv("GOPATH")

func main() {
	t2t := converter.NewTable2Struct()
    f :=`E:\go\admin_paltform\app\model\`
	err := t2t.SavePath(f).
		Dsn("root:123456@tcp(10.10.35.21:3306)/platform?charset=utf8").
		RealNameMethod("TableName").
		EnableJsonTag(true).
		Config(&converter.T2tConfig{
			StructNameToHump: true,
			TagToLower:       true,
			UcFirstOnly:      true,
			SeperatFile:      true,
		}).
		TagKey("gorm").
		TagCallBack(func(columnName string, isPrimaryKey bool) string {
			var tag string
			if isPrimaryKey {
				tag = "primary_key;"
			}
			tag += columnName
			return tag
		}).
		Run()
	fmt.Println(err)
}
