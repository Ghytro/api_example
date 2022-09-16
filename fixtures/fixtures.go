package fixtures

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
)

func DefineDBModel(db *pg.DB) {
	basePath := "./"
	fileNames := []string{
		"liquids.sql",
		"users_and_orders.sql",
	}
	for _, fileName := range fileNames {
		file, err := os.Open(basePath + fileName)
		if err != nil {
			log.Fatal(err)
		}
		b := make([]byte, 1)
		var buf bytes.Buffer
		for _, err := file.Read(b); err != nil; _, err = file.Read(b) {
			if b[0] == ';' {
				if _, err = db.Exec(buf.String()); err != nil {
					log.Fatal(err)
				}
				buf.Reset()
			}
			buf.WriteByte(b[0])
		}
		if err != io.EOF {
			log.Fatal(err)
		}
	}
}
