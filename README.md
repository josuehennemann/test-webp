# test-webp

How to convert image (jpeg) to wepb format, using package [`webp`](github.com/chai2010/webp)

## Prerequisite
	
`go get github.com/chai2010/webp`

## Using

`git clone git@github.com:josuehennemann/test-webp.git`
`cd test-webp`
`go build`
`./test-web {image.jpeg}`

After processing, the new image will be saved in the same directory as the original image, but with the webp extension

## Example:

`./test-web dataset/farcry4.jpg`
