package main

import (
	"io"
	"strings"
	"text/template"
)

// BadgeSvg is the definition for the badge.
type BadgeSvg struct {
	Name            string
	Version         string
	HighlightColor  string
	HeaderTextStart int
	BodyTextStart   int
	HeaderWidth     int
	BodyWidth       int
	TotalWidth      int
}

const (
	characterWidth = 8
	bufferWidth    = 4
)

// NewBadgeSvg will create a new badge.
func NewBadgeSvg(status Status, returnType string) (badge BadgeSvg) {

	badge = BadgeSvg{
		HighlightColor: "#4c1",
		Name:           "Unknown",
		Version:        "---",
	}

	switch strings.ToLower(returnType) {
	case "ert":
		badge.Version = status.ErtVersion
		badge.Name = "ERT"
	case "", "opsman":
		badge.Version = status.OpsManVersion
		badge.Name = "Ops Man"
	default:
		break
	}

	badge.HeaderWidth = (2 * bufferWidth) + (len(badge.Name) * characterWidth)
	badge.BodyWidth = (2 * bufferWidth) + (len(badge.Version) * characterWidth)

	badge.TotalWidth = badge.HeaderWidth + badge.BodyWidth
	badge.HeaderTextStart = badge.HeaderWidth / 2
	badge.BodyTextStart = (badge.BodyWidth / 2) + badge.HeaderWidth

	return
}

// Write writes the badge to a writer.
func (svg BadgeSvg) Write(writer io.Writer) (err error) {
	structure := `<svg xmlns="http://www.w3.org/2000/svg" width="{{.TotalWidth}}" height="20">
    <linearGradient id="b" x2="0" y2="100%">
      <stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
      <stop offset="1" stop-opacity=".1"/>
    </linearGradient>
    <mask id="a">
      <rect width="180" height="20" rx="3" fill="#fff"/>
    </mask>
    <g>
      <path fill="#555" d="M0 0h{{.HeaderWidth}}v20H0z"/>
      <path fill="{{.HighlightColor}}" d="M{{.HeaderWidth}} 0h{{.BodyWidth}}v20H{{.HeaderWidth}}z"/>
    </g>
    <g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11">
      <text x="{{.HeaderTextStart}}" y="15" fill="#010101" fill-opacity=".3">{{.Name}}</text>
      <text x="{{.HeaderTextStart}}" y="14">{{.Name}}</text>
      <text x="{{.BodyTextStart}}" y="15" fill="#010101" fill-opacity=".3">{{.Version}}</text>
      <text x="{{.BodyTextStart}}" y="14">{{.Version}}</text>
    </g>
  </svg>`

	tmpl := template.New("badge")

	if tmpl, err = tmpl.Parse(structure); err == nil {
		err = tmpl.Execute(writer, svg)
	}

	return
}
