"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.formatFee = exports.formatDuration = exports.padString = void 0;
function padString(n) {
    return n < 10 ? '0' + n.toFixed(0) : n.toFixed(0);
}
exports.padString = padString;
// 格式化时间，将秒转换为时分秒
function formatDuration(sec) {
    const h = Math.floor(sec / 3600);
    sec -= 3600 * h;
    const m = Math.floor(sec / 60);
    sec -= 60 * m;
    const s = Math.floor(sec);
    return {
        hh: padString(h),
        mm: padString(m),
        ss: padString(s),
    };
}
exports.formatDuration = formatDuration;
// 格式化费用，将分转换为元
function formatFee(cents) {
    return (cents / 100).toFixed(2);
}
exports.formatFee = formatFee;
