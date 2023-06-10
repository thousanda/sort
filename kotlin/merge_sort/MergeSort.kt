fun MergeSort(array: MutableList<Int>) {
    val n = array.size

    // 配列の要素数が0, 1の場合は何もしない (ソート済み)
    if (n < 2) {
        return
    }

    // 配列を2つに分割
    val mid = n / 2
    // subList()は参照なので、toMutableList()でコピーを作成
    val left = array.subList(0, mid).toMutableList()
    val right = array.subList(mid, n).toMutableList()

    // それぞれの配列を再帰的にソート
    MergeSort(left)
    MergeSort(right)

    // 2つの配列をマージ
    var il = 0
    var ir = 0
    for (i in 0 until n) {
        if (il == left.size) {  // 左側の配列の要素をすべて使い切った場合
            // 右側の配列の要素を追加
            array[i] = right[ir]
            ir++
        } else if (ir == right.size) {  // 右側の配列の要素をすべて使い切った場合
            // 左側の配列の要素を追加
            array[i] = left[il]
            il++
        } else if (left[il] <= right[ir]) {  // 左側の配列の要素が小さい場合
            // 左側の配列の要素を追加
            array[i] = left[il]
            il++
        } else {  // 右側の配列の要素が小さい場合
            // 右側の配列の要素を追加
            array[i] = right[ir]
            ir++
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
        MergeSort(numbers)
        println("After sorting : $numbers")
        println("---------------------------")
    }
}
