package cmd

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"

	"github.com/spf13/cobra"
	"github.com/tma5/otaks/config"
)

var (
	manifestCmd = &cobra.Command{
		Use:     "manifest",
		Short:   "otaks manifest",
		Aliases: []string{"s"},
		Run:     manifest,
	}
)

// ManifestParameter ...
type ManifestParameter struct {
	Text  string `xml:",chardata"`
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

// ManifestConfiguration ...
type ManifestConfiguration struct {
	Text       string              `xml:",chardata"`
	Parameters []ManifestParameter `xml:"Parameter"`
}

// ManifestContent ...
type ManifestContent struct {
	Text     string `xml:",chardata"`
	Ignore   string `xml:"ignore,attr"`
	ZipEntry string `xml:"zipEntry,attr"`
}

// ManifestContents ...
type ManifestContents struct {
	Text    string            `xml:",chardata"`
	Content []ManifestContent `xml:"Content"`
}

// MissionPackageManifest ...
type MissionPackageManifest struct {
	XMLName       xml.Name              `xml:"MissionPackageManifest"`
	Text          string                `xml:",chardata"`
	Version       string                `xml:"version,attr"`
	Configuration ManifestConfiguration `xml:"Configuration"`
	Contents      ManifestContents      `xml:"Contents"`
}

// Entry ...
type Entry struct {
	Text  string `xml:",chardata"`
	Key   string `xml:"key,attr"`
	Class string `xml:"class,attr"`
}

// Preference ...
type Preference struct {
	Text    string  `xml:",chardata"`
	Version string  `xml:"version,attr"`
	Name    string  `xml:"name,attr"`
	Entries []Entry `xml:"entry"`
}

// Preferences ...
type Preferences struct {
	XMLName    xml.Name     `xml:"preferences"`
	Text       string       `xml:",chardata"`
	Preference []Preference `xml:"preference"`
}

func init() {

}

func generateCertificates() (truststore []byte, p12 []byte, err error) {

	return nil, nil, nil
}

func manifest(cmd *cobra.Command, args []string) {
	config, err := config.NewConfig(configLocation)
	if err != nil {
		log.Fatal(err)
	}

	if logLevel != defaultLogLevel {
		config.Server.Logging.Level = logLevel
	}

	connectionString := fmt.Sprintf("%s:%d", config.Server.Domain, config.Server.App.Port)
	muid := uuid.New().String()
	pref := Preferences{
		Preference: []Preference{
			Preference{
				Version: "1",
				Name:    "cot_streams",
				Entries: []Entry{
					Entry{Key: "count", Class: "class java.lang.Integer", Text: "1"},
					Entry{Key: "description0", Class: "class java.lang.String", Text: config.Server.Description},
					Entry{Key: "enabled0", Class: "class java.lang.Boolean", Text: "true"},
					Entry{Key: "connectString0", Class: "class java.lang.String", Text: connectionString},
				},
			},
			Preference{
				Version: "1",
				Name:    "com.atakmap.app_preferences",
				Entries: []Entry{
					Entry{Key: "displayServerConnectionWidget", Class: "class java.lang.Boolean", Text: "true"},
				},
			},
		},
	}

	manifest := MissionPackageManifest{
		Version: "2",
		Configuration: ManifestConfiguration{
			Parameters: []ManifestParameter{
				ManifestParameter{Name: "uid", Value: muid},
				ManifestParameter{Name: "name", Value: config.Server.Name},
				ManifestParameter{Name: "onRecieveDelete", Value: "true"},
			},
		},
		Contents: ManifestContents{
			Content: []ManifestContent{
				ManifestContent{
					Ignore:   "false",
					ZipEntry: fmt.Sprintf("%s/%s.pref", muid, config.Server.Name),
				},
			},
		},
	}
	/*
		/MANIFEST/manifest.xml
		/{UUID}/{NAME}.pref
		/{UUID}/{NAME}.p12
		/{UUID}/truststore.p12
	*/

	if config.Server.TLS.Enabled {
		for _, e := range pref.Preference {
			if e.Name == "com.atakmap.app_preferences" {
				e.Entries = append(e.Entries, Entry{
					Key:   "caLocation",
					Class: "class java.lang.String",
					Text:  "/storage/emulated/0/atak/cert/truststore.p12",
				})
				e.Entries = append(e.Entries, Entry{
					Key:   "certificateLocation",
					Class: "class java.lang.String",
					Text:  fmt.Sprintf("/storage/emulated/0/atak/cert/%s_01.p12", config.Server.Name),
				})
			}
		}

		manifest.Contents.Content = append(manifest.Contents.Content, ManifestContent{
			Ignore: "false", ZipEntry: muid + "/truststore.pk12",
		})
		manifest.Contents.Content = append(manifest.Contents.Content, ManifestContent{
			Ignore: "false", ZipEntry: muid + "/" + config.Server.Name + "_01.pk12",
		})
	}

	out, err := xml.MarshalIndent(manifest, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(out)

	out, err = xml.MarshalIndent(pref, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(out)
}
