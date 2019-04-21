package main

import (
	"fmt"
	"log"

	"github.com/coreos/go-semver/semver"
	"github.com/j-and-j-global/wordpress-parser"
)

var (
	provenanceParserMap = PPVs{
		PPV{wordpressParser.Parser{}, "wordpress", semver.Must(semver.NewVersion("1.0.0"))},
	}
)

// PPV represents a Parser Provenance Version mapping
type PPV struct {
	Parser     Parser
	Provenance string
	Version    *semver.Version
}

type PPVs []PPV

func (ppvs PPVs) Find(provenance, version string) (p Parser, err error) {
	v, err := semver.NewVersion(version)
	if err != nil {
		return
	}

	for _, ppv := range ppvs {
		if ppv.Provenance == provenance && ppv.Version.Equal(*v) {
			p = ppv.Parser

			return
		}
	}

	err = fmt.Errorf("Could not find a parser for %q %s", provenance, version)

	return
}

func (ppvs PPVs) TransformerLoop(input, output chan MessageWithEnvelope) (err error) {
	for m := range input {
		parser, err := ppvs.Find(m.Provenance, m.Version)
		if err != nil {
			log.Printf("Transformer Loop: %+v", err)

			continue
		}

		out, err := parser.Parse([]byte(m.Message.Body))
		if err != nil {
			log.Printf("Transformer Loop: %+v", err)

			continue
		}

		m.Message.Body = string(out)
		output <- m
	}

	return fmt.Errorf("Transformer Loop: input channel closed")
}
