/**
 * 时间日期相关操作
 */


/**
 * 时间格式化
 * 将 2018-09-23T11:54:16.000+0000 格式化成 2018/09/23 11:54:16
 * @param datetime 国际化日期格式
 */
export function format (datetime) {
  return formatWithSeperator(datetime, "/", ":");
}

/**
 * 时间格式化
 * 将 2018-09-23T11:54:16.000+0000 格式化成类似 2018/09/23 11:54:16
 * 可以指定日期和时间分隔符
 * @param datetime 国际化日期格式
 */
export function formatWithSeperator (datetime, dateSeprator, timeSeprator) {
  if (datetime != null) {
    const dateMat = new Date(datetime);
    const year = dateMat.getFullYear();
    const month = dateMat.getMonth() + 1;
    const day = dateMat.getDate();
    const hh = dateMat.getHours();
    const mm = dateMat.getMinutes();
    const ss = dateMat.getSeconds();
    const timeFormat = year + dateSeprator + month + dateSeprator + day + " " + hh + timeSeprator + mm + timeSeprator + ss;
    return timeFormat;
  }
}

export function strTime() {
	var today = new Date();
	var month = today.getMonth() + 1;
	month = month < 10 ? '0'+month : month;
	var day = today.getDate() < 10 ? '0'+today.getDate() : today.getDate();
	var hours = today.getHours() < 10 ? '0'+today.getHours() : today.getHours();
	var mins = today.getMinutes() < 10 ? '0'+today.getMinutes() : today.getMinutes();
	var secs = today.getSeconds() < 10 ? '0'+today.getSeconds() : today.getSeconds();
	var len = 2;
	var $chars = 'ABCDEFGHJKMNPQRSTWXYZ';
	var maxPos = $chars.length;
	var wrd = '';
	for (var i = 0; i < len; i++) {
		wrd += $chars.charAt(Math.floor(Math.random() * maxPos));
	}

	return '' + today.getFullYear() + month + day + hours + mins + secs + wrd;
}