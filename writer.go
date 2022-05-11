package writer

import (
	"context"
	"fmt"
	"github.com/paulmach/orb/geojson"
	_ "github.com/sfomuseum/go-sfomuseum-export/v2"
	"github.com/whosonfirst/go-whosonfirst-export/v2"
	wof_writer "github.com/whosonfirst/go-whosonfirst-writer"
	go_writer "github.com/whosonfirst/go-writer"
)

// WriteFeature will serialize and write 'f' using 'wr' using a SFO Museum specific `whosonfirst/go-whosonfirst-export/v2.Exporter` instance.
func WriteFeature(ctx context.Context, wr go_writer.Writer, f *geojson.Feature) error {

	ex, err := export.NewExporter(ctx, "sfomuseum://")

	if err != nil {
		return fmt.Errorf("Failed to create SFO Museum exporter, %w", err)
	}

	return wof_writer.WriteFeatureWithExporter(ctx, wr, ex, f)
}

// WriteBytes will write 'body' using 'wr' using a SFO Musuem specific `whosonfirst/go-whosonfirst-export/v2.Exporter` instance.
func WriteBytes(ctx context.Context, wr go_writer.Writer, body []byte) error {

	ex, err := export.NewExporter(ctx, "sfomuseum://")

	if err != nil {
		return fmt.Errorf("Failed to create SFO Museum exporter, %w", err)
	}

	return wof_writer.WriteBytesWithExporter(ctx, wr, ex, body)
}
