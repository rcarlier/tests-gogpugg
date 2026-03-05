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

## context_clip.go

-   Clip
-   ClipPreserve
-   ClipRect
-   ResetClip

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

## context_layer.go

-   PopLayer
-   PushLayer
-   SetBlendMode

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

## gpu/gpu.go

-   SetDeviceProvider

## gradient_linear.go

-   AddColorStop
-   ColorAt
-   NewLinearGradientBrush
-   SetExtend

## gradient_radial.go

-   AddColorStop
-   ColorAt
-   NewRadialGradientBrush
-   SetExtend
-   SetFocus

## gradient_sweep.go

-   AddColorStop
-   ColorAt
-   NewSweepGradientBrush
-   SetEndAngle
-   SetExtend

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

## integration/ggcanvas/render.go

-   DefaultRenderOptions
-   RenderTo
-   RenderToEx
-   RenderToPosition
-   RenderToScaled

## logger.go

-   Enabled
-   Handle
-   Logger
-   SetLogger
-   WithAttrs
-   WithGroup

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

## render/software.go

-   Capabilities
-   Flush
-   NewSoftwareRenderer
-   Render

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

## scene/filter.go

-   Add
-   Apply
-   ExpandBounds
-   ExpandsOutput
-   IsEmpty
-   Len
-   NewFilterChain
-   String

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

## sdf.go

-   SDFCircleCoverage
-   SDFFilledCircleCoverage
-   SDFFilledRRectCoverage
-   SDFRRectCoverage

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

## shape_detect.go

-   DetectShape

## shapes.go

-   DrawRegularPolygon

## software.go

-   Fill
-   NewSoftwareRenderer
-   Resize
-   Stroke

## solver.go

-   SolveCubic
-   SolveCubicInUnitInterval
-   SolveQuadratic
-   SolveQuadraticInUnitInterval

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

## text/cache.go

-   Clear
-   Get
-   GetOrCreate
-   Len
-   NewCache
-   Set

## text/cache/lru.go

-   Clear
-   Len
-   MoveToFront
-   Oldest
-   PushFront
-   Remove
-   RemoveOldest

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

## text/color_font.go

-   DetectGlyphType
-   GetBitmapGlyph
-   GetCOLRGlyph
-   HasAnyColorTable
-   PreferredColorFormat

## text/draw.go

-   Draw
-   Measure

## text/draw_emoji.go

-   Clear
-   DrawWithEmoji
-   Get
-   NewBitmapGlyphCache
-   Put
-   Size

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

## text/options.go

-   WithCacheLimit
-   WithDirection
-   WithHinting
-   WithLanguage
-   WithParser

## text/parser.go

-   Height
-   RegisterParser

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

## text/segment.go

-   IsPunctuation
-   IsWhitespace
-   NewBuiltinSegmenter
-   NewBuiltinSegmenterWithDirection
-   RuneCount
-   Segment
-   SegmentText
-   SegmentTextRTL

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

## text/source.go

-   Close
-   Face
-   Name
-   NewFontSource
-   NewFontSourceFromFile
-   Parsed

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

## text/tab.go

-   SetTabWidth
-   TabWidth

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

