# github.com/gogpu/gg

## accelerator.go

-   Accelerator
-   CloseAccelerator
-   RegisterAccelerator
-   SetAcceleratorDeviceProvider
-   SetAcceleratorSurfaceTarget

## brush.go

-   BrushFromPattern
-   ColorAt
-   ColorAt
-   Lerp
-   Opaque
-   PatternFromBrush
-   Solid
-   SolidHex
-   SolidRGB
-   SolidRGBA
-   Transparent
-   WithAlpha
-   brushMarker

## brush_custom.go

-   Checkerboard
-   ColorAt
-   HorizontalGradient
-   LinearGradient
-   NewCustomBrush
-   RadialGradient
-   Stripes
-   VerticalGradient
-   WithName
-   brushMarker
-   clampT
-   toCustomBrush

## cmd/ggdemo/main.go

-   drawGradientBackground
-   drawPathDemo
-   drawShapesDemo
-   drawTransformDemo
-   main

## color.go

-   Color
-   FromColor
-   HSL
-   Hex
-   Lerp
-   Premultiply
-   RGB
-   RGBA
-   RGBA2
-   Unpremultiply
-   clamp255
-   clamp65535
-   parseHex

## context.go

-   Clear
-   ClearDash
-   ClearPath
-   ClearWithColor
-   Close
-   ClosePath
-   CubicTo
-   DrawArc
-   DrawCircle
-   DrawEllipse
-   DrawEllipticalArc
-   DrawLine
-   DrawPoint
-   DrawRectangle
-   DrawRoundedRectangle
-   EncodeJPEG
-   EncodePNG
-   Fill
-   FillBrush
-   FillPreserve
-   FlushGPU
-   GetCurrentPoint
-   GetStroke
-   GetTransform
-   Height
-   Identity
-   Image
-   InvertY
-   IsDashed
-   LineTo
-   MoveTo
-   NewContext
-   NewContextForImage
-   NewSubPath
-   PipelineMode
-   Pop
-   Push
-   QuadraticTo
-   RasterizerMode
-   Resize
-   ResizeTarget
-   Rotate
-   RotateAbout
-   SavePNG
-   Scale
-   SetColor
-   SetDash
-   SetDashOffset
-   SetFillBrush
-   SetFillRule
-   SetHexColor
-   SetLineCap
-   SetLineJoin
-   SetLineWidth
-   SetMiterLimit
-   SetPipelineMode
-   SetPixel
-   SetRGB
-   SetRGBA
-   SetRasterizerMode
-   SetStroke
-   SetStrokeBrush
-   SetTransform
-   Shear
-   Stroke
-   StrokeBrush
-   StrokePreserve
-   Transform
-   TransformPoint
-   Translate
-   Width
-   applyClipToPaint
-   arcSegment
-   currentColor
-   doFill
-   doStroke
-   flushGPUAccelerator
-   gpuRenderTarget
-   sdfAccelForShape
-   setForceSDF
-   tryGPUFill
-   tryGPUOp
-   tryGPUStroke

## context_clip.go

-   Clip
-   ClipPreserve
-   ClipRect
-   ResetClip
-   convertPathElements
-   initClipStack

## context_image.go

-   ColorAt
-   CreateImagePattern
-   DrawImage
-   DrawImageCircular
-   DrawImageEx
-   DrawImageRounded
-   ImageBufFromImage
-   LoadImage
-   LoadWebP
-   NewImageBuf
-   SetAnchor
-   SetClamp
-   SetFillPattern
-   SetOpacity
-   SetScale
-   SetStrokePattern
-   pixmapToImageBuf

## context_layer.go

-   PopLayer
-   PushLayer
-   SetBlendMode
-   compositeLayer
-   newLayerStack

## context_mask.go

-   AsMask
-   ClearMask
-   GetMask
-   InvertMask
-   SetMask

## coverage_filler.go

-   GetCoverageFiller
-   RegisterCoverageFiller

## curve.go

-   BoundingBox
-   BoundingBox
-   BoundingBox
-   Contains
-   Deriv
-   End
-   End
-   End
-   Eval
-   Eval
-   Eval
-   Extrema
-   Extrema
-   Height
-   Inflections
-   Length
-   Midpoint
-   NewCubicBez
-   NewLine
-   NewQuadBez
-   NewRect
-   Normal
-   Raise
-   Reversed
-   Start
-   Start
-   Start
-   Subdivide
-   Subdivide
-   Subdivide
-   Subsegment
-   Subsegment
-   Subsegment
-   Tangent
-   Union
-   Width

## dash.go

-   Clone
-   IsDashed
-   NewDash
-   NormalizedOffset
-   PatternLength
-   Scale
-   WithOffset
-   effectiveArray

## examples/basic/main.go

-   main

## examples/bezier_test/main.go

-   main

## examples/clipping/main.go

-   drawStar
-   example1CircularClip
-   example2RectClip
-   example3ClipPreserve
-   example4NestedClips
-   example5ComplexClip
-   example6ResetClip
-   main

## examples/color_emoji/main.go

-   findFont
-   generateBitmapGrid
-   generateColorPalette
-   generateSampleEmoji
-   getTable
-   main

## examples/compute_pipeline/main.go

-   buildDemoScene
-   buildTriptych
-   circleCubics
-   comparePixels
-   main
-   rectLines
-   renderGPU
-   savePNG
-   starLines
-   triangleLines

## examples/gpu/main.go

-   drawRegularPolygon
-   drawStar
-   init
-   main

## examples/images/main.go

-   createTestImage
-   drawExamples
-   drawText
-   main

## examples/recording/main.go

-   exportPDF
-   exportRaster
-   exportSVG
-   main

## examples/scene/main.go

-   buildScene
-   demonstrateIncrementalRender
-   main

## examples/shapes/main.go

-   drawShapes
-   main

## examples/text/main.go

-   findSystemFont
-   main

## examples/text_fallback/main.go

-   drawMultiFaceDemo
-   drawNoEmojiFallback
-   findEmojiFont
-   findMainFont
-   main

## examples/text_transform/main.go

-   findSystemFont
-   main

## gpu/gpu.go

-   SetDeviceProvider
-   init

## gradient.go

-   applyExtendMode
-   clamp01
-   colorAtOffset
-   interpolateColorLinear
-   sortStops

## gradient_linear.go

-   AddColorStop
-   ColorAt
-   NewLinearGradientBrush
-   SetExtend
-   brushMarker
-   firstStopColor

## gradient_radial.go

-   AddColorStop
-   ColorAt
-   NewRadialGradientBrush
-   SetExtend
-   SetFocus
-   brushMarker
-   computeT
-   computeTFocal
-   computeTSimple

## gradient_sweep.go

-   AddColorStop
-   ColorAt
-   NewSweepGradientBrush
-   SetEndAngle
-   SetExtend
-   angleToT
-   brushMarker
-   normalizeAngle

## integration/ggcanvas/canvas.go

-   Close
-   Context
-   Draw
-   Flush
-   Height
-   IsDirty
-   MarkDirty
-   MustNew
-   New
-   Provider
-   RenderDirect
-   Resize
-   Size
-   Texture
-   Width
-   createTexture
-   deferTextureDestruction
-   destroyTexture

## integration/ggcanvas/render.go

-   DefaultRenderOptions
-   RenderTo
-   RenderToEx
-   RenderToPosition
-   RenderToScaled
-   promotePendingTexture

## internal/blend/advanced.go

-   blendColorBurn
-   blendColorDodge
-   blendDarken
-   blendDifference
-   blendExclusion
-   blendHardLight
-   blendLighten
-   blendMultiply
-   blendOverlay
-   blendScreen
-   blendSoftLight
-   maxByte
-   separableBlend

## internal/blend/batch_advanced.go

-   DarkenBatch
-   DifferenceBatch
-   ExclusionBatch
-   LightenBatch
-   MultiplyBatch
-   OverlayBatch
-   ScreenBatch

## internal/blend/batch_porter_duff.go

-   ClearBatch
-   DestinationAtopBatch
-   DestinationBatch
-   DestinationInBatch
-   DestinationOutBatch
-   DestinationOverBatch
-   GetBatchBlendFunc
-   ModulateBatch
-   PlusBatch
-   SourceAtopBatch
-   SourceBatch
-   SourceInBatch
-   SourceOutBatch
-   SourceOverBatch
-   XorBatch

## internal/blend/blend.go

-   Blend
-   destinationIn
-   destinationOut
-   sourceOver

## internal/blend/blend_batch.go

-   BlendBatch
-   BlendBatchAligned
-   BlendSourceOverBatch
-   BlendSourceOverBatchAligned

## internal/blend/hsl.go

-   ClipColor
-   Lum
-   Sat
-   SetLum
-   SetSat
-   blendColor
-   blendHue
-   blendLuminosity
-   blendSaturation
-   hslBlendColor
-   hslBlendHue
-   hslBlendLuminosity
-   hslBlendSaturation
-   max3
-   min3
-   nonSeparableBlend
-   sortRGB

## internal/blend/layer.go

-   BlendMode
-   Bounds
-   Buffer
-   Clear
-   Current
-   CurrentBlendMode
-   Depth
-   NewLayer
-   NewLayerStack
-   Opacity
-   Pop
-   Push
-   SetOpacity
-   compositeLayer

## internal/blend/linear.go

-   BlendLinear
-   GetBlendFuncLinear

## internal/blend/math.go

-   addClamp
-   clamp255
-   div255
-   div255Exact
-   inv255
-   mulDiv255
-   mulDiv255Exact
-   subClamp

## internal/blend/porter_duff.go

-   GetBlendFunc
-   addDiv255
-   blendClear
-   blendDestination
-   blendDestinationAtop
-   blendDestinationIn
-   blendDestinationOut
-   blendDestinationOver
-   blendModulate
-   blendPlus
-   blendSource
-   blendSourceAtop
-   blendSourceIn
-   blendSourceOut
-   blendSourceOver
-   blendXor
-   clampAdd
-   minByte

## internal/cache/cache.go

-   Capacity
-   Clear
-   Delete
-   Get
-   GetOrCreate
-   Len
-   New
-   Set
-   Stats
-   evictOldest

## internal/cache/lru.go

-   Clear
-   Len
-   MoveToFront
-   Oldest
-   PushFront
-   Remove
-   RemoveOldest
-   newLRUList
-   unlink

## internal/cache/sharded.go

-   Capacity
-   Clear
-   Delete
-   Get
-   GetOrCreate
-   IntHasher
-   Len
-   NewSharded
-   ResetStats
-   Set
-   ShardLen
-   Stats
-   StringHasher
-   TotalCapacity
-   Uint64Hasher
-   getShard

## internal/clip/edge_clipper.go

-   Clip
-   ClipCubic
-   ClipLine
-   ClipQuadratic
-   NewEdgeClipper
-   allInside
-   boundsIntersect
-   cbrt
-   chopCubicAt
-   chopCubicAtExtremaRoots
-   chopCubicAtXExtrema
-   chopCubicAtYExtrema
-   chopQuadAt
-   chopQuadAtXExtrema
-   chopQuadAtYExtrema
-   clipMonoCubic
-   clipMonoQuad
-   cubicIntersectX
-   cubicIntersectY
-   evalCubic
-   evalQuad
-   filterAndSort
-   isNotMonotonic
-   outcode
-   quadIntersectX
-   quadIntersectY
-   solveCubic
-   solveQuadratic
-   subdivideCubicRange
-   subdivideQuadRange

