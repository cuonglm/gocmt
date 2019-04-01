package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func parseFile(fset *token.FileSet, path, template string) (*ast.File, error) {
	af, err := parser.ParseFile(
		fset,
		path,
		nil,
		parser.ParseComments|parser.AllErrors,
	)

	if err != nil {
		return nil, err
	}

	commentTemplate := commentBase + template

	// Inject first comment to prevent nil comment map
	if len(af.Comments) == 0 {
		af.Comments = []*ast.CommentGroup{{List: []*ast.Comment{{Slash: -1, Text: "// gocmt"}}}}
		defer func() {
			// Remove the injected comment
			af.Comments = af.Comments[1:]
		}()
	}
	cmap := ast.NewCommentMap(fset, af, af.Comments)

	for _, d := range af.Decls {
		switch d.(type) {
		case *ast.FuncDecl:
			fd := d.(*ast.FuncDecl)

			if !fd.Name.IsExported() {
				continue
			}

			addFuncDeclComment(fd, commentTemplate)
			cmap[fd] = []*ast.CommentGroup{fd.Doc}

		case *ast.GenDecl:
			gd := d.(*ast.GenDecl)

			switch gd.Tok {
			case token.CONST, token.VAR:
				if gd.Lparen == token.NoPos && gd.Rparen == token.NoPos{
					vs := gd.Specs[0].(*ast.ValueSpec)
					if !vs.Names[0].IsExported() {
						continue
					}
					addValueSpecComment(gd, vs, commentTemplate)
				} else {
					// if there's a () add comment for each sub entry
					for _, spec := range gd.Specs{
						vs := spec.(*ast.ValueSpec)
						if !vs.Names[0].IsExported() {
							continue
						}
						addParenValueSpecComment( vs, commentTemplate)
						cmap[vs] = []*ast.CommentGroup{vs.Doc}
					}
					continue
				}

			case token.TYPE:
				ts := gd.Specs[0].(*ast.TypeSpec)
				if !ts.Name.IsExported() {
					continue
				}
				addTypeSpecComment(gd, ts, commentTemplate)
			default:
				continue
			}

			cmap[gd] = []*ast.CommentGroup{gd.Doc}

		default:
			continue
		}
	}

	// Rebuild comments
	af.Comments = cmap.Filter(af).Comments()

	return af, nil
}

func addFuncDeclComment(fd *ast.FuncDecl, commentTemplate string) {
	if fd.Doc == nil || strings.TrimSpace(fd.Doc.Text()) == fd.Name.Name {
		text := fmt.Sprintf(commentTemplate, fd.Name)
		pos := fd.Pos() - token.Pos(1)
		fd.Doc = &ast.CommentGroup{List: []*ast.Comment{{Slash: pos, Text: text}}}
	}

}

func addValueSpecComment(gd *ast.GenDecl, vs *ast.ValueSpec, commentTemplate string) {
	if gd.Doc == nil || strings.TrimSpace(gd.Doc.Text()) == vs.Names[0].Name {
		text := fmt.Sprintf(commentTemplate, vs.Names[0].Name)
		pos := gd.Pos() - token.Pos(1)
		gd.Doc = &ast.CommentGroup{List: []*ast.Comment{{Slash: pos, Text: text}}}
	}
}


func addParenValueSpecComment( vs *ast.ValueSpec, commentTemplate string) {
	if vs.Doc == nil || strings.TrimSpace(vs.Doc.Text()) == vs.Names[0].Name {
		commentTemplate = strings.Replace(commentTemplate, commentBase, commentIndentedBase, 1)
		text := fmt.Sprintf(commentTemplate, vs.Names[0].Name)
		pos := vs.Pos() - token.Pos(1)
		vs.Doc = &ast.CommentGroup{List: []*ast.Comment{{Slash: pos, Text: text}}}
	}
}

func addTypeSpecComment(gd *ast.GenDecl, ts *ast.TypeSpec, commentTemplate string) {
	if gd.Doc == nil || strings.TrimSpace(gd.Doc.Text()) == ts.Name.Name {
		text := fmt.Sprintf(commentTemplate, ts.Name.Name)
		pos := gd.Pos() - token.Pos(1)
		gd.Doc = &ast.CommentGroup{List: []*ast.Comment{{Slash: pos, Text: text}}}
	}
}
