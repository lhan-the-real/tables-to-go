package custom

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fraenky8/tables-to-go/internal/cli"
	"github.com/fraenky8/tables-to-go/pkg/database"
	"github.com/fraenky8/tables-to-go/pkg/output"
	"github.com/fraenky8/tables-to-go/pkg/settings"
)

func MySQLTableToStruct(user, password, dbName, schema, tableName, host, port string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		dir = "."
	}

	settings := settings.Settings{
		Verbose:  false,
		VVerbose: false,
		Force:    false,

		DbType:         settings.DBTypeMySQL,
		User:           user,
		Pswd:           password,
		DbName:         dbName,
		Schema:         schema,
		TableName:      tableName,
		Host:           host,
		Port:           port, // left blank, automatically determined if not set
		Socket:         "",
		OutputFilePath: dir,
		OutputFormat:   settings.OutputFormatCamelCase,
		FileNameFormat: settings.FileNameFormatCamelCase,
		PackageName:    "main",
		Prefix:         "",
		Suffix:         "",
		Null:           settings.NullTypeSQL,

		NoInitialism: false,

		TagsNoDb: false,

		TagsMastermindStructable:       false,
		TagsMastermindStructableOnly:   false,
		IsMastermindStructableRecorder: false,

		TagsGorm: false,
	}
	db := database.New(&settings)
	if err := db.Connect(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	writer := output.NewFileWriter(settings.OutputFilePath)

	if err := cli.Run(&settings, db, writer); err != nil {
		fmt.Printf("run error: %v\n", err)
		os.Exit(1)
	}
}
