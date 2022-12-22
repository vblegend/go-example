export class DateUtil {


    public static format(date: Date, fmt: string): string {
        const o = {
            "M+": date.getMonth() + 1,                 //月份 
            "d+": date.getDate(),                    //日 
            "h+": date.getHours(),                   //小时 
            "m+": date.getMinutes(),                 //分 
            "s+": date.getSeconds(),                 //秒 
            "q+": Math.floor((date.getMonth() + 3) / 3), //季度 
            "S": date.getMilliseconds()             //毫秒 
        };
        if (/(y+)/.test(fmt)) {
            fmt = fmt.replace(RegExp.$1, (date.getFullYear() + "").substr(4 - RegExp.$1.length));
        }
        for (const k in o) {
            if (new RegExp("(" + k + ")").test(fmt)) {
                const value = o[<"S">k].toString();
                fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? value : (("00" + value).substr(("" + value).length)));
            }
        }
        return fmt;
    }

}