## internal/clip/mask.go

-   ApplyCoverage
-   Bounds
-   Coverage
-   Mask
-   NewMaskClipper
-   evalCubicBezier
-   evalQuadraticBezier
-   flattenPath
-   isPathElement
-   isPathElement
-   isPathElement
-   isPathElement
-   isPathElement
-   makeEdge
-   rasterizePath
-   rasterizeScanline
-   rasterizeScanlineAA
-   sortFloats

## internal/clip/stack.go

-   Bounds
-   Coverage
-   Depth
-   IsVisible
-   NewClipStack
-   Pop
-   PushPath
-   PushRect
-   Reset

## internal/clip/types.go

-   Add
-   Bottom
-   Bounds
-   Bounds
-   Contains
-   Intersect
-   Intersects
-   IsEmpty
-   Lerp
-   Mul
-   NewRect
-   Pt
-   Right
-   Sub

## internal/color/convert.go

-   F32ToU8
-   LinearToSRGB
-   LinearToSRGBColor
-   SRGBToLinear
-   SRGBToLinearColor
-   U8ToF32
-   clampAndRound

## internal/color/lut.go

-   LinearToSRGBFast
-   LinearToSRGBSlow
-   SRGBToLinearFast
-   SRGBToLinearSlow
-   init

## internal/filter/blur.go

-   Apply
-   ExpandBounds
-   NewBlurFilter
-   NewBlurFilterXY
-   blurHorizontal
-   blurVertical
-   clampInt
-   clampUint8
-   copyFromTemp
-   copyPixmapRegion
-   copyToTemp
-   getTempBuffer
-   putTempBuffer

## internal/filter/colormatrix.go

-   Apply
-   ExpandBounds
-   Multiply
-   NewBrightnessFilter
-   NewColorMatrixFilter
-   NewColorTintFilter
-   NewContrastFilter
-   NewGrayscaleFilter
-   NewHueRotateFilter
-   NewIdentityColorMatrix
-   NewInvertFilter
-   NewOpacityFilter
-   NewSaturationFilter
-   NewSepiaFilter
-   cosf
-   modTwoPi
-   sinf

## internal/filter/kernel.go

-   BoxKernel
-   CachedGaussianKernel
-   GaussianKernel
-   KernelCenter
-   OptimalKernelSize
-   get
-   newKernelCache

## internal/filter/shadow.go

-   Apply
-   ExpandBounds
-   NewDropShadowFilter
-   NewSimpleDropShadow
-   blurAlphaChannel
-   clamp255f
-   compositeShadow
-   extractAlpha
-   getAlphaBuffer
-   putAlphaBuffer

## internal/gpu/adapter.go

-   BeginComputePass
-   CreateBindGroup
-   CreateBindGroupLayout
-   CreateBuffer
-   CreateComputePipeline
-   CreatePipelineLayout
-   CreateShaderModule
-   CreateTexture
-   DestroyBindGroup
-   DestroyBindGroupLayout
-   DestroyBuffer
-   DestroyComputePipeline
-   DestroyPipelineLayout
-   DestroyShaderModule
-   DestroyTexture
-   Dispatch
-   End
-   MaxBufferSize
-   MaxWorkgroupSize
-   NewHALAdapter
-   ReadBuffer
-   ReadTexture
-   SetBindGroup
-   SetPipeline
-   Submit
-   SupportsCompute
-   WaitIdle
-   WriteBuffer
-   WriteTexture
-   convertBindGroupEntry
-   convertBindGroupLayoutEntry
-   convertBufferUsage
-   convertTextureFormat
-   newID

## internal/gpu/adaptive_filler.go

-   ComputeFiller
-   FillCoverage
-   SparseFiller

## internal/gpu/analytic_filler.go

-   AlphaRuns
-   Coverage
-   Fill
-   FillPath
-   FillToBuffer
-   Height
-   NewAnalyticFiller
-   Reset
-   Width
-   accumulateCoverageSubpixel
-   applyFillRule
-   clamp32
-   computeSegmentCoverage
-   coverageToRuns
-   max32f
-   min32f
-   processScanlineWithScale
-   stepCurveSegment

## internal/gpu/analytic_filler_vello.go

-   Fill
-   NewAnalyticFillerVello
-   Reset
-   collectSegments
-   processScanlineVello
-   velloAbsF32

## internal/gpu/atlas.go

-   AllocCount
-   AllocCount
-   Allocate
-   Allocate
-   AllocateAndUpload
-   Close
-   Contains
-   Height
-   IsClosed
-   IsValid
-   NewRectAllocator
-   NewTextureAtlas
-   Reset
-   Reset
-   String
-   Texture
-   Upload
-   UsedArea
-   Utilization
-   Utilization
-   Width
-   allocateNewShelf
-   allocateOnShelf
-   fitsOnShelf

## internal/gpu/backend.go

-   Close
-   Close
-   Device
-   Fill
-   GPUInfo
-   Height
-   Init
-   IsInitialized
-   Name
-   NewBackend
-   NewRenderer
-   Queue
-   RenderScene
-   Stroke
-   Width
-   newGPURenderer

## internal/gpu/buffer.go

-   CreateBuffer
-   CreateBufferSimple
-   CreateStagingBuffer
-   Descriptor
-   Destroy
-   GetMappedRange
-   IsDestroyed
-   Label
-   MapAsync
-   MapState
-   NewBuffer
-   PollMapAsync
-   Raw
-   Size
-   String
-   String
-   Unmap
-   Usage

## internal/gpu/coarse.go

-   BinCurveEdges
-   CalculateBackdrop
-   Entries
-   EntriesAtLocation
-   Grid
-   HasNext
-   NewCoarseRasterizer
-   NewIterator
-   Next
-   Rasterize
-   Reset
-   Reset
-   Segments
-   SortEntries
-   TileColumns
-   TileRows
-   addEntry
-   binSingleCurveEdge
-   calculateCurveBackdrops
-   clampU16
-   cloneCurveEdge
-   getCurveXBounds
-   isSorted
-   processRow
-   rasterizeLine
-   rasterizeSlopedLine
-   rasterizeVerticalLine

## internal/gpu/command_encoder.go

-   BeginComputePass
-   BeginRenderPass
-   ClearBuffer
-   CopyBufferToBuffer
-   CopyBufferToTexture
-   CopyTextureToBuffer
-   CopyTextureToTexture
-   CoreBuffer
-   Finish
-   Label
-   Label
-   NewCoreCommandEncoder
-   NewCoreCommandEncoderWithDevice
-   Status
-   checkRecording
-   checkRecordingLocked
-   endComputePass
-   endRenderPass
-   statusLocked
-   toCoreDescriptor

## internal/gpu/commands.go

-   BeginComputePass
-   BeginRenderPass
-   CopyTextureToBuffer
-   CopyTextureToTexture
-   Dispatch
-   DispatchForSize
-   DispatchWorkgroups
-   DispatchWorkgroupsForSize
-   Draw
-   Draw
-   DrawFullScreen
-   DrawFullScreenTriangle
-   DrawIndexed
-   End
-   End
-   Finish
-   Finish
-   Finish
-   ID
-   NewCommandBuffer
-   NewCommandEncoder
-   NewComputeCommandBuilder
-   NewQueueSubmitter
-   NewRenderCommandBuilder
-   PassCount
-   SetBindGroup
-   SetBindGroup
-   SetBindGroup
-   SetBindGroup
-   SetIndexBuffer
-   SetPipeline
-   SetPipeline
-   SetPipeline
-   SetPipeline
-   SetVertexBuffer
-   Submit
-   Target
-   WriteBuffer
-   WriteTexture
-   endPass

## internal/gpu/compute_pass.go

-   Destroy
-   DispatchCount
-   DispatchWorkgroups
-   DispatchWorkgroupsIndirect
-   End
-   ID
-   IsDestroyed
-   IsEnded
-   Label
-   SetBindGroup
-   SetPipeline
-   Size
-   Size
-   Size
-   State
-   String
-   WorkgroupSize
-   checkRecording

## internal/gpu/convex_renderer.go

-   BuildConvexVertices
-   Destroy
-   NewConvexRenderer
-   RecordDraws
-   buildConvexVerticesReuse
-   convexVertexCount
-   convexVertexLayout
-   createPipeline
-   destroy
-   destroyPipeline
-   ensurePipeline
-   ensurePipelineWithStencil
-   writeConvexVertex

## internal/gpu/convexity.go

-   AnalyzeConvexity
-   IsConvex

## internal/gpu/device.go

-   CheckDeviceLimits
-   String
-   createDevice
-   getDeviceQueue
-   getGPUInfo
-   logGPUInfo
-   releaseAdapter
-   releaseDevice

## internal/gpu/fine.go

-   Alphas
-   FillRule
-   Grid
-   NewFineRasterizer
-   NewStripRenderer
-   Rasterize
-   RasterizeCurves
-   RenderTiles
-   RenderToBuffer
-   Reset
-   Reset
-   SetAliasMode
-   SetFillRule
-   SetFillRule
-   Strips
-   clampf32
-   emitSentinel
-   fillTileWithBackdrop
-   finalizeTile
-   finalizeTileFromCurves
-   finalizeTileToStrip
-   initTileWinding
-   initTileWindingForStrip
-   lineEdgeToSegment
-   processCubicEdgeForTile
-   processEdgeForTile
-   processLineEdgeForTile
-   processQuadraticEdgeForTile
-   processSegment
-   processSegmentForStrip
-   processTileWithCurves
-   startStrip

## internal/gpu/flatten.go

-   FlattenPath
-   FlattenPathTo
-   NewFlattenContext
-   Reset
-   Segments
-   absf32
-   addMonotonicLine
-   findQuadraticRoots
-   flattenCubic
-   flattenQuadratic
-   lerp
-   maxf32
-   minf32
-   splitCubicAt
-   splitQuadraticAt
-   subdivideCubic
-   subdivideQuadratic

## internal/gpu/gpu_coarse.go

-   Destroy
-   GetTileEntries
-   IsInitialized
-   IsShaderReady
-   NewGPUCoarseRasterizer
-   Rasterize
-   SPIRVCode
-   TileColumns
-   TileRows
-   coarseConfigToBytes
-   computeEntriesCPU
-   convertSegments
-   createBindGroupLayouts
-   createPipelineLayout
-   createPipelines
-   init
-   processRowCPU
-   processSlopedLineCPU
-   processVerticalLineCPU
-   tileEntriesToBytes

## internal/gpu/gpu_fine.go

