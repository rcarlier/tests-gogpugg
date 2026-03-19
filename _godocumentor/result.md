# github.com/gogpu/gg

## accelerator.go

-   `func Accelerator() GPUAccelerator`
-   `func CloseAccelerator()`
-   `func RegisterAccelerator(a GPUAccelerator) error`
-   `func SetAcceleratorDeviceProvider(provider any) error`
-   `func SetAcceleratorSurfaceTarget(view any, width, height uint32)`

## brush.go

-   `func BrushFromPattern(p Pattern) Brush`
-   `func (SolidBrush) (b SolidBrush) ColorAt(_, _ float64) RGBA`
-   `func (*brushPattern) (p *brushPattern) ColorAt(x, y float64) RGBA`
-   `func (SolidBrush) (b SolidBrush) Lerp(other SolidBrush, t float64) SolidBrush`
-   `func (SolidBrush) (b SolidBrush) Opaque() SolidBrush`
-   `func PatternFromBrush(b Brush) Pattern`
-   `func Solid(c RGBA) SolidBrush`
-   `func SolidHex(hex string) SolidBrush`
-   `func SolidRGB(r, g, b float64) SolidBrush`
-   `func SolidRGBA(r, g, b, a float64) SolidBrush`
-   `func (SolidBrush) (b SolidBrush) Transparent() SolidBrush`
-   `func (SolidBrush) (b SolidBrush) WithAlpha(alpha float64) SolidBrush`

## brush_custom.go

-   `func Checkerboard(c0, c1 RGBA, size float64) CustomBrush`
-   `func (CustomBrush) (b CustomBrush) ColorAt(x, y float64) RGBA`
-   `func HorizontalGradient(c0, c1 RGBA, x0, x1 float64) CustomBrush`
-   `func LinearGradient(c0, c1 RGBA, x0, y0, x1, y1 float64) CustomBrush`
-   `func NewCustomBrush(fn ColorFunc) CustomBrush`
-   `func RadialGradient(c0, c1 RGBA, cx, cy, r float64) CustomBrush`
-   `func Stripes(c0, c1 RGBA, width, angle float64) CustomBrush`
-   `func VerticalGradient(c0, c1 RGBA, y0, y1 float64) CustomBrush`
-   `func (CustomBrush) (b CustomBrush) WithName(name string) CustomBrush`

## color.go

-   `func (RGBA) (c RGBA) Color() color.Color`
-   `func FromColor(c color.Color) RGBA`
-   `func HSL(h, s, l float64) RGBA`
-   `func Hex(hex string) RGBA`
-   `func (RGBA) (c RGBA) Lerp(other RGBA, t float64) RGBA`
-   `func (RGBA) (c RGBA) Premultiply() RGBA`
-   `func RGB(r, g, b float64) RGBA`
-   `func (RGBA) (c RGBA) RGBA() (r, g, b, a uint32)`
-   `func RGBA2(r, g, b, a float64) RGBA`
-   `func (RGBA) (c RGBA) Unpremultiply() RGBA`

## context.go

-   `func (*Context) (c *Context) Clear()`
-   `func (*Context) (c *Context) ClearDash()`
-   `func (*Context) (c *Context) ClearPath()`
-   `func (*Context) (c *Context) ClearWithColor(col RGBA)`
-   `func (*Context) (c *Context) Close() error`
-   `func (*Context) (c *Context) ClosePath()`
-   `func (*Context) (c *Context) CubicTo(c1x, c1y, c2x, c2y, x, y float64)`
-   `func (*Context) (c *Context) DrawArc(x, y, r, angle1, angle2 float64)`
-   `func (*Context) (c *Context) DrawCircle(x, y, r float64)`
-   `func (*Context) (c *Context) DrawEllipse(x, y, rx, ry float64)`
-   `func (*Context) (c *Context) DrawEllipticalArc(x, y, rx, ry, angle1, angle2 float64)`
-   `func (*Context) (c *Context) DrawLine(x1, y1, x2, y2 float64)`
-   `func (*Context) (c *Context) DrawPoint(x, y, r float64)`
-   `func (*Context) (c *Context) DrawRectangle(x, y, w, h float64)`
-   `func (*Context) (c *Context) DrawRoundedRectangle(x, y, w, h, r float64)`
-   `func (*Context) (c *Context) EncodeJPEG(w io.Writer, quality int) error`
-   `func (*Context) (c *Context) EncodePNG(w io.Writer) error`
-   `func (*Context) (c *Context) Fill() error`
-   `func (*Context) (c *Context) FillBrush() Brush`
-   `func (*Context) (c *Context) FillPreserve() error`
-   `func (*Context) (c *Context) FlushGPU() error`
-   `func (*Context) (c *Context) GetCurrentPoint() (x, y float64, ok bool)`
-   `func (*Context) (c *Context) GetStroke() Stroke`
-   `func (*Context) (c *Context) GetTransform() Matrix`
-   `func (*Context) (c *Context) Height() int`
-   `func (*Context) (c *Context) Identity()`
-   `func (*Context) (c *Context) Image() image.Image`
-   `func (*Context) (c *Context) InvertY()`
-   `func (*Context) (c *Context) IsDashed() bool`
-   `func (*Context) (c *Context) LineTo(x, y float64)`
-   `func (*Context) (c *Context) MoveTo(x, y float64)`
-   `func NewContext(width, height int, opts ...ContextOption) *Context`
-   `func NewContextForImage(img image.Image, opts ...ContextOption) *Context`
-   `func (*Context) (c *Context) NewSubPath()`
-   `func (*Context) (c *Context) PipelineMode() PipelineMode`
-   `func (*Context) (c *Context) Pop()`
-   `func (*Context) (c *Context) Push()`
-   `func (*Context) (c *Context) QuadraticTo(cx, cy, x, y float64)`
-   `func (*Context) (c *Context) RasterizerMode() RasterizerMode`
-   `func (*Context) (c *Context) Resize(width, height int) error`
-   `func (*Context) (c *Context) ResizeTarget() *Pixmap`
-   `func (*Context) (c *Context) Rotate(angle float64)`
-   `func (*Context) (c *Context) RotateAbout(angle, x, y float64)`
-   `func (*Context) (c *Context) SavePNG(path string) error`
-   `func (*Context) (c *Context) Scale(x, y float64)`
-   `func (*Context) (c *Context) SetColor(col color.Color)`
-   `func (*Context) (c *Context) SetDash(lengths ...float64)`
-   `func (*Context) (c *Context) SetDashOffset(offset float64)`
-   `func (*Context) (c *Context) SetFillBrush(b Brush)`
-   `func (*Context) (c *Context) SetFillRule(rule FillRule)`
-   `func (*Context) (c *Context) SetHexColor(hex string)`
-   `func (*Context) (c *Context) SetLineCap(lineCap LineCap)`
-   `func (*Context) (c *Context) SetLineJoin(join LineJoin)`
-   `func (*Context) (c *Context) SetLineWidth(width float64)`
-   `func (*Context) (c *Context) SetMiterLimit(limit float64)`
-   `func (*Context) (c *Context) SetPipelineMode(mode PipelineMode)`
-   `func (*Context) (c *Context) SetPixel(x, y int, col RGBA)`
-   `func (*Context) (c *Context) SetRGB(r, g, b float64)`
-   `func (*Context) (c *Context) SetRGBA(r, g, b, a float64)`
-   `func (*Context) (c *Context) SetRasterizerMode(mode RasterizerMode)`
-   `func (*Context) (c *Context) SetStroke(stroke Stroke)`
-   `func (*Context) (c *Context) SetStrokeBrush(b Brush)`
-   `func (*Context) (c *Context) SetTransform(m Matrix)`
-   `func (*Context) (c *Context) Shear(x, y float64)`
-   `func (*Context) (c *Context) Stroke() error`
-   `func (*Context) (c *Context) StrokeBrush() Brush`
-   `func (*Context) (c *Context) StrokePreserve() error`
-   `func (*Context) (c *Context) Transform(m Matrix)`
-   `func (*Context) (c *Context) TransformPoint(x, y float64) (float64, float64)`
-   `func (*Context) (c *Context) Translate(x, y float64)`
-   `func (*Context) (c *Context) Width() int`

## context_clip.go

-   `func (*Context) (c *Context) Clip()`
-   `func (*Context) (c *Context) ClipPreserve()`
-   `func (*Context) (c *Context) ClipRect(x, y, w, h float64)`
-   `func (*Context) (c *Context) ResetClip()`

## context_image.go

-   `func (*ImagePattern) (p *ImagePattern) ColorAt(x, y float64) RGBA`
-   `func (*Context) (c *Context) CreateImagePattern(img *ImageBuf, x, y, w, h int) Pattern`
-   `func (*Context) (c *Context) DrawImage(img *ImageBuf, x, y float64)`
-   `func (*Context) (c *Context) DrawImageCircular(img *ImageBuf, cx, cy, radius float64)`
-   `func (*Context) (c *Context) DrawImageEx(img *ImageBuf, opts DrawImageOptions)`
-   `func (*Context) (c *Context) DrawImageRounded(img *ImageBuf, x, y, radius float64)`
-   `func ImageBufFromImage(img image.Image) *ImageBuf`
-   `func LoadImage(path string) (*ImageBuf, error)`
-   `func LoadWebP(path string) (*ImageBuf, error)`
-   `func NewImageBuf(width, height int, format ImageFormat) (*ImageBuf, error)`
-   `func (*ImagePattern) (p *ImagePattern) SetAnchor(x, y float64)`
-   `func (*ImagePattern) (p *ImagePattern) SetClamp(clamp bool)`
-   `func (*Context) (c *Context) SetFillPattern(pattern Pattern)`
-   `func (*ImagePattern) (p *ImagePattern) SetOpacity(opacity float64)`
-   `func (*ImagePattern) (p *ImagePattern) SetScale(sx, sy float64)`
-   `func (*Context) (c *Context) SetStrokePattern(pattern Pattern)`

## context_layer.go

-   `func (*Context) (c *Context) PopLayer()`
-   `func (*Context) (c *Context) PushLayer(blendMode BlendMode, opacity float64)`
-   `func (*Context) (c *Context) SetBlendMode(_ BlendMode)`

## context_mask.go

-   `func (*Context) (c *Context) AsMask() *Mask`
-   `func (*Context) (c *Context) ClearMask()`
-   `func (*Context) (c *Context) GetMask() *Mask`
-   `func (*Context) (c *Context) InvertMask()`
-   `func (*Context) (c *Context) SetMask(mask *Mask)`

## coverage_filler.go

-   `func GetCoverageFiller() CoverageFiller`
-   `func RegisterCoverageFiller(f CoverageFiller)`

## curve.go

-   `func (Line) (l Line) BoundingBox() Rect`
-   `func (CubicBez) (c CubicBez) BoundingBox() Rect`
-   `func (QuadBez) (q QuadBez) BoundingBox() Rect`
-   `func (Rect) (r Rect) Contains(p Point) bool`
-   `func (CubicBez) (c CubicBez) Deriv() QuadBez`
-   `func (CubicBez) (c CubicBez) End() Point`
-   `func (Line) (l Line) End() Point`
-   `func (QuadBez) (q QuadBez) End() Point`
-   `func (QuadBez) (q QuadBez) Eval(t float64) Point`
-   `func (CubicBez) (c CubicBez) Eval(t float64) Point`
-   `func (Line) (l Line) Eval(t float64) Point`
-   `func (CubicBez) (c CubicBez) Extrema() []float64`
-   `func (QuadBez) (q QuadBez) Extrema() []float64`
-   `func (Rect) (r Rect) Height() float64`
-   `func (CubicBez) (c CubicBez) Inflections() []float64`
-   `func (Line) (l Line) Length() float64`
-   `func (Line) (l Line) Midpoint() Point`
-   `func NewCubicBez(p0, p1, p2, p3 Point) CubicBez`
-   `func NewLine(p0, p1 Point) Line`
-   `func NewQuadBez(p0, p1, p2 Point) QuadBez`
-   `func NewRect(p1, p2 Point) Rect`
-   `func (CubicBez) (c CubicBez) Normal(t float64) Vec2`
-   `func (QuadBez) (q QuadBez) Raise() CubicBez`
-   `func (Line) (l Line) Reversed() Line`
-   `func (Line) (l Line) Start() Point`
-   `func (QuadBez) (q QuadBez) Start() Point`
-   `func (CubicBez) (c CubicBez) Start() Point`
-   `func (QuadBez) (q QuadBez) Subdivide() (QuadBez, QuadBez)`
-   `func (CubicBez) (c CubicBez) Subdivide() (CubicBez, CubicBez)`
-   `func (Line) (l Line) Subdivide() (Line, Line)`
-   `func (CubicBez) (c CubicBez) Subsegment(t0, t1 float64) CubicBez`
-   `func (Line) (l Line) Subsegment(t0, t1 float64) Line`
-   `func (QuadBez) (q QuadBez) Subsegment(t0, t1 float64) QuadBez`
-   `func (CubicBez) (c CubicBez) Tangent(t float64) Vec2`
-   `func (Rect) (r Rect) Union(other Rect) Rect`
-   `func (Rect) (r Rect) Width() float64`

## dash.go

-   `func (*Dash) (d *Dash) Clone() *Dash`
-   `func (*Dash) (d *Dash) IsDashed() bool`
-   `func NewDash(lengths ...float64) *Dash`
-   `func (*Dash) (d *Dash) NormalizedOffset() float64`
-   `func (*Dash) (d *Dash) PatternLength() float64`
-   `func (*Dash) (d *Dash) Scale(factor float64) *Dash`
-   `func (*Dash) (d *Dash) WithOffset(offset float64) *Dash`

## gpu/gpu.go

-   `func SetDeviceProvider(provider any) error`

## gradient_linear.go

-   `func (*LinearGradientBrush) (g *LinearGradientBrush) AddColorStop(offset float64, c RGBA) *LinearGradientBrush`
-   `func (*LinearGradientBrush) (g *LinearGradientBrush) ColorAt(x, y float64) RGBA`
-   `func NewLinearGradientBrush(x0, y0, x1, y1 float64) *LinearGradientBrush`
-   `func (*LinearGradientBrush) (g *LinearGradientBrush) SetExtend(mode ExtendMode) *LinearGradientBrush`

## gradient_radial.go

-   `func (*RadialGradientBrush) (g *RadialGradientBrush) AddColorStop(offset float64, c RGBA) *RadialGradientBrush`
-   `func (*RadialGradientBrush) (g *RadialGradientBrush) ColorAt(x, y float64) RGBA`
-   `func NewRadialGradientBrush(cx, cy, startRadius, endRadius float64) *RadialGradientBrush`
-   `func (*RadialGradientBrush) (g *RadialGradientBrush) SetExtend(mode ExtendMode) *RadialGradientBrush`
-   `func (*RadialGradientBrush) (g *RadialGradientBrush) SetFocus(fx, fy float64) *RadialGradientBrush`

## gradient_sweep.go

-   `func (*SweepGradientBrush) (g *SweepGradientBrush) AddColorStop(offset float64, c RGBA) *SweepGradientBrush`
-   `func (*SweepGradientBrush) (g *SweepGradientBrush) ColorAt(x, y float64) RGBA`
-   `func NewSweepGradientBrush(cx, cy, startAngle float64) *SweepGradientBrush`
-   `func (*SweepGradientBrush) (g *SweepGradientBrush) SetEndAngle(endAngle float64) *SweepGradientBrush`
-   `func (*SweepGradientBrush) (g *SweepGradientBrush) SetExtend(mode ExtendMode) *SweepGradientBrush`

## integration/ggcanvas/canvas.go

