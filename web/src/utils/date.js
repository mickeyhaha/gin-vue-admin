// 对Date的扩展，将 Date 转化为指定格式的String
// 月(M)、日(d)、小时(h)、分(m)、秒(s)、季度(q) 可以用 1-2 个占位符，
// 年(y)可以用 1-4 个占位符，毫秒(S)只能用 1 个占位符(是 1-3 位的数字)
// (new Date()).Format("yyyy-MM-dd hh:mm:ss.S") ==> 2006-07-02 08:09:04.423
// (new Date()).Format("yyyy-M-d h:m:s.S")      ==> 2006-7-2 8:9:4.18
Date.prototype.Format = function(fmt) {
    var o = {
        "M+": this.getMonth() + 1, //月份
        "d+": this.getDate(), //日
        "h+": this.getHours(), //小时
        "m+": this.getMinutes(), //分
        "s+": this.getSeconds(), //秒
        "q+": Math.floor((this.getMonth() + 3) / 3), //季度
        "S": this.getMilliseconds() //毫秒
    };
    if (/(y+)/.test(fmt))
        fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
    for (var k in o)
        if (new RegExp("(" + k + ")").test(fmt))
            fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
    return fmt;
}

export function formatTimeToStr(times, pattern) {
    var d = new Date(times).Format("yyyy-MM-dd hh:mm:ss");
    if (pattern) {
        d = new Date(times).Format(pattern);
    }
    return d.toLocaleString();
}

export function DateAdd(interval, number, date) {
    //確保為date類型:
    date = convertToDate(date);
    switch (interval.toLowerCase()) {
        case "y": return new Date(date.setFullYear(date.getFullYear() + number));
        case "m": return new Date(date.setMonth(date.getMonth() + number));
        case "d": return new Date(date.setDate(date.getDate() + number));
        case "w": return new Date(date.setDate(date.getDate() + 7 * number));
        case "h": return new Date(date.setHours(date.getHours() + number));
        case "n": return new Date(date.setMinutes(date.getMinutes() + number));
        case "s": return new Date(date.setSeconds(date.getSeconds() + number));
        case "l": return new Date(date.setMilliseconds(date.getMilliseconds() + number));
    }
};
export function dateFormat(date) {
    //確保為date類型:
    date = convertToDate(date);
    var defyear = parseInt(date.getFullYear());//當前年
    var defmonth = parseInt(date.getMonth() + 1, 10); //當前月
    var defday = date.getDate();//當前日
    var result = "";
    if (defmonth < 10 && defday < 10) {
        result = defyear + '-0' + defmonth + '-0' + defday;
    } else if (defmonth < 10) {
        result = defyear + '-0' + defmonth + '-' + defday;
    } else if (defday < 10) {
        result = defyear + '-' + defmonth + '-0' + defday;
    } else {
        result = defyear + '-' + defmonth + '-' + defday;
    }
    return result;
};
//javascript中定義的replaceAll()
String.prototype.replaceAll = function (s1, s2) {
    return this.replace(new RegExp(s1, "gm"), s2);
};
//將日期類型格式的字符串轉化為日期類型:
export function convertToDate(expr) {
    if (typeof expr == 'string') {
        expr = expr.replaceAll('-', '/');//將字符中的-替換為/,原因是IE或其它瀏覽器不支持-符號的Date.parse()
        return new Date(Date.parse(expr));
    } else {
        return expr;
    }
};