-   Destroy
-   FillRuleToGPU
-   IsInitialized
-   IsShaderReady
-   NewGPUFineRasterizer
-   Rasterize
-   SPIRVCode
-   boolToUint32
-   buildTileData
-   computeCoverageCPU
-   computePixelArea
-   convertSegments
-   convertTileInfos
-   convertTileSegmentRefs
-   createBindGroupLayouts
-   createPipelineLayout
-   createPipelines
-   init
-   segmentsToBytes
-   tileRefsToBytes
-   tilesToBytes
-   windingToCoverage
-   writeFloat32
-   writeInt32
-   writeUint32

## internal/gpu/gpu_flatten.go

-   ComputeCursorStates
-   ConvertAffineToGPU
-   ConvertPathToGPU
-   Destroy
-   EstimateSegmentCount
-   Flatten
-   FlattenWithContext
-   IsInitialized
-   IsShaderReady
-   MaxPaths
-   MaxSegments
-   NewGPUFlattenRasterizer
-   SPIRVCode
-   SetTolerance
-   Tolerance
-   affineTransformToBytes
-   createBindGroupLayouts
-   createPipelineLayout
-   createPipelines
-   cursorStatesToBytes
-   flattenCPU
-   flattenConfigToBytes
-   init
-   pathElementsToBytes
-   pointsToBytes
-   segmentCountsToBytes
-   wangCubic
-   wangQuadratic

## internal/gpu/gpu_text.go

-   AtlasRGBAData
-   AtlasSize
-   DirtyAtlases
-   GlyphCount
-   LayoutText
-   MarkClean
-   NewGPUTextEngine
-   PxRange
-   computeFontID

## internal/gpu/gpu_texture.go

-   BytesPerPixel
-   Close
-   CreateTexture
-   CreateTextureFromPixmap
-   DownloadPixmap
-   Format
-   Height
-   IsReleased
-   Label
-   SetMemoryManager
-   SizeBytes
-   String
-   String
-   TextureID
-   ToWGPUFormat
-   UploadPixmap
-   UploadRegion
-   ViewID
-   Width

## internal/gpu/gpu_textures.go

-   destroyTextures
-   ensureSurfaceTextures
-   ensureTextures

## internal/gpu/logger.go

-   Enabled
-   Handle
-   WithAttrs
-   WithGroup
-   init
-   setLogger
-   slogger

## internal/gpu/memory.go

-   AllocTexture
-   Close
-   Contains
-   FreeTexture
-   NewMemoryManager
-   SetBudget
-   Stats
-   String
-   Textures
-   TouchTexture
-   evictIfNeeded
-   registerTextureLocked
-   removeTextureLocked
-   unregisterTexture

## internal/gpu/path_convert.go

-   clampU8
-   convertPathToPathDef
-   convertPathToStrokeElements
-   convertShapeToPathDef
-   extractColorU8
-   flushCubics

## internal/gpu/pipeline.go

-   BlendPipelineCount
-   Build
-   Close
-   CreateBlendBindGroup
-   CreateBlitBindGroup
-   CreateStripBindGroup
-   GetBlendLayout
-   GetBlendPipeline
-   GetBlitLayout
-   GetBlitPipeline
-   GetCompositePipeline
-   GetStripLayout
-   GetStripPipeline
-   IsInitialized
-   NewBindGroupBuilder
-   NewPipelineCache
-   WarmupBlendPipelines
-   createBlendPipeline
-   createBlitPipeline
-   createCompositePipeline
-   createStripPipeline

## internal/gpu/pipeline_cache_core.go

-   Clear
-   CodeHash
-   ComputePipelineCount
-   Destroy
-   DestroyAll
-   GetOrCreateComputePipeline
-   GetOrCreateRenderPipeline
-   HashComputePipelineDescriptor
-   HashRenderPipelineDescriptor
-   HitRate
-   ID
-   IsDestroyed
-   Label
-   NewPipelineCacheCore
-   NewShaderModule
-   Raw
-   RenderPipelineCount
-   Size
-   Stats
-   createComputePipeline
-   createRenderPipeline
-   hashBytes
-   hashWriteBool
-   hashWriteString
-   hashWriteUint32
-   hashWriteUint64
-   nextPipelineID

## internal/gpu/render_pass.go

-   Destroy
-   Destroy
-   Draw
-   DrawIndexed
-   DrawIndexedIndirect
-   DrawIndirect
-   End
-   ID
-   ID
-   IsDestroyed
-   IsDestroyed
-   IsEnded
-   Label
-   Label
-   SetBindGroup
-   SetBlendConstant
-   SetIndexBuffer
-   SetPipeline
-   SetScissorRect
-   SetStencilReference
-   SetVertexBuffer
-   SetViewport
-   State
-   String
-   checkRecording

## internal/gpu/render_session.go

-   ConvexRendererRef
-   Destroy
-   EnsureTextures
-   NewGPURenderSession
-   RenderFrame
-   RenderMode
-   SDFPipeline
-   SetConvexRenderer
-   SetSDFPipeline
-   SetStencilRenderer
-   SetSurfaceTarget
-   SetTextAtlas
-   SetTextPipeline
-   Size
-   StencilRendererRef
-   TextAtlasView
-   TextPipelineRef
-   aggregateTextBatches
-   buildConvexResources
-   buildSDFResources
-   buildStencilResourcesBatch
-   buildTextResources
-   copySubmitAndReadback
-   destroyPersistentBuffers
-   drainQueue
-   encodeSubmitReadback
-   encodeSubmitSurface
-   ensurePipelines
-   ensureTextPipeline
-   prepareTextResources

## internal/gpu/renderer.go

-   Close
-   Height
-   LayerDepth
-   MemoryStats
-   NewGPUSceneRenderer
-   RenderScene
-   RenderSceneWithContext
-   RenderToPixmap
-   RenderToPixmapWithContext
-   Resize
-   Width
-   applyTileCoverage
-   applyTileToMask
-   blendTextures
-   clearTexture
-   compositeToTarget
-   downloadToPixmap
-   expandStroke
-   getCurrentTarget
-   popClip
-   popLayer
-   processClipCommand
-   processCommandsWithContext
-   processDrawCommand
-   processLayerCommand
-   processPathCommand
-   pushClip
-   pushLayer
-   rasterizeClipFromGrid
-   rasterizeFromGrid
-   renderClipMask
-   renderFill
-   renderImage
-   renderStroke

## internal/gpu/scene_bridge.go

-   BuildEdgesFromScenePath
-   IsEmpty
-   Points
-   TransformPoint
-   Verbs

## internal/gpu/sdf_gpu.go

-   CanAccelerate
-   CanCompute
-   Close
-   DrawText
-   FillPath
-   FillShape
-   Flush
-   Init
-   Name
-   PendingCount
-   SceneStats
-   SetDeviceProvider
-   SetForceSDF
-   SetLogger
-   SetPipelineMode
-   SetSurfaceTarget
-   StrokePath
-   StrokeShape
-   effectivePipelineMode
-   extractConvexPolygon
-   flushLocked
-   getColorFromPaint
-   initGPU
-   initVelloAccelerator
-   queueShape
-   sameTarget
-   syncTextAtlases

## internal/gpu/sdf_render.go

-   Destroy
-   DetectedShapeToRenderShape
-   NewSDFRenderPipeline
-   RecordDraws
-   RenderShapes
-   Size
-   buildSDFRenderVertices
-   buildSDFRenderVerticesReuse
-   createAndUploadBuffer
-   createPipeline
-   destroyPipeline
-   destroyTextures
-   encodeAndReadback
-   ensurePipelineWithStencil
-   ensureReady
-   ensureTextures
-   makeSDFRenderUniform
-   sdfRenderVertexLayout
-   writeSDFRenderVertex

## internal/gpu/segment.go

-   Add
-   AddLine
-   Bounds
-   Bounds
-   CrossesTileRow
-   DeltaX
-   DeltaY
-   InverseSlope
-   IsHorizontal
-   IsVertical
-   Len
-   NewLineSegment
-   NewSegmentList
-   Reset
-   Segments
-   SegmentsInTileRow
-   Slope
-   SortByTileY
-   TileXRange
-   TileYRange
-   XAtY
-   YAtX

## internal/gpu/shader_helper.go

-   CompileShaderToSPIRV
-   CreateShaderModule
-   Destroy

## internal/gpu/shaders.go

-   BlendModeToShader
-   CompileShaders
-   GetBlendShaderSource
-   GetBlitShaderSource
-   GetCompositeShaderSource
-   GetStripShaderSource
-   IsValid
-   ShaderToBlendMode
-   ValidateBlendModeMapping

## internal/gpu/sparse_strips.go

-   DefaultConfig
-   FillRule
-   Get
-   GetStats
-   Grid
-   Height
-   NewSparseStripsPool
-   NewSparseStripsRasterizer
-   Put
-   RasterizePath
-   RasterizePath
-   RasterizeToStrips
-   RenderStripsToBuffer
-   RenderToBuffer
-   Reset
-   Segments
-   SetFillRule
-   SetSize
-   Strips
-   Width
-   calculateStripEndX
-   fillGap
-   renderStripAlphas

## internal/gpu/sparse_strips_filler.go

-   FillCoverage
-   convertGGToScenePath
-   convertToSceneFillRule

## internal/gpu/sparse_strips_gpu.go

-   CheckGPUComputeSupport
-   Destroy
-   Destroy
-   FillRule
-   Grid
-   Grid
-   IsGPUAvailable
-   IsGPUAvailable
-   IsStageGPUAvailable
-   NewHybridFineRasterizer
-   NewHybridPipeline
-   Rasterize
-   RasterizePath
-   Reset
-   Reset
-   ResetStats
-   SetFillRule
-   SetFillRule
-   SetTolerance
-   Stats
-   Stats
-   String
-   coverageToFineGrid
-   coverageToGrid
-   gpuEntriesToCPUCoarse
-   initGPURasterizers
-   runCoarseStage
-   runFineStage
-   runFlattenStage
-   shouldUseGPU
-   shouldUseGPU

## internal/gpu/sparse_strips_simd.go

-   BlendTileSIMD
-   FinalizeTileSIMD
-   InitTileWindingSIMD
-   ProcessMultipleSegmentsSIMD
-   ProcessSegmentSIMD
-   RenderToBufferSIMD
-   finalizeTileEvenOddSIMD
-   finalizeTileNonZeroSIMD
-   processRowsSIMD

## internal/gpu/stencil_pipeline.go

-   createPipelines
-   destroyPipelines

## internal/gpu/stencil_renderer.go

-   Destroy
-   EnsureTextures
-   NewStencilRenderer
-   RecordPath
-   RenderPassDescriptor
-   RenderPath
-   ResolveTexture
-   Size
-   compositeBGRAOverRGBA
-   convertBGRAToRGBA
-   createAndUploadVertexBuffer
-   createRenderBuffers
-   createUniformAndBindGroup
-   destroy
-   destroyTextures
-   encodeAndReadback
-   ensureReady
-   float32SliceToBytes
-   makeCoverUniform
-   makeStencilFillUniform
-   submitAndReadback

## internal/gpu/tessellate.go

-   Bounds
-   CoverQuad
-   NewFanTessellator
-   Reset
-   TessellatePath
-   TriangleCount
-   Vertices
-   emitFanTriangle
-   flattenCubicFan
-   flattenQuadFan
-   updateBounds