-   `func (*Canvas) (c *Canvas) Close() error`
-   `func (*Canvas) (c *Canvas) Context() *gg.Context`
-   `func (*Canvas) (c *Canvas) Draw(fn func(*gg.Context)) error`
-   `func (*Canvas) (c *Canvas) Flush() (any, error)`
-   `func (*Canvas) (c *Canvas) Height() int`
-   `func (*Canvas) (c *Canvas) IsDirty() bool`
-   `func (*Canvas) (c *Canvas) MarkDirty()`
-   `func MustNew(provider gpucontext.DeviceProvider, width, height int) *Canvas`
-   `func New(provider gpucontext.DeviceProvider, width, height int) (*Canvas, error)`
-   `func (*Canvas) (c *Canvas) Provider() gpucontext.DeviceProvider`
-   `func (*Canvas) (c *Canvas) RenderDirect(surfaceView any, width, height uint32) error`
-   `func (*Canvas) (c *Canvas) Resize(width, height int) error`
-   `func (*Canvas) (c *Canvas) Size() (width, height int)`
-   `func (*Canvas) (c *Canvas) Texture() any`
-   `func (*Canvas) (c *Canvas) Width() int`

## integration/ggcanvas/render.go

-   `func DefaultRenderOptions() RenderOptions`
-   `func (*Canvas) (c *Canvas) RenderTo(dc gpucontext.TextureDrawer) error`
-   `func (*Canvas) (c *Canvas) RenderToEx(dc gpucontext.TextureDrawer, opts RenderOptions) error`
-   `func (*Canvas) (c *Canvas) RenderToPosition(dc gpucontext.TextureDrawer, x, y float32) error`
-   `func (*Canvas) (c *Canvas) RenderToScaled(dc gpucontext.TextureDrawer, scale float32) error`

## logger.go

-   `func (nopHandler) (nopHandler) Enabled(context.Context, slog.Level) bool`
-   `func (nopHandler) (nopHandler) Handle(context.Context, slog.Record) error`
-   `func Logger() *slog.Logger`
-   `func SetLogger(l *slog.Logger)`
-   `func (nopHandler) (nopHandler) WithAttrs([]slog.Attr) slog.Handler`
-   `func (nopHandler) (nopHandler) WithGroup(string) slog.Handler`

## mask.go

-   `func (*Mask) (m *Mask) At(x, y int) uint8`
-   `func (*Mask) (m *Mask) Bounds() image.Rectangle`
-   `func (*Mask) (m *Mask) Clear()`
-   `func (*Mask) (m *Mask) Clone() *Mask`
-   `func (*Mask) (m *Mask) Data() []uint8`
-   `func (*Mask) (m *Mask) Fill(value uint8)`
-   `func (*Mask) (m *Mask) Height() int`
-   `func (*Mask) (m *Mask) Invert()`
-   `func NewMask(width, height int) *Mask`
-   `func NewMaskFromAlpha(img image.Image) *Mask`
-   `func (*Mask) (m *Mask) Set(x, y int, value uint8)`
-   `func (*Mask) (m *Mask) Width() int`

## matrix.go

-   `func Identity() Matrix`
-   `func (Matrix) (m Matrix) Invert() Matrix`
-   `func (Matrix) (m Matrix) IsIdentity() bool`
-   `func (Matrix) (m Matrix) IsScaleOnly() bool`
-   `func (Matrix) (m Matrix) IsTranslation() bool`
-   `func (Matrix) (m Matrix) IsTranslationOnly() bool`
-   `func (Matrix) (m Matrix) MaxScaleFactor() float64`
-   `func (Matrix) (m Matrix) Multiply(other Matrix) Matrix`
-   `func Rotate(angle float64) Matrix`
-   `func Scale(x, y float64) Matrix`
-   `func (Matrix) (m Matrix) ScaleFactor() float64`
-   `func Shear(x, y float64) Matrix`
-   `func (Matrix) (m Matrix) TransformPoint(p Point) Point`
-   `func (Matrix) (m Matrix) TransformVector(p Point) Point`
-   `func Translate(x, y float64) Matrix`

## options.go

-   `func WithPipelineMode(mode PipelineMode) ContextOption`
-   `func WithPixmap(pm *Pixmap) ContextOption`
-   `func WithRenderer(r Renderer) ContextOption`

## paint.go

-   `func (*Paint) (p *Paint) Clone() *Paint`
-   `func (*Paint) (p *Paint) ColorAt(x, y float64) RGBA`
-   `func (*Paint) (p *Paint) EffectiveDash() *Dash`
-   `func (*Paint) (p *Paint) EffectiveLineCap() LineCap`
-   `func (*Paint) (p *Paint) EffectiveLineJoin() LineJoin`
-   `func (*Paint) (p *Paint) EffectiveLineWidth() float64`
-   `func (*Paint) (p *Paint) EffectiveMiterLimit() float64`
-   `func (*Paint) (p *Paint) GetBrush() Brush`
-   `func (*Paint) (p *Paint) GetStroke() Stroke`
-   `func (*Paint) (p *Paint) IsDashed() bool`
-   `func NewPaint() *Paint`
-   `func (*Paint) (p *Paint) SetBrush(b Brush)`
-   `func (*Paint) (p *Paint) SetStroke(s Stroke)`

## painter.go

-   `func (*SolidPainter) (p *SolidPainter) PaintSpan(dest []RGBA, _, _ int, length int)`
-   `func (*FuncPainter) (p *FuncPainter) PaintSpan(dest []RGBA, x, y, length int)`
-   `func PainterFromPaint(paint *Paint) Painter`

## path.go

-   `func (*Path) (p *Path) Arc(cx, cy, r, angle1, angle2 float64)`
-   `func (*Path) (p *Path) Circle(cx, cy, r float64)`
-   `func (*Path) (p *Path) Clear()`
-   `func (*Path) (p *Path) Clone() *Path`
-   `func (*Path) (p *Path) Close()`
-   `func (*Path) (p *Path) CubicTo(c1x, c1y, c2x, c2y, x, y float64)`
-   `func (*Path) (p *Path) CurrentPoint() Point`
-   `func (*Path) (p *Path) Elements() []PathElement`
-   `func (*Path) (p *Path) Ellipse(cx, cy, rx, ry float64)`
-   `func (*Path) (p *Path) HasCurrentPoint() bool`
-   `func (*Path) (p *Path) LineTo(x, y float64)`
-   `func (*Path) (p *Path) MoveTo(x, y float64)`
-   `func NewPath() *Path`
-   `func (*Path) (p *Path) QuadraticTo(cx, cy, x, y float64)`
-   `func (*Path) (p *Path) Rectangle(x, y, w, h float64)`
-   `func (*Path) (p *Path) RoundedRectangle(x, y, w, h, r float64)`
-   `func (*Path) (p *Path) Transform(m Matrix) *Path`

## path_builder.go

-   `func (*PathBuilder) (b *PathBuilder) Build() *Path`
-   `func BuildPath() *PathBuilder`
-   `func (*PathBuilder) (b *PathBuilder) Circle(cx, cy, r float64) *PathBuilder`
-   `func (*PathBuilder) (b *PathBuilder) Close() *PathBuilder`
-   `func (*PathBuilder) (b *PathBuilder) CubicTo(c1x, c1y, c2x, c2y, x, y float64) *PathBuilder`
-   `func (*PathBuilder) (b *PathBuilder) Ellipse(cx, cy, rx, ry float64) *PathBuilder`
-   `func (*PathBuilder) (b *PathBuilder) LineTo(x, y float64) *PathBuilder`
-   `func (*PathBuilder) (b *PathBuilder) MoveTo(x, y float64) *PathBuilder`
-   `func (*PathBuilder) (b *PathBuilder) Path() *Path`
-   `func (*PathBuilder) (b *PathBuilder) Polygon(cx, cy, radius float64, sides int) *PathBuilder`
-   `func (*PathBuilder) (b *PathBuilder) QuadTo(cx, cy, x, y float64) *PathBuilder`
-   `func (*PathBuilder) (b *PathBuilder) Rect(x, y, w, h float64) *PathBuilder`
-   `func (*PathBuilder) (b *PathBuilder) RoundRect(x, y, w, h, r float64) *PathBuilder`
-   `func (*PathBuilder) (b *PathBuilder) Star(cx, cy, outerRadius, innerRadius float64, points int) *PathBuilder`

## path_ops.go

-   `func (*Path) (p *Path) Area() float64`
-   `func (*Path) (p *Path) BoundingBox() Rect`
-   `func (*Path) (p *Path) Contains(pt Point) bool`
-   `func (*Path) (p *Path) Flatten(tolerance float64) []Point`
-   `func (*Path) (p *Path) FlattenCallback(tolerance float64, fn func(pt Point))`
-   `func (*Path) (p *Path) Length(accuracy float64) float64`
-   `func (*Path) (p *Path) Reversed() *Path`
-   `func (*Path) (p *Path) Winding(pt Point) int`

## pattern.go

-   `func (*SolidPattern) (p *SolidPattern) ColorAt(x, y float64) RGBA`
-   `func NewSolidPattern(color RGBA) *SolidPattern`

## pipeline_mode.go

-   `func SelectPipeline(stats SceneStats, hasComputeSupport bool) PipelineMode`
-   `func (PipelineMode) (m PipelineMode) String() string`

## pixmap.go

-   `func (*Pixmap) (p *Pixmap) At(x, y int) color.Color`
-   `func (*Pixmap) (p *Pixmap) Bounds() image.Rectangle`
-   `func (*Pixmap) (p *Pixmap) Clear(c RGBA)`
-   `func (*Pixmap) (p *Pixmap) ColorModel() color.Model`
-   `func (*Pixmap) (p *Pixmap) Data() []uint8`
-   `func (*Pixmap) (p *Pixmap) EncodeJPEG(w io.Writer, quality int) error`
-   `func (*Pixmap) (p *Pixmap) EncodePNG(w io.Writer) error`
-   `func (*Pixmap) (p *Pixmap) FillSpan(x1, x2, y int, c RGBA)`
-   `func (*Pixmap) (p *Pixmap) FillSpanBlend(x1, x2, y int, c RGBA)`
-   `func FromImage(img image.Image) *Pixmap`
-   `func (*Pixmap) (p *Pixmap) GetPixel(x, y int) RGBA`
-   `func (*Pixmap) (p *Pixmap) Height() int`
-   `func NewPixmap(width, height int) *Pixmap`
-   `func (*Pixmap) (p *Pixmap) SavePNG(path string) error`
-   `func (*Pixmap) (p *Pixmap) Set(x, y int, c color.Color)`
-   `func (*Pixmap) (p *Pixmap) SetPixel(x, y int, c RGBA)`
-   `func (*Pixmap) (p *Pixmap) SetPixelPremul(x, y int, r, g, b, a uint8)`
-   `func (*Pixmap) (p *Pixmap) ToImage() *image.RGBA`
-   `func (*Pixmap) (p *Pixmap) Width() int`

## point.go

-   `func (Point) (p Point) Add(q Point) Point`
-   `func (Point) (p Point) Cross(q Point) float64`
-   `func (Point) (p Point) Distance(q Point) float64`
-   `func (Point) (p Point) Div(s float64) Point`
-   `func (Point) (p Point) Dot(q Point) float64`
-   `func (Point) (p Point) Length() float64`
-   `func (Point) (p Point) LengthSquared() float64`
-   `func (Point) (p Point) Lerp(q Point, t float64) Point`
-   `func (Point) (p Point) Mul(s float64) Point`
-   `func (Point) (p Point) Normalize() Point`
-   `func Pt(x, y float64) Point`
-   `func (Point) (p Point) Rotate(angle float64) Point`
-   `func (Point) (p Point) Sub(q Point) Point`

## rasterizer_mode.go

-   `func (RasterizerMode) (m RasterizerMode) String() string`

## recording/backends/raster/backend.go

-   `func (*Backend) (b *Backend) Begin(width, height int) error`
-   `func (*Backend) (b *Backend) ClearClip()`
-   `func (*Backend) (b *Backend) DrawImage(img image.Image, src, dst recording.Rect, _ recording.ImageOptions)`
-   `func (*Backend) (b *Backend) DrawText(_ string, _, _ float64, _ text.Face, brush recording.Brush)`
-   `func (*Backend) (b *Backend) End() error`
-   `func (*Backend) (b *Backend) FillPath(path *gg.Path, brush recording.Brush, rule recording.FillRule)`
-   `func (*Backend) (b *Backend) FillRect(rect recording.Rect, brush recording.Brush)`
-   `func (*Backend) (b *Backend) Height() int`
-   `func (*Backend) (b *Backend) Image() image.Image`
-   `func NewBackend() *Backend`
-   `func (*Backend) (b *Backend) Pixmap() *gg.Pixmap`
-   `func (*Backend) (b *Backend) Restore()`
-   `func (*Backend) (b *Backend) Save()`
-   `func (*Backend) (b *Backend) SavePNG(path string) error`
-   `func (*Backend) (b *Backend) SaveToFile(path string) error`
-   `func (*Backend) (b *Backend) SetClip(path *gg.Path, rule recording.FillRule)`
-   `func (*Backend) (b *Backend) SetTransform(m recording.Matrix)`
-   `func (*Backend) (b *Backend) StrokePath(path *gg.Path, brush recording.Brush, stroke recording.Stroke)`
-   `func (*Backend) (b *Backend) Width() int`
-   `func (*countingWriter) (cw *countingWriter) Write(p []byte) (int, error)`
-   `func (*Backend) (b *Backend) WriteTo(w io.Writer) (int64, error)`

## recording/brush.go

-   `func (*RadialGradientBrush) (g *RadialGradientBrush) AddColorStop(offset float64, color gg.RGBA) *RadialGradientBrush`
-   `func (*LinearGradientBrush) (g *LinearGradientBrush) AddColorStop(offset float64, color gg.RGBA) *LinearGradientBrush`
-   `func (*SweepGradientBrush) (g *SweepGradientBrush) AddColorStop(offset float64, color gg.RGBA) *SweepGradientBrush`
-   `func BrushFromGG(b gg.Brush) Brush`
-   `func NewLinearGradientBrush(x0, y0, x1, y1 float64) *LinearGradientBrush`
-   `func NewPatternBrush(imageRef ImageRef) *PatternBrush`
-   `func NewRadialGradientBrush(cx, cy, startRadius, endRadius float64) *RadialGradientBrush`
-   `func NewSolidBrush(color gg.RGBA) SolidBrush`
-   `func NewSweepGradientBrush(cx, cy, startAngle float64) *SweepGradientBrush`
-   `func (*SweepGradientBrush) (g *SweepGradientBrush) SetEndAngle(endAngle float64) *SweepGradientBrush`
-   `func (*RadialGradientBrush) (g *RadialGradientBrush) SetExtend(mode ExtendMode) *RadialGradientBrush`
-   `func (*SweepGradientBrush) (g *SweepGradientBrush) SetExtend(mode ExtendMode) *SweepGradientBrush`
-   `func (*LinearGradientBrush) (g *LinearGradientBrush) SetExtend(mode ExtendMode) *LinearGradientBrush`
-   `func (*RadialGradientBrush) (g *RadialGradientBrush) SetFocus(fx, fy float64) *RadialGradientBrush`
-   `func (*PatternBrush) (b *PatternBrush) SetRepeat(mode RepeatMode) *PatternBrush`
-   `func (*PatternBrush) (b *PatternBrush) SetTransform(m gg.Matrix) *PatternBrush`

## recording/command.go

-   `func (Stroke) (s Stroke) Clone() Stroke`
-   `func DefaultImageOptions() ImageOptions`
-   `func DefaultStroke() Stroke`
-   `func (PathRef) (r PathRef) IsValid() bool`
-   `func (BrushRef) (r BrushRef) IsValid() bool`
-   `func (ImageRef) (r ImageRef) IsValid() bool`
-   `func (CommandType) (c CommandType) String() string`
-   `func (StrokeRectCommand) (StrokeRectCommand) Type() CommandType`
-   `func (SetStrokeStyleCommand) (SetStrokeStyleCommand) Type() CommandType`
-   `func (FillPathCommand) (FillPathCommand) Type() CommandType`
-   `func (StrokePathCommand) (StrokePathCommand) Type() CommandType`
-   `func (FillRectCommand) (FillRectCommand) Type() CommandType`
-   `func (SetClipCommand) (SetClipCommand) Type() CommandType`
-   `func (DrawImageCommand) (DrawImageCommand) Type() CommandType`
-   `func (DrawTextCommand) (DrawTextCommand) Type() CommandType`
-   `func (SetFillStyleCommand) (SetFillStyleCommand) Type() CommandType`
-   `func (ClearClipCommand) (ClearClipCommand) Type() CommandType`
-   `func (SetLineWidthCommand) (SetLineWidthCommand) Type() CommandType`
-   `func (SetLineCapCommand) (SetLineCapCommand) Type() CommandType`
-   `func (SetLineJoinCommand) (SetLineJoinCommand) Type() CommandType`
-   `func (SetMiterLimitCommand) (SetMiterLimitCommand) Type() CommandType`
-   `func (SetDashCommand) (SetDashCommand) Type() CommandType`
-   `func (SetFillRuleCommand) (SetFillRuleCommand) Type() CommandType`
-   `func (SetTransformCommand) (SetTransformCommand) Type() CommandType`
-   `func (RestoreCommand) (RestoreCommand) Type() CommandType`
-   `func (SaveCommand) (SaveCommand) Type() CommandType`

