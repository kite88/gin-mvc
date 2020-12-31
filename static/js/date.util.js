class DateUtil {
    formatUnix(timestamp) {
        if (timestamp === '' || timestamp === 0 || timestamp === null || timestamp === undefined) {
            return ''
        }
        var date = new Date(timestamp * 1000)
        var year = date.getFullYear(),
            m = date.getMonth() + 1,
            month = m < 10 ? '0' + m : m,
            day = (date.getDate() < 10) ? '0' + date.getDate() : date.getDate(),
            hours = (date.getHours() < 10) ? '0' + date.getHours() : date.getHours(),
            minutes = (date.getMinutes() < 10) ? '0' + date.getMinutes() : date.getMinutes(),
            seconds = (date.getSeconds() < 10) ? '0' + date.getSeconds() : date.getSeconds()
        return year + '/' + month + '/' + day + ' ' + hours + ':' + minutes + ':' + seconds
    }
}
