// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func AboutUsPage() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<main><div class=\"container mx-auto py-4\"><article class=\"flex flex-col gap-4\"><h1 class=\"as-h1 text-gray-900\">درباره ما</h1><div class=\"text-base text-gray-900\"><p>کتابدوزک پروژه\u200cای بود که سال\u200cها پیش تصمیم به ساخت آن داشتم، ولی هیچوقت این تصمیم را عملی نکردم. یه روز بحثی در گروه تلگرامی <a href=\"https://t.me/GolangEngineers/53516\" target=\"_blank\" class=\"underline\">Go Engineers</a> حول محور کتاب بود و قبل از اون هم چندبار مشاهده کردم که دوستانی علاقه\u200cمند به خواندن کتاب\u200cهایی هستند ولی یا امکان دسترسی به آن کتاب خاص را ندارند یا قیمت کتاب\u200cها بسیار بالاست. پس تصمیم گرفتم این پروژه رو اجرایی کنم که بهونه\u200cای باشه برای اشتراک کتاب و دانش بین افراد.</p><p>همچنین تصمیم گرفتم سورس این پروژه رو متن\u200cباز بزارم تا هم امکان مشارکت علاقه\u200cمندان به اون وجود داشته باشه، هم افرادی از بخش\u200cهایی از کد اون استفاده کنند یا الهام بگیرند.</p><p>در حال حاضر هزینه سرور و غیره شخصی پرداخت میشه که قصد دارم لینکی برای حمایت مالی قرار بدم، ولی پروژه رو همیشه رایگان، متن\u200cباز و عام المنفعه نگه\u200cدارم.</p><p>پروژه تازه آغاز شده و قطعا نقص\u200cهایی خواهد داشت. می\u200cتوانید هرگونه پیشنهاد رو به شناسه تلگرام بنده (<a href=\"https://t.me/nasermirzaei89\" dir=\"ltr\" target=\"_blank\" class=\"underline\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs("@nasermirzaei89")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/about-us-page.templ`, Line: 24, Col: 108}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</a>) ارسال بفرمایید، و یا ایرادات رو در مخزن گیت\u200cهاب اعلام بفرمائید.<br>از هرگونه مشارکت در کد هم استقبال میشه.</p><p><br>با احترام<br>ناصر میرزائی</p></div></article></div></main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
