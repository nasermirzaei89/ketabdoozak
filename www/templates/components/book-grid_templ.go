// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/nasermirzaei89/ketabdoozak/listing"
	"github.com/nasermirzaei89/ketabdoozak/www/templates/icons"
	. "github.com/nasermirzaei89/ketabdoozak/www/templates/utils"
)

func BookGrid(title string, items []*listing.Item, actions templ.Component) templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"flex flex-col gap-4 px-2\"><div class=\"flex flex-row justify-between pt-2\"><div class=\"text-2xl\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/components/book-grid.templ`, Line: 12, Col: 32}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</div><div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = actions.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(items) > 0 {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<div class=\"grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 gap-4\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for i := range items {
				templ_7745c5c3_Err = BookGridItem(items[i]).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</div><div class=\"flex justify-center hidden\"><button class=\"as-button variant-outlined size-md\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = icons.MdiDownload(6).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "بارگزاری بیشتر</button></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "<div class=\"flex flex-col gap-2 py-4 items-center\"><svg width=\"100\" height=\"100\" viewBox=\"0 0 100 100\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M93.5081 30.614C91.4331 28.489 87.1811 27.695 83.6501 29.535C83.4731 29.312 83.2921 29.106 83.1071 28.922C82.9101 28.725 82.7001 28.549 82.4821 28.383C84.2951 24.811 83.3021 20.355 80.9841 18.09C78.9391 16.093 76.1701 15.976 73.7571 17.785C73.0951 18.282 72.9601 19.223 73.4571 19.886C73.9551 20.549 74.8951 20.682 75.5581 20.186C76.8051 19.248 77.8951 19.267 78.8871 20.237C80.3991 21.714 80.9741 24.854 79.7071 27.173C79.2141 27.081 78.6951 27.031 78.1471 27.031C75.7461 27.031 73.3531 27.928 71.9241 28.465L71.6241 28.577C70.3351 29.055 69.1031 29.557 67.9211 30.083C64.8151 29.271 61.4791 28.681 58.4901 28.681C54.6821 28.681 51.8711 29.628 50.1341 31.496C48.6671 33.074 48.0221 35.232 48.2171 37.913C48.4141 40.609 49.5391 42.55 51.4701 43.632C51.1231 44.259 50.7971 44.895 50.5151 45.546C47.0471 53.559 49.8531 60.4 49.9741 60.689L50.2061 61.241L50.7551 61.481C50.8801 61.535 53.8601 62.815 58.1841 62.815C61.2351 62.815 64.7461 62.201 68.3251 60.33C68.3431 60.366 68.3551 60.405 68.3741 60.441C69.1151 61.882 70.7301 63.655 74.2131 63.909C74.5281 63.931 74.8331 63.943 75.1281 63.943C78.1841 63.943 80.5361 62.708 81.9301 60.372C84.2511 56.481 83.5471 50.08 82.1131 44.493C82.8191 43.014 83.5021 41.426 84.1551 39.711C85.3811 36.491 85.4501 34.061 84.9811 32.239C87.2931 30.994 90.0901 31.411 91.3631 32.713C92.1771 33.546 92.1221 34.594 91.1971 35.825C90.7001 36.488 90.8351 37.429 91.4971 37.926C91.7681 38.128 92.0831 38.226 92.3961 38.226C92.8521 38.226 93.3021 38.019 93.5971 37.626C95.7611 34.738 94.9701 32.111 93.5081 30.614ZM79.3541 58.833C78.4011 60.432 76.8291 61.092 74.4321 60.915C72.1611 60.75 71.4101 59.781 71.0421 59.067C70.9911 58.968 70.9631 58.85 70.9191 58.745C70.5241 57.794 70.4171 56.594 70.5311 55.229C70.6091 54.286 70.7811 53.275 71.0411 52.211C71.4581 50.508 72.0781 48.693 72.7931 46.892C73.3261 45.551 73.9091 44.222 74.5081 42.951C75.3171 41.234 76.1461 39.627 76.8941 38.265L76.9751 38.346C78.1021 41.124 79.2811 44.906 79.9311 48.584C80.6461 52.63 80.7201 56.543 79.3541 58.833ZM51.2101 37.694C51.0761 35.853 51.4421 34.494 52.3311 33.538C53.4781 32.305 55.5511 31.68 58.4911 31.68C60.1241 31.68 61.8901 31.88 63.6671 32.202C67.3831 32.877 71.1371 34.093 73.7951 35.167L73.8511 35.223C72.6061 35.905 70.9811 36.755 69.1681 37.612C67.9241 38.201 66.5921 38.79 65.2301 39.331C63.4541 40.038 61.6331 40.652 59.8991 41.07C58.7271 41.352 57.5991 41.539 56.5501 41.599C56.3131 41.613 56.0711 41.63 55.8431 41.63C54.6971 41.63 53.7871 41.463 53.0861 41.109C51.9231 40.521 51.3361 39.416 51.2101 37.694ZM71.4211 39.864L72.2461 40.689C71.6541 41.897 71.0511 43.21 70.4781 44.578L67.5311 41.631C68.8981 41.057 70.2141 40.453 71.4211 39.864ZM62.3411 43.511L68.6081 49.779C68.1411 51.371 67.7871 52.951 67.6191 54.446L57.6861 44.514C59.1691 44.342 60.7451 43.985 62.3411 43.511ZM60.1551 59.711L51.9571 51.513C52.1381 50.115 52.5081 48.606 53.1471 47.046L64.7341 58.633C63.1511 59.225 61.6101 59.56 60.1551 59.711ZM81.3521 38.641C81.1661 39.128 80.9781 39.602 80.7881 40.067C80.0011 37.763 79.1951 35.849 78.6101 34.679C78.5361 34.399 78.3861 34.148 78.1721 33.952C77.9771 33.738 77.7231 33.588 77.4421 33.515C76.2541 32.914 74.5181 32.182 72.4681 31.466C72.5361 31.441 72.6011 31.413 72.6691 31.388L72.9791 31.272C74.2181 30.807 76.2911 30.03 78.1481 30.03C79.3761 30.03 80.3041 30.361 80.9861 31.043C82.5511 32.609 82.6751 35.165 81.3521 38.641Z\" fill=\"#111827\"></path> <path d=\"M44.7351 66.188C44.1601 66.188 43.6101 65.855 43.3631 65.295C43.0271 64.537 43.3691 63.651 44.1271 63.316C45.4971 62.709 46.3121 62.302 46.3291 62.294C47.0711 61.924 47.9721 62.226 48.3411 62.965C48.7121 63.705 48.4121 64.607 47.6721 64.978C47.6721 64.978 46.8071 65.411 45.3421 66.059C45.1451 66.146 44.9381 66.188 44.7351 66.188Z\" fill=\"#111827\"></path> <path d=\"M15.4262 90.381C14.7982 90.381 14.2132 89.985 14.0032 89.357C13.7402 88.572 14.1642 87.721 14.9502 87.459C16.5582 86.921 18.0632 86.302 19.4242 85.618C20.1622 85.244 21.0672 85.543 21.4382 86.284C21.8102 87.024 21.5122 87.926 20.7722 88.298C19.2842 89.046 17.6462 89.721 15.9022 90.305C15.7442 90.356 15.5832 90.381 15.4262 90.381ZM24.3442 85.732C23.9192 85.732 23.4983 85.553 23.2013 85.205C22.6643 84.575 22.7393 83.628 23.3703 83.091C23.5763 82.915 23.7712 82.738 23.9552 82.56C24.9522 81.589 25.8933 80.603 26.7493 79.63C27.2943 79.008 28.2442 78.946 28.8652 79.494C29.4872 80.041 29.5492 80.988 29.0012 81.61C28.0952 82.642 27.1002 83.685 26.0462 84.711C25.8172 84.934 25.5722 85.155 25.3162 85.374C25.0332 85.614 24.6882 85.732 24.3442 85.732ZM30.9802 78.137C30.6892 78.137 30.3962 78.053 30.1372 77.877C29.4522 77.411 29.2752 76.479 29.7412 75.793C30.6942 74.391 31.4852 72.99 32.0912 71.63C32.4282 70.871 33.3192 70.53 34.0702 70.869C34.8282 71.206 35.1683 72.092 34.8313 72.848C34.1533 74.371 33.2752 75.93 32.2222 77.48C31.9322 77.907 31.4602 78.137 30.9802 78.137ZM25.3792 71.772C24.6082 71.772 23.9522 71.181 23.8862 70.398C23.8172 69.573 24.4292 68.847 25.2552 68.777C26.7412 68.652 28.3592 68.404 30.0632 68.04C30.8692 67.869 31.6703 68.383 31.8433 69.193C32.0163 70.003 31.5003 70.8 30.6903 70.973C28.8633 71.364 27.1202 71.63 25.5072 71.767C25.4642 71.771 25.4222 71.772 25.3792 71.772ZM20.3363 71.707C20.2713 71.707 20.2043 71.703 20.1383 71.694C18.0783 71.422 16.2802 70.834 14.7912 69.945C14.0792 69.52 13.8472 68.599 14.2712 67.888C14.6972 67.175 15.6182 66.945 16.3282 67.368C17.4642 68.046 18.8783 68.501 20.5303 68.719C21.3513 68.827 21.9292 69.581 21.8212 70.403C21.7222 71.158 21.0773 71.707 20.3363 71.707ZM35.2672 69.73C34.6242 69.73 34.0303 69.313 33.8323 68.667C33.8023 68.567 33.7822 68.467 33.7732 68.367C33.5132 68.047 33.3863 67.622 33.4583 67.182C33.5783 66.438 33.6392 65.698 33.6392 64.982C33.6392 64.2 33.5663 63.424 33.4223 62.678C33.2653 61.865 33.7973 61.077 34.6103 60.921C35.4173 60.765 36.2103 61.295 36.3673 62.109C36.5483 63.043 36.6383 64.008 36.6383 64.981C36.6383 65.394 36.6213 65.813 36.5893 66.234C37.5553 65.914 38.5373 65.566 39.5273 65.194C40.2993 64.903 41.1662 65.296 41.4592 66.071C41.7502 66.846 41.3573 67.712 40.5823 68.003C38.9263 68.625 37.2842 69.185 35.7042 69.665C35.5592 69.71 35.4112 69.73 35.2672 69.73ZM12.1812 66.5C11.6302 66.5 11.1002 66.195 10.8372 65.669C10.1402 64.269 9.69125 62.625 9.50825 60.785C9.48825 60.599 9.47325 60.412 9.46125 60.227C9.40725 59.401 10.0342 58.687 10.8612 58.633C11.6992 58.574 12.4012 59.206 12.4552 60.033C12.4652 60.18 12.4762 60.329 12.4922 60.481C12.6402 61.961 12.9872 63.255 13.5222 64.33C13.8912 65.071 13.5892 65.972 12.8482 66.342C12.6342 66.449 12.4052 66.5 12.1812 66.5ZM33.0092 59.24C32.5202 59.24 32.0392 59.001 31.7512 58.559C30.9462 57.325 29.9063 56.218 28.6583 55.27C27.9983 54.769 27.8703 53.829 28.3713 53.168C28.8723 52.507 29.8122 52.379 30.4732 52.881C31.9952 54.037 33.2722 55.397 34.2652 56.92C34.7172 57.614 34.5212 58.544 33.8272 58.996C33.5742 59.162 33.2902 59.24 33.0092 59.24ZM11.7632 56.698C11.5482 56.698 11.3292 56.652 11.1222 56.553C10.3732 56.199 10.0542 55.305 10.4082 54.556C11.2772 52.719 12.7352 51.267 14.6252 50.356C15.3692 49.995 16.2683 50.309 16.6273 51.055C16.9873 51.801 16.6742 52.698 15.9282 53.058C14.6462 53.676 13.7002 54.612 13.1192 55.839C12.8632 56.382 12.3242 56.698 11.7632 56.698ZM25.1362 53.179C24.9672 53.179 24.7953 53.151 24.6273 53.09C23.1533 52.559 21.6092 52.263 20.1602 52.235C19.3322 52.218 18.6743 51.534 18.6903 50.705C18.7073 49.877 19.3862 49.268 20.2202 49.235C21.9912 49.27 23.8662 49.628 25.6442 50.268C26.4232 50.549 26.8282 51.409 26.5462 52.188C26.3262 52.799 25.7502 53.179 25.1362 53.179Z\" fill=\"#111827\"></path> <path d=\"M7.99823 92.136C7.26423 92.136 6.62223 91.596 6.51523 90.848C6.39823 90.028 6.96823 89.268 7.78823 89.151C8.57123 89.04 9.37123 88.901 10.1652 88.739C10.9772 88.575 11.7692 89.096 11.9352 89.909C12.1002 90.721 11.5772 91.513 10.7652 91.679C9.91323 91.852 9.05423 92 8.21223 92.121C8.14123 92.131 8.06823 92.136 7.99823 92.136Z\" fill=\"#111827\"></path></svg><div>کتابی پیدا نشد</div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func BookGridItem(item *listing.Item) templ.Component {
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
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "<a href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 templ.SafeURL = GetItemURL(ctx, item.ID)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var4)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "\" class=\"flex flex-col gap-2 p-2 rounded-md border border-gray-300 shadow-xs relative\"><img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(item.ThumbnailURL)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/components/book-grid.templ`, Line: 45, Col: 30}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "\" class=\"rounded-sm aspect-square w-full bg-gray-500\" alt=\"\"><div class=\"line-clamp-2\" dir=\"auto\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(item.Title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/components/book-grid.templ`, Line: 46, Col: 51}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "</div><div class=\"inline-flex items-center gap-1 text-sm text-gray-700\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.MdiMapMarker(5).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(item.LocationTitle)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/components/book-grid.templ`, Line: 49, Col: 23}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "</div><div class=\"flex gap-1\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i := range item.Types {
			templ_7745c5c3_Err = BookTypeBadge(item.Types[i]).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if item.Status != listing.ItemStatusPublished {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "<div class=\"as-badge absolute top-4 start-4\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = ItemStatusText(item.Status).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "</a>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func BookTypeBadge(typ listing.ItemType) templ.Component {
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
		templ_7745c5c3_Var8 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var8 == nil {
			templ_7745c5c3_Var8 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "<span class=\"as-badge color-primary\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		switch typ {
		case listing.ItemTypeDonate:
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "اهدا")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case listing.ItemTypeExchange:
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "معاوضه")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case listing.ItemTypeLend:
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "امانت")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case listing.ItemTypeSell:
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, "فروشی")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, "</span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func ItemStatusText(status listing.ItemStatus) templ.Component {
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
		templ_7745c5c3_Var9 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var9 == nil {
			templ_7745c5c3_Var9 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		switch status {
		case listing.ItemStatusDraft:
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, "<span>پیش\u200cنویس</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case listing.ItemStatusPendingReview:
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 25, "<span>منتظر بررسی</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case listing.ItemStatusPublished:
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 26, "<span>منتشر شده</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case listing.ItemStatusRejected:
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 27, "<span>رد شده</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case listing.ItemStatusExpired:
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 28, "<span>منقضی شده</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case listing.ItemStatusArchived:
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 29, "<span>آرشیو شده</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case listing.ItemStatusDeleted:
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 30, "<span>حذف شده</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		default:
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 31, "<span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(string(status))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `www/templates/components/book-grid.templ`, Line: 96, Col: 25}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 32, "</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
