"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g = Object.create((typeof Iterator === "function" ? Iterator : Object).prototype);
    return g.next = verb(0), g["throw"] = verb(1), g["return"] = verb(2), typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
Object.defineProperty(exports, "__esModule", { value: true });
var core_1 = require("https://esm.sh/@tiptap/core");
var extension_underline_1 = require("https://esm.sh/@tiptap/extension-underline");
var extension_bullet_list_1 = require("https://esm.sh/@tiptap/extension-bullet-list");
var extension_text_1 = require("https://esm.sh/@tiptap/extension-text");
var extension_document_1 = require("https://esm.sh/@tiptap/extension-document");
var extension_blockquote_1 = require("https://esm.sh/@tiptap/extension-blockquote");
var extension_code_block_1 = require("https://esm.sh/@tiptap/extension-code-block");
var extension_hard_break_1 = require("https://esm.sh/@tiptap/extension-hard-break");
var extension_heading_1 = require("https://esm.sh/@tiptap/extension-heading");
var extension_horizontal_rule_1 = require("https://esm.sh/@tiptap/extension-horizontal-rule");
var extension_list_item_1 = require("https://esm.sh/@tiptap/extension-list-item");
var extension_ordered_list_1 = require("https://esm.sh/@tiptap/extension-ordered-list");
var extension_paragraph_1 = require("https://esm.sh/@tiptap/extension-paragraph");
var extension_bold_1 = require("https://esm.sh/@tiptap/extension-bold");
var extension_code_1 = require("https://esm.sh/@tiptap/extension-code");
var extension_italic_1 = require("https://esm.sh/@tiptap/extension-italic");
var extension_strike_1 = require("https://esm.sh/@tiptap/extension-strike");
var extension_dropcursor_1 = require("https://esm.sh/@tiptap/extension-dropcursor");
var extension_gapcursor_1 = require("https://esm.sh/@tiptap/extension-gapcursor");
var extension_history_1 = require("https://esm.sh/@tiptap/extension-history");
function chooseThumbnailUrl() {
    var fileInput = document.createElement("input");
    fileInput.type = "file";
    fileInput.accept = "image/png, image/jpeg";
    fileInput.click();
    fileInput.onchange = function (changeEvent) {
        return __awaiter(this, void 0, void 0, function () {
            var file, formData, response, json, textInput, thumbnailPreview, error_1;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        file = changeEvent.target.files[0];
                        if (!file)
                            return [2 /*return*/];
                        formData = new FormData();
                        formData.append("file", file);
                        _a.label = 1;
                    case 1:
                        _a.trys.push([1, 4, , 5]);
                        return [4 /*yield*/, fetch("/www/upload-item-thumbnail", {
                                method: "POST",
                                body: formData,
                            })];
                    case 2:
                        response = _a.sent();
                        if (!response.ok)
                            throw new Error("Upload failed");
                        return [4 /*yield*/, response.json()];
                    case 3:
                        json = _a.sent();
                        textInput = document.getElementById("thumbnailUrl");
                        textInput.value = (new URL("/filemanager/files/".concat(json.filename), location.toString())).toString();
                        thumbnailPreview = document.getElementById("thumbnailPreview");
                        thumbnailPreview.src = textInput.value;
                        return [3 /*break*/, 5];
                    case 4:
                        error_1 = _a.sent();
                        console.error("error uploading file:", error_1);
                        return [3 /*break*/, 5];
                    case 5: return [2 /*return*/];
                }
            });
        });
    };
}
window.chooseThumbnailUrl = chooseThumbnailUrl;
var initWysiwygEditor = function (element) {
    var editorContainer = document.createElement("div");
    editorContainer.setAttribute("id", "".concat(element.id, "-wysiwyg-editor"));
    editorContainer.classList.add("wysiwyg-editor");
    var editorToolbar = document.createElement("div");
    editorToolbar.setAttribute("id", "".concat(element.id, "-wysiwyg-editor-toolbar"));
    editorToolbar.classList.add("wysiwyg-editor-toolbar");
    editorContainer.appendChild(editorToolbar);
    var editorContent = document.createElement("div");
    editorContent.setAttribute("id", "".concat(element.id, "-wysiwyg-editor-content"));
    editorContainer.appendChild(editorContent);
    element.insertAdjacentElement("afterend", editorContainer);
    element.classList.add("hidden");
    var editor = new core_1.Editor({
        element: editorContent,
        extensions: [
            extension_code_block_1.default,
            extension_document_1.default,
            extension_hard_break_1.default,
            extension_horizontal_rule_1.default,
            extension_text_1.default,
            extension_code_1.default,
            extension_dropcursor_1.default,
            extension_gapcursor_1.default,
            extension_history_1.default,
            extension_heading_1.default,
            extension_paragraph_1.default,
            extension_bold_1.default,
            extension_italic_1.default,
            extension_underline_1.default,
            extension_strike_1.default,
            extension_bullet_list_1.default,
            extension_ordered_list_1.default,
            extension_list_item_1.default,
            extension_blockquote_1.default,
        ],
        content: element.value,
        class: "w-full",
        editorProps: {
            attributes: {
                class: "prose max-w-none min-h-32 px-3 py-2"
            },
        },
        onUpdate: function (_a) {
            var editor = _a.editor;
            element.value = editor.getHTML();
        },
    });
    // Add toolbar
    var toolbar = [
        [
            {
                title: "Heading 2",
                content: "<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\"><title>format-header-2</title><path d=\"M3,4H5V10H9V4H11V18H9V12H5V18H3V4M21,18H15A2,2 0 0,1 13,16C13,15.47 13.2,15 13.54,14.64L18.41,9.41C18.78,9.05 19,8.55 19,8A2,2 0 0,0 17,6A2,2 0 0,0 15,8H13A4,4 0 0,1 17,4A4,4 0 0,1 21,8C21,9.1 20.55,10.1 19.83,10.83L15,16H21V18Z\" /></svg>",
                onClick: function () { return editor.chain().focus().toggleHeading({ level: 2 }).run(); }
            },
            {
                title: "Heading 3",
                content: "<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\"><title>format-header-3</title><path d=\"M3,4H5V10H9V4H11V18H9V12H5V18H3V4M15,4H19A2,2 0 0,1 21,6V16A2,2 0 0,1 19,18H15A2,2 0 0,1 13,16V15H15V16H19V12H15V10H19V6H15V7H13V6A2,2 0 0,1 15,4Z\" /></svg>",
                onClick: function () { return editor.chain().focus().toggleHeading({ level: 3 }).run(); }
            },
        ],
        [
            {
                title: "Bold",
                content: "<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\"><title>format-bold</title><path d=\"M13.5,15.5H10V12.5H13.5A1.5,1.5 0 0,1 15,14A1.5,1.5 0 0,1 13.5,15.5M10,6.5H13A1.5,1.5 0 0,1 14.5,8A1.5,1.5 0 0,1 13,9.5H10M15.6,10.79C16.57,10.11 17.25,9 17.25,8C17.25,5.74 15.5,4 13.25,4H7V18H14.04C16.14,18 17.75,16.3 17.75,14.21C17.75,12.69 16.89,11.39 15.6,10.79Z\" /></svg>",
                onClick: function () { return editor.chain().focus().toggleBold().run(); }
            },
            {
                title: "Italic",
                content: "<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\"><title>format-italic</title><path d=\"M10,4V7H12.21L8.79,15H6V18H14V15H11.79L15.21,7H18V4H10Z\" /></svg>",
                onClick: function () { return editor.chain().focus().toggleItalic().run(); }
            },
            {
                title: "Underline",
                content: "<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\"><title>format-underline</title><path d=\"M5,21H19V19H5V21M12,17A6,6 0 0,0 18,11V3H15.5V11A3.5,3.5 0 0,1 12,14.5A3.5,3.5 0 0,1 8.5,11V3H6V11A6,6 0 0,0 12,17Z\" /></svg>",
                onClick: function () { return editor.chain().focus().toggleUnderline().run(); }
            },
            {
                title: "Strikethrough",
                content: "<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\"><title>format-strikethrough-variant</title><path d=\"M7.2 9.8C6 7.5 7.7 4.8 10.1 4.3C13.2 3.3 17.7 4.7 17.6 8.5H14.6C14.6 8.2 14.5 7.9 14.5 7.7C14.3 7.1 13.9 6.8 13.3 6.6C12.5 6.3 11.2 6.4 10.5 6.9C9 8.2 10.4 9.5 12 10H7.4C7.3 9.9 7.3 9.8 7.2 9.8M21 13V11H3V13H12.6C12.8 13.1 13 13.1 13.2 13.2C13.8 13.5 14.3 13.7 14.5 14.3C14.6 14.7 14.7 15.2 14.5 15.6C14.3 16.1 13.9 16.3 13.4 16.5C11.6 17 9.4 16.3 9.5 14.1H6.5C6.4 16.7 8.6 18.5 11 18.8C14.8 19.6 19.3 17.2 17.3 12.9L21 13Z\" /></svg>",
                onClick: function () { return editor.chain().focus().toggleStrike().run(); }
            },
        ],
        [
            {
                title: "Bulleted List",
                content: "<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\"><title>format-list-bulleted</title><path d=\"M7,5H21V7H7V5M7,13V11H21V13H7M4,4.5A1.5,1.5 0 0,1 5.5,6A1.5,1.5 0 0,1 4,7.5A1.5,1.5 0 0,1 2.5,6A1.5,1.5 0 0,1 4,4.5M4,10.5A1.5,1.5 0 0,1 5.5,12A1.5,1.5 0 0,1 4,13.5A1.5,1.5 0 0,1 2.5,12A1.5,1.5 0 0,1 4,10.5M7,19V17H21V19H7M4,16.5A1.5,1.5 0 0,1 5.5,18A1.5,1.5 0 0,1 4,19.5A1.5,1.5 0 0,1 2.5,18A1.5,1.5 0 0,1 4,16.5Z\" /></svg>",
                onClick: function () { return editor.chain().focus().toggleBulletList().run(); }
            },
            {
                title: "Numbered List",
                content: "<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\"><title>format-list-numbered</title><path d=\"M7,13V11H21V13H7M7,19V17H21V19H7M7,7V5H21V7H7M3,8V5H2V4H4V8H3M2,17V16H5V20H2V19H4V18.5H3V17.5H4V17H2M4.25,10A0.75,0.75 0 0,1 5,10.75C5,10.95 4.92,11.14 4.79,11.27L3.12,13H5V14H2V13.08L4,11H2V10H4.25Z\" /></svg>",
                onClick: function () { return editor.chain().focus().toggleOrderedList().run(); }
            },
        ],
        [
            {
                title: "Blockquote",
                content: "<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\"><title>format-quote-open</title><path d=\"M10,7L8,11H11V17H5V11L7,7H10M18,7L16,11H19V17H13V11L15,7H18Z\" /></svg>",
                onClick: function () { return editor.chain().focus().toggleBlockquote().run(); }
            },
        ],
    ];
    // TODO: Check active
    toolbar.forEach(function (toolset) {
        var div = document.createElement("div");
        div.classList.add("toolset");
        toolset.forEach(function (tool) {
            var button = document.createElement("button");
            button.innerHTML = tool.content;
            button.type = "button";
            button.title = tool.title;
            button.onclick = tool.onClick;
            div.appendChild(button);
        });
        editorToolbar.appendChild(div);
    });
};
var selector = "textarea[data-wysiwyg-editor]";
var observer = new MutationObserver(function (mutations) {
    mutations.forEach(function (mutation) {
        mutation.addedNodes.forEach(function (node) {
            var _a;
            if (node.nodeType === 1) { // Ensure it's an element node
                if (node.matches(selector)) {
                    initWysiwygEditor(node);
                }
                else {
                    // Check inside the node for any matching textarea
                    (_a = node.querySelectorAll) === null || _a === void 0 ? void 0 : _a.call(node, selector).forEach(initWysiwygEditor);
                }
            }
        });
    });
});
// Observe the whole document for added elements
observer.observe(document.body, { childList: true, subtree: true });
var elements = document.querySelectorAll(selector);
elements.forEach(initWysiwygEditor);
