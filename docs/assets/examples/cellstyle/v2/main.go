package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/color"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts"
	"github.com/johnfercher/maroto/v2/pkg/grid/row"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/text"
)

func main() {
	cfg := config.NewBuilder().
		WithDebug(false).
		Build()

	colStyle := &props.Style{
		BackgroundColor: &color.Color{80, 80, 80},
		Border:          true,
		BorderColor:     &color.Color{200, 0, 0},
	}

	whiteText := props.Text{
		Color: &color.Color{255, 255, 255},
		Style: consts.Bold,
		Size:  12,
		Align: consts.Center,
		Top:   2,
	}

	rowStyle := &props.Style{
		BackgroundColor: &color.Color{220, 220, 220},
		Border:          true,
		BorderColor:     &color.Color{0, 0, 200},
	}

	blackText := props.Text{
		Style: consts.Bold,
		Size:  12,
		Align: consts.Center,
		Top:   2,
	}

	mrt := pkg.NewMaroto(cfg)
	m := pkg.NewMetricsDecorator(mrt)

	for i := 0; i < 15; i++ {
		m.AddRows(
			row.New(10).Add(
				text.NewCol(4, "string", whiteText).WithStyle(colStyle),
				text.NewCol(4, "string", whiteText).WithStyle(colStyle),
				text.NewCol(4, "string", whiteText).WithStyle(colStyle),
			),
		)

		m.AddRows(
			row.New(10).WithStyle(rowStyle).Add(
				text.NewCol(4, "string", blackText),
				text.NewCol(4, "string", blackText),
				text.NewCol(4, "string", blackText),
			),
		)
	}

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/cellstylev2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}