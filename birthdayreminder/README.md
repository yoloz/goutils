# birthday reminder

add somebody birthday and near by it you will get an email notification.

## catalog
.
├── conf
├── go.mod
├── main.go
├── README.md
├── start.sh
├── stop.sh
└── web

## web

参考：https://www.mithriljs.net/simple-application.html

* 坑点一

样例的 UserForm.js中` oninput: m.withAttr("value", function(value) {User.current.firstName = value})`
报错：`TypeError: m.withAttr is not a function`

解决可见：https://github.com/MithrilJS/mithril.js/issues/2295

即使用：https://mithril.js.org/simple-application.html
`oninput: function (e) { User.current.lastName = e.target.value }`

* 坑点二

路由连接要添加hashbang,如：

用户编辑中` return m("a.user-list-item", {href: "/edit/" + user.id, oncreate: m.route.link}, user.firstName + " " + user.lastName)`,换成`#!/edit/`

layout中的返回用户列表` m("a[href='/list']", {oncreate: m.route.link}, "Users")`，换成`#!/list`

