package quickPerm

func GeneratePermutationsString(data []string) <-chan []string {  
    c := make(chan []string)
    go func(c chan []string) {
        defer close(c) 
        permutateString(c, data)
    }(c)
    return c 
}
func permutateString(c chan []string, inputs []string){  
    output := make([]string, len(inputs))
    copy(output, inputs)
    c <- output

    size := len(inputs)
    p := make([]int, size + 1)
    for i := 0; i < size + 1; i++ {
        p[i] = i;
    }
    for i := 1; i < size;{
        p[i]--
        j := 0
        if i % 2 == 1 {
            j = p[i]
        }
        tmp := inputs[j]
        inputs[j] = inputs[i]
        inputs[i] = tmp
        output := make([]string, len(inputs))
        copy(output, inputs)
        c <- output
        for i = 1; p[i] == 0; i++{
            p[i] = i
        }
    }
}

func GeneratePermutationsInt(data []int) <-chan []int {  
    c := make(chan []int)
    go func(c chan []int) {
        defer close(c) 
        permutateInt(c, data)
    }(c)
    return c 
}
func permutateInt(c chan []int, inputs []int){  
    output := make([]int, len(inputs))
    copy(output, inputs)
    c <- output

    size := len(inputs)
    p := make([]int, size + 1)
    for i := 0; i < size + 1; i++ {
        p[i] = i;
    }
    for i := 1; i < size;{
        p[i]--
        j := 0
        if i % 2 == 1 {
            j = p[i]
        }
        tmp := inputs[j]
        inputs[j] = inputs[i]
        inputs[i] = tmp
        output := make([]int, len(inputs))
        copy(output, inputs)
        c <- output
        for i = 1; p[i] == 0; i++{
            p[i] = i
        }
    }
}


func GeneratePermutationsInt8(data []int8) <-chan []int8 {  
    c := make(chan []int8)
    go func(c chan []int8) {
        defer close(c) 
        permutateInt8(c, data)
    }(c)
    return c 
}
func permutateInt8(c chan []int8, inputs []int8){  
    output := make([]int8, len(inputs))
    copy(output, inputs)
    c <- output

    size := len(inputs)
    p := make([]int, size + 1)
    for i := 0; i < size + 1; i++ {
        p[i] = i;
    }
    for i := 1; i < size;{
        p[i]--
        j := 0
        if i % 2 == 1 {
            j = p[i]
        }
        tmp := inputs[j]
        inputs[j] = inputs[i]
        inputs[i] = tmp
        output := make([]int8, len(inputs))
        copy(output, inputs)
        c <- output
        for i = 1; p[i] == 0; i++{
            p[i] = i
        }
    }
}


func GeneratePermutationsInt16(data []int16) <-chan []int16 {  
    c := make(chan []int16)
    go func(c chan []int16) {
        defer close(c) 
        permutateInt16(c, data)
    }(c)
    return c 
}
func permutateInt16(c chan []int16, inputs []int16){  
    output := make([]int16, len(inputs))
    copy(output, inputs)
    c <- output

    size := len(inputs)
    p := make([]int, size + 1)
    for i := 0; i < size + 1; i++ {
        p[i] = i;
    }
    for i := 1; i < size;{
        p[i]--
        j := 0
        if i % 2 == 1 {
            j = p[i]
        }
        tmp := inputs[j]
        inputs[j] = inputs[i]
        inputs[i] = tmp
        output := make([]int16, len(inputs))
        copy(output, inputs)
        c <- output
        for i = 1; p[i] == 0; i++{
            p[i] = i
        }
    }
}


func GeneratePermutationsInt32(data []int32) <-chan []int32 {  
    c := make(chan []int32)
    go func(c chan []int32) {
        defer close(c) 
        permutateInt32(c, data)
    }(c)
    return c 
}
func permutateInt32(c chan []int32, inputs []int32){  
    output := make([]int32, len(inputs))
    copy(output, inputs)
    c <- output

    size := len(inputs)
    p := make([]int, size + 1)
    for i := 0; i < size + 1; i++ {
        p[i] = i;
    }
    for i := 1; i < size;{
        p[i]--
        j := 0
        if i % 2 == 1 {
            j = p[i]
        }
        tmp := inputs[j]
        inputs[j] = inputs[i]
        inputs[i] = tmp
        output := make([]int32, len(inputs))
        copy(output, inputs)
        c <- output
        for i = 1; p[i] == 0; i++{
            p[i] = i
        }
    }
}


