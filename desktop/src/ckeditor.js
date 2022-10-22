const ClassicEditor = require("@ckeditor/ckeditor5-build-classic");

const editor = ClassicEditor;

export default editor;

let token = window.axios.defaults.headers.common["Authorization"];

let baseURL = window.axios.defaults.baseURL;

editor.defaultConfig = {
  toolbar: [
    "heading",
    "bold",
    "italic",
    "link",
    "bulletedList",
    "numberedList",
    "blockQuote",
    "undo",
    "redo",
    "imageTextAlternative",
    "imageUpload",
    "imageStyle:inline",
    "imageStyle:block",
    "imageStyle:side",
  ],
  ckfinder: {
    uploadUrl: baseURL + "/api/v1/image/upload?token=" + token,
  },
};
