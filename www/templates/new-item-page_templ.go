// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/nasermirzaei89/ketabdoozak/listing"
import "github.com/nasermirzaei89/ketabdoozak/www/templates/icons"
import . "github.com/nasermirzaei89/ketabdoozak/www/templates/utils"

func NewItemPage() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<main><div class=\"container mx-auto py-4\"><form class=\"flex flex-col gap-4\"><h1 class=\"text-2xl font-semibold text-gray-900\">افزودن کتاب</h1><div class=\"flex flex-row gap-2 justify-end\"><button class=\"as-button variant-filled is-primary size-md\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.MdiSend(6).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "ارسال برای انتشار</button> <button class=\"as-button variant-outlined size-md\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.MdiContentSave(6).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "ذخیره پیش\u200cنویس</button> <a class=\"as-button variant-outlined size-md\" href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 templ.SafeURL = GetURL(ctx, "my/items")
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var2)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "\">لغو</a></div><div class=\"flex flex-row gap-4\"><div class=\"grow flex flex-col gap-4\"><div class=\"flex flex-col gap-1\"><label for=\"title\" class=\"text-xl font-semibold\">عنوان کتاب</label><div class=\"flex flex-row border rounded-md border-gray-300 focus-within:border-primary-500 focus-within:ring ring-primary-500 gap-2 px-2 py-1 text-gray-700 w-full\"><input type=\"text\" class=\"text-base w-full pe-2 focus:outline-none focus:ring-0\" id=\"title\" name=\"title\" placeholder=\"\"></div></div><div class=\"flex flex-col gap-1\"><label for=\"location\" class=\"text-xl font-semibold\">محل دریافت</label><div class=\"flex flex-row border rounded-md border-gray-300 focus-within:border-primary-500 focus-within:ring ring-primary-500 gap-2 px-2 py-1 text-gray-700 w-full max-w-80\"><select type=\"text\" class=\"text-base w-full pe-2 focus:outline-none focus:ring-0\" id=\"location\" name=\"location\"><option value=\"tehran\">تهران</option> <option value=\"mashhad\">مشهد</option></select></div></div><div class=\"flex flex-col gap-1\"><div class=\"text-xl font-semibold\">انواع ارائه</div><div class=\"flex flex-col gap-2 px-2\"><label class=\"inline-flex gap-2\"><input type=\"checkbox\" name=\"type\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(string(listing.ItemTypeDonate))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 48, Col: 82}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\"> اهدا</label> <label class=\"inline-flex gap-2\"><input type=\"checkbox\" name=\"type\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(string(listing.ItemTypeExchange))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 52, Col: 84}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "\"> معاوضه</label> <label class=\"inline-flex gap-2\"><input type=\"checkbox\" name=\"type\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(string(listing.ItemTypeLend))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 56, Col: 80}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "\"> امانت</label> <label class=\"inline-flex gap-2\"><input type=\"checkbox\" name=\"type\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(string(listing.ItemTypeSell))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 60, Col: 80}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "\"> فروشی</label></div></div><div class=\"flex flex-col gap-1\"><div class=\"text-xl font-semibold\">اطلاعات تماس</div><div class=\"flex flex-col gap-2 px-2\"><label class=\"inline-flex gap-2\"><input type=\"checkbox\" name=\"contactInfo\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(string(listing.ItemContactInfoTypePhoneNumber))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 69, Col: 105}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\"> شماره تماس</label> <label class=\"inline-flex gap-2\"><input type=\"checkbox\" name=\"contactInfo\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(string(listing.ItemContactInfoTypeSMS))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 73, Col: 97}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "\"> شماره پیامک</label> <label class=\"inline-flex gap-2\"><input type=\"checkbox\" name=\"contactInfo\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(string(listing.ItemContactInfoTypeTelegram))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 77, Col: 102}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "\"> شناسه تلگرام</label> <label class=\"inline-flex gap-2\"><input type=\"checkbox\" name=\"contactInfo\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(string(listing.ItemContactInfoTypeWhatsapp))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 81, Col: 102}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "\"> شماره واتساپ</label></div></div><div class=\"flex flex-col gap-1\"><label for=\"description\" class=\"text-xl font-semibold\">توضیحات</label> <textarea id=\"description\" name=\"description\" class=\"rounded-md border border-black/60 outline-none focus:ring-2 ring-black/30 min-h-10 px-4 py-2 text-base\" data-wysiwyg-editor></textarea></div></div><div class=\"grow max-w-80 flex flex-col gap-4\"><div class=\"flex flex-col gap-1\"><div class=\"text-xl font-semibold\">وضعیت انتشار</div><div>ذخیره نشده</div></div><div class=\"flex flex-col gap-1\"><div class=\"text-xl font-semibold\">تصویر کتاب</div><img src=\"https://placehold.co/300x300?text=No Thumbnail\" id=\"thumbnailPreview\" alt=\"\" class=\"rounded-sm aspect-square w-full bg-gray-500\"> <input type=\"hidden\" id=\"thumbnailUrl\" name=\"thumbnailUrl\"><div class=\"flex flex-row gap-2\"><div class=\"as-button variant-filled size-md\" role=\"button\" onclick=\"chooseThumbnailUrl()\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.MdiUpload(6).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "بارگزاری تصویر</div><div class=\"as-button variant-outlined size-md\" role=\"button\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.MdiDelete(6).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "حذف تصویر</div></div><script>\n                                function chooseThumbnailUrl() {\n                                    const fileInput = document.createElement(\"input\");\n                                    fileInput.type = \"file\"\n                                    fileInput.accept = \"image/png, image/jpeg\"\n                                    fileInput.click();\n                                    fileInput.onchange = async function(changeEvent) {\n                                        const file = changeEvent.target.files[0];\n                                        if (!file) return;\n\n                                        const formData = new FormData();\n                                        formData.append(\"file\", file);\n\n                                        try {\n                                            const response = await fetch(\"/www/upload-item-thumbnail\", {\n                                                method: \"POST\",\n                                                body: formData,\n                                            });\n\n                                            if (!response.ok) throw new Error(\"Upload failed\");\n\n                                            const json = await response.json();\n\n                                            textInput = document.getElementById(\"thumbnailUrl\");\n\n                                            textInput.value = (new URL(`/filemanager/files/${json.filename}`, location)).toString();\n\n                                            thumbnailPreview = document.getElementById(\"thumbnailPreview\");\n\n                                            thumbnailPreview.src = textInput.value;\n                                        } catch (error) {\n                                            console.error(\"error uploading file:\", error);\n                                        }\n                                    }\n                                }\n                            </script></div></div></div></form></div></main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = InitWysiwygEditor().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
