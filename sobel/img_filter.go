package image_processing

import (
	"gocv.io/x/gocv"
)

func get_sobel_image(input_image gocv.Mat) gocv.Mat {

	input_img := input_image.Clone()
	blur_img := gocv.NewMat()
	sobel_img := gocv.NewMat()
	sobel := gocv.NewMat()

	//sobelX := gocv.NewMat()
	//sobelY := gocv.NewMat()

	//sobel_abs_grad_x := gocv.NewMat()
	//sobel_abs_grad_y := gocv.NewMat()

	//BilateralFilter
	gocv.BilateralFilter(input_img, &blur_img, -1, 75, 75)

	//Laplacian
	//laplacian := gocv.NewMat()
	//gocv.Laplacian(sobelimage, &laplacian, sobelimage.Channels(), 5.0, 1.0, 0.0, gocv.BorderDefault)

	//Sobel

	gocv.Sobel(blur_img, &sobel_img, gocv.MatTypeCV16S, 1, 0, 3.0, 1.0, 0.0, gocv.BorderDefault)
	//gocv.Sobel(sobelimage, &sobelX, gocv.MatTypeCV16S, 1, 0, 3.0, 1.0, 0.0, gocv.BorderDefault)
	//gocv.Sobel(sobelimage, &sobelY, gocv.MatTypeCV16S, 0, 1, 3.0, 1.0, 0.0, gocv.BorderDefault)

	//convert image data to absolute value
	gocv.ConvertScaleAbs(sobel_img, &sobel, 1, 0)
	//gocv.ConvertScaleAbs(sobelY, &sobel_abs_grad_y, 1, 0)

	//addweitghted : merge X-Y
	//gocv.AddWeighted(sobel_abs_grad_x, 0.5, sobel_abs_grad_y, 0.5, 0, &sobel)

	//GaussianBlur(result_blurred, result_blurred, Size(3, 3), 3)
	//bilateralFilter()

	return sobel
}
