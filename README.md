# Variants

Create reusable functions for generating CSS class names based on variants.

Like [cva](https://github.com/joe-bell/cva), but for Go.

> [!WARNING]
> I built this little library as something that I can share across projects as I try out building web applications with Go.
> I wouldn't consider it production ready by any means, but feel free to snoop around the code if you are interested.

## Install

```sh
go get -u github.com/brattonross/variants
```

## Usage

```go
package main

import (
    "github.com/brattonross/variants"
)

type ButtonProps struct {
    Color string
    Size  string
}

var Button = variants.New("font-medium rounded", variants.Options[ButtonProps]{
    Variants: map[string]map[any]string{
        "Color": {
            "primary":   "text-white bg-blue-500 hover:bg-blue-600",
            "secondary": "text-gray-800 bg-gray-200 hover:bg-gray-300",
        },
        "Size": {
            "small":  "h-6 px-2",
            "medium": "h-8 px-3",
            "large":  "h-10 px-4",
        },
    },
	CompoundVariants: map[ButtonProps]string{
        {
            Color: "primary",
            Size:  "medium",
        }: "uppercase",
    },
})

func main() {
    _ = Button(ButtonProps{
        Color: "primary",
        Size:  "medium",
    })
    // => "font-medium rounded text-white bg-blue-500 hover:bg-blue-600 h-8 px-3 uppercase"
}
```

## Prior Art

- Heavily inspired by [cva](https://github.com/joe-bell/cva).
- [Stitches](https://stitches.dev/docs/variants)
- [Vanilla Extract](https://vanilla-extract.style/documentation/api/style-variants/)
