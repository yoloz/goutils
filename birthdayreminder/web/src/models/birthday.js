var m = require("mithril")

var birthday = {
    list: [],
    loadList: function () {
        return m.request({
            method: "GET",
            url: "/query",
            withCredentials: true,
        }).then(function (result) {
            birthday.list = result.data
        })
    },
    current: {},
    // update: function () {
    //     return m.request({
    //         method: "POST",
    //         url: "/update",
    //         data: birthday.current,
    //         withCredentials: true,
    //     })
    // },
}

module.exports = birthday