## recording/matrix.go

-   `func (Rect) (r Rect) Contains(x, y float64) bool`
-   `func (Matrix) (m Matrix) Determinant() float64`
-   `func (Rect) (r Rect) Height() float64`
-   `func Identity() Matrix`
-   `func (Rect) (r Rect) Inset(dx, dy float64) Rect`
-   `func (Rect) (r Rect) Intersect(other Rect) Rect`
-   `func (Matrix) (m Matrix) Invert() Matrix`
-   `func (Rect) (r Rect) IsEmpty() bool`
-   `func (Matrix) (m Matrix) IsIdentity() bool`
-   `func (Matrix) (m Matrix) IsTranslation() bool`
-   `func (Matrix) (m Matrix) Multiply(other Matrix) Matrix`
-   `func NewRect(x, y, width, height float64) Rect`
-   `func NewRectFromPoints(x1, y1, x2, y2 float64) Rect`
-   `func (Rect) (r Rect) Offset(dx, dy float64) Rect`
-   `func Rotate(angle float64) Matrix`
-   `func Scale(sx, sy float64) Matrix`
-   `func (Matrix) (m Matrix) ScaleFactor() float64`
-   `func Shear(x, y float64) Matrix`
-   `func (Matrix) (m Matrix) TransformPoint(x, y float64) (float64, float64)`
-   `func (Matrix) (m Matrix) TransformVector(x, y float64) (float64, float64)`
-   `func Translate(x, y float64) Matrix`
-   `func (Matrix) (m Matrix) Translation() (x, y float64)`
-   `func (Rect) (r Rect) Union(other Rect) Rect`
-   `func (Rect) (r Rect) Width() float64`
-   `func (Rect) (r Rect) X() float64`
-   `func (Rect) (r Rect) Y() float64`

## recording/pool.go

-   `func (*ResourcePool) (p *ResourcePool) AddBrush(brush Brush) BrushRef`
-   `func (*ResourcePool) (p *ResourcePool) AddFont(face text.Face) FontRef`
-   `func (*ResourcePool) (p *ResourcePool) AddImage(img image.Image) ImageRef`
-   `func (*ResourcePool) (p *ResourcePool) AddPath(path *gg.Path) PathRef`
-   `func (*ResourcePool) (p *ResourcePool) BrushCount() int`
-   `func (*ResourcePool) (p *ResourcePool) Clear()`
-   `func (*ResourcePool) (p *ResourcePool) Clone() *ResourcePool`
-   `func (*ResourcePool) (p *ResourcePool) FontCount() int`
-   `func (*ResourcePool) (p *ResourcePool) GetBrush(ref BrushRef) Brush`
-   `func (*ResourcePool) (p *ResourcePool) GetFont(ref FontRef) text.Face`
-   `func (*ResourcePool) (p *ResourcePool) GetImage(ref ImageRef) image.Image`
-   `func (*ResourcePool) (p *ResourcePool) GetPath(ref PathRef) *gg.Path`
-   `func (*ResourcePool) (p *ResourcePool) ImageCount() int`
-   `func (FontRef) (r FontRef) IsValid() bool`
-   `func NewResourcePool() *ResourcePool`
-   `func (*ResourcePool) (p *ResourcePool) PathCount() int`

## recording/recorder.go

-   `func (*Recorder) (r *Recorder) Clear()`
-   `func (*Recorder) (r *Recorder) ClearDash()`
-   `func (*Recorder) (r *Recorder) ClearPath()`
-   `func (*Recorder) (r *Recorder) ClearWithColor(c gg.RGBA)`
-   `func (*Recorder) (r *Recorder) Clip()`
-   `func (*Recorder) (r *Recorder) ClipPreserve()`
-   `func (*Recorder) (r *Recorder) ClosePath()`
-   `func (*Recording) (r *Recording) Commands() []Command`
-   `func (*Recorder) (r *Recorder) CubicTo(c1x, c1y, c2x, c2y, x, y float64)`
-   `func (*Recorder) (r *Recorder) DrawArc(x, y, radius, angle1, angle2 float64)`
-   `func (*Recorder) (r *Recorder) DrawCircle(x, y, radius float64)`
-   `func (*Recorder) (r *Recorder) DrawEllipse(x, y, rx, ry float64)`
-   `func (*Recorder) (r *Recorder) DrawEllipticalArc(x, y, rx, ry, angle1, angle2 float64)`
-   `func (*Recorder) (r *Recorder) DrawImage(img image.Image, x, y int)`
-   `func (*Recorder) (r *Recorder) DrawImageAnchored(img image.Image, x, y int, ax, ay float64)`
-   `func (*Recorder) (r *Recorder) DrawImageScaled(img image.Image, x, y, w, h float64)`
-   `func (*Recorder) (r *Recorder) DrawLine(x1, y1, x2, y2 float64)`
-   `func (*Recorder) (r *Recorder) DrawPoint(x, y, radius float64)`
-   `func (*Recorder) (r *Recorder) DrawRectangle(x, y, w, h float64)`
-   `func (*Recorder) (r *Recorder) DrawRoundedRectangle(x, y, w, h, radius float64)`
-   `func (*Recorder) (r *Recorder) DrawString(s string, x, y float64)`
-   `func (*Recorder) (r *Recorder) DrawStringAnchored(s string, x, y, ax, ay float64)`
-   `func (*Recorder) (r *Recorder) DrawStringWrapped(s string, x, y, ax, ay, width, lineSpacing float64, align text.Alignment)`
-   `func (*Recorder) (r *Recorder) Fill()`
-   `func (*Recorder) (r *Recorder) FillPreserve()`
-   `func (*Recorder) (r *Recorder) FillRectangle(x, y, w, h float64)`
-   `func (*Recorder) (r *Recorder) FillStroke()`
-   `func (*Recorder) (r *Recorder) FinishRecording() *Recording`
-   `func (*Recorder) (r *Recorder) GetCurrentPoint() (x, y float64, ok bool)`
-   `func (*Recorder) (r *Recorder) GetTransform() Matrix`
-   `func (*Recorder) (r *Recorder) Height() int`
-   `func (*Recording) (r *Recording) Height() int`
-   `func (*Recorder) (r *Recorder) Identity()`
-   `func (*Recorder) (r *Recorder) InvertY()`
-   `func (*Recorder) (r *Recorder) LineTo(x, y float64)`
-   `func (*Recorder) (r *Recorder) MeasureMultilineString(s string, lineSpacing float64) (width, height float64)`
-   `func (*Recorder) (r *Recorder) MeasureString(s string) (w, h float64)`
-   `func (*Recorder) (r *Recorder) MoveTo(x, y float64)`
-   `func NewRecorder(width, height int) *Recorder`
-   `func (*Recorder) (r *Recorder) NewSubPath()`
-   `func (*Recording) (r *Recording) Playback(backend Backend) error`
-   `func (*Recorder) (r *Recorder) Pop()`
-   `func (*Recorder) (r *Recorder) Push()`
-   `func (*Recorder) (r *Recorder) QuadraticTo(cx, cy, x, y float64)`
-   `func (*Recorder) (r *Recorder) ResetClip()`
-   `func (*Recording) (r *Recording) Resources() *ResourcePool`
-   `func (*Recorder) (r *Recorder) Restore()`
-   `func (*Recorder) (r *Recorder) Rotate(angle float64)`
-   `func (*Recorder) (r *Recorder) RotateAbout(angle, x, y float64)`
-   `func (*Recorder) (r *Recorder) Save()`
-   `func (*Recorder) (r *Recorder) Scale(sx, sy float64)`
-   `func (*Recorder) (r *Recorder) SetColor(c gg.RGBA)`
-   `func (*Recorder) (r *Recorder) SetDash(lengths ...float64)`
-   `func (*Recorder) (r *Recorder) SetDashOffset(offset float64)`
-   `func (*Recorder) (r *Recorder) SetFillBrush(brush gg.Brush)`
-   `func (*Recorder) (r *Recorder) SetFillRGB(red, green, blue float64)`
-   `func (*Recorder) (r *Recorder) SetFillRGBA(red, green, blue, alpha float64)`
-   `func (*Recorder) (r *Recorder) SetFillRule(rule FillRule)`
-   `func (*Recorder) (r *Recorder) SetFillRuleGG(rule gg.FillRule)`
-   `func (*Recorder) (r *Recorder) SetFillStyle(brush Brush)`
-   `func (*Recorder) (r *Recorder) SetFont(face text.Face)`
-   `func (*Recorder) (r *Recorder) SetFontFamily(family string)`
-   `func (*Recorder) (r *Recorder) SetFontSize(size float64)`
-   `func (*Recorder) (r *Recorder) SetHexColor(hex string)`
-   `func (*Recorder) (r *Recorder) SetLineCap(lc LineCap)`
-   `func (*Recorder) (r *Recorder) SetLineCapGG(lc gg.LineCap)`
-   `func (*Recorder) (r *Recorder) SetLineJoin(join LineJoin)`
-   `func (*Recorder) (r *Recorder) SetLineJoinGG(join gg.LineJoin)`
-   `func (*Recorder) (r *Recorder) SetLineWidth(width float64)`
-   `func (*Recorder) (r *Recorder) SetMiterLimit(limit float64)`
-   `func (*Recorder) (r *Recorder) SetRGB(red, green, blue float64)`
-   `func (*Recorder) (r *Recorder) SetRGBA(red, green, blue, alpha float64)`
-   `func (*Recorder) (r *Recorder) SetStrokeBrush(brush gg.Brush)`
-   `func (*Recorder) (r *Recorder) SetStrokeRGB(red, green, blue float64)`
-   `func (*Recorder) (r *Recorder) SetStrokeRGBA(red, green, blue, alpha float64)`
-   `func (*Recorder) (r *Recorder) SetStrokeStyle(brush Brush)`
-   `func (*Recorder) (r *Recorder) SetTransform(m Matrix)`
-   `func (*Recorder) (r *Recorder) Shear(x, y float64)`
-   `func (*Recorder) (r *Recorder) Stroke()`
-   `func (*Recorder) (r *Recorder) StrokePreserve()`
-   `func (*Recorder) (r *Recorder) StrokeRectangle(x, y, w, h float64)`
-   `func (*Recorder) (r *Recorder) Transform(m Matrix)`
-   `func (*Recorder) (r *Recorder) TransformPoint(x, y float64) (float64, float64)`
-   `func (*Recorder) (r *Recorder) Translate(x, y float64)`
-   `func (*Recorder) (r *Recorder) Width() int`
-   `func (*Recording) (r *Recording) Width() int`
-   `func (*Recorder) (r *Recorder) WordWrap(s string, w float64) []string`

## recording/registry.go

-   `func Backends() []string`
-   `func Count() int`
-   `func IsRegistered(name string) bool`
-   `func MustBackend(name string) Backend`
-   `func NewBackend(name string) (Backend, error)`
-   `func Register(name string, factory BackendFactory)`
-   `func Unregister(name string)`

## render/device.go

-   `func (NullDeviceHandle) (NullDeviceHandle) Adapter() gpucontext.Adapter`
-   `func DefaultTextureDescriptor(width, height uint32, format gputypes.TextureFormat) TextureDescriptor`
-   `func (NullDeviceHandle) (NullDeviceHandle) Device() gpucontext.Device`
-   `func (NullDeviceHandle) (NullDeviceHandle) Queue() gpucontext.Queue`
-   `func (NullDeviceHandle) (NullDeviceHandle) SurfaceFormat() gputypes.TextureFormat`

## render/gpu_renderer.go

-   `func (*GPURenderer) (r *GPURenderer) Capabilities() RendererCapabilities`
-   `func (*GPURenderer) (r *GPURenderer) CreateTextureTarget(width, height int) (*TextureTarget, error)`
-   `func (*GPURenderer) (r *GPURenderer) DeviceHandle() DeviceHandle`
-   `func (*GPURenderer) (r *GPURenderer) Flush() error`
-   `func NewGPURenderer(handle DeviceHandle) (*GPURenderer, error)`
-   `func (*GPURenderer) (r *GPURenderer) Render(target RenderTarget, scene *Scene) error`

## render/layers.go

-   `func (*LayeredPixmapTarget) (t *LayeredPixmapTarget) Clear(c color.Color)`
-   `func (*LayeredPixmapTarget) (t *LayeredPixmapTarget) ClearLayer(z int, c color.Color) error`
-   `func (*LayeredPixmapTarget) (t *LayeredPixmapTarget) Composite()`
-   `func (*LayeredPixmapTarget) (t *LayeredPixmapTarget) CreateLayer(z int) (RenderTarget, error)`
-   `func (*LayeredPixmapTarget) (t *LayeredPixmapTarget) Format() gputypes.TextureFormat`
-   `func (*LayeredPixmapTarget) (t *LayeredPixmapTarget) GetLayer(z int) RenderTarget`
-   `func (*LayeredPixmapTarget) (t *LayeredPixmapTarget) Height() int`
-   `func (*LayeredPixmapTarget) (t *LayeredPixmapTarget) Image() *image.RGBA`
-   `func (*LayeredPixmapTarget) (t *LayeredPixmapTarget) Layers() []int`
-   `func NewLayeredPixmapTarget(width, height int) *LayeredPixmapTarget`
-   `func (*LayeredPixmapTarget) (t *LayeredPixmapTarget) Pixels() []byte`
-   `func (*LayeredPixmapTarget) (t *LayeredPixmapTarget) RemoveLayer(z int) error`
-   `func (*LayeredPixmapTarget) (t *LayeredPixmapTarget) SetLayerVisible(z int, visible bool)`
-   `func (*LayeredPixmapTarget) (t *LayeredPixmapTarget) Stride() int`
-   `func (*LayeredPixmapTarget) (t *LayeredPixmapTarget) TextureView() TextureView`
-   `func (*LayeredPixmapTarget) (t *LayeredPixmapTarget) Width() int`

## render/scene.go

-   `func (*Scene) (s *Scene) Circle(cx, cy, r float64)`
-   `func (*Scene) (s *Scene) Clear(c color.Color)`
-   `func (*Scene) (s *Scene) ClearDirty()`
-   `func (*Scene) (s *Scene) ClosePath()`
-   `func (*Scene) (s *Scene) CommandCount() int`
-   `func (*Scene) (s *Scene) CubicTo(c1x, c1y, c2x, c2y, x, y float64)`
-   `func (*Scene) (s *Scene) DirtyRects() []DirtyRect`
-   `func (*Scene) (s *Scene) Fill()`
-   `func (*Scene) (s *Scene) HasDirtyRegions() bool`
-   `func (*Scene) (s *Scene) Invalidate(rect DirtyRect)`
-   `func (*Scene) (s *Scene) InvalidateAll()`
-   `func (*Scene) (s *Scene) IsEmpty() bool`
-   `func (*pathBuilder) (p *pathBuilder) IsEmpty() bool`
-   `func (*Scene) (s *Scene) LineTo(x, y float64)`
-   `func (*Scene) (s *Scene) MoveTo(x, y float64)`
-   `func (*Scene) (s *Scene) NeedsFullRedraw() bool`
-   `func NewScene() *Scene`
-   `func (*pathBuilder) (p *pathBuilder) Points() []float32`
-   `func (*Scene) (s *Scene) QuadTo(cx, cy, x, y float64)`
-   `func (*Scene) (s *Scene) Rectangle(x, y, width, height float64)`
-   `func (*Scene) (s *Scene) Reset()`
-   `func (*Scene) (s *Scene) SetFillColor(c color.Color)`
-   `func (*Scene) (s *Scene) SetFillRule(rule raster.FillRule)`
-   `func (*Scene) (s *Scene) SetStrokeColor(c color.Color)`
-   `func (*Scene) (s *Scene) SetStrokeWidth(width float64)`
-   `func (*Scene) (s *Scene) Stroke()`
-   `func (*pathBuilder) (p *pathBuilder) Verbs() []raster.PathVerb`