func GeneratePermutationsInt64(data []int64) <-chan []int64 {  
    c := make(chan []int64)
    go func(c chan []int64) {
        defer close(c) 
        permutateInt64(c, data)
    }(c)
    return c 
}
func permutateInt64(c chan []int64, inputs []int64){  
    output := make([]int64, len(inputs))
    copy(output, inputs)
    c <- output

    size := len(inputs)
    p := make([]int, size + 1)
    for i := 0; i < size + 1; i++ {
        p[i] = i;
    }
    for i := 1; i < size;{
        p[i]--
        j := 0
        if i % 2 == 1 {
            j = p[i]
        }
        tmp := inputs[j]
        inputs[j] = inputs[i]
        inputs[i] = tmp
        output := make([]int64, len(inputs))
        copy(output, inputs)
        c <- output
        for i = 1; p[i] == 0; i++{
            p[i] = i
        }
    }
}



func GeneratePermutationsUint(data []uint) <-chan []uint {  
    c := make(chan []uint)
    go func(c chan []uint) {
        defer close(c) 
        permutateUint(c, data)
    }(c)
    return c 
}
func permutateUint(c chan []uint, inputs []uint){  
    output := make([]uint, len(inputs))
    copy(output, inputs)
    c <- output

    size := len(inputs)
    p := make([]int, size + 1)
    for i := 0; i < size + 1; i++ {
        p[i] = i;
    }
    for i := 1; i < size;{
        p[i]--
        j := 0
        if i % 2 == 1 {
            j = p[i]
        }
        tmp := inputs[j]
        inputs[j] = inputs[i]
        inputs[i] = tmp
        output := make([]uint, len(inputs))
        copy(output, inputs)
        c <- output
        for i = 1; p[i] == 0; i++{
            p[i] = i
        }
    }
}



func GeneratePermutationsUint8(data []uint8) <-chan []uint8 {  
    c := make(chan []uint8)
    go func(c chan []uint8) {
        defer close(c) 
        permutateUint8(c, data)
    }(c)
    return c 
}
func permutateUint8(c chan []uint8, inputs []uint8){  
    output := make([]uint8, len(inputs))
    copy(output, inputs)
    c <- output

    size := len(inputs)
    p := make([]int, size + 1)
    for i := 0; i < size + 1; i++ {
        p[i] = i;
    }
    for i := 1; i < size;{
        p[i]--
        j := 0
        if i % 2 == 1 {
            j = p[i]
        }
        tmp := inputs[j]
        inputs[j] = inputs[i]
        inputs[i] = tmp
        output := make([]uint8, len(inputs))
        copy(output, inputs)
        c <- output
        for i = 1; p[i] == 0; i++{
            p[i] = i
        }
    }
}



func GeneratePermutationsUint16(data []uint16) <-chan []uint16 {  
    c := make(chan []uint16)
    go func(c chan []uint16) {
        defer close(c) 
        permutateUint16(c, data)
    }(c)
    return c 
}
func permutateUint16(c chan []uint16, inputs []uint16){  
    output := make([]uint16, len(inputs))
    copy(output, inputs)
    c <- output

    size := len(inputs)
    p := make([]int, size + 1)
    for i := 0; i < size + 1; i++ {
        p[i] = i;
    }
    for i := 1; i < size;{
        p[i]--
        j := 0
        if i % 2 == 1 {
            j = p[i]
        }
        tmp := inputs[j]
        inputs[j] = inputs[i]
        inputs[i] = tmp
        output := make([]uint16, len(inputs))
        copy(output, inputs)
        c <- output
        for i = 1; p[i] == 0; i++{
            p[i] = i
        }
    }
}

func GeneratePermutationsUint32(data []uint32) <-chan []uint32 {  
    c := make(chan []uint32)
    go func(c chan []uint32) {
        defer close(c) 
        permutateUint32(c, data)
    }(c)
    return c 
}
func permutateUint32(c chan []uint32, inputs []uint32){  
    output := make([]uint32, len(inputs))
    copy(output, inputs)
    c <- output

    size := len(inputs)
    p := make([]int, size + 1)
    for i := 0; i < size + 1; i++ {
        p[i] = i;
    }
    for i := 1; i < size;{
        p[i]--
        j := 0
        if i % 2 == 1 {
            j = p[i]
        }
        tmp := inputs[j]
        inputs[j] = inputs[i]
        inputs[i] = tmp
        output := make([]uint32, len(inputs))
        copy(output, inputs)
        c <- output
        for i = 1; p[i] == 0; i++{
            p[i] = i
        }
    }
}

