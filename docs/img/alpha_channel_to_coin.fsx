#r "nuget: SixLabors.ImageSharp"

open SixLabors.ImageSharp

open SixLabors.ImageSharp.PixelFormats
open SixLabors.ImageSharp.Processing
open SixLabors.ImageSharp.Processing.Processors.Transforms

let baseName = "hexagon_black"
// Example usage
let inputImagePath = $"{baseName}.png"
let outputImagePath = $"{baseName}_edited.png"

let pixelate (imagePath: string) (outputPath: string) =
  use image = Image.Load<Rgba32>(imagePath)

  let modifyPixel (img: Image<Rgba32>) =
    img.Mutate(fun ctx -> ctx.Pixelate 5 |> ignore)

  modifyPixel image
  image.Save(outputPath)

let color = Rgba32(0uy, 0uy, 0uy, 0uy)


let cropToFigure (image: Image<Rgba32>) =
  let xs: Processors.IImageProcessor array = [| EntropyCropProcessor() |]
  image.Mutate(xs)

let transparencyOutsideRadius (image: Image<Rgba32>) (radius: int) =
  let width = image.Width
  let height = image.Height
  let centerX = width / 2
  let centerY = height / 2
  let radiusSquared = radius * radius

  // Process each pixel
  for y in 0 .. height - 1 do
    for x in 0 .. width - 1 do
      let dx = x - centerX
      let dy = y - centerY

      if (dx * dx + dy * dy > radiusSquared) then
        // Set pixel outside the circle to transparent
        //image.Mutate(fun ctx -> ctx.ApplyProcessor())
        let z = image[x, y]

        image[x, y] <- color

let transparentBackground (imagePath: string) (outputPath: string) =
  // Load the image
  use image = Image.Load<Rgba32>(imagePath)

  cropToFigure image
  let radius = image.Width / 2
  transparencyOutsideRadius image radius

  let encoder =
    Formats.Png.PngEncoder(
      TransparentColorMode = Formats.Png.PngTransparentColorMode.Preserve,
      ColorType = Formats.Png.PngColorType.RgbWithAlpha
    )

  image.SaveAsPng(outputPath, encoder)

transparentBackground inputImagePath outputImagePath
