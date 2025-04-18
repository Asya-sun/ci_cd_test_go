package quadratic

import (
    "testing"
)

func TestSolve(t *testing.T) {
    tests := []struct {
        a, b, c          float64
        expectedX1       float64
        expectedX2       float64
        expectedHasRoots bool
    }{
        {1, -5, 6, 3, 2, true},    // x² -5x +6 =0 → x=3, x=2
        {1, 4, 4, -2, -2, true},    // x² +4x +4=0 → x=-2 (кратный)
        {1, 0, 1, 0, 0, false},     // x² +1=0 → нет корней
        {0, 2, 3, 0, 0, false},     // 0x² +2x +3=0 → ошибка
        {2, 5, -3, 0.5, -3, true},  // 2x² +5x -3=0 → x=0.5, x=-3
    }

    for _, test := range tests {
        x1, x2, hasRoots := Solve(test.a, test.b, test.c)
        
        if hasRoots != test.expectedHasRoots {
            t.Errorf("a=%.1f, b=%.1f, c=%.1f: ожидалось hasRoots=%v, получено %v",
                test.a, test.b, test.c, test.expectedHasRoots, hasRoots)
            continue
        }

        if hasRoots {
            if (x1 != test.expectedX1 || x2 != test.expectedX2) &&
               (x1 != test.expectedX2 || x2 != test.expectedX1) {
                t.Errorf("a=%.1f, b=%.1f, c=%.1f: ожидались корни %.1f и %.1f, получено %.1f и %.1f",
                    test.a, test.b, test.c, test.expectedX1, test.expectedX2, x1, x2)
            }
        }
    }
}