## render/software.go

-   `func (*SoftwareRenderer) (r *SoftwareRenderer) Capabilities() RendererCapabilities`
-   `func (*SoftwareRenderer) (r *SoftwareRenderer) Flush() error`
-   `func NewSoftwareRenderer() *SoftwareRenderer`
-   `func (*SoftwareRenderer) (r *SoftwareRenderer) Render(target RenderTarget, scene *Scene) error`

## render/target.go

-   `func (*PixmapTarget) (t *PixmapTarget) Clear(c color.Color)`
-   `func (*TextureTarget) (t *TextureTarget) Destroy()`
-   `func (*TextureTarget) (t *TextureTarget) Format() gputypes.TextureFormat`
-   `func (*PixmapTarget) (t *PixmapTarget) Format() gputypes.TextureFormat`
-   `func (*SurfaceTarget) (t *SurfaceTarget) Format() gputypes.TextureFormat`
-   `func (*PixmapTarget) (t *PixmapTarget) GetPixel(x, y int) color.Color`
-   `func (*TextureTarget) (t *TextureTarget) Height() int`
-   `func (*PixmapTarget) (t *PixmapTarget) Height() int`
-   `func (*SurfaceTarget) (t *SurfaceTarget) Height() int`
-   `func (*PixmapTarget) (t *PixmapTarget) Image() *image.RGBA`
-   `func NewPixmapTarget(width, height int) *PixmapTarget`
-   `func NewPixmapTargetFromImage(img *image.RGBA) *PixmapTarget`
-   `func NewSurfaceTarget(width, height int, format gputypes.TextureFormat, view TextureView) *SurfaceTarget`
-   `func NewTextureTarget(handle DeviceHandle, width, height int, format gputypes.TextureFormat) (*TextureTarget, error)`
-   `func (*PixmapTarget) (t *PixmapTarget) Pixels() []byte`
-   `func (*SurfaceTarget) (t *SurfaceTarget) Pixels() []byte`
-   `func (*TextureTarget) (t *TextureTarget) Pixels() []byte`
-   `func (*PixmapTarget) (t *PixmapTarget) Resize(width, height int)`
-   `func (*PixmapTarget) (t *PixmapTarget) SetPixel(x, y int, c color.Color)`
-   `func (*PixmapTarget) (t *PixmapTarget) Stride() int`
-   `func (*TextureTarget) (t *TextureTarget) Stride() int`
-   `func (*SurfaceTarget) (t *SurfaceTarget) Stride() int`
-   `func (*TextureTarget) (t *TextureTarget) TextureView() TextureView`
-   `func (*PixmapTarget) (t *PixmapTarget) TextureView() TextureView`
-   `func (*SurfaceTarget) (t *SurfaceTarget) TextureView() TextureView`
-   `func (*SurfaceTarget) (t *SurfaceTarget) Width() int`
-   `func (*TextureTarget) (t *TextureTarget) Width() int`
-   `func (*PixmapTarget) (t *PixmapTarget) Width() int`

## scene/blend_integration.go

-   `func AdvancedModes() []BlendMode`
-   `func AllBlendModes() []BlendMode`
-   `func BlendModeFromInternal(internal blend.BlendMode) BlendMode`
-   `func (BlendMode) (mode BlendMode) GetBlendFunc() blend.BlendFunc`
-   `func HSLModes() []BlendMode`
-   `func PorterDuffModes() []BlendMode`
-   `func (BlendMode) (mode BlendMode) ToInternalBlendMode() blend.BlendMode`

## scene/builder.go

-   `func (*SceneBuilder) (b *SceneBuilder) Build() *Scene`
-   `func (*SceneBuilder) (b *SceneBuilder) Clip(shape Shape, fn func(*SceneBuilder)) *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) CurrentTransform() Affine`
-   `func (*SceneBuilder) (b *SceneBuilder) DrawLine(x1, y1, x2, y2 float32, brush Brush, lineWidth float32) *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) Fill(shape Shape, brush Brush) *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) FillCircle(cx, cy, r float32, brush Brush) *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) FillPath(path *Path, brush Brush) *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) FillRect(x, y, width, height float32, brush Brush) *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) FillWith(shape Shape, brush Brush, style FillStyle) *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) Group(fn func(*SceneBuilder)) *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) Image(img *Image, rect Rect) *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) Layer(blend BlendMode, alpha float32, clip Shape, fn func(*SceneBuilder)) *SceneBuilder`
-   `func NewSceneBuilder() *SceneBuilder`
-   `func NewSceneBuilderFrom(scene *Scene) *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) Reset() *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) ResetTransform() *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) Rotate(angle float32) *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) Scale(x, y float32) *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) Scene() *Scene`
-   `func (*SceneBuilder) (b *SceneBuilder) Stroke(shape Shape, brush Brush, width float32) *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) StrokeCircle(cx, cy, r float32, brush Brush, lineWidth float32) *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) StrokePath(path *Path, brush Brush, lineWidth float32) *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) StrokeRect(x, y, width, height float32, brush Brush, lineWidth float32) *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) StrokeWith(shape Shape, brush Brush, style *StrokeStyle) *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) Transform(t Affine) *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) Translate(x, y float32) *SceneBuilder`
-   `func (*SceneBuilder) (b *SceneBuilder) WithTransform(t Affine, fn func(*SceneBuilder)) *SceneBuilder`

## scene/cache.go

-   `func (*LayerCache) (c *LayerCache) Contains(hash uint64) bool`
-   `func DefaultLayerCache() *LayerCache`
-   `func (*LayerCache) (c *LayerCache) EntryCount() int`
-   `func (*LayerCache) (c *LayerCache) Get(hash uint64) (*gg.Pixmap, bool)`
-   `func (*LayerCache) (c *LayerCache) GetVersion(hash uint64) (uint64, bool)`
-   `func (*LayerCache) (c *LayerCache) Invalidate(hash uint64)`
-   `func (*LayerCache) (c *LayerCache) InvalidateAll()`
-   `func (*LayerCache) (c *LayerCache) MaxSize() int64`
-   `func NewLayerCache(maxSizeMB int) *LayerCache`
-   `func (*LayerCache) (c *LayerCache) Put(hash uint64, pixmap *gg.Pixmap, version uint64)`
-   `func (*LayerCache) (c *LayerCache) ResetStats()`
-   `func (*LayerCache) (c *LayerCache) SetMaxSize(mb int)`
-   `func (*LayerCache) (c *LayerCache) Size() int64`
-   `func (*LayerCache) (c *LayerCache) Stats() CacheStats`
-   `func (*LayerCache) (c *LayerCache) Trim(targetSize int64)`

## scene/decoder.go

-   `func (*Decoder) (d *Decoder) Brush() (r, g, b, a float32)`
-   `func (*Decoder) (d *Decoder) CollectPath() *Path`
-   `func (*Decoder) (d *Decoder) CubicTo() (c1x, c1y, c2x, c2y, x, y float32)`
-   `func (*Decoder) (d *Decoder) Encoding() *Encoding`
-   `func (*Decoder) (d *Decoder) Fill() (brush Brush, style FillStyle)`
-   `func (*Decoder) (d *Decoder) HasMore() bool`
-   `func (*Decoder) (d *Decoder) Image() (imageIndex uint32, transform Affine)`
-   `func (*Decoder) (d *Decoder) LineTo() (x, y float32)`
-   `func (*Decoder) (d *Decoder) MoveTo() (x, y float32)`
-   `func NewDecoder(enc *Encoding) *Decoder`
-   `func (*Decoder) (d *Decoder) Next() bool`
-   `func (*Decoder) (d *Decoder) Peek() Tag`
-   `func (*Decoder) (d *Decoder) Position() int`
-   `func (*Decoder) (d *Decoder) PushLayer() (blend BlendMode, alpha float32)`
-   `func (*Decoder) (d *Decoder) QuadTo() (cx, cy, x, y float32)`
-   `func (*Decoder) (d *Decoder) Reset(enc *Encoding)`
-   `func (*Decoder) (d *Decoder) SkipPath()`
-   `func (*Decoder) (d *Decoder) Stroke() (brush Brush, style *StrokeStyle)`
-   `func (*Decoder) (d *Decoder) Tag() Tag`
-   `func (*Decoder) (d *Decoder) Transform() Affine`

## scene/encoding.go

-   `func AffineFromMatrix(m gg.Matrix) Affine`
-   `func (*Encoding) (e *Encoding) Append(other *Encoding)`
-   `func (*Encoding) (e *Encoding) Bounds() Rect`
-   `func (*Encoding) (e *Encoding) Brushes() []Brush`
-   `func (*Encoding) (e *Encoding) Capacity() int`
-   `func (*Encoding) (e *Encoding) Clone() *Encoding`
-   `func DefaultStrokeStyle() *StrokeStyle`
-   `func (*Encoding) (e *Encoding) DrawData() []uint32`
-   `func EmptyRect() Rect`
-   `func (*Encoding) (e *Encoding) EncodeBeginClip()`
-   `func (*Encoding) (e *Encoding) EncodeBrush(brush Brush) int`
-   `func (*Encoding) (e *Encoding) EncodeEndClip()`
-   `func (*Encoding) (e *Encoding) EncodeFill(brush Brush, style FillStyle)`
-   `func (*Encoding) (e *Encoding) EncodeImage(imageIndex uint32, transform Affine)`
-   `func (*Encoding) (e *Encoding) EncodePath(p *gg.Path)`
-   `func (*Encoding) (e *Encoding) EncodePopLayer()`
-   `func (*Encoding) (e *Encoding) EncodePushLayer(blend BlendMode, alpha float32)`
-   `func (*Encoding) (e *Encoding) EncodeStroke(brush Brush, style *StrokeStyle)`
-   `func (*Encoding) (e *Encoding) EncodeTransform(t Affine)`
-   `func (*Encoding) (e *Encoding) EncodeTransformFromMatrix(m gg.Matrix)`
-   `func (*Iterator) (it *Iterator) GetBrush(idx uint32) (Brush, bool)`
-   `func (*Encoding) (e *Encoding) Hash() uint64`
-   `func (Rect) (r Rect) Height() float32`
-   `func IdentityAffine() Affine`
-   `func (BlendMode) (mode BlendMode) IsAdvanced() bool`
-   `func (Rect) (r Rect) IsEmpty() bool`
-   `func (*Encoding) (e *Encoding) IsEmpty() bool`
-   `func (BlendMode) (mode BlendMode) IsHSL() bool`
-   `func (Affine) (a Affine) IsIdentity() bool`
-   `func (BlendMode) (mode BlendMode) IsPorterDuff() bool`
-   `func (Affine) (a Affine) Multiply(b Affine) Affine`
-   `func NewEncoding() *Encoding`
-   `func (*Encoding) (e *Encoding) NewIterator() *Iterator`
-   `func (*Iterator) (it *Iterator) Next() (Tag, bool)`
-   `func (*Encoding) (e *Encoding) PathCount() int`
-   `func (*Encoding) (e *Encoding) PathData() []float32`
-   `func (*Iterator) (it *Iterator) ReadDrawData(n int) []uint32`
-   `func (*Iterator) (it *Iterator) ReadPathData(n int) []float32`
-   `func (*Iterator) (it *Iterator) ReadTransform() (Affine, bool)`
-   `func (*Encoding) (e *Encoding) Reset()`
-   `func (*Iterator) (it *Iterator) Reset()`
-   `func RotateAffine(angle float32) Affine`
-   `func ScaleAffine(x, y float32) Affine`
-   `func (*Encoding) (e *Encoding) ShapeCount() int`
-   `func (*Encoding) (e *Encoding) Size() int`
-   `func SolidBrush(c gg.RGBA) Brush`
-   `func (BlendMode) (mode BlendMode) String() string`
-   `func (*Encoding) (e *Encoding) Tags() []Tag`
-   `func (Affine) (a Affine) TransformPoint(x, y float32) (float32, float32)`
-   `func (*Encoding) (e *Encoding) Transforms() []Affine`
-   `func TranslateAffine(x, y float32) Affine`
-   `func (Rect) (r Rect) Union(other Rect) Rect`
-   `func (Rect) (r Rect) UnionPoint(x, y float32) Rect`
-   `func (*Encoding) (e *Encoding) UpdateBounds(bounds Rect)`
-   `func (Rect) (r Rect) Width() float32`

## scene/filter.go

-   `func (*FilterChain) (fc *FilterChain) Add(f Filter)`
-   `func (*FilterChain) (fc *FilterChain) Apply(src, dst *gg.Pixmap, bounds Rect)`
-   `func (*FilterChain) (fc *FilterChain) ExpandBounds(input Rect) Rect`
-   `func (FilterType) (ft FilterType) ExpandsOutput() bool`
-   `func (*FilterChain) (fc *FilterChain) IsEmpty() bool`
-   `func (*FilterChain) (fc *FilterChain) Len() int`
-   `func NewFilterChain(filters ...Filter) *FilterChain`
-   `func (FilterType) (ft FilterType) String() string`

## scene/layer.go

-   `func (*LayerStack) (s *LayerStack) AcquireLayer() *LayerState`
-   `func (*LayerStack) (s *LayerStack) All() []*LayerState`
-   `func (*ClipStack) (s *ClipStack) CombinedBounds() Rect`
-   `func (*ClipState) (cs *ClipState) Contains(x, y float32) bool`
-   `func (*ClipStack) (s *ClipStack) Contains(x, y float32) bool`
-   `func (*ClipStack) (s *ClipStack) Depth() int`
-   `func (*LayerStack) (s *LayerStack) Depth() int`
-   `func (*LayerState) (ls *LayerState) HasClip() bool`
-   `func (*ClipState) (cs *ClipState) Intersects(r Rect) bool`
-   `func (*ClipStack) (s *ClipStack) Intersects(r Rect) bool`
-   `func (LayerKind) (k LayerKind) IsClipOnly() bool`
-   `func (*LayerState) (ls *LayerState) IsEmpty() bool`
-   `func (*ClipStack) (s *ClipStack) IsEmpty() bool`
-   `func (*LayerStack) (s *LayerStack) IsRoot() bool`
-   `func (LayerKind) (k LayerKind) NeedsOffscreen() bool`
-   `func NewClipLayer(clip Shape) *LayerState`
-   `func NewClipStack() *ClipStack`
-   `func NewClipState(shape Shape, transform Affine) *ClipState`
-   `func NewFilteredLayer(blend BlendMode, alpha float32) *LayerState`
-   `func NewLayerStack() *LayerStack`
-   `func NewLayerState(kind LayerKind, blend BlendMode, alpha float32) *LayerState`
-   `func (*ClipStack) (s *ClipStack) Pop() *ClipState`
-   `func (*LayerStack) (s *LayerStack) Pop() *LayerState`
-   `func (*ClipStack) (s *ClipStack) Push(clip *ClipState)`
-   `func (*LayerStack) (s *LayerStack) Push(layer *LayerState)`
-   `func (*LayerStack) (s *LayerStack) ReleaseLayer(layer *LayerState)`
-   `func (*LayerStack) (s *LayerStack) Reset()`
-   `func (*LayerState) (ls *LayerState) Reset()`
-   `func (*ClipStack) (s *ClipStack) Reset()`
-   `func (*LayerStack) (s *LayerStack) Root() *LayerState`
-   `func (LayerKind) (k LayerKind) String() string`
-   `func (*ClipStack) (s *ClipStack) Top() *ClipState`
-   `func (*LayerStack) (s *LayerStack) Top() *LayerState`
-   `func (*LayerState) (ls *LayerState) UpdateBounds(r Rect)`

## scene/path.go

