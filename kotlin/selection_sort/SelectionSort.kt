fun selectionSort(array: MutableList<Int>) {
    val n = array.size
    for (i in 0 until n - 1) {
        // 最小値を見つけるための初期インデックスを設定
        var minIndex = i
        for (j in i + 1 until n) {
            // より小さい値を見つけたら、最小インデックスを更新
            if (array[j] < array[minIndex]) {
                minIndex = j
            }
        }

        // 最小値を現在のインデックス（i）に交換
        if (minIndex != i) {
            val tmp = array[minIndex]
            array[minIndex] = array[i]
            array[i] = tmp
        }
    }
}

fun main() {
    val numbersList = listOf(
        mutableListOf(64, 34, 25, 12, 22, 11, 90),
        mutableListOf(5, 8, 1, 3, 7, 9, 2),
        mutableListOf(99, 45, 38, 22, 12, 19, 91, 88)
    )
    for (numbers in numbersList) {
        println("Before sorting : $numbers")
        selectionSort(numbers)
        println("After sorting : $numbers")
        println("---------------------------")
    }
}
