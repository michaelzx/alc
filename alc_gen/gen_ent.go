package alc_gen

import (
	"alchemy/alc/alc_fs"
	"alchemy/alc/alc_print"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"go/format"
	"io/ioutil"
	"path/filepath"
)

type baseTplData struct {
}

type entTplData struct {
	TableName   string
	StructName  string
	PackageName string
	Imports     []string
	Fields      []string
}

func (g *Generator) genEnt() {
	entDir := filepath.Join(g.OutDir, g.PackageEnt)
	alc_fs.CreateIfNotExist(entDir)
	tmpl := g.loadTpl(g.TplEnt)
	for tableName, meta := range g.tableMetaMap {
		var doc bytes.Buffer
		err := tmpl.Execute(&doc, &entTplData{
			TableName:   meta.TableName,
			StructName:  meta.StructName,
			PackageName: g.PackageEnt,
			Fields:      meta.Fields,
			Imports:     meta.EntImports,
		})
		if err != nil {
			panic(errors.Wrap(err, "模板执行失败"+tableName))
		}

		reFormatDoc, err := format.Source(doc.Bytes())
		if err != nil {
			panic(errors.Wrap(err, "格式化go文件失败"+tableName))
			return
		}
		fileName := tableName + "_ent.go"
		if g.FileNameProcessor != nil {
			fileName = g.FileNameProcessor(fileName)
		}
		outPath := filepath.Join(g.OutDir, g.PackageEnt, fileName)
		ioutil.WriteFile(filepath.Join(outPath), reFormatDoc, 0777)
		alc_print.Green(fmt.Sprintf("成功生成 : %s", filepath.Join(g.rootPath, outPath)))
	}

}
