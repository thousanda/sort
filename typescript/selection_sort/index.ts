// main function to test the selection sort
const main = (): void => {
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
        selectionSort(arr);
        // print the array
        console.log(arr);
    }
}

// sort the array using selection sort
const selectionSort = (arr: number[]): void => {
    // iterate through the array
    for (let i = 0; i < arr.length; i++) {
        // set the current index as minimum
        let min = i;
        let temp = arr[i];
        // iterate to the right of the current index
        for (let j = i + 1; j < arr.length; j++) {
            // check if the current index is greater than the next index
            if (arr[j] < arr[min]) {
                // set the next index as the minimum
                min = j;
            }
        }
        // swap the current index with the minimum
        arr[i] = arr[min];
        arr[min] = temp;
    }
}

// call the main function
main();