## internal/gpu/text_pipeline.go

-   AtlasManager
-   Close
-   Close
-   Config
-   DefaultTextPipelineConfig
-   DefaultTextRendererConfig
-   Destroy
-   GetMSDFTextShaderSource
-   Init
-   Init
-   IsInitialized
-   NewMSDFTextPipeline
-   NewTextPipeline
-   NewTextPipelineDefault
-   NewTextRenderer
-   Pipeline
-   RecordDraws
-   RenderText
-   RenderTextBatch
-   SyncAtlases
-   buildTextIndexData
-   buildTextVertexData
-   createPipeline
-   destroyPipeline
-   ensurePipelineWithStencil
-   generateQuadIndices
-   makeTextUniform
-   prepareUniforms
-   quadsToVertices
-   rgbToRGBA
-   textVertexLayout
-   writeTextVertex

## internal/gpu/texture.go

-   ArrayLayerCount
-   Aspect
-   BaseArrayLayer
-   BaseMipLevel
-   CreateCoreTexture
-   CreateCoreTextureSimple
-   CreateView
-   DepthOrArrayLayers
-   Descriptor
-   Descriptor
-   Destroy
-   Destroy
-   Dimension
-   Dimension
-   Format
-   Format
-   GetDefaultView
-   Height
-   IsDefault
-   IsDestroyed
-   IsDestroyed
-   Label
-   Label
-   MipLevelCount
-   MipLevelCount
-   NewTexture
-   Raw
-   Raw
-   SampleCount
-   Size
-   Texture
-   Usage
-   Width
-   createDefaultView
-   destroy
-   halViewDescToViewDesc
-   textureViewDimensionFromTexture

## internal/gpu/tile.go

-   Bounds
-   FillRule
-   FillSolid
-   ForEach
-   ForEachInRow
-   ForEachSorted
-   Get
-   Get
-   GetCoverage
-   GetOrCreate
-   Has
-   IsEmpty
-   IsSolid
-   Key
-   NewTileGrid
-   NewTilePool
-   PixelBounds
-   PixelToTile
-   PixelToTileF
-   PixelX
-   PixelY
-   Put
-   Reset
-   Reset
-   SetCoverage
-   SetFillRule
-   TileCount

## internal/gpu/tilecompute/coarse.go

-   CoarseRasterize
-   computeLineBBox
-   emitDrawToTiles
-   extractPathFillRules
-   generatePTCLs
-   groupLinesByPath
-   resolveDrawParams
-   runPathStages
-   tileSegRange

## internal/gpu/tilecompute/compositor.go

-   blendSourceOver

## internal/gpu/tilecompute/draw_leaf.go

-   combine
-   drawLeafScan
-   drawReduce
-   newDrawMonoid

## internal/gpu/tilecompute/euler.go

-   cubicParamsFromPointsDerivs
-   espcIntInvApprox
-   eulerParamsFromAngles
-   eval
-   evalTh
-   evalWithOffset
-   evalWithOffset
-   integEuler10

## internal/gpu/tilecompute/fine.go

-   fillPath
-   fineRasterizeTile
-   premulToStraightU8

## internal/gpu/tilecompute/flatten.go

-   FlattenFill
-   cubeSignedSqrt
-   evalCubicAndDeriv
-   flattenEulerFill
-   trailingZeros32

## internal/gpu/tilecompute/path_count.go

-   pathCountMain

## internal/gpu/tilecompute/path_tiling.go

-   pathTilingMain

## internal/gpu/tilecompute/pathtag.go

-   combine
-   newPathMonoid
-   numPathTagWords
-   pathtagReduce
-   pathtagScan

## internal/gpu/tilecompute/ptcl.go

-   NewPTCL
-   ReadCmd
-   ReadColorData
-   ReadEndClipData
-   ReadFillData
-   WriteBeginClip
-   WriteColor
-   WriteEnd
-   WriteEndClip
-   WriteFill
-   WriteSolid

## internal/gpu/tilecompute/rasterizer.go

-   LineSoupFromVelloLine
-   NewRasterizer
-   Rasterize
-   RasterizeScene
-   RasterizeScenePTCL
-   computePath

## internal/gpu/tilecompute/scene_encode.go

-   EncodeScene
-   PackScene
-   packPathTags

## internal/gpu/tilecompute/util.go

-   abs32
-   add
-   atan2
-   ceil32
-   clamp32
-   copysign32
-   cos32
-   dot
-   floor32
-   length
-   lengthSq
-   max32
-   maxi32
-   maxu32
-   min32
-   mini32
-   minu32
-   mul
-   newVec2
-   pow32
-   round32
-   signum32
-   sin32
-   span
-   sub
-   toArray
-   vec2FromArray

## internal/gpu/tilecompute_filler.go

-   FillCoverage
-   segmentsToLineSoup

## internal/gpu/vello_accelerator.go

-   CanAccelerate
-   CanCompute
-   Close
-   FillPath
-   FillShape
-   Flush
-   Init
-   InitStandalone
-   Name
-   PendingCount
-   RenderSceneCompute
-   SetDeviceProvider
-   SetLogger
-   StrokePath
-   StrokeShape
-   buildPathMetadata
-   clampF
-   compositeOver
-   computePathBBox
-   dispatchComputeScene
-   flushLocked
-   initGPU
-   logPipelineDiagnostics
-   readbackBuffer
-   uint32SliceToBytes
-   unpackPixels
-   uploadPathAuxData
-   velloSameTarget

## internal/gpu/vello_compute.go

-   AllocateBuffers
-   Close
-   ComputeWorkgroupCount
-   DestroyBuffers
-   Dispatch
-   Init
-   NewVelloComputeDispatcher
-   String
-   cleanup
-   computeBufferSizes
-   createVelloBuffer
-   destroyPartialInit
-   encodeComputeStages
-   sizeInBytes
-   stageBindGroupEntries
-   stageBindGroupLayoutEntries
-   submitAndWait
-   toBytes

## internal/gpu/vello_tiles.go

-   Fill
-   NewTileRasterizer
-   Reset
-   abs32
-   addSegmentToTile
-   binSegments
-   computeBackdropPrefixSum
-   emitScanline
-   fillPath
-   fillTileScanline
-   velloSpan

## internal/gpucore/pipeline.go

-   Config
-   Destroy
-   Execute
-   ExecuteWithStats
-   IsInitialized
-   NewHybridPipeline
-   Resize
-   SetTolerance
-   SetUseCPUFallback
-   TileColumns
-   TileCount
-   TileRows
-   Tolerance
-   UseGPU
-   init

## internal/image/affine.go

-   Identity
-   Invert
-   Multiply
-   Rotate
-   RotateAt
-   Scale
-   ScaleAt
-   Shear
-   TransformPoint
-   Translate

## internal/image/buf.go

-   Bounds
-   ByteSize
-   Clear
-   Clone
-   Data
-   Fill
-   Format
-   FromRaw
-   GetRGBA
-   Height
-   InvalidatePremulCache
-   IsEmpty
-   IsPremulCached
-   NewImageBuf
-   NewImageBufWithStride
-   PixelBytes
-   PixelOffset
-   PremultipliedData
-   RowBytes
-   SetPixelBytes
-   SetRGBA
-   Stride
-   SubImage
-   Width
-   computePremultiplied
-   premulPixel

## internal/image/draw.go

-   DrawImage
-   String
-   blend
-   blendMultiply
-   blendNormal
-   blendOverlay
-   blendScreen
-   overlayChannel

## internal/image/format.go

-   BitsPerChannel
-   BytesPerPixel
-   Channels
-   HasAlpha
-   ImageBytes
-   Info
-   IsGrayscale
-   IsPremultiplied
-   IsValid
-   PremultipliedVersion
-   RowBytes
-   String
-   UnpremultipliedVersion

## internal/image/interp.go

-   Sample
-   SampleBicubic
-   SampleBilinear
-   SampleNearest
-   String
-   bicubicInterp
-   clamp
-   clampFloat
-   cubicWeight
-   lerp
-   lerp2D

## internal/image/io.go

-   Decode
-   DecodeJPEG
-   DecodePNG
-   DecodeWebP
-   EncodeJPEG
-   EncodePNG
-   EncodeToBytes
-   EncodeToJPEGBytes
-   FromStdImage
-   LoadImage
-   LoadImageFromBytes
-   LoadJPEG
-   LoadPNG
-   LoadWebP
-   SaveJPEG
-   SavePNG
-   ToStdImage

## internal/image/mipmap.go

-   GenerateMipmaps
-   Level
-   LevelForScale
-   NumLevels
-   Release
-   downsample

## internal/image/pattern.go

-   Image
-   Interpolation
-   Mipmaps
-   NewImagePattern
-   Opacity
-   Sample
-   SampleWithScale
-   SpreadMode
-   String
-   Transform
-   WithInterpolation
-   WithMipmaps
-   WithOpacity
-   WithSpreadMode
-   WithTransform
-   applySpreadMode
-   reflectCoord

## internal/image/pool.go

-   Get
-   GetFromDefault
-   NewPool
-   Put
-   PutToDefault

## internal/parallel/dirty.go

-   Clear
-   Count
-   ForEachDirty
-   GetAndClear
-   IsDirty
-   IsEmpty
-   Mark
-   MarkAll
-   MarkRect
-   NewDirtyRegion
-   Resize
-   TilesX
-   TilesY
-   TotalTiles

## internal/parallel/pool.go

-   Close
-   ExecuteAll
-   ExecuteAsync
-   IsRunning
-   NewWorkerPool
-   QueuedWork
-   Submit
-   Workers
-   drainQueue
-   steal
-   worker

## internal/parallel/rasterizer.go

-   Clear
-   ClearDirty
-   Close
-   Composite
-   CompositeDirty
-   DirtyTileCount
-   FillRect
-   FillTiles
-   Grid
-   Height
-   MarkAllDirty
-   MarkRectDirty
-   NewParallelRasterizer
-   NewParallelRasterizerWithWorkers
-   Resize
-   TileCount
-   Width
-   clearTile
-   colorToRGBA
-   compositeTile
-   fillRectInTile

## internal/parallel/tile.go

-   Bounds
-   ByteSize
-   CanvasPixelOffset
-   Contains
-   PixelOffset
-   Reset
-   Stride

## internal/parallel/tile_grid.go

-   AllTiles
-   ClearDirty
-   Close
-   DirtyTiles
-   ForEach
-   ForEachDirty
-   Height
-   MarkAllDirty
-   MarkDirty
-   MarkRectDirty
-   NewTileGrid
-   Resize
-   TileAt
-   TileAtPixel
-   TileCount
-   TilesInRect
-   TilesX
-   TilesY
-   Width
-   allocateTiles

## internal/parallel/tile_pool.go

-   Get
-   GetTile
-   NewTilePool
-   Put
-   PutTile
-   getOrCreatePool
-   poolKey

## internal/path/edge_iter.go

-   CollectEdges
-   NewEdgeIter
-   Next
-   closeLine
-   handleCubicTo
-   handleLineTo
-   handleMoveTo
-   handleQuadTo
-   insertElements
-   processFlattened

