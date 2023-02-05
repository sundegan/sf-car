"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.routing = void 0;
var routing;
(function (routing) {
    function drving(o) {
        return `/pages/driving/driving?trip_id=${o.trip_id}`;
    }
    routing.drving = drving;
    function lock(o) {
        return `/pages/lock/lock?car_id=${o.car_id}`;
    }
    routing.lock = lock;
    function register(p) {
        const page = '/pages/register/register';
        if (!p) {
            return page;
        }
        return `${page}?redirect=${encodeURIComponent(p.redirectURL)}`;
    }
    routing.register = register;
    function mytrips() {
        return '/pages/mytrips/mytrips';
    }
    routing.mytrips = mytrips;
})(routing = exports.routing || (exports.routing = {}));
