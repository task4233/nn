package mnist

// makeMatrix makes slice (first arugument) x (second argument)
func makeMatrix(row int, column int) [][]float64 {
    resArr := make([][]float64, row)
    for ri:=0; ri<row; ri++ {
        resArr[ri] = make([]float64, column)
    }
    return resArr
}

// add is the function which add first array and second array
func add(lArr [][]float64, rArr [][]float64) [][]float64 {
    for ri:= 0; ri<len(lArr); ri++ {
        for ci:=0; ci<len(lArr[ri]); ci++ {
            lArr[ri][ci] += rArr[ri][ci]
        }
    }
    return lArr
}

// mul is the function which multiply first array and second array
func mul(lArr [][]float64, rArr [][]float64) [][]float64 {
    resArr := makeMatrix(len(lArr), len(rArr[0]))
    for ri:= 0; ri<len(resArr); ri++ {
        for ci:=0; ci<len(resArr[0]); ci++ {
            sum := 0.0
            for rj := 0; rj<len(lArr); rj++ {
                for cj := 0; cj < len(lArr[0]); cj++ {
                    sum += lArr[rj][cj] * rArr[cj][rj]
                }
            }
            resArr[ri][ci] = sum
        }
    }
    return resArr
}
