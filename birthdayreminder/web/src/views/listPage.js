var m = require("mithril")
var birthday = require("../models/birthday")

module.exports = {
    oninit: birthday.loadList,
    view: function () {
        return [
            m("table.birthday-list-headline[width='100%'][border='0'][cellspacing='0'][cellpadding='0'][align='center']",
                m("tbody",
                    m("tr",
                        m("td[align='center'][height='60']",
                            "生日备忘录"
                        )
                    )
                )
            ),
            m("div", { "style": { "margin-top": "10px" } },
                m("a.bfont[href='#!/addNew']", { oncreate: m.route.link }, "添加"),
            ),
            m("table.birthday-list[width='100%'][border='0'][cellspacing='1'][cellpadding='4'][bgcolor='#cccccc'][align='center']",
                m("tbody",
                    [
                        m("tr",
                            [
                                m("th.btbg.titfont[width='15%']",
                                    "姓名"
                                ),
                                m("th.btbg.titfont[width='10%']",
                                    "类型"
                                ),
                                m("th.btbg.titfont[width='15%']",
                                    "日期"
                                ),
                                m("th.btbg.titfont[width='50%']",
                                    "发送邮箱"
                                ),
                                m("th.btbg.titfont[width='10%']",
                                    "操作"
                                )
                            ]
                        ),
                        birthday.list.map(function (bd) {
                            return m("tr",
                                [
                                    m("td", bd.name),
                                    m("td", bd.timeType),
                                    m("td", bd.timeText),
                                    m("td", bd.sendEmail),
                                    m("td", [
                                        m("a[href='#!/edit/" + bd.id + "/" + bd.name + "/" + bd.timeType + "/" + bd.timeText + "/"
                                            + bd.sendEmail + "']", { oncreate: m.route.link }, "修改"),
                                        m.trust("&nbsp;"),
                                        m.trust("&nbsp;"),
                                        m("a[href='/delete?id=" + bd.id + "']", { oncreate: m.route.link }, "删除"),
                                    ])
                                ]
                            )
                        })
                    ]
                )
            )
        ]
    }
}