package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

//https://github.com/zhangshy/netfiles/blob/master/pcshow/treefile.go

type treestate struct {
	Opened   bool `json:"opened"`
	Disabled bool `json:"disabled"`
	Selected bool `json:"selected"`
}

type filetree struct {
	Id       string            `json:"id"`
	Text     string            `json:"text"`
	Icon     string            `json:"icon,omitempty"`
	State    *treestate        `json:"state,omitempty"`
	Children bool              `json:"children,omitempty"`
	Li_attr  map[string]string `json:"li_attr,omitempty"`
	A_attr   map[string]string `json:"a_attr,omitempty"`
}

func treeFileHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	dirpath := r.Form["id"][0]
	log.Println("dirpath:" + dirpath)
	if strings.HasPrefix(dirpath, "C:") {
		log.Println("Windows C do not browse!")
		io.WriteString(w, "[{\"text\" : \"Do not use C:\"}]")
		return
	}
	filetrees := make([]filetree, 0)
	fileinfos, err := ioutil.ReadDir(dirpath)
	if err != nil {
		log.Println("ReadDir error!" + dirpath)
	}
	for _, info := range fileinfos {
		tree := filetree{}
		// 使用文件路径作为id
		tree.Id = filepath.Join(dirpath, info.Name())
		tree.Text = info.Name()
		if info.IsDir() {
			// 是文件夹
			tree.Icon = "jstree-folder"
			tree.Children = true
		} else {
			// Use the rel attribute
			// http://stackoverflow.com/questions/4899520/jstree-types-plugin-does-not-display-custom-icons
			// tree.Li_attr = map[string]string{"rel": "file"}
			tree.Icon = "jstree-file"
			// 设置文件的href
			tree.A_attr = map[string]string{"href": "/download?file=" + tree.Id}
		}

		// 如果文件名为test的话，就选中这个文件
		if info.Name() == "test" {
			state := &treestate{Selected: true}
			tree.State = state
		}
		filetrees = append(filetrees, tree)
	}

	// filetreeData := "[ { \"text\" : \"Root node\", \"children\" : [ \"Child node 1\", \"Child node 2\" ] } ]"
	// io.WriteString(w, filetreeData)

	filetreeData, err := json.Marshal(filetrees)
	io.Copy(w, bytes.NewReader(filetreeData))
}

// <!DOCTYPE html>
// <html>
// <head>
//   <title>文件浏览</title>
//   <link rel="stylesheet" href="/js/jstree-3.2.1/themes/default/style.min.css">
// </head>
// <body>
//   <div id="jstree_demo_div">
//     <ul>
//       <li id="node_1">Root node 1
//         <ul>
//           <li>Child node 1</li>
//           <li><a href="#">Child node 2</a></li>
//         </ul>
//       </li>
//     </ul>
//   </div>
//   <h2>using json</h2>
//   <div id="using_json"></div>
//   <h2>using json2</h2>
//   <div id="using_json2"></div>
//   <h2>using axax</h2>
//   <input type="text" id="dir_input" placeholder="请输入要浏览的文件路径">
//   <input type="button" id="btn_tree" value="确定">
//   <div id="using_ajax"></div>
//   <script src="/js/jquery/1.11.3/jquery.min.js"></script>
//   <script src="/js/jstree-3.2.1/jstree.min.js"></script>
//   <script type="text/javascript">
//     $(document).ready(function() {
//       $('#jstree_demo_div').jstree();
//       $('#jstree_demo_div').on("changed.jstree", function (e, data) {
//         console.log(data.selected);
//       });
//       $('#using_json').jstree({ 'core' : {
//           'data' : [
//              'Simple root node',
//              {
//                'text' : 'Root node 2',
//                'state' : {
//                  'opened' : true,
//                  'selected' : true
//                },
//                'children' : [
//                  { 'text' : 'Child 1' },
//                  'Child 2'
//                ]
//             }
//           ]
//       } });
//       $('#using_json2').jstree({ 'core' : {
//           'data' : [
//              { "id" : "ajson1", "parent" : "#", "text" : "Simple root node" },
//              { "id" : "ajson2", "parent" : "#", "text" : "Root node 2" },
//              { "id" : "ajson3", "parent" : "ajson2", "text" : "Child 1" },
//              { "id" : "ajson4", "parent" : "ajson2", "text" : "Child 2" },
//           ]
//       } });
//       $('#btn_tree').click(function() {
//         root = $('#dir_input').val();
//         $('#root_warn').remove();
//         console.log("root path: " + root);
//         if (root.length<=0) {
//           // 在被选元素之前插入内容
//           $('#dir_input').before("<div id='root_warn'>*请输入要浏览的文件路径*</div>");
//           // alert("请输入要浏览的文件路径");
//           return;
//         }
//         if (root.substring(0, 2)=="C:") {
//           $('#dir_input').before("<div id='root_warn'>*不允许浏览C盘*</div>");
//           return;
//         }
//         $('#using_ajax').jstree("refresh"); // 刷新jstree
//         $('#using_ajax').jstree({
//           'core' : {
//             'data' : {
//               'url' : "/tree_file",
//               "dataType" : "json",
//               'data' : function (node) {
//                 console.log("using_ajax node id: " + node.id + " root:" + root)
//                 if (node.id==='#') {
//                   return { 'id' : root };
//                 } else {
//                   return { 'id' : node.id };
//                 }
//               }
//             }
//           }
//         });
//       });
//     });
//   </script>
// </body>
// </html>
