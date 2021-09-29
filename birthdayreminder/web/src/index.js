var m = require("mithril")

var listPage = require("./views/listPage")
var updatePage = require("./views/updatePage")
var newPage = require("./views/newPage")
var errPage = require("./views/errPage")
var layout = require("./views/layout")

m.route(document.body, "/list", {
    "/list": {
        render: function () {
            return m(layout, m(listPage))
        }
    },
    "/err": {
        render: function () {
            return m(layout, m(errPage))
        }
    },
    "/addNew": {
        render: function () {
            return m(layout, m(newPage))
        }
    },
    "/edit/:id/:name/:timeType/:timeText/:sendEmail": {
        render: function (vnode) {
            return m(layout, m(updatePage, vnode.attrs))
        }
    },
})