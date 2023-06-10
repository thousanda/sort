#include <stdio.h>

// length of test arrays
const int n = 5;

// function prototypes
void selection_sort(int arr[], int n);
void print_array(int arr[], int n);

// main function to test selection sort
int main(void) {
    // list of test arrays
    int arr_list[][n] = {
		{5, 4, 3, 2, 1},
		{1, 2, 3, 4, 5},
		{1, 3, 2, 5, 4},
		{5, 1, 4, 2, 3},
	};

	// test selection sort
	for (int i = 0; i < sizeof(n); i++) {
	    int arr[n];
	    for (int j = 0; j < n; j++) {
            arr[j] = arr_list[i][j];
	    }
        selection_sort(arr, n);
        print_array(arr, n);
    }

    return 0;
}

// selection sort function
void selection_sort(int arr[], int n) {
    int i, j, min, temp;
    for (i = 0; i < n - 1; i++) {
        min = i; // set min to current index
        for (j = i + 1; j < n; j++) {
            if (arr[j] < arr[min]) { // if current element is less than min
                min = j; // set min to current index
            }
        }
        // swap min with current element
        temp = arr[i];
        arr[i] = arr[min];
        arr[min] = temp;
    }
}

// function to print array
void print_array(int arr[], int n) {
    for (int i = 0; i < n; i++) {
        // new line after last element
        if (i == n - 1) {
            printf("%d\n", arr[i]);
            break;
        }
        printf("%d ", arr[i]);
    }
}
