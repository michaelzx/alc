package alc_gen

import (
	"bytes"
	"database/sql"
	_ "embed"
	"github.com/jimsmart/schema"
	"github.com/michaelzx/alc/alc_fs"
	"github.com/pkg/errors"
	"go/format"
	"io/ioutil"
	"path/filepath"
)

type Table struct {
	Name    string
	DB      *sql.DB
	Columns []*sql.ColumnType
	Meta    *TableMeta
}

func NewTable(name string, DB *sql.DB) *Table {
	return &Table{Name: name, DB: DB}
}

func (t *Table) Check(g *Gen) error {
	cols, err := schema.Table(t.DB, t.Name)
	if err != nil {
		return err
	}
	t.Columns = cols
	t.Meta, err = NewTableMeta(g.db, g.dbCfg.DbName, t.Name, g.tablePrefix)
	if err != nil {
		return err
	}
	return nil
}

func (t *Table) gen(g *Gen) error {
	if g.genModel {
		err := t.genModelInternal(g)
		if err != nil {
			return err
		}
		err = t.genModel(g)
		if err != nil {
			return err
		}
	}
	return nil
}

//go:embed tpl/model_internal.tpl
var modelInternalTpl string

func (t *Table) genModelInternal(g *Gen) error {
	var doc bytes.Buffer
	// 加载模板字符串
	tmpl, err := loadTpl(modelInternalTpl)
	if err != nil {
		return errors.Wrap(err, "model模板加载失败"+t.Name)
	}
	err = tmpl.Execute(&doc, &map[string]interface{}{
		"ModelTableName": t.Meta.ModelTableName,
		"StructName":     t.Meta.StructName,
		"Fields":         t.Meta.Fields,
		"Imports":        t.Meta.Imports,
	})
	if err != nil {
		return errors.Wrap(err, "模板执行失败"+t.Name)
	}
	reFormatDoc, err := format.Source(doc.Bytes())
	if err != nil {
		return errors.Wrap(err, "格式化go文件失败"+t.Name)
	}
	fileName := t.Name + ".go"
	outPath := filepath.Join(g.rootPath, "internal", "model", "internal", fileName)
	g.logger.Info("生成model internal:" + outPath)
	outDir := filepath.Dir(outPath)
	alc_fs.CreateIfNotExist(outDir)
	ioutil.WriteFile(outPath, reFormatDoc, 0777)
	return nil
}

//go:embed tpl/model.tpl
var modelTpl string

func (t *Table) genModel(g *Gen) error {
	var doc bytes.Buffer
	// 加载模板字符串
	tmpl, err := loadTpl(modelTpl)
	if err != nil {
		return errors.Wrap(err, "model模板加载失败"+t.Name)
	}
	err = tmpl.Execute(&doc, &map[string]interface{}{
		"Package":    `"` + filepath.Join(g.rootPackage, "internal", "model", "internal") + `"`,
		"StructName": t.Meta.StructName,
	})
	if err != nil {
		return errors.Wrap(err, "模板执行失败"+t.Name)
	}
	reFormatDoc, err := format.Source(doc.Bytes())
	if err != nil {
		return errors.Wrap(err, "格式化go文件失败"+t.Name)
	}
	fileName := t.Name + ".go"
	outPath := filepath.Join(g.rootPath, "internal", "model", fileName)
	g.logger.Info("生成model:" + outPath)
	outDir := filepath.Dir(outPath)
	alc_fs.CreateIfNotExist(outDir)
	ioutil.WriteFile(outPath, reFormatDoc, 0777)
	return nil
}
