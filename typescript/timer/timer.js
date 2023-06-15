"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Timer = void 0;
var Timer = /** @class */ (function () {
    function Timer() {
        this.start = 0;
        this.stop = 0;
    }
    Timer.prototype.Start = function () {
        this.start = new Date().getTime();
    };
    Timer.prototype.Stop = function () {
        this.stop = new Date().getTime();
    };
    Timer.prototype.Print = function () {
        var elapsed = this.stop - this.start;
        console.log("Execution time: " + elapsed + "ms");
    };
    return Timer;
}());
exports.Timer = Timer;
