package clear

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"wayne/clear/dal/query"
)

func TestConfigMap(t *testing.T) {
	rep := query.ConfigMapTemplate
	list, err := rep.Where(rep.ID.Gt(0)).Find()

	if err != nil {
		panic(err)
	}

	for _, item := range list {

		filePath := fmt.Sprintf("./configmap/%d.json", item.ID)

		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("文件打开失败", err)
		}

		//及时关闭file句柄
		defer file.Close()
		//写入文件时，使用带缓存的 *Writer
		write := bufio.NewWriter(file)

		write.WriteString(item.Template)
		write.Flush()

	}
}

func TestUpdateConfigMap(t *testing.T) {

	rep := query.ConfigMapTemplate

	root := "./configmap"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		name := strings.Replace(info.Name(), ".json", "", -1)

		id, err := strconv.ParseInt(name, 10, 64)

		if err != nil {
			panic(err)
		}

		if id > 0 {
			txt, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			tpl := string(txt)
			res, err := rep.Where(rep.ID.Eq(id)).Update(rep.Template, tpl)
			if err != nil {
				fmt.Printf("ID:%d \n", id)
				panic(err)
			}

			fmt.Printf("id=%d,update:%d \n", id, res.RowsAffected)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

}
