#include <stdio.h>
#include <stdlib.h>

// length of test arrays
const int n = 5;

// function prototypes
void merge_sort(int arr[], int n);
void print_array(int arr[], int n);

// main function to test merge sort
int main(void) {
    // list of test arrays
    int arr_list[][n] = {
		{5, 4, 3, 2, 1},
		{1, 2, 3, 4, 5},
		{1, 3, 2, 5, 4},
		{5, 1, 4, 2, 3},
	};

	// test merge sort on each array
	for (int i = 0; i < sizeof(n); i++) {
	    int arr[n];
	    for (int j = 0; j < n; j++) {
            arr[j] = arr_list[i][j];
	    }
        merge_sort(arr, n);
        print_array(arr, n);
    }

    return 0;
}

// merge sort function
void merge_sort(int arr[], int n) {
    // base case
    if (n < 2) {
        return;
    }

    /* split */
    // find middle index
    int mid = n / 2;
    int len_left = mid;
    int len_right = n - mid;

    // allocate memory for left and right halves
    int *left;
    int *right;
    left = (int*) malloc(mid * sizeof(int));
    right = (int*) malloc((n - mid) * sizeof(int));

    // copy left half
    for (int i = 0; i < len_left; i++) {
        left[i] = arr[i];
    }
    // copy right half
    for (int i = 0; i < len_right; i++) {
        right[i] = arr[mid + i];
    }

    /* sort */
    // sort left half
    merge_sort(left, len_left);
    // sort right half
    merge_sort(right, len_right);

    /* merge */
    // merge left and right halves
    int il = 0;
    int ir = 0;
    for (int i = 0; i < n; i++) {
        // if left half is empty, copy right half
        if (il == len_left) {
            arr[i] = right[ir];
            ir++;
        }
        // if right half is empty, copy left half
        else if (ir == len_right) {
            arr[i] = left[il];
            il++;
        }
        // if left half element is smaller, copy it
        else if (left[il] < right[ir]) {
            arr[i] = left[il];
            il++;
        }
        // if right half element is smaller, copy it
        else {
            arr[i] = right[ir];
            ir++;
        }
    }
    free(left);
    free(right);
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