-   `func (*Path) (p *Path) Arc(cx, cy, rx, ry, startAngle, endAngle float32, sweepClockwise bool) *Path`
-   `func (*Path) (p *Path) Bounds() Rect`
-   `func (*Path) (p *Path) Circle(cx, cy, r float32) *Path`
-   `func (*Path) (p *Path) Clone() *Path`
-   `func (*Path) (p *Path) Close() *Path`
-   `func (*Path) (p *Path) Contains(px, py float32) bool`
-   `func (*Path) (p *Path) CubicTo(c1x, c1y, c2x, c2y, x, y float32) *Path`
-   `func (*Path) (p *Path) Elements() iter.Seq[PathElement]`
-   `func (*Path) (p *Path) ElementsWithCursor() iter.Seq2[Point, PathElement]`
-   `func (*Path) (p *Path) Ellipse(cx, cy, rx, ry float32) *Path`
-   `func (*PathPool) (pp *PathPool) Get() *Path`
-   `func (*Path) (p *Path) IsEmpty() bool`
-   `func (*Path) (p *Path) LineTo(x, y float32) *Path`
-   `func (*Path) (p *Path) MoveTo(x, y float32) *Path`
-   `func NewPath() *Path`
-   `func NewPathPool() *PathPool`
-   `func (*Path) (p *Path) PointCount() int`
-   `func (PathVerb) (v PathVerb) PointCount() int`
-   `func (*Path) (p *Path) Points() []float32`
-   `func (*PathPool) (pp *PathPool) Put(p *Path)`
-   `func (*Path) (p *Path) QuadTo(cx, cy, x, y float32) *Path`
-   `func (*Path) (p *Path) Rectangle(x, y, w, h float32) *Path`
-   `func (*Path) (p *Path) Reset()`
-   `func (*Path) (p *Path) Reverse() *Path`
-   `func (*Path) (p *Path) RoundedRectangle(x, y, w, h, r float32) *Path`
-   `func (PathVerb) (v PathVerb) String() string`
-   `func (*Path) (p *Path) Transform(t Affine) *Path`
-   `func (*Path) (p *Path) VerbCount() int`
-   `func (*Path) (p *Path) Verbs() []PathVerb`

## scene/pool.go

-   `func (*EncodingPool) (p *EncodingPool) Get() *Encoding`
-   `func GetEncoding() *Encoding`
-   `func NewEncodingPool() *EncodingPool`
-   `func (*EncodingPool) (p *EncodingPool) Put(enc *Encoding)`
-   `func PutEncoding(enc *Encoding)`
-   `func (*EncodingPool) (p *EncodingPool) Warmup(count int)`

## scene/renderer.go

-   `func (*Renderer) (r *Renderer) Cache() *LayerCache`
-   `func (*Renderer) (r *Renderer) CacheStats() CacheStats`
-   `func (*Renderer) (r *Renderer) Close()`
-   `func (*Renderer) (r *Renderer) DirtyTileCount() int`
-   `func (*Renderer) (r *Renderer) Height() int`
-   `func (*Renderer) (r *Renderer) MarkAllDirty()`
-   `func (*Renderer) (r *Renderer) MarkDirty(x, y, w, h int)`
-   `func NewRenderer(width, height int, opts ...RendererOption) *Renderer`
-   `func (*Renderer) (r *Renderer) Render(target *gg.Pixmap, scene *Scene) error`
-   `func (*Renderer) (r *Renderer) RenderDirty(target *gg.Pixmap, scene *Scene, dirty *parallel.DirtyRegion) error`
-   `func (*Renderer) (r *Renderer) RenderDirtyWithContext(ctx context.Context, target *gg.Pixmap, scene *Scene, dirty *parallel.DirtyRegion) error`
-   `func (*Renderer) (r *Renderer) RenderWithContext(ctx context.Context, target *gg.Pixmap, scene *Scene) error`
-   `func (*Renderer) (r *Renderer) Resize(width, height int)`
-   `func (*Renderer) (r *Renderer) Stats() RenderStats`
-   `func (*Renderer) (r *Renderer) TileCount() int`
-   `func (*Renderer) (r *Renderer) Width() int`
-   `func WithCache(cache *LayerCache) RendererOption`
-   `func WithCacheSize(mb int) RendererOption`
-   `func WithTileSize(size int) RendererOption`
-   `func WithWorkers(n int) RendererOption`
-   `func (*Renderer) (r *Renderer) Workers() int`

## scene/scene.go

-   `func (*Scene) (s *Scene) Bounds() Rect`
-   `func (*Image) (img *Image) Bounds() Rect`
-   `func (*Scene) (s *Scene) ClipBounds() Rect`
-   `func (*Scene) (s *Scene) ClipDepth() int`
-   `func (*Scene) (s *Scene) DrawImage(img *Image, transform Affine)`
-   `func (*Scene) (s *Scene) Encoding() *Encoding`
-   `func (*Scene) (s *Scene) Fill(style FillStyle, transform Affine, brush Brush, shape Shape)`
-   `func (*Scene) (s *Scene) Flatten() *Encoding`
-   `func (*ScenePool) (sp *ScenePool) Get() *Scene`
-   `func (*Scene) (s *Scene) Images() []*Image`
-   `func (*Scene) (s *Scene) IsEmpty() bool`
-   `func (*Image) (img *Image) IsEmpty() bool`
-   `func (*Scene) (s *Scene) LayerDepth() int`
-   `func NewImage(width, height int) *Image`
-   `func NewScene() *Scene`
-   `func NewScenePool() *ScenePool`
-   `func (*Scene) (s *Scene) PopClip() bool`
-   `func (*Scene) (s *Scene) PopLayer() bool`
-   `func (*Scene) (s *Scene) PopTransform() bool`
-   `func (*Scene) (s *Scene) PushClip(shape Shape)`
-   `func (*Scene) (s *Scene) PushLayer(blend BlendMode, alpha float32, clip Shape)`
-   `func (*Scene) (s *Scene) PushTransform(t Affine)`
-   `func (*ScenePool) (sp *ScenePool) Put(scene *Scene)`
-   `func (*Scene) (s *Scene) Reset()`
-   `func (*Scene) (s *Scene) Rotate(angle float32)`
-   `func (*Scene) (s *Scene) Scale(x, y float32)`
-   `func (*Scene) (s *Scene) SetTransform(t Affine)`
-   `func (*Scene) (s *Scene) Stroke(style *StrokeStyle, transform Affine, brush Brush, shape Shape)`
-   `func (*Scene) (s *Scene) Transform() Affine`
-   `func (*Scene) (s *Scene) TransformDepth() int`
-   `func (*Scene) (s *Scene) Translate(x, y float32)`
-   `func (*Scene) (s *Scene) Version() uint64`
-   `func (*ScenePool) (sp *ScenePool) Warmup(count int)`

## scene/shape.go

-   `func (*CompositeShape) (cs *CompositeShape) AddShape(shape Shape)`
-   `func (*TransformShape) (ts *TransformShape) Bounds() Rect`
-   `func (*PathShape) (ps *PathShape) Bounds() Rect`
-   `func (*LineShape) (l *LineShape) Bounds() Rect`
-   `func (*CompositeShape) (cs *CompositeShape) Bounds() Rect`
-   `func (*PolygonShape) (p *PolygonShape) Bounds() Rect`
-   `func (*RoundedRectShape) (r *RoundedRectShape) Bounds() Rect`
-   `func (*RegularPolygonShape) (rp *RegularPolygonShape) Bounds() Rect`
-   `func (*StarShape) (s *StarShape) Bounds() Rect`
-   `func (*RectShape) (r *RectShape) Bounds() Rect`
-   `func (*PieShape) (p *PieShape) Bounds() Rect`
-   `func (*CircleShape) (c *CircleShape) Bounds() Rect`
-   `func (*ArcShape) (a *ArcShape) Bounds() Rect`
-   `func (*EllipseShape) (e *EllipseShape) Bounds() Rect`
-   `func (*CircleShape) (c *CircleShape) Contains(px, py float32) bool`
-   `func (*EllipseShape) (e *EllipseShape) Contains(px, py float32) bool`
-   `func (*RectShape) (r *RectShape) Contains(px, py float32) bool`
-   `func (*LineShape) (l *LineShape) Length() float32`
-   `func NewArcShape(cx, cy, rx, ry, startAngle, endAngle float32, sweepClockwise bool) *ArcShape`
-   `func NewCircleShape(cx, cy, r float32) *CircleShape`
-   `func NewCompositeShape(shapes ...Shape) *CompositeShape`
-   `func NewEllipseShape(cx, cy, rx, ry float32) *EllipseShape`
-   `func NewLineShape(x1, y1, x2, y2 float32) *LineShape`
-   `func NewPathShape(path *Path) *PathShape`
-   `func NewPieShape(cx, cy, r, startAngle, endAngle float32, sweepClockwise bool) *PieShape`
-   `func NewPolygonShape(points ...float32) *PolygonShape`
-   `func NewRectShape(x, y, width, height float32) *RectShape`
-   `func NewRegularPolygonShape(cx, cy, r float32, sides int, rotation float32) *RegularPolygonShape`
-   `func NewRoundedRectShape(x, y, width, height, radius float32) *RoundedRectShape`
-   `func NewStarShape(cx, cy, outerRadius, innerRadius float32, points int, rotation float32) *StarShape`
-   `func NewTransformShape(shape Shape, transform Affine) *TransformShape`
-   `func (*PolygonShape) (p *PolygonShape) Point(i int) (x, y float32, ok bool)`
-   `func (*PolygonShape) (p *PolygonShape) PointCount() int`
-   `func (*CompositeShape) (cs *CompositeShape) ShapeCount() int`
-   `func (*ArcShape) (a *ArcShape) ToPath() *Path`
-   `func (*EllipseShape) (e *EllipseShape) ToPath() *Path`
-   `func (*LineShape) (l *LineShape) ToPath() *Path`
-   `func (*PieShape) (p *PieShape) ToPath() *Path`
-   `func (*StarShape) (s *StarShape) ToPath() *Path`
-   `func (*RegularPolygonShape) (rp *RegularPolygonShape) ToPath() *Path`
-   `func (*TransformShape) (ts *TransformShape) ToPath() *Path`
-   `func (*CircleShape) (c *CircleShape) ToPath() *Path`
-   `func (*RoundedRectShape) (r *RoundedRectShape) ToPath() *Path`
-   `func (*PathShape) (ps *PathShape) ToPath() *Path`
-   `func (*CompositeShape) (cs *CompositeShape) ToPath() *Path`
-   `func (*RectShape) (r *RectShape) ToPath() *Path`
-   `func (*PolygonShape) (p *PolygonShape) ToPath() *Path`

## scene/tag.go

-   `func (Tag) (t Tag) DataSize() int`
-   `func (Tag) (t Tag) IsClipCommand() bool`
-   `func (Tag) (t Tag) IsDrawCommand() bool`
-   `func (Tag) (t Tag) IsLayerCommand() bool`
-   `func (Tag) (t Tag) IsPathCommand() bool`
-   `func (Tag) (t Tag) String() string`

## scene/text.go

-   `func (*TextShape) (ts *TextShape) Bounds() Rect`
-   `func (*TextRenderer) (r *TextRenderer) Config() TextRendererConfig`
-   `func DefaultTextRendererConfig() TextRendererConfig`
-   `func (*Scene) (s *Scene) DrawGlyphs(glyphs []text.ShapedGlyph, face text.Face, brush Brush) error`
-   `func (*Scene) (s *Scene) DrawText(str string, face text.Face, x, y float32, brush Brush) error`
-   `func (*TextRendererPool) (p *TextRendererPool) Get() *TextRenderer`
-   `func NewTextRenderer() *TextRenderer`
-   `func NewTextRendererPool() *TextRendererPool`
-   `func NewTextRendererWithConfig(config TextRendererConfig) *TextRenderer`
-   `func NewTextShape(str string, face text.Face, x, y float32) (*TextShape, error)`
-   `func (*TextRendererPool) (p *TextRendererPool) Put(r *TextRenderer)`
-   `func (*TextRenderer) (r *TextRenderer) RenderGlyph(glyph text.ShapedGlyph, face text.Face) (*RenderedGlyph, error)`
-   `func (*TextRenderer) (r *TextRenderer) RenderGlyphs(glyphs []text.ShapedGlyph, face text.Face) ([]*RenderedGlyph, error)`
-   `func (*TextRenderer) (r *TextRenderer) RenderRun(run *text.ShapedRun) ([]*RenderedGlyph, error)`
-   `func (*TextRenderer) (r *TextRenderer) RenderText(str string, face text.Face) ([]*RenderedGlyph, error)`
-   `func (*TextRenderer) (r *TextRenderer) RenderTextToScene(s *Scene, str string, face text.Face, x, y float32, brush Brush) error`
-   `func (*TextRenderer) (r *TextRenderer) RenderToScene(s *Scene, glyphs []text.ShapedGlyph, face text.Face, brush Brush) error`
-   `func (*TextRenderer) (r *TextRenderer) SetConfig(config TextRendererConfig)`
-   `func TextAdvance(glyphs []*RenderedGlyph) float32`
-   `func TextBounds(glyphs []*RenderedGlyph) Rect`
-   `func (*TextRenderer) (r *TextRenderer) ToCompositePath(glyphs []*RenderedGlyph) *Path`
-   `func (*TextShape) (ts *TextShape) ToPath() *Path`

## sdf.go

-   `func SDFCircleCoverage(px, py, cx, cy, radius, halfStrokeWidth float64) float64`
-   `func SDFFilledCircleCoverage(px, py, cx, cy, radius float64) float64`
-   `func SDFFilledRRectCoverage(px, py, cx, cy, halfW, halfH, cornerRadius float64) float64`
-   `func SDFRRectCoverage(px, py, cx, cy, halfW, halfH, cornerRadius, halfStrokeWidth float64) float64`

## sdf_accelerator.go

-   `func (*SDFAccelerator) (a *SDFAccelerator) CanAccelerate(op AcceleratedOp) bool`
-   `func (*SDFAccelerator) (a *SDFAccelerator) Close()`
-   `func (*SDFAccelerator) (a *SDFAccelerator) FillPath(_ GPURenderTarget, _ *Path, _ *Paint) error`
-   `func (*SDFAccelerator) (a *SDFAccelerator) FillShape(target GPURenderTarget, shape DetectedShape, paint *Paint) error`
-   `func (*SDFAccelerator) (a *SDFAccelerator) Flush(_ GPURenderTarget) error`
-   `func (*SDFAccelerator) (a *SDFAccelerator) Init() error`
-   `func (*SDFAccelerator) (a *SDFAccelerator) Name() string`
-   `func (*SDFAccelerator) (a *SDFAccelerator) SetForceSDF(force bool)`
-   `func (*SDFAccelerator) (a *SDFAccelerator) StrokePath(_ GPURenderTarget, _ *Path, _ *Paint) error`
-   `func (*SDFAccelerator) (a *SDFAccelerator) StrokeShape(target GPURenderTarget, shape DetectedShape, paint *Paint) error`

## shape_detect.go

-   `func DetectShape(path *Path) DetectedShape`

## shapes.go

-   `func (*Context) (c *Context) DrawRegularPolygon(n int, x, y, r, rotation float64)`

## software.go

-   `func (*SoftwareRenderer) (r *SoftwareRenderer) Fill(pixmap *Pixmap, p *Path, paint *Paint) error`
-   `func NewSoftwareRenderer(width, height int) *SoftwareRenderer`
-   `func (*SoftwareRenderer) (r *SoftwareRenderer) Resize(width, height int)`
-   `func (*SoftwareRenderer) (r *SoftwareRenderer) Stroke(pixmap *Pixmap, p *Path, paint *Paint) error`

## solver.go

-   `func SolveCubic(a, b, c, d float64) []float64`
-   `func SolveCubicInUnitInterval(a, b, c, d float64) []float64`
-   `func SolveQuadratic(a, b, c float64) []float64`
-   `func SolveQuadraticInUnitInterval(a, b, c float64) []float64`

## stroke.go

