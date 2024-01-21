package main

import (
	"fmt"
	"math"
	"time"

	"github.com/NHollmann/Raytracer-Challenge-Go/camera"
	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/intersection"
	"github.com/NHollmann/Raytracer-Challenge-Go/light"
	"github.com/NHollmann/Raytracer-Challenge-Go/matrix"
	"github.com/NHollmann/Raytracer-Challenge-Go/pattern"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

func main() {
	fmt.Println("Raytracer")
	fmt.Println("=========")
	fmt.Println("Begin rendering...")

	startTime := time.Now()
	renderImage()
	duration := time.Since(startTime)

	fmt.Println("Done!")
	fmt.Printf("Rendering took %s\n", duration)
}

func renderImage() {
	world := intersection.NewWorld()

	floor := intersection.NewPlane()
	patA := pattern.NewStripePatternColor(color.New(0.8, 0, 0), color.New(0.4, 0, 0))
	patB := pattern.NewStripePatternColor(color.New(0, 0.8, 0), color.New(0, 0.4, 0))
	patA.SetTransform(matrix.RotationY(math.Pi / 4.0).Mul(matrix.Scaling(0.1, 0.1, 0.1)))
	patB.SetTransform(matrix.RotationY(math.Pi / -4.0).Mul(matrix.Scaling(0.1, 0.1, 0.1)))
	floor.Material.Pattern = pattern.NewCheckerPattern(patA, patB)
	world.AddObject(floor)

	/*
		floor := intersection.NewSphere()
		floor.Transform = matrix.Scaling(10, 0.01, 10)
		floor.Material.Color = color.New(1, 0.9, 0.9)
		floor.Material.Specular = 0
		world.AddObject(floor)

		leftWall := intersection.NewSphere()
		leftWall.Transform = matrix.
			Translation(0, 0, 5).
			Mul(matrix.RotationY(-math.Pi / 4.0)).
			Mul(matrix.RotationX(math.Pi / 2.0)).
			Mul(matrix.Scaling(10, 0.01, 10))
		leftWall.Material = floor.Material
		world.AddObject(leftWall)

		rightWall := intersection.NewSphere()
		rightWall.Transform = matrix.
			Translation(0, 0, 5).
			Mul(matrix.RotationY(math.Pi / 4.0)).
			Mul(matrix.RotationX(math.Pi / 2.0)).
			Mul(matrix.Scaling(10, 0.01, 10))
		rightWall.Material = floor.Material
		world.AddObject(rightWall)
	*/

	middle := intersection.NewSphere()
	middle.Transform = matrix.Translation(-0.5, 1, 0.5)
	middle.Material.Pattern = pattern.NewGradientPatternColor(color.New(0.2, 0.8, 0.0), color.New(1, 0.1, 0))
	middle.Material.Pattern.SetTransform(matrix.RotationY(math.Pi / 8.0).Mul(matrix.Scaling(2, 2, 2).Mul(matrix.Translation(0.5, 0, 0))))
	middle.Material.Color = color.New(0.1, 1, 0.5)
	middle.Material.Diffuse = 0.7
	middle.Material.Specular = 0.3
	middle.Material.Reflective = 0.5
	world.AddObject(middle)

	right := intersection.NewSphere()
	right.Transform = matrix.Translation(1.5, 0.5, -0.5).Mul(matrix.Scaling(0.5, 0.5, 0.5))
	right.Material.Color = color.New(0.5, 1, 0.1)
	right.Material.Diffuse = 0.7
	right.Material.Specular = 0.3
	right.Material.Reflective = 0.3
	world.AddObject(right)

	left := intersection.NewSphere()
	left.Transform = matrix.Translation(-1.5, 0.33, -0.75).Mul(matrix.Scaling(0.33, 0.33, 0.33))
	left.Material.Color = color.New(1, 0.8, 0.1)
	left.Material.Diffuse = 0.7
	left.Material.Specular = 0.1
	world.AddObject(left)

	lightSource := light.NewPoint(tuple.Point(-10, 10, -10), color.New(1, 1, 1))
	world.AddLight(lightSource)

	cam := camera.New(400, 200, math.Pi/3.0)
	cam.Transform = matrix.ViewTransform(
		tuple.Point(0, 1.5, -5),
		tuple.Point(0, 1, 0),
		tuple.Vector(0, 1, 0),
	)

	renderCanvas := cam.Render(world)
	renderCanvas.SavePpmToFile("rendering.ppm")
}
