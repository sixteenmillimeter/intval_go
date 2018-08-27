package main

import (
    //"flag"
    "fmt"
    "image"
    "image/color"
    "image/draw"
    _ "image/jpeg"
    _ "image/png"
    //"golang.org/x/image/tiff"
    //"os"
    "os/exec"
    "log"
    //"path/filepath"
    "github.com/google/gxui"
    //"github.com/google/gxui/math"
    "github.com/google/gxui/drivers/gl"
    "github.com/google/gxui/themes/dark"
    "github.com/google/gxui/samples/flags"
)

func ffmpeg_check() {
    cmd := exec.Command("ffmpeg", "-h")
    out, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatalf("cmd.Run() failed with %s\n", err)
    }
    fmt.Printf("combined out:\n%s\n", string(out))
}

func ffmpeg_info () {}

func ffmpeg_export () {
    //ffmpeg -i input -compression_algo raw -pix_fmt rgb24 output.tiff
}

func appMain(driver gxui.Driver) {
    width, height := 640, 480
    m := image.NewRGBA(image.Rect(0, 0, width, height))
    gray := color.RGBA{200, 200, 200, 255}
    draw.Draw(m, m.Bounds(), &image.Uniform{gray}, image.ZP, draw.Src)

    // The themes create the content. Currently only a dark theme is offered for GUI elements.
    theme := dark.CreateTheme(driver)
    img := theme.CreateImage()
    window := theme.CreateWindow(width, height, "INTVAL GO")
    texture := driver.CreateTexture(m, 1.0)
    img.SetTexture(texture)
    window.AddChild(img)
    window.OnClose(driver.Terminate)
    ffmpeg_check()

    window.SetScale(flags.DefaultScaleFactor)
    button := theme.CreateButton()
    button.SetHorizontalAlignment(gxui.AlignCenter)
    //button.SetSizeMode(gxui.Fill)
    toggle := func() {
        fullscreen := !window.Fullscreen()
        window.SetFullscreen(fullscreen)
        if fullscreen {
            button.SetText("Make windowed")
        } else {
            button.SetText("Make fullscreen")
        }
    }
    button.SetText("Make fullscreen")
    button.OnClick(func(gxui.MouseEvent) { toggle() })
    window.AddChild(button)
}

func main() {
    gl.StartDriver(appMain)
}