## internal/path/flatten.go

-   Add
-   Distance
-   Dot
-   Flatten
-   Length
-   Lerp
-   Mul
-   Sub
-   distanceToLine
-   flattenCubic
-   flattenCubicRec
-   flattenQuadratic
-   flattenQuadraticRec
-   isPathElement
-   isPathElement
-   isPathElement
-   isPathElement
-   isPathElement

## internal/raster/alpha_runs.go

-   Add
-   AddWithCoverage
-   Clear
-   CopyTo
-   GetAlpha
-   IsEmpty
-   Iter
-   IterRuns
-   NewAlphaRuns
-   Reset
-   SetOffset
-   Width
-   addWithMaxValue
-   breakRun
-   catchOverflow

## internal/raster/analytic_filler.go

-   AlphaRuns
-   Coverage
-   Fill
-   FillPath
-   FillToBuffer
-   Height
-   NewAnalyticFiller
-   Reset
-   Width
-   accumulateCoverageSubpixel
-   applyFillRule
-   clamp32
-   computeSegmentCoverage
-   coverageToRuns
-   offscreenLeftWinding
-   processScanlineWithScale
-   stepCurveSegment

## internal/raster/curve_aet.go

-   AdvanceX
-   ComputeSpans
-   EdgeAt
-   Edges
-   ForEach
-   Insert
-   IsEmpty
-   Len
-   NewCurveAwareAET
-   RemoveExpired
-   RemoveExpiredSubpixel
-   Reset
-   SortByX
-   StepCurves
-   String
-   computeCoverageForWinding
-   emitSpanIfNeeded
-   nonZeroCoverage

## internal/raster/curve_edge.go

-   AsLine
-   BottomY
-   CurveCount
-   CurveCount
-   IsVertical
-   Line
-   Line
-   NewCubicEdge
-   NewCubicEdgeVariant
-   NewLineEdge
-   NewLineEdgeVariant
-   NewQuadraticEdge
-   NewQuadraticEdgeVariant
-   TopY
-   Update
-   Update
-   Update
-   Winding
-   Winding
-   cheapDistance
-   computeDY
-   cubicDeltaFromLine
-   diffToShift
-   newCubicEdgeSetup
-   newQuadraticEdgeSetup
-   update

## internal/raster/edge.go

-   Active
-   Add
-   AddLine
-   Bounds
-   ContainsY
-   Edges
-   Height
-   InsertEdge
-   IsActiveAt
-   Len
-   Len
-   NewEdge
-   NewEdgeList
-   NewEdgeWithWinding
-   NewSimpleAET
-   RemoveExpired
-   Reset
-   Reset
-   SortByX
-   SortByYMin
-   UpdateX
-   XAtY

## internal/raster/edge_builder.go

-   AAShift
-   AllEdges
-   Bounds
-   BuildFromPath
-   ClipRect
-   CubicEdgeCount
-   CubicEdges
-   EdgeCount
-   EmptyRect
-   FlattenCurves
-   IsEmpty
-   IsEmpty
-   LineEdgeCount
-   LineEdges
-   NewEdgeBuilder
-   QuadraticEdgeCount
-   QuadraticEdges
-   Reset
-   SetClipRect
-   SetFlattenCurves
-   TransformPoint
-   VelloLines
-   addCubic
-   addLine
-   addLineUnclipped
-   addQuad
-   clipAndAddLine
-   clipLineX
-   combineVertical
-   curveBBoxInsideClip
-   emitClippedSegments
-   emitSegment
-   findXSplits
-   flattenCubicRecursive
-   flattenCubicToLines
-   flattenQuadRecursive
-   flattenQuadToLines
-   newEmptyBounds
-   unionPoint

## internal/raster/fixed.go

-   FDot16CeilToInt
-   FDot16Div
-   FDot16FastDiv
-   FDot16FloorToInt
-   FDot16FromFloat32
-   FDot16FromFloat64
-   FDot16Mul
-   FDot16RoundToInt
-   FDot16ToFloat32
-   FDot16ToFloat64
-   FDot6CanConvertToFDot16
-   FDot6Ceil
-   FDot6Div
-   FDot6Floor
-   FDot6FromFloat32
-   FDot6FromFloat64
-   FDot6FromInt
-   FDot6Round
-   FDot6SmallScale
-   FDot6ToFDot16
-   FDot6ToFixedDiv2
-   FDot6ToFloat32
-   FDot6ToFloat64
-   FDot6UpShift
-   FDot8FromFDot16
-   absInt32
-   leftShift
-   leftShift64
-   maxInt32
-   saturateInt32

## internal/raster/path_geometry.go

-   ChopCubicAtYExtrema
-   ChopQuadAtYExtrema
-   CubicIsYMonotonic
-   QuadIsYMonotonic
-   ZeroGeomPoint
-   absF32
-   chopCubicAt
-   chopCubicAtSingle
-   chopQuadAt
-   findCubicExtrema
-   findUnitQuadRoots
-   isNotMonotonic
-   lerpPoint
-   maxF32
-   minF32
-   validUnitDivide

## internal/raster/scene_adapter.go

-   IsEmpty
-   NewScenePathAdapter
-   Points
-   Verbs

## internal/stroke/expander.go

-   Add
-   Add
-   Angle
-   Cross
-   DefaultStroke
-   Distance
-   Dot
-   Expand
-   Length
-   LengthSquared
-   Lerp
-   Neg
-   NewStrokeExpander
-   Normalize
-   Perp
-   Scale
-   SetTolerance
-   Sub
-   Sub
-   ToPoint
-   Vec2
-   appendPath
-   appendReversed
-   applyBevelJoin
-   applyCap
-   applyMiterJoin
-   applyRoundJoin
-   arcSegment
-   build
-   close
-   computeMiterPoint
-   cubicTo
-   distanceToLine
-   doCubic
-   doJoin
-   doLine
-   doQuad
-   finish
-   finishClosed
-   flattenCubic
-   flattenCubicRec
-   flattenQuad
-   flattenQuadRec
-   getEndPoint
-   isEmpty
-   isPathElement
-   isPathElement
-   isPathElement
-   isPathElement
-   isPathElement
-   joinWithPrevious
-   lineTo
-   moveTo
-   newPathBuilder
-   quadTo
-   reset
-   roundCap
-   roundJoin
-   roundJoinRev
-   squareCap
-   startFirstSegment
-   transformPoint

## internal/wide/batch.go

-   LoadDst
-   LoadSrc
-   StoreDst

## internal/wide/batch_aa.go

-   BlendBatchAA
-   BlendSolidColorBatchAA
-   BlendSolidColorSpanAA
-   SourceOverBatchAA
-   addClamp
-   mulDiv255Fast

## internal/wide/f32x8.go

-   Add
-   Clamp
-   Div
-   Lerp
-   Max
-   Min
-   Mul
-   SplatF32
-   Sqrt
-   Sub

## internal/wide/u16x16.go

-   Add
-   Clamp
-   Div255
-   Inv
-   Mul
-   MulDiv255
-   SplatU16
-   Sub

## logger.go

-   Enabled
-   Handle
-   Logger
-   SetLogger
-   WithAttrs
-   WithGroup
-   init
-   newNopLogger
-   propagateLogger

## mask.go

-   At
-   Bounds
-   Clear
-   Clone
-   Data
-   Fill
-   Height
-   Invert
-   NewMask
-   NewMaskFromAlpha
-   Set
-   Width

## matrix.go

-   Identity
-   Invert
-   IsIdentity
-   IsScaleOnly
-   IsTranslation
-   IsTranslationOnly
-   MaxScaleFactor
-   Multiply
-   Rotate
-   Scale
-   ScaleFactor
-   Shear
-   TransformPoint
-   TransformVector
-   Translate

## options.go

-   WithPipelineMode
-   WithPixmap
-   WithRenderer
-   defaultOptions

## paint.go

-   Clone
-   ColorAt
-   EffectiveDash
-   EffectiveLineCap
-   EffectiveLineJoin
-   EffectiveLineWidth
-   EffectiveMiterLimit
-   GetBrush
-   GetStroke
-   IsDashed
-   NewPaint
-   SetBrush
-   SetStroke

## painter.go

-   PaintSpan
-   PaintSpan
-   PainterFromPaint

## path.go

-   Arc
-   Circle
-   Clear
-   Clone
-   Close
-   CubicTo
-   CurrentPoint
-   Elements
-   Ellipse
-   HasCurrentPoint
-   LineTo
-   MoveTo
-   NewPath
-   QuadraticTo
-   Rectangle
-   RoundedRectangle
-   Transform
-   arcSegment
-   isEmpty
-   isPathElement
-   isPathElement
-   isPathElement
-   isPathElement
-   isPathElement

## path_builder.go

-   Build
-   BuildPath
-   Circle
-   Close
-   CubicTo
-   Ellipse
-   LineTo
-   MoveTo
-   Path
-   Polygon
-   QuadTo
-   Rect
-   RoundRect
-   Star

## path_ops.go

-   Area
-   BoundingBox
-   Contains
-   Flatten
-   FlattenCallback
-   Length
-   Reversed
-   Winding
-   collectSubpaths
-   cubicArea
-   cubicBBox
-   cubicFlatness
-   cubicLength
-   cubicLengthRecursive
-   cubicWinding
-   expandBBox
-   flattenCubic
-   flattenCubicRecursive
-   flattenCubicWinding
-   flattenCubicWindingRecursive
-   flattenQuad
-   flattenQuadRecursive
-   flattenQuadWinding
-   flattenQuadWindingRecursive
-   getElementStartPoint
-   getSubpathEndpoint
-   isLeft
-   lineArea
-   lineWinding
-   quadArea
-   quadBBox
-   quadLength
-   quadLengthRecursive
-   quadWinding
-   reverseSubpath

## pattern.go

-   ColorAt
-   NewSolidPattern

## pipeline_mode.go

-   SelectPipeline
-   String

## pixmap.go

-   At
-   Bounds
-   Clear
-   ColorModel
-   Data
-   EncodeJPEG
-   EncodePNG
-   FillSpan
-   FillSpanBlend
-   FromImage
-   GetPixel
-   Height
-   NewPixmap
-   SavePNG
-   Set
-   SetPixel
-   SetPixelPremul
-   ToImage
-   Width
-   getPremul
-   setPremul

## point.go

-   Add
-   Cross
-   Distance
-   Div
-   Dot
-   Length
-   LengthSquared
-   Lerp
-   Mul
-   Normalize
-   Pt
-   Rotate
-   Sub

## raster/raster.go

-   init

## rasterizer_mode.go

-   String

## recording/backends/raster/backend.go

-   Begin
-   ClearClip
-   DrawImage
-   DrawText
-   End
-   FillPath
-   FillRect
-   Height
-   Image
-   NewBackend
-   Pixmap
-   Restore
-   Save
-   SavePNG
-   SaveToFile
-   SetClip
-   SetTransform
-   StrokePath
-   Width
-   Write
-   WriteTo
-   applyBrush
-   applyStroke
-   convertFillRule
-   convertLineCap
-   convertLineJoin
-   init
-   setPath
-   setPathFromElements

