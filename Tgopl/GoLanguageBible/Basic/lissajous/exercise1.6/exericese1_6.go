package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// 扩展调色板，加入更多颜色
var palette = []color.Color{
	color.White,                        // 白色
	color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // 红色
	color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // 绿色
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // 蓝色
	color.RGBA{0xFF, 0xFF, 0x00, 0xFF}, // 黄色
	color.RGBA{0xFF, 0x00, 0xFF, 0xFF}, // 紫色
	color.RGBA{0x00, 0xFF, 0xFF, 0xFF}, // 青色
}

const (
	whiteIndex = 0 // 白色的索引
)

func main() {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 完整的x振荡器循环次数
		res     = 0.001 // 角度分辨率
		size    = 100   // 图片画布的大小 [-size..+size]
		nframes = 64    // 动画帧数
		delay   = 8     // 每帧之间的延迟，以10ms为单位
	)

	freq := rand.Float64() * 3.0 // y振荡器的频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 相位差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		colorIndex := uint8(i % len(palette)) // 根据帧数选择不同的颜色
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// 根据不同帧的数值设置颜色
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 输出动画GIF
}
