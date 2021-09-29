var m = require("mithril")
var birthday = require("../models/birthday")


module.exports = {
    oninit: function (vnode) {
        birthday.current.id = vnode.attrs.id;
        birthday.current.name = vnode.attrs.name;
        birthday.current.timeType = vnode.attrs.timeType;
        birthday.current.timeText = vnode.attrs.timeText;
        birthday.current.sendEmail = vnode.attrs.sendEmail;
    },
    view: function () {
        return m("form", { "style": { "margin-left": "30%", "margin-top": "10%" } }, [
            m("input[type=text][name=id][hidden=hidden]", {
                // oninput: function (e) { birthday.current.id = e.target.value },
                value: birthday.current.id
            }),
            m("label.label", "姓名"),
            m("input.inputbox[type=text][name=name][placeholder=姓名][required=required]", {
                oninput: function (e) { birthday.current.name = e.target.value },
                value: birthday.current.name
            }),
            m("label.label", "类型"),
            m("select.inputbox[name='timeType']",
                { value: birthday.current.timeType, onchange() { birthday.current.timeType = this.value } },
                [
                    m("option[value='1']",
                        "公历"
                    ),
                    m("option[value='0']",
                        "农历"
                    )
                ]
            ),
            m("label.label", "日期"),
            m("input.inputbox[type='text'][name='timeText'][placeholder='2-12'][required=required]", {
                oninput: function (e) { birthday.current.timeText = e.target.value },
                value: birthday.current.timeText
            }),
            m("label.label", "通知邮箱"),
            m("input.inputbox[type='text'][name='sendEmail'][placeholder='t1@abc.com;t2@abc.com'][required=required]", {
                oninput: function (e) { birthday.current.sendEmail = e.target.value },
                value: birthday.current.sendEmail
            }),
            // m("button.button[type=button]", { onclick: birthday.update }, "提交"),
            m("button.button[type='submit'][formaction='/update']", "提交")
        ])
    }
}