## recording/brush.go

-   AddColorStop
-   AddColorStop
-   AddColorStop
-   BrushFromGG
-   NewLinearGradientBrush
-   NewPatternBrush
-   NewRadialGradientBrush
-   NewSolidBrush
-   NewSweepGradientBrush
-   SetEndAngle
-   SetExtend
-   SetExtend
-   SetExtend
-   SetFocus
-   SetRepeat
-   SetTransform
-   brushMarker
-   brushMarker
-   brushMarker
-   brushMarker
-   brushMarker
-   linearGradientFromGG
-   radialGradientFromGG
-   sweepGradientFromGG

## recording/command.go

-   Clone
-   DefaultImageOptions
-   DefaultStroke
-   IsValid
-   IsValid
-   IsValid
-   String
-   Type
-   Type
-   Type
-   Type
-   Type
-   Type
-   Type
-   Type
-   Type
-   Type
-   Type
-   Type
-   Type
-   Type
-   Type
-   Type
-   Type
-   Type
-   Type

## recording/matrix.go

-   Contains
-   Determinant
-   Height
-   Identity
-   Inset
-   Intersect
-   Invert
-   IsEmpty
-   IsIdentity
-   IsTranslation
-   Multiply
-   NewRect
-   NewRectFromPoints
-   Offset
-   Rotate
-   Scale
-   ScaleFactor
-   Shear
-   TransformPoint
-   TransformVector
-   Translate
-   Translation
-   Union
-   Width
-   X
-   Y

## recording/pool.go

-   AddBrush
-   AddFont
-   AddImage
-   AddPath
-   BrushCount
-   Clear
-   Clone
-   FontCount
-   GetBrush
-   GetFont
-   GetImage
-   GetPath
-   ImageCount
-   IsValid
-   NewResourcePool
-   PathCount

## recording/recorder.go

-   Clear
-   ClearDash
-   ClearPath
-   ClearWithColor
-   Clip
-   ClipPreserve
-   ClosePath
-   Commands
-   CubicTo
-   DrawArc
-   DrawCircle
-   DrawEllipse
-   DrawEllipticalArc
-   DrawImage
-   DrawImageAnchored
-   DrawImageScaled
-   DrawLine
-   DrawPoint
-   DrawRectangle
-   DrawRoundedRectangle
-   DrawString
-   DrawStringAnchored
-   DrawStringWrapped
-   Fill
-   FillPreserve
-   FillRectangle
-   FillStroke
-   FinishRecording
-   GetCurrentPoint
-   GetTransform
-   Height
-   Height
-   Identity
-   InvertY
-   LineTo
-   MeasureMultilineString
-   MeasureString
-   MoveTo
-   NewRecorder
-   NewSubPath
-   Playback
-   Pop
-   Push
-   QuadraticTo
-   ResetClip
-   Resources
-   Restore
-   Rotate
-   RotateAbout
-   Save
-   Scale
-   SetColor
-   SetDash
-   SetDashOffset
-   SetFillBrush
-   SetFillRGB
-   SetFillRGBA
-   SetFillRule
-   SetFillRuleGG
-   SetFillStyle
-   SetFont
-   SetFontFamily
-   SetFontSize
-   SetHexColor
-   SetLineCap
-   SetLineCapGG
-   SetLineJoin
-   SetLineJoinGG
-   SetLineWidth
-   SetMiterLimit
-   SetRGB
-   SetRGBA
-   SetStrokeBrush
-   SetStrokeRGB
-   SetStrokeRGBA
-   SetStrokeStyle
-   SetTransform
-   Shear
-   Stroke
-   StrokePreserve
-   StrokeRectangle
-   Transform
-   TransformPoint
-   Translate
-   Width
-   Width
-   WordWrap
-   arcSegment
-   drawArcPath
-   recorderSplitLines

## recording/registry.go

-   Backends
-   Count
-   IsRegistered
-   MustBackend
-   NewBackend
-   Register
-   Unregister

## render/device.go

-   Adapter
-   DefaultTextureDescriptor
-   Device
-   Queue
-   SurfaceFormat

## render/gpu_renderer.go

-   Capabilities
-   CreateTextureTarget
-   DeviceHandle
-   Flush
-   NewGPURenderer
-   Render

## render/layers.go

-   Clear
-   ClearLayer
-   Composite
-   CreateLayer
-   Format
-   GetLayer
-   Height
-   Image
-   Layers
-   NewLayeredPixmapTarget
-   Pixels
-   RemoveLayer
-   SetLayerVisible
-   Stride
-   TextureView
-   Width

## render/scene.go

-   Circle
-   Clear
-   ClearDirty
-   ClosePath
-   CommandCount
-   CubicTo
-   DirtyRects
-   Fill
-   HasDirtyRegions
-   Invalidate
-   InvalidateAll
-   IsEmpty
-   IsEmpty
-   LineTo
-   MoveTo
-   NeedsFullRedraw
-   NewScene
-   Points
-   QuadTo
-   Rectangle
-   Reset
-   SetFillColor
-   SetFillRule
-   SetStrokeColor
-   SetStrokeWidth
-   Stroke
-   Verbs
-   drawCommands
-   snapshotPath

## render/software.go

-   Capabilities
-   Flush
-   NewSoftwareRenderer
-   Render
-   addLineStroke
-   ensureFiller
-   expandStroke
-   renderClear
-   renderFill
-   renderStroke
-   sqrt32

## render/target.go

-   Clear
-   Destroy
-   Format
-   Format
-   Format
-   GetPixel
-   Height
-   Height
-   Height
-   Image
-   NewPixmapTarget
-   NewPixmapTargetFromImage
-   NewSurfaceTarget
-   NewTextureTarget
-   Pixels
-   Pixels
-   Pixels
-   Resize
-   SetPixel
-   Stride
-   Stride
-   Stride
-   TextureView
-   TextureView
-   TextureView
-   Width
-   Width
-   Width

## scene/blend_integration.go

-   AdvancedModes
-   AllBlendModes
-   BlendModeFromInternal
-   GetBlendFunc
-   HSLModes
-   PorterDuffModes
-   ToInternalBlendMode

## scene/builder.go

-   Build
-   Clip
-   CurrentTransform
-   DrawLine
-   Fill
-   FillCircle
-   FillPath
-   FillRect
-   FillWith
-   Group
-   Image
-   Layer
-   NewSceneBuilder
-   NewSceneBuilderFrom
-   Reset
-   ResetTransform
-   Rotate
-   Scale
-   Scene
-   Stroke
-   StrokeCircle
-   StrokePath
-   StrokeRect
-   StrokeWith
-   Transform
-   Translate
-   WithTransform

## scene/cache.go

-   Contains
-   DefaultLayerCache
-   EntryCount
-   Get
-   GetVersion
-   Invalidate
-   InvalidateAll
-   MaxSize
-   NewLayerCache
-   Put
-   ResetStats
-   SetMaxSize
-   Size
-   Stats
-   Trim
-   evictUntilSize
-   pixmapSize

## scene/decoder.go

-   Brush
-   CollectPath
-   CubicTo
-   Encoding
-   Fill
-   HasMore
-   Image
-   LineTo
-   MoveTo
-   NewDecoder
-   Next
-   Peek
-   Position
-   PushLayer
-   QuadTo
-   Reset
-   SkipPath
-   Stroke
-   Tag
-   Transform

## scene/encoding.go

-   AffineFromMatrix
-   Append
-   Bounds
-   Brushes
-   Capacity
-   Clone
-   DefaultStrokeStyle
-   DrawData
-   EmptyRect
-   EncodeBeginClip
-   EncodeBrush
-   EncodeEndClip
-   EncodeFill
-   EncodeImage
-   EncodePath
-   EncodePopLayer
-   EncodePushLayer
-   EncodeStroke
-   EncodeTransform
-   EncodeTransformFromMatrix
-   GetBrush
-   Hash
-   Height
-   IdentityAffine
-   IsAdvanced
-   IsEmpty
-   IsEmpty
-   IsHSL
-   IsIdentity
-   IsPorterDuff
-   Multiply
-   NewEncoding
-   NewIterator
-   Next
-   PathCount
-   PathData
-   ReadDrawData
-   ReadPathData
-   ReadTransform
-   Reset
-   Reset
-   RotateAffine
-   ScaleAffine
-   ShapeCount
-   Size
-   SolidBrush
-   String
-   Tags
-   TransformPoint
-   Transforms
-   TranslateAffine
-   Union
-   UnionPoint
-   UpdateBounds
-   Width
-   encodeCubicTo
-   encodeLineTo
-   encodeMoveTo
-   encodeQuadTo
-   max32
-   min32

## scene/filter.go

-   Add
-   Apply
-   ExpandBounds
-   ExpandsOutput
-   IsEmpty
-   Len
-   NewFilterChain
-   String
-   copyPixmap

## scene/layer.go

-   AcquireLayer
-   All
-   CombinedBounds
-   Contains
-   Contains
-   Depth
-   Depth
-   HasClip
-   Intersects
-   Intersects
-   IsClipOnly
-   IsEmpty
-   IsEmpty
-   IsRoot
-   NeedsOffscreen
-   NewClipLayer
-   NewClipStack
-   NewClipState
-   NewFilteredLayer
-   NewLayerStack
-   NewLayerState
-   Pop
-   Pop
-   Push
-   Push
-   ReleaseLayer
-   Reset
-   Reset
-   Reset
-   Root
-   String
-   Top
-   Top
-   UpdateBounds
-   clampAlpha
-   get
-   newLayerPool
-   put

## scene/path.go

-   Arc
-   Bounds
-   Circle
-   Clone
-   Close
-   Contains
-   CubicTo
-   Elements
-   ElementsWithCursor
-   Ellipse
-   Get
-   IsEmpty
-   LineTo
-   MoveTo
-   NewPath
-   NewPathPool
-   PointCount
-   PointCount
-   Points
-   Put
-   QuadTo
-   Rectangle
-   Reset
-   Reverse
-   RoundedRectangle
-   String
-   Transform
-   VerbCount
-   Verbs
-   arcSegment
-   isLeft
-   reverseSubpath
-   windingCubic
-   windingQuad
-   windingSegment

## scene/pool.go

-   Get
-   GetEncoding
-   NewEncodingPool
-   Put
-   PutEncoding
-   Warmup

## scene/renderer.go

-   Cache
-   CacheStats
-   Close
-   DirtyTileCount
-   Height
-   MarkAllDirty
-   MarkDirty
-   NewRenderer
-   Render
-   RenderDirty
-   RenderDirtyWithContext
-   RenderWithContext
-   Resize
-   Stats
-   TileCount
-   Width
-   WithCache
-   WithCacheSize
-   WithTileSize
-   WithWorkers
-   Workers
-   compositeTile
-   compositeTiles
-   convertFillPaint
-   convertPath
-   convertStrokePaint
-   executeEncodingOnTile
-   getPixmap
-   getRenderer
-   putPixmap
-   putRenderer
-   rectIntersects
-   renderTile
-   renderTilesWithContext
-   updateStats

## scene/scene.go