-   `func Bold() Stroke`
-   `func (Stroke) (s Stroke) Clone() Stroke`
-   `func DashedStroke(lengths ...float64) Stroke`
-   `func DefaultStroke() Stroke`
-   `func DottedStroke() Stroke`
-   `func (Stroke) (s Stroke) IsDashed() bool`
-   `func RoundStroke() Stroke`
-   `func SquareStroke() Stroke`
-   `func Thick() Stroke`
-   `func Thin() Stroke`
-   `func (Stroke) (s Stroke) WithCap(lineCap LineCap) Stroke`
-   `func (Stroke) (s Stroke) WithDash(dash *Dash) Stroke`
-   `func (Stroke) (s Stroke) WithDashOffset(offset float64) Stroke`
-   `func (Stroke) (s Stroke) WithDashPattern(lengths ...float64) Stroke`
-   `func (Stroke) (s Stroke) WithJoin(join LineJoin) Stroke`
-   `func (Stroke) (s Stroke) WithMiterLimit(limit float64) Stroke`
-   `func (Stroke) (s Stroke) WithWidth(w float64) Stroke`

## surface/gpu_surface.go

-   `func (*GPUSurface) (s *GPUSurface) Backend() GPUBackend`
-   `func (*GPUSurface) (s *GPUSurface) Capabilities() Capabilities`
-   `func (*GPUSurface) (s *GPUSurface) Clear(c color.Color)`
-   `func (*GPUSurface) (s *GPUSurface) Close() error`
-   `func (*GPUSurface) (s *GPUSurface) DrawImage(img image.Image, at Point, opts *DrawImageOptions)`
-   `func (*GPUSurface) (s *GPUSurface) Fill(path *Path, style FillStyle)`
-   `func (*GPUSurface) (s *GPUSurface) Flush() error`
-   `func (*GPUSurface) (s *GPUSurface) Height() int`
-   `func NewGPUSurface(width, height int, backend GPUBackend) (*GPUSurface, error)`
-   `func (*GPUSurface) (s *GPUSurface) Snapshot() *image.RGBA`
-   `func (*GPUSurface) (s *GPUSurface) Stroke(path *Path, style StrokeStyle)`
-   `func (*GPUSurface) (s *GPUSurface) Width() int`

## surface/image_surface.go

-   `func (*ImageSurface) (s *ImageSurface) Capabilities() Capabilities`
-   `func (*ImageSurface) (s *ImageSurface) Clear(c color.Color)`
-   `func (*ImageSurface) (s *ImageSurface) Close() error`
-   `func (*ImageSurface) (s *ImageSurface) DrawImage(img image.Image, at Point, opts *DrawImageOptions)`
-   `func (*ImageSurface) (s *ImageSurface) Fill(path *Path, style FillStyle)`
-   `func (*ImageSurface) (s *ImageSurface) Flush() error`
-   `func (*ImageSurface) (s *ImageSurface) Height() int`
-   `func (*ImageSurface) (s *ImageSurface) Image() *image.RGBA`
-   `func NewImageSurface(width, height int) *ImageSurface`
-   `func NewImageSurfaceFromImage(img *image.RGBA) *ImageSurface`
-   `func (*ImageSurface) (s *ImageSurface) Snapshot() *image.RGBA`
-   `func (*ImageSurface) (s *ImageSurface) Stroke(path *Path, style StrokeStyle)`
-   `func (*ImageSurface) (s *ImageSurface) Width() int`

## surface/path.go

-   `func (*Path) (p *Path) Arc(cx, cy, r, angle1, angle2 float64)`
-   `func (*Path) (p *Path) Bounds() (minX, minY, maxX, maxY float64)`
-   `func (*Path) (p *Path) Circle(cx, cy, r float64)`
-   `func (*Path) (p *Path) Clear()`
-   `func (*Path) (p *Path) Clone() *Path`
-   `func (*Path) (p *Path) Close()`
-   `func (*Path) (p *Path) CubicTo(c1x, c1y, c2x, c2y, x, y float64)`
-   `func (*Path) (p *Path) CurrentPoint() Point`
-   `func (*Path) (p *Path) Ellipse(cx, cy, rx, ry float64)`
-   `func (*Path) (p *Path) IsEmpty() bool`
-   `func (*Path) (p *Path) LineTo(x, y float64)`
-   `func (*Path) (p *Path) MoveTo(x, y float64)`
-   `func NewPath() *Path`
-   `func (*Path) (p *Path) Points() []float32`
-   `func (*Path) (p *Path) QuadTo(cx, cy, x, y float64)`
-   `func (*Path) (p *Path) Rectangle(x, y, w, h float64)`
-   `func (*Path) (p *Path) RoundedRectangle(x, y, w, h, r float64)`
-   `func (*Path) (p *Path) Verbs() []raster.PathVerb`

## surface/registry.go

-   `func (*Registry) (r *Registry) Available() []string`
-   `func Available() []string`
-   `func (*BackendUnavailableError) (e *BackendUnavailableError) Error() string`
-   `func (*BackendNotFoundError) (e *BackendNotFoundError) Error() string`
-   `func (*Registry) (r *Registry) Get(name string) (*RegistryEntry, bool)`
-   `func Get(name string) (*RegistryEntry, bool)`
-   `func (*Registry) (r *Registry) List() []string`
-   `func List() []string`
-   `func NewRegistry() *Registry`
-   `func NewSurface(width, height int) (Surface, error)`
-   `func (*Registry) (r *Registry) NewSurface(opts Options) (Surface, error)`
-   `func NewSurfaceByName(name string, width, height int) (Surface, error)`
-   `func (*Registry) (r *Registry) NewSurfaceByName(name string, opts Options) (Surface, error)`
-   `func NewSurfaceByNameWithOptions(name string, opts Options) (Surface, error)`
-   `func NewSurfaceWithOptions(opts Options) (Surface, error)`
-   `func (*Registry) (r *Registry) Register(name string, priority int, factory SurfaceFactory, available func() bool)`
-   `func Register(name string, priority int, factory SurfaceFactory, available func() bool)`
-   `func (*Registry) (r *Registry) Unregister(name string)`
-   `func Unregister(name string)`

## surface/types.go

-   `func (SolidPattern) (p SolidPattern) ColorAt(_, _ float64) color.Color`
-   `func DefaultDrawImageOptions() *DrawImageOptions`
-   `func DefaultFillStyle() FillStyle`
-   `func DefaultOptions(width, height int) Options`
-   `func DefaultStrokeStyle() StrokeStyle`
-   `func (StrokeStyle) (s StrokeStyle) IsDashed() bool`
-   `func Pt(x, y float64) Point`
-   `func (StrokeStyle) (s StrokeStyle) WithCap(lineCap LineCap) StrokeStyle`
-   `func (StrokeStyle) (s StrokeStyle) WithColor(c color.Color) StrokeStyle`
-   `func (FillStyle) (f FillStyle) WithColor(c color.Color) FillStyle`
-   `func (StrokeStyle) (s StrokeStyle) WithDash(pattern []float64, offset float64) StrokeStyle`
-   `func (StrokeStyle) (s StrokeStyle) WithJoin(join LineJoin) StrokeStyle`
-   `func (StrokeStyle) (s StrokeStyle) WithMiterLimit(limit float64) StrokeStyle`
-   `func (StrokeStyle) (s StrokeStyle) WithPattern(p Pattern) StrokeStyle`
-   `func (FillStyle) (f FillStyle) WithPattern(p Pattern) FillStyle`
-   `func (FillStyle) (f FillStyle) WithRule(r FillRule) FillStyle`
-   `func (StrokeStyle) (s StrokeStyle) WithWidth(w float64) StrokeStyle`

## text.go

-   `func (*Context) (c *Context) DrawString(s string, x, y float64)`
-   `func (*Context) (c *Context) DrawStringAnchored(s string, x, y, ax, ay float64)`
-   `func (*Context) (c *Context) DrawStringWrapped(s string, x, y, ax, ay, width, lineSpacing float64, align Align)`
-   `func (*Context) (c *Context) Font() text.Face`
-   `func (*Context) (c *Context) LoadFontFace(path string, points float64) error`
-   `func (*Context) (c *Context) MeasureMultilineString(s string, lineSpacing float64) (width, height float64)`
-   `func (*Context) (c *Context) MeasureString(s string) (w, h float64)`
-   `func (*Context) (c *Context) SetFont(face text.Face)`
-   `func (*Context) (c *Context) WordWrap(s string, w float64) []string`

## text/cache.go

-   `func (*Cache[K, V]) (c *Cache[K, V]) Clear()`
-   `func (*Cache[K, V]) (c *Cache[K, V]) Get(key K) (V, bool)`
-   `func (*Cache[K, V]) (c *Cache[K, V]) GetOrCreate(key K, create func() V) V`
-   `func (*Cache[K, V]) (c *Cache[K, V]) Len() int`
-   `func NewCache[K comparable, V any](softLimit int) *Cache[K, V]`
-   `func (*Cache[K, V]) (c *Cache[K, V]) Set(key K, value V)`

## text/cache/lru.go

-   `func (*lruList[K]) (l *lruList[K]) Clear()`
-   `func (*lruList[K]) (l *lruList[K]) Len() int`
-   `func (*lruList[K]) (l *lruList[K]) MoveToFront(node *lruNode[K])`
-   `func (*lruList[K]) (l *lruList[K]) Oldest() (K, bool)`
-   `func (*lruList[K]) (l *lruList[K]) PushFront(key K) *lruNode[K]`
-   `func (*lruList[K]) (l *lruList[K]) Remove(node *lruNode[K])`
-   `func (*lruList[K]) (l *lruList[K]) RemoveOldest() (K, bool)`

## text/cache/shaping.go

-   `func (*ShapingCache) (c *ShapingCache) Capacity() int`
-   `func (*ShapingCache) (c *ShapingCache) Clear()`
-   `func DefaultShapingCache() *ShapingCache`
-   `func (*ShapingCache) (c *ShapingCache) Delete(key ShapingKey) bool`
-   `func (*ShapingCache) (c *ShapingCache) Get(key ShapingKey) (*text.ShapedRun, bool)`
-   `func (*ShapingCache) (c *ShapingCache) GetOrCreate(key ShapingKey, create func() *text.ShapedRun) *text.ShapedRun`
-   `func HashFeatures(features map[string]int) uint64`
-   `func (*ShapingCache) (c *ShapingCache) Len() int`
-   `func NewShapingCache(capacity int) *ShapingCache`
-   `func NewShapingKey(textStr string, fontID uint64, size float32, direction text.Direction, features uint64) ShapingKey`
-   `func (*ShapingCache) (c *ShapingCache) ResetStats()`
-   `func (*ShapingCache) (c *ShapingCache) Set(key ShapingKey, value *text.ShapedRun)`
-   `func (*ShapingCache) (c *ShapingCache) ShardLen() [DefaultShardCount]int`
-   `func (*ShapingCache) (c *ShapingCache) Stats() CacheStats`
-   `func (*ShapingCache) (c *ShapingCache) TotalCapacity() int`

## text/color_font.go

-   `func DetectGlyphType(font ParsedFont, glyphID uint16) GlyphType`
-   `func GetBitmapGlyph(font ParsedFont, glyphID uint16, ppem uint16) (*emoji.BitmapGlyph, error)`
-   `func GetCOLRGlyph(font ParsedFont, glyphID uint16, paletteIndex int) (*emoji.COLRGlyph, error)`
-   `func (ColorFontInfo) (info ColorFontInfo) HasAnyColorTable() bool`
-   `func (ColorFontInfo) (info ColorFontInfo) PreferredColorFormat() string`

## text/draw.go

-   `func Draw(dst draw.Image, text string, face Face, x, y float64, col color.Color)`
-   `func Measure(text string, face Face) (width, height float64)`

## text/draw_emoji.go

-   `func (*BitmapGlyphCache) (c *BitmapGlyphCache) Clear()`
-   `func DrawWithEmoji(dst draw.Image, text string, face Face, x, y float64, col color.Color)`
-   `func (*BitmapGlyphCache) (c *BitmapGlyphCache) Get(fontID uintptr, glyphID, ppem uint16) *CachedBitmap`
-   `func NewBitmapGlyphCache(maxSize int) *BitmapGlyphCache`
-   `func (*BitmapGlyphCache) (c *BitmapGlyphCache) Put(fontID uintptr, glyphID, ppem uint16, bitmap *emoji.BitmapGlyph, img image.Image)`
-   `func (*BitmapGlyphCache) (c *BitmapGlyphCache) Size() int`

## text/emoji/bitmap.go

-   `func (*SBIXParser) (p *SBIXParser) BestStrikeForPPEM(ppem uint16) int`
-   `func (*BitmapGlyph) (b *BitmapGlyph) Decode() (image.Image, error)`
-   `func (*SBIXParser) (p *SBIXParser) GetGlyph(glyphID, strikeIndex int) (*BitmapGlyph, error)`
-   `func (*SBIXParser) (p *SBIXParser) HasGlyph(glyphID, strikeIndex int) bool`
-   `func (*CBDTParser) (p *CBDTParser) HasTable() bool`
-   `func NewCBDTParser(cbdtData, cblcData []byte) (*CBDTParser, error)`
-   `func NewSBIXParser(data []byte, numGlyphs uint16) (*SBIXParser, error)`
-   `func (*SBIXParser) (p *SBIXParser) NumStrikes() int`
-   `func (*SBIXParser) (p *SBIXParser) StrikePPEM(strikeIndex int) uint16`
-   `func (BitmapFormat) (f BitmapFormat) String() string`

## text/emoji/cbdt_extractor.go

-   `func (*CBDTExtractor) (e *CBDTExtractor) AvailablePPEMs() []uint16`
-   `func (*CBDTExtractor) (e *CBDTExtractor) GetGlyph(glyphID uint16, ppem uint16) (*BitmapGlyph, error)`
-   `func (*CBDTExtractor) (e *CBDTExtractor) GetGlyphAtStrike(glyphID uint16, strikeIndex int) (*BitmapGlyph, error)`
-   `func (*CBDTExtractor) (e *CBDTExtractor) GetGlyphWithStrategy(glyphID uint16, ppem uint16, strategy StrikeStrategy) (*BitmapGlyph, error)`
-   `func (*CBDTExtractor) (e *CBDTExtractor) HasGlyph(glyphID uint16) bool`
-   `func (*CBDTExtractor) (e *CBDTExtractor) HasGlyphInStrike(glyphID uint16, strikeIndex int) bool`
-   `func NewCBDTExtractor(cbdtData, cblcData []byte) (*CBDTExtractor, error)`
-   `func (*CBDTExtractor) (e *CBDTExtractor) NumStrikes() int`
-   `func (*CBDTExtractor) (e *CBDTExtractor) SelectStrike(ppem uint16, strategy StrikeStrategy) int`
-   `func (*CBDTExtractor) (e *CBDTExtractor) StrikeBitDepth(index int) uint8`
-   `func (*CBDTExtractor) (e *CBDTExtractor) StrikePPEM(index int) uint16`
-   `func (StrikeStrategy) (s StrikeStrategy) String() string`

## text/emoji/colr.go

-   `func (Rect) (r Rect) Empty() bool`
-   `func (*COLRParser) (p *COLRParser) GetGlyph(glyphID uint16, paletteIndex int) (*COLRGlyph, error)`
-   `func (*COLRParser) (p *COLRParser) HasGlyph(glyphID uint16) bool`
-   `func (Rect) (r Rect) Height() float64`
-   `func (ColorLayer) (l ColorLayer) IsForeground() bool`
-   `func NewCOLRParser(colrData, cpalData []byte) (*COLRParser, error)`
-   `func (*COLRParser) (p *COLRParser) NumPalettes() int`
-   `func (*COLRParser) (p *COLRParser) PaletteColors(paletteIndex int) []Color`
-   `func (Color) (c Color) RGBA() (r, g, b, a uint32)`
-   `func RenderCOLRToImage(
	glyph *COLRGlyph,
	renderLayer func(glyphID uint16) *image.Alpha,
	width, height int,
	foreground color.RGBA,
) *image.RGBA`
-   `func (Color) (c Color) ToRGBA() color.RGBA`
-   `func (Rect) (r Rect) Width() float64`

## text/emoji/detect.go

