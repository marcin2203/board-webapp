// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ProfilePage(email string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html lang=\"en\"><header data-testid=\"headerTemplate\"><title>profile</title><script src=\"https://unpkg.com/htmx.org@1.9.11\" integrity=\"sha384-0gxUXCCR8yv9FM2b+U3FDbsKthCI66oH5IA9fHppQq9DDMHuMauqq1ZHBpJxQ0J0\" crossorigin=\"anonymous\"></script><style>\r\n        .column {\r\n        float: left;\r\n        }\r\n\r\n        .left {\r\n        width: 25%;\r\n        }\r\n\r\n        .right {\r\n        width: 75%;\r\n        }\r\n\r\n        .row:after {\r\n        content: \"\";\r\n        display: table;\r\n        clear: both;\r\n        }\r\n        footer {\r\n            width: 100%;\r\n            position:fixed;\r\n            text-align: center;\r\n            padding: 3px;\r\n            bottom:0px;\r\n            background-color: black;\r\n            color: white;\r\n        }   \r\n</style></header><body><header>Profil: ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(email)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates\profile.templ`, Line: 39, Col: 23}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</header><hr><main><div class=\"column left\">Menu-col1\r<hr><table><tr><button type=\"button\">POKAŻ DANE</button></tr><tr><button type=\"button\">ZMIEN HASLO</button></tr><tr><button type=\"button\">ZMIEN EMAIL</button></tr><tr><button type=\"button\">USUN</button></tr></table></div><div class=\"column right\">Menu-col2\r</div></main><hr><footer><p>Author: Hege Refsnes<br><a>hege@example.com</a></p></footer></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
