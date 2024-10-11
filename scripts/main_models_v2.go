package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TableColumn defines the structure for table columns
type TableColumn struct {
	ColumnName       string
	DataType         string
	IsNullable       string
	ColumnDefault    string
	IsPrimaryKey     bool
	IsForeignKey     bool
	RelatedModelName string
	RelatedModelType string
	RelatedModelJson string
	References       string
	GormTag          string
}

// TableTemplate defines the structure for table metadata
type TableTemplate struct {
	TableName string
	Columns   []TableColumn
}

// Map SQL data types to Go data types and determine GORM tags
func mapSQLTypeToGoType(sqlType, isNullable string, isPrimaryKey bool, columnDefault string) (string, string) {
	var goType string
	var gormTagParts []string

	switch sqlType {
	case "integer", "serial":
		goType = "int"
	case "bigint", "bigserial":
		goType = "int64"
	case "smallint":
		goType = "int16"
	case "boolean":
		goType = "bool"
	case "character varying", "varchar", "character", "char", "text", "uuid":
		goType = "string"
	case "timestamp", "timestamp with time zone", "timestamp without time zone", "date", "time":
		goType = "time.Time"
	case "numeric", "decimal", "double precision", "real":
		goType = "float64"
	default:
		goType = "interface{}"
	}

	if isPrimaryKey {
		gormTagParts = append(gormTagParts, "primaryKey")
		if strings.Contains(columnDefault, "nextval") {
			gormTagParts = append(gormTagParts, "autoIncrement")
		}
	}
	if isNullable == "NO" {
		gormTagParts = append(gormTagParts, "not null")
	}
	if columnDefault != "" && !strings.Contains(columnDefault, "nextval") {
		defaultValue := strings.Trim(columnDefault, "'")
		gormTagParts = append(gormTagParts, fmt.Sprintf("default:%s", defaultValue))
	}

	return goType, strings.Join(gormTagParts, ";")
}

// Generate Go model from table structure
func generateModel(table TableTemplate) (string, error) {
	tpl := `package models

type {{.TableName | ToCamelCase}} struct {
{{range .Columns}}	{{.ColumnName | ToCamelCase}} {{.DataType}} ` + "`json:\"{{.ColumnName | ToSnakeCase}}\" gorm:\"{{.GormTag}}\"`" + `
{{if .IsForeignKey}}	{{.RelatedModelName | ToCamelCase}} {{.RelatedModelType}} ` + "`json:\"{{.RelatedModelJson}}\" gorm:\"foreignKey:{{.ColumnName | ToCamelCase}};references:{{.References}}\"`" + `
{{end}}{{end}}}

func ({{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

	funcMap := template.FuncMap{
		"ToCamelCase": func(str string) string {
			words := strings.Split(str, "_")
			for i := range words {
				words[i] = strings.Title(words[i])
			}
			return strings.Join(words, "")
		},
		"ToSnakeCase": func(str string) string {
			return strings.ToLower(str)
		},
		// Función para ajustar los nombres relacionados eliminando el Fk y capitalizando
		"FormatForeignKey": func(columnName string) string {
			if strings.Contains(columnName, "Fk") {
				parts := strings.Split(columnName, "Fk")
				// Toma la parte después de Fk y la pone en CamelCase
				if len(parts) > 1 {
					return strings.Title(parts[1])
				}
			}
			return columnName
		},
	}

	t, err := template.New("model").Funcs(funcMap).Parse(tpl)
	if err != nil {
		return "", err
	}

	for i := range table.Columns {
		// Ajusta el nombre para relaciones de llaves foráneas eliminando "Fk"
		if strings.Contains(strings.ToLower(table.Columns[i].ColumnName), "fk") {
			table.Columns[i].IsForeignKey = true
			// Remueve "Fk" y capitaliza la primera letra restante
			table.Columns[i].RelatedModelName = strings.Replace(table.Columns[i].ColumnName, "Fk", "", 1)
			table.Columns[i].RelatedModelName = strings.Title(table.Columns[i].RelatedModelName)
			table.Columns[i].RelatedModelType = table.Columns[i].RelatedModelName
			table.Columns[i].RelatedModelJson = strings.ToLower(table.Columns[i].RelatedModelName)
			table.Columns[i].References = table.Columns[i].RelatedModelName + "Id"
		}
	}

	var result strings.Builder
	err = t.Execute(&result, table)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}

func main() {
	dsn := "host=34.238.243.7 user=postgres password=1234 dbname=api_zap port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	tables := []string{}
	db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Scan(&tables)

	modelDir := "./models"
	if err := os.MkdirAll(modelDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	errorChannel := make(chan error, len(tables))

	for _, table := range tables {
		wg.Add(1)
		go func(table string) {
			defer wg.Done()

			columns := []TableColumn{}
			primaryKeys := []string{}

			db.Raw("SELECT a.attname FROM pg_index i JOIN pg_attribute a ON a.attnum = ANY(i.indkey) WHERE i.indrelid = ?::regclass AND i.indisprimary", table).Scan(&primaryKeys)

			db.Raw("SELECT column_name, data_type, is_nullable, column_default FROM information_schema.columns WHERE table_name = ?", table).Scan(&columns)

			for i := range columns {
				columns[i].IsPrimaryKey = contains(primaryKeys, columns[i].ColumnName)
				columns[i].DataType, columns[i].GormTag = mapSQLTypeToGoType(columns[i].DataType, columns[i].IsNullable, columns[i].IsPrimaryKey, columns[i].ColumnDefault)
			}

			modelTemplate := TableTemplate{
				TableName: table,
				Columns:   columns,
			}

			model, err := generateModel(modelTemplate)
			if err != nil {
				errorChannel <- err
				return
			}

			fileName := filepath.Join(modelDir, strings.ToLower(table)+".go")
			err = os.WriteFile(fileName, []byte(model), 0644)
			if err != nil {
				errorChannel <- err
				return
			}

			fmt.Printf("Generated model for table %s: %s\n", table, fileName)
		}(table)
	}

	wg.Wait()
	close(errorChannel)

	for err := range errorChannel {
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Todos los modelos fueron generados correctamente.")
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