func GeneratePermutationsUint64(data []uint64) <-chan []uint64 {  
    c := make(chan []uint64)
    go func(c chan []uint64) {
        defer close(c) 
        permutateUint64(c, data)
    }(c)
    return c 
}
func permutateUint64(c chan []uint64, inputs []uint64){  
    output := make([]uint64, len(inputs))
    copy(output, inputs)
    c <- output

    size := len(inputs)
    p := make([]int, size + 1)
    for i := 0; i < size + 1; i++ {
        p[i] = i;
    }
    for i := 1; i < size;{
        p[i]--
        j := 0
        if i % 2 == 1 {
            j = p[i]
        }
        tmp := inputs[j]
        inputs[j] = inputs[i]
        inputs[i] = tmp
        output := make([]uint64, len(inputs))
        copy(output, inputs)
        c <- output
        for i = 1; p[i] == 0; i++{
            p[i] = i
        }
    }
}


func GeneratePermutationsFloat32(data []float32) <-chan []float32 {  
    c := make(chan []float32)
    go func(c chan []float32) {
        defer close(c) 
        permutateFloat32(c, data)
    }(c)
    return c 
}
func permutateFloat32(c chan []float32, inputs []float32){  
    output := make([]float32, len(inputs))
    copy(output, inputs)
    c <- output

    size := len(inputs)
    p := make([]int, size + 1)
    for i := 0; i < size + 1; i++ {
        p[i] = i;
    }
    for i := 1; i < size;{
        p[i]--
        j := 0
        if i % 2 == 1 {
            j = p[i]
        }
        tmp := inputs[j]
        inputs[j] = inputs[i]
        inputs[i] = tmp
        output := make([]float32, len(inputs))
        copy(output, inputs)
        c <- output
        for i = 1; p[i] == 0; i++{
            p[i] = i
        }
    }
}


func GeneratePermutationsFloat64(data []float64) <-chan []float64 {  
    c := make(chan []float64)
    go func(c chan []float64) {
        defer close(c) 
        permutateFloat64(c, data)
    }(c)
    return c 
}
func permutateFloat64(c chan []float64, inputs []float64){  
    output := make([]float64, len(inputs))
    copy(output, inputs)
    c <- output

    size := len(inputs)
    p := make([]int, size + 1)
    for i := 0; i < size + 1; i++ {
        p[i] = i;
    }
    for i := 1; i < size;{
        p[i]--
        j := 0
        if i % 2 == 1 {
            j = p[i]
        }
        tmp := inputs[j]
        inputs[j] = inputs[i]
        inputs[i] = tmp
        output := make([]float64, len(inputs))
        copy(output, inputs)
        c <- output
        for i = 1; p[i] == 0; i++{
            p[i] = i
        }
    }
}

func GeneratePermutationsComplex64(data []complex64) <-chan []complex64 {  
    c := make(chan []complex64)
    go func(c chan []complex64) {
        defer close(c) 
        permutateComplex64(c, data)
    }(c)
    return c 
}
func permutateComplex64(c chan []complex64, inputs []complex64){  
    output := make([]complex64, len(inputs))
    copy(output, inputs)
    c <- output

    size := len(inputs)
    p := make([]int, size + 1)
    for i := 0; i < size + 1; i++ {
        p[i] = i;
    }
    for i := 1; i < size;{
        p[i]--
        j := 0
        if i % 2 == 1 {
            j = p[i]
        }
        tmp := inputs[j]
        inputs[j] = inputs[i]
        inputs[i] = tmp
        output := make([]complex64, len(inputs))
        copy(output, inputs)
        c <- output
        for i = 1; p[i] == 0; i++{
            p[i] = i
        }
    }
}

func GeneratePermutationsComplex128(data []complex128) <-chan []complex128 {  
    c := make(chan []complex128)
    go func(c chan []complex128) {
        defer close(c) 
        permutateComplex128(c, data)
    }(c)
    return c 
}
func permutateComplex128(c chan []complex128, inputs []complex128){  
    output := make([]complex128, len(inputs))
    copy(output, inputs)
    c <- output

    size := len(inputs)
    p := make([]int, size + 1)
    for i := 0; i < size + 1; i++ {
        p[i] = i;
    }
    for i := 1; i < size;{
        p[i]--
        j := 0
        if i % 2 == 1 {
            j = p[i]
        }
        tmp := inputs[j]
        inputs[j] = inputs[i]
        inputs[i] = tmp
        output := make([]complex128, len(inputs))
        copy(output, inputs)
        c <- output
        for i = 1; p[i] == 0; i++{
            p[i] = i
        }
    }
}
