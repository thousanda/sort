import { Timer } from "../timer/timer";

const main = (): void => {
    const timer = new Timer();

    // 計測開始
    timer.Start();

    // 計測したい処理
    runMergeSort();

    // 計測終了
    timer.Stop();
    // 計測結果出力
    timer.Print();
};

const runMergeSort = (): void => {
    // create an array of numbers
    const arrList: number[][] = [
        [5, 4, 3, 2, 1],
        [1, 2, 3, 4, 5],
        [3, 2, 1, 5, 4],
        [1, 3, 2, 5, 4],
        [1, 2, 3, 5, 4],
    ];
    for (let i = 0; i < arrList.length; i++) {
        let arr = arrList[i];
        // sort the array
        const sorted = mergeSort(arr);
        // print the array
        console.log(sorted);
    }
}

const mergeSort = (arr: number[]): number[] => {
    // if the array is empty or has one element, return the array
    if (arr.length <= 1) {
        return arr;
    }

    // find the middle index
    const middle = Math.floor(arr.length / 2);
    // split the array into left and right
    const left = arr.slice(0, middle);
    const right = arr.slice(middle);

    // recursively sort the left and right arrays
    const sortedLeft = mergeSort(left);
    const sortedRight = mergeSort(right);

    // new array to store the sorted numbers to return
    let sorted: number[] = [];

    // merge the sorted left and right arrays
    let i_left: number = 0;
    let i_right: number = 0;
    for (let i = 0; i < arr.length; i++) {
        if (i_left >= sortedLeft.length) { // if the left array is empty
            sorted[i] = sortedRight[i_right];
            i_right++;
        } else if (i_right >= sortedRight.length) { // if the right array is empty
            sorted[i] = sortedLeft[i_left];
            i_left++;
        } else if (sortedLeft[i_left] < sortedRight[i_right]) { // if the left array has the smaller number
            sorted[i] = sortedLeft[i_left];
            i_left++;
        } else { // if the right array has the smaller number
            sorted[i] = sortedRight[i_right];
            i_right++;
        }
    }

    return sorted;
}

// call the main function
main();
