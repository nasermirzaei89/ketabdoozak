// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func HeaderLogo() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<svg width=\"48\" height=\"48\" viewBox=\"0 0 48 48\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M15.9041 6.47046L14.8435 7.53101L17.5505 10.238L18.6111 9.17749L15.9041 6.47046Z\" class=\"fill-slate-600 dark:fill-amber-300\"></path> <path d=\"M17.9695 8.53711C17.592 8.86694 17.2481 9.23081 16.9324 9.62109L17.5505 10.2393L18.6111 9.17871L17.9695 8.53711Z\" class=\"fill-slate-700 dark:fill-amber-400\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M30.5935 6.47046L27.8865 9.17749L28.947 10.238L31.6541 7.53101L30.5935 6.47046Z\" class=\"fill-slate-600 dark:fill-amber-300\"></path> <path d=\"M28.531 8.53418L27.8865 9.17871L28.947 10.2393L29.5652 9.62109C29.2495 9.23117 28.9084 8.86373 28.531 8.53418Z\" class=\"fill-slate-700 dark:fill-amber-400\"></path> <path d=\"M15.6138 15.2016C15.6153 12.2576 17.0735 9.53841 19.438 8.06879C21.8024 6.59901 24.7141 6.60215 27.0757 8.07776C29.4373 9.55274 30.8899 12.2752 30.886 15.2192C26.0512 16.4294 20.9984 16.6029 15.6138 15.2016Z\" class=\"fill-slate-600 dark:fill-amber-300\"></path> <path d=\"M15.6193 14.8982C10.0663 17.5716 7.50177 22.9233 7.5 28.7657C7.5 37.382 14.0398 43.5 23.25 43.5C32.4602 43.5 39 37.382 39 28.7657C38.994 22.9181 36.4182 17.565 30.8553 14.8982C25.7016 16.0056 20.8067 15.7331 15.6193 14.8982Z\" class=\"fill-red-400 dark:fill-red-500\"></path> <path d=\"M19.5 21.0005C19.5 21.3983 19.3419 21.7798 19.0606 22.0611C18.7793 22.3424 18.3978 22.5005 18 22.5005C17.6022 22.5005 17.2206 22.3424 16.9393 22.0611C16.658 21.7798 16.5 21.3983 16.5 21.0005C16.5 20.6027 16.658 20.2212 16.9393 19.9398C17.2206 19.6585 17.6022 19.5005 18 19.5005C18.3978 19.5005 18.7793 19.6585 19.0606 19.9398C19.3419 20.2212 19.5 20.6027 19.5 21.0005Z\" class=\"fill-slate-600 dark:fill-amber-300\"></path> <path d=\"M16.5 28.125C16.5 28.7217 16.263 29.294 15.841 29.716C15.419 30.138 14.8467 30.375 14.25 30.375C13.6533 30.375 13.081 30.138 12.659 29.716C12.237 29.294 12 28.7217 12 28.125C12 27.5283 12.237 26.956 12.659 26.534C13.081 26.112 13.6533 25.875 14.25 25.875C14.8467 25.875 15.419 26.112 15.841 26.534C16.263 26.956 16.5 27.5283 16.5 28.125Z\" class=\"fill-slate-600 dark:fill-amber-300\"></path> <path d=\"M19.125 36C19.125 36.4973 18.9275 36.9742 18.5758 37.3259C18.2242 37.6775 17.7473 37.875 17.25 37.875C16.7527 37.875 16.2758 37.6775 15.9242 37.3259C15.5725 36.9742 15.375 36.4973 15.375 36C15.375 35.5027 15.5725 35.0258 15.9242 34.6741C16.2758 34.3225 16.7527 34.125 17.25 34.125C17.7473 34.125 18.2242 34.3225 18.5758 34.6741C18.9275 35.0258 19.125 35.5027 19.125 36Z\" class=\"fill-slate-600 dark:fill-amber-300\"></path> <path d=\"M27 21.0005C27 21.3983 27.158 21.7798 27.4393 22.0611C27.7206 22.3424 28.1022 22.5005 28.5 22.5005C28.8978 22.5005 29.2793 22.3424 29.5606 22.0611C29.8419 21.7798 30 21.3983 30 21.0005C30 20.6027 29.8419 20.2212 29.5606 19.9398C29.2793 19.6585 28.8978 19.5005 28.5 19.5005C28.1022 19.5005 27.7206 19.6585 27.4393 19.9398C27.158 20.2212 27 20.6027 27 21.0005Z\" class=\"fill-slate-600 dark:fill-amber-300\"></path> <path d=\"M30 28.125C30 28.7217 30.237 29.294 30.659 29.716C31.081 30.138 31.6533 30.375 32.25 30.375C32.8467 30.375 33.419 30.138 33.841 29.716C34.263 29.294 34.5 28.7217 34.5 28.125C34.5 27.5283 34.263 26.956 33.841 26.534C33.419 26.112 32.8467 25.875 32.25 25.875C31.6533 25.875 31.081 26.112 30.659 26.534C30.237 26.956 30 27.5283 30 28.125Z\" class=\"fill-slate-600 dark:fill-amber-300\"></path> <path d=\"M27.375 36C27.375 36.4973 27.5725 36.9742 27.9242 37.3259C28.2758 37.6775 28.7527 37.875 29.25 37.875C29.7473 37.875 30.2242 37.6775 30.5758 37.3259C30.9275 36.9742 31.125 36.4973 31.125 36C31.125 35.5027 30.9275 35.0258 30.5758 34.6741C30.2242 34.3225 29.7473 34.125 29.25 34.125C28.7527 34.125 28.2758 34.3225 27.9242 34.6741C27.5725 35.0258 27.375 35.5027 27.375 36Z\" class=\"fill-slate-600 dark:fill-amber-300\"></path> <path d=\"M26.625 28.125C26.625 28.9704 26.2892 29.7812 25.6914 30.3789C25.0936 30.9767 24.2829 31.3125 23.4375 31.3125C22.5921 31.3125 21.7814 30.9767 21.1836 30.3789C20.5858 29.7812 20.25 28.9704 20.25 28.125C20.25 27.2796 20.5858 26.4688 21.1836 25.8711C21.7814 25.2733 22.5921 24.9375 23.4375 24.9375C24.2829 24.9375 25.0936 25.2733 25.6914 25.8711C26.2892 26.4688 26.625 27.2796 26.625 28.125Z\" class=\"fill-slate-600 dark:fill-amber-300\"></path> <path d=\"M15.9041 6.47168L14.8435 7.53223L15.3005 7.98926C15.7559 7.74832 16.1147 7.35845 16.3171 6.88477L15.9041 6.47168Z\" class=\"fill-slate-700 dark:fill-amber-400\"></path> <path d=\"M30.5935 6.47168L30.1833 6.88184C30.3845 7.35614 30.7423 7.74705 31.197 7.98926L31.6541 7.53223L30.5935 6.47168Z\" class=\"fill-slate-700 dark:fill-amber-400\"></path> <path d=\"M34.1478 6C34.1478 6.24921 34.0987 6.49596 34.0033 6.7262C33.9079 6.95644 33.7682 7.16567 33.5919 7.34189C33.4157 7.51811 33.2064 7.65787 32.9762 7.75323C32.7459 7.8486 32.4992 7.89771 32.2499 7.89771C32.0007 7.89771 31.7539 7.8486 31.5237 7.75323C31.2934 7.65787 31.0841 7.51811 30.9079 7.34189C30.7317 7.16567 30.5919 6.95644 30.4965 6.7262C30.4011 6.49596 30.352 6.24921 30.3521 6C30.352 5.75079 30.4011 5.50405 30.4965 5.2738C30.5919 5.04356 30.7317 4.83433 30.9079 4.65811C31.0841 4.48189 31.2934 4.34214 31.5237 4.24677C31.7539 4.1514 32.0007 4.10229 32.2499 4.10229C32.4992 4.10229 32.7459 4.1514 32.9762 4.24677C33.2064 4.34214 33.4157 4.48189 33.5919 4.65811C33.7682 4.83433 33.9079 5.04356 34.0033 5.2738C34.0987 5.50405 34.1478 5.75079 34.1478 6Z\" class=\"fill-red-400 dark:fill-red-500\"></path> <path d=\"M16.1478 6C16.1478 6.24921 16.0987 6.49596 16.0033 6.7262C15.9079 6.95644 15.7682 7.16567 15.5919 7.34189C15.4157 7.51811 15.2064 7.65787 14.9762 7.75323C14.7459 7.8486 14.4992 7.89771 14.2499 7.89771C14.0007 7.89771 13.7539 7.8486 13.5237 7.75323C13.2934 7.65787 13.0841 7.51811 12.9079 7.34189C12.7317 7.16567 12.5919 6.95644 12.4965 6.7262C12.4011 6.49596 12.352 6.24921 12.3521 6C12.352 5.75079 12.4011 5.50405 12.4965 5.2738C12.5919 5.04356 12.7317 4.83433 12.9079 4.65811C13.0841 4.48189 13.2934 4.34214 13.5237 4.24677C13.7539 4.1514 14.0007 4.10229 14.2499 4.10229C14.4992 4.10229 14.7459 4.1514 14.9762 4.24677C15.2064 4.34214 15.4157 4.48189 15.5919 4.65811C15.7682 4.83433 15.9079 5.04356 16.0033 5.2738C16.0987 5.50405 16.1478 5.75079 16.1478 6Z\" class=\"fill-red-400 dark:fill-red-500\"></path> <path d=\"M15.6198 14.8975C15.3593 15.0229 15.1114 15.1618 14.864 15.2988C14.861 15.3884 14.8464 15.4755 14.8464 15.5654V15.5684C20.7729 17.1108 26.3327 16.918 31.6541 15.586C31.6542 15.4921 31.639 15.4011 31.6365 15.3077C31.3814 15.1664 31.1233 15.0265 30.8542 14.8975C25.7005 16.005 20.8073 15.7324 15.6198 14.8975Z\" class=\"fill-red-500 dark:fill-red-600\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