-   `func IsBlackFlag(r rune) bool`
-   `func IsCancelTag(r rune) bool`
-   `func IsCombiningEnclosingKeycap(r rune) bool`
-   `func IsEmoji(r rune) bool`
-   `func IsEmojiModifier(r rune) bool`
-   `func IsEmojiModifierBase(r rune) bool`
-   `func IsEmojiPresentation(r rune) bool`
-   `func IsEmojiVariation(r rune) bool`
-   `func IsKeycapBase(r rune) bool`
-   `func IsRegionalIndicator(r rune) bool`
-   `func IsTagCharacter(r rune) bool`
-   `func IsTextPresentation(r rune) bool`
-   `func IsVariationSelector(r rune) bool`
-   `func IsZWJ(r rune) bool`
-   `func Segment(text string) []Run`

## text/emoji/sequence.go

-   `func GetFlagCode(seq Sequence) string`
-   `func GetSkinTone(r rune) SkinTone`
-   `func GetTagSequenceCode(seq Sequence) string`
-   `func (Sequence) (s Sequence) HasModifier() bool`
-   `func IsValidSequence(seq Sequence) bool`
-   `func (Sequence) (s Sequence) Len() int`
-   `func Normalize(seq Sequence) Sequence`
-   `func Parse(runes []rune) []Sequence`
-   `func ParseString(text string) []Sequence`
-   `func SkinToneRune(tone SkinTone) rune`
-   `func (SequenceType) (t SequenceType) String() string`
-   `func (SkinTone) (t SkinTone) String() string`
-   `func (Sequence) (s Sequence) String() string`

## text/errors.go

-   `func (*DirectionMismatchError) (e *DirectionMismatchError) Error() string`

## text/face.go

-   `func (*sourceFace) (f *sourceFace) Advance(text string) float64`
-   `func (*sourceFace) (f *sourceFace) AppendGlyphs(dst []Glyph, text string) []Glyph`
-   `func (*sourceFace) (f *sourceFace) Direction() Direction`
-   `func (*sourceFace) (f *sourceFace) Glyphs(text string) iter.Seq[Glyph]`
-   `func (*sourceFace) (f *sourceFace) HasGlyph(r rune) bool`
-   `func (*sourceFace) (f *sourceFace) Metrics() Metrics`
-   `func (*sourceFace) (f *sourceFace) Size() float64`
-   `func (*sourceFace) (f *sourceFace) Source() *FontSource`

## text/filtered.go

-   `func (*FilteredFace) (f *FilteredFace) Advance(text string) float64`
-   `func (*FilteredFace) (f *FilteredFace) AppendGlyphs(dst []Glyph, text string) []Glyph`
-   `func (UnicodeRange) (ur UnicodeRange) Contains(r rune) bool`
-   `func (*FilteredFace) (f *FilteredFace) Direction() Direction`
-   `func (*FilteredFace) (f *FilteredFace) Glyphs(text string) iter.Seq[Glyph]`
-   `func (*FilteredFace) (f *FilteredFace) HasGlyph(r rune) bool`
-   `func (*FilteredFace) (f *FilteredFace) Metrics() Metrics`
-   `func NewFilteredFace(face Face, ranges ...UnicodeRange) *FilteredFace`
-   `func (*FilteredFace) (f *FilteredFace) Size() float64`
-   `func (*FilteredFace) (f *FilteredFace) Source() *FontSource`

## text/glyph_cache.go

-   `func (*GlyphCache) (c *GlyphCache) Clear()`
-   `func (*GlyphCache) (c *GlyphCache) CurrentFrame() uint64`
-   `func DefaultGlyphCacheConfig() GlyphCacheConfig`
-   `func (*GlyphCache) (c *GlyphCache) Delete(key OutlineCacheKey)`
-   `func (*GlyphCache) (c *GlyphCache) Get(key OutlineCacheKey) *GlyphOutline`
-   `func (*GlyphCachePool) (p *GlyphCachePool) Get() *GlyphCache`
-   `func GetGlobalGlyphCache() *GlyphCache`
-   `func (*GlyphCache) (c *GlyphCache) GetOrCreate(key OutlineCacheKey, create func() *GlyphOutline) *GlyphOutline`
-   `func (*GlyphCache) (c *GlyphCache) HitRate() float64`
-   `func (*GlyphCache) (c *GlyphCache) Len() int`
-   `func (*GlyphCache) (c *GlyphCache) Maintain()`
-   `func NewGlyphCache() *GlyphCache`
-   `func NewGlyphCachePool() *GlyphCachePool`
-   `func NewGlyphCacheWithConfig(config GlyphCacheConfig) *GlyphCache`
-   `func (*GlyphCachePool) (p *GlyphCachePool) Put(c *GlyphCache)`
-   `func (*GlyphCache) (c *GlyphCache) ResetStats()`
-   `func (*GlyphCache) (c *GlyphCache) Set(key OutlineCacheKey, outline *GlyphOutline)`
-   `func SetGlobalGlyphCache(cache *GlyphCache) *GlyphCache`
-   `func (*GlyphCache) (c *GlyphCache) Stats() (hits, misses, evictions, insertions uint64)`

## text/glyph_outline.go

-   `func (*GlyphOutline) (o *GlyphOutline) Clone() *GlyphOutline`
-   `func (*FontError) (e *FontError) Error() string`
-   `func (*OutlineExtractor) (e *OutlineExtractor) ExtractOutline(font ParsedFont, gid GlyphID, size float64) (*GlyphOutline, error)`
-   `func IdentityTransform() *AffineTransform`
-   `func (*GlyphOutline) (o *GlyphOutline) IsEmpty() bool`
-   `func (*AffineTransform) (m *AffineTransform) Multiply(other *AffineTransform) *AffineTransform`
-   `func NewOutlineExtractor() *OutlineExtractor`
-   `func (*GlyphOutline) (o *GlyphOutline) Scale(factor float32) *GlyphOutline`
-   `func ScaleTransform(sx, sy float32) *AffineTransform`
-   `func (*GlyphOutline) (o *GlyphOutline) SegmentCount() int`
-   `func (OutlineOp) (op OutlineOp) String() string`
-   `func (*GlyphOutline) (o *GlyphOutline) Transform(m *AffineTransform) *GlyphOutline`
-   `func (*AffineTransform) (m *AffineTransform) TransformPoint(x, y float32) (float32, float32)`
-   `func (*GlyphOutline) (o *GlyphOutline) Translate(dx, dy float32) *GlyphOutline`
-   `func TranslateTransform(tx, ty float32) *AffineTransform`

## text/glyph_renderer.go

-   `func (*GlyphRenderer) (r *GlyphRenderer) Cache() *GlyphCache`
-   `func DefaultRenderParams() RenderParams`
-   `func GetGlobalTextRenderer() *TextRenderer`
-   `func (*TextRenderer) (tr *TextRenderer) GlyphRenderer() *GlyphRenderer`
-   `func NewGlyphRenderer() *GlyphRenderer`
-   `func NewGlyphRendererWithCache(cache *GlyphCache) *GlyphRenderer`
-   `func NewTextRenderer() *TextRenderer`
-   `func (*GlyphRenderer) (r *GlyphRenderer) RenderGlyph(
	glyph *ShapedGlyph,
	font ParsedFont,
	size float64,
	params RenderParams,
) *GlyphOutline`
-   `func (*GlyphRenderer) (r *GlyphRenderer) RenderGlyphs(
	glyphs []ShapedGlyph,
	font ParsedFont,
	size float64,
	params RenderParams,
) []*GlyphOutline`
-   `func (*GlyphRenderer) (r *GlyphRenderer) RenderLayout(layout *Layout, params RenderParams) [][]*GlyphOutline`
-   `func (*GlyphRenderer) (r *GlyphRenderer) RenderRun(run *ShapedRun, params RenderParams) []*GlyphOutline`
-   `func RotateTransform(angle float32) *AffineTransform`
-   `func ScaleTransformXY(sx, sy float32) *AffineTransform`
-   `func (*GlyphRenderer) (r *GlyphRenderer) SetCache(cache *GlyphCache)`
-   `func (*TextRenderer) (tr *TextRenderer) SetDefaultColor(c color.RGBA)`
-   `func (*TextRenderer) (tr *TextRenderer) SetDefaultFace(face Face)`
-   `func (*TextRenderer) (tr *TextRenderer) SetDefaultSize(size float64)`
-   `func (*TextRenderer) (tr *TextRenderer) ShapeAndRender(text string) ([]*GlyphOutline, error)`
-   `func (*TextRenderer) (tr *TextRenderer) ShapeAndRenderAt(text string, x, y float64) ([]*GlyphOutline, error)`
-   `func (RenderParams) (p RenderParams) WithColor(c color.RGBA) RenderParams`
-   `func (RenderParams) (p RenderParams) WithOpacity(opacity float64) RenderParams`
-   `func (RenderParams) (p RenderParams) WithTransform(t *AffineTransform) RenderParams`

## text/glyph_run.go

-   `func (*GlyphRunBuilder) (b *GlyphRunBuilder) AddGlyph(fontID uint64, glyphID GlyphID, pos Point, size float32)`
-   `func (*GlyphRunBuilder) (b *GlyphRunBuilder) AddShapedGlyph(fontID uint64, glyph *ShapedGlyph, size float32)`
-   `func (*GlyphRunBuilder) (b *GlyphRunBuilder) AddShapedGlyphs(fontID uint64, glyphs []ShapedGlyph, origin Point, size float32)`
-   `func (*GlyphRunBuilder) (b *GlyphRunBuilder) AddShapedRun(run *ShapedRun, origin Point)`
-   `func (*GlyphRunBuilder) (b *GlyphRunBuilder) Build(createGlyph func(fontID uint64, glyphID GlyphID, size float32) *GlyphOutline) []DrawCommand`
-   `func (*GlyphRunBuilder) (b *GlyphRunBuilder) BuildTransformed(
	createGlyph func(fontID uint64, glyphID GlyphID, size float32) *GlyphOutline,
	userTransform *AffineTransform,
) []DrawCommand`
-   `func (*GlyphRunBuilder) (b *GlyphRunBuilder) Cache() *GlyphCache`
-   `func (*GlyphRunBuilder) (b *GlyphRunBuilder) Clear()`
-   `func (*GlyphRunBuilderPool) (p *GlyphRunBuilderPool) Get() *GlyphRunBuilder`
-   `func (*GlyphRunBuilder) (b *GlyphRunBuilder) Instances() []GlyphInstance`
-   `func (*GlyphRunBuilder) (b *GlyphRunBuilder) Len() int`
-   `func NewGlyphRunBuilder(cache *GlyphCache) *GlyphRunBuilder`
-   `func NewGlyphRunBuilderPool(cache *GlyphCache) *GlyphRunBuilderPool`
-   `func (*GlyphRunBuilderPool) (p *GlyphRunBuilderPool) Put(builder *GlyphRunBuilder)`
-   `func (*GlyphRunBuilder) (b *GlyphRunBuilder) SetCache(cache *GlyphCache)`

## text/glyph_type.go

-   `func (GlyphFlags) (f GlyphFlags) Has(flag GlyphFlags) bool`
-   `func (GlyphType) (t GlyphType) String() string`
-   `func (GlyphFlags) (f GlyphFlags) String() string`

## text/layout.go

-   `func DefaultLayoutOptions() LayoutOptions`
-   `func (*Line) (l *Line) Height() float64`
-   `func LayoutText(text string, face Face, opts LayoutOptions) *Layout`
-   `func LayoutTextSimple(text string, face Face) *Layout`
-   `func LayoutTextWithContext(ctx context.Context, text string, face Face, opts LayoutOptions) (*Layout, error)`
-   `func (Alignment) (a Alignment) String() string`

## text/metrics.go

-   `func (Metrics) (m Metrics) LineHeight() float64`

## text/msdf/atlas.go

-   `func (*AtlasManager) (m *AtlasManager) AtlasCount() int`
-   `func (*AtlasManager) (m *AtlasManager) AtlasInfos() []AtlasInfo`
-   `func (*AtlasManager) (m *AtlasManager) Clear()`
-   `func (*ConcurrentAtlasManager) (c *ConcurrentAtlasManager) Clear()`
-   `func (*AtlasManager) (m *AtlasManager) Compact() int`
-   `func (*AtlasManager) (m *AtlasManager) Config() AtlasConfig`
-   `func DefaultAtlasConfig() AtlasConfig`
-   `func (*AtlasManager) (m *AtlasManager) DirtyAtlases() []int`
-   `func (*AtlasConfigError) (e *AtlasConfigError) Error() string`
-   `func (*AtlasFullError) (e *AtlasFullError) Error() string`
-   `func (*AtlasManager) (m *AtlasManager) Generator() *Generator`
-   `func (*ConcurrentAtlasManager) (c *ConcurrentAtlasManager) Get(key GlyphKey, outline *text.GlyphOutline) (Region, error)`
-   `func (*AtlasManager) (m *AtlasManager) Get(key GlyphKey, outline *text.GlyphOutline) (Region, error)`
-   `func (*AtlasManager) (m *AtlasManager) GetAtlas(index int) *Atlas`
-   `func (*AtlasManager) (m *AtlasManager) GetBatch(keys []GlyphKey, outlines []*text.GlyphOutline) ([]Region, error)`
-   `func (*ConcurrentAtlasManager) (c *ConcurrentAtlasManager) GlyphCount() int`
-   `func (*AtlasManager) (m *AtlasManager) GlyphCount() int`
-   `func (*Atlas) (a *Atlas) GlyphCount() int`
-   `func (*AtlasManager) (m *AtlasManager) HasGlyph(key GlyphKey) bool`
-   `func (*ConcurrentAtlasManager) (c *ConcurrentAtlasManager) HasGlyph(key GlyphKey) bool`
-   `func (*Atlas) (a *Atlas) IsDirty() bool`
-   `func (*Atlas) (a *Atlas) IsFull() bool`
-   `func (*AtlasManager) (m *AtlasManager) MarkAllClean()`
-   `func (*AtlasManager) (m *AtlasManager) MarkClean(index int)`
-   `func (*AtlasManager) (m *AtlasManager) MemoryUsage() int64`
-   `func (*ConcurrentAtlasManager) (c *ConcurrentAtlasManager) MemoryUsage() int64`
-   `func NewAtlasManager(config AtlasConfig) (*AtlasManager, error)`
-   `func NewAtlasManagerDefault() *AtlasManager`
-   `func NewConcurrentAtlasManager(config AtlasConfig, numShards int) (*ConcurrentAtlasManager, error)`
-   `func (*AtlasManager) (m *AtlasManager) Remove(key GlyphKey) bool`
-   `func (*AtlasManager) (m *AtlasManager) SetGenerator(g *Generator)`
-   `func (*ConcurrentAtlasManager) (c *ConcurrentAtlasManager) Stats() (hits, misses uint64, atlasCount int)`
-   `func (*AtlasManager) (m *AtlasManager) Stats() (hits, misses uint64, atlasCount int)`
-   `func (*Atlas) (a *Atlas) Utilization() float64`
-   `func (*AtlasConfig) (c *AtlasConfig) Validate() error`

## text/msdf/contour.go

-   `func (*Shape) (s *Shape) AddContour(c *Contour)`
-   `func (*Contour) (c *Contour) AddEdge(e Edge)`
-   `func AssignColors(shape *Shape, angleThreshold float64)`
-   `func (*Contour) (c *Contour) Bounds() Rect`
-   `func (*Shape) (s *Shape) CalculateBounds()`
-   `func (*Contour) (c *Contour) CalculateWinding()`
-   `func (*Contour) (c *Contour) Clone() *Contour`
-   `func (*Shape) (s *Shape) EdgeCount() int`
-   `func FromOutline(outline *text.GlyphOutline) *Shape`
-   `func (*Contour) (c *Contour) IsClockwise() bool`
-   `func NewContour() *Contour`
-   `func NewShape() *Shape`
-   `func SelectBlue(color EdgeColor) bool`
-   `func SelectGreen(color EdgeColor) bool`
-   `func SelectRed(color EdgeColor) bool`
-   `func SwitchColor(current EdgeColor, seed int) EdgeColor`
-   `func (*Shape) (s *Shape) Validate() bool`

## text/msdf/edge.go

