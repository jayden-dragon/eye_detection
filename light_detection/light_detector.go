package main

import (
	"image"
	c "image/color"

	"gocv.io/x/gocv"
)

func Lightdetector(buffer gocv.Mat) image.Rectangle {
	///read Image
	var rect image.Rectangle
	var flag float64

	input := gocv.NewMat()

	new_black := gocv.NewMatWithSizeFromScalar(gocv.NewScalar(0, 0, 0, 0), buffer.Rows(), buffer.Cols(), gocv.MatTypeCV8U)
	gocv.Rectangle(&new_black, image.Rect(0, 120, 639, 479), c.RGBA{255, 255, 255, 255}, -1)
	gocv.Rectangle(&new_black, image.Rect(0, 0, 310, 479), c.RGBA{255, 255, 255, 255}, -1)

	img2 := gocv.NewMatWithSize(buffer.Rows(), buffer.Cols(), gocv.MatTypeCV8U)

	//fmt.Println("Rows : ", img2.Rows(), "Cols : ", img2.Cols())

	gocv.CvtColor(buffer, &input, gocv.ColorBGRToGray)

	gocv.BitwiseAnd(input, new_black, &img2)

	gocv.GaussianBlur(img2, &img2, image.Point{X: 11, Y: 11}, 0, 0, gocv.BorderReflect101)
	gocv.Threshold(img2, &img2, 180, 255, gocv.ThresholdBinary)
	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(3, 3))

	gocv.Erode(img2, &img2, kernel)
	gocv.Dilate(img2, &img2, kernel)

	contours := gocv.FindContours(img2, gocv.RetrievalExternal, gocv.ChainApproxSimple)
	flag = float64(0)
	for _, c := range contours {
		area := gocv.ContourArea(c)
		if flag < area {
			flag = area
			rect = gocv.BoundingRect(c)
		}
	}
	return rect
}