-   Bounds
-   Bounds
-   ClipBounds
-   ClipDepth
-   DrawImage
-   Encoding
-   Fill
-   Flatten
-   Get
-   Images
-   IsEmpty
-   IsEmpty
-   LayerDepth
-   NewImage
-   NewScene
-   NewScenePool
-   PopClip
-   PopLayer
-   PopTransform
-   PushClip
-   PushLayer
-   PushTransform
-   Put
-   Reset
-   Rotate
-   Scale
-   SetTransform
-   Stroke
-   Transform
-   TransformDepth
-   Translate
-   Version
-   Warmup
-   currentEncoding
-   encodeScenePath
-   flattenLayers
-   registerImage
-   transformBounds

## scene/shape.go

-   AddShape
-   Bounds
-   Bounds
-   Bounds
-   Bounds
-   Bounds
-   Bounds
-   Bounds
-   Bounds
-   Bounds
-   Bounds
-   Bounds
-   Bounds
-   Bounds
-   Contains
-   Contains
-   Contains
-   Length
-   NewArcShape
-   NewCircleShape
-   NewCompositeShape
-   NewEllipseShape
-   NewLineShape
-   NewPathShape
-   NewPieShape
-   NewPolygonShape
-   NewRectShape
-   NewRegularPolygonShape
-   NewRoundedRectShape
-   NewStarShape
-   NewTransformShape
-   Point
-   PointCount
-   ShapeCount
-   ToPath
-   ToPath
-   ToPath
-   ToPath
-   ToPath
-   ToPath
-   ToPath
-   ToPath
-   ToPath
-   ToPath
-   ToPath
-   ToPath
-   ToPath

## scene/tag.go

-   DataSize
-   IsClipCommand
-   IsDrawCommand
-   IsLayerCommand
-   IsPathCommand
-   String

## scene/text.go

-   Bounds
-   Config
-   DefaultTextRendererConfig
-   DrawGlyphs
-   DrawText
-   Get
-   NewTextRenderer
-   NewTextRendererPool
-   NewTextRendererWithConfig
-   NewTextShape
-   Put
-   RenderGlyph
-   RenderGlyphs
-   RenderRun
-   RenderText
-   RenderTextToScene
-   RenderToScene
-   SetConfig
-   TextAdvance
-   TextBounds
-   ToCompositePath
-   ToPath
-   appendPath
-   outlineToPath

## sdf.go

-   SDFCircleCoverage
-   SDFFilledCircleCoverage
-   SDFFilledRRectCoverage
-   SDFRRectCoverage
-   sdfRRect
-   smoothstepCoverage

## sdf_accelerator.go

-   CanAccelerate
-   Close
-   FillPath
-   FillShape
-   Flush
-   Init
-   Name
-   SetForceSDF
-   StrokePath
-   StrokeShape
-   blendPixel
-   fillCircleSDF
-   fillEllipseSDF
-   fillRRectSDF
-   getColorFromPaint
-   shapeTooSmallForSDF
-   strokeCircleSDF
-   strokeEllipseSDF
-   strokeRRectSDF

## shape_detect.go

-   DetectShape
-   checkCP
-   detectCircleOrEllipse
-   detectRRect
-   detectRect
-   pointsClose
-   verifyEllipseControlPoints

## shapes.go

-   DrawRegularPolygon

## software.go

-   Fill
-   NewSoftwareRenderer
-   Resize
-   Stroke
-   adaptiveThreshold
-   applyClipCoverage
-   blendAlphaRunsFromCoreRuns
-   blendAlphaRunsFromCoreRunsPaint
-   blendCoveragePaint
-   blendCoverageSolid
-   convertGGPathToCorePath
-   convertLineCap
-   convertLineJoin
-   convertPathToStrokeElements
-   convertStrokeElementsToPath
-   dashCubic
-   dashLine
-   dashPath
-   dashQuad
-   dashStateAtOffset
-   fillCoveragePaintPath
-   fillCoverageSolidPath
-   fillWithCoverageFiller
-   flattenCubicForDash
-   flattenCubicRecForDash
-   flattenQuadForDash
-   flattenQuadRecForDash
-   forcedFiller
-   pathBounds
-   pathEndAt
-   pointLineDistance
-   shouldUseTileRasterizer
-   solidColorFromPaint

## solver.go

-   SolveCubic
-   SolveCubicInUnitInterval
-   SolveQuadratic
-   SolveQuadraticInUnitInterval
-   filterRootsToUnitInterval
-   isFinite
-   solveQuadraticLinear
-   solveQuadraticNormal
-   solveQuadraticOverflow

## stroke.go

-   Bold
-   Clone
-   DashedStroke
-   DefaultStroke
-   DottedStroke
-   IsDashed
-   RoundStroke
-   SquareStroke
-   Thick
-   Thin
-   WithCap
-   WithDash
-   WithDashOffset
-   WithDashPattern
-   WithJoin
-   WithMiterLimit
-   WithWidth

## surface/gpu_surface.go

-   Backend
-   Capabilities
-   Clear
-   Close
-   DrawImage
-   Fill
-   Flush
-   Height
-   NewGPUSurface
-   Snapshot
-   Stroke
-   Width

## surface/image_surface.go

-   Capabilities
-   Clear
-   Close
-   DrawImage
-   Fill
-   Flush
-   Height
-   Image
-   NewImageSurface
-   NewImageSurfaceFromImage
-   Snapshot
-   Stroke
-   Width
-   applyAlpha
-   blendPixel
-   blendPixelAlpha
-   blendRow
-   buildStrokeOutline
-   expandStroke
-   flattenCubicForStroke
-   flattenCubicRecursive
-   flattenQuadForStroke
-   flattenQuadRecursive
-   resolveColor
-   sqrtf32

## surface/path.go

-   Arc
-   Bounds
-   Circle
-   Clear
-   Clone
-   Close
-   CubicTo
-   CurrentPoint
-   Ellipse
-   IsEmpty
-   LineTo
-   MoveTo
-   NewPath
-   Points
-   QuadTo
-   Rectangle
-   RoundedRectangle
-   Verbs
-   arcSegment

## surface/registry.go

-   Available
-   Available
-   Error
-   Error
-   Get
-   Get
-   List
-   List
-   NewRegistry
-   NewSurface
-   NewSurface
-   NewSurfaceByName
-   NewSurfaceByName
-   NewSurfaceByNameWithOptions
-   NewSurfaceWithOptions
-   Register
-   Register
-   Unregister
-   Unregister
-   init
-   sortedNames

## surface/types.go

-   ColorAt
-   DefaultDrawImageOptions
-   DefaultFillStyle
-   DefaultOptions
-   DefaultStrokeStyle
-   IsDashed
-   Pt
-   WithCap
-   WithColor
-   WithColor
-   WithDash
-   WithJoin
-   WithMiterLimit
-   WithPattern
-   WithPattern
-   WithRule
-   WithWidth

## text.go

-   DrawString
-   DrawStringAnchored
-   DrawStringWrapped
-   Font
-   LoadFontFace
-   MeasureMultilineString
-   MeasureString
-   SetFont
-   WordWrap
-   drawStringAsOutlines
-   drawStringBitmap
-   drawStringCPU
-   drawStringScaled
-   ensureOutlineExtractor
-   fontHeight
-   splitLines
-   tryGPUText

## text/cache.go

-   Clear
-   Get
-   GetOrCreate
-   Len
-   NewCache
-   Set
-   evictOldest

## text/cache/lru.go

-   Clear
-   Len
-   MoveToFront
-   Oldest
-   PushFront
-   Remove
-   RemoveOldest
-   newLRUList
-   unlink

## text/cache/shaping.go

-   Capacity
-   Clear
-   DefaultShapingCache
-   Delete
-   Get
-   GetOrCreate
-   HashFeatures
-   Len
-   NewShapingCache
-   NewShapingKey
-   ResetStats
-   Set
-   ShardLen
-   Stats
-   TotalCapacity
-   getShard
-   hashString
-   keyHash

## text/color_font.go

-   DetectGlyphType
-   GetBitmapGlyph
-   GetCOLRGlyph
-   HasAnyColorTable
-   PreferredColorFormat

## text/draw.go

-   Draw
-   Measure
-   drawFilteredFace
-   drawMultiFace
-   drawSourceFace
-   mapHinting

## text/draw_emoji.go

-   Clear
-   DrawWithEmoji
-   Get
-   NewBitmapGlyphCache
-   Put
-   Size
-   drawBitmapGlyph
-   drawSingleOutlineGlyph
-   drawWithColorSupport

## text/emoji/bitmap.go

-   BestStrikeForPPEM
-   Decode
-   GetGlyph
-   HasGlyph
-   HasTable
-   NewCBDTParser
-   NewSBIXParser
-   NumStrikes
-   StrikePPEM
-   String
-   absDiff
-   parse
-   parseGraphicType
-   parseStrike

## text/emoji/cbdt_extractor.go

-   AvailablePPEMs
-   GetGlyph
-   GetGlyphAtStrike
-   GetGlyphWithStrategy
-   HasGlyph
-   HasGlyphInStrike
-   NewCBDTExtractor
-   NumStrikes
-   SelectStrike
-   StrikeBitDepth
-   StrikePPEM
-   String
-   calculateGlyphLocation
-   extractGlyphFromSubtable
-   extractImageData
-   hasGlyphInStrike
-   parseBigGlyphMetrics
-   parseBitmapSizeRecord
-   parseCBLC
-   parseIndexSubtable
-   parseIndexSubtables
-   parseSbitLineMetrics
-   parseSmallGlyphMetrics

## text/emoji/colr.go

-   Empty
-   GetGlyph
-   HasGlyph
-   Height
-   IsForeground
-   NewCOLRParser
-   NumPalettes
-   PaletteColors
-   RGBA
-   RenderCOLRToImage
-   ToRGBA
-   Width
-   findBaseGlyph
-   parseBaseGlyphs
-   parseCOLRHeader
-   parseCPAL
-   parseLayers

## text/emoji/detect.go

-   IsBlackFlag
-   IsCancelTag
-   IsCombiningEnclosingKeycap
-   IsEmoji
-   IsEmojiModifier
-   IsEmojiModifierBase
-   IsEmojiPresentation
-   IsEmojiVariation
-   IsKeycapBase
-   IsRegionalIndicator
-   IsTagCharacter
-   IsTextPresentation
-   IsVariationSelector
-   IsZWJ
-   Segment
-   appendOrMergeRun
-   computeByteOffsets
-   isEmojiBase
-   isEmojiComponent
-   isEmojiOrEmojiComponent
-   isEmojiPresentation
-   isTextPresentationEmoji
-   parseEmojiSequence
-   parseEmojiSequenceAfterZWJ
-   parseExtendedSequence
-   parseKeycapSequence
-   parseTagSequence

## text/emoji/sequence.go

-   GetFlagCode
-   GetSkinTone
-   GetTagSequenceCode
-   HasModifier
-   IsValidSequence
-   Len
-   Normalize
-   Parse
-   ParseString
-   SkinToneRune
-   String
-   String
-   String
-   parseAfterZWJ
-   parseExtendedSequenceAt
-   parseKeycapSequenceAt
-   parseSequenceAt
-   parseTagSequenceAt

