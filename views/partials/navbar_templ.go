// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package partials

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Navbar(fromProtected bool, username string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<nav class=\"navbar fixed top-0 z-10\"><div class=\"navbar-start\"><a hx-swap=\"transition:true\" class=\"btn btn-ghost text-xl\" href=\"/\">TicketTracker</a></div><div class=\"navbar-end\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if fromProtected {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"font-bold mr-8\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(username)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/partials/navbar.templ`, Line: 13, Col: 22}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <a hx-swap=\"transition:true\" class=\"btn btn-ghost text-lg\" href=\"/incidents\">Incidents</a> <button hx-swap=\"transition:true\" hx-post=\"/logout\" hx-confirm=\"Are you sure you want to log out?\" onClick=\"this.addEventListener(&#39;htmx:confirm&#39;, (e) =&gt; {\n\t\t\t\t\t\te.preventDefault()\n\t\t\t\t\t\tSwal.fire({\n\t\t\t\t\t\t\ttitle: &#39;Do you want to perform this action?&#39;,\n\t\t\t\t\t\t\ttext: `${e.detail.question}`,\n\t\t\t\t\t\t\ticon: &#39;warning&#39;,\n\t\t\t\t\t\t\tbackground: &#39;#1D232A&#39;,\n\t\t\t\t\t\t\tcolor: &#39;#A6ADBA&#39;,\n\t\t\t\t\t\t\tshowCancelButton: true,\n\t\t\t\t\t\t\tconfirmButtonColor: &#39;#3085d6&#39;,\n\t\t\t\t\t\t\tcancelButtonColor: &#39;#d33&#39;,\n\t\t\t\t\t\t\tconfirmButtonText: &#39;Yes&#39;\n\t\t\t\t\t\t}).then((result) =&gt; {\n\t\t\t\t\t\t\tif(result.isConfirmed) e.detail.issueRequest(true);\n\t\t\t\t\t\t})\n\t\t\t\t\t})\" hx-target=\"body\" hx-push-url=\"true\" class=\"btn btn-ghost text-lg\">Logout</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a hx-swap=\"transition:true\" class=\"btn btn-ghost text-lg\" href=\"/register\">Register</a> <a hx-swap=\"transition:true\" class=\"btn btn-ghost text-lg\" href=\"/login\">Login</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate