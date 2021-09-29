var m = require("mithril")

module.exports = {
    view: function () {
        return m("div.txtbg404",
            m("div.txtbox",
                [
                    m("p",
                        "对不起，出现异常错误"
                    ),
                    m("p.paddingbox",
                        "请点击以下链接继续"
                    ),
                    m("p",
                        [
                            "》",
                            m("a", { href: "#!/list", oncreate: m.route.link }, "返回网站首页")
                        ]
                    )
                ]
            )
        )
    }
}