package quadratic

import (
    "math"
)

// Solve решает квадратное уравнение ax² + bx + c = 0
// Возвращает корни и флаг наличия действительных корней
func Solve(a, b, c float64) (x1, x2 float64, hasRoots bool) {
    if a == 0 {
        return 0, 0, false
    }

    discriminant := b*b - 4*a*c
    if discriminant < 0 {
        return 0, 0, false
    }

    sqrtD := math.Sqrt(discriminant)
    x1 = (-b + sqrtD) / (2 * a)
    x2 = (-b - sqrtD) / (2 * a)

    // Для случая одного корня (x1 == x2)
    return x1, x2, true
}