# Sobel

- Input : input_img
- Blur layer : blur_img
- Sobel layer : sobel_img
- Output : sobel

### Step 1. Clone input image

`input_img := input_image.Clone()`

### Step 2. BilateralFilter  :  Noise reduction

`gocv.BilateralFilter(input_img, &blur_img, 5, 75, 75)`

- src : input_img
- dst : blur_img
- kernel_size : 5
- degree : 75

### Step 3. Sobel  :  Edge detector

`gocv.Sobel(blur_img, &sobel_img, gocv.MatTypeCV16S, 1, 0, 3.0, 1.0, 0.0, gocv.BorderDefault)`

- src : blur_img
- dst : sobel_img
- mat_type : gocv.MatTypeCV16S
- axis : 1,0 → X
- kernel_size : 3
- out of boundary : gocv.BorderDefault

***Edge detector : Canny, Sobel, Laplacian ...***

### Step 4. Generalization

`gocv.ConvertScaleAbs(sobel_img, &sobel, 1, 0)`

- src : sobel_img
- dst : sobel
- range : 1, 0