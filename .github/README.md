# in3D
A relatively simple Go powered OpenGL Graphics Engine

Create a new Window, Get OpenGL Context, Setup Camera Projection, create 3D object, Draw!  
Go Ahead, you can do it yourself... `go get github.com/vekjja/in3D`
```go
package main

import (
	"github.com/vekjja/in3D"
)

func main() {

	in3D.Init(800, 600, "Simple Cube in3D")
	in3D.NewLight().Position =
		in3D.Position{X: 10, Y: 1, Z: 10}

	in3D.SetRelPath("../assets/textures")
	texture := in3D.NewTexture("vekjja.jpg")
	color := []float32{1, 1, 1}

	obj := in3D.NewPointsObject(
		in3D.NewPosition(0, 0, -7),
		in3D.Cube,
		texture,
		color,
		in3D.Shader["phong"],
	)
	obj.SceneLogic = func(s *in3D.SceneData) {
		s.XRotation++
		s.YRotation++
	}

	for !in3D.ShouldClose() {
		in3D.Update()
		obj.Draw()
		in3D.SwapBuffers()
	}
}

```
![Simple Rotating Cude in3D](./examples/assets/textures/readme.png)
### ME-TODO:
  *  Optimize all the things!  
  *  Add Shadows, Ambient Occulsion and other light related things  
  *  Have fun making examples!  

### YOU-TODO:
  * Checkout the other examples to see some more basic functionality  
  * `go run examples/mesh/mesh.go`

### Installation on Ubuntu
```
sudo apt install libgl1-mesa-dev libxrandr-dev libxcursor-dev libxi-dev libxinerama-dev

go get github.com/vekjja/in3D
```  
### Installation on Fedora
```sh
sudo dnf install -y libXxf86vm-devel  mesa-libGL-devel mesa-libGLES-devel libXcursor-devel libXrandr-devel libXinerama-devel libXi-devel libX11-devel

go get github.com/vekjja/in3D
```

##### Make sure OpenGL 4.1 is supported by your system, drivers and hardware


