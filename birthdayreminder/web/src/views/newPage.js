var m = require("mithril")

module.exports = {
    view: function () {
        return m("form[method='post']", { "style": { "margin-left": "30%" } }, [
            m("label.label", "姓名"),
            m("input.inputbox[type=text][name=name][placeholder=姓名][required=required]"),
            m("label.label", "类型"),
            m("select.inputbox[name='timeType']",
                [
                    m("option[value='1']",
                        "公历"
                    ),
                    m("option[value='0']",
                        "农历"
                    )
                ]),
            m("label.label", "日期"),
            m("input.inputbox[type='text'][name='timeText'][placeholder='2-12'][required=required]"),
            m("label.label", "通知邮箱"),
            m("input.inputbox[type='text'][name='sendEmail'][placeholder='t1@abc.com;t2@abc.com'][required=required]"),
            m("br"),
            m("button.button[type='submit'][formaction='/add']", "提交")
        ])
    }
}