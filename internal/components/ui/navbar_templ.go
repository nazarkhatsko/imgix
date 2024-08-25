// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package ui

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "imgix/internal/common/constants"

type NavbarOptions struct {
	AppName  string
	PageName string
}

func Navbar(options NavbarOptions) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<nav class=\"navbar navbar-expand-lg bg-body-tertiary\"><div class=\"container\"><!-- Brand --><a class=\"navbar-brand\" href=\"\"><img src=\"/assets/icons/favicon.png\" alt=\"Logo\" width=\"30\" height=\"24\" class=\"d-inline-block align-text-top\"> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(options.AppName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/components/ui/navbar.templ`, Line: 16, Col: 21}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a><!-- Toggle navigation --><button class=\"navbar-toggler\" type=\"button\" data-bs-toggle=\"collapse\" data-bs-target=\"#navbarSupportedContent\" aria-controls=\"navbarSupportedContent\" aria-expanded=\"false\" aria-label=\"Toggle navigation\"><span class=\"navbar-toggler-icon\"></span></button><!-- Menu --><div class=\"collapse navbar-collapse\" id=\"navbarSupportedContent\"><ul class=\"navbar-nav me-auto mb-2 mb-lg-0\"><li class=\"nav-item\"><a class=\"nav-link active\" href=\"/welcome\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if options.PageName == constants.WELLCOME_PAGE_NAME {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" aria-current=\"page\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">Welcome</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/convert\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if options.PageName == constants.CONVERT_PAGE_NAME {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" aria-current=\"page\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">Convert</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/compress\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if options.PageName == constants.COMPRESS_PAGE_NAME {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" aria-current=\"page\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">Compress</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/resize\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if options.PageName == constants.RESIZE_PAGE_NAME {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" aria-current=\"page\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">Resize</a></li></ul></div><!-- Links --><ul class=\"navbar-nav flex-row flex-wrap ms-md-auto\"><li class=\"nav-item col-6 col-lg-auto\"><a class=\"nav-link py-2 px-0 px-lg-2\" href=\"https://github.com/nazarkhatsko/imgix\" target=\"_blank\" rel=\"noopener\"><i class=\"bi bi-github\"></i> <small class=\"d-lg-none ms-2\">GitHub</small></a></li><li class=\"nav-item py-2 py-lg-1 col-12 col-lg-auto\"><div class=\"vr d-none d-lg-flex h-100 mx-lg-2\"></div><hr class=\"d-lg-none my-2\"></li><li class=\"nav-item dropdown\"><button class=\"btn btn-link nav-link py-2 px-0 px-lg-2 dropdown-toggle d-flex align-items-center\" id=\"bd-theme\" type=\"button\" aria-expanded=\"false\" data-bs-toggle=\"dropdown\" data-bs-display=\"static\" aria-label=\"Toggle theme (dark)\"><i class=\"bi bi-moon-stars-fill\"></i> <span class=\"d-lg-none ms-2\" id=\"bd-theme-text\">Toggle theme</span></button><ul class=\"dropdown-menu dropdown-menu-end\" aria-labelledby=\"bd-theme-text\"><li><button type=\"button\" class=\"dropdown-item d-flex align-items-center\" data-bs-theme-value=\"light\" aria-pressed=\"true\"><i class=\"bi bi-sun-fill me-2\"></i> Light <i class=\"bi bi-check2 ms-auto\"></i></button></li><li><button type=\"button\" class=\"dropdown-item d-flex align-items-center active\" data-bs-theme-value=\"dark\" aria-pressed=\"false\"><i class=\"bi bi-moon-stars-fill me-2\"></i> Dark <i class=\"bi bi-check2 ms-auto d-none\"></i></button></li><li><button type=\"button\" class=\"dropdown-item d-flex align-items-center\" data-bs-theme-value=\"auto\" aria-pressed=\"false\"><i class=\"bi bi-circle-half me-2\"></i> Auto <i class=\"bi bi-check2 ms-auto d-none\"></i></button></li></ul></li></ul></div></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
