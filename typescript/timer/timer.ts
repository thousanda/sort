export class Timer {
    private start: number;
    private stop: number;

    constructor() {
        this.start = 0;
        this.stop = 0;
    }

    Start() {
        this.start = new Date().getTime();
    }

    Stop() {
        this.stop = new Date().getTime();
    }

    Print() {
        const elapsed = this.stop - this.start;
        console.log("Execution time: " + elapsed + "ms");
    }
}