## text/errors.go

-   Error

## text/face.go

-   Advance
-   AppendGlyphs
-   Direction
-   Glyphs
-   HasGlyph
-   Metrics
-   Size
-   Source
-   private

## text/filtered.go

-   Advance
-   AppendGlyphs
-   Contains
-   Direction
-   Glyphs
-   HasGlyph
-   Metrics
-   NewFilteredFace
-   Size
-   Source
-   inRanges
-   private

## text/glyph_cache.go

-   Clear
-   CurrentFrame
-   DefaultGlyphCacheConfig
-   Delete
-   Get
-   Get
-   GetGlobalGlyphCache
-   GetOrCreate
-   HitRate
-   Len
-   Maintain
-   NewGlyphCache
-   NewGlyphCachePool
-   NewGlyphCacheWithConfig
-   Put
-   ResetStats
-   Set
-   SetGlobalGlyphCache
-   Stats
-   addToFront
-   getShard
-   moveToFront
-   remove
-   removeTail

## text/glyph_outline.go

-   Clone
-   Error
-   ExtractOutline
-   IdentityTransform
-   IsEmpty
-   Multiply
-   NewOutlineExtractor
-   Scale
-   ScaleTransform
-   SegmentCount
-   String
-   Transform
-   TransformPoint
-   Translate
-   TranslateTransform
-   extractFromSFNT
-   fixedPointToOutline
-   getGlyphAdvance
-   updateBounds
-   updateMinMax

## text/glyph_renderer.go

-   Cache
-   DefaultRenderParams
-   GetGlobalTextRenderer
-   GlyphRenderer
-   NewGlyphRenderer
-   NewGlyphRendererWithCache
-   NewTextRenderer
-   RenderGlyph
-   RenderGlyphs
-   RenderLayout
-   RenderRun
-   RotateTransform
-   ScaleTransformXY
-   SetCache
-   SetDefaultColor
-   SetDefaultFace
-   SetDefaultSize
-   ShapeAndRender
-   ShapeAndRenderAt
-   WithColor
-   WithOpacity
-   WithTransform
-   clampOpacity
-   computeFontID
-   cosf64
-   cosineApprox
-   sineApprox
-   sinf64
-   sizeToInt16
-   transformOutline

## text/glyph_run.go

-   AddGlyph
-   AddShapedGlyph
-   AddShapedGlyphs
-   AddShapedRun
-   Build
-   BuildTransformed
-   Cache
-   Clear
-   Get
-   Instances
-   Len
-   NewGlyphRunBuilder
-   NewGlyphRunBuilderPool
-   Put
-   SetCache
-   float32ToSizeBits
-   sizeBitsToFloat32

## text/glyph_type.go

-   Has
-   String
-   String

## text/layout.go

-   DefaultLayoutOptions
-   Height
-   LayoutText
-   LayoutTextSimple
-   LayoutTextWithContext
-   String
-   applyAlignment
-   calculateBreakPosition
-   createLine
-   createLineFromGlyphs
-   isCJK
-   isWordBreakRune
-   layoutParagraph
-   markBreakOpportunities
-   shapeSegments
-   splitParagraphs
-   wrapLinesWithMode

## text/metrics.go

-   LineHeight

## text/msdf/atlas.go

-   AtlasCount
-   AtlasInfos
-   Clear
-   Clear
-   Compact
-   Config
-   DefaultAtlasConfig
-   DirtyAtlases
-   Error
-   Error
-   Generator
-   Get
-   Get
-   GetAtlas
-   GetBatch
-   GlyphCount
-   GlyphCount
-   GlyphCount
-   HasGlyph
-   HasGlyph
-   IsDirty
-   IsFull
-   MarkAllClean
-   MarkClean
-   MemoryUsage
-   MemoryUsage
-   NewAtlasManager
-   NewAtlasManagerDefault
-   NewConcurrentAtlasManager
-   Remove
-   SetGenerator
-   Stats
-   Stats
-   Utilization
-   Validate
-   copyMSDF
-   findOrCreateAtlas
-   getShard
-   newAtlas

## text/msdf/contour.go

-   AddContour
-   AddEdge
-   AssignColors
-   Bounds
-   CalculateBounds
-   CalculateWinding
-   Clone
-   EdgeCount
-   FromOutline
-   IsClockwise
-   NewContour
-   NewShape
-   SelectBlue
-   SelectGreen
-   SelectRed
-   SwitchColor
-   Validate
-   assignContourColors

## text/msdf/edge.go

-   Bounds
-   Clone
-   DirectionAt
-   EndPoint
-   HasBlue
-   HasGreen
-   HasRed
-   NewCubicEdge
-   NewLinearEdge
-   NewQuadraticEdge
-   PointAt
-   SignedDistance
-   StartPoint
-   String
-   String
-   cbrt
-   cubicBounds
-   cubicDerivative
-   cubicSecondDerivative
-   cubicSignedDistance
-   evaluateCubic
-   evaluateQuadratic
-   linearBounds
-   linearSignedDistance
-   newtonRefineCubic
-   quadraticBounds
-   quadraticDerivative
-   quadraticSignedDistance
-   solveCubic
-   solveCubicCardano
-   solveCubicOneRoot
-   solveCubicRepeatedRoots
-   solveCubicThreeRoots
-   solveLinear
-   solveQuadratic
-   solveQuadraticFull

## text/msdf/generator.go

-   Config
-   DefaultGenerator
-   ErrorCorrection
-   Generate
-   Generate
-   GenerateBatch
-   GenerateWithMetrics
-   Get
-   MedianFilter
-   NewGenerator
-   NewGeneratorPool
-   Put
-   SetConfig
-   calculateScale
-   channelDistance
-   clampInt
-   collectNeighborhood
-   correctChannel
-   distanceToPixel
-   generateDistanceField
-   generateEmpty
-   median3Byte
-   median9
-   processRows

## text/msdf/shelf.go

-   Allocate
-   Allocate
-   AllocateFixed
-   Allocated
-   CanFit
-   Capacity
-   CellSize
-   CurrentShelfRemainingWidth
-   GridDimensions
-   IsFull
-   NewGridAllocator
-   NewShelfAllocator
-   Remaining
-   RemainingHeight
-   Reset
-   Reset
-   ShelfCount
-   TotalArea
-   UsedArea
-   Utilization
-   Utilization

## text/msdf/types.go

-   Add
-   Angle
-   AngleBetween
-   Center
-   Combine
-   Contains
-   Cross
-   DefaultConfig
-   Dot
-   Error
-   Expand
-   GetPixel
-   Height
-   Infinite
-   IsCloserThan
-   IsEmpty
-   Length
-   LengthSquared
-   Lerp
-   Mul
-   NewSignedDistance
-   Normalized
-   OutlineToPixel
-   Perpendicular
-   PixelOffset
-   PixelToOutline
-   SetPixel
-   Sub
-   Union
-   Validate
-   Width

## text/multi.go

-   Advance
-   AppendGlyphs
-   Direction
-   Glyphs
-   HasGlyph
-   Metrics
-   NewMultiFace
-   Size
-   Source
-   faceForRune
-   private

## text/options.go

-   WithCacheLimit
-   WithDirection
-   WithHinting
-   WithLanguage
-   WithParser
-   defaultFaceConfig
-   defaultSourceConfig

## text/parser.go

-   Height
-   RegisterParser
-   getParser

## text/parser_ximage.go

-   FullName
-   GlyphAdvance
-   GlyphBounds
-   GlyphIndex
-   Metrics
-   Name
-   NumGlyphs
-   Parse
-   UnitsPerEm
-   fixedToFloat64

## text/rasterize.go

-   RasterizeGlyph

## text/runetobool.go

-   Clear
-   Get
-   NewRuneToBoolMap
-   Set

## text/script.go

-   DetectScript
-   IsRTL
-   RequiresComplexShaping
-   String
-   detectASCII
-   detectCommonSymbols
-   detectEastAsian
-   detectEuropean
-   detectLatin1Supplement
-   detectMiddleEastern
-   detectOther
-   detectSouthAsian
-   detectSoutheastAsian

## text/segment.go

-   IsPunctuation
-   IsWhitespace
-   NewBuiltinSegmenter
-   NewBuiltinSegmenterWithDirection
-   RuneCount
-   Segment
-   SegmentText
-   SegmentTextRTL
-   buildSegments
-   computeBidiLevels
-   computeByteOffsets
-   detectScripts
-   findNextConcreteScript
-   makeSegment
-   resolveCommonScript
-   resolveInheritedScripts

## text/shaped_run.go

-   Bounds
-   Height
-   LineHeight
-   Width

## text/shaper.go

-   GetShaper
-   SetShaper
-   Shape

## text/shaper_builtin.go

-   Shape

## text/shaper_gotext.go

-   ClearCache
-   NewGoTextShaper
-   RemoveSource
-   Shape
-   convertGlyphs
-   detectScript
-   fixedToFloat
-   floatToFixed
-   getOrCreateFont
-   mapDirection

## text/source.go

-   Close
-   Face
-   Name
-   NewFontSource
-   NewFontSourceFromFile
-   Parsed
-   copyCheck
-   extractFontName

## text/subpixel.go

-   CacheMultiplier
-   Clear
-   Config
-   DefaultSubpixelConfig
-   Delete
-   Divisions
-   Get
-   GetGlobalSubpixelCache
-   GetOrCreate
-   HighQualitySubpixelConfig
-   HitRate
-   IsEnabled
-   IsEnabled
-   Len
-   Maintain
-   MakeSubpixelKey
-   NewSubpixelCache
-   NewSubpixelCacheWithConfig
-   NoSubpixelConfig
-   Quantize
-   QuantizePoint
-   ResetStats
-   Set
-   SetConfig
-   SetGlobalSubpixelCache
-   Stats
-   String
-   SubpixelOffset
-   SubpixelOffsets
-   toOutlineKey

## text/tab.go

-   SetTabWidth
-   TabWidth
-   expandTabs
-   fixTabGlyphs
-   spaceGIDAndAdvance
-   tabAdvance

## text/types.go

-   Empty
-   Height
-   IsHorizontal
-   IsVertical
-   String
-   String
-   Width

## text/wrap.go

-   MeasureText
-   String
-   WrapText
-   calculateLineBreakPosition
-   canBreakAt
-   classifyRune
-   classifySpecificRune
-   computeBreak
-   computeWordBreak
-   findBreakOpportunities
-   findLineEnd
-   findRuneIndexForCluster
-   isBreakBetweenCategories
-   isCJKRune
-   markBreakOpportunitiesEnhanced
-   measureRune
-   mustBreakAt
-   newWrapTextInfo
-   runeToByteOffset
-   substring
-   wrapParagraph

## vec.go

-   Add
-   Angle
-   Approx
-   Atan2
-   Cross
-   Div
-   Dot
-   IsZero
-   Length
-   LengthSq
-   Lerp
-   Mul
-   Neg
-   Normalize
-   Perp
-   PointToVec2
-   Rotate
-   Sub
-   ToPoint
-   V2

