package main

import (
    "fmt"
    "math"
)

type vector struct {
    x float64
    y float64
}

type line struct {
    v1 vector
    v2 vector
}

type triangle struct {
    points [3]vector
}

func main() {
    trianglePoints := [3]vector{{ x: 0,  y: 10 },
                                { x: 30, y: 20 },
                                { x: 20, y: 0 }}
    myTriangle := triangle { points: trianglePoints }
    fmt.Println("Is inside triangle?: ", isInTriagle(myTriangle, vector { x: 25, y: 10 }))
}

func isUnder(line line, point vector) bool {
    leftVector  := line.v1
    rightVector := line.v2
    if (line.v2.x < line.v1.x) {
        leftVector  =  line.v2
        rightVector = line.v1
    }

    xDiff := point.x - leftVector.x
    yDiff := point.y - leftVector.y
    if (leftVector.y > rightVector.y) { // Negative slope
        xDiff = rightVector.x - point.x
        yDiff = point.y - rightVector.y
    }

    xLength := math.Abs(line.v1.x - line.v2.x)
    yLength := math.Abs(line.v1.y - line.v2.y)

    return xDiff / xLength > yDiff / yLength
}

func isInTriagle(triangle triangle, vector vector) bool {
    for i := 0; i < len(triangle.points); i++ {
        point := triangle.points[i]
        nextPoint := triangle.points[(i + 1) % 3]
        thirdPoint := triangle.points[(i + 2) % 3]

        currentLine := line { v1: point, v2: nextPoint }
        isThirdUnder := isUnder(currentLine, thirdPoint)
        if (isThirdUnder  && !isUnder(currentLine, vector) ||
            !isThirdUnder && isUnder(currentLine, vector)) {
            return false
        }
    }
    return true
}