-   `func (*Edge) (e *Edge) Bounds() Rect`
-   `func (*Edge) (e *Edge) Clone() Edge`
-   `func (*Edge) (e *Edge) DirectionAt(t float64) Point`
-   `func (*Edge) (e *Edge) EndPoint() Point`
-   `func (EdgeColor) (c EdgeColor) HasBlue() bool`
-   `func (EdgeColor) (c EdgeColor) HasGreen() bool`
-   `func (EdgeColor) (c EdgeColor) HasRed() bool`
-   `func NewCubicEdge(start, control1, control2, end Point) Edge`
-   `func NewLinearEdge(start, end Point) Edge`
-   `func NewQuadraticEdge(start, control, end Point) Edge`
-   `func (*Edge) (e *Edge) PointAt(t float64) Point`
-   `func (*Edge) (e *Edge) SignedDistance(p Point) SignedDistance`
-   `func (*Edge) (e *Edge) StartPoint() Point`
-   `func (EdgeType) (t EdgeType) String() string`
-   `func (EdgeColor) (c EdgeColor) String() string`

## text/msdf/generator.go

-   `func (*Generator) (g *Generator) Config() Config`
-   `func DefaultGenerator() *Generator`
-   `func ErrorCorrection(msdf *MSDF, threshold float64)`
-   `func (*Generator) (g *Generator) Generate(outline *text.GlyphOutline) (*MSDF, error)`
-   `func (*GeneratorPool) (p *GeneratorPool) Generate(outline *text.GlyphOutline) (*MSDF, error)`
-   `func (*Generator) (g *Generator) GenerateBatch(outlines []*text.GlyphOutline) ([]*MSDF, error)`
-   `func (*Generator) (g *Generator) GenerateWithMetrics(outline *text.GlyphOutline) (*MSDF, *Metrics, error)`
-   `func (*GeneratorPool) (p *GeneratorPool) Get() *Generator`
-   `func MedianFilter(msdf *MSDF) *MSDF`
-   `func NewGenerator(config Config) *Generator`
-   `func NewGeneratorPool(config Config) *GeneratorPool`
-   `func (*GeneratorPool) (p *GeneratorPool) Put(g *Generator)`
-   `func (*Generator) (g *Generator) SetConfig(config Config)`

## text/msdf/shelf.go

-   `func (*ShelfAllocator) (a *ShelfAllocator) Allocate(w, h int) (x, y int, ok bool)`
-   `func (*GridAllocator) (g *GridAllocator) Allocate() (x, y int, ok bool)`
-   `func (*ShelfAllocator) (a *ShelfAllocator) AllocateFixed(cellSize int) (x, y int, ok bool)`
-   `func (*GridAllocator) (g *GridAllocator) Allocated() int`
-   `func (*ShelfAllocator) (a *ShelfAllocator) CanFit(w, h int) bool`
-   `func (*GridAllocator) (g *GridAllocator) Capacity() int`
-   `func (*GridAllocator) (g *GridAllocator) CellSize() int`
-   `func (*ShelfAllocator) (a *ShelfAllocator) CurrentShelfRemainingWidth() int`
-   `func (*GridAllocator) (g *GridAllocator) GridDimensions() (cols, rows int)`
-   `func (*GridAllocator) (g *GridAllocator) IsFull() bool`
-   `func NewGridAllocator(width, height, cellSize, padding int) *GridAllocator`
-   `func NewShelfAllocator(width, height, padding int) *ShelfAllocator`
-   `func (*GridAllocator) (g *GridAllocator) Remaining() int`
-   `func (*ShelfAllocator) (a *ShelfAllocator) RemainingHeight() int`
-   `func (*GridAllocator) (g *GridAllocator) Reset()`
-   `func (*ShelfAllocator) (a *ShelfAllocator) Reset()`
-   `func (*ShelfAllocator) (a *ShelfAllocator) ShelfCount() int`
-   `func (*ShelfAllocator) (a *ShelfAllocator) TotalArea() int`
-   `func (*ShelfAllocator) (a *ShelfAllocator) UsedArea() int`
-   `func (*ShelfAllocator) (a *ShelfAllocator) Utilization() float64`
-   `func (*GridAllocator) (g *GridAllocator) Utilization() float64`

## text/msdf/types.go

-   `func (Point) (p Point) Add(q Point) Point`
-   `func (Point) (p Point) Angle() float64`
-   `func AngleBetween(a, b Point) float64`
-   `func (Rect) (r Rect) Center() Point`
-   `func (SignedDistance) (d SignedDistance) Combine(other SignedDistance) SignedDistance`
-   `func (Rect) (r Rect) Contains(p Point) bool`
-   `func (Point) (p Point) Cross(q Point) float64`
-   `func DefaultConfig() Config`
-   `func (Point) (p Point) Dot(q Point) float64`
-   `func (*ConfigError) (e *ConfigError) Error() string`
-   `func (Rect) (r Rect) Expand(margin float64) Rect`
-   `func (*MSDF) (m *MSDF) GetPixel(x, y int) (r, g, b byte)`
-   `func (Rect) (r Rect) Height() float64`
-   `func Infinite() SignedDistance`
-   `func (SignedDistance) (d SignedDistance) IsCloserThan(other SignedDistance) bool`
-   `func (Rect) (r Rect) IsEmpty() bool`
-   `func (Point) (p Point) Length() float64`
-   `func (Point) (p Point) LengthSquared() float64`
-   `func (Point) (p Point) Lerp(q Point, t float64) Point`
-   `func (Point) (p Point) Mul(s float64) Point`
-   `func NewSignedDistance(distance, dot float64) SignedDistance`
-   `func (Point) (p Point) Normalized() Point`
-   `func (*MSDF) (m *MSDF) OutlineToPixel(ox, oy float64) (px, py float64)`
-   `func (Point) (p Point) Perpendicular() Point`
-   `func (*MSDF) (m *MSDF) PixelOffset(x, y int) int`
-   `func (*MSDF) (m *MSDF) PixelToOutline(px, py float64) (ox, oy float64)`
-   `func (*MSDF) (m *MSDF) SetPixel(x, y int, r, g, b byte)`
-   `func (Point) (p Point) Sub(q Point) Point`
-   `func (Rect) (r Rect) Union(s Rect) Rect`
-   `func (*Config) (c *Config) Validate() error`
-   `func (Rect) (r Rect) Width() float64`

## text/multi.go

-   `func (*MultiFace) (m *MultiFace) Advance(text string) float64`
-   `func (*MultiFace) (m *MultiFace) AppendGlyphs(dst []Glyph, text string) []Glyph`
-   `func (*MultiFace) (m *MultiFace) Direction() Direction`
-   `func (*MultiFace) (m *MultiFace) Glyphs(text string) iter.Seq[Glyph]`
-   `func (*MultiFace) (m *MultiFace) HasGlyph(r rune) bool`
-   `func (*MultiFace) (m *MultiFace) Metrics() Metrics`
-   `func NewMultiFace(faces ...Face) (*MultiFace, error)`
-   `func (*MultiFace) (m *MultiFace) Size() float64`
-   `func (*MultiFace) (m *MultiFace) Source() *FontSource`

## text/options.go

-   `func WithCacheLimit(n int) SourceOption`
-   `func WithDirection(d Direction) FaceOption`
-   `func WithHinting(h Hinting) FaceOption`
-   `func WithLanguage(lang string) FaceOption`
-   `func WithParser(name string) SourceOption`

## text/parser.go

-   `func (FontMetrics) (m FontMetrics) Height() float64`
-   `func RegisterParser(name string, parser FontParser)`

## text/parser_ximage.go

-   `func (*ximageParsedFont) (f *ximageParsedFont) FullName() string`
-   `func (*ximageParsedFont) (f *ximageParsedFont) GlyphAdvance(glyphIndex uint16, ppem float64) float64`
-   `func (*ximageParsedFont) (f *ximageParsedFont) GlyphBounds(glyphIndex uint16, ppem float64) Rect`
-   `func (*ximageParsedFont) (f *ximageParsedFont) GlyphIndex(r rune) uint16`
-   `func (*ximageParsedFont) (f *ximageParsedFont) Metrics(ppem float64) FontMetrics`
-   `func (*ximageParsedFont) (f *ximageParsedFont) Name() string`
-   `func (*ximageParsedFont) (f *ximageParsedFont) NumGlyphs() int`
-   `func (*ximageParser) (p *ximageParser) Parse(data []byte) (ParsedFont, error)`
-   `func (*ximageParsedFont) (f *ximageParsedFont) UnitsPerEm() int`

## text/rasterize.go

-   `func RasterizeGlyph(parsed ParsedFont, glyphID GlyphID, ppem float64) *GlyphImage`

## text/runetobool.go

-   `func (*RuneToBoolMap) (m *RuneToBoolMap) Clear()`
-   `func (*RuneToBoolMap) (m *RuneToBoolMap) Get(r rune) (hasGlyph, checked bool)`
-   `func NewRuneToBoolMap() *RuneToBoolMap`
-   `func (*RuneToBoolMap) (m *RuneToBoolMap) Set(r rune, hasGlyph bool)`

## text/script.go

-   `func DetectScript(r rune) Script`
-   `func (Script) (s Script) IsRTL() bool`
-   `func (Script) (s Script) RequiresComplexShaping() bool`
-   `func (Script) (s Script) String() string`

## text/segment.go

-   `func IsPunctuation(r rune) bool`
-   `func IsWhitespace(r rune) bool`
-   `func NewBuiltinSegmenter() *BuiltinSegmenter`
-   `func NewBuiltinSegmenterWithDirection(dir Direction) *BuiltinSegmenter`
-   `func (Segment) (s Segment) RuneCount() int`
-   `func (*BuiltinSegmenter) (s *BuiltinSegmenter) Segment(text string) []Segment`
-   `func SegmentText(text string) []Segment`
-   `func SegmentTextRTL(text string) []Segment`

## text/shaped_run.go

-   `func (*ShapedRun) (r *ShapedRun) Bounds() (x, y, width, height float64)`
-   `func (*ShapedRun) (r *ShapedRun) Height() float64`
-   `func (*ShapedRun) (r *ShapedRun) LineHeight() float64`
-   `func (*ShapedRun) (r *ShapedRun) Width() float64`

## text/shaper.go

-   `func GetShaper() Shaper`
-   `func SetShaper(s Shaper)`
-   `func Shape(text string, face Face) []ShapedGlyph`

## text/shaper_builtin.go

-   `func (*BuiltinShaper) (s *BuiltinShaper) Shape(text string, face Face) []ShapedGlyph`

## text/shaper_gotext.go

-   `func (*GoTextShaper) (s *GoTextShaper) ClearCache()`
-   `func NewGoTextShaper() *GoTextShaper`
-   `func (*GoTextShaper) (s *GoTextShaper) RemoveSource(source *FontSource)`
-   `func (*GoTextShaper) (s *GoTextShaper) Shape(text string, face Face) []ShapedGlyph`

## text/source.go

-   `func (*FontSource) (s *FontSource) Close() error`
-   `func (*FontSource) (s *FontSource) Face(size float64, opts ...FaceOption) Face`
-   `func (*FontSource) (s *FontSource) Name() string`
-   `func NewFontSource(data []byte, opts ...SourceOption) (*FontSource, error)`
-   `func NewFontSourceFromFile(path string, opts ...SourceOption) (*FontSource, error)`
-   `func (*FontSource) (s *FontSource) Parsed() ParsedFont`

## text/subpixel.go

-   `func (SubpixelConfig) (c SubpixelConfig) CacheMultiplier() int`
-   `func (*SubpixelCache) (c *SubpixelCache) Clear()`
-   `func (*SubpixelCache) (c *SubpixelCache) Config() SubpixelConfig`
-   `func DefaultSubpixelConfig() SubpixelConfig`
-   `func (*SubpixelCache) (c *SubpixelCache) Delete(key SubpixelKey)`
-   `func (SubpixelMode) (m SubpixelMode) Divisions() int`
-   `func (*SubpixelCache) (c *SubpixelCache) Get(key SubpixelKey) *GlyphOutline`
-   `func GetGlobalSubpixelCache() *SubpixelCache`
-   `func (*SubpixelCache) (c *SubpixelCache) GetOrCreate(
	key SubpixelKey,
	create func(offsetX, offsetY float64) *GlyphOutline,
) *GlyphOutline`
-   `func HighQualitySubpixelConfig() SubpixelConfig`
-   `func (*SubpixelCache) (c *SubpixelCache) HitRate() float64`
-   `func (SubpixelConfig) (c SubpixelConfig) IsEnabled() bool`
-   `func (SubpixelMode) (m SubpixelMode) IsEnabled() bool`
-   `func (*SubpixelCache) (c *SubpixelCache) Len() int`
-   `func (*SubpixelCache) (c *SubpixelCache) Maintain()`
-   `func MakeSubpixelKey(baseKey OutlineCacheKey, x, y float64, config SubpixelConfig) SubpixelKey`
-   `func NewSubpixelCache(config SubpixelConfig) *SubpixelCache`
-   `func NewSubpixelCacheWithConfig(config SubpixelConfig, glyphConfig GlyphCacheConfig) *SubpixelCache`
-   `func NoSubpixelConfig() SubpixelConfig`
-   `func Quantize(pos float64, mode SubpixelMode) (intPos int, subPos uint8)`
-   `func QuantizePoint(x, y float64, config SubpixelConfig) (intX, intY int, subX, subY uint8)`
-   `func (*SubpixelCache) (c *SubpixelCache) ResetStats()`
-   `func (*SubpixelCache) (c *SubpixelCache) Set(key SubpixelKey, outline *GlyphOutline)`
-   `func (*SubpixelCache) (c *SubpixelCache) SetConfig(config SubpixelConfig)`
-   `func SetGlobalSubpixelCache(cache *SubpixelCache) *SubpixelCache`
-   `func (*SubpixelCache) (c *SubpixelCache) Stats() SubpixelCacheStats`
-   `func (SubpixelMode) (m SubpixelMode) String() string`
-   `func SubpixelOffset(subPos uint8, mode SubpixelMode) float64`
-   `func SubpixelOffsets(subX, subY uint8, config SubpixelConfig) (offsetX, offsetY float64)`

## text/tab.go

-   `func SetTabWidth(n int)`
-   `func TabWidth() int`

## text/types.go

-   `func (Rect) (r Rect) Empty() bool`
-   `func (Rect) (r Rect) Height() float64`
-   `func (Direction) (d Direction) IsHorizontal() bool`
-   `func (Direction) (d Direction) IsVertical() bool`
-   `func (Direction) (d Direction) String() string`
-   `func (Hinting) (h Hinting) String() string`
-   `func (Rect) (r Rect) Width() float64`

## text/wrap.go

-   `func MeasureText(text string, face Face) float64`
-   `func (WrapMode) (m WrapMode) String() string`
-   `func WrapText(text string, face Face, maxWidth float64, mode WrapMode) []WrapResult`

## vec.go

-   `func (Vec2) (v Vec2) Add(w Vec2) Vec2`
-   `func (Vec2) (v Vec2) Angle(w Vec2) float64`
-   `func (Vec2) (v Vec2) Approx(w Vec2, epsilon float64) bool`
-   `func (Vec2) (v Vec2) Atan2() float64`
-   `func (Vec2) (v Vec2) Cross(w Vec2) float64`
-   `func (Vec2) (v Vec2) Div(s float64) Vec2`
-   `func (Vec2) (v Vec2) Dot(w Vec2) float64`
-   `func (Vec2) (v Vec2) IsZero() bool`
-   `func (Vec2) (v Vec2) Length() float64`
-   `func (Vec2) (v Vec2) LengthSq() float64`
-   `func (Vec2) (v Vec2) Lerp(w Vec2, t float64) Vec2`
-   `func (Vec2) (v Vec2) Mul(s float64) Vec2`
-   `func (Vec2) (v Vec2) Neg() Vec2`
-   `func (Vec2) (v Vec2) Normalize() Vec2`
-   `func (Vec2) (v Vec2) Perp() Vec2`
-   `func PointToVec2(p Point) Vec2`
-   `func (Vec2) (v Vec2) Rotate(angle float64) Vec2`
-   `func (Vec2) (v Vec2) Sub(w Vec2) Vec2`
-   `func (Vec2) (v Vec2) ToPoint() Point`
-   `func V2(x, y float64) Vec2`

