var m = require("mithril")

module.exports = {
    view: function (vnode) {
        return m("main.layout", [
            m("nav.menu", [
                m("a.bfont[href='#!/list']", { oncreate: m.route.link }, "首页")
            ]),
            m("section", vnode.children)
        ])
    }
}