// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.850
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/nasermirzaei89/ketabdoozak/listing"
	"github.com/nasermirzaei89/ketabdoozak/www/templates/icons"
	. "github.com/nasermirzaei89/ketabdoozak/www/templates/utils"
	"html/template"
	"slices"
)

func NewItemPage(item *listing.Item, locations []*listing.Location, csrfField template.HTML) templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<main><div class=\"container mx-auto py-4\"><form class=\"flex flex-col gap-4 px-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.Raw(string(csrfField)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<h1 class=\"text-2xl font-semibold text-gray-900\">افزودن کتاب</h1><div class=\"flex flex-col sm:flex-row order-last sm:order-none gap-2 justify-end\"><button class=\"as-button variant-filled is-primary size-md\" hx-post=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(string(GetURL(ctx, "items")))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 20, Col: 103}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.MdiSend(6).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "ارسال برای انتشار</button> <button class=\"as-button variant-outlined size-md\" hx-post=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(string(GetURL(ctx, "items?as-draft")))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 24, Col: 103}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.MdiContentSave(6).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "ذخیره پیش\u200cنویس</button> <a class=\"as-button variant-outlined size-md\" href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 templ.SafeURL = GetURL(ctx, "my/items")
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var4)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "\">لغو</a></div><div class=\"flex flex-col md:flex-row gap-4\"><div class=\"grow flex flex-col gap-4\"><div class=\"flex flex-col gap-1\"><label for=\"title\" class=\"text-xl font-semibold\">عنوان کتاب</label><div class=\"flex flex-row border rounded-md border-gray-300 focus-within:border-primary-500 focus-within:ring ring-primary-500 gap-2 px-2 py-1 text-gray-700 w-full\"><input type=\"text\" class=\"text-base w-full pe-2 focus:outline-none focus:ring-0\" id=\"title\" name=\"title\" x-data=\"{ title: &#39;&#39; }\" x-model=\"title\" :dir=\"title.trim() === &#39;&#39; ? &#39;rtl&#39; : &#39;auto&#39;\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(item.Title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 45, Col: 27}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "\"></div></div><div class=\"flex flex-col gap-1\"><label for=\"ownerName\" class=\"text-xl font-semibold\">نام صاحب کتاب</label><div class=\"flex flex-row border rounded-md border-gray-300 focus-within:border-primary-500 focus-within:ring ring-primary-500 gap-2 px-2 py-1 text-gray-700 w-full sm:max-w-80\"><input type=\"text\" class=\"text-base w-full pe-2 focus:outline-none focus:ring-0\" id=\"ownerName\" name=\"ownerName\" placeholder=\"نام شما\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(item.OwnerName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 52, Col: 171}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\"></div></div><div class=\"flex flex-col gap-1\"><label for=\"location\" class=\"text-xl font-semibold\">محل دریافت</label><div class=\"flex flex-row border rounded-md border-gray-300 focus-within:border-primary-500 focus-within:ring ring-primary-500 gap-2 px-2 py-1 text-gray-700 w-full sm:max-w-80\"><select type=\"text\" class=\"text-base w-full pe-2 focus:outline-none focus:ring-0\" id=\"location\" name=\"locationId\"><option value=\"\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if item.LocationID == "" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "></option> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i := range locations {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(locations[i].ID)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 61, Col: 41}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if item.LocationID == locations[i].ID {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, " selected")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, ">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(locations[i].Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 62, Col: 31}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "</select></div></div><div class=\"flex flex-col gap-1\"><div class=\"text-xl font-semibold\">انواع ارائه</div><div class=\"flex flex-col gap-2 px-2\"><label class=\"inline-flex gap-2\"><input type=\"checkbox\" name=\"types\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(string(listing.ItemTypeDonate))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 72, Col: 83}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if slices.Contains(item.Types, listing.ItemTypeDonate) {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, " checked")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "> اهدا</label> <label class=\"inline-flex gap-2\"><input type=\"checkbox\" name=\"types\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(string(listing.ItemTypeExchange))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 76, Col: 85}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if slices.Contains(item.Types, listing.ItemTypeExchange) {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, " checked")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, "> معاوضه</label> <label class=\"inline-flex gap-2\"><input type=\"checkbox\" name=\"types\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(string(listing.ItemTypeLend))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 80, Col: 81}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, "\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if slices.Contains(item.Types, listing.ItemTypeLend) {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 25, " checked")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 26, "> امانت</label> <label class=\"inline-flex gap-2\"><input type=\"checkbox\" name=\"types\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(string(listing.ItemTypeSell))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 84, Col: 81}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 27, "\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if slices.Contains(item.Types, listing.ItemTypeSell) {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 28, " checked")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 29, "> فروشی</label></div></div><div class=\"flex flex-col gap-1\"><div class=\"text-xl font-semibold\">اطلاعات تماس</div><div class=\"flex flex-col gap-2 items-start\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, v := range item.ContactInfo {
			templ_7745c5c3_Err = ContactInfoFormItem(v.Type, v.Value).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 30, "<button type=\"button\" class=\"as-button variant-outlined size-md\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(string(GetURL(ctx, "/new-contact-info-item")))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 98, Col: 63}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 31, "\" hx-swap=\"beforebegin\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.MdiPlus(6).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 32, "افزودن</button></div></div><div class=\"flex flex-col gap-1\"><label for=\"description\" class=\"text-xl font-semibold\">توضیحات</label> <textarea id=\"description\" name=\"description\" class=\"rounded-md border border-black/60 outline-none focus:ring-2 ring-black/30 min-h-10 px-4 py-2 text-base hidden\" data-wysiwyg-editor>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var14 string
		templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(string(item.Description))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 113, Col: 34}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 33, "</textarea></div></div><div class=\"grow sm:max-w-80 flex flex-col gap-4\"><div class=\"flex flex-col gap-1\"><div class=\"text-xl font-semibold\">وضعیت انتشار</div><div>ذخیره نشده</div></div><div class=\"flex flex-col gap-1\"><div class=\"text-xl font-semibold\">تصویر کتاب</div><img src=\"https://placehold.co/300x300?text=No Thumbnail\" id=\"thumbnailPreview\" alt=\"\" class=\"rounded-sm aspect-square w-full bg-gray-500 object-contain\"> <input type=\"hidden\" id=\"thumbnailUrl\" name=\"thumbnailUrl\"><div class=\"flex flex-row gap-2 justify-center\"><button class=\"as-button variant-filled size-md\" type=\"button\" tabindex=\"0\" x-data @click=\"chooseThumbnailUrl()\" @keydown.enter=\"chooseThumbnailUrl()\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.MdiUpload(6).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 34, "بارگزاری تصویر</button> <button class=\"as-button variant-outlined size-md\" type=\"button\" tabindex=\"0\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if item.ThumbnailURL == "" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 35, " disabled")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 36, ">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.MdiDelete(6).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 37, "حذف تصویر</button></div></div></div></div></form></div></main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func ContactInfoFormItem(typ listing.ItemContactInfoType, val string) templ.Component {
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
		templ_7745c5c3_Var15 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var15 == nil {
			templ_7745c5c3_Var15 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 38, "<div class=\"flex flex-row gap-2 w-full\" x-data><label class=\"flex flex-row border rounded-md border-gray-300 focus-within:border-primary-500 focus-within:ring ring-primary-500 gap-2 px-2 py-1 text-gray-700 w-full max-w-80\"><select class=\"text-base w-full pe-2 focus:outline-none focus:ring-0\" name=\"contactInfoType\"><option value=\"phoneNumber\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if typ == listing.ItemContactInfoTypePhoneNumber {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 39, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 40, ">شماره تماس</option> <option value=\"sms\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if typ == listing.ItemContactInfoTypeSMS {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 41, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 42, ">شماره پیامک</option> <option value=\"telegram\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if typ == listing.ItemContactInfoTypeTelegram {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 43, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 44, ">شناسه تلگرام</option> <option value=\"whatsapp\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if typ == listing.ItemContactInfoTypeWhatsapp {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 45, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 46, ">شماره واتساپ</option></select></label> <label class=\"flex flex-row border rounded-md border-gray-300 focus-within:border-primary-500 focus-within:ring ring-primary-500 gap-2 px-2 py-1 text-gray-700 w-full max-w-80\"><input type=\"text\" dir=\"ltr\" class=\"text-base w-full pe-2 focus:outline-none focus:ring-0\" name=\"contactInfoValue\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var16 string
		templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(val)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/new-item-page.templ`, Line: 167, Col: 129}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 47, "\"></label> <button type=\"button\" class=\"as-button variant-outlined size-md\" @click=\"$el.parentElement.remove()\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.MdiDelete(6).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 48, "<span class=\"hidden sm:inline\">حذف</span></button></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
