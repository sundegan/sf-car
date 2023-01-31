export function padString(n: number) {
    return n < 10 ? '0'+n.toFixed(0) : n.toFixed(0)
}

// 格式化时间，将秒转换为时分秒
export function formatDuration(sec: number) {
    const h = Math.floor(sec/3600)
    sec -= 3600 * h
    const m = Math.floor(sec / 60)
    sec -= 60 * m
    const s = Math.floor(sec)
    return {
        hh: padString(h),
        mm: padString(m),
        ss: padString(s),
    }
}

// 格式化费用，将分转换为元
export function formatFee(cents: number) {
    return (cents / 100).toFixed(2)
}
