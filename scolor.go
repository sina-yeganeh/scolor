package color

import (
    "fmt"
)

func TextColor(color, text string) {
    switch (color) {
    case "Red":
        fmt.Println(string("\033[31m"), text)
    case "Blue":
        fmt.Println(string("\033[34m"), text)
    case "Green":
        fmt.Println(string("\033[34m"), text)
    case "Cyan":
        fmt.Println(string("\033[34m"), text)
    case "White":
        fmt.Println(string("\033[34m"), text)
    case "Yellow":
        fmt.Println(string("\033[34m"), text)
    